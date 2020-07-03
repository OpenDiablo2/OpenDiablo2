package d2enum

// ObjectType is the type of an object
type ObjectType int

const (
	ObjectTypePlayer ObjectType = iota
	ObjectTypeCharacter
	ObjectTypeItem
)
