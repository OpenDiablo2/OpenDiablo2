package d2records

// ComponentCodes is a lookup table for DCC Animation ComponentID Subtype,
// it links hardcoded data with the txt files
type ComponentCodes map[string]*ComponentCodeRecord

// ComponentCodeRecord represents a single row from compcode.txt
type ComponentCodeRecord struct {
	Component string
	Code      string
}
