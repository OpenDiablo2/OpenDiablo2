package d2player

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
)

// NewPlayerStateFactory creates a new PlayerStateFactory and initializes it.
func NewPlayerStateFactory(asset *d2asset.AssetManager) (*PlayerStateFactory, error) {
	factory := &PlayerStateFactory{
		asset: asset,
	}

	return factory, nil
}

// PlayerStateFactory is responsible for creating player state objects
type PlayerStateFactory struct {
	asset *d2asset.AssetManager
}

// CreatePlayerState creates a PlayerState instance and returns a pointer to it
func (f *PlayerStateFactory) CreatePlayerState(
	heroName string,
	hero d2enum.Hero,
	classStats *d2datadict.CharStatsRecord,
) *PlayerState {
	result := &PlayerState{
		HeroName: heroName,
		HeroType: hero,
		Act:      1,
		Stats:    d2hero.CreateHeroStatsState(hero, classStats),
		FilePath: "",
	}

	if err := result.Save(); err != nil {
		fmt.Printf("failed to save game state!, err: %v\n", err)
		return nil
	}

	return result
}
