package d2records

// LevelPresets stores all of the LevelPresetRecords
type LevelPresets map[int]LevelPresetRecord

// LevelPresetRecord is a representation of a row from lvlprest.txt
// these records define parameters for the preset level map generator
type LevelPresetRecord struct {
	Files        [6]string
	Name         string
	DefinitionID int
	LevelID      int
	SizeX        int
	SizeY        int
	Pops         int
	PopPad       int
	FileCount    int
	Dt1Mask      uint
	Populate     bool
	Logicals     bool
	Outdoors     bool
	Animate      bool
	KillEdge     bool
	FillBlanks   bool
	AutoMap      bool
	Scan         bool
	Beta         bool
	Expansion    bool
}
