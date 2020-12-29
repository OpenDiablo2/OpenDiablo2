package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
)

// SkillListRow represents a row of skills that is shown when the skill select menu is rendered.
type SkillListRow struct {
	Rectangle   d2geom.Rectangle
	Skills      []*d2hero.HeroSkill
	cachedImage d2interface.Surface
}

// AddSkill appends to the skills of the row.
func (s *SkillListRow) AddSkill(skill *d2hero.HeroSkill) {
	s.Skills = append(s.Skills, skill)
}

// GetWidth returns the width based on the size of the skills.
func (s *SkillListRow) GetWidth() int {
	return skillIconWidth * len(s.Skills)
}

// GetRectangle returns the rectangle of the list.
func (s *SkillListRow) GetRectangle() d2geom.Rectangle {
	return s.Rectangle
}

// IsInRect returns true when the list has any skills and coordinates are in the rectangle of the list.
func (s *SkillListRow) IsInRect(x, y int) bool {
	// if there are no skills, row won't be rendered and it shouldn't be considered visible
	return len(s.Skills) > 0 && s.Rectangle.IsInRect(x, y)
}
