package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// Drawable represents an instance that can be drawn
type Drawable interface {
	Render(target d2common.Surface)
	Advance(elapsed float64)
	GetSize() (width, height int)
	SetPosition(x, y int)
	GetPosition() (x, y int)
	GetVisible() bool
	SetVisible(visible bool)
}
