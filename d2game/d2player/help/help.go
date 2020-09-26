package help

import (
	"fmt"
	"image/color"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
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
		f, err := h.uiManager.NewSprite(d2resource.HelpBorder, d2resource.PaletteSky)
		if err != nil {
			log.Print(err)
		}

		err = f.SetCurrentFrame(frameIndex)
		if err != nil {
			log.Print(err)
		}

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

	text := d2tbl.TranslateString("Strhelp1") // "Diablo II Help"
	newLabel := h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)
	ww, _ := newLabel.GetSize()
	newLabel.SetPosition((800/2)-(ww/2)-30, 0)
	h.text = append(h.text, newLabel)

	// Close

	h.closeButton = h.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	h.closeButton.SetPosition(685, 25)
	h.closeButton.SetVisible(false)
	h.closeButton.OnActivated(func() { h.close() })

	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(d2tbl.TranslateString("strClose")) // "Close"
	newLabel.SetPosition(680, 60)
	h.text = append(h.text, newLabel)

	// Bullets

	yOffset := 59
	h.createBullet(callout{
		LabelText: fmt.Sprintf(d2tbl.TranslateString("StrHelp2"), "Ctrl"), // TODO "Ctrl" should be hotkey // "Hold Down <%s> to Run"
		LabelX:    100,
		LabelY:    yOffset - 10,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 20
	h.createBullet(callout{
		LabelText: fmt.Sprintf(d2tbl.TranslateString("StrHelp3"), "Alt"), // TODO "Alt" should be hotkey // "Hold down <%s> to highlight items on the ground"
		LabelX:    100,
		LabelY:    yOffset - 10,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 20
	h.createBullet(callout{
		LabelText: fmt.Sprintf(d2tbl.TranslateString("StrHelp4"), "Shift"), // TODO "Shift" should be hotkey // "Hold down <%s> to attack while standing still"
		LabelX:    100,
		LabelY:    yOffset - 10,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 20
	h.createBullet(callout{
		LabelText: fmt.Sprintf(d2tbl.TranslateString("StrHelp5"), "Tab"), // TODO "Tab" should be hotkey // "Hit <%s> to toggle the automap on and off"
		LabelX:    100,
		LabelY:    yOffset - 10,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 20
	h.createBullet(callout{
		LabelText: d2tbl.TranslateString("StrHelp6"), // "Hit <Esc> to bring up the Game Menu"
		LabelX:    100,
		LabelY:    yOffset - 10,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 20
	h.createBullet(callout{
		LabelText: d2tbl.TranslateString("StrHelp7"), // "Hit <Enter> to go into chat mode"
		LabelX:    100,
		LabelY:    yOffset - 10,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 20
	h.createBullet(callout{
		LabelText: d2tbl.TranslateString("StrHelp8"), // "Hit F1-F8 to set your Left or Right Mouse Buttton Skills."
		LabelX:    100,
		LabelY:    yOffset - 10,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	yOffset += 20
	h.createBullet(callout{
		LabelText: fmt.Sprintf(d2tbl.TranslateString("StrHelp8a"), "H"), // TODO "H" should be hotkey
		LabelX:    100,
		LabelY:    yOffset - 10,
		DotX:      100 - 12,
		DotY:      yOffset,
	})

	// Callouts

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("strlvlup"), // "New Stats"
		LabelX:    222,
		LabelY:    355,
		DotX:      217,
		DotY:      574,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("strnewskl"), // "New Skill"
		LabelX:    578,
		LabelY:    355,
		DotX:      573,
		DotY:      574,
	})

	// Some of the help fonts require mulktiple lines.
	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp10"), // "Left Mouse-"
		LabelX:    135,
		LabelY:    382,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp11"), // "Button Skill"
		LabelX:    135,
		LabelY:    397,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp12"), // "(Click to Change)"
		LabelX:    135,
		LabelY:    412,
		DotX:      130,
		DotY:      565,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp13"), // "Right Mouse"
		LabelX:    675,
		LabelY:    381,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp11"), // "Button Skill"
		LabelX:    675,
		LabelY:    396,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp12"), // "(Click to Change)"
		LabelX:    675,
		LabelY:    411,
		DotX:      670,
		DotY:      562,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp17"), // "Mini-Panel"
		LabelX:    450,
		LabelY:    371,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp18"), // "(Opens Character,"
		LabelX:    450,
		LabelY:    386,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp19"), // "inventory, and"
		LabelX:    450,
		LabelY:    401,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp20"), // "other screens)"
		LabelX:    450,
		LabelY:    417,
		DotX:      445,
		DotY:      539,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp9"), // "Life Orb"
		LabelX:    65,
		LabelY:    451,
		DotX:      60,
		DotY:      538,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp15"), // "Stamina Bar"
		LabelX:    315,
		LabelY:    450,
		DotX:      310,
		DotY:      583,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp22"), // "Mana Orb"
		LabelX:    745,
		LabelY:    451,
		DotX:      740,
		DotY:      538,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp14"), // "Run/Walk"
		LabelX:    264,
		LabelY:    480,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp14a"), // "Toggle"
		LabelX:    264,
		LabelY:    495,
		DotX:      259,
		DotY:      583,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp16"), // "Experience"
		LabelX:    370,
		LabelY:    476,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp16a"), // "Bar"
		LabelX:    370,
		LabelY:    493,
		DotX:      365,
		DotY:      565,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp21"), // "Belt"
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
	newLabel := h.uiManager.NewLabel(d2resource.FontFormal12, d2resource.PaletteSky)
	newLabel.SetText(c.LabelText)
	//ww, hh = newLabel.GetSize()
	newLabel.SetPosition(c.LabelX, c.LabelY)
	h.text = append(h.text, newLabel)

	newDot, err := h.uiManager.NewSprite(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	err = newDot.SetCurrentFrame(0)
	if err != nil {
		log.Print(err)
	}
	newDot.SetPosition(c.DotX, c.DotY+14)
	h.frames = append(h.frames, newDot)
}

func (h *Overlay) createLabel(c callout) {
	newLabel := h.uiManager.NewLabel(d2resource.FontFormal12, d2resource.PaletteSky)
	newLabel.SetText(c.LabelText)
	//ww, hh = newLabel.GetSize()
	newLabel.SetPosition(c.LabelX, c.LabelY)
	h.text = append(h.text, newLabel)
	newLabel.Alignment = d2gui.HorizontalAlignCenter
}

func (h *Overlay) createCallout(c callout) {
	newLabel := h.uiManager.NewLabel(d2resource.FontFormal12, d2resource.PaletteSky)
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

	newDot, err := h.uiManager.NewSprite(d2resource.HelpWhiteBullet, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	err = newDot.SetCurrentFrame(0)
	if err != nil {
		log.Print(err)
	}

	newDot.SetPosition(c.DotX, c.DotY)
	h.frames = append(h.frames, newDot)
}

func (h *Overlay) Render(target d2interface.Surface) error {
	if !h.isOpen {
		return nil
	}

	for _, f := range h.frames {
		err := f.Render(target)
		if err != nil {
			return err
		}
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
