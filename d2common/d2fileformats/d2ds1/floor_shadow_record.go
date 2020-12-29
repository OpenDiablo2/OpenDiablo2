package d2ds1

// FloorShadowRecord represents a floor or shadow record in a DS1 file.
type FloorShadowRecord struct {
	Prop1       byte
	Sequence    byte
	Unknown1    byte
	Style       byte
	Unknown2    byte
	Hidden      bool
	RandomIndex byte
	Animated    bool
	YAdjust     int
}
