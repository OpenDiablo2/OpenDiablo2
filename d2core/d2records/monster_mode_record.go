package d2records

// MonModes stores all of the GemsRecords
type MonModes map[string]*MonModeRecord

// MonModeRecord is a representation of a single row of Monmode.txt
type MonModeRecord struct {
	Name  string
	Token string
	Code  string
}
