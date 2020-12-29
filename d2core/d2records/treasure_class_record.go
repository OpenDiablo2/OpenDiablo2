package d2records

// TreasureClass contains all of the TreasureClassRecords
type TreasureClass map[string]*TreasureClassRecord

// TreasureClassRecord represents a rule for item drops in diablo 2
type TreasureClassRecord struct {
	Name       string
	Group      int
	Level      int
	NumPicks   int
	FreqUnique int
	FreqSet    int
	FreqRare   int
	FreqMagic  int
	FreqNoDrop int
	Treasures  []*Treasure
}

// Treasure describes a treasure to drop
// the Name is either a reference to an item, or to another treasure class
type Treasure struct {
	Code        string
	Probability int
}
