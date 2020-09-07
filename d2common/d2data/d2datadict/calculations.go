package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// CalculationRecord The skillcalc.txt and misscalc.txt files are essentially lookup tables
// for the Skills.txt and Missiles.txt Calc functions To avoid duplication (since they have
// identical fields) they are both represented by the CalculationRecord type
type CalculationRecord struct {
	Code        string
	Description string
}

// SkillCalculations is where calculation records are stored
var SkillCalculations map[string]*CalculationRecord //nolint:gochecknoglobals // Currently global by design

// MissileCalculations is where missile calculations are stored
var MissileCalculations map[string]*CalculationRecord //nolint:gochecknoglobals // Currently global by design

// LoadSkillCalculations loads skill calculation records from skillcalc.txt
func LoadSkillCalculations(file []byte) {
	SkillCalculations = make(map[string]*CalculationRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &CalculationRecord{
			Code:        d.String("code"),
			Description: d.String("*desc"),
		}
		SkillCalculations[record.Code] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Skill Calculation records", len(SkillCalculations))
}

// LoadMissileCalculations loads missile calculation records from misscalc.txt
func LoadMissileCalculations(file []byte) {
	MissileCalculations = make(map[string]*CalculationRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &CalculationRecord{
			Code:        d.String("code"),
			Description: d.String("*desc"),
		}
		MissileCalculations[record.Code] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Missile Calculation records", len(MissileCalculations))
}
