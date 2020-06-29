package d2dt1

// SubTileFlags represent the sub-tile flags for a DT1
type SubTileFlags struct {
	BlockWalk       bool
	BlockLOS        bool
	BlockJump       bool
	BlockPlayerWalk bool
	Unknown1        bool
	BlockLight      bool
	Unknown2        bool
	Unknown3        bool
}

// DebugString returns the debug string
func (s *SubTileFlags) DebugString() string {
	result := ""

	if s.BlockWalk {
		result += "BlockWalk "
	}

	if s.BlockLOS {
		result += "BlockLOS "
	}

	if s.BlockJump {
		result += "BlockJump "
	}

	if s.BlockPlayerWalk {
		result += "BlockPlayerWalk "
	}

	if s.Unknown1 {
		result += "Unknown1 "
	}

	if s.BlockLight {
		result += "BlockLight "
	}

	if s.Unknown2 {
		result += "Unknown2 "
	}

	if s.Unknown3 {
		result += "Unknown3 "
	}

	return result
}

// NewSubTileFlags returns a list of new subtile flags
//nolint:gomnd binary flags
func NewSubTileFlags(data byte) SubTileFlags {
	return SubTileFlags{
		BlockWalk:       data&1 == 1,
		BlockLOS:        data&2 == 2,
		BlockJump:       data&4 == 4,
		BlockPlayerWalk: data&8 == 8,
		Unknown1:        data&16 == 16,
		BlockLight:      data&32 == 32,
		Unknown2:        data&64 == 64,
		Unknown3:        data&128 == 128,
	}
}
