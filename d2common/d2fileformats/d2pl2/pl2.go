package d2pl2

import (
	"encoding/binary"

	"github.com/go-restruct/restruct"
)

// PL2 represents a palette file.
type PL2 struct {
	BasePalette PL2Palette

	LightLevelVariations [32]PL2PaletteTransform
	InvColorVariations   [16]PL2PaletteTransform
	SelectedUintShift    PL2PaletteTransform
	AlphaBlend           [3][256]PL2PaletteTransform
	AdditiveBlend        [256]PL2PaletteTransform
	MultiplicativeBlend  [256]PL2PaletteTransform
	HueVariations        [111]PL2PaletteTransform
	RedTones             PL2PaletteTransform
	GreenTones           PL2PaletteTransform
	BlueTones            PL2PaletteTransform
	UnknownVariations    [14]PL2PaletteTransform
	MaxComponentBlend    [256]PL2PaletteTransform
	DarkendColorShift    PL2PaletteTransform

	TextColors      [13]PL2Color24Bits
	TextColorShifts [13]PL2PaletteTransform
}

// Load uses restruct to read the binary pl2 data into structs
func Load(data []byte) (*PL2, error) {
	result := &PL2{}

	restruct.EnableExprBeta()

	err := restruct.Unpack(data, binary.LittleEndian, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
