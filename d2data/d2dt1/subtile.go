package d2dt1

type SubTileFlags struct {
	BlockWalk bool
	BlockLOS bool
	BlockJump bool
	BlockPlayerWalk bool
	Unknown1 bool
	BlockLight bool
	Unknown2 bool
	Unknown3 bool
}

func NewSubTileFlags(data byte) SubTileFlags {
	return SubTileFlags{
		BlockWalk:       data & 1 == 1,
		BlockLOS:        data & 2 == 2,
		BlockJump:       data & 4 == 4,
		BlockPlayerWalk: data & 8 == 8,
		Unknown1:        data & 16 == 16,
		BlockLight:      data & 32 == 32,
		Unknown2:        data & 64 == 64,
		Unknown3:        data & 128 == 128,
	}
}
