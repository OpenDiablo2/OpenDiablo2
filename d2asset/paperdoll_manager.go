package d2asset

import (
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
)

type paperdollManager struct {
	cache *cache
}

func createPaperdollManager() *paperdollManager {
	return &paperdollManager{cache: createCache(PaperdollBudget)}
}

func (pm *paperdollManager) loadPaperdoll(object *d2datadict.ObjectLookupRecord, palettePath string) (*Paperdoll, error) {
	palette, err := loadPalette(palettePath)
	if err != nil {
		return nil, err
	}

	return createPaperdoll(object, palette), nil
}
