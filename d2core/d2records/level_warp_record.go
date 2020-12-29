package d2records

// LevelWarps loaded from txt records
type LevelWarps map[int]*LevelWarpRecord

// LevelWarpRecord is a representation of a row from lvlwarp.txt
// it describes the warp graphics offsets and dimensions for levels
type LevelWarpRecord struct {
	Name       string
	ID         int
	SelectX    int
	SelectY    int
	SelectDX   int
	SelectDY   int
	ExitWalkX  int
	ExitWalkY  int
	OffsetX    int
	OffsetY    int
	LitVersion bool
	Tiles      int
	Direction  string
}
