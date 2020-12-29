package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func skillCalcLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := loadCalculations(r, d, "Skill")
	if err != nil {
		return err
	}

	r.Calculation.Skills = records

	return nil
}

func missileCalcLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := loadCalculations(r, d, "Missile")
	if err != nil {
		return err
	}

	r.Calculation.Missiles = records

	return nil
}

func loadCalculations(r *RecordManager, d *d2txt.DataDictionary, name string) (Calculations, error) {
	records := make(Calculations)

	for d.Next() {
		record := &CalculationRecord{
			Code:        d.String("code"),
			Description: d.String("*desc"),
		}
		records[record.Code] = record
	}

	if d.Err != nil {
		return nil, d.Err
	}

	r.Debugf("Loaded %d %s Calculation records", len(records), name)

	return records, nil
}
