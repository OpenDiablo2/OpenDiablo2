package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// LevelSubstitutionRecord is a representation of a row from lvlsub.txt
// these records are parameters for levels and describe substitution rules
type LevelSubstitutionRecord struct {
	// Description, reference only.
	Name string // Name

	// This value is used in Levels.txt, in the column 'SubType'. You'll notice
	// that in LvlSub.txt some rows use the same value, we can say they forms
	// groups. If you count each row of a group starting from 0, then you'll
	// obtain what is written in Levels.txt, columns 'SubTheme', 'SubWaypoint'
	// and 'SubShrine'. (added by Paul Siramy)
	ID int // Type

	// What .ds1 is being used.
	File string // File

	// 0 for classic, 1 for Expansion.
	IsExpansion bool // Expansion

	// Unknown as all have 0.
	// CheckAll

	// this field can contain values ranging from -1 to 2
	// NOTE: wall types have 0, 1 or 2, while Non-wall types have -1.
	BorderType int // BordType

	// Set it to 1 or 2 I'm assuming this means a block of tiles ie: 4x4.
	GridSize int // GridSize

	// For some rows, this is their place in LvlTypes.txt. The Dt1 mask also
	// includes the mask for the Floor.Dt1 of that level. (see Trials0 below)
	Mask int // Dt1Mask

	// The probability of the Dt1 being spawned.
	ChanceSpawn0 int // Prob0
	ChanceSpawn1 int // Prob1
	ChanceSpawn2 int // Prob2
	ChanceSpawn3 int // Prob3
	ChanceSpawn4 int // Prob4

	// This appears to be a chance of either a floor tile being spawned or the
	// actual Dt1..
	ChanceFloor0 int // Trials0
	ChanceFloor1 int // Trials1
	ChanceFloor2 int // Trials2
	ChanceFloor3 int // Trials3
	ChanceFloor4 int // Trials4

	// This appears to be how much will spawn in the Grid.
	GridMax0 int // Max0
	GridMax1 int // Max1
	GridMax2 int // Max2
	GridMax3 int // Max3
	GridMax4 int // Max4

	// Beta
}

// LevelSubstitutions stores all of the LevelSubstitutionRecords
//nolint:gochecknoglobals // Currently global by design
var LevelSubstitutions map[int]*LevelSubstitutionRecord

// LoadLevelSubstitutions loads lvlsub.txt and parses into records
func LoadLevelSubstitutions(file []byte) {
	dict := d2common.LoadDataDictionary(string(file))
	numRecords := len(dict.Data)
	LevelSubstitutions = make(map[int]*LevelSubstitutionRecord, numRecords)

	for idx := range dict.Data {
		record := &LevelSubstitutionRecord{
			Name:         dict.GetString("Name", idx),
			ID:           dict.GetNumber("Type", idx),
			File:         dict.GetString("File", idx),
			IsExpansion:  dict.GetNumber("Expansion", idx) > 0,
			BorderType:   dict.GetNumber("BordType", idx),
			GridSize:     dict.GetNumber("GridSize", idx),
			Mask:         dict.GetNumber("Dt1Mask", idx),
			ChanceSpawn0: dict.GetNumber("Prob0", idx),
			ChanceSpawn1: dict.GetNumber("Prob1", idx),
			ChanceSpawn2: dict.GetNumber("Prob2", idx),
			ChanceSpawn3: dict.GetNumber("Prob3", idx),
			ChanceSpawn4: dict.GetNumber("Prob4", idx),
			ChanceFloor0: dict.GetNumber("Trials0", idx),
			ChanceFloor1: dict.GetNumber("Trials1", idx),
			ChanceFloor2: dict.GetNumber("Trials2", idx),
			ChanceFloor3: dict.GetNumber("Trials3", idx),
			ChanceFloor4: dict.GetNumber("Trials4", idx),
			GridMax0:     dict.GetNumber("Max0", idx),
			GridMax1:     dict.GetNumber("Max1", idx),
			GridMax2:     dict.GetNumber("Max2", idx),
			GridMax3:     dict.GetNumber("Max3", idx),
			GridMax4:     dict.GetNumber("Max4", idx),
		}

		LevelSubstitutions[record.ID] = record
	}

	log.Printf("Loaded %d LevelSubstitution records", len(LevelSubstitutions))
}
