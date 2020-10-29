package d2player

import (
	"fmt"
	"log"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
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

	skillIconXOff  = 346
	skillIconYOff  = 59
	skillIconDistX = 69
	skillIconDistY = 68

	skillLabelXOffset = 49
	skillLabelYOffset = -4

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
	st.button.SetVisible(false)
	st.button.SetPosition(x, y)
}

type skillTreeHeroTypeResources struct {
	skillIcon      *d2ui.Sprite
	skillIconPath  string
	skillPanel     *d2ui.Sprite
	skillPanelPath string
}

type skillTree struct {
	resources     *skillTreeHeroTypeResources
	asset         *d2asset.AssetManager
	renderer      d2interface.Renderer
	guiManager    *d2gui.GuiManager
	uiManager     *d2ui.UIManager
	layout        *d2gui.Layout
	skills        map[int]*d2hero.HeroSkill
	heroClass     d2enum.Hero
	frame         *d2ui.UIFrame
	availSPLabel  *d2ui.Label
	skillLvlLabel *d2ui.Label
	closeButton   *d2ui.Button
	tab           [numTabs]*skillTreeTab
	isOpen        bool
	originX       int
	originY       int
	selectedTab   int
	onCloseCb     func()
}

func newSkillTree(
	skills map[int]*d2hero.HeroSkill,
	heroClass d2enum.Hero,
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
) *skillTree {
	st := &skillTree{
		skills:     skills,
		heroClass:  heroClass,
		asset:      asset,
		renderer:   renderer,
		uiManager:  ui,
		guiManager: guiManager,
		originX:    skillTreePanelX,
		originY:    skillTreePanelY,
		tab: [numTabs]*skillTreeTab{
			{},
			{},
			{},
		},
	}

	return st
}

func (s *skillTree) load() {
	s.frame = d2ui.NewUIFrame(s.asset, s.uiManager, d2ui.FrameRight)
	s.closeButton = s.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	s.closeButton.SetVisible(false)
	s.closeButton.OnActivated(func() { s.Close() })

	s.skillLvlLabel = s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)

	s.setHeroTypeResourcePath()
	s.loadForHeroType()
	s.setTab(0)
}

func (s *skillTree) loadForHeroType() {
	sp, err := s.uiManager.NewSprite(s.resources.skillPanelPath, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	s.resources.skillPanel = sp

	si, err := s.uiManager.NewSprite(s.resources.skillIconPath, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	s.resources.skillIcon = si

	s.tab[firstTab].createButton(s.uiManager, tabButtonX, tabButton0Y)
	s.tab[firstTab].button.OnActivated(func() { s.setTab(firstTab) })

	s.tab[secondTab].createButton(s.uiManager, tabButtonX, tabButton1Y)
	s.tab[secondTab].button.OnActivated(func() { s.setTab(secondTab) })

	s.tab[thirdTab].createButton(s.uiManager, tabButtonX, tabButton2Y)
	s.tab[thirdTab].button.OnActivated(func() { s.setTab(thirdTab) })

	s.availSPLabel = s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	s.availSPLabel.SetPosition(availSPLabelX, availSPLabelY)
	s.availSPLabel.Alignment = d2gui.HorizontalAlignCenter
	s.availSPLabel.SetText(makeTabString("StrSklTree1", "StrSklTree2", "StrSklTree3"))
}

type heroTabData struct {
	resources        *skillTreeHeroTypeResources
	str1, str2, str3 string
	closeButtonPos   [numTabs]int
}

func makeTabString(keys ...interface{}) string {
	translations := make([]interface{}, len(keys))

	token := "%s"
	format := token

	for idx, key := range keys {
		if idx > 0 {
			format += "\n" + token
		}

		translations[idx] = d2tbl.TranslateString(key.(string))
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
			makeTabString("StrSklTree21", "StrSklTree4"),
			makeTabString("StrSklTree21", "StrSklTree22"),
			makeTabString("StrSklTree20"),
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
			makeTabString("StrSklTree19"),
			makeTabString("StrSklTree17", "StrSklTree18", "StrSklTree5"),
			makeTabString("StrSklTree16", "StrSklTree5"),
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
			makeTabString("StrSklTree15", "StrSklTree4"),
			makeTabString("StrSklTree14", "StrSklTree13"),
			makeTabString("StrSklTree12", "StrSklTree13"),
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
			makeTabString("StrSklTree30"),
			makeTabString("StrSklTree31", "StrSklTree32"),
			makeTabString("StrSklTree33", "StrSklTree34"),
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
			makeTabString("StrSklTree25", "StrSklTree5"),
			makeTabString("StrSklTree24", "StrSklTree5"),
			makeTabString("StrSklTree23", "StrSklTree5"),
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
			makeTabString("StrSklTree10", "StrSklTree11", "StrSklTree4"),
			makeTabString("StrSklTree8", "StrSklTree9", "StrSklTree4"),
			makeTabString("StrSklTree6", "StrSklTree7", "StrSklTree4"),
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
			makeTabString("StrSklTree26"),
			makeTabString("StrSklTree27", "StrSklTree28"),
			makeTabString("StrSklTree29"),
			makeCloseButtonPos(
				skillCloseButtonXRight,
				skillCloseButtonXRight,
				skillCloseButtonXRight),
		},
	}

	return tabMap[class]
}

func (s *skillTree) setHeroTypeResourcePath() {
	entry := s.getTab(s.heroClass)
	if entry == nil {
		log.Fatal("Unknown Hero Type")
	}

	s.resources = entry.resources
	s.tab[firstTab].buttonText = entry.str1
	s.tab[secondTab].buttonText = entry.str2
	s.tab[thirdTab].buttonText = entry.str3

	for i := 0; i < numTabs; i++ {
		s.tab[i].closeButtonPosX = entry.closeButtonPos[i]
	}
}

// Toggle the skill tree visibility
func (s *skillTree) Toggle() {
	fmt.Println("SkillTree toggled")

	if s.isOpen {
		s.Close()
	} else {
		s.Open()
	}
}

// Close the skill tree
func (s *skillTree) Close() {
	s.isOpen = false
	s.guiManager.SetLayout(nil)
	s.closeButton.SetVisible(false)

	for i := 0; i < numTabs; i++ {
		s.tab[i].button.SetVisible(false)
	}

	s.onCloseCb()
}

// Open the skill tree
func (s *skillTree) Open() {
	s.isOpen = true
	if s.layout == nil {
		s.layout = d2gui.CreateLayout(s.renderer, d2gui.PositionTypeHorizontal, s.asset)
	}

	s.closeButton.SetVisible(true)

	for i := 0; i < numTabs; i++ {
		s.tab[i].button.SetVisible(true)
	}

	s.guiManager.SetLayout(s.layout)
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
}

func (s *skillTree) renderPanelSegment(
	target d2interface.Surface,
	frame int) error {
	if err := s.resources.skillPanel.SetCurrentFrame(frame); err != nil {
		return err
	}

	s.resources.skillPanel.Render(target)

	return nil
}

func (s *skillTree) renderTabCommon(target d2interface.Surface) error {
	skillPanel := s.resources.skillPanel
	x, y := s.originX, s.originY

	// top
	w, h, err := skillPanel.GetFrameSize(frameCommonTabTopLeft)
	if err != nil {
		return err
	}

	y += h

	skillPanel.SetPosition(x, y)

	err = s.renderPanelSegment(target, frameCommonTabTopLeft)
	if err != nil {
		return err
	}

	skillPanel.SetPosition(x+w, y)

	err = s.renderPanelSegment(target, frameCommonTabTopRight)
	if err != nil {
		return err
	}

	// bottom
	_, h, err = skillPanel.GetFrameSize(frameCommonTabBottomLeft)
	if err != nil {
		return err
	}

	y += h

	skillPanel.SetPosition(x, y)

	err = s.renderPanelSegment(target, frameCommonTabBottomLeft)
	if err != nil {
		return err
	}

	skillPanel.SetPosition(x+w, y)

	err = s.renderPanelSegment(target, frameCommonTabBottomRight)
	if err != nil {
		return err
	}

	// available skill points label
	s.availSPLabel.Render(target)

	return nil
}

func (s *skillTree) renderTab(target d2interface.Surface, tab int) error {
	topFrame := frameOffsetTop + (tabIndexOffset * tab)
	bottomFrame := frameOffsetBottom + (tabIndexOffset * tab)

	skillPanel := s.resources.skillPanel
	x, y := s.originX, s.originY

	// top
	_, h0, err := skillPanel.GetFrameSize(topFrame)
	if err != nil {
		return err
	}

	y += h0

	skillPanel.SetPosition(x, y)

	err = s.renderPanelSegment(target, topFrame)
	if err != nil {
		return err
	}

	// bottom
	w, h1, err := skillPanel.GetFrameSize(bottomFrame)
	if err != nil {
		return err
	}

	skillPanel.SetPosition(x, y+h1)

	if err := s.renderPanelSegment(target, bottomFrame); err != nil {
		return err
	}

	// tab button highlighted
	switch tab {
	case firstTab:
		skillPanel.SetPosition(x+w, y+h1)

		if err := s.renderPanelSegment(target, frameSelectedTab1Full); err != nil {
			return err
		}
	case secondTab:
		x += w
		skillPanel.SetPosition(x, s.originY+h0)

		if err := s.renderPanelSegment(target, frameSelectedTab2Top); err != nil {
			return err
		}

		skillPanel.SetPosition(x, y+h1)

		if err := s.renderPanelSegment(target, frameSelectedTab2Bottom); err != nil {
			return err
		}
	case thirdTab:
		skillPanel.SetPosition(x+w, y)

		if err := s.renderPanelSegment(target, frameSelectedTab3Full); err != nil {
			return err
		}
	}

	return nil
}

func (s *skillTree) renderSkillIcon(target d2interface.Surface, skill *d2hero.HeroSkill) error {
	skillIcon := s.resources.skillIcon
	if err := skillIcon.SetCurrentFrame(skill.IconCel); err != nil {
		return err
	}

	x := skillIconXOff + skill.SkillColumn*skillIconDistX
	y := skillIconYOff + skill.SkillRow*skillIconDistY

	skillIcon.SetPosition(x, y)

	if skill.SkillPoints == 0 {
		target.PushSaturation(skillIconGreySat)
		defer target.Pop()

		target.PushBrightness(skillIconGreyBright)
		defer target.Pop()
	}

	skillIcon.Render(target)

	return nil
}

func (s *skillTree) renderSkillIconLabel(target d2interface.Surface, skill *d2hero.HeroSkill) {
	if skill.SkillPoints == 0 {
		return
	}

	s.skillLvlLabel.SetText(strconv.Itoa(skill.SkillPoints))
	x := skillIconXOff + skill.SkillColumn*skillIconDistX + skillLabelXOffset
	y := skillIconYOff + skill.SkillRow*skillIconDistY + skillLabelYOffset
	s.skillLvlLabel.SetPosition(x, y)
	s.skillLvlLabel.Render(target)
}

func (s *skillTree) renderSkillIcons(target d2interface.Surface, tab int) error {
	for idx := range s.skills {
		skill := s.skills[idx]
		if skill.SkillPage != tab+1 {
			continue
		}

		if err := s.renderSkillIcon(target, skill); err != nil {
			return err
		}

		s.renderSkillIconLabel(target, skill)
	}

	return nil
}

// Render the skill tree panel
func (s *skillTree) Render(target d2interface.Surface) error {
	if !s.isOpen {
		return nil
	}

	if err := s.frame.Render(target); err != nil {
		return err
	}

	if err := s.renderTabCommon(target); err != nil {
		return err
	}

	if err := s.renderTab(target, s.selectedTab); err != nil {
		return err
	}

	if err := s.renderSkillIcons(target, s.selectedTab); err != nil {
		return err
	}

	return nil
}
