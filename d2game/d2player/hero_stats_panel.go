package d2player

import (
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type PanelText struct {
	X           int
	Y           int
	Height      int
	Text        string
	Font        string
	AlignCenter bool
}

type StatsPanelLabels struct {
	Level        d2ui.Label
	Experience   d2ui.Label
	NextLevelExp d2ui.Label
	Strength     d2ui.Label
	Dexterity    d2ui.Label
	Vitality     d2ui.Label
	Energy       d2ui.Label
	Health       d2ui.Label
	MaxHealth    d2ui.Label
	Mana         d2ui.Label
	MaxMana      d2ui.Label
	MaxStamina   d2ui.Label
	Stamina      d2ui.Label
}

var StaticTextLabels = []PanelText{
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

// stores all the labels that can change during gameplay(e.g. current level, current hp, mana, etc.)
var StatValueLabels = make([]d2ui.Label, 13)

type HeroStatsPanel struct {
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

func NewHeroStatsPanel(renderer d2interface.Renderer, heroName string, heroClass d2enum.Hero,
	heroState d2hero.HeroStatsState) *HeroStatsPanel {
	originX := 0
	originY := 0

	return &HeroStatsPanel{
		renderer:  renderer,
		originX:   originX,
		originY:   originY,
		heroState: &heroState,
		heroName:  heroName,
		heroClass: heroClass,
		labels:    &StatsPanelLabels{},
	}
}

func (s *HeroStatsPanel) Load() {
	animation, _ := d2asset.LoadAnimation(d2resource.Frame, d2resource.PaletteSky)
	s.frame, _ = d2ui.LoadSprite(animation)
	animation, _ = d2asset.LoadAnimation(d2resource.InventoryCharacterPanel, d2resource.PaletteSky)
	s.panel, _ = d2ui.LoadSprite(animation)
	s.initStatValueLabels()
}

func (s *HeroStatsPanel) IsOpen() bool {
	return s.isOpen
}

func (s *HeroStatsPanel) Toggle() {
	s.isOpen = !s.isOpen
}

func (s *HeroStatsPanel) Open() {
	s.isOpen = true
}

func (s *HeroStatsPanel) Close() {
	s.isOpen = false
}

func (s *HeroStatsPanel) Render(target d2interface.Surface) {
	if !s.isOpen {
		return
	}

	if s.staticMenuImageCache == nil {
		frameWidth, frameHeight := s.frame.GetFrameBounds()
		framesCount := s.frame.GetFrameCount()
		surface, err := s.renderer.NewSurface(frameWidth*framesCount, frameHeight*framesCount, d2interface.FilterNearest)

		if err != nil {
			return
		}
		s.staticMenuImageCache = &surface
		s.renderStaticMenu(*s.staticMenuImageCache)
	}
	target.Render(*s.staticMenuImageCache)
	s.renderStatValues(target)
}

func (s *HeroStatsPanel) renderStaticMenu(target d2interface.Surface) {
	x, y := s.originX, s.originY

	// Frame
	// Top left
	s.frame.SetCurrentFrame(0)
	w, h := s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, y+h)
	s.frame.Render(target)
	x += w
	y += h

	// Top right
	s.frame.SetCurrentFrame(1)
	w, h = s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, s.originY+h)
	s.frame.Render(target)
	x = s.originX

	// Right
	s.frame.SetCurrentFrame(2)
	w, h = s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, y+h)
	s.frame.Render(target)
	y += h

	// Bottom left
	s.frame.SetCurrentFrame(3)
	w, h = s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, y+h)
	s.frame.Render(target)
	x += w

	// Bottom right
	s.frame.SetCurrentFrame(4)
	w, h = s.frame.GetCurrentFrameSize()
	s.frame.SetPosition(x, y+h)
	s.frame.Render(target)

	x, y = s.originX, s.originY
	y += 64
	x += 80

	// Panel
	// Top left
	s.panel.SetCurrentFrame(0)
	w, h = s.panel.GetCurrentFrameSize()
	s.panel.SetPosition(x, y+h)
	s.panel.Render(target)
	x += w

	// Top right
	s.panel.SetCurrentFrame(1)
	w, h = s.panel.GetCurrentFrameSize()
	s.panel.SetPosition(x, y+h)
	s.panel.Render(target)
	y += h

	// Bottom right
	s.panel.SetCurrentFrame(3)
	w, h = s.panel.GetCurrentFrameSize()
	s.panel.SetPosition(x, y+h)
	s.panel.Render(target)

	// Bottom left
	s.panel.SetCurrentFrame(2)
	w, h = s.panel.GetCurrentFrameSize()
	s.panel.SetPosition(x-w, y+h)
	s.panel.Render(target)

	var label d2ui.Label
	// all static labels are not stored since we use them only once to generate the image cache
	for _, textElement := range StaticTextLabels {
		label = s.createTextLabel(textElement)
		label.Render(target)
	}
	// hero name and class are part of the static image cache since they don't change after we enter the world
	label = s.createTextLabel(PanelText{X: 165, Y: 72, Text: s.heroName, Font: d2resource.Font16, AlignCenter: true})
	label.Render(target)
	label = s.createTextLabel(PanelText{X: 330, Y: 72, Text: s.heroClass.String(), Font: d2resource.Font16, AlignCenter: true})
	label.Render(target)
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

func (s *HeroStatsPanel) renderStatValueNum(label d2ui.Label, value int, target d2interface.Surface) {
	label.SetText(strconv.Itoa(value))
	label.Render(target)
}

func (s *HeroStatsPanel) createStatValueLabel(stat int, x int, y int) d2ui.Label {
	text := strconv.Itoa(stat)
	return s.createTextLabel(PanelText{X: x, Y: y, Text: text, Font: d2resource.Font16, AlignCenter: true})
}

func (s *HeroStatsPanel) createTextLabel(element PanelText) d2ui.Label {
	label := d2ui.CreateLabel(s.renderer, element.Font, d2resource.PaletteStatic)
	if element.AlignCenter {
		label.Alignment = d2ui.LabelAlignCenter
	}

	label.SetText(element.Text)
	label.SetPosition(element.X, element.Y)
	return label
}
