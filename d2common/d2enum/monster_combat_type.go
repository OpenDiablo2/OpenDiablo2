package d2enum

// MonsterCombatType is used for setting the monster as melee or ranged
type MonsterCombatType int

const (
	// MonsterMelee is a flag that sets the monster as melee-only
	MonsterMelee MonsterCombatType = iota

	// MonsterRanged is a flag that sets the monster as ranged-only
	MonsterRanged
)
