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

// PanelText represents text on the panel
type PanelText struct {
	X           int
	Y           int
	Height      int
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
	frame                *d2ui.Sprite
	panel                *d2ui.Sprite
	heroState            *d2hero.HeroStatsState
	heroName             string
	heroClass            d2enum.Hero
	renderer             d2interface.Renderer
	staticMenuImageCache *d2interface.Surface
	labels               *StatsPanelLabels

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
	s.frame, err = s.uiManager.NewSprite(d2resource.Frame, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

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
	s.isOpen = !s.isOpen
}

// Open opens the hero status panel
func (s *HeroStatsPanel) Open() {
	s.isOpen = true
}

// Close closed the hero status panel
func (s *HeroStatsPanel) Close() {
	s.isOpen = false
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
	x, y := s.originX, s.originY

	// Frame
	// Top left
	if err := s.frame.SetCurrentFrame(0); err != nil {
		return err
	}

	w, h := s.frame.GetCurrentFrameSize()

	s.frame.SetPosition(x, y+h)

	if err := s.frame.Render(target); err != nil {
		return err
	}

	x += w
	y += h

	// Top right
	if err := s.frame.SetCurrentFrame(1); err != nil {
		return err
	}

	_, h = s.frame.GetCurrentFrameSize()

	s.frame.SetPosition(x, s.originY+h)

	if err := s.frame.Render(target); err != nil {
		return err
	}

	x = s.originX

	// Right
	if err := s.frame.SetCurrentFrame(2); err != nil {
		return err
	}

	_, h = s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, y+h)

	if err := s.frame.Render(target); err != nil {
		return err
	}

	y += h

	// Bottom left
	if err := s.frame.SetCurrentFrame(3); err != nil {
		return err
	}

	w, h = s.frame.GetCurrentFrameSize()

	s.frame.SetPosition(x, y+h)

	if err := s.frame.Render(target); err != nil {
		return err
	}

	x += w

	// Bottom right
	if err := s.frame.SetCurrentFrame(4); err != nil {
		return err
	}

	_, h = s.frame.GetCurrentFrameSize()

	s.frame.SetPosition(x, y+h)

	if err := s.frame.Render(target); err != nil {
		return err
	}

	x, y = s.originX, s.originY
	y += 64
	x += 80

	// Panel
	// Top left
	if err := s.panel.SetCurrentFrame(0); err != nil {
		return err
	}

	w, h = s.panel.GetCurrentFrameSize()

	s.panel.SetPosition(x, y+h)

	if err := s.panel.Render(target); err != nil {
		return err
	}

	x += w

	// Top right
	if err := s.panel.SetCurrentFrame(1); err != nil {
		return err
	}

	_, h = s.panel.GetCurrentFrameSize()

	s.panel.SetPosition(x, y+h)

	if err := s.panel.Render(target); err != nil {
		return err
	}

	y += h

	// Bottom right
	if err := s.panel.SetCurrentFrame(3); err != nil {
		return err
	}

	_, h = s.panel.GetCurrentFrameSize()

	s.panel.SetPosition(x, y+h)

	if err := s.panel.Render(target); err != nil {
		return err
	}

	// Bottom left
	if err := s.panel.SetCurrentFrame(2); err != nil {
		return err
	}

	w, h = s.panel.GetCurrentFrameSize()

	s.panel.SetPosition(x-w, y+h)

	if err := s.panel.Render(target); err != nil {
		return err
	}

	var label *d2ui.Label

	// all static labels are not stored since we use them only once to generate the image cache

	//nolint:gomnd
	var staticTextLabels = []PanelText{
		{X: 110, Y: 100, Text: "Level", Font: d2resource.Font6, AlignCenter: true},
		{X: 200, Y: 100, Text: "Experience", Font: d2resource.Font6, AlignCenter: true},
		{X: 330, Y: 100, Text: "Next Level", Font: d2resource.Font6, AlignCenter: true},
		{X: 100, Y: 150, Text: "Strength", Font: d2resource.Font6},
		{X: 100, Y: 213, Text: "Dexterity", Font: d2resource.Font6},
		{X: 100, Y: 300, Text: "Vitality", Font: d2resource.Font6},
		{X: 100, Y: 360, Text: "Energy", Font: d2resource.Font6},
		{X: 280, Y: 260, Text: "Defense", Font: d2resource.Font6},
		{X: 280, Y: 300, Text: "Stamina", Font: d2resource.Font6, AlignCenter: true},
		{X: 280, Y: 322, Text: "Life", Font: d2resource.Font6, AlignCenter: true},
		{X: 280, Y: 360, Text: "Mana", Font: d2resource.Font6, AlignCenter: true},

		// can't use "Fire\nResistance" because line spacing is too big and breaks the layout
		{X: 310, Y: 395, Text: "Fire", Font: d2resource.Font6, AlignCenter: true},
		{X: 310, Y: 402, Text: "Resistance", Font: d2resource.Font6, AlignCenter: true},

		{X: 310, Y: 420, Text: "Cold", Font: d2resource.Font6, AlignCenter: true},
		{X: 310, Y: 427, Text: "Resistance", Font: d2resource.Font6, AlignCenter: true},

		{X: 310, Y: 445, Text: "Lightning", Font: d2resource.Font6, AlignCenter: true},
		{X: 310, Y: 452, Text: "Resistance", Font: d2resource.Font6, AlignCenter: true},

		{X: 310, Y: 468, Text: "Poison", Font: d2resource.Font6, AlignCenter: true},
		{X: 310, Y: 477, Text: "Resistance", Font: d2resource.Font6, AlignCenter: true},
	}

	for _, textElement := range staticTextLabels {
		label = s.createTextLabel(textElement)
		label.Render(target)
	}
	// hero name and class are part of the static image cache since they don't change after we enter the world
	label = s.createTextLabel(PanelText{X: 165, Y: 72, Text: s.heroName, Font: d2resource.Font16, AlignCenter: true})

	label.Render(target)

	label = s.createTextLabel(PanelText{X: 330, Y: 72, Text: s.heroClass.String(), Font: d2resource.Font16, AlignCenter: true})

	label.Render(target)

	return nil
}

func (s *HeroStatsPanel) initStatValueLabels() {
	s.labels.Level = s.createStatValueLabel(s.heroState.Level, 112, 110)
	s.labels.Experience = s.createStatValueLabel(s.heroState.Experience, 200, 110)
	s.labels.NextLevelExp = s.createStatValueLabel(s.heroState.NextLevelExp, 330, 110)

	s.labels.Strength = s.createStatValueLabel(s.heroState.Strength, 175, 147)
	s.labels.Dexterity = s.createStatValueLabel(s.heroState.Dexterity, 175, 207)
	s.labels.Vitality = s.createStatValueLabel(s.heroState.Vitality, 175, 295)
	s.labels.Energy = s.createStatValueLabel(s.heroState.Energy, 175, 355)

	s.labels.MaxStamina = s.createStatValueLabel(s.heroState.MaxStamina, 330, 295)
	s.labels.Stamina = s.createStatValueLabel(s.heroState.Stamina, 370, 295)

	s.labels.MaxHealth = s.createStatValueLabel(s.heroState.MaxHealth, 330, 320)
	s.labels.Health = s.createStatValueLabel(s.heroState.Health, 370, 320)

	s.labels.MaxMana = s.createStatValueLabel(s.heroState.MaxMana, 330, 355)
	s.labels.Mana = s.createStatValueLabel(s.heroState.Mana, 370, 355)
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
	s.renderStatValueNum(s.labels.Stamina, s.heroState.Stamina, target)

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
