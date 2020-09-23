package d2ui

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// Sprite is a positioned visual object.
type Sprite struct {
	x         int
	y         int
	animation d2interface.Animation
}

const (
	errNoAnimation = "no animation was specified"
)

// NewSprite creates a new Sprite
func (ui *UIManager) NewSprite(animationPath, palettePath string) (*Sprite, error) {
	animation, err := ui.asset.LoadAnimation(animationPath, palettePath)
	if animation == nil || err != nil {
		return nil, fmt.Errorf(errNoAnimation)
	}

	err = animation.BindRenderer(ui.renderer)
	if err != nil {
		return nil, err
	}

	return &Sprite{animation: animation}, nil
}

// Render renders the sprite on the given surface
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

// RenderSegmented renders a sprite that is internally segmented as frames
func (s *Sprite) RenderSegmented(target d2interface.Surface, segmentsX, segmentsY, frameOffset int) error {
	var currentY int

	for y := 0; y < segmentsY; y++ {
		var currentX, maxFrameHeight int

		for x := 0; x < segmentsX; x++ {
			idx := x + y*segmentsX + frameOffset*segmentsX*segmentsY
			if err := s.animation.SetCurrentFrame(idx); err != nil {
				return err
			}

			target.PushTranslation(s.x+currentX, s.y+currentY)
			err := s.animation.Render(target)
			target.Pop()

			if err != nil {
				return err
			}

			frameWidth, frameHeight := s.GetCurrentFrameSize()
			maxFrameHeight = d2math.MaxInt(maxFrameHeight, frameHeight)
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
func (s *Sprite) GetPosition() (x, y int) {
	return s.x, s.y
}

// GetFrameSize gets the Size(width, height) of a indexed frame.
func (s *Sprite) GetFrameSize(frameIndex int) (x, y int, err error) {
	return s.animation.GetFrameSize(frameIndex)
}

// GetCurrentFrameSize gets the Size(width, height) of the current frame.
func (s *Sprite) GetCurrentFrameSize() (width, height int) {
	return s.animation.GetCurrentFrameSize()
}

// GetFrameBounds gets maximum Size(width, height) of all frame.
func (s *Sprite) GetFrameBounds() (width, height int) {
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
	err := s.animation.SetCurrentFrame(0)
	if err != nil {
		log.Print(err)
	}
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

// SetPlayLength sets the play length of the sprite animation
func (s *Sprite) SetPlayLength(playLength float64) {
	s.animation.SetPlayLength(playLength)
}

// SetPlayLengthMs sets the play length of the sprite animation in milliseconds
func (s *Sprite) SetPlayLengthMs(playLengthMs int) {
	s.animation.SetPlayLengthMs(playLengthMs)
}

// SetColorMod sets the color modifier
func (s *Sprite) SetColorMod(c color.Color) {
	s.animation.SetColorMod(c)
}

// Advance advances the animation
func (s *Sprite) Advance(elapsed float64) error {
	return s.animation.Advance(elapsed)
}

// SetEffect sets the draw effect type
func (s *Sprite) SetEffect(e d2enum.DrawEffect) {
	s.animation.SetEffect(e)
}
