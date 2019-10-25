package Common

import "log"

// BitStream is a utility class for reading groups of bits from a stream
type BitStream struct {
	data         []byte
	dataPosition int
	current      int
	bitCount     int
}

// CreateBitStream creates a new BitStream
func CreateBitStream(newData []byte) *BitStream {
	result := &BitStream{
		data:         newData,
		dataPosition: 0,
		current:      0,
		bitCount:     0,
	}
	return result
}

// ReadBits reads the specified number of bits and returns the value
func (v *BitStream) ReadBits(bitCount int) int {
	if bitCount > 16 {
		log.Panic("Maximum BitCount is 16")
	}
	if !v.EnsureBits(bitCount) {
		return -1
	}
	result := v.current & (0xffff >> (16 - bitCount))
	v.WasteBits(bitCount)
	return result
}

// PeekByte returns the current byte without adjusting the position
func (v *BitStream) PeekByte() int {
	if !v.EnsureBits(8) {
		return -1
	}
	return v.current & 0xff
}

// EnsureBits ensures that the specified number of bits are available
func (v *BitStream) EnsureBits(bitCount int) bool {
	if bitCount <= v.bitCount {
		return true
	}
	if v.dataPosition >= len(v.data) {
		return false
	}
	nextvalue := v.data[v.dataPosition]
	v.dataPosition++
	v.current |= int(nextvalue) << v.bitCount
	v.bitCount += 8
	return true
}

// WasteBits dry-reads the specified number of bits
func (v *BitStream) WasteBits(bitCount int) {
	v.current >>= bitCount
	v.bitCount -= bitCount
}
