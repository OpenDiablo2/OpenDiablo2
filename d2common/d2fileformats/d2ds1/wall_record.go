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
	Hidden      bool
	RandomIndex byte
	YAdjust     int
}
