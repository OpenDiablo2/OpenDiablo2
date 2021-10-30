package d2records

// Sets contain the set records from sets.txt
type Sets map[string]*SetRecord

// SetRecord describes the set bonus for a group of set items
type SetRecord struct {
	Key            string
	StringTableKey string
	Properties     struct {
		PartialA []*SetProperty
		PartialB []*SetProperty
		Full     []*SetProperty
	}
	Version int
	Level   int
}

// SetProperty represents a property possessed by the set
type SetProperty = PropertyDescriptor
