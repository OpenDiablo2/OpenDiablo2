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

// PrepareTile selects which graphic to use and updates the tiles subtileflags
func (t *MapTile) PrepareTile(x, y int, me *MapEngine) {
	for wIdx := range t.Components.Walls {
		wall := &t.Components.Walls[wIdx]
		options := me.GetTiles(int(wall.Style), int(wall.Sequence), wall.Type)

		if options == nil {
			break
		}

		wall.RandomIndex = getRandomTile(options, x, y, me.seed)

		for i := range t.SubTiles {
			t.SubTiles[i].Combine(options[wall.RandomIndex].SubTileFlags[i])
		}
	}

	for fIdx := range t.Components.Floors {
		floor := &t.Components.Floors[fIdx]
		options := me.GetTiles(int(floor.Style), int(floor.Sequence), 0)

		if options == nil {
			break
		}

		if options[0].MaterialFlags.Lava {
			floor.Animated = true
			floor.RandomIndex = 0
		} else {
			floor.RandomIndex = getRandomTile(options, x, y, me.seed)
		}

		for i := range t.SubTiles {
			t.SubTiles[i].Combine(options[floor.RandomIndex].SubTileFlags[i])
		}
	}

	for sIdx := range t.Components.Shadows {
		shadow := &t.Components.Shadows[sIdx]
		options := me.GetTiles(int(shadow.Style), int(shadow.Sequence), 13)

		if options == nil {
			break
		}

		shadow.RandomIndex = getRandomTile(options, x, y, me.seed)

		for i := range t.SubTiles {
			t.SubTiles[i].Combine(options[shadow.RandomIndex].SubTileFlags[i])
		}
	}
}

// Walker's Alias Method for weighted random selection with xorshifting for random numbers
// Selects a random tile from the slice, rest of args just used for seeding
func getRandomTile(tiles []d2dt1.Tile, x, y int, seed int64) byte {
	var tileSeed uint64
	tileSeed = uint64(seed) + uint64(x)
	tileSeed *= uint64(y)

	const (
		xorshiftA = 13
		xorshiftB = 17
		xorshiftC = 5
	)

	tileSeed ^= tileSeed << xorshiftA
	tileSeed ^= tileSeed >> xorshiftB
	tileSeed ^= tileSeed << xorshiftC

	weightSum := 0

	for i := range tiles {
		weightSum += int(tiles[i].RarityFrameIndex)
	}

	if weightSum == 0 {
		return 0
	}

	random := tileSeed % uint64(weightSum)

	sum := 0

	for i := range tiles {
		sum += int(tiles[i].RarityFrameIndex)
		if sum >= int(random) {
			return byte(i)
		}
	}

	// This return shouldn't be hit
	return 0
}
