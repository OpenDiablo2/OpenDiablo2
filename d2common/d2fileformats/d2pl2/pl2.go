package d2pl2

import (
	"encoding/binary"

	"github.com/go-restruct/restruct"
)

type PL2File struct {
	BasePalette					PL2Palette

	LightLevelVariations		[32]PL2PaletteTransform
	InvColorVariations			[16]PL2PaletteTransform
	SelectedUintShift			PL2PaletteTransform
	AlphaBlend					[3][256]PL2PaletteTransform
	AdditiveBlend				[256]PL2PaletteTransform
	MultiplicativeBlend			[256]PL2PaletteTransform
	HueVariations				[111]PL2PaletteTransform
	RedTones					PL2PaletteTransform
	GreenTones					PL2PaletteTransform
	BlueTones					PL2PaletteTransform
	UnknownVariations			[14]PL2PaletteTransform
	MaxComponentBlend			[256]PL2PaletteTransform
	DarkendColorShift			PL2PaletteTransform

	TextColors					[13]PL2Color24Bits
	TextColorShifts				[13]PL2PaletteTransform
}

type PL2Color struct {
	R	uint8
	G	uint8
	B	uint8
	_	uint8
}

type PL2Color24Bits struct {
	R	uint8
	G	uint8
	B	uint8
}

type PL2Palette struct {
	Colors		[256]PL2Color
}

type PL2PaletteTransform struct {
	Indices		[256]uint8
}

// uses restruct to read the binary dc6 data into structs
func LoadPL2(data []byte) (*PL2File, error) {
	result := &PL2File{}

	restruct.EnableExprBeta()

	err := restruct.Unpack(data, binary.LittleEndian, &result)
	if err != nil {
	  return nil, err
	}

	return result, nil
}
