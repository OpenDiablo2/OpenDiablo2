package d2asset

import (
	"errors"
	"image"
	"image/color"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	d2iface "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
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

	image d2iface.Surface
}

type animationDirection struct {
	decoded bool
	frames  []*animationFrame
}

// animation has directionality, play modes, and frame counting
type animation struct {
	directions       []animationDirection
	effect           d2enum.DrawEffect
	colorMod         color.Color
	frameIndex       int
	directionIndex   int
	lastFrameTime    float64
	playedCount      int
	playMode         playMode
	playLength       float64
	subStartingFrame int
	subEndingFrame   int
	originAtBottom   bool
	playLoop         bool
	hasSubLoop       bool // runs after first animation ends
}

// SetSubLoop sets a sub loop for the animation
func (a *animation) SetSubLoop(startFrame, endFrame int) {
	a.subStartingFrame = startFrame
	a.subEndingFrame = endFrame
	a.hasSubLoop = true
}

// Advance advances the animation state
func (a *animation) Advance(elapsed float64) error {
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

func (a *animation) renderShadow(target d2iface.Surface) error {
	//_, height := a.GetFrameBounds()
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]
	target.PushFilter(d2enum.FilterLinear)
	target.PushTranslation(
		frame.offsetX,
		int(float64(frame.offsetY)*0.5))
	target.PushScale(1.0, 0.5)
	target.PushSkew(0.5, 0.0)
	target.PushEffect(d2enum.DrawEffectPctTransparency25)
	target.PushBrightness(0.0)

	defer target.PopN(6)

	return target.Render(frame.image)
}

// Render renders the animation to the given surface
func (a *animation) Render(target d2iface.Surface) error {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushTranslation(frame.offsetX, frame.offsetY)
	defer target.Pop()

	target.PushEffect(a.effect)
	defer target.Pop()

	target.PushColor(a.colorMod)
	defer target.Pop()

	return target.Render(frame.image)
}

// RenderFromOrigin renders the animation from the animation origin
func (a *animation) RenderFromOrigin(target d2iface.Surface, shadow bool) error {
	if a.originAtBottom {
		direction := a.directions[a.directionIndex]
		frame := direction.frames[a.frameIndex]
		target.PushTranslation(0, -frame.height)

		defer target.Pop()
	}

	if shadow {
		_, height := a.GetFrameBounds()
		height = int(math.Abs(float64(height)))
		target.PushTranslation((-height / 2), 0)
		defer target.Pop()
		return a.renderShadow(target)
	}

	return a.Render(target)
}

// RenderSection renders the section of the animation frame enclosed by bounds
func (a *animation) RenderSection(sfc d2iface.Surface, bound image.Rectangle) error {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	sfc.PushTranslation(frame.offsetX, frame.offsetY)
	sfc.PushEffect(a.effect)
	sfc.PushColor(a.colorMod)

	defer sfc.PopN(3)

	return sfc.RenderSection(frame.image, bound)
}

// GetFrameSize gets the Size(width, height) of a indexed frame.
func (a *animation) GetFrameSize(frameIndex int) (width, height int, err error) {
	direction := a.directions[a.directionIndex]
	if frameIndex >= len(direction.frames) {
		return 0, 0, errors.New("invalid frame index")
	}

	frame := direction.frames[frameIndex]

	return frame.width, frame.height, nil
}

// GetCurrentFrameSize gets the Size(width, height) of the current frame.
func (a *animation) GetCurrentFrameSize() (width, height int) {
	width, height, _ = a.GetFrameSize(a.frameIndex)
	return width, height
}

// GetFrameBounds gets maximum Size(width, height) of all frame.
func (a *animation) GetFrameBounds() (maxWidth, maxHeight int) {
	maxWidth, maxHeight = 0, 0

	direction := a.directions[a.directionIndex]
	for _, frame := range direction.frames {
		maxWidth = d2common.MaxInt(maxWidth, frame.width)
		maxHeight = d2common.MaxInt(maxHeight, frame.height)
	}

	return maxWidth, maxHeight
}

// GetCurrentFrame gets index of current frame in animation
func (a *animation) GetCurrentFrame() int {
	return a.frameIndex
}

// GetFrameCount gets number of frames in animation
func (a *animation) GetFrameCount() int {
	direction := a.directions[a.directionIndex]
	return len(direction.frames)
}

// IsOnFirstFrame gets if the animation on its first frame
func (a *animation) IsOnFirstFrame() bool {
	return a.frameIndex == 0
}

// IsOnLastFrame gets if the animation on its last frame
func (a *animation) IsOnLastFrame() bool {
	return a.frameIndex == a.GetFrameCount()-1
}

// GetDirectionCount gets the number of animation direction
func (a *animation) GetDirectionCount() int {
	return len(a.directions)
}

// SetDirection places the animation in the direction of an animation
func (a *animation) SetDirection(directionIndex int) error {
	const smallestInvalidDirectionIndex = 64
	if directionIndex >= smallestInvalidDirectionIndex {
		return errors.New("invalid direction index")
	}

	a.directionIndex = d2dcc.Dir64ToDcc(directionIndex, len(a.directions))
	a.frameIndex = 0

	return nil
}

// GetDirection get the current animation direction
func (a *animation) GetDirection() int {
	return a.directionIndex
}

// SetCurrentFrame sets animation at a specific frame
func (a *animation) SetCurrentFrame(frameIndex int) error {
	if frameIndex >= a.GetFrameCount() {
		return errors.New("invalid frame index")
	}

	a.frameIndex = frameIndex
	a.lastFrameTime = 0

	return nil
}

// Rewind animation to beginning
func (a *animation) Rewind() {
	_ = a.SetCurrentFrame(0)
}

// PlayForward plays animation forward
func (a *animation) PlayForward() {
	a.playMode = playModeForward
	a.lastFrameTime = 0
}

// PlayBackward plays animation backward
func (a *animation) PlayBackward() {
	a.playMode = playModeBackward
	a.lastFrameTime = 0
}

// Pause animation
func (a *animation) Pause() {
	a.playMode = playModePause
	a.lastFrameTime = 0
}

// SetPlayLoop sets whether to loop the animation
func (a *animation) SetPlayLoop(loop bool) {
	a.playLoop = loop
}

// SetPlaySpeed sets play speed of the animation
func (a *animation) SetPlaySpeed(playSpeed float64) {
	a.SetPlayLength(playSpeed * float64(a.GetFrameCount()))
}

// SetPlayLength sets the Animation's play length in seconds
func (a *animation) SetPlayLength(playLength float64) {
	// TODO refactor to use time.Duration instead of float64
	a.playLength = playLength
	a.lastFrameTime = 0
}

// SetPlayLengthMs sets the Animation's play length in milliseconds
func (a *animation) SetPlayLengthMs(playLengthMs int) {
	// TODO remove this method
	const millisecondsPerSecond = 1000.0

	a.SetPlayLength(float64(playLengthMs) / millisecondsPerSecond)
}

// SetColorMod sets the Animation's color mod
func (a *animation) SetColorMod(colorMod color.Color) {
	a.colorMod = colorMod
}

// GetPlayedCount gets the number of times the application played
func (a *animation) GetPlayedCount() int {
	return a.playedCount
}

// ResetPlayedCount resets the play count
func (a *animation) ResetPlayedCount() {
	a.playedCount = 0
}

func (a *animation) SetEffect(e d2enum.DrawEffect) {
	a.effect = e
}
