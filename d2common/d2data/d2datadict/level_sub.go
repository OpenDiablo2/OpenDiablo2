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
	LevelSubstitutions = make(map[int]*LevelSubstitutionRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &LevelSubstitutionRecord{
			Name:         d.String("Name"),
			ID:           d.Number("Type"),
			File:         d.String("File"),
			IsExpansion:  d.Number("Expansion") > 0,
			BorderType:   d.Number("BordType"),
			GridSize:     d.Number("GridSize"),
			Mask:         d.Number("Dt1Mask"),
			ChanceSpawn0: d.Number("Prob0"),
			ChanceSpawn1: d.Number("Prob1"),
			ChanceSpawn2: d.Number("Prob2"),
			ChanceSpawn3: d.Number("Prob3"),
			ChanceSpawn4: d.Number("Prob4"),
			ChanceFloor0: d.Number("Trials0"),
			ChanceFloor1: d.Number("Trials1"),
			ChanceFloor2: d.Number("Trials2"),
			ChanceFloor3: d.Number("Trials3"),
			ChanceFloor4: d.Number("Trials4"),
			GridMax0:     d.Number("Max0"),
			GridMax1:     d.Number("Max1"),
			GridMax2:     d.Number("Max2"),
			GridMax3:     d.Number("Max3"),
			GridMax4:     d.Number("Max4"),
		}
		LevelSubstitutions[record.ID] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d LevelSubstitution records", len(LevelSubstitutions))
}
