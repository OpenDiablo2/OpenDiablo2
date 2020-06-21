package d2ds1

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type TileRecord struct {
	Floors        []FloorShadowRecord
	Walls         []WallRecord
	Shadows       []FloorShadowRecord
	Substitutions []SubstitutionRecord

	// This is set and used internally by the engine to determine what region this map is from
	RegionType d2enum.RegionIdType
}
