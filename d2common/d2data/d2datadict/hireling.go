package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// HirelingRecord is a representation of rows in hireling.txt
// these records describe mercenaries
type HirelingRecord struct {
	Hireling        string
	SubType         string
	ID              int
	Class           int
	Act             int
	Difficulty      int
	Level           int
	Seller          int
	NameFirst       string
	NameLast        string
	Gold            int
	ExpPerLvl       int
	HP              int
	HPPerLvl        int
	Defense         int
	DefPerLvl       int
	Str             int
	StrPerLvl       int
	Dex             int
	DexPerLvl       int
	AR              int
	ARPerLvl        int
	Share           int
	DmgMin          int
	DmgMax          int
	DmgPerLvl       int
	Resist          int
	ResistPerLvl    int
	WType1          string
	WType2          string
	HireDesc        string
	DefaultChance   int
	Skill1          string
	Mode1           int
	Chance1         int
	ChancePerLevel1 int
	Level1          int
	LvlPerLvl1      int
	Skill2          string
	Mode2           int
	Chance2         int
	ChancePerLevel2 int
	Level2          int
	LvlPerLvl2      int
	Skill3          string
	Mode3           int
	Chance3         int
	ChancePerLevel3 int
	Level3          int
	LvlPerLvl3      int
	Skill4          string
	Mode4           int
	Chance4         int
	ChancePerLevel4 int
	Level4          int
	LvlPerLvl4      int
	Skill5          string
	Mode5           int
	Chance5         int
	ChancePerLevel5 int
	Level5          int
	LvlPerLvl5      int
	Skill6          string
	Mode6           int
	Chance6         int
	ChancePerLevel6 int
	Level6          int
	LvlPerLvl6      int
	Head            int
	Torso           int
	Weapon          int
	Shield          int
}

// Hirelings stores hireling (mercenary) records
//nolint:gochecknoglobals // Currently global by design, only written once
var Hirelings []*HirelingRecord

// LoadHireling loads hireling data into []*HirelingRecord
func LoadHireling(file []byte) {
	d := d2common.LoadDataDictionary(string(file))

	Hirelings = make([]*HirelingRecord, len(d.Data))

	for idx := range d.Data {
		hireling := &HirelingRecord{
			Hireling:        d.GetString("Hireling", idx),
			SubType:         d.GetString("SubType", idx),
			ID:              d.GetNumber("Id", idx),
			Class:           d.GetNumber("Class", idx),
			Act:             d.GetNumber("Act", idx),
			Difficulty:      d.GetNumber("Difficulty", idx),
			Level:           d.GetNumber("Level", idx),
			Seller:          d.GetNumber("Seller", idx),
			NameFirst:       d.GetString("NameFirst", idx),
			NameLast:        d.GetString("NameLast", idx),
			Gold:            d.GetNumber("Gold", idx),
			ExpPerLvl:       d.GetNumber("Exp/Lvl", idx),
			HP:              d.GetNumber("HP", idx),
			HPPerLvl:        d.GetNumber("HP/Lvl", idx),
			Defense:         d.GetNumber("Defense", idx),
			DefPerLvl:       d.GetNumber("Id", idx),
			Str:             d.GetNumber("Str", idx),
			StrPerLvl:       d.GetNumber("Str/Lvl", idx),
			Dex:             d.GetNumber("Dex", idx),
			DexPerLvl:       d.GetNumber("Dex/Lvl", idx),
			AR:              d.GetNumber("AR", idx),
			ARPerLvl:        d.GetNumber("AR/Lvl", idx),
			Share:           d.GetNumber("Share", idx),
			DmgMin:          d.GetNumber("Dmg-Min", idx),
			DmgMax:          d.GetNumber("Dmg-Max", idx),
			DmgPerLvl:       d.GetNumber("Dmg/Lvl", idx),
			Resist:          d.GetNumber("Resist", idx),
			ResistPerLvl:    d.GetNumber("Resist/Lvl", idx),
			WType1:          d.GetString("WType1", idx),
			WType2:          d.GetString("WType2", idx),
			HireDesc:        d.GetString("HireDesc", idx),
			DefaultChance:   d.GetNumber("DefaultChance", idx),
			Skill1:          d.GetString("Skill1", idx),
			Mode1:           d.GetNumber("Mode1", idx),
			Chance1:         d.GetNumber("Chance1", idx),
			ChancePerLevel1: d.GetNumber("ChancePerLvl1", idx),
			Level1:          d.GetNumber("Level1", idx),
			LvlPerLvl1:      d.GetNumber("LvlPerLvl1", idx),
			Skill2:          d.GetString("Skill2", idx),
			Mode2:           d.GetNumber("Mode2", idx),
			Chance2:         d.GetNumber("Chance2", idx),
			ChancePerLevel2: d.GetNumber("ChancePerLvl2", idx),
			Level2:          d.GetNumber("Level2", idx),
			LvlPerLvl2:      d.GetNumber("LvlPerLvl2", idx),
			Skill3:          d.GetString("Skill3", idx),
			Mode3:           d.GetNumber("Mode3", idx),
			Chance3:         d.GetNumber("Chance3", idx),
			ChancePerLevel3: d.GetNumber("ChancePerLvl3", idx),
			Level3:          d.GetNumber("Level3", idx),
			LvlPerLvl3:      d.GetNumber("LvlPerLvl3", idx),
			Skill4:          d.GetString("Skill4", idx),
			Mode4:           d.GetNumber("Mode4", idx),
			Chance4:         d.GetNumber("Chance4", idx),
			ChancePerLevel4: d.GetNumber("ChancePerLvl4", idx),
			Level4:          d.GetNumber("Level4", idx),
			LvlPerLvl4:      d.GetNumber("LvlPerLvl4", idx),
			Skill5:          d.GetString("Skill5", idx),
			Mode5:           d.GetNumber("Mode5", idx),
			Chance5:         d.GetNumber("Chance5", idx),
			ChancePerLevel5: d.GetNumber("ChancePerLvl5", idx),
			Level5:          d.GetNumber("Level5", idx),
			LvlPerLvl5:      d.GetNumber("LvlPerLvl5", idx),
			Skill6:          d.GetString("Skill6", idx),
			Mode6:           d.GetNumber("Mode6", idx),
			Chance6:         d.GetNumber("Chance6", idx),
			ChancePerLevel6: d.GetNumber("ChancePerLvl6", idx),
			Level6:          d.GetNumber("Level6", idx),
			LvlPerLvl6:      d.GetNumber("LvlPerLvl6", idx),
			Head:            d.GetNumber("Head", idx),
			Torso:           d.GetNumber("Torso", idx),
			Weapon:          d.GetNumber("Weapon", idx),
			Shield:          d.GetNumber("Shield", idx),
		}
		Hirelings = append(Hirelings, hireling)
	}

	log.Printf("Loaded %d Hireling records", len(Hirelings))
}
