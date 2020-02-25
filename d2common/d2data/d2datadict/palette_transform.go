package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
)

// PaletteTransformType represents a palette
type PaletteTransformRec = d2pl2.PL2File

var PaletteTransforms []PaletteTransformRec

// CreatePaletteTransform creates a palette
func CreatePaletteTransform(data []byte) (PaletteTransformRec, error) {
	result, err := d2pl2.LoadPL2(data)

	return result, err
}

func LoadPaletteTransform(file []byte) {
	if transform, err := CreatePaletteTransform(file); err == nil {
		PaletteTransforms = append(PaletteTransforms, transform)
	}
}
