package d2player

type EquipmentSlot struct {
	Item InventoryItem
	x    int
	y    int
}

func genEquipmentSlotsMap() map[string]EquipmentSlot {
	return map[string]EquipmentSlot{
		"LeftHand": {
			Item: nil,
			x:    438,
			y:    200,
		},
		"RightHand": {
			Item: nil,
			x:    665,
			y:    200,
		},
		"Helmet": {
			Item: nil,
			x:    549,
			y:    133,
		},
		"Amulet": {
			Item: nil,
			x:    613,
			y:    103,
		},
		"Chest": {
			Item: nil,
			x:    546,
			y:    205,
		},
		"Belt": {
			Item: nil,
			x:    548,
			y:    250,
		},
		"LeftRing": {
			Item: nil,
			x:    501,
			y:    250,
		},
		"RightRing": {
			Item: nil,
			x:    613,
			y:    250,
		},
		"Gloves": {
			Item: nil,
			x:    429,
			y:    252,
		},
		"Shoes": {
			Item: nil,
			x:    661,
			y:    252,
		},
	}
}
