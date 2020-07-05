package d2asset

import (
	"errors"
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	fontBudget = 64
)

type fontManager struct {
	assetManager d2interface.AssetManager
	cache        d2interface.Cache
}

func createFontManager() d2interface.ArchivedFontManager {
	return &fontManager{
		cache: d2common.CreateCache(fontBudget),
	}
}

// Bind to an asset manager
func (fm *fontManager) Bind(manager d2interface.AssetManager) error {
	if fm.assetManager != nil {
		return errors.New("font manager already bound to an asset manager")
	}

	fm.assetManager = manager

	return nil
}

// LoadFont loads a font from the archives managed by the ArchiveManager
func (fm *fontManager) LoadFont(tablePath, spritePath, palettePath string) (d2interface.Font,
	error) {
	cachePath := fmt.Sprintf("%s;%s;%s", tablePath, spritePath, palettePath)
	if font, found := fm.cache.Retrieve(cachePath); found {
		return font.(d2interface.Font).Clone(), nil
	}

	font, err := loadFont(tablePath, spritePath, palettePath)
	if err != nil {
		return nil, err
	}

	if err := fm.cache.Insert(cachePath, font.Clone(), 1); err != nil {
		return nil, err
	}

	return font, nil
}

// ClearCache clears the font cache
func (fm *fontManager) ClearCache() {
	fm.cache.Clear()
}

// GetCache returns the font managers cache
func (fm *fontManager) GetCache() d2interface.Cache {
	return fm.cache
}
