package d2records

// BodyLocations contains the body location records
type BodyLocations map[string]*BodyLocationRecord

// BodyLocationRecord describes a body location that items can be equipped to
type BodyLocationRecord struct {
	Name string
	Code string
}
