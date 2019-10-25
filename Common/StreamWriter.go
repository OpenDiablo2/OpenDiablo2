package Common

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

// PushWord writes an uint16 word to the stream
func (v *StreamWriter) PushWord(val uint16) {
	v.data = append(v.data, byte(val&0xFF))
	v.data = append(v.data, byte((val>>8)&0xFF))
}

// PushSWord writes a int16 word to the stream
func (v *StreamWriter) PushSWord(val int16) {
	v.data = append(v.data, byte(val&0xFF))
	v.data = append(v.data, byte((val>>8)&0xFF))
}

// PushDword writes a uint32 dword to the stream
func (v *StreamWriter) PushDword(val uint32) {
	v.data = append(v.data, byte(val&0xFF))
	v.data = append(v.data, byte((val>>8)&0xFF))
	v.data = append(v.data, byte((val>>16)&0xFF))
	v.data = append(v.data, byte((val>>24)&0xFF))
}

// GetBytes returns the the byte slice of the underlying data
func (v *StreamWriter) GetBytes() []byte {
	return v.data
}
