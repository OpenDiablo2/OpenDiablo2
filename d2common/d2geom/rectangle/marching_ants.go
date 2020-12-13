package rectangle

import (
	"math"

	"github.com/gravestench/pho/geom/point"
)

const (
	AutoStep     = -1
	AutoQuantity = -1
)

// Returns an array of points from the perimeter of the Rectangle,
// where each point is spaced out based on either the `step` value, or the `quantity`.
func MarchingAnts(r *Rectangle, step float64, quantity int, out []*point.Point) []*point.Point {
	if step <= 0 {
		step = AutoStep
	}

	if quantity <= 0 {
		quantity = AutoQuantity
	}

	if out == nil {
		out = make([]*point.Point, quantity)
	}

	if step == AutoStep && quantity == AutoQuantity {
		return /* bail */ out
	}

	if step == AutoStep {
		step = Perimeter(r) / float64(quantity)
	} else {
		quantity = int(math.Round(Perimeter(r) / step))
	}

	const (
		top = iota
		right
		bottom
		left
		numFaces
	)

	x, y := r.X, r.Y
	face := top

	for idx := 0; idx < quantity; idx++ {
		out = append(out, point.New(x, y))

		switch face {
		case top:
			x += step

			if x >= r.Right() {
				face = (face + 1) % numFaces
				y += x - r.Right()
				x = r.Right()
			}
		case right:
			y += step

			if y >= r.Bottom() {
				face = (face + 1) % numFaces
				x -= y - r.Bottom()
				y = r.Bottom()
			}
		case bottom:
			x -= step

			if x <= r.Left() {
				face = (face + 1) % numFaces
				y -= r.Left() - x
				x = r.Left()
			}
		case left:
			y -= step

			if y <= r.Top() {
				face = (face + 1) % numFaces
				y = r.Top()
			}
		}
	}

	return out
}
