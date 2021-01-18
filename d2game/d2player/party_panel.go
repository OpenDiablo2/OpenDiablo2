package d2player

import (
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
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
	// nolint:deadcode,varcheck,unused // will be used
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
	levelLabelX               = 383
	baseBarY                  = 134
	baseRelationshipSwitcherY = 150
	baseSeeingSwitcherY       = 140
	baseListeningSwitcherY    = 140
	baseNameLabelY            = 145
	baseClassLabelY           = 158
	baseLevelLabelY           = 160
	nextBar                   = 52
)

// newPartyIndex creates new party index
func (s *PartyPanel) newPartyIndex(player *d2mapentity.Player, idx int, relations d2enum.PlayersRelationships) *partyIndex {
	result := &partyIndex{
		hero: player,
	}

	nameLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	nameLabel.SetText(player.Name())
	result.name = nameLabel

	classLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	classLabel.SetText(s.asset.TranslateString(player.Class.String()))
	result.class = classLabel

	levelLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	levelLabel.SetText(s.asset.TranslateString("level") + ": " + strconv.Itoa(player.Stats.Level))
	levelLabel.Alignment = d2ui.HorizontalAlignRight
	result.level = levelLabel

	relationships := s.createSwitcher(relationshipsFrame)
	result.relationshipSwitcher = relationships

	seeing := s.createSwitcher(seeingButtonFrame)
	result.seeingSwitcher = seeing

	listening := s.createSwitcher(listeningButtonFrame)
	result.listeningSwitcher = listening

	result.relationships = relations

	result.setColor(relations)
	result.setPositions(idx)

	return result
}

type partyIndex struct {
	hero                 *d2mapentity.Player
	name                 *d2ui.Label
	class                *d2ui.Label
	level                *d2ui.Label
	relationshipSwitcher *d2ui.SwitchableButton
	seeingSwitcher       *d2ui.SwitchableButton
	listeningSwitcher    *d2ui.SwitchableButton
	relationships        d2enum.PlayersRelationships
}

func (pi *partyIndex) setColor(relations d2enum.PlayersRelationships) {
	var color = d2util.Color(d2gui.ColorWhite)

	switch relations {
	case d2enum.PlayerRelationEnemy:
		color = d2util.Color(d2gui.ColorRed)

		pi.relationshipSwitcher.SetState(false)
	case d2enum.PlayerRelationFriend:
		color = d2util.Color(d2gui.ColorGreen)
	}

	pi.name.Color[0] = color
	pi.class.Color[0] = color
	pi.level.Color[0] = color
}

func (pi *partyIndex) setPositions(idx int) {
	pi.name.SetPosition(nameLabelX, baseNameLabelY+nextBar*idx)
	pi.class.SetPosition(classLabelX, baseClassLabelY+nextBar*idx)
	pi.level.SetPosition(levelLabelX, baseLevelLabelY+nextBar*idx)
	pi.relationshipSwitcher.SetPosition(relationshipSwitcherX, baseRelationshipSwitcherY+nextBar*idx)
	pi.seeingSwitcher.SetPosition(seeingSwitcherX, baseSeeingSwitcherY+idx*nextBar)
	pi.listeningSwitcher.SetPosition(listeningSwitcherX, baseListeningSwitcherY+idx*nextBar)
}

// DeletePlayer deletes player from PartyIndexes
func (s *PartyPanel) DeletePlayer(player *d2mapentity.Player) bool {
	for n, i := range s.partyIndexes {
		if i.hero == player {
			s.partyIndexes[n] = nil
			s.Sort()

			return true
		}
	}

	return false
}

// Sort sorts party indexes
func (s *PartyPanel) Sort() {
	var sorted [maxPlayersInGame]*partyIndex

	idx := 0

	for _, i := range s.partyIndexes {
		if i != nil {
			sorted[idx] = i
			idx++
		}
	}

	s.partyIndexes = sorted

	for n, i := range s.partyIndexes {
		if i != nil {
			i.setPositions(n)
		}
	}
}

// UpdatePlayer updates party-index
func (s *PartyPanel) UpdatePlayer(oldPlayer, newPlayer *d2mapentity.Player) bool {
	for n, i := range s.partyIndexes {
		if i.hero == oldPlayer {
			s.partyIndexes[n].hero = newPlayer

			return true
		}
	}

	return false
}

func (s *PartyPanel) setBarPosition() {
	for n, i := range s.partyIndexes {
		currentN := n

		if i == nil {
			s.barX, s.barY = barX, baseBarY+currentN*nextBar
			break
		}
	}
}

// NewPartyPanel creates a new party panel
func NewPartyPanel(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	heroName string,
	l d2util.LogLevel,

	// example Player structure (me)
	testPlayer *d2mapentity.Player,

	heroState *d2hero.HeroStatsState) *PartyPanel {
	originX := 0
	originY := 0

	var v [maxPlayersInGame]*partyIndex
	for i := 0; i < maxPlayersInGame; i++ {
		v[i] = nil
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

		testPlayer: testPlayer,
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

	testPlayer *d2mapentity.Player
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

	// example data
	p0 := s.testPlayer
	s.partyIndexes[0] = s.newPartyIndex(p0, 0, d2enum.PlayerRelationEnemy)
	p1 := s.testPlayer
	// nolint:gomnd // only test
	p1.Stats.Level = 99
	p1.Class = d2enum.HeroNecromancer
	s.partyIndexes[1] = s.newPartyIndex(p1, 1, d2enum.PlayerRelationFriend)

	if !s.DeletePlayer(p0) {
		s.Warning("Cannot remove player: Player Not Found")
	}

	for _, i := range s.partyIndexes {
		if i != nil {
			continue
		}

		s.panelGroup.AddWidget(i.name)
		s.panelGroup.AddWidget(i.class)
		s.panelGroup.AddWidget(i.relationshipSwitcher)
		s.panelGroup.AddWidget(i.seeingSwitcher)
		s.panelGroup.AddWidget(i.listeningSwitcher)
		s.panelGroup.AddWidget(i.level)
	}

	w, h = s.bar.GetCurrentFrameSize()
	v := s.uiManager.NewCustomWidget(s.renderBar, w, h)
	s.panelGroup.AddWidget(v)

	s.setBarPosition()

	s.panelGroup.SetVisible(false)
}

func (s *PartyPanel) createSwitcher(frame int) *d2ui.SwitchableButton {
	active := s.uiManager.NewCustomButton(d2resource.PartyBoxes, frame)
	inactive := s.uiManager.NewCustomButton(d2resource.PartyBoxes, frame+nextButtonFrame)
	switcher := s.uiManager.NewSwitchableButton(active, inactive, true)
	switcher.SetVisible(false)

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
