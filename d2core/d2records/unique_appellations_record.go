package d2records

// UniqueAppellations contains all of the UniqueAppellationRecords
type UniqueAppellations map[string]*UniqueAppellationRecord

// UniqueAppellationRecord described the extra suffix of a unique monster name
type UniqueAppellationRecord struct {
	// The title
	Name string
}
