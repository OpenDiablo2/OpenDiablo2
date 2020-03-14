package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/keyboard"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/mouse"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

const (
	// CursorButtonLeft represents the left mouse button
	CursorButtonLeft = mouse.ButtonLeft
	// CursorButtonRight represents the right mouse button
	CursorButtonRight = mouse.ButtonRight
)

var widgets []Widget
var cursorButtons mouse.MouseButton
var pressedIndex int
var CursorX int
var CursorY int
var clickSfx d2audio.SoundEffect
var keyState map[keyboard.Key]bool
var mouseState map[mouse.MouseButton]bool

type uiInstance struct{}

func (ui *uiInstance) OnMouseButtonDown(event d2input.MouseEvent) bool {
	mouseState[event.Button] = true
	if event.Button == mouse.ButtonLeft {
		handleLeftMousePress()
		return true
	}
	if event.Button == mouse.ButtonRight {
		handleRightMousePress()
		return true
	}
	return false
}

func (ui *uiInstance) OnMouseButtonUp(event d2input.MouseEvent) bool {
	mouseState[event.Button] = false
	if event.Button == mouse.ButtonLeft {
		handleLeftMouseRelease()
		return true
	}
	return false
}

func (ui *uiInstance) OnMouseMove(event d2input.MouseMoveEvent) bool {
	CursorX = event.X
	CursorY = event.Y
	return true
}

func Initialize() error {
	keyState = make(map[keyboard.Key]bool)
	mouseState = make(map[mouse.MouseButton]bool)

	ui := &uiInstance{}
	if err := d2input.BindHandler(ui); err != nil {
		return err
	}

	pressedIndex = -1
	clickSfx, _ = d2audio.LoadSoundEffect(d2resource.SFXButtonClick)
	return nil
}

// Reset resets the state of the UI manager. Typically called for new scenes
func Reset() {
	widgets = make([]Widget, 0)
	pressedIndex = -1
}

// AddWidget adds a widget to the UI manager
func AddWidget(widget Widget) {
	widgets = append(widgets, widget)
}

// Render renders all of the UI elements
func Render(target d2render.Surface) {
	for _, widget := range widgets {
		if widget.GetVisible() {
			widget.Render(target)
		}
	}
}

// Update updates all of the UI elements
func Advance(elapsed float64) {
	for _, widget := range widgets {
		if widget.GetVisible() {
			widget.Advance(elapsed)
		}
	}

}

func handleLeftMousePress() {
	found := false
	for i, widget := range widgets {
		if !widget.GetVisible() || !widget.GetEnabled() {
			continue
		}
		wx, wy := widget.GetPosition()
		ww, wh := widget.GetSize()
		if CursorX >= wx && CursorX <= wx+ww && CursorY >= wy && CursorY <= wy+wh {
			widget.SetPressed(true)
			if pressedIndex == -1 {
				found = true
				pressedIndex = i
				clickSfx.Play()
			} else if pressedIndex > -1 && pressedIndex != i {
				widgets[i].SetPressed(false)
			} else {
				found = true
			}
		} else {
			widget.SetPressed(false)
		}
	}
	if !found {
		if pressedIndex > -1 {
			widgets[pressedIndex].SetPressed(false)
		} else {
			pressedIndex = -2
		}
	}
}

func handleLeftMouseRelease() {
	if pressedIndex > -1 {
		widget := widgets[pressedIndex]
		wx, wy := widget.GetPosition()
		ww, wh := widget.GetSize()
		if CursorX >= wx && CursorX <= wx+ww && CursorY >= wy && CursorY <= wy+wh {
			widget.Activate()
		}
	} else {
		for _, widget := range widgets {
			if !widget.GetVisible() || !widget.GetEnabled() {
				continue
			}
			widget.SetPressed(false)
		}
	}
	pressedIndex = -1
}

func handleRightMousePress() {
	cursorButtons |= CursorButtonRight
}

func CursorButtonPressed(button mouse.MouseButton) bool {
	if val, found := mouseState[button]; found {
		return val
	}
	mouseState[button] = false
	return false
}

func KeyPressed(key keyboard.Key) bool {
	if val, found := keyState[key]; found {
		return val
	}
	keyState[key] = false
	return false
}
