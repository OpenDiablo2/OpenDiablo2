package d2datautils

import "bytes"

// StreamWriter allows you to create a byte array by streaming in writes of various sizes
type StreamWriter struct {
	data      *bytes.Buffer
	bitOffset int
	bitCache  byte
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

// PushBytes writes a bytes to the stream
func (v *StreamWriter) PushBytes(b ...byte) {
	for _, i := range b {
		v.data.WriteByte(i)
	}
}

func (v *StreamWriter) PushBit(b bool) {
	if b {
		v.bitCache |= (1 << v.bitOffset)
	}
	v.bitOffset++

	if v.bitOffset != bitsPerByte {
		return
	}

	v.PushBytes(v.bitCache)
	v.bitCache = 0
	v.bitOffset = 0
}

func (v *StreamWriter) PushBits(b byte, bits int) {
	val := b
	for i := 0; i < bits; i++ {
		if val&1 == 1 {
			v.PushBit(true)
		} else {
			v.PushBit(false)
		}

		val >>= 1
	}
}

func (v *StreamWriter) PushBits16(b uint16, bits int) {
	val := b
	for i := 0; i < bits; i++ {
		if val&1 == 1 {
			v.PushBit(true)
		} else {
			v.PushBit(false)
		}
		val >>= 1
	}
}

func (v *StreamWriter) PushBits32(b uint32, bits int) {
	val := b
	for i := 0; i < bits; i++ {
		if val&1 == 1 {
			v.PushBit(true)
		} else {
			v.PushBit(false)
		}
		val >>= 1
	}
}

func (v *StreamWriter) ForcePushBits() {
	for i := 0; i < bitsPerByte-v.bitOffset; i++ {
		v.PushBit(0 != 0)
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
