package d2item

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// Item describes all types of item that can be placed in the
// player inventory grid (not just things that can be equipped!)
type Item interface {
	Context() StatContext
	SetContext(StatContext)

	ItemType() d2enum.InventoryItemType
	SlotType() d2enum.EquippedSlot
	StatList() d2stats.StatList

	Description() string
}
