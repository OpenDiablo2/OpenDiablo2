package d2player

import (
	"fmt"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	white = 0xffffffff
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
	iconOffsetY                = 88
	questOffsetX, questOffsetY = 4, 4
	socket1X                   = 100
	socket2X                   = 200
	socket3X                   = 300
	socketUpY                  = 95
	socketDownY                = 190
)

const (
	questLogCloseButtonX, questLogCloseButtonY = 358, 455
	questLogDescrButtonX, questLogDescrButtonY = 308, 457
	questNameLabelX, questNameLabelY           = 240, 297
	questDescrLabelX, questDescrLabelY         = 90, 317
)

const (
	questTabY       = 66
	questTabYOffset = 31
	questTabBaseX   = 86
	questTabXOffset = 61
)

const questCompleteAnimationDuration = 3

func (s *QuestLog) getPositionForSocket(number int) (x, y int) {
	pos := []struct {
		x int
		y int
	}{
		{socket1X, socketUpY},
		{socket2X, socketUpY},
		{socket3X, socketUpY},
		{socket1X, socketDownY},
		{socket2X, socketDownY},
		{socket3X, socketDownY},
	}

	return pos[number].x, pos[number].y
}

// NewQuestLog creates a new quest log
func NewQuestLog(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	l d2util.LogLevel,
	audioProvider d2interface.AudioProvider,
	act int) *QuestLog {
	originX := 0
	originY := 0

	//nolint:gomnd // this is only test
	qs := map[int]int{
		0:  -2,
		1:  -2,
		2:  -2,
		3:  0,
		4:  1,
		5:  4,
		6:  3,
		7:  -1,
		8:  0,
		9:  0,
		10: 0,
		11: 0,
		12: 0,
		13: 0,
		14: 0,
		15: 0,
		16: 0,
		17: 0,
		18: 0,
		19: 0,
		20: 0,
		21: 0,
		22: 0,
		23: 0,
		24: 0,
		25: 0,
		26: 0,
	}

	var quests [d2enum.ActsNumber]*questEntire
	for i := 0; i < d2enum.ActsNumber; i++ {
		quests[i] = &questEntire{WidgetGroup: ui.NewWidgetGroup(d2ui.RenderPriorityQuestLog)}
	}

	var tabs [d2enum.ActsNumber]questLogTab
	for i := 0; i < d2enum.ActsNumber; i++ {
		tabs[i] = questLogTab{}
	}

	// nolint:gomnd // this is only test, it also should come from save file
	mpa := 2

	ql := &QuestLog{
		asset:         asset,
		uiManager:     ui,
		originX:       originX,
		originY:       originY,
		act:           act,
		tab:           tabs,
		quests:        quests,
		questStatus:   qs,
		maxPlayersAct: mpa,
		audioProvider: audioProvider,
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
	tab           [d2enum.ActsNumber]questLogTab
	audioProvider d2interface.AudioProvider
	completeSound d2interface.SoundEffect

	questName     *d2ui.Label
	questDescr    *d2ui.Label
	quests        [d2enum.ActsNumber]*questEntire
	questStatus   map[int]int
	maxPlayersAct int

	originX int
	originY int
	isOpen  bool

	*d2util.Logger
}

type questEntire struct {
	*d2ui.WidgetGroup
	icons   []*d2ui.Sprite
	buttons []*d2ui.Button
	sockets []*d2ui.Sprite
}

/* questIconTab returns path to quest animation using its
act and number. From d2resource:
        QuestLogAQuestAnimation = "/data/global/ui/MENU/a%dq%d.dc6"*/
func (s *QuestLog) questIconsTable(act, number int) string {
	return fmt.Sprintf(d2resource.QuestLogAQuestAnimation, act, number+1)
}

const (
	completedFrame  = 24
	inProgresFrame  = 25
	notStartedFrame = 26
)

const (
	socketNormalFrame      = 0
	socketHighlightedFrame = 1
)

const questDescriptionLenght = 30

type questLogTab struct {
	sprite          *d2ui.Sprite
	invisibleButton *d2ui.Button
}

// Load the data for the hero status panel
func (s *QuestLog) Load() {
	var err error

	s.panelGroup = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)

	// quest completion sound.
	s.completeSound, err = s.audioProvider.LoadSound(d2resource.QuestLogDoneSfx, false, false)
	if err != nil {
		s.Error(err.Error())
	}

	frame := d2ui.NewUIFrame(s.asset, s.uiManager, d2ui.FrameLeft)
	s.panelGroup.AddWidget(frame)

	s.panel, err = s.uiManager.NewSprite(d2resource.QuestLogBg, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	w, h := frame.GetSize()
	staticPanel := s.uiManager.NewCustomWidgetCached(s.renderStaticPanelFrames, w, h)
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
	s.questName.Color[0] = d2util.Color(white)
	s.questName.SetPosition(questNameLabelX, questNameLabelY)
	s.panelGroup.AddWidget(s.questName)

	s.questDescr = s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteStatic)
	s.questDescr.Alignment = d2ui.HorizontalAlignLeft
	s.questDescr.Color[0] = d2util.Color(white)
	s.questDescr.SetPosition(questDescrLabelX, questDescrLabelY)
	s.panelGroup.AddWidget(s.questDescr)

	s.loadTabs()

	// creates quest boards for each act
	for i := 0; i < d2enum.ActsNumber; i++ {
		item, icons, buttons, sockets := s.loadQuestBoard(i + 1)
		s.quests[i] = &questEntire{item, icons, buttons, sockets}
	}

	s.panelGroup.SetVisible(false)
}

// loadTabs loads quest log tabs
func (s *QuestLog) loadTabs() {
	var err error

	tabsResource := d2resource.WPTabs

	// create tabs only for 'discovered' acts
	for i := 0; i < s.maxPlayersAct; i++ {
		currentValue := i

		s.tab[i].sprite, err = s.uiManager.NewSprite(tabsResource, d2resource.PaletteSky)
		if err != nil {
			s.Error(err.Error())
		}

		// nolint:gomnd // it's constant.
		// each tab has two frames (active / inactive)
		frame := 2 * currentValue

		err := s.tab[i].sprite.SetCurrentFrame(frame)
		if err != nil {
			s.Errorf("Tabs sprite (%s) hasn't frame %d. %s", tabsResource, frame, err.Error())
		}

		s.tab[i].sprite.SetPosition(questTabBaseX+i*questTabXOffset, questTabY+questTabYOffset)

		s.tab[i].invisibleButton = s.uiManager.NewButton(d2ui.ButtonTypeTabBlank, "")
		s.tab[i].invisibleButton.SetPosition(questTabBaseX+i*questTabXOffset, questTabY)
		s.tab[i].invisibleButton.OnActivated(func() { s.setTab(currentValue) })

		s.panelGroup.AddWidget(s.tab[i].sprite)
		s.panelGroup.AddWidget(s.tab[i].invisibleButton)
	}

	// sets tab to current player's act.
	s.setTab(s.act - 1)
}

// loadQuestBoard creates quest fields (socket, button, icon) for specified act
func (s *QuestLog) loadQuestBoard(act int) (wg *d2ui.WidgetGroup, icons []*d2ui.Sprite, buttons []*d2ui.Button, sockets []*d2ui.Sprite) {
	wg = s.uiManager.NewWidgetGroup(d2ui.RenderPriorityQuestLog)

	// sets number of quests in act (for act 4 it's only 3, else 6)
	var questsInAct int
	if act == d2enum.Act4 {
		questsInAct = d2enum.HalfQuestsNumber
	} else {
		questsInAct = d2enum.NormalActQuestsNumber
	}

	for n := 0; n < questsInAct; n++ {
		cw := n
		x, y := s.getPositionForSocket(n)

		socket, err := s.uiManager.NewSprite(d2resource.QuestLogSocket, d2resource.PaletteSky)
		if err != nil {
			s.Error(err.Error())
		}

		socket.SetPosition(x, y+iconOffsetY+questOffsetY)
		sockets = append(sockets, socket)

		icon, err := s.makeQuestIconForAct(act, n, x, y)
		if err != nil {
			s.Error(err.Error())
		}

		icons = append(icons, icon)

		button := s.uiManager.NewButton(d2ui.ButtonTypeBlankQuestBtn, "")
		button.SetPosition(x+questOffsetX, y+questOffsetY)
		button.SetEnabled(s.questStatus[s.cordsToQuestID(act, cw)] != d2enum.QuestStatusNotStarted)
		buttons = append(buttons, button)
	}

	for i := 0; i < questsInAct; i++ {
		currentQuest := i

		// creates callback for quest button
		buttons[i].OnActivated(func() {
			var err error

			// set normal (not-highlighted) frame for each quest socket
			for j := 0; j < questsInAct; j++ {
				err = sockets[j].SetCurrentFrame(socketNormalFrame)
				if err != nil {
					s.Error(err.Error())
				}
			}

			// highlights appropriate socket
			err = sockets[currentQuest].SetCurrentFrame(socketHighlightedFrame)
			if err != nil {
				s.Error(err.Error())
			}

			// sets quest labels
			s.onQuestClicked(currentQuest + 1)
		})
	}

	// adds sockets to widget group
	for _, s := range sockets {
		wg.AddWidget(s)
	}

	// adds buttons to widget group
	for _, b := range buttons {
		wg.AddWidget(b)
	}

	// adds icons to widget group
	for _, i := range icons {
		wg.AddWidget(i)
	}

	wg.SetVisible(false)

	return wg, icons, buttons, sockets
}

func (s *QuestLog) makeQuestIconForAct(act, n, x, y int) (*d2ui.Sprite, error) {
	iconResource := s.questIconsTable(act, n)

	icon, err := s.uiManager.NewSprite(iconResource, d2resource.PaletteSky)
	if err != nil {
		s.Fatalf("during creating new quest icons for act %d (icon sprite %s doesn't exist). %s", act, iconResource, err.Error())
	}

	switch s.questStatus[s.cordsToQuestID(act, n)] {
	case d2enum.QuestStatusCompleted:
		err = icon.SetCurrentFrame(completedFrame)
	case d2enum.QuestStatusCompleting:
		// animation will be played after quest-log panel is opened (see s.playQuestAnimation)
		err = icon.SetCurrentFrame(0)
	case d2enum.QuestStatusNotStarted:
		err = icon.SetCurrentFrame(notStartedFrame)
	default:
		err = icon.SetCurrentFrame(inProgresFrame)
	}

	icon.SetPosition(x+questOffsetX, y+questOffsetY+iconOffsetY)

	return icon, err
}

// playQuestAnimations plays animations for quests (when status=questStatusCompleting)
func (s *QuestLog) playQuestAnimations() {
	for j, i := range s.quests[s.selectedTab].icons {
		questID := s.cordsToQuestID(s.selectedTab+1, j)
		if s.questStatus[questID] == d2enum.QuestStatusCompleting {
			s.completeSound.Play()

			// quest should be highlighted and it's label should be displayed
			s.quests[s.selectedTab].buttons[j].Activate()

			i.SetPlayLength(questCompleteAnimationDuration)
			i.PlayForward()
			i.SetPlayLoop(false)
		}
	}
}

// stopPlayedAnimation stops currently played animations and sets quests in
// completing state to completed (should be used, when quest log is closing)
func (s *QuestLog) stopPlayedAnimations() {
	// stops all played animations
	for j, i := range s.quests[s.selectedTab].icons {
		questID := s.cordsToQuestID(s.selectedTab+1, j)
		if s.questStatus[questID] == d2enum.QuestStatusCompleting {
			s.questStatus[questID] = d2enum.QuestStatusCompleted

			err := i.SetCurrentFrame(completedFrame)
			if err != nil {
				s.Error(err.Error())
			}
		}
	}
}

// setQuestLabel loads quest labels text (title and description)
func (s *QuestLog) setQuestLabel() {
	if s.selectedQuest == d2enum.QuestNone {
		s.questName.SetText("")
		s.questDescr.SetText("")

		return
	}

	s.questName.SetText(s.asset.TranslateString(fmt.Sprintf("qstsa%dq%d", s.selectedTab+1, s.selectedQuest)))

	status := s.questStatus[s.cordsToQuestID(s.selectedTab+1, s.selectedQuest)-1]
	switch status {
	case d2enum.QuestStatusCompleted, d2enum.QuestStatusCompleting:
		s.questDescr.SetText(
			strings.Join(
				d2util.SplitIntoLinesWithMaxWidth(
					s.asset.TranslateString("qstsprevious"),
					questDescriptionLenght),
				"\n"),
		)
	case d2enum.QuestStatusNotStarted:
		s.questDescr.SetText("")
	default:
		str := fmt.Sprintf("qstsa%dq%d%d", s.selectedTab+1, s.selectedQuest, status)
		descr := s.asset.TranslateString(str)

		// if description not found
		if str == descr {
			s.questDescr.SetText("")
		} else {
			s.questDescr.SetText(strings.Join(
				d2util.SplitIntoLinesWithMaxWidth(
					descr, questDescriptionLenght),
				"\n"),
			)
		}
	}
}

// switch all socket (in current tab) to normal state
func (s *QuestLog) clearHighlightment() {
	for _, i := range s.quests[s.selectedTab].sockets {
		err := i.SetCurrentFrame(socketNormalFrame)
		if err != nil {
			s.Error(err.Error())
		}
	}
}

func (s *QuestLog) setTab(tab int) {
	var mod int

	// before we leafe current tab, we need to switch highlighted
	// quest socket to normal frame
	s.clearHighlightment()

	s.selectedTab = tab
	s.selectedQuest = d2enum.QuestNone
	s.setQuestLabel()
	s.playQuestAnimations()

	// displays appropriate quests board
	for i := 0; i < s.maxPlayersAct; i++ {
		s.quests[i].SetVisible(tab == i)
	}

	// "highlights" appropriate tab
	for i := 0; i < s.maxPlayersAct; i++ {
		cv := i

		// converts bool to 1/0
		if cv == s.selectedTab {
			mod = 0
		} else {
			mod = 1
		}

		// sets tab sprite to highlighted/non-highlighted
		err := s.tab[cv].sprite.SetCurrentFrame(2*cv + mod)
		if err != nil {
			s.Error(err.Error())
		}
	}
}

func (s *QuestLog) onQuestClicked(number int) {
	s.selectedQuest = number
	s.setQuestLabel()
	s.Infof("Quest number %d in tab %d clicked", number, s.selectedTab)
}

//
func (s *QuestLog) onDescrClicked() {
	s.Info("Quest description button clicked")
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
	s.playQuestAnimations()
}

// Close closed the hero status panel
func (s *QuestLog) Close() {
	s.isOpen = false
	s.panelGroup.SetVisible(false)

	for i := 0; i < s.maxPlayersAct; i++ {
		s.quests[i].SetVisible(false)
	}

	s.stopPlayedAnimations()

	s.onCloseCb()
}

// SetOnCloseCb the callback run on closing the HeroStatsPanel
func (s *QuestLog) SetOnCloseCb(cb func()) {
	s.onCloseCb = cb
}

// Advance updates labels on the panel
func (s *QuestLog) Advance(elapsed float64) {
	if !s.IsOpen() {
		return
	}

	for j, i := range s.quests[s.selectedTab].icons {
		questID := s.cordsToQuestID(s.selectedTab+1, j)
		if s.questStatus[questID] == d2enum.QuestStatusCompleting {
			if err := i.Advance(elapsed); err != nil {
				s.Error(err.Error())
			}

			if i.GetCurrentFrame() == completedFrame {
				s.questStatus[questID] = d2enum.QuestStatusCompleted
			}
		}
	}
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

func (s *QuestLog) cordsToQuestID(act, number int) int {
	key := (act-1)*d2enum.NormalActQuestsNumber + number
	if act > d2enum.Act4 {
		key -= d2enum.HalfQuestsNumber
	}

	return key
}
