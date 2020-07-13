package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// Static checks to confirm struct conforms to interface
var _ d2interface.ArchivedPaletteManager = &paletteManager{}
var _ d2interface.Cacher = &paletteManager{}

type paletteManager struct {
	cache d2interface.Cache
}

const (
	paletteBudget = 64
)

func createPaletteManager() d2interface.ArchivedPaletteManager {
	return &paletteManager{d2common.CreateCache(paletteBudget)}
}

// LoadPalette loads a palette from archives managed by the ArchiveManager
func (pm *paletteManager) LoadPalette(palettePath string) (d2interface.Palette, error) {
	if palette, found := pm.cache.Retrieve(palettePath); found {
		return palette.(d2interface.Palette), nil
	}

	paletteData, err := LoadFile(palettePath)
	if err != nil {
		return nil, err
	}

	palette, err := d2dat.Load(paletteData)
	if err != nil {
		return nil, err
	}

	if err := pm.cache.Insert(palettePath, palette, 1); err != nil {
		return nil, err
	}

	return palette, nil
}

// ClearCache clears the palette cache
func (pm *paletteManager) ClearCache() {
	pm.cache.Clear()
}

// GetCache returns the palette managers cache
func (pm *paletteManager) GetCache() d2interface.Cache {
	return pm.cache
}
