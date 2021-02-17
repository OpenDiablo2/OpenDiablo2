package d2ds1

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

const (
	prop1Bitmask = 0x000000FF
	prop1Offset  = 0
	prop1Length  = 8

	sequenceBitmask = 0x00003F00
	sequenceOffset  = 8
	sequenceLength  = 6

	unknown1Bitmask = 0x000FC000
	unknown1Offset  = 14
	unknown1Length  = 6

	styleBitmask = 0x03F00000
	styleOffset  = 20
	styleLength  = 6

	unknown2Bitmask = 0x7C000000
	unknown2Offset  = 26
	unknown2Length  = 5

	hiddenBitmask = 0x80000000
	hiddenOffset  = 31
	hiddenLength  = 1
)

type floorShadow struct {
	Prop1       byte
	Sequence    byte
	Unknown1    byte
	Style       byte
	Unknown2    byte
	HiddenBytes byte
	RandomIndex byte
	Animated    bool
	YAdjust     int
}

// Floor represents a floor record in a DS1 file. (it is just an alias of floorShadow).
type Floor = floorShadow

// Shadow represents a shadow record in a DS1 file. (it is just an alias of floorShadow).
type Shadow = floorShadow

// Hidden returns if floor/shadow is hidden
func (f *Floor) Hidden() bool {
	return f.HiddenBytes > 0
}

// Decode decodes floor-shadow record
func (f *Floor) Decode(dw uint32) {
	f.Prop1 = byte((dw & prop1Bitmask) >> prop1Offset)
	f.Sequence = byte((dw & sequenceBitmask) >> sequenceOffset)
	f.Unknown1 = byte((dw & unknown1Bitmask) >> unknown1Offset)
	f.Style = byte((dw & styleBitmask) >> styleOffset)
	f.Unknown2 = byte((dw & unknown2Bitmask) >> unknown2Offset)
	f.HiddenBytes = byte((dw & hiddenBitmask) >> hiddenOffset)
}

// Encode adds Floor's bits to stream writter given
func (f *Floor) Encode(sw *d2datautils.StreamWriter) {
	sw.PushBits32(uint32(f.Prop1), prop1Length)
	sw.PushBits32(uint32(f.Sequence), sequenceLength)
	sw.PushBits32(uint32(f.Unknown1), unknown1Length)
	sw.PushBits32(uint32(f.Style), styleLength)
	sw.PushBits32(uint32(f.Unknown2), unknown2Length)
	sw.PushBits32(uint32(f.HiddenBytes), hiddenLength)
}
