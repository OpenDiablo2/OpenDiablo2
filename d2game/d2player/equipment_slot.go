package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// EquipmentSlot represents an equipment slot for a player
type EquipmentSlot struct {
	item   InventoryItem
	x      int
	y      int
	width  int
	height int
}

func genEquipmentSlotsMap(record *d2datadict.InventoryRecord) map[d2enum.EquippedSlot]EquipmentSlot {
	slotMap := map[d2enum.EquippedSlot]EquipmentSlot{}

	slots := []d2enum.EquippedSlot{
		d2enum.EquippedSlotHead,
		d2enum.EquippedSlotTorso,
		d2enum.EquippedSlotLegs,
		d2enum.EquippedSlotRightArm,
		d2enum.EquippedSlotLeftArm,
		d2enum.EquippedSlotLeftHand,
		d2enum.EquippedSlotRightHand,
		d2enum.EquippedSlotNeck,
		d2enum.EquippedSlotBelt,
		d2enum.EquippedSlotGloves,
	}
	
	for _, slot := range slots {
		box := record.Slots[slot]
		equipmentSlot := EquipmentSlot{
			nil,
			box.Left,
			box.Bottom + cellPadding,
			box.Width,
			box.Height,
		}
		slotMap[slot] = equipmentSlot
	}

	return slotMap
}
