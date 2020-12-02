package rectangle

import "github.com/gravestench/pho/geom/point"

// GetPoint calculates the coordinates of a point at a certain `position` on the
// Rectangle's perimeter, assigns to and returns the given point, or creates a point if nil.
func GetPoint(r *Rectangle, position float64, p *point.Point) *point.Point {
	if p == nil {
		p = point.New(0, 0)
	}

	if position <= 0 || position >= 1 {
		p.X, p.Y = r.X, r.Y
		return p
	}

	perimeter := Perimeter(r) * position

	if position > 0.5 {
		perimeter -= r.Width + r.Height

		if perimeter <= r.Width {
			// face 3
			p.X, p.Y = r.Right()-perimeter, r.Bottom()
		} else {
			// face 4
			p.X, p.Y = r.X, r.Bottom()-(perimeter-r.Width)
		}
	} else if position <= r.Width {
		// face 1
		p.X, p.Y = r.X+perimeter, r.Y
	} else {
		// face 2
		p.X, p.Y = r.Right(), r.Y+(perimeter-r.Width)
	}

	return p
}
