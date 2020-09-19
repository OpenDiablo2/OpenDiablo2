package d2records

// ElemTypes stores the ElemTypeRecords
type ElemTypes map[string]*ElemTypeRecord

// ElemTypeRecord represents a single line in ElemType.txt
type ElemTypeRecord struct {
	// ElemType Elemental damage type name
	ElemType string

	// Code Elemental damage type code
	Code string
}
