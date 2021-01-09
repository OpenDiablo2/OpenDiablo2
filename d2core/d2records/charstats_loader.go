package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// nolint:funlen // cant reduce
func charStatsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(CharStats)

	stringMap := map[string]d2enum.Hero{
		"Amazon":      d2enum.HeroAmazon,
		"Barbarian":   d2enum.HeroBarbarian,
		"Druid":       d2enum.HeroDruid,
		"Assassin":    d2enum.HeroAssassin,
		"Necromancer": d2enum.HeroNecromancer,
		"Paladin":     d2enum.HeroPaladin,
		"Sorceress":   d2enum.HeroSorceress,
	}

	tokenMap := map[string]d2enum.WeaponClass{
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

	for d.Next() {
		record := &CharStatsRecord{
			Class: stringMap[d.String("class")],

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

			BaseWeaponClass: tokenMap[d.String("baseWClass")],

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
		records[record.Class] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d CharStats records", len(records))

	r.Character.Stats = records

	return nil
}
