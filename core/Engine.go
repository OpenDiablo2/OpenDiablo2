package core

import (
	"log"
	"math"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/mpq"

	"github.com/OpenDiablo2/OpenDiablo2/palettedefs"

	"github.com/OpenDiablo2/OpenDiablo2/sound"

	"github.com/OpenDiablo2/OpenDiablo2/common"
	"github.com/OpenDiablo2/OpenDiablo2/resourcepaths"
	"github.com/OpenDiablo2/OpenDiablo2/scenes"
	"github.com/OpenDiablo2/OpenDiablo2/ui"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Engine is the core OpenDiablo2 engine
type Engine struct {
	Settings        *common.Configuration // Engine configuration settings from json file
	Files           map[string]string     // Map that defines which files are in which MPQs
	CheckedPatch    map[string]bool       // First time we check a file, we'll check if it's in the patch. This notes that we've already checked that.
	LoadingSprite   *common.Sprite        // The sprite shown when loading stuff
	loadingProgress float64               // LoadingProcess is a range between 0.0 and 1.0. If set, loading screen displays.
	loadingIndex    int                   // Determines which load function is currently being called
	thingsToLoad    []func()              // The load functions for the next scene
	stepLoadingSize float64               // The size for each loading step
	CurrentScene    scenes.Scene          // The current scene being rendered
	UIManager       *ui.Manager           // The UI manager
	SoundManager    *sound.Manager        // The sound manager
	nextScene       scenes.Scene          // The next scene to be loaded at the end of the game loop
	fullscreenKey   bool                  // When true, the fullscreen toggle is still being pressed
	lastTime        float64               // Last time we updated the scene
	showFPS         bool
}

// CreateEngine creates and instance of the OpenDiablo2 engine
func CreateEngine() *Engine {
	result := &Engine{
		CurrentScene: nil,
		nextScene:    nil,
	}
	result.loadConfigurationFile()
	resourcepaths.LanguageCode = result.Settings.Language
	result.mapMpqFiles()
	common.LoadPalettes(result.Files, result)
	common.LoadTextDictionary(result)
	common.LoadLevelTypes(result)
	common.LoadLevelPresets(result)
	common.LoadLevelWarps(result)
	common.LoadObjectTypes(result)
	common.LoadObjects(result)
	common.LoadWeapons(result)
	common.LoadArmors(result)
	common.LoadUniqueItems(result)
	common.LoadMissiles(result)
	common.LoadSounds(result)
	common.LoadObjectLookups()
	common.LoadAnimationData(result)
	common.LoadMonStats(result)
	result.SoundManager = sound.CreateManager(result)
	result.SoundManager.SetVolumes(result.Settings.BgmVolume, result.Settings.SfxVolume)
	result.UIManager = ui.CreateManager(result, *result.SoundManager)
	result.LoadingSprite = result.LoadSprite(resourcepaths.LoadingScreen, palettedefs.Loading)
	loadingSpriteSizeX, loadingSpriteSizeY := result.LoadingSprite.GetSize()
	result.LoadingSprite.MoveTo(int(400-(loadingSpriteSizeX/2)), int(300+(loadingSpriteSizeY/2)))
	result.SetNextScene(scenes.CreateMainMenu(result, result, result.UIManager, result.SoundManager))
	//result.SetNextScene(Scenes.CreateBlizzardIntro(result, result))
	return result
}

func (v *Engine) loadConfigurationFile() {
	log.Println("Loading configuration file")
	v.Settings = common.LoadConfiguration()
}

func (v *Engine) mapMpqFiles() {
	v.Files = make(map[string]string)
}

var mutex sync.Mutex

func (v *Engine) LoadFile(fileName string) []byte {
	fileName = strings.ReplaceAll(fileName, "{LANG}", resourcepaths.LanguageCode)
	fileName = strings.ToLower(fileName)
	fileName = strings.ReplaceAll(fileName, `/`, "\\")
	if fileName[0] == '\\' {
		fileName = fileName[1:]
	}
	mutex.Lock()
	defer mutex.Unlock()
	// TODO: May want to cache some things if performance becomes an issue
	cachedMpqFile, cacheExists := v.Files[fileName]
	if cacheExists {
		archive, _ := mpq.Load(cachedMpqFile)
		result, _ := archive.ReadFile(fileName)
		return result
	}
	for _, mpqFile := range v.Settings.MpqLoadOrder {
		archive, _ := mpq.Load(path.Join(v.Settings.MpqPath, mpqFile))
		if !archive.FileExists(fileName) {
			continue
		}
		result, _ := archive.ReadFile(fileName)
		if len(result) == 0 {
			continue
		}
		v.Files[fileName] = path.Join(v.Settings.MpqPath, mpqFile)
		return result
	}
	log.Fatalf("Could not load %s from MPQs", fileName)
	return []byte{}
}

// IsLoading returns true if the engine is currently in a loading state
func (v *Engine) IsLoading() bool {
	return v.loadingProgress < 1.0
}

// LoadSprite loads a sprite from the game's data files
func (v *Engine) LoadSprite(fileName string, palette palettedefs.PaletteType) *common.Sprite {
	data := v.LoadFile(fileName)
	sprite := common.CreateSprite(data, common.Palettes[palette])
	return sprite
}

// updateScene handles the scene maintenance for the engine
func (v *Engine) updateScene() {
	if v.nextScene == nil {
		if v.thingsToLoad != nil {
			if v.loadingIndex < len(v.thingsToLoad) {
				v.thingsToLoad[v.loadingIndex]()
				v.loadingIndex++
				if v.loadingIndex < len(v.thingsToLoad) {
					v.StepLoading()
				} else {
					v.FinishLoading()
					v.thingsToLoad = nil
				}
				return
			}
		}
		return
	}
	if v.CurrentScene != nil {
		v.CurrentScene.Unload()
		runtime.GC()
	}
	v.CurrentScene = v.nextScene
	v.nextScene = nil
	v.UIManager.Reset()
	v.thingsToLoad = v.CurrentScene.Load()
	v.loadingIndex = 0
	v.SetLoadingStepSize(1.0 / float64(len(v.thingsToLoad)))
	v.ResetLoading()
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

	if inpututil.IsKeyJustPressed(ebiten.KeyF6) {
		v.showFPS = !v.showFPS
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF8) {
		ebiten.SetVsyncEnabled(!ebiten.IsVsyncEnabled())
	}

	v.updateScene()
	if v.CurrentScene == nil {
		log.Fatal("no scene loaded")
	}

	if v.IsLoading() {
		return
	}

	currentTime := float64(time.Now().UnixNano()) / float64(time.Second)

	deltaTime := math.Min((currentTime - v.lastTime), 0.1)
	v.lastTime = currentTime

	v.CurrentScene.Update(deltaTime)
	v.UIManager.Update()
}

// Draw draws the game
func (v *Engine) Draw(screen *ebiten.Image) {
	if v.loadingProgress < 1.0 {
		v.LoadingSprite.Frame = uint8(common.Max(0, common.Min(uint32(len(v.LoadingSprite.Frames)-1), uint32(float64(len(v.LoadingSprite.Frames)-1)*v.loadingProgress))))
		v.LoadingSprite.Draw(screen)
	} else {
		if v.CurrentScene == nil {
			log.Fatal("no scene loaded")
		}
		v.CurrentScene.Render(screen)
		v.UIManager.Draw(screen)
	}
	if v.showFPS == true {
		ebitenutil.DebugPrintAt(screen, "vsync:"+strconv.FormatBool(ebiten.IsVsyncEnabled())+"\nFPS:"+strconv.Itoa(int(ebiten.CurrentFPS())), 5, 565)
	}
}

// SetNextScene tells the engine what scene to load on the next update cycle
func (v *Engine) SetNextScene(nextScene scenes.Scene) {
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
