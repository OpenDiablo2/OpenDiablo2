package d2enum

// MonUModConstIndex is used as an index into d2datadict.MonsterUniqueModifierConstants
type MonUModConstIndex int

const (
	//Champion chance
	ChampionChance MonUModConstIndex = iota

	//MinionHPBonus is the HP bonus received by minions
	MinionHPBonus
	MinionHPBonusNightmare
	MinionHPBonusHell

	//ChampionHPBonus is the HP bonus received by champions
	ChampionHPBonus
	ChampionHPBonusNightmare
	ChampionHPBonusHell

	//UniqueHPBonus is the HP bonus received by random uniques
	UniqueHPBonus
	UniqueHPBonusNightmare
	UniqueHPBonusHell

	//Attack rating and damage bonus of champions
	ChampionAttackRatingBonus
	ChampionDamageBonus

	//Attack rating and damage bonus of strong minions
	StrongMinionAttackRatingBonus
	StrongMinionDamageBonus

	//Minion elemental damage bonus
	MinionElementalDamageMinBonus
	MinionElementalDamageMinBonusNightmare
	MinionElementalDamageMinBonusHell

	MinionElementalDamageMaxBonus
	MinionElementalDamageMaxBonusNightmare
	MinionElementalDamageMaxBonusHell

	//Minion elemental damage bonus
	ChampionElementalDamageMinBonus
	ChampionElementalDamageMinBonusNightmare
	ChampionElementalDamageMinBonusHell

	ChampionElementalDamageMaxBonus
	ChampionElementalDamageMaxBonusNightmare
	ChampionElementalDamageMaxBonusHell

	//Unique elemental damage bonus
	UniqueElementalDamageMinBonus
	UniqueElementalDamageMinBonusNightmare
	UniqueElementalDamageMinBonusHell

	UniqueElementalDamageMaxBonus
	UniqueElementalDamageMaxBonusNightmare
	UniqueElementalDamageMaxBonusHell
)
