package d2records

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
type AutoMaps []*AutoMapRecord
