package d2enum

// EquippedTypeSlot represents the type of equipment slot
type EquippedSlotType int

const (
	Head      EquippedSlotType = iota + 1
	Torso     EquippedSlotType = 2
	Legs      EquippedSlotType = 3
	RightArm  EquippedSlotType = 4
	LeftArm   EquippedSlotType = 5
	LeftHand  EquippedSlotType = 6
	RightHand EquippedSlotType = 7
	Neck      EquippedSlotType = 8
	Belt      EquippedSlotType = 9
	Gloves    EquippedSlotType = 10
)
