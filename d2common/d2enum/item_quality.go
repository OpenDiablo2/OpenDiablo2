package d2enum

//ItemQuality is used for enumerating item quality values
type ItemQuality int

const (
	LowQuality ItemQuality = iota + 1
	Normal
	Superior
	Magic
	Set
	Rare
	Unique
	Crafted
	Tempered
)
