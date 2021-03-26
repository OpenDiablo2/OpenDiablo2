package d2enum

// TileType represents a tile type
type TileType byte

// Tile types
const (
	TileFloor TileType = iota
	TileLeftWall
	TileRightWall
	TileRightPartOfNorthCornerWall
	TileLeftPartOfNorthCornerWall
	TileLeftEndWall
	TileRightEndWall
	TileSouthCornerWall
	TileLeftWallWithDoor
	TileRightWallWithDoor
	TileSpecialTile1
	TileSpecialTile2
	TilePillarsColumnsAndStandaloneObjects
	TileShadow
	TileTree
	TileRoof
	TileLowerWallsEquivalentToLeftWall
	TileLowerWallsEquivalentToRightWall
	TileLowerWallsEquivalentToRightLeftNorthCornerWall
	TileLowerWallsEquivalentToSouthCornerwall
)

// LowerWall checks for lower wall tiles
func (tile TileType) LowerWall() bool {
	switch tile {
	case TileLowerWallsEquivalentToLeftWall,
		TileLowerWallsEquivalentToRightWall,
		TileLowerWallsEquivalentToRightLeftNorthCornerWall,
		TileLowerWallsEquivalentToSouthCornerwall:
		return true
	}

	return false
}

// UpperWall checks for upper wall tiles
func (tile TileType) UpperWall() bool {
	switch tile {
	case TileLeftWall,
		TileRightWall,
		TileRightPartOfNorthCornerWall,
		TileLeftPartOfNorthCornerWall,
		TileLeftEndWall,
		TileRightEndWall,
		TileSouthCornerWall,
		TileLeftWallWithDoor,
		TileRightWallWithDoor,
		TilePillarsColumnsAndStandaloneObjects,
		TileTree:
		return true
	}

	return false
}

// Special checks for special tiles
func (tile TileType) Special() bool {
	switch tile {
	case TileSpecialTile1, TileSpecialTile2:
		return true
	}

	return false
}

func (tile TileType) String() string {
	strings := map[TileType]string{
		TileFloor:                              "floor",
		TileLeftWall:                           "Left Wall",
		TileRightWall:                          "Upper Wall",
		TileRightPartOfNorthCornerWall:         "Upper part of an Upper-Left corner",
		TileLeftPartOfNorthCornerWall:          "Left part of an Upper-Left corner",
		TileLeftEndWall:                        "Upper-Right corner",
		TileRightEndWall:                       "Lower-Left corner",
		TileSouthCornerWall:                    "Lower-Right corner",
		TileLeftWallWithDoor:                   "Left Wall with Door object, but not always",
		TileRightWallWithDoor:                  "Upper Wall with Door object, but not always",
		TileSpecialTile1:                       "special",
		TileSpecialTile2:                       "special",
		TilePillarsColumnsAndStandaloneObjects: "billars, collumns or standalone object",
		TileShadow:                             "shadow",
		TileTree:                               "wall/object",
		TileRoof:                               "roof",
		TileLowerWallsEquivalentToLeftWall:     "lower wall (left wall)",
		TileLowerWallsEquivalentToRightWall:    "lower wall (right wall)",
		TileLowerWallsEquivalentToRightLeftNorthCornerWall: "lower wall (north corner wall)",
		TileLowerWallsEquivalentToSouthCornerwall:          "lower wall (south corner wall)",
	}

	str, found := strings[tile]

	if !found {
		str = "unknown"
	}

	return str
}
