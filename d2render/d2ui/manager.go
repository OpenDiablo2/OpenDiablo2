package d2ui

import (
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
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

// Manager represents the UI manager
type Manager struct {
	widgets            []Widget
	cursorSprite       *d2render.Sprite
	cursorButtons      CursorButton
	pressedIndex       int
	CursorX            int
	CursorY            int
	clickSfx           *d2audio.SoundEffect
	waitForLeftMouseUp bool
}

// CreateManager creates a new instance of a UI manager
func CreateManager(soundManager d2audio.Manager) *Manager {
	cursorSprite, _ := d2render.LoadSprite(d2resource.CursorDefault, d2resource.PaletteUnits)
	result := &Manager{
		pressedIndex:       -1,
		cursorSprite:       cursorSprite,
		clickSfx:           soundManager.LoadSoundEffect(d2resource.SFXButtonClick),
		waitForLeftMouseUp: false,
	}
	return result
}

// Reset resets the state of the UI manager. Typically called for new scenes
func (v *Manager) Reset() {
	v.widgets = make([]Widget, 0)
	v.pressedIndex = -1
	v.waitForLeftMouseUp = true
}

// AddWidget adds a widget to the UI manager
func (v *Manager) AddWidget(widget Widget) {
	v.widgets = append(v.widgets, widget)
}

func (v *Manager) WaitForMouseRelease() {
	v.waitForLeftMouseUp = true
}

// Render renders all of the UI elements
func (v *Manager) Render(screen *ebiten.Image) {
	for _, widget := range v.widgets {
		if !widget.GetVisible() {
			continue
		}
		widget.Render(screen)
	}

	cx, cy := ebiten.CursorPosition()
	v.cursorSprite.SetPosition(cx, cy)
	v.cursorSprite.Render(screen)
}

// Update updates all of the UI elements
func (v *Manager) Update() {
	v.cursorButtons = 0
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !v.waitForLeftMouseUp {
			v.cursorButtons |= CursorButtonLeft
		}
	} else {
		if v.waitForLeftMouseUp {
			v.waitForLeftMouseUp = false
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		v.cursorButtons |= CursorButtonRight
	}
	v.CursorX, v.CursorY = ebiten.CursorPosition()
	if v.CursorButtonPressed(CursorButtonLeft) {
		found := false
		for i, widget := range v.widgets {
			if !widget.GetVisible() || !widget.GetEnabled() {
				continue
			}
			wx, wy := widget.GetPosition()
			ww, wh := widget.GetSize()
			if v.CursorX >= wx && v.CursorX <= wx+int(ww) && v.CursorY >= wy && v.CursorY <= wy+int(wh) {
				widget.SetPressed(true)
				if v.pressedIndex == -1 {
					found = true
					v.pressedIndex = i
					v.clickSfx.Play()
				} else if v.pressedIndex > -1 && v.pressedIndex != i {
					v.widgets[i].SetPressed(false)
				} else {
					found = true
				}
			} else {
				widget.SetPressed(false)
			}
		}
		if !found {
			if v.pressedIndex > -1 {
				v.widgets[v.pressedIndex].SetPressed(false)
			} else {
				v.pressedIndex = -2
			}
		}
	} else {
		if v.pressedIndex > -1 {
			widget := v.widgets[v.pressedIndex]
			wx, wy := widget.GetPosition()
			ww, wh := widget.GetSize()
			if v.CursorX >= wx && v.CursorX <= wx+int(ww) && v.CursorY >= wy && v.CursorY <= wy+int(wh) {
				widget.Activate()
			}
		} else {
			for _, widget := range v.widgets {
				if !widget.GetVisible() || !widget.GetEnabled() {
					continue
				}
				widget.SetPressed(false)
			}
		}
		v.pressedIndex = -1
	}
}

// CursorButtonPressed determines if the specified button has been pressed
func (v *Manager) CursorButtonPressed(button CursorButton) bool {
	return v.cursorButtons&button > 0
}

func (v *Manager) KeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}
