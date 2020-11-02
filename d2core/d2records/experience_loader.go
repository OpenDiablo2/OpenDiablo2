package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
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

	the rest are the breakpoints records
*/

func experienceLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	breakpoints := make(ExperienceBreakpoints)

	d.Next() // move to the first row, the max level data

	// parse the max level data
	maxLevels := ExperienceMaxLevels{
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
		breakpoints[record.Level] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d Experience Breakpoint records", len(breakpoints))

	r.Character.MaxLevel = maxLevels
	r.Character.Experience = breakpoints

	return nil
}
