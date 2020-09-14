package d2asset

import (
	"errors"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var _ d2interface.Animation = &DCCAnimation{} // Static check to confirm struct conforms to
// interface

// DCCAnimation represents an animation decoded from DCC
type DCCAnimation struct {
	animation
	*AssetManager
	dccPath  string
	palette  d2interface.Palette
	renderer d2interface.Renderer
}

// Clone creates a copy of the animation
func (a *DCCAnimation) Clone() d2interface.Animation {
	animation := *a
	return &animation
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

func (a *DCCAnimation) decodeDirection(directionIndex int) error {
	dcc, err := a.loadDCC(a.dccPath)
	if err != nil {
		return err
	}

	direction := dcc.DecodeDirection(directionIndex)

	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for _, dccFrame := range direction.Frames {
		minX = d2math.MinInt(minX, dccFrame.Box.Left)
		minY = d2math.MinInt(minY, dccFrame.Box.Top)
		maxX = d2math.MaxInt(maxX, dccFrame.Box.Right())
		maxY = d2math.MaxInt(maxY, dccFrame.Box.Bottom())
	}

	frameWidth := maxX - minX
	frameHeight := maxY - minY

	for _, dccFrame := range direction.Frames {
		pixels := d2util.ImgIndexToRGBA(dccFrame.PixelData, a.palette)

		sfc, err := a.renderer.NewSurface(frameWidth, frameHeight, d2enum.FilterNearest)
		if err != nil {
			return err
		}

		if err := sfc.ReplacePixels(pixels); err != nil {
			return err
		}

		a.directions[directionIndex].decoded = true
		a.directions[directionIndex].frames = append(a.directions[directionIndex].frames, &animationFrame{
			width:   frameWidth,
			height:  frameHeight,
			offsetX: minX,
			offsetY: minY,
			image:   sfc,
		})
	}

	return nil
}
