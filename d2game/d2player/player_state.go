package d2player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
)

// PlayerState stores the state of the player
type PlayerState struct {
	HeroName  string                         `json:"heroName"`
	HeroType  d2enum.Hero                    `json:"heroType"`
	HeroLevel int                            `json:"heroLevel"`
	Act       int                            `json:"act"`
	FilePath  string                         `json:"-"`
	Equipment d2inventory.CharacterEquipment `json:"equipment"`
	Stats     *d2hero.HeroStatsState         `json:"stats"`
	Skills    *d2hero.HeroSkillsState        `json:"skills"`
	X         float64                        `json:"x"`
	Y         float64                        `json:"y"`
}

// HasGameStates returns true if the player has any previously saved game
func HasGameStates() bool {
	basePath, _ := getGameBaseSavePath()
	files, _ := ioutil.ReadDir(basePath)

	return len(files) > 0
}

// GetAllPlayerStates returns all player saves
func GetAllPlayerStates() []*PlayerState {
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
			continue
		} else if gameState.Stats == nil || gameState.Skills == nil {
			// temporarily loading default class stats if the character was created before saving stats/skills was introduced
			// to be removed in the future
			classStats := d2datadict.CharStats[gameState.HeroType]
			gameState.Stats = d2hero.CreateHeroStatsState(gameState.HeroType, classStats)
			gameState.Skills = d2hero.CreateHeroSkillsState(classStats)

			if err := gameState.Save(); err != nil {
				fmt.Printf("failed to save game state!, err: %v\n", err)
			}
		}
		result = append(result, gameState)

	}

	return result
}

// CreateTestGameState is used for the map engine previewer
func CreateTestGameState() *PlayerState {
	result := &PlayerState{}
	return result
}

// LoadPlayerState loads the player state from the file
func LoadPlayerState(filePath string) *PlayerState {
	strData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}

	result := &PlayerState{
		FilePath: filePath,
	}

	err = json.Unmarshal(strData, result)
	if err != nil {
		return nil
	}

	return result
}

// CreatePlayerState creates a PlayerState instance and returns a pointer to it
func CreatePlayerState(heroName string, hero d2enum.Hero, classStats *d2datadict.CharStatsRecord) *PlayerState {
	result := &PlayerState{
		HeroName:  heroName,
		HeroType:  hero,
		Act:       1,
		Stats:     d2hero.CreateHeroStatsState(hero, classStats),
		Skills:    d2hero.CreateHeroSkillsState(classStats),
		Equipment: d2inventory.HeroObjects[hero],
		FilePath:  "",
	}

	if err := result.Save(); err != nil {
		fmt.Printf("failed to save game state!, err: %v\n", err)
		return nil
	}

	return result
}

func getGameBaseSavePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(configDir, "OpenDiablo2/Saves"), nil
}

func getFirstFreeFileName() string {
	i := 0
	basePath, _ := getGameBaseSavePath()

	for {
		filePath := path.Join(basePath, strconv.Itoa(i)+".od2")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return filePath
		}
		i++
	}
}

// Save saves the player state to a file
func (v *PlayerState) Save() error {
	if v.FilePath == "" {
		v.FilePath = getFirstFreeFileName()
	}
	if err := os.MkdirAll(path.Dir(v.FilePath), 0755); err != nil {
		return err
	}

	fileJSON, _ := json.MarshalIndent(v, "", "   ")
	if err := ioutil.WriteFile(v.FilePath, fileJSON, 0644); err != nil {
		return err
	}

	return nil
}
