package d2player

import (
	"fmt"
	"sort"

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
	skillIconWidth    = 48
	screenWidth       = 800
	screenHeight      = 600
	skillIconHeight   = 48
	rightPanelEndX    = 720
	leftPanelStartX   = 90
	skillPanelOffsetY = 465
	skillListsLength  = 5 // 0 to 4. 0 - General Skills, 1 to 3 - Class-specific skills(based on the 3 different skill trees), 4 - Other skills
)

// NewHeroSkillsPanel creates a new hero status panel
func NewHeroSkillsPanel(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	hero *d2mapentity.Player,
	l d2util.LogLevel,
	isLeftPanel bool) *SkillPanel {
	var activeSkill *d2hero.HeroSkill
	if isLeftPanel {
		activeSkill = hero.LeftSkill
	} else {
		activeSkill = hero.RightSkill
	}

	hoverTooltip := ui.NewTooltip(d2resource.Font16, d2resource.PaletteStatic, d2ui.TooltipXLeft, d2ui.TooltipYTop)

	skillPanel := &SkillPanel{
		asset:        asset,
		activeSkill:  activeSkill,
		ui:           ui,
		isOpen:       false,
		ListRows:     make([]*SkillListRow, skillListsLength),
		renderer:     ui.Renderer(),
		isLeftPanel:  isLeftPanel,
		hero:         hero,
		hoverTooltip: hoverTooltip,
	}

	skillPanel.Logger = d2util.NewLogger()
	skillPanel.Logger.SetLevel(l)
	skillPanel.Logger.SetPrefix(logPrefix)

	return skillPanel
}

// SkillPanel represents a skill select menu popup that is displayed when the player left clicks on his active left/right skill.
type SkillPanel struct {
	asset                *d2asset.AssetManager
	activeSkill          *d2hero.HeroSkill
	hero                 *d2mapentity.Player
	ListRows             []*SkillListRow
	renderer             d2interface.Renderer
	ui                   *d2ui.UIManager
	hoveredSkill         *d2hero.HeroSkill
	hoverTooltip         *d2ui.Tooltip
	isOpen               bool
	regenerateImageCache bool
	isLeftPanel          bool

	*d2util.Logger
}

// Open opens the hero skills panel
func (s *SkillPanel) Open() {
	s.isOpen = true
	s.regenerateImageCache = true
}

// Close the hero skills panel
func (s *SkillPanel) Close() {
	s.isOpen = false
}

// IsInRect returns whether the X Y coordinates are in some of the list rows of the panel.
func (s *SkillPanel) IsInRect(x, y int) bool {
	for _, listRow := range s.ListRows {
		if listRow != nil && listRow.IsInRect(x, y) {
			return true
		}
	}

	return false
}

// GetListRowByPos returns the skill list row for a given X and Y, based on the width and height of the skills list.
func (s *SkillPanel) GetListRowByPos(x, y int) *SkillListRow {
	for _, listRow := range s.ListRows {
		if listRow.IsInRect(x, y) {
			return listRow
		}
	}

	return nil
}

// Render gets called on every tick
func (s *SkillPanel) Render(target d2interface.Surface) error {
	if !s.isOpen {
		return nil
	}

	if s.regenerateImageCache {
		if err := s.generateSkillRowImageCache(); err != nil {
			return err
		}

		s.regenerateImageCache = false
	}

	renderedRows := 0

	for _, skillListRow := range s.ListRows {
		if len(skillListRow.Skills) == 0 {
			continue
		}

		startX := s.getRowStartX(skillListRow)
		rowOffsetY := skillPanelOffsetY - (renderedRows * skillIconHeight)

		target.PushTranslation(startX, rowOffsetY)
		target.Render(skillListRow.cachedImage)
		target.Pop()

		renderedRows++
	}

	if s.hoveredSkill != nil {
		s.hoverTooltip.Render(target)
	}

	return nil
}

// RegenerateImageCache will force re-generating the cached menu image on next Render.
// Somewhat expensive operation, should not be called often.
// Currently called every time the panel is opened or when the player learns a new skill.
func (s *SkillPanel) RegenerateImageCache() {
	s.regenerateImageCache = true
}

// IsOpen returns true if the hero skills panel is open
func (s *SkillPanel) IsOpen() bool {
	return s.isOpen
}

// Toggle toggles the visibility of the hero status panel
func (s *SkillPanel) Toggle() {
	if s.isOpen {
		s.Close()
	} else {
		s.Open()
	}
}

func (s *SkillPanel) generateSkillRowImageCache() error {
	for idx := range s.ListRows {
		s.ListRows[idx] = &SkillListRow{Skills: make([]*d2hero.HeroSkill, 0), Rectangle: d2geom.Rectangle{Height: 0, Width: 0}}
	}

	for _, skill := range s.hero.Skills {
		// left panel with an incompatible skill(e.g. Paladin auras cant be used as a left skill)
		if s.isLeftPanel && !skill.Leftskill {
			continue
		}

		// ListRow is -1 for other skills that should not be shown in the panel(e.g. Kick)
		if skill.ListRow == -1 || skill.Passive {
			continue
		}

		s.ListRows[skill.ListRow].AddSkill(skill)
	}

	visibleRows := 0

	for idx, skillListRow := range s.ListRows {
		// row won't be considered as visible
		if len(skillListRow.Skills) == 0 {
			continue
		}

		skillListRow.Rectangle = d2geom.Rectangle{
			Height: skillIconHeight,
			Width:  skillListRow.GetWidth(),
			Left:   s.getRowStartX(skillListRow),
			Top:    skillPanelOffsetY - (visibleRows * skillIconHeight),
		}

		skillRow := skillListRow

		sort.SliceStable(skillListRow.Skills, func(a, b int) bool {
			// left panel skills are aligned by ID (low to high), right panel is the opposite
			if s.isLeftPanel {
				return skillRow.Skills[a].ID < skillRow.Skills[b].ID
			}

			return skillRow.Skills[a].ID > skillRow.Skills[b].ID
		})

		cachedImage, err := s.createSkillListImage(skillListRow)

		if err != nil {
			s.Error(err.Error())
			return err
		}

		s.ListRows[idx].cachedImage = cachedImage
		visibleRows++
	}

	return nil
}

func (s *SkillPanel) createSkillListImage(skillsListRow *SkillListRow) (d2interface.Surface, error) {
	surface := s.renderer.NewSurface(len(skillsListRow.Skills)*skillIconWidth, skillIconHeight)

	lastSkillResourcePath := d2resource.GenericSkills
	skillSprite, _ := s.ui.NewSprite(s.getSkillResourceByClass(""), d2resource.PaletteSky)

	for idx, skill := range skillsListRow.Skills {
		currentResourcePath := s.getSkillResourceByClass(skill.Charclass)
		// only load a new sprite if the DCC file path changed
		if currentResourcePath != lastSkillResourcePath {
			lastSkillResourcePath = currentResourcePath
			skillSprite, _ = s.ui.NewSprite(currentResourcePath, d2resource.PaletteSky)
		}

		if skillSprite.GetFrameCount() <= skill.IconCel {
			// happens for non-player skills, since they do not have an icon
			s.Errorf("Invalid IconCel(sprite frame index) [%d] - Skill name: %s, skipping.", skill.IconCel, skill.Name)
			continue
		}

		if err := skillSprite.SetCurrentFrame(skill.IconCel); err != nil {
			return nil, err
		}

		surface.PushTranslation(idx*skillIconWidth, 50)
		skillSprite.Render(surface)
		surface.Pop()
	}

	return surface, nil
}

func (s *SkillPanel) getRowStartX(skillRow *SkillListRow) int {
	if s.isLeftPanel {
		return leftPanelStartX
	}

	// for the right panel, we only know where it should end, so we calculate the start based on the width of the list row
	return rightPanelEndX - skillRow.GetWidth()
}

func (s *SkillPanel) getSkillAtPos(x, y int) *d2hero.HeroSkill {
	listRow := s.GetListRowByPos(x, y)

	if listRow == nil {
		return nil
	}

	skillIndex := (x - s.getRowStartX(listRow)) / skillIconWidth
	skill := listRow.Skills[skillIndex]

	return skill
}

func (s *SkillPanel) getSkillIdxAtPos(x, y int) int {
	listRow := s.GetListRowByPos(x, y)

	if listRow == nil {
		return -1
	}

	skillIndex := (x - s.getRowStartX(listRow)) / skillIconWidth

	return skillIndex
}

// HandleClick will change the hero's active(left or right) skill and return true.
// Returns false if the given X, Y is out of panel boundaries.
func (s *SkillPanel) HandleClick(x, y int) bool {
	if !s.isOpen || !s.IsInRect(x, y) {
		return false
	}

	clickedSkill := s.getSkillAtPos(x, y)

	if clickedSkill == nil {
		return false
	}

	if s.isLeftPanel {
		s.hero.LeftSkill = clickedSkill
	} else {
		s.hero.RightSkill = clickedSkill
	}

	return true
}

// HandleMouseMove will process a mouse move event, if inside the panel.
func (s *SkillPanel) HandleMouseMove(x, y int) bool {
	if !s.isOpen {
		return false
	}

	if !s.IsInRect(x, y) {
		// panel still open but player hovered outside panel - hide the previously hovered skill(if any)
		s.hoveredSkill = nil
		return false
	}

	previousHovered := s.hoveredSkill
	s.hoveredSkill = s.getSkillAtPos(x, y)

	if previousHovered != s.hoveredSkill && s.hoveredSkill != nil {
		skillDescription := s.asset.TranslateString(s.hoveredSkill.ShortKey)
		s.hoverTooltip.SetText(fmt.Sprintf("%s\n%s", s.hoveredSkill.Skill, skillDescription))

		listRow := s.GetListRowByPos(x, y)

		tooltipX := (s.getSkillIdxAtPos(x, y) * skillIconWidth) + s.getRowStartX(listRow)
		tooltipY := listRow.Rectangle.Top + listRow.Rectangle.Height
		s.hoverTooltip.SetPosition(tooltipX, tooltipY)
	}

	return true
}

func (s *SkillPanel) getSkillResourceByClass(class string) string {
	resource := ""

	switch class {
	case "":
		resource = d2resource.GenericSkills
	case "bar":
		resource = d2resource.BarbarianSkills
	case "nec":
		resource = d2resource.NecromancerSkills
	case "pal":
		resource = d2resource.PaladinSkills
	case "ass":
		resource = d2resource.AssassinSkills
	case "sor":
		resource = d2resource.SorcererSkills
	case "ama":
		resource = d2resource.AmazonSkills
	case "dru":
		resource = d2resource.DruidSkills
	default:
		s.Errorf("Unknown class token: '%s'", class)
	}

	return resource
}
