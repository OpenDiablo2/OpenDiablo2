package d2dcc

// DCCPixelBufferEntry represents a single entry in the pixel buffer.
type DCCPixelBufferEntry struct {
	Value          [4]byte
	Frame          int
	FrameCellIndex int
}
