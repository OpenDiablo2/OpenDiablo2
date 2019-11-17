package d2core

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/OpenDiablo2/D2Shared/d2common"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
)

/*
	File Spec
	--------------------------------------------
	UINT32 GameState Version
	INT64  Game Seed
    BYTE   Hero Type
	BYTE   Hero Level
    BYTE   Act
    BYTE   Hero Name Length
    BYTE[] Hero Name
	--------------------------------------------
*/

type GameState struct {
	Seed      int64
	HeroName  string
	HeroType  d2enum.Hero
	HeroLevel int
	Act       int
	FilePath  string
	Equipment CharacterEquipment
}

const GameStateVersion = uint32(2) // Update this when you make breaking changes

func HasGameStates() bool {
	files, _ := ioutil.ReadDir(getGameBaseSavePath())
	return len(files) > 0
}

func GetAllGameStates() []*GameState {
	// TODO: Make this not crash tf out on bad files
	basePath := getGameBaseSavePath()
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
	result := &GameState{
		FilePath: path,
	}
	f, err := os.Open(path)
	if err != nil {
		log.Panicf(err.Error())
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Panicf(err.Error())
	}
	defer f.Close()
	sr := d2common.CreateStreamReader(bytes)
	if sr.GetUInt32() != GameStateVersion {
		// Unknown game version
		return nil
	}
	result.Seed = sr.GetInt64()
	result.HeroType = d2enum.Hero(sr.GetByte())
	result.HeroLevel = int(sr.GetByte())
	result.Act = int(sr.GetByte())
	heroNameLen := sr.GetByte()
	heroName, _ := sr.ReadBytes(int(heroNameLen))
	result.HeroName = string(heroName)
	return result
}

func CreateGameState(heroName string, hero d2enum.Hero, hardcore bool) *GameState {
	result := &GameState{
		HeroName: heroName,
		HeroType: hero,
		Act:      1,
		Seed:     time.Now().UnixNano(),
		FilePath: "",
	}
	result.Save()
	return result
}

func getGameBaseSavePath() string {
	if runtime.GOOS == "windows" {
		appDataPath := os.Getenv("APPDATA")
		basePath := path.Join(appDataPath, "OpenDiablo2", "Saves")
		if err := os.MkdirAll(basePath, os.ModeDir); err != nil {
			log.Panicf(err.Error())
		}
		return basePath
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicf(err.Error())
	}
	basePath := path.Join(homeDir, ".OpenDiablo2", "Saves")
	if err := os.MkdirAll(basePath, 0755); err != nil {
		log.Panicf(err.Error())
	}
	// TODO: Is mac supposed to have a super special place for the save games?
	return basePath
}

func getFirstFreeFileName() string {
	i := 0
	basePath := getGameBaseSavePath()
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
	f, err := os.Create(v.FilePath)
	if err != nil {
		log.Panicf(err.Error())
	}
	defer f.Close()
	sr := d2common.CreateStreamWriter()
	sr.PushUint32(GameStateVersion)
	sr.PushInt64(v.Seed)
	sr.PushByte(byte(v.HeroType))
	sr.PushByte(byte(v.HeroLevel))
	sr.PushByte(byte(v.Act))
	sr.PushByte(byte(len(v.HeroName)))
	for _, ch := range v.HeroName {
		sr.PushByte(byte(ch))
	}
	if _, err := f.Write(sr.GetBytes()); err != nil {
		log.Panicf(err.Error())
	}
}
