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

// Decode decodes floor-shadow record
func (f *FloorShadowRecord) Decode(dw uint32) {
	f.Prop1 = byte(dw & 0x000000FF)            //nolint:gomnd // Bitmask
	f.Sequence = byte((dw & 0x00003F00) >> 8)  //nolint:gomnd // Bitmask
	f.Unknown1 = byte((dw & 0x000FC000) >> 14) //nolint:gomnd // Bitmask
	f.Style = byte((dw & 0x03F00000) >> 20)    //nolint:gomnd // Bitmask
	f.Unknown2 = byte((dw & 0x7C000000) >> 26) //nolint:gomnd // Bitmask
	f.hidden = byte((dw & 0x80000000) >> 31)   //nolint:gomnd // Bitmask
}

// Encode encodes floor-shadow record
func (f *FloorShadowRecord) Encode() (dw uint32) {
	dw |= uint32(f.Prop1) & 0xFF            //nolint:gomnd // Bitmask
	dw |= (uint32(f.Sequence) & 0x3F) << 8  //nolint:gomnd // Bitmask
	dw |= (uint32(f.Unknown1) & 0xFC) << 14 //nolint:gomnd // Bitmask
	dw |= (uint32(f.Style) & 0x3F) << 20    //nolint:gomnd // Bitmask
	dw |= (uint32(f.Unknown2) & 0x7C) << 26 //nolint:gomnd // Bitmask
	dw |= (uint32(f.hidden) & 0x01) << 31   //nolint:gomnd // Bitmask

	return dw
}
