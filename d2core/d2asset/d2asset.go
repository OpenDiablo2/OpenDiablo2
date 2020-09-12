package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

// NewAssetManager creates and assigns all necessary dependencies for the AssetManager top-level functions to work correctly
func NewAssetManager(renderer d2interface.Renderer, config *d2config.Configuration,
	term d2interface.Terminal) (*AssetManager, error) {
	manager := &AssetManager{
		renderer,
		d2loader.NewLoader(config),
		d2cache.CreateCache(animationBudget),
		d2cache.CreateCache(tableBudget),
		d2cache.CreateCache(fontBudget),
		d2cache.CreateCache(paletteBudget),
		d2cache.CreateCache(paletteTransformBudget),
	}

	if term != nil {
		err := manager.BindTerminalCommands(term)
		return manager, err
	}

	return manager, nil
}
