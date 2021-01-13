package d2datautils

import (
	"io"
)

const (
	bytesPerint16 = 2
	bytesPerint32 = 4
	bytesPerint64 = 8
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

// ReadByte reads a byte from the stream
func (v *StreamReader) ReadByte() (byte, error) {
	if v.position >= v.Size() {
		return 0, io.EOF
	}

	result := v.data[v.position]
	v.position++

	return result, nil
}

// ReadInt16 returns a int16 word from the stream
func (v *StreamReader) ReadInt16() (int16, error) {
	b, err := v.ReadUInt16()
	return int16(b), err
}

// ReadUInt16 returns a uint16 word from the stream
func (v *StreamReader) ReadUInt16() (uint16, error) {
	b, err := v.ReadBytes(bytesPerint16)
	if err != nil {
		return 0, err
	}

	return uint16(b[0]) | uint16(b[1])<<8, err
}

// ReadInt32 returns an int32 dword from the stream
func (v *StreamReader) ReadInt32() (int32, error) {
	b, err := v.ReadUInt32()
	return int32(b), err
}

// ReadUInt32 returns a uint32 dword from the stream
//nolint
func (v *StreamReader) ReadUInt32() (uint32, error) {
	b, err := v.ReadBytes(bytesPerint32)
	if err != nil {
		return 0, err
	}

	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24, err
}

// ReadInt64 returns a uint64 qword from the stream
func (v *StreamReader) ReadInt64() (int64, error) {
	b, err := v.ReadUInt64()
	return int64(b), err
}

// ReadUInt64 returns a uint64 qword from the stream
//nolint
func (v *StreamReader) ReadUInt64() (uint64, error) {
	b, err := v.ReadBytes(bytesPerint64)
	if err != nil {
		return 0, err
	}


	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56, err
}

// Position returns the current stream position
func (v *StreamReader) Position() uint64 {
	return v.position
}

// SetPosition sets the stream position with the given position
func (v *StreamReader) SetPosition(newPosition uint64) {
	v.position = newPosition
}

// Size returns the total size of the stream in bytes
func (v *StreamReader) Size() uint64 {
	return uint64(len(v.data))
}

// ReadBytes reads multiple bytes
func (v *StreamReader) ReadBytes(count int) ([]byte, error) {
	size := v.Size()
	if v.position >= size || v.position+uint64(count) > size {
		return nil, io.EOF
	}

	result := v.data[v.position : v.position+uint64(count)]
	v.position += uint64(count)

	return result, nil
}

// SkipBytes moves the stream position forward by the given amount
func (v *StreamReader) SkipBytes(count int) {
	v.position += uint64(count)
}

// Read implements io.Reader
func (v *StreamReader) Read(p []byte) (n int, err error) {
	streamLength := v.Size()

	for i := 0; ; i++ {
		if v.Position() >= streamLength {
			return i, io.EOF
		}

		if i >= len(p) {
			return i, nil
		}

		p[i], err = v.ReadByte()
		if err != nil {
			return i, err
		}
	}
}

// EOF returns if the stream position is reached to the end of the data, or not
func (v *StreamReader) EOF() bool {
	return v.position >= uint64(len(v.data))
}
