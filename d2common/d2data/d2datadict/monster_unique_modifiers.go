package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

const (
	numModifierConstants = 34
)

// MonsterUniqueModifierRecord represents a single line in monumod.txt
// Information gathered from [https://d2mods.info/forum/kb/viewarticle?a=161]
type MonsterUniqueModifierRecord struct {
	// Name of modifer, not used by other files
	Name string

	// ID of the modifier,
	// the Mod fields of SuperUniqueRecord refer to these ID's
	ID int

	// Enabled boolean for whether this modifier can be applied
	Enabled bool

	// ExpansionOnly boolean for whether this modifier can only be applied in an expansion game.
	// In the file, the value 100 represents expansion only
	ExpansionOnly bool

	// If true, "Minion" will be displayed below the life bar of minions of
	// the monster with this modifier
	Xfer bool

	// Champion boolean, only usable by champion monsters
	Champion bool

	// FPick Unknown
	FPick int

	// Exclude1 monster type code that cannot have this modifier
	Exclude1 string

	// Exclude2 monster type code that cannot have this modifier
	Exclude2 string

	PickFrequencies struct {
		Normal    *pickFreq
		Nightmare *pickFreq
		Hell      *pickFreq
	}
}

type pickFreq struct {
	// Champion pick frequency
	Champion int

	// Unique pick frequency
	Unique int
}

// MonsterUniqueModifiers stores the MonsterUniqueModifierRecords
var MonsterUniqueModifiers map[string]*MonsterUniqueModifierRecord //nolint:gochecknoglobals // Currently global by design

// MonsterUniqueModifierConstants contains constants from monumod.txt,
// can be accessed with indices from d2enum.MonUModConstIndex
var MonsterUniqueModifierConstants []int //nolint:gochecknoglobals // currently global by design

// See [https://d2mods.info/forum/kb/viewarticle?a=161] for more info

// LoadMonsterUniqueModifiers loads MonsterUniqueModifierRecords into MonsterUniqueModifiers
func LoadMonsterUniqueModifiers(file []byte) {
	MonsterUniqueModifiers = make(map[string]*MonsterUniqueModifierRecord)
	MonsterUniqueModifierConstants = make([]int, 0, numModifierConstants)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &MonsterUniqueModifierRecord{
			Name:          d.String("uniquemod"),
			ID:            d.Number("id"),
			Enabled:       d.Bool("enabled"),
			ExpansionOnly: d.Number("version") == expansionCode,
			Xfer:          d.Bool("xfer"),
			Champion:      d.Bool("champion"),
			FPick:         d.Number("fpick"),
			Exclude1:      d.String("exclude1"),
			Exclude2:      d.String("exclude2"),
			PickFrequencies: struct {
				Normal    *pickFreq
				Nightmare *pickFreq
				Hell      *pickFreq
			}{
				Normal: &pickFreq{
					Champion: d.Number("cpick"),
					Unique:   d.Number("upick"),
				},
				Nightmare: &pickFreq{
					Champion: d.Number("cpick (N)"),
					Unique:   d.Number("upick (N)"),
				},
				Hell: &pickFreq{
					Champion: d.Number("cpick (H)"),
					Unique:   d.Number("upick (H)"),
				},
			},
		}

		MonsterUniqueModifiers[record.Name] = record

		if len(MonsterUniqueModifierConstants) < numModifierConstants {
			MonsterUniqueModifierConstants = append(MonsterUniqueModifierConstants, d.Number("constants"))
		}
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonsterUniqueModifier records", len(MonsterUniqueModifiers))
}
