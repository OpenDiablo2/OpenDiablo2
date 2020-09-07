package d2datautils

import "bytes"

const (
	byteMask = 0xFF
)

// StreamWriter allows you to create a byte array by streaming in writes of various sizes
type StreamWriter struct {
	data *bytes.Buffer
}

// CreateStreamWriter creates a new StreamWriter instance
func CreateStreamWriter() *StreamWriter {
	result := &StreamWriter{
		data: new(bytes.Buffer),
	}

	return result
}

// PushByte writes a byte to the stream
func (v *StreamWriter) PushByte(val byte) {
	v.data.WriteByte(val)
}

// PushUint16 writes an uint16 word to the stream
func (v *StreamWriter) PushUint16(val uint16) {
	for count := 0; count < bytesPerInt16; count++ {
		shift := count * bitsPerByte
		v.data.WriteByte(byte(val>>shift) & byteMask)
	}
}

// PushInt16 writes a int16 word to the stream
func (v *StreamWriter) PushInt16(val int16) {
	for count := 0; count < bytesPerInt16; count++ {
		shift := count * bitsPerByte
		v.data.WriteByte(byte(val>>shift) & byteMask)
	}
}

// PushUint32 writes a uint32 dword to the stream
func (v *StreamWriter) PushUint32(val uint32) {
	for count := 0; count < bytesPerInt32; count++ {
		shift := count * bitsPerByte
		v.data.WriteByte(byte(val>>shift) & byteMask)
	}
}

// PushUint64 writes a uint64 qword to the stream
func (v *StreamWriter) PushUint64(val uint64) {
	for count := 0; count < bytesPerInt64; count++ {
		shift := count * bitsPerByte
		v.data.WriteByte(byte(val>>shift) & byteMask)
	}
}

// PushInt64 writes a uint64 qword to the stream
func (v *StreamWriter) PushInt64(val int64) {
	v.PushUint64(uint64(val))
}

// GetBytes returns the the byte slice of the underlying data
func (v *StreamWriter) GetBytes() []byte {
	return v.data.Bytes()
}
