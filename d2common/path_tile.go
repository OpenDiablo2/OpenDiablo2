package d2common

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2astar"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// PathTile represents a node in path finding
type PathTile struct {
	Walkable            bool
	Up, Down            *PathTile
	Left, Right         *PathTile
	UpLeft, UpRight     *PathTile
	DownLeft, DownRight *PathTile
	Position            d2vector.Position
}

// PathNeighbors returns the direct neighboring nodes of this node which can be pathed to
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

// PathNeighborCost calculates the exact movement cost to neighbor nodes
func (t *PathTile) PathNeighborCost(_ d2astar.Pather) float64 {
	return 1 // No cost specifics currently...
}

// PathEstimatedCost is a heuristic method for estimating movement costs between non-adjacent nodes
func (t *PathTile) PathEstimatedCost(to d2astar.Pather) float64 {
	delta := to.(*PathTile).Position.Clone()
	delta.Subtract(&t.Position.Vector)
	delta.Abs()

	return delta.X() + delta.Y()
}
