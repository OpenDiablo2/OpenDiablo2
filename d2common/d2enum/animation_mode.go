package d2enum

// PlayerAnimationMode represents player animation modes
type PlayerAnimationMode int

// MonsterAnimationMode represents monster animation modes
type MonsterAnimationMode int

// ObjectAnimationMode represents object animation modes
type ObjectAnimationMode int

const (
	// AnimationModePlayerDeath represents DT
	AnimationModePlayerDeath PlayerAnimationMode = iota

	// AnimationModePlayerNeutral represents NU
	AnimationModePlayerNeutral

	// AnimationModePlayerWalk represents WL
	AnimationModePlayerWalk

	// AnimationModePlayerRun represents RN
	AnimationModePlayerRun

	// AnimationModePlayerGetHit represents GH
	AnimationModePlayerGetHit

	// AnimationModePlayerTownNeutral represents TN
	AnimationModePlayerTownNeutral

	// AnimationModePlayerTownWalk represents TW
	AnimationModePlayerTownWalk

	// AnimationModePlayerAttack1 represents A1
	AnimationModePlayerAttack1

	// AnimationModePlayerAttack2 represents A2
	AnimationModePlayerAttack2

	// AnimationModePlayerBlock represents BL
	AnimationModePlayerBlock

	// AnimationModePlayerCast represents SC
	AnimationModePlayerCast

	// AnimationModePlayerThrow represents TH
	AnimationModePlayerThrow

	// AnimationModePlayerKick represents KK
	AnimationModePlayerKick

	// AnimationModePlayerSkill1 represents S1
	AnimationModePlayerSkill1

	// AnimationModePlayerSkill2 represents S2
	AnimationModePlayerSkill2

	// AnimationModePlayerSkill3 represents S3
	AnimationModePlayerSkill3

	// AnimationModePlayerSkill4 represents S4
	AnimationModePlayerSkill4

	// AnimationModePlayerDead represents DD
	AnimationModePlayerDead

	// AnimationModePlayerSequence represents GH
	AnimationModePlayerSequence

	// AnimationModePlayerKnockBack represents GH
	AnimationModePlayerKnockBack
)
const (

	// AnimationModeMonsterDeath represents DT
	AnimationModeMonsterDeath MonsterAnimationMode = iota

	// AnimationModeMonsterNeutral represents NU
	AnimationModeMonsterNeutral

	// AnimationModeMonsterWalk represents WL
	AnimationModeMonsterWalk

	// AnimationModeMonsterGetHit represents GH
	AnimationModeMonsterGetHit

	// AnimationModeMonsterAttack1 represents A1
	AnimationModeMonsterAttack1

	// AnimationModeMonsterAttack2 represents A2
	AnimationModeMonsterAttack2

	// AnimationModeMonsterBlock represents BL
	AnimationModeMonsterBlock

	// AnimationModeMonsterCast represents SC
	AnimationModeMonsterCast

	// AnimationModeMonsterSkill1 represents S1
	AnimationModeMonsterSkill1

	// AnimationModeMonsterSkill2 represents S2
	AnimationModeMonsterSkill2

	// AnimationModeMonsterSkill3 represents S3
	AnimationModeMonsterSkill3

	// AnimationModeMonsterSkill4 represents S4
	AnimationModeMonsterSkill4

	// AnimationModeMonsterDead represents DD
	AnimationModeMonsterDead

	// AnimationModeMonsterKnockback represents GH
	AnimationModeMonsterKnockback

	// AnimationModeMonsterSequence represents xx
	AnimationModeMonsterSequence

	// AnimationModeMonsterRun represents RN
	AnimationModeMonsterRun
)
const (

	// AnimationModeObjectNeutral represents NU
	AnimationModeObjectNeutral ObjectAnimationMode = iota

	// AnimationModeObjectOperating represents OP
	AnimationModeObjectOperating

	// AnimationModeObjectOpened represents ON
	AnimationModeObjectOpened

	// AnimationModeObjectSpecial1 represents S1
	AnimationModeObjectSpecial1

	// AnimationModeObjectSpecial2 represents S2
	AnimationModeObjectSpecial2

	// AnimationModeObjectSpecial3 represents S3
	AnimationModeObjectSpecial3

	// AnimationModeObjectSpecial4 represents S4
	AnimationModeObjectSpecial4

	// AnimationModeObjectSpecial5 represents S5
	AnimationModeObjectSpecial5
)

//go:generate stringer -linecomment -type PlayerAnimationMode
//go:generate stringer -linecomment -type MonsterAnimationMode
//go:generate stringer -linecomment -type ObjectAnimationMode
//go:generate string2enum -samepkg -linecomment -type ObjectAnimationMode
