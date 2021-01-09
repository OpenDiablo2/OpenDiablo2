package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// ExperienceBreakpoints describes the required experience
// for each level for each character class
type ExperienceBreakpoints map[int]*ExperienceBreakpointsRecord

// ExperienceMaxLevels defines the max character levels
type ExperienceMaxLevels map[d2enum.Hero]int

// ExperienceBreakpointsRecord describes the experience points required to
// gain a level for all character classes
type ExperienceBreakpointsRecord struct {
	Level           int
	HeroBreakpoints map[d2enum.Hero]int
	Ratio           int
}
