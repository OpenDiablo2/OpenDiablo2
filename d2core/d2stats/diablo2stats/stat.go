package diablo2stats

import (
	"fmt"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// static check that diablo2Stat implements Stat
var _ d2stats.Stat = &diablo2Stat{}

type descValPosition int

const (
	descValHide descValPosition = iota
	descValPrefix
	descValPostfix
)

const (
	maxSkillTabIndex = 2
	oneValue         = 1
	twoValue         = 2
	threeValue       = 3
	fourValue        = 4
)

const (
	twoComponentStr   = "%s %s"
	threeComponentStr = "%s %s %s"
	fourComponentStr  = "%s %s %s %s"
)

const (
	intVal = d2stats.StatValueInt
	sum    = d2stats.StatValueCombineSum
	static = d2stats.StatValueCombineStatic
)

// diablo2Stat is an implementation of an OpenDiablo2 Stat, with a set of values.
// It is pretty tightly coupled to the data files for d2
type diablo2Stat struct {
	factory *StatFactory
	record  *d2records.ItemStatCostRecord
	values  []d2stats.StatValue
}

// depending on the stat record, sets up the proper number of values,
// as well as set up the stat value number types, value combination types, and
// the value stringer functions used
func (s *diablo2Stat) init(numbers ...float64) { //nolint:funlen,gocyclo // can't reduce
	if s.record == nil {
		return
	}

	//nolint:gomnd // introducing a const for these would be worse
	switch s.record.DescFnID {
	case 0:
		// special case for poisonlength, or other stats, which have a
		// 0-value descfnID field but need to store values
		s.values = make([]d2stats.StatValue, len(numbers))
		for idx := range s.values {
			s.values[idx] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerIntSigned)
		}
	case 1:
		// +31 to Strength
		// Replenish Life +20 || Drain Life -8
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerIntSigned)
	case 2:
		// +16% Increased Chance of Blocking
		// Lightning Absorb +10%
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal,
			sum).SetStringer(s.factory.stringerIntPercentageSigned)
	case 3:
		// Damage Reduced by 25
		// Slain Monsters Rest in Peace
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal, sum)
	case 4:
		// Poison Resist +25%
		// +25% Faster Run/Walk
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal,
			sum).SetStringer(s.factory.stringerIntPercentageSigned)
	case 5:
		// Hit Causes Monster to Flee 25%
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal, sum)
		s.values[0].SetStringer(s.factory.stringerIntPercentageUnsigned)
	case 6:
		// +25 to Life (Based on Character Level)
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerIntSigned)
	case 7:
		// Lightning Resist +25% (Based on Character Level)
		// +25% Better Chance of Getting Magic Items (Based on Character Level)
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal,
			sum).SetStringer(s.factory.stringerIntPercentageSigned)
	case 8:
		// +25% Enhanced Defense (Based on Character Level)
		// Heal Stamina Plus +25% (Based on Character Level)
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal,
			sum).SetStringer(s.factory.stringerIntPercentageSigned)
	case 9:
		// Attacker Takes Damage of 25 (Based on Character Level)
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal, sum)
	case 11:
		// Repairs 2 durability per second
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal, sum)
	case 12:
		// Hit Blinds Target +5
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerIntSigned)
	case 13:
		// +5 to Paladin Skill Levels
		s.values = make([]d2stats.StatValue, twoValue)
		s.values[0] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerIntSigned)
		s.values[1] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerClassAllSkills)
	case 14:
		// +5 to Combat Skills (Paladin Only)
		s.values = make([]d2stats.StatValue, threeValue)
		s.values[0] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerIntSigned)
		s.values[1] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerClassOnly)
		s.values[2] = s.factory.NewValue(intVal, static)
	case 15:
		//  5% Chance to cast level 7 Frozen Orb on attack
		s.values = make([]d2stats.StatValue, threeValue)
		s.values[0] = s.factory.NewValue(intVal, sum)
		s.values[1] = s.factory.NewValue(intVal, static)
		s.values[2] = s.factory.NewValue(intVal, static).SetStringer(s.factory.stringerSkillName)
	case 16:
		// Level 3 Warmth Aura When Equipped
		s.values = make([]d2stats.StatValue, twoValue)
		s.values[0] = s.factory.NewValue(intVal, sum)
		s.values[1] = s.factory.NewValue(intVal, static).SetStringer(s.factory.stringerSkillName)
	case 20:
		// -25% Target Defense
		s.values = make([]d2stats.StatValue, oneValue)
		s.values[0] = s.factory.NewValue(intVal,
			sum).SetStringer(s.factory.stringerIntPercentageSigned)
	case 22:
		// 25% to Attack Rating versus Specter
		s.values = make([]d2stats.StatValue, twoValue)
		s.values[0] = s.factory.NewValue(intVal,
			sum).SetStringer(s.factory.stringerIntPercentageUnsigned)
		s.values[1] = s.factory.NewValue(intVal, static).SetStringer(s.factory.stringerMonsterName)
	case 23:
		//  25% Reanimate as: Specter
		s.values = make([]d2stats.StatValue, twoValue)
		s.values[0] = s.factory.NewValue(intVal,
			sum).SetStringer(s.factory.stringerIntPercentageUnsigned)
		s.values[1] = s.factory.NewValue(intVal, static).SetStringer(s.factory.stringerMonsterName)
	case 24:
		// Level 25 Frozen Orb (19/20 Charges)
		s.values = make([]d2stats.StatValue, fourValue)
		s.values[0] = s.factory.NewValue(intVal, static)
		s.values[1] = s.factory.NewValue(intVal, static).SetStringer(s.factory.stringerSkillName)
		s.values[2] = s.factory.NewValue(intVal, static)
		s.values[3] = s.factory.NewValue(intVal, static)
	case 27:
		// +25 to Frozen Orb (Paladin Only)
		s.values = make([]d2stats.StatValue, threeValue)
		s.values[0] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerIntSigned)
		s.values[1] = s.factory.NewValue(intVal, static).SetStringer(s.factory.stringerSkillName)
		s.values[2] = s.factory.NewValue(intVal, static).SetStringer(s.factory.stringerClassOnly)
	case 28:
		// +25 to Frozen Orb
		s.values = make([]d2stats.StatValue, twoValue)
		s.values[0] = s.factory.NewValue(intVal, sum).SetStringer(s.factory.stringerIntSigned)
		s.values[1] = s.factory.NewValue(intVal, static).SetStringer(s.factory.stringerSkillName)
	default:
		return
	}

	for idx := range numbers {
		if idx > len(s.values)-1 {
			break
		}

		s.values[idx].SetFloat(numbers[idx])
	}
}

// Name returns the name of the stat (the key in itemstatcosts)
func (s *diablo2Stat) Name() string {
	return s.record.Name
}

// Priority returns the description printing priority
func (s *diablo2Stat) Priority() int {
	return s.record.DescPriority
}

// Values returns the stat values of the stat
func (s *diablo2Stat) Values() []d2stats.StatValue {
	return s.values
}

// SetValues sets the stat values
func (s *diablo2Stat) SetValues(values ...d2stats.StatValue) {
	s.values = make([]d2stats.StatValue, len(values))
	for idx := range values {
		s.values[idx] = values[idx]
	}
}

// Clone returns a deep copy of the diablo2Stat
func (s *diablo2Stat) Clone() d2stats.Stat {
	clone := &diablo2Stat{
		factory: s.factory,
		record:  s.record,
	}

	clone.init()

	for idx := range s.values {
		srcVal := s.values[idx]
		dstVal := &Diablo2StatValue{
			numberType:  srcVal.NumberType(),
			combineType: srcVal.CombineType(),
		}

		dstVal.SetStringer(srcVal.Stringer())

		switch srcVal.NumberType() {
		case d2stats.StatValueInt:
			dstVal.SetInt(srcVal.Int())
		case d2stats.StatValueFloat:
			dstVal.SetFloat(srcVal.Float())
		}

		if len(clone.values) < len(s.values) {
			clone.values = make([]d2stats.StatValue, len(s.values))
		}

		clone.values[idx] = dstVal
	}

	return clone
}

// Copy to this stat value the values of the given stat value
func (s *diablo2Stat) Copy(from d2stats.Stat) d2stats.Stat {
	srcValues := from.Values()
	s.values = make([]d2stats.StatValue, len(srcValues))

	for idx := range srcValues {
		src := srcValues[idx]
		valType := src.NumberType()
		dst := &Diablo2StatValue{numberType: valType}
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
func (s *diablo2Stat) Combine(other d2stats.Stat) (result d2stats.Stat, err error) {
	cantBeCombinedErr := fmt.Errorf("cannot combine %s with %s", s.Name(), other.Name())

	if !s.canBeCombinedWith(other) {
		return nil, cantBeCombinedErr
	}

	result = s.Clone()
	srcValues, dstValues := other.Values(), result.Values()

	for idx := range result.Values() {
		v1, v2 := dstValues[idx], srcValues[idx]
		combinationRule := v1.CombineType()

		if combinationRule == d2stats.StatValueCombineStatic {
			// we do not add the values, they remain the same
			// for things like monster/class/skill index or on proc stats
			// where the level of a skill isn't summed, but the
			// chance to cast values are
			continue
		}

		if combinationRule == d2stats.StatValueCombineSum {
			valType := v1.NumberType()
			switch valType {
			case d2stats.StatValueInt:
				v1.SetInt(v1.Int() + v2.Int())
			case d2stats.StatValueFloat:
				v1.SetFloat(v1.Float() + v2.Float())
			}
		}
	}

	return result, nil
}

func (s *diablo2Stat) canBeCombinedWith(other d2stats.Stat) bool {
	if s.Name() != other.Name() {
		return false
	}

	values1, values2 := s.Values(), other.Values()
	if len(values1) != len(values2) {
		return false
	}

	for idx := range values1 {
		if values1[idx].NumberType() != values2[idx].NumberType() {
			return false
		}

		if values1[idx].CombineType() != values2[idx].CombineType() {
			return false
		}

		// in the case that we are trying to combine stats like:
		// 		+1 to Paladin Skills
		// 		+1 to Sorceress Skills
		// the numeric value (an index) that denotes the class skill type knows not to be summed
		// with the other index, even though the format of the stat and stat value is pretty much
		// the same.
		if values1[idx].CombineType() == d2stats.StatValueCombineStatic {
			if values1[idx].Float() != values2[idx].Float() {
				return false
			}
		}
	}

	return true
}

// String returns the formatted description string
func (s *diablo2Stat) String() string { //nolint:gocyclo // switch statement is not so bad
	var result string

	for idx := range s.values {
		if s.values[idx].Stringer() == nil {
			s.values[idx].SetStringer(s.factory.stringerUnsignedInt)
		}
	}

	//nolint:gomnd // introducing a const for these would be worse
	switch s.record.DescFnID {
	case 1, 2, 3, 4, 5, 12, 20:
		result = s.descFn1()
	case 6, 7, 8:
		result = s.descFn6()
	case 9:
		result = s.descFn9()
	case 11:
		result = s.descFn11()
	case 13:
		result = s.descFn13()
	case 14:
		result = s.descFn14()
	case 15:
		result = s.descFn15()
	case 16:
		result = s.descFn16()
	case 22, 23:
		result = s.descFn22()
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

func (s *diablo2Stat) descFn1() string {
	var stringTableKey, result string

	value := s.values[0]

	formatString := twoComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	stringTableString := s.factory.asset.TranslateString(stringTableKey)

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

func (s *diablo2Stat) descFn6() string {
	var stringTableKey, result string

	value := s.values[0]

	formatString := threeComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	str1 := s.factory.asset.TranslateString(stringTableKey)
	str2 := s.factory.asset.TranslateString(s.record.DescStr2)

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

func (s *diablo2Stat) descFn9() string {
	var stringTableKey, result string

	value := s.values[0]

	formatString := threeComponentStr

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	str1 := s.factory.asset.TranslateString(stringTableKey)
	str2 := s.factory.asset.TranslateString(s.record.DescStr2)

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

func (s *diablo2Stat) descFn11() string {
	var stringTableKey string

	value := s.values[0]

	if value.Int() < 0 {
		stringTableKey = s.record.DescStrNeg
	} else {
		stringTableKey = s.record.DescStrPos
	}

	str1 := s.factory.asset.TranslateString(stringTableKey)

	formatString := str1

	return fmt.Sprintf(formatString, value)
}

func (s *diablo2Stat) descFn13() string {
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

func (s *diablo2Stat) descFn14() string {
	// strings come out like `+5 to Combat Skills (Paladin Only)`
	numSkills, hero, skillTab := s.values[0], s.values[1], s.values[2]
	heroMap := s.factory.getHeroMap()
	heroIndex := hero.Int()
	classRecord := s.factory.asset.Records.Character.Stats[heroMap[heroIndex]]

	// diablo 2 is hardcoded to have only 3 skill tabs
	skillTabIndex := skillTab.Int()
	if skillTabIndex < 0 || skillTabIndex > maxSkillTabIndex {
		skillTabIndex = 0
	}

	// `+5`
	numSkillsStr := numSkills.String()

	// `to Combat Skills`
	skillTabKey := classRecord.SkillStrTab[skillTabIndex]
	skillTabStr := s.factory.asset.TranslateString(skillTabKey)
	skillTabStr = strings.ReplaceAll(skillTabStr, "+%d ", "") // has a token we dont need

	// `(Paladin Only)`
	classOnlyStr := hero.String()

	return fmt.Sprintf(threeComponentStr, numSkillsStr, skillTabStr, classOnlyStr)
}

func (s *diablo2Stat) descFn15() string {
	chance, lvl, skill := s.values[0], s.values[1], s.values[2]

	// Special case, `chance to cast` format is actually in the string table!
	chanceToCastStr := s.factory.asset.TranslateString(s.record.DescStrPos)

	return fmt.Sprintf(chanceToCastStr, chance.Int(), lvl.Int(), skill)
}

func (s *diablo2Stat) descFn16() string {
	skillLevel, skillIndex := s.values[0], s.values[1]

	// Special case, `Level # XYZ Aura When Equipped`, format is actually in the string table!
	format := s.factory.asset.TranslateString(s.record.DescStrPos)

	return fmt.Sprintf(format, skillLevel.Int(), skillIndex)
}

func (s *diablo2Stat) descFn22() string {
	arBonus, monsterIndex := s.values[0], s.values[1]
	arVersus := s.factory.asset.TranslateString(s.record.DescStrPos)

	return fmt.Sprintf(threeComponentStr, arBonus, arVersus, monsterIndex)
}

func (s *diablo2Stat) descFn24() string {
	// Special case formatting
	format := "Level " + threeComponentStr

	lvl, skill, chargeMax, chargeCurrent := s.values[0],
		s.values[1],
		s.values[2].Int(),
		s.values[3].Int()

	chargeStr := s.factory.asset.TranslateString(s.record.DescStrPos)
	chargeStr = fmt.Sprintf(chargeStr, chargeCurrent, chargeMax)

	return fmt.Sprintf(format, lvl, skill, chargeStr)
}

func (s *diablo2Stat) descFn27() string {
	// property "skill-rand" will try to make an instance with an invalid hero index
	// in this case, we use descfn 28
	if s.values[2].Int() == -1 {
		return s.descFn28()
	}

	amount, skill, hero := s.values[0], s.values[1], s.values[2]

	return fmt.Sprintf(fourComponentStr, amount, "to", skill, hero)
}

func (s *diablo2Stat) descFn28() string {
	amount, skill := s.values[0], s.values[1]

	return fmt.Sprintf(threeComponentStr, amount, "to", skill)
}

// DescGroupString return a string based on the DescGroupFuncID
func (s *diablo2Stat) DescGroupString(a ...interface{}) string {
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
