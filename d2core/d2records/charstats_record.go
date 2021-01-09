package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// CharStats holds all of the CharStatsRecords
type CharStats map[d2enum.Hero]*CharStatsRecord

// CharStatsRecord is a struct that represents a single row from charstats.txt
type CharStatsRecord struct {
	Class d2enum.Hero

	// the initial stats at character level 1
	InitStr     int // initial strength
	InitDex     int // initial dexterity
	InitVit     int // initial vitality
	InitEne     int // initial energy
	InitStamina int // initial stamina

	ManaRegen   int // number of seconds to regen mana completely
	ToHitFactor int // added to basic AR of character class

	VelocityWalk    int // velocity of the character while walking
	VelocityRun     int // velocity of the character while running
	StaminaRunDrain int // rate of stamina loss, lower is longer drain time

	// NOTE: Each point of Life/Mana/Stamina is divided by 256 for precision.
	// value is in fourths, lowest possible is 64/256
	LifePerLevel    int // amount of life per character level
	ManaPerLevel    int // amount of mana per character level
	StaminaPerLevel int // amount of stamina per character level

	LifePerVit    int // life per point of vitality
	ManaPerEne    int // mana per point of energy
	StaminaPerVit int // stamina per point of vitality

	StatPerLevel int // amount of stat points per level

	BlockFactor int // added to base shield block% in armor.txt (display & calc)

	// appears on starting weapon
	StartSkillBonus string // a key that points to a property

	// The skills the character class starts with (always available)
	BaseSkill [10]string // the base skill keys of the character, always available

	// string for bonus to class skills (ex: +1 to all Amazon skills).
	SkillStrAll       string    // string for bonus to all skills
	SkillStrTab       [3]string // string for bonus per skill tabs
	SkillStrClassOnly string    // string for class-exclusive skills

	BaseWeaponClass d2enum.WeaponClass // controls animation when unarmed

	StartItem         [10]string // tokens for the starting items
	StartItemLocation [10]string // locations of the starting items
	StartItemCount    [10]int    // amount of the starting items
}
