package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

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
	AutoMaps = make([]*AutoMapRecord, 0)

	var frameFields = []string{"Cel1", "Cel2", "Cel3", "Cel4"}

	// Split file by newlines and tabs
	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &AutoMapRecord{
			LevelName: d.String("LevelName"),
			TileName:  d.String("TileName"),

			Style:         d.Number("Style"),
			StartSequence: d.Number("StartSequence"),
			EndSequence:   d.Number("EndSequence"),

			//Type1: d.String("Type1"),
			//Type2: d.String("Type2"),
			//Type3: d.String("Type3"),
			//Type4: d.String("Type4"),
			// Note: I commented these out for now because they supposedly
			// aren't useful see the AutoMapRecord struct.
		}
		record.Frames = make([]int, len(frameFields))

		for i := range frameFields {
			record.Frames[i] = d.Number(frameFields[i])
		}

		AutoMaps = append(AutoMaps, record)
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d AutoMapRecord records", len(AutoMaps))
}
