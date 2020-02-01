package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// PaletteRGB represents a color in a palette
type PaletteRGB struct {
	R, G, B uint8
}

// PaletteType represents a palette
type PaletteRec struct {
	Name   d2enum.PaletteType
	Colors [256]PaletteRGB
}

var Palettes map[d2enum.PaletteType]PaletteRec

// CreatePalette creates a palette
func CreatePalette(name d2enum.PaletteType, data []byte) PaletteRec {
	result := PaletteRec{Name: name}

	for i := 0; i <= 255; i++ {
		result.Colors[i] = PaletteRGB{
			B: data[i*3],
			G: data[(i*3)+1],
			R: data[(i*3)+2],
		}
	}
	return result
}

func LoadPalette(paletteType d2enum.PaletteType, file []byte) {
	if Palettes == nil {
		Palettes = make(map[d2enum.PaletteType]PaletteRec)
	}
	palette := CreatePalette(paletteType, file)
	Palettes[paletteType] = palette

}
