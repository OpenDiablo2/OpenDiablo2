package d2enum

// DifficultyType is an enum for the possible difficulties
type DifficultyType int

const (
	// DifficultyNormal is the normal difficulty
	DifficultyNormal DifficultyType = iota
	// DifficultyNightmare is the nightmare difficulty
	DifficultyNightmare
	// DifficultyHell is the hell difficulty
	DifficultyHell
)
