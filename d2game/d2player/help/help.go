// Package help contains the in-game diablo2 help panel
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

const (
	inHalf = 2 // when we divide by 2
)

const (
	// all in pixels
	windowWidth = 800

	bulletOffsetY = 14

	// the title of the panel
	titleLabelOffsetX = -30

	// for the bulleted list near the top of the screen
	listRootX              = 100
	listRootY              = 59
	listBulletOffsetY      = 10
	listBulletOffsetX      = 12
	listItemVerticalOffset = 20
	listBulletRootY        = listRootY - listBulletOffsetY + listItemVerticalOffset
	listBulletX            = listRootX - listBulletOffsetX

	// the close button for the help panel
	closeButtonX = 685
	closeButtonY = 25

	// the rest of these are for text with a line and dot, towards the bottom of the screen
	newStatsLabelX = 222
	newStatsLabelY = 355
	newStatsDotX   = 217
	newStatsDotY   = 574

	newSkillLabelX = 578
	newSkillLabelY = 355
	newSkillDotX   = 573
	newSkillDotY   = 574

	leftMouseLabelX = 135
	leftMouseLabelY = 382

	leftButtonSkillLabelX = 135
	leftButtonSkillLabelY = 397

	leftSkillClickToChangeLabelX = 135
	leftSkillClickToChangeLabelY = 412
	leftSkillClickToChangeDotX   = 130
	leftSkillClickToChangeDotY   = 565

	rightMouseLabelX = 675
	rightMouseLabelY = 381

	rightButtonSkillLabelX = 675
	rightButtonSkillLabelY = 396

	rightSkillClickToChangeLabelX = 675
	rightSkillClickToChangeLabelY = 411
	rightSkillClickToChangeDotX   = 670
	rightSkillClickToChangeDotY   = 562

	miniPanelLabelX = 450
	miniPanelLabelY = 371

	characterLabelX = 450
	characterLabelY = 386

	inventoryLabelX = 450
	inventoryLabelY = 401

	otherScreensLabelX = 450
	otherScreensLabelY = 417
	otherScreensDotX   = 445
	otherScreensDotY   = 539

	lifeOrbLabelX = 65
	lifeOrbLabelY = 451
	lifeOrbDotX   = 60
	lifeOrbDotY   = 538

	staminaBarLabelX = 315
	staminaBarLabelY = 450
	staminaBarDotX   = 310
	staminaBarDotY   = 583

	manaOrbLabelX = 745
	manaOrbLabelY = 451
	manaOrbDotX   = 740
	manaOrbDotY   = 538

	runWalkButtonLabelX = 264
	runWalkButtonLabelY = 480

	toggleLabelX = 264
	toggleLabelY = 495
	toggleDotX   = 259
	toggleDotY   = 583

	experienceLabelX = 370
	experienceLabelY = 476

	barLabelX = 370
	barLabelY = 493
	barDotX   = 365
	barDotY   = 565

	beltLabelX = 535
	beltLabelY = 490
	beltDotX   = 530
	beltDotY   = 568
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

// NewHelpOverlay creates a new HelpOverlay instance
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

// Toggle the visibility state of the overlay
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

// IsOpen returns whether or not the overlay is visible/open
func (h *Overlay) IsOpen() bool {
	return h.isOpen
}

// IsInRect checks if the given point is within the overlay layout rectangle
func (h *Overlay) IsInRect(px, py int) bool {

	ww, hh := h.layout.GetSize()
	x, y := h.layout.GetPosition()

	if px >= x && px <= x+ww && py >= y && py <= y+hh {
		return true
	}

	return false
}

// Load the overlay graphical assets
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
			x = windowWidth - ww - 245
		case trCornerFrameTHalf:
			y = hh
			x = windowWidth - ww - 20
		case trCornerFrameRHalf:
			y = hh
			x = windowWidth - ww
		case rFrame:
			y = hh + prevY
			x = windowWidth - ww
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

	titleLabelWidth, _ := newLabel.GetSize()

	newLabel.SetPosition((windowWidth/inHalf)-(titleLabelWidth/inHalf)+titleLabelOffsetX, 0)
	h.text = append(h.text, newLabel)

	// Close

	h.closeButton = h.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	h.closeButton.SetPosition(closeButtonX, closeButtonY)
	h.closeButton.SetVisible(false)
	h.closeButton.OnActivated(func() { h.close() })

	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(d2tbl.TranslateString("strClose")) // "Close"
	newLabel.SetPosition(680, 60)
	h.text = append(h.text, newLabel)

	// Bullets

	callouts := []struct{ text string }{
		// TODO "Ctrl" should be hotkey // "Hold Down <%s> to Run"
		{text: fmt.Sprintf(d2tbl.TranslateString("StrHelp2"), "Ctrl")},

		// TODO "Alt" should be hotkey // "Hold down <%s> to highlight items on the ground"
		{text: fmt.Sprintf(d2tbl.TranslateString("StrHelp3"), "Alt")},
		// TODO "Shift" should be hotkey // "Hold down <%s> to attack while standing still"
		{text: fmt.Sprintf(d2tbl.TranslateString("StrHelp4"), "Shift")},

		// TODO "Tab" should be hotkey // "Hit <%s> to toggle the automap on and off"
		{text: fmt.Sprintf(d2tbl.TranslateString("StrHelp5"), "Tab")},

		// "Hit <Esc> to bring up the Game Menu"
		{text: d2tbl.TranslateString("StrHelp6")},

		// "Hit <Enter> to go into chat mode"
		{text: d2tbl.TranslateString("StrHelp7")},

		// "Hit F1-F8 to set your Left or Right Mouse Buttton Skills."
		{text: d2tbl.TranslateString("StrHelp8")},

		// TODO "H" should be hotkey,
		{text: fmt.Sprintf(d2tbl.TranslateString("StrHelp8a"), "H")},
	}

	for idx := range callouts {
		listItemOffsetY := idx * listItemVerticalOffset

		h.createBullet(callout{
			LabelText: callouts[idx].text,
			LabelX:    listRootX,
			LabelY:    listRootY + listItemOffsetY,
			DotX:      listBulletX,
			DotY:      listBulletRootY + listItemOffsetY,
		})
	}

	// Callouts

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("strlvlup"), // "New Stats"
		LabelX:    newStatsLabelX,
		LabelY:    newStatsLabelY,
		DotX:      newStatsDotX,
		DotY:      newStatsDotY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("strnewskl"), // "New Skill"
		LabelX:    newSkillLabelX,
		LabelY:    newSkillLabelY,
		DotX:      newSkillDotX,
		DotY:      newSkillDotY,
	})

	// Some of the help fonts require mulktiple lines.
	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp10"), // "Left Mouse-"
		LabelX:    leftMouseLabelX,
		LabelY:    leftMouseLabelY,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp11"), // "Button Skill"
		LabelX:    leftButtonSkillLabelX,
		LabelY:    leftButtonSkillLabelY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp12"), // "(Click to Change)"
		LabelX:    leftSkillClickToChangeLabelX,
		LabelY:    leftSkillClickToChangeLabelY,
		DotX:      leftSkillClickToChangeDotX,
		DotY:      leftSkillClickToChangeDotY,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp13"), // "Right Mouse"
		LabelX:    rightMouseLabelX,
		LabelY:    rightMouseLabelY,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp11"), // "Button Skill"
		LabelX:    rightButtonSkillLabelX,
		LabelY:    rightButtonSkillLabelY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp12"), // "(Click to Change)"
		LabelX:    rightSkillClickToChangeLabelX,
		LabelY:    rightSkillClickToChangeLabelY,
		DotX:      rightSkillClickToChangeDotX,
		DotY:      rightSkillClickToChangeDotY,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp17"), // "Mini-Panel"
		LabelX:    miniPanelLabelX,
		LabelY:    miniPanelLabelY,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp18"), // "(Opens Character,"
		LabelX:    characterLabelX,
		LabelY:    characterLabelY,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp19"), // "inventory, and"
		LabelX:    inventoryLabelX,
		LabelY:    inventoryLabelY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp20"), // "other screens)"
		LabelX:    otherScreensLabelX,
		LabelY:    otherScreensLabelY,
		DotX:      otherScreensDotX,
		DotY:      otherScreensDotY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp9"), // "Life Orb"
		LabelX:    lifeOrbLabelX,
		LabelY:    lifeOrbLabelY,
		DotX:      lifeOrbDotX,
		DotY:      lifeOrbDotY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp15"), // "Stamina Bar"
		LabelX:    staminaBarLabelX,
		LabelY:    staminaBarLabelY,
		DotX:      staminaBarDotX,
		DotY:      staminaBarDotY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp22"), // "Mana Orb"
		LabelX:    manaOrbLabelX,
		LabelY:    manaOrbLabelY,
		DotX:      manaOrbDotX,
		DotY:      manaOrbDotY,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp14"), // "Run/Walk"
		LabelX:    runWalkButtonLabelX,
		LabelY:    runWalkButtonLabelY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp14a"), // "Toggle"
		LabelX:    toggleLabelX,
		LabelY:    toggleLabelY,
		DotX:      toggleDotX,
		DotY:      toggleDotY,
	})

	h.createLabel(callout{
		LabelText: d2tbl.TranslateString("StrHelp16"), // "Experience"
		LabelX:    experienceLabelX,
		LabelY:    experienceLabelY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp16a"), // "Bar"
		LabelX:    barLabelX,
		LabelY:    barLabelY,
		DotX:      barDotX,
		DotY:      barDotY,
	})

	h.createCallout(callout{
		LabelText: d2tbl.TranslateString("StrHelp21"), // "Belt"
		LabelX:    beltLabelX,
		LabelY:    beltLabelY,
		DotX:      beltDotX,
		DotY:      beltDotY,
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

	newDot.SetPosition(c.DotX, c.DotY+bulletOffsetY)
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

// Render the overlay to the given surface
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
