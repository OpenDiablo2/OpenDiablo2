package Common

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

// StreamReader allows you to read data from a byte array in various formats
type StreamReader struct {
	data     []byte
	position uint64
}

// CreateStreamReader creates an instance of the stream reader
func CreateStreamReader(source []byte) *StreamReader {
	result := &StreamReader{
		data:     source,
		position: 0,
	}
	return result
}

// GetPosition returns the current stream position
func (v *StreamReader) GetPosition() uint64 {
	return v.position
}

// GetSize returns the total size of the stream in bytes
func (v *StreamReader) GetSize() uint64 {
	return uint64(len(v.data))
}

// GetByte returns a byte from the stream
func (v *StreamReader) GetByte() byte {
	result := v.data[v.position]
	v.position++
	return result
}

// GetUInt16 returns a uint16 word from the stream
func (v *StreamReader) GetUInt16() uint16 {
	result := uint16(v.data[v.position])
	result += uint16(v.data[v.position+1]) << 8
	v.position += 2
	return result
}

// GetInt16 returns a int16 word from the stream
func (v *StreamReader) GetInt16() int16 {
	var result int16
	err := binary.Read(bytes.NewReader([]byte{v.data[v.position], v.data[v.position+1]}), binary.LittleEndian, &result)
	if err != nil {
		log.Panic(err)
	}
	v.position += 2
	return result
}

func (v *StreamReader) SetPosition(newPosition uint64) {
	v.position = newPosition
}

// GetUInt32 returns a uint32 word from the stream
func (v *StreamReader) GetUInt32() uint32 {
	var result uint32
	err := binary.Read(bytes.NewReader(
		[]byte{
			v.data[v.position],
			v.data[v.position+1],
			v.data[v.position+2],
			v.data[v.position+3],
		}), binary.LittleEndian, &result)
	if err != nil {
		log.Panic(err)
	}
	v.position += 4
	return result
}

// GetInt32 returns an int32 word from the stream
func (v *StreamReader) GetInt32() int32 {
	var result int32
	err := binary.Read(bytes.NewReader(
		[]byte{
			v.data[v.position],
			v.data[v.position+1],
			v.data[v.position+2],
			v.data[v.position+3],
		}), binary.LittleEndian, &result)
	if err != nil {
		log.Panic(err)
	}
	v.position += 4
	return result
}

// ReadByte implements io.ByteReader
func (v *StreamReader) ReadByte() (byte, error) {
	return v.GetByte(), nil
}

// ReadBytes reads multiple bytes
func (v *StreamReader) ReadBytes(count int) ([]byte, error) {
	result := make([]byte, count)
	for i := 0; i < count; i++ {
		result[i] = v.GetByte()
	}
	return result, nil
}

func (v *StreamReader) SkipBytes(count int) {
	v.position += uint64(count)
}

// Read implements io.Reader
func (v *StreamReader) Read(p []byte) (n int, err error) {
	streamLength := v.GetSize()
	for i := 0; ; i++ {
		if v.GetPosition() >= streamLength {
			return i, io.EOF
		}
		if i >= len(p) {
			return i, nil
		}
		p[i] = v.GetByte()
	}
}

func (v *StreamReader) Eof() bool {
	return v.position >= uint64(len(v.data))
}
