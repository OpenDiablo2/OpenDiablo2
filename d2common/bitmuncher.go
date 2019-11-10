package d2common

type BitMuncher struct {
	data     []byte
	Offset   int
	BitsRead int
}

func CreateBitMuncher(data []byte, offset int) *BitMuncher {
	return &BitMuncher{
		data:     data,
		Offset:   offset,
		BitsRead: 0,
	}
}

func CopyBitMuncher(source *BitMuncher) *BitMuncher {
	return &BitMuncher{
		source.data,
		source.Offset,
		0,
	}
}

func (v *BitMuncher) GetBit() uint32 {
	result := uint32(v.data[v.Offset/8]>>uint(v.Offset%8)) & 0x01
	v.Offset++
	v.BitsRead++
	return result
}

func (v *BitMuncher) SkipBits(bits int) {
	v.Offset += bits
	v.BitsRead += bits
}

func (v *BitMuncher) GetByte() byte {
	return byte(v.GetBits(8))
}

func (v *BitMuncher) GetInt32() int32 {
	return v.MakeSigned(v.GetBits(32), 32)
}

func (v *BitMuncher) GetUInt32() uint32 {
	return v.GetBits(32)
}

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

func (v *BitMuncher) GetSignedBits(bits int) int {
	return int(v.MakeSigned(v.GetBits(bits), bits))
}

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
	// We need to extend the signed bit out so that the negative value  representation still works with the 2s compliment rule.
	result := uint32(4294967295)
	for i := byte(0); i < byte(bits); i++ {
		if ((value >> uint(i)) & 1) == 0 {
			result -= uint32(1 << uint(i))
		}
	}
	// Force casting to a signed value
	return int32(result)
}
