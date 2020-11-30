package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const ( // for the dc6 frames
	questLogTopLeft = iota
	questLogTopRight
	questLogBottomLeft
	questLogBottomRight
)

const (
	questLogOffsetX, questLogOffsetY = 80, 64
)

const (
	q1SocketX, q1SocketY = 100, 95
	q1X, q1Y             = 100, 95
	q2SocketX, q2SocketY = 200, 95
	q2X, q2Y             = 100, 95
	q3SocketX, q3SocketY = 300, 95
	q3X, q3Y             = 100, 95
	q4SocketX, q4SocketY = 100, 190
	q4X, q4Y             = 100, 95
	q5SocketX, q5SocketY = 200, 190
	q5X, q5Y             = 100, 95
	q6SocketX, q6SocketY = 300, 190
	q6X, q6Y             = 100, 95
)

const (
	questLogCloseButtonX, questLogCloseButtonY = 358, 455
	questLogDescrButtonX, questLogDescrButtonY = 308, 457
)

// toset
const (
	questNameX, questNameY               = 150, 220
	questDescriptionX, questDescriptionY = 50, 250
)

// toset
const (
	questTabY  = 66
	questTab1X = 85
	questTab2X = 143
	questTab3X = 201
	questTab4X = 259
	questTab5X = 317
	questTab6X = 375
)

//toset
const (
	questTabSelectedFrame1 = 1
	questTabSelectedFrame2 = 2
	questTabSelectedFrame3 = 3
	questTabSelectedFrame4 = 4
	questTabSelectedFrame5 = 5
	questTabSelectedFrame6 = 6
)

const (
	questLogTab1 = iota
	questLogTab2
	questLogTab3
	questLogTab4
	questLogTab5
	questLogTab6
	questLogNumTabs
)

// NewQuestLog creates a new quest log
func NewQuestLog(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	act int,
	l d2util.LogLevel) *QuestLog {
	originX := 0
	originY := 0

	ql := &QuestLog{
		asset:     asset,
		uiManager: ui,
		originX:   originX,
		originY:   originY,
		//act:       act,
		act: 5,
		tab: [questLogNumTabs]*questLogTab{
			{},
			{},
			{},
			{},
			{},
		},
	}

	ql.Logger = d2util.NewLogger()
	ql.Logger.SetLevel(l)
	ql.Logger.SetPrefix(logPrefix)

	return ql
}

// QuestLog represents the quest log
type QuestLog struct {
	asset      *d2asset.AssetManager
	uiManager  *d2ui.UIManager
	panel      *d2ui.Sprite
	labels     *StatsPanelLabels
	onCloseCb  func()
	panelGroup *d2ui.WidgetGroup
	act        int
	tab        [questLogNumTabs]*questLogTab

	originX int
	originY int
	isOpen  bool

	*d2util.Logger
}

type questField struct {
	name              *d2ui.Label
	status            int // for now -1 = complete, 0 = not started > 0 in progress
	notStarted        *d2ui.Sprite
	inProgress        *d2ui.Sprite
	completed         *d2ui.Sprite
	completeAnimation *d2ui.Sprite
	description       *d2ui.Label
}

type questLogTab struct {
	button *d2ui.Button
}

func (qt *questLogTab) createButton(uiManager *d2ui.UIManager, x, y int) {
}

// IsAct4 returns true, when game act is act 4 (in this act, there are only 3 quests)
func (s *QuestLog) IsAct4() bool {
	return s.act == 4
}

// Load the data for the hero status panel
func (s *QuestLog) Load() {
	var err error

	s.panelGroup = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)

	frame := d2ui.NewUIFrame(s.asset, s.uiManager, d2ui.FrameLeft)
	s.panelGroup.AddWidget(frame)

	s.panel, err = s.uiManager.NewSprite(d2resource.QuestLogBg, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	w, h := frame.GetSize()
	staticPanel := s.uiManager.NewCustomWidgetCached(s.renderStaticMenu, w, h)
	s.panelGroup.AddWidget(staticPanel)

	closeButton := s.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	closeButton.SetVisible(false)
	closeButton.SetPosition(questLogCloseButtonX, questLogCloseButtonY)
	closeButton.OnActivated(func() { s.Close() })
	s.panelGroup.AddWidget(closeButton)

	descrButton := s.uiManager.NewButton(d2ui.ButtonTypeQuestDescr, "")
	descrButton.SetVisible(false)
	descrButton.SetPosition(questLogDescrButtonX, questLogDescrButtonY)
	descrButton.OnActivated(s.onDescrClicked)
	s.panelGroup.AddWidget(descrButton)

	s.loadTabs()
	s.initStatValueLabels()
	s.panelGroup.SetVisible(false)
}

func (s *QuestLog) loadTabs() {
	s.tab[questLogTab1].button = s.uiManager.NewButton(d2ui.ButtonTypeTab1, "")
	s.tab[questLogTab1].button.SetPosition(questTab1X, questTabY)
	s.tab[questLogTab1].button.OnActivated(func() { s.setTab(questLogTab1) })
	//s.tab[questLogTab1].button.SetEnabled(false)
	s.panelGroup.AddWidget(s.tab[questLogTab1].button)

	s.tab[questLogTab2].button = s.uiManager.NewButton(d2ui.ButtonTypeTab2, "")
	s.tab[questLogTab2].button.SetPosition(questTab2X, questTabY)
	s.tab[questLogTab2].button.OnActivated(func() { s.setTab(questLogTab2) })
	//s.tab[questLogTab2].button.SetEnabled(false)
	s.panelGroup.AddWidget(s.tab[questLogTab2].button)

	s.tab[questLogTab3].button = s.uiManager.NewButton(d2ui.ButtonTypeTab3, "")
	s.tab[questLogTab3].button.SetPosition(questTab3X, questTabY)
	s.tab[questLogTab3].button.OnActivated(func() { s.setTab(questLogTab3) })
	//s.tab[questLogTab1].button.SetEnabled(false)
	s.panelGroup.AddWidget(s.tab[questLogTab3].button)

	s.tab[questLogTab4].button = s.uiManager.NewButton(d2ui.ButtonTypeTab4, "")
	s.tab[questLogTab4].button.SetPosition(questTab4X, questTabY)
	s.tab[questLogTab4].button.OnActivated(func() { s.setTab(questLogTab4) })
	//s.tab[questLogTab1].button.SetEnabled(false)
	s.panelGroup.AddWidget(s.tab[questLogTab4].button)

	s.tab[questLogTab5].button = s.uiManager.NewButton(d2ui.ButtonTypeTab5, "")
	s.tab[questLogTab5].button.SetPosition(questTab5X, questTabY)
	s.tab[questLogTab5].button.OnActivated(func() { s.setTab(questLogTab5) })
	//s.tab[questLogTab1].button.SetEnabled(false)
	s.panelGroup.AddWidget(s.tab[questLogTab5].button)

	s.setTab(1)
}

func (s *QuestLog) setTab(tab int) {
	for i := 0; i < questLogNumTabs-1; i++ {
		s.tab[i].button.SetEnabled(i == tab-1)
		//s.tab[i].button.SetPressed(!(i == tab-1))
	}
}

func (s *QuestLog) onDescrClicked() {
	//
}

// IsOpen returns true if the hero status panel is open
func (s *QuestLog) IsOpen() bool {
	return s.isOpen
}

// Toggle toggles the visibility of the hero status panel
func (s *QuestLog) Toggle() {
	if s.isOpen {
		s.Close()
	} else {
		s.Open()
	}
}

// Open opens the hero status panel
func (s *QuestLog) Open() {
	s.isOpen = true
	s.panelGroup.SetVisible(true)
}

// Close closed the hero status panel
func (s *QuestLog) Close() {
	s.isOpen = false
	s.panelGroup.SetVisible(false)
	s.onCloseCb()
}

// SetOnCloseCb the callback run on closing the HeroStatsPanel
func (s *QuestLog) SetOnCloseCb(cb func()) {
	s.onCloseCb = cb
}

/*
// Advance updates labels on the panel
func (s *HeroStatsPanel) Advance(elapsed float64) {
	if !s.isOpen {
		return
	}

	s.setStatValues()
}*/

func (s *QuestLog) renderStaticMenu(target d2interface.Surface) {
	s.renderStaticPanelFrames(target)
	s.renderSockets(target)
}
func (s *QuestLog) renderStaticPanelFrames(target d2interface.Surface) {
	frames := []int{
		questLogTopLeft,
		questLogTopRight,
		questLogBottomRight,
		questLogBottomLeft,
	}

	currentX := s.originX + questLogOffsetX
	currentY := s.originY + questLogOffsetY

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

func (s *QuestLog) renderSockets(target d2interface.Surface) {
	var socket *d2ui.Sprite

	var err error

	var sockets []struct{ x, y int }

	socketPath := d2resource.QuestLogSocket

	// all static labels are not stored since we use them only once to generate the image cache
	if s.IsAct4() {
		sockets = []struct {
			x, y int
		}{
			{q1SocketX, q1SocketY},
			{q2SocketX, q2SocketY},
			{q3SocketX, q3SocketY},
		}
	} else {
		sockets = []struct {
			x, y int
		}{
			{q1SocketX, q1SocketY},
			{q2SocketX, q2SocketY},
			{q3SocketX, q3SocketY},
			{q4SocketX, q4SocketY},
			{q5SocketX, q5SocketY},
			{q6SocketX, q6SocketY},
		}
	}

	for _, cfg := range sockets {
		socket, err = s.uiManager.NewSprite(socketPath, d2resource.PaletteSky)
		if err != nil {
			s.Error(err.Error())
		}

		socket.SetPosition(cfg.x, cfg.y)

		socket.RenderSegmented(target, 1, 1, 0)
	}
}

func (s *QuestLog) initStatValueLabels() {
	/*valueLabelConfigs := []struct {
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
	}*/
}

/*
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
*/
/*
func (s *QuestLog) createStatValueLabel(stat, x, y int) *d2ui.Label {
	text := strconv.Itoa(stat)
	return s.createTextLabel(PanelText{X: x, Y: y, Text: text, Font: d2resource.Font16, AlignCenter: true})
}*/

/*
func (s *QuestLog) createTextLabel(element PanelText) *d2ui.Label {
	label := s.uiManager.NewLabel(element.Font, d2resource.PaletteStatic)
	if element.AlignCenter {
		label.Alignment = d2ui.HorizontalAlignCenter
	}

	label.SetText(element.Text)
	label.SetPosition(element.X, element.Y)
	s.panelGroup.AddWidget(label)

	return label
}*/
