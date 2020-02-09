package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Sprite struct {
	x int
	y int

	segmentsX   int
	segmentsY   int
	frameOffset int

	visible   bool
	animation *d2asset.Animation
}

func CreateSprite(imagePath, palettePath string) (*Sprite, error) {
	animation, err := d2asset.LoadAnimation(imagePath, palettePath)
	if err != nil {
		return nil, err
	}

	sprite := &Sprite{
		animation: animation,
		visible:   true,
	}

	return sprite, nil
}

func (s *Sprite) SetSegmented(segmentsX, segmentsY, frameOffset int) {
	s.segmentsX = segmentsX
	s.segmentsY = segmentsY
	s.frameOffset = frameOffset
}

func (s *Sprite) SetPosition(x, y int) {
	s.x = x
	s.y = y
}

func (s *Sprite) Show() {
	s.visible = true
}

func (s *Sprite) Hide() {
	s.visible = false
}

func (s *Sprite) getPosition() (int, int) {
	return s.x, s.y
}

func (s *Sprite) getSize() (int, int) {
	return s.animation.GetCurrentFrameSize()
}

func (s *Sprite) render(target d2render.Surface) error {
	if !s.visible {
		return nil
	}

	_, height := s.animation.GetCurrentFrameSize()
	target.PushTranslation(s.x, s.y-height)
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
	return s.animation.Advance(elapsed)
}
