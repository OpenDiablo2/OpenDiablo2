package d2player

import (
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	overlayLayoutID int = 1
)

const (
	tlCornerFrame = iota
	lFrame
	tFrameLHalf
	tFrameRHalf
	trCornerFrameTHalf
	trCornerFrameRHalf
	rFrame
)

type line struct {
	StartX int
	StartY int
	MoveX  int
	MoveY  int
	Color  color.Color
}

// HelpOverlay represents the in-game overlay that toggles visibility when the h key is pressed
type HelpOverlay struct {
	isOpen    bool
	renderer  d2interface.Renderer
	frames    []*d2ui.Sprite
	text      []*d2ui.Label
	lines     []line
	uiManager *d2ui.UIManager
	originX   int
	originY   int
	layout    *d2gui.Layout
}

func NewHelpOverlay(renderer d2interface.Renderer) *HelpOverlay {
	h := &HelpOverlay{
		renderer: renderer,
	}

	return h
}

func (h *HelpOverlay) onHoverElement(n int) {

}

func (h *HelpOverlay) setLayout(id int) {
	h.onHoverElement(0)
}

func (h *HelpOverlay) Toggle() {
	fmt.Print("Help overlay toggled\n")
	if h.isOpen {
		h.close()
	} else {
		h.open()
	}
}

func (h *HelpOverlay) close() {
	h.isOpen = false
	d2gui.SetLayout(nil)
}

func (h *HelpOverlay) open() {
	h.isOpen = true
	if h.layout == nil {
		h.layout = d2gui.CreateLayout(h.renderer, d2gui.PositionTypeHorizontal)
		// layoutLeft := h.layout.AddLayout(d2gui.PositionTypeVertical)

		// tlCorner, _ := layoutLeft.AddSprite(d2resource.HelpBorder, d2resource.PaletteSky)
		// tlCorner.SetSegmented(0, 0, 0)

		//layoutLeft.AddSprite(imagePath, palettePath)
	}
	d2gui.SetLayout(h.layout)
}

func (h *HelpOverlay) Load() {

	var (
		x     = 0
		y     = 0
		prevX = 0
		prevY = 0
	)
	for frameIndex := 0; frameIndex < 7; frameIndex++ {
		animation, _ := d2asset.LoadAnimation(d2resource.HelpBorder, d2resource.PaletteSky)
		animation.SetCurrentFrame(frameIndex)
		f, _ := h.uiManager.NewSprite(animation)

		ww, hh := f.GetCurrentFrameSize()
		//fmt.Printf("Help frame %d size: %d, %d\n", frameIndex, ww, hh)

		switch frameIndex {
		case tlCornerFrame:
			y = hh
		case lFrame:
			y = hh + prevY
		case tFrameLHalf:
			y = hh
			x = 65
		case tFrameRHalf:
			y = hh
			x = 800 - ww - 245
		case trCornerFrameTHalf:
			y = hh
			x = 800 - ww - 20
		case trCornerFrameRHalf:
			y = hh
			x = 800 - ww
		case rFrame:
			y = hh + prevY
			x = 800 - ww
		}

		//y += 50

		_ = prevX

		prevX = x
		prevY = y
		f.SetPosition(x, y)
		h.frames = append(h.frames, f)
	}

	// Title

	text := "Diablo II Help"
	newLabel := h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, hh := newLabel.GetSize()
	newLabel.SetPosition((800/2)-(ww/2)-30, 0)
	h.text = append(h.text, newLabel)

	// Bullets

	yOffset := 60
	text = "Hold Down <Ctrl> to Run"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, hh = newLabel.GetSize()
	newLabel.SetPosition(100, yOffset)
	h.text = append(h.text, newLabel)

	anim, _ := d2asset.LoadAnimation(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	newDot, _ := h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(100-12, yOffset+14)
	h.frames = append(h.frames, newDot)

	yOffset += hh
	text = "Hold Down <Alt> to highlight items on the ground"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, hh = newLabel.GetSize()
	newLabel.SetPosition(100, yOffset)
	h.text = append(h.text, newLabel)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(100-12, yOffset+14)
	h.frames = append(h.frames, newDot)

	yOffset += hh
	text = "Hold Down <Shift> to attack while standing still"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, hh = newLabel.GetSize()
	newLabel.SetPosition(100, yOffset)
	h.text = append(h.text, newLabel)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(100-12, yOffset+14)
	h.frames = append(h.frames, newDot)

	yOffset += hh
	text = "Hit <Tab> to toggle the automap on and off"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, hh = newLabel.GetSize()
	newLabel.SetPosition(100, yOffset)
	h.text = append(h.text, newLabel)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(100-12, yOffset+14)
	h.frames = append(h.frames, newDot)

	yOffset += hh
	text = "Hit <Esc> to bring up the Game Menu"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, hh = newLabel.GetSize()
	newLabel.SetPosition(100, yOffset)
	h.text = append(h.text, newLabel)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(100-12, yOffset+14)
	h.frames = append(h.frames, newDot)

	yOffset += hh
	text = "Hit <Enter> to go into chat mode"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, hh = newLabel.GetSize()
	newLabel.SetPosition(100, yOffset)
	h.text = append(h.text, newLabel)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(100-12, yOffset+14)
	h.frames = append(h.frames, newDot)

	yOffset += hh
	text = "Use F1-F8 to set your Left or Right Mouse Button Skills"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, hh = newLabel.GetSize()
	newLabel.SetPosition(100, yOffset)
	h.text = append(h.text, newLabel)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(100-12, yOffset+14)
	h.frames = append(h.frames, newDot)

	yOffset += hh
	text = "Hit <H> to toggle this screen open and closed"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, hh = newLabel.GetSize()
	newLabel.SetPosition(100, yOffset)
	h.text = append(h.text, newLabel)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(100-12, yOffset+14)
	h.frames = append(h.frames, newDot)

	// Callouts

	text = "New Stats"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	newLabel.SetPosition(180, 350)
	h.text = append(h.text, newLabel)

	l := line{
		StartX: 219,
		StartY: 350 + 30,
		MoveX:  0,
		MoveY:  575 - 350 - 30,
		Color:  color.White,
	}
	h.lines = append(h.lines, l)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpWhiteBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(215, 575)
	h.frames = append(h.frames, newDot)

	text = "New Skill"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	newLabel.SetPosition(540, 350)
	h.text = append(h.text, newLabel)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpWhiteBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(570, 575)
	h.frames = append(h.frames, newDot)

	text = "Left Mouse-\nButton Skill\n(Click to Change)"
	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	newLabel.Alignment = d2gui.HorizontalAlignCenter
	newLabel.SetPosition(135, 380)
	h.text = append(h.text, newLabel)

	anim, _ = d2asset.LoadAnimation(d2resource.HelpWhiteBullet, d2resource.PaletteSky)
	newDot, _ = h.uiManager.NewSprite(anim)
	newDot.SetCurrentFrame(0)
	newDot.SetPosition(130, 565)
	h.frames = append(h.frames, newDot)

}

func (h *HelpOverlay) Render(target d2interface.Surface) error {
	if !h.isOpen {
		return nil
	}

	for _, f := range h.frames {
		f.Render(target)
	}

	for _, t := range h.text {
		t.Render(target)
	}

	for _, l := range h.lines {

		target.PushTranslation(l.StartX, l.StartY)
		target.DrawLine(l.MoveX, l.MoveY, l.Color)
		target.Pop()

		// target.DrawLine(0, entityHeight, color.White)
		// target.DrawLine(entityWidth, 0, color.White)
	}

	// x, y := h.originX, h.originY

	// y += 20

	// frameIndex := 0

	// // Frame
	// // Top left
	// if err := h.frame.SetCurrentFrame(frameIndex); err != nil {
	// 	return err
	// }

	// h.frame.SetCurrentFrame(frameIndex)
	// h.frame.Render(target)
	// frameIndex++

	// width, height := h.frame.GetCurrentFrameSize()
	// _ = width
	// y += height
	// h.frame.SetPosition(x, y)

	// if err := h.frame.SetCurrentFrame(frameIndex); err != nil {
	// 	return err
	// }

	// h.frame.SetCurrentFrame(frameIndex)
	// h.frame.Render(target)
	// frameIndex++

	// // _ = width

	// h.frame.SetPosition(x, y)

	// if err := h.frame.Render(target); err != nil {
	// 	return err
	// }

	//d2gui.Render(target)

	return nil
}
