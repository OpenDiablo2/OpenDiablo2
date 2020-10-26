package d2records

// RarePrefixes is where all RareItemPrefixRecords are stored
type RarePrefixes []*RareItemPrefixRecord

// RareItemPrefixRecord is a name prefix for rare items (items with more than 2 affixes)
type RareItemPrefixRecord = RareItemAffix
