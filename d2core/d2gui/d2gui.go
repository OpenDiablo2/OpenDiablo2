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

func RenderL(l *Layout, target d2render.Surface) error {
	return l.render(target)
}

func AdvanceL(l *Layout, elapsed float64) error {
	return l.advance(elapsed)
}

func Render(target d2render.Surface) error {
	verifyWasInit()
	return singleton.render(target)
}

func Advance(elapsed float64) error {
	verifyWasInit()
	return singleton.advance(elapsed)
}

func CreateLayout(positionType PositionType) *Layout {
	verifyWasInit()
	return createLayout(positionType)
}

func SetLayout(layout *Layout) {
	verifyWasInit()
	singleton.SetLayout(layout)
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
