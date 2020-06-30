package d2enum

// ObjectType is the type of an object
type ObjectType int

const (
	ObjectTypeCharacter ObjectType = iota + 1
	ObjectTypeItem
)
