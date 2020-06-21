package d2enum

// from levels.txt

type LevelGenerationType int

const (
	LevelTypeRandomMaze LevelGenerationType = iota
	LevelTypePreset
	LevelTypeWilderness
)
