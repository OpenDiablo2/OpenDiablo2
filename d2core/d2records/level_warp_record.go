package d2records

// LevelWarps loaded from txt records
type LevelWarps map[int]*LevelWarpRecord

// LevelWarpRecord is a representation of a row from lvlwarp.txt
// it describes the warp graphics offsets and dimensions for levels
type LevelWarpRecord struct {
	Name       string
	Direction  string
	SelectX    int
	SelectY    int
	SelectDX   int
	SelectDY   int
	ID         int
	ExitWalkY  int
	OffsetX    int
	OffsetY    int
	Tiles      int
	ExitWalkX  int
	LitVersion bool
}
