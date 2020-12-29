package diablo2item

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

const (
	noValue = iota
	oneValue
	twoValue
	threeValue
)

const (
	skillTabsPerClass = 3
)

// these come from properties.txt, the types of functions that properties can use to evaluate args
const (
	fnNone = iota
	fnValuesToStat
	fnArmorPercent
	fnRepeatPreviousWithMinMax // repeat only with min and max
	fnUnused
	fnDamageMin
	fnDamageMax
	fnDamagePercent
	fnSpeedRelated
	fnRepeatPreviousWithParamMinMax // repeat with param, man, and max
	fnClassSkillTab
	fnProcs
	fnRandomSkill
	fnMaxDurability
	fnNumSockets
	fnStatMin
	fnStatMax
	fnStatParam
	fnTimeRelated
	fnChargeRelated
	fnIndestructable
	fnClassSkills
	fnSingleSkill
	fnEthereal
	fnStateApplyToTarget
)

// PropertyType describes what kind of property this is
type PropertyType int

// Property types
// Not all properties contain stats, some are just used to compute a value
// examples are:
//		min/max
//		% damage
//		indestructable and etheral flags
const (
	PropertyComputeStats   = iota // for properties that do compute stats
	PropertyComputeInteger        // for properties that compute an integer value
	PropertyComputeBoolean        // for properties that compute a boolean
)

const (
	fnRandClassSkill = 36
)

// Property is an item property.
type Property struct {
	factory      *ItemFactory
	record       *d2records.PropertyRecord
	stats        []d2stats.Stat
	PropertyType PropertyType

	// the inputValues that were passed initially when calling `NewProperty`
	inputParams []int

	// some properties are statless and used only for computing a value
	computedInt  int
	computedBool bool
}

func (p *Property) init() *Property {
	p.stats = make([]d2stats.Stat, 0)

	// some property functions need to be able to repeat last function
	// this is for properties with multiple stats that want to repeat the same
	// initialization step with the same min/max params
	var lastFnCalled int

	var stat d2stats.Stat

	for idx := range p.record.Stats {
		if p.record.Stats[idx] == nil {
			continue
		}

		stat, lastFnCalled = p.eval(idx, lastFnCalled)

		// some property stats don't actually have a stat
		// but they have functions on the first stat entry
		if stat != nil {
			p.stats = append(p.stats, stat)
		}
	}

	return p
}

// eval will attempt to create a stat, and will return the function id that was last run.
// this is because some of the properties have a func index which indicates that it should
// repeat the previous fn with the same parameters, but for a different stat.
func (p *Property) eval(propStatIdx, previousFnID int) (stat d2stats.Stat, funcID int) {
	pStatRecord := p.record.Stats[propStatIdx]
	iscRecord := p.factory.asset.Records.Item.Stats[pStatRecord.StatCode]

	funcID = pStatRecord.FunctionID

	switch funcID {
	case fnRepeatPreviousWithMinMax, fnRepeatPreviousWithParamMinMax:
		funcID = previousFnID
		fallthrough
	case fnValuesToStat, fnSpeedRelated, fnMaxDurability, fnNumSockets,
		fnStatMin, fnStatMax, fnSingleSkill, fnArmorPercent:
		p.PropertyType = PropertyComputeStats
		stat = p.fnValuesToStat(iscRecord)
	case fnDamageMin, fnDamageMax, fnDamagePercent:
		p.PropertyType = PropertyComputeInteger
		p.computedInt = p.fnComputeInteger()
	case fnClassSkillTab:
		p.PropertyType = PropertyComputeStats
		stat = p.fnClassSkillTab(iscRecord)
	case fnProcs:
		p.PropertyType = PropertyComputeStats
		stat = p.fnProcs(iscRecord)
	case fnRandomSkill:
		p.PropertyType = PropertyComputeStats
		stat = p.fnRandomSkill(iscRecord)
	case fnStatParam:
		p.PropertyType = PropertyComputeStats
		stat = p.fnStatParam(iscRecord)
	case fnChargeRelated:
		p.PropertyType = PropertyComputeStats
		stat = p.fnChargeRelated(iscRecord)
	case fnIndestructable, fnEthereal:
		p.PropertyType = PropertyComputeBoolean
		p.computedBool = p.fnBoolean()
	case fnClassSkills:
		p.PropertyType = PropertyComputeStats
		stat = p.fnClassSkills(pStatRecord, iscRecord)
	case fnStateApplyToTarget:
		p.PropertyType = PropertyComputeStats
		stat = p.fnStateApplyToTarget(iscRecord)
	case fnRandClassSkill:
		p.PropertyType = PropertyComputeStats
		stat = p.fnRandClassSkill(iscRecord)
	case fnNone, fnUnused, fnTimeRelated:
	default:
	}

	return stat, funcID
}

// fnValuesToStat Applies a value to a stat, can use SetX parameter.
func (p *Property) fnValuesToStat(iscRecord *d2records.ItemStatCostRecord) d2stats.Stat {
	// the only special case to handle for this function is for
	// property "color", which corresponds to ISC record "item_lightcolor"
	// I'm not yet sure how to handle this special case... it is likely
	// and index into one of the colors in colors.txt
	var min, max int

	var propParam, statValue float64

	switch len(p.inputParams) {
	case noValue, oneValue:
		return nil
	case twoValue:
		min, max = p.inputParams[0], p.inputParams[1]
	case threeValue:
		propParam = float64(p.inputParams[0])
		min, max = p.inputParams[1], p.inputParams[2]
	default:
		min, max = p.inputParams[0], p.inputParams[1]
	}

	if max < min {
		min, max = max, min
	}

	// nolint:gosec // not concerned with crypto-strong randomness
	statValue = float64(rand.Intn(max-min+1) + min)

	return p.factory.stat.NewStat(iscRecord.Name, statValue, propParam)
}

// fnComputeInteger Dmg-min related ???
func (p *Property) fnComputeInteger() int {
	var min, max int

	switch len(p.inputParams) {
	case noValue, oneValue:
		return 0
	default:
		min, max = p.inputParams[0], p.inputParams[1]
	}

	// nolint:gosec // not concerned with crypto-strong randomness
	statValue := rand.Intn(max-min+1) + min

	return statValue
}

// fnClassSkillTab skilltab skill group ???
func (p *Property) fnClassSkillTab(iscRecord *d2records.ItemStatCostRecord) d2stats.Stat {
	// from here: https://d2mods.info/forum/kb/viewarticle?a=45
	// Amazon
	// 0 - Bow & Crossbow
	// 1 - Passive & Magic
	// 2 - Spear & Javelin
	// Sorceress
	// 3 - Fire
	// 4 - Lightning
	// 5 - Cold
	// Necromancer
	// 6 - Curses
	// 7 - Poison & Bone
	// 8 - Summoning
	// Paladin
	// 9 - Offensive Auras
	// 10 - Combat Skills
	// 11 - Defensive Auras
	// Barbarian
	// 12 - Masteries
	// 13 - Combat Skills
	// 14 - Warcries
	// Druid
	// 15 - Summoning
	// 16 - Shapeshifting
	// 17 - Elemental
	// Assassin
	// 18 - Traps
	// 19 - Shadow Disciplines
	// 20 - Martial Arts
	param, min, max := p.inputParams[0], p.inputParams[1], p.inputParams[2]
	skillTabIdx := float64(param % skillTabsPerClass)
	classIdx := float64(param / skillTabsPerClass)

	// nolint:gosec // not concerned with crypto-strong randomness
	level := float64(rand.Intn(max-min+1) + min)

	return p.factory.stat.NewStat(iscRecord.Name, level, classIdx, skillTabIdx)
}

// fnProcs event-based skills ???
func (p *Property) fnProcs(iscRecord *d2records.ItemStatCostRecord) d2stats.Stat {
	var skillID, chance, skillLevel float64

	switch len(p.inputParams) {
	case noValue, oneValue, twoValue:
		return nil
	default:
		skillID = float64(p.inputParams[0])
		chance = float64(p.inputParams[1])
		skillLevel = float64(p.inputParams[2])
	}

	return p.factory.stat.NewStat(iscRecord.Name, chance, skillLevel, skillID)
}

// fnRandomSkill random selection of parameters for parameter-based stat ???
func (p *Property) fnRandomSkill(iscRecord *d2records.ItemStatCostRecord) d2stats.Stat {
	var skillLevel, skillID float64

	invalidHeroIndex := -1.0

	switch len(p.inputParams) {
	case noValue, oneValue, twoValue:
		return nil
	default:
		skillLevel = float64(p.inputParams[0])
		min, max := p.inputParams[1], p.inputParams[2]
		// nolint:gosec // not concerned with crypto-strong randomness
		skillID = float64(rand.Intn(max-min+1) + min)
	}

	return p.factory.stat.NewStat(iscRecord.Name, skillLevel, skillID, invalidHeroIndex)
}

// fnStatParam use param field only
func (p *Property) fnStatParam(iscRecord *d2records.ItemStatCostRecord) d2stats.Stat {
	switch len(p.inputParams) {
	case noValue:
		return nil
	default:
		val := float64(p.inputParams[0])
		return p.factory.stat.NewStat(iscRecord.Name, val)
	}
}

// fnChargeRelated Related to charged item.
func (p *Property) fnChargeRelated(iscRecord *d2records.ItemStatCostRecord) d2stats.Stat {
	var lvl, skill, charges float64

	switch len(p.inputParams) {
	case noValue, oneValue, twoValue:
		return nil
	default:
		lvl = float64(p.inputParams[2])
		skill = float64(p.inputParams[0])
		charges = float64(p.inputParams[1])

		return p.factory.stat.NewStat(iscRecord.Name, lvl, skill, charges, charges)
	}
}

// fnIndestructable Simple boolean stuff. Use by indestruct.
func (p *Property) fnBoolean() bool {
	var min, max int

	switch len(p.inputParams) {
	case noValue, oneValue:
		return false
	default:
		min, max = p.inputParams[0], p.inputParams[1]
	}

	// nolint:gosec // not concerned with crypto-strong randomness
	statValue := rand.Intn(max-min+1) + min

	return statValue > 0
}

// fnClassSkills Add to group of skills, group determined by stat ID, uses ValX parameter.
func (p *Property) fnClassSkills(
	propStatRecord *d2records.PropertyStatRecord,
	iscRecord *d2records.ItemStatCostRecord,
) d2stats.Stat {
	// in order 0..6
	// Amazon
	// Sorceress
	// Necromancer
	// Paladin
	// Druid
	// Assassin
	var min, max, classIdx int

	switch len(p.inputParams) {
	case noValue, oneValue:
		return nil
	default:
		min, max = p.inputParams[0], p.inputParams[1]
	}

	// nolint:gosec // not concerned with crypto-strong randomness
	statValue := rand.Intn(max-min+1) + min
	classIdx = propStatRecord.Value

	return p.factory.stat.NewStat(iscRecord.Name, float64(statValue), float64(classIdx))
}

// fnStateApplyToTarget property applied to character or target monster ???
func (p *Property) fnStateApplyToTarget(_ *d2records.ItemStatCostRecord) d2stats.Stat {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/818
	return nil
}

// fnRandClassSkill property applied to character or target monster ???
func (p *Property) fnRandClassSkill(_ *d2records.ItemStatCostRecord) d2stats.Stat {
	return nil
}
