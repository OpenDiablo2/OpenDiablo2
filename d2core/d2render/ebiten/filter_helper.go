package ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/hajimehoshi/ebiten"
)

func d2ToEbitenFilter(filter d2interface.Filter) ebiten.Filter {
	switch filter {
	case d2interface.FilterDefault:
		return ebiten.FilterDefault
	case d2interface.FilterLinear:
		return ebiten.FilterLinear
	case d2interface.FilterNearest:
		return ebiten.FilterNearest
	}

	return ebiten.FilterDefault
}

func ebitenToD2Filter(filter ebiten.Filter) d2interface.Filter {
	switch filter {
	case ebiten.FilterDefault:
		return d2interface.FilterDefault
	case ebiten.FilterLinear:
		return d2interface.FilterLinear
	case ebiten.FilterNearest:
		return d2interface.FilterNearest
	}

	return d2interface.FilterDefault
}
