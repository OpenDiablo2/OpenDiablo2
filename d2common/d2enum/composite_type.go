package d2enum

// CompositeType represents a composite type
type CompositeType int

const (

	// CompositeTypeHead is a composite type for heads
	CompositeTypeHead CompositeType = iota // HD

	// CompositeTypeTorso is a composite type for torsos
	CompositeTypeTorso // TR

	// CompositeTypeLegs is a composite type for legs
	CompositeTypeLegs // LG

	// CompositeTypeRightArm is a composite type for right arms
	CompositeTypeRightArm // RA

	// CompositeTypeLeftArm is a composite type for left arms
	CompositeTypeLeftArm // LA

	// CompositeTypeRightHand is a composite type for right hands
	CompositeTypeRightHand // RH

	// CompositeTypeLeftHand is a composite type for left hands
	CompositeTypeLeftHand // LH

	// CompositeTypeShield is a composite type for shields
	CompositeTypeShield // SH

	// CompositeTypeSpecial1 is a composite type for special type 1s
	CompositeTypeSpecial1 // S1

	// CompositeTypeSpecial2 is a composite type for special type 2s
	CompositeTypeSpecial2 // S2

	// CompositeTypeSpecial3 is a composite type for special type 3s
	CompositeTypeSpecial3 // S3

	// CompositeTypeSpecial4 is a composite type for special type 4s
	CompositeTypeSpecial4 // S4

	// CompositeTypeSpecial5 is a composite type for special type 5s
	CompositeTypeSpecial5 // S5

	// CompositeTypeSpecial6 is a composite type for special type 6s
	CompositeTypeSpecial6 // S6

	// CompositeTypeSpecial7 is a composite type for special type 7s
	CompositeTypeSpecial7 // S7

	// CompositeTypeSpecial8 is a composite type for special type 8s
	CompositeTypeSpecial8 // S8

	// CompositeTypeMax is used to determine the max number of composite types
	CompositeTypeMax
)

//go:generate stringer -linecomment -type CompositeType
