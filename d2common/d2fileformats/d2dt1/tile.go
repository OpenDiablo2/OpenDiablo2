package d2dt1

type Tile struct {
	Direction          int32
	RoofHeight         int16
	MaterialFlags      MaterialFlags
	Height             int32
	Width              int32
	Type               int32
	Style              int32
	Sequence           int32
	RarityFrameIndex   int32
	SubTileFlags       [25]SubTileFlags
	blockHeaderPointer int32
	blockHeaderSize    int32
	Blocks             []Block
}

var subtileLookup = [5][5]int{
	{20, 21, 22, 23, 24},
	{15, 16, 17, 18, 19},
	{10, 11, 12, 13, 14},
	{5, 6, 7, 8, 9},
	{0, 1, 2, 3, 4},
}

func (t *Tile) GetSubTileFlags(x, y int) *SubTileFlags {

	return &t.SubTileFlags[subtileLookup[y][x]]
}
