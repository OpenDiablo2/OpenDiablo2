package d2enum

//go:generate stringer -linecomment -type WeaponClass -output weapon_class_string.go
//go:generate string2enum -samepkg -linecomment -type WeaponClass -output weapon_class_string2enum.go

// WeaponClass represents a weapon class
type WeaponClass int

// Weapon classes
const (
	WeaponClassNone                 WeaponClass = iota //
	WeaponClassHandToHand                              // hth
	WeaponClassBow                                     // bow
	WeaponClassOneHandSwing                            // 1hs
	WeaponClassOneHandThrust                           // 1ht
	WeaponClassStaff                                   // stf
	WeaponClassTwoHandSwing                            // 2hs
	WeaponClassTwoHandThrust                           // 2ht
	WeaponClassCrossbow                                // xbw
	WeaponClassLeftJabRightSwing                       // 1js
	WeaponClassLeftJabRightThrust                      // 1jt
	WeaponClassLeftSwingRightSwing                     // 1ss
	WeaponClassLeftSwingRightThrust                    // 1st
	WeaponClassOneHandToHand                           // ht1
	WeaponClassTwoHandToHand                           // ht2
)

func (w WeaponClass) Name() string {
	strings := map[WeaponClass]string{
		WeaponClassNone:                 "None",
		WeaponClassHandToHand:           "Hand To Hand",
		WeaponClassBow:                  "Bow",
		WeaponClassOneHandSwing:         "One Hand Swing",
		WeaponClassOneHandThrust:        "One Hand Thrust",
		WeaponClassStaff:                "Staff",
		WeaponClassTwoHandSwing:         "Two Hand Swing",
		WeaponClassTwoHandThrust:        "Two Hand Thrust",
		WeaponClassCrossbow:             "Crossbow",
		WeaponClassLeftJabRightSwing:    "Left Jab Right Swing",
		WeaponClassLeftJabRightThrust:   "Left Jab Right Thrust",
		WeaponClassLeftSwingRightSwing:  "Left Swing Right Swing",
		WeaponClassLeftSwingRightThrust: "Left Swing Right Thrust",
		WeaponClassOneHandToHand:        "One Hand To Hand",
		WeaponClassTwoHandToHand:        "Two Hand To Hand",
	}

	weaponClass, found := strings[w]
	if !found {
		return "Unknown"
	}

	return weaponClass
}
