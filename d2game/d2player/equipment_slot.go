package d2player

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

type EquipmentSlot struct {
	item   InventoryItem
	x      int
	y      int
	width  int
	height int
}

func genEquipmentSlotsMap() map[d2enum.EquippedSlotType]EquipmentSlot {
	return map[d2enum.EquippedSlotType]EquipmentSlot{
		d2enum.LeftArm: {
			item: nil,
			x:    438,
			y:    200,
		},
		d2enum.RightArm: {
			item: nil,
			x:    665,
			y:    200,
		},
		d2enum.Head: {
			item: nil,
			x:    536,
			y:    122,
		},
		d2enum.Neck: {
			item: nil,
			x:    609,
			y:    122,
		},
		d2enum.Torso: {
			item: nil,
			x:    536,
			y:    205,
		},
		d2enum.Belt: {
			item: nil,
			x:    536,
			y:    270,
		},
		d2enum.LeftHand: {
			item: nil,
			x:    491,
			y:    270,
		},
		d2enum.RightHand: {
			item: nil,
			x:    609,
			y:    270,
		},
		d2enum.Gloves: {
			item: nil,
			x:    420,
			y:    295,
		},
		d2enum.Legs: {
			item: nil,
			x:    652,
			y:    295,
		},
	}
}
