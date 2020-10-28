package d2records

// PlayerTypes is a map of PlayerTypeRecords
type PlayerTypes map[string]*PlayerTypeRecord

// PlayerTypeRecord is used for changing character animation modes.
type PlayerTypeRecord struct {
	Name  string
	Token string
}
