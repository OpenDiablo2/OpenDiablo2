package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

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

// CharStats holds all of the CharStatsRecords
//nolint:gochecknoglobals // Currently global by design, only written once
var CharStats map[d2enum.Hero]*CharStatsRecord
var charStringMap map[string]d2enum.Hero         //nolint:gochecknoglobals // Currently global by design
var weaponTokenMap map[string]d2enum.WeaponClass //nolint:gochecknoglobals // Currently global by design

// LoadCharStats loads charstats.txt file contents into map[d2enum.Hero]*CharStatsRecord
//nolint:funlen // Makes no sense to split
// LoadCharStats loads charstats.txt file contents into map[d2enum.Hero]*CharStatsRecord
func LoadCharStats(file []byte) {
	CharStats = make(map[d2enum.Hero]*CharStatsRecord)

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

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &CharStatsRecord{
			Class: charStringMap[d.String("class")],

			InitStr:     d.Number("str"),
			InitDex:     d.Number("dex"),
			InitVit:     d.Number("vit"),
			InitEne:     d.Number("int"),
			InitStamina: d.Number("stamina"),

			ManaRegen:   d.Number("ManaRegen"),
			ToHitFactor: d.Number("ToHitFactor"),

			VelocityWalk:    d.Number("WalkVelocity"),
			VelocityRun:     d.Number("RunVelocity"),
			StaminaRunDrain: d.Number("RunDrain"),

			LifePerLevel:    d.Number("LifePerLevel"),
			ManaPerLevel:    d.Number("ManaPerLevel"),
			StaminaPerLevel: d.Number("StaminaPerLevel"),

			LifePerVit:    d.Number("LifePerVitality"),
			ManaPerEne:    d.Number("ManaPerMagic"),
			StaminaPerVit: d.Number("StaminaPerVitality"),

			StatPerLevel: d.Number("StatPerLevel"),
			BlockFactor:  d.Number("BlockFactor"),

			StartSkillBonus:   d.String("StartSkill"),
			SkillStrAll:       d.String("StrAllSkills"),
			SkillStrClassOnly: d.String("StrClassOnly"),

			BaseSkill: [10]string{
				d.String("Skill 1"),
				d.String("Skill 2"),
				d.String("Skill 3"),
				d.String("Skill 4"),
				d.String("Skill 5"),
				d.String("Skill 6"),
				d.String("Skill 7"),
				d.String("Skill 8"),
				d.String("Skill 9"),
				d.String("Skill 10"),
			},

			SkillStrTab: [3]string{
				d.String("StrSkillTab1"),
				d.String("StrSkillTab2"),
				d.String("StrSkillTab3"),
			},

			BaseWeaponClass: weaponTokenMap[d.String("baseWClass")],

			StartItem: [10]string{
				d.String("item1"),
				d.String("item2"),
				d.String("item3"),
				d.String("item4"),
				d.String("item5"),
				d.String("item6"),
				d.String("item7"),
				d.String("item8"),
				d.String("item9"),
				d.String("item10"),
			},

			StartItemLocation: [10]string{
				d.String("item1loc"),
				d.String("item2loc"),
				d.String("item3loc"),
				d.String("item4loc"),
				d.String("item5loc"),
				d.String("item6loc"),
				d.String("item7loc"),
				d.String("item8loc"),
				d.String("item9loc"),
				d.String("item10loc"),
			},

			StartItemCount: [10]int{
				d.Number("item1count"),
				d.Number("item2count"),
				d.Number("item3count"),
				d.Number("item4count"),
				d.Number("item5count"),
				d.Number("item6count"),
				d.Number("item7count"),
				d.Number("item8count"),
				d.Number("item9count"),
				d.Number("item10count"),
			},
		}
		CharStats[record.Class] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d CharStats records", len(CharStats))
}
