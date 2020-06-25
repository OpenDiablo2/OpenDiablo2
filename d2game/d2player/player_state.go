package d2player

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
)

type PlayerState struct {
	HeroName  string                         `json:"heroName"`
	HeroType  d2enum.Hero                    `json:"heroType"`
	HeroLevel int                            `json:"heroLevel"`
	Act       int                            `json:"act"`
	FilePath  string                         `json:"-"`
	Equipment d2inventory.CharacterEquipment `json:"equipment"`
	Stats     *d2hero.HeroStatsState          `json:"stats"`
	X         float64                        `json:"x"`
	Y         float64                        `json:"y"`
}

func HasGameStates() bool {
	basePath, _ := getGameBaseSavePath()
	files, _ := ioutil.ReadDir(basePath)
	return len(files) > 0
}

func GetAllPlayerStates() []*PlayerState {
	basePath, _ := getGameBaseSavePath()
	files, _ := ioutil.ReadDir(basePath)
	result := make([]*PlayerState, 0)
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() || len(fileName) < 5 || strings.ToLower(fileName[len(fileName)-4:]) != ".od2" {
			continue
		}
		gameState := LoadPlayerState(path.Join(basePath, file.Name()))
		if gameState == nil {
			continue
		// temporarily loading default class stats if the character was created before saving stats was introduced
		// to be removed in the future
		} else if gameState.Stats == nil {
			gameState.Stats = d2hero.CreateHeroStatsState(gameState.HeroType, *d2datadict.CharStats[gameState.HeroType], 1, 0)
			gameState.Save()
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

func LoadPlayerState(path string) *PlayerState {
	strData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	result := &PlayerState{
		FilePath: path,
	}
	err = json.Unmarshal(strData, result)
	if err != nil {
		return nil
	}
	return result
}

func CreatePlayerState(heroName string, hero d2enum.Hero, classStats d2datadict.CharStatsRecord, hardcore bool) *PlayerState {
	result := &PlayerState{
		HeroName:  heroName,
		HeroType:  hero,
		Act:       1,
		Stats: d2hero.CreateHeroStatsState(hero, classStats, 1, 0),
		Equipment: d2inventory.HeroObjects[hero],
		FilePath:  "",
	}

	result.Save()
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

func (v *PlayerState) Save() {
	if v.FilePath == "" {
		v.FilePath = getFirstFreeFileName()
	}
	if err := os.MkdirAll(path.Dir(v.FilePath), 0755); err != nil {
		log.Panic(err.Error())
	}
	fileJson, _ := json.MarshalIndent(v, "", "   ")
	ioutil.WriteFile(v.FilePath, fileJson, 0644)
}
