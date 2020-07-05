package d2asset

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type paletteTransformManager struct {
	assetManager d2interface.AssetManager
	cache d2interface.Cache
}

const (
	paletteTransformBudget = 64
)

func createPaletteTransformManager() *paletteTransformManager {
	return &paletteTransformManager{cache: d2common.CreateCache(paletteTransformBudget)}
}

// Bind to an asset manager
func (pm *paletteTransformManager) Bind(manager d2interface.AssetManager) error {
	if pm.assetManager != nil {
		return errors.New("palette transfrom manager already bound to an asset manager")
	}
	pm.assetManager = manager
	return nil
}

func (pm *paletteTransformManager) LoadPaletteTransform(path string) (*d2pl2.PL2, error) {
	if pl2, found := pm.cache.Retrieve(path); found {
		return pl2.(*d2pl2.PL2), nil
	}

	data, err := pm.assetManager.LoadFile(path)
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
