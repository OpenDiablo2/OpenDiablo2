package d2gui

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
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

func Render(target d2interface.Surface) error {
	verifyWasInit()
	return singleton.render(target)
}

func Advance(elapsed float64) error {
	verifyWasInit()
	return singleton.advance(elapsed)
}

func CreateLayout(renderer d2interface.Renderer, positionType PositionType) *Layout {
	verifyWasInit()
	return createLayout(renderer, positionType)
}

func SetLayout(layout *Layout) {
	verifyWasInit()
	singleton.SetLayout(layout)
}

// ShowLoadScreen renders the loading progress screen. The provided progress argument defines the loading animation's state in the range `[0, 1]`, where `0` is initial frame and `1` is the final frame
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
