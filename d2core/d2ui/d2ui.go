package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/hajimehoshi/ebiten"
)

// CursorButton represents a mouse button
type CursorButton uint8

const (
	// CursorButtonLeft represents the left mouse button
	CursorButtonLeft CursorButton = 1
	// CursorButtonRight represents the right mouse button
	CursorButtonRight CursorButton = 2
)

var widgets []Widget
var cursorSprite *Sprite
var cursorButtons CursorButton
var pressedIndex int
var CursorX int
var CursorY int
var clickSfx d2audio.SoundEffect
var waitForLeftMouseUp bool

func Initialize(curSprite *Sprite) {
	cursorSprite = curSprite
	pressedIndex = -1
	clickSfx, _ = d2audio.LoadSoundEffect(d2resource.SFXButtonClick)
	waitForLeftMouseUp = false
}

// Reset resets the state of the UI manager. Typically called for new scenes
func Reset() {
	widgets = make([]Widget, 0)
	pressedIndex = -1
	waitForLeftMouseUp = true
}

// AddWidget adds a widget to the UI manager
func AddWidget(widget Widget) {
	widgets = append(widgets, widget)
}

func WaitForMouseRelease() {
	waitForLeftMouseUp = true
}

// Render renders all of the UI elements
func Render(target d2render.Surface) {
	for _, widget := range widgets {
		if widget.GetVisible() {
			widget.Render(target)
		}
	}

	cx, cy := ebiten.CursorPosition()
	cursorSprite.SetPosition(cx, cy)
	cursorSprite.Render(target)
}

// Update updates all of the UI elements
func Advance(elapsed float64) {
	for _, widget := range widgets {
		if widget.GetVisible() {
			widget.Advance(elapsed)
		}
	}

	cursorButtons = 0
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !waitForLeftMouseUp {
			cursorButtons |= CursorButtonLeft
		}
	} else {
		if waitForLeftMouseUp {
			waitForLeftMouseUp = false
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		cursorButtons |= CursorButtonRight
	}
	CursorX, CursorY = ebiten.CursorPosition()
	if CursorButtonPressed(CursorButtonLeft) {
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
	} else {
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
}

// CursorButtonPressed determines if the specified button has been pressed
func CursorButtonPressed(button CursorButton) bool {
	return cursorButtons&button > 0
}

func KeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

func GetCursorSprite() *Sprite {
	return cursorSprite
}
