package d2enum

//go:generate stringer -linecomment -type MonsterAnimationMode -output monster_animation_mode_string.go

// MonsterAnimationMode represents monster animation modes
type MonsterAnimationMode int

// Monster animation modes
const (
	MonsterAnimationModeDeath     MonsterAnimationMode = iota // DT
	MonsterAnimationModeNeutral                               // NU
	MonsterAnimationModeWalk                                  // WL
	MonsterAnimationModeGetHit                                // GH
	MonsterAnimationModeAttack1                               // A1
	MonsterAnimationModeAttack2                               // A2
	MonsterAnimationModeBlock                                 // BL
	MonsterAnimationModeCast                                  // SC
	MonsterAnimationModeSkill1                                // S1
	MonsterAnimationModeSkill2                                // S2
	MonsterAnimationModeSkill3                                // S3
	MonsterAnimationModeSkill4                                // S4
	MonsterAnimationModeDead                                  // DD
	MonsterAnimationModeKnockback                             // GH
	MonsterAnimationModeSequence                              // xx
	MonsterAnimationModeRun                                   // RN
)
