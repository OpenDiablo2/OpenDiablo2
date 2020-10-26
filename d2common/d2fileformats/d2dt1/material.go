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
