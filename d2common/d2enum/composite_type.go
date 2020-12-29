package d2enum

//go:generate stringer -linecomment -type CompositeType -output composite_type_string.go

// CompositeType represents a composite type
type CompositeType int

// Composite types
const (
	CompositeTypeHead      CompositeType = iota // HD
	CompositeTypeTorso                          // TR
	CompositeTypeLegs                           // LG
	CompositeTypeRightArm                       // RA
	CompositeTypeLeftArm                        // LA
	CompositeTypeRightHand                      // RH
	CompositeTypeLeftHand                       // LH
	CompositeTypeShield                         // SH
	CompositeTypeSpecial1                       // S1
	CompositeTypeSpecial2                       // S2
	CompositeTypeSpecial3                       // S3
	CompositeTypeSpecial4                       // S4
	CompositeTypeSpecial5                       // S5
	CompositeTypeSpecial6                       // S6
	CompositeTypeSpecial7                       // S7
	CompositeTypeSpecial8                       // S8
	CompositeTypeMax
)
