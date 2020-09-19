package d2player

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
)

// NewPlayerStateFactory creates a new PlayerStateFactory and initializes it.
func NewPlayerStateFactory(asset *d2asset.AssetManager) (*PlayerStateFactory, error) {
	inventoryItemFactory, err := d2inventory.NewInventoryItemFactory(asset)
	if err != nil {
		return nil, err
	}

	factory := &PlayerStateFactory{
		asset:                asset,
		InventoryItemFactory: inventoryItemFactory,
	}

	return factory, nil
}

// PlayerStateFactory is responsible for creating player state objects
type PlayerStateFactory struct {
	asset *d2asset.AssetManager
	*d2hero.HeroStateFactory
	*d2inventory.InventoryItemFactory
}

// CreatePlayerState creates a PlayerState instance and returns a pointer to it
func (f *PlayerStateFactory) CreatePlayerState(
	heroName string,
	hero d2enum.Hero,
	classStats *d2datadict.CharStatsRecord,
) *PlayerState {
	result := &PlayerState{
		HeroName:  heroName,
		HeroType:  hero,
		Act:       1,
		Stats:     f.CreateHeroStatsState(hero, classStats),
		Skills:    f.CreateHeroSkillsState(classStats),
		Equipment: f.DefaultHeroItems[hero],
		FilePath:  "",
	}

	if err := result.Save(); err != nil {
		fmt.Printf("failed to save game state!, err: %v\n", err)
		return nil
	}

	return result
}
