package d2common

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2astar"

type PathTile struct {
	Walkable                                                    bool
	Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight *PathTile
	X, Y                                                        float64
}

func (t *PathTile) PathNeighbors() []d2astar.Pather {
	result := make([]d2astar.Pather, 0, 8)
	if t.Up != nil {
		result = append(result, t.Up)
	}
	if t.Right != nil {
		result = append(result, t.Right)
	}
	if t.Down != nil {
		result = append(result, t.Down)
	}
	if t.Left != nil {
		result = append(result, t.Left)
	}
	if t.UpLeft != nil {
		result = append(result, t.UpLeft)
	}
	if t.UpRight != nil {
		result = append(result, t.UpRight)
	}
	if t.DownLeft != nil {
		result = append(result, t.DownLeft)
	}
	if t.DownRight != nil {
		result = append(result, t.DownRight)
	}

	return result
}

func (t *PathTile) PathNeighborCost(to d2astar.Pather) float64 {
	return 1 // No cost specifics currently...
}

func (t *PathTile) PathEstimatedCost(to d2astar.Pather) float64 {
	toT := to.(*PathTile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	r := absX + absY

	return r
}
