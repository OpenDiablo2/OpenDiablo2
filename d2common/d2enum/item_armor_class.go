package d2enum

// ArmorClass is a 3-character token for the armor. It's used for speed calculations.
type ArmorClass string

//  Armor classes
const (
	ArmorClassLite   = "lit"
	ArmorClassMedium = "med"
	ArmorClassHeavy  = "hvy"
)
