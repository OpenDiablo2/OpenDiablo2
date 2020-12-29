package d2records

// RareSuffixes is where all RareItemSuffixRecords are stored
type RareSuffixes []*RareItemSuffixRecord

// RareItemSuffixRecord is a name suffix for rare items (items with more than 2 affixes)
type RareItemSuffixRecord = RareItemAffix
