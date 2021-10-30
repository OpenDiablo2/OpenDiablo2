package d2records

// AutoMapRecord represents one row from d2data.mpq/AutoMap.txt.
// Based on the information here https://d2mods.info/forum/kb/viewarticle?a=419
type AutoMapRecord struct {
	LevelName     string
	TileName      string
	Frames        []int
	Style         int
	StartSequence int
	EndSequence   int
}

// AutoMaps contains all data in AutoMap.txt.
type AutoMaps []*AutoMapRecord
