package help

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
	tlCornerFrame = iota
	lFrame
	tFrameLHalf
	tFrameRHalf
	trCornerFrameTHalf
	trCornerFrameRHalf
	rFrame
)

// Overlay represents the in-game overlay that toggles visibility when the h key is pressed
type Overlay struct {
	asset       *d2asset.AssetManager
	isOpen      bool
	renderer    d2interface.Renderer
	frames      []*d2ui.Sprite
	text        []*d2ui.Label
	lines       []line
	uiManager   *d2ui.UIManager
	originX     int
	originY     int
	layout      *d2gui.Layout
	closeButton *d2ui.Button
	guiManager  *d2gui.GuiManager
}

func NewHelpOverlay(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
) *Overlay {
	h := &Overlay{
		asset:      asset,
		renderer:   renderer,
		uiManager:  ui,
		guiManager: guiManager,
	}

	return h
}

func (h *Overlay) onMouseDown() {
	// If mouse over close button
	// close()
}

func (h *Overlay) Toggle() {
	fmt.Print("Help overlay toggled\n")
	if h.isOpen {
		h.close()
	} else {
		h.open()
	}
}

func (h *Overlay) close() {
	h.isOpen = false
	h.closeButton.SetVisible(false)
	h.guiManager.SetLayout(nil)
}

func (h *Overlay) open() {
	h.isOpen = true
	if h.layout == nil {
		h.layout = d2gui.CreateLayout(h.renderer, d2gui.PositionTypeHorizontal, h.asset)
	}

	h.closeButton.SetVisible(true)
	h.closeButton.SetPressed(false)

	h.guiManager.SetLayout(h.layout)
}

func (h *Overlay) IsOpen() bool {
	return h.isOpen
}

func (h *Overlay) IsInRect(px, py int) bool {
	ww, hh := h.closeButton.GetSize()
	x, y := h.closeButton.GetPosition()

	if px >= x && px <= x+ww && py >= y && py <= y+hh {
		return true
	}
	return false
}

func (h *Overlay) Load() {

	var (
		x     = 0
		y     = 0
		prevX = 0
		prevY = 0
	)
	for frameIndex := 0; frameIndex < 7; frameIndex++ {
		f, _ := h.uiManager.NewSprite(d2resource.HelpBorder, d2resource.PaletteSky)
		_ = f.SetCurrentFrame(frameIndex)

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
	ww, _ := newLabel.GetSize()
	newLabel.SetPosition((800/2)-(ww/2)-30, 0)
	h.text = append(h.text, newLabel)

	// Close

	//close, _ := h.uiManager.NewSprite(d2resource.SquareButton, d2resource.PaletteSky)
	//_ = close.SetCurrentFrame(0)
	//close.SetPosition(685, 57)
	//h.frames = append(h.frames, close)

	h.closeButton = h.uiManager.NewButton(d2ui.ButtonTypeClose, "0")
	h.closeButton.SetPosition(685, 25)
	h.closeButton.SetVisible(false)
	h.closeButton.OnActivated(func() { h.close() })

	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText("Close")
	newLabel.SetPosition(680, 60)
	h.text = append(h.text, newLabel)

	newLabel = h.uiManager.NewLabel(d2resource.Font30, d2resource.PaletteSky)
	newLabel.SetText("0")
	newLabel.SetPosition(695, 32)
	h.text = append(h.text, newLabel)

	// Bullets

	yOffset := 60
	h.createBullet(callout{
		LabelText: "Hold Down <Ctrl> to Run",
		LabelX:    100,
		LabelY:    yOffset,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 12
	h.createBullet(callout{
		LabelText: "Hold down <Alt> to highlight items on the ground",
		LabelX:    100,
		LabelY:    yOffset,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 12
	h.createBullet(callout{
		LabelText: "Hold down <Shift> to attack while standing still",
		LabelX:    100,
		LabelY:    yOffset,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 12
	h.createBullet(callout{
		LabelText: "Hit <Tab> to toggle the automap on and off",
		LabelX:    100,
		LabelY:    yOffset,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 12
	h.createBullet(callout{
		LabelText: "Hit <Esc> to bring up the Game Menu",
		LabelX:    100,
		LabelY:    yOffset,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 12
	h.createBullet(callout{
		LabelText: "Hit <Enter> to go into chat mode",
		LabelX:    100,
		LabelY:    yOffset,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 12
	h.createBullet(callout{
		LabelText: "Use F1-F8 to set your Left or Right Mouse Button Skills",
		LabelX:    100,
		LabelY:    yOffset,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 12
	h.createBullet(callout{
		LabelText: "Hit <H> to toggle this screen open and closed",
		LabelX:    100,
		LabelY:    yOffset,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	// Callouts

	h.createCallout(callout{
		LabelText: "New Stats",
		LabelX:    220,
		LabelY:    350,
		DotX:      215,
		DotY:      575,
	})

	h.createCallout(callout{
		LabelText: "New Skill",
		LabelX:    575,
		LabelY:    350,
		DotX:      570,
		DotY:      575,
	})

	h.createCallout(callout{
		LabelText: "Left Mouse-\nButton Skill\n(Click to Change)",
		LabelX:    135,
		LabelY:    380,
		DotX:      130,
		DotY:      565,
	})

	h.createCallout(callout{
		LabelText: "Mini-Panel\n(Opens Character,\ninventory, and\nother screens)",
		LabelX:    450,
		LabelY:    365,
		DotX:      445,
		DotY:      540,
	})

	h.createCallout(callout{
		LabelText: "Right Mouse-\nButton Skill\n(Click to Change)",
		LabelX:    675,
		LabelY:    375,
		DotX:      670,
		DotY:      560,
	})

	h.createCallout(callout{
		LabelText: "Life Orb",
		LabelX:    65,
		LabelY:    460,
		DotX:      60,
		DotY:      535,
	})

	h.createCallout(callout{
		LabelText: "Stamina Bar",
		LabelX:    315,
		LabelY:    460,
		DotX:      310,
		DotY:      585,
	})

	h.createCallout(callout{
		LabelText: "Mana Orb",
		LabelX:    745,
		LabelY:    460,
		DotX:      740,
		DotY:      535,
	})

	h.createCallout(callout{
		LabelText: "Run/Walk\nToggle",
		LabelX:    263,
		LabelY:    480,
		DotX:      258,
		DotY:      585,
	})

	h.createCallout(callout{
		LabelText: "Experience\nBar",
		LabelX:    370,
		LabelY:    480,
		DotX:      365,
		DotY:      565,
	})

	h.createCallout(callout{
		LabelText: "Belt",
		LabelX:    535,
		LabelY:    490,
		DotX:      530,
		DotY:      568,
	})

}

type line struct {
	StartX int
	StartY int
	MoveX  int
	MoveY  int
	Color  color.Color
}

type callout struct {
	LabelText string
	LabelX    int
	LabelY    int
	DotX      int
	DotY      int
}

func (h *Overlay) createBullet(c callout) {
	newLabel := h.uiManager.NewLabel(d2resource.FontFormal11, d2resource.PaletteSky)
	newLabel.SetText(c.LabelText)
	//ww, hh = newLabel.GetSize()
	newLabel.SetPosition(c.LabelX, c.LabelY)
	h.text = append(h.text, newLabel)

	newDot, _ := h.uiManager.NewSprite(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	_ = newDot.SetCurrentFrame(0)
	newDot.SetPosition(c.DotX, c.DotY+14)
	h.frames = append(h.frames, newDot)
}

func (h *Overlay) createCallout(c callout) {
	newLabel := h.uiManager.NewLabel(d2resource.FontFormal11, d2resource.PaletteSky)
	newLabel.Color[0] = color.White
	newLabel.SetText(c.LabelText)
	newLabel.SetPosition(c.LabelX, c.LabelY)
	newLabel.Alignment = d2gui.HorizontalAlignCenter
	ww, hh := newLabel.GetTextMetrics(c.LabelText)
	h.text = append(h.text, newLabel)
	_ = ww

	l := line{
		StartX: c.LabelX,
		StartY: c.LabelY + hh + 5,
		MoveX:  0,
		MoveY:  c.DotY - c.LabelY - hh - 5,
		Color:  color.White,
	}
	h.lines = append(h.lines, l)

	newDot, _ := h.uiManager.NewSprite(d2resource.HelpWhiteBullet, d2resource.PaletteSky)
	_ = newDot.SetCurrentFrame(0)
	newDot.SetPosition(c.DotX, c.DotY)
	h.frames = append(h.frames, newDot)
}

func (h *Overlay) Render(target d2interface.Surface) error {
	if !h.isOpen {
		return nil
	}

	for _, f := range h.frames {
		_ = f.Render(target)
	}

	for _, t := range h.text {
		t.Render(target)
	}

	for _, l := range h.lines {
		target.PushTranslation(l.StartX, l.StartY)
		target.DrawLine(l.MoveX, l.MoveY, l.Color)
		target.Pop()
	}

	return nil
}
