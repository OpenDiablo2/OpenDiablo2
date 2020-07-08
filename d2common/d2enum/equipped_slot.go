package d2enum

// EquippedSlot represents the type of equipment slot
type EquippedSlot int

// Equipped slot ID's
const (
	EquippedSlotHead EquippedSlot = iota + 1
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
