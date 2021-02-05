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
