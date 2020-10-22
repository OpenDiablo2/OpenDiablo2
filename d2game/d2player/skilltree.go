package d2player

import (
	"fmt"
	"log"

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
	TabButtonX  = 628
	TabButton0Y = 385
	TabButton1Y = 277
	TabButton2Y = 170

	AvailSPLabelX = 677
	AvailSPLabelY = 72
)

type SkillTreeTab struct {
	buttonText  string
	button      *d2ui.Button
}

func (st *SkillTreeTab) CreateButton(uiManager *d2ui.UIManager, x int, y int) {
	st.button = uiManager.NewButton(d2ui.ButtonTypeSkillTreeTab, st.buttonText)
	st.button.SetVisible(false)
	st.button.SetPosition(x, y)
}

type SkillTreeHeroTypeResources struct {
	skillIcon     *d2ui.Sprite
	skillIconPath string
	skillPanel    *d2ui.Sprite
	skillPanelPath string
}

type SkillTree struct {
	resources    *SkillTreeHeroTypeResources
	asset        *d2asset.AssetManager
	renderer     d2interface.Renderer
	guiManager   *d2gui.GuiManager
	uiManager    *d2ui.UIManager
	layout       *d2gui.Layout
	skills       map[int]*d2hero.HeroSkill
	heroClass    d2enum.Hero
	frame        *d2ui.UIFrame
	availSPLabel *d2ui.Label
	tab          [3]*SkillTreeTab
	isOpen       bool
	originX      int
	originY      int
	selectedTab  int
}

func NewSkillTree(
	skills map[int]*d2hero.HeroSkill,
	heroClass  d2enum.Hero,
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
) *SkillTree {
	st := &SkillTree {
		skills: skills,
		heroClass: heroClass,
		asset: asset,
		renderer: renderer,
		uiManager: ui,
		guiManager: guiManager,
		originX: 401,
		originY: 64,
		tab: [3]*SkillTreeTab{
			{},
			{},
			{},
		},
	}
	return st
}

func (s *SkillTree) Load() {

	s.frame = d2ui.NewUIFrame(s.asset, s.uiManager, d2ui.FrameRight)

	s.setHeroTypeResourcePath()
	s.LoadForHeroType()
	s.setTab(0)
}

func (s *SkillTree) LoadForHeroType() {
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

	s.tab[0].CreateButton(s.uiManager, TabButtonX, TabButton0Y)
	s.tab[0].button.OnActivated(func() { s.setTab(0) })
	s.tab[1].CreateButton(s.uiManager, TabButtonX, TabButton1Y)
	s.tab[1].button.OnActivated(func() { s.setTab(1) })
	s.tab[2].CreateButton(s.uiManager, TabButtonX, TabButton2Y)
	s.tab[2].button.OnActivated(func() { s.setTab(2) })

	s.availSPLabel = s.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteSky)
	s.availSPLabel.SetPosition(AvailSPLabelX, AvailSPLabelY)
	s.availSPLabel.Alignment = d2gui.HorizontalAlignCenter
	s.availSPLabel.SetText(fmt.Sprintf("%s\n%s\n%s",
		d2tbl.TranslateString("StrSklTree1"),
		d2tbl.TranslateString("StrSklTree2"),
		d2tbl.TranslateString("StrSklTree3"),
	))
}

func (s *SkillTree) setHeroTypeResourcePath() {
	var res *SkillTreeHeroTypeResources

	switch s.heroClass {
	case d2enum.HeroBarbarian:
		res = &SkillTreeHeroTypeResources {
			skillPanelPath: d2resource.SkillsPanelBarbarian,
			skillIconPath: d2resource.BarbarianSkills,
		}

		s.tab[0].buttonText = fmt.Sprintf("%s\n%s",
			d2tbl.TranslateString("StrSklTree21"),
			d2tbl.TranslateString("StrSklTree4"))
		s.tab[1].buttonText = fmt.Sprintf("%s\n%s",
			d2tbl.TranslateString("StrSklTree21"),
			d2tbl.TranslateString("StrSklTree22"))
		s.tab[2].buttonText = d2tbl.TranslateString("StrSklTree20")

		case d2enum.HeroNecromancer:
			res = &SkillTreeHeroTypeResources {
			skillPanelPath: d2resource.SkillsPanelNecromancer,
				skillIconPath: d2resource.NecromancerSkills,
			}

		        s.tab[0].buttonText = d2tbl.TranslateString("StrSklTree19")
			s.tab[1].buttonText = fmt.Sprintf("%s\n%s\n%s",
					d2tbl.TranslateString("StrSklTree17"),
					d2tbl.TranslateString("StrSklTree18"),
					d2tbl.TranslateString("StrSklTree5"))
			s.tab[2].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree16"),
					d2tbl.TranslateString("StrSklTree5"))
		case d2enum.HeroPaladin:
			res = &SkillTreeHeroTypeResources {
			skillPanelPath: d2resource.SkillsPanelPaladin,
				skillIconPath: d2resource.PaladinSkills,
			}

			s.tab[0].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree15"),
					d2tbl.TranslateString("StrSklTree4"))
			s.tab[1].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree14"),
					d2tbl.TranslateString("StrSklTree13"))
			s.tab[2].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree12"),
					d2tbl.TranslateString("StrSklTree13"))
		case d2enum.HeroAssassin:
			res = &SkillTreeHeroTypeResources {
				skillPanelPath: d2resource.SkillsPanelAssassin,
				skillIconPath: d2resource.AssassinSkills,
			}

			s.tab[0].buttonText = d2tbl.TranslateString("StrSklTree30")
			s.tab[1].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree31"),
					d2tbl.TranslateString("StrSklTree32"))
			s.tab[2].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree33"),
					d2tbl.TranslateString("StrSklTree34"))
		case d2enum.HeroSorceress:
			res = &SkillTreeHeroTypeResources {
				skillPanelPath: d2resource.SkillsPanelSorcerer,
					skillIconPath: d2resource.SorcererSkills,
			}
			s.tab[0].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree25"),
					d2tbl.TranslateString("StrSklTree5"))
			s.tab[1].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree24"),
					d2tbl.TranslateString("StrSklTree5"))
			s.tab[2].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree23"),
					d2tbl.TranslateString("StrSklTree5"))
		case d2enum.HeroAmazon:
			res = &SkillTreeHeroTypeResources {
				skillPanelPath: d2resource.SkillsPanelAmazon,
				skillIconPath: d2resource.AmazonSkills,
			}
			s.tab[0].buttonText = fmt.Sprintf("%s\n%s\n%s",
					d2tbl.TranslateString("StrSklTree10"),
					d2tbl.TranslateString("StrSklTree11"),
					d2tbl.TranslateString("StrSklTree4"))
			s.tab[1].buttonText = fmt.Sprintf("%s\n%s\n%s",
					d2tbl.TranslateString("StrSklTree8"),
					d2tbl.TranslateString("StrSklTree9"),
					d2tbl.TranslateString("StrSklTree4"))
			s.tab[2].buttonText = fmt.Sprintf("%s\n%s\n%s",
					d2tbl.TranslateString("StrSklTree6"),
					d2tbl.TranslateString("StrSklTree7"),
					d2tbl.TranslateString("StrSklTree4"))
		case d2enum.HeroDruid:
			res = &SkillTreeHeroTypeResources {
				skillPanelPath: d2resource.SkillsPanelDruid,
				skillIconPath: d2resource.DruidSkills,
			}
			s.tab[0].buttonText = d2tbl.TranslateString("StrSklTree26")
			s.tab[1].buttonText = fmt.Sprintf("%s\n%s",
					d2tbl.TranslateString("StrSklTree27"),
					d2tbl.TranslateString("StrSklTree28"))
			s.tab[2].buttonText = d2tbl.TranslateString("StrSklTree29")
	default:
		log.Fatal("Unknown Hero Type")
	}
	s.resources = res
}

func (s *SkillTree) Toggle() {
	fmt.Println("SkillTree toggled")
	if s.isOpen {
		s.Close()
	} else {
		s.Open()
	}
}

func (s *SkillTree) Close() {
	s.isOpen = false
	s.guiManager.SetLayout(nil)
	for i:=0; i < 3; i++ {
		s.tab[i].button.SetVisible(false)
	}
}

func (s *SkillTree) Open() {
	s.isOpen = true
	if s.layout == nil {
		s.layout = d2gui.CreateLayout(s.renderer, d2gui.PositionTypeHorizontal, s.asset)
	}

	for i:=0; i < 3; i++ {
		s.tab[i].button.SetVisible(true)
	}
	s.guiManager.SetLayout(s.layout)

}

func (s *SkillTree) IsOpen() bool {
	return s.isOpen
}

func (s *SkillTree) setTab(tab int) {
	s.selectedTab = tab
}

func (s *SkillTree) renderPanelSegment(
	target d2interface.Surface,
	frame int) error {
	if err := s.resources.skillPanel.SetCurrentFrame(frame); err != nil {
		return err
	}
	if err := s.resources.skillPanel.Render(target); err != nil {
		return err
	}
	return nil
}

func (s *SkillTree) renderTabCommon (target d2interface.Surface) error {
	skillPanel := s.resources.skillPanel
	x, y := s.originX, s.originY

	// top
	w, h, err := skillPanel.GetFrameSize(0)
	if err != nil {
		return err
	}
	y += h
	skillPanel.SetPosition(x, y)
	if err:= s.renderPanelSegment(target, 0); err != nil {
		return err
	}

	skillPanel.SetPosition(x+w, y)
	if err:= s.renderPanelSegment(target, 1); err != nil {
		return err
	}

	// bottom
	_, h, err = skillPanel.GetFrameSize(2)
	if err != nil {
		return err
	}
	y += h
	skillPanel.SetPosition(x, y)
	if err:= s.renderPanelSegment(target, 2); err != nil {
		return err
	}

	skillPanel.SetPosition(x+w, y)
	if err:= s.renderPanelSegment(target, 3); err != nil {
		return err
	}

	// available skill points label
	s.availSPLabel.Render(target)
	return nil
}

func (s *SkillTree) renderTab (target d2interface.Surface, tab int) error {
	var frameID [2]int

	frameID[0] = 4 + (4*tab)
	frameID[1] = 6 + (4*tab)

	skillPanel := s.resources.skillPanel
	x, y := s.originX, s.originY

	// top
	_, h0, err := skillPanel.GetFrameSize(frameID[0])
	if err != nil {
		return err
	}
	y += h0
	skillPanel.SetPosition(x, y)
	if err:= s.renderPanelSegment(target, frameID[0]); err != nil {
		return err
	}

	// bottom
	w, h1, err := skillPanel.GetFrameSize(frameID[1])
	if err != nil {
		return err
	}
	skillPanel.SetPosition(x, y+h1)
	if err:= s.renderPanelSegment(target, frameID[1]); err != nil {
		return err
	}

	// tab button highlighted
	switch tab {
	case 0:
		skillPanel.SetPosition(x+w, y+h1)
		if err:= s.renderPanelSegment(target, 7); err != nil {
			return err
		}
	case 1:
		x += w
		skillPanel.SetPosition(x, s.originY + h0)
		if err:= s.renderPanelSegment(target, 9); err != nil {
			return err
		}
		skillPanel.SetPosition(x, y + h1)
		if err:= s.renderPanelSegment(target, 11); err != nil {
			return err
		}
	case 2:
		skillPanel.SetPosition(x+w, y)
		if err:= s.renderPanelSegment(target, 13); err != nil {
			return err
		}
	}
	return nil
}

func (s *SkillTree) Render (target d2interface.Surface) error {
	if !s.isOpen {
		return nil
	}
	s.frame.Render(target)
	s.renderTabCommon(target)
	s.renderTab(target, s.selectedTab)
	return nil
}
