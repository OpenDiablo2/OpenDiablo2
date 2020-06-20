package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const (
	expansion = "Expansion" // blizzard put this in the txt where expansion data starts
)

//nolint:gochecknoglobals // Currently global by design, only written once
var frameFields = []string{"Cel1", "Cel2", "Cel3", "Cel4"}

// AutoMapRecord represents one row from d2data.mpq/AutoMap.txt.
// Based on the information here https://d2mods.info/forum/kb/viewarticle?a=419
type AutoMapRecord struct {
	// LevelName is a string with an act number followed
	// by a level type, separated by a space. For example:
	// '1 Barracks' is the barracks level in act 1.
	LevelName string

	// TileName refers to a certain tile orientation.
	// See https://d2mods.info/forum/kb/viewarticle?a=468
	TileName string

	// Style is the top index in a 2D tile array.
	Style int // tiles[autoMapRecord.Style][]

	// StartSequence and EndSequence are sub indices the
	// same 2D array as Style. They describe a range of
	// tiles for which covered by this AutoMapRecord.
	// In some rows you can find a value of -1. This means
	// the game will only look at Style and TileName to
	// determine which tiles are addressed.
	StartSequence int // tiles[][autoMapRecord.StartSequence]
	EndSequence   int // tiles[][autoMapRecord.EndSequence]

	// Type values are described as:
	// "...just comment fields, as far as I know. Put in
	// whatever you like..."
	// The values seem functional but naming conventions
	// vary between LevelNames.
	// Type1 string
	// Type2 string
	// Type3 string
	// Type4 string // Note: I commented these out for now because they supposedly aren't useful see the LoadAutoMaps function.

	// Frames determine the frame of the MaxiMap(s).dc6 that
	// will be applied to the specified tiles. The frames
	// are in rows, if each row holds 20 images (when you
	// re-extract the chart with Dc6Table, you can specify
	// how many graphics a line can hold), line 1 includes
	// icons 0-19, line 2 from 20 to 39 etc.
	// Multiple values exist for Cel (and Type) to enable
	// variation. Presumably game chooses randomly between
	// any of the 4 values which are not set to -1.
	Frames []int
}

// AutoMaps contains all data in AutoMap.txt.
//nolint:gochecknoglobals // Current design is to have these global
var AutoMaps []*AutoMapRecord

// LoadAutoMaps populates AutoMaps with the data from AutoMap.txt.
// It also amends a duplicate field (column) name in that data.
func LoadAutoMaps(file []byte) {
	// Fix the error in the original file
	fileString := fixDuplicateFieldName(string(file))

	// Split file by newlines and tabs
	d := d2common.LoadDataDictionary(fileString)

	// Construct records
	AutoMaps = make([]*AutoMapRecord, len(d.Data))

	for idx := range d.Data {
		if d.GetString("LevelName", idx) == expansion {
			continue
		}

		AutoMaps[idx] = &AutoMapRecord{
			LevelName: d.GetString("LevelName", idx),
			TileName:  d.GetString("TileName", idx),

			Style:         d.GetNumber("Style", idx),
			StartSequence: d.GetNumber("StartSequence", idx),
			EndSequence:   d.GetNumber("EndSequence", idx),

			//Type1: d.GetString("Type1", idx),
			//Type2: d.GetString("Type2", idx),
			//Type3: d.GetString("Type3", idx),
			//Type4: d.GetString("Type4", idx),
			// Note: I commented these out for now because they supposedly
			// aren't useful see the AutoMapRecord struct.
		}

		AutoMaps[idx].Frames = make([]int, len(frameFields))
		for i := range frameFields {
			AutoMaps[idx].Frames[i] = d.GetNumber(frameFields[i], idx)
		}
	}

	log.Printf("Loaded %d AutoMapRecord records", len(AutoMaps))
}

// fixDuplicateFieldName changes one of the two 'Type2' fields
// in AutoMap.txt to 'Type3'. An error in the file can be seen
// by looking at the lists of 'Type' and 'Cel' fields:
//
//	Type1	Type2	Type2*	Type4
// 	Cel1	Cel2	Cel3	Cel4
//
// LoadDataDictionary uses a set of field names. The duplicate
// is omitted resulting in all rows being skipped because their
// counts are different from the field names count.
func fixDuplicateFieldName(fileString string) string {
	// Split rows
	rows := strings.Split(fileString, "\r\n")

	// Split the field names row and correct the duplicate
	fieldNames := strings.Split(rows[0], "\t")
	fieldNames[9] = "Type3"

	// Join the field names back up and assign to the first row
	rows[0] = strings.Join(fieldNames, "\t")

	// Return the rows, joined back into one string
	return strings.Join(rows, "\r\n")
}
