package d2inventory

type CharacterEquipment struct {
	Head      *InventoryItemArmor  // Head
	Torso     *InventoryItemArmor  // TR
	Legs      *InventoryItemArmor  // Legs
	RightArm  *InventoryItemArmor  // RA
	LeftArm   *InventoryItemArmor  // LA
	LeftHand  *InventoryItemWeapon // LH
	RightHand *InventoryItemWeapon // RH
	Shield    *InventoryItemArmor  // SH
	// S1-S8?
}
