package UI

import (
	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/ResourcePaths"
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
	widgets       []Widget
	cursorSprite  *Common.Sprite
	cursorButtons CursorButton
	pressedIndex  int
	CursorX       int
	CursorY       int
}

// CreateManager creates a new instance of a UI manager
func CreateManager(provider Common.FileProvider) *Manager {
	result := &Manager{
		pressedIndex: -1,
		widgets:      make([]Widget, 0),
		cursorSprite: provider.LoadSprite(ResourcePaths.CursorDefault, Palettes.Units),
	}
	return result
}

// Reset resets the state of the UI manager. Typically called for new scenes
func (v *Manager) Reset() {
	v.widgets = make([]Widget, 0)
	v.pressedIndex = -1
}

// AddWidget adds a widget to the UI manager
func (v *Manager) AddWidget(widget Widget) {
	v.widgets = append(v.widgets, widget)
}

// Draw renders all of the UI elements
func (v *Manager) Draw(screen *ebiten.Image) {
	for _, widget := range v.widgets {
		if !widget.GetVisible() {
			continue
		}
		widget.Draw(screen)
	}

	cx, cy := ebiten.CursorPosition()
	v.cursorSprite.MoveTo(cx, cy)
	v.cursorSprite.Draw(screen)
}

// Update updates all of the UI elements
func (v *Manager) Update() {
	v.cursorButtons = 0
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		v.cursorButtons |= CursorButtonLeft
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
			wx, wy := widget.GetLocation()
			ww, wh := widget.GetSize()
			if v.CursorX >= wx && v.CursorX <= wx+int(ww) && v.CursorY >= wy && v.CursorY <= wy+int(wh) {
				widget.SetPressed(true)
				if v.pressedIndex == -1 {
					found = true
					v.pressedIndex = i
				} else if v.pressedIndex > -1 && v.pressedIndex != i {
					v.widgets[i].SetPressed(false)
				} else {
					v.widgets[i].SetPressed(true)
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
			wx, wy := widget.GetLocation()
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
