package d2stats

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// CreateStat creates a stat instance with the given ID and number of values
func CreateStat(record *d2datadict.ItemStatCostRecord, values ...int) *Stat {
	if record == nil {
		return nil
	}

	stat := &Stat{
		Record: record,
		Values: values,
	}

	return stat
}

// Stat is an instance of a Stat, with a set of Values
type Stat struct {
	Record *d2datadict.ItemStatCostRecord
	Values []int
}

// Clone returns a deep copy of the Stat
func (s Stat) Clone() *Stat {
	clone := &Stat{
		Record: s.Record,
		Values: make([]int, len(s.Values)),
	}

	for idx := range s.Values {
		clone.Values[idx] = s.Values[idx]
	}

	return clone
}

// Description returns the formatted description string
func (s *Stat) Description() string {
	return s.DescString(s.Values...)
}

// StatDescriptionFormatStrings is an array of the base format strings used
// by the `descfn` methods for stats. The records in itemstatcost.txt have a
// number field which denotes which of these functions is used for formatting
// the stat description.
// These came from phrozen keep:
// https://d2mods.info/forum/kb/viewarticle?a=448
//nolint:gochecknoglobals // better for lookup
var StatDescriptionFormatStrings = []string{
	"",
	"%v %s",
	"%v%% %s",
	"%v %s",
	"%v%% %s",
	"%v%% %s",
	"%v %s %s",
	"%v%% %v %s",
	"%v%% %s %s",
	"%v %s %s",
	"%v %s %s",
	"Repairs 1 Durability In %v Seconds",
	"%v +%v",
	"+%v to %s",
	"+%v to %s %s",
	"%v%% %s",
	"%v %s",
	"%v %s (Increases near %v)",
	"%v%% %s (Increases near %v)",
	"",
	"%v%% %s",
	"%v %s",
	"%v%% %s %s",
	"%v%% %s %s",
	"Level %v %s %s",
	"",
	"",
	"+%v to %s %s",
	"+%v to %s",
}

var statValueCountLookup map[int]int //nolint:gochecknoglobals // lookup

// DescString return a string based on the DescFnID
func (s *Stat) DescString(values ...int) string {
	if s.Record.DescFnID < 0 || s.Record.DescFnID > len(StatDescriptionFormatStrings) {
		return ""
	}

	var result string
	switch s.Record.DescFnID {
	case 1:
		result = s.descFn1(values...)
	case 2:
		result = s.descFn2(values...)
	case 3:
		result = s.descFn3(values...)
	case 4:
		result = s.descFn4(values...)
	case 5:
		result = s.descFn5(values...)
	case 6:
		result = s.descFn6(values...)
	case 7:
		result = s.descFn7(values...)
	case 8:
		result = s.descFn8(values...)
	case 9:
		result = s.descFn9(values...)
	case 11:
		result = s.descFn11(values...)
	case 12:
		result = s.descFn12(values...)
	case 13:
		result = s.descFn13(values...)
	case 14:
		result = s.descFn14(values...)
	case 15:
		result = s.descFn15(values...)
	case 16:
		result = s.descFn16(values...)
	case 20:
		result = s.descFn20(values...)
	case 22:
		result = s.descFn22(values...)
	case 23:
		result = s.descFn23(values...)
	case 24:
		result = s.descFn24(values...)
	case 27:
		result = s.descFn27(values...)
	case 28:
		result = s.descFn28(values...)
	}

	return result
}

func (s *Stat) descFn1(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = s.Record.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey = s.Record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch s.Record.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	result = strings.ReplaceAll(result, "+-", "-")
	result = strings.ReplaceAll(result, " +%d", "")

	return result
}

func (s *Stat) descFn2(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = s.Record.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey = s.Record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch s.Record.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.ReplaceAll(result, " +%d", "")

	return result
}

func (s *Stat) descFn3(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = s.Record.DescStrNeg
	} else {
		stringTableKey = s.Record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch s.Record.DescVal {
	case 0:
		format = strings.Split(format, " ")[0]
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.ReplaceAll(result, "+-", "-")
	result = strings.ReplaceAll(result, " +%d", "")

	return result
}

func (s *Stat) descFn4(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = s.Record.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey = s.Record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch s.Record.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.ReplaceAll(result, " +%d", "")

	return result
}

func (s *Stat) descFn5(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = s.Record.DescStrNeg
	} else {
		stringTableKey = s.Record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch s.Record.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.ReplaceAll(result, " +%d", "")

	return result
}

func (s *Stat) descFn6(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey1 string
	if value < 0 {
		stringTableKey1 = s.Record.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey1 = s.Record.DescStrPos
	}

	stringTableStr1 := d2common.TranslateString(stringTableKey1)

	// this stat has an additional string (Based on Character Level)
	stringTableStr2 := d2common.TranslateString(s.Record.DescStr2)

	var result string

	switch s.Record.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableStr1, stringTableStr2)
	case 1:
		result = fmt.Sprintf(format, value, stringTableStr1, stringTableStr2)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableStr1, value)
	default:
		result = ""
	}

	// bugs
	result = strings.ReplaceAll(result, "+-", "-")
	result = strings.ReplaceAll(result, " +%d", "")

	return result
}

func (s *Stat) descFn7(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = s.Record.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey = s.Record.DescStrPos
	}

	stringTableStr1 := d2common.TranslateString(stringTableKey)

	// this stat has an additional string (Based on Character Level)
	stringTableStr2 := d2common.TranslateString(s.Record.DescStr2)

	var result string

	switch s.Record.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableStr1, stringTableStr2)
	case 1:
		result = fmt.Sprintf(format, value, stringTableStr1, stringTableStr2)
	case 2:
		formatSplit := strings.Split(format, " ")
		// formatSplit = reverseStringSlice(formatSplit)
		formatSplit[0], formatSplit[1] = formatSplit[1], formatSplit[0]
		format = strings.Join(formatSplit, " ")
		result = fmt.Sprintf(format, stringTableStr1, value, stringTableStr2)
	default:
		result = ""
	}

	return result
}

func (s *Stat) descFn8(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey1 string
	if value < 0 {
		stringTableKey1 = s.Record.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey1 = s.Record.DescStrPos
	}

	stringTableStr1 := d2common.TranslateString(stringTableKey1)

	// this stat has an additional string (Based on Character Level)
	stringTableStr2 := d2common.TranslateString(s.Record.DescStr2)

	var result string

	switch s.Record.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableStr1, stringTableStr2)
	case 1:
		result = fmt.Sprintf(format, value, stringTableStr1, stringTableStr2)
	case 2:
		formatSplit := strings.Split(format, " ")
		formatSplit[0], formatSplit[1] = formatSplit[1], formatSplit[0]
		format = strings.Join(formatSplit, " ")
		result = fmt.Sprintf(format, stringTableStr1, value, stringTableStr2)
	default:
		result = ""
	}

	return result
}

func (s *Stat) descFn9(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey1 string
	if value < 0 {
		stringTableKey1 = s.Record.DescStrNeg
	} else {
		stringTableKey1 = s.Record.DescStrPos
	}

	stringTableStr1 := d2common.TranslateString(stringTableKey1)

	// this stat has an additional string (Based on Character Level)
	stringTableStr2 := d2common.TranslateString(s.Record.DescStr2)

	var result string

	switch s.Record.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableStr1, stringTableStr2)
	case 1:
		result = fmt.Sprintf(format, value, stringTableStr1, stringTableStr2)
	case 2:
		formatSplit := strings.Split(format, " ")
		formatSplit[0], formatSplit[1] = formatSplit[1], formatSplit[0]
		format = strings.Join(formatSplit, " ")
		result = fmt.Sprintf(format, stringTableStr1, value, stringTableStr2)
	default:
		result = ""
	}

	return result
}

func (s *Stat) descFn11(values ...int) string {
	// we know there is only one value for this stat
	value := values[0]

	// the only stat to use this fn is "Repairs durability in X seconds"
	format := d2common.TranslateString(s.Record.DescStrPos)

	return fmt.Sprintf(format, value)
}

func (s *Stat) descFn12(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	str1 := d2common.TranslateString(s.Record.DescStrPos)

	return fmt.Sprintf(format, str1, value)
}

func (s *Stat) descFn13(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]
	numSkills, heroIndex := values[0], values[1]

	heroMap := map[int]d2enum.Hero{
		int(d2enum.HeroAmazon):      d2enum.HeroAmazon,
		int(d2enum.HeroSorceress):   d2enum.HeroSorceress,
		int(d2enum.HeroNecromancer): d2enum.HeroNecromancer,
		int(d2enum.HeroPaladin):     d2enum.HeroPaladin,
		int(d2enum.HeroBarbarian):   d2enum.HeroBarbarian,
		int(d2enum.HeroDruid):       d2enum.HeroDruid,
		int(d2enum.HeroAssassin):    d2enum.HeroAssassin,
	}

	classRecord := d2datadict.CharStats[heroMap[heroIndex]]
	descStr1 := d2common.TranslateString(classRecord.SkillStrAll)
	result := fmt.Sprintf(format, numSkills, descStr1)

	result = strings.ReplaceAll(result, "+-", "-")

	return result
}

func (s *Stat) descFn14(values ...int) string {
	numSkills, heroIndex, skillTabIndex := values[0], values[1], values[2]

	if skillTabIndex > 2 || skillTabIndex < 0 {
		skillTabIndex = 0
	}

	heroMap := map[int]d2enum.Hero{
		int(d2enum.HeroAmazon):      d2enum.HeroAmazon,
		int(d2enum.HeroSorceress):   d2enum.HeroSorceress,
		int(d2enum.HeroNecromancer): d2enum.HeroNecromancer,
		int(d2enum.HeroPaladin):     d2enum.HeroPaladin,
		int(d2enum.HeroBarbarian):   d2enum.HeroBarbarian,
		int(d2enum.HeroDruid):       d2enum.HeroDruid,
		int(d2enum.HeroAssassin):    d2enum.HeroAssassin,
	}

	classRecord := d2datadict.CharStats[heroMap[heroIndex]]
	skillTabKey := classRecord.SkillStrTab[skillTabIndex]
	classOnlyKey := classRecord.SkillStrClassOnly

	skillTabStr := d2common.TranslateString(skillTabKey) + " %v"
	skillTabStr = strings.ReplaceAll(skillTabStr, "%d", "%v")
	classOnlyStr := d2common.TranslateString(classOnlyKey)
	result := fmt.Sprintf(skillTabStr, numSkills, classOnlyStr)

	// bugs
	result = strings.ReplaceAll(result, "+-", "-")

	return result
}

func within(n, min, max int) int {
	if n < min {
		n = min
	} else if n > max {
		n = max
	}

	return n
}

func (s *Stat) descFn15(values ...int) string {
	format := d2common.TranslateString(s.Record.DescStrPos)
	chanceToCast, skillLevel, skillIndex := values[0], values[1], values[2]

	chanceToCast = within(chanceToCast, 0, 100)
	skillLevel = within(skillLevel, 1, 1<<8)
	skillLevel = within(skillLevel, 0, len(d2datadict.SkillDetails)-1)

	skillRecord := d2datadict.SkillDetails[skillIndex]
	//skillDescKey := skillRecord.Skilldesc
	//skillDescRecord := SkillDescriptions[skillDescKey]
	//skillShortName := d2common.TranslateString(skillDescRecord.ShortKey)

	result := fmt.Sprintf(format, chanceToCast, skillLevel, skillRecord.Skill)

	// bugs
	result = strings.ReplaceAll(result, "+-", "-")

	return result
}

func (s *Stat) descFn16(values ...int) string {
	skillLevel, skillIndex := values[0], values[1]

	str1 := d2common.TranslateString(s.Record.DescStrPos)

	skillRecord := d2datadict.SkillDetails[skillIndex]

	result := fmt.Sprintf(str1, skillLevel, skillRecord.Skill)

	// bugs
	result = strings.ReplaceAll(result, "+-", "-")

	return result
}

/*
func (s *Stat) descFn17(values ...int) string {
	// these were not implemented in original D2
	// leaving them out for now as I don't know how to
	// write a test for them, nor do I think vanilla content uses them
	// but these are the stat keys which point to this func...
	// item_armor_bytime
	// item_hp_bytime
	// item_mana_bytime
	// item_maxdamage_bytime
	// item_strength_bytime
	// item_dexterity_bytime
	// item_energy_bytime
	// item_vitality_bytime
	// item_tohit_bytime
	// item_cold_damagemax_bytime
	// item_fire_damagemax_bytime
	// item_ltng_damagemax_bytime
	// item_pois_damagemax_bytime
	// item_stamina_bytime
	// item_tohit_demon_bytime
	// item_tohit_undead_bytime
	// item_kick_damage_bytime
}

func (s *Stat) descFn18(values ...int) string {
	// ... same with these ...
	// item_armorpercent_bytime
	// item_maxdamage_percent_bytime
	// item_tohitpercent_bytime
	// item_resist_cold_bytime
	// item_resist_fire_bytime
	// item_resist_ltng_bytime
	// item_resist_pois_bytime
	// item_absorb_cold_bytime
	// item_absorb_fire_bytime
	// item_absorb_ltng_bytime
	// item_find_gold_bytime
	// item_find_magic_bytime
	// item_regenstamina_bytime
	// item_damage_demon_bytime
	// item_damage_undead_bytime
	// item_crushingblow_bytime
	// item_openwounds_bytime
	// item_deadlystrike_bytime
}
*/

func (s *Stat) descFn20(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = s.Record.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey = s.Record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch s.Record.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.ReplaceAll(result, " +%d", "")

	return result
}

func (s *Stat) descFn22(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]
	statAgainst, monsterIndex := values[0], values[1]

	var monsterKey string

	for key := range d2datadict.MonStats {
		if d2datadict.MonStats[key].Id == monsterIndex {
			monsterKey = key
			break
		}
	}

	str1 := d2common.TranslateString(s.Record.DescStrPos)
	monsterName := d2datadict.MonStats[monsterKey].NameString

	result := fmt.Sprintf(format, statAgainst, str1, monsterName)

	result = strings.ReplaceAll(result, "+-", "-")

	return result
}

func (s *Stat) descFn23(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]
	chanceReanimate, monsterIndex := values[0], values[1]

	var monsterKey string

	for key := range d2datadict.MonStats {
		if d2datadict.MonStats[key].Id == monsterIndex {
			monsterKey = key
			break
		}
	}

	str1 := d2common.TranslateString(s.Record.DescStrPos)
	monsterName := d2datadict.MonStats[monsterKey].NameString

	result := fmt.Sprintf(format, chanceReanimate, str1, monsterName)
	result = strings.ReplaceAll(result, "+-", "-")

	return result
}

func (s *Stat) descFn24(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]
	lvl, skillID, chargeMax, chargeCurrent := values[0], values[1], values[2], values[3]

	charges := d2common.TranslateString(s.Record.DescStrPos)
	charges = fmt.Sprintf(charges, chargeCurrent, chargeMax)

	skillName := d2datadict.SkillDetails[skillID].Skill

	result := fmt.Sprintf(format, lvl, skillName, charges)

	return result
}

func (s *Stat) descFn27(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]
	amount, skillID, heroIndex := values[0], values[1], values[2]

	skillName := d2datadict.SkillDetails[skillID].Skill

	heroMap := map[int]d2enum.Hero{
		int(d2enum.HeroAmazon):      d2enum.HeroAmazon,
		int(d2enum.HeroSorceress):   d2enum.HeroSorceress,
		int(d2enum.HeroNecromancer): d2enum.HeroNecromancer,
		int(d2enum.HeroPaladin):     d2enum.HeroPaladin,
		int(d2enum.HeroBarbarian):   d2enum.HeroBarbarian,
		int(d2enum.HeroDruid):       d2enum.HeroDruid,
		int(d2enum.HeroAssassin):    d2enum.HeroAssassin,
	}

	classRecord := d2datadict.CharStats[heroMap[heroIndex]]
	classOnlyStr := d2common.TranslateString(classRecord.SkillStrClassOnly)

	return fmt.Sprintf(format, amount, skillName, classOnlyStr)
}

func (s *Stat) descFn28(values ...int) string {
	format := StatDescriptionFormatStrings[s.Record.DescFnID]
	amount, skillID := values[0], values[1]

	skillName := d2datadict.SkillDetails[skillID].Skill

	return fmt.Sprintf(format, amount, skillName)
}

// DescGroupString return a string based on the DescGroupFuncID
func (s *Stat) DescGroupString(a ...interface{}) string {
	if s.Record.DescGroupFuncID < 0 || s.Record.DescGroupFuncID > len(
		StatDescriptionFormatStrings) {
		return ""
	}

	format := StatDescriptionFormatStrings[s.Record.DescGroupFuncID]

	return fmt.Sprintf(format, a...)
}

// NumStatValues returns the number of values a stat instance for this
// record should have
func (s *Stat) NumStatValues() int {
	if num, found := statValueCountLookup[s.Record.DescGroupFuncID]; found {
		return num
	}

	if statValueCountLookup == nil {
		statValueCountLookup = make(map[int]int)
	}

	format := StatDescriptionFormatStrings[s.Record.DescGroupFuncID]
	pattern := regexp.MustCompile("%v")
	matches := pattern.FindAllStringIndex(format, -1)
	num := len(matches)
	statValueCountLookup[s.Record.DescGroupFuncID] = num

	return num
}

func reverseStringSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
