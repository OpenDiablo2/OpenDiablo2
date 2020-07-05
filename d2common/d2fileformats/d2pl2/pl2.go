package d2pl2

import (
	"encoding/binary"

	"github.com/go-restruct/restruct"
)

// PL2 represents a palette file.
type PL2 struct {
	basePalette PL2Palette

	lightLevelVariations [32]PL2PaletteTransform
	invColorVariations   [16]PL2PaletteTransform
	selectedUintShift    PL2PaletteTransform
	alphaBlend           [3][256]PL2PaletteTransform
	additiveBlend        [256]PL2PaletteTransform
	multiplicativeBlend  [256]PL2PaletteTransform
	hueVariations        [111]PL2PaletteTransform
	redTones             PL2PaletteTransform
	greenTones           PL2PaletteTransform
	blueTones            PL2PaletteTransform
	unknownVariations    [14]PL2PaletteTransform
	maxComponentBlend    [256]PL2PaletteTransform
	darkendColorShift    PL2PaletteTransform

	textColors      [13]PL2Color24Bits
	textColorShifts [13]PL2PaletteTransform
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
