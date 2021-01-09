package d2mpq

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/JoshVarga/blast"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2compression"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// Stream represents a stream of data in an MPQ archive
type Stream struct {
	BlockTableEntry   BlockTableEntry
	BlockPositions    []uint32
	CurrentData       []byte
	FileName          string
	MPQData           *MPQ
	EncryptionSeed    uint32
	CurrentPosition   uint32
	CurrentBlockIndex uint32
	BlockSize         uint32
}

// CreateStream creates an MPQ stream
func CreateStream(mpq *MPQ, blockTableEntry BlockTableEntry, fileName string) (*Stream, error) {
	result := &Stream{
		MPQData:           mpq,
		BlockTableEntry:   blockTableEntry,
		CurrentBlockIndex: 0xFFFFFFFF, //nolint:gomnd // MPQ magic
	}
	fileSegs := strings.Split(fileName, `\`)
	result.EncryptionSeed = hashString(fileSegs[len(fileSegs)-1], 3)

	if result.BlockTableEntry.HasFlag(FileFixKey) {
		result.EncryptionSeed = (result.EncryptionSeed + result.BlockTableEntry.FilePosition) ^ result.BlockTableEntry.UncompressedFileSize
	}

	result.BlockSize = 0x200 << result.MPQData.data.BlockSize //nolint:gomnd // MPQ magic

	if result.BlockTableEntry.HasFlag(FilePatchFile) {
		log.Fatal("Patching is not supported")
	}

	var err error

	if (result.BlockTableEntry.HasFlag(FileCompress) || result.BlockTableEntry.HasFlag(FileImplode)) &&
		!result.BlockTableEntry.HasFlag(FileSingleUnit) {
		err = result.loadBlockOffsets()
	}

	return result, err
}

func (v *Stream) loadBlockOffsets() error {
	blockPositionCount := ((v.BlockTableEntry.UncompressedFileSize + v.BlockSize - 1) / v.BlockSize) + 1
	v.BlockPositions = make([]uint32, blockPositionCount)

	_, err := v.MPQData.file.Seek(int64(v.BlockTableEntry.FilePosition), 0)
	if err != nil {
		return err
	}

	mpqBytes := make([]byte, blockPositionCount*4) //nolint:gomnd // MPQ magic

	_, err = v.MPQData.file.Read(mpqBytes)
	if err != nil {
		return err
	}

	for i := range v.BlockPositions {
		idx := i * 4 //nolint:gomnd // MPQ magic
		v.BlockPositions[i] = binary.LittleEndian.Uint32(mpqBytes[idx : idx+4])
	}

	blockPosSize := blockPositionCount << 2 //nolint:gomnd // MPQ magic

	if v.BlockTableEntry.HasFlag(FileEncrypted) {
		decrypt(v.BlockPositions, v.EncryptionSeed-1)

		if v.BlockPositions[0] != blockPosSize {
			log.Println("Decryption of MPQ failed!")
			return errors.New("decryption of MPQ failed")
		}

		if v.BlockPositions[1] > v.BlockSize+blockPosSize {
			log.Println("Decryption of MPQ failed!")
			return errors.New("decryption of MPQ failed")
		}
	}

	return nil
}

func (v *Stream) Read(buffer []byte, offset, count uint32) uint32 {
	if v.BlockTableEntry.HasFlag(FileSingleUnit) {
		return v.readInternalSingleUnit(buffer, offset, count)
	}

	toRead := count
	readTotal := uint32(0)

	for toRead > 0 {
		read := v.readInternal(buffer, offset, toRead)

		if read == 0 {
			break
		}

		readTotal += read
		offset += read
		toRead -= read
	}

	return readTotal
}

func (v *Stream) readInternalSingleUnit(buffer []byte, offset, count uint32) uint32 {
	if len(v.CurrentData) == 0 {
		v.loadSingleUnit()
	}

	bytesToCopy := d2math.Min(uint32(len(v.CurrentData))-v.CurrentPosition, count)

	copy(buffer[offset:offset+bytesToCopy], v.CurrentData[v.CurrentPosition:v.CurrentPosition+bytesToCopy])

	v.CurrentPosition += bytesToCopy

	return bytesToCopy
}

func (v *Stream) readInternal(buffer []byte, offset, count uint32) uint32 {
	v.bufferData()

	localPosition := v.CurrentPosition % v.BlockSize
	bytesToCopy := d2math.MinInt32(int32(len(v.CurrentData))-int32(localPosition), int32(count))

	if bytesToCopy <= 0 {
		return 0
	}

	copy(buffer[offset:offset+uint32(bytesToCopy)], v.CurrentData[localPosition:localPosition+uint32(bytesToCopy)])

	v.CurrentPosition += uint32(bytesToCopy)

	return uint32(bytesToCopy)
}

func (v *Stream) bufferData() {
	requiredBlock := v.CurrentPosition / v.BlockSize

	if requiredBlock == v.CurrentBlockIndex {
		return
	}

	expectedLength := d2math.Min(v.BlockTableEntry.UncompressedFileSize-(requiredBlock*v.BlockSize), v.BlockSize)
	v.CurrentData = v.loadBlock(requiredBlock, expectedLength)
	v.CurrentBlockIndex = requiredBlock
}

func (v *Stream) loadSingleUnit() {
	fileData := make([]byte, v.BlockSize)

	_, err := v.MPQData.file.Seek(int64(v.MPQData.data.HeaderSize), 0)
	if err != nil {
		log.Print(err)
	}

	_, err = v.MPQData.file.Read(fileData)
	if err != nil {
		log.Print(err)
	}

	if v.BlockSize == v.BlockTableEntry.UncompressedFileSize {
		v.CurrentData = fileData
		return
	}

	v.CurrentData = decompressMulti(fileData, v.BlockTableEntry.UncompressedFileSize)
}

func (v *Stream) loadBlock(blockIndex, expectedLength uint32) []byte {
	var (
		offset uint32
		toRead uint32
	)

	if v.BlockTableEntry.HasFlag(FileCompress) || v.BlockTableEntry.HasFlag(FileImplode) {
		offset = v.BlockPositions[blockIndex]
		toRead = v.BlockPositions[blockIndex+1] - offset
	} else {
		offset = blockIndex * v.BlockSize
		toRead = expectedLength
	}

	offset += v.BlockTableEntry.FilePosition
	data := make([]byte, toRead)

	_, err := v.MPQData.file.Seek(int64(offset), 0)
	if err != nil {
		log.Print(err)
	}

	_, err = v.MPQData.file.Read(data)
	if err != nil {
		log.Print(err)
	}

	if v.BlockTableEntry.HasFlag(FileEncrypted) && v.BlockTableEntry.UncompressedFileSize > 3 {
		if v.EncryptionSeed == 0 {
			panic("Unable to determine encryption key")
		}

		decryptBytes(data, blockIndex+v.EncryptionSeed)
	}

	if v.BlockTableEntry.HasFlag(FileCompress) && (toRead != expectedLength) {
		if !v.BlockTableEntry.HasFlag(FileSingleUnit) {
			data = decompressMulti(data, expectedLength)
		} else {
			data = pkDecompress(data)
		}
	}

	if v.BlockTableEntry.HasFlag(FileImplode) && (toRead != expectedLength) {
		data = pkDecompress(data)
	}

	return data
}

//nolint:gomnd // Will fix enum values later
func decompressMulti(data []byte /*expectedLength*/, _ uint32) []byte {
	compressionType := data[0]

	switch compressionType {
	case 1: // Huffman
		panic("huffman decompression not supported")
	case 2: // ZLib/Deflate
		return deflate(data[1:])
	case 8: // PKLib/Impode
		return pkDecompress(data[1:])
	case 0x10: // BZip2
		panic("bzip2 decompression not supported")
	case 0x80: // IMA ADPCM Stereo
		return d2compression.WavDecompress(data[1:], 2)
	case 0x40: // IMA ADPCM Mono
		return d2compression.WavDecompress(data[1:], 1)
	case 0x12:
		panic("lzma decompression not supported")
	// Combos
	case 0x22:
		// sparse then zlib
		panic("sparse decompression + deflate decompression not supported")
	case 0x30:
		// sparse then bzip2
		panic("sparse decompression + bzip2 decompression not supported")
	case 0x41:
		sinput := d2compression.HuffmanDecompress(data[1:])
		sinput = d2compression.WavDecompress(sinput, 1)
		tmp := make([]byte, len(sinput))

		copy(tmp, sinput)

		return tmp
	case 0x48:
		// byte[] result = PKDecompress(sinput, outputLength);
		// return MpqWavCompression.Decompress(new MemoryStream(result), 1);
		panic("pk + mpqwav decompression not supported")
	case 0x81:
		sinput := d2compression.HuffmanDecompress(data[1:])
		sinput = d2compression.WavDecompress(sinput, 2)
		tmp := make([]byte, len(sinput))
		copy(tmp, sinput)

		return tmp
	case 0x88:
		// byte[] result = PKDecompress(sinput, outputLength);
		// return MpqWavCompression.Decompress(new MemoryStream(result), 2);
		panic("pk + wav decompression not supported")
	default:
		panic(fmt.Sprintf("decompression not supported for unknown compression type %X", compressionType))
	}
}

func deflate(data []byte) []byte {
	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)

	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)

	_, err = buffer.ReadFrom(r)
	if err != nil {
		log.Panic(err)
	}

	err = r.Close()
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}

func pkDecompress(data []byte) []byte {
	b := bytes.NewReader(data)
	r, err := blast.NewReader(b)

	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)

	_, err = buffer.ReadFrom(r)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}
