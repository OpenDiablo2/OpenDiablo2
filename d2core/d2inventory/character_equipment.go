package d2inventory

// CharacterEquipment stores equipments of a character
type CharacterEquipment struct {
	Head      *InventoryItemArmor  `json:"head"`      // Head
	Torso     *InventoryItemArmor  `json:"torso"`     // TR
	Legs      *InventoryItemArmor  `json:"legs"`      // Legs
	RightArm  *InventoryItemArmor  `json:"rightArm"`  // RA
	LeftArm   *InventoryItemArmor  `json:"leftArm"`   // LA
	LeftHand  *InventoryItemWeapon `json:"leftHand"`  // LH
	RightHand *InventoryItemWeapon `json:"rightHand"` // RH
	Shield    *InventoryItemArmor  `json:"shield"`    // SH
	// S1-S8?
}
