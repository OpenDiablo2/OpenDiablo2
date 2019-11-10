package d2enum

type AnimationFrame int

const (
	AnimationFrameNoEvent AnimationFrame = 0
	AnimationFrameAttack  AnimationFrame = 1
	AnimationFrameMissile AnimationFrame = 2
	AnimationFrameSound   AnimationFrame = 3
	AnimationFrameSkill   AnimationFrame = 4
)
