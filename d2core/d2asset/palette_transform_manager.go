package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
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

func (pm *paletteTransformManager) loadPaletteTransform(path string) (*d2pl2.PL2File, error) {
	if pl2, found := pm.cache.Retrieve(path); found {
		return pl2.(*d2pl2.PL2File), nil
	}

	data, err := LoadFile(path); 
	if err != nil {
		return nil, err
	}

	pl2, err := d2pl2.LoadPL2(data)
	if err != nil {
		return nil, err
	}

	pm.cache.Insert(path, pl2, 1)
	return pl2, nil
}
