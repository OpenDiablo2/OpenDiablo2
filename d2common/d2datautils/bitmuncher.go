package d2datautils

// BitMuncher is used for parsing files that are not byte-aligned such as the DCC files.
type BitMuncher struct {
	data     []byte
	offset   int
	bitsRead int
}

const (
	twosComplimentNegativeOne = 4294967295
	byteLen                   = 8
	oneBit                    = 0x01
	fourBytes                 = byteLen * 4
)

// CreateBitMuncher Creates a BitMuncher
func CreateBitMuncher(data []byte, offset int) *BitMuncher {
	return (&BitMuncher{}).Init(data, offset)
}

// CopyBitMuncher Creates a copy of the source BitMuncher
func CopyBitMuncher(source *BitMuncher) *BitMuncher {
	return source.Copy()
}

// Init initializes the BitMuncher with data and an offset
func (v *BitMuncher) Init(data []byte, offset int) *BitMuncher {
	v.data = data
	v.offset = offset
	v.bitsRead = 0

	return v
}

// Copy returns a copy of a BitMuncher
func (v BitMuncher) Copy() *BitMuncher {
	v.bitsRead = 0
	return &v
}

// Offset returns the offset of the BitMuncher
func (v *BitMuncher) Offset() int {
	return v.offset
}

// SetOffset sets the offset of the BitMuncher
func (v *BitMuncher) SetOffset(n int) {
	v.offset = n
}

// BitsRead returns the number of bits the BitMuncher has read
func (v *BitMuncher) BitsRead() int {
	return v.bitsRead
}

// SetBitsRead sets the number of bits the BitMuncher has read
func (v *BitMuncher) SetBitsRead(n int) {
	v.bitsRead = n
}

// GetBit reads a bit and returns it as uint32
func (v *BitMuncher) GetBit() uint32 {
	result := uint32(v.data[v.offset/byteLen]>>uint(v.offset%byteLen)) & oneBit
	v.offset++
	v.bitsRead++

	return result
}

// SkipBits skips bits, incrementing the offset and bits read
func (v *BitMuncher) SkipBits(bits int) {
	v.offset += bits
	v.bitsRead += bits
}

// GetByte reads a byte from data
func (v *BitMuncher) GetByte() byte {
	return byte(v.GetBits(byteLen))
}

// GetInt32 reads an int32 from data
func (v *BitMuncher) GetInt32() int32 {
	return v.MakeSigned(v.GetBits(fourBytes), fourBytes)
}

// GetUInt32 reads an unsigned uint32 from data
func (v *BitMuncher) GetUInt32() uint32 {
	return v.GetBits(fourBytes)
}

// GetBits given a number of bits to read, reads that number of
// bits and retruns as a uint32
func (v *BitMuncher) GetBits(bits int) uint32 {
	if bits == 0 {
		return 0
	}

	result := uint32(0)
	for i := 0; i < bits; i++ {
		result |= v.GetBit() << uint(i)
	}

	return result
}

// GetSignedBits Given a number of bits, reads that many bits and returns as int
func (v *BitMuncher) GetSignedBits(bits int) int {
	return int(v.MakeSigned(v.GetBits(bits), bits))
}

// MakeSigned converts a uint32 value into an int32
func (v *BitMuncher) MakeSigned(value uint32, bits int) int32 {
	if bits == 0 {
		return 0
	}
	// If its a single bit, a value of 1 is -1 automagically
	if bits == 1 {
		return -int32(value)
	}
	// If there is no sign bit, return the value as is
	if (value & (1 << uint(bits-1))) == 0 {
		return int32(value)
	}

	// We need to extend the signed bit out so that the negative value
	// representation still works with the 2s compliment rule.
	result := uint32(twosComplimentNegativeOne)

	for i := byte(0); i < byte(bits); i++ {
		if ((value >> uint(i)) & 1) == 0 {
			result -= uint32(1 << uint(i))
		}
	}

	// Force casting to a signed value
	return int32(result)
}
