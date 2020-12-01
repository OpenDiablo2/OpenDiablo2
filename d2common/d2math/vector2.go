package d2math

import "math"

// Vector2Like is something that has an XY method that returns x and y coordinate values
type Vector2Like interface {
	XY() (float64, float64)
}

// static check that Vector2 is Vector2Like
var _ Vector2Like = &Vector2{}

// NewVector2 creates a new Vector2
func NewVector2(x, y float64) *Vector2 {
	return &Vector2{
		X: x,
		Y: y,
	}
}

// Vector2 is a representation of a vector in 2D space.
type Vector2 struct {
	X, Y float64
}

// XY returns the x and y values
func (v *Vector2) XY() (x, y float64) {
	return v.X, v.Y
}

// Clone makes a clone of this Vector2.
func (v *Vector2) Clone() *Vector2 {
	return NewVector2(v.X, v.Y)
}

// Copy makes a clone of this Vector2.
func (v *Vector2) Copy(source *Vector2) *Vector2 {
	return v.Set(source.X, source.Y)
}

// SetFromVectorLike sets the x, y values of this Vector from a given Vector2Like object.
func (v *Vector2) SetFromVectorLike(l Vector2Like) *Vector2 {
	return v.Set(l.XY())
}

// Set the `x` and `y` components of the this Vector to the given `x` and `y` values.
func (v *Vector2) Set(x, y float64) *Vector2 {
	v.X, v.Y = x, y
	return v
}

// SetTo is an alias for `Vector2.Set`.
func (v *Vector2) SetTo(x, y float64) *Vector2 {
	return v.Set(x, y)
}

// SetToPolar sets the `x` and `y` values of this object from a given polar coordinate.
func (v *Vector2) SetToPolar(azimuth, radius float64) *Vector2 {
	return v.Set(math.Cos(azimuth)*radius, math.Sin(azimuth)*radius)
}

// Equals checks whether this Vector is equal to a given Vector.
func (v *Vector2) Equals(other *Vector2) bool {
	return math.Abs(v.X-other.X) < Epsilon && math.Abs(v.Y-other.Y) < Epsilon
}

// Angle calculates the angle between this Vector and the positive x-v.Xis, in radians.
func (v *Vector2) Angle() float64 {
	angle := math.Atan2(v.XY())

	if angle < 0 {
		angle = PI2
	}

	return angle
}

// SetAngle sets the angle of this Vector.
func (v *Vector2) SetAngle(angle float64) *Vector2 {
	return v.SetToPolar(angle, v.Length())
}

// Add the given Vector to this Vector.
// Addition is component-wise and mutates the vector.
func (v *Vector2) Add(source *Vector2) *Vector2 {
	return v.Set(v.X+source.X, v.Y+source.Y)
}

// Subtract the given Vector from this Vector.
// Subtraction is component-wise and mutates the vector.
func (v *Vector2) Subtract(source *Vector2) *Vector2 {
	return v.Set(v.X-source.X, v.Y-source.Y)
}

// Multiply this vector with the given Vector.
// Multiplication is component-wise and mutates the vector.
func (v *Vector2) Multiply(source *Vector2) *Vector2 {
	return v.Set(v.X*source.X, v.Y*source.Y)
}

// Scale this Vector with a scalar.
// Multiplication is component-wise and mutates the vector.
func (v *Vector2) Scale(scalar float64) *Vector2 {
	return v.Set(v.X*scalar, v.Y*scalar)
}

// Divide this vector by the given Vector.
// Division is component-wise and mutates the vector.
func (v *Vector2) Divide(source *Vector2) *Vector2 {
	return v.Set(v.X/source.X, v.Y/source.Y)
}

// Negate the x/y values of this vector.
func (v *Vector2) Negate() *Vector2 {
	return v.Set(-v.X, -v.Y)
}

// Distance calculate the distance between this Vector and the given Vector.
func (v *Vector2) Distance(source *Vector2) float64 {
	return math.Sqrt(v.DistanceSquared(source))
}

// DistanceSquared calculate the distance squared between this Vector and the given Vector.
func (v *Vector2) DistanceSquared(source *Vector2) float64 {
	dx, dy := source.X-v.X, source.Y-v.Y
	return dx*dx + dy*dy
}

// Length calculates the length (or magnitude) of this Vector.
func (v *Vector2) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared calculates the length squared of this Vector.
func (v *Vector2) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Length calculates the length (or magnitude) of this Vector.
func (v *Vector2) SetLength(l float64) *Vector2 {
	return v.Normalize().Scale(l)
}

// Normalize this Vector to length of 1
func (v *Vector2) Normalize() *Vector2 {
	l := v.Length()

	if l > 0 {
		l = 1 / math.Sqrt(l)
		v.Scale(l)
	}

	return v
}

// NormalizeRightHand rotates this Vector to its perpendicular, in the positive direction.
func (v *Vector2) NormalizeRightHand() *Vector2 {
	return v.Set(v.Y*-1, v.X)
}

// NormalizeLeftHand rotates this Vector to its perpendicular, in the negative direction.
func (v *Vector2) NormalizeLeftHand() *Vector2 {
	return v.Set(v.Y, v.X*-1)
}

// Dot calculate the dot product of this Vector and the given Vector.
func (v *Vector2) Dot(other *Vector2) float64 {
	return v.X*other.X + v.Y + other.Y
}

// Cross calculate the dot product of this Vector and the given Vector.
func (v *Vector2) Cross(other *Vector2) float64 {
	return v.X*other.X - v.Y + other.Y
}

// Lerp linearly interpolates between this Vector and the given Vector.
func (v *Vector2) Lerp(other *Vector2, t float64) *Vector2 {
	return v.Set(v.X+t*(other.X-v.X), v.Y+t*(other.Y-v.Y))
}

// TransformMat3 transforms this Vector with the given Matrix3.
func (v *Vector2) TransformMat3(m3 *Matrix3) *Vector2 {
	m := m3.Values
	return v.Set(m[0]*v.X+m[3]*v.Y+m[6], m[1]*v.X+m[4]*v.Y+m[7])
}

// TransformMat4 transforms this Vector with the given Matrix4.
func (v *Vector2) TransformMat4(m4 *Matrix4) *Vector2 {
	m := m4.Values
	return v.Set(m[0]*v.X+m[4]*v.Y+m[12], m[1]*v.X+m[5]*v.Y+m[13])
}

// Reset makes this Vector the zero vector (0, 0).
func (v *Vector2) Reset() *Vector2 {
	return v.Set(0, 0)
}

// Limit the length (or magnitude) of this Vector.
func (v *Vector2) Limit(l float64) *Vector2 {
	if v.Length() > l {
		v.SetLength(l)
	}

	return v
}

// Reflect this Vector off a line defined by a normal.
func (v *Vector2) Reflect(other *Vector2) *Vector2 {
	normal := other.Clone().Normalize()

	return v.Subtract(normal.Scale(2 * v.Dot(normal)))
}

// Mirror reflects this Vector across another.
func (v *Vector2) Mirror(axis *Vector2) *Vector2 {
	return v.Reflect(axis).Negate()
}

// Rotate this Vector by an angle amount.
func (v *Vector2) Rotate(radians float64) *Vector2 {
	c, s := math.Cos(radians), math.Sin(radians)

	return v.Set(c*v.X-s*v.Y, s*v.X+c*v.Y)
}
