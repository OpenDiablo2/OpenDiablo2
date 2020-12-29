package d2records

// ItemRatios holds all of the ItemRatioRecords from ItemRatio.txt
type ItemRatios map[string]*ItemRatioRecord

// DropRatioInfo is a helper struct for item drop calculation
type DropRatioInfo struct {
	Frequency  int
	Divisor    int
	DivisorMin int
}

// ItemRatioRecord encapsulates information found in ItemRatio.txt, it specifies drop ratios
// for various types of items
// The information has been gathered from [https://d2mods.info/forum/kb/viewarticle?a=387]
type ItemRatioRecord struct {
	Function string
	// 0 for classic, 1 for LoD
	Version bool

	// 0 for normal, 1 for exceptional
	Uber          bool
	ClassSpecific bool

	// All following fields are used in item drop calculation
	UniqueDropInfo    DropRatioInfo
	RareDropInfo      DropRatioInfo
	SetDropInfo       DropRatioInfo
	MagicDropInfo     DropRatioInfo
	HiQualityDropInfo DropRatioInfo
	NormalDropInfo    DropRatioInfo
}
