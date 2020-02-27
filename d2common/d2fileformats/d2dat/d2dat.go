package d2dat

type DATColor struct {
	R uint8
	G uint8
	B uint8
}

type DATPalette struct {
	Colors [256]DATColor
}

func LoadDAT(data []byte) (*DATPalette, error) {
	palette := &DATPalette{}

	for i := 0; i < 256; i++ {
		palette.Colors[i] = DATColor{B: data[i*3], G: data[i*3+1], R: data[i*3+2]}
	}

	return palette, nil
}
