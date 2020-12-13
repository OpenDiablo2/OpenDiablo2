package rectangle

import (
	"math"

	"github.com/gravestench/pho/geom/intersects"
)

// Takes two Rectangles and first checks to see if they intersect.
// If they intersect it will return the area of intersection in the `out` Rectangle.
// If they do not intersect, the `out` Rectangle will have a width and height of zero.
// The given `output` rectangle will be assigned the intsersect values and returned.
// A new rectangle will be created if it is nil.
func Intersection(a, b, output *Rectangle) *Rectangle {
	if output == nil {
		output = New(0, 0, 0, 0)
	}

	if intersects.RectangleToRectangle(a, b) {
		output.X = math.Max(a.X, b.X)
		output.Y = math.Max(a.Y, b.Y)
		output.Width = math.Min(a.Right(), b.Right()) - output.X
		output.Height = math.Min(a.Bottom(), b.Bottom()) - output.Y
	} else {
		output.SetEmpty()
	}

	if output.Width < 0 {
		output.X += output.Width
		output.Width *= -1
	}

	if output.Height < 0 {
		output.Y += output.Height
		output.Height *= -1
	}

	return output
}
