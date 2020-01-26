package d2enum

type WeaponClass int

const (
	WeaponClassNone                 WeaponClass = 0  //
	WeaponClassHandToHand           WeaponClass = 1  // hth
	WeaponClassBow                  WeaponClass = 2  // bow
	WeaponClassOneHandSwing         WeaponClass = 3  // 1hs
	WeaponClassOneHandThrust        WeaponClass = 4  // 1ht
	WeaponClassStaff                WeaponClass = 5  // stf
	WeaponClassTwoHandSwing         WeaponClass = 6  // 2hs
	WeaponClassTwoHandThrust        WeaponClass = 7  // 2ht
	WeaponClassCrossbow             WeaponClass = 8  // xbw
	WeaponClassLeftJabRightSwing    WeaponClass = 9  // 1js
	WeaponClassLeftJabRightThrust   WeaponClass = 10 // 1jt
	WeaponClassLeftSwingRightSwing  WeaponClass = 11 // 1ss
	WeaponClassLeftSwingRightThrust WeaponClass = 12 // 1st
	WeaponClassOneHandToHand        WeaponClass = 13 // ht1
	WeaponClassTwoHandToHand        WeaponClass = 14 // ht2
)

//go:generate stringer -linecomment -type WeaponClass
//go:generate string2enum -samepkg -linecomment -type WeaponClass
