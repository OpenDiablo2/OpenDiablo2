package d2mpq

// HashEntryMap represents a hash entry map
type HashEntryMap struct {
	entries map[uint64]HashTableEntry
}

// Insert inserts a hash entry into the table
func (hem *HashEntryMap) Insert(entry *HashTableEntry) {
	if hem.entries == nil {
		hem.entries = make(map[uint64]HashTableEntry)
	}

	hem.entries[uint64(entry.NamePartA)<<32|uint64(entry.NamePartB)] = *entry
}

// Find finds a hash entry
func (hem *HashEntryMap) Find(fileName string) (*HashTableEntry, bool) {
	if hem.entries == nil {
		return nil, false
	}

	hashA := hashString(fileName, 1)
	hashB := hashString(fileName, 2)

	entry, found := hem.entries[uint64(hashA)<<32|uint64(hashB)]

	return &entry, found
}

// Contains returns true if the hash entry contains the values
func (hem *HashEntryMap) Contains(fileName string) bool {
	_, found := hem.Find(fileName)
	return found
}
