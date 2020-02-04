package d2gui

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

var (
	ErrWasInit = errors.New("gui system is already initialized")
	ErrNotInit = errors.New("gui system is not initialized")
)

var singleton *guiManager

func Initialize() error {
	assertNotInit()

	var err error
	if singleton, err = createGuiManager(); err != nil {
		return err
	}

	return nil
}

func Shutdown() {
	singleton = nil
}

func Render(target d2render.Surface) error {
	assertWasInit()
	return singleton.render(target)
}

func Advance(elapsed float64) error {
	assertWasInit()
	return singleton.advance(elapsed)
}

func Clear() {
	assertWasInit()
	singleton.clear()
}

func ShowLoadScreen(progress float64) {
	assertWasInit()
	singleton.showLoadScreen(progress)
}

func HideLoadScreen() {
	assertWasInit()
	singleton.hideLoadScreen()
}

func ShowCursor() {
	assertWasInit()
	singleton.showCursor()
}

func HideCursor() {
	assertWasInit()
	singleton.hideCursor()
}

func assertWasInit() {
	if singleton == nil {
		panic(ErrNotInit)
	}
}

func assertNotInit() {
	if singleton != nil {
		panic(ErrWasInit)
	}
}
