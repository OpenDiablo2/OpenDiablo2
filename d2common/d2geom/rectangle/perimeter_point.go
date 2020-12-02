package rectangle

import (
	"math"

	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath"
)

func PerimeterPoint(r *Rectangle, angle float64, out *point.Point) *point.Point {
	if out == nil {
		out = point.New(0, 0)
	}

	angle = phomath.DegToRad(angle)
	polarity := map[bool]float64{true: 1, false: -1}
	s, c := math.Sin(angle), math.Cos(angle)
	dx, dy := r.Width/2*polarity[c > 0], r.Height/2*polarity[s > 0]

	out.X = dx + r.CenterX()
	out.Y = dy + r.CenterY()

	return out
}
