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
