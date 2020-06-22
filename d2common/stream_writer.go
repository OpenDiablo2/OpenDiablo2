package d2common

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

// PushByte writes a byte to the stream
func (v *StreamWriter) PushByte(val byte) {
	v.data.WriteByte(val)
}

// PushUint16 writes an uint16 word to the stream
func (v *StreamWriter) PushUint16(val uint16) {
	v.data.WriteByte(byte(val) & 0xFF)
	v.data.WriteByte(byte(val>>8) & 0xFF)
}

// PushInt16 writes a int16 word to the stream
func (v *StreamWriter) PushInt16(val int16) {
	v.data.WriteByte(byte(val) & 0xFF)
	v.data.WriteByte(byte(val>>8) & 0xFF)
}

// PushUint32 writes a uint32 dword to the stream
func (v *StreamWriter) PushUint32(val uint32) {
	v.data.WriteByte(byte(val) & 0xFF)
	v.data.WriteByte(byte(val>>8) & 0xFF)
	v.data.WriteByte(byte(val>>16) & 0xFF)
	v.data.WriteByte(byte(val>>24) & 0xFF)
}

// PushUint64 writes a uint64 qword to the stream
func (v *StreamWriter) PushUint64(val uint64) {
	v.data.WriteByte(byte(val) & 0xFF)
	v.data.WriteByte(byte(val>>8) & 0xFF)
	v.data.WriteByte(byte(val>>16) & 0xFF)
	v.data.WriteByte(byte(val>>24) & 0xFF)
	v.data.WriteByte(byte(val>>32) & 0xFF)
	v.data.WriteByte(byte(val>>40) & 0xFF)
	v.data.WriteByte(byte(val>>48) & 0xFF)
	v.data.WriteByte(byte(val>>56) & 0xFF)
}

// PushInt64 writes a uint64 qword to the stream
func (v *StreamWriter) PushInt64(val int64) {
	v.PushUint64(uint64(val))
}

// GetBytes returns the the byte slice of the underlying data
func (v *StreamWriter) GetBytes() []byte {
	return v.data.Bytes()
}
