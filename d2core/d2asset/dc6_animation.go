package d2asset

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var _ d2interface.Animation = &DC6Animation{} // Static check to confirm struct conforms to
// interface

func newDC6Animation(
	dc6 *d2dc6.DC6,
	pal d2interface.Palette,
	effect d2enum.DrawEffect,
) (d2interface.Animation, error) {
	DC6 := &DC6Animation{
		dc6:     dc6,
		palette: pal,
	}

	anim := &Animation{
		playLength:     defaultPlayLength,
		playLoop:       true,
		originAtBottom: true,
		effect:         effect,
		onBindRenderer: func(r d2interface.Renderer) error {
			if DC6.renderer != r {
				DC6.renderer = r
				return DC6.createSurfaces()
			}

			return nil
		},
	}

	DC6.Animation = *anim

	err := DC6.init()
	if err != nil {
		return nil, err
	}

	return DC6, nil
}

// DC6Animation is an animation made from a DC6 file
type DC6Animation struct {
	Animation
	dc6     *d2dc6.DC6
	palette d2interface.Palette
}

func (a *DC6Animation) init() error {
	a.directions = make([]animationDirection, a.dc6.Directions)

	for directionIndex := range a.directions {
		a.directions[directionIndex].frames = make([]animationFrame, a.dc6.FramesPerDirection)
	}

	err := a.decode()

	return err
}

// SetDirection decodes and sets the direction
func (a *DC6Animation) SetDirection(directionIndex int) error {
	const smallestInvalidDirectionIndex = 64
	if directionIndex >= smallestInvalidDirectionIndex {
		return errors.New("invalid direction index")
	}

	direction := d2dcc.Dir64ToDcc(directionIndex, len(a.directions))

	if !a.directions[directionIndex].decoded {
		err := a.decodeDirection(direction)
		if err != nil {
			return err
		}
	}

	a.directionIndex = direction

	return nil
}

func (a *DC6Animation) decode() error {
	for directionIndex := 0; directionIndex < len(a.directions); directionIndex++ {
		err := a.decodeDirection(directionIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *DC6Animation) decodeDirection(directionIndex int) error {
	for frameIndex := 0; frameIndex < int(a.dc6.FramesPerDirection); frameIndex++ {
		frame := a.decodeFrame(directionIndex, frameIndex)
		a.directions[directionIndex].frames[frameIndex] = frame
	}

	a.directions[directionIndex].decoded = true

	return nil
}

func (a *DC6Animation) decodeFrame(directionIndex, frameIndex int) animationFrame {
	startFrame := directionIndex * int(a.dc6.FramesPerDirection)

	dc6Frame := a.dc6.Frames[startFrame+frameIndex]

	frame := animationFrame{
		width:   int(dc6Frame.Width),
		height:  int(dc6Frame.Height),
		offsetX: int(dc6Frame.OffsetX),
		offsetY: int(dc6Frame.OffsetY),
	}

	a.directions[directionIndex].frames[frameIndex].decoded = true

	return frame
}

func (a *DC6Animation) createSurfaces() error {
	for directionIndex := 0; directionIndex < len(a.directions); directionIndex++ {
		err := a.createDirectionSurfaces(directionIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *DC6Animation) createDirectionSurfaces(directionIndex int) error {
	for frameIndex := 0; frameIndex < int(a.dc6.FramesPerDirection); frameIndex++ {
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

func (a *DC6Animation) createFrameSurface(directionIndex, frameIndex int) (d2interface.Surface, error) {
	if !a.directions[directionIndex].frames[frameIndex].decoded {
		frame := a.decodeFrame(directionIndex, frameIndex)
		a.directions[directionIndex].frames[frameIndex] = frame
	}

	startFrame := directionIndex * int(a.dc6.FramesPerDirection)
	dc6Frame := a.dc6.Frames[startFrame+frameIndex]
	indexData := a.dc6.DecodeFrame(startFrame + frameIndex)
	colorData := d2util.ImgIndexToRGBA(indexData, a.palette)

	if a.renderer == nil {
		return nil, errors.New("no renderer")
	}

	sfc := a.renderer.NewSurface(int(dc6Frame.Width), int(dc6Frame.Height))

	sfc.ReplacePixels(colorData)

	return sfc, nil
}

// Clone creates a copy of the animation
func (a *DC6Animation) Clone() d2interface.Animation {
	clone := &DC6Animation{}
	clone.Animation = *a.Animation.Clone().(*Animation)
	clone.dc6 = a.dc6.Clone()
	clone.palette = a.palette

	return clone
}
