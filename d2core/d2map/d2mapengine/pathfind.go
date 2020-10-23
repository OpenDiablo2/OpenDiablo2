package d2mapengine

import (
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// PathFind finds a path between given start and dest positions and returns the positions of the path
func (m *MapEngine) PathFind(start, dest d2vector.Position) []d2vector.Position {
	points := make([]d2vector.Position, 0)
	_, point := m.checkLos(start, dest)
	points = append(points, point)

	return points
}

// checkLos finds out if there is a clear line of sight between two points
func (m *MapEngine) checkLos(start, end d2vector.Position) (bool, d2vector.Position) {
	dv := d2vector.Position{Vector: *end.Clone()}
	dv.Subtract(&start.Vector)
	dx := dv.X()
	dy := dv.Y()
	N := math.Max(math.Abs(dx), math.Abs(dy))

	var divN float64
	if N == 0 {
		divN = 0.0
	} else {
		divN = 1.0 / N // nolint:gomnd // we're just taking inverse...
	}

	xstep := dx * divN
	ystep := dy * divN
	x := start.X()
	y := start.Y()

	for i := 0; i <= int(N); i++ {
		x += xstep
		y += ystep

		if m.SubTileAt(int(math.Floor(x)), int(math.Floor(y))).BlockWalk {
			return false, d2vector.NewPosition(x-xstep, y-ystep)
		}
	}

	return true, end
}
