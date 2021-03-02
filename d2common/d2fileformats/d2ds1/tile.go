package d2ds1

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
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

type tileCommonFields struct {
	Prop1       byte
	Sequence    byte
	Unknown1    byte
	Style       byte
	Unknown2    byte
	HiddenBytes byte
	RandomIndex byte
	YAdjust     int
}

type tileFloorShadowFields struct {
	Animated bool
}

type tileSubstitutionFields struct {
	Substitution uint32 // unknown
}

type tileWallFields struct {
	Type d2enum.TileType
	Zero byte
}

// Tile represents a tile record in a DS1 file.
type Tile struct {
	tileCommonFields
	tileFloorShadowFields
	tileSubstitutionFields
	tileWallFields
}

// Hidden returns if wall is hidden
func (t *Tile) Hidden() bool {
	return t.HiddenBytes > 0
}

// DecodeWall decodes as a wall record
func (t *Tile) DecodeWall(dw uint32) {
	t.Prop1 = byte((dw & prop1Bitmask) >> prop1Offset)
	t.Sequence = byte((dw & sequenceBitmask) >> sequenceOffset)
	t.Unknown1 = byte((dw & unknown1Bitmask) >> unknown1Offset)
	t.Style = byte((dw & styleBitmask) >> styleOffset)
	t.Unknown2 = byte((dw & unknown2Bitmask) >> unknown2Offset)
	t.HiddenBytes = byte((dw & hiddenBitmask) >> hiddenOffset)
}

// EncodeWall adds wall's record's bytes into stream writer given
func (t *Tile) EncodeWall(sw *d2datautils.StreamWriter) {
	sw.PushBits32(uint32(t.Prop1), prop1Length)
	sw.PushBits32(uint32(t.Sequence), sequenceLength)
	sw.PushBits32(uint32(t.Unknown1), unknown1Length)
	sw.PushBits32(uint32(t.Style), styleLength)
	sw.PushBits32(uint32(t.Unknown2), unknown2Length)
	sw.PushBits32(uint32(t.HiddenBytes), hiddenLength)
}

func (t *Tile) decodeFloorShadow(dw uint32) {
	t.Prop1 = byte((dw & prop1Bitmask) >> prop1Offset)
	t.Sequence = byte((dw & sequenceBitmask) >> sequenceOffset)
	t.Unknown1 = byte((dw & unknown1Bitmask) >> unknown1Offset)
	t.Style = byte((dw & styleBitmask) >> styleOffset)
	t.Unknown2 = byte((dw & unknown2Bitmask) >> unknown2Offset)
	t.HiddenBytes = byte((dw & hiddenBitmask) >> hiddenOffset)
}

func (t *Tile) encodeFloorShadow(sw *d2datautils.StreamWriter) {
	sw.PushBits32(uint32(t.Prop1), prop1Length)
	sw.PushBits32(uint32(t.Sequence), sequenceLength)
	sw.PushBits32(uint32(t.Unknown1), unknown1Length)
	sw.PushBits32(uint32(t.Style), styleLength)
	sw.PushBits32(uint32(t.Unknown2), unknown2Length)
	sw.PushBits32(uint32(t.HiddenBytes), hiddenLength)
}

// DecodeFloor decodes as a floor record
func (t *Tile) DecodeFloor(dw uint32) {
	t.decodeFloorShadow(dw)
}

// EncodeFloor adds Floor's bits to stream writer given
func (t *Tile) EncodeFloor(sw *d2datautils.StreamWriter) {
	t.encodeFloorShadow(sw)
}

// DecodeShadow decodes as a shadow record
func (t *Tile) DecodeShadow(dw uint32) {
	t.decodeFloorShadow(dw)
}

// EncodeShadow adds shadow's bits to stream writer given
func (t *Tile) EncodeShadow(sw *d2datautils.StreamWriter) {
	t.encodeFloorShadow(sw)
}
