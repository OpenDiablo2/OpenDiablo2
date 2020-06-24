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

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2compression"
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
		CurrentBlockIndex: 0xFFFFFFFF,
	}
	fileSegs := strings.Split(fileName, `\`)
	result.EncryptionSeed = hashString(fileSegs[len(fileSegs)-1], 3)
	if result.BlockTableEntry.HasFlag(FileFixKey) {
		result.EncryptionSeed = (result.EncryptionSeed + result.BlockTableEntry.FilePosition) ^ result.BlockTableEntry.UncompressedFileSize
	}
	result.BlockSize = 0x200 << result.MPQData.Data.BlockSize

	if result.BlockTableEntry.HasFlag(FilePatchFile) {
		log.Fatal("Patching is not supported")
	}

	var err error
	if (result.BlockTableEntry.HasFlag(FileCompress) || result.BlockTableEntry.HasFlag(FileImplode)) && !result.BlockTableEntry.HasFlag(FileSingleUnit) {
		err = result.loadBlockOffsets()
	}
	return result, err
}

func (v *Stream) loadBlockOffsets() error {
	blockPositionCount := ((v.BlockTableEntry.UncompressedFileSize + v.BlockSize - 1) / v.BlockSize) + 1
	v.BlockPositions = make([]uint32, blockPositionCount)
	v.MPQData.File.Seek(int64(v.BlockTableEntry.FilePosition), 0)
	mpqBytes := make([]byte, blockPositionCount*4)
	v.MPQData.File.Read(mpqBytes)
	for i := range v.BlockPositions {
		idx := i * 4
		v.BlockPositions[i] = binary.LittleEndian.Uint32(mpqBytes[idx : idx+4])
	}
	//binary.Read(v.MPQData.File, binary.LittleEndian, &v.BlockPositions)
	blockPosSize := blockPositionCount << 2
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
		read := v.readInternal(buffer, offset, count)
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

	bytesToCopy := d2common.Min(uint32(len(v.CurrentData))-v.CurrentPosition, count)
	copy(buffer[offset:offset+bytesToCopy], v.CurrentData[v.CurrentPosition:v.CurrentPosition+bytesToCopy])
	v.CurrentPosition += bytesToCopy
	return bytesToCopy
}

func (v *Stream) readInternal(buffer []byte, offset, count uint32) uint32 {
	v.bufferData()
	localPosition := v.CurrentPosition % v.BlockSize
	bytesToCopy := d2common.MinInt32(int32(len(v.CurrentData))-int32(localPosition), int32(count))
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
	expectedLength := d2common.Min(v.BlockTableEntry.UncompressedFileSize-(requiredBlock*v.BlockSize), v.BlockSize)
	v.CurrentData = v.loadBlock(requiredBlock, expectedLength)
	v.CurrentBlockIndex = requiredBlock
}

func (v *Stream) loadSingleUnit() {
	fileData := make([]byte, v.BlockSize)
	v.MPQData.File.Seek(int64(v.MPQData.Data.HeaderSize), 0)
	v.MPQData.File.Read(fileData)
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
	v.MPQData.File.Seek(int64(offset), 0)
	v.MPQData.File.Read(data)
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

func decompressMulti(data []byte, expectedLength uint32) []byte {
	copmressionType := data[0]
	switch copmressionType {
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
		// TODO: sparse then zlib
		panic("sparse decompression + deflate decompression not supported")
	case 0x30:
		// TODO: sparse then bzip2
		panic("sparse decompression + bzip2 decompression not supported")
	case 0x41:
		sinput := d2compression.HuffmanDecompress(data[1:])
		sinput = d2compression.WavDecompress(sinput, 1)
		tmp := make([]byte, len(sinput))
		copy(tmp, sinput)
		return tmp
	case 0x48:
		//byte[] result = PKDecompress(sinput, outputLength);
		//return MpqWavCompression.Decompress(new MemoryStream(result), 1);
		panic("pk + mpqwav decompression not supported")
	case 0x81:
		sinput := d2compression.HuffmanDecompress(data[1:])
		sinput = d2compression.WavDecompress(sinput, 2)
		tmp := make([]byte, len(sinput))
		copy(tmp, sinput)
		return tmp
	case 0x88:
		//byte[] result = PKDecompress(sinput, outputLength);
		//return MpqWavCompression.Decompress(new MemoryStream(result), 2);
		panic("pk + wav decompression not supported")
	default:
		panic(fmt.Sprintf("decompression not supported for unknown compression type %X", copmressionType))
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
