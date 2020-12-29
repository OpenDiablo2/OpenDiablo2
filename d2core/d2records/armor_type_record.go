package d2records

// ArmorTypes is a map of ArmorTypeRecords
type ArmorTypes map[string]*ArmorTypeRecord

// ArmorTypeRecord describes an armor type. It has a name and 3-character token.
// The token is used to change the character animation mode.
type ArmorTypeRecord struct {
	Name  string
	Token string
}
