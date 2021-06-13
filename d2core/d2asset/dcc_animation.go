package d2asset

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var _ d2interface.Animation = &DCCAnimation{} // Static check to confirm struct conforms to
// interface

func newDCCAnimation(
	dcc *d2dcc.DCC,
	pal d2interface.Palette,
	effect d2enum.DrawEffect,
) (d2interface.Animation, error) {
	DCC := &DCCAnimation{
		dcc:     dcc,
		palette: pal,
	}

	anim := &Animation{
		playLength: defaultPlayLength,
		playLoop:   true,
		effect:     effect,
		onBindRenderer: func(r d2interface.Renderer) error {
			if DCC.renderer != r {
				DCC.renderer = r
				return DCC.createSurfaces()
			}

			return nil
		},
	}

	DCC.Animation = *anim

	err := DCC.init()
	if err != nil {
		return nil, err
	}

	return DCC, nil
}

// DCCAnimation represents an animation decoded from DCC
type DCCAnimation struct {
	palette d2interface.Palette
	dcc     *d2dcc.DCC
	Animation
}

func (a *DCCAnimation) init() error {
	a.directions = make([]animationDirection, a.dcc.NumberOfDirections)

	for directionIndex := range a.directions {
		a.directions[directionIndex].frames = make([]animationFrame, a.dcc.FramesPerDirection)
	}

	err := a.decode()

	return err
}

// Clone creates a copy of the animation
func (a *DCCAnimation) Clone() d2interface.Animation {
	clone := &DCCAnimation{}
	clone.Animation = *a.Animation.Clone().(*Animation)
	clone.dcc = a.dcc.Clone()
	clone.palette = a.palette

	return clone
}

// SetDirection places the animation in the direction of an animation
func (a *DCCAnimation) SetDirection(directionIndex int) error {
	const smallestInvalidDirectionIndex = 64
	if directionIndex >= smallestInvalidDirectionIndex {
		return errors.New("invalid direction index")
	}

	direction := d2dcc.Dir64ToDcc(directionIndex, len(a.directions))
	if !a.directions[direction].decoded {
		err := a.decodeDirection(direction)
		if err != nil {
			return err
		}
	}

	a.directionIndex = direction
	a.frameIndex = 0

	return nil
}

func (a *DCCAnimation) decode() error {
	for directionIndex := 0; directionIndex < len(a.directions); directionIndex++ {
		err := a.decodeDirection(directionIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *DCCAnimation) decodeDirection(directionIndex int) error {
	dccDirection := a.dcc.Directions[directionIndex]

	for frameIndex := range dccDirection.Frames {
		if a.directions[directionIndex].frames == nil {
			a.directions[directionIndex].frames = make([]animationFrame, a.dcc.FramesPerDirection)
		}

		a.directions[directionIndex].decoded = true

		frame := a.decodeFrame(directionIndex)
		a.directions[directionIndex].frames[frameIndex] = frame
	}

	return nil
}

func (a *DCCAnimation) decodeFrame(directionIndex int) animationFrame {
	dccDirection := a.dcc.Directions[directionIndex]

	frame := animationFrame{
		width:   dccDirection.Box.Width,
		height:  dccDirection.Box.Height,
		offsetX: dccDirection.Box.Left,
		offsetY: dccDirection.Box.Top,
		decoded: true,
	}

	return frame
}

func (a *DCCAnimation) createSurfaces() error {
	for directionIndex := 0; directionIndex < len(a.directions); directionIndex++ {
		err := a.createDirectionSurfaces(directionIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *DCCAnimation) createDirectionSurfaces(directionIndex int) error {
	for frameIndex := 0; frameIndex < a.dcc.FramesPerDirection; frameIndex++ {
		if !a.directions[directionIndex].decoded {
			err := a.decodeDirection(directionIndex)
			if err != nil {
				return err
			}
		}

		surface, err := a.createFrameSurface(directionIndex, frameIndex)
		if err != nil {
			return err
		}

		a.directions[directionIndex].frames[frameIndex].image = surface
	}

	return nil
}

func (a *DCCAnimation) createFrameSurface(directionIndex, frameIndex int) (d2interface.Surface, error) {
	if !a.directions[directionIndex].frames[frameIndex].decoded {
		frame := a.decodeFrame(directionIndex)
		a.directions[directionIndex].frames[frameIndex] = frame
	}

	dccFrame := a.dcc.Directions[directionIndex].Frames[frameIndex]
	animFrame := a.directions[directionIndex].frames[frameIndex]
	indexData := dccFrame.PixelData

	if len(indexData) != (animFrame.width * animFrame.height) {
		return nil, errors.New("pixel data incorrect")
	}

	colorData := d2util.ImgIndexToRGBA(indexData, a.palette)

	if a.renderer == nil {
		return nil, errors.New("no renderer")
	}

	sfc := a.renderer.NewSurface(animFrame.width, animFrame.height)

	sfc.ReplacePixels(colorData)

	return sfc, nil
}
