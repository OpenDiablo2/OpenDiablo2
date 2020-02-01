package d2game

import (
	"log"
	"runtime"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2scenemanager"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

var loadingSprite *d2render.Sprite // The sprite shown when loading stuff
var lastTime float64               // Last time we updated the scene
var showFPS bool
var timeScale float64

type bsForInputHanding struct {
}

var bsHandler *bsForInputHanding

func Initialize(loadingSpr *d2render.Sprite) error {
	bsHandler = &bsForInputHanding{}
	loadingSprite = loadingSpr
	timeScale = 1.0
	lastTime = d2helper.Now()
	d2input.BindHandler(bsHandler)

	return nil
}

func Run(gitBranch string) error {
	if err := d2render.Run(update, 800, 600, "OpenDiablo 2 ("+gitBranch+")"); err != nil {
		log.Fatal(err)
	}
	return nil
}

func SetTimeScale(scale float64) {
	timeScale = scale
}

func GetTimeScale() float64 {
	return timeScale
}

func (bs *bsForInputHanding) OnKeyDown(event d2input.KeyEvent) bool {
	if event.Key == d2input.KeyEnter && event.KeyMod == d2input.KeyModAlt {
		isFullScreen, _ := d2render.IsFullScreen()
		d2render.SetFullScreen(!isFullScreen)
		return true
	}

	if event.Key == d2input.KeyF6 {
		showFPS = !showFPS
		return true
	}

	if event.Key == d2input.KeyF8 {
		enabled, _ := d2render.GetVSyncEnabled()
		d2render.SetVSyncEnabled(!enabled)
		return true
	}

	return false
}

// Advance updates the internal state of the engine
func Advance() {
	d2scenemanager.UpdateScene()
	if d2scenemanager.GetCurrentScene() == nil {
		log.Fatal("no scene loaded")
	}

	if d2scenemanager.IsLoading() {
		return
	}

	currentTime := d2helper.Now()
	deltaTime := (currentTime - lastTime) * timeScale
	lastTime = currentTime

	d2scenemanager.Advance(deltaTime)
	d2ui.Advance(deltaTime)
	d2term.Advance(deltaTime)
	d2input.Advance(deltaTime)
}

// Draw draws the game
func render(target d2common.Surface) {
	if d2scenemanager.GetLoadingProgress() < 1.0 {
		loadingSprite.SetCurrentFrame(int(d2helper.Max(0,
			d2helper.Min(uint32(loadingSprite.GetFrameCount()-1),
				uint32(float64(loadingSprite.GetFrameCount()-1)*d2scenemanager.GetLoadingProgress())))))
		loadingSprite.Render(target)
	} else {
		if d2scenemanager.GetCurrentScene() == nil {
			log.Fatal("no scene loaded")
		}
		d2scenemanager.Render(target)
		d2ui.Render(target)
	}
	if showFPS {
		target.PushTranslation(5, 565)
		vsyncEnabled, _ := d2render.GetVSyncEnabled()
		fps, _ := d2render.CurrentFPS()
		target.DrawText("vsync:" + strconv.FormatBool(vsyncEnabled) + "\nFPS:" + strconv.Itoa(int(fps)))
		target.Pop()

		cx, cy, _ := d2render.GetCursorPos()

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

func update(screen d2common.Surface) error {
	Advance()
	err, drawingSkipped := d2render.IsDrawingSkipped()
	if err != nil {
		return err
	}
	if !drawingSkipped {
		_, surface := d2render.CreateSurface(screen)
		render(surface)
		if surface.GetDepth() > 0 {
			panic("detected surface stack leak")
		}
	}

	return nil
}
