package d2records

// LowQualities is a slice of LowQualityRecords
type LowQualities []*LowQualityRecord

// LowQualityRecord is a name prefix that can be used for low quality item names
type LowQualityRecord struct {
	Name string
}
