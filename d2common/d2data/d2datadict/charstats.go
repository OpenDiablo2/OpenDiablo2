package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type CharStatsRecord struct {
	Class d2enum.Hero

	// the initial stats at character level 1
	InitStr     int
	InitDex     int
	InitVit     int
	InitEne     int
	InitStamina int

	ManaRegen   int // number of seconds to regen mana completely
	ToHitFactor int // added to basic AR of character class

	VelocityWalk    int
	VelocityRun     int
	StaminaRunDrain int // rate of stamina loss, lower is longer drain time

	// NOTE: Each point of Life/Mana/Stamina is divided by 256 for precision.
	LifePerLevel    int // value is in fourths, lowest possible is 64/256
	ManaPerLevel    int
	StaminaPerLevel int

	LifePerVit    int // life per point of vitality
	ManaPerEne    int // mana per point of energy
	StaminaPerVit int

	StatPerLevel int // stat points per level

	BlockFactor int // added to base shield block% in armor.txt (display & calc)

	// appears on starting weapon
	StartSkillBonus string

	// The skills the character class starts with (always available)
	BaseSkill [10]string

	// string for bonus to class skills (ex: +1 to all Amazon skills).
	SkillStrAll       string
	SkillStrTab       [3]string // string for bonus per skill tabs
	SkillStrClassOnly string    // string for class-exclusive skills

	BaseWeaponClass d2enum.WeaponClass // controls animation when unarmed

	StartItem         [10]string
	StartItemLocation [10]string
	StartItemCount    [10]int
}

var CharStats map[d2enum.Hero]*CharStatsRecord
var charStringMap map[string]d2enum.Hero
var weaponTokenMap map[string]d2enum.WeaponClass

func LoadCharStats(file []byte) {

	charStringMap = map[string]d2enum.Hero{
		"Amazon":      d2enum.HeroAmazon,
		"Barbarian":   d2enum.HeroBarbarian,
		"Druid":       d2enum.HeroDruid,
		"Assassin":    d2enum.HeroAssassin,
		"Necromancer": d2enum.HeroNecromancer,
		"Paladin":     d2enum.HeroPaladin,
		"Sorceress":   d2enum.HeroSorceress,
	}

	weaponTokenMap = map[string]d2enum.WeaponClass{
		"":    d2enum.WeaponClassNone,
		"hth": d2enum.WeaponClassHandToHand,
		"bow": d2enum.WeaponClassBow,
		"1hs": d2enum.WeaponClassOneHandSwing,
		"1ht": d2enum.WeaponClassOneHandThrust,
		"stf": d2enum.WeaponClassStaff,
		"2hs": d2enum.WeaponClassTwoHandSwing,
		"2ht": d2enum.WeaponClassTwoHandThrust,
		"xbw": d2enum.WeaponClassCrossbow,
		"1js": d2enum.WeaponClassLeftJabRightSwing,
		"1jt": d2enum.WeaponClassLeftJabRightThrust,
		"1ss": d2enum.WeaponClassLeftSwingRightSwing,
		"1st": d2enum.WeaponClassLeftSwingRightThrust,
		"ht1": d2enum.WeaponClassOneHandToHand,
		"ht2": d2enum.WeaponClassTwoHandToHand,
	}

	d := d2common.LoadDataDictionary(string(file))
	CharStats = make(map[d2enum.Hero]*CharStatsRecord, len(d.Data))

	for idx := range d.Data {
		record := &CharStatsRecord{
			Class: charStringMap[d.GetString("class", idx)],

			InitStr:     d.GetNumber("str", idx),
			InitDex:     d.GetNumber("dex", idx),
			InitVit:     d.GetNumber("vit", idx),
			InitEne:     d.GetNumber("int", idx),
			InitStamina: d.GetNumber("stamina", idx),

			ManaRegen:   d.GetNumber("ManaRegen", idx),
			ToHitFactor: d.GetNumber("ToHitFactor", idx),

			VelocityWalk:    d.GetNumber("WalkVelocity", idx),
			VelocityRun:     d.GetNumber("RunVelocity", idx),
			StaminaRunDrain: d.GetNumber("RunDrain", idx),

			LifePerLevel:    d.GetNumber("LifePerLevel", idx),
			ManaPerLevel:    d.GetNumber("ManaPerLevel", idx),
			StaminaPerLevel: d.GetNumber("StaminaPerLevel", idx),

			LifePerVit:    d.GetNumber("LifePerVitality", idx),
			ManaPerEne:    d.GetNumber("ManaPerMagic", idx),
			StaminaPerVit: d.GetNumber("StaminaPerVitality", idx),

			StatPerLevel: d.GetNumber("StatPerLevel", idx),
			BlockFactor:  d.GetNumber("BlockFactor", idx),

			StartSkillBonus:   d.GetString("StartSkill", idx),
			SkillStrAll:       d.GetString("StrAllSkills", idx),
			SkillStrClassOnly: d.GetString("StrClassOnly", idx),

			BaseSkill: [10]string{
				d.GetString("Skill 1", idx),
				d.GetString("Skill 2", idx),
				d.GetString("Skill 3", idx),
				d.GetString("Skill 4", idx),
				d.GetString("Skill 5", idx),
				d.GetString("Skill 6", idx),
				d.GetString("Skill 7", idx),
				d.GetString("Skill 8", idx),
				d.GetString("Skill 9", idx),
				d.GetString("Skill 10", idx),
			},

			SkillStrTab: [3]string{
				d.GetString("StrSkillTab1", idx),
				d.GetString("StrSkillTab2", idx),
				d.GetString("StrSkillTab3", idx),
			},

			BaseWeaponClass: weaponTokenMap[d.GetString("baseWClass", idx)],

			StartItem: [10]string{
				d.GetString("item1", idx),
				d.GetString("item2", idx),
				d.GetString("item3", idx),
				d.GetString("item4", idx),
				d.GetString("item5", idx),
				d.GetString("item6", idx),
				d.GetString("item7", idx),
				d.GetString("item8", idx),
				d.GetString("item9", idx),
				d.GetString("item10", idx),
			},

			StartItemLocation: [10]string{
				d.GetString("item1loc", idx),
				d.GetString("item2loc", idx),
				d.GetString("item3loc", idx),
				d.GetString("item4loc", idx),
				d.GetString("item5loc", idx),
				d.GetString("item6loc", idx),
				d.GetString("item7loc", idx),
				d.GetString("item8loc", idx),
				d.GetString("item9loc", idx),
				d.GetString("item10loc", idx),
			},

			StartItemCount: [10]int{
				d.GetNumber("item1count", idx),
				d.GetNumber("item2count", idx),
				d.GetNumber("item3count", idx),
				d.GetNumber("item4count", idx),
				d.GetNumber("item5count", idx),
				d.GetNumber("item6count", idx),
				d.GetNumber("item7count", idx),
				d.GetNumber("item8count", idx),
				d.GetNumber("item9count", idx),
				d.GetNumber("item10count", idx),
			},
		}
		CharStats[record.Class] = record
	}
	log.Printf("Loaded %d CharStats records", len(CharStats))
}
