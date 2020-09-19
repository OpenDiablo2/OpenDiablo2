package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

// NewAssetManager creates and assigns all necessary dependencies for the AssetManager top-level functions to work correctly
func NewAssetManager(config *d2config.Configuration) (*AssetManager, error) {
	manager := &AssetManager{
		d2loader.NewLoader(config),
		d2cache.CreateCache(tableBudget),
		d2cache.CreateCache(animationBudget),
		d2cache.CreateCache(fontBudget),
		d2cache.CreateCache(paletteBudget),
		d2cache.CreateCache(paletteTransformBudget),
		&d2records.RecordManager{},
	}

	err := manager.init()
	if err != nil {
		return nil, err
	}

	return manager, nil
}
