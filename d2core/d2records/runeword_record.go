package d2records

// Runewords stores all of the RunesRecords
type Runewords map[string]*RunesRecord

// RunesRecord is a representation of a single row of runes.txt. It defines
// runewords available in the game.
type RunesRecord struct {
	Name     string
	RuneName string // More of a note - the actual name should be read from the TBL files.
	Complete bool   // An enabled/disabled flag. Only "Complete" runewords work in game.
	Server   bool   // Marks a runeword as only available on ladder, not single player or tcp/ip.

	// The item types for includsion/exclusion for this runeword record
	ItemTypes struct {
		Include []string
		Exclude []string
	}

	// Runes slice of ID pointers from Misc.txt, controls what runes are
	// required to make the rune word and in what order they are to be socketed.
	Runes []string

	Properties []*RunewordProperty
}

type RunewordProperty struct {
	// Code is the property code
	Code string

	// Param is either string or int, parameter for the property
	Param string

	// Min is the minimum value for the property
	Min int

	// Max is the maximum value for the property
	Max int
}
