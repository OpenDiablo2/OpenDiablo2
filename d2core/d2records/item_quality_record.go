package d2records

// ItemQualities stores all of the QualityRecords
type ItemQualities map[string]*ItemQualityRecord

// ItemQualityRecord represents a single row of ItemQualities.txt, which controls
// properties for superior quality items
type ItemQualityRecord struct {
	NumMods   int
	Mod1Code  string
	Mod1Param int
	Mod1Min   int
	Mod1Max   int
	Mod2Code  string
	Mod2Param int
	Mod2Min   int
	Mod2Max   int

	// The following fields determine this row's applicability to
	// categories of item.
	Armor   bool
	Weapon  bool
	Shield  bool
	Thrown  bool
	Scepter bool
	Wand    bool
	Staff   bool
	Bow     bool
	Boots   bool
	Gloves  bool
	Belt    bool

	Level    int
	Multiply int
	Add      int
}
