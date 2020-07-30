package diablo2item

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"math/rand"
	"regexp"
	"strconv"
)

const (
	DropModifierBaseProbability = 1024 // base DropModifier probability total
)

type DropModifier int

const (
	DropModifierNone DropModifier = iota
	DropModifierUnique
	DropModifierSet
	DropModifierRare
	DropModifierMagic
)

const (
	// DynamicItemLevelRange for treasure codes like `armo33`, this code is used to
	// select all equivalent items (matching `armo` in this case) with item levels 33,34,35
	DynamicItemLevelRange = 3
)

const (
	goldItemCodeWithMult = "gld,mul="
	goldItemCode         = "gld"
)

// ItemGenerator is a diablo 2 implementation of an item generator
type ItemGenerator struct {
	rand   *rand.Rand
	source rand.Source
	Seed   int64
}

// SetSeed sets the item generator seed
func (ig *ItemGenerator) SetSeed(seed int64) {
	if ig.rand == nil || ig.source == nil {
		ig.source = rand.NewSource(seed)
		ig.rand = rand.New(ig.source)
	}
	ig.Seed = seed
}

func (ig *ItemGenerator) rollDropModifier(tcr *d2datadict.TreasureClassRecord) DropModifier {
	modMap := map[int]DropModifier{
		0: DropModifierNone,
		1: DropModifierUnique,
		2: DropModifierSet,
		3: DropModifierRare,
		4: DropModifierMagic,
	}

	dropModifiers := []int{
		DropModifierBaseProbability,
		tcr.FreqUnique,
		tcr.FreqSet,
		tcr.FreqRare,
		tcr.FreqMagic,
	}

	for idx := range dropModifiers {
		if idx == 0 {
			continue
		}

		dropModifiers[idx] += dropModifiers[idx-1]
	}

	roll := ig.rand.Intn(dropModifiers[len(dropModifiers)-1])

	for idx := range dropModifiers {
		if roll < dropModifiers[idx] {
			return modMap[idx]
		}
	}

	return DropModifierNone
}

func (ig *ItemGenerator) rollTreasurePick(tcr *d2datadict.TreasureClassRecord) *d2datadict.Treasure {
	// treasure probabilities
	tprob := make([]int, len(tcr.Treasures)+1)
	total := tcr.FreqNoDrop
	tprob[0] = total

	for idx := range tcr.Treasures {
		total += tcr.Treasures[idx].Probability
		tprob[idx+1] = total
	}

	roll := ig.rand.Intn(total)

	for idx := range tprob {
		if roll < tprob[idx] {
			if idx == 0 {
				break
			}

			return tcr.Treasures[idx-1]
		}
	}

	return nil
}

// ItemsFromTreasureClass rolls for and creates items using a treasure class record
func (ig *ItemGenerator) ItemsFromTreasureClass(tcr *d2datadict.TreasureClassRecord) []*Item {
	result := make([]*Item, 0)

	treasurePicks := make([]*d2datadict.Treasure, 0)

	// if tcr.NumPicks is negative, each item probability is instead a count for how many
	// of that treasure to drop
	if tcr.NumPicks < 0 {
		picksLeft := tcr.NumPicks

		// for each of the treasures, we pick it N times, where N is the count for the item
		// we do this until we run out of picks
		for idx := range tcr.Treasures {
			howMany := tcr.Treasures[idx].Probability
			for count := 0; count < howMany && picksLeft < 0; count++ {
				treasurePicks = append(treasurePicks, tcr.Treasures[idx])
				picksLeft++
			}
		}
	} else {
		// for N picks, we roll for a treasure and append to our treasures if it isn't a NoDrop
		for picksLeft := tcr.NumPicks; picksLeft > 0; picksLeft-- {
			rolledTreasure := ig.rollTreasurePick(tcr)

			if rolledTreasure == nil {
				continue
			}

			treasurePicks = append(treasurePicks, rolledTreasure)
		}
	}

	// for each of our picked/rolled treasures, we will attempt to generate an item.
	// The treasure may actually be a reference to another treasure class, in which
	// case we will roll that treasure class, eventually getting a slice of items
	for idx := range treasurePicks {
		picked := treasurePicks[idx]
		if record, found := d2datadict.TreasureClass[picked.Code]; found {
			// the code is for a treasure class, we roll again using that TC
			itemSlice := ig.ItemsFromTreasureClass(record)
			for itemIdx := range itemSlice {
				itemSlice[itemIdx].applyDropModifier(ig.rollDropModifier(tcr))
				itemSlice[itemIdx].generateAllProperties()
				itemSlice[itemIdx].updateItemAttributes()
				result = append(result, itemSlice[itemIdx])
			}
		} else {
			// the code is not for a treasure class, but for an item
			item := ig.ItemFromTreasure(picked)
			if item != nil {
				item.applyDropModifier(ig.rollDropModifier(tcr))
				item.generateAllProperties()
				item.updateItemAttributes()
				result = append(result, item)
			}
		}
	}

	return result
}

// ItemFromTreasure rolls for a ig.rand.m item using the Treasure struct (from d2datadict)
func (ig *ItemGenerator) ItemFromTreasure(treasure *d2datadict.Treasure) *Item {
	result := &Item{
		rand: rand.New(rand.NewSource(ig.Seed)),
	}

	// in this case, the treasure code is a code used by an ItemCommonRecord
	commonRecord := d2datadict.CommonItems[treasure.Code]
	if commonRecord != nil {
		result.CommonCode = commonRecord.Code
		return result
	}

	// next, we check if the treasure code is a generic type like `armo`
	equivList := d2datadict.ItemEquivalenciesByTypeCode[treasure.Code]
	if equivList != nil {
		result.CommonCode = equivList[ig.rand.Intn(len(equivList))].Code
		return result
	}

	// in this case, the treasure code is something like `armo23` and needs to
	// be resolved to ItemCommonRecords for armors with levels 23,24,25
	matches := resolveDynamicTreasureCode(treasure.Code)
	if matches != nil {
		numItems := len(matches)
		if numItems < 1 {
			return nil
		}

		result.CommonCode = matches[ig.rand.Intn(numItems)].Code

		return result
	}

	return nil
}

func resolveDynamicTreasureCode(code string) []*d2datadict.ItemCommonRecord {
	numericComponent := getNumericComponent(code)
	stringComponent := getStringComponent(code)

	if stringComponent == goldItemCodeWithMult {
		// todo need to do something with the numeric component (the gold multiplier)
		stringComponent = goldItemCode
	}

	result := make([]*d2datadict.ItemCommonRecord, 0)
	equivList := d2datadict.ItemEquivalenciesByTypeCode[stringComponent]

	for idx := range equivList {
		record := equivList[idx]
		minLevel := numericComponent
		maxLevel := minLevel + DynamicItemLevelRange

		if record.Level >= minLevel && record.Level < maxLevel {
			result = append(result, record)
		}
	}

	return result
}

func getStringComponent(code string) string {
	re := regexp.MustCompile(`\d+`)
	return string(re.ReplaceAll([]byte(code), []byte("")))
}

func getNumericComponent(code string) int {
	result := 0

	re := regexp.MustCompile(`[^\d]`)
	numStr := string(re.ReplaceAll([]byte(code), []byte("")))

	if number, err := strconv.ParseInt(numStr, 10, 32); err == nil {
		result = int(number)
	}

	return result
}
