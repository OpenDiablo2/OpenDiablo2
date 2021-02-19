// Package d2fontglyph represents a single font glyph
package d2fontglyph

// Create creates a new font glyph
func Create(frame, width, height int) *FontGlyph {
	// nolint:gomnd // thes bytes are constant
	// comes from https://d2mods.info/forum/viewtopic.php?t=42044
	result := &FontGlyph{
		unknown1: []byte{0},
		unknown2: []byte{1, 0, 0},
		unknown3: []byte{0, 0, 0, 0, 0},
		frame:    frame,
		width:    width,
		height:   height,
	}

	return result
}

// FontGlyph represents a single font glyph
type FontGlyph struct {
	unknown1 []byte
	unknown2 []byte
	unknown3 []byte
	frame    int
	width    int
	height   int
}

// SetSize sets glyph's size to w, h
func (fg *FontGlyph) SetSize(w, h int) {
	fg.width, fg.height = w, h
}

// Size returns glyph's size
func (fg *FontGlyph) Size() (w, h int) {
	return fg.width, fg.height
}

// Width returns font width
func (fg *FontGlyph) Width() int {
	return fg.width
}

// Height returns glyph's height
func (fg *FontGlyph) Height() int {
	return fg.height
}

// SetFrameIndex sets frame index to idx
func (fg *FontGlyph) SetFrameIndex(idx int) {
	fg.frame = idx
}

// FrameIndex returns glyph's frame
func (fg *FontGlyph) FrameIndex() int {
	return fg.frame
}

// Unknown1 returns unknowns bytes
func (fg *FontGlyph) Unknown1() []byte {
	return fg.unknown1
}

// Unknown2 returns unknowns bytes
func (fg *FontGlyph) Unknown2() []byte {
	return fg.unknown2
}

// Unknown3 returns unknowns bytes
func (fg *FontGlyph) Unknown3() []byte {
	return fg.unknown3
}
