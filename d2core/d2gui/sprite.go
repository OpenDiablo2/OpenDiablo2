package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
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

func createSprite(imagePath, palettePath string) *Sprite {
	sprite := new(Sprite)
	sprite.animation, _ = d2asset.LoadAnimation(imagePath, palettePath)
	sprite.visible = true
	return sprite
}

func (s *Sprite) SetSegmented(segmentsX, segmentsY, frameOffset int) {
	s.segmentsX = segmentsX
	s.segmentsY = segmentsY
	s.frameOffset = frameOffset
}

func (s *Sprite) render(target d2render.Surface) error {
	if s.animation == nil {
		return nil
	}

	_, height := s.animation.GetCurrentFrameSize()
	target.PushTranslation(0, -height)
	defer target.Pop()

	if s.segmentsX == 0 && s.segmentsY == 0 {
		return s.animation.Render(target)
	}

	var currentY int
	for y := 0; y < s.segmentsY; y++ {
		var currentX int
		var maxHeight int
		for x := 0; x < s.segmentsX; x++ {
			if err := s.animation.SetCurrentFrame(x + y*s.segmentsX + s.frameOffset*s.segmentsX*s.segmentsY); err != nil {
				return err
			}

			target.PushTranslation(s.x+currentX, s.y+currentY)
			err := s.animation.Render(target)
			target.Pop()
			if err != nil {
				return err
			}

			width, height := s.animation.GetCurrentFrameSize()
			maxHeight = d2common.MaxInt(maxHeight, height)
			currentX += width
		}

		currentY += maxHeight
	}

	return nil
}

func (s *Sprite) advance(elapsed float64) error {
	if s.animation == nil {
		return nil
	}

	return s.animation.Advance(elapsed)
}

func (s *Sprite) getSize() (int, int) {
	if s.animation == nil {
		return 0, 0
	}

	return s.animation.GetCurrentFrameSize()
}
