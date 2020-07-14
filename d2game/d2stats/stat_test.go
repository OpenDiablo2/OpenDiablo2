package d2stats

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

func TestStat_AssetInit(t *testing.T) {
	if err := d2config.Load(); err != nil {
		panic(err)
	}
	if err := d2asset.Initialize(nil, nil); err != nil {
		panic(err)
	}

	tablePaths := []string{
		d2resource.PatchStringTable,
		d2resource.ExpansionStringTable,
		d2resource.StringTable,
	}

	for _, tablePath := range tablePaths {
		data, err := d2asset.LoadFile(tablePath)
		if err != nil {
			panic(err)
		}

		d2common.LoadTextDictionary(data)
	}

	if data, err := d2asset.LoadFile(d2resource.ItemStatCost); err != nil {
		panic(err)
	} else {
		d2datadict.LoadItemStatCosts(data)
	}

	// necessary for +skills stats descriptions
	if data, err := d2asset.LoadFile(d2resource.CharStats); err != nil {
		panic(err)
	} else {
		d2datadict.LoadCharStats(data)
	}

	// necessary for `chance to cast` skills
	if data, err := d2asset.LoadFile(d2resource.Skills); err != nil {
		panic(err)
	} else {
		d2datadict.LoadSkills(data)
	}

	if data, err := d2asset.LoadFile(d2resource.SkillDesc); err != nil {
		panic(err)
	} else {
		d2datadict.LoadSkillDescriptions(data)
	}

	// for attack/damage vs monster type
	if data, err := d2asset.LoadFile(d2resource.MonStats); err != nil {
		panic(err)
	} else {
		d2datadict.LoadMonStats(data)
	}
}

func TestStat_Clone(t *testing.T) {
	r := d2datadict.ItemStatCosts["strength"]
	s1 := CreateStat(r, 5)
	s2 := s1.Clone()

	// make sure the stats are distinct
	if &s1 == &s2 {
		t.Errorf("stats share the same pointer %d == %d", &s1, &s2)
	}

	// make sure the stat values are unique
	vs1, vs2 := s1.Values, s2.Values
	if &vs1 == &vs2 {
		t.Errorf("stat values share the same pointer %d == %d", &s1, &s2)
	}

	s2.Values[0] = 6
	v1, v2 := s1.Values[0], s2.Values[0]

	// make sure the value ranges are distinct
	if v1 == v2 {
		t.Errorf("stat value ranges should not be equal")
	}
}

// func TestStat_Description(t *testing.T) {
//	r := d2datadict.ItemStatCosts["strength"]
//
//	s1 := CreateStat(r, 5)
//	desc := s1.Description()
//
//	if desc != "+5 to Strength" {
//		t.Errorf("unexpected description string: %s", desc)
//	}
//
//	s1.Values[0] = -5
//	desc = s1.Description()
//
//	if desc != "-5 to Strength" {
//		t.Errorf("unexpected description string: %s", desc)
//	}
//}

// func TestStat_DescriptionAll(t *testing.T) {
//	records := d2datadict.ItemStatCosts
//	random := func(a, b float64) int {
//		return int(rand.Float64()*a - b)
//	}
//
//	for idx := range d2datadict.StatDescriptionFormatStrings {
//		for key := range records {
//			if strings.Contains(key, "bytime") {
//				// i don't think these are used in d2
//				continue
//			}
//			fnID := records[key].DescFnID
//
//			if fnID != idx {
//				continue
//			}
//
//			val := random(50, -1)
//			statPositive := CreateStat(records[key], val)
//			statNegative := CreateStat(records[key], val * -1)
//			fmt.Printf("\"%s\", %d, \"%s\",\n",
//				key,
//				val,
//				statPositive.Description(),
//			)
//			fmt.Printf("\"%s\", %d, \"%s\",\n",
//				key,
//				val * -1,
//				statNegative.Description(),
//			)
//
//		}
//	}
// }

const (
	errStr string = "incorrect description string format for stat"
	errFmt string = "%v:\n\tKey: %v\n\tVal: %+v\n\texpected: %v\n\tgot: %v\n\n"
)

//nolint:funlen // these are the stats that use DescFn1, not making this shorter
func TestStat_DescriptionFn1(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_damagetargetac", 31, "+31 to Monster Defense Per Hit"},
		{"item_damagetargetac", -31, "-31 to Monster Defense Per Hit"},
		{"maxdamage", 48, "+48 to Maximum Damage"},
		{"maxdamage", -48, "-48 to Maximum Damage"},
		{"magicmindam", 34, "+34 magic damage"},
		{"magicmindam", -34, "-34 magic damage"},
		{"firemindam", 22, "+22 to Minimum Fire Damage"},
		{"firemindam", -22, "-22 to Minimum Fire Damage"},
		{"magicmaxdam", 22, "+22 magic damage"},
		{"magicmaxdam", -22, "-22 magic damage"},
		{"poisonmaxdam", 35, "+35 to Maximum Poison Damage"},
		{"poisonmaxdam", -35, "-35 to Maximum Poison Damage"},
		{"item_manaafterkill", 4, "+4 to Mana after each Kill"},
		{"item_manaafterkill", -4, "-4 to Mana after each Kill"},
		{"item_absorbcold", 8, "+8 Cold Absorb"},
		{"item_absorbcold", -8, "-8 Cold Absorb"},
		{"poisonmindam", 5, "+5 to Minimum Poison Damage"},
		{"poisonmindam", -5, "-5 to Minimum Poison Damage"},
		{"hpregen", 16, "Replenish Life +16"},
		{"hpregen", -16, "Drain Life -16"},
		{"item_healafterkill", 26, "+26 Life after each Kill"},
		{"item_healafterkill", -26, "-26 Life after each Kill"},
		{"item_demon_tohit", 41, "+41 to Attack Rating against Demons"},
		{"item_demon_tohit", -41, "-41 to Attack Rating against Demons"},
		{"item_absorbmagic", 11, "+11 Magic Absorb"},
		{"item_absorbmagic", -11, "-11 Magic Absorb"},
		{"item_absorblight", 20, "+20 Lightning Absorb"},
		{"item_absorblight", -20, "-20 Lightning Absorb"},
		{"item_absorbfire", 16, "+16 Fire Absorb"},
		{"item_absorbfire", -16, "-16 Fire Absorb"},
		{"secondary_maxdamage", 24, "+24 to Maximum Damage"},
		{"secondary_maxdamage", -24, "-24 to Maximum Damage"},
		{"armorclass_vs_hth", 15, "+15 Defense vs. Melee"},
		{"armorclass_vs_hth", -15, "-15 Defense vs. Melee"},
		{"vitality", 15, "+15 to Vitality"},
		{"vitality", -15, "-15 to Vitality"},
		{"coldmaxdam", 34, "+34 to Maximum Cold Damage"},
		{"coldmaxdam", -34, "-34 to Maximum Cold Damage"},
		{"item_lightradius", 11, "+11 to Light Radius"},
		{"item_lightradius", -11, "-11 to Light Radius"},
		{"maxmana", 11, "+11 to Mana"},
		{"maxmana", -11, "-11 to Mana"},
		{"item_healafterdemonkill", 19, "+19 Life after each Demon Kill"},
		{"item_healafterdemonkill", -19, "-19 Life after each Demon Kill"},
		{"lightmindam", 29, "+29 to Minimum Lightning Damage"},
		{"lightmindam", -29, "-29 to Minimum Lightning Damage"},
		{"item_elemskill", 44, "+44 to Fire Skills"},
		{"item_elemskill", -44, "-44 to Fire Skills"},
		{"strength", 15, "+15 to Strength"},
		{"strength", -15, "-15 to Strength"},
		{"lightmaxdam", 15, "+15 to Maximum Lightning Damage"},
		{"lightmaxdam", -15, "-15 to Maximum Lightning Damage"},
		{"armorclass_vs_missile", 38, "+38 Defense vs. Missile"},
		{"armorclass_vs_missile", -38, "-38 Defense vs. Missile"},
		{"secondary_mindamage", 11, "+11 to Minimum Damage"},
		{"secondary_mindamage", -11, "-11 to Minimum Damage"},
		{"maxhp", 44, "+44 to Life"},
		{"maxhp", -44, "-44 to Life"},
		{"tohit", 35, "+35 to Attack Rating"},
		{"tohit", -35, "-35 to Attack Rating"},
		{"item_normaldamage", 27, "Damage +27"},
		{"item_normaldamage", -27, "Damage -27"},
		{"maxstamina", 2, "+2 Maximum Stamina"},
		{"maxstamina", -2, "-2 Maximum Stamina"},
		{"firemaxdam", 8, "+8 to Maximum Fire Damage"},
		{"firemaxdam", -8, "-8 to Maximum Fire Damage"},
		{"item_undead_tohit", 31, "+31 to Attack Rating against Undead"},
		{"item_undead_tohit", -31, "-31 to Attack Rating against Undead"},
		{"armorclass", 49, "+49 Defense"},
		{"armorclass", -49, "-49 Defense"},
		{"dexterity", 4, "+4 to Dexterity"},
		{"dexterity", -4, "-4 to Dexterity"},
		{"coldmindam", 30, "+30 to Minimum Cold Damage"},
		{"coldmindam", -30, "-30 to Minimum Cold Damage"},
		{"item_kickdamage", 3, "+3 Kick Damage"},
		{"item_kickdamage", -3, "-3 Kick Damage"},
		{"item_allskills", 35, "+35 to All Skills"},
		{"item_allskills", -35, "-35 to All Skills"},
		{"energy", 16, "+16 to Energy"},
		{"energy", -16, "-16 to Energy"},
		{"mindamage", 9, "+9 to Minimum Damage"},
		{"mindamage", -9, "-9 to Minimum Damage"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn2(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"toblock", 25, "+25% Increased Chance of Blocking"},
		{"manarecoverybonus", 25, "Regenerate Mana +25%"},
		{"staminarecoverybonus", 25, "Heal Stamina Plus +25%"},
		{"damageresist", 25, "Damage Reduced by +25%"},
		{"lifedrainmindam", 25, "+25% Life stolen per hit"},
		{"manadrainmindam", 25, "+25% Mana stolen per hit"},
		{"item_maxdurability_percent", 25, "Increase Maximum Durability +25%"},
		{"item_maxhp_percent", 25, "Increase Maximum Life +25%"},
		{"item_maxmana_percent", 25, "Increase Maximum Mana +25%"},
		{"item_goldbonus", 25, "+25% Extra Gold from Monsters"},
		{"item_magicbonus", 25, "+25% Better Chance of Getting Magic Items"},
		{"item_reducedprices", 25, "Reduces all Vendor Prices +25%"},
		{"item_poisonlengthresist", 25, "Poison Length Reduced by +25%"},
		{"item_damagetomana", 25, "+25% Damage Taken Goes To Mana"},
		{"item_tohit_percent", 25, "+25% Bonus to Attack Rating"},
		{"item_openwounds", 25, "+25% Chance of Open Wounds"},
		{"item_crushingblow", 25, "+25% Chance of Crushing Blow"},
		{"item_deadlystrike", 25, "+25% Deadly Strike"},
		{"item_absorbfire_percent", 25, "Fire Absorb +25%"},
		{"item_absorblight_percent", 25, "Lightning Absorb +25%"},
		{"item_absorbmagic_percent", 25, "Magic Absorb +25%"},
		{"item_absorbcold_percent", 25, "Cold Absorb +25%"},
		{"item_slow", 25, "Slows Target by +25%"},
		{"item_staminadrainpct", 25, "+25% Slower Stamina Drain"},

		// would be cool to fix the wording on some of these in the future
		// the negatives don't seen to use different strings table entries
		// better --> worse
		// reduced --> increased
		// increased --> reduced
		{"toblock", -25, "-25% Increased Chance of Blocking"},
		{"manarecoverybonus", -25, "Regenerate Mana -25%"},
		{"staminarecoverybonus", -25, "Heal Stamina Plus -25%"},
		{"damageresist", -25, "Damage Reduced by -25%"},
		{"lifedrainmindam", -25, "-25% Life stolen per hit"},
		{"manadrainmindam", -25, "-25% Mana stolen per hit"},
		{"item_maxdurability_percent", -25, "Increase Maximum Durability -25%"},
		{"item_maxhp_percent", -25, "Increase Maximum Life -25%"},
		{"item_maxmana_percent", -25, "Increase Maximum Mana -25%"},
		{"item_goldbonus", -25, "-25% Extra Gold from Monsters"},
		{"item_magicbonus", -25, "-25% Better Chance of Getting Magic Items"},
		{"item_reducedprices", -25, "Reduces all Vendor Prices -25%"},
		{"item_poisonlengthresist", -25, "Poison Length Reduced by -25%"},
		{"item_damagetomana", -25, "-25% Damage Taken Goes To Mana"},
		{"item_tohit_percent", -25, "-25% Bonus to Attack Rating"},
		{"item_openwounds", -25, "-25% Chance of Open Wounds"},
		{"item_crushingblow", -25, "-25% Chance of Crushing Blow"},
		{"item_deadlystrike", -25, "-25% Deadly Strike"},
		{"item_absorbfire_percent", -25, "Fire Absorb -25%"},
		{"item_absorblight_percent", -25, "Lightning Absorb -25%"},
		{"item_absorbmagic_percent", -25, "Magic Absorb -25%"},
		{"item_absorbcold_percent", -25, "Cold Absorb -25%"},
		{"item_slow", -25, "Slows Target by -25%"},
		{"item_staminadrainpct", -25, "-25% Slower Stamina Drain"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn3(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_maxdamage_percent", 25, "Enhanced Maximum Damage"},
		{"item_mindamage_percent", 25, "Enhanced Minimum Damage"},
		{"normal_damage_reduction", 25, "Damage Reduced by 25"},
		{"magic_damage_reduction", 25, "Magic Damage Reduced by 25"},
		{"item_attackertakesdamage", 25, "Attacker Takes Damage of 25"},
		{"item_knockback", 25, "Knockback"},
		{"item_restinpeace", 25, "Slain Monsters Rest in Peace"},
		{"item_ignoretargetac", 25, "Ignore Target's Defense"},
		{"item_preventheal", 25, "Prevent Monster Heal"},
		{"item_halffreezeduration", 25, "Half Freeze Duration"},
		{"item_throwable", 25, "Throwable"},
		{"item_attackertakeslightdamage", 25, "Attacker Takes Lightning Damage of 25"},
		{"item_indesctructible", 25, "Indestructible"},
		{"item_cannotbefrozen", 25, "Cannot Be Frozen"},
		{"item_pierce", 25, "Piercing Attack"},
		{"item_magicarrow", 25, "Fires Magic Arrows"},
		{"item_explosivearrow", 25, "Fires Explosive Arrows or Bolts"},
		{"item_replenish_quantity", 25, "Replenishes quantity"},
		{"item_extra_stack", 25, "Increased Stack Size"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn4(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_armor_percent", 25, "+25% Enhanced Defense"},
		{"magicresist", 25, "Magic Resist +25%"},
		{"maxmagicresist", 25, "+25% to Maximum Magic Resist"},
		{"fireresist", 25, "Fire Resist +25%"},
		{"maxfireresist", 25, "+25% to Maximum Fire Resist"},
		{"lightresist", 25, "Lightning Resist +25%"},
		{"maxlightresist", 25, "+25% to Maximum Lightning Resist"},
		{"coldresist", 25, "Cold Resist +25%"},
		{"maxcoldresist", 25, "+25% to Maximum Cold Resist"},
		{"poisonresist", 25, "Poison Resist +25%"},
		{"maxpoisonresist", 25, "+25% to Maximum Poison Resist"},
		{"item_addexperience", 25, "+25% to Experience Gained"},
		{"item_req_percent", 25, "Requirements +25%"},
		{"item_fasterattackrate", 25, "+25% Increased Attack Speed"},
		{"item_fastermovevelocity", 25, "+25% Faster Run/Walk"},
		{"item_fastergethitrate", 25, "+25% Faster Hit Recovery"},
		{"item_fasterblockrate", 25, "+25% Faster Block Rate"},
		{"item_fastercastrate", 25, "+25% Faster Cast Rate"},
		{"item_demondamage_percent", 25, "+25% Damage to Demons"},
		{"item_undeaddamage_percent", 25, "+25% Damage to Undead"},
		{"passive_fire_mastery", 25, "+25% to Fire Skill Damage"},
		{"passive_ltng_mastery", 25, "+25% to Lightning Skill Damage"},
		{"passive_cold_mastery", 25, "+25% to Cold Skill Damage"},
		{"passive_pois_mastery", 25, "+25% to Poison Skill Damage"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn5(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_howl", 25, "Hit Causes Monster to Flee 25%"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn6(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_armor_perlevel", 25, "+25 Defense (Based on Character Level)"},
		{"item_hp_perlevel", 25, "+25 to Life (Based on Character Level)"},
		{"item_mana_perlevel", 25, "+25 to Mana (Based on Character Level)"},
		{"item_maxdamage_perlevel", 25, "+25 to Maximum Damage (Based on Character Level)"},
		{"item_strength_perlevel", 25, "+25 to Strength (Based on Character Level)"},
		{"item_dexterity_perlevel", 25, "+25 to Dexterity (Based on Character Level)"},
		{"item_energy_perlevel", 25, "+25 to Energy (Based on Character Level)"},
		{"item_vitality_perlevel", 25, "+25 to Vitality (Based on Character Level)"},
		{"item_tohit_perlevel", 25, "+25 to Attack Rating (Based on Character Level)"},
		{"item_cold_damagemax_perlevel", 25, "+25 to Maximum Cold Damage (Based on Character Level)"},
		{"item_fire_damagemax_perlevel", 25, "+25 to Maximum Fire Damage (Based on Character Level)"},
		{"item_ltng_damagemax_perlevel", 25, "+25 to Maximum Lightning Damage (Based on Character Level)"},
		{"item_pois_damagemax_perlevel", 25, "+25 to Maximum Poison Damage (Based on Character Level)"},
		{"item_absorb_cold_perlevel", 25, "+25 Absorbs Cold Damage (Based on Character Level)"},
		{"item_absorb_fire_perlevel", 25, "+25 Absorbs Fire Damage (Based on Character Level)"},
		{"item_absorb_ltng_perlevel", 25, "+25 Absorbs Lightning Damage (Based on Character Level)"},
		{"item_stamina_perlevel", 25, "+25 Maximum Stamina (Based on Character Level)"},
		{"item_tohit_demon_perlevel", 25, "+25 to Attack Rating against Demons (Based on Character Level)"},
		{"item_tohit_undead_perlevel", 25, "+25 to Attack Rating against Undead (Based on Character Level)"},
		{"item_kick_damage_perlevel", 25, "+25 Kick Damage (Based on Character Level)"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn7(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_tohitpercent_perlevel", 25, "+25% Bonus to Attack Rating (Based on Character Level)"},
		{"item_resist_cold_perlevel", 25, "Cold Resist +25% (Based on Character Level)"},
		{"item_resist_fire_perlevel", 25, "Fire Resist +25% (Based on Character Level)"},
		{"item_resist_ltng_perlevel", 25, "Lightning Resist +25% (Based on Character Level)"},
		{"item_resist_pois_perlevel", 25, "Poison Resist +25% (Based on Character Level)"},
		{"item_find_gold_perlevel", 25, "+25% Extra Gold from Monsters (Based on Character Level)"},
		{"item_find_magic_perlevel", 25, "+25% Better Chance of Getting Magic Items (Based on Character Level)"},
		{"item_crushingblow_perlevel", 25, "+25% Chance of Crushing Blow (Based on Character Level)"},
		{"item_openwounds_perlevel", 25, "+25% Chance of Open Wounds (Based on Character Level)"},
		{"item_deadlystrike_perlevel", 25, "+25% Deadly Strike (Based on Character Level)"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn8(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_armorpercent_perlevel", 25, "+25% Enhanced Defense (Based on Character Level)"},
		{"item_maxdamage_percent_perlevel", 25, "+25% Enhanced Maximum Damage (" +
			"Based on Character Level)"},
		{"item_regenstamina_perlevel", 25, "Heal Stamina Plus +25% (Based on Character Level)"},
		{"item_damage_demon_perlevel", 25, "+25% Damage to Demons (Based on Character Level)"},
		{"item_damage_undead_perlevel", 25, "+25% Damage to Undead (Based on Character Level)"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn9(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_thorns_perlevel", 25, "Attacker Takes Damage of 25 (Based on Character Level)"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn11(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_replenish_durability", 2, "Repairs 2 durability per second"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn12(t *testing.T) {
	tests := []struct {
		recordKey string
		val       int
		expect    string
	}{
		{"item_stupidity", 5, "Hit Blinds Target +5"},
		{"item_freeze", 5, "Freezes target +5"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		val := test.val
		expect := test.expect
		stat := CreateStat(record, val)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, val, expect, got)
		}
	}
}

func TestStat_DescriptionFn13(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{
			"item_addclassskills",
			[]int{5, 1},
			"+5 to to Barbarian Skill Levels",
		},
		{
			"item_addclassskills",
			[]int{5, 2},
			"+5 to to Necromancer Skill Levels",
		},
		{
			"item_addclassskills",
			[]int{5, 3},
			"+5 to to Paladin Skill Levels",
		},
		{
			"item_addclassskills",
			[]int{5, 4},
			"+5 to to Assassin Skills",
		},
		{
			"item_addclassskills",
			[]int{5, 5},
			"+5 to to Sorceress Skill Levels",
		},
		{
			"item_addclassskills",
			[]int{5, 6},
			"+5 to to Amazon Skill Levels",
		},
		{
			"item_addclassskills",
			[]int{5, 7},
			"+5 to to Druid Skills",
		},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}

func TestStat_DescriptionFn14(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{
			"item_addskill_tab",
			[]int{12, 1, 0},
			"+12 to Combat Skills (Barbarian Only)",
		},
		{
			"item_addskill_tab",
			[]int{-6, 1, 1},
			"-6 to Masteries (Barbarian Only)",
		},
		{
			"item_addskill_tab",
			[]int{3, 1, 2},
			"+3 to Warcries (Barbarian Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 2, 0},
			"+5 to Curses (Necromancer Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 2, 1},
			"+5 to Poison and Bone Skills (Necromancer Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 2, 2},
			"+5 to Summoning Skills (Necromancer Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 3, 0},
			"+5 to Combat Skills (Paladin Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 3, 1},
			"+5 to Offensive Auras (Paladin Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 3, 2},
			"+5 to Defensive Auras (Paladin Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 4, 0},
			"+5 to Traps (Assassin Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 4, 1},
			"+5 to Shadow Disciplines (Assassin Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 4, 2},
			"+5 to Martial Arts (Assassin Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 5, 0},
			"+5 to Fire Skills (Sorceress Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 5, 1},
			"+5 to Lightning Skills (Sorceress Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 5, 2},
			"+5 to Cold Skills (Sorceress Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 6, 0},
			"+5 to Bow and Crossbow Skills (Amazon Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 6, 1},
			"+5 to Passive and Magic Skills (Amazon Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 6, 2},
			"+5 to Javelin and Spear Skills (Amazon Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 7, 0},
			"+5 to Summoning Skills (Druid Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 7, 1},
			"+5 to Shape Shifting Skills (Druid Only)",
		},
		{
			"item_addskill_tab",
			[]int{5, 7, 2},
			"+5 to Elemental Skills (Druid Only)",
		},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}

func TestStat_DescriptionFn15(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{"item_skillonattack", []int{5, 7, 260}, "5% Chance to cast level 7 Dragon Claw on attack"},
		{"item_skillonkill", []int{5, 7, 261}, "5% Chance to cast level 7 Charged Bolt Sentry when you Kill an Enemy"},
		{"item_skillondeath", []int{5, 7, 262}, "5% Chance to cast level 7 Wake of Fire Sentry when you Die"},
		{"item_skillonhit", []int{5, 7, 263}, "5% Chance to cast level 7 Weapon Block on striking"},
		{"item_skillonlevelup", []int{5, 7, 264}, "5% Chance to cast level 7 Cloak of Shadows when you Level-Up"},
		{"item_skillongethit", []int{5, 7, 265}, "5% Chance to cast level 7 Cobra Strike when struck"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}

func TestStat_DescriptionFn16(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{"item_aura", []int{3, 37}, "Level 3 Warmth Aura When Equipped"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}

func TestStat_DescriptionFn20(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{"item_fractionaltargetac", []int{-25}, "-25% Target Defense"},
		{"item_pierce_cold", []int{-25}, "-25% to Enemy Cold Resistance"},
		{"item_pierce_fire", []int{-25}, "-25% to Enemy Fire Resistance"},
		{"item_pierce_ltng", []int{-25}, "-25% to Enemy Lightning Resistance"},
		{"item_pierce_pois", []int{-25}, "-25% to Enemy Poison Resistance"},
		{"passive_fire_pierce", []int{-25}, "-25% to Enemy Fire Resistance"},
		{"passive_ltng_pierce", []int{-25}, "-25% to Enemy Lightning Resistance"},
		{"passive_cold_pierce", []int{-25}, "-25% to Enemy Cold Resistance"},
		{"passive_pois_pierce", []int{-25}, "-25% to Enemy Poison Resistance"},

	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}

func TestStat_DescriptionFn22(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{"attack_vs_montype", []int{25, 40}, "25% to Attack Rating versus Specter"},
		{"damage_vs_montype", []int{25, 41}, "25% to Damage versus Apparition"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}

func TestStat_DescriptionFn23(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{"item_reanimate", []int{25, 40}, "25% Reanimate as: Specter"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}

func TestStat_DescriptionFn24(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{"item_charged_skill", []int{25, 40, 20, 19}, "Level 25 Frozen Armor (19/20 Charges)"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}

func TestStat_DescriptionFn27(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{"item_singleskill", []int{25, 40, 5}, "+25 to Frozen Armor (Sorceress Only)"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}

func TestStat_DescriptionFn28(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []int
		expect    string
	}{
		{"item_nonclassskill", []int{25, 64}, "+25 to Frozen Orb"},
	}

	for idx := range tests {
		test := tests[idx]
		record := d2datadict.ItemStatCosts[test.recordKey]
		expect := test.expect
		stat := CreateStat(record, test.vals...)

		if got := stat.Description(); got != expect {
			t.Errorf(errFmt, errStr, test.recordKey, test.vals, expect, got)
		}
	}
}
