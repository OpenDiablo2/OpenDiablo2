package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
)

// PaletteType represents a palette
type PaletteRec struct {
	Name   d2enum.PaletteType
	Colors [256]d2dat.DATColor
}

var Palettes map[d2enum.PaletteType]PaletteRec

// CreatePalette creates a palette
func CreatePalette(name d2enum.PaletteType, data []byte) PaletteRec {
	palette, _ := d2dat.LoadDAT(data)
	return PaletteRec{Name: name, Colors: palette.Colors}
}

func LoadPalette(paletteType d2enum.PaletteType, file []byte) {
	if Palettes == nil {
		Palettes = make(map[d2enum.PaletteType]PaletteRec)
	}
	palette := CreatePalette(paletteType, file)
	Palettes[paletteType] = palette
}
