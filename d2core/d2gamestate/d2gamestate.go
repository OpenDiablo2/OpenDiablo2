package d2gamestate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type GameState struct {
	Seed      int64                          `json:"seed"` // TODO: Seed needs to be regenerated every time the game starts
	HeroName  string                         `json:"heroName"`
	HeroType  d2enum.Hero                    `json:"heroType"`
	HeroLevel int                            `json:"heroLevel"`
	Act       int                            `json:"act"`
	FilePath  string                         `json:"-"`
	Equipment d2inventory.CharacterEquipment `json:"equipment"`
}

const GameStateVersion = uint32(2) // Update this when you make breaking changes

func HasGameStates() bool {
	basePath, _ := getGameBaseSavePath()
	files, _ := ioutil.ReadDir(basePath)
	return len(files) > 0
}

func GetAllGameStates() []*GameState {
	basePath, _ := getGameBaseSavePath()
	files, _ := ioutil.ReadDir(basePath)
	result := make([]*GameState, 0)
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() || len(fileName) < 5 || strings.ToLower(fileName[len(fileName)-4:]) != ".od2" {
			continue
		}
		gameState := LoadGameState(path.Join(basePath, file.Name()))
		if gameState == nil {
			continue
		}
		result = append(result, gameState)
	}
	return result
}

// CreateTestGameState is used for the map engine previewer
func CreateTestGameState() *GameState {
	result := &GameState{
		Seed: time.Now().UnixNano(),
	}
	return result
}

func LoadGameState(path string) *GameState {
	strData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	result := &GameState{
		FilePath: path,
	}
	err = json.Unmarshal(strData, result)
	if err != nil {
		return nil
	}
	return result
}

func CreateGameState(heroName string, hero d2enum.Hero, hardcore bool) *GameState {
	result := &GameState{
		HeroName:  heroName,
		HeroType:  hero,
		Act:       1,
		Seed:      time.Now().UnixNano(),
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

func (v *GameState) Save() {
	if v.FilePath == "" {
		v.FilePath = getFirstFreeFileName()
	}
	if err := os.MkdirAll(path.Dir(v.FilePath), 0755); err != nil {
		log.Panic(err.Error())
	}
	fileJson, _ := json.MarshalIndent(v, "", "   ")
	ioutil.WriteFile(v.FilePath, fileJson, 0644)
}
