package MapEngine

type Orientation int32

const (
	Floors                                         Orientation = 0
	LeftWall                                       Orientation = 1
	RightWall                                      Orientation = 2
	RightPartOfNorthCornerWall                     Orientation = 3
	LeftPartOfNorthCornerWall                      Orientation = 4
	LeftEndWall                                    Orientation = 5
	RightEndWall                                   Orientation = 6
	SouthCornerWall                                Orientation = 7
	LeftWallWithDoor                               Orientation = 8
	RightWallWithDoor                              Orientation = 9
	SpecialTile1                                   Orientation = 10
	SpecialTile2                                   Orientation = 11
	PillarsColumnsAndStandaloneObjects             Orientation = 12
	Shadows                                        Orientation = 13
	Trees                                          Orientation = 14
	Roofs                                          Orientation = 15
	LowerWallsEquivalentToLeftWall                 Orientation = 16
	LowerWallsEquivalentToRightWall                Orientation = 17
	LowerWallsEquivalentToRightLeftNorthCornerWall Orientation = 18
	LowerWallsEquivalentToSouthCornerwall          Orientation = 19
)
