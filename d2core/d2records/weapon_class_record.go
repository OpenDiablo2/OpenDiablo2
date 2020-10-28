package d2records

// WeaponClasses is a map of WeaponClassRecords
type WeaponClasses map[string]*WeaponClassRecord

// WeaponClassRecord describes a weapon class. It has a name and 3-character token.
// The token is used to change the character animation mode.
type WeaponClassRecord struct {
	Name  string
	Token string
}
