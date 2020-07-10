package d2enum

// AnimationFrame represents a single frame of animation.
type AnimationFrame int

// AnimationFrame types
const (
	AnimationFrameNoEvent AnimationFrame = iota
	AnimationFrameAttack
	AnimationFrameMissile
	AnimationFrameSound
	AnimationFrameSkill
)
