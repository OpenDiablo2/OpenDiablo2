package d2records

// MonsterTypes stores all of the MonTypeRecords
type MonsterTypes map[string]*MonTypeRecord

// MonTypeRecord is a representation of a single row of MonType.txt.
type MonTypeRecord struct {
	Type   string
	Equiv1 string
	Equiv2 string
	Equiv3 string
	// StrSing is the string displayed for the singular form (Skeleton), note
	// that this is unused in the original engine, since the only modifier
	// display code that accesses MonType uses StrPlur.
	StrSing   string
	StrPlural string
}
