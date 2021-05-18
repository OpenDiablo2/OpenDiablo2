package d2datautils

import (
	"bytes"
	"log"
)

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

// Offset returns current bit offset
func (v *StreamWriter) Offset() int {
	return v.bitOffset
}

// PushBytes writes a bytes to the stream
func (v *StreamWriter) PushBytes(b ...byte) {
	for _, i := range b {
		v.data.WriteByte(i)
	}
}

// PushBit pushes single bit into stream
// WARNING: if you'll use PushBit, offset'll be less than 8, and if you'll
// use another Push... method, bits'll not be pushed
func (v *StreamWriter) PushBit(b bool) {
	if b {
		v.bitCache |= 1 << v.bitOffset
	}
	v.bitOffset++

	if v.bitOffset != bitsPerByte {
		return
	}

	v.PushBytes(v.bitCache)
	v.bitCache = 0
	v.bitOffset = 0
}

// PushBits pushes bits (with max range 8)
func (v *StreamWriter) PushBits(b byte, bits int) {
	if bits > bitsPerByte {
		log.Print("input bits number must be less (or equal) than 8")
	}

	val := b

	for i := 0; i < bits; i++ {
		v.PushBit(val&1 == 1)
		val >>= 1
	}
}

// PushBits16 pushes bits (with max range 16)
func (v *StreamWriter) PushBits16(b uint16, bits int) {
	if bits > bitsPerByte*bytesPerint16 {
		log.Print("input bits number must be less (or equal) than 16")
	}

	val := b

	for i := 0; i < bits; i++ {
		v.PushBit(val&1 == 1)
		val >>= 1
	}
}

// PushBits32 pushes bits (with max range 32)
func (v *StreamWriter) PushBits32(b uint32, bits int) {
	if bits > bitsPerByte*bytesPerint32 {
		log.Print("input bits number must be less (or equal) than 32")
	}

	val := b

	for i := 0; i < bits; i++ {
		v.PushBit(val&1 == 1)
		val >>= 1
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

// Align aligns stream writer to bytes
func (v *StreamWriter) Align() {
	if o := v.bitOffset % bitsPerByte; o > 0 {
		v.PushBits(0, bitsPerByte-o)
	}
}
