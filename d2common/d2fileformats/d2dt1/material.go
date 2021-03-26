package d2dt1

// MaterialFlags represents the material flags. Lots of unknowns for now...
type MaterialFlags struct {
	Other        bool
	Water        bool
	WoodObject   bool
	InsideStone  bool
	OutsideStone bool
	Dirt         bool
	Sand         bool
	Wood         bool
	Lava         bool
	Snow         bool
}

// NewMaterialFlags  represents the material flags
// nolint:gomnd // Binary values
func NewMaterialFlags(data uint16) MaterialFlags {
	return MaterialFlags{
		Other:        data&0x0001 == 0x0001,
		Water:        data&0x0002 == 0x0002,
		WoodObject:   data&0x0004 == 0x0004,
		InsideStone:  data&0x0008 == 0x0008,
		OutsideStone: data&0x0010 == 0x0010,
		Dirt:         data&0x0020 == 0x0020,
		Sand:         data&0x0040 == 0x0040,
		Wood:         data&0x0080 == 0x0080,
		Lava:         data&0x0100 == 0x0100,
		Snow:         data&0x0400 == 0x0400,
	}
}

// Encode encodes MaterialFlags back to uint16
func (m *MaterialFlags) Encode() uint16 {
	var b uint16 = 0x000

	if m.Other {
		b |= 0x0001
	}

	if m.Water {
		b |= 0x0002
	}

	if m.WoodObject {
		b |= 0x0004
	}

	if m.InsideStone {
		b |= 0x0008
	}

	if m.OutsideStone {
		b |= 0x0010
	}

	if m.Dirt {
		b |= 0x0020
	}

	if m.Sand {
		b |= 0x0040
	}

	if m.Wood {
		b |= 0x0080
	}

	if m.Lava {
		b |= 0x0100
	}

	if m.Snow {
		b |= 0x0400
	}

	return b
}
