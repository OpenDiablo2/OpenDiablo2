package d2asset

import (
	"errors"
	"image/color"

	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"
	"github.com/OpenDiablo2/D2Shared/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2corehelper"

	"github.com/hajimehoshi/ebiten"
)

type playMode int

const (
	playModePause playMode = iota
	playModeForward
	playModeBackward
)

type animationFrame struct {
	width   int
	height  int
	offsetX int
	offsetY int

	image *ebiten.Image
}

type animationDirection struct {
	frames []*animationFrame
}

type Animation struct {
	directions     []*animationDirection
	frameIndex     int
	directionIndex int
	lastFrameTime  float64

	compositeMode ebiten.CompositeMode
	colorMod      color.Color

	playMode   playMode
	playLength float64
	playLoop   bool
}

func createAnimationFromDC6(dc6 *d2dc6.DC6File) (*Animation, error) {
	animation := &Animation{
		playLength: 1.0,
		playLoop:   true,
	}

	for frameIndex, frame := range dc6.Frames {
		image, err := ebiten.NewImage(int(frame.Width), int(frame.Height), ebiten.FilterNearest)
		if err != nil {
			return nil, err
		}

		if err := image.ReplacePixels(frame.ColorData()); err != nil {
			return nil, err
		}

		directionIndex := frameIndex / int(dc6.FramesPerDirection)
		if directionIndex >= len(animation.directions) {
			animation.directions = append(animation.directions, new(animationDirection))
		}

		direction := animation.directions[directionIndex]
		direction.frames = append(direction.frames, &animationFrame{
			width:   int(frame.Width),
			height:  int(frame.Height),
			offsetX: int(frame.OffsetX),
			offsetY: int(frame.OffsetY),
			image:   image,
		})
	}

	return animation, nil
}

func (a *Animation) clone() *Animation {
	animation := *a
	return &animation
}

func (a *Animation) Advance(elapsed float64) error {
	if a.playMode == playModePause {
		return nil
	}

	frameCount := a.GetFrameCount()
	frameLength := a.playLength / float64(frameCount)
	a.lastFrameTime += elapsed
	framesAdvanced := int(a.lastFrameTime / frameLength)
	a.lastFrameTime -= float64(framesAdvanced) * frameLength

	for i := 0; i < framesAdvanced; i++ {
		switch a.playMode {
		case playModeForward:
			a.frameIndex++
			if a.frameIndex >= frameCount {
				if a.playLoop {
					a.frameIndex = 0
				} else {
					a.frameIndex = frameCount - 1
					break
				}
			}
		case playModeBackward:
			a.frameIndex--
			if a.frameIndex < 0 {
				if a.playLoop {
					a.frameIndex = frameCount - 1
				} else {
					a.frameIndex = 0
					break
				}
			}
		}
	}

	return nil
}

func (a *Animation) Render(target *ebiten.Image, offsetX, offsetY int) error {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(frame.offsetX+offsetX), float64(frame.offsetY+offsetY))
	opts.CompositeMode = a.compositeMode
	if a.colorMod != nil {
		opts.ColorM = d2corehelper.ColorToColorM(a.colorMod)
	}

	return target.DrawImage(frame.image, opts)
}

func (a *Animation) GetFrameSize(frameIndex int) (int, int, error) {
	direction := a.directions[a.directionIndex]
	if frameIndex >= len(direction.frames) {
		return 0, 0, errors.New("invalid frame index")
	}

	frame := direction.frames[frameIndex]
	return frame.width, frame.height, nil
}

func (a *Animation) GetCurrentFrameSize() (int, int) {
	width, height, _ := a.GetFrameSize(a.frameIndex)
	return width, height
}

func (a *Animation) GetFrameBounds() (int, int) {
	maxWidth, maxHeight := 0, 0

	direction := a.directions[a.directionIndex]
	for _, frame := range direction.frames {
		maxWidth = d2helper.MaxInt(maxWidth, frame.width)
		maxHeight = d2helper.MaxInt(maxHeight, frame.height)
	}

	return maxWidth, maxHeight
}

func (a *Animation) GetCurrentFrame() int {
	return a.frameIndex
}

func (a *Animation) GetFrameCount() int {
	direction := a.directions[a.directionIndex]
	return len(direction.frames)
}

func (a *Animation) IsOnFirstFrame() bool {
	return a.frameIndex == 0
}

func (a *Animation) IsOnLastFrame() bool {
	return a.frameIndex == a.GetFrameCount()-1
}

func (a *Animation) GetDirectionCount() int {
	return len(a.directions)
}

func (a *Animation) SetDirection(directionIndex int) error {
	if directionIndex >= len(a.directions) {
		return errors.New("invalid direction index")
	}

	a.directionIndex = directionIndex
	a.frameIndex = 0
	return nil
}

func (a *Animation) GetDirection() int {
	return a.directionIndex
}

func (a *Animation) SetCurrentFrame(frameIndex int) error {
	if frameIndex >= a.GetFrameCount() {
		return errors.New("invalid frame index")
	}

	a.frameIndex = frameIndex
	a.lastFrameTime = 0
	return nil
}

func (a *Animation) Rewind() {
	a.SetCurrentFrame(0)
}

func (a *Animation) PlayForward() {
	a.playMode = playModeForward
	a.lastFrameTime = 0
}

func (a *Animation) PlayBackward() {
	a.playMode = playModeBackward
	a.lastFrameTime = 0
}

func (a *Animation) Pause() {
	a.playMode = playModePause
	a.lastFrameTime = 0
}

func (a *Animation) SetPlayLoop(loop bool) {
	a.playLoop = true
}

func (a *Animation) SetPlayLength(playLength float64) {
	a.playLength = playLength
	a.lastFrameTime = 0
}

func (a *Animation) SetPlayLengthMs(playLengthMs int) {
	a.SetPlayLength(float64(playLengthMs) / 1000.0)
}

func (a *Animation) SetColorMod(color color.Color) {
	a.colorMod = color
}

func (a *Animation) SetBlend(blend bool) {
	if blend {
		a.compositeMode = ebiten.CompositeModeLighter
	} else {
		a.compositeMode = ebiten.CompositeModeSourceOver
	}
}
