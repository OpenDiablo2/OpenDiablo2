package d2records

// ItemQualities stores all of the QualityRecords
type ItemQualities map[string]*ItemQualityRecord

// ItemQualityRecord represents a single row of ItemQualities.txt, which controls
// properties for superior quality items
type ItemQualityRecord struct {
	Mod1Code  string
	Mod2Code  string
	NumMods   int
	Mod1Param int
	Mod1Min   int
	Mod1Max   int
	Mod2Param int
	Mod2Min   int
	Mod2Max   int
	Multiply  int
	Level     int
	Add       int
	Thrown    bool
	Scepter   bool
	Wand      bool
	Staff     bool
	Bow       bool
	Boots     bool
	Gloves    bool
	Belt      bool
	Weapon    bool
	Armor     bool
	Shield    bool
}
