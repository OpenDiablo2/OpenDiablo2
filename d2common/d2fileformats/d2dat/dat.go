package d2dat

// Load loads a DAT file.
func Load(data []byte) (*DATPalette, error) {
	palette := &DATPalette{}

	for i := 0; i < 256; i++ {
		palette.Colors[i] = DATColor{B: data[i*3], G: data[i*3+1], R: data[i*3+2]}
	}

	return palette, nil
}
