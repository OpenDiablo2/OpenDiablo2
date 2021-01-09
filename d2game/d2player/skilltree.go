package d2player

import (
	"errors"
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	tabButtonX  = 628
	tabButton0Y = 385
	tabButton1Y = 277
	tabButton2Y = 170

	availSPLabelX = 677
	availSPLabelY = 72

	skillCloseButtonXLeft   = 416
	skillCloseButtonXMiddle = 501
	skillCloseButtonXRight  = 572
	skillCloseButtonY       = 449
)

const (
	firstTab = iota
	secondTab
	thirdTab
	numTabs
)

const (
	tabIndexOffset = 4

	frameOffsetTop    = 4
	frameOffsetBottom = 6
)

const (
	frameCommonTabTopLeft = iota
	frameCommonTabTopRight
	frameCommonTabBottomLeft
	frameCommonTabBottomRight
)

const (
	frameSelectedTab1Full   = 7
	frameSelectedTab2Top    = 9 // tab2 top and bottom portions are in 2 frames :(
	frameSelectedTab2Bottom = 11
	frameSelectedTab3Full   = 13
)

const (
	skillTreePanelX = 401
	skillTreePanelY = 64
)

const (
	skillIconGreySat    = 0.2
	skillIconGreyBright = 0.44
)

type skillTreeTab struct {
	buttonText      string
	button          *d2ui.Button
	closeButtonPosX int
}

func (st *skillTreeTab) createButton(uiManager *d2ui.UIManager, x, y int) {
	st.button = uiManager.NewButton(d2ui.ButtonTypeSkillTreeTab, st.buttonText)
	st.button.SetPosition(x, y)
}

type skillTreeHeroTypeResources struct {
	skillSprite    *d2ui.Sprite
	skillIconPath  string
	skillPanel     *d2ui.Sprite
	skillPanelPath string
}

func newSkillTree(
	skills map[int]*d2hero.HeroSkill,
	heroClass d2enum.Hero,
	asset *d2asset.AssetManager,
	l d2util.LogLevel,
	ui *d2ui.UIManager,
) *skillTree {
	st := &skillTree{
		skills:    skills,
		heroClass: heroClass,
		asset:     asset,
		uiManager: ui,
		originX:   skillTreePanelX,
		originY:   skillTreePanelY,
		tab: [numTabs]*skillTreeTab{
			{},
			{},
			{},
		},
		l: l,
	}

	st.Logger = d2util.NewLogger()
	st.Logger.SetLevel(l)
	st.Logger.SetPrefix(logPrefix)

	return st
}

type skillTree struct {
	resources    *skillTreeHeroTypeResources
	asset        *d2asset.AssetManager
	uiManager    *d2ui.UIManager
	skills       map[int]*d2hero.HeroSkill
	skillIcons   []*skillIcon
	heroClass    d2enum.Hero
	frame        *d2ui.UIFrame
	availSPLabel *d2ui.Label
	closeButton  *d2ui.Button
	tab          [numTabs]*skillTreeTab
	isOpen       bool
	originX      int
	originY      int
	selectedTab  int
	onCloseCb    func()
	panelGroup   *d2ui.WidgetGroup
	iconGroup    *d2ui.WidgetGroup
	panel        *d2ui.CustomWidget

	*d2util.Logger
	l d2util.LogLevel
}

func (s *skillTree) load() {
	s.panelGroup = s.uiManager.NewWidgetGroup(d2ui.RenderPrioritySkilltree)
	s.iconGroup = s.uiManager.NewWidgetGroup(d2ui.RenderPrioritySkilltreeIcon)

	s.panel = s.uiManager.NewCustomWidget(s.Render, 400, 600)
	s.panelGroup.AddWidget(s.panel)

	s.frame = d2ui.NewUIFrame(s.asset, s.uiManager, d2ui.FrameRight)
	s.panelGroup.AddWidget(s.frame)

	s.closeButton = s.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	s.closeButton.SetVisible(false)
	s.closeButton.OnActivated(func() { s.Close() })
	s.panelGroup.AddWidget(s.closeButton)

	if err := s.setHeroTypeResourcePath(); err != nil {
		s.Error(err.Error())
	}

	s.loadForHeroType()

	for _, skill := range s.skills {
		si := newSkillIcon(s.uiManager, s.resources.skillSprite, s.l, skill)
		s.skillIcons = append(s.skillIcons, si)
		s.iconGroup.AddWidget(si)
	}

	s.panelGroup.SetVisible(false)
	s.setTab(0)
	s.iconGroup.SetVisible(false)
}

func (s *skillTree) loadForHeroType() {
	sp, err := s.uiManager.NewSprite(s.resources.skillPanelPath, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	s.resources.skillPanel = sp

	si, err := s.uiManager.NewSprite(s.resources.skillIconPath, d2resource.PaletteSky)
	if err != nil {
		s.Error(err.Error())
	}

	s.resources.skillSprite = si

	s.tab[firstTab].createButton(s.uiManager, tabButtonX, tabButton0Y)
	s.tab[firstTab].button.OnActivated(func() { s.setTab(firstTab) })
	s.panelGroup.AddWidget(s.tab[firstTab].button)

	s.tab[secondTab].createButton(s.uiManager, tabButtonX, tabButton1Y)
	s.tab[secondTab].button.OnActivated(func() { s.setTab(secondTab) })
	s.panelGroup.AddWidget(s.tab[secondTab].button)

	s.tab[thirdTab].createButton(s.uiManager, tabButtonX, tabButton2Y)
	s.tab[thirdTab].button.OnActivated(func() { s.setTab(thirdTab) })
	s.panelGroup.AddWidget(s.tab[thirdTab].button)

	s.availSPLabel = s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	s.availSPLabel.SetPosition(availSPLabelX, availSPLabelY)
	s.availSPLabel.Alignment = d2ui.HorizontalAlignCenter
	s.availSPLabel.SetText(s.makeTabString("StrSklTree1", "StrSklTree2", "StrSklTree3"))
	s.panelGroup.AddWidget(s.availSPLabel)
}

type heroTabData struct {
	resources        *skillTreeHeroTypeResources
	str1, str2, str3 string
	closeButtonPos   [numTabs]int
}

func (s *skillTree) makeTabString(keys ...interface{}) string {
	translations := make([]interface{}, len(keys))

	token := "%s"
	format := token

	for idx, key := range keys {
		if idx > 0 {
			format += "\n" + token
		}

		translations[idx] = s.asset.TranslateString(key.(string))
	}

	return fmt.Sprintf(format, translations...)
}

func makeCloseButtonPos(close1, close2, close3 int) [numTabs]int {
	return [numTabs]int{close1, close2, close3}
}

func (s *skillTree) getTab(class d2enum.Hero) *heroTabData {
	tabMap := map[d2enum.Hero]*heroTabData{
		d2enum.HeroBarbarian: {
			&skillTreeHeroTypeResources{
				skillPanelPath: d2resource.SkillsPanelBarbarian,
				skillIconPath:  d2resource.BarbarianSkills,
			},
			s.makeTabString("StrSklTree21", "StrSklTree4"),
			s.makeTabString("StrSklTree21", "StrSklTree22"),
			s.makeTabString("StrSklTree20"),
			makeCloseButtonPos(
				skillCloseButtonXRight,
				skillCloseButtonXLeft,
				skillCloseButtonXRight),
		},
		d2enum.HeroNecromancer: {
			&skillTreeHeroTypeResources{
				skillPanelPath: d2resource.SkillsPanelNecromancer,
				skillIconPath:  d2resource.NecromancerSkills,
			},
			s.makeTabString("StrSklTree19"),
			s.makeTabString("StrSklTree17", "StrSklTree18", "StrSklTree5"),
			s.makeTabString("StrSklTree16", "StrSklTree5"),
			makeCloseButtonPos(
				skillCloseButtonXLeft,
				skillCloseButtonXRight,
				skillCloseButtonXLeft),
		},
		d2enum.HeroPaladin: {
			&skillTreeHeroTypeResources{
				skillPanelPath: d2resource.SkillsPanelPaladin,
				skillIconPath:  d2resource.PaladinSkills,
			},
			s.makeTabString("StrSklTree15", "StrSklTree4"),
			s.makeTabString("StrSklTree14", "StrSklTree13"),
			s.makeTabString("StrSklTree12", "StrSklTree13"),
			makeCloseButtonPos(
				skillCloseButtonXLeft,
				skillCloseButtonXMiddle,
				skillCloseButtonXLeft),
		},
		d2enum.HeroAssassin: {
			&skillTreeHeroTypeResources{
				skillPanelPath: d2resource.SkillsPanelAssassin,
				skillIconPath:  d2resource.AssassinSkills,
			},
			s.makeTabString("StrSklTree30"),
			s.makeTabString("StrSklTree31", "StrSklTree32"),
			s.makeTabString("StrSklTree33", "StrSklTree34"),
			makeCloseButtonPos(
				skillCloseButtonXMiddle,
				skillCloseButtonXRight,
				skillCloseButtonXLeft),
		},
		d2enum.HeroSorceress: {
			&skillTreeHeroTypeResources{
				skillPanelPath: d2resource.SkillsPanelSorcerer,
				skillIconPath:  d2resource.SorcererSkills,
			},
			s.makeTabString("StrSklTree25", "StrSklTree5"),
			s.makeTabString("StrSklTree24", "StrSklTree5"),
			s.makeTabString("StrSklTree23", "StrSklTree5"),
			makeCloseButtonPos(
				skillCloseButtonXLeft,
				skillCloseButtonXLeft,
				skillCloseButtonXRight),
		},
		d2enum.HeroAmazon: {
			&skillTreeHeroTypeResources{
				skillPanelPath: d2resource.SkillsPanelAmazon,
				skillIconPath:  d2resource.AmazonSkills,
			},
			s.makeTabString("StrSklTree10", "StrSklTree11", "StrSklTree4"),
			s.makeTabString("StrSklTree8", "StrSklTree9", "StrSklTree4"),
			s.makeTabString("StrSklTree6", "StrSklTree7", "StrSklTree4"),
			makeCloseButtonPos(
				skillCloseButtonXRight,
				skillCloseButtonXMiddle,
				skillCloseButtonXLeft),
		},
		d2enum.HeroDruid: {
			&skillTreeHeroTypeResources{
				skillPanelPath: d2resource.SkillsPanelDruid,
				skillIconPath:  d2resource.DruidSkills,
			},
			s.makeTabString("StrSklTree26"),
			s.makeTabString("StrSklTree27", "StrSklTree28"),
			s.makeTabString("StrSklTree29"),
			makeCloseButtonPos(
				skillCloseButtonXRight,
				skillCloseButtonXRight,
				skillCloseButtonXRight),
		},
	}

	return tabMap[class]
}

func (s *skillTree) setHeroTypeResourcePath() error {
	entry := s.getTab(s.heroClass)

	if entry == nil {
		return errors.New("unknown hero type")
	}

	s.resources = entry.resources
	s.tab[firstTab].buttonText = entry.str1
	s.tab[secondTab].buttonText = entry.str2
	s.tab[thirdTab].buttonText = entry.str3

	for i := 0; i < numTabs; i++ {
		s.tab[i].closeButtonPosX = entry.closeButtonPos[i]
	}

	return nil
}

// Toggle the skill tree visibility
func (s *skillTree) Toggle() {
	s.Info("SkillTree toggled")

	if s.isOpen {
		s.Close()
	} else {
		s.Open()
	}
}

// Close the skill tree
func (s *skillTree) Close() {
	s.isOpen = false

	s.panelGroup.SetVisible(false)
	s.iconGroup.SetVisible(false)

	s.onCloseCb()
}

// Open the skill tree
func (s *skillTree) Open() {
	s.isOpen = true

	s.panelGroup.SetVisible(true)
	s.iconGroup.SetVisible(true)

	// we only want to enable the icons of our current tab again
	s.setTab(s.selectedTab)
}

func (s *skillTree) IsOpen() bool {
	return s.isOpen
}

// Set the callback run on closing the skilltree
func (s *skillTree) SetOnCloseCb(cb func()) {
	s.onCloseCb = cb
}

func (s *skillTree) setTab(tab int) {
	s.selectedTab = tab
	s.closeButton.SetPosition(s.tab[tab].closeButtonPosX, skillCloseButtonY)

	for _, si := range s.skillIcons {
		si.SetVisible(si.skill.SkillPage == tab+1)
	}
}

func (s *skillTree) renderPanelSegment(
	target d2interface.Surface,
	frame int) {
	if err := s.resources.skillPanel.SetCurrentFrame(frame); err != nil {
		s.Error(err.Error())
		return
	}

	s.resources.skillPanel.Render(target)
}

func (s *skillTree) renderTabCommon(target d2interface.Surface) {
	skillPanel := s.resources.skillPanel
	x, y := s.originX, s.originY

	// top
	w, h, err := skillPanel.GetFrameSize(frameCommonTabTopLeft)
	if err != nil {
		s.Error(err.Error())
		return
	}

	y += h

	skillPanel.SetPosition(x, y)
	s.renderPanelSegment(target, frameCommonTabTopLeft)

	skillPanel.SetPosition(x+w, y)
	s.renderPanelSegment(target, frameCommonTabTopRight)

	// bottom
	_, h, err = skillPanel.GetFrameSize(frameCommonTabBottomLeft)
	if err != nil {
		s.Error(err.Error())
		return
	}

	y += h

	skillPanel.SetPosition(x, y)
	s.renderPanelSegment(target, frameCommonTabBottomLeft)

	skillPanel.SetPosition(x+w, y)
	s.renderPanelSegment(target, frameCommonTabBottomRight)
}

func (s *skillTree) renderTab(target d2interface.Surface, tab int) {
	topFrame := frameOffsetTop + (tabIndexOffset * tab)
	bottomFrame := frameOffsetBottom + (tabIndexOffset * tab)

	skillPanel := s.resources.skillPanel
	x, y := s.originX, s.originY

	// top
	_, h0, err := skillPanel.GetFrameSize(topFrame)
	if err != nil {
		s.Error(err.Error())
		return
	}

	y += h0

	skillPanel.SetPosition(x, y)
	s.renderPanelSegment(target, topFrame)

	// bottom
	w, h1, err := skillPanel.GetFrameSize(bottomFrame)
	if err != nil {
		s.Error(err.Error())
		return
	}

	skillPanel.SetPosition(x, y+h1)

	s.renderPanelSegment(target, bottomFrame)

	// tab button highlighted
	switch tab {
	case firstTab:
		skillPanel.SetPosition(x+w, y+h1)
		s.renderPanelSegment(target, frameSelectedTab1Full)
	case secondTab:
		x += w
		skillPanel.SetPosition(x, s.originY+h0)
		s.renderPanelSegment(target, frameSelectedTab2Top)

		skillPanel.SetPosition(x, y+h1)
		s.renderPanelSegment(target, frameSelectedTab2Bottom)
	case thirdTab:
		skillPanel.SetPosition(x+w, y)
		s.renderPanelSegment(target, frameSelectedTab3Full)
	}
}

// Render the skill tree panel
func (s *skillTree) Render(target d2interface.Surface) {
	s.renderTabCommon(target)
	s.renderTab(target, s.selectedTab)
}
