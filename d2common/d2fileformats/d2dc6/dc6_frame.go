package d2dc6

// DC6Frame represents a single frame in a DC6.
type DC6Frame struct {
	FrameData  []byte
	Terminator []byte
	Flipped    uint32
	Width      uint32
	OffsetY    int32
	Unknown    uint32
	NextBlock  uint32
	Length     uint32
	Height     uint32
	OffsetX    int32
}
