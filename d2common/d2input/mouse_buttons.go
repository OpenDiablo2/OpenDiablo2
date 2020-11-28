package d2input

import "github.com/hajimehoshi/ebiten/v2"

type MouseButton = int

const (
	MouseButtonLeft   = MouseButton(ebiten.MouseButtonLeft)
	MouseButtonRight  = MouseButton(ebiten.MouseButtonRight)
	MouseButtonMiddle = MouseButton(ebiten.MouseButtonMiddle)
)
