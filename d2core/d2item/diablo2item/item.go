package diablo2item

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// PropertyPool is used for separating properties by their source
type PropertyPool int

const (
	PropertyPoolAffix PropertyPool = iota
	PropertyPoolAutoMagic
	PropertyPoolUnique
	PropertyPoolSetItem
	PropertyPoolSet
)

// static check to ensure Item implements Item
var _ d2item.Item = &Item{}

type Item struct {
	code     string
	modifier DropModifier
	slotType d2enum.EquippedSlot
	itemType d2enum.InventoryItemType

	recordItemType   *d2datadict.ItemTypeRecord
	recordItemCommon *d2datadict.ItemCommonRecord
	recordItemUnique *d2datadict.UniqueItemRecord
	recordSet        *d2datadict.SetRecord
	recordSetItem    *d2datadict.SetItemRecord

	recordPrefix [3]*d2datadict.ItemAffixCommonRecord
	recordSuffix [3]*d2datadict.ItemAffixCommonRecord

	propertyPools  map[PropertyPool][]*Property
	statContext    d2item.StatContext
	statList       d2stats.StatList
	BaseStatList   d2stats.StatList
	UniqueStatList d2stats.StatList

	indestructable bool
	ethereal       bool
	itemLevel      int

	numSockets int
	sockets    []*d2item.Item // there will be checks for handling the craziness this might entail

	worldSprite     *d2ui.Sprite
	inventorySprite *d2ui.Sprite
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
func (i *Item) ItemType() d2enum.InventoryItemType {
	return i.itemType
}

// ItemLevel returns the level of item
func (i *Item) ItemLevel() int {
	return i.itemLevel
}

// CommonRecord returns the ItemCommonRecord of the item
func (i *Item) CommonRecord() *d2datadict.ItemCommonRecord {
	return i.recordItemCommon
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
