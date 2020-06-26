package d2enum

// from levels.txt, field `DrlgType`
// https://d2mods.info/forum/kb/viewarticle?a=301

// Setting for Level Generation: You have 3 possibilities here:
// 1 Random Maze
// 2 Preset Area
// 3 Wilderness level

type LevelGenerationType int

const (
	LevelTypeRandomMaze LevelGenerationType = iota
	LevelTypePreset
	LevelTypeWilderness
)
