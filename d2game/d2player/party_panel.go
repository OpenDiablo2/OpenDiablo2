package d2player

import (
	"log"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
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
	lightRed   = 0xdb3f3dff
	orange     = 0xffa800ff
)

const ( // for the dc6 frames
	partyPanelTopLeft = iota
	partyPanelTopRight
	partyPanelBottomLeft
	partyPanelBottomRight
)

const ( // for bar's dc6 frames
	barLeft = iota
	barRight
)

const (
	partyPanelOffsetX, partyPanelOffsetY = 80, 64
)

const (
	partyPanelCloseButtonX, partyPanelCloseButtonY = 358, 453
	partyPanelHeroNameX, partyPanelHeroNameY       = 180, 80
)

const (
	buttonSize = 19
)

const (
	barX, baseBarY                                   = 90, 134
	relationshipSwitcherX, baseRelationshipSwitcherY = 95, 150
	listeningSwitcherX, baseListeningSwitcherY       = 342, 140
	seeingSwitcherX, baseSeeingSwitcherY             = 365, 140
	nameLabelX, baseNameLabelY                       = 115, 144
	nameTooltipX, baseNameTooltipY                   = 100, 120
	classLabelX, baseClassLabelY                     = 115, 158
	levelLabelX, baseLevelLabelY                     = 386, 160
	inviteAcceptButtonX, baseInviteAcceptButtonY     = 265, 147
	indexOffset                                      = 52
)

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

	var partyIndexes [d2enum.MaxPlayersInGame]*partyIndex

	var indexes [d2enum.MaxPlayersInGame]*d2ui.WidgetGroup

	for i := 0; i < d2enum.MaxPlayersInGame; i++ {
		partyIndexes[i] = pp.newPartyIndex()
		indexes[i] = pp.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)
	}

	pp.partyIndexes = partyIndexes
	pp.indexes = indexes

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

	partyIndexes [d2enum.MaxPlayersInGame]*partyIndex
	indexes      [d2enum.MaxPlayersInGame]*d2ui.WidgetGroup

	players map[string]*d2mapentity.Player
	me      *d2mapentity.Player

	originX int
	originY int
	isOpen  bool
	barX    int
	barY    int

	*d2util.Logger
}

// newPartyIndex creates new party index
func (s *PartyPanel) newPartyIndex() *partyIndex {
	result := &partyIndex{
		asset: s.asset,
		me:    s.me,
	}

	nameLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	result.nameTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.name = nameLabel

	classLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	result.class = classLabel

	result.nameRect = d2geom.Rectangle{}

	levelLabel := s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	levelLabel.Alignment = d2ui.HorizontalAlignRight
	result.level = levelLabel

	relationships := s.createSwitcher(d2enum.PartyButtonRelationshipsFrame)
	relationships.SetDisabledColor(lightRed)

	result.relationshipsActiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.relationshipsActiveTooltip.SetText(s.asset.TranslateString("strParty7") + "\n" + s.asset.TranslateString("strParty8"))
	relationships.SetActiveTooltip(result.relationshipsActiveTooltip)

	result.relationshipsInactiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky,
		d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.relationshipsInactiveTooltip.SetText(s.asset.TranslateString("strParty9") + "\n" + s.asset.TranslateString("strParty8"))
	relationships.SetInactiveTooltip(result.relationshipsInactiveTooltip)

	result.relationshipSwitcher = relationships

	seeing := s.createSwitcher(d2enum.PartyButtonSeeingFrame)

	result.seeingActiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.seeingActiveTooltip.SetText(s.asset.TranslateString("strParty19"))
	seeing.SetActiveTooltip(result.seeingActiveTooltip)

	result.seeingInactiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.seeingInactiveTooltip.SetText(s.asset.TranslateString("strParty22"))
	seeing.SetInactiveTooltip(result.seeingInactiveTooltip)

	result.seeingSwitcher = seeing

	listening := s.createSwitcher(d2enum.PartyButtonListeningFrame)

	result.listeningActiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.listeningActiveTooltip.SetText(s.asset.TranslateString("strParty17") + "\n" + s.asset.TranslateString("strParty18"))
	listening.SetActiveTooltip(result.listeningActiveTooltip)

	result.listeningInactiveTooltip = s.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	result.listeningInactiveTooltip.SetText(s.asset.TranslateString("strParty11") + "\n" + s.asset.TranslateString("strParty16"))
	listening.SetInactiveTooltip(result.listeningInactiveTooltip)

	result.listeningSwitcher = listening

	result.inviteAcceptButton = s.uiManager.NewButton(d2ui.ButtonTypePartyButton, s.asset.TranslateString("Invite"))
	result.inviteAcceptButton.SetVisible(false)

	return result
}

// partyIndex represents a party index
type partyIndex struct {
	asset *d2asset.AssetManager
	me    *d2mapentity.Player

	hero                         *d2mapentity.Player
	name                         *d2ui.Label
	nameTooltip                  *d2ui.Tooltip
	nameRect                     d2geom.Rectangle
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
	inviteAcceptButton           *d2ui.Button
	relationships                d2enum.PlayersRelationships
}

func (pi *partyIndex) setNameTooltipText() {
	switch pi.relationships {
	case d2enum.PlayerRelationNeutral, d2enum.PlayerRelationFriend:
		pi.nameTooltip.SetText(pi.asset.TranslateString("Party17"))
	case d2enum.PlayerRelationEnemy:
		pi.nameTooltip.SetText(pi.asset.TranslateString("Party12"))
	}
}

// setColor sets appropriate labels' colors
func (pi *partyIndex) setColor(relations d2enum.PlayersRelationships) {
	color := d2util.Color(white)

	switch relations {
	case d2enum.PlayerRelationEnemy:
		color = d2util.Color(red)

		pi.relationshipSwitcher.SetState(false)
	case d2enum.PlayerRelationFriend:
		color = d2util.Color(lightGreen)
	case d2enum.PlayerRelationNeutral:
		if pi.CanGoHostile() {
			color = d2util.Color(white)
		} else {
			color = d2util.Color(orange)
			pi.relationshipSwitcher.SetEnabled(false)
		}
	}

	pi.name.Color[0] = color
	pi.class.Color[0] = color
	pi.level.Color[0] = color
}

// setPositions sets party-index's position to given
func (pi *partyIndex) setPositions(idx int) {
	var w, h int

	pi.name.SetPosition(nameLabelX, baseNameLabelY+indexOffset*idx)
	pi.nameTooltip.SetPosition(nameTooltipX, baseNameTooltipY+indexOffset*idx)
	pi.class.SetPosition(classLabelX, baseClassLabelY+indexOffset*idx)
	pi.level.SetPosition(levelLabelX, baseLevelLabelY+indexOffset*idx)

	w, h1 := pi.class.GetSize()

	_, h = pi.name.GetSize()

	pi.nameRect = d2geom.Rectangle{
		Left:   nameLabelX,
		Top:    baseNameLabelY + idx*indexOffset,
		Width:  w,
		Height: h + h1,
	}

	pi.relationshipSwitcher.SetPosition(relationshipSwitcherX, baseRelationshipSwitcherY+indexOffset*idx)
	_, h = pi.relationshipsActiveTooltip.GetSize()
	pi.relationshipsActiveTooltip.SetPosition(relationshipSwitcherX+buttonSize, baseRelationshipSwitcherY+idx*indexOffset-h)
	_, h = pi.relationshipsInactiveTooltip.GetSize()
	pi.relationshipsInactiveTooltip.SetPosition(relationshipSwitcherX+buttonSize, baseRelationshipSwitcherY+idx*indexOffset-h)

	pi.seeingSwitcher.SetPosition(seeingSwitcherX, baseSeeingSwitcherY+idx*indexOffset)
	_, h = pi.seeingActiveTooltip.GetSize()
	pi.seeingActiveTooltip.SetPosition(seeingSwitcherX+buttonSize, baseSeeingSwitcherY+idx*indexOffset-h)
	_, h = pi.seeingInactiveTooltip.GetSize()
	pi.seeingInactiveTooltip.SetPosition(seeingSwitcherX+buttonSize, baseSeeingSwitcherY+idx*indexOffset-h)

	pi.listeningSwitcher.SetPosition(listeningSwitcherX, baseListeningSwitcherY+idx*indexOffset)
	_, h = pi.listeningActiveTooltip.GetSize()
	pi.listeningActiveTooltip.SetPosition(listeningSwitcherX+buttonSize, baseListeningSwitcherY+idx*indexOffset-h)
	_, h = pi.listeningInactiveTooltip.GetSize()
	pi.listeningInactiveTooltip.SetPosition(listeningSwitcherX+buttonSize, baseListeningSwitcherY+idx*indexOffset-h)

	pi.inviteAcceptButton.SetPosition(inviteAcceptButtonX, baseInviteAcceptButtonY+idx*indexOffset)
}

func (pi *partyIndex) CanGoHostile() bool {
	return pi.hero.Stats.Level >= d2enum.PlayersHostileLevel && pi.me.Stats.Level >= d2enum.PlayersHostileLevel
}

// Load the data for the hero status panel
func (s *PartyPanel) Load() {
	var err error

	var w, h int

	// create widgetGroups
	s.panelGroup = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)
	for i := 0; i < d2enum.MaxPlayersInGame; i++ {
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
		s.indexes[n].AddWidget(i.inviteAcceptButton)
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
	active := s.uiManager.NewDefaultButton(d2resource.PartyBoxes, frame)
	inactive := s.uiManager.NewDefaultButton(d2resource.PartyBoxes, frame+d2enum.PartyButtonNextButtonFrame)
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

// AddPlayer adds a new player to the party panel
func (s *PartyPanel) AddPlayer(player *d2mapentity.Player, relations d2enum.PlayersRelationships) {
	idx := 0

	// search for free index
	for n, i := range s.partyIndexes {
		if i.hero == nil {
			idx = n
			break
		}
	}

	s.partyIndexes[idx].hero = player

	s.partyIndexes[idx].name.SetText(player.Name())

	s.partyIndexes[idx].class.SetText(s.asset.TranslateString(player.Class.String()))

	s.partyIndexes[idx].level.SetText(s.asset.TranslateString("Level") + ":" + strconv.Itoa(player.Stats.Level))

	s.partyIndexes[idx].relationships = relations

	s.partyIndexes[idx].setColor(relations)

	s.partyIndexes[idx].setPositions(idx)

	s.partyIndexes[idx].setNameTooltipText()
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
	var sortedWG [d2enum.MaxPlayersInGame]*d2ui.WidgetGroup
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
			s.barX, s.barY = barX, baseBarY+currentN*indexOffset
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
			// s.Open contains appropriate code to do that.
			if s.IsOpen() {
				s.Open()
			}
		}
	}
}

// UpdatePlayersList updates internal players list
func (s *PartyPanel) UpdatePlayersList(list map[string]*d2mapentity.Player) {
	s.players = list
}

// Advance advances panel
func (s *PartyPanel) Advance(_ float64) {
	if !s.IsOpen() {
		return
	}

	s.UpdatePanel()
}

// OnMouseMove handles mouse movement events
func (s *PartyPanel) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	mx, my := event.X(), event.Y()

	for _, i := range s.partyIndexes {
		// Mouse over a game control element
		if i.nameRect.IsInRect(mx, my) {
			i.nameTooltip.SetVisible(true)
		} else {
			i.nameTooltip.SetVisible(false)
		}
	}

	return true
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
		barLeft,
		barRight,
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
