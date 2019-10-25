package UI

import (
	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/hajimehoshi/ebiten"
)

// Manager represents the UI manager
type Manager struct {
	widgets      []*Widget
	cursorSprite *Common.Sprite
}

// CreateManager creates a new instance of a UI manager
func CreateManager(provider Common.SpriteProvider) *Manager {
	result := &Manager{
		widgets:      make([]*Widget, 0),
		cursorSprite: provider.LoadSprite(ResourcePaths.CursorDefault, Palettes.Units),
	}
	return result
}

// Reset resets the state of the UI manager. Typically called for new scenes
func (v *Manager) Reset() {
	v.widgets = make([]*Widget, 0)
}

// AddWidget adds a widget to the UI manager
func (v *Manager) AddWidget(widget *Widget) {
	v.widgets = append(v.widgets, widget)
}

// Draw renders all of the UI elements
func (v *Manager) Draw(screen *ebiten.Image) {
	cx, cy := ebiten.CursorPosition()
	v.cursorSprite.MoveTo(cx, cy)
	v.cursorSprite.Draw(screen)
}

// Update updates all of the UI elements
func (v *Manager) Update() {

}
