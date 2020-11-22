package d2player

import (
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

/*
	the 800x600 help screen dc6 file frames look like this
	the position we set for frames is the lower-left corner x,y
	+----+------------------+-------------------+------------+----+
	| 1  | 3                | 4                 | 5          | 6  |
	|    |------------------+-------------------|            |    |
	|    |                                      |            |    |
	|    |                                      |            |    |
	+----+                                      +------------+----+
	| 2  |                                                   | 7  |
	|    |                                                   |    |
	|    |                                                   |    |
	+----+                                                   +----+
*/
const (
	// if you add up frame widths 1,3,4,5,6 you get (65+255+255+245+20) = 840
	magicHelpBorderOffsetX = -40
)

const (
	frameTopLeft = iota
	frameBottomLeft
	frameTopMiddleLeft
	frameTopMiddleRight
	frameTopRightNoCorner
	frameTopRight
	frameBottomRight
)

const (
	inHalf = 2 // when we divide by 2
)

const (
	// all in pixels
	windowWidth = 800

	bulletOffsetY = 14
	lineOffset    = 5

	// the title of the panel
	titleLabelOffsetX = -37

	// for the bulleted list near the top of the screen
	listRootX              = 100
	listRootY              = 59
	listBulletOffsetY      = 10
	listBulletOffsetX      = 12
	listItemVerticalOffset = 20
	listBulletRootY        = listRootY - listBulletOffsetY + listItemVerticalOffset
	listBulletX            = listRootX - listBulletOffsetX

	// the close button for the help panel
	closeButtonX      = 685
	closeButtonY      = 25
	closeButtonLabelX = 702
	closeButtonLabelY = 60

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

// NewHelpOverlay creates a new HelpOverlay instance
func NewHelpOverlay(
	asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	l d2util.LogLevel,
	keyMap *KeyMap,
) *HelpOverlay {
	h := &HelpOverlay{
		asset:     asset,
		uiManager: ui,
		keyMap:    keyMap,
	}

	h.Logger = d2util.NewLogger()
	h.Logger.SetLevel(l)
	h.Logger.SetPrefix(logPrefix)

	return h
}

// HelpOverlay represents the in-game overlay that toggles visibility when the h key is pressed
type HelpOverlay struct {
	asset            *d2asset.AssetManager
	isOpen           bool
	frames           []*d2ui.Sprite
	text             []*d2ui.Label
	lines            []line
	uiManager        *d2ui.UIManager
	closeButton      *d2ui.Button
	keyMap           *KeyMap
	onCloseCb        func()
	panelGroup       *d2ui.WidgetGroup
	backgroundWidget *d2ui.CustomWidget

	*d2util.Logger
}

// Toggle the visibility state of the overlay
func (h *HelpOverlay) Toggle() {
	h.Info("Help overlay toggled")

	if h.isOpen {
		h.Close()
	} else {
		h.open()
	}
}

// Close will hide the help overlay
func (h *HelpOverlay) Close() {
	h.isOpen = false
	h.panelGroup.SetVisible(false)
	h.onCloseCb()
}

// SetOnCloseCb sets the callback run when Close() is called
func (h *HelpOverlay) SetOnCloseCb(cb func()) {
	h.onCloseCb = cb
}

func (h *HelpOverlay) open() {
	h.isOpen = true
	h.panelGroup.SetVisible(true)
}

// IsOpen returns whether or not the overlay is visible/open
func (h *HelpOverlay) IsOpen() bool {
	return h.isOpen
}

// IsInRect checks if the given point is within the overlay layout rectangle
func (h *HelpOverlay) IsInRect(px, py int) bool {
	return h.panelGroup.Contains(px, py)
}

// Load the overlay graphical assets
func (h *HelpOverlay) Load() {
	h.panelGroup = h.uiManager.NewWidgetGroup(d2ui.RenderPriorityHelpPanel)

	h.setupOverlayFrame()
	h.setupTitleAndButton()
	h.setupBulletedList()
	h.setupLabelsWithLines()

	h.backgroundWidget = h.uiManager.NewCustomWidgetCached(h.Render, screenWidth, screenHeight)
	h.panelGroup.AddWidget(h.backgroundWidget)
	h.panelGroup.SetVisible(false)
}

func (h *HelpOverlay) setupOverlayFrame() {
	frames := []int{
		frameTopLeft,
		frameBottomLeft,
		frameTopMiddleLeft,
		frameTopMiddleRight,
		frameTopRightNoCorner,
		frameTopRight,
		frameBottomRight,
	}

	left, top := 0, 0
	firstFrameWidth := 0
	prevY := 0
	prevWidth := 0
	currentX, currentY := left, top

	for _, frameIndex := range frames {
		f, err := h.uiManager.NewSprite(d2resource.HelpBorder, d2resource.PaletteSky)
		if err != nil {
			h.Error(err.Error())
		}

		err = f.SetCurrentFrame(frameIndex)
		if err != nil {
			h.Error(err.Error())
		}

		frameWidth, frameHeight := f.GetCurrentFrameSize()

		switch frameIndex {
		case frameTopLeft:
			currentY += frameHeight
			firstFrameWidth = frameWidth
		case frameBottomLeft:
			currentY += frameHeight
		case frameTopMiddleLeft:
			currentX = firstFrameWidth
			currentY = top + frameHeight
		case frameTopMiddleRight:
			currentY = top + frameHeight
			currentX += prevWidth
			currentX += magicHelpBorderOffsetX
		case frameTopRightNoCorner:
			currentY = top + frameHeight
			currentX += prevWidth
		case frameTopRight:
			currentY = top + frameHeight
			currentX += prevWidth
		case frameBottomRight:
			currentY = prevY + frameHeight
		}

		prevY = currentY
		prevWidth = frameWidth

		f.SetPosition(currentX, currentY)

		h.frames = append(h.frames, f)
	}
}

func (h *HelpOverlay) setupTitleAndButton() {
	// Title
	text := h.asset.TranslateString("Strhelp1") // "Diablo II Help"
	newLabel := h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(text)

	titleLabelWidth, _ := newLabel.GetSize()

	newLabel.SetPosition((windowWidth/inHalf)-(titleLabelWidth/inHalf)+titleLabelOffsetX, 0)
	h.text = append(h.text, newLabel)

	// Button
	h.closeButton = h.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	h.closeButton.SetPosition(closeButtonX, closeButtonY)
	h.closeButton.SetVisible(false)
	h.closeButton.OnActivated(func() { h.Close() })
	h.closeButton.SetRenderPriority(d2ui.RenderPriorityForeground)
	h.panelGroup.AddWidget(h.closeButton)

	newLabel = h.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	newLabel.SetText(h.asset.TranslateString("strClose")) // "Close"
	newLabel.SetPosition(closeButtonLabelX, closeButtonLabelY)
	newLabel.Alignment = d2ui.HorizontalAlignCenter
	h.text = append(h.text, newLabel)
}

func (h *HelpOverlay) setupBulletedList() {
	// Bullets
	// the hotkeys displayed here should be pulled from a mapping of input events to game events
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/793
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/794
	callouts := []struct{ text string }{
		// "Ctrl" should be hotkey // "Hold Down <%s> to Run"
		{text: fmt.Sprintf(
			h.asset.TranslateString("StrHelp2"),
			h.keyMap.KeyToString(h.keyMap.GetKeysForGameEvent(d2enum.HoldRun).Primary),
		)},

		// "Alt" should be hotkey // "Hold down <%s> to highlight items on the ground"
		{text: fmt.Sprintf(
			h.asset.TranslateString("StrHelp3"),
			h.keyMap.KeyToString(h.keyMap.GetKeysForGameEvent(d2enum.HoldShowGroundItems).Primary),
		)},

		// "Shift" should be hotkey // "Hold down <%s> to attack while standing still"
		{text: fmt.Sprintf(
			h.asset.TranslateString("StrHelp4"),
			h.keyMap.KeyToString(h.keyMap.GetKeysForGameEvent(d2enum.HoldStandStill).Primary),
		)},

		// "Tab" should be hotkey // "Hit <%s> to toggle the automap on and off"
		{text: fmt.Sprintf(
			h.asset.TranslateString("StrHelp5"),
			h.keyMap.KeyToString(h.keyMap.GetKeysForGameEvent(d2enum.ToggleAutomap).Primary),
		)},

		// "Hit <Esc> to bring up the Game Menu"
		{text: h.asset.TranslateString("StrHelp6")},

		// "Hit <Enter> to go into chat mode"
		{text: h.asset.TranslateString("StrHelp7")},

		// "Hit F1-F8 to set your Left or Right Mouse Buttton Skills."
		{text: h.asset.TranslateString("StrHelp8")},

		// "H" should be hotkey,
		{text: fmt.Sprintf(
			h.asset.TranslateString("StrHelp8a"),
			h.keyMap.KeyToString(h.keyMap.GetKeysForGameEvent(d2enum.ToggleHelpScreen).Primary),
		)},
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
}

// nolint:funlen // can't reduce
func (h *HelpOverlay) setupLabelsWithLines() {
	h.createCallout(callout{
		LabelText: h.asset.TranslateString("strlvlup"), // "New Stats"
		LabelX:    newStatsLabelX,
		LabelY:    newStatsLabelY,
		DotX:      newStatsDotX,
		DotY:      newStatsDotY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("strnewskl"), // "New Skill"
		LabelX:    newSkillLabelX,
		LabelY:    newSkillLabelY,
		DotX:      newSkillDotX,
		DotY:      newSkillDotY,
	})

	// Some of the help fonts require mulktiple lines.
	h.createLabel(callout{
		LabelText: h.asset.TranslateString("StrHelp10"), // "Left Mouse-"
		LabelX:    leftMouseLabelX,
		LabelY:    leftMouseLabelY,
	})

	h.createLabel(callout{
		LabelText: h.asset.TranslateString("StrHelp11"), // "Button Skill"
		LabelX:    leftButtonSkillLabelX,
		LabelY:    leftButtonSkillLabelY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("StrHelp12"), // "(Click to Change)"
		LabelX:    leftSkillClickToChangeLabelX,
		LabelY:    leftSkillClickToChangeLabelY,
		DotX:      leftSkillClickToChangeDotX,
		DotY:      leftSkillClickToChangeDotY,
	})

	h.createLabel(callout{
		LabelText: h.asset.TranslateString("StrHelp13"), // "Right Mouse"
		LabelX:    rightMouseLabelX,
		LabelY:    rightMouseLabelY,
	})

	h.createLabel(callout{
		LabelText: h.asset.TranslateString("StrHelp11"), // "Button Skill"
		LabelX:    rightButtonSkillLabelX,
		LabelY:    rightButtonSkillLabelY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("StrHelp12"), // "(Click to Change)"
		LabelX:    rightSkillClickToChangeLabelX,
		LabelY:    rightSkillClickToChangeLabelY,
		DotX:      rightSkillClickToChangeDotX,
		DotY:      rightSkillClickToChangeDotY,
	})

	h.createLabel(callout{
		LabelText: h.asset.TranslateString("StrHelp17"), // "Mini-Panel"
		LabelX:    miniPanelLabelX,
		LabelY:    miniPanelLabelY,
	})

	h.createLabel(callout{
		LabelText: h.asset.TranslateString("StrHelp18"), // "(Opens Character,"
		LabelX:    characterLabelX,
		LabelY:    characterLabelY,
	})

	h.createLabel(callout{
		LabelText: h.asset.TranslateString("StrHelp19"), // "inventory, and"
		LabelX:    inventoryLabelX,
		LabelY:    inventoryLabelY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("StrHelp20"), // "other screens)"
		LabelX:    otherScreensLabelX,
		LabelY:    otherScreensLabelY,
		DotX:      otherScreensDotX,
		DotY:      otherScreensDotY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("StrHelp9"), // "Life Orb"
		LabelX:    lifeOrbLabelX,
		LabelY:    lifeOrbLabelY,
		DotX:      lifeOrbDotX,
		DotY:      lifeOrbDotY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("StrHelp15"), // "Stamina Bar"
		LabelX:    staminaBarLabelX,
		LabelY:    staminaBarLabelY,
		DotX:      staminaBarDotX,
		DotY:      staminaBarDotY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("StrHelp22"), // "Mana Orb"
		LabelX:    manaOrbLabelX,
		LabelY:    manaOrbLabelY,
		DotX:      manaOrbDotX,
		DotY:      manaOrbDotY,
	})

	h.createLabel(callout{
		LabelText: h.asset.TranslateString("StrHelp14"), // "Run/Walk"
		LabelX:    runWalkButtonLabelX,
		LabelY:    runWalkButtonLabelY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("StrHelp14a"), // "Toggle"
		LabelX:    toggleLabelX,
		LabelY:    toggleLabelY,
		DotX:      toggleDotX,
		DotY:      toggleDotY,
	})

	h.createLabel(callout{
		LabelText: h.asset.TranslateString("StrHelp16"), // "Experience"
		LabelX:    experienceLabelX,
		LabelY:    experienceLabelY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("StrHelp16a"), // "Bar"
		LabelX:    barLabelX,
		LabelY:    barLabelY,
		DotX:      barDotX,
		DotY:      barDotY,
	})

	h.createCallout(callout{
		LabelText: h.asset.TranslateString("StrHelp21"), // "Belt"
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

func (h *HelpOverlay) createBullet(c callout) {
	newLabel := h.uiManager.NewLabel(d2resource.FontFormal12, d2resource.PaletteSky)
	newLabel.SetText(c.LabelText)
	newLabel.SetPosition(c.LabelX, c.LabelY)
	h.text = append(h.text, newLabel)

	newDot, err := h.uiManager.NewSprite(d2resource.HelpYellowBullet, d2resource.PaletteSky)
	if err != nil {
		h.Error(err.Error())
	}

	err = newDot.SetCurrentFrame(0)
	if err != nil {
		h.Error(err.Error())
	}

	newDot.SetPosition(c.DotX, c.DotY+bulletOffsetY)
	h.frames = append(h.frames, newDot)
}

func (h *HelpOverlay) createLabel(c callout) {
	newLabel := h.uiManager.NewLabel(d2resource.FontFormal12, d2resource.PaletteSky)
	newLabel.SetText(c.LabelText)
	newLabel.SetPosition(c.LabelX, c.LabelY)
	h.text = append(h.text, newLabel)
	newLabel.Alignment = d2ui.HorizontalAlignCenter
}

func (h *HelpOverlay) createCallout(c callout) {
	newLabel := h.uiManager.NewLabel(d2resource.FontFormal12, d2resource.PaletteSky)
	newLabel.Color[0] = color.White
	newLabel.SetText(c.LabelText)
	newLabel.SetPosition(c.LabelX, c.LabelY)
	newLabel.Alignment = d2ui.HorizontalAlignCenter
	ww, hh := newLabel.GetTextMetrics(c.LabelText)
	h.text = append(h.text, newLabel)
	_ = ww

	l := line{
		StartX: c.LabelX,
		StartY: c.LabelY + hh + lineOffset,
		MoveX:  0,
		MoveY:  c.DotY - c.LabelY - hh - lineOffset,
		Color:  color.White,
	}

	h.lines = append(h.lines, l)

	newDot, err := h.uiManager.NewSprite(d2resource.HelpWhiteBullet, d2resource.PaletteSky)
	if err != nil {
		h.Error(err.Error())
	}

	err = newDot.SetCurrentFrame(0)
	if err != nil {
		h.Error(err.Error())
	}

	newDot.SetPosition(c.DotX, c.DotY)
	h.frames = append(h.frames, newDot)
}

// Render the overlay to the given surface
func (h *HelpOverlay) Render(target d2interface.Surface) {
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
	}
}
