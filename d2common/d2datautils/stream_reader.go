package d2datautils

import (
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

// GetByte returns a byte from the stream
func (v *StreamReader) GetByte() byte {
	result := v.data[v.position]
	v.position++

	return result
}

// GetInt16 returns a int16 word from the stream
func (v *StreamReader) GetInt16() int16 {
	return int16(v.GetUInt16())
}

// GetUInt16 returns a uint16 word from the stream
//nolint
func (v *StreamReader) GetUInt16() uint16 {
	b := v.ReadBytes(2)
	return uint16(b[0]) | uint16(b[1])<<8
}

// GetInt32 returns an int32 dword from the stream
func (v *StreamReader) GetInt32() int32 {
	return int32(v.GetUInt32())
}

// GetUInt32 returns a uint32 dword from the stream
//nolint
func (v *StreamReader) GetUInt32() uint32 {
	b := v.ReadBytes(4)
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

// GetInt64 returns a uint64 qword from the stream
func (v *StreamReader) GetInt64() int64 {
	return int64(v.GetUInt64())
}

// GetUInt64 returns a uint64 qword from the stream
//nolint
func (v *StreamReader) GetUInt64() uint64 {
	b := v.ReadBytes(8)
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

// GetPosition returns the current stream position
func (v *StreamReader) GetPosition() uint64 {
	return v.position
}

// SetPosition sets the stream position with the given position
func (v *StreamReader) SetPosition(newPosition uint64) {
	v.position = newPosition
}

// GetSize returns the total size of the stream in bytes
func (v *StreamReader) GetSize() uint64 {
	return uint64(len(v.data))
}

// ReadByte implements io.ByteReader
func (v *StreamReader) ReadByte() (byte, error) {
	return v.GetByte(), nil
}

// ReadBytes reads multiple bytes
func (v *StreamReader) ReadBytes(count int) []byte {
	result := v.data[v.position : v.position+uint64(count)]
	v.position += uint64(count)

	return result
}

// SkipBytes moves the stream position forward by the given amount
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

// EOF returns if the stream position is reached to the end of the data, or not
func (v *StreamReader) EOF() bool {
	return v.position >= uint64(len(v.data))
}
