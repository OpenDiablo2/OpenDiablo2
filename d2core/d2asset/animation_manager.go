package d2asset

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	animationBudget = 64
)

// Static checks to confirm struct conforms to interface
var _ d2interface.ArchivedAnimationManager = &animationManager{}
var _ d2interface.Cacher = &animationManager{}

type animationManager struct {
	cache    d2interface.Cache
	renderer d2interface.Renderer
}

func (am *animationManager) ClearCache() {
	am.cache.Clear()
}

func (am *animationManager) GetCache() d2interface.Cache {
	return am.cache
}

func createAnimationManager(renderer d2interface.Renderer) *animationManager {
	return &animationManager{
		renderer: renderer,
		cache:    d2common.CreateCache(animationBudget),
	}
}

func (am *animationManager) LoadAnimation(
	animationPath, palettePath string,
	effect d2enum.DrawEffect) (d2interface.Animation, error) {
	cachePath := fmt.Sprintf("%s;%s;%d", animationPath, palettePath, effect)
	if animation, found := am.cache.Retrieve(cachePath); found {
		return animation.(d2interface.Animation).Clone(), nil
	}

	var animation d2interface.Animation

	ext := strings.ToLower(filepath.Ext(animationPath))
	switch ext {
	case ".dc6":
		palette, err := LoadPalette(palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = CreateDC6Animation(am.renderer, animationPath, palette, d2enum.DrawEffectNone)
		if err != nil {
			return nil, err
		}
	case ".dcc":
		palette, err := LoadPalette(palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = CreateDCCAnimation(am.renderer, animationPath, palette, effect)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown animation format: %s", ext)
	}

	if err := am.cache.Insert(cachePath, animation.Clone(), 1); err != nil {
		return nil, err
	}

	return animation, nil
}
