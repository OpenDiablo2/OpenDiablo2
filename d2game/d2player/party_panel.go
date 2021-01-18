package d2player

import (
	"log"
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
func (s *PartyPanel) newPartyIndex() *partyIndex {
	result := &partyIndex{}

	nameLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	result.name = nameLabel

	classLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	result.class = classLabel

	levelLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	levelLabel.Alignment = d2ui.HorizontalAlignRight
	result.level = levelLabel

	relationships := s.createSwitcher(relationshipsFrame)
	result.relationshipSwitcher = relationships

	seeing := s.createSwitcher(seeingButtonFrame)
	result.seeingSwitcher = seeing

	listening := s.createSwitcher(listeningButtonFrame)
	result.listeningSwitcher = listening

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

// AddPlayer adds a new player to the party panel
func (s *PartyPanel) AddPlayer(player *d2mapentity.Player, idx int, relations d2enum.PlayersRelationships) {
	s.partyIndexes[idx].hero = player

	s.partyIndexes[idx].name.SetText(player.Name())

	s.partyIndexes[idx].class.SetText(s.asset.TranslateString(player.Class.String()))

	s.partyIndexes[idx].level.SetText(s.asset.TranslateString("level") + ": " + strconv.Itoa(player.Stats.Level))

	s.partyIndexes[idx].relationships = relations

	s.partyIndexes[idx].setColor(relations)
	s.partyIndexes[idx].setPositions(idx)
}

// DeletePlayer deletes player from PartyIndexes
func (s *PartyPanel) DeletePlayer(player *d2mapentity.Player) bool {
	for n, i := range s.partyIndexes {
		if i.hero == player {
			s.Debugf("removing player at index %d", n)

			s.partyIndexes[n].hero = nil
			s.Sort()

			return true
		}
	}

	return false
}

// Sort sorts party indexes
func (s *PartyPanel) Sort() {
	var emptySlots []*partyIndex

	var emptySlotsNumbers []int

	var fullSlots []*partyIndex

	var fullSlotsNumbers []int

	// split s.partyIndexes to empty and non-empty
	for n, i := range s.partyIndexes {
		if i.hero == nil {
			emptySlots = append(emptySlots, i)
			emptySlotsNumbers = append(emptySlotsNumbers, n)
		} else {
			fullSlots = append(fullSlots, i)
			fullSlotsNumbers = append(fullSlotsNumbers, n)
		}
	}

	// adds non-empty indexes befor empty indexes
	for n, i := range fullSlots {
		s.partyIndexes[n] = i
	}

	// adds empty indexes
	for n, i := range emptySlots {
		s.partyIndexes[len(fullSlots)+n] = i
	}

	// sorts widget groups
	var sortedWG [maxPlayersInGame]*d2ui.WidgetGroup
	// first add non empty WG's
	for n, i := range fullSlotsNumbers {
		sortedWG[n] = s.indexes[i]
	}

	// after that, adds empty WG's
	for n, i := range emptySlotsNumbers {
		sortedWG[len(fullSlotsNumbers)+n] = s.indexes[i]
	}

	// overwrite existing order
	s.indexes = sortedWG

	// sets appropriate positions
	for n, i := range s.partyIndexes {
		if i.hero != nil {
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

		if i.hero == nil {
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
	log.Print("OpenDiablo2 - Party Panel - development")

	originX := 0
	originY := 0

	pp := &PartyPanel{
		asset:     asset,
		uiManager: ui,
		originX:   originX,
		originY:   originY,
		heroState: heroState,
		heroName:  heroName,
		labels:    &StatsPanelLabels{},
		barX:      barX,
		barY:      baseBarY,

		testPlayer: testPlayer,
	}

	var partyIndexes [maxPlayersInGame]*partyIndex

	var indexes [maxPlayersInGame]*d2ui.WidgetGroup

	for i := 0; i < maxPlayersInGame; i++ {
		partyIndexes[i] = pp.newPartyIndex()
		indexes[i] = pp.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)
	}

	pp.partyIndexes = partyIndexes

	pp.Logger = d2util.NewLogger()
	pp.Logger.SetLevel(l)
	pp.Logger.SetPrefix(logPrefix)

	return pp
}

// PartyPanel represents the party panel
type PartyPanel struct {
	asset      *d2asset.AssetManager
	uiManager  *d2ui.UIManager
	panel      *d2ui.Sprite
	bar        *d2ui.Sprite
	heroState  *d2hero.HeroStatsState
	heroName   string
	labels     *StatsPanelLabels
	onCloseCb  func()
	panelGroup *d2ui.WidgetGroup

	partyIndexes [maxPlayersInGame]*partyIndex
	indexes      [maxPlayersInGame]*d2ui.WidgetGroup

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
	for i := 0; i < maxPlayersInGame; i++ {
		s.indexes[i] = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)
	}

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
	s.AddPlayer(p0, 0, d2enum.PlayerRelationEnemy)
	p1 := s.testPlayer
	// nolint:gomnd // only test
	p1.Stats.Level = 99
	p1.Class = d2enum.HeroNecromancer
	s.AddPlayer(p1, 1, d2enum.PlayerRelationFriend)

	for n, i := range s.partyIndexes {
		s.indexes[n].AddWidget(i.name)
		s.indexes[n].AddWidget(i.class)
		s.indexes[n].AddWidget(i.relationshipSwitcher)
		s.indexes[n].AddWidget(i.seeingSwitcher)
		s.indexes[n].AddWidget(i.listeningSwitcher)
		s.indexes[n].AddWidget(i.level)
	}

	if !s.DeletePlayer(p0) {
		s.Warning("cannot remove player: DeletePlayer returned false")
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

	for n, i := range s.indexes {
		if s.partyIndexes[n].hero != nil {
			i.SetVisible(true)
		}
	}
}

// Close closed the hero status panel
func (s *PartyPanel) Close() {
	s.isOpen = false
	s.panelGroup.SetVisible(false)

	for _, i := range s.indexes {
		i.SetVisible(false)
	}
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
