package Common

import (
	"bytes"
	"encoding/binary"
	"io"
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

// GetWord returns a uint16 word from the stream
func (v *StreamReader) GetWord() uint16 {
	result := uint16(v.data[v.position])
	result += (uint16(v.data[v.position+1]) << 8)
	v.position += 2
	return result
}

// GetSWord returns a int16 word from the stream
func (v *StreamReader) GetSWord() int16 {
	var result int16
	binary.Read(bytes.NewReader([]byte{v.data[v.position], v.data[v.position+1]}), binary.LittleEndian, &result)
	v.position += 2
	return result
}

// GetDword returns a uint32 dword from the stream
func (v *StreamReader) GetDword() uint32 {
	result := uint32(v.data[v.position])
	result += (uint32(v.data[v.position+1]) << 8)
	result += (uint32(v.data[v.position+2]) << 16)
	result += (uint32(v.data[v.position+3]) << 24)
	v.position += 4
	return result
}

// ReadByte implements io.ByteReader
func (v *StreamReader) ReadByte() (byte, error) {
	return v.GetByte(), nil
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
