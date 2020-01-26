package d2core

import (
	"log"
	"math"
	"runtime"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon/d2coreinterface"
	"github.com/OpenDiablo2/OpenDiablo2/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2term"

	"github.com/OpenDiablo2/OpenDiablo2/d2render"

	"github.com/OpenDiablo2/OpenDiablo2/d2data"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"

	"github.com/OpenDiablo2/OpenDiablo2/d2audio"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"

	"github.com/hajimehoshi/ebiten"
)

// Engine is the core OpenDiablo2 engine
type Engine struct {
	Settings        *d2corecommon.Configuration // Engine configuration settings from json file
	CheckedPatch    map[string]bool             // First time we check a file, we'll check if it's in the patch. This notes that we've already checked that.
	LoadingSprite   *d2render.Sprite            // The sprite shown when loading stuff
	loadingProgress float64                     // LoadingProcess is a range between 0.0 and 1.0. If set, loading screen displays.
	loadingIndex    int                         // Determines which load function is currently being called
	thingsToLoad    []func()                    // The load functions for the next scene
	stepLoadingSize float64                     // The size for each loading step
	CurrentScene    d2coreinterface.Scene       // The current scene being rendered
	UIManager       *d2ui.Manager               // The UI manager
	SoundManager    *d2audio.Manager            // The sound manager
	nextScene       d2coreinterface.Scene       // The next scene to be loaded at the end of the game loop
	lastTime        float64                     // Last time we updated the scene
	showFPS         bool
	timeScale       float64
}

// CreateEngine creates and instance of the OpenDiablo2 engine
func CreateEngine() *Engine {
	result := &Engine{
		lastTime:  d2helper.Now(),
		timeScale: 1.0,
	}

	result.Settings = d2corecommon.LoadConfiguration()
	if err := result.Settings.Save(); err != nil {
		log.Printf("could not load settings: %v", err)
	}

	d2asset.Initialize(result.Settings)
	d2resource.LanguageCode = result.Settings.Language
	d2datadict.LoadPalettes(nil, result)
	d2common.LoadTextDictionary(result)
	d2datadict.LoadLevelTypes(result)
	d2datadict.LoadLevelPresets(result)
	d2datadict.LoadLevelWarps(result)
	d2datadict.LoadObjectTypes(result)
	d2datadict.LoadObjects(result)
	d2datadict.LoadWeapons(result)
	d2datadict.LoadArmors(result)
	d2datadict.LoadMiscItems(result)
	d2datadict.LoadUniqueItems(result)
	d2datadict.LoadMissiles(result)
	d2datadict.LoadSounds(result)
	d2data.LoadAnimationData(result)
	d2datadict.LoadMonStats(result)
	LoadHeroObjects()
	result.SoundManager = d2audio.CreateManager()
	result.SoundManager.SetVolumes(result.Settings.BgmVolume, result.Settings.SfxVolume)
	result.UIManager = d2ui.CreateManager(*result.SoundManager)
	result.LoadingSprite, _ = d2render.LoadSprite(d2resource.LoadingScreen, d2resource.PaletteLoading)
	loadingSpriteSizeX, loadingSpriteSizeY := result.LoadingSprite.GetCurrentFrameSize()
	result.LoadingSprite.SetPosition(int(400-(loadingSpriteSizeX/2)), int(300+(loadingSpriteSizeY/2)))

	d2term.BindAction("timescale", "set scalar for elapsed time", func(scale float64) {
		if scale <= 0 {
			d2term.OutputError("invalid time scale value")
		} else {
			d2term.OutputInfo("timescale changed from %f to %f", result.timeScale, scale)
			result.timeScale = scale
		}
	})

	d2input.BindHandler(result)

	return result
}

func (v *Engine) LoadFile(fileName string) []byte {
	data, _ := d2asset.LoadFile(fileName)
	return data
}

// IsLoading returns true if the engine is currently in a loading state
func (v Engine) IsLoading() bool {
	return v.loadingProgress < 1.0
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

func (e *Engine) OnKeyDown(event d2input.KeyEvent) bool {
	if event.Key == d2input.KeyEnter && event.KeyMod == d2input.KeyModAlt {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
		return true
	}

	if event.Key == d2input.KeyF6 {
		e.showFPS = !e.showFPS
		return true
	}

	if event.Key == d2input.KeyF8 {
		ebiten.SetVsyncEnabled(!ebiten.IsVsyncEnabled())
		return true
	}

	return false
}

// Advance updates the internal state of the engine
func (v *Engine) Advance() {
	v.updateScene()
	if v.CurrentScene == nil {
		log.Fatal("no scene loaded")
	}

	if v.IsLoading() {
		return
	}

	currentTime := d2helper.Now()
	deltaTime := (currentTime - v.lastTime) * v.timeScale
	v.lastTime = currentTime

	v.CurrentScene.Advance(deltaTime)
	v.UIManager.Advance(deltaTime)

	d2term.Advance(deltaTime)
	d2input.Advance(deltaTime)
}

// Draw draws the game
func (v Engine) Render(target *d2surface.Surface) {
	if v.loadingProgress < 1.0 {
		v.LoadingSprite.SetCurrentFrame(int(d2helper.Max(0, d2helper.Min(uint32(v.LoadingSprite.GetFrameCount()-1), uint32(float64(v.LoadingSprite.GetFrameCount()-1)*v.loadingProgress)))))
		v.LoadingSprite.Render(target)
	} else {
		if v.CurrentScene == nil {
			log.Fatal("no scene loaded")
		}
		v.CurrentScene.Render(target)
		v.UIManager.Render(target)
	}
	if v.showFPS {
		target.PushTranslation(5, 565)
		target.DrawText("vsync:" + strconv.FormatBool(ebiten.IsVsyncEnabled()) + "\nFPS:" + strconv.Itoa(int(ebiten.CurrentFPS())))
		target.Pop()

		cx, cy := ebiten.CursorPosition()

		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		target.PushTranslation(680, 0)
		target.DrawText("Alloc   " + strconv.FormatInt(int64(m.Alloc)/1024/1024, 10))
		target.PushTranslation(0, 16)
		target.DrawText("Pause   " + strconv.FormatInt(int64(m.PauseTotalNs/1024/1024), 10))
		target.PushTranslation(0, 16)
		target.DrawText("HeapSys " + strconv.FormatInt(int64(m.HeapSys/1024/1024), 10))
		target.PushTranslation(0, 16)
		target.DrawText("NumGC   " + strconv.FormatInt(int64(m.NumGC), 10))
		target.PushTranslation(0, 16)
		target.DrawText("Coords  " + strconv.FormatInt(int64(cx), 10) + "," + strconv.FormatInt(int64(cy), 10))
		target.PopN(5)
	}

	d2term.Render(target)
}

// SetNextScene tells the engine what scene to load on the next update cycle
func (v *Engine) SetNextScene(nextScene d2coreinterface.Scene) {
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
