package d2ds1

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// WallRecord represents a wall record.
type WallRecord struct {
	Type        d2enum.TileType
	Zero        byte
	Prop1       byte
	Sequence    byte
	Unknown1    byte
	Style       byte
	Unknown2    byte
	hidden      byte
	RandomIndex byte
	YAdjust     int
}

// Hidden returns if wall is hidden
func (w *WallRecord) Hidden() bool {
	return w.hidden > 0
}

// Encode encodes wall record
func (w *WallRecord) Encode() (dw uint32) {
	dw |= uint32(w.Prop1) & 0xFF            //nolint:gomnd // Bitmask
	dw |= (uint32(w.Sequence) & 0x3F) << 8  //nolint:gomnd // Bitmask
	dw |= (uint32(w.Unknown1) & 0xFC) << 14 //nolint:gomnd // Bitmask
	dw |= (uint32(w.Style) & 0x3F) << 20    //nolint:gomnd // Bitmask
	dw |= (uint32(w.Unknown2) & 0x7C) << 26 //nolint:gomnd // Bitmask
	dw |= (uint32(w.hidden) & 0x01) << 31   //nolint:gomnd // Bitmask

	return dw
}

// Decode decodes wall record
func (w *WallRecord) Decode(dw uint32) {
	w.Prop1 = byte(dw & 0x000000FF)            //nolint:gomnd // Bitmask
	w.Sequence = byte((dw & 0x00003F00) >> 8)  //nolint:gomnd // Bitmask
	w.Unknown1 = byte((dw & 0x000FC000) >> 14) //nolint:gomnd // Bitmask
	w.Style = byte((dw & 0x03F00000) >> 20)    //nolint:gomnd // Bitmask
	w.Unknown2 = byte((dw & 0x7C000000) >> 26) //nolint:gomnd // Bitmask
	w.hidden = byte((dw & 0x80000000) >> 31)   //nolint:gomnd // Bitmask
}
