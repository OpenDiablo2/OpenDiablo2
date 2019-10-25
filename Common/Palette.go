package Common

import "github.com/essial/OpenDiablo2/Palettes"

// PaletteRGB represents a color in a palette
type PaletteRGB struct {
	R, G, B uint8
}

// Palette represents a palette
type Palette struct {
	Name   Palettes.Palette
	Colors [256]PaletteRGB
}

// CreatePalette creates a palette
func CreatePalette(name Palettes.Palette, data []byte) Palette {
	result := Palette{Name: name}

	for i := 0; i <= 255; i++ {
		result.Colors[i] = PaletteRGB{
			B: data[i*3],
			G: data[(i*3)+1],
			R: data[(i*3)+2],
		}
	}

	return result
}
