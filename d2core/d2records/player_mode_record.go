package d2records

// PlayerModes stores the PlayerModeRecords
type PlayerModes map[string]*PlayerModeRecord

// PlayerModeRecord represents a single line in PlayerMode.txt
type PlayerModeRecord struct {
	// Player animation mode name
	Name string

	// Player animation mode token
	Token string
}
