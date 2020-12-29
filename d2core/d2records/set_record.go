package d2records

// Sets contain the set records from sets.txt
type Sets map[string]*SetRecord

// SetRecord describes the set bonus for a group of set items
type SetRecord struct {
	// index
	// String key linked to by the set field in SetItems.
	// txt - used to tie all of the set's items to the same set.
	Key string

	// name
	// String key to item's name in a .tbl file.
	StringTableKey string

	// Version 0 for vanilla, 100 for LoD expansion
	Version int

	// Level
	// set level, perhaps intended as a minimum level for partial or full attributes to appear
	// (reference only, not loaded into game).
	Level int

	// Properties contains the partial and full set bonus properties.
	Properties struct {
		PartialA []*SetProperty
		PartialB []*SetProperty
		Full     []*SetProperty
	}
}

// SetProperty represents a property possessed by the set
type SetProperty = PropertyDescriptor
