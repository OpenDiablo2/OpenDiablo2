package d2enum

// CompositeType represents a composite type
type CompositeType int

const (

	// CompositeTypeHead is a composite type for heads
	CompositeTypeHead CompositeType = iota

	// CompositeTypeTorso is a composite type for torsos
	CompositeTypeTorso

	// CompositeTypeLegs is a composite type for legs
	CompositeTypeLegs

	// CompositeTypeRightArm is a composite type for right arms
	CompositeTypeRightArm

	// CompositeTypeLeftArm is a composite type for left arms
	CompositeTypeLeftArm

	// CompositeTypeRightHand is a composite type for right hands
	CompositeTypeRightHand

	// CompositeTypeLeftHand is a composite type for left hands
	CompositeTypeLeftHand

	// CompositeTypeShield is a composite type for shields
	CompositeTypeShield

	// CompositeTypeSpecial1 is a composite type for special type 1s
	CompositeTypeSpecial1

	// CompositeTypeSpecial2 is a composite type for special type 2s
	CompositeTypeSpecial2

	// CompositeTypeSpecial3 is a composite type for special type 3s
	CompositeTypeSpecial3

	// CompositeTypeSpecial4 is a composite type for special type 4s
	CompositeTypeSpecial4

	// CompositeTypeSpecial5 is a composite type for special type 5s
	CompositeTypeSpecial5

	// CompositeTypeSpecial6 is a composite type for special type 6s
	CompositeTypeSpecial6

	// CompositeTypeSpecial7 is a composite type for special type 7s
	CompositeTypeSpecial7

	// CompositeTypeSpecial8 is a composite type for special type 8s
	CompositeTypeSpecial8

	// CompositeTypeMax is used to determine the max number of composite types
	CompositeTypeMax
)
