package d2records

// CompositeTypes is a map of CompositeTypeRecords
type CompositeTypes map[string]*CompositeTypeRecord

// CompositeTypeRecord describes a layer for an animated composite (multi-sprite entities).
// The token is used for changing character animation modes.
type CompositeTypeRecord struct {
	Name  string
	Token string
}
