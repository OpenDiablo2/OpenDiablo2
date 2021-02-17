package d2ds1

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Wall represents a wall record.
type Wall struct {
	Type        d2enum.TileType
	Zero        byte
	Prop1       byte
	Sequence    byte
	Unknown1    byte
	Style       byte
	Unknown2    byte
	HiddenBytes byte
	RandomIndex byte
	YAdjust     int
}

// Hidden returns if wall is hidden
func (w *Wall) Hidden() bool {
	return w.HiddenBytes > 0
}

// Decode decodes wall record
func (w *Wall) Decode(dw uint32) {
	w.Prop1 = byte((dw & prop1Bitmask) >> prop1Offset)
	w.Sequence = byte((dw & sequenceBitmask) >> sequenceOffset)
	w.Unknown1 = byte((dw & unknown1Bitmask) >> unknown1Offset)
	w.Style = byte((dw & styleBitmask) >> styleOffset)
	w.Unknown2 = byte((dw & unknown2Bitmask) >> unknown2Offset)
	w.HiddenBytes = byte((dw & hiddenBitmask) >> hiddenOffset)
}

// Encode adds wall's record's bytes into stream writer given
func (w *Wall) Encode(sw *d2datautils.StreamWriter) {
	sw.PushBits32(uint32(w.Prop1), prop1Length)
	sw.PushBits32(uint32(w.Sequence), sequenceLength)
	sw.PushBits32(uint32(w.Unknown1), unknown1Length)
	sw.PushBits32(uint32(w.Style), styleLength)
	sw.PushBits32(uint32(w.Unknown2), unknown2Length)
	sw.PushBits32(uint32(w.HiddenBytes), hiddenLength)
}
