package d2player

import (
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const ( // for the dc6 frames
	statsPanelTopLeft = iota
	statsPanelTopRight
	statsPanelBottomLeft
	statsPanelBottomRight
)

const (
	statsPanelOffsetX, statsPanelOffsetY = 80, 64
)

const (
	labelLevelX, labelLevelY = 110, 100

	labelHeroNameX, labelHeroNameY   = 165, 72
	labelHeroClassX, labelHeroClassY = 330, 74

	labelExperienceX, labelExperienceY = 200, 100
	labelNextLevelX, labelNextLevelY   = 330, 100

	labelStrengthX, labelStrengthY   = 100, 150
	labelDexterityX, labelDexterityY = 100, 213
	labelVitalityX, labelVitalityY   = 95, 300
	labelEnergyX, labelEnergyY       = 100, 360

	labelDefenseX, labelDefenseY = 280, 260
	labelStaminaX, labelStaminaY = 280, 300
	labelLifeX, labelLifeY       = 280, 322
	labelManaX, labelManaY       = 280, 360

	labelResFireLine1X, labelResFireLine1Y   = 310, 396
	labelResFireLine2X, labelResFireLine2Y   = 310, 403
	labelResColdLine1X, labelResColdLine1Y   = 310, 444
	labelResColdLine2X, labelResColdLine2Y   = 310, 452
	labelResLightLine1X, labelResLightLine1Y = 310, 420
	labelResLightLine2X, labelResLightLine2Y = 310, 428
	labelResPoisLine1X, labelResPoisLine1Y   = 310, 468
	labelResPoisLine2X, labelResPoisLine2Y   = 310, 476
)

const (
	heroStatsCloseButtonX, heroStatsCloseButtonY = 208, 453
)

// PanelText represents text on the panel
type PanelText struct {
	X           int
	Y           int
	Text        string
	Font        string
	AlignCenter bool
}

// StatsPanelLabels represents the labels in the status panel
type StatsPanelLabels struct {
	Level        *d2ui.Label
	Experience   *d2ui.Label
	NextLevelExp *d2ui.Label
	Strength     *d2ui.Label
	Dexterity    *d2ui.Label
	Vitality     *d2ui.Label
	Energy       *d2ui.Label
	Health       *d2ui.Label
	MaxHealth    *d2ui.Label
	Mana         *d2ui.Label
	MaxMana      *d2ui.Label
	MaxStamina   *d2ui.Label
	Stamina      *d2ui.Label
}

// NewHeroStatsPanel creates a new hero status panel
func NewHeroStatsPanel(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	heroName string,
	heroClass d2enum.Hero,
	l d2util.LogLevel,
	heroState *d2hero.HeroStatsState) *HeroStatsPanel {
	originX := 0
	originY := 0

	hsp := &HeroStatsPanel{
		asset:     asset,
		uiManager: ui,
		originX:   originX,
		originY:   originY,
		heroState: heroState,
		heroName:  heroName,
		heroClass: heroClass,
		labels:    &StatsPanelLabels{},
	}

	hsp.Logger = d2util.NewLogger()
	hsp.Logger.SetLevel(l)
	hsp.Logger.SetPrefix(logPrefix)

	return hsp
}

// HeroStatsPanel represents the hero status panel
type HeroStatsPanel struct {
	asset      *d2asset.AssetManager
	uiManager  *d2ui.UIManager
	panel      *d2ui.Sprite
	heroState  *d2hero.HeroStatsState
	heroName   string
	heroClass  d2enum.Hero
	labels     *StatsPanelLabels
	onCloseCb  func()
	panelGroup *d2ui.WidgetGroup

	originX int
	originY int
	isOpen  bool

	*d2util.Logger
}

// Load the data for the hero status panel
func (s *HeroStatsPanel) Load() {
	var err error

	s.panelGroup = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)

	frame := d2ui.NewUIFrame(s.asset, s.uiManager, d2ui.FrameLeft)
	s.panelGroup.AddWidget(frame)

	s.panel, err = s.uiManager.NewSprite(d2resource.InventoryCharacterPanel, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	w, h := frame.GetSize()
	staticPanel := s.uiManager.NewCustomWidgetCached(s.renderStaticMenu, w, h)
	s.panelGroup.AddWidget(staticPanel)

	closeButton := s.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	closeButton.SetVisible(false)
	closeButton.SetPosition(heroStatsCloseButtonX, heroStatsCloseButtonY)
	closeButton.OnActivated(func() { s.Close() })
	s.panelGroup.AddWidget(closeButton)

	s.initStatValueLabels()
	s.panelGroup.SetVisible(false)
}

// IsOpen returns true if the hero status panel is open
func (s *HeroStatsPanel) IsOpen() bool {
	return s.isOpen
}

// Toggle toggles the visibility of the hero status panel
func (s *HeroStatsPanel) Toggle() {
	if s.isOpen {
		s.Close()
	} else {
		s.Open()
	}
}

// Open opens the hero status panel
func (s *HeroStatsPanel) Open() {
	s.isOpen = true
	s.panelGroup.SetVisible(true)
}

// Close closed the hero status panel
func (s *HeroStatsPanel) Close() {
	s.isOpen = false
	s.panelGroup.SetVisible(false)
	s.onCloseCb()
}

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

// nolint:dupl // see quest_log.go.renderStaticPanelFrames comment
func (s *HeroStatsPanel) renderStaticPanelFrames(target d2interface.Surface) {
	frames := []int{
		statsPanelTopLeft,
		statsPanelTopRight,
		statsPanelBottomRight,
		statsPanelBottomLeft,
	}

	currentX := s.originX + statsPanelOffsetX
	currentY := s.originY + statsPanelOffsetY

	for _, frameIndex := range frames {
		if err := s.panel.SetCurrentFrame(frameIndex); err != nil {
			s.Error(err.Error())
		}

		w, h := s.panel.GetCurrentFrameSize()

		switch frameIndex {
		case statsPanelTopLeft:
			s.panel.SetPosition(currentX, currentY+h)
			currentX += w
		case statsPanelTopRight:
			s.panel.SetPosition(currentX, currentY+h)
			currentY += h
		case statsPanelBottomRight:
			s.panel.SetPosition(currentX, currentY+h)
		case statsPanelBottomLeft:
			s.panel.SetPosition(currentX-w, currentY+h)
		}

		s.panel.Render(target)
	}
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
}
