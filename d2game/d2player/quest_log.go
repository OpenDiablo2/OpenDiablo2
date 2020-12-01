package d2player

import (
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const white = 0xffffffff

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
	iconOffsetY                = 88
	questOffsetX, questOffsetY = 4, 4
	q1SocketX, q1SocketY       = 100, 95
	q2SocketX, q2SocketY       = 200, 95
	q3SocketX, q3SocketY       = 300, 95
	q4SocketX, q4SocketY       = 100, 190
	q5SocketX, q5SocketY       = 200, 190
	q6SocketX, q6SocketY       = 300, 190
)

const (
	questLogCloseButtonX, questLogCloseButtonY = 358, 455
	questLogDescrButtonX, questLogDescrButtonY = 308, 457
	questLabelX, questLabelY                   = 240, 297
)

// toset
const (
	questTabY  = 66
	questTab1X = 85
	questTab2X = 143
	questTab3X = 201
	questTab4X = 259
	questTab5X = 317
)

const (
	questLogTab1 = iota
	questLogTab2
	questLogTab3
	questLogTab4
	questLogTab5
	questLogNumTabs
)

const (
	questNone = iota
	quest1
	quest2
	quest3
	quest4
	quest5
	quest6
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
		act: 1,
		tab: [questLogNumTabs]*questLogTab{
			{},
			{},
			{},
			{},
			{},
		},
		selectedQuest: 1,
	}

	ql.Logger = d2util.NewLogger()
	ql.Logger.SetLevel(l)
	ql.Logger.SetPrefix(logPrefix)

	return ql
}

// QuestLog represents the quest log
type QuestLog struct {
	asset         *d2asset.AssetManager
	uiManager     *d2ui.UIManager
	panel         *d2ui.Sprite
	onCloseCb     func()
	panelGroup    *d2ui.WidgetGroup
	selectedTab   int
	selectedQuest int
	act           int
	tab           [questLogNumTabs]*questLogTab

	q1 *d2ui.Button
	q2 *d2ui.Button
	q3 *d2ui.Button
	q4 *d2ui.Button
	q5 *d2ui.Button
	q6 *d2ui.Button
	// quests    []*questField
	questsa1  *d2ui.WidgetGroup
	questName *d2ui.Label
	/*questsa2    *d2ui.WidgetGroup
	questsa3    *d2ui.WidgetGroup
	questsa4    *d2ui.WidgetGroup
	questsa5    *d2ui.WidgetGroup*/

	originX int
	originY int
	isOpen  bool

	*d2util.Logger
}

type questField struct {
	descr  *d2ui.Label
	sprite *d2ui.Sprite
	status int // for now -1 = complete, 0 = not started > 0 in progress
}

type questLogTab struct {
	button          *d2ui.Button
	invisibleButton *d2ui.Button
}

func (q *questLogTab) newTab(ui *d2ui.UIManager, tabType d2ui.ButtonType, x int) {
	q.button = ui.NewButton(tabType, "")
	q.invisibleButton = ui.NewButton(d2ui.ButtonTypeTabBlank, "")
	q.button.SetPosition(x, questTabY)
	q.invisibleButton.SetPosition(x, questTabY)
}

// IsAct4 returns true, when game act is act 4 (in this act, there are only 3 quests)
func (s *QuestLog) IsAct4() bool {
	return s.act == 4
}

// Load the data for the hero status panel
func (s *QuestLog) Load() {
	var err error

	s.questsa1 = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)
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

	s.questName = s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteStatic)
	s.questName.Alignment = d2ui.HorizontalAlignCenter
	s.questName.Color[0] = rgbaColor(white)
	s.questName.SetPosition(questLabelX, questLabelY)
	s.panelGroup.AddWidget(s.questName)

	s.loadTabs()
	s.setQuestButtons()
	s.loadQuestIcons()

	s.questsa1.SetVisible(false)
	s.panelGroup.SetVisible(false)
}

func (s *QuestLog) questTable(act, number int) struct {
	name           string
	numberOfDescrs int
	status         int
	frame          int
	x              int
	y              int
} {
	var quests = []struct {
		name           string
		numberOfDescrs int
		status         int
		frame          int
		x, y           int
	}{
		{"qstsa1q1", 5, 0, 0, q1SocketX, q1SocketY},
		{"qstsa1q2", 0, 0, 1, q2SocketX, q2SocketY},
		{"qstsa1q3", 0, 0, 2, q3SocketX, q3SocketY},
		{"qstsa1q4", 0, 0, 3, q4SocketX, q4SocketY},
		{"qstsa1q5", 0, 0, 4, q5SocketX, q5SocketY},
		{"qstsa1q6", 0, 0, 5, q6SocketX, q6SocketY},
	}

	return quests[(act-1)*6+(number)]
}

func (s *QuestLog) loadQuestIcons() {
	var quest questField

	var err error

	for a := 0; a < 5; a++ {
		for n := 0; n < 6; n++ {
			q := s.questTable(1, n)

			quest.sprite, err = s.uiManager.NewSprite(d2resource.QuestLogDone, d2resource.PaletteSky)
			if err != nil {
				s.Error(err.Error())
			}

			err = quest.sprite.SetCurrentFrame(q.frame)
			if err != nil {
				s.Error(err.Error())
			}

			quest.sprite.SetPosition(q.x+questOffsetX, q.y+questOffsetY+iconOffsetY)

			s.panelGroup.AddWidget(quest.sprite)
		}
	}
}

func (s *QuestLog) loadQuestLabels() {
	if s.selectedQuest == 0 {
		s.questName.SetText("")
		return
	}

	s.questName.SetText(s.asset.TranslateString(fmt.Sprintf("qstsa%dq%d", s.selectedTab+1, s.selectedQuest)))
}

// copy from character select
func rgbaColor(rgba uint32) color.RGBA {
	result := color.RGBA{}
	a, b, g, r := 0, 1, 2, 3
	byteWidth := 8
	byteMask := 0xff

	for idx := 0; idx < 4; idx++ {
		shift := idx * byteWidth
		component := uint8(rgba>>shift) & uint8(byteMask)

		switch idx {
		case a:
			result.A = component
		case b:
			result.B = component
		case g:
			result.G = component
		case r:
			result.R = component
		}
	}

	return result
}

func (s *QuestLog) loadTabs() {
	s.tab[questLogTab1].newTab(s.uiManager, d2ui.ButtonTypeTab1, questTab1X)
	s.tab[questLogTab1].invisibleButton.OnActivated(func() { s.setTab(questLogTab1) })
	s.panelGroup.AddWidget(s.tab[questLogTab1].button)
	s.panelGroup.AddWidget(s.tab[questLogTab1].invisibleButton)

	s.tab[questLogTab2].newTab(s.uiManager, d2ui.ButtonTypeTab2, questTab2X)
	s.tab[questLogTab2].invisibleButton.OnActivated(func() { s.setTab(questLogTab2) })
	s.panelGroup.AddWidget(s.tab[questLogTab2].button)
	s.panelGroup.AddWidget(s.tab[questLogTab2].invisibleButton)

	s.tab[questLogTab3].newTab(s.uiManager, d2ui.ButtonTypeTab3, questTab3X)
	s.tab[questLogTab3].invisibleButton.OnActivated(func() { s.setTab(questLogTab3) })
	s.panelGroup.AddWidget(s.tab[questLogTab3].button)
	s.panelGroup.AddWidget(s.tab[questLogTab3].invisibleButton)

	s.tab[questLogTab4].newTab(s.uiManager, d2ui.ButtonTypeTab4, questTab4X)
	s.tab[questLogTab4].invisibleButton.OnActivated(func() { s.setTab(questLogTab4) })
	s.panelGroup.AddWidget(s.tab[questLogTab4].button)
	s.panelGroup.AddWidget(s.tab[questLogTab4].invisibleButton)

	s.tab[questLogTab5].newTab(s.uiManager, d2ui.ButtonTypeTab5, questTab5X)
	s.tab[questLogTab5].invisibleButton.OnActivated(func() { s.setTab(questLogTab5) })
	s.panelGroup.AddWidget(s.tab[questLogTab5].button)
	s.panelGroup.AddWidget(s.tab[questLogTab5].invisibleButton)

	s.setTab(s.act - 1)
}

func (s *QuestLog) setTab(tab int) {
	s.selectedTab = tab
	s.selectedQuest = questNone
	s.loadQuestLabels()

	for i := 0; i < questLogNumTabs; i++ {
		s.tab[i].button.SetEnabled(i == tab)
	}
}

func (s *QuestLog) setQuestButtons() {
	s.q1 = s.uiManager.NewButton(d2ui.ButtonTypeBlangQuestBtn, "")
	s.q1.SetPosition(q1SocketX+questOffsetX, q1SocketY+questOffsetY)
	s.q1.OnActivated(func() { s.onQuestClicked(quest1) })
	s.panelGroup.AddWidget(s.q1)

	s.q2 = s.uiManager.NewButton(d2ui.ButtonTypeBlangQuestBtn, "")
	s.q2.SetPosition(q2SocketX+questOffsetX, q2SocketY+questOffsetY)
	s.q2.OnActivated(func() { s.onQuestClicked(quest2) })
	s.panelGroup.AddWidget(s.q2)

	s.q3 = s.uiManager.NewButton(d2ui.ButtonTypeBlangQuestBtn, "")
	s.q3.SetPosition(q3SocketX+questOffsetX, q3SocketY+questOffsetY)
	s.q3.OnActivated(func() { s.onQuestClicked(quest3) })
	s.panelGroup.AddWidget(s.q3)

	s.q4 = s.uiManager.NewButton(d2ui.ButtonTypeBlangQuestBtn, "")
	s.q4.SetPosition(q4SocketX+questOffsetX, q4SocketY+questOffsetY)
	s.q4.OnActivated(func() { s.onQuestClicked(quest4) })
	s.panelGroup.AddWidget(s.q4)

	s.q5 = s.uiManager.NewButton(d2ui.ButtonTypeBlangQuestBtn, "")
	s.q5.SetPosition(q5SocketX+questOffsetX, q5SocketY+questOffsetY)
	s.q5.OnActivated(func() { s.onQuestClicked(quest5) })
	s.panelGroup.AddWidget(s.q5)

	s.q6 = s.uiManager.NewButton(d2ui.ButtonTypeBlangQuestBtn, "")
	s.q6.SetPosition(q6SocketX+questOffsetX, q6SocketY+questOffsetY)
	s.q6.OnActivated(func() { s.onQuestClicked(quest6) })
	s.panelGroup.AddWidget(s.q6)
}

func (s *QuestLog) onQuestClicked(number int) {
	s.selectedQuest = number
	s.loadQuestLabels()
	fmt.Printf("\nQuest %d clicked", number)
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
	s.questsa1.SetVisible(true)
}

// Close closed the hero status panel
func (s *QuestLog) Close() {
	s.isOpen = false
	s.panelGroup.SetVisible(false)
	s.questsa1.SetVisible(false)
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
