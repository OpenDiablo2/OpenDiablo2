package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

// NewAssetManager creates and assigns all necessary dependencies for the AssetManager top-level functions to work correctly
func NewAssetManager(renderer d2interface.Renderer,
	term d2interface.Terminal) (*AssetManager, error) {

	manager := &AssetManager{}

	manager.archiveManager = &archiveManager{
		AssetManager: manager,
		cache:        d2cache.CreateCache(archiveBudget),
		config:       d2config.Config,
	}

	manager.archivedFileManager = &fileManager{
		manager,
		d2cache.CreateCache(fileBudget),
		manager.archiveManager,
		d2config.Config,
	}

	manager.paletteManager = &paletteManager{
		manager,
		d2cache.CreateCache(paletteBudget),
	}

	manager.paletteTransformManager = &paletteTransformManager{
		manager,
		d2cache.CreateCache(paletteTransformBudget),
	}

	manager.animationManager = &animationManager{
		AssetManager: manager,
		renderer:     renderer,
		cache:        d2cache.CreateCache(animationBudget),
	}

	manager.fontManager = &fontManager{manager, d2cache.CreateCache(fontBudget)}

	if term != nil {
		return manager, manager.BindTerminalCommands(term)
	}

	return manager, nil
}
