package OpenDiablo2

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"path"
	"strings"
	"sync"

	"github.com/essial/OpenDiablo2/Sound"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/MPQ"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/essial/OpenDiablo2/Scenes"
	"github.com/essial/OpenDiablo2/UI"

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
	Settings        EngineConfig                        // Engine configuration settings from json file
	Files           map[string]string                   // Map that defines which files are in which MPQs
	Palettes        map[Palettes.Palette]Common.Palette // Color palettes
	SoundEntries    map[string]SoundEntry               // Sound configurations
	LoadingSprite   *Common.Sprite                      // The sprite shown when loading stuff
	loadingProgress float64                             // LoadingProcess is a range between 0.0 and 1.0. If set, loading screen displays.
	stepLoadingSize float64                             // The size for each loading step
	CurrentScene    Scenes.Scene                        // The current scene being rendered
	UIManager       *UI.Manager                         // The UI manager
	SoundManager    *Sound.Manager                      // The sound manager
	nextScene       Scenes.Scene                        // The next scene to be loaded at the end of the game loop
	fullscreenKey   bool                                // When true, the fullscreen toggle is still being pressed
}

// CreateEngine creates and instance of the OpenDiablo2 engine
func CreateEngine() *Engine {
	result := &Engine{
		CurrentScene: nil,
		nextScene:    nil,
	}
	result.loadConfigurationFile()
	result.mapMpqFiles()
	result.loadPalettes()
	result.loadSoundEntries()
	result.SoundManager = Sound.CreateManager(result)
	result.UIManager = UI.CreateManager(result)
	result.LoadingSprite = result.LoadSprite(ResourcePaths.LoadingScreen, Palettes.Loading)
	loadingSpriteSizeX, loadingSpriteSizeY := result.LoadingSprite.GetSize()
	result.LoadingSprite.MoveTo(int(400-(loadingSpriteSizeX/2)), int(300+(loadingSpriteSizeY/2)))
	result.SetNextScene(Scenes.CreateMainMenu(result, result, result.UIManager, result.SoundManager))
	return result
}

func (v *Engine) loadConfigurationFile() {
	log.Println("loading configuration file")
	configJSON, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
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
		mpq, err := MPQ.Load(mpqPath)
		if err != nil {
			log.Fatal(err)
		}
		fileListText, err := mpq.ReadFile("(listfile)")
		if err != nil {
			log.Fatal(err)
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

var mutex sync.Mutex

// LoadFile loads a file from the specified mpq and returns the data as a byte array
func (v *Engine) LoadFile(fileName string) []byte {
	mutex.Lock()
	// TODO: May want to cache some things if performance becomes an issue
	mpqFile := v.Files[strings.ToLower(fileName)]
	mpq, err := MPQ.Load(mpqFile)
	if err != nil {
		log.Fatal(err)
	}
	fileName = strings.ReplaceAll(fileName, `/`, `\`)[1:]
	blockTableEntry, err := mpq.GetFileBlockData(fileName)
	if err != nil {
		log.Fatal(err)
	}
	mpqStream := MPQ.CreateStream(mpq, blockTableEntry, fileName)
	result := make([]byte, blockTableEntry.UncompressedFileSize)
	mpqStream.Read(result, 0, blockTableEntry.UncompressedFileSize)
	mutex.Unlock()
	return result
}

// IsLoading returns true if the engine is currently in a loading state
func (v *Engine) IsLoading() bool {
	return v.loadingProgress < 1.0
}

func (v *Engine) loadPalettes() {
	v.Palettes = make(map[Palettes.Palette]Common.Palette)
	log.Println("loading palettes")
	for file := range v.Files {
		if strings.Index(file, "/data/global/palette/") != 0 || strings.Index(file, ".dat") != len(file)-4 {
			continue
		}
		nameParts := strings.Split(file, `/`)
		paletteName := Palettes.Palette(nameParts[len(nameParts)-2])
		palette := Common.CreatePalette(paletteName, v.LoadFile(file))
		v.Palettes[paletteName] = palette
	}
}

func (v *Engine) loadSoundEntries() {
	log.Println("loading sound configurations")
	v.SoundEntries = make(map[string]SoundEntry)
	soundData := strings.Split(string(v.LoadFile(ResourcePaths.SoundSettings)), "\r\n")[1:]
	for _, line := range soundData {
		if len(line) == 0 {
			continue
		}
		soundEntry := CreateSoundEntry(line)
		v.SoundEntries[soundEntry.Handle] = soundEntry
	}
}

// LoadSprite loads a sprite from the game's data files
func (v *Engine) LoadSprite(fileName string, palette Palettes.Palette) *Common.Sprite {
	data := v.LoadFile(fileName)
	sprite := Common.CreateSprite(data, v.Palettes[palette])
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
	v.UIManager.Reset()
	thingsToLoad := v.CurrentScene.Load()
	v.SetLoadingStepSize(1.0 / float64(len(thingsToLoad)))
	v.ResetLoading()
	go func() {
		for _, f := range thingsToLoad {
			f()
			v.StepLoading()
		}
		v.FinishLoading()
	}()
}

// Update updates the internal state of the engine
func (v *Engine) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyAlt) && ebiten.IsKeyPressed(ebiten.KeyEnter) {
		if !v.fullscreenKey {
			ebiten.SetFullscreen(!ebiten.IsFullscreen())
		}
		v.fullscreenKey = true
	} else {
		v.fullscreenKey = false
	}

	v.updateScene()
	if v.CurrentScene == nil {
		log.Fatal("no scene loaded")
	}

	if v.IsLoading() {
		return
	}

	v.CurrentScene.Update()
	v.UIManager.Update()
}

// Draw draws the game
func (v *Engine) Draw(screen *ebiten.Image) {
	if v.loadingProgress < 1.0 {
		v.LoadingSprite.Frame = uint8(Common.Max(0, Common.Min(uint32(len(v.LoadingSprite.Frames)-1), uint32(float64(len(v.LoadingSprite.Frames)-1)*v.loadingProgress))))
		v.LoadingSprite.Draw(screen)
	} else {
		if v.CurrentScene == nil {
			log.Fatal("no scene loaded")
		}
		v.CurrentScene.Render(screen)
		v.UIManager.Draw(screen)
	}
}

// SetNextScene tells the engine what scene to load on the next update cycle
func (v *Engine) SetNextScene(nextScene Scenes.Scene) {
	v.nextScene = nextScene
}

// SetLoadingStepSize sets the size of the loading step
func (v *Engine) SetLoadingStepSize(size float64) {
	v.stepLoadingSize = size
}

// ResetLoading resets the loading progress
func (v *Engine) ResetLoading() {
	v.loadingProgress = 0.0
}

// StepLoading increments the loading progress
func (v *Engine) StepLoading() {
	v.loadingProgress = math.Min(1.0, v.loadingProgress+v.stepLoadingSize)
}

// FinishLoading terminates the loading phase
func (v *Engine) FinishLoading() {
	v.loadingProgress = 1.0
}
