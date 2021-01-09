package d2records

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func monsterPropertiesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonsterProperties)

	for d.Next() {
		record := &MonPropRecord{
			ID: d.String("Id"),

			Properties: struct {
				Normal    [numMonProps]*MonProp
				Nightmare [numMonProps]*MonProp
				Hell      [numMonProps]*MonProp
			}{
				[numMonProps]*MonProp{},
				[numMonProps]*MonProp{},
				[numMonProps]*MonProp{},
			},
		}

		for idx := 1; idx <= numMonProps; idx++ {
			record.Properties.Normal[idx-1] = &MonProp{
				Code:   d.String(fmt.Sprintf(fmtProp, idx, fmtNormal)),
				Param:  d.String(fmt.Sprintf(fmtPar, idx, fmtNormal)),
				Chance: d.Number(fmt.Sprintf(fmtChance, idx, fmtNormal)),
				Min:    d.Number(fmt.Sprintf(fmtMin, idx, fmtNormal)),
				Max:    d.Number(fmt.Sprintf(fmtMax, idx, fmtNormal)),
			}

			record.Properties.Nightmare[idx-1] = &MonProp{
				Code:   d.String(fmt.Sprintf(fmtProp, idx, fmtNightmare)),
				Param:  d.String(fmt.Sprintf(fmtPar, idx, fmtNightmare)),
				Chance: d.Number(fmt.Sprintf(fmtChance, idx, fmtNightmare)),
				Min:    d.Number(fmt.Sprintf(fmtMin, idx, fmtNightmare)),
				Max:    d.Number(fmt.Sprintf(fmtMax, idx, fmtNightmare)),
			}

			record.Properties.Hell[idx-1] = &MonProp{
				Code:   d.String(fmt.Sprintf(fmtProp, idx, fmtHell)),
				Param:  d.String(fmt.Sprintf(fmtPar, idx, fmtHell)),
				Chance: d.Number(fmt.Sprintf(fmtChance, idx, fmtHell)),
				Min:    d.Number(fmt.Sprintf(fmtMin, idx, fmtHell)),
				Max:    d.Number(fmt.Sprintf(fmtMax, idx, fmtHell)),
			}
		}

		records[record.ID] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d MonProp records", len(records))

	r.Monster.Props = records

	return nil
}
