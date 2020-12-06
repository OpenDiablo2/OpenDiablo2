package d2player

import (
	"fmt"
	"strconv"
	//"strings"

	//"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	//"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	//"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
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

// NewHeroStatsPanel creates a new hero status panel
func NewMoveGoldPanel(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	gold int,
	l d2util.LogLevel,
) *moveGoldPanel {
	originX := 0
	originY := 0

	mgp := &moveGoldPanel{
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

// HeroStatsPanel represents the hero status panel
type moveGoldPanel struct {
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

// Load the data for the hero status panel
func (s *moveGoldPanel) Load() {
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

func (s *moveGoldPanel) incrose() {
	currentValue, err := strconv.Atoi(s.value.GetText())
	if err != nil {
		s.Errorf("Incorrect value in textbox (cannot be converted into intager) %s", err)
	} else {
		if currentValue < s.gold {
			s.value.SetText(fmt.Sprintln(currentValue + 1))
		}
	}
}

func (s *moveGoldPanel) decrose() {
	currentValue, err := strconv.Atoi(s.value.GetText())
	if err != nil {
		s.Errorf("Incorrect value in textbox (cannot be converted into intager) %s", err)
	} else {
		if currentValue > 0 {
			s.value.SetText(fmt.Sprintln(currentValue - 1))
		}
	}
}

func (s *moveGoldPanel) setActionText() {
	dropGoldStr := d2util.SplitIntoLinesWithMaxWidth(s.asset.TranslateString("strDropGoldHowMuch"), 20)
	//if s.isChest {
	if true {
		s.actionLabel1.SetText(d2ui.ColorTokenize(dropGoldStr[0], d2ui.ColorTokenGold))
		s.actionLabel2.SetText(d2ui.ColorTokenize(dropGoldStr[1], d2ui.ColorTokenGold))
	}
}

// IsOpen returns true if the hero status panel is open
func (s *moveGoldPanel) IsOpen() bool {
	return s.isOpen
}

// Toggle toggles the visibility of the hero status panel
func (s *moveGoldPanel) Toggle() {
	if s.isOpen {
		s.Close()
	} else {
		s.Open()
	}
}

// Open opens the hero status panel
func (s *moveGoldPanel) Open() {
	s.isOpen = true
	s.panelGroup.SetVisible(true)
}

// Close closed the hero status panel
func (s *moveGoldPanel) Close() {
	s.isOpen = false
	s.panelGroup.SetVisible(false)
	//s.onCloseCb()
}

/*
// SetOnCloseCb the callback run on closing the HeroStatsPanel
func (s *HeroStatsPanel) SetOnCloseCb(cb func()) {
	s.onCloseCb = cb
}

// Advance updates labels on the panel
func (s *HeroStatsPanel) Advance(elapsed float64) {
	if !s.isOpen {
		return
	}

	s.setStatValues()
}

func (s *HeroStatsPanel) renderStaticMenu(target d2interface.Surface) {
	s.renderStaticPanelFrames(target)
	s.renderStaticLabels(target)
}

func (s *HeroStatsPanel) renderStaticLabels(target d2interface.Surface) {
	var label *d2ui.Label

	fr := strings.Split(s.asset.TranslateString("strchrfir"), "\n")
	lr := strings.Split(s.asset.TranslateString("strchrlit"), "\n")
	cr := strings.Split(s.asset.TranslateString("strchrcol"), "\n")
	pr := strings.Split(s.asset.TranslateString("strchrpos"), "\n")
	// all static labels are not stored since we use them only once to generate the image cache
	var staticLabelConfigs = []struct {
		x, y        int
		txt         string
		font        string
		centerAlign bool
	}{
		{labelHeroNameX, labelHeroNameY, s.heroName, d2resource.Font16, true},
		{labelHeroClassX, labelHeroClassY, s.asset.TranslateString(s.heroClass), d2resource.Font16, true},

		{labelLevelX, labelLevelY, s.asset.TranslateString("strchrlvl"), d2resource.Font6, true},
		{labelExperienceX, labelExperienceY, s.asset.TranslateString("strchrexp"), d2resource.Font6, true},
		{labelNextLevelX, labelNextLevelY, s.asset.TranslateString("strchrnxtlvl"), d2resource.Font6, true},
		{labelStrengthX, labelStrengthY, s.asset.TranslateString("strchrstr"), d2resource.Font6, false},
		{labelDexterityX, labelDexterityY, s.asset.TranslateString("strchrdex"), d2resource.Font6, false},
		{labelVitalityX, labelVitalityY, s.asset.TranslateString("strchrvit"), d2resource.Font6, false},
		{labelEnergyX, labelEnergyY, s.asset.TranslateString("strchreng"), d2resource.Font6, false},
		{labelDefenseX, labelDefenseY, s.asset.TranslateString("strchrdef"), d2resource.Font6, false},
		{labelStaminaX, labelStaminaY, s.asset.TranslateString("strchrstm"), d2resource.Font6, true},
		{labelLifeX, labelLifeY, s.asset.TranslateString("strchrlif"), d2resource.Font6, true},
		{labelManaX, labelManaY, s.asset.TranslateString("strchrman"), d2resource.Font6, true},

		// can't use "Fire\nResistance" because line spacing is too big and breaks the layout
		{labelResFireLine1X, labelResFireLine1Y, fr[0], d2resource.Font6, true},
		{labelResFireLine2X, labelResFireLine2Y, fr[len(fr)-1], d2resource.Font6, true},

		{labelResColdLine1X, labelResColdLine1Y, cr[0], d2resource.Font6, true},
		{labelResColdLine2X, labelResColdLine2Y, cr[len(cr)-1], d2resource.Font6, true},

		{labelResLightLine1X, labelResLightLine1Y, lr[0], d2resource.Font6, true},
		{labelResLightLine2X, labelResLightLine2Y, lr[len(lr)-1], d2resource.Font6, true},

		{labelResPoisLine1X, labelResPoisLine1Y, pr[0], d2resource.Font6, true},
		{labelResPoisLine2X, labelResPoisLine2Y, pr[len(pr)-1], d2resource.Font6, true},
	}

	for _, cfg := range staticLabelConfigs {
		label = s.createTextLabel(PanelText{
			cfg.x, cfg.y,
			cfg.txt,
			cfg.font,
			cfg.centerAlign,
		})

		label.Render(target)
	}
}

func (s *HeroStatsPanel) initStatValueLabels() {
	valueLabelConfigs := []struct {
		assignTo **d2ui.Label
		value    int
		x, y     int
	}{
		{&s.labels.Level, s.heroState.Level, 112, 110},
		{&s.labels.Experience, s.heroState.Experience, 200, 110},
		{&s.labels.NextLevelExp, s.heroState.NextLevelExp, 330, 110},
		{&s.labels.Strength, s.heroState.Strength, 175, 147},
		{&s.labels.Dexterity, s.heroState.Dexterity, 175, 207},
		{&s.labels.Vitality, s.heroState.Vitality, 175, 295},
		{&s.labels.Energy, s.heroState.Energy, 175, 355},
		{&s.labels.MaxStamina, s.heroState.MaxStamina, 330, 295},
		{&s.labels.Stamina, int(s.heroState.Stamina), 370, 295},
		{&s.labels.MaxHealth, s.heroState.MaxHealth, 330, 320},
		{&s.labels.Health, s.heroState.Health, 370, 320},
		{&s.labels.MaxMana, s.heroState.MaxMana, 330, 355},
		{&s.labels.Mana, s.heroState.Mana, 370, 355},
	}

	for _, cfg := range valueLabelConfigs {
		*cfg.assignTo = s.createStatValueLabel(cfg.value, cfg.x, cfg.y)
	}
}

func (s *HeroStatsPanel) setStatValues() {
	s.labels.Level.SetText(strconv.Itoa(s.heroState.Level))
	s.labels.Experience.SetText(strconv.Itoa(s.heroState.Experience))
	s.labels.NextLevelExp.SetText(strconv.Itoa(s.heroState.NextLevelExp))

	s.labels.Strength.SetText(strconv.Itoa(s.heroState.Strength))
	s.labels.Dexterity.SetText(strconv.Itoa(s.heroState.Dexterity))
	s.labels.Vitality.SetText(strconv.Itoa(s.heroState.Vitality))
	s.labels.Energy.SetText(strconv.Itoa(s.heroState.Energy))

	s.labels.MaxHealth.SetText(strconv.Itoa(s.heroState.MaxHealth))
	s.labels.Health.SetText(strconv.Itoa(s.heroState.Health))

	s.labels.MaxStamina.SetText(strconv.Itoa(s.heroState.MaxStamina))
	s.labels.Stamina.SetText(strconv.Itoa(int(s.heroState.Stamina)))

	s.labels.MaxMana.SetText(strconv.Itoa(s.heroState.MaxMana))
	s.labels.Mana.SetText(strconv.Itoa(s.heroState.Mana))
}

func (s *HeroStatsPanel) createStatValueLabel(stat, x, y int) *d2ui.Label {
	text := strconv.Itoa(stat)
	return s.createTextLabel(PanelText{X: x, Y: y, Text: text, Font: d2resource.Font16, AlignCenter: true})
}

func (s *HeroStatsPanel) createTextLabel(element PanelText) *d2ui.Label {
	label := s.uiManager.NewLabel(element.Font, d2resource.PaletteStatic)
	if element.AlignCenter {
		label.Alignment = d2ui.HorizontalAlignCenter
	}

	label.SetText(element.Text)
	label.SetPosition(element.X, element.Y)
	s.panelGroup.AddWidget(label)

	return label
}*/
