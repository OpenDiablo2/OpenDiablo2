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

func (t *Tile) GetSubTileFlags(x, y int) *SubTileFlags {
	if x < 0 || x > 4 || y < 0 || y > 4 {
		return &SubTileFlags{}
	}
	return &t.SubTileFlags[x + (y * 5)]
}
