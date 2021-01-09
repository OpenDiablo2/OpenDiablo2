package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func monsterSuperUniqeLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(SuperUniques)

	for d.Next() {
		record := &SuperUniqueRecord{
			Key:      d.String("Superunique"),
			Name:     d.String("Name"),
			Class:    d.String("Class"),
			HcIdx:    d.String("hcIdx"),
			MonSound: d.String("MonSound"),
			Mod: [3]int{
				d.Number("Mod1"),
				d.Number("Mod2"),
				d.Number("Mod3"),
			},
			MinGrp:                 d.Number("MinGrp"),
			MaxGrp:                 d.Number("MaxGrp"),
			IsExpansion:            d.Bool("EClass"),
			AutoPosition:           d.Bool("AutoPos"),
			Stacks:                 d.Bool("Stacks"),
			TreasureClassNormal:    d.String("TC"),
			TreasureClassNightmare: d.String("TC(N)"),
			TreasureClassHell:      d.String("TC(H)"),
			UTransNormal:           d.String("Utrans"),
			UTransNightmare:        d.String("Utrans(N)"),
			UTransHell:             d.String("Utrans(H)"),
		}
		records[record.Key] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Monster.Unique.Super = records

	r.Logger.Infof("Loaded %d SuperUnique records", len(records))

	return nil
}
