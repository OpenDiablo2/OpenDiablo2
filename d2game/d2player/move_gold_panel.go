package d2player

import (
	"fmt"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	moveGoldX, moveGoldY                       = 300, 350
	moveGoldCloseButtonX, moveGoldCloseButtonY = moveGoldX + 35, moveGoldY - 42
	moveGoldOkButtonX, moveGoldOkButtonY       = moveGoldX + 140, moveGoldY - 42
	moveGoldValueX, moveGoldValueY             = moveGoldX + 29, moveGoldY - 90
	moveGoldActionLabelX, moveGoldActionLabelY = moveGoldX + 105, moveGoldY - 150
	moveGoldActionLabelOffsetY                 = 25
	moveGoldUpArrowX, moveGoldUpArrowY         = 314, 259
	moveGoldDownArrowX, moveGoldDownArrowY     = 314, 274
)

const goldValueFilter = "0123456789"

// NewMoveGoldPanel creates a new move gold panel
func NewMoveGoldPanel(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	gold int,
	l d2util.LogLevel,
) *MoveGoldPanel {
	originX := 0
	originY := 0

	mgp := &MoveGoldPanel{
		asset:     asset,
		uiManager: ui,
		originX:   originX,
		originY:   originY,
		gold:      gold,
	}

	mgp.Logger = d2util.NewLogger()
	mgp.Logger.SetLevel(l)
	mgp.Logger.SetPrefix(logPrefix)

	return mgp
}

// MoveGoldPanel represents the move gold panel
type MoveGoldPanel struct {
	asset        *d2asset.AssetManager
	uiManager    *d2ui.UIManager
	panel        *d2ui.Sprite
	onCloseCb    func()
	panelGroup   *d2ui.WidgetGroup
	gold         int
	actionLabel1 *d2ui.Label
	actionLabel2 *d2ui.Label
	value        *d2ui.TextBox

	originX int
	originY int
	isOpen  bool

	*d2util.Logger
}

// Load the data for the move gold panel
func (s *MoveGoldPanel) Load() {
	var err error

	s.panelGroup = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityInventory)

	s.panel, err = s.uiManager.NewSprite(d2resource.MoveGoldDialog, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	s.panel.SetPosition(moveGoldX, moveGoldY)
	s.panelGroup.AddWidget(s.panel)

	closeButton := s.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	closeButton.SetVisible(false)
	closeButton.SetPosition(moveGoldCloseButtonX, moveGoldCloseButtonY)
	closeButton.OnActivated(func() { s.Close() })
	s.panelGroup.AddWidget(closeButton)

	okButton := s.uiManager.NewButton(d2ui.ButtonTypeSquareOk, "")
	okButton.SetVisible(false)
	okButton.SetPosition(moveGoldOkButtonX, moveGoldOkButtonY)
	okButton.OnActivated(func() { s.Close() })
	s.panelGroup.AddWidget(okButton)

	s.value = s.uiManager.NewTextbox()
	s.value.SetFilter(goldValueFilter)
	s.value.SetText(fmt.Sprintln(s.gold))
	s.value.Activate()
	s.value.SetNumberOnly(s.gold)
	s.value.SetPosition(moveGoldValueX, moveGoldValueY)
	s.panelGroup.AddWidget(s.value)

	s.actionLabel1 = s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteStatic)
	s.actionLabel1.Alignment = d2ui.HorizontalAlignCenter
	s.actionLabel1.SetPosition(moveGoldActionLabelX, moveGoldActionLabelY)
	s.panelGroup.AddWidget(s.actionLabel1)

	s.actionLabel2 = s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteStatic)
	s.actionLabel2.Alignment = d2ui.HorizontalAlignCenter
	s.actionLabel2.SetPosition(moveGoldActionLabelX, moveGoldActionLabelY+moveGoldActionLabelOffsetY)
	s.panelGroup.AddWidget(s.actionLabel2)

	incrose := s.uiManager.NewButton(d2ui.ButtonTypeUpArrow, d2resource.PaletteSky)
	incrose.SetPosition(moveGoldUpArrowX, moveGoldUpArrowY)
	incrose.SetVisible(false)
	incrose.OnActivated(func() { s.incrose() })
	s.panelGroup.AddWidget(incrose)

	decrose := s.uiManager.NewButton(d2ui.ButtonTypeDownArrow, d2resource.PaletteSky)
	decrose.SetPosition(moveGoldDownArrowX, moveGoldDownArrowY)
	decrose.SetVisible(false)
	decrose.OnActivated(func() { s.decrose() })
	s.panelGroup.AddWidget(decrose)

	s.setActionText()

	s.panelGroup.SetVisible(false)
}

func (s *MoveGoldPanel) incrose() {
	currentValue, err := strconv.Atoi(s.value.GetText())
	if err != nil {
		s.Errorf("Incorrect value in textbox (cannot be converted into intager) %s", err)
		return
	}

	if currentValue < s.gold {
		s.value.SetText(fmt.Sprintln(currentValue + 1))
	}
}

func (s *MoveGoldPanel) decrose() {
	currentValue, err := strconv.Atoi(s.value.GetText())
	if err != nil {
		s.Errorf("Incorrect value in textbox (cannot be converted into intager) %s", err)
		return
	}

	if currentValue > 0 {
		s.value.SetText(fmt.Sprintln(currentValue - 1))
	}
}

func (s *MoveGoldPanel) setActionText() {
	dropGoldStr := d2util.SplitIntoLinesWithMaxWidth(s.asset.TranslateString("strDropGoldHowMuch"), 20)
	//if s.isChest {
	if true {
		s.actionLabel1.SetText(d2ui.ColorTokenize(dropGoldStr[0], d2ui.ColorTokenGold))
		s.actionLabel2.SetText(d2ui.ColorTokenize(dropGoldStr[1], d2ui.ColorTokenGold))
	}
}

// IsOpen returns true if the move gold panel is opened
func (s *MoveGoldPanel) IsOpen() bool {
	return s.isOpen
}

// Toggle toggles the visibility of the move gold panel
func (s *MoveGoldPanel) Toggle() {
	if s.isOpen {
		s.Close()
	} else {
		s.Open()
	}
}

// Open opens the move gold panel
func (s *MoveGoldPanel) Open() {
	s.isOpen = true
	s.panelGroup.SetVisible(true)
}

// Close closed the move gold panel
func (s *MoveGoldPanel) Close() {
	s.isOpen = false
	s.panelGroup.SetVisible(false)
	s.onCloseCb()
}

// SetOnCloseCb the callback run on closing the HeroStatsPanel
func (s *MoveGoldPanel) SetOnCloseCb(cb func()) {
	s.onCloseCb = cb
}
