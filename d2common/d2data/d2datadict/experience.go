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
var ExperienceBreakpoints map[int]*ExperienceBreakpointsRecord

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
	ExperienceBreakpoints = make(map[int]*ExperienceBreakpointsRecord)

	d := d2common.LoadDataDictionary(file)
	d.Next()

	// the first row describes the max level of char classes
	maxLevels = map[d2enum.Hero]int{
		d2enum.HeroAmazon:      d.Number("Amazon"),
		d2enum.HeroBarbarian:   d.Number("Barbarian"),
		d2enum.HeroDruid:       d.Number("Druid"),
		d2enum.HeroAssassin:    d.Number("Assassin"),
		d2enum.HeroNecromancer: d.Number("Necromancer"),
		d2enum.HeroPaladin:     d.Number("Paladin"),
		d2enum.HeroSorceress:   d.Number("Sorceress"),
	}

	for d.Next() {
		record := &ExperienceBreakpointsRecord{
			Level: d.Number("Level"),
			HeroBreakpoints: map[d2enum.Hero]int{
				d2enum.HeroAmazon:      d.Number("Amazon"),
				d2enum.HeroBarbarian:   d.Number("Barbarian"),
				d2enum.HeroDruid:       d.Number("Druid"),
				d2enum.HeroAssassin:    d.Number("Assassin"),
				d2enum.HeroNecromancer: d.Number("Necromancer"),
				d2enum.HeroPaladin:     d.Number("Paladin"),
				d2enum.HeroSorceress:   d.Number("Sorceress"),
			},
			Ratio: d.Number("ExpRatio"),
		}
		ExperienceBreakpoints[record.Level] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d ExperienceBreakpoint records", len(ExperienceBreakpoints))
}
