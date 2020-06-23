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

type ExperienceBreakpointsRecord struct {
	Level           int
	HeroBreakpoints map[d2enum.Hero]int
	Ratio           int
}

var experienceStringMap map[string]d2enum.Hero
var experienceHeroMap map[d2enum.Hero]string

var ExperienceBreakpoints []*ExperienceBreakpointsRecord
var maxLevels map[d2enum.Hero]int

func GetMaxLevelByHero(heroType d2enum.Hero) int {
	return maxLevels[heroType]
}

func GetExperienceBreakpoint(heroType d2enum.Hero, level int) int {
	return ExperienceBreakpoints[level].HeroBreakpoints[heroType]
}

func LoadExperienceBreakpoints(file []byte) {
	d := d2common.LoadDataDictionary(string(file))

	experienceStringMap = map[string]d2enum.Hero{
		"Amazon":      d2enum.HeroAmazon,
		"Barbarian":   d2enum.HeroBarbarian,
		"Druid":       d2enum.HeroDruid,
		"Assassin":    d2enum.HeroAssassin,
		"Necromancer": d2enum.HeroNecromancer,
		"Paladin":     d2enum.HeroPaladin,
		"Sorceress":   d2enum.HeroSorceress,
	}

	experienceHeroMap = map[d2enum.Hero]string{
		d2enum.HeroAmazon:      "Amazon",
		d2enum.HeroBarbarian:   "Barbarian",
		d2enum.HeroDruid:       "Druid",
		d2enum.HeroAssassin:    "Assassin",
		d2enum.HeroNecromancer: "Necromancer",
		d2enum.HeroPaladin:     "Paladin",
		d2enum.HeroSorceress:   "Sorceress",
	}

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
