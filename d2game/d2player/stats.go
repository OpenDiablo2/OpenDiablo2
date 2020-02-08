package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type HeroStats struct {
	frame   *d2ui.Sprite
	panel   *d2ui.Sprite
	originX int
	originY int
	isOpen  bool
}

func NewHeroStats() *HeroStats {
	originX := 0
	originY := 0
	return &HeroStats{
		originX: originX,
		originY: originY,
	}
}

func (s *HeroStats) Load() {
	animation, _ := d2asset.LoadAnimation(d2resource.Frame, d2resource.PaletteSky)
	s.frame, _ = d2ui.LoadSprite(animation)
	animation, _ = d2asset.LoadAnimation(d2resource.InventoryCharacterPanel, d2resource.PaletteSky)
	s.panel, _ = d2ui.LoadSprite(animation)
}

func (s *HeroStats) IsOpen() bool {
	return s.isOpen
}

func (s *HeroStats) Toggle() {
	s.isOpen = !s.isOpen
}

func (s *HeroStats) Open() {
	s.isOpen = true
}

func (s *HeroStats) Close() {
	s.isOpen = false
}

func (s *HeroStats) Render(target d2render.Surface) {
	if !s.isOpen {
		return
	}

	x, y := s.originX, s.originY

	// Frame
	// Top left
	s.frame.SetCurrentFrame(0)
	w, h := s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, y+h)
	s.frame.Render(target)
	x += w
	y += h

	// Top right
	s.frame.SetCurrentFrame(1)
	w, h = s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, s.originY+h)
	s.frame.Render(target)
	x = s.originX

	// Right
	s.frame.SetCurrentFrame(2)
	w, h = s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, y+h)
	s.frame.Render(target)
	y += h

	// Bottom left
	s.frame.SetCurrentFrame(3)
	w, h = s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, y+h)
	s.frame.Render(target)
	x += w

	// Bottom right
	s.frame.SetCurrentFrame(4)
	w, h = s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, y+h)
	s.frame.Render(target)

	x, y = s.originX, s.originY
	y += 64
	x += 80

	// Panel
	// Top left
	s.panel.SetCurrentFrame(0)
	w, h = s.panel.GetCurrentFrameSize()
	s.panel.SetPosition(x, y+h)
	s.panel.Render(target)
	x += w

	// Top right
	s.panel.SetCurrentFrame(1)
	w, h = s.panel.GetCurrentFrameSize()
	s.panel.SetPosition(x, y+h)
	s.panel.Render(target)
	y += h

	// Bottom right
	s.panel.SetCurrentFrame(3)
	w, h = s.panel.GetCurrentFrameSize()
	s.panel.SetPosition(x, y+h)
	s.panel.Render(target)

	// Bottom left
	s.panel.SetCurrentFrame(2)
	w, h = s.panel.GetCurrentFrameSize()
	s.panel.SetPosition(x-w, y+h)
	s.panel.Render(target)

}
