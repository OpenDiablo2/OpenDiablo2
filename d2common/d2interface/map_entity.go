package d2interface

// MapEntity is something that can be positioned on and rendered on the game map
type MapEntity interface {
	Render(target Surface)
	Advance(tickTime float64)
	GetPosition() (float64, float64)
	GetLayer() int
	GetPositionF() (float64, float64)
	Name() string
	Selectable() bool
	Highlight()
}
