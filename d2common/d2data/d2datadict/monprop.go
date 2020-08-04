package d2datadict

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const (
	numMonProps  = 6
	fmtProp      = "prop%d%s"
	fmtChance    = "chance%d%s"
	fmtPar       = "par%d%s"
	fmtMin       = "min%d%s"
	fmtMax       = "max%d%s"
	fmtNormal    = ""
	fmtNightmare = " (N)"
	fmtHell      = " (H)"
)

// MonPropRecord is a representation of a single row of monprop.txt
type MonPropRecord struct {
	ID string

	Properties struct {
		Normal    [numMonProps]*monProp
		Nightmare [numMonProps]*monProp
		Hell      [numMonProps]*monProp
	}
}

type monProp struct {
	Code   string
	Param  string
	Chance int
	Min    int
	Max    int
}

// MonProps stores all of the MonPropRecords
var MonProps map[string]*MonPropRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadMonProps loads monster property records into a map[string]*MonPropRecord
func LoadMonProps(file []byte) {
	MonProps = make(map[string]*MonPropRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &MonPropRecord{
			ID: d.String("Id"),

			Properties: struct {
				Normal    [numMonProps]*monProp
				Nightmare [numMonProps]*monProp
				Hell      [numMonProps]*monProp
			}{
				[numMonProps]*monProp{},
				[numMonProps]*monProp{},
				[numMonProps]*monProp{},
			},
		}

		for idx := 1; idx <= numMonProps; idx++ {
			record.Properties.Normal[idx-1] = &monProp{
				Code:   d.String(fmt.Sprintf(fmtProp, idx, fmtNormal)),
				Param:  d.String(fmt.Sprintf(fmtPar, idx, fmtNormal)),
				Chance: d.Number(fmt.Sprintf(fmtChance, idx, fmtNormal)),
				Min:    d.Number(fmt.Sprintf(fmtMin, idx, fmtNormal)),
				Max:    d.Number(fmt.Sprintf(fmtMax, idx, fmtNormal)),
			}

			record.Properties.Nightmare[idx-1] = &monProp{
				Code:   d.String(fmt.Sprintf(fmtProp, idx, fmtNightmare)),
				Param:  d.String(fmt.Sprintf(fmtPar, idx, fmtNightmare)),
				Chance: d.Number(fmt.Sprintf(fmtChance, idx, fmtNightmare)),
				Min:    d.Number(fmt.Sprintf(fmtMin, idx, fmtNightmare)),
				Max:    d.Number(fmt.Sprintf(fmtMax, idx, fmtNightmare)),
			}

			record.Properties.Hell[idx-1] = &monProp{
				Code:   d.String(fmt.Sprintf(fmtProp, idx, fmtHell)),
				Param:  d.String(fmt.Sprintf(fmtPar, idx, fmtHell)),
				Chance: d.Number(fmt.Sprintf(fmtChance, idx, fmtHell)),
				Min:    d.Number(fmt.Sprintf(fmtMin, idx, fmtHell)),
				Max:    d.Number(fmt.Sprintf(fmtMax, idx, fmtHell)),
			}
		}

		MonProps[record.ID] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonProp records", len(MonProps))
}
