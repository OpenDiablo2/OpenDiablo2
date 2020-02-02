package d2ui

import (
	"errors"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Sprite struct {
	x         int
	y         int
	animation *d2asset.Animation
}

var (
	ErrNoAnimation = errors.New("no animation was specified")
)

func LoadSprite(animation *d2asset.Animation) (*Sprite, error) {
	if animation == nil {
		return nil, ErrNoAnimation
	}
	return &Sprite{animation: animation}, nil
}

func (s *Sprite) Render(target d2render.Surface) error {
	_, frameHeight := s.animation.GetCurrentFrameSize()

	target.PushTranslation(s.x, s.y-frameHeight)
	defer target.Pop()
	return s.animation.Render(target)
}

func (s *Sprite) RenderSegmented(target d2render.Surface, segmentsX, segmentsY, frameOffset int) error {
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

func (s *Sprite) SetPosition(x, y int) {
	s.x = x
	s.y = y
}

func (s *Sprite) GetPosition() (int, int) {
	return s.x, s.y
}

func (s *Sprite) GetFrameSize(frameIndex int) (int, int, error) {
	return s.animation.GetFrameSize(frameIndex)
}

func (s *Sprite) GetCurrentFrameSize() (int, int) {
	return s.animation.GetCurrentFrameSize()
}

func (s *Sprite) GetFrameBounds() (int, int) {
	return s.animation.GetFrameBounds()
}

func (s *Sprite) GetCurrentFrame() int {
	return s.animation.GetCurrentFrame()
}

func (s *Sprite) GetFrameCount() int {
	return s.animation.GetFrameCount()
}

func (s *Sprite) IsOnFirstFrame() bool {
	return s.animation.IsOnFirstFrame()
}

func (s *Sprite) IsOnLastFrame() bool {
	return s.animation.IsOnLastFrame()
}

func (s *Sprite) GetDirectionCount() int {
	return s.animation.GetDirectionCount()
}

func (s *Sprite) SetDirection(directionIndex int) error {
	return s.animation.SetDirection(directionIndex)
}

func (s *Sprite) GetDirection() int {
	return s.animation.GetDirection()
}

func (s *Sprite) SetCurrentFrame(frameIndex int) error {
	return s.animation.SetCurrentFrame(frameIndex)
}

func (s *Sprite) Rewind() {
	s.animation.SetCurrentFrame(0)
}

func (s *Sprite) PlayForward() {
	s.animation.PlayForward()
}

func (s *Sprite) PlayBackward() {
	s.animation.PlayBackward()
}

func (s *Sprite) Pause() {
	s.animation.Pause()
}

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

func (s *Sprite) SetBlend(blend bool) {
	s.animation.SetBlend(blend)
}

func (s *Sprite) Advance(elapsed float64) error {
	return s.animation.Advance(elapsed)
}
