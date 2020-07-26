package d2enum

// EquippedSlot represents the type of equipment slot
type EquippedSlot int

// Equipped slot ID's
const (
	EquippedSlotNone EquippedSlot = iota
	EquippedSlotHead
	EquippedSlotTorso
	EquippedSlotLegs
	EquippedSlotRightArm
	EquippedSlotLeftArm
	EquippedSlotLeftHand
	EquippedSlotRightHand
	EquippedSlotNeck
	EquippedSlotBelt
	EquippedSlotGloves
)
