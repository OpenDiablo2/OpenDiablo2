package d2asset

import (
	"errors"
	"image/color"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"

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
	playedCount    int

	compositeMode ebiten.CompositeMode
	colorMod      color.Color

	playMode   playMode
	playLength float64
	playLoop   bool
}

func createAnimationFromDCC(dcc *d2dcc.DCC, palette *d2datadict.PaletteRec, transparency int) (*Animation, error) {
	animation := &Animation{
		playLength: 1.0,
		playLoop:   true,
	}

	for directionIndex, dccDirection := range dcc.Directions {
		for _, dccFrame := range dccDirection.Frames {
			minX, minY := math.MaxInt32, math.MaxInt32
			maxX, maxY := math.MinInt32, math.MinInt32
			for _, dccFrame := range dccDirection.Frames {
				minX = d2helper.MinInt(minX, dccFrame.Box.Left)
				minY = d2helper.MinInt(minY, dccFrame.Box.Top)
				maxX = d2helper.MaxInt(maxX, dccFrame.Box.Right())
				maxY = d2helper.MaxInt(maxY, dccFrame.Box.Bottom())
			}

			frameWidth := maxX - minX
			frameHeight := maxY - minY

			pixels := make([]byte, frameWidth*frameHeight*4)
			for y := 0; y < frameHeight; y++ {
				for x := 0; x < frameWidth; x++ {
					if paletteIndex := dccFrame.PixelData[y*frameWidth+x]; paletteIndex != 0 {
						color := palette.Colors[paletteIndex]
						offset := (x + y*frameWidth) * 4
						pixels[offset] = color.R
						pixels[offset+1] = color.G
						pixels[offset+2] = color.B
						pixels[offset+3] = byte(transparency)
					}
				}
			}

			image, err := ebiten.NewImage(frameWidth, frameHeight, ebiten.FilterNearest)
			if err != nil {
				return nil, err
			}

			if err := image.ReplacePixels(pixels); err != nil {
				return nil, err
			}

			if directionIndex >= len(animation.directions) {
				animation.directions = append(animation.directions, new(animationDirection))
			}

			direction := animation.directions[directionIndex]
			direction.frames = append(direction.frames, &animationFrame{
				width:   int(dccFrame.Width),
				height:  int(dccFrame.Height),
				offsetX: minX,
				offsetY: minY,
				image:   image,
			})

		}
	}

	return animation, nil
}

func createAnimationFromDC6(dc6 *d2dc6.DC6File) (*Animation, error) {
	animation := &Animation{
		playLength: 1.0,
		playLoop:   true,
	}

	for frameIndex, dc6Frame := range dc6.Frames {
		image, err := ebiten.NewImage(int(dc6Frame.Width), int(dc6Frame.Height), ebiten.FilterNearest)
		if err != nil {
			return nil, err
		}

		if err := image.ReplacePixels(dc6Frame.ColorData()); err != nil {
			return nil, err
		}

		directionIndex := frameIndex / int(dc6.FramesPerDirection)
		if directionIndex >= len(animation.directions) {
			animation.directions = append(animation.directions, new(animationDirection))
		}

		direction := animation.directions[directionIndex]
		direction.frames = append(direction.frames, &animationFrame{
			width:   int(dc6Frame.Width),
			height:  int(dc6Frame.Height),
			offsetX: int(dc6Frame.OffsetX),
			offsetY: int(dc6Frame.OffsetY),
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
				a.playedCount++
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
				a.playedCount++
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

func (a *Animation) Render(target *d2surface.Surface) error {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushTranslation(frame.offsetX, frame.offsetY)
	target.PushCompositeMode(a.compositeMode)
	target.PushColor(a.colorMod)
	defer target.PopN(3)
	return target.Render(frame.image)
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

func (a *Animation) SetPlaySpeed(playSpeed float64) {
	a.SetPlayLength(playSpeed * float64(a.GetFrameCount()))
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

func (a *Animation) GetPlayedCount() int {
	return a.playedCount
}

func (a *Animation) ResetPlayedCount() {
	a.playedCount = 0
}

func (a *Animation) SetBlend(blend bool) {
	if blend {
		a.compositeMode = ebiten.CompositeModeLighter
	} else {
		a.compositeMode = ebiten.CompositeModeSourceOver
	}
}
