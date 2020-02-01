package ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/hajimehoshi/ebiten"
)

func d2ToEbitenCompositeMode(comp d2common.CompositeMode) ebiten.CompositeMode {
	switch comp {
	case d2common.CompositeModeSourceOver:
		return ebiten.CompositeModeSourceOver
	case d2common.CompositeModeClear:
		return ebiten.CompositeModeClear
	case d2common.CompositeModeCopy:
		return ebiten.CompositeModeCopy
	case d2common.CompositeModeDestination:
		return ebiten.CompositeModeDestination
	case d2common.CompositeModeDestinationOver:
		return ebiten.CompositeModeDestinationOver
	case d2common.CompositeModeSourceIn:
		return ebiten.CompositeModeSourceIn
	case d2common.CompositeModeDestinationIn:
		return ebiten.CompositeModeDestinationIn
	case d2common.CompositeModeSourceOut:
		return ebiten.CompositeModeSourceOut
	case d2common.CompositeModeDestinationOut:
		return ebiten.CompositeModeDestinationOut
	case d2common.CompositeModeSourceAtop:
		return ebiten.CompositeModeSourceAtop
	case d2common.CompositeModeDestinationAtop:
		return ebiten.CompositeModeDestinationAtop
	case d2common.CompositeModeXor:
		return ebiten.CompositeModeXor
	case d2common.CompositeModeLighter:
		return ebiten.CompositeModeLighter
	}

	return ebiten.CompositeModeSourceOver
}

func ebitenToD2CompositeMode(comp ebiten.CompositeMode) d2common.CompositeMode {
	switch comp {
	case ebiten.CompositeModeSourceOver:
		return d2common.CompositeModeSourceOver
	case ebiten.CompositeModeClear:
		return d2common.CompositeModeClear
	case ebiten.CompositeModeCopy:
		return d2common.CompositeModeCopy
	case ebiten.CompositeModeDestination:
		return d2common.CompositeModeDestination
	case ebiten.CompositeModeDestinationOver:
		return d2common.CompositeModeDestinationOver
	case ebiten.CompositeModeSourceIn:
		return d2common.CompositeModeSourceIn
	case ebiten.CompositeModeDestinationIn:
		return d2common.CompositeModeDestinationIn
	case ebiten.CompositeModeSourceOut:
		return d2common.CompositeModeSourceOut
	case ebiten.CompositeModeDestinationOut:
		return d2common.CompositeModeDestinationOut
	case ebiten.CompositeModeSourceAtop:
		return d2common.CompositeModeSourceAtop
	case ebiten.CompositeModeDestinationAtop:
		return d2common.CompositeModeDestinationAtop
	case ebiten.CompositeModeXor:
		return d2common.CompositeModeXor
	case ebiten.CompositeModeLighter:
		return d2common.CompositeModeLighter
	}

	return d2common.CompositeModeSourceOver
}
