package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

//MonsterUniqueModifierRecord represents a single line in monumod.txt
//Information gathered from [https://d2mods.info/forum/kb/viewarticle?a=161]
type MonsterUniqueModifierRecord struct {
	//Name of modifer, not used by other files
	Name string

	//ID of the modifier
	//The Mod fields of SuperUniqueRecord refer to these ID's
	ID int

	//If true, this modifier can be applied
	Enabled bool

	//If true, this modifier can only be applied in an expansion game
	//In the file, the value 100 represents expansion only
	ExpansionOnly bool

	//If true, "Minion" will be displayed below the life bar of minions of the monster with this modifier
	Xfer bool

	//If true, only usable by champion monsters
	Champion bool

	//Unknown
	FPick int

	//Names of monster types that cannot have this modifier
	//Refer to Type in MonType.txt
	Exclude1 string
	Exclude2 string

	//Determines the frequency this modifier appears on champion monsters
	//If empty, it will not appear on champions
	ChampionPickFrequency          int
	ChampionPickFrequencyNightmare int
	ChampionPickFrequencyHell      int

	//Determines the frequency this modifier appears on random unique monsters
	//If empty, it will not appear on random unique monsters
	UniquePickFrequency          int
	UniquePickFrequencyNightmare int
	UniquePickFrequencyHell      int

	//FInit int: Unused
}

//MonsterUniqueModifiers stores the MonsterUniqueModifierRecords
var MonsterUniqueModifiers map[string]*MonsterUniqueModifierRecord //nolint:gochecknoglobals // Currently global by design

//MonsterUniqueModifierConstants contains constants from monumod.txt
//See [https://d2mods.info/forum/kb/viewarticle?a=161] for more info
//Can be accessed with indices from d2enum.MonUModConstIndex
var MonsterUniqueModifierConstants []int //nolint:gochecknoglobals

//LoadMonsterUniqueModifiers loads MonsterUniqueModifierRecords into MonsterUniqueModifiers
func LoadMonsterUniqueModifiers(file []byte) {
	MonsterUniqueModifiers = make(map[string]*MonsterUniqueModifierRecord)
	MonsterUniqueModifierConstants = make([]int, 0, 34)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &MonsterUniqueModifierRecord{
			Name:                           d.String("uniquemod"),
			ID:                             d.Number("id"),
			Enabled:                        d.Bool("enabled"),
			ExpansionOnly:                  d.Number("version") == 100,
			Xfer:                           d.Bool("xfer"),
			Champion:                       d.Bool("champion"),
			FPick:                          d.Number("fpick"),
			Exclude1:                       d.String("exclude1"),
			Exclude2:                       d.String("exclude2"),
			ChampionPickFrequency:          d.Number("cpick"),
			ChampionPickFrequencyNightmare: d.Number("cpick (N)"),
			ChampionPickFrequencyHell:      d.Number("cpick (H)"),
			UniquePickFrequency:            d.Number("upick"),
			UniquePickFrequencyNightmare:   d.Number("upick (N)"),
			UniquePickFrequencyHell:        d.Number("upick (H)"),
		}
		MonsterUniqueModifiers[record.Name] = record
		if len(MonsterUniqueModifierConstants) < 34 {
			MonsterUniqueModifierConstants = append(MonsterUniqueModifierConstants, (d.Number("constants")))
		}
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonsterUniqueModifier records", len(MonsterUniqueModifiers))
}
