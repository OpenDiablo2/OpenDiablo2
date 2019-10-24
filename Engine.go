package OpenDiablo2

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"sync"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/ResourcePaths"

	"github.com/hajimehoshi/ebiten"
)

// EngineConfig defines the configuration for the engine, loaded from config.json
type EngineConfig struct {
	FullScreen      bool
	Scale           float64
	RunInBackground bool
	TicksPerSecond  int
	VsyncEnabled    bool
	MpqPath         string
	MpqLoadOrder    []string
}

// Engine is the core OpenDiablo2 engine
type Engine struct {
	Settings        EngineConfig          // Engine configuration settings from json file
	Files           map[string]string     // Map that defines which files are in which MPQs
	Palettes        map[string]Palette    // Color palettes
	SoundEntries    map[string]SoundEntry // Sound configurations
	CursorSprite    Sprite                // The sprite shown for cursors
	LoadingSprite   Sprite                // The sprite shown when loading stuff
	CursorX         int                   // X position of the cursor
	CursorY         int                   // Y position of the cursor
	LoadingProgress float64               // LoadingProcess is a range between 0.0 and 1.0. If set, loading screen displays.
	CurrentScene    Common.SceneInterface // The current scene being rendered
	nextScene       Common.SceneInterface // The next scene to be loaded at the end of the game loop
	fontCache       map[string]*MPQFont   // The font cash
}

// CreateEngine creates and instance of the OpenDiablo2 engine
func CreateEngine() *Engine {
	result := &Engine{
		LoadingProgress: float64(0.0),
		CurrentScene:    nil,
		nextScene:       nil,
		fontCache:       make(map[string]*MPQFont),
	}
	result.loadConfigurationFile()
	result.mapMpqFiles()
	result.loadPalettes()
	result.loadSoundEntries()
	result.CursorSprite = result.LoadSprite(ResourcePaths.CursorDefault, result.Palettes["units"])
	result.LoadingSprite = result.LoadSprite(ResourcePaths.LoadingScreen, result.Palettes["loading"])
	loadingSpriteSizeX, loadingSpriteSizeY := result.LoadingSprite.GetSize()
	result.LoadingSprite.MoveTo(int(400-(loadingSpriteSizeX/2)), int(300+(loadingSpriteSizeY/2)))
	result.SetNextScene(CreateMainMenu(result))
	return result
}

func (v *Engine) loadConfigurationFile() {
	log.Println("loading configuration file")
	configJSON, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config EngineConfig

	json.Unmarshal(configJSON, &config)
	v.Settings = config
}

func (v *Engine) mapMpqFiles() {
	log.Println("mapping mpq file structure")
	v.Files = make(map[string]string)
	lock := sync.RWMutex{}
	for _, mpqFileName := range v.Settings.MpqLoadOrder {
		mpqPath := path.Join(v.Settings.MpqPath, mpqFileName)
		mpq, err := LoadMPQ(mpqPath)
		if err != nil {
			panic(err)
		}
		fileListText, err := mpq.ReadFile("(listfile)")
		if err != nil {
			panic(err)
		}
		fileList := strings.Split(string(fileListText), "\r\n")
		for _, filePath := range fileList {
			if _, exists := v.Files[strings.ToLower(filePath)]; exists {
				lock.RUnlock()
				continue
			}
			v.Files[`/`+strings.ReplaceAll(strings.ToLower(filePath), `\`, `/`)] = mpqPath
		}
	}
}

// GetFile loads a file from the specified mpq and returns the data as a byte array
func (v *Engine) GetFile(fileName string) []byte {
	// TODO: May want to cache some things if performance becomes an issue
	mpqFile := v.Files[strings.ToLower(fileName)]
	mpq, err := LoadMPQ(mpqFile)
	if err != nil {
		panic(err)
	}
	blockTableEntry, err := mpq.getFileBlockData(strings.ReplaceAll(fileName, `/`, `\`)[1:])
	if err != nil {
		panic(err)
	}
	mpqStream := CreateMPQStream(mpq, blockTableEntry, fileName)
	result := make([]byte, blockTableEntry.UncompressedFileSize)
	mpqStream.Read(result, 0, blockTableEntry.UncompressedFileSize)

	return result
}

// IsLoading returns true if the engine is currently in a loading state
func (v *Engine) IsLoading() bool {
	return v.LoadingProgress < 1.0
}

func (v *Engine) loadPalettes() {
	v.Palettes = make(map[string]Palette)
	log.Println("loading palettes")
	for file := range v.Files {
		if strings.Index(file, "/data/global/palette/") != 0 || strings.Index(file, ".dat") != len(file)-4 {
			continue
		}
		nameParts := strings.Split(file, `/`)
		paletteName := nameParts[len(nameParts)-2]
		palette := CreatePalette(paletteName, v.GetFile(file))
		v.Palettes[paletteName] = palette
	}
}

func (v *Engine) loadSoundEntries() {
	log.Println("loading sound configurations")
	v.SoundEntries = make(map[string]SoundEntry)
	soundData := strings.Split(string(v.GetFile(ResourcePaths.SoundSettings)), "\r\n")[1:]
	for _, line := range soundData {
		if len(line) == 0 {
			continue
		}
		soundEntry := CreateSoundEntry(line)
		v.SoundEntries[soundEntry.Handle] = soundEntry
	}
}

// LoadSprite loads a sprite from the game's data files
func (v *Engine) LoadSprite(fileName string, palette Palette) Sprite {
	data := v.GetFile(fileName)
	sprite := CreateSprite(data, palette)
	return sprite
}

// updateScene handles the scene maintenance for the engine
func (v *Engine) updateScene() {
	if v.nextScene == nil {
		return
	}
	if v.CurrentScene != nil {
		v.CurrentScene.Unload()
	}
	v.CurrentScene = v.nextScene
	v.nextScene = nil
	v.CurrentScene.Load()
}

// Update updates the internal state of the engine
func (v *Engine) Update() {
	v.updateScene()
	if v.CurrentScene == nil {
		panic("no scene loaded")
	}
	v.CurrentScene.Update()
}

// Draw draws the game
func (v *Engine) Draw(screen *ebiten.Image) {
	v.CursorX, v.CursorY = ebiten.CursorPosition()
	if v.LoadingProgress < 1.0 {
		v.LoadingSprite.Frame = uint8(Max(0, Min(uint32(len(v.LoadingSprite.Frames)-1), uint32(float64(len(v.LoadingSprite.Frames)-1)*v.LoadingProgress))))
		v.LoadingSprite.Draw(screen)
	} else {
		if v.CurrentScene == nil {
			panic("no scene loaded")
		}
		v.CurrentScene.Render(screen)
	}

	v.CursorSprite.MoveTo(v.CursorX, v.CursorY)
	v.CursorSprite.Draw(screen)
}

// SetNextScene tells the engine what scene to load on the next update cycle
func (v *Engine) SetNextScene(nextScene Common.SceneInterface) {
	v.nextScene = nextScene
}

// GetFont creates or loads an existing font
func (v *Engine) GetFont(font, palette string) *MPQFont {
	cacheItem, exists := v.fontCache[font+"_"+palette]
	if exists {
		return cacheItem
	}
	newFont := CreateMPQFont(v, font, v.Palettes[palette])
	v.fontCache[font+"_"+palette] = newFont
	return newFont
}
