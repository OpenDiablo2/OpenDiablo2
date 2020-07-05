package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type paletteManager struct {
	cache d2interface.Cache
}

const (
	paletteBudget = 64
)

func createPaletteManager() *paletteManager {
	return &paletteManager{d2common.CreateCache(paletteBudget)}
}

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
