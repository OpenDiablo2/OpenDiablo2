package ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/hajimehoshi/ebiten"
)

func d2ToEbitenFilter(filter d2common.Filter) ebiten.Filter {
	switch filter {
	case d2common.FilterDefault:
		return ebiten.FilterDefault
	case d2common.FilterLinear:
		return ebiten.FilterLinear
	case d2common.FilterNearest:
		return ebiten.FilterNearest
	}

	return ebiten.FilterDefault
}

func ebitenToD2Filter(filter ebiten.Filter) d2common.Filter {
	switch filter {
	case ebiten.FilterDefault:
		return d2common.FilterDefault
	case ebiten.FilterLinear:
		return d2common.FilterLinear
	case ebiten.FilterNearest:
		return d2common.FilterNearest
	}

	return d2common.FilterDefault
}
