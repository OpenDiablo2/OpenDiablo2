package d2mapengine

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
)

// MapTile is a tile placed on the map
type MapTile struct {
	Components d2ds1.TileRecord
	RegionType d2enum.RegionIdType
	SubTiles   [25]d2dt1.SubTileFlags
}

// GetSubTileFlags returns the tile flags for the given subtile
func (t *MapTile) GetSubTileFlags(x, y int) *d2dt1.SubTileFlags {
	var subtileLookup = [5][5]int{
		{20, 21, 22, 23, 24},
		{15, 16, 17, 18, 19},
		{10, 11, 12, 13, 14},
		{5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4},
	}

	return &t.SubTiles[subtileLookup[y][x]]
}
