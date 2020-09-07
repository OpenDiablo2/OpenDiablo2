package d2datautils

import (
	"io"
)

const (
	bytesPerInt16 = 2
	bytesPerInt32 = 4
	bytesPerInt64 = 8
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
	var result uint16

	for offset := uint64(0); offset < bytesPerInt16; offset++ {
		shift := uint8(bitsPerByte * offset)
		result += uint16(v.data[v.position+offset]) << shift
	}

	v.position += bytesPerInt16

	return result
}

// GetInt16 returns a int16 word from the stream
func (v *StreamReader) GetInt16() int16 {
	var result int16

	for offset := uint64(0); offset < bytesPerInt16; offset++ {
		shift := uint8(bitsPerByte * offset)
		result += int16(v.data[v.position+offset]) << shift
	}

	v.position += bytesPerInt16

	return result
}

// SetPosition sets the stream position with the given position
func (v *StreamReader) SetPosition(newPosition uint64) {
	v.position = newPosition
}

// GetUInt32 returns a uint32 dword from the stream
func (v *StreamReader) GetUInt32() uint32 {
	var result uint32

	for offset := uint64(0); offset < bytesPerInt32; offset++ {
		shift := uint8(bitsPerByte * offset)
		result += uint32(v.data[v.position+offset]) << shift
	}

	v.position += bytesPerInt32

	return result
}

// GetInt32 returns an int32 dword from the stream
func (v *StreamReader) GetInt32() int32 {
	var result int32

	for offset := uint64(0); offset < bytesPerInt32; offset++ {
		shift := uint8(bitsPerByte * offset)
		result += int32(v.data[v.position+offset]) << shift
	}

	v.position += bytesPerInt32

	return result
}

// GetUint64 returns a uint64 qword from the stream
func (v *StreamReader) GetUint64() uint64 {
	var result uint64

	for offset := uint64(0); offset < bytesPerInt64; offset++ {
		shift := uint8(bitsPerByte * offset)
		result += uint64(v.data[v.position+offset]) << shift
	}

	v.position += bytesPerInt64

	return result
}

// GetInt64 returns a uint64 qword from the stream
func (v *StreamReader) GetInt64() int64 {
	return int64(v.GetUint64())
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
