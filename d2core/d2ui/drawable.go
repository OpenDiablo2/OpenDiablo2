package d2ui

import "github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"

// Drawable represents an instance that can be drawn
type Drawable interface {
	Render(target d2render.Surface)
	Advance(elapsed float64)
	GetSize() (width, height int)
	SetPosition(x, y int)
	GetPosition() (x, y int)
	GetVisible() bool
	SetVisible(visible bool)
}
