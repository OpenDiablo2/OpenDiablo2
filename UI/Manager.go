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
	CursorX       int
	CursorY       int
}

// CreateManager creates a new instance of a UI manager
func CreateManager(provider Common.FileProvider) *Manager {
	result := &Manager{
		widgets:      make([]Widget, 0),
		cursorSprite: provider.LoadSprite(ResourcePaths.CursorDefault, Palettes.Units),
	}
	return result
}

// Reset resets the state of the UI manager. Typically called for new scenes
func (v *Manager) Reset() {
	v.widgets = make([]Widget, 0)
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
}

// CursorButtonPressed determines if the specified button has been pressed
func (v *Manager) CursorButtonPressed(button CursorButton) bool {
	return v.cursorButtons&button > 0
}
