package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Sprite struct {
	widgetBase

	segmentsX   int
	segmentsY   int
	frameOffset int

	animation *d2asset.Animation
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

func (s *Sprite) SetSegmented(segmentsX, segmentsY, frameOffset int) {
	s.segmentsX = segmentsX
	s.segmentsY = segmentsY
	s.frameOffset = frameOffset
}

func (s *Sprite) render(target d2render.Surface) error {
	return renderSegmented(s.animation, s.segmentsX, s.segmentsY, s.frameOffset, target)
}

func (s *Sprite) advance(elapsed float64) error {
	return s.animation.Advance(elapsed)
}

func (s *Sprite) getSize() (int, int) {
	return s.animation.GetCurrentFrameSize()
}
