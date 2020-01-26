package d2enum

type TileType byte

const (
	Floor                                          TileType = 0
	LeftWall                                       TileType = 1
	RightWall                                      TileType = 2
	RightPartOfNorthCornerWall                     TileType = 3
	LeftPartOfNorthCornerWall                      TileType = 4
	LeftEndWall                                    TileType = 5
	RightEndWall                                   TileType = 6
	SouthCornerWall                                TileType = 7
	LeftWallWithDoor                               TileType = 8
	RightWallWithDoor                              TileType = 9
	SpecialTile1                                   TileType = 10
	SpecialTile2                                   TileType = 11
	PillarsColumnsAndStandaloneObjects             TileType = 12
	Shadow                                         TileType = 13
	Tree                                           TileType = 14
	Roof                                           TileType = 15
	LowerWallsEquivalentToLeftWall                 TileType = 16
	LowerWallsEquivalentToRightWall                TileType = 17
	LowerWallsEquivalentToRightLeftNorthCornerWall TileType = 18
	LowerWallsEquivalentToSouthCornerwall          TileType = 19
)

func (tile TileType) LowerWall() bool {
	switch tile {
	case 16, 17, 18, 19:
		return true
	default:
		return false
	}
}

func (tile TileType) UpperWall() bool {
	switch tile {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 14:
		return true
	default:
		return false
	}
}
