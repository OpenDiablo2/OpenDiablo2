package d2enum

type PlayerAnimationMode int
type MonsterAnimationMode int
type ObjectAnimationMode int

const (
	AnimationModePlayerDeath       PlayerAnimationMode = iota // DT
	AnimationModePlayerNeutral                                // NU
	AnimationModePlayerWalk                                   // WL
	AnimationModePlayerRun                                    // RN
	AnimationModePlayerGetHit                                 // GH
	AnimationModePlayerTownNeutral                            // TN
	AnimationModePlayerTownWalk                               // TW
	AnimationModePlayerAttack1                                // A1
	AnimationModePlayerAttack2                                // A2
	AnimationModePlayerBlock                                  // BL
	AnimationModePlayerCast                                   // SC
	AnimationModePlayerThrow                                  // TH
	AnimationModePlayerKick                                   // KK
	AnimationModePlayerSkill1                                 // S1
	AnimationModePlayerSkill2                                 // S2
	AnimationModePlayerSkill3                                 // S3
	AnimationModePlayerSkill4                                 // S4
	AnimationModePlayerDead                                   // DD
	AnimationModePlayerSequence                               // GH
	AnimationModePlayerKnockBack                              // GH
)
const (
	AnimationModeMonsterDeath     MonsterAnimationMode = iota // DT
	AnimationModeMonsterNeutral                               // NU
	AnimationModeMonsterWalk                                  // WL
	AnimationModeMonsterGetHit                                // GH
	AnimationModeMonsterAttack1                               // A1
	AnimationModeMonsterAttack2                               // A2
	AnimationModeMonsterBlock                                 // BL
	AnimationModeMonsterCast                                  // SC
	AnimationModeMonsterSkill1                                // S1
	AnimationModeMonsterSkill2                                // S2
	AnimationModeMonsterSkill3                                // S3
	AnimationModeMonsterSkill4                                // S4
	AnimationModeMonsterDead                                  // DD
	AnimationModeMonsterKnockback                             // GH
	AnimationModeMonsterSequence                              // xx
	AnimationModeMonsterRun                                   // RN
)
const (
	AnimationModeObjectNeutral   ObjectAnimationMode = iota // NU
	AnimationModeObjectOperating                            // OP
	AnimationModeObjectOpened                               // ON
	AnimationModeObjectSpecial1                             // S1
	AnimationModeObjectSpecial2                             // S2
	AnimationModeObjectSpecial3                             // S3
	AnimationModeObjectSpecial4                             // S4
	AnimationModeObjectSpecial5                             // S5
)

//go:generate stringer -linecomment -type PlayerAnimationMode
//go:generate stringer -linecomment -type MonsterAnimationMode
//go:generate stringer -linecomment -type ObjectAnimationMode
