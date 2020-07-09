package d2player

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// EquipmentSlot represents an equipment slot for a player
type EquipmentSlot struct {
	item   InventoryItem
	x      int
	y      int
	width  int
	height int
}

//nolint:gomnd Magic numbers are necessary for this file
func genEquipmentSlotsMap() map[d2enum.EquippedSlot]EquipmentSlot {
	return map[d2enum.EquippedSlot]EquipmentSlot{
		d2enum.EquippedSlotLeftArm: {
			item:   nil,
			x:      418,
			y:      224,
			width:  61,
			height: 116,
		},
		d2enum.EquippedSlotRightArm: {
			item:   nil,
			x:      648,
			y:      224,
			width:  61,
			height: 116,
		},
		d2enum.EquippedSlotHead: {
			item:   nil,
			x:      532,
			y:      125,
			width:  62,
			height: 62,
		},
		d2enum.EquippedSlotNeck: {
			item:   nil,
			x:      604,
			y:      125,
			width:  32,
			height: 32,
		},
		d2enum.EquippedSlotTorso: {
			item:   nil,
			x:      532,
			y:      224,
			width:  62,
			height: 90,
		},
		d2enum.EquippedSlotBelt: {
			item:   nil,
			x:      533,
			y:      269,
			width:  62,
			height: 32,
		},
		d2enum.EquippedSlotLeftHand: {
			item:   nil,
			x:      491,
			y:      270,
			width:  32,
			height: 32,
		},
		d2enum.EquippedSlotRightHand: {
			item:   nil,
			x:      606,
			y:      270,
			width:  32,
			height: 32,
		},
		d2enum.EquippedSlotGloves: {
			item:   nil,
			x:      417,
			y:      299,
			width:  62,
			height: 62,
		},
		d2enum.EquippedSlotLegs: {
			item:   nil,
			x:      648,
			y:      299,
			width:  62,
			height: 62,
		},
	}
}
