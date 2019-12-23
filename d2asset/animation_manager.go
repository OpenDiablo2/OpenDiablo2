package d2asset

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

type animationManager struct {
	cache *cache
}

func createAnimationManager() *animationManager {
	return &animationManager{cache: createCache(AnimationBudget)}
}

func (sm *animationManager) loadAnimation(animationPath, palettePath string, transparency int) (*Animation, error) {
	cachePath := fmt.Sprintf("%s;%s;%d", animationPath, palettePath, transparency)
	if animation, found := sm.cache.retrieve(cachePath); found {
		return animation.(*Animation).clone(), nil
	}

	var animation *Animation
	switch strings.ToLower(filepath.Ext(animationPath)) {
	case ".dc6":
		dc6, err := loadDC6(animationPath, palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = createAnimationFromDC6(dc6)
		if err != nil {
			return nil, err
		}
	case ".dcc":
		dcc, err := loadDCC(animationPath)
		if err != nil {
			return nil, err
		}

		palette, err := loadPalette(palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = createAnimationFromDCC(dcc, palette, transparency)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("unknown animation format")
	}

	if err := sm.cache.insert(cachePath, animation.clone(), 1); err != nil {
		return nil, err
	}

	return animation, nil
}
