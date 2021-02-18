package d2font

// NewFontGlyph creates a new font glyph
func NewFontGlyph(frame, width, height int) *FontGlyph {
	// nolint:gomnd // thes bytes are constant
	// comes from https://d2mods.info/forum/viewtopic.php?t=42044
	result := &FontGlyph{
		unknown1: []byte{0, 0},
		unknown2: []byte{1, 0, 0},
		unknown3: []byte{0, 0, 0, 0, 0},
		frame:    frame,
		width:    width,
		height:   height,
	}

	return result
}
