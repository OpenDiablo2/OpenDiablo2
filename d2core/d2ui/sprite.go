package d2ui

import (
	"fmt"
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

// Sprite is a positioned visual object.
type Sprite struct {
	*BaseWidget
	animation d2interface.Animation

	*d2util.Logger
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

	animation.BindRenderer(ui.renderer)

	base := NewBaseWidget(ui)

	return &Sprite{
		BaseWidget: base,
		animation:  animation,
		Logger:     ui.Logger,
	}, nil
}

// Render renders the sprite on the given surface
func (s *Sprite) Render(target d2interface.Surface) {
	_, frameHeight := s.animation.GetCurrentFrameSize()

	target.PushTranslation(s.x, s.y-frameHeight)
	defer target.Pop()

	s.animation.Render(target)
}

// GetSurface returns the surface of the sprite at the given frame
func (s *Sprite) GetSurface() d2interface.Surface {
	return s.animation.GetCurrentFrameSurface()
}

// RenderSection renders the section of the sprite enclosed by bounds
func (s *Sprite) RenderSection(sfc d2interface.Surface, bound image.Rectangle) {
	sfc.PushTranslation(s.x, s.y-bound.Dy())
	defer sfc.Pop()

	s.animation.RenderSection(sfc, bound)
}

// RenderSegmented renders a sprite that is internally segmented as frames
func (s *Sprite) RenderSegmented(target d2interface.Surface, segmentsX, segmentsY, frameOffset int) {
	var currentY int

	for y := 0; y < segmentsY; y++ {
		var currentX, maxFrameHeight int

		for x := 0; x < segmentsX; x++ {
			idx := x + y*segmentsX + frameOffset*segmentsX*segmentsY
			if err := s.animation.SetCurrentFrame(idx); err != nil {
				s.Error("SetCurrentFrame error" + err.Error())
			}

			target.PushTranslation(s.x+currentX, s.y+currentY)
			s.animation.Render(target)
			target.Pop()

			frameWidth, frameHeight := s.GetCurrentFrameSize()
			maxFrameHeight = d2math.MaxInt(maxFrameHeight, frameHeight)
			currentX += frameWidth
		}

		currentY += maxFrameHeight
	}
}

// GetSize returns the size of the current frame
func (s *Sprite) GetSize() (width, height int) {
	return s.GetCurrentFrameSize()
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
		s.Error(err.Error())
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
