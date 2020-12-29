package d2records

// Gamble is a map of GambleRecords
type Gamble map[string]*GambleRecord

// GambleRecord is a representation of an item type that can be gambled for at vendors
type GambleRecord struct {
	Name string
	Code string
}
