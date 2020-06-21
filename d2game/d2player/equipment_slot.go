package d2player

type EquipmentSlot struct {
	item InventoryItem
	x    int
	y    int
}

func genEquipmentSlotsMap() map[string]EquipmentSlot {
	return map[string]EquipmentSlot{
		"LeftHand": {
			item: nil,
			x:    438,
			y:    200,
		},
		"RightHand": {
			item: nil,
			x:    665,
			y:    200,
		},
		"Helmet": {
			item: nil,
			x:    536,
			y:    122,
		},
		"Amulet": {
			item: nil,
			x:    609,
			y:    122,
		},
		"Chest": {
			item: nil,
			x:    536,
			y:    205,
		},
		"Belt": {
			item: nil,
			x:    536,
			y:    270,
		},
		"LeftRing": {
			item: nil,
			x:    491,
			y:    270,
		},
		"RightRing": {
			item: nil,
			x:    609,
			y:    270,
		},
		"Gloves": {
			item: nil,
			x:    420,
			y:    295,
		},
		"Shoes": {
			item: nil,
			x:    652,
			y:    295,
		},
	}
}
