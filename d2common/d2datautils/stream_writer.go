package d2datautils

import "bytes"

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

// GetBytes returns the the byte slice of the underlying data
func (v *StreamWriter) GetBytes() []byte {
	return v.data.Bytes()
}

// PushByte writes a byte to the stream
func (v *StreamWriter) PushByte(val byte) {
	v.data.WriteByte(val)
}

// PushBytes writes a byte slince to the stream
func (v *StreamWriter) PushBytes(b []byte) {
	for _, i := range b {
		v.PushByte(i)
	}
}

// PushInt16 writes a int16 word to the stream
func (v *StreamWriter) PushInt16(val int16) {
	v.PushUint16(uint16(val))
}

// PushUint16 writes an uint16 word to the stream
//nolint
func (v *StreamWriter) PushUint16(val uint16) {
	v.data.WriteByte(byte(val))
	v.data.WriteByte(byte(val >> 8))
}

// PushInt32 writes a int32 dword to the stream
func (v *StreamWriter) PushInt32(val int32) {
	v.PushUint32(uint32(val))
}

// PushUint32 writes a uint32 dword to the stream
//nolint
func (v *StreamWriter) PushUint32(val uint32) {
	v.data.WriteByte(byte(val))
	v.data.WriteByte(byte(val >> 8))
	v.data.WriteByte(byte(val >> 16))
	v.data.WriteByte(byte(val >> 24))
}

// PushInt64 writes a uint64 qword to the stream
func (v *StreamWriter) PushInt64(val int64) {
	v.PushUint64(uint64(val))
}

// PushUint64 writes a uint64 qword to the stream
//nolint
func (v *StreamWriter) PushUint64(val uint64) {
	v.data.WriteByte(byte(val))
	v.data.WriteByte(byte(val >> 8))
	v.data.WriteByte(byte(val >> 16))
	v.data.WriteByte(byte(val >> 24))
	v.data.WriteByte(byte(val >> 32))
	v.data.WriteByte(byte(val >> 40))
	v.data.WriteByte(byte(val >> 48))
	v.data.WriteByte(byte(val >> 56))
}
