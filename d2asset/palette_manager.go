package d2asset

import (
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
)

type paletteManager struct {
	cache *cache
}

func createPaletteManager() *paletteManager {
	return &paletteManager{createCache(PaletteBudget)}
}

func (pm *paletteManager) loadPalette(palettePath string) (*d2datadict.PaletteRec, error) {
	if palette, found := pm.cache.retrieve(palettePath); found {
		return palette.(*d2datadict.PaletteRec), nil
	}

	paletteData, err := LoadFile(palettePath)
	if err != nil {
		return nil, err
	}

	palette := d2datadict.CreatePalette("", paletteData)
	pm.cache.insert(palettePath, &palette, 1)
	return &palette, nil
}
