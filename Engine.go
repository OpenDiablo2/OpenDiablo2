package OpenDiablo2

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"sync"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/essial/OpenDiablo2/UI"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
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

// CursorButton represents a mouse button
type CursorButton uint8

const (
	// CursorButtonLeft represents the left mouse button
	CursorButtonLeft CursorButton = 1
	// CursorButtonRight represents the right mouse button
	CursorButtonRight CursorButton = 2
)

// Engine is the core OpenDiablo2 engine
type Engine struct {
	Settings        EngineConfig                        // Engine configuration settings from json file
	Files           map[string]string                   // Map that defines which files are in which MPQs
	Palettes        map[Palettes.Palette]Common.Palette // Color palettes
	SoundEntries    map[string]SoundEntry               // Sound configurations
	LoadingSprite   *Common.Sprite                      // The sprite shown when loading stuff
	CursorX         int                                 // X position of the cursor
	CursorY         int                                 // Y position of the cursor
	CursorButtons   CursorButton                        // The buttons that are currently being pressed
	LoadingProgress float64                             // LoadingProcess is a range between 0.0 and 1.0. If set, loading screen displays.
	CurrentScene    Common.SceneInterface               // The current scene being rendered
	UIManager       *UI.Manager                         // The UI manager
	nextScene       Common.SceneInterface               // The next scene to be loaded at the end of the game loop
	fontCache       map[string]*MPQFont                 // The font cash
	audioContext    *audio.Context                      // The Audio context
	bgmAudio        *audio.Player                       // The audio player
	fullscreenKey   bool                                // When true, the fullscreen toggle is still being pressed
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
	result.UIManager = UI.CreateManager(result)
	audioContext, err := audio.NewContext(22050)
	if err != nil {
		log.Fatal(err)
	}
	result.audioContext = audioContext
	result.LoadingSprite = result.LoadSprite(ResourcePaths.LoadingScreen, Palettes.Loading)
	loadingSpriteSizeX, loadingSpriteSizeY := result.LoadingSprite.GetSize()
	result.LoadingSprite.MoveTo(int(400-(loadingSpriteSizeX/2)), int(300+(loadingSpriteSizeY/2)))
	result.SetNextScene(CreateMainMenu(result))
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
		mpq, err := LoadMPQ(mpqPath)
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

// GetFile loads a file from the specified mpq and returns the data as a byte array
func (v *Engine) GetFile(fileName string) []byte {
	// TODO: May want to cache some things if performance becomes an issue
	mpqFile := v.Files[strings.ToLower(fileName)]
	mpq, err := LoadMPQ(mpqFile)
	if err != nil {
		log.Fatal(err)
	}
	fileName = strings.ReplaceAll(fileName, `/`, `\`)[1:]
	blockTableEntry, err := mpq.getFileBlockData(fileName)
	if err != nil {
		log.Fatal(err)
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
	v.Palettes = make(map[Palettes.Palette]Common.Palette)
	log.Println("loading palettes")
	for file := range v.Files {
		if strings.Index(file, "/data/global/palette/") != 0 || strings.Index(file, ".dat") != len(file)-4 {
			continue
		}
		nameParts := strings.Split(file, `/`)
		paletteName := Palettes.Palette(nameParts[len(nameParts)-2])
		palette := Common.CreatePalette(paletteName, v.GetFile(file))
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
func (v *Engine) LoadSprite(fileName string, palette Palettes.Palette) *Common.Sprite {
	data := v.GetFile(fileName)
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
	v.CurrentScene.Load()
}

// CursorButtonPressed determines if the specified button has been pressed
func (v *Engine) CursorButtonPressed(button CursorButton) bool {
	return v.CursorButtons&button > 0
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
	v.CursorButtons = 0
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		v.CursorButtons |= CursorButtonLeft
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		v.CursorButtons |= CursorButtonRight
	}
	v.CurrentScene.Update()
	v.UIManager.Update()
}

// Draw draws the game
func (v *Engine) Draw(screen *ebiten.Image) {
	v.CursorX, v.CursorY = ebiten.CursorPosition()
	if v.LoadingProgress < 1.0 {
		v.LoadingSprite.Frame = uint8(Common.Max(0, Common.Min(uint32(len(v.LoadingSprite.Frames)-1), uint32(float64(len(v.LoadingSprite.Frames)-1)*v.LoadingProgress))))
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
func (v *Engine) SetNextScene(nextScene Common.SceneInterface) {
	v.nextScene = nextScene
}

// GetFont creates or loads an existing font
func (v *Engine) GetFont(font string, palette Palettes.Palette) *MPQFont {
	cacheItem, exists := v.fontCache[font+"_"+string(palette)]
	if exists {
		return cacheItem
	}
	newFont := CreateMPQFont(v, font, palette)
	v.fontCache[font+"_"+string(palette)] = newFont
	return newFont
}

// PlayBGM plays an infinitely looping background track
func (v *Engine) PlayBGM(song string) {
	go func() {
		if v.bgmAudio != nil {
			v.bgmAudio.Close()
		}
		audioData := v.GetFile(song)
		d, err := wav.Decode(v.audioContext, audio.BytesReadSeekCloser(audioData))
		if err != nil {
			log.Fatal(err)
		}
		s := audio.NewInfiniteLoop(d, int64(len(audioData)))

		v.bgmAudio, err = audio.NewPlayer(v.audioContext, s)
		if err != nil {
			log.Fatal(err)
		}
		// Play the infinite-length stream. This never ends.
		v.bgmAudio.Rewind()
		v.bgmAudio.Play()
	}()
}
