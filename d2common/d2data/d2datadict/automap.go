package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// AutoMapRecord represents one row from d2data.mpq/AutoMap.txt.
type AutoMapRecord struct {
	// LevelName is a string with an act number followed by a level
	// type, separated by a space.
	//
	// For example: '1 Barracks' is the barracks level in act 1.
	//
	// The following list shows level names by their act:
	// Act 1: 	Barracks Catacombs Cathedral Cave Courtyard Crypt
	// 		Jail Monestary Town Tristram Wilderness
	// Act 2: 	Arcane Basement Desert Harem Lair Sewer Tomb Town
	// Act 3: 	Dungeon Jungle Kurast Sewer Spider Town
	// Act 4: 	Lava Mesa Town
	// Act 5: 	Baal Barricade Ice Lava Siege Temple
	LevelName string

	//TileName
	//Style
	//StartSequence
	//EndSequence
	//Type1
	//Cel1
	//Type2
	//Cel2
	//Type3
	//Cel3
	//Type4
	//Cel4
}

// AutoMaps contains all data in AutoMap.txt.
var AutoMaps []*AutoMapRecord

// LoadAutoMaps populates AutoMaps with the data from AutoMap.txt.
// It also amends a duplicate field name in the AutoMap.txt file.
func LoadAutoMaps(file []byte) {
	// Fix the error in the original file
	fileString := fixDuplicateFieldName(string(file))

	// Load data
	d := d2common.LoadDataDictionary(fileString)

	// Populate slice items
	AutoMaps = make([]*AutoMapRecord, len(d.Data))
	for idx := range d.Data {
		// This file contains a line with empty field values, used
		// as a visual separator.
		//
		// TODO: Should that be handled in LoadDataDictionary since it also
		// needs handling here: https://github.com/OpenDiablo2/OpenDiablo2/pull/347/commits/2413aada25b1e23d742f9e966a77648626b69138
		// ..and possibly elsewhere.
		if idx == 2603 {
			continue
		}

		AutoMaps[idx] = &AutoMapRecord{
			LevelName: d.GetString("LevelName", idx),
			//TileName:
			//Style:
			//StartSequence:
			//EndSequence:
			//Type1:
			//Cel1:
			//Type2:
			//Cel2:
			//Type2:
			//Cel3:
			//Type4:
			//Cel4:

		}
	}

	log.Printf( /*"Loaded %d AutoMapRecord records"*/ "LoadAutoMaps ran - %d", len(AutoMaps))
}

// fixDuplicateFieldName changes one of the two 'Type2'
// fields in AutoMap.txt to 'Type3'. An error in the file
// is immediately obvious looking at the lists of 'Type'
// and 'Cel' fields:
//
//	Type1	Type2	Type2*	Type4
// 	Cel1	Cel2	Cel3	Cel4
//
// The duplicate trips up d2common.LoadDataDictionary which
// creates a set of unique field name strings and ignores
// any row with a different count to that set.
// The field name count is 12, due to the duplicate, and the
// row value counts are all 13 resulting in empty arrays.
func fixDuplicateFieldName(fileString string) string {
	// Split lines
	lines := strings.Split(fileString, "\r\n")

	// Split the field names line correct the duplicate Type2 field to Type3
	fieldNames := strings.Split(lines[0], "\t")
	fieldNames[9] = "Type3"

	// Join the field names back up and assign to the first line
	lines[0] = strings.Join(fieldNames, "\t")

	// Return the lines, joined back into a single string
	return strings.Join(lines, "\r\n")
}
