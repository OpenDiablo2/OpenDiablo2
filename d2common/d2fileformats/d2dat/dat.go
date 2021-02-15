package d2dat

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	// index offset helpers
	b = iota
	g
	r
	o
)

// Load loads a DAT file.
func Load(data []byte) (d2interface.Palette, error) {
	palette := &DATPalette{}

	for i := 0; i < 256; i++ {
		// offsets look like i*3+n, where n is 0,1,2
		palette.colors[i] = &DATColor{b: data[i*o+b], g: data[i*o+g], r: data[i*o+r]}
	}

	return palette, nil
}

// Marshal encodes data palette back into byte slice
func (p *DATPalette) Marshal() []byte {
	result := make([]byte, len(p.colors))

	for _, i := range &p.colors {
		result = append(result, i.B(), i.G(), i.R())
	}

	return result
}
