package d2pl2

import (
	"testing"
)

func exampleData() *PL2 {
	result := &PL2{
		BasePalette:       PL2Palette{},
		SelectedUintShift: PL2PaletteTransform{},
		RedTones:          PL2PaletteTransform{},
		GreenTones:        PL2PaletteTransform{},
		BlueTones:         PL2PaletteTransform{},
		DarkendColorShift: PL2PaletteTransform{},
	}

	result.BasePalette.Colors[0].R = 8
	result.DarkendColorShift.Indices[0] = 123

	return result
}

func TestPL2_MarshalUnmarshal(t *testing.T) {
	pl2 := exampleData()

	data := pl2.Marshal()

	newPL2, err := Load(data)
	if err != nil {
		t.Error(err)
	}

	if newPL2.BasePalette.Colors[0] != pl2.BasePalette.Colors[0] {
		t.Fatal("unexpected length")
	}

	if pl2.DarkendColorShift.Indices[0] != newPL2.DarkendColorShift.Indices[0] {
	}
}
