package common

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/palettedefs"
)

// PaletteRGB represents a color in a palette
type PaletteRGB struct {
	R, G, B uint8
}

// PaletteType represents a palette
type PaletteRec struct {
	Name   palettedefs.PaletteType
	Colors [256]PaletteRGB
}

var Palettes map[palettedefs.PaletteType]PaletteRec

// CreatePalette creates a palette
func CreatePalette(name palettedefs.PaletteType, data []byte) PaletteRec {
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

func LoadPalettes(mpqFiles map[string]string, fileProvider FileProvider) {
	Palettes = make(map[palettedefs.PaletteType]PaletteRec)
	for _, pal := range []string{
		"act1", "act2", "act3", "act4", "act5", "endgame", "endgame2", "fechar", "loading",
		"menu0", "menu1", "menu2", "menu3", "menu4", "sky", "static", "trademark", "units",
	} {
		filePath := `data\global\palette\` + pal + `\pal.dat`
		paletteName := palettedefs.PaletteType(pal)
		palette := CreatePalette(paletteName, fileProvider.LoadFile(filePath))
		Palettes[paletteName] = palette
	}
	log.Printf("Loaded %d palettes", len(Palettes))
}
