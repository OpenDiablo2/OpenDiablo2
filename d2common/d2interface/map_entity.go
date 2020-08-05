package d2interface

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"

// MapEntity is something that can be positioned on and rendered on the game map
type MapEntity interface {
	ID() string
	Render(target Surface)
	Advance(tickTime float64)
	GetPosition() d2vector.Position
	GetVelocity() d2vector.Vector
	GetSize() (width, height int)
	GetLayer() int
	GetPositionF() (float64, float64)
	Label() string
	Selectable() bool
	Highlight()
}
