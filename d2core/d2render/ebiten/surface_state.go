package ebiten

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type surfaceState struct {
	color      color.Color
	x          int
	y          int
	filter     ebiten.Filter
	brightness float64
	saturation float64
	effect     d2enum.DrawEffect
	skewX      float64
	skewY      float64
	scaleX     float64
	scaleY     float64
}
