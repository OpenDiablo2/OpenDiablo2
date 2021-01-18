package d2player

import (
	"log"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	lightGreen = 0x18ff00ff
	red        = 0xff0000ff
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
	listeningButtonFrame = iota * 4
	relationshipsFrame
	seeingButtonFrame
	// nolint:deadcode,varcheck,unused // will be used
	lockButtonFrame

	nextButtonFrame = 2
)

const (
	buttonSize = 19
)

const (
	maxPlayersInGame          = 8
	barX                      = 90
	relationshipSwitcherX     = 95
	listeningSwitcherX        = 345
	seeingSwitcherX           = 365
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

	result.relationshipsActiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.relationshipsActiveTooltip.SetText(s.asset.TranslateString("strParty7") + "\n" + s.asset.TranslateString("strParty8"))
	relationships.SetActiveTooltip(result.relationshipsActiveTooltip)

	result.relationshipsInactiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky,
		d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.relationshipsInactiveTooltip.SetText(s.asset.TranslateString("strParty9") + "\n" + s.asset.TranslateString("strParty8"))
	relationships.SetInactiveTooltip(result.relationshipsInactiveTooltip)

	result.relationshipSwitcher = relationships

	seeing := s.createSwitcher(seeingButtonFrame)

	result.seeingActiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.seeingActiveTooltip.SetText(s.asset.TranslateString("strParty19"))
	seeing.SetActiveTooltip(result.seeingActiveTooltip)

	result.seeingInactiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.seeingInactiveTooltip.SetText(s.asset.TranslateString("strParty22"))
	seeing.SetInactiveTooltip(result.seeingInactiveTooltip)

	result.seeingSwitcher = seeing

	listening := s.createSwitcher(listeningButtonFrame)

	result.listeningActiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.listeningActiveTooltip.SetText(s.asset.TranslateString("strParty17") + "\n" + s.asset.TranslateString("strParty18"))
	listening.SetActiveTooltip(result.listeningActiveTooltip)

	result.listeningInactiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.listeningInactiveTooltip.SetText(s.asset.TranslateString("strParty11") + "\n" + s.asset.TranslateString("strParty16"))
	listening.SetInactiveTooltip(result.listeningInactiveTooltip)

	result.listeningSwitcher = listening

	return result
}

// partyIndex represents a party index
type partyIndex struct {
	hero                         *d2mapentity.Player
	name                         *d2ui.Label
	class                        *d2ui.Label
	level                        *d2ui.Label
	relationshipSwitcher         *d2ui.SwitchableButton
	relationshipsActiveTooltip   *d2ui.Tooltip
	relationshipsInactiveTooltip *d2ui.Tooltip
	seeingSwitcher               *d2ui.SwitchableButton
	seeingActiveTooltip          *d2ui.Tooltip
	seeingInactiveTooltip        *d2ui.Tooltip
	listeningSwitcher            *d2ui.SwitchableButton
	listeningActiveTooltip       *d2ui.Tooltip
	listeningInactiveTooltip     *d2ui.Tooltip
	relationships                d2enum.PlayersRelationships
}

// setColor sets appropriate labels' colors
func (pi *partyIndex) setColor(relations d2enum.PlayersRelationships) {
	var color = d2util.Color(white)

	switch relations {
	case d2enum.PlayerRelationEnemy:
		color = d2util.Color(red)

		pi.relationshipSwitcher.SetState(false)
	case d2enum.PlayerRelationFriend:
		color = d2util.Color(lightGreen)
	}

	pi.name.Color[0] = color
	pi.class.Color[0] = color
	pi.level.Color[0] = color
}

// setPositions sets party-index's position to given
func (pi *partyIndex) setPositions(idx int) {
	var h int

	pi.name.SetPosition(nameLabelX, baseNameLabelY+nextBar*idx)
	pi.class.SetPosition(classLabelX, baseClassLabelY+nextBar*idx)
	pi.level.SetPosition(levelLabelX, baseLevelLabelY+nextBar*idx)

	pi.relationshipSwitcher.SetPosition(relationshipSwitcherX, baseRelationshipSwitcherY+nextBar*idx)
	_, h = pi.relationshipsActiveTooltip.GetSize()
	pi.relationshipsActiveTooltip.SetPosition(relationshipSwitcherX+buttonSize, baseRelationshipSwitcherY+idx*nextBar-h)
	_, h = pi.relationshipsInactiveTooltip.GetSize()
	pi.relationshipsInactiveTooltip.SetPosition(relationshipSwitcherX+buttonSize, baseRelationshipSwitcherY+idx*nextBar-h)

	pi.seeingSwitcher.SetPosition(seeingSwitcherX, baseSeeingSwitcherY+idx*nextBar)
	_, h = pi.seeingActiveTooltip.GetSize()
	pi.seeingActiveTooltip.SetPosition(seeingSwitcherX+buttonSize, baseSeeingSwitcherY+idx*nextBar-h)
	_, h = pi.seeingInactiveTooltip.GetSize()
	pi.seeingInactiveTooltip.SetPosition(seeingSwitcherX+buttonSize, baseSeeingSwitcherY+idx*nextBar-h)

	pi.listeningSwitcher.SetPosition(listeningSwitcherX, baseListeningSwitcherY+idx*nextBar)
	_, h = pi.listeningActiveTooltip.GetSize()
	pi.listeningActiveTooltip.SetPosition(listeningSwitcherX+buttonSize, baseListeningSwitcherY+idx*nextBar-h)
	_, h = pi.listeningInactiveTooltip.GetSize()
	pi.listeningInactiveTooltip.SetPosition(listeningSwitcherX+buttonSize, baseListeningSwitcherY+idx*nextBar-h)
}

// AddPlayer adds a new player to the party panel
func (s *PartyPanel) AddPlayer(player *d2mapentity.Player, relations d2enum.PlayersRelationships) {
	idx := 0

	for n, i := range s.partyIndexes {
		if i.hero == nil {
			idx = n
			break
		}
	}

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

// IsInPanel returns true if player given already exists in panel
func (s *PartyPanel) IsInPanel(player *d2mapentity.Player) bool {
	for _, i := range s.partyIndexes {
		if i.hero == player {
			return true
		}
	}

	return false
}

// IsMe returns true if player given is "me"
func (s *PartyPanel) IsMe(player *d2mapentity.Player) bool {
	return player == s.me
}

// setBarPosition sets party-panel bar's position
func (s *PartyPanel) setBarPosition() {
	for n, i := range s.partyIndexes {
		currentN := n

		if i.hero == nil {
			s.barX, s.barY = barX, baseBarY+currentN*nextBar
			break
		}
	}
}

// UpdatePanel updates panel indexes with players list
func (s *PartyPanel) UpdatePanel() {
	for _, i := range s.players {
		if !s.IsInPanel(i) && !s.IsMe(i) {
			s.AddPlayer(i, d2enum.PlayerRelationNeutral)

			// we need to switch all hidden widgets to be visible
			if s.IsOpen() {
				s.Open()
			}
		}
	}
}

// NewPartyPanel creates a new party panel
func NewPartyPanel(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	heroName string,
	l d2util.LogLevel,
	me *d2mapentity.Player,
	heroState *d2hero.HeroStatsState,
	players map[string]*d2mapentity.Player) *PartyPanel {
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
		players:   players,
		me:        me,
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

	players map[string]*d2mapentity.Player
	me      *d2mapentity.Player

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

	var w, h int

	// create widgetGroups
	s.panelGroup = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)
	for i := 0; i < maxPlayersInGame; i++ {
		s.indexes[i] = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)
	}

	// create frame
	frame := s.uiManager.NewUIFrame(d2ui.FrameLeft)
	s.panelGroup.AddWidget(frame)

	s.panel, err = s.uiManager.NewSprite(d2resource.PartyPanel, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	// create panel
	w, h = frame.GetSize()
	staticPanel := s.uiManager.NewCustomWidgetCached(s.renderStaticPanelFrames, w, h)
	s.panelGroup.AddWidget(staticPanel)

	// create close button
	closeButton := s.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	closeButton.SetVisible(false)
	closeButton.SetPosition(partyPanelCloseButtonX, partyPanelCloseButtonY)
	closeButton.OnActivated(func() { s.Close() })
	s.panelGroup.AddWidget(closeButton)

	// our name label
	heroName := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	heroName.SetText(s.heroName)
	heroName.SetPosition(partyPanelHeroNameX, partyPanelHeroNameY)
	heroName.Alignment = d2ui.HorizontalAlignCenter
	s.panelGroup.AddWidget(heroName)

	// create WidgetGroups of party indexes
	for n, i := range s.partyIndexes {
		s.indexes[n].AddWidget(i.name)
		s.indexes[n].AddWidget(i.class)
		s.indexes[n].AddWidget(i.relationshipSwitcher)
		s.indexes[n].AddWidget(i.seeingSwitcher)
		s.indexes[n].AddWidget(i.listeningSwitcher)
		s.indexes[n].AddWidget(i.level)
	}

	// create bar
	s.bar, err = s.uiManager.NewSprite(d2resource.PartyBar, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	w, h = s.bar.GetCurrentFrameSize()
	v := s.uiManager.NewCustomWidget(s.renderBar, w, h)
	s.panelGroup.AddWidget(v)

	s.setBarPosition()

	s.panelGroup.SetVisible(false)
}

// createSwitcher creates party-panel switcher using frame given
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
	if !s.IsOpen() {
		return
	}

	s.UpdatePanel()
}

// UpdatePlayersList updates internal players list
func (s *PartyPanel) UpdatePlayersList(list map[string]*d2mapentity.Player) {
	s.players = list
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

// renderBar renders party panel's bar
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
