package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

/*	first column of experience.txt
	Level
	Amazon
	Sorceress
	Necromancer
	Paladin
	Barbarian
	Druid
	Assassin
	ExpRatio

	second row is special case, shows max levels

	MaxLvl
	99
	99
	99
	99
	99
	99
	99
	10
*/

// ExperienceBreakpointsRecord describes the experience points required to
// gain a level for all character classes
type ExperienceBreakpointsRecord struct {
	Level           int
	HeroBreakpoints map[d2enum.Hero]int
	Ratio           int
}

// ExperienceBreakpoints describes the required experience
// for each level for each character class
//nolint:gochecknoglobals // Currently global by design, only written once
var ExperienceBreakpoints []*ExperienceBreakpointsRecord

//nolint:gochecknoglobals // Currently global by design
var maxLevels map[d2enum.Hero]int

// GetMaxLevelByHero returns the highest level attainable for a hero type
func GetMaxLevelByHero(heroType d2enum.Hero) int {
	return maxLevels[heroType]
}

// GetExperienceBreakpoint given a hero type and a level, returns the experience required for the level
func GetExperienceBreakpoint(heroType d2enum.Hero, level int) int {
	return ExperienceBreakpoints[level].HeroBreakpoints[heroType]
}

// LoadExperienceBreakpoints loads experience.txt into a map
// ExperienceBreakpoints []*ExperienceBreakpointsRecord
func LoadExperienceBreakpoints(file []byte) {
	d := d2common.LoadDataDictionary(string(file))

	// we skip the second row because that describes max level of char classes
	ExperienceBreakpoints = make([]*ExperienceBreakpointsRecord, len(d.Data)-1)

	for idx := range d.Data {
		if idx == 0 {
			// max levels are a special case
			maxLevels = map[d2enum.Hero]int{
				d2enum.HeroAmazon:      d.GetNumber("Amazon", idx),
				d2enum.HeroBarbarian:   d.GetNumber("Barbarian", idx),
				d2enum.HeroDruid:       d.GetNumber("Druid", idx),
				d2enum.HeroAssassin:    d.GetNumber("Assassin", idx),
				d2enum.HeroNecromancer: d.GetNumber("Necromancer", idx),
				d2enum.HeroPaladin:     d.GetNumber("Paladin", idx),
				d2enum.HeroSorceress:   d.GetNumber("Sorceress", idx),
			}

			continue
		}

		record := &ExperienceBreakpointsRecord{
			Level: d.GetNumber("Level", idx),
			HeroBreakpoints: map[d2enum.Hero]int{
				d2enum.HeroAmazon:      d.GetNumber("Amazon", idx),
				d2enum.HeroBarbarian:   d.GetNumber("Barbarian", idx),
				d2enum.HeroDruid:       d.GetNumber("Druid", idx),
				d2enum.HeroAssassin:    d.GetNumber("Assassin", idx),
				d2enum.HeroNecromancer: d.GetNumber("Necromancer", idx),
				d2enum.HeroPaladin:     d.GetNumber("Paladin", idx),
				d2enum.HeroSorceress:   d.GetNumber("Sorceress", idx),
			},
			Ratio: d.GetNumber("ExpRatio", idx),
		}

		ExperienceBreakpoints[record.Level] = record
	}

	log.Printf("Loaded %d ExperienceBreakpoint records", len(ExperienceBreakpoints))
}
