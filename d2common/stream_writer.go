package d2common

// StreamWriter allows you to create a byte array by streaming in writes of various sizes
type StreamWriter struct {
	data []byte
}

// CreateStreamWriter creates a new StreamWriter instance
func CreateStreamWriter() *StreamWriter {
	result := &StreamWriter{
		data: make([]byte, 0),
	}
	return result
}

// PushByte writes a byte to the stream
func (v *StreamWriter) PushByte(val byte) {
	v.data = append(v.data, val)
}

// PushUint16 writes an uint16 word to the stream
func (v *StreamWriter) PushUint16(val uint16) {
	v.data = append(v.data, byte(val&0xFF))
	v.data = append(v.data, byte((val>>8)&0xFF))
}

// PushInt16 writes a int16 word to the stream
func (v *StreamWriter) PushInt16(val int16) {
	v.data = append(v.data, byte(val&0xFF))
	v.data = append(v.data, byte((val>>8)&0xFF))
}

// PushUint32 writes a uint32 dword to the stream
func (v *StreamWriter) PushUint32(val uint32) {
	v.data = append(v.data, byte(val&0xFF))
	v.data = append(v.data, byte((val>>8)&0xFF))
	v.data = append(v.data, byte((val>>16)&0xFF))
	v.data = append(v.data, byte((val>>24)&0xFF))
}

// PushUint64 writes a uint64 qword to the stream
func (v *StreamWriter) PushUint64(val uint64) {
	v.data = append(v.data, byte(val&0xFF))
	v.data = append(v.data, byte((val>>8)&0xFF))
	v.data = append(v.data, byte((val>>16)&0xFF))
	v.data = append(v.data, byte((val>>24)&0xFF))
	v.data = append(v.data, byte((val>>32)&0xFF))
	v.data = append(v.data, byte((val>>40)&0xFF))
	v.data = append(v.data, byte((val>>48)&0xFF))
	v.data = append(v.data, byte((val>>56)&0xFF))
}

// PushInt64 writes a uint64 qword to the stream
func (v *StreamWriter) PushInt64(val int64) {
	result := uint64(val)
	v.data = append(v.data, byte(result&0xFF))
	v.data = append(v.data, byte((result>>8)&0xFF))
	v.data = append(v.data, byte((result>>16)&0xFF))
	v.data = append(v.data, byte((result>>24)&0xFF))
	v.data = append(v.data, byte((result>>32)&0xFF))
	v.data = append(v.data, byte((result>>40)&0xFF))
	v.data = append(v.data, byte((result>>48)&0xFF))
	v.data = append(v.data, byte((result>>56)&0xFF))
}

// GetBytes returns the the byte slice of the underlying data
func (v *StreamWriter) GetBytes() []byte {
	return v.data
}
