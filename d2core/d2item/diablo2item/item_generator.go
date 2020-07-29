package diablo2item

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const (
	DropModifierBaseProbability = 1024 // base modifier probability total
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

// ItemGenerator is a diablo 2 implementation of an item generator
type ItemGenerator struct {
	Seed int64
}

// SetSeed sets the item generator seed
func (ig *ItemGenerator) SetSeed(seed int64) {
	rand.Seed(seed)
	ig.Seed = seed
}

func rollDropModifier(tcr *d2datadict.TreasureClassRecord) DropModifier {
	modMap := map[int]DropModifier{
		0 : DropModifierNone,
		1 : DropModifierUnique,
		2 : DropModifierSet,
		3 : DropModifierRare,
		4 : DropModifierMagic,
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

	roll := rand.Intn(dropModifiers[len(dropModifiers)-1])

	for idx := range dropModifiers {
		if roll < dropModifiers[idx] {
			return modMap[idx]
		}
	}

	return DropModifierNone
}

func rollTreasurePick(tcr *d2datadict.TreasureClassRecord) *d2datadict.Treasure {
	// treasure probabilities
	tprob := make([]int, len(tcr.Treasures)+1)
	total := tcr.FreqNoDrop
	tprob[0] = total

	for idx := range tcr.Treasures {
		total += tcr.Treasures[idx].Probability
		tprob[idx+1] = total
	}

	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(total)

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
			rolledTreasure := rollTreasurePick(tcr)

			if rolledTreasure == nil {
				continue
			}

			treasurePicks = append(treasurePicks, rolledTreasure)
		}
	}

	// for each of our picked/rolled treasures, we will attempt to generate an item
	// the treasure pick may actually be a reference to another treasure class, in which
	// case we will roll that treasure class, eventually getting items
	for idx := range treasurePicks {
		picked := treasurePicks[idx]
		if record, found := d2datadict.TreasureClass[picked.Code]; found {
			// the code is for a treasure class, we roll again using that TC
			// this may result in more than one item, as well
			moreTreasures := ig.ItemsFromTreasureClass(record)
			result = append(result, moreTreasures...)
		} else {
			// the code is not for a treasure class, but for an item!
			item := ig.ItemFromTreasure(picked)
			if item != nil {
				item.modifier = rollDropModifier(tcr)
				result = append(result, item)
			}
		}
	}

	return result
}

// ItemFromTreasure rolls for a random item using the Treasure struct (from d2datadict)
func (ig *ItemGenerator) ItemFromTreasure(treasure *d2datadict.Treasure) *Item {
	result := &Item{code: treasure.Code}

	commonRecord := d2datadict.CommonItems[treasure.Code]
	if commonRecord != nil {
		result.recordItemCommon = commonRecord
		return result
	}

	// doesn't have a ItemCommonRecord, it may be a generic (equivalency)
	equivList := d2datadict.ItemEquivalencies[treasure.Code]
	if equivList != nil {
		result.recordItemCommon = equivList[rand.Intn(len(equivList))] // random pick
		return result
	}

	dynamicList := resolveDynamicTreasureCode(treasure.Code)
	if dynamicList != nil {
		numItems := len(dynamicList)
		if numItems < 1 {
			return nil
		}

		result.recordItemCommon = dynamicList[rand.Intn(numItems)] // random pick

		return result
	}

	return nil
}

func resolveDynamicTreasureCode(code string) []*d2datadict.ItemCommonRecord {
	numericComponent := getNumericComponentFromCode(code)
	stringComponent := getStringComponentFromCode(code)

	result := make([]*d2datadict.ItemCommonRecord, 0)
	equivList := d2datadict.ItemEquivalencies[stringComponent]

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

func getStringComponentFromCode(code string) string {
	re := regexp.MustCompile(`\d+`)
	return string(re.ReplaceAll([]byte(code), []byte("")))
}

func getNumericComponentFromCode(code string) int {
	result := 0

	re := regexp.MustCompile(`[^\d]`)
	numStr := string(re.ReplaceAll([]byte(code), []byte("")))

	if number, err := strconv.ParseInt(numStr, 10, 32); err == nil {
		result = int(number)
	}

	return result
}
