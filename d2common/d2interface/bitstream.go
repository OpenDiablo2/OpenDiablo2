package d2interface

type BitStream interface {
	ReadBits(bitCount int) int
	PeekByte() int
	EnsureBits(bitCount int) bool
	WasteBits(bitCount int)
}
