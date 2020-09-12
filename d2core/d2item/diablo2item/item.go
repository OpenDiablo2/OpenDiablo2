package diablo2item

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats/diablo2stats"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// PropertyPool is used for separating properties by their source
type PropertyPool int

// Property pools
const (
	PropertyPoolPrefix PropertyPool = iota
	PropertyPoolSuffix
	PropertyPoolUnique
	PropertyPoolSetItem
	PropertyPoolSet
)

// for handling special cases
const (
	jewelItemCode          = "jew"
	propertyEthereal       = "ethereal"
	propertyIndestructable = "indestruct"
)

const (
	magicItemPrefixMax = 1
	magicItemSuffixMax = 1
	rareItemPrefixMax  = 3
	rareItemSuffixMax  = 3
	rareJewelPrefixMax = 3
	rareJewelSuffixMax = 3
	rareJewelAffixMax  = 4
)

// static check to ensure Item implements Item
var _ d2item.Item = &Item{}

type Item struct {
	name string
	Seed int64
	rand *rand.Rand // non-global rand instance for re-generating the item

	slotType d2enum.EquippedSlot

	TypeCode    string
	CommonCode  string
	UniqueCode  string
	SetCode     string
	SetItemCode string
	PrefixCodes []string
	SuffixCodes []string

	properties      map[PropertyPool][]*Property
	statContext     d2item.StatContext
	statList        d2stats.StatList
	uniqueStatList  d2stats.StatList
	setItemStatList d2stats.StatList

	attributes *itemAttributes

	GridX int
	GridY int

	sockets []*d2item.Item // there will be checks for handling the craziness this might entail
}

type itemAttributes struct {
	worldSprite     *d2ui.Sprite
	inventorySprite *d2ui.Sprite

	damageOneHand minMaxEnhanceable
	damageTwoHand minMaxEnhanceable
	damageMissile minMaxEnhanceable
	stackSize     minMaxEnhanceable
	durability    minMaxEnhanceable

	personalization string

	quality                 int
	defense                 int
	currentStackSize        int
	currentDurability       int
	baseItemLevel           int
	requiredLevel           int
	numSockets              int
	requirementsEnhancement int
	requiredStrength        int
	requiredDexterity       int
	classSpecific           d2enum.Hero

	identitified   bool
	crafted        bool
	durable        bool // some items specify that they have no durability
	indestructable bool
	ethereal       bool
	throwable      bool
}

type minMaxEnhanceable struct {
	min     int
	max     int
	enhance int
}

// Label returns the item name
func (i *Item) Label() string {
	str := i.name

	if !i.attributes.identitified {
		str = d2tbl.TranslateString(i.CommonRecord().NameString)
	}

	if i.attributes.crafted {
		return d2ui.ColorTokenize(str, d2ui.ColorTokenCraftedItem)
	}

	if i.SetItemRecord() != nil {
		return d2ui.ColorTokenize(str, d2ui.ColorTokenSetItem)
	}

	if i.UniqueRecord() != nil {
		return d2ui.ColorTokenize(str, d2ui.ColorTokenUniqueItem)
	}

	numAffixes := len(i.PrefixRecords()) + len(i.SuffixRecords())

	if numAffixes > 0 && numAffixes < 3 {
		return d2ui.ColorTokenize(str, d2ui.ColorTokenMagicItem)
	}

	if numAffixes > 2 {
		return d2ui.ColorTokenize(str, d2ui.ColorTokenRareItem)
	}

	if i.sockets != nil {
		if len(i.sockets) > 0 {
			return d2ui.ColorTokenize(str, d2ui.ColorTokenSocketedItem)
		}
	}

	return d2ui.ColorTokenize(str, d2ui.ColorTokenNormalItem)
}

// Context returns the statContext that is being used to evaluate stats. for example,
// stats which are based on character level will be evaluated with the player
// as the statContext, as the player stat list will contain stats that describe the
// character level
func (i *Item) Context() d2item.StatContext {
	return i.statContext
}

// SetContext sets the statContext for evaluating item stats
func (i *Item) SetContext(ctx d2item.StatContext) {
	i.statContext = ctx
}

// ItemType returns the type of item
func (i *Item) ItemType() string {
	return i.TypeCode
}

// ItemLevel returns the level of item
func (i *Item) ItemLevel() int {
	return i.attributes.baseItemLevel
}

// TypeRecord returns the ItemTypeRecord of the item
func (i *Item) TypeRecord() *d2datadict.ItemTypeRecord {
	return d2datadict.ItemTypes[i.TypeCode]
}

// CommonRecord returns the ItemCommonRecord of the item
func (i *Item) CommonRecord() *d2datadict.ItemCommonRecord {
	return d2datadict.CommonItems[i.CommonCode]
}

// UniqueRecord returns the UniqueItemRecord of the item
func (i *Item) UniqueRecord() *d2datadict.UniqueItemRecord {
	return d2datadict.UniqueItems[i.UniqueCode]
}

// SetRecord returns the SetRecord of the item
func (i *Item) SetRecord() *d2datadict.SetRecord {
	return d2datadict.SetRecords[i.SetCode]
}

// SetItemRecord returns the SetRecord of the item
func (i *Item) SetItemRecord() *d2datadict.SetItemRecord {
	return d2datadict.SetItems[i.SetItemCode]
}

// PrefixRecords returns the ItemAffixCommonRecords of the prefixes of the item
func (i *Item) PrefixRecords() []*d2datadict.ItemAffixCommonRecord {
	return affixRecords(i.PrefixCodes, d2datadict.MagicPrefix)
}

// SuffixRecords returns the ItemAffixCommonRecords of the prefixes of the item
func (i *Item) SuffixRecords() []*d2datadict.ItemAffixCommonRecord {
	return affixRecords(i.SuffixCodes, d2datadict.MagicSuffix)
}

func affixRecords(
	fromCodes []string,
	affixes map[string]*d2datadict.ItemAffixCommonRecord,
) []*d2datadict.ItemAffixCommonRecord {
	if len(fromCodes) < 1 {
		return nil
	}

	result := make([]*d2datadict.ItemAffixCommonRecord, len(fromCodes))

	for idx, code := range fromCodes {
		rec := affixes[code]
		result[idx] = rec
	}

	return result
}

// SlotType returns the slot type (where it can be equipped)
func (i *Item) SlotType() d2enum.EquippedSlot {
	return i.slotType
}

// StatList returns the evaluated stat list
func (i *Item) StatList() d2stats.StatList {
	return i.statList
}

// Description returns the full description string for the item
func (i *Item) Description() string {
	return ""
}

// applyDropModifier attempts to find the necessary set, unique, or
// affix records, depending on the drop modifier given. If an unsupported
// drop modifier is supplied, it will attempt to reconcile by picked
// magic affixes as if it were a rare.
func (i *Item) applyDropModifier(modifier DropModifier) {
	modifier = i.sanitizeDropModifier(modifier)

	switch modifier {
	case DropModifierUnique:
		i.pickUniqueRecord()

		if i.UniqueRecord() == nil {
			i.applyDropModifier(DropModifierRare)
			return
		}
	case DropModifierSet:
		i.pickSetRecords()

		if i.SetRecord() == nil || i.SetItemRecord() == nil {
			i.applyDropModifier(DropModifierRare)
			return
		}
	case DropModifierRare, DropModifierMagic:
		// the method of picking stays the same for magic/rare
		// but magic gets to pick more, and jewels have a special
		// way of picking affixes
		i.pickMagicAffixes(modifier)
	case DropModifierNone:
	default:
		return
	}
}

func (i *Item) sanitizeDropModifier(modifier DropModifier) DropModifier {
	if i.TypeRecord() == nil {
		i.TypeCode = i.CommonRecord().Type
	}

	// should this item always be normal?
	if i.TypeRecord().Normal {
		modifier = DropModifierNone
	}

	// should this item always be magic?
	if i.TypeRecord().Magic {
		modifier = DropModifierMagic
	}

	// if it isn't allowed to be rare, force it to be magic
	if modifier == DropModifierRare && !i.TypeRecord().Rare {
		modifier = DropModifierMagic
	}

	return modifier
}

func (i *Item) pickUniqueRecord() {
	matches := findMatchingUniqueRecords(i.CommonRecord())
	if len(matches) > 0 {
		match := matches[i.rand.Intn(len(matches))]
		i.UniqueCode = match.Code
	}
}

func (i *Item) pickSetRecords() {
	if matches := findMatchingSetItemRecords(i.CommonRecord()); len(matches) > 0 {
		picked := matches[i.rand.Intn(len(matches))]
		i.SetItemCode = picked.SetItemKey

		if rec := i.SetItemRecord(); rec != nil {
			i.SetCode = rec.SetKey
		}
	}
}

func (i *Item) pickMagicAffixes(mod DropModifier) {
	if i.PrefixCodes == nil {
		i.PrefixCodes = make([]string, 0)
	}

	if i.SuffixCodes == nil {
		i.SuffixCodes = make([]string, 0)
	}

	totalAffixes, numSuffixes, numPrefixes := 0, 0, 0

	switch mod {
	case DropModifierRare:
		if i.CommonRecord().Type == jewelItemCode {
			numPrefixes, numSuffixes = rareJewelPrefixMax, rareJewelSuffixMax
			totalAffixes = rareJewelAffixMax
		} else {
			numPrefixes, numSuffixes = rareItemPrefixMax, rareItemSuffixMax
			totalAffixes = numPrefixes + numSuffixes
		}
	case DropModifierMagic:
		numPrefixes, numSuffixes = magicItemPrefixMax, magicItemSuffixMax
		totalAffixes = numPrefixes + numSuffixes
	}

	i.PrefixCodes = i.pickRandomAffixes(numPrefixes, totalAffixes, d2datadict.MagicPrefix)
	i.SuffixCodes = i.pickRandomAffixes(numSuffixes, totalAffixes, d2datadict.MagicSuffix)
}

func (i *Item) pickRandomAffixes(max, totalMax int, affixMap map[string]*d2datadict.ItemAffixCommonRecord) []string {
	pickedCodes := make([]string, 0)

	for numPicks := 0; numPicks < max; numPicks++ {
		matches := findMatchingAffixes(i.CommonRecord(), affixMap)

		if rollPrefix := i.rand.Intn(2); rollPrefix > 0 {
			affixCount := len(i.PrefixRecords()) + len(i.SuffixRecords())
			if len(i.PrefixRecords()) > max || affixCount > totalMax {
				break
			}

			if len(matches) > 0 {
				picked := matches[i.rand.Intn(len(matches))]
				pickedCodes = append(pickedCodes, picked.Name)
			}
		}
	}

	return pickedCodes
}

// SetSeed sets the item generator seed
func (i *Item) SetSeed(seed int64) {
	if i.rand == nil {
		source := rand.NewSource(seed)
		i.rand = rand.New(source)
	}

	i.Seed = seed
}

func (i *Item) init() *Item {
	if i.rand == nil {
		i.SetSeed(0)
	}

	i.generateAllProperties()
	i.updateItemAttributes()

	return i
}

func (i *Item) generateAllProperties() {
	if i.attributes == nil {
		i.attributes = &itemAttributes{}
	}

	// these will get updated by any generated properties
	i.attributes.ethereal = false
	i.attributes.indestructable = false

	pools := []PropertyPool{
		PropertyPoolPrefix,
		PropertyPoolSuffix,
		PropertyPoolUnique,
		PropertyPoolSetItem,
		PropertyPoolSet,
	}

	for _, pool := range pools {
		i.generateProperties(pool)
	}
}

func (i *Item) generateProperties(pool PropertyPool) {
	var props []*Property

	switch pool {
	case PropertyPoolPrefix, PropertyPoolSuffix:
		if generated := i.generateAffixProperties(pool); generated != nil {
			props = generated
		}
	case PropertyPoolUnique:
		if generated := i.generateUniqueProperties(); generated != nil {
			props = generated
		}
	case PropertyPoolSetItem:
		if generated := i.generateSetItemProperties(); generated != nil {
			props = generated
		}
	case PropertyPoolSet: // todo set bonus handling, needs player/equipment context
	}

	if props == nil {
		return
	}

	if i.properties == nil {
		i.properties = make(map[PropertyPool][]*Property)
	}

	i.properties[pool] = props

	// in the case one of the properties is a stat-less prop for indestructable/ethereal
	// we need to set the item attributes to the rolled values. we use `||` here just in
	// case another property has already set the flag
	for propIdx := range props {
		prop := props[propIdx]
		switch prop.record.Code {
		case propertyEthereal:
			i.attributes.ethereal = i.attributes.ethereal || prop.computedBool
		case propertyIndestructable:
			i.attributes.indestructable = i.attributes.ethereal || prop.computedBool
		}
	}
}

func (i *Item) updateItemAttributes() {
	i.generateName()

	r := i.CommonRecord()
	i.attributes = &itemAttributes{
		damageOneHand: minMaxEnhanceable{
			min: r.MinDamage,
			max: r.MaxDamage,
		},

		damageTwoHand: minMaxEnhanceable{
			min: r.Min2HandDamage,
			max: r.Max2HandDamage,
		},

		damageMissile: minMaxEnhanceable{
			min: r.MinMissileDamage,
			max: r.MaxMissileDamage,
		},
		stackSize: minMaxEnhanceable{
			min: r.MinStack,
			max: r.MaxStack,
		},
		durability: minMaxEnhanceable{
			min: r.Durability,
			max: r.Durability,
		},

		baseItemLevel:     r.Level,
		requiredLevel:     r.RequiredLevel,
		requiredStrength:  r.RequiredStrength,
		requiredDexterity: r.RequiredDexterity,
		durable:           !r.NoDurability,
		throwable:         r.Throwable,
	}

	def, minDef, maxDef := 0, r.MinAC, r.MaxAC

	if maxDef < minDef {
		minDef, maxDef = maxDef, minDef
	}

	if minDef > 1 && maxDef > 1 {
		def = i.rand.Intn(maxDef-minDef+1) + minDef
	}

	i.attributes.defense = def
}

func (i *Item) generateAffixProperties(pool PropertyPool) []*Property {
	var affixRecords []*d2datadict.ItemAffixCommonRecord

	switch pool {
	case PropertyPoolPrefix:
		affixRecords = i.PrefixRecords()
	case PropertyPoolSuffix:
		affixRecords = i.SuffixRecords()
	default:
		return nil
	}

	if affixRecords == nil || len(affixRecords) < 1 {
		return nil
	}

	result := make([]*Property, 0)

	// for each prefix
	for recIdx := range affixRecords {
		affix := affixRecords[recIdx]
		// for each modifier
		for modIdx := range affix.Modifiers {
			mod := affix.Modifiers[modIdx]

			prop := NewProperty(mod.Code, mod.Parameter, mod.Min, mod.Max)
			if prop == nil {
				continue
			}

			result = append(result, prop)
		}
	}

	return result
}

func (i *Item) generateUniqueProperties() []*Property {
	if i.UniqueRecord() == nil {
		return nil
	}

	result := make([]*Property, 0)

	for propIdx := range i.UniqueRecord().Properties {
		propInfo := i.UniqueRecord().Properties[propIdx]

		// sketchy ass unique records, the param should be an int but sometimes it's the name
		// of a skill, which needs to be converted to the skill index
		paramStr := getStringComponent(propInfo.Parameter)
		paramInt := getNumericComponent(propInfo.Parameter)

		if paramStr != "" {
			for skillID := range d2datadict.SkillDetails {
				if d2datadict.SkillDetails[skillID].Skill == paramStr {
					paramInt = skillID
				}
			}
		}

		prop := NewProperty(propInfo.Code, paramInt, propInfo.Min, propInfo.Max)
		if prop == nil {
			continue
		}

		result = append(result, prop)
	}

	return result
}

func (i *Item) generateSetItemProperties() []*Property {
	if i.SetItemRecord() == nil {
		return nil
	}

	result := make([]*Property, 0)

	for propIdx := range i.SetItemRecord().Properties {
		setProp := i.SetItemRecord().Properties[propIdx]

		// like with unique records, the property param is sometimes a skill name
		// as a string, not an integer index
		paramStr := getStringComponent(setProp.Parameter)
		paramInt := getNumericComponent(setProp.Parameter)

		if paramStr != "" {
			for skillID := range d2datadict.SkillDetails {
				if d2datadict.SkillDetails[skillID].Skill == paramStr {
					paramInt = skillID
				}
			}
		}

		prop := NewProperty(setProp.Code, paramInt, setProp.Min, setProp.Max)
		if prop == nil {
			continue
		}

		result = append(result, prop)
	}

	return result
}

func (i *Item) generateName() {
	if i.SetItemRecord() != nil {
		i.name = d2tbl.TranslateString(i.SetItemRecord().SetItemKey)
		return
	}

	if i.UniqueRecord() != nil {
		i.name = d2tbl.TranslateString(i.UniqueRecord().Name)
		return
	}

	name := d2tbl.TranslateString(i.CommonRecord().NameString)

	numAffixes := 0
	if prefixes := i.PrefixRecords(); prefixes != nil {
		numAffixes += len(prefixes)
	}

	if suffixes := i.SuffixRecords(); suffixes != nil {
		numAffixes += len(suffixes)
	}

	// if it has 1 to 2 affixes, it's a magic item, and we just put the current item
	// name between the prefix and suffix strings
	if numAffixes > 0 && numAffixes < 3 {
		if len(i.PrefixRecords()) > 0 {
			affix := i.PrefixRecords()[i.rand.Intn(len(i.PrefixRecords()))]
			name = fmt.Sprintf("%s %s", affix.Name, name)
		}

		if len(i.SuffixRecords()) > 0 {
			affix := i.SuffixRecords()[i.rand.Intn(len(i.SuffixRecords()))]
			name = fmt.Sprintf("%s %s", name, affix.Name)
		}
	}

	// if it has more than 2 affixes, it's a rare item
	// rare items use entries from rareprefix.txt and raresuffix.txt to make their names,
	// and the prefix and suffix actually go before thec current item name
	if numAffixes >= 3 {
		i.rand.Seed(i.Seed)

		numPrefix, numSuffix := len(d2datadict.RarePrefixes), len(d2datadict.RareSuffixes)
		preIdx, sufIdx := i.rand.Intn(numPrefix), i.rand.Intn(numSuffix)
		prefix := d2datadict.RarePrefixes[preIdx].Name
		suffix := d2datadict.RareSuffixes[sufIdx].Name
		name = fmt.Sprintf("%s %s\n%s", strings.Title(prefix), strings.Title(suffix), name)
	}

	i.name = name
}

// GetStatStrings is a test function for getting all stat strings
func (i *Item) GetStatStrings() []string {
	result := make([]string, 0)
	stats := make([]d2stats.Stat, 0)

	for pool := range i.properties {
		propPool := i.properties[pool]
		if propPool == nil {
			continue
		}

		for propIdx := range propPool {
			if propPool[propIdx] == nil {
				continue
			}

			prop := propPool[propIdx]

			for statIdx := range prop.stats {
				stats = append(stats, prop.stats[statIdx])
			}
		}
	}

	if len(stats) > 0 {
		stats = diablo2stats.NewStatList(stats...).ReduceStats().Stats()
	}

	sort.Slice(stats, func(i, j int) bool { return stats[i].Priority() > stats[j].Priority() })

	for statIdx := range stats {
		statStr := stats[statIdx].String()
		if statStr != "" {
			result = append(result, statStr)
		}
	}

	return result
}

func findMatchingUniqueRecords(icr *d2datadict.ItemCommonRecord) []*d2datadict.UniqueItemRecord {
	result := make([]*d2datadict.UniqueItemRecord, 0)

	c1, c2, c3, c4 := icr.Code, icr.NormalCode, icr.UberCode, icr.UltraCode

	for uCode := range d2datadict.UniqueItems {
		uRec := d2datadict.UniqueItems[uCode]

		switch uCode {
		case c1, c2, c3, c4:
			result = append(result, uRec)
		}
	}

	return result
}

// find possible SetItemRecords that the given ItemCommonRecord can have
func findMatchingSetItemRecords(icr *d2datadict.ItemCommonRecord) []*d2datadict.SetItemRecord {
	result := make([]*d2datadict.SetItemRecord, 0)

	c1, c2, c3, c4 := icr.Code, icr.NormalCode, icr.UberCode, icr.UltraCode

	for setItemIdx := range d2datadict.SetItems {
		switch d2datadict.SetItems[setItemIdx].ItemCode {
		case c1, c2, c3, c4:
			result = append(result, d2datadict.SetItems[setItemIdx])
		}
	}

	return result
}

// for a given ItemCommonRecord, find all possible affixes that can spawn
func findMatchingAffixes(
	icr *d2datadict.ItemCommonRecord,
	fromAffixes map[string]*d2datadict.ItemAffixCommonRecord,
) []*d2datadict.ItemAffixCommonRecord {
	result := make([]*d2datadict.ItemAffixCommonRecord, 0)

	equivItemTypes := d2datadict.FindEquivalentTypesByItemCommonRecord(icr)

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

// these functions are to satisfy the inventory grid item interface

// GetInventoryItemName returns the item name
func (i *Item) GetInventoryItemName() string {
	return i.Label()
}

// GetInventoryItemType returns whether the item is a weapon, armor, or misc item
func (i *Item) GetInventoryItemType() d2enum.InventoryItemType {
	typeCode := i.TypeRecord().Code

	armorEquiv := d2datadict.ItemEquivalenciesByTypeCode["armo"]
	weaponEquiv := d2datadict.ItemEquivalenciesByTypeCode["weap"]

	for idx := range armorEquiv {
		if armorEquiv[idx].Code == typeCode {
			return d2enum.InventoryItemTypeArmor
		}
	}

	for idx := range weaponEquiv {
		if weaponEquiv[idx].Code == typeCode {
			return d2enum.InventoryItemTypeWeapon
		}
	}

	return d2enum.InventoryItemTypeItem
}

func (i *Item) InventoryGridSize() (width, height int) {
	r := i.CommonRecord()
	return r.InventoryWidth, r.InventoryHeight
}

func (i *Item) GetItemCode() string {
	return i.CommonRecord().Code
}

func (i *Item) Serialize() []byte {
	panic("item serialization not yet implemented")
}

func (i *Item) InventoryGridSlot() (x, y int) {
	return i.GridX, i.GridY
}

func (i *Item) SetInventoryGridSlot(x, y int) {
	i.GridX, i.GridY = x, y
}

func (i *Item) GetInventoryGridSize() (x, y int) {
	return i.GridX, i.GridY
}

func (i *Item) Identify() *Item {
	i.attributes.identitified = true
	return i
}

// from a string table
const (
	reqNotMet    = "ItemStats1a" // "Requirements not met",
	unidentified = "ItemStats1b" // "Unidentified",
	charges      = "ItemStats1c" // "Charges:",
	durability   = "ItemStats1d" // "Durability:",
	reqStrength  = "ItemStats1e" // "Required Strength:",
	reqDexterity = "ItemStats1f" // "Required Dexterity:",
	damage       = "ItemStats1g" // "Damage:",
	defense      = "ItemStats1h" // "Defense:",
	quantity     = "ItemStats1i" // "Quantity:",
	of           = "ItemStats1j" // "of",
	to           = "to"          // "to"
	damage1h     = "ItemStats1l" // "One-Hand Damage:",
	damage2h     = "ItemStats1m" // "Two-Hand Damage:",
	damageThrow  = "ItemStats1n" // "Throw Damage:",
	damageSmite  = "ItemStats1o" // "Smite Damage:",
	reqLevel     = "ItemStats1p" // "Required Level:",
)

func (i *Item) GetItemDescription() []string {
	lines := make([]string, 0)

	common := i.CommonRecord()

	lines = append(lines, i.Label())

	str := ""

	if common.MinAC > 0 {
		min, max := common.MinAC, common.MaxAC
		str = fmt.Sprintf("%s %v %s %v", d2tbl.TranslateString(defense), min, d2tbl.TranslateString(to), max)
		str = d2ui.ColorTokenize(str, d2ui.ColorTokenWhite)
		lines = append(lines, str)
	}

	if common.MinDamage > 0 {
		min, max := common.MinDamage, common.MaxDamage
		str = fmt.Sprintf("%s %v %s %v", d2tbl.TranslateString(damage1h), min, d2tbl.TranslateString(to), max)
		str = d2ui.ColorTokenize(str, d2ui.ColorTokenWhite)
		lines = append(lines, str)
	}

	if common.Min2HandDamage > 0 {
		min, max := common.Min2HandDamage, common.Max2HandDamage
		str = fmt.Sprintf("%s %v %s %v", d2tbl.TranslateString(damage2h), min, d2tbl.TranslateString(to), max)
		str = d2ui.ColorTokenize(str, d2ui.ColorTokenWhite)
		lines = append(lines, str)
	}

	if common.MinMissileDamage > 0 {
		min, max := common.MinMissileDamage, common.MaxMissileDamage
		str = fmt.Sprintf("%s %v %s %v", d2tbl.TranslateString(damageThrow), min, d2tbl.TranslateString(to), max)
		str = d2ui.ColorTokenize(str, d2ui.ColorTokenWhite)
		lines = append(lines, str)
	}

	if common.RequiredStrength > 1 {
		str = fmt.Sprintf("%s %v", d2tbl.TranslateString(reqStrength), common.RequiredStrength)
		str = d2ui.ColorTokenize(str, d2ui.ColorTokenWhite)
		lines = append(lines, str)
	}

	if common.RequiredDexterity > 1 {
		str = fmt.Sprintf("%s %v", d2tbl.TranslateString(reqDexterity), common.RequiredDexterity)
		str = d2ui.ColorTokenize(str, d2ui.ColorTokenWhite)
		lines = append(lines, str)
	}

	if common.RequiredLevel > 1 {
		str = fmt.Sprintf("%s %v", d2tbl.TranslateString(reqLevel), common.RequiredLevel)
		str = d2ui.ColorTokenize(str, d2ui.ColorTokenWhite)
		lines = append(lines, str)
	}

	statStrings := i.GetStatStrings()

	for _, statStr := range statStrings {
		str = d2ui.ColorTokenize(statStr, d2ui.ColorTokenBlue)
		lines = append(lines, str)
	}

	return lines
}
