package d2dt1

// Tile is a representation of a map tile
type Tile struct {
	unknown2           []byte
	Blocks             []Block
	Sequence           int32
	RarityFrameIndex   int32
	Height             int32
	Width              int32
	Type               int32
	Direction          int32
	blockHeaderPointer int32
	blockHeaderSize    int32
	Style              int32
	RoofHeight         int16
	SubTileFlags       [25]SubTileFlags
	MaterialFlags      MaterialFlags
}

func (t *Tile) unknown1() []byte {
	return make([]byte, numUnknownTileBytes1)
}

func (t *Tile) unknown3() []byte {
	return make([]byte, numUnknownTileBytes3)
}

func (t *Tile) unknown4() []byte {
	return make([]byte, numUnknownTileBytes4)
}
