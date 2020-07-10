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
