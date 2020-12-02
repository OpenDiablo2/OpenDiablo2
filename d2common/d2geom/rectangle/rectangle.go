package rectangle

import (
	"github.com/gravestench/pho/geom"
	"github.com/gravestench/pho/geom/line"
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath"
)

// New creates a new rectangle
func New(x, y, w, h float64) *Rectangle {
	return &Rectangle{
		Type:   geom.Rectangle,
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}

// Encapsulates a 2D rectangle defined by its corner point in the top-left and its extends
// in x (width) and y (height)
type Rectangle struct {
	Type          geom.ShapeType
	X, Y          float64
	Width, Height float64
}

// Area calculates the area of the given rectangle.
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Left returns the left position of the rectangle
func (r *Rectangle) Left() float64 {
	return r.X
}

// SetLeft sets the left position of the rectangle, which also sets the x coordinate
func (r *Rectangle) SetLeft(value float64) *Rectangle {
	if value >= r.Right() {
		r.Width = 0
	} else {
		r.Width = r.Right() - value
	}

	r.X = value
	return r
}

// Right returns the right position of the rectangle
func (r *Rectangle) Right() float64 {
	return r.X + r.Width
}

// SetRight sets the right position of the rectangle, which adjusts the width
func (r *Rectangle) SetRight(value float64) *Rectangle {
	if value <= r.X {
		r.Width = 0
	} else {
		r.Width = value - r.X
	}

	r.X = value

	return r
}

// Top returns the top position of the rectangle
func (r *Rectangle) Top() float64 {
	return r.Y
}

// SetTop sets the top position of the rectangle, which also adjusts the Y coordinate
func (r *Rectangle) SetTop(value float64) *Rectangle {
	if value >= r.Bottom() {
		r.Height = 0
	} else {
		r.Height = r.Bottom() - value
	}

	return r
}

// Bottom returns the bottom position of the rectangle
func (r *Rectangle) Bottom() float64 {
	return r.Y + r.Height
}

// SetBottom sets the bottom position of the rectangle, which also adjusts the height
func (r *Rectangle) SetBottom(value float64) *Rectangle {
	if value <= r.Y {
		r.Height = 0
	} else {
		r.Height = value - r.Y
	}

	return r
}

func (r *Rectangle) CenterX() float64 {
	return r.X + r.Width/2
}

func (r *Rectangle) SetCenterX(value float64) *Rectangle {
	r.X = value - r.Width/2
	return r
}

func (r *Rectangle) CenterY() float64 {
	return r.Y + r.Height/2
}

func (r *Rectangle) SetCenterY(value float64) *Rectangle {
	r.Y = value - r.Height/2

	return r
}

// CenterOn moves the top-left corner of a Rectangle so that its center is at the given coordinates.
func (r *Rectangle) CenterOn(x, y float64) *Rectangle {
	return CenterOn(r, x, y)
}

// Contains checks if the given x, y is inside the Rectangle's bounds.
func (r *Rectangle) Contains(x, y float64) bool {
	return Contains(r, x, y)
}

// GetPoint calculates the coordinates of a point at a certain `position` on the
// Rectangle's perimeter, assigns to and returns the given point, or creates a point if nil.
func (r *Rectangle) GetPoint(position float64, p *point.Point) *point.Point {
	return GetPoint(r, position, p)
}

// GetPoints returns a slice of points from the perimeter of the Rectangle,
// each spaced out based on the quantity or step required.
func (r *Rectangle) GetPoints(quantity int, stepRate float64, points []*point.Point) []*point.Point {
	return GetPoints(r, quantity, stepRate, points)
}

// GetRandomPoint returns a random point within the Rectangle's bounds.
func (r *Rectangle) GetRandomPoint(p *point.Point) *point.Point {
	return GetRandomPoint(r, p)
}

// SetTo sets the position, width, and height of the Rectangle.
func (r *Rectangle) SetTo(x, y, w, h float64) *Rectangle {
	r.X, r.Y, r.Width, r.Height = x, y, w, h
	return r
}

// SetEmpty resets the position, width, and height of the Rectangle to 0.
func (r *Rectangle) SetEmpty() *Rectangle {
	return r.SetTo(0, 0, 0, 0)
}

//  SetPosition sets the position of the rectangle.
func (r *Rectangle) SetPosition(x, y float64) *Rectangle {
	r.X, r.Y = x, y
	return r
}

// SetSize sets the width and height of the rectangle.
func (r *Rectangle) SetSize(w, h float64) *Rectangle {
	r.Width, r.Height = w, h
	return r
}

// IsEmpty determines if the Rectangle is empty.
// A Rectangle is empty if its width or height is less than or equal to 0.
func (r *Rectangle) IsEmpty() bool {
	return r.Width <= phomath.Epsilon || r.Height <= phomath.Epsilon
}

// GetLineA returns a line object that corresponds to the top side of this rectangle.
// Assigns to the given line and returns it, or creates a new line if nil.
func (r *Rectangle) GetLineA(l *line.Line) *line.Line {
	if l == nil {
		l = line.New(0, 0, 0, 0)
	}

	l.SetTo(r.X, r.Y, r.Right(), r.Y)

	return l
}

// GetLineB returns a line object that corresponds to the right side of this rectangle.
// Assigns to the given line and returns it, or creates a new line if nil.
func (r *Rectangle) GetLineB(l *line.Line) *line.Line {
	if l == nil {
		l = line.New(0, 0, 0, 0)
	}

	l.SetTo(r.Right(), r.Y, r.Right(), r.Bottom())

	return l
}

// GetLineC returns a line object that corresponds to the bottom side of this rectangle.
// Assigns to the given line and returns it, or creates a new line if nil.
func (r *Rectangle) GetLineC(l *line.Line) *line.Line {
	if l == nil {
		l = line.New(0, 0, 0, 0)
	}

	l.SetTo(r.X, r.Bottom(), r.X, r.Y)

	return l
}

// GetLineD returns a line object that corresponds to the left side of this rectangle.
// Assigns to the given line and returns it, or creates a new line if nil.
func (r *Rectangle) GetLineD(l *line.Line) *line.Line {
	if l == nil {
		l = line.New(0, 0, 0, 0)
	}

	l.SetTo(r.X, r.Y, r.X, r.Bottom())

	return l
}

// Ceil rounds a Rectangle's position up to the smallest integer greater than or equal to each
// current coordinate.
func (r *Rectangle) Ceil() *Rectangle {
	return Ceil(r)
}

// CeilAll rounds a Rectangle's position and size up to the smallest
// integer greater than or equal to each respective value.
func (r *Rectangle) CeilAll() *Rectangle {
	return CeilAll(r)
}

// Clone this rectangle to a new rectangle instance
func (r *Rectangle) Clone() *Rectangle {
	return New(r.X, r.Y, r.Width, r.Height)
}

// ContainsPoint checks if a given point is inside a Rectangle's bounds.
func (r *Rectangle) ContainsPoint(p *point.Point) bool {
	return Contains(r, p.X, p.Y)
}

// ContainsRect checks if a given point is inside a Rectangle's bounds.
func (r *Rectangle) ContainsRectangle(other *Rectangle) bool {
	return ContainsRectangle(r, other)
}

// CopyFrom copies the values of the given rectangle.
func (r *Rectangle) CopyFrom(source *Rectangle) *Rectangle {
	return CopyFrom(source, r)
}

// Deconstruct creates a slice of points for each corner of a Rectangle.
// If a slice is specified, each point object will be added to the end of the slice,
// otherwise a new slice will be created.
func (r *Rectangle) Deconstruct(to []*point.Point) []*point.Point {
	return Deconstruct(r, to)
}

// Equals compares the `x`, `y`, `width` and `height` properties of two rectangles.
func (r *Rectangle) Equals(other *Rectangle) bool {
	return Equals(r, other)
}

// Adjusts rectangle, changing its width, height and position,
// so that it fits inside the area of the source rectangle, while maintaining its original
// aspect ratio.
func (r *Rectangle) FitInside(other *Rectangle) *Rectangle {
	return FitInside(r, other)
}

// GetCenter returns the center of the Rectangle as a Point.
func (r *Rectangle) GetCenter() *point.Point {
	return GetCenter(r)
}

// GetSize returns the size of the Rectangle, expressed as a Point object.
// With the value of the `width` as the `x` property and the `height` as the `y` property.
func (r *Rectangle) GetSize() *point.Point {
	return GetSize(r)
}

func (r *Rectangle) Inflate(x, y float64) *Rectangle {
	return Inflate(r, x, y)
}

// Takes two Rectangles and first checks to see if they intersect.
// If they intersect it will return the area of intersection in the `out` Rectangle.
// If they do not intersect, the `out` Rectangle will have a width and height of zero.
// The given `intersect` rectangle will be assigned the intsersect values and returned.
// A new rectangle will be created if `intersect` is nil.
func (r *Rectangle) Intersection(other, intersect *Rectangle) *Rectangle {
	return Intersection(r, other, intersect)
}

// MergePoints adjusts this rectangle using a list of points by repositioning and/or resizing
// it such that all points are located on or within its bounds.
func (r *Rectangle) MergePoints(points []*point.Point) *Rectangle {
	return MergePoints(r, points)
}

// MergeRectangle merges the given rectangle into this rectangle and returns this rectangle.
// Neither rectangle should have a negative width or height.
func (r *Rectangle) MergeRectangle(other *Rectangle) *Rectangle {
	return MergeRectangle(r, other)
}

// MergeXY merges this rectangle with a point by repositioning and/or resizing it so that the
// point is/on or/within its bounds.
func (r *Rectangle) MergeXY(x, y float64) *Rectangle {
	return MergeXY(r, x, y)
}

// Offset nudges (translates) the top left corner of this Rectangle by a given offset.
func (r *Rectangle) Offset(x, y float64) *Rectangle {
	return Offset(r, x, y)
}

// OffsetPoint nudges (translates) the top left corner of this Rectangle by the coordinates of a
// point.
func (r *Rectangle) OffsetPoint(p *point.Point) *Rectangle {
	return OffsetPoint(r, p)
}

// Checks if this Rectangle overlaps with another rectangle.
func (r *Rectangle) Overlaps(other *Rectangle) bool {
	return Overlaps(r, other)
}

// PerimeterPoint returns a Point from the perimeter of the Rectangle based on the given angle.
func (r *Rectangle) PerimeterPoint(angle float64, p *point.Point) *point.Point {
	return PerimeterPoint(r, angle, p)
}

// Calculates a random point that lies within the `outer` Rectangle, but outside of the `inner`
// Rectangle. The inner Rectangle must be fully contained within the outer rectangle.
func (r *Rectangle) GetRandomPointOutside(other *Rectangle, out *point.Point) *point.Point {
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
func (r *Rectangle) SameDimensions(other *Rectangle) bool {
	return SameDimensions(r, other)
}

// Scale the width and height of this Rectangle by the given amounts.
func (r *Rectangle) Scale(x, y float64) *Rectangle {
	return Scale(r, x, y)
}

// Union creates a new Rectangle or repositions and/or resizes an existing Rectangle so that it
// encompasses the two given Rectangles, i.e. calculates their union.
func (r *Rectangle) Union(other *Rectangle) *Rectangle {
	return Union(r, other, r)
}
