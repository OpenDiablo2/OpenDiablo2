package ebiten

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func d2ToEbitenFilter(filter d2enum.Filter) ebiten.Filter {
	switch filter {
	case d2enum.FilterDefault:
		return ebiten.FilterDefault
	case d2enum.FilterLinear:
		return ebiten.FilterLinear
	case d2enum.FilterNearest:
		return ebiten.FilterNearest
	}

	return ebiten.FilterDefault
}

// func ebitenToD2Filter(filter ebiten.Filter) d2enum.Filter {
// 	switch filter {
// 	case ebiten.FilterDefault:
// 		return d2enum.FilterDefault
// 	case ebiten.FilterLinear:
// 		return d2enum.FilterLinear
// 	case ebiten.FilterNearest:
// 		return d2enum.FilterNearest
// 	}
//
// 	return d2enum.FilterDefault
// }
