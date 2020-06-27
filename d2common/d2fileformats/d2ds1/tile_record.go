package d2ds1

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// TileRecord represents a tile record in a DS1 file.
type TileRecord struct {
	Floors        []FloorShadowRecord  // Collection of floor records
	Walls         []WallRecord         // Collection of wall records
	Shadows       []FloorShadowRecord  // Collection of shadow records
	Substitutions []SubstitutionRecord // Collection of substitutions

	// This is set and used internally by the engine to determine what region this map is from
	RegionType d2enum.RegionIdType
}
