package d2assetmanager

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

type PaletteManager struct {
	cache *d2common.Cache
}

const (
	PaletteBudget = 64
)

func CreatePaletteManager() *PaletteManager {
	return &PaletteManager{d2common.CreateCache(PaletteBudget)}
}

func (pm *PaletteManager) SetCacheVerbose(verbose bool) {
	pm.cache.SetCacheVerbose(verbose)
}

func (pm *PaletteManager) ClearCache() {
	pm.cache.Clear()
}

func (pm *PaletteManager) GetCacheWeight() int {
	return pm.cache.GetWeight()
}

func (pm *PaletteManager) GetCacheBudget() int {
	return pm.cache.GetBudget()
}

func (pm *PaletteManager) LoadPalette(palettePath string) (*d2datadict.PaletteRec, error) {
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
