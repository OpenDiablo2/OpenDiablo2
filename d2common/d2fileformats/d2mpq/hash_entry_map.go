package d2mpq

type HashEntryMap struct {
	entries map[uint64]*HashTableEntry
}

func (hem *HashEntryMap) Insert(entry *HashTableEntry) {
	if hem.entries == nil {
		hem.entries = make(map[uint64]*HashTableEntry)
	}

	hem.entries[uint64(entry.NamePartA)<<32|uint64(entry.NamePartB)] = entry
}

func (hem *HashEntryMap) Find(fileName string) (*HashTableEntry, bool) {
	if hem.entries == nil {
		return nil, false
	}

	hashA := hashString(fileName, 1)
	hashB := hashString(fileName, 2)

	entry, found := hem.entries[uint64(hashA)<<32|uint64(hashB)]
	return entry, found
}

func (hem *HashEntryMap) Contains(fileName string) bool {
	_, found := hem.Find(fileName)
	return found
}
