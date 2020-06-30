package d2interface

type BitMuncher interface {
	Init(data []byte, offset int) BitMuncher
	Copy() BitMuncher
	Offset() int
	SetOffset(int)
	BitsRead() int
	SetBitsRead(int)
	GetBit() uint32
	SkipBits(bits int)
	GetByte() byte
	GetInt32() int32
	GetUInt32() uint32
	GetBits(bits int) uint32
	GetSignedBits(bits int) int
	MakeSigned(value uint32, bits int) int32
}
