package ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/hajimehoshi/ebiten"
)

func d2ToEbitenCompositeMode(comp d2enum.CompositeMode) ebiten.CompositeMode {
	switch comp {
	case d2enum.CompositeModeSourceOver:
		return ebiten.CompositeModeSourceOver
	case d2enum.CompositeModeClear:
		return ebiten.CompositeModeClear
	case d2enum.CompositeModeCopy:
		return ebiten.CompositeModeCopy
	case d2enum.CompositeModeDestination:
		return ebiten.CompositeModeDestination
	case d2enum.CompositeModeDestinationOver:
		return ebiten.CompositeModeDestinationOver
	case d2enum.CompositeModeSourceIn:
		return ebiten.CompositeModeSourceIn
	case d2enum.CompositeModeDestinationIn:
		return ebiten.CompositeModeDestinationIn
	case d2enum.CompositeModeSourceOut:
		return ebiten.CompositeModeSourceOut
	case d2enum.CompositeModeDestinationOut:
		return ebiten.CompositeModeDestinationOut
	case d2enum.CompositeModeSourceAtop:
		return ebiten.CompositeModeSourceAtop
	case d2enum.CompositeModeDestinationAtop:
		return ebiten.CompositeModeDestinationAtop
	case d2enum.CompositeModeXor:
		return ebiten.CompositeModeXor
	case d2enum.CompositeModeLighter:
		return ebiten.CompositeModeLighter
	}

	return ebiten.CompositeModeSourceOver
}

func ebitenToD2CompositeMode(comp ebiten.CompositeMode) d2enum.CompositeMode {
	switch comp {
	case ebiten.CompositeModeSourceOver:
		return d2enum.CompositeModeSourceOver
	case ebiten.CompositeModeClear:
		return d2enum.CompositeModeClear
	case ebiten.CompositeModeCopy:
		return d2enum.CompositeModeCopy
	case ebiten.CompositeModeDestination:
		return d2enum.CompositeModeDestination
	case ebiten.CompositeModeDestinationOver:
		return d2enum.CompositeModeDestinationOver
	case ebiten.CompositeModeSourceIn:
		return d2enum.CompositeModeSourceIn
	case ebiten.CompositeModeDestinationIn:
		return d2enum.CompositeModeDestinationIn
	case ebiten.CompositeModeSourceOut:
		return d2enum.CompositeModeSourceOut
	case ebiten.CompositeModeDestinationOut:
		return d2enum.CompositeModeDestinationOut
	case ebiten.CompositeModeSourceAtop:
		return d2enum.CompositeModeSourceAtop
	case ebiten.CompositeModeDestinationAtop:
		return d2enum.CompositeModeDestinationAtop
	case ebiten.CompositeModeXor:
		return d2enum.CompositeModeXor
	case ebiten.CompositeModeLighter:
		return d2enum.CompositeModeLighter
	}

	return d2enum.CompositeModeSourceOver
}
