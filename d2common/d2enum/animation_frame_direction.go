package d2enum

// AnimationFrameDirection enumerates animation frame directions used in d2datadict.MonsterSequenceFrame
type AnimationFrameDirection int

// Sprite frame directions
const (
	SouthWest AnimationFrameDirection = iota
	NorthWest
	NorthEast
	SouthEast
	South
	West
	North
	East
)
