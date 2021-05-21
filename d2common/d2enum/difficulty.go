package d2enum

//go:generate stringer -linecomment -type DifficultyType -output difficulty_string.go

// DifficultyType is an enum for the possible difficulties
type DifficultyType uint8

// Difficulty levels
const (
	DifficultyNormal    DifficultyType = iota // normal
	DifficultyNightmare                       // nightmare
	DifficultyHell                            // hell
)
