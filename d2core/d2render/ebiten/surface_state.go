package ebiten

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type surfaceState struct {
	x              int
	y              int
	filter         ebiten.Filter
	color          color.Color
	brightness     float64
	saturation     float64
	effect         d2enum.DrawEffect
	skewX, skewY   float64
	scaleX, scaleY float64
	rotate         float64
}
