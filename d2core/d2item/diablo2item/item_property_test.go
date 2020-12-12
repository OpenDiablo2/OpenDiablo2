package diablo2item

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

// nolint:gochecknoglobals // just a test
var itemStatCosts = map[string]*d2records.ItemStatCostRecord{
	"strength": {
		Name:       "strength",
		DescFnID:   1,
		DescVal:    1,
		DescStrPos: "to Strength",
		DescStrNeg: "to Strength",
	},
	"dexterity": {
		Name:       "dexterity",
		DescFnID:   1,
		DescVal:    1,
		DescStrPos: "to Dexterity",
		DescStrNeg: "to Dexterity",
	},
	"vitality": {
		Name:       "vitality",
		DescFnID:   1,
		DescVal:    1,
		DescStrPos: "to Vitality",
		DescStrNeg: "to Vitality",
	},
	"energy": {
		Name:       "energy",
		DescFnID:   1,
		DescVal:    1,
		DescStrPos: "to Energy",
		DescStrNeg: "to Energy",
	},
	"hpregen": {
		Name:       "hpregen",
		DescFnID:   1,
		DescVal:    2,
		DescStrPos: "Replenish Life",
		DescStrNeg: "Drain Life",
	},
	"toblock": {
		Name:       "toblock",
		DescFnID:   2,
		DescVal:    1,
		DescStrPos: "Increased Chance of Blocking",
		DescStrNeg: "Increased Chance of Blocking",
	},
	"item_absorblight_percent": {
		Name:       "item_absorblight_percent",
		DescFnID:   2,
		DescVal:    2,
		DescStrPos: "Lightning Absorb",
		DescStrNeg: "Lightning Absorb",
	},
	"item_maxdurability_percent": {
		Name:       "item_maxdurability_percent",
		DescFnID:   2,
		DescVal:    2,
		DescStrPos: "Increase Maximum Durability",
		DescStrNeg: "Increase Maximum Durability",
	},
	"item_restinpeace": {
		Name:       "item_restinpeace",
		DescFnID:   3,
		DescVal:    0,
		DescStrPos: "Slain Monsters Rest in Peace",
		DescStrNeg: "Slain Monsters Rest in Peace",
	},
	"normal_damage_reduction": {
		Name:       "normal_damage_reduction",
		DescFnID:   3,
		DescVal:    2,
		DescStrPos: "Damage Reduced by",
		DescStrNeg: "Damage Reduced by",
	},
	"poisonresist": {
		Name:       "poisonresist",
		DescFnID:   4,
		DescVal:    2,
		DescStrPos: "Poison Resist",
		DescStrNeg: "Poison Resist",
	},
	"item_fastermovevelocity": {
		Name:       "item_fastermovevelocity",
		DescFnID:   4,
		DescVal:    1,
		DescStrPos: "Faster Run/Walk",
		DescStrNeg: "Faster Run/Walk",
	},
	"item_howl": {
		Name:       "item_howl",
		DescFnID:   5,
		DescVal:    2,
		DescStrPos: "Hit Causes Monster to Flee",
		DescStrNeg: "Hit Causes Monster to Flee",
	},
	"item_hp_perlevel": {
		Name:       "item_hp_perlevel",
		DescFnID:   6,
		DescVal:    1,
		DescStrPos: "to Life",
		DescStrNeg: "to Life",
		DescStr2:   "(Based on Character Level)",
	},
	"item_resist_ltng_perlevel": {
		Name:       "item_resist_ltng_perlevel",
		DescFnID:   7,
		DescVal:    2,
		DescStrPos: "Lightning Resist",
		DescStrNeg: "Lightning Resist",
		DescStr2:   "(Based on Character Level)",
	},
	"item_find_magic_perlevel": {
		Name:       "item_find_magic_perlevel",
		DescFnID:   7,
		DescVal:    1,
		DescStrPos: "Better Chance of Getting Magic Items",
		DescStrNeg: "Better Chance of Getting Magic Items",
		DescStr2:   "(Based on Character Level)",
	},
	"item_armorpercent_perlevel": {
		Name:       "item_armorpercent_perlevel",
		DescFnID:   8,
		DescVal:    1,
		DescStrPos: "Enhanced Defense",
		DescStrNeg: "Enhanced Defense",
		DescStr2:   "(Based on Character Level)",
	},
	"item_regenstamina_perlevel": {
		Name:       "item_regenstamina_perlevel",
		DescFnID:   8,
		DescVal:    2,
		DescStrPos: "Heal Stamina Plus",
		DescStrNeg: "Heal Stamina Plus",
		DescStr2:   "(Based on Character Level)",
	},
	"item_thorns_perlevel": {
		Name:       "item_thorns_perlevel",
		DescFnID:   9,
		DescVal:    2,
		DescStrPos: "Attacker Takes Damage of",
		DescStrNeg: "Attacker Takes Damage of",
		DescStr2:   "(Based on Character Level)",
	},
	"item_replenish_durability": {
		Name:       "item_replenish_durability",
		DescFnID:   11,
		DescVal:    1,
		DescStrPos: "Repairs %v durability per second",
		DescStrNeg: "Repairs %v durability per second",
		DescStr2:   "",
	},
	"item_stupidity": {
		Name:       "item_stupidity",
		DescFnID:   12,
		DescVal:    2,
		DescStrPos: "Hit Blinds Target",
		DescStrNeg: "Hit Blinds Target",
	},
	"item_addclassskills": {
		Name:     "item_addclassskills",
		DescFnID: 13,
		DescVal:  1,
	},
	"item_addskill_tab": {
		Name:     "item_addskill_tab",
		DescFnID: 14,
		DescVal:  1,
	},
	"item_skillonattack": {
		Name:       "item_skillonattack",
		DescFnID:   15,
		DescVal:    1,
		DescStrPos: "%d%% Chance to cast level %d %s on attack",
		DescStrNeg: "%d%% Chance to cast level %d %s on attack",
	},
	"item_aura": {
		Name:       "item_aura",
		DescFnID:   16,
		DescVal:    1,
		DescStrPos: "Level %d %s Aura When Equipped",
		DescStrNeg: "Level %d %s Aura When Equipped",
	},
	"item_fractionaltargetac": {
		Name:       "item_fractionaltargetac",
		DescFnID:   20,
		DescVal:    1,
		DescStrPos: "Target Defense",
		DescStrNeg: "Target Defense",
	},
	"attack_vs_montype": {
		Name:       "item_fractionaltargetac",
		DescFnID:   22,
		DescVal:    1,
		DescStrPos: "to Attack Rating versus",
		DescStrNeg: "to Attack Rating versus",
	},
	"item_reanimate": {
		Name:       "item_reanimate",
		DescFnID:   23,
		DescVal:    2,
		DescStrPos: "Reanimate as:",
		DescStrNeg: "Reanimate as:",
	},
	"item_charged_skill": {
		Name:       "item_charged_skill",
		DescFnID:   24,
		DescVal:    2,
		DescStrPos: "(%d/%d Charges)",
		DescStrNeg: "(%d/%d Charges)",
	},
	"item_singleskill": {
		Name:     "item_singleskill",
		DescFnID: 27,
		DescVal:  0,
	},
	"item_nonclassskill": {
		Name:       "item_nonclassskill",
		DescFnID:   28,
		DescVal:    2,
		DescStrPos: "(%d/%d Charges)",
		DescStrNeg: "(%d/%d Charges)",
	},
	"item_armor_percent": {
		Name:       "item_armor_percent",
		DescFnID:   4,
		DescVal:    1,
		DescStrPos: "Enhanced Defense",
		DescStrNeg: "Enhanced Defense",
	},
	"item_fastercastrate": {
		Name:       "item_fastercastrate",
		DescFnID:   4,
		DescVal:    1,
		DescStrPos: "Faster Cast Rate",
		DescStrNeg: "Faster Cast Rate",
	},
	"item_skillonlevelup": {
		Name:       "item_skillonlevelup",
		DescFnID:   15,
		DescVal:    0,
		DescStrPos: "%d%% Chance to cast level %d %s when you Level-Up",
		DescStrNeg: "%d%% Chance to cast level %d %s when you Level-Up",
	},
	"item_numsockets": {
		Name: "item_numsockets",
	},
	"poisonmindam": {
		Name:       "poisonmindam",
		DescFnID:   1,
		DescVal:    1,
		DescStrPos: "to Minimum Poison Damage",
		DescStrNeg: "to Minimum Poison Damage",
	},
	"poisonmaxdam": {
		Name:       "poisonmaxdam",
		DescFnID:   1,
		DescVal:    1,
		DescStrPos: "to Maximum Poison Damage",
		DescStrNeg: "to Maximum Poison Damage",
	},
	"poisonlength": {
		Name: "poisonlength",
	},
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
var skillDetails = map[int]*d2records.SkillRecord{
	37: {Skill: "Warmth"},
	64: {Skill: "Frozen Orb"},
}

// nolint:gochecknoglobals // just a test
var monStats = map[string]*d2records.MonStatsRecord{
	"Specter": {NameString: "Specter", ID: 40},
}

// nolint:gochecknoglobals // just a test
var properties = map[string]*d2records.PropertyRecord{
	"allstats": {
		Code: "allstats",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 1, StatCode: "strength"},
			{FunctionID: 3, StatCode: "dexterity"},
			{FunctionID: 3, StatCode: "vitality"},
			{FunctionID: 3, StatCode: "energy"},
		},
	},
	"ac%": {
		Code: "ac%",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 2, StatCode: "item_armor_percent"},
		},
	},
	"dmg-min": {
		Code: "dmg-min",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 5},
		},
	},
	"dmg-max": {
		Code: "dmg-max",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 6},
		},
	},
	"dmg%": {
		Code: "dmg%",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 7},
		},
	},
	"cast1": {
		Code: "cast1",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 8, StatCode: "item_fastercastrate"},
		},
	},
	"skilltab": {
		Code: "skilltab",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 10, StatCode: "item_addskill_tab"},
		},
	},
	"levelup-skill": {
		Code: "levelup-skill",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 11, StatCode: "item_skillonlevelup"},
		},
	},
	"skill-rand": {
		Code: "skill-rand",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 12, StatCode: "item_singleskill"},
		},
	},
	"dur%": {
		Code: "dur%",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 13, StatCode: "item_maxdurability_percent"},
		},
	},
	"sock": {
		Code: "sock",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 14, StatCode: "item_numsockets"},
		},
	},
	"dmg-pois": {
		Code: "dmg-pois",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 15, StatCode: "poisonmindam"},
			{FunctionID: 16, StatCode: "poisonmaxdam"},
			{FunctionID: 17, StatCode: "poisonlength"},
		},
	},
	"charged": {
		Code: "charged",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 19, StatCode: "item_charged_skill"},
		},
	},
	"indestruct": {
		Code: "indestruct",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 20},
		},
	},
	"pal": {
		Code: "pal",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 21, StatCode: "item_addclassskills", Value: 3},
		},
	},
	"oskill": {
		Code: "oskill",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 22, StatCode: "item_nonclassskill"},
		},
	},
	"ethereal": {
		Code: "ethereal",
		Stats: [7]*d2records.PropertyStatRecord{
			{FunctionID: 23},
		},
	},
}

// nolint:gochecknoglobals // just a test
var testAssetManager *d2asset.AssetManager

// nolint:gochecknoglobals // just a test
var testItemFactory *ItemFactory

func TestSetup(t *testing.T) {
	var err error

	testAssetManager = &d2asset.AssetManager{}
	testAssetManager.Records = &d2records.RecordManager{}

	testItemFactory, err = NewItemFactory(testAssetManager)
	if err != nil {
		t.Error(err)
		return
	}

	testAssetManager.Records.Item.Stats = itemStatCosts
	testAssetManager.Records.Character.Stats = charStats
	testAssetManager.Records.Skill.Details = skillDetails
	testAssetManager.Records.Monster.Stats = monStats
	testAssetManager.Records.Properties = properties
}

func TestNewProperty(t *testing.T) { //nolint:funlen // it's mostly test-case definitions
	rand.Seed(time.Now().UTC().UnixNano())

	tests := []struct {
		propKey        string
		inputValues    []int
		expectNumStats int
		expectStr      []string
	}{
		{ // fnId 1 + 3
			"allstats",
			[]int{1, 10},
			4,
			[]string{
				"+# to Strength",
				"+# to Dexterity",
				"+# to Vitality",
				"+# to Energy",
			},
		},
		{ // fnId 2
			"ac%",
			[]int{1, 10},
			1,
			[]string{"+#% Enhanced Defense"},
		},
		{ // fnId 5
			// dmg-min, dmg-max, dmg%, indestructable, and ethereal dont have stats!
			"dmg-min",
			[]int{1, 10},
			0,
			[]string{""},
		},
		{ // fnId 6
			// dmg-min, dmg-max, dmg%, indestructable, and ethereal dont have stats!
			"dmg-max",
			[]int{1, 10},
			0,
			[]string{""},
		},
		{ // fnId 7
			// dmg-min, dmg-max, dmg%, indestructable, and ethereal dont have stats!
			"dmg%",
			[]int{1, 10},
			0,
			[]string{""},
		},
		{ // fnId 8
			"cast1",
			[]int{1, 10},
			1,
			[]string{"+#% Faster Cast Rate"},
		},
		{
			"indestruct",
			[]int{0, 1},
			0,
			[]string{""},
		},
		{
			"ethereal",
			[]int{0, 1},
			0,
			[]string{""},
		},
		{ // fnId 10
			"skilltab",
			[]int{10, 1, 3},
			1,
			[]string{"+# to Offensive Auras (Paladin Only)"},
		},
		{ // fnId 11
			"levelup-skill",
			[]int{64, 100, 3},
			1,
			[]string{"#% Chance to cast level # Frozen Orb when you Level-Up"},
		},
		{ // fnId 12
			"skill-rand",
			[]int{10, 64, 64},
			1,
			[]string{"+# to Frozen Orb"},
		},
		{ // fnId 13
			"dur%",
			[]int{1, 10},
			1,
			[]string{"Increase Maximum Durability +#%"},
		},
		{ // fnId 14
			"sock",
			[]int{0, 6},
			1,
			[]string{""},
		},
		{ // fnId 15, 16, 17
			"dmg-pois",
			[]int{100, 5, 10},
			3,
			[]string{
				"+# to Minimum Poison Damage",
				"+# to Maximum Poison Damage",
				"", // length, non-printing
			},
		},
		{ // fnId 19
			"charged",
			[]int{64, 20, 10},
			1,
			[]string{"Level # Frozen Orb (#/# Charges)"},
		},
		{ // fnId 21
			"pal",
			[]int{1, 5},
			1,
			[]string{"+# to Paladin Skill Levels"},
		},
		{ // fnId 22
			"oskill",
			[]int{64, 1, 5},
			1,
			[]string{"+# to Frozen Orb"},
		},
	}

	numericToken := "#"
	re := regexp.MustCompile(`\d+`)

	for testIdx := range tests {
		test := &tests[testIdx]
		prop := testItemFactory.NewProperty(test.propKey, test.inputValues...)

		if prop == nil {
			t.Error("property is nil")
			continue
		}

		infoFmt := "\r\nProperty `%s`, arguments %v"
		infoStr := fmt.Sprintf(infoFmt, prop.record.Code, test.inputValues)
		fmt.Println(infoStr)

		if len(prop.stats) != test.expectNumStats {
			errFmt := "unexpected property stat count: want %v, have %v"
			t.Errorf(errFmt, test.expectNumStats, len(prop.stats))

			continue
		}

		switch prop.PropertyType {
		case PropertyComputeBoolean:
			fmtStr := "\tGot: [Non-printing boolean property] [Bool Value: %v]"
			got := fmt.Sprintf(fmtStr, prop.computedBool)
			fmt.Println(got)
		case PropertyComputeInteger:
			fmtStr := "\tGot: [Non-printing integer property] [Int Value: %v]"
			got := fmt.Sprintf(fmtStr, prop.computedInt)
			fmt.Println(got)
		case PropertyComputeStats:
			for statIdx := range prop.stats {
				stat := prop.stats[statIdx]
				expectStr := test.expectStr[statIdx]
				statStr := stat.String()
				stripped := string(re.ReplaceAll([]byte(statStr), []byte(numericToken)))

				if expectStr == "" {
					statFmt := "[Non-printing stat] Code: %v, inputValues: %+v"

					vals := stat.Values()
					valInts := make([]int, len(vals))

					for idx := range vals {
						valInts[idx] = vals[idx].Int()
					}

					statStr = fmt.Sprintf(statFmt, stat.Name(), valInts)
					got := fmt.Sprintf("\tGot: %s", statStr)
					fmt.Println(got)
				} else {
					got := fmt.Sprintf("\tGot: %s", statStr)
					fmt.Println(got)
				}

				if stripped != expectStr {
					expected := fmt.Sprintf("\tExpected: %s", test.expectStr)
					t.Error(expected)
				}
			}
		}
	}
}
