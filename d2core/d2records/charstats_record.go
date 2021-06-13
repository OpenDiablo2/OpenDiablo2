package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// CharStats holds all of the CharStatRecords
type CharStats map[d2enum.Hero]*CharStatRecord

// CharStatRecord is a struct that represents a single row from charstats.txt
type CharStatRecord struct {
	BaseSkill         [10]string
	StartItemLocation [10]string
	StartItem         [10]string
	SkillStrTab       [3]string
	SkillStrClassOnly string
	SkillStrAll       string
	StartSkillBonus   string
	StartItemCount    [10]int
	LifePerVit        int
	VelocityRun       int
	StaminaRunDrain   int
	LifePerLevel      int
	ManaPerLevel      int
	StaminaPerLevel   int
	VelocityWalk      int
	ManaPerEne        int
	StaminaPerVit     int
	StatPerLevel      int
	BlockFactor       int
	ToHitFactor       int
	ManaRegen         int
	InitStamina       int
	InitEne           int
	InitVit           int
	BaseWeaponClass   d2enum.WeaponClass
	InitDex           int
	InitStr           int
	Class             d2enum.Hero
}
