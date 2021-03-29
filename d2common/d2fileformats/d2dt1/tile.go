package d2dt1

// Tile is a representation of a map tile
type Tile struct {
	unknown2           []byte
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

func (t *Tile) unknown1() []byte {
	return make([]byte, numUnknownTileBytes1)
}

func (t *Tile) unknown3() []byte {
	return make([]byte, numUnknownTileBytes3)
}

func (t *Tile) unknown4() []byte {
	return make([]byte, numUnknownTileBytes4)
}
