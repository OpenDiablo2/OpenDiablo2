package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type paletteTransformManager struct {
	*AssetManager
	cache d2interface.Cache
}

const (
	paletteTransformBudget = 64
)

func (pm *paletteTransformManager) loadPaletteTransform(path string) (*d2pl2.PL2, error) {
	if pl2, found := pm.cache.Retrieve(path); found {
		return pl2.(*d2pl2.PL2), nil
	}

	data, err := pm.LoadFile(path)
	if err != nil {
		return nil, err
	}

	pl2, err := d2pl2.Load(data)
	if err != nil {
		return nil, err
	}

	if err := pm.cache.Insert(path, pl2, 1); err != nil {
		return nil, err
	}

	return pl2, nil
}
