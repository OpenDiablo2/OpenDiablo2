package d2asset

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const (
	fontBudget = 64
)

type fontManager struct {
	cache d2interface.Cache
}

func createFontManager() *fontManager {
	return &fontManager{d2common.CreateCache(fontBudget)}
}

func (fm *fontManager) loadFont(tablePath, spritePath, palettePath string) (*Font, error) {
	cachePath := fmt.Sprintf("%s;%s;%s", tablePath, spritePath, palettePath)
	if font, found := fm.cache.Retrieve(cachePath); found {
		return font.(*Font).Clone(), nil
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
