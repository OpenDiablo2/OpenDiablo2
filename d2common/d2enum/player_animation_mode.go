package d2enum

//go:generate stringer -linecomment -type PlayerAnimationMode -output player_animation_mode_string.go

// PlayerAnimationMode represents player animation modes
type PlayerAnimationMode int

// Player animation modes
const (
	PlayerAnimationModeDeath       PlayerAnimationMode = iota // DT
	PlayerAnimationModeNeutral                                // NU
	PlayerAnimationModeWalk                                   // WL
	PlayerAnimationModeRun                                    // RN
	PlayerAnimationModeGetHit                                 // GH
	PlayerAnimationModeTownNeutral                            // TN
	PlayerAnimationModeTownWalk                               // TW
	PlayerAnimationModeAttack1                                // A1
	PlayerAnimationModeAttack2                                // A2
	PlayerAnimationModeBlock                                  // BL
	PlayerAnimationModeCast                                   // SC
	PlayerAnimationModeThrow                                  // TH
	PlayerAnimationModeKick                                   // KK
	PlayerAnimationModeSkill1                                 // S1
	PlayerAnimationModeSkill2                                 // S2
	PlayerAnimationModeSkill3                                 // S3
	PlayerAnimationModeSkill4                                 // S4
	PlayerAnimationModeDead                                   // DD
	PlayerAnimationModeSequence                               // GH
	PlayerAnimationModeKnockBack                              // GH
	PlayerAnimationModeNone                                   // "" - aura skills, e.g. Paladin's Concentration Aura
)
