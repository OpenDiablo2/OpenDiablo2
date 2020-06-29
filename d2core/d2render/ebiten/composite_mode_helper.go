package ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/hajimehoshi/ebiten"
)

func d2ToEbitenCompositeMode(comp d2interface.CompositeMode) ebiten.CompositeMode {
	switch comp {
	case d2interface.CompositeModeSourceOver:
		return ebiten.CompositeModeSourceOver
	case d2interface.CompositeModeClear:
		return ebiten.CompositeModeClear
	case d2interface.CompositeModeCopy:
		return ebiten.CompositeModeCopy
	case d2interface.CompositeModeDestination:
		return ebiten.CompositeModeDestination
	case d2interface.CompositeModeDestinationOver:
		return ebiten.CompositeModeDestinationOver
	case d2interface.CompositeModeSourceIn:
		return ebiten.CompositeModeSourceIn
	case d2interface.CompositeModeDestinationIn:
		return ebiten.CompositeModeDestinationIn
	case d2interface.CompositeModeSourceOut:
		return ebiten.CompositeModeSourceOut
	case d2interface.CompositeModeDestinationOut:
		return ebiten.CompositeModeDestinationOut
	case d2interface.CompositeModeSourceAtop:
		return ebiten.CompositeModeSourceAtop
	case d2interface.CompositeModeDestinationAtop:
		return ebiten.CompositeModeDestinationAtop
	case d2interface.CompositeModeXor:
		return ebiten.CompositeModeXor
	case d2interface.CompositeModeLighter:
		return ebiten.CompositeModeLighter
	}

	return ebiten.CompositeModeSourceOver
}

func ebitenToD2CompositeMode(comp ebiten.CompositeMode) d2interface.CompositeMode {
	switch comp {
	case ebiten.CompositeModeSourceOver:
		return d2interface.CompositeModeSourceOver
	case ebiten.CompositeModeClear:
		return d2interface.CompositeModeClear
	case ebiten.CompositeModeCopy:
		return d2interface.CompositeModeCopy
	case ebiten.CompositeModeDestination:
		return d2interface.CompositeModeDestination
	case ebiten.CompositeModeDestinationOver:
		return d2interface.CompositeModeDestinationOver
	case ebiten.CompositeModeSourceIn:
		return d2interface.CompositeModeSourceIn
	case ebiten.CompositeModeDestinationIn:
		return d2interface.CompositeModeDestinationIn
	case ebiten.CompositeModeSourceOut:
		return d2interface.CompositeModeSourceOut
	case ebiten.CompositeModeDestinationOut:
		return d2interface.CompositeModeDestinationOut
	case ebiten.CompositeModeSourceAtop:
		return d2interface.CompositeModeSourceAtop
	case ebiten.CompositeModeDestinationAtop:
		return d2interface.CompositeModeDestinationAtop
	case ebiten.CompositeModeXor:
		return d2interface.CompositeModeXor
	case ebiten.CompositeModeLighter:
		return d2interface.CompositeModeLighter
	}

	return d2interface.CompositeModeSourceOver
}
