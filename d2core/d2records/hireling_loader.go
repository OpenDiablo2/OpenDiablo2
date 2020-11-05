package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadHireling loads hireling data into []*HirelingRecord
func hirelingLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make([]*HirelingRecord, 0)

	for d.Next() {
		hireling := &HirelingRecord{
			Hireling:        d.String("Hireling"),
			SubType:         d.String("SubType"),
			ID:              d.Number("Id"),
			Class:           d.Number("Class"),
			Act:             d.Number("Act"),
			Difficulty:      d.Number("Difficulty"),
			Level:           d.Number("Level"),
			Seller:          d.Number("Seller"),
			NameFirst:       d.String("NameFirst"),
			NameLast:        d.String("NameLast"),
			Gold:            d.Number("Gold"),
			ExpPerLvl:       d.Number("Exp/Lvl"),
			HP:              d.Number("HP"),
			HPPerLvl:        d.Number("HP/Lvl"),
			Defense:         d.Number("Defense"),
			DefPerLvl:       d.Number("Id"),
			Str:             d.Number("Str"),
			StrPerLvl:       d.Number("Str/Lvl"),
			Dex:             d.Number("Dex"),
			DexPerLvl:       d.Number("Dex/Lvl"),
			AR:              d.Number("AR"),
			ARPerLvl:        d.Number("AR/Lvl"),
			Share:           d.Number("Share"),
			DmgMin:          d.Number("Dmg-Min"),
			DmgMax:          d.Number("Dmg-Max"),
			DmgPerLvl:       d.Number("Dmg/Lvl"),
			Resist:          d.Number("Resist"),
			ResistPerLvl:    d.Number("Resist/Lvl"),
			WType1:          d.String("WType1"),
			WType2:          d.String("WType2"),
			HireDesc:        d.String("HireDesc"),
			DefaultChance:   d.Number("DefaultChance"),
			Skill1:          d.String("Skill1"),
			Mode1:           d.Number("Mode1"),
			Chance1:         d.Number("Chance1"),
			ChancePerLevel1: d.Number("ChancePerLvl1"),
			Level1:          d.Number("Level1"),
			LvlPerLvl1:      d.Number("LvlPerLvl1"),
			Skill2:          d.String("Skill2"),
			Mode2:           d.Number("Mode2"),
			Chance2:         d.Number("Chance2"),
			ChancePerLevel2: d.Number("ChancePerLvl2"),
			Level2:          d.Number("Level2"),
			LvlPerLvl2:      d.Number("LvlPerLvl2"),
			Skill3:          d.String("Skill3"),
			Mode3:           d.Number("Mode3"),
			Chance3:         d.Number("Chance3"),
			ChancePerLevel3: d.Number("ChancePerLvl3"),
			Level3:          d.Number("Level3"),
			LvlPerLvl3:      d.Number("LvlPerLvl3"),
			Skill4:          d.String("Skill4"),
			Mode4:           d.Number("Mode4"),
			Chance4:         d.Number("Chance4"),
			ChancePerLevel4: d.Number("ChancePerLvl4"),
			Level4:          d.Number("Level4"),
			LvlPerLvl4:      d.Number("LvlPerLvl4"),
			Skill5:          d.String("Skill5"),
			Mode5:           d.Number("Mode5"),
			Chance5:         d.Number("Chance5"),
			ChancePerLevel5: d.Number("ChancePerLvl5"),
			Level5:          d.Number("Level5"),
			LvlPerLvl5:      d.Number("LvlPerLvl5"),
			Skill6:          d.String("Skill6"),
			Mode6:           d.Number("Mode6"),
			Chance6:         d.Number("Chance6"),
			ChancePerLevel6: d.Number("ChancePerLvl6"),
			Level6:          d.Number("Level6"),
			LvlPerLvl6:      d.Number("LvlPerLvl6"),
			Head:            d.Number("Head"),
			Torso:           d.Number("Torso"),
			Weapon:          d.Number("Weapon"),
			Shield:          d.Number("Shield"),
		}
		records = append(records, hireling)
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d Hireling records", len(records))

	r.Hireling.Details = records

	return nil
}
