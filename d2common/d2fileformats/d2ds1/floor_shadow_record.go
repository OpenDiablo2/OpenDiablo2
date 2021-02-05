package d2ds1

// FloorShadowRecord represents a floor or shadow record in a DS1 file.
type FloorShadowRecord struct {
	Prop1       byte
	Sequence    byte
	Unknown1    byte
	Style       byte
	Unknown2    byte
	hidden      byte
	RandomIndex byte
	Animated    bool
	YAdjust     int
}

// Hidden returns if floor/shadow is hidden
func (f *FloorShadowRecord) Hidden() bool {
	return f.hidden > 0
}
