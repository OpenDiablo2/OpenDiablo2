package d2render

import (
	"image/color"

	"github.com/OpenDiablo2/D2Shared/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2asset"

	"github.com/hajimehoshi/ebiten"
)

type Sprite struct {
	x             int
	y             int
	lastFrameTime float64
	animation     *d2asset.Animation
}

func LoadSprite(animationPath, palettePath string) (*Sprite, error) {
	animation, err := d2asset.LoadAnimation(animationPath, palettePath)
	if err != nil {
		return nil, err
	}

	return &Sprite{lastFrameTime: d2helper.Now(), animation: animation}, nil
}

func MustLoadSprite(animationPath, palettePath string) *Sprite {
	sprite, err := LoadSprite(animationPath, palettePath)
	if err != nil {
		panic(err)
	}

	return sprite
}

func (s *Sprite) Render(target *ebiten.Image) error {
	if err := s.advance(); err != nil {
		return err
	}

	_, frameHeight := s.animation.GetCurrentFrameSize()

	if err := s.animation.Render(target, s.x, s.y-frameHeight); err != nil {
		return err
	}

	return nil
}

func (s *Sprite) RenderSegmented(target *ebiten.Image, segmentsX, segmentsY, frameOffset int) error {
	if err := s.advance(); err != nil {
		return err
	}

	var currentY int
	for y := 0; y < segmentsY; y++ {
		var currentX int
		var maxFrameHeight int
		for x := 0; x < segmentsX; x++ {
			if err := s.animation.SetCurrentFrame(x + y*segmentsX + frameOffset*segmentsX*segmentsY); err != nil {
				return err
			}

			if err := s.animation.Render(target, currentX, currentY); err != nil {
				return err
			}

			frameWidth, frameHeight := s.GetCurrentFrameSize()
			maxFrameHeight = d2helper.MaxInt(maxFrameHeight, frameHeight)
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
	s.lastFrameTime = d2helper.Now()
	return s.animation.SetCurrentFrame(frameIndex)
}

func (s *Sprite) Rewind() {
	s.lastFrameTime = d2helper.Now()
	s.animation.SetCurrentFrame(0)
}

func (s *Sprite) PlayForward() {
	s.lastFrameTime = d2helper.Now()
	s.animation.PlayForward()
}

func (s *Sprite) PlayBackward() {
	s.lastFrameTime = d2helper.Now()
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

func (s *Sprite) advance() error {
	lastFrameTime := d2helper.Now()
	if err := s.animation.Advance(lastFrameTime - s.lastFrameTime); err != nil {
		return err
	}
	s.lastFrameTime = lastFrameTime
	return nil
}
