package d2enum

// AnimationFrameEvent enumerates events used in d2datadict.MonsterSequenceFrame
type AnimationFrameEvent int

// Animation frame events
const (
	NoEvent AnimationFrameEvent = iota
	MeleeAttack
	MissileAttack
	PlaySound
	LaunchSpell
)
