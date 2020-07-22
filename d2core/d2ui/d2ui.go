package d2ui

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

// CursorButton represents a mouse button
type CursorButton uint8

const (
	// CursorButtonLeft represents the left mouse button
	CursorButtonLeft CursorButton = 1
	// CursorButtonRight represents the right mouse button
	CursorButtonRight CursorButton = 2
)

type UI struct {
	inputManager  d2interface.InputManager
	widgets       []Widget
	cursorButtons CursorButton // TODO (carrelld) convert dependent code and remove
	CursorX       int          // TODO (carrelld) convert dependent code and remove
	CursorY       int          // TODO (carrelld) convert dependent code and remove
	pressedWidget Widget
	clickSfx      d2interface.SoundEffect
}

var singleton UI

func Initialize(inputManager d2interface.InputManager, audioProvider d2interface.AudioProvider) {
	sfx, err := audioProvider.LoadSoundEffect(d2resource.SFXButtonClick)
	if err != nil {
		log.Fatalf("failed to initialize ui: %v", err)
	}
	singleton.clickSfx = sfx

	singleton.inputManager = inputManager
	if err := singleton.inputManager.BindHandler(&singleton); err != nil {
		log.Fatalf("failed to initialize ui: %v", err)
	}
}

// Reset resets the state of the UI manager. Typically called for new screens
func Reset() {
	singleton.widgets = nil
	singleton.pressedWidget = nil
}

// AddWidget adds a widget to the UI manager
func AddWidget(widget Widget) {
	singleton.inputManager.BindHandler(widget)
	singleton.widgets = append(singleton.widgets, widget)
}

func (u *UI) OnMouseButtonUp(event d2interface.MouseEvent) bool {
	singleton.CursorX, singleton.CursorY = event.X(), event.Y()
	if event.Button() == d2enum.MouseButtonLeft {
		singleton.cursorButtons |= CursorButtonLeft
		// activate previously pressed widget if cursor is still hovering
		w := singleton.pressedWidget
		if w != nil && contains(w, singleton.CursorX, singleton.CursorY) && w.GetVisible() && w.GetEnabled() {
			w.Activate()
		}
		// unpress all widgets that are pressed
		for _, w := range singleton.widgets {
			w.SetPressed(false)
		}
	}
	return false
}

func (u *UI) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	singleton.CursorX, singleton.CursorY = event.X(), event.Y()
	if event.Button() == d2enum.MouseButtonLeft {
		// find and press a widget on screen
		singleton.pressedWidget = nil
		for _, w := range singleton.widgets {
			if contains(w, singleton.CursorX, singleton.CursorY) && w.GetVisible() && w.GetEnabled() {
				w.SetPressed(true)
				singleton.pressedWidget = w
				u.clickSfx.Play()
				break
			}
		}
	}
	if event.Button() == d2enum.MouseButtonRight {
		singleton.cursorButtons |= CursorButtonRight
	}
	return false
}

// Render renders all of the UI elements
func Render(target d2interface.Surface) {
	for _, widget := range singleton.widgets {
		if widget.GetVisible() {
			widget.Render(target)
		}
	}
}

// contains determines whether a given x,y coordinate lands within a Widget
func contains(w Widget, x, y int) bool {
	wx, wy := w.GetPosition()
	ww, wh := w.GetSize()
	return x >= wx && x <= wx+ww && y >= wy && y <= wy+wh
}

// Update updates all of the UI elements
func Advance(elapsed float64) {
	for _, widget := range singleton.widgets {
		if widget.GetVisible() {
			widget.Advance(elapsed)
		}
	}
}

// CursorButtonPressed determines if the specified button has been pressed
func CursorButtonPressed(button CursorButton) bool {
	return singleton.cursorButtons&button > 0
}

func CursorPosition() (x, y int) {
	return singleton.CursorX, singleton.CursorY
}
