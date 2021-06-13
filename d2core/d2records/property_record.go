package d2records

// Properties stores all of the PropertyRecords
type Properties map[string]*PropertyRecord

// PropertyStatRecord contains stat information for a property
type PropertyStatRecord struct {
	StatCode   string
	SetID      int
	Value      int
	FunctionID int
}

// PropertyRecord is a representation of a single row of properties.txt
type PropertyRecord struct {
	Stats  [7]*PropertyStatRecord
	Code   string
	Active string
}
