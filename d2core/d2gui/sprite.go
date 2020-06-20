package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type AnimationDirection int

const (
	DirectionForward  AnimationDirection = 0
	DirectionBackward                    = 1
)

type Sprite struct {
	widgetBase

	segmentsX   int
	segmentsY   int
	frameOffset int

	animation *d2asset.Animation
}

type AnimatedSprite struct {
	*Sprite
}

func createSprite(imagePath, palettePath string) (*Sprite, error) {
	animation, err := d2asset.LoadAnimation(imagePath, palettePath)
	if err != nil {
		return nil, err
	}

	sprite := &Sprite{}
	sprite.animation = animation
	sprite.SetVisible(true)

	return sprite, nil
}

func createAnimatedSprite(imagePath, palettePath string, direction AnimationDirection) (*AnimatedSprite, error) {
	animation, err := d2asset.LoadAnimation(imagePath, palettePath)
	if err != nil {
		return nil, err
	}
	sprite := &AnimatedSprite{
		&Sprite{},
	}
	sprite.animation = animation
	if direction == DirectionForward {
		sprite.animation.PlayForward()
	} else {
		sprite.animation.PlayBackward()
	}
	sprite.animation.SetBlend(false)
	sprite.SetVisible(true)

	return sprite, nil
}

func (s *AnimatedSprite) render(target d2interface.Surface) error {
	_, frameHeight := s.animation.GetCurrentFrameSize()

	target.PushTranslation(s.x, s.y-frameHeight)
	defer target.Pop()
	return s.animation.Render(target)
}

func (s *Sprite) SetSegmented(segmentsX, segmentsY, frameOffset int) {
	s.segmentsX = segmentsX
	s.segmentsY = segmentsY
	s.frameOffset = frameOffset
}

func (s *Sprite) render(target d2interface.Surface) error {
	return renderSegmented(s.animation, s.segmentsX, s.segmentsY, s.frameOffset, target)
}

func (s *Sprite) advance(elapsed float64) error {
	return s.animation.Advance(elapsed)
}

func (s *Sprite) getSize() (int, int) {
	return s.animation.GetCurrentFrameSize()
}
