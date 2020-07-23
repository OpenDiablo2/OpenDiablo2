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
func (s *Diablo2Stat) Combine(other d2stats.Stat) (result d2stats.Stat, err error) {
	cantBeCombinedErr := fmt.Errorf("cannot combine %s with %s", s.Name(), other.Name())

	if !s.canBeCombinedWith(other) {
		return nil, cantBeCombinedErr
	}

	result = s.Clone()
	srcValues, dstValues := other.Values(), result.Values()

	for idx := range result.Values() {
		v1, v2 := dstValues[idx], srcValues[idx]

		valType := v1.Type()
		switch valType {
		case d2stats.StatValueInt:
			v1.SetInt(v1.Int() + v2.Int())
		case d2stats.StatValueFloat:
			v1.SetFloat(v1.Float() + v2.Float())
		}
	}

	return result, nil
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
		s.values[0].SetStringer(stringerIntSigned)
		result = s.descFn1()
	case 2:
		s.values[0].SetStringer(stringerIntPercentageSigned)
		result = s.descFn2()
	case 3:
		result = s.descFn3()
	case 4:
		s.values[0].SetStringer(stringerIntPercentageSigned)
		result = s.descFn4()
	case 5:
		s.values[0].SetStringer(stringerIntPercentageUnsigned)
		result = s.descFn5()
	case 6:
		s.values[0].SetStringer(stringerIntSigned)
		result = s.descFn6()
	case 7:
		s.values[0].SetStringer(stringerIntPercentageSigned)
		result = s.descFn7()
	case 8:
		s.values[0].SetStringer(stringerIntPercentageSigned)
		result = s.descFn8()
	case 9:
		result = s.descFn9()
	case 11:
		result = s.descFn11()
	case 12:
		s.values[0].SetStringer(stringerIntSigned)
		result = s.descFn12()
	case 13:
		s.values[0].SetStringer(stringerIntSigned)
		s.values[1].SetStringer(stringerClassAllSkills)
		result = s.descFn13()
	case 14:
		s.values[0].SetStringer(stringerIntSigned)
		s.values[1].SetStringer(stringerClassOnly)
		result = s.descFn14()
	case 15:
		s.values[2].SetStringer(stringerSkillName)
		result = s.descFn15()
	case 16:
		s.values[1].SetStringer(stringerSkillName)
		result = s.descFn16()
	case 20:
		s.values[0].SetStringer(stringerIntPercentageSigned)
		result = s.descFn20()
	case 22:
		s.values[0].SetStringer(stringerIntPercentageUnsigned)
		s.values[1].SetStringer(stringerMonsterName)
		result = s.descFn22()
	case 23:
		s.values[0].SetStringer(stringerIntPercentageUnsigned)
		s.values[1].SetStringer(stringerMonsterName)
		result = s.descFn23()
	case 24:
		s.values[1].SetStringer(stringerSkillName)
		result = s.descFn24()
	case 27:
		s.values[0].SetStringer(stringerIntSigned)
		s.values[1].SetStringer(stringerSkillName)
		s.values[2].SetStringer(stringerClassOnly)
		result = s.descFn27()
	case 28:
		s.values[0].SetStringer(stringerIntSigned)
		s.values[1].SetStringer(stringerSkillName)
		result = s.descFn28()
	default:
		result = ""
	}

	return result
}

// +31 to Strength
// Replenish Life +20 || Drain Life -8
func (s *Diablo2Stat) descFn1() string {
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
		result = fmt.Sprintf(formatString, value, stringTableString)
	case descValPostfix:
		result = fmt.Sprintf(formatString, stringTableString, value)
	case descValHide:
		result = stringTableString
	default:
		result = ""
	}

	return result
}

// +16% Increased Chance of Blocking
// Lightning Absorb +10%
func (s *Diablo2Stat) descFn2() string {
	// for now, same as fn1
	return s.descFn1()
}

// Damage Reduced by 25
// Slain Monsters Rest in Peace
func (s *Diablo2Stat) descFn3() string {
	// for now, same as fn1
	return s.descFn1()
}

// Poison Resist +25%
// +25% Faster Run/Walk
func (s *Diablo2Stat) descFn4() string {
	// for now, same as fn1
	return s.descFn1()
}

// Hit Causes Monster to Flee 25%
func (s *Diablo2Stat) descFn5() string {
	// for now, same as fn1
	return s.descFn1()
}

// +25 to Life (Based on Character Level)
func (s *Diablo2Stat) descFn6() string {
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
		result = fmt.Sprintf(formatString, value, str1, str2)
	case descValPostfix:
		result = fmt.Sprintf(formatString, str1, value, str2)
	case descValHide:
		formatString = twoComponentStr
		result = fmt.Sprintf(formatString, value, str2)
	default:
		result = ""
	}

	return result
}

// Lightning Resist +25% (Based on Character Level)
// +25% Better Chance of Getting Magic Items (Based on Character Level)
func (s *Diablo2Stat) descFn7() string {
	// for now, same as fn6
	return s.descFn6()
}

// +25% Enhanced Defense (Based on Character Level)
// Heal Stamina Plus +25% (Based on Character Level)
func (s *Diablo2Stat) descFn8() string {
	// for now, same as fn6
	return s.descFn6()
}

// Attacker Takes Damage of 25 (Based on Character Level)
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
		result = fmt.Sprintf(formatString, value, str1, str2)
	case descValPostfix:
		result = fmt.Sprintf(formatString, str1, value, str2)
	case descValHide:
		result = fmt.Sprintf(twoComponentStr, value, str2)
	default:
		result = ""
	}

	return result
}

// Repairs 2 durability per second
func (s *Diablo2Stat) descFn11() string {
	var stringTableKey string

	value := s.values[0]

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	str1 := d2common.TranslateString(stringTableKey)

	formatString := str1

	return fmt.Sprintf(formatString, value)
}

// Hit Blinds Target +5
func (s *Diablo2Stat) descFn12() string {
	// for now, same as fn1
	return s.descFn1()
}

// +5 to Paladin Skill Levels
func (s *Diablo2Stat) descFn13() string {
	value := s.values[0]
	allSkills := s.values[1]

	formatString := twoComponentStr

	switch descValPosition(s.record.DescVal) {
	case descValPrefix:
		return fmt.Sprintf(formatString, value, allSkills)
	case descValPostfix:
		return fmt.Sprintf(formatString, allSkills, value)
	case descValHide:
		return allSkills.String()
	default:
		return ""
	}
}

//  +5 to Combat Skills (Paladin Only)
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
	numSkillsStr := numSkills.String()

	// `to Combat Skills`
	skillTabKey := classRecord.SkillStrTab[skillTabIndex]
	skillTabStr := d2common.TranslateString(skillTabKey)
	skillTabStr = strings.ReplaceAll(skillTabStr, "+%d ", "") // has a token we dont need

	// `(Paladin Only)`
	classOnlyStr := hero.String()

	return fmt.Sprintf(threeComponentStr, numSkillsStr, skillTabStr, classOnlyStr)
}

//  5% Chance to cast level 7 Frozen Orb on attack
func (s *Diablo2Stat) descFn15() string {
	chance, lvl, skill := s.values[0], s.values[1], s.values[2]

	// Special case, `chance to cast` format is actually in the string table!
	chanceToCastStr := d2common.TranslateString(s.record.DescStrPos)

	return fmt.Sprintf(chanceToCastStr, chance.Int(), lvl.Int(), skill)
}

// Level 3 Warmth Aura When Equipped
func (s *Diablo2Stat) descFn16() string {
	skillLevel, skillIndex := s.values[0], s.values[1]

	// Special case, `Level # XYZ Aura When Equipped`, format is actually in the string table!
	format := d2common.TranslateString(s.record.DescStrPos)

	return fmt.Sprintf(format, skillLevel.Int(), skillIndex)
}

// -25% Target Defense
func (s *Diablo2Stat) descFn20() string {
	// for now, same as fn2
	return s.descFn2()
}

// 25% to Attack Rating versus Specter
func (s *Diablo2Stat) descFn22() string {
	arBonus, monsterIndex := s.values[0], s.values[1]
	arVersus := d2common.TranslateString(s.record.DescStrPos)

	return fmt.Sprintf(threeComponentStr, arBonus, arVersus, monsterIndex)
}

//  25% Reanimate as: Specter
func (s *Diablo2Stat) descFn23() string {
	// for now, same as fn22
	return s.descFn22()
}

// Level 25 Frozen Orb (19/20 Charges)
func (s *Diablo2Stat) descFn24() string {
	// Special case formatting
	format := "Level " + threeComponentStr

	lvl, skill, chargeMax, chargeCurrent := s.values[0],
		s.values[1],
		s.values[2].Int(),
		s.values[3].Int()

	chargeStr := d2common.TranslateString(s.record.DescStrPos)
	chargeStr = fmt.Sprintf(chargeStr, chargeCurrent, chargeMax)

	return fmt.Sprintf(format, lvl, skill, chargeStr)
}

// +25 to Frozen Orb (Paladin Only)
func (s *Diablo2Stat) descFn27() string {
	amount, skill, hero := s.values[0], s.values[1], s.values[2]

	return fmt.Sprintf(fourComponentStr, amount, "to", skill, hero)
}

// +25 to Frozen Orb
func (s *Diablo2Stat) descFn28() string {
	amount, skill := s.values[0], s.values[1]

	return fmt.Sprintf(threeComponentStr, amount, "to", skill)
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
