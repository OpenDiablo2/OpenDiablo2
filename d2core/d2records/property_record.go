package d2records

// Properties stores all of the PropertyRecords
type Properties map[string]*PropertyRecord

// PropertyStatRecord contains stat information for a property
type PropertyStatRecord struct {
	SetID      int
	Value      int
	FunctionID int
	StatCode   string
}

// PropertyRecord is a representation of a single row of properties.txt
type PropertyRecord struct {
	Code   string
	Active string
	Stats  [7]*PropertyStatRecord
}
