package d2player

import (
	"log"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
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
	labelHeroClassX, labelHeroClassY = 330, 72

	labelExperienceX, labelExperienceY = 200, 100
	labelNextLevelX, labelNextLevelY   = 330, 100

	labelStrengthX, labelStrengthY   = 100, 150
	labelDexterityX, labelDexterityY = 100, 213
	labelVitalityX, labelVitalityY   = 100, 300
	labelEnergyX, labelEnergyY       = 100, 360

	labelDefenseX, labelDefenseY = 280, 260
	labelStaminaX, labelStaminaY = 280, 300
	labelLifeX, labelLifeY       = 280, 322
	labelManaX, labelManaY       = 280, 360

	labelResFireLine1X, labelResFireLine1Y   = 310, 395
	labelResFireLine2X, labelResFireLine2Y   = 310, 402
	labelResColdLine1X, labelResColdLine1Y   = 310, 445
	labelResColdLine2X, labelResColdLine2Y   = 310, 452
	labelResLightLine1X, labelResLightLine1Y = 310, 420
	labelResLightLine2X, labelResLightLine2Y = 310, 427
	labelResPoisLine1X, labelResPoisLine1Y   = 310, 468
	labelResPoisLine2X, labelResPoisLine2Y   = 310, 477
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

// HeroStatsPanel represents the hero status panel
type HeroStatsPanel struct {
	asset                *d2asset.AssetManager
	uiManager            *d2ui.UIManager
	frame                *d2ui.UIFrame
	panel                *d2ui.Sprite
	heroState            *d2hero.HeroStatsState
	heroName             string
	heroClass            d2enum.Hero
	renderer             d2interface.Renderer
	staticMenuImageCache *d2interface.Surface
	labels               *StatsPanelLabels
	closeButton          *d2ui.Button
	onCloseCb            func()

	originX int
	originY int
	isOpen  bool
}

// NewHeroStatsPanel creates a new hero status panel
func NewHeroStatsPanel(asset *d2asset.AssetManager, ui *d2ui.UIManager, heroName string, heroClass d2enum.Hero,
	heroState *d2hero.HeroStatsState) *HeroStatsPanel {
	originX := 0
	originY := 0

	return &HeroStatsPanel{
		asset:     asset,
		uiManager: ui,
		renderer:  ui.Renderer(),
		originX:   originX,
		originY:   originY,
		heroState: heroState,
		heroName:  heroName,
		heroClass: heroClass,
		labels:    &StatsPanelLabels{},
	}
}

// Load the data for the hero status panel
func (s *HeroStatsPanel) Load() {
	var err error

	s.frame = d2ui.NewUIFrame(s.asset, s.uiManager, d2ui.FrameLeft)

	s.closeButton = s.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	s.closeButton.SetVisible(false)
	s.closeButton.SetPosition(heroStatsCloseButtonX, heroStatsCloseButtonY)
	s.closeButton.OnActivated(func() { s.Close() })

	s.panel, err = s.uiManager.NewSprite(d2resource.InventoryCharacterPanel, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	s.initStatValueLabels()
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
	s.closeButton.SetVisible(true)
}

// Close closed the hero status panel
func (s *HeroStatsPanel) Close() {
	s.isOpen = false
	s.closeButton.SetVisible(false)
	s.onCloseCb()
}

// SetOnCloseCb the callback run on closing the HeroStatsPanel
func (s *HeroStatsPanel) SetOnCloseCb(cb func()) {
	s.onCloseCb = cb
}

// Render renders the hero status panel
func (s *HeroStatsPanel) Render(target d2interface.Surface) error {
	if !s.isOpen {
		return nil
	}

	if s.staticMenuImageCache == nil {
		frameWidth, frameHeight := s.frame.GetFrameBounds()
		framesCount := s.frame.GetFrameCount()
		surface, err := s.renderer.NewSurface(frameWidth*framesCount, frameHeight*framesCount, d2enum.FilterNearest)

		if err != nil {
			return err
		}

		s.staticMenuImageCache = &surface

		if err := s.renderStaticMenu(*s.staticMenuImageCache); err != nil {
			return err
		}
	}

	if err := target.Render(*s.staticMenuImageCache); err != nil {
		return err
	}

	s.renderStatValues(target)

	return nil
}

func (s *HeroStatsPanel) renderStaticMenu(target d2interface.Surface) error {
	if err := s.renderStaticPanelFrames(target); err != nil {
		return err
	}

	s.renderStaticLabels(target)

	return nil
}

func (s *HeroStatsPanel) renderStaticPanelFrames(target d2interface.Surface) error {
	if err := s.frame.Render(target); err != nil {
		return err
	}

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
			return err
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

		if err := s.panel.Render(target); err != nil {
			return err
		}
	}

	return nil
}

func (s *HeroStatsPanel) renderStaticLabels(target d2interface.Surface) {
	var label *d2ui.Label

	// all static labels are not stored since we use them only once to generate the image cache
	var staticLabelConfigs = []struct {
		x, y        int
		txt         string
		font        string
		centerAlign bool
	}{
		{labelHeroNameX, labelHeroNameY, s.heroName, d2resource.Font16, true},
		{labelHeroClassX, labelHeroClassY, s.heroClass.String(), d2resource.Font16, true},

		{labelLevelX, labelLevelY, "Level", d2resource.Font6, true},
		{labelExperienceX, labelExperienceY, "Experience", d2resource.Font6, true},
		{labelNextLevelX, labelNextLevelY, "Next Level", d2resource.Font6, true},
		{labelStrengthX, labelStrengthY, "Strength", d2resource.Font6, false},
		{labelDexterityX, labelDexterityY, "Dexterity", d2resource.Font6, false},
		{labelVitalityX, labelVitalityY, "Vitality", d2resource.Font6, false},
		{labelEnergyX, labelEnergyY, "Energy", d2resource.Font6, false},
		{labelDefenseX, labelDefenseY, "Defense", d2resource.Font6, false},
		{labelStaminaX, labelStaminaY, "Stamina", d2resource.Font6, true},
		{labelLifeX, labelLifeY, "Life", d2resource.Font6, true},
		{labelManaX, labelManaY, "Mana", d2resource.Font6, true},

		// can't use "Fire\nResistance" because line spacing is too big and breaks the layout
		{labelResFireLine1X, labelResFireLine1Y, "Fire", d2resource.Font6, true},
		{labelResFireLine2X, labelResFireLine2Y, "Resistance", d2resource.Font6, true},

		{labelResColdLine1X, labelResColdLine1Y, "Cold", d2resource.Font6, true},
		{labelResColdLine2X, labelResColdLine2Y, "Resistance", d2resource.Font6, true},

		{labelResLightLine1X, labelResLightLine1Y, "Lightning", d2resource.Font6, true},
		{labelResLightLine2X, labelResLightLine2Y, "Resistance", d2resource.Font6, true},

		{labelResPoisLine1X, labelResPoisLine1Y, "Poison", d2resource.Font6, true},
		{labelResPoisLine2X, labelResPoisLine2Y, "Resistance", d2resource.Font6, true},
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

func (s *HeroStatsPanel) renderStatValues(target d2interface.Surface) {
	s.renderStatValueNum(s.labels.Level, s.heroState.Level, target)
	s.renderStatValueNum(s.labels.Experience, s.heroState.Experience, target)
	s.renderStatValueNum(s.labels.NextLevelExp, s.heroState.NextLevelExp, target)

	s.renderStatValueNum(s.labels.Strength, s.heroState.Strength, target)
	s.renderStatValueNum(s.labels.Dexterity, s.heroState.Dexterity, target)
	s.renderStatValueNum(s.labels.Vitality, s.heroState.Vitality, target)
	s.renderStatValueNum(s.labels.Energy, s.heroState.Energy, target)

	s.renderStatValueNum(s.labels.MaxHealth, s.heroState.MaxHealth, target)
	s.renderStatValueNum(s.labels.Health, s.heroState.Health, target)

	s.renderStatValueNum(s.labels.MaxStamina, s.heroState.MaxStamina, target)
	s.renderStatValueNum(s.labels.Stamina, int(s.heroState.Stamina), target)

	s.renderStatValueNum(s.labels.MaxMana, s.heroState.MaxMana, target)
	s.renderStatValueNum(s.labels.Mana, s.heroState.Mana, target)
}

func (s *HeroStatsPanel) renderStatValueNum(label *d2ui.Label, value int,
	target d2interface.Surface) {
	label.SetText(strconv.Itoa(value))
	label.Render(target)
}

func (s *HeroStatsPanel) createStatValueLabel(stat, x, y int) *d2ui.Label {
	text := strconv.Itoa(stat)
	return s.createTextLabel(PanelText{X: x, Y: y, Text: text, Font: d2resource.Font16, AlignCenter: true})
}

func (s *HeroStatsPanel) createTextLabel(element PanelText) *d2ui.Label {
	label := s.uiManager.NewLabel(element.Font, d2resource.PaletteStatic)
	if element.AlignCenter {
		label.Alignment = d2gui.HorizontalAlignCenter
	}

	label.SetText(element.Text)
	label.SetPosition(element.X, element.Y)

	return label
}
