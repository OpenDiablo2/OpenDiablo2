package d2data

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/datadict"
)

type Object struct {
	Type       int32
	Id         int32
	X          int32
	Y          int32
	Flags      int32
	Paths      []d2common.Path
	Lookup     *datadict.ObjectLookupRecord
	ObjectInfo *datadict.ObjectRecord
}
