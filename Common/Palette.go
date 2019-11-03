package Common

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/PaletteDefs"
)

// PaletteRGB represents a color in a palette
type PaletteRGB struct {
	R, G, B uint8
}

// PaletteType represents a palette
type PaletteRec struct {
	Name   PaletteDefs.PaletteType
	Colors [256]PaletteRGB
}

var Palettes map[PaletteDefs.PaletteType]PaletteRec

// CreatePalette creates a palette
func CreatePalette(name PaletteDefs.PaletteType, data []byte) PaletteRec {
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

func LoadPalettes(mpqFiles map[string]*MpqFileRecord, fileProvider FileProvider) {
	Palettes = make(map[PaletteDefs.PaletteType]PaletteRec)
	for file := range mpqFiles {
		if strings.Index(file, "/data/global/palette/") != 0 || strings.Index(file, ".dat") != len(file)-4 {
			continue
		}
		nameParts := strings.Split(file, `/`)
		paletteName := PaletteDefs.PaletteType(nameParts[len(nameParts)-2])
		palette := CreatePalette(paletteName, fileProvider.LoadFile(file))
		Palettes[paletteName] = palette
	}
	log.Printf("Loaded %d palettes", len(Palettes))
}
