package d2asset

import (
	"errors"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	d2iface "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var _ d2iface.Animation = &DCCAnimation{} // Static check to confirm struct conforms to interface

// DCCAnimation represents an animation decoded from DCC
type DCCAnimation struct {
	animation
	dccPath  string
	palette  d2iface.Palette
	renderer d2iface.Renderer
}

// CreateDCCAnimation creates an animation from d2dcc.DCC and d2dat.DATPalette
func CreateDCCAnimation(renderer d2iface.Renderer, dccPath string, palette d2iface.Palette,
	effect d2enum.DrawEffect) (d2iface.Animation, error) {
	dcc, err := loadDCC(dccPath)
	if err != nil {
		return nil, err
	}

	anim := animation{
		playLength: defaultPlayLength,
		playLoop:   true,
		directions: make([]animationDirection, dcc.NumberOfDirections),
		effect:     effect,
	}

	DCC := DCCAnimation{
		animation: anim,
		dccPath:   dccPath,
		palette:   palette,
		renderer:  renderer,
	}

	err = DCC.SetDirection(0)
	if err != nil {
		return nil, err
	}

	return &DCC, nil
}

// Clone creates a copy of the animation
func (a *DCCAnimation) Clone() d2iface.Animation {
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
	dcc, err := loadDCC(a.dccPath)
	if err != nil {
		return err
	}

	direction := dcc.DecodeDirection(directionIndex)

	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for _, dccFrame := range direction.Frames {
		minX = d2common.MinInt(minX, dccFrame.Box.Left)
		minY = d2common.MinInt(minY, dccFrame.Box.Top)
		maxX = d2common.MaxInt(maxX, dccFrame.Box.Right())
		maxY = d2common.MaxInt(maxY, dccFrame.Box.Bottom())
	}

	for _, dccFrame := range direction.Frames {
		frameWidth := maxX - minX
		frameHeight := maxY - minY

		pixels := ImgIndexToRGBA(dccFrame.PixelData, a.palette)

		sfc, err := a.renderer.NewSurface(frameWidth, frameHeight, d2enum.FilterNearest)
		if err != nil {
			return err
		}

		if err := sfc.ReplacePixels(pixels); err != nil {
			return err
		}

		a.directions[directionIndex].decoded = true
		a.directions[directionIndex].frames = append(a.directions[directionIndex].frames, &animationFrame{
			width:   dccFrame.Width,
			height:  dccFrame.Height,
			offsetX: minX,
			offsetY: minY,
			image:   sfc,
		})
	}

	return nil
}
