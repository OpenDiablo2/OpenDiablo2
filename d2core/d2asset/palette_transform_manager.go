package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

type paletteTransformManager struct {
	cache *d2common.Cache
}

const (
	paletteTransformBudget = 64
)

func createPaletteTransformManager() *paletteTransformManager {
	return &paletteTransformManager{d2common.CreateCache(paletteTransformBudget)}
}

func (pm *paletteTransformManager) loadPaletteTransform(path string) (*d2datadict.PaletteTransformRec, error) {
	if paletteTransform, found := pm.cache.Retrieve(path); found {
		return paletteTransform.(*d2datadict.PaletteTransformRec), nil
	}

	paletteTransformData, err := LoadFile(path)
	if err != nil {
		return nil, err
	}

	paletteTransform := d2datadict.CreatePaletteTransform("", paletteTransformData)
	pm.cache.Insert(path, &paletteTransform, 1)
	return &paletteTransform, nil
}
