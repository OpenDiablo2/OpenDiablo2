package d2records

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func monsterPropertiesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonsterProperties)

	for d.Next() {
		record := &MonPropRecord{
			ID: d.String("Id"),

			Properties: struct {
				Normal    [NumMonProps]*MonProp
				Nightmare [NumMonProps]*MonProp
				Hell      [NumMonProps]*MonProp
			}{
				[NumMonProps]*MonProp{},
				[NumMonProps]*MonProp{},
				[NumMonProps]*MonProp{},
			},
		}

		for idx := 1; idx <= NumMonProps; idx++ {
			record.Properties.Normal[idx-1] = &MonProp{
				Code:   d.String(fmt.Sprintf(FmtProp, idx, FmtNormal)),
				Param:  d.String(fmt.Sprintf(FmtPar, idx, FmtNormal)),
				Chance: d.Number(fmt.Sprintf(FmtChance, idx, FmtNormal)),
				Min:    d.Number(fmt.Sprintf(FmtMin, idx, FmtNormal)),
				Max:    d.Number(fmt.Sprintf(FmtMax, idx, FmtNormal)),
			}

			record.Properties.Nightmare[idx-1] = &MonProp{
				Code:   d.String(fmt.Sprintf(FmtProp, idx, FmtNightmare)),
				Param:  d.String(fmt.Sprintf(FmtPar, idx, FmtNightmare)),
				Chance: d.Number(fmt.Sprintf(FmtChance, idx, FmtNightmare)),
				Min:    d.Number(fmt.Sprintf(FmtMin, idx, FmtNightmare)),
				Max:    d.Number(fmt.Sprintf(FmtMax, idx, FmtNightmare)),
			}

			record.Properties.Hell[idx-1] = &MonProp{
				Code:   d.String(fmt.Sprintf(FmtProp, idx, FmtHell)),
				Param:  d.String(fmt.Sprintf(FmtPar, idx, FmtHell)),
				Chance: d.Number(fmt.Sprintf(FmtChance, idx, FmtHell)),
				Min:    d.Number(fmt.Sprintf(FmtMin, idx, FmtHell)),
				Max:    d.Number(fmt.Sprintf(FmtMax, idx, FmtHell)),
			}
		}

		records[record.ID] = record
	}

	if d.Err != nil {
		return d.Err
	}

	log.Printf("Loaded %d MonProp records", len(records))

	r.Monster.Props = records

	return nil
}
