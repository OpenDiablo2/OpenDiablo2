package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const ( // for the dc6 frames
	partyPanelTopLeft = iota
	partyPanelTopRight
	partyPanelBottomLeft
	partyPanelBottomRight
)

const (
	partyPanelOffsetX, partyPanelOffsetY = 80, 64
)

const (
	partyPanelCloseButtonX, partyPanelCloseButtonY = 358, 453
	partyPanelHeroNameX, partyPanelHeroNameY       = 180, 80
)

const (
	seeingButtonFrame = iota * 4
	relationshipsFrame
	listeningButtonFrame
	lockButtonFrame

	nextButtonFrame = 2
)

const (
	maxPlayersInGame          = 8
	barX                      = 90
	relationshipSwitcherX     = 95
	seeingSwitcherX           = 345
	listeningSwitcherX        = 365
	nameLabelX                = 115
	classLabelX               = 115
	baseBarY                  = 134
	baseRelationshipSwitcherY = 150
	baseSeeingSwitcherY       = 145
	baseListeningSwitcherY    = 145
	baseNameLabelY            = 145
	baseClassLabelY           = 160
	nextBar                   = 52
)

type partyIndex struct {
	name                 *d2ui.Label
	class                *d2ui.Label
	level                *d2ui.Label
	relationshipSwitcher *d2ui.SwitchableButton
	seeingSwitcher       *d2ui.SwitchableButton
	listeningSwitcher    *d2ui.SwitchableButton
}

func (s *PartyPanel) NewPartyIndex(name string, class d2enum.Hero, level int, idx int) *partyIndex {
	result := &partyIndex{}

	nameLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	nameLabel.SetText(name)
	nameLabel.SetPosition(nameLabelX, baseNameLabelY+nextBar*idx)
	result.name = nameLabel

	classLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	classLabel.SetText(s.asset.TranslateString(class.String()))
	classLabel.SetPosition(classLabelX, baseClassLabelY+nextBar*idx)
	result.class = classLabel

	relationships := s.createSwitcher(relationshipsFrame, relationshipSwitcherX, baseRelationshipSwitcherY+nextBar*idx)
	result.relationshipSwitcher = relationships

	seeing := s.createSwitcher(seeingButtonFrame, seeingSwitcherX, baseSeeingSwitcherY+nextBar*idx)
	result.seeingSwitcher = seeing

	listening := s.createSwitcher(listeningButtonFrame, listeningSwitcherX, baseListeningSwitcherY+nextBar*idx)
	result.listeningSwitcher = listening

	return result
}

// NewPartyPanel creates a new party panel
func NewPartyPanel(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	heroName string,
	l d2util.LogLevel,
	heroState *d2hero.HeroStatsState) *PartyPanel {
	originX := 0
	originY := 0

	var v [maxPlayersInGame]*partyIndex
	for i := 0; i < maxPlayersInGame; i++ {
		v[i] = &partyIndex{}
	}

	hsp := &PartyPanel{
		asset:        asset,
		uiManager:    ui,
		originX:      originX,
		originY:      originY,
		heroState:    heroState,
		heroName:     heroName,
		labels:       &StatsPanelLabels{},
		partyIndexes: v,
		barX:         barX,
		barY:         baseBarY,
	}

	hsp.Logger = d2util.NewLogger()
	hsp.Logger.SetLevel(l)
	hsp.Logger.SetPrefix(logPrefix)

	return hsp
}

// PartyPanel represents the party panel
type PartyPanel struct {
	asset        *d2asset.AssetManager
	uiManager    *d2ui.UIManager
	panel        *d2ui.Sprite
	bar          *d2ui.Sprite
	heroState    *d2hero.HeroStatsState
	heroName     string
	labels       *StatsPanelLabels
	onCloseCb    func()
	panelGroup   *d2ui.WidgetGroup
	partyIndexes [maxPlayersInGame]*partyIndex

	originX int
	originY int
	isOpen  bool
	barX    int
	barY    int

	*d2util.Logger
}

// Load the data for the hero status panel
func (s *PartyPanel) Load() {
	var err error

	s.panelGroup = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)

	frame := s.uiManager.NewUIFrame(d2ui.FrameLeft)
	s.panelGroup.AddWidget(frame)

	s.panel, err = s.uiManager.NewSprite(d2resource.PartyPanel, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	w, h := frame.GetSize()
	staticPanel := s.uiManager.NewCustomWidgetCached(s.renderStaticPanelFrames, w, h)
	s.panelGroup.AddWidget(staticPanel)

	closeButton := s.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	closeButton.SetVisible(false)
	closeButton.SetPosition(partyPanelCloseButtonX, partyPanelCloseButtonY)
	closeButton.OnActivated(func() { s.Close() })
	s.panelGroup.AddWidget(closeButton)

	heroName := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	heroName.SetText(s.heroName)
	heroName.SetPosition(partyPanelHeroNameX, partyPanelHeroNameY)
	heroName.Alignment = d2ui.HorizontalAlignCenter
	s.panelGroup.AddWidget(heroName)

	s.bar, err = s.uiManager.NewSprite(d2resource.PartyBar, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	s.barX, s.barY = barX, baseBarY+1*nextBar
	w, h = s.bar.GetCurrentFrameSize()
	v := s.uiManager.NewCustomWidget(s.renderBar, w, h)
	s.panelGroup.AddWidget(v)

	// example data
	s.partyIndexes[0] = s.NewPartyIndex("PartyMember", d2enum.HeroPaladin, 5, 0)
	for _, i := range s.partyIndexes {
		// needed for "developing time" to avoit panic
		if i.name != nil {
			s.panelGroup.AddWidget(i.name)
		}

		if i.class != nil {
			s.panelGroup.AddWidget(i.class)
		}

		if i.relationshipSwitcher != nil {
			s.panelGroup.AddWidget(i.relationshipSwitcher)
		}

		if i.seeingSwitcher != nil {
			s.panelGroup.AddWidget(i.seeingSwitcher)
		}

		if i.listeningSwitcher != nil {
			s.panelGroup.AddWidget(i.listeningSwitcher)
		}
	}

	s.panelGroup.SetVisible(false)
}

func (s *PartyPanel) createSwitcher(frame, x, y int) *d2ui.SwitchableButton {
	active := s.uiManager.NewCustomButton(d2resource.PartyBoxes, frame)
	inactive := s.uiManager.NewCustomButton(d2resource.PartyBoxes, frame+nextButtonFrame)
	switcher := s.uiManager.NewSwitchableButton(active, inactive, true)
	switcher.SetPosition(x, y)

	return switcher
}

// IsOpen returns true if the hero status panel is open
func (s *PartyPanel) IsOpen() bool {
	return s.isOpen
}

// Toggle toggles the visibility of the hero status panel
func (s *PartyPanel) Toggle() {
	if s.isOpen {
		s.Close()
	} else {
		s.Open()
	}
}

// Open opens the hero status panel
func (s *PartyPanel) Open() {
	s.isOpen = true
	s.panelGroup.SetVisible(true)
}

// Close closed the hero status panel
func (s *PartyPanel) Close() {
	s.isOpen = false
	s.panelGroup.SetVisible(false)
}

// SetOnCloseCb the callback run on closing the PartyPanel
func (s *PartyPanel) SetOnCloseCb(cb func()) {
	s.onCloseCb = cb
}

// Advance advances panel
func (s *PartyPanel) Advance(_ float64) {
	// noop
}

// nolint:dupl // see quest_log.go.renderStaticPanelFrames comment
func (s *PartyPanel) renderStaticPanelFrames(target d2interface.Surface) {
	frames := []int{
		partyPanelTopLeft,
		partyPanelTopRight,
		partyPanelBottomRight,
		partyPanelBottomLeft,
	}

	currentX := s.originX + partyPanelOffsetX
	currentY := s.originY + partyPanelOffsetY

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

func (s *PartyPanel) renderBar(target d2interface.Surface) {
	frames := []int{
		partyPanelTopLeft,
		partyPanelTopRight,
	}

	currentX := s.originX + s.barX
	currentY := s.originY + s.barY

	for _, frameIndex := range frames {
		if err := s.bar.SetCurrentFrame(frameIndex); err != nil {
			s.Error(err.Error())
		}

		w, h := s.bar.GetCurrentFrameSize()

		switch frameIndex {
		case statsPanelTopLeft:
			s.bar.SetPosition(currentX, currentY)
			currentX += w
		case statsPanelTopRight:
			s.bar.SetPosition(currentX, currentY)
			currentY += h
		}

		s.bar.Render(target)
	}
}
