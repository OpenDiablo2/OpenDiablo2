package d2records

// ObjectModes is a map of ObjectModeRecords
type ObjectModes map[string]*ObjectModeRecord

// ObjectModeRecord is a representation of an animation mode for an object
type ObjectModeRecord struct {
	Name  string
	Token string
}
