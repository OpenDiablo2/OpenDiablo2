package ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/hajimehoshi/ebiten"
)

func d2ToEbitenCompositeMode(comp d2render.CompositeMode) ebiten.CompositeMode {
	switch comp {
	case d2render.CompositeModeSourceOver:
		return ebiten.CompositeModeSourceOver
	case d2render.CompositeModeClear:
		return ebiten.CompositeModeClear
	case d2render.CompositeModeCopy:
		return ebiten.CompositeModeCopy
	case d2render.CompositeModeDestination:
		return ebiten.CompositeModeDestination
	case d2render.CompositeModeDestinationOver:
		return ebiten.CompositeModeDestinationOver
	case d2render.CompositeModeSourceIn:
		return ebiten.CompositeModeSourceIn
	case d2render.CompositeModeDestinationIn:
		return ebiten.CompositeModeDestinationIn
	case d2render.CompositeModeSourceOut:
		return ebiten.CompositeModeSourceOut
	case d2render.CompositeModeDestinationOut:
		return ebiten.CompositeModeDestinationOut
	case d2render.CompositeModeSourceAtop:
		return ebiten.CompositeModeSourceAtop
	case d2render.CompositeModeDestinationAtop:
		return ebiten.CompositeModeDestinationAtop
	case d2render.CompositeModeXor:
		return ebiten.CompositeModeXor
	case d2render.CompositeModeLighter:
		return ebiten.CompositeModeLighter
	}

	return ebiten.CompositeModeSourceOver
}

func ebitenToD2CompositeMode(comp ebiten.CompositeMode) d2render.CompositeMode {
	switch comp {
	case ebiten.CompositeModeSourceOver:
		return d2render.CompositeModeSourceOver
	case ebiten.CompositeModeClear:
		return d2render.CompositeModeClear
	case ebiten.CompositeModeCopy:
		return d2render.CompositeModeCopy
	case ebiten.CompositeModeDestination:
		return d2render.CompositeModeDestination
	case ebiten.CompositeModeDestinationOver:
		return d2render.CompositeModeDestinationOver
	case ebiten.CompositeModeSourceIn:
		return d2render.CompositeModeSourceIn
	case ebiten.CompositeModeDestinationIn:
		return d2render.CompositeModeDestinationIn
	case ebiten.CompositeModeSourceOut:
		return d2render.CompositeModeSourceOut
	case ebiten.CompositeModeDestinationOut:
		return d2render.CompositeModeDestinationOut
	case ebiten.CompositeModeSourceAtop:
		return d2render.CompositeModeSourceAtop
	case ebiten.CompositeModeDestinationAtop:
		return d2render.CompositeModeDestinationAtop
	case ebiten.CompositeModeXor:
		return d2render.CompositeModeXor
	case ebiten.CompositeModeLighter:
		return d2render.CompositeModeLighter
	}

	return d2render.CompositeModeSourceOver
}
