package Common

import "github.com/hajimehoshi/ebiten"

// Drawable represents an instance that can be drawn
type Drawable interface {
	Draw(target *ebiten.Image)
	GetSize() (width, height uint32)
	MoveTo(x, y int)
	GetLocation() (x, y int)
	GetVisible() bool
	SetVisible(visible bool)
}
