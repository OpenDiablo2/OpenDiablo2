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

type SetProperty struct {
	// Code is an ID pointer of a property from Properties.txt,
	// these columns control each of the eight different full set modifiers a set item can grant you
	// at most.
	Code string

	// Param is the passed on to the associated property, this is used to pass skill IDs, state IDs,
	// monster IDs, montype IDs and the like on to the properties that require them,
	// these fields support calculations.
	Param string

	// Min value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guides for Properties.txt and ItemStatCost.
	// txt for further details.
	Min int

	// Max value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guides for Properties.txt and ItemStatCost.
	// txt for further details.
	Max int
}
