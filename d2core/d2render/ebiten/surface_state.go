package ebiten

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/hajimehoshi/ebiten"
)

type surfaceState struct {
	x          int
	y          int
	filter     ebiten.Filter
	color      color.Color
	brightness float64
	effect     d2enum.DrawEffect
}
