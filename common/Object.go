package common

type Object struct {
	Type       int32
	Id         int32
	X          int32
	Y          int32
	Flags      int32
	Paths      []Path
	Lookup     *ObjectLookupRecord
	ObjectInfo *ObjectRecord
}
