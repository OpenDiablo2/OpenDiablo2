package diablo2stats

import (
	"fmt"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

const (
	errStr string = "incorrect description string format for stat"
	errFmt string = "%v:\n\tDescFnID: %v\n\tKey: %v\n\tVal: %+v\n\texpected: %v\n\tgot: %v\n\n"
)

// nolint:gochecknoglobals // just a test
var itemStatCosts = map[string]*d2records.ItemStatCostRecord{
	"strength": {
		Name:       "strength",
		DescFnID:   1,
		DescVal:    int(descValPrefix),
		DescStrPos: "to Strength",
		DescStrNeg: "to Strength",
	},
	"dexterity": {
		Name:       "dexterity",
		DescFnID:   1,
		DescVal:    int(descValPrefix),
		DescStrPos: "to Dexterity",
		DescStrNeg: "to Dexterity",
	},
	"vitality": {
		Name:       "vitality",
		DescFnID:   1,
		DescVal:    int(descValPrefix),
		DescStrPos: "to Vitality",
		DescStrNeg: "to Vitality",
	},
	"energy": {
		Name:       "energy",
		DescFnID:   1,
		DescVal:    int(descValPrefix),
		DescStrPos: "to Energy",
		DescStrNeg: "to Energy",
	},
	"hpregen": {
		Name:       "hpregen",
		DescFnID:   1,
		DescVal:    int(descValPostfix),
		DescStrPos: "Replenish Life",
		DescStrNeg: "Drain Life",
	},
	"toblock": {
		Name:       "toblock",
		DescFnID:   2,
		DescVal:    int(descValPrefix),
		DescStrPos: "Increased Chance of Blocking",
		DescStrNeg: "Increased Chance of Blocking",
	},
	"item_absorblight_percent": {
		Name:       "item_absorblight_percent",
		DescFnID:   2,
		DescVal:    int(descValPostfix),
		DescStrPos: "Lightning Absorb",
		DescStrNeg: "Lightning Absorb",
	},
	"item_restinpeace": {
		Name:       "item_restinpeace",
		DescFnID:   3,
		DescVal:    int(descValHide),
		DescStrPos: "Slain Monsters Rest in Peace",
		DescStrNeg: "Slain Monsters Rest in Peace",
	},
	"normal_damage_reduction": {
		Name:       "normal_damage_reduction",
		DescFnID:   3,
		DescVal:    int(descValPostfix),
		DescStrPos: "Damage Reduced by",
		DescStrNeg: "Damage Reduced by",
	},
	"poisonresist": {
		Name:       "poisonresist",
		DescFnID:   4,
		DescVal:    int(descValPostfix),
		DescStrPos: "Poison Resist",
		DescStrNeg: "Poison Resist",
	},
	"item_fastermovevelocity": {
		Name:       "item_fastermovevelocity",
		DescFnID:   4,
		DescVal:    int(descValPrefix),
		DescStrPos: "Faster Run/Walk",
		DescStrNeg: "Faster Run/Walk",
	},
	"item_howl": {
		Name:       "item_howl",
		DescFnID:   5,
		DescVal:    int(descValPostfix),
		DescStrPos: "Hit Causes Monster to Flee",
		DescStrNeg: "Hit Causes Monster to Flee",
	},
	"item_hp_perlevel": {
		Name:       "item_hp_perlevel",
		DescFnID:   6,
		DescVal:    int(descValPrefix),
		DescStrPos: "to Life",
		DescStrNeg: "to Life",
		DescStr2:   "(Based on Character Level)",
	},
	"item_resist_ltng_perlevel": {
		Name:       "item_resist_ltng_perlevel",
		DescFnID:   7,
		DescVal:    int(descValPostfix),
		DescStrPos: "Lightning Resist",
		DescStrNeg: "Lightning Resist",
		DescStr2:   "(Based on Character Level)",
	},
	"item_find_magic_perlevel": {
		Name:       "item_find_magic_perlevel",
		DescFnID:   7,
		DescVal:    int(descValPrefix),
		DescStrPos: "Better Chance of Getting Magic Items",
		DescStrNeg: "Better Chance of Getting Magic Items",
		DescStr2:   "(Based on Character Level)",
	},
	"item_armorpercent_perlevel": {
		Name:       "item_armorpercent_perlevel",
		DescFnID:   8,
		DescVal:    int(descValPrefix),
		DescStrPos: "Enhanced Defense",
		DescStrNeg: "Enhanced Defense",
		DescStr2:   "(Based on Character Level)",
	},
	"item_regenstamina_perlevel": {
		Name:       "item_regenstamina_perlevel",
		DescFnID:   8,
		DescVal:    int(descValPostfix),
		DescStrPos: "Heal Stamina Plus",
		DescStrNeg: "Heal Stamina Plus",
		DescStr2:   "(Based on Character Level)",
	},
	"item_thorns_perlevel": {
		Name:       "item_thorns_perlevel",
		DescFnID:   9,
		DescVal:    int(descValPostfix),
		DescStrPos: "Attacker Takes Damage of",
		DescStrNeg: "Attacker Takes Damage of",
		DescStr2:   "(Based on Character Level)",
	},
	"item_replenish_durability": {
		Name:       "item_replenish_durability",
		DescFnID:   11,
		DescVal:    int(descValPrefix),
		DescStrPos: "Repairs %v durability per second",
		DescStrNeg: "Repairs %v durability per second",
		DescStr2:   "",
	},
	"item_stupidity": {
		Name:       "item_stupidity",
		DescFnID:   12,
		DescVal:    int(descValPostfix),
		DescStrPos: "Hit Blinds Target",
		DescStrNeg: "Hit Blinds Target",
	},
	"item_addclassskills": {
		Name:     "item_addclassskills",
		DescFnID: 13,
		DescVal:  int(descValPrefix),
	},
	"item_addskill_tab": {
		Name:     "item_addskill_tab",
		DescFnID: 14,
		DescVal:  int(descValPrefix),
	},
	"item_skillonattack": {
		Name:       "item_skillonattack",
		DescFnID:   15,
		DescVal:    int(descValPrefix),
		DescStrPos: "%d%% Chance to cast level %d %s on attack",
		DescStrNeg: "%d%% Chance to cast level %d %s on attack",
	},
	"item_aura": {
		Name:       "item_aura",
		DescFnID:   16,
		DescVal:    int(descValPrefix),
		DescStrPos: "Level %d %s Aura When Equipped",
		DescStrNeg: "Level %d %s Aura When Equipped",
	},
	"item_fractionaltargetac": {
		Name:       "item_fractionaltargetac",
		DescFnID:   20,
		DescVal:    int(descValPrefix),
		DescStrPos: "Target Defense",
		DescStrNeg: "Target Defense",
	},
	"attack_vs_montype": {
		Name:       "item_fractionaltargetac",
		DescFnID:   22,
		DescVal:    int(descValPrefix),
		DescStrPos: "to Attack Rating versus",
		DescStrNeg: "to Attack Rating versus",
	},
	"item_reanimate": {
		Name:       "item_reanimate",
		DescFnID:   23,
		DescVal:    int(descValPostfix),
		DescStrPos: "Reanimate as:",
		DescStrNeg: "Reanimate as:",
	},
	"item_charged_skill": {
		Name:       "item_charged_skill",
		DescFnID:   24,
		DescVal:    int(descValPostfix),
		DescStrPos: "(%d/%d Charges)",
		DescStrNeg: "(%d/%d Charges)",
	},
	"item_singleskill": {
		Name:       "item_singleskill",
		DescFnID:   27,
		DescVal:    int(descValPostfix),
		DescStrPos: "(%d/%d Charges)",
		DescStrNeg: "(%d/%d Charges)",
	},
	"item_nonclassskill": {
		Name:       "item_nonclassskill",
		DescFnID:   28,
		DescVal:    int(descValPostfix),
		DescStrPos: "(%d/%d Charges)",
		DescStrNeg: "(%d/%d Charges)",
	},
}

// nolint:gochecknoglobals // just a test
var skillDetails = map[int]*d2records.SkillRecord{
	37: {Skill: "Warmth"},
	64: {Skill: "Frozen Orb"},
}

// nolint:gochecknoglobals // just a test
var monStats = map[string]*d2records.MonStatsRecord{
	"Specter": {NameString: "Specter", ID: 40},
}

// nolint:gochecknoglobals // just a test
var charStats = map[d2enum.Hero]*d2records.CharStatsRecord{
	d2enum.HeroPaladin: {
		Class:             d2enum.HeroPaladin,
		SkillStrAll:       "to Paladin Skill Levels",
		SkillStrClassOnly: "(Paladin Only)",
		SkillStrTab: [3]string{
			"+%d to Combat Skills",
			"+%d to Offensive Auras",
			"+%d to Defensive Auras",
		},
	},
}

// nolint:gochecknoglobals // just a test
var testAssetManager2 *d2asset.AssetManager

// nolint:gochecknoglobals // just a test
var testStatFactory2 *StatFactory

func TestSetup_Stat(t *testing.T) {
	testAssetManager2 = &d2asset.AssetManager{}
	testAssetManager2.Records = &d2records.RecordManager{}

	testStatFactory2, _ = NewStatFactory(testAssetManager2)

	testAssetManager2.Records.Item.Stats = itemStatCosts
	testAssetManager2.Records.Character.Stats = charStats
	testAssetManager2.Records.Skill.Details = skillDetails
	testAssetManager2.Records.Monster.Stats = monStats
}

func TestStat_Clone(t *testing.T) {
	s1 := testStatFactory2.NewStat("strength", 5)
	s2 := s1.Clone()

	// make sure the stats are distinct
	if &s1 == &s2 {
		t.Errorf("stats share the same pointer %d == %d", &s1, &s2)
	}

	// make sure the stat values are unique
	vs1, vs2 := s1.Values(), s2.Values()
	if &vs1 == &vs2 {
		t.Errorf("stat values share the same pointer %d == %d", &s1, &s2)
	}

	s2.Values()[0].SetInt(6)
	v1, v2 := s1.Values()[0].Int(), s2.Values()[0].Int()

	// make sure the value ranges are distinct
	if v1 == v2 {
		t.Errorf("clones should not share stat values")
	}
}

func TestStat_Descriptions(t *testing.T) {
	tests := []struct {
		recordKey string
		vals      []float64
		expect    string
	}{
		// DescFn1
		{"strength", []float64{31}, "+31 to Strength"},
		{"hpregen", []float64{20}, "Replenish Life +20"},
		{"hpregen", []float64{-8}, "Drain Life -8"},

		// DescFn2
		{"toblock", []float64{16}, "+16% Increased Chance of Blocking"},
		{"item_absorblight_percent", []float64{10}, "Lightning Absorb +10%"},

		// DescFn3
		{"normal_damage_reduction", []float64{25}, "Damage Reduced by 25"},
		{"item_restinpeace", []float64{25}, "Slain Monsters Rest in Peace"},

		// DescFn4
		{"poisonresist", []float64{25}, "Poison Resist +25%"},
		{"item_fastermovevelocity", []float64{25}, "+25% Faster Run/Walk"},

		// DescFn5
		{"item_howl", []float64{25}, "Hit Causes Monster to Flee 25%"},

		// DescFn6
		{"item_hp_perlevel", []float64{25}, "+25 to Life (Based on Character Level)"},

		// DescFn7
		{"item_resist_ltng_perlevel", []float64{25},
			"Lightning Resist +25% (Based on Character Level)"},
		{"item_find_magic_perlevel", []float64{25}, "+25% Better Chance of Getting Magic Items (" +
			"Based on Character Level)"},

		// DescFn8
		{"item_armorpercent_perlevel", []float64{25},
			"+25% Enhanced Defense (Based on Character Level)"},
		{"item_regenstamina_perlevel", []float64{25},
			"Heal Stamina Plus +25% (Based on Character Level)"},

		// DescFn9
		{"item_thorns_perlevel", []float64{25}, "Attacker Takes Damage of 25 (" +
			"Based on Character Level)"},

		// DescFn11
		{"item_replenish_durability", []float64{2}, "Repairs 2 durability per second"},

		// DescFn12
		{"item_stupidity", []float64{5}, "Hit Blinds Target +5"},

		// DescFn13
		{"item_addclassskills", []float64{5, 3}, "+5 to Paladin Skill Levels"},

		// DescFn14
		{"item_addskill_tab", []float64{5, 3, 0}, "+5 to Combat Skills (Paladin Only)"},
		{"item_addskill_tab", []float64{5, 3, 1}, "+5 to Offensive Auras (Paladin Only)"},
		{"item_addskill_tab", []float64{5, 3, 2}, "+5 to Defensive Auras (Paladin Only)"},

		// DescFn15
		{"item_skillonattack", []float64{5, 7, 64},
			"5% Chance to cast level 7 Frozen Orb on attack"},

		// DescFn16
		{"item_aura", []float64{3, 37}, "Level 3 Warmth Aura When Equipped"},

		// DescFn20
		{"item_fractionaltargetac", []float64{-25}, "-25% Target Defense"},

		// DescFn22
		{"attack_vs_montype", []float64{25, 40}, "25% to Attack Rating versus Specter"},

		// DescFn23
		{"item_reanimate", []float64{25, 40}, "25% Reanimate as: Specter"},

		// DescFn24
		{"item_charged_skill", []float64{25, 64, 20, 19}, "Level 25 Frozen Orb (19/20 Charges)"},

		// DescFn27
		{"item_singleskill", []float64{25, 64, 3}, "+25 to Frozen Orb (Paladin Only)"},

		// DescFn28
		{"item_nonclassskill", []float64{25, 64}, "+25 to Frozen Orb"},
	}

	for idx := range tests {
		test := tests[idx]
		key := test.recordKey
		record := itemStatCosts[key]
		expect := test.expect
		stat := testStatFactory2.NewStat(key, test.vals...)

		if got := stat.String(); got != expect {
			t.Errorf(errFmt, errStr, record.DescFnID, test.recordKey, test.vals, expect, got)
		} else {
			success := "[Desc Func %d][%s %+v] %s"
			success = fmt.Sprintf(success, record.DescFnID, record.Name, test.vals, got)
			fmt.Println(success)
		}
	}
}

func TestDiablo2Stat_Combine(t *testing.T) {
	a := testStatFactory2.NewStat("item_nonclassskill", 25, 64) // "+25 to Frozen Orb"
	b := testStatFactory2.NewStat("item_nonclassskill", 5, 64)  // "+5 to Frozen Orb"

	c, err := a.Combine(b)

	if err != nil || c.String() != "+30 to Frozen Orb" {
		t.Errorf("stats combination failed\r%s", err)
	}

	d := testStatFactory2.NewStat("item_nonclassskill", 5, 37) // "+5 to Warmth"
	_, err = c.Combine(d)

	if err == nil {
		t.Error("stats were combined when they should not have been.")
	}
}
