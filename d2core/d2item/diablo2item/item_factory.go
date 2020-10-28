package diablo2item

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats/diablo2stats"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

const (
	defaultSeed = 0
)

const (
	dropModifierBaseProbability = 1024 // base dropModifier probability total
)

type dropModifier int

const (
	dropModifierNone dropModifier = iota
	dropModifierUnique
	dropModifierSet
	dropModifierRare
	dropModifierMagic
)

const (
	// dynamicItemLevelRange for treasure codes like `armo33`, this code is used to
	// select all equivalent items (matching `armo` in this case) with item levels 33,34,35
	dynamicItemLevelRange = 3
)

const (
	goldItemCodeWithMult = "gld,mul="
	goldItemCode         = "gld"
)

// NewItemFactory creates a new ItemFactory instance
func NewItemFactory(asset *d2asset.AssetManager) (*ItemFactory, error) {
	itemFactory := &ItemFactory{
		asset: asset,
		Seed:  0,
	}

	itemFactory.SetSeed(defaultSeed)

	statFactory, err := diablo2stats.NewStatFactory(asset)
	if err != nil {
		return nil, err
	}

	itemFactory.stat = statFactory

	return itemFactory, nil
}

// ItemFactory is a diablo 2 implementation of an item generator
type ItemFactory struct {
	asset  *d2asset.AssetManager
	stat   *diablo2stats.StatFactory
	rand   *rand.Rand
	source rand.Source
	Seed   int64
}

// SetSeed sets the item generator seed
func (f *ItemFactory) SetSeed(seed int64) {
	if f.rand == nil || f.source == nil {
		// nolint:gosec // we're not concerned with crypto-strong randomness
		f.rand = rand.New(rand.NewSource(seed))
	}

	f.Seed = seed
}

// NewItem creates a new item instance from the given codes
func (f *ItemFactory) NewItem(codes ...string) (*Item, error) {
	var common, set, unique string

	prefixes, suffixes := make([]string, 0), make([]string, 0)

	for _, code := range codes {
		if found := f.asset.Records.Item.All[code]; found != nil {
			common = code
			continue
		}

		if found := f.asset.Records.Item.SetItems[code]; found != nil {
			set = code
			continue
		}

		if found := f.asset.Records.Item.Unique[code]; found != nil {
			unique = code
			continue
		}

		if found := f.asset.Records.Item.Magic.Prefix[code]; found != nil {
			prefixes = append(prefixes, code)
			continue
		}

		if found := f.asset.Records.Item.Magic.Suffix[code]; found != nil {
			suffixes = append(suffixes, code)
			continue
		}
	}

	if common == "" {
		return nil, errors.New("cannot create item")
	}

	item := &Item{
		factory:    f,
		CommonCode: common,
	}

	if set != "" { // it's a set item
		item.SetItemCode = set
		return item.init(), nil
	}

	if unique != "" { // it's a unique item
		item.UniqueCode = unique
		return item.init(), nil
	}

	if len(prefixes) > 0 {
		item.PrefixCodes = prefixes
	}

	if len(suffixes) > 0 {
		item.SuffixCodes = suffixes
	}

	return item.init(), nil
}

// NewProperty creates a property
func (f *ItemFactory) NewProperty(code string, values ...int) *Property {
	record := f.asset.Records.Properties[code]

	if record == nil {
		return nil
	}

	result := &Property{
		factory:     f,
		record:      record,
		inputParams: values,
	}

	return result.init()
}

func (f *ItemFactory) rollDropModifier(tcr *d2records.TreasureClassRecord) dropModifier {
	modMap := map[int]dropModifier{
		0: dropModifierNone,
		1: dropModifierUnique,
		2: dropModifierSet,
		3: dropModifierRare,
		4: dropModifierMagic,
	}

	dropModifiers := []int{
		dropModifierBaseProbability,
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

	roll := f.rand.Intn(dropModifiers[len(dropModifiers)-1])

	for idx := range dropModifiers {
		if roll < dropModifiers[idx] {
			return modMap[idx]
		}
	}

	return dropModifierNone
}

func (f *ItemFactory) rollTreasurePick(tcr *d2records.TreasureClassRecord) *d2records.Treasure {
	// treasure probabilities
	tprob := make([]int, len(tcr.Treasures)+1)
	total := tcr.FreqNoDrop
	tprob[0] = total

	for idx := range tcr.Treasures {
		total += tcr.Treasures[idx].Probability
		tprob[idx+1] = total
	}

	roll := f.rand.Intn(total)

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
func (f *ItemFactory) ItemsFromTreasureClass(tcr *d2records.TreasureClassRecord) []*Item {
	result := make([]*Item, 0)

	treasurePicks := make([]*d2records.Treasure, 0)

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
			rolledTreasure := f.rollTreasurePick(tcr)

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
		if record, found := f.asset.Records.Item.Treasure.Normal[picked.Code]; found {
			// the code is for a treasure class, we roll again using that TC
			itemSlice := f.ItemsFromTreasureClass(record)
			for itemIdx := range itemSlice {
				itemSlice[itemIdx].applyDropModifier(f.rollDropModifier(tcr))
				itemSlice[itemIdx].init()
				result = append(result, itemSlice[itemIdx])
			}
		} else {
			// the code is not for a treasure class, but for an item
			item := f.ItemFromTreasure(picked)
			if item != nil {
				item.applyDropModifier(f.rollDropModifier(tcr))
				item.init()
				result = append(result, item)
			}
		}
	}

	return result
}

// ItemFromTreasure rolls for a f.rand.m item using the Treasure struct (from d2datadict)
func (f *ItemFactory) ItemFromTreasure(treasure *d2records.Treasure) *Item {
	result := &Item{
		// nolint:gosec // we're not concerned with crypto-strong randomness
		rand: rand.New(rand.NewSource(f.Seed)),
	}

	// in this case, the treasure code is a code used by an ItemCommonRecord
	commonRecord := f.asset.Records.Item.All[treasure.Code]
	if commonRecord != nil {
		result.CommonCode = commonRecord.Code
		return result
	}

	// next, we check if the treasure code is a generic type like `armo`
	equivList := f.asset.Records.Item.Equivalency[treasure.Code]
	if equivList != nil {
		result.CommonCode = equivList[f.rand.Intn(len(equivList))].Code
		return result
	}

	// in this case, the treasure code is something like `armo23` and needs to
	// be resolved to ItemCommonRecords for armors with levels 23,24,25
	matches := f.resolveDynamicTreasureCode(treasure.Code)
	if matches != nil {
		numItems := len(matches)
		if numItems < 1 {
			return nil
		}

		result.CommonCode = matches[f.rand.Intn(numItems)].Code

		return result
	}

	return nil
}

// FindMatchingAffixes for a given ItemCommonRecord, find all possible affixes that can spawn
func (f *ItemFactory) FindMatchingAffixes(
	icr *d2records.ItemCommonRecord,
	fromAffixes map[string]*d2records.ItemAffixCommonRecord,
) []*d2records.ItemAffixCommonRecord {
	result := make([]*d2records.ItemAffixCommonRecord, 0)

	equivItemTypes := f.asset.Records.FindEquivalentTypesByItemCommonRecord(icr)

	for prefixIdx := range fromAffixes {
		include, exclude := false, false
		affix := fromAffixes[prefixIdx]

		for itemTypeIdx := range equivItemTypes {
			itemType := equivItemTypes[itemTypeIdx]

			for _, excludedType := range affix.ItemExclude {
				if itemType == excludedType {
					exclude = true
					break
				}
			}

			if exclude {
				break
			}

			for _, includedType := range affix.ItemInclude {
				if itemType == includedType {
					include = true
					break
				}
			}

			if !include {
				continue
			}

			if icr.Level < affix.Level {
				continue
			}

			result = append(result, affix)
		}
	}

	return result
}

func (f *ItemFactory) resolveDynamicTreasureCode(code string) []*d2records.ItemCommonRecord {
	numericComponent := getNumericComponent(code)
	stringComponent := getStringComponent(code)

	if stringComponent == goldItemCodeWithMult {
		// need to do something with the numeric component (the gold multiplier)
		stringComponent = goldItemCode
	}

	result := make([]*d2records.ItemCommonRecord, 0)
	equivList := f.asset.Records.Item.Equivalency[stringComponent]

	for idx := range equivList {
		record := equivList[idx]
		minLevel := numericComponent
		maxLevel := minLevel + dynamicItemLevelRange

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

	re := regexp.MustCompile(`\D`)
	numStr := string(re.ReplaceAll([]byte(code), []byte("")))

	if number, err := strconv.ParseInt(numStr, 10, 32); err == nil {
		result = int(number)
	}

	return result
}
