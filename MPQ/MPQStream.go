package MPQ

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/JoshVarga/blast"
	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Compression"
)

// Stream represents a stream of data in an MPQ archive
type Stream struct {
	MPQData           MPQ
	BlockTableEntry   BlockTableEntry
	FileName          string
	EncryptionSeed    uint32
	BlockPositions    []uint32
	CurrentPosition   uint32
	CurrentData       []byte
	CurrentBlockIndex uint32
	BlockSize         uint32
}

// CreateStream creates an MPQ stream
func CreateStream(mpq MPQ, blockTableEntry BlockTableEntry, fileName string) *Stream {
	result := &Stream{
		MPQData:           mpq,
		BlockTableEntry:   blockTableEntry,
		CurrentBlockIndex: 0xFFFFFFFF,
	}
	fileSegs := strings.Split(fileName, `\`)
	result.EncryptionSeed = hashString(fileSegs[len(fileSegs)-1], 3)
	if result.BlockTableEntry.HasFlag(MpqFileFixKey) {
		result.EncryptionSeed = (result.EncryptionSeed + result.BlockTableEntry.FilePosition) ^ result.BlockTableEntry.UncompressedFileSize
	}
	result.BlockSize = 0x200 << result.MPQData.Data.BlockSize

	if (result.BlockTableEntry.HasFlag(MpqFileCompress) || result.BlockTableEntry.HasFlag(MpqFileImplode)) && !result.BlockTableEntry.HasFlag(MpqFileSingleUnit) {
		result.loadBlockOffsets()
	}
	return result
}

func (v *Stream) loadBlockOffsets() {
	blockPositionCount := ((v.BlockTableEntry.UncompressedFileSize + v.BlockSize - 1) / v.BlockSize) + 1
	v.BlockPositions = make([]uint32, blockPositionCount)
	v.MPQData.File.Seek(int64(v.BlockTableEntry.FilePosition), 0)
	binary.Read(v.MPQData.File, binary.LittleEndian, &v.BlockPositions)
	blockPosSize := blockPositionCount << 2
	if v.BlockTableEntry.HasFlag(MpqFileEncrypted) {
		decrypt(v.BlockPositions, v.EncryptionSeed-1)
		if v.BlockPositions[0] != blockPosSize {
			panic("Decryption of MPQ failed!")
		}
		if v.BlockPositions[1] > v.BlockSize+blockPosSize {
			panic("Decryption of MPQ failed!")
		}
	}
}

func (v *Stream) Read(buffer []byte, offset, count uint32) uint32 {
	if v.BlockTableEntry.HasFlag(MpqFileSingleUnit) {
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

	bytesToCopy := Common.Min(uint32(len(v.CurrentData))-v.CurrentPosition, count)
	copy(buffer[offset:offset+bytesToCopy], v.CurrentData[v.CurrentPosition:v.CurrentPosition+bytesToCopy])
	v.CurrentPosition += bytesToCopy
	return bytesToCopy
}

func (v *Stream) readInternal(buffer []byte, offset, count uint32) uint32 {
	v.bufferData()
	localPosition := v.CurrentPosition % v.BlockSize
	bytesToCopy := Common.Min(uint32(len(v.CurrentData))-localPosition, count)
	if bytesToCopy <= 0 {
		return 0
	}
	copy(buffer[offset:offset+bytesToCopy], v.CurrentData[localPosition:localPosition+bytesToCopy])
	v.CurrentPosition += bytesToCopy
	return bytesToCopy
}

func (v *Stream) bufferData() {
	requiredBlock := uint32(v.CurrentPosition / v.BlockSize)
	if requiredBlock == v.CurrentBlockIndex {
		return
	}
	expectedLength := Common.Min(v.BlockTableEntry.UncompressedFileSize-(requiredBlock*v.BlockSize), v.BlockSize)
	v.CurrentData = v.loadBlock(requiredBlock, expectedLength)
	v.CurrentBlockIndex = requiredBlock
}

func (v *Stream) loadSingleUnit() {
	fileData := make([]byte, v.BlockSize)
	v.MPQData.File.Seek(int64(v.MPQData.Data.HeaderSize), 0)
	binary.Read(v.MPQData.File, binary.LittleEndian, &fileData)
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
	if v.BlockTableEntry.HasFlag(MpqFileCompress) || v.BlockTableEntry.HasFlag(MpqFileImplode) {
		offset = v.BlockPositions[blockIndex]
		toRead = v.BlockPositions[blockIndex+1] - offset
	} else {
		offset = blockIndex * v.BlockSize
		toRead = expectedLength
	}
	offset += v.BlockTableEntry.FilePosition
	data := make([]byte, toRead)
	v.MPQData.File.Seek(int64(offset), 0)
	binary.Read(v.MPQData.File, binary.LittleEndian, &data)
	if v.BlockTableEntry.HasFlag(MpqFileEncrypted) && v.BlockTableEntry.UncompressedFileSize > 3 {
		if v.EncryptionSeed == 0 {
			panic("Unable to determine encryption key")
		}

		decryptBytes(data, blockIndex+v.EncryptionSeed)
	}
	if v.BlockTableEntry.HasFlag(MpqFileCompress) && (toRead != expectedLength) {
		if !v.BlockTableEntry.HasFlag(MpqFileSingleUnit) {
			data = decompressMulti(data, expectedLength)
		} else {
			data = pkDecompress(data)
		}
	}
	if v.BlockTableEntry.HasFlag(MpqFileImplode) && (toRead != expectedLength) {
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
		//return MpqWavCompression.Decompress(sinput, 2);
		panic("ima adpcm sterio decompression not supported")
	case 0x40: // IMA ADPCM Mono
		//return MpqWavCompression.Decompress(sinput, 1)
		panic("mpq wav decompression not supported")
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
		sinput := Compression.HuffmanDecompress(data[1:])
		sinput = Compression.WavDecompress(sinput, 1)
		tmp := make([]byte, len(sinput))
		copy(tmp, sinput)
		return tmp
	case 0x48:
		//byte[] result = PKDecompress(sinput, outputLength);
		//return MpqWavCompression.Decompress(new MemoryStream(result), 1);
		panic("pk + mpqwav decompression not supported")
	case 0x81:
		sinput := Compression.HuffmanDecompress(data[1:])
		sinput = Compression.WavDecompress(sinput, 2)
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
	defer r.Close()
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r)
	return buffer.Bytes()
}

func pkDecompress(data []byte) []byte {
	b := bytes.NewReader(data)
	r, err := blast.NewReader(b)
	defer r.Close()
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r)
	return buffer.Bytes()
}
