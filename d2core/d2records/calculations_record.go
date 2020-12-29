package d2records

// Calculations is where calculation records are stored
type Calculations map[string]*CalculationRecord

// CalculationRecord The skillcalc.txt and misscalc.txt files are essentially lookup tables
// for the Skills.txt and Missiles.txt Calc functions To avoid duplication (since they have
// identical fields) they are both represented by the CalculationRecord type
type CalculationRecord struct {
	Code        string
	Description string
}
