package d2data

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

type Object struct {
	Type       int
	Id         int
	X          int
	Y          int
	Flags      int
	Paths      []d2common.Path
	Lookup     *d2datadict.ObjectLookupRecord
	ObjectInfo *d2datadict.ObjectRecord
}
