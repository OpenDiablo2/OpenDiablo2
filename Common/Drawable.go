package Common

import "github.com/hajimehoshi/ebiten"

// Drawable represents an instance that can be drawn
type Drawable interface {
	Draw(target *ebiten.Image)
	GetSize() (uint32, uint32)
	MoveTo(x, y int)
	GetLocation() (int, int)
	GetVisible() bool
	SetVisible(bool)
}
