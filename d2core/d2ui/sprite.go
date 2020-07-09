package d2ui

import (
	"errors"
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// Sprite is a positioned visual object.
type Sprite struct {
	x         int
	y         int
	animation d2interface.Animation
}

var (
	ErrNoAnimation = errors.New("no animation was specified")
)

func LoadSprite(animation d2interface.Animation) (*Sprite, error) {
	if animation == nil {
		return nil, ErrNoAnimation
	}
	return &Sprite{animation: animation}, nil
}

func (s *Sprite) Render(target d2interface.Surface) error {
	_, frameHeight := s.animation.GetCurrentFrameSize()

	target.PushTranslation(s.x, s.y-frameHeight)
	defer target.Pop()

	return s.animation.Render(target)
}

// RenderSection renders the section of the sprite enclosed by bounds
func (s *Sprite) RenderSection(sfc d2interface.Surface, bound image.Rectangle) error {
	sfc.PushTranslation(s.x, s.y-bound.Dy())
	defer sfc.Pop()

	return s.animation.RenderSection(sfc, bound)
}

func (s *Sprite) RenderSegmented(target d2interface.Surface, segmentsX, segmentsY, frameOffset int) error {
	var currentY int

	for y := 0; y < segmentsY; y++ {
		var currentX int
		var maxFrameHeight int

		for x := 0; x < segmentsX; x++ {
			if err := s.animation.SetCurrentFrame(x + y*segmentsX + frameOffset*segmentsX*segmentsY); err != nil {
				return err
			}

			target.PushTranslation(s.x+currentX, s.y+currentY)
			err := s.animation.Render(target)
			target.Pop()

			if err != nil {
				return err
			}

			frameWidth, frameHeight := s.GetCurrentFrameSize()
			maxFrameHeight = d2common.MaxInt(maxFrameHeight, frameHeight)
			currentX += frameWidth
		}

		currentY += maxFrameHeight
	}

	return nil
}

// SetPosition places the sprite in 2D
func (s *Sprite) SetPosition(x, y int) {
	s.x = x
	s.y = y
}

// GetPosition retrieves the 2D position of the sprite
func (s *Sprite) GetPosition() (int, int) {
	return s.x, s.y
}

// GetFrameSize gets the Size(width, height) of a indexed frame.
func (s *Sprite) GetFrameSize(frameIndex int) (int, int, error) {
	return s.animation.GetFrameSize(frameIndex)
}

// GetCurrentFrameSize gets the Size(width, height) of the current frame.
func (s *Sprite) GetCurrentFrameSize() (int, int) {
	return s.animation.GetCurrentFrameSize()
}

// GetFrameBounds gets maximum Size(width, height) of all frame.
func (s *Sprite) GetFrameBounds() (int, int) {
	return s.animation.GetFrameBounds()
}

// GetCurrentFrame gets index of current frame in animation
func (s *Sprite) GetCurrentFrame() int {
	return s.animation.GetCurrentFrame()
}

// GetFrameCount gets number of frames in animation
func (s *Sprite) GetFrameCount() int {
	return s.animation.GetFrameCount()
}

// IsOnFirstFrame gets if the animation on its first frame
func (s *Sprite) IsOnFirstFrame() bool {
	return s.animation.IsOnFirstFrame()
}

// IsOnLastFrame gets if the animation on its last frame
func (s *Sprite) IsOnLastFrame() bool {
	return s.animation.IsOnLastFrame()
}

// GetDirectionCount gets the number of animation direction
func (s *Sprite) GetDirectionCount() int {
	return s.animation.GetDirectionCount()
}

// SetDirection places the animation in the direction of an animation
func (s *Sprite) SetDirection(directionIndex int) error {
	return s.animation.SetDirection(directionIndex)
}

// GetDirection get the current animation direction
func (s *Sprite) GetDirection() int {
	return s.animation.GetDirection()
}

// SetCurrentFrame sets animation at a specific frame
func (s *Sprite) SetCurrentFrame(frameIndex int) error {
	return s.animation.SetCurrentFrame(frameIndex)
}

// Rewind sprite to beginning
func (s *Sprite) Rewind() {
	s.animation.SetCurrentFrame(0)
}

// PlayForward plays sprite forward
func (s *Sprite) PlayForward() {
	s.animation.PlayForward()
}

// PlayBackward play sprites backward
func (s *Sprite) PlayBackward() {
	s.animation.PlayBackward()
}

// Pause animation
func (s *Sprite) Pause() {
	s.animation.Pause()
}

// SetPlayLoop sets whether to loop the animation
func (s *Sprite) SetPlayLoop(loop bool) {
	s.animation.SetPlayLoop(loop)
}

func (s *Sprite) SetPlayLength(playLength float64) {
	s.animation.SetPlayLength(playLength)
}

func (s *Sprite) SetPlayLengthMs(playLengthMs int) {
	s.animation.SetPlayLengthMs(playLengthMs)
}

func (s *Sprite) SetColorMod(color color.Color) {
	s.animation.SetColorMod(color)
}

func (s *Sprite) Advance(elapsed float64) error {
	return s.animation.Advance(elapsed)
}

func (s *Sprite) SetEffect(e d2enum.DrawEffect) {
	s.animation.SetEffect(e)
}
