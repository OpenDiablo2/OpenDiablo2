package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

type paletteManager struct {
	cache *d2common.Cache
}

const (
	paletteBudget = 64
)

func createPaletteManager() *paletteManager {
	return &paletteManager{d2common.CreateCache(paletteBudget)}
}

func (pm *paletteManager) loadPalette(palettePath string) (*d2datadict.PaletteRec, error) {
	if palette, found := pm.cache.Retrieve(palettePath); found {
		return palette.(*d2datadict.PaletteRec), nil
	}

	paletteData, err := LoadFile(palettePath)
	if err != nil {
		return nil, err
	}

	palette := d2datadict.CreatePalette("", paletteData)
	pm.cache.Insert(palettePath, &palette, 1)
	return &palette, nil
}
