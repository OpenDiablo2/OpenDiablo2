package d2enum

// PetIconType determines the pet icon type
type PetIconType int

// Pet icon types
// The information has been gathered from [https:// d2mods.info/forum/kb/viewarticle?a=355]
const (
	NoIcon PetIconType = iota
	ShowIconOnly
	ShowIconAndQuantity // Quantity, such as the number of skeletons
	ShowIconOnly2
)
