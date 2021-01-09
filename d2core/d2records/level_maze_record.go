package d2records

// LevelMazeDetails stores all of the LevelMazeDetailsRecords
type LevelMazeDetails map[int]*LevelMazeDetailsRecord

// LevelMazeDetailsRecord is a representation of a row from lvlmaze.txt
// these records define the parameters passed to the maze level generator
type LevelMazeDetailsRecord struct {
	// descriptive, not loaded in game. Corresponds with Name field in
	// Levels.txt
	Name string // Name

	// ID from Levels.txt
	// NOTE: Cave 1 is the Den of Evil, its associated treasure level is quest
	// only.
	LevelID int // Level

	// the minimum number of .ds1 map sections that will make up the maze in
	// Normal, Nightmare and Hell difficulties.
	NumRoomsNormal    int // Rooms
	NumRoomsNightmare int // Rooms(N)
	NumRoomsHell      int // Rooms(H)

	// the size in the X\Y direction of any component ds1 map section.
	SizeX int // SizeX
	SizeY int // SizeY

	// Possibly related to how adjacent .ds1s are connected with each other,
	// but what the different values are for is unknown.
	// Merge int // Merge

	// Included in the original Diablo II beta tests and in the demo version.
	// Beta
}
