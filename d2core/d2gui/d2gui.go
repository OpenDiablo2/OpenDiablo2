package d2gui

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

var (
	ErrWasInit = errors.New("gui system is already initialized")
	ErrNotInit = errors.New("gui system is not initialized")
)

var singleton *manager

func Initialize() error {
	verifyNotInit()

	var err error
	if singleton, err = createGuiManager(); err != nil {
		return err
	}

	return nil
}

func Render(target d2render.Surface) error {
	verifyWasInit()
	return singleton.render(target)
}

func Advance(elapsed float64) error {
	verifyWasInit()
	return singleton.advance(elapsed)
}

func AddLayout() *Layout {
	return singleton.addLayout()
}

func AddSprite(imagePath, palettePath string) *Sprite {
	return singleton.addSprite(imagePath, palettePath)
}

func AddLabel(text string, fontStyle FontStyle) *Label {
	return singleton.addLabel(text, fontStyle)
}

func Clear() {
	verifyWasInit()
	singleton.clear()
}

func ShowLoadScreen(progress float64) {
	verifyWasInit()
	singleton.showLoadScreen(progress)
}

func HideLoadScreen() {
	verifyWasInit()
	singleton.hideLoadScreen()
}

func ShowCursor() {
	verifyWasInit()
	singleton.showCursor()
}

func HideCursor() {
	verifyWasInit()
	singleton.hideCursor()
}

func verifyWasInit() {
	if singleton == nil {
		panic(ErrNotInit)
	}
}

func verifyNotInit() {
	if singleton != nil {
		panic(ErrWasInit)
	}
}
