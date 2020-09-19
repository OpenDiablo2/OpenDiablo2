package d2hero

import "github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

func NewHeroStateFactory(asset *d2asset.AssetManager) (*HeroStateFactory, error) {
	factory := &HeroStateFactory{asset: asset}

	return factory, nil
}

type HeroStateFactory struct {
	asset *d2asset.AssetManager
}
