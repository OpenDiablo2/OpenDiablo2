package d2records

// Runewords stores all of the RuneRecords
type Runewords map[string]*RuneRecord

// RuneRecord is a representation of a single row of runes.txt. It defines
// runewords available in the game.
type RuneRecord struct {
	Name      string
	RuneName  string
	ItemTypes struct {
		Include []string
		Exclude []string
	}
	Runes      []string
	Properties []*RunewordProperty
	Complete   bool
	Server     bool
}

// RunewordProperty is a representation of a stat possessed by this runeword
type RunewordProperty = PropertyDescriptor
