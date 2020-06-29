package d2dc6

// DC6Frame represents a single frame in a DC6.
type DC6Frame struct {
	Flipped    uint32 `struct:"uint32"`
	Width      uint32 `struct:"uint32"`
	Height     uint32 `struct:"uint32"`
	OffsetX    int32  `struct:"int32"`
	OffsetY    int32  `struct:"int32"`
	Unknown    uint32 `struct:"uint32"`
	NextBlock  uint32 `struct:"uint32"`
	Length     uint32 `struct:"uint32,sizeof=FrameData"`
	FrameData  []byte
	Terminator []byte `struct:"[3]byte"`
}
