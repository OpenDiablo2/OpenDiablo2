package d2gui

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var (
	errWasInit = errors.New("gui system is already initialized")
	errNotInit = errors.New("gui system is not initialized")
)

var singleton *manager

// Initialize creates a singleton gui manager
func Initialize(inputManager d2interface.InputManager) error {
	verifyNotInit()

	var err error
	if singleton, err = createGuiManager(inputManager); err != nil {
		return err
	}

	return nil
}

// Render all of the gui elements
func Render(target d2interface.Surface) error {
	verifyWasInit()
	return singleton.render(target)
}

// Advance all of the gui elements
func Advance(elapsed float64) error {
	verifyWasInit()
	return singleton.advance(elapsed)
}

// CreateLayout creates a dynamic layout
func CreateLayout(renderer d2interface.Renderer, positionType PositionType) *Layout {
	verifyWasInit()
	return createLayout(renderer, positionType)
}

// SetLayout sets the gui manager's layout
func SetLayout(layout *Layout) {
	verifyWasInit()
	singleton.SetLayout(layout)
}

// ShowLoadScreen renders the loading progress screen.
// The provided progress argument defines the loading animation's state in the range `[0, 1]`,
// where `0` is initial frame and `1` is the final frame
func ShowLoadScreen(progress float64) {
	verifyWasInit()
	singleton.showLoadScreen(progress)
}

// HideLoadScreen hides the loading screen
func HideLoadScreen() {
	verifyWasInit()
	singleton.hideLoadScreen()
}

// ShowCursor shows the in-game mouse cursor
func ShowCursor() {
	verifyWasInit()
	singleton.showCursor()
}

// HideCursor hides the in-game mouse cursor
func HideCursor() {
	verifyWasInit()
	singleton.hideCursor()
}

func verifyWasInit() {
	if singleton == nil {
		panic(errNotInit)
	}
}

func verifyNotInit() {
	if singleton != nil {
		panic(errWasInit)
	}
}
