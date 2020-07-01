package d2asset

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const (
	animationBudget = 64
)

type animationManager struct {
	cache *d2common.Cache
}

func createAnimationManager() *animationManager {
	return &animationManager{d2common.CreateCache(animationBudget)}
}

// Dispose of all the animations to prevent ebiten storing the images
func (am *animationManager) ClearCache() {
	keys := am.cache.GetKeys()
	for idx := range keys {
		key := keys[idx]

		//TODO: move fonts in another cache ?
		if strings.HasPrefix(key, "/data/local/FONT/") {
			continue;
		}
		animation, _:= am.cache.Retrieve(key)
		animation.(*Animation).Dispose()
	}

	am.cache.Clear()
}

func (am *animationManager) loadAnimation(animationPath, palettePath string, transparency int) (*Animation, error) {
	cachePath := fmt.Sprintf("%s;%s;%d", animationPath, palettePath, transparency)
	if animation, found := am.cache.Retrieve(cachePath); found {
			return animation.(*Animation), nil
	}

	var animation *Animation

	ext := strings.ToLower(filepath.Ext(animationPath))
	switch ext {
	case ".dc6":
		dc6, err := loadDC6(animationPath)
		if err != nil {
			return nil, err
		}

		palette, err := LoadPalette(palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = CreateAnimationFromDC6(dc6, palette)
		if err != nil {
			return nil, err
		}
	case ".dcc":
		dcc, err := loadDCC(animationPath)
		if err != nil {
			return nil, err
		}

		palette, err := LoadPalette(palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = CreateAnimationFromDCC(dcc, palette, transparency)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown animation format: %s", ext)
	}

	if err := am.cache.Insert(cachePath, animation, 1); err != nil {
		return nil, err
	}

	animation.SetKey(cachePath)

	return animation, nil
}
