package ebiten

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type surfaceState struct {
	x      int
	y      int
	mode   ebiten.CompositeMode
	filter ebiten.Filter
	color  color.Color
	brightness float64
}
