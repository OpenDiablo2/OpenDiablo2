package d2game

import (
	"log"
	"runtime"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2scene"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

var singleton struct {
	loadingSprite *d2ui.Sprite // The sprite shown when loading stuff
	lastTime      float64      // Last time we updated the scene
	showFPS       bool
	timeScale     float64
}

func Initialize(loadingSpr *d2ui.Sprite) error {
	singleton.loadingSprite = loadingSpr
	singleton.timeScale = 1.0
	singleton.lastTime = d2common.Now()

	d2term.BindAction("fullscreen", "toggles fullscreen", func() {
		fullscreen, err := d2render.IsFullScreen()
		if err == nil {
			fullscreen = !fullscreen
			d2render.SetFullScreen(fullscreen)
			d2term.OutputInfo("fullscreen is now: %v", fullscreen)
		} else {
			d2term.OutputError(err.Error())
		}
	})
	d2term.BindAction("vsync", "toggles vsync", func() {
		vsync, err := d2render.GetVSyncEnabled()
		if err == nil {
			vsync = !vsync
			d2render.SetVSyncEnabled(vsync)
			d2term.OutputInfo("vsync is now: %v", vsync)
		} else {
			d2term.OutputError(err.Error())
		}
	})
	d2term.BindAction("fps", "toggle fps counter", func() {
		singleton.showFPS = !singleton.showFPS
		d2term.OutputInfo("fps counter is now: %v", singleton.showFPS)
	})
	d2term.BindAction("timescale", "set scalar for elapsed time", func(timeScale float64) {
		if timeScale <= 0 {
			d2term.OutputError("invalid time scale value")
		} else {
			d2term.OutputInfo("timescale changed from %f to %f", singleton.timeScale, timeScale)
		}
	})

	return nil
}

func Run(gitBranch string) error {
	if err := d2render.Run(update, 800, 600, "OpenDiablo 2 ("+gitBranch+")"); err != nil {
		log.Fatal(err)
	}
	return nil
}

// Advance updates the internal state of the engine
func Advance() {
	d2scene.UpdateScene()
	if d2scene.GetCurrentScene() == nil {
		log.Fatal("no scene loaded")
	}

	if d2scene.IsLoading() {
		return
	}

	currentTime := d2common.Now()
	deltaTime := (currentTime - singleton.lastTime) * singleton.timeScale
	singleton.lastTime = currentTime

	d2scene.Advance(deltaTime)
	d2ui.Advance(deltaTime)
	d2term.Advance(deltaTime)
	d2input.Advance(deltaTime)
}

// Draw draws the game
func render(target d2render.Surface) {
	if d2scene.GetLoadingProgress() < 1.0 {
		singleton.loadingSprite.SetCurrentFrame(
			int(d2common.Max(0, d2common.Min(
				uint32(singleton.loadingSprite.GetFrameCount()-1),
				uint32(float64(singleton.loadingSprite.GetFrameCount()-1)*d2scene.GetLoadingProgress()),
			))),
		)
		singleton.loadingSprite.Render(target)
	} else {
		if d2scene.GetCurrentScene() == nil {
			log.Fatal("no scene loaded")
		}
		d2scene.Render(target)
		d2ui.Render(target)
	}
	if singleton.showFPS {
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

func update(screen d2render.Surface) error {
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
