// Package d2player contains the information necessary for managing players
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
func genEquipmentSlotsMap() map[d2enum.EquippedSlotType]EquipmentSlot {
	return map[d2enum.EquippedSlotType]EquipmentSlot{
		d2enum.LeftArm: {
			item:   nil,
			x:      418,
			y:      224,
			width:  61,
			height: 116,
		},
		d2enum.RightArm: {
			item:   nil,
			x:      648,
			y:      224,
			width:  61,
			height: 116,
		},
		d2enum.Head: {
			item:   nil,
			x:      532,
			y:      125,
			width:  62,
			height: 62,
		},
		d2enum.Neck: {
			item:   nil,
			x:      604,
			y:      125,
			width:  32,
			height: 32,
		},
		d2enum.Torso: {
			item:   nil,
			x:      532,
			y:      224,
			width:  62,
			height: 90,
		},
		d2enum.Belt: {
			item:   nil,
			x:      533,
			y:      269,
			width:  62,
			height: 32,
		},
		d2enum.LeftHand: {
			item:   nil,
			x:      491,
			y:      270,
			width:  32,
			height: 32,
		},
		d2enum.RightHand: {
			item:   nil,
			x:      606,
			y:      270,
			width:  32,
			height: 32,
		},
		d2enum.Gloves: {
			item:   nil,
			x:      417,
			y:      299,
			width:  62,
			height: 62,
		},
		d2enum.Legs: {
			item:   nil,
			x:      648,
			y:      299,
			width:  62,
			height: 62,
		},
	}
}
