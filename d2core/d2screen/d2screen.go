package d2screen

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Screen interface{}

type ScreenLoadHandler interface {
	// OnLoad performs all necessary loading to prepare a screen to be shown such as loading assets, placing and binding
	// of ui elements, etc. This loading is done asynchronously. The provided channel will allow implementations to
	// provide progress via Error, Progress, or Done
	OnLoad(loading LoadingState)
}

type ScreenUnloadHandler interface {
	OnUnload() error
}

type ScreenRenderHandler interface {
	Render(target d2render.Surface) error
}

type ScreenAdvanceHandler interface {
	Advance(elapsed float64) error
}

var singleton struct {
	nextScreen    Screen
	loadingScreen Screen
	loadingState  LoadingState
	currentScreen Screen
}

func SetNextScreen(screen Screen) {
	singleton.nextScreen = screen
}

func Advance(elapsed float64) error {
	switch {
	case singleton.loadingScreen != nil:
		// this call blocks execution and could lead to deadlock if a screen implements OnLoad incorreclty
		load, ok := <-singleton.loadingState.updates
		if !ok {
			log.Println("loadingState chan should not be closed while in a loading screen")
		}
		if load.err != nil {
			log.Printf("PROBLEM LOADING THE SCREEN: %v", load.err)
			return load.err
		}
		d2gui.ShowLoadScreen(load.progress)
		if load.done {
			singleton.currentScreen = singleton.loadingScreen
			singleton.loadingScreen = nil
			d2gui.ShowCursor()
			d2gui.HideLoadScreen()
		}
	case singleton.nextScreen != nil:
		if handler, ok := singleton.currentScreen.(ScreenUnloadHandler); ok {
			if err := handler.OnUnload(); err != nil {
				return err
			}
		}

		d2ui.Reset()
		d2gui.SetLayout(nil)

		if handler, ok := singleton.nextScreen.(ScreenLoadHandler); ok {
			d2gui.ShowLoadScreen(0)
			d2gui.HideCursor()
			singleton.loadingState = LoadingState{updates: make(chan loadingUpdate)}
			go func() {
				handler.OnLoad(singleton.loadingState)
				singleton.loadingState.Done()
			}()
			singleton.currentScreen = nil
			singleton.loadingScreen = singleton.nextScreen
		} else {
			singleton.currentScreen = singleton.nextScreen
			singleton.loadingScreen = nil
		}
		singleton.nextScreen = nil
	case singleton.currentScreen != nil:
		if handler, ok := singleton.currentScreen.(ScreenAdvanceHandler); ok {
			if err := handler.Advance(elapsed); err != nil {
				return err
			}
		}
	}

	return nil
}

func Render(surface d2render.Surface) error {
	if handler, ok := singleton.currentScreen.(ScreenRenderHandler); ok {
		if err := handler.Render(surface); err != nil {
			return err
		}
	}

	return nil
}

type LoadingState struct {
	updates chan loadingUpdate
}

type loadingUpdate struct {
	progress float64
	err      error
	done     bool
}

// Error provides a way for callers to report an error during loading.
// This is meant to be delivered via the progress channel in OnLoad implementations.
func (l *LoadingState) Error(err error) {
	l.updates <- loadingUpdate{err: err}
}

// Progress provides a way for callers to report the ratio between `0` and `1` of the progress made loading a screen.
// This is meant to be delivered via the progress channel in OnLoad implementations.
func (l *LoadingState) Progress(ratio float64) {
	l.updates <- loadingUpdate{progress: ratio}
}

// Done provides a way for callers to report that screen loading has been completed.
// This is meant to be delivered via the progress channel in OnLoad implementations.
func (l *LoadingState) Done() {
	l.updates <- loadingUpdate{progress: 1.0}
	l.updates <- loadingUpdate{done: true}
}
