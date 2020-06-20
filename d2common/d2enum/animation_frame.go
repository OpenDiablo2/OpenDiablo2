package d2enum

// AnimationFrame represents a single frame of animation.
type AnimationFrame int

const (
	// AnimationFrameNoEvent represents an animation frame with no event
	AnimationFrameNoEvent AnimationFrame = iota

	// AnimationFrameAttack represents an animation frame with an attack event
	AnimationFrameAttack

	// AnimationFrameMissile represents an animation frame with a missile event
	AnimationFrameMissile

	// AnimationFrameSound represents an animation frame with a sound event
	AnimationFrameSound

	// AnimationFrameSkill represents an animation frame with a skill event
	AnimationFrameSkill
)
