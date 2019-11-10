package d2enum

type AnimationMode int

const (
	AnimationModePlayerDeath       AnimationMode = 0  // DT
	AnimationModePlayerNeutral     AnimationMode = 1  // NU
	AnimationModePlayerWalk        AnimationMode = 2  // WL
	AnimationModePlayerRun         AnimationMode = 3  // RN
	AnimationModePlayerGetHit      AnimationMode = 4  // GH
	AnimationModePlayerTownNeutral AnimationMode = 5  // TN
	AnimationModePlayerTownWalk    AnimationMode = 6  // TW
	AnimationModePlayerAttack1     AnimationMode = 7  // A1
	AnimationModePlayerAttack2     AnimationMode = 8  // A2
	AnimationModePlayerBlock       AnimationMode = 9  // BL
	AnimationModePlayerCast        AnimationMode = 10 // SC
	AnimationModePlayerThrow       AnimationMode = 11 // TH
	AnimationModePlayerKick        AnimationMode = 12 // KK
	AnimationModePlayerSkill1      AnimationMode = 13 // S1
	AnimationModePlayerSkill2      AnimationMode = 14 // S2
	AnimationModePlayerSkill3      AnimationMode = 15 // S3
	AnimationModePlayerSkill4      AnimationMode = 16 // S4
	AnimationModePlayerDead        AnimationMode = 17 // DD
	AnimationModePlayerSequence    AnimationMode = 18 // GH
	AnimationModePlayerKnockBack   AnimationMode = 19 // GH
	AnimationModeMonsterDeath      AnimationMode = 20 // DT
	AnimationModeMonsterNeutral    AnimationMode = 21 // NU
	AnimationModeMonsterWalk       AnimationMode = 22 // WL
	AnimationModeMonsterGetHit     AnimationMode = 23 // GH
	AnimationModeMonsterAttack1    AnimationMode = 24 // A1
	AnimationModeMonsterAttack2    AnimationMode = 25 // A2
	AnimationModeMonsterBlock      AnimationMode = 26 // BL
	AnimationModeMonsterCast       AnimationMode = 27 // SC
	AnimationModeMonsterSkill1     AnimationMode = 28 // S1
	AnimationModeMonsterSkill2     AnimationMode = 29 // S2
	AnimationModeMonsterSkill3     AnimationMode = 30 // S3
	AnimationModeMonsterSkill4     AnimationMode = 31 // S4
	AnimationModeMonsterDead       AnimationMode = 32 // DD
	AnimationModeMonsterKnockback  AnimationMode = 33 // GH
	AnimationModeMonsterSequence   AnimationMode = 34 // xx
	AnimationModeMonsterRun        AnimationMode = 35 // RN
	AnimationModeObjectNeutral     AnimationMode = 36 // NU
	AnimationModeObjectOperating   AnimationMode = 37 // OP
	AnimationModeObjectOpened      AnimationMode = 38 // ON
	AnimationModeObjectSpecial1    AnimationMode = 39 // S1
	AnimationModeObjectSpecial2    AnimationMode = 40 // S2
	AnimationModeObjectSpecial3    AnimationMode = 41 // S3
	AnimationModeObjectSpecial4    AnimationMode = 42 // S4
	AnimationModeObjectSpecial5    AnimationMode = 43 // S5
)

//go:generate stringer -linecomment -type AnimationMode
