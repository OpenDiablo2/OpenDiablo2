package d2sprite

import (
	"errors"
	"image"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

type playMode int

const (
	playModePause playMode = iota
	playModeForward
	playModeBackward
)

const fps25 = 25

type spriteFrame struct {
	decoded bool

	width   int
	height  int
	offsetX int
	offsetY int

	image d2interface.Surface
}

type spriteDirection struct {
	decoded bool
	frames  []spriteFrame
}

// static check that we implement the sprite interface
var _ d2interface.Sprite = &Sprite{}

// Sprite has directionality, play modes, and frame counting
type Sprite struct {
	renderer         d2interface.Renderer
	onBindRenderer   func(renderer d2interface.Renderer) error
	directions       []spriteDirection
	effect           d2enum.DrawEffect
	colorMod         color.Color
	frameIndex       int
	directionIndex   int
	lastFrameTime    time.Duration
	playedCount      int
	playMode         playMode
	playLength       time.Duration // https://github.com/OpenDiablo2/OpenDiablo2/issues/813
	subStartingFrame int
	subEndingFrame   int
	originAtBottom   bool
	playLoop         bool
	hasSubLoop       bool // runs after first sprite ends
	hasShadow        bool
}

// SetSubLoop sets a sub loop for the sprite
func (a *Sprite) SetSubLoop(startFrame, endFrame int) {
	a.subStartingFrame = startFrame
	a.subEndingFrame = endFrame
	a.hasSubLoop = true
}

// Advance advances the sprite state
func (a *Sprite) Advance(elapsed time.Duration) error {
	if a.playMode == playModePause {
		return nil
	}

	frameCount := a.GetFrameCount()
	frameLength := time.Duration(float64(a.playLength) / float64(frameCount))

	a.lastFrameTime += elapsed

	framesAdvanced := int(float64(a.lastFrameTime) / float64(frameLength))
	a.lastFrameTime -= time.Duration(float64(framesAdvanced)) * frameLength

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

const (
	full = 1.0
	half = 0.5
	zero = 0.0
)

func (a *Sprite) renderShadow(target d2interface.Surface) {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushFilter(d2enum.FilterLinear)
	defer target.Pop()

	target.PushTranslation(frame.offsetX, int(float64(frame.offsetY)*half))
	defer target.Pop()

	target.PushScale(full, half)
	defer target.Pop()

	target.PushSkew(half, zero)
	defer target.Pop()

	target.PushEffect(d2enum.DrawEffectPctTransparency25)
	defer target.Pop()

	target.PushBrightness(zero)
	defer target.Pop()

	target.Render(frame.image)
}

// GetCurrentFrameSurface returns the surface for the current frame of the
// sprite
func (a *Sprite) GetCurrentFrameSurface() d2interface.Surface {
	return a.directions[a.directionIndex].frames[a.frameIndex].image
}

// Render renders the sprite to the given surface
func (a *Sprite) Render(target d2interface.Surface) {
	if a.renderer == nil {
		a.BindRenderer(target.Renderer())
	}

	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushTranslation(frame.offsetX, frame.offsetY)
	defer target.Pop()

	target.PushEffect(a.effect)
	defer target.Pop()

	target.PushColor(a.colorMod)
	defer target.Pop()

	target.Render(frame.image)
}

// BindRenderer binds the given renderer to the sprite so that it can initialize
// the required surfaces
func (a *Sprite) BindRenderer(r d2interface.Renderer) {
	if a.onBindRenderer == nil {
		return
	}

	if err := a.onBindRenderer(r); err != nil {
		log.Println(err)
	}
}

// RenderFromOrigin renders the sprite from the sprite origin
func (a *Sprite) RenderFromOrigin(target d2interface.Surface, shadow bool) {
	if a.renderer == nil {
		a.BindRenderer(target.Renderer())
	}

	if a.originAtBottom {
		direction := a.directions[a.directionIndex]
		frame := direction.frames[a.frameIndex]
		target.PushTranslation(0, -frame.height)

		defer target.Pop()
	}

	if shadow && !a.effect.Transparent() && a.hasShadow {
		_, height := a.GetFrameBounds()
		height = int(math.Abs(float64(height)))
		halfHeight := height / 2 //nolint:gomnd // this ain't rocket surgery...

		target.PushTranslation(-halfHeight, 0)
		defer target.Pop()

		a.renderShadow(target)

		return
	}

	a.Render(target)
}

// RenderSection renders the section of the sprite frame enclosed by bounds
func (a *Sprite) RenderSection(target d2interface.Surface, bound image.Rectangle) {
	if a.renderer == nil {
		a.BindRenderer(target.Renderer())
	}

	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushTranslation(frame.offsetX, frame.offsetY)
	defer target.Pop()

	target.PushEffect(a.effect)
	defer target.Pop()

	target.PushColor(a.colorMod)
	defer target.Pop()

	target.RenderSection(frame.image, bound)
}

// GetFrameSize gets the Size(width, height) of a indexed frame.
func (a *Sprite) GetFrameSize(frameIndex int) (width, height int, err error) {
	direction := a.directions[a.directionIndex]
	if frameIndex >= len(direction.frames) {
		return 0, 0, errors.New("invalid frame index")
	}

	frame := direction.frames[frameIndex]

	return frame.width, frame.height, nil
}

// GetCurrentFrameSize gets the Size(width, height) of the current frame.
func (a *Sprite) GetCurrentFrameSize() (width, height int) {
	width, height, err := a.GetFrameSize(a.frameIndex)
	if err != nil {
		log.Print(err)
	}

	return width, height
}

func (a *Sprite) GetCurrentFrameOffset() (int, int) {
	f := a.directions[a.directionIndex].frames[a.frameIndex]
	return f.offsetX, f.offsetY
}

// GetFrameBounds gets maximum Size(width, height) of all frame.
func (a *Sprite) GetFrameBounds() (maxWidth, maxHeight int) {
	maxWidth, maxHeight = 0, 0

	direction := a.directions[a.directionIndex]

	for _, frame := range direction.frames {
		maxWidth = d2math.MaxInt(maxWidth, frame.width)
		maxHeight = d2math.MaxInt(maxHeight, frame.height)
	}

	return maxWidth, maxHeight
}

// GetCurrentFrame gets index of current frame in sprite
func (a *Sprite) GetCurrentFrame() int {
	return a.frameIndex
}

// GetFrameCount gets number of frames in sprite
func (a *Sprite) GetFrameCount() int {
	direction := a.directions[a.directionIndex]
	return len(direction.frames)
}

// IsOnFirstFrame gets if the sprite on its first frame
func (a *Sprite) IsOnFirstFrame() bool {
	return a.frameIndex == 0
}

// IsOnLastFrame gets if the sprite on its last frame
func (a *Sprite) IsOnLastFrame() bool {
	return a.frameIndex == a.GetFrameCount()-1
}

// GetDirectionCount gets the number of sprite direction
func (a *Sprite) GetDirectionCount() int {
	return len(a.directions)
}

// SetDirection places the sprite in the direction of an sprite
func (a *Sprite) SetDirection(directionIndex int) error {
	const smallestInvalidDirectionIndex = 64
	if directionIndex >= smallestInvalidDirectionIndex {
		return errors.New("invalid direction index")
	}

	a.directionIndex = d2dcc.Dir64ToDcc(directionIndex, len(a.directions))
	a.frameIndex = 0

	return nil
}

// GetDirection get the current sprite direction
func (a *Sprite) GetDirection() int {
	return a.directionIndex
}

// SetCurrentFrame sets sprite at a specific frame
func (a *Sprite) SetCurrentFrame(frameIndex int) error {
	if frameIndex >= a.GetFrameCount() {
		return errors.New("invalid frame index")
	}

	a.frameIndex = frameIndex
	a.lastFrameTime = 0

	return nil
}

// Rewind sprite to beginning
func (a *Sprite) Rewind() {
	err := a.SetCurrentFrame(0)
	if err != nil {
		log.Print(err)
	}
}

// PlayForward plays sprite forward
func (a *Sprite) PlayForward() {
	a.playMode = playModeForward
	a.lastFrameTime = 0
}

// PlayBackward plays sprite backward
func (a *Sprite) PlayBackward() {
	a.playMode = playModeBackward
	a.lastFrameTime = 0
}

// Pause sprite
func (a *Sprite) Pause() {
	a.playMode = playModePause
	a.lastFrameTime = 0
}

// SetPlayLoop sets whether to loop the sprite
func (a *Sprite) SetPlayLoop(loop bool) {
	a.playLoop = loop
}

// SetPlayFPS sets play speed of the sprite
func (a *Sprite) SetPlayFPS(fps float64) {
	a.SetPlayLength(time.Duration(float64(time.Second) / fps * float64(a.GetFrameCount())))
}

// SetPlaySpeed sets play speed of the sprite
func (a *Sprite) SetPlaySpeed(playSpeed time.Duration) {
	frameDuration := time.Duration(float64(a.GetFrameCount()) / float64(playSpeed))
	a.SetPlayLength(frameDuration)
}

// SetPlayLength sets the Sprite's play length in seconds
func (a *Sprite) SetPlayLength(playLength time.Duration) {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/813
	a.playLength = playLength
	a.lastFrameTime = 0
}

// SetColorMod sets the Sprite's color mod
func (a *Sprite) SetColorMod(colorMod color.Color) {
	a.colorMod = colorMod
}

// GetPlayedCount gets the number of times the application played
func (a *Sprite) GetPlayedCount() int {
	return a.playedCount
}

// ResetPlayedCount resets the play count
func (a *Sprite) ResetPlayedCount() {
	a.playedCount = 0
}

// SetEffect sets the draw effect for the sprite
func (a *Sprite) SetEffect(e d2enum.DrawEffect) {
	a.effect = e
}

// SetShadow sets bool for whether or not to draw a shadow
func (a *Sprite) SetShadow(shadow bool) {
	a.hasShadow = shadow
}

// Clone creates a copy of the Sprite
func (a *Sprite) Clone() d2interface.Sprite {
	clone := *a
	copy(clone.directions, a.directions)

	return &clone
}
