package d2mpq

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/JoshVarga/blast"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2compression"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// Stream represents a stream of data in an MPQ archive
type Stream struct {
	Data      []byte
	Positions []uint32
	MPQ       *MPQ
	Block     *Block
	Index     uint32
	Size      uint32
	Position  uint32
}

// CreateStream creates an MPQ stream
func CreateStream(mpq *MPQ, block *Block, fileName string) (*Stream, error) {
	s := &Stream{
		MPQ:   mpq,
		Block: block,
		Index: 0xFFFFFFFF, //nolint:gomnd // MPQ magic
	}

	if s.Block.HasFlag(FileFixKey) {
		s.Block.calculateEncryptionSeed(fileName)
	}

	s.Size = 0x200 << s.MPQ.header.BlockSize //nolint:gomnd // MPQ magic

	if s.Block.HasFlag(FilePatchFile) {
		return nil, errors.New("patching is not supported")
	}

	if (s.Block.HasFlag(FileCompress) || s.Block.HasFlag(FileImplode)) && !s.Block.HasFlag(FileSingleUnit) {
		if err := s.loadBlockOffsets(); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (v *Stream) loadBlockOffsets() error {
	if _, err := v.MPQ.file.Seek(int64(v.Block.FilePosition), io.SeekStart); err != nil {
		return err
	}

	blockPositionCount := ((v.Block.UncompressedFileSize + v.Size - 1) / v.Size) + 1
	v.Positions = make([]uint32, blockPositionCount)

	if err := binary.Read(v.MPQ.file, binary.LittleEndian, &v.Positions); err != nil {
		return err
	}

	if v.Block.HasFlag(FileEncrypted) {
		decrypt(v.Positions, v.Block.EncryptionSeed-1)

		blockPosSize := blockPositionCount << 2 //nolint:gomnd // MPQ magic
		if v.Positions[0] != blockPosSize {
			return errors.New("decryption of MPQ failed")
		}

		if v.Positions[1] > v.Size+blockPosSize {
			return errors.New("decryption of MPQ failed")
		}
	}

	return nil
}

func (v *Stream) Read(buffer []byte, offset, count uint32) (readTotal uint32, err error) {
	if v.Block.HasFlag(FileSingleUnit) {
		return v.readInternalSingleUnit(buffer, offset, count)
	}

	var read uint32

	toRead := count
	for toRead > 0 {
		if read, err = v.readInternal(buffer, offset, toRead); err != nil {
			return readTotal, err
		}

		if read == 0 {
			break
		}

		readTotal += read
		offset += read
		toRead -= read
	}

	return readTotal, nil
}

func (v *Stream) readInternalSingleUnit(buffer []byte, offset, count uint32) (uint32, error) {
	if len(v.Data) == 0 {
		if err := v.loadSingleUnit(); err != nil {
			return 0, err
		}
	}

	return v.copy(buffer, offset, v.Position, count)
}

func (v *Stream) readInternal(buffer []byte, offset, count uint32) (uint32, error) {
	if err := v.bufferData(); err != nil {
		return 0, err
	}

	localPosition := v.Position % v.Size

	return v.copy(buffer, offset, localPosition, count)
}

func (v *Stream) copy(buffer []byte, offset, pos, count uint32) (uint32, error) {
	bytesToCopy := d2math.Min(uint32(len(v.Data))-pos, count)
	if bytesToCopy <= 0 {
		return 0, io.EOF
	}

	copy(buffer[offset:offset+bytesToCopy], v.Data[pos:pos+bytesToCopy])
	v.Position += bytesToCopy

	return bytesToCopy, nil
}

func (v *Stream) bufferData() (err error) {
	blockIndex := v.Position / v.Size

	if blockIndex == v.Index {
		return nil
	}

	expectedLength := d2math.Min(v.Block.UncompressedFileSize-(blockIndex*v.Size), v.Size)
	if v.Data, err = v.loadBlock(blockIndex, expectedLength); err != nil {
		return err
	}

	v.Index = blockIndex

	return nil
}

func (v *Stream) loadSingleUnit() (err error) {
	if _, err = v.MPQ.file.Seek(int64(v.MPQ.header.HeaderSize), io.SeekStart); err != nil {
		return err
	}

	fileData := make([]byte, v.Size)

	if _, err = v.MPQ.file.Read(fileData); err != nil {
		return err
	}

	if v.Size == v.Block.UncompressedFileSize {
		v.Data = fileData
		return nil
	}

	v.Data, err = decompressMulti(fileData, v.Block.UncompressedFileSize)

	return err
}

func (v *Stream) loadBlock(blockIndex, expectedLength uint32) ([]byte, error) {
	var (
		offset uint32
		toRead uint32
	)

	if v.Block.HasFlag(FileCompress) || v.Block.HasFlag(FileImplode) {
		offset = v.Positions[blockIndex]
		toRead = v.Positions[blockIndex+1] - offset
	} else {
		offset = blockIndex * v.Size
		toRead = expectedLength
	}

	offset += v.Block.FilePosition
	data := make([]byte, toRead)

	if _, err := v.MPQ.file.Seek(int64(offset), io.SeekStart); err != nil {
		return []byte{}, err
	}

	if _, err := v.MPQ.file.Read(data); err != nil {
		return []byte{}, err
	}

	if v.Block.HasFlag(FileEncrypted) && v.Block.UncompressedFileSize > 3 {
		if v.Block.EncryptionSeed == 0 {
			return []byte{}, errors.New("unable to determine encryption key")
		}

		decryptBytes(data, blockIndex+v.Block.EncryptionSeed)
	}

	if v.Block.HasFlag(FileCompress) && (toRead != expectedLength) {
		if !v.Block.HasFlag(FileSingleUnit) {
			return decompressMulti(data, expectedLength)
		}

		return pkDecompress(data)
	}

	if v.Block.HasFlag(FileImplode) && (toRead != expectedLength) {
		return pkDecompress(data)
	}

	return data, nil
}

//nolint:gomnd,funlen,gocyclo // Will fix enum values later, can't help function length
func decompressMulti(data []byte /*expectedLength*/, _ uint32) ([]byte, error) {
	compressionType := data[0]

	switch compressionType {
	case 1: // Huffman
		return []byte{}, errors.New("huffman decompression not supported")
	case 2: // ZLib/Deflate
		return deflate(data[1:])
	case 8: // PKLib/Impode
		return pkDecompress(data[1:])
	case 0x10: // BZip2
		return []byte{}, errors.New("bzip2 decompression not supported")
	case 0x80: // IMA ADPCM Stereo
		return d2compression.WavDecompress(data[1:], 2)
	case 0x40: // IMA ADPCM Mono
		return d2compression.WavDecompress(data[1:], 1)
	case 0x12:
		return []byte{}, errors.New("lzma decompression not supported")
	// Combos
	case 0x22:
		// sparse then zlib
		return []byte{}, errors.New("sparse decompression + deflate decompression not supported")
	case 0x30:
		// sparse then bzip2
		return []byte{}, errors.New("sparse decompression + bzip2 decompression not supported")
	case 0x41:
		sinput, err := d2compression.WavDecompress(d2compression.HuffmanDecompress(data[1:]), 1)
		if err != nil {
			return nil, err
		}

		tmp := make([]byte, len(sinput))

		copy(tmp, sinput)

		return tmp, nil
	case 0x48:
		// byte[] result = PKDecompress(sinput, outputLength);
		// return MpqWavCompression.Decompress(new MemoryStream(result), 1);
		return []byte{}, errors.New("pk + mpqwav decompression not supported")
	case 0x81:
		sinput, err := d2compression.WavDecompress(d2compression.HuffmanDecompress(data[1:]), 2)
		if err != nil {
			return nil, err
		}

		tmp := make([]byte, len(sinput))
		copy(tmp, sinput)

		return tmp, nil
	case 0x88:
		// byte[] result = PKDecompress(sinput, outputLength);
		// return MpqWavCompression.Decompress(new MemoryStream(result), 2);
		return []byte{}, errors.New("pk + wav decompression not supported")
	}

	return []byte{}, fmt.Errorf("decompression not supported for unknown compression type %X", compressionType)
}

func deflate(data []byte) ([]byte, error) {
	b := bytes.NewReader(data)

	r, err := zlib.NewReader(b)
	if err != nil {
		return []byte{}, err
	}

	buffer := new(bytes.Buffer)

	_, err = buffer.ReadFrom(r)
	if err != nil {
		return []byte{}, err
	}

	err = r.Close()
	if err != nil {
		return []byte{}, err
	}

	return buffer.Bytes(), nil
}

func pkDecompress(data []byte) ([]byte, error) {
	b := bytes.NewReader(data)

	r, err := blast.NewReader(b)
	if err != nil {
		return []byte{}, err
	}

	buffer := new(bytes.Buffer)

	if _, err = buffer.ReadFrom(r); err != nil {
		return []byte{}, err
	}

	err = r.Close()
	if err != nil {
		return []byte{}, err
	}

	return buffer.Bytes(), nil
}
