package d2ui

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
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
	widgets            []Widget
	cursorButtons      CursorButton
	pressedIndex       int
	CursorX            int
	CursorY            int
	clickSfx           d2audio.SoundEffect
	waitForLeftMouseUp bool
}

var singleton UI

func Initialize() {
	clickSfx, err := d2audio.LoadSoundEffect(d2resource.SFXButtonClick)
	if err != nil {
		log.Fatalf("failed to initialize ui: %v", err)
	}
	singleton = UI{}
	singleton.pressedIndex = -1
	singleton.clickSfx = clickSfx

	d2input.BindHandler(singleton)
}

// Reset resets the state of the UI manager. Typically called for new screens
func Reset() {
	singleton.widgets = make([]Widget, 0)
	singleton.pressedIndex = -1
	singleton.waitForLeftMouseUp = true
}

// AddWidget adds a widget to the UI manager
func AddWidget(widget Widget) {
	singleton.widgets = append(singleton.widgets, widget)
}

func (u *UI) OnMouseButtonDown(event d2input.MouseEvent) bool {
	if event.Button == d2input.MouseButtonLeft {
		if !singleton.waitForLeftMouseUp {
			singleton.cursorButtons |= CursorButtonLeft
		}
	} else {
		if singleton.waitForLeftMouseUp {
			singleton.waitForLeftMouseUp = false
		}
	}
	if event.Button == d2input.MouseButtonRight {
		singleton.cursorButtons |= CursorButtonRight
	}
	singleton.CursorX, singleton.CursorY = event.X, event.Y
	if CursorButtonPressed(CursorButtonLeft) {
		found := false
		for i, widget := range singleton.widgets {
			if !widget.GetVisible() || !widget.GetEnabled() {
				continue
			}
			wx, wy := widget.GetPosition()
			ww, wh := widget.GetSize()
			if singleton.CursorX >= wx && singleton.CursorX <= wx+ww && singleton.CursorY >= wy && singleton.CursorY <= wy+wh {
				widget.SetPressed(true)
				if singleton.pressedIndex == -1 {
					found = true
					singleton.pressedIndex = i
					singleton.clickSfx.Play()
				} else if singleton.pressedIndex > -1 && singleton.pressedIndex != i {
					singleton.widgets[i].SetPressed(false)
				} else {
					found = true
				}
			} else {
				widget.SetPressed(false)
			}
		}
		if !found {
			if singleton.pressedIndex > -1 {
				singleton.widgets[singleton.pressedIndex].SetPressed(false)
			} else {
				singleton.pressedIndex = -2
			}
		}
	} else {
		if singleton.pressedIndex > -1 {
			widget := singleton.widgets[singleton.pressedIndex]
			wx, wy := widget.GetPosition()
			ww, wh := widget.GetSize()
			if singleton.CursorX >= wx && singleton.CursorX <= wx+ww && singleton.CursorY >= wy && singleton.CursorY <= wy+wh {
				widget.Activate()
			}
			return true
		} else {
			for _, widget := range singleton.widgets {
				if !widget.GetVisible() || !widget.GetEnabled() {
					continue
				}
				widget.SetPressed(false)
			}
		}
		singleton.pressedIndex = -1
	}
	return false
}

func WaitForMouseRelease() {
	singleton.waitForLeftMouseUp = true
}

// Render renders all of the UI elements
func Render(target d2render.Surface) {
	for _, widget := range singleton.widgets {
		if widget.GetVisible() {
			widget.Render(target)
		}
	}
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
