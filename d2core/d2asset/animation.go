package d2asset

import (
	"errors"
	"image"
	"image/color"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type playMode int

const (
	playModePause playMode = iota
	playModeForward
	playModeBackward
)

const defaultPlayLength = 1.0

type animationFrame struct {
	width   int
	height  int
	offsetX int
	offsetY int

	image d2interface.Surface
}

type animationDirection struct {
	frames []*animationFrame
}

// Animation has directionality, play modes, and frame counting
type Animation struct {
	directions     []*animationDirection
	frameIndex     int
	directionIndex int
	lastFrameTime  float64
	playedCount    int

	compositeMode  d2enum.CompositeMode
	colorMod       color.Color
	originAtBottom bool

	playMode         playMode
	playLength       float64
	playLoop         bool
	hasSubLoop       bool // runs after first animation ends
	subStartingFrame int
	subEndingFrame   int
}

// CreateAnimationFromDCC creates an animation from d2dcc.DCC and d2dat.DATPalette
func CreateAnimationFromDCC(dcc *d2dcc.DCC, palette *d2dat.DATPalette, transparency int) (*Animation, error) {
	animation := &Animation{
		playLength: defaultPlayLength,
		playLoop:   true,
	}

	for directionIndex, dccDirection := range dcc.Directions {
		for _, dccFrame := range dccDirection.Frames {
			minX, minY := math.MaxInt32, math.MaxInt32
			maxX, maxY := math.MinInt32, math.MinInt32

			for _, dccFrame := range dccDirection.Frames {
				minX = d2common.MinInt(minX, dccFrame.Box.Left)
				minY = d2common.MinInt(minY, dccFrame.Box.Top)
				maxX = d2common.MaxInt(maxX, dccFrame.Box.Right())
				maxY = d2common.MaxInt(maxY, dccFrame.Box.Bottom())
			}

			frameWidth := maxX - minX
			frameHeight := maxY - minY

			const bytesPerPixel = 4
			pixels := make([]byte, frameWidth*frameHeight*bytesPerPixel)

			for y := 0; y < frameHeight; y++ {
				for x := 0; x < frameWidth; x++ {
					if paletteIndex := dccFrame.PixelData[y*frameWidth+x]; paletteIndex != 0 {
						palColor := palette.Colors[paletteIndex]
						offset := (x + y*frameWidth) * bytesPerPixel
						pixels[offset] = palColor.R
						pixels[offset+1] = palColor.G
						pixels[offset+2] = palColor.B
						pixels[offset+3] = byte(transparency)
					}
				}
			}

			sfc, err := d2render.NewSurface(frameWidth, frameHeight, d2interface.FilterNearest)
			if err != nil {
				return nil, err
			}

			if err := sfc.ReplacePixels(pixels); err != nil {
				return nil, err
			}

			if directionIndex >= len(animation.directions) {
				animation.directions = append(animation.directions, new(animationDirection))
			}

			direction := animation.directions[directionIndex]
			direction.frames = append(direction.frames, &animationFrame{
				width:   dccFrame.Width,
				height:  dccFrame.Height,
				offsetX: minX,
				offsetY: minY,
				image:   sfc,
			})
		}
	}

	return animation, nil
}

// CreateAnimationFromDC6 creates an Animation from d2dc6.DC6 and d2dat.DATPalette
func CreateAnimationFromDC6(dc6 *d2dc6.DC6, palette *d2dat.DATPalette) (*Animation, error) {
	animation := &Animation{
		playLength:     defaultPlayLength,
		playLoop:       true,
		originAtBottom: true,
	}

	for frameIndex, dc6Frame := range dc6.Frames {
		sfc, err := d2render.NewSurface(int(dc6Frame.Width), int(dc6Frame.Height), d2interface.FilterNearest)
		if err != nil {
			return nil, err
		}

		indexData := make([]int, dc6Frame.Width*dc6Frame.Height)
		for i := range indexData {
			indexData[i] = -1
		}

		x := 0
		y := int(dc6Frame.Height) - 1
		offset := 0

		for {
			b := int(dc6Frame.FrameData[offset])
			offset++

			if b == 0x80 {
				if y == 0 {
					break
				}
				y--
				x = 0
			} else if b&0x80 > 0 {
				transparentPixels := b & 0x7f
				for i := 0; i < transparentPixels; i++ {
					indexData[x+y*int(dc6Frame.Width)+i] = -1
				}
				x += transparentPixels
			} else {
				for i := 0; i < b; i++ {
					indexData[x+y*int(dc6Frame.Width)+i] = int(dc6Frame.FrameData[offset])
					offset++
				}
				x += b
			}
		}

		bytesPerPixel := 4
		colorData := make([]byte, int(dc6Frame.Width)*int(dc6Frame.Height)*bytesPerPixel)

		for i := 0; i < int(dc6Frame.Width*dc6Frame.Height); i++ {
			if indexData[i] < 1 { // TODO: Is this == -1 or < 1?
				continue
			}

			colorData[i*bytesPerPixel] = palette.Colors[indexData[i]].R
			colorData[i*bytesPerPixel+1] = palette.Colors[indexData[i]].G
			colorData[i*bytesPerPixel+2] = palette.Colors[indexData[i]].B
			colorData[i*bytesPerPixel+3] = 0xff
		}

		if err := sfc.ReplacePixels(colorData); err != nil {
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
			image:   sfc,
		})
	}

	return animation, nil
}

func (a *Animation) Clone() *Animation {
	animation := *a
	return &animation
}

func (a *Animation) SetSubLoop(startFrame, EndFrame int) {
	a.subStartingFrame = startFrame
	a.subEndingFrame = EndFrame
	a.hasSubLoop = true
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
		startIndex := 0
		endIndex := frameCount
		if a.hasSubLoop && a.playedCount > 0 {
			startIndex = a.subStartingFrame
			endIndex = a.subEndingFrame
		}

		switch a.playMode {
		case playModeForward:
			a.frameIndex++
			if a.frameIndex >= endIndex {
				a.playedCount++
				if a.playLoop {
					a.frameIndex = startIndex
				} else {
					a.frameIndex = endIndex - 1
					break
				}
			}
		case playModeBackward:
			a.frameIndex--
			if a.frameIndex < startIndex {
				a.playedCount++
				if a.playLoop {
					a.frameIndex = endIndex - 1
				} else {
					a.frameIndex = startIndex
					break
				}
			}
		}
	}

	return nil
}

func (a *Animation) Render(target d2interface.Surface) error {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushTranslation(frame.offsetX, frame.offsetY)
	defer target.Pop()

	target.PushCompositeMode(a.compositeMode)
	defer target.Pop()

	target.PushColor(a.colorMod)
	defer target.Pop()

	return target.Render(frame.image)
}

func (a *Animation) RenderFromOrigin(target d2interface.Surface) error {
	if a.originAtBottom {
		direction := a.directions[a.directionIndex]
		frame := direction.frames[a.frameIndex]
		target.PushTranslation(0, -frame.height)
		defer target.Pop()
	}

	return a.Render(target)
}

// RenderSection renders the section of the animation frame enclosed by bounds
func (a *Animation) RenderSection(sfc d2interface.Surface, bound image.Rectangle) error {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	sfc.PushTranslation(frame.offsetX, frame.offsetY)
	sfc.PushCompositeMode(a.compositeMode)
	sfc.PushColor(a.colorMod)
	defer sfc.PopN(3)
	return sfc.RenderSection(frame.image, bound)
}

// GetFrameSize gets the Size(width, height) of a indexed frame.
func (a *Animation) GetFrameSize(frameIndex int) (int, int, error) {
	direction := a.directions[a.directionIndex]
	if frameIndex >= len(direction.frames) {
		return 0, 0, errors.New("invalid frame index")
	}

	frame := direction.frames[frameIndex]
	return frame.width, frame.height, nil
}

// GetCurrentFrameSize gets the Size(width, height) of the current frame.
func (a *Animation) GetCurrentFrameSize() (int, int) {
	width, height, _ := a.GetFrameSize(a.frameIndex)
	return width, height
}

// GetFrameBounds gets maximum Size(width, height) of all frame.
func (a *Animation) GetFrameBounds() (int, int) {
	maxWidth, maxHeight := 0, 0

	direction := a.directions[a.directionIndex]
	for _, frame := range direction.frames {
		maxWidth = d2common.MaxInt(maxWidth, frame.width)
		maxHeight = d2common.MaxInt(maxHeight, frame.height)
	}

	return maxWidth, maxHeight
}

// GetCurrentFrame gets index of current frame in animation
func (a *Animation) GetCurrentFrame() int {
	return a.frameIndex
}

// GetFrameCount gets number of frames in animation
func (a *Animation) GetFrameCount() int {
	direction := a.directions[a.directionIndex]
	return len(direction.frames)
}

// IsOnFirstFrame gets if the animation on its first frame
func (a *Animation) IsOnFirstFrame() bool {
	return a.frameIndex == 0
}

// IsOnLastFrame gets if the animation on its last frame
func (a *Animation) IsOnLastFrame() bool {
	return a.frameIndex == a.GetFrameCount()-1
}

// GetDirectionCount gets the number of animation direction
func (a *Animation) GetDirectionCount() int {
	return len(a.directions)
}

// SetDirection places the animation in the direction of an animation
func (a *Animation) SetDirection(directionIndex int) error {
	const smallestInvalidDirectionIndex = 64
	if directionIndex >= smallestInvalidDirectionIndex {
		return errors.New("invalid direction index")
	}

	a.directionIndex = d2dcc.Dir64ToDcc(directionIndex, len(a.directions))
	a.frameIndex = 0

	return nil
}

// GetDirection get the current animation direction
func (a *Animation) GetDirection() int {
	return a.directionIndex
}

// SetCurrentFrame sets animation at a specific frame
func (a *Animation) SetCurrentFrame(frameIndex int) error {
	if frameIndex >= a.GetFrameCount() {
		return errors.New("invalid frame index")
	}

	a.frameIndex = frameIndex
	a.lastFrameTime = 0

	return nil
}

// Rewind animation to beginning
func (a *Animation) Rewind() {
	a.SetCurrentFrame(0)
}

// PlayForward plays animation forward
func (a *Animation) PlayForward() {
	a.playMode = playModeForward
	a.lastFrameTime = 0
}

// PlayBackward plays animation backward
func (a *Animation) PlayBackward() {
	a.playMode = playModeBackward
	a.lastFrameTime = 0
}

// Pause animation
func (a *Animation) Pause() {
	a.playMode = playModePause
	a.lastFrameTime = 0
}

// SetPlayLoop sets whether to loop the animation
func (a *Animation) SetPlayLoop(loop bool) {
	a.playLoop = loop
}

// SetPlaySpeed sets play speed of the animation
func (a *Animation) SetPlaySpeed(playSpeed float64) {
	a.SetPlayLength(playSpeed * float64(a.GetFrameCount()))
}

// SetPlayLength sets the Animation's play length in seconds
func (a *Animation) SetPlayLength(playLength float64) { // TODO refactor to use time.Duration instead of float64
	a.playLength = playLength
	a.lastFrameTime = 0
}

// SetPlayLengthMs sets the Animation's play length in milliseconds
func (a *Animation) SetPlayLengthMs(playLengthMs int) { // TODO remove this method
	const millisecondsPerSecond = 1000.0
	a.SetPlayLength(float64(playLengthMs) / millisecondsPerSecond)
}

// SetColorMod sets the Animation's color mod
func (a *Animation) SetColorMod(colorMod color.Color) {
	a.colorMod = colorMod
}

// GetPlayedCount gets the number of times the application played
func (a *Animation) GetPlayedCount() int {
	return a.playedCount
}

// ResetPlayedCount resets the play count
func (a *Animation) ResetPlayedCount() {
	a.playedCount = 0
}

// SetBlend sets the Animation alpha blending status
func (a *Animation) SetBlend(blend bool) {
	if blend {
		a.compositeMode = d2enum.CompositeModeLighter
	} else {
		a.compositeMode = d2enum.CompositeModeSourceOver
	}
}
