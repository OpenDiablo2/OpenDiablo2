package rectangle

import (
	"github.com/gravestench/pho/geom/point"
)

// ByStepRate is a special value that tells GetPoints to use the stepRate instead of quantity
// for generating perimeter points
const ByStepRate = -1

// GetPoints returns a slice of points from the perimeter of the Rectangle,
// each spaced out based on the quantity or step required.
func GetPoints(r *Rectangle, quantity int, stepRate float64, points []*point.Point) []*point.Point {
	if quantity == ByStepRate {
		quantity = int(Perimeter(r) / stepRate)
	}

	if points == nil {
		points = make([]*point.Point, 0)
	}

	for len(points) < quantity {
		points = append(points, nil)
	}

	for idx := 0; idx < quantity; idx++ {
		position := float64(idx) / float64(quantity)

		points[idx] = GetPoint(r, position, nil)
	}

	return points
}
