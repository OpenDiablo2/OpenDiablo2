package d2player

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

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

	heroState, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	factory := &PlayerStateFactory{
		asset:                asset,
		InventoryItemFactory: inventoryItemFactory,
		HeroStateFactory:     heroState,
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
	classStats *d2records.CharStatsRecord,
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

// GetAllPlayerStates returns all player saves
func (f *PlayerStateFactory) GetAllPlayerStates() []*PlayerState {
	basePath, _ := getGameBaseSavePath()
	files, _ := ioutil.ReadDir(basePath)
	result := make([]*PlayerState, 0)

	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() || len(fileName) < 5 || !strings.EqualFold(fileName[len(fileName)-4:], ".od2") {
			continue
		}

		gameState := LoadPlayerState(path.Join(basePath, file.Name()))
		if gameState == nil || gameState.HeroType == d2enum.HeroNone {

		} else if gameState.Stats == nil || gameState.Skills == nil {
			// temporarily loading default class stats if the character was created before saving stats/skills was introduced
			// to be removed in the future
			classStats := f.asset.Records.Character.Stats[gameState.HeroType]
			gameState.Stats = f.CreateHeroStatsState(gameState.HeroType, classStats)
			gameState.Skills = f.CreateHeroSkillsState(classStats)
			if err := gameState.Save(); err != nil {
				fmt.Printf("failed to save game state!, err: %v\n", err)
			}
		}

		result = append(result, gameState)
	}

	return result
}
