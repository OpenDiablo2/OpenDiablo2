package rectangle

import "github.com/gravestench/pho/geom/point"

type RectangleNamespace interface {
	New(x, y, w, h float64) *Rectangle
	Contains(r *Rectangle, x, y float64) bool
	GetPoint(r *Rectangle, position float64, p *point.Point) *point.Point
	GetPoints(r *Rectangle, quantity int, stepRate float64, points []*point.Point) []*point.Point
	GetRandomPoint(r *Rectangle, p *point.Point) *point.Point
	ContainsPoint(r *Rectangle, p *point.Point) bool
	ContainsRectangle(r *Rectangle, other *Rectangle) bool
	Deconstruct(r *Rectangle, to []*point.Point) []*point.Point
	Equals(r *Rectangle, other *Rectangle) bool
	FitInside(r *Rectangle, other *Rectangle) *Rectangle
	Inflate(r *Rectangle, x, y float64) *Rectangle
	Intersection(r *Rectangle, other, intersect *Rectangle) *Rectangle
	MergePoints(r *Rectangle, points []*point.Point) *Rectangle
	MergeRectangle(r *Rectangle, other *Rectangle) *Rectangle
	MergeXY(r *Rectangle, x, y float64) *Rectangle
	Offset(r *Rectangle, x, y float64) *Rectangle
	OffsetPoint(r *Rectangle, p *point.Point) *Rectangle
	Overlaps(r *Rectangle, other *Rectangle) bool
	PerimeterPoint(r *Rectangle, angle float64, p *point.Point) *point.Point
	GetRandomPointOutside(r *Rectangle, other *Rectangle, out *point.Point) *point.Point
	SameDimensions(r *Rectangle, other *Rectangle) bool
	Scale(r *Rectangle, x, y float64) *Rectangle
	Union(r *Rectangle, other *Rectangle) *Rectangle
}

type Namespace struct{}

// New creates a new Rectangle instance.
func (*Namespace) New(x, y, w, h float64) *Rectangle {
	return New(x, y, w, h)
}

// Contains checks if the given x, y is inside the Rectangle's bounds.
func (*Namespace) Contains(r *Rectangle, x, y float64) bool {
	return Contains(r, x, y)
}

// GetPoint calculates the coordinates of a point at a certain `position` on the
// Rectangle's perimeter, assigns to and returns the given point, or creates a point if nil.
func (*Namespace) GetPoint(r *Rectangle, position float64, p *point.Point) *point.Point {
	return GetPoint(r, position, p)
}

// GetPoints returns a slice of points from the perimeter of the Rectangle,
// each spaced out based on the quantity or step required.
func (*Namespace) GetPoints(r *Rectangle, quantity int, stepRate float64,
	points []*point.Point) []*point.Point {
	return GetPoints(r, quantity, stepRate, points)
}

// GetRandomPoint returns a random point within the Rectangle's bounds.
func (*Namespace) GetRandomPoint(r *Rectangle, p *point.Point) *point.Point {
	return GetRandomPoint(r, p)
}

// Ceil rounds a Rectangle's position up to the smallest integer greater than or equal to each
// current coordinate.
func (*Namespace) Ceil(r *Rectangle) *Rectangle {
	return Ceil(r)
}

// CeilAll rounds a Rectangle's position and size up to the smallest
// integer greater than or equal to each respective value.
func (*Namespace) CeilAll(r *Rectangle) *Rectangle {
	return CeilAll(r)
}

// ContainsPoint checks if a given point is inside a Rectangle's bounds.
func (*Namespace) ContainsPoint(r *Rectangle, p *point.Point) bool {
	return Contains(r, p.X, p.Y)
}

// ContainsRect checks if a given point is inside a Rectangle's bounds.
func (*Namespace) ContainsRectangle(r, other *Rectangle) bool {
	return ContainsRectangle(r, other)
}

// CopyFrom copies the values of the given rectangle.
func (*Namespace) CopyFrom(target, source *Rectangle) *Rectangle {
	return CopyFrom(source, target)
}

// Deconstruct creates a slice of points for each corner of a Rectangle.
// If a slice is specified, each point object will be added to the end of the slice,
// otherwise a new slice will be created.
func (*Namespace) Deconstruct(r *Rectangle, to []*point.Point) []*point.Point {
	return Deconstruct(r, to)
}

// Equals compares the `x`, `y`, `width` and `height` properties of two rectangles.
func (*Namespace) Equals(a, b *Rectangle) bool {
	return Equals(a, b)
}

// Adjusts rectangle, changing its width, height and position,
// so that it fits inside the area of the source rectangle, while maintaining its original
// aspect ratio.
func (*Namespace) FitInside(inner, outer *Rectangle) *Rectangle {
	return FitInside(inner, outer)
}

func (*Namespace) Inflate(r *Rectangle, x, y float64) *Rectangle {
	return Inflate(r, x, y)
}

// Takes two Rectangles and first checks to see if they intersect.
// If they intersect it will return the area of intersection in the `out` Rectangle.
// If they do not intersect, the `out` Rectangle will have a width and height of zero.
// The given `intersect` rectangle will be assigned the intsersect values and returned.
// A new rectangle will be created if `intersect` is nil.
func (*Namespace) Intersection(r, other, intersect *Rectangle) *Rectangle {
	return Intersection(r, other, intersect)
}

// MergePoints adjusts this rectangle using a list of points by repositioning and/or resizing
// it such that all points are located on or within its bounds.
func (*Namespace) MergePoints(r *Rectangle, points []*point.Point) *Rectangle {
	return MergePoints(r, points)
}

// MergeRectangle merges the given rectangle into this rectangle and returns this rectangle.
// Neither rectangle should have a negative width or height.
func (*Namespace) MergeRectangle(r *Rectangle, other *Rectangle) *Rectangle {
	return MergeRectangle(r, other)
}

// MergeXY merges this rectangle with a point by repositioning and/or resizing it so that the
// point is/on or/within its bounds.
func (*Namespace) MergeXY(r *Rectangle, x, y float64) *Rectangle {
	return MergeXY(r, x, y)
}

// Offset nudges (translates) the top left corner of this Rectangle by a given offset.
func (*Namespace) Offset(r *Rectangle, x, y float64) *Rectangle {
	return Offset(r, x, y)
}

// OffsetPoint nudges (translates) the top left corner of this Rectangle by the coordinates of a
// point.
func (*Namespace) OffsetPoint(r *Rectangle, p *point.Point) *Rectangle {
	return OffsetPoint(r, p)
}

// Checks if this Rectangle overlaps with another rectangle.
func (*Namespace) Overlaps(r *Rectangle, other *Rectangle) bool {
	return Overlaps(r, other)
}

// PerimeterPoint returns a Point from the perimeter of the Rectangle based on the given angle.
func (*Namespace) PerimeterPoint(r *Rectangle, angle float64, p *point.Point) *point.Point {
	return PerimeterPoint(r, angle, p)
}

// Calculates a random point that lies within the `outer` Rectangle, but outside of the `inner`
// Rectangle. The inner Rectangle must be fully contained within the outer rectangle.
func (*Namespace) GetRandomPointOutside(r, other *Rectangle, out *point.Point) *point.Point {
	var outer, inner *Rectangle

	if r.ContainsRectangle(other) {
		outer, inner = r, other
	} else {
		outer, inner = other, r
	}

	return GetRandomPointOutside(outer, inner, out)
}

// SameDimensions determines if the two objects (either Rectangles or Rectangle-like) have the same
// width and height values under strict equality.
func (*Namespace) SameDimensions(r, other *Rectangle) bool {
	return SameDimensions(r, other)
}

// Scale the width and height of this Rectangle by the given amounts.
func (*Namespace) Scale(r *Rectangle, x, y float64) *Rectangle {
	return Scale(r, x, y)
}

// Union creates a new Rectangle or repositions and/or resizes an existing Rectangle so that it
// encompasses the two given Rectangles, i.e. calculates their union.
func (*Namespace) Union(r *Rectangle, other *Rectangle) *Rectangle {
	return Union(r, other, r)
}
