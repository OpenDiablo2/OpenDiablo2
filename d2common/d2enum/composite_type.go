package d2enum

const (
	unknown = "Unknown"
)

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

// Name returns a full name of layer
func (i CompositeType) Name() string {
	strings := map[CompositeType]string{
		CompositeTypeHead:      "Head",
		CompositeTypeTorso:     "Torso",
		CompositeTypeLegs:      "Legs",
		CompositeTypeRightArm:  "Right Arm",
		CompositeTypeLeftArm:   "Left Arm",
		CompositeTypeRightHand: "Right Hand",
		CompositeTypeLeftHand:  "Left Hand",
		CompositeTypeShield:    "Shield",
		CompositeTypeSpecial1:  "Special 1",
		CompositeTypeSpecial2:  "Special 2",
		CompositeTypeSpecial3:  "Special 3",
		CompositeTypeSpecial4:  "Special 4",
		CompositeTypeSpecial5:  "Special 5",
		CompositeTypeSpecial6:  "Special 6",
		CompositeTypeSpecial7:  "Special 7",
		CompositeTypeSpecial8:  "Special 8",
	}

	layerName, found := strings[i]
	if !found {
		return unknown
	}

	return layerName
}
