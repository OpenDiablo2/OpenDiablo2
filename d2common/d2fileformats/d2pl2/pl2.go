package d2pl2

import (
	"log"
	"encoding/binary"

	"github.com/go-restruct/restruct"
)

type PL2File struct {
	basePalette					PL2Palette

	lightLevelVariations		[32]PL2PaletteTransform
	invColorVariations			[16]PL2PaletteTransform
	selectedUintShift			PL2PaletteTransform
	alphaBlend					[3][256]PL2PaletteTransform
	additiveBlend				[256]PL2PaletteTransform
	multiplicativeBlend			[256]PL2PaletteTransform
	hueVariations				[111]PL2PaletteTransform
	redTones					PL2PaletteTransform
	greenTones					PL2PaletteTransform
	blueTones					PL2PaletteTransform
	unknownVariations			[14]PL2PaletteTransform
	maxComponentBlend			[256]PL2PaletteTransform
	darkendColorShift			PL2PaletteTransform

	textColors					[13]PL2Color24Bits
	textColorShifts				[13]PL2PaletteTransform
}

type PL2Color struct {
	r	uint8
	g	uint8
	b	uint8
	_	uint8
}

type PL2Color24Bits struct {
	r	uint8
	g	uint8
	b	uint8
}

type PL2Palette struct {
	colors		[256]PL2Color
}

type PL2PaletteTransform struct {
	indices		[256]uint8
}

// uses restruct to read the binary dc6 data into structs
func LoadPL2(data []byte) (PL2File, error) {
	result := &PL2File{}

	restruct.EnableExprBeta()
	err := restruct.Unpack(data, binary.LittleEndian, &result)
	if err != nil {
		log.Printf("failed to read pl2: %v", err)
	}

	return *result, err
}
