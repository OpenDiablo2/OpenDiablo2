package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// The skillcalc.txt and misscalc.txt files are essentially lookup tables for the Skills.txt and Missiles.txt Calc functions
// To avoid duplication (since they have identical fields) they are both represented by the CalculationRecord type
type CalculationRecord struct {
	Code        string
	Description string
}

var SkillCalculations map[string]*CalculationRecord
var MissileCalculations map[string]*CalculationRecord

func LoadSkillCalculations(file []byte) {
	SkillCalculations = make(map[string]*CalculationRecord)

	d := d2common.LoadDataDictionary(file)
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

func LoadMissileCalculations(file []byte) {
	MissileCalculations = make(map[string]*CalculationRecord)

	d := d2common.LoadDataDictionary(file)
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
