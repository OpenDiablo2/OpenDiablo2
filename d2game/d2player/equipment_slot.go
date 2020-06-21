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
			x:    549,
			y:    133,
		},
		"Amulet": {
			item: nil,
			x:    613,
			y:    103,
		},
		"Chest": {
			item: nil,
			x:    546,
			y:    205,
		},
		"Belt": {
			item: nil,
			x:    548,
			y:    250,
		},
		"LeftRing": {
			item: nil,
			x:    501,
			y:    250,
		},
		"RightRing": {
			item: nil,
			x:    613,
			y:    250,
		},
		"Gloves": {
			item: nil,
			x:    429,
			y:    252,
		},
		"Shoes": {
			item: nil,
			x:    661,
			y:    252,
		},
	}
}
