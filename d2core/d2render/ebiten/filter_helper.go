package ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/hajimehoshi/ebiten"
)

func d2ToEbitenFilter(filter d2render.Filter) ebiten.Filter {
	switch filter {
	case d2render.FilterDefault:
		return ebiten.FilterDefault
	case d2render.FilterLinear:
		return ebiten.FilterLinear
	case d2render.FilterNearest:
		return ebiten.FilterNearest
	}

	return ebiten.FilterDefault
}

func ebitenToD2Filter(filter ebiten.Filter) d2render.Filter {
	switch filter {
	case ebiten.FilterDefault:
		return d2render.FilterDefault
	case ebiten.FilterLinear:
		return d2render.FilterLinear
	case ebiten.FilterNearest:
		return d2render.FilterNearest
	}

	return d2render.FilterDefault
}
