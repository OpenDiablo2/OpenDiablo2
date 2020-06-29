package d2enum

// MonsterAlignmentType determines the hostility of the monster towards players
type MonsterAlignmentType int

const (
	// MonsterEnemy flag will make monsters hostile towards players
	MonsterEnemy MonsterAlignmentType = iota

	// MonsterFriend will make monsters friendly towards players
	// this is likely used by NPC's and summons
	MonsterFriend

	// MonsterNeutral will make monsters not care about players or monsters
	// this flag is used for `critter` type monsters
	MonsterNeutral
)
