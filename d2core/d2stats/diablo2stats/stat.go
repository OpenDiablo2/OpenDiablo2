package diablo2stats

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// static check that Diablo2Stat implements Stat
var _ d2stats.Stat = &Diablo2Stat{}

type descValPosition int

const (
	descValHide descValPosition = iota
	descValPrefix
	descValPostfix
)

const (
	maxSkillTabIndex = 2
)

const (
	twoComponentStr   = "%s %s"
	threeComponentStr = "%s %s %s"
	fourComponentStr  = "%s %s %s %s"
)

// Diablo2Stat is an instance of a Diablo2Stat, with a set of values
type Diablo2Stat struct {
	record *d2datadict.ItemStatCostRecord
	values []d2stats.StatValue
}

// Name returns the name of the stat (the key in itemstatcosts)
func (s *Diablo2Stat) Name() string {
	return s.record.Name
}

// Priority returns the description printing priority
func (s *Diablo2Stat) Priority() int {
	return s.record.DescPriority
}

// Values returns the stat values of the stat
func (s *Diablo2Stat) Values() []d2stats.StatValue {
	return s.values
}

// SetValues sets the stat values
func (s *Diablo2Stat) SetValues(values ...d2stats.StatValue) {
	s.values = make([]d2stats.StatValue, len(values))
	for idx := range values {
		s.values[idx] = values[idx]
	}
}

// Clone returns a deep copy of the Diablo2Stat
func (s *Diablo2Stat) Clone() d2stats.Stat {
	clone := &Diablo2Stat{
		record: s.record,
		values: make([]d2stats.StatValue, len(s.Values())),
	}

	for idx := range s.values {
		srcVal := s.values[idx]
		dstVal := reflect.New(reflect.ValueOf(srcVal).Elem().Type()).Interface().(d2stats.StatValue)

		switch srcVal.Type() {
		case d2stats.StatValueInt:
			dstVal.SetInt(srcVal.Int())
		case d2stats.StatValueFloat:
			dstVal.SetFloat(srcVal.Float())
		}

		clone.values[idx] = dstVal
	}

	return clone
}

// Copy to this stat value the values of the given stat value
func (s *Diablo2Stat) Copy(from d2stats.Stat) d2stats.Stat {
	srcValues := from.Values()
	s.values = make([]d2stats.StatValue, len(srcValues))

	for idx := range srcValues {
		src := srcValues[idx]
		valType := src.Type()
		dst := &Diablo2StatValue{_type: valType}
		dst.SetStringer(src.Stringer())

		switch valType {
		case d2stats.StatValueInt:
			dst.SetInt(src.Int())
		case d2stats.StatValueFloat:
			dst.SetFloat(src.Float())
		}

		s.values[idx] = dst
	}

	return s
}

// Combine sums the other stat with this one (does not alter this stat, returns altered clone!)
func (s *Diablo2Stat) Combine(other d2stats.Stat) (combined d2stats.Stat, err error) {
	cantBeCombinedErr := fmt.Errorf("cannot combine %s with %s", s.Name(), other.Name())

	if !s.canBeCombinedWith(other) {
		return nil, cantBeCombinedErr
	}

	combined = s.Clone()
	srcValues, dstValues := combined.Values(), other.Values()

	for idx := range s.values {
		v1, v2 := dstValues[idx], srcValues[idx]

		valType := v1.Type()
		switch valType {
		case d2stats.StatValueInt:
			v1.SetInt(v1.Int() + v2.Int())
		case d2stats.StatValueFloat:
			v1.SetFloat(v1.Float() + v2.Float())
		}
	}

	combined.SetValues(dstValues...)

	return combined, nil
}

func (s *Diablo2Stat) canBeCombinedWith(other d2stats.Stat) bool {
	if s.Name() != other.Name() {
		return false
	}

	values1, values2 := s.Values(), other.Values()
	if len(values1) != len(values2) {
		return false
	}

	for idx := range values1 {
		if values1[idx].Type() != values2[idx].Type() {
			return false
		}
	}

	return true
}

// String returns the formatted description string
func (s *Diablo2Stat) String() string { //nolint:gocyclo switch statement is not so bad
	var result string

	//nolint:gomdn introducing a const for these would be worse
	switch s.record.DescFnID {
	case 1:
		result = s.descFn1()
	case 2:
		result = s.descFn2()
	case 3:
		result = s.descFn3()
	case 4:
		result = s.descFn4()
	case 5:
		result = s.descFn5()
	case 6:
		result = s.descFn6()
	case 7:
		result = s.descFn7()
	case 8:
		result = s.descFn8()
	case 9:
		result = s.descFn9()
	case 11:
		result = s.descFn11()
	case 12:
		result = s.descFn12()
	case 13:
		result = s.descFn13()
	case 14:
		result = s.descFn14()
	case 15:
		result = s.descFn15()
	case 16:
		result = s.descFn16()
	case 20:
		result = s.descFn20()
	case 22:
		result = s.descFn22()
	case 23:
		result = s.descFn23()
	case 24:
		result = s.descFn24()
	case 27:
		result = s.descFn27()
	case 28:
		result = s.descFn28()
	default:
		result = ""
	}

	return result
}

func (s *Diablo2Stat) descFn1() string {
	var stringTableKey, result string

	value := s.values[0]

	value.SetStringer(stringerIntSigned)

	formatString := twoComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	switch descValPosition(s.record.DescVal) {
	case descValPrefix:
		result = fmt.Sprintf(formatString, value.String(), stringTableString)
	case descValPostfix:
		result = fmt.Sprintf(formatString, stringTableString, value.String())
	case descValHide:
		result = stringTableString
	default:
		result = ""
	}

	return result
}

func (s *Diablo2Stat) descFn2() string {
	var stringTableKey, result string

	value := s.values[0]

	value.SetStringer(stringerIntPercentageSigned)

	formatString := twoComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	switch descValPosition(s.record.DescVal) {
	case descValPrefix:
		result = fmt.Sprintf(formatString, value.String(), stringTableString)
	case descValPostfix:
		result = fmt.Sprintf(formatString, stringTableString, value.String())
	case descValHide:
		result = stringTableString
	default:
		result = ""
	}

	return result
}

func (s *Diablo2Stat) descFn3() string {
	var stringTableKey, result string

	value := s.values[0]

	formatString := twoComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	switch descValPosition(s.record.DescVal) {
	case descValPrefix:
		result = fmt.Sprintf(formatString, value.String(), stringTableString)
	case descValPostfix:
		result = fmt.Sprintf(formatString, stringTableString, value.String())
	case descValHide:
		result = stringTableString
	default:
		result = ""
	}

	return result
}

func (s *Diablo2Stat) descFn4() string {
	// for now, same as fn2
	return s.descFn2()
}

func (s *Diablo2Stat) descFn5() string {
	var stringTableKey, result string

	value := s.values[0]

	value.SetStringer(stringerIntPercentageUnsigned)

	formatString := twoComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	switch descValPosition(s.record.DescVal) {
	case descValPrefix:
		result = fmt.Sprintf(formatString, value.String(), stringTableString)
	case descValPostfix:
		result = fmt.Sprintf(formatString, stringTableString, value.String())
	case descValHide:
		result = stringTableString
	default:
		result = ""
	}

	return result
}

func (s *Diablo2Stat) descFn6() string {
	var stringTableKey, result string

	value := s.values[0]

	value.SetStringer(stringerIntSigned)

	formatString := threeComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	str1 := d2common.TranslateString(stringTableKey)
	str2 := d2common.TranslateString(s.record.DescStr2)

	switch descValPosition(s.record.DescVal) {
	case descValPrefix:
		result = fmt.Sprintf(formatString, value.String(), str1, str2)
	case descValPostfix:
		result = fmt.Sprintf(formatString, str1, value.String(), str2)
	case descValHide:
		formatString = twoComponentStr
		result = fmt.Sprintf(formatString, value.String(), str2)
	default:
		result = ""
	}

	return result
}

func (s *Diablo2Stat) descFn7() string {
	var stringTableKey, result string

	value := s.values[0]

	value.SetStringer(stringerIntPercentageSigned)

	formatString := threeComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	str1 := d2common.TranslateString(stringTableKey)
	str2 := d2common.TranslateString(s.record.DescStr2)

	switch descValPosition(s.record.DescVal) {
	case descValPrefix:
		result = fmt.Sprintf(formatString, value.String(), str1, str2)
	case descValPostfix:
		result = fmt.Sprintf(formatString, str1, value.String(), str2)
	case descValHide:
		formatString = twoComponentStr
		result = fmt.Sprintf(formatString, value.String(), str2)
	default:
		result = ""
	}

	return result
}

func (s *Diablo2Stat) descFn8() string {
	// for now, same as fn7
	return s.descFn7()
}

func (s *Diablo2Stat) descFn9() string {
	var stringTableKey, result string

	value := s.values[0]

	formatString := threeComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	str1 := d2common.TranslateString(stringTableKey)
	str2 := d2common.TranslateString(s.record.DescStr2)

	switch descValPosition(s.record.DescVal) {
	case descValPrefix:
		result = fmt.Sprintf(formatString, value.String(), str1, str2)
	case descValPostfix:
		result = fmt.Sprintf(formatString, str1, value.String(), str2)
	case descValHide:
		result = fmt.Sprintf(twoComponentStr, value.String(), str2)
	default:
		result = ""
	}

	return result
}

func (s *Diablo2Stat) descFn11() string {
	var stringTableKey string

	var result string

	value := s.values[0]

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	str1 := d2common.TranslateString(stringTableKey)

	formatString := str1

	result = fmt.Sprintf(formatString, value.String())

	return result
}

func (s *Diablo2Stat) descFn12() string {
	return s.descFn1()
}

func (s *Diablo2Stat) descFn13() string {
	var result string

	value := s.values[0]
	allSkills := s.values[1]

	value.SetStringer(stringerIntSigned)
	allSkills.SetStringer(stringerClassAllSkills)

	formatString := twoComponentStr

	switch descValPosition(s.record.DescVal) {
	case descValPrefix:
		result = fmt.Sprintf(formatString, value.String(), allSkills.String())
	case descValPostfix:
		result = fmt.Sprintf(formatString, allSkills.String(), value.String())
	case descValHide:
		result = allSkills.String()
	default:
		result = ""
	}

	return result
}

func (s *Diablo2Stat) descFn14() string {
	// strings come out like `+5 to Combat Skills (Paladin Only)`
	numSkills, hero, skillTab := s.values[0], s.values[1], s.values[2]
	heroMap := getHeroMap()
	heroIndex := hero.Int()
	classRecord := d2datadict.CharStats[heroMap[heroIndex]]

	// diablo 2 is hardcoded to have only 3 skill tabs
	skillTabIndex := skillTab.Int()
	if skillTabIndex < 0 || skillTabIndex > maxSkillTabIndex {
		skillTabIndex = 0
	}

	// `+5`
	numSkills.SetStringer(stringerIntSigned)
	numSkillsStr := numSkills.String()

	// `to Combat Skills`
	skillTabKey := classRecord.SkillStrTab[skillTabIndex]
	skillTabStr := d2common.TranslateString(skillTabKey)
	skillTabStr = strings.ReplaceAll(skillTabStr, "+%d ", "") // has a token we dont need

	// `(Paladin Only)`
	hero.SetStringer(stringerClassOnly)
	classOnlyStr := hero.String()

	return fmt.Sprintf(threeComponentStr, numSkillsStr, skillTabStr, classOnlyStr)
}

func (s *Diablo2Stat) descFn15() string {
	chance, lvl, skill := s.values[0], s.values[1], s.values[2]

	chance.SetStringer(stringerIntPercentageUnsigned)
	skill.SetStringer(stringerSkillName)

	// Special case, `chance to cast` format is actually in the string table!
	chanceToCastStr := d2common.TranslateString(s.record.DescStrPos)

	return fmt.Sprintf(chanceToCastStr, chance.Int(), lvl.Int(), skill.String())
}

func (s *Diablo2Stat) descFn16() string {
	skillLevel, skillIndex := s.values[0], s.values[1]

	skillIndex.SetStringer(stringerSkillName)

	// Special case, `Level # XYZ Aura When Equipped`, format is actually in the string table!
	format := d2common.TranslateString(s.record.DescStrPos)

	return fmt.Sprintf(format, skillLevel.Int(), skillIndex.String())
}

/*
func (s *Diablo2Stat) descFn17() string {
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

func (s *Diablo2Stat) descFn18() string {
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

func (s *Diablo2Stat) descFn20() string {
	// for now, same as fn2
	return s.descFn2()
}

func (s *Diablo2Stat) descFn22() string {
	arBonus, monsterIndex := s.values[0], s.values[1]
	arVersus := d2common.TranslateString(s.record.DescStrPos)

	arBonus.SetStringer(stringerIntPercentageUnsigned)
	monsterIndex.SetStringer(stringerMonsterName)

	return fmt.Sprintf(threeComponentStr, arBonus.String(), arVersus, monsterIndex.String())
}

func (s *Diablo2Stat) descFn23() string {
	// for now, same as fn22
	return s.descFn22()
}

func (s *Diablo2Stat) descFn24() string {
	// Special case formatting
	format := "Level " + threeComponentStr

	lvl, skillIdx, chargeMax, chargeCurrent := s.values[0],
		s.values[1],
		s.values[2].Int(),
		s.values[3].Int()

	skillIdx.SetStringer(stringerSkillName)

	chargeStr := d2common.TranslateString(s.record.DescStrPos)
	chargeStr = fmt.Sprintf(chargeStr, chargeCurrent, chargeMax)

	return fmt.Sprintf(format, lvl.String(), skillIdx.String(), chargeStr)
}

func (s *Diablo2Stat) descFn27() string {
	amount, skillIdx, heroIdx := s.values[0], s.values[1], s.values[2]

	amount.SetStringer(stringerIntSigned)
	skillIdx.SetStringer(stringerSkillName)
	heroIdx.SetStringer(stringerClassOnly)

	return fmt.Sprintf(fourComponentStr, amount.String(), "to", skillIdx.String(), heroIdx.String())
}

func (s *Diablo2Stat) descFn28() string {
	amount, skillIdx := s.values[0], s.values[1]

	amount.SetStringer(stringerIntSigned)
	skillIdx.SetStringer(stringerSkillName)

	return fmt.Sprintf(threeComponentStr, amount.String(), "to", skillIdx.String())
}

// DescGroupString return a string based on the DescGroupFuncID
func (s *Diablo2Stat) DescGroupString(a ...interface{}) string {
	if s.record.DescGroupFuncID < 0 {
		return ""
	}

	format := ""
	for range a {
		format += "%s "
	}

	format = strings.Trim(format, " ")

	return fmt.Sprintf(format, a...)
}
