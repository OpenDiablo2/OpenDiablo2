package d2ds1

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2path"
)

// Object is a game world object
type Object struct {
	Type  int
	ID    int
	X     int
	Y     int
	Flags int
	Paths []d2path.Path
}

// Equals checks if this Object is equivalent to the given Object
func (o *Object) Equals(other *Object) bool {
	return o.Type == other.Type &&
		o.ID == other.ID &&
		o.X == other.X &&
		o.Y == other.Y &&
		o.Flags == other.Flags &&
		len(o.Paths) == len(other.Paths)
}
