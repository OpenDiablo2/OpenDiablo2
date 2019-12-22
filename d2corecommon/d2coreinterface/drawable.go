package d2coreinterface

import "github.com/hajimehoshi/ebiten"

// Drawable represents an instance that can be drawn
type Drawable interface {
	Render(target *ebiten.Image)
	GetSize() (width, height int)
	SetPosition(x, y int)
	GetPosition() (x, y int)
	GetVisible() bool
	SetVisible(visible bool)
}
