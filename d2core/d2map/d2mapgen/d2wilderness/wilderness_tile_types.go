// Package d2wilderness provides an enumeration of wilderness types
package d2wilderness

// nolint:golint // these don't require individual explanations.
const (
	TreeBorderSouth int = iota + 4
	TreeBorderWest
	TreeBorderNorth
	TreeBorderEast
	TreeBorderSouthWest
	TreeBorderNorthWest
	TreeBorderNorthEast
	TreeBorderSouthEast
	TreeBoxNorthEast
	TreeBoxSouthEast
	TreeBoxSouthWest
	TreeBoxNorthWest
	WallBorderWest
	WallBorderEast
	WallBorderWestFenceNorth
	WallBorderNorthWest
	WallBorderWestFenceSouth
	WallBorderNorthFenceEast
	WallBorderNorthFenceWest
	WallBoxSouthEast
	WallBorderNorthUndergroundPassageEntrance
	WallBorderWestUndergroundPassageEntrance
	WaterBorderEast
	WaterBorderWest
	WaterBridgeEast
	StoneFill1
	StoneFill2
	CorralFill
	RandomTreesAndWallBoxLarge
	TreeBoxNorthSouth
	ShrineWithFenceAndTrees
	RandomTreesAndWallBoxSmall
	TreeBoxWestEastWithNorthSouthPath
	TreeBoxNorthSouthWithEastWestPath
	SwampFill1
	SwampFill2
	TreeFill
	Ruin
	FallenCamp1
	FallenCamp2
	FallenCampBishbosh
	Camp
	Pond
	Cottages1
	Cottages2
	Cottages3
	Bivouac
	CaveEntrance
	DenOfEvilEntrance
)
