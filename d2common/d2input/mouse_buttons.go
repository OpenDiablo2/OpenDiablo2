package d2input

import "github.com/hajimehoshi/ebiten/v2"

// MouseButton represents a button on a mouse
type MouseButton = int

// MouseButtons
const (
	MouseButtonLeft   = MouseButton(ebiten.MouseButtonLeft)
	MouseButtonRight  = MouseButton(ebiten.MouseButtonRight)
	MouseButtonMiddle = MouseButton(ebiten.MouseButtonMiddle)
)
