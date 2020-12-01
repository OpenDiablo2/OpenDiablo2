package d2sprite

import (
	"errors"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var _ d2interface.Sprite = &DCCSprite{} // Static check to confirm struct conforms to
// interface

// NewDCCSprite creates an sprite from a dcc file
func NewDCCSprite(
	dcc *d2dcc.DCC,
	pal d2interface.Palette,
	effect d2enum.DrawEffect,
) (d2interface.Sprite, error) {
	DCC := &DCCSprite{
		dcc:     dcc,
		palette: pal,
	}

	anim := &Sprite{
		playLoop: true,
		effect:   effect,
		onBindRenderer: func(r d2interface.Renderer) error {
			if DCC.renderer != r {
				DCC.renderer = r
				return DCC.createSurfaces()
			}

			return nil
		},
	}

	DCC.Sprite = *anim

	err := DCC.init()
	if err != nil {
		return nil, err
	}

	return DCC, nil
}

// DCCSprite represents an sprite decoded from DCC
type DCCSprite struct {
	Sprite
	dcc     *d2dcc.DCC
	palette d2interface.Palette
}

func (a *DCCSprite) init() error {
	a.directions = make([]spriteDirection, a.dcc.NumberOfDirections)

	for directionIndex := range a.directions {
		a.directions[directionIndex].frames = make([]spriteFrame, a.dcc.FramesPerDirection)
	}

	err := a.decode()

	a.SetPlayFPS(fps25)

	return err
}

// Clone creates a copy of the sprite
func (a *DCCSprite) Clone() d2interface.Sprite {
	clone := &DCCSprite{}
	clone.Sprite = *a.Sprite.Clone().(*Sprite)
	clone.dcc = a.dcc.Clone()
	clone.palette = a.palette

	return clone
}

// SetDirection places the sprite in the direction of an sprite
func (a *DCCSprite) SetDirection(directionIndex int) error {
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

func (a *DCCSprite) decode() error {
	for directionIndex := 0; directionIndex < len(a.directions); directionIndex++ {
		err := a.decodeDirection(directionIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *DCCSprite) decodeDirection(directionIndex int) error {
	dccDirection := a.dcc.Directions[directionIndex]

	for frameIndex := range dccDirection.Frames {
		if a.directions[directionIndex].frames == nil {
			a.directions[directionIndex].frames = make([]spriteFrame, a.dcc.FramesPerDirection)
		}

		a.directions[directionIndex].decoded = true

		frame := a.decodeFrame(directionIndex)
		a.directions[directionIndex].frames[frameIndex] = frame
	}

	return nil
}

func (a *DCCSprite) decodeFrame(directionIndex int) spriteFrame {
	dccDirection := a.dcc.Directions[directionIndex]

	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for _, dccFrame := range dccDirection.Frames {
		minX = d2math.MinInt(minX, dccFrame.Box.Left)
		minY = d2math.MinInt(minY, dccFrame.Box.Top)
		maxX = d2math.MaxInt(maxX, dccFrame.Box.Right())
		maxY = d2math.MaxInt(maxY, dccFrame.Box.Bottom())
	}

	frameWidth := maxX - minX
	frameHeight := maxY - minY

	frame := spriteFrame{
		width:   frameWidth,
		height:  frameHeight,
		offsetX: minX,
		offsetY: minY,
		decoded: true,
	}

	return frame
}

func (a *DCCSprite) createSurfaces() error {
	for directionIndex := 0; directionIndex < len(a.directions); directionIndex++ {
		err := a.createDirectionSurfaces(directionIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *DCCSprite) createDirectionSurfaces(directionIndex int) error {
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

func (a *DCCSprite) createFrameSurface(directionIndex, frameIndex int) (d2interface.Surface, error) {
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
