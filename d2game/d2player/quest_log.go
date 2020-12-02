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

const (
	act1 = iota + 1
	act2
	act3
	act4
	act5
)

const ( // for the dc6 frames
	questLogTopLeft = iota
	questLogTopRight
	questLogBottomLeft
	questLogBottomRight
)

const (
	normalActQuestsNumber = 6
	act4QuestsNumber      = 3
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
	questNone = 0
)


func (s *QuestLog) questTable(act, number int) struct {
	name           string
	numberOfDescrs int
	status         int
	frame          int
	x              int
	y              int
} {
	var quests = []struct {
		name           string // name of quest in string table
		numberOfDescrs int    // number of possible descriptions (not used yet)
		status         int    // status of quest (not used yet)
		frame          int    // frame of quest
		x, y           int    // position of quest
	}{
		{"qstsa1q1", 5, 0, 0, q1SocketX, q1SocketY},
		{"qstsa1q2", 0, 0, 1, q2SocketX, q2SocketY},
		{"qstsa1q3", 0, 0, 2, q3SocketX, q3SocketY},
		{"qstsa1q4", 0, 0, 3, q4SocketX, q4SocketY},
		{"qstsa1q5", 0, 0, 4, q5SocketX, q5SocketY},
		{"qstsa1q6", 0, 0, 5, q6SocketX, q6SocketY},
		{"qstsa2q1", 0, 0, 6, q1SocketX, q1SocketY},
		{"qstsa2q2", 0, 0, 7, q2SocketX, q2SocketY},
		{"qstsa2q3", 0, 0, 8, q3SocketX, q3SocketY},
		{"qstsa2q4", 0, 0, 9, q4SocketX, q4SocketY},
		{"qstsa2q5", 0, 0, 10, q5SocketX, q5SocketY},
		{"qstsa2q6", 0, 0, 11, q6SocketX, q6SocketY},
		{"qstsa3q1", 0, 0, 12, q1SocketX, q1SocketY},
		{"qstsa3q2", 0, 0, 13, q2SocketX, q2SocketY},
		{"qstsa3q3", 0, 0, 14, q3SocketX, q3SocketY},
		{"qstsa3q4", 0, 0, 15, q4SocketX, q4SocketY},
		{"qstsa3q5", 0, 0, 16, q5SocketX, q5SocketY},
		{"qstsa3q6", 0, 0, 17, q6SocketX, q6SocketY},
		{"qstsa4q1", 0, 0, 18, q1SocketX, q1SocketY},
		{"qstsa4q2", 0, 0, 19, q2SocketX, q2SocketY},
		{"qstsa4q3", 0, 0, 20, q3SocketX, q3SocketY},
		{"qstsa5q1", 0, 0, 21, q1SocketX, q1SocketY},
		{"qstsa5q2", 0, 0, 22, q2SocketX, q2SocketY},
		{"qstsa5q3", 0, 0, 23, q3SocketX, q3SocketY},
		{"qstsa5q4", 0, 0, 24, q4SocketX, q4SocketY},
		{"qstsa5q5", 0, 0, 25, q5SocketX, q5SocketY},
		{"qstsa5q6", 0, 0, 26, q6SocketX, q6SocketY},
	}

	key := (act-1)*normalActQuestsNumber + number
	if act > act4 {
		key -= act4QuestsNumber
	}

	return quests[key]
}

// NewQuestLog creates a new quest log
func NewQuestLog(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	l d2util.LogLevel,
	act int) *QuestLog {
	originX := 0
	originY := 0

	ql := &QuestLog{
		asset:     asset,
		uiManager: ui,
		originX:   originX,
		originY:   originY,
		act:       act,
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
	asset         *d2asset.AssetManager
	uiManager     *d2ui.UIManager
	panel         *d2ui.Sprite
	onCloseCb     func()
	panelGroup    *d2ui.WidgetGroup
	selectedTab   int
	selectedQuest int
	act           int
	tab           [questLogNumTabs]*questLogTab

	questName *d2ui.Label
	questsa1  *d2ui.WidgetGroup
	questsa2  *d2ui.WidgetGroup
	questsa3  *d2ui.WidgetGroup
	questsa4  *d2ui.WidgetGroup
	questsa5  *d2ui.WidgetGroup

	originX int
	originY int
	isOpen  bool

	*d2util.Logger
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

// Load the data for the hero status panel
func (s *QuestLog) Load() {
	var err error

	s.questsa1 = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)
	s.questsa2 = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)
	s.questsa3 = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)
	s.questsa4 = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)
	s.questsa5 = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)
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
	s.loadQuestIcons()

	s.panelGroup.SetVisible(false)
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

func (s *QuestLog) loadQuestIcons() {
	s.questsa1 = s.loadQuestIconsForAct(act1)
	s.questsa2 = s.loadQuestIconsForAct(act2)
	s.questsa3 = s.loadQuestIconsForAct(act3)
	s.questsa4 = s.loadQuestIconsForAct(act4)
	s.questsa5 = s.loadQuestIconsForAct(act5)
}

func (s *QuestLog) loadQuestIconsForAct(act int) *d2ui.WidgetGroup {
	wg := s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)

	var questsInAct int
	if act == act4 {
		questsInAct = act4QuestsNumber
	} else {
		questsInAct = normalActQuestsNumber
	}

	for n := 0; n < questsInAct; n++ {
		q := s.questTable(act, n)

		button := s.uiManager.NewButton(d2ui.ButtonTypeBlankQuestBtn, "")
		button.SetPosition(q.x+questOffsetX, q.y+questOffsetY)
		button.OnActivated(s.makeQuestCallback(n))

		socket, err := s.uiManager.NewSprite(d2resource.QuestLogSocket, d2resource.PaletteSky)
		if err != nil {
			s.Error(err.Error())
		}

		socket.SetPosition(q.x+questOffsetX, q.y+iconOffsetY+2*questOffsetY)

		icon, err := s.uiManager.NewSprite(d2resource.QuestLogDone, d2resource.PaletteSky)
		if err != nil {
			s.Error(err.Error())
		}

		err = icon.SetCurrentFrame(q.frame)
		if err != nil {
			s.Error(err.Error())
		}

		icon.SetPosition(q.x+questOffsetX, q.y+questOffsetY+iconOffsetY)

		wg.AddWidget(icon)
		wg.AddWidget(socket)
		wg.AddWidget(button)
	}
	wg.SetVisible(false)

	return wg
}


func (s *QuestLog) makeQuestCallback(n int) func() {
	return func() {
		s.onQuestClicked(n + 1)
	}
}

func (s *QuestLog) setQuestLabels() {
	if s.selectedQuest == 0 {
		s.questName.SetText("")
		return
	}

	s.questName.SetText(s.asset.TranslateString(fmt.Sprintf("qstsa%dq%d", s.selectedTab+1, s.selectedQuest)))
}

func (s *QuestLog) setTab(tab int) {
	s.selectedTab = tab
	s.selectedQuest = questNone
	s.setQuestLabels()
  
	s.questsa1.SetVisible(tab == questLogTab1)
	s.questsa2.SetVisible(tab == questLogTab2)
	s.questsa3.SetVisible(tab == questLogTab3)
	s.questsa4.SetVisible(tab == questLogTab4)
	s.questsa5.SetVisible(tab == questLogTab5)

	for i := 0; i < questLogNumTabs; i++ {
		s.tab[i].button.SetEnabled(i == tab)
	}
}

func (s *QuestLog) onQuestClicked(number int) {
	s.selectedQuest = number
	s.setQuestLabels()
	s.Infof("Quest number %d in tab %d clicked", number, s.selectedTab)
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
	s.setTab(s.selectedTab)
}

// Close closed the hero status panel
func (s *QuestLog) Close() {
	s.isOpen = false
	s.panelGroup.SetVisible(false)
	s.questsa1.SetVisible(false)
	s.questsa2.SetVisible(false)
	s.questsa3.SetVisible(false)
	s.questsa4.SetVisible(false)
	s.questsa5.SetVisible(false)
	s.onCloseCb()
}

// SetOnCloseCb the callback run on closing the HeroStatsPanel
func (s *QuestLog) SetOnCloseCb(cb func()) {
	s.onCloseCb = cb
}

// Advance updates labels on the panel
func (s *QuestLog) Advance(elapsed float64) {
	//
}

func (s *QuestLog) renderStaticMenu(target d2interface.Surface) {
	s.renderStaticPanelFrames(target)
}

// nolint:dupl // I think it is OK, to duplicate this function
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
		case questLogTopLeft:
			s.panel.SetPosition(currentX, currentY+h)
			currentX += w
		case questLogTopRight:
			s.panel.SetPosition(currentX, currentY+h)
			currentY += h
		case questLogBottomRight:
			s.panel.SetPosition(currentX, currentY+h)
		case questLogBottomLeft:
			s.panel.SetPosition(currentX-w, currentY+h)
		}

		s.panel.Render(target)
	}
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
