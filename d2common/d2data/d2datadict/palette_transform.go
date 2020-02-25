package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// PaletteTransformType represents a paletteTransform
type PaletteTransformRec struct {
	Name	d2enum.PaletteTransformType
	Data	d2pl2.PL2File
}

var PaletteTransforms map[d2enum.PaletteTransformType]PaletteTransformRec

// CreatePaletteTransform creates a paletteTransform transform
func CreatePaletteTransform(name d2enum.PaletteTransformType, data []byte) PaletteTransformRec {
	result := PaletteTransformRec{Name: name}
	// TODO: where should this possible error go... ? LoadPL2 returns 2 args
	pl2, _:= d2pl2.LoadPL2(data)
	result.Data = pl2
	return result
}

func LoadPaletteTransform(paletteTransformType d2enum.PaletteTransformType, file []byte) {
	if PaletteTransforms == nil {
		PaletteTransforms = make(map[d2enum.PaletteTransformType]PaletteTransformRec)
	}
	paletteTransform := CreatePaletteTransform(paletteTransformType, file)
	PaletteTransforms[paletteTransformType] = paletteTransform

}
