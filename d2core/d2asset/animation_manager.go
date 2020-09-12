package d2asset

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	animationBudget = 64
)

// Static checks to confirm struct conforms to interface
var _ d2interface.AnimationManager = &animationManager{}
var _ d2interface.Cacher = &animationManager{}

type animationManager struct {
	*AssetManager
	cache    d2interface.Cache
	renderer d2interface.Renderer
}

func (am *animationManager) ClearCache() {
	am.cache.Clear()
}

func (am *animationManager) GetCache() d2interface.Cache {
	return am.cache
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
		palette, err := am.LoadPalette(palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = am.CreateDC6Animation(animationPath, palette, d2enum.DrawEffectNone)
		if err != nil {
			return nil, err
		}
	case ".dcc":
		palette, err := am.LoadPalette(palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = am.CreateDCCAnimation(animationPath, palette, effect)
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

// CreateDC6Animation creates an Animation from d2dc6.DC6 and d2dat.DATPalette
func (am *animationManager) CreateDC6Animation(dc6Path string,
	palette d2interface.Palette, effect d2enum.DrawEffect) (d2interface.Animation, error) {
	dc6, err := am.loadDC6(dc6Path)
	if err != nil {
		return nil, err
	}

	anim := DC6Animation{
		animation: animation{
			directions:     make([]animationDirection, dc6.Directions),
			playLength:     defaultPlayLength,
			playLoop:       true,
			originAtBottom: true,
			effect:         effect,
		},
		dc6Path:  dc6Path,
		dc6:      dc6,
		palette:  palette,
		renderer: am.renderer,
	}

	err = anim.SetDirection(0)

	return &anim, err
}

// CreateDCCAnimation creates an animation from d2dcc.DCC and d2dat.DATPalette
func (am *animationManager) CreateDCCAnimation(dccPath string,
	palette d2interface.Palette,
	effect d2enum.DrawEffect) (d2interface.Animation, error) {
	dcc, err := am.loadDCC(dccPath)
	if err != nil {
		return nil, err
	}

	anim := animation{
		playLength: defaultPlayLength,
		playLoop:   true,
		directions: make([]animationDirection, dcc.NumberOfDirections),
		effect:     effect,
	}

	DCC := DCCAnimation{
		animation:        anim,
		animationManager: am,
		dccPath:          dccPath,
		palette:          palette,
		renderer:         am.renderer,
	}

	err = DCC.SetDirection(0)
	if err != nil {
		return nil, err
	}

	return &DCC, nil
}

func (am *animationManager) loadDC6(path string) (*d2dc6.DC6, error) {
	dc6Data, err := am.LoadFile(path)
	if err != nil {
		return nil, err
	}

	dc6, err := d2dc6.Load(dc6Data)
	if err != nil {
		return nil, err
	}

	return dc6, nil
}

func (am *animationManager) loadDCC(path string) (*d2dcc.DCC, error) {
	dccData, err := am.LoadFile(path)
	if err != nil {
		return nil, err
	}

	return d2dcc.Load(dccData)
}
