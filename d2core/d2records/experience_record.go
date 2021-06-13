package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// ExperienceBreakpoints describes the required experience
// for each level for each character class
type ExperienceBreakpoints map[int]*ExperienceBreakpointRecord

// ExperienceMaxLevels defines the max character levels
type ExperienceMaxLevels map[d2enum.Hero]int

// ExperienceBreakpointRecord describes the experience points required to
// gain a level for all character classes
type ExperienceBreakpointRecord struct {
	HeroBreakpoints map[d2enum.Hero]int
	Level           int
	Ratio           int
}
