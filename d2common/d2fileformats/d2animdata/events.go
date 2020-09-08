package d2animdata

// AnimationEvent represents an event that can happen on a frame of animation
type AnimationEvent byte

// Animation events
const (
	AnimationEventNone AnimationEvent = iota
	AnimationEventAttack
	AnimationEventMissile
	AnimationEventSound
	AnimationEventSkill
)
