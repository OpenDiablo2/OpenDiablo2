// Package d2vector provides an implementation of a 2D Euclidean vector using float64 to store the two values.
package d2vector

import (
	"fmt"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// Vector is an implementation of a Euclidean vector using float64 with common vector convenience methods.
type Vector struct {
	x, y float64
}

const (
	two                  float64 = 2
	entityDirectionCount float64 = 64 // 64 is the diablo equivalent of 360 degrees when dealing with entity rotation.
)

// NewVector creates a new Vector with the given x and y values.
func NewVector(x, y float64) Vector {
	return Vector{x, y}
}

// X returns the x value of this vector.
func (v *Vector) X() float64 {
	return v.x
}

// Y returns the y value of this vector.
func (v *Vector) Y() float64 {
	return v.y
}

// Equals returns true if the float64 values of this vector are exactly equal to the given Vector.
func (v *Vector) Equals(o Vector) bool {
	return v.x == o.x && v.y == o.y
}

// EqualsApprox returns true if the values of this Vector are approximately equal to those of the given Vector. If the
// difference between either of the value pairs is smaller than d2math.Epsilon, they will be considered equal.
func (v *Vector) EqualsApprox(o Vector) bool {
	x, y := v.CompareApprox(o)
	return x == 0 && y == 0
}

// CompareApprox returns 2 ints describing the difference between the vectors. If the difference between either of the
// value pairs is smaller than d2math.Epsilon, they will be considered equal.
func (v *Vector) CompareApprox(o Vector) (x, y int) {
	return d2math.CompareFloat64Fuzzy(v.x, o.x),
		d2math.CompareFloat64Fuzzy(v.y, o.y)
}

// TODO: untested method
// IsZero returns true if this vector's values are both exactly zero.
func (v *Vector) IsZero() bool {
	return v.x == 0 && v.y == 0
}

// Set the vector values to the given float64 values.
func (v *Vector) Set(x, y float64) *Vector {
	v.x = x
	v.y = y

	return v
}

// Clone returns a new a copy of this Vector.
func (v *Vector) Clone() Vector {
	return NewVector(v.x, v.y)
}

// Copy sets this vector's values to those of the given vector.
func (v *Vector) Copy(o *Vector) *Vector {
	v.x = o.x
	v.y = o.y

	return v
}

// Floor rounds the vector down to the nearest whole numbers.
func (v *Vector) Floor() *Vector {
	v.x = math.Floor(v.x)
	v.y = math.Floor(v.y)

	return v
}

// Clamp limits the values of v to those of a and b. If the values of v are between those of a and b they will be
// unchanged.
func (v *Vector) Clamp(a, b *Vector) *Vector {
	v.x = d2math.ClampFloat64(v.x, a.x, b.x)
	v.y = d2math.ClampFloat64(v.y, a.y, b.y)

	return v
}

// Add the given vector to this vector.
func (v *Vector) Add(o *Vector) *Vector {
	v.x += o.x
	v.y += o.y

	return v
}

// AddScalar the given vector to this vector.
func (v *Vector) AddScalar(s float64) *Vector {
	v.x += s
	v.y += s

	return v
}

// Subtract the given vector from this vector.
func (v *Vector) Subtract(o *Vector) *Vector {
	v.x -= o.x
	v.y -= o.y

	return v
}

// Multiply this Vector by the given Vector.
func (v *Vector) Multiply(o *Vector) *Vector {
	v.x *= o.x
	v.y *= o.y

	return v
}

// Scale multiplies both values of this vector by a single given value.
func (v *Vector) Scale(s float64) *Vector {
	v.x *= s
	v.y *= s

	return v
}

// Divide this vector by the given vector.
func (v *Vector) Divide(o *Vector) *Vector {
	v.x /= o.x
	v.y /= o.y

	return v
}

// DivideScalar divides this vector by the given float64 value.
func (v *Vector) DivideScalar(s float64) *Vector {
	v.x /= s
	v.y /= s

	return v
}

// Abs sets the vector to it's absolute (positive) equivalent.
func (v *Vector) Abs() *Vector {
	xm, ym := 1.0, 1.0
	if v.x < 0 {
		xm = -1
	}

	if v.y < 0 {
		ym = -1
	}

	v.x *= xm
	v.y *= ym

	return v
}

// Negate multiplies this vector by -1.
func (v *Vector) Negate() *Vector {
	return v.Scale(-1)
}

// Distance between this Vector's position and that of the given Vector.
func (v *Vector) Distance(o Vector) float64 {
	delta := o.Clone()
	delta.Subtract(v)

	return delta.Length()
}

// Length (magnitude/quantity) of this Vector.
func (v *Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

// SetLength sets the length of this Vector without changing the direction.
func (v *Vector) SetLength(length float64) *Vector {
	v.Normalize()
	v.Scale(length)

	return v
}

// Lerp sets this vector to the linear interpolation between this and the given vector. The interp argument determines
// the distance between the two vectors. An interp of 0 will return this vector and 1 will return the given vector.
func (v *Vector) Lerp(o *Vector, interp float64) *Vector {
	v.x = d2math.Lerp(v.x, o.x, interp)
	v.y = d2math.Lerp(v.y, o.y, interp)

	return v
}

// Dot returns the dot product of this Vector and the given Vector.
func (v *Vector) Dot(o *Vector) float64 {
	return v.x*o.x + v.y*o.y
}

// Cross returns the cross product of this Vector and the given Vector. Note: Cross product is specific to 3D space.
// This a not cross product. It is the Z component of a 3D vector cross product calculation. The X and Y components use
// the value of z which doesn't exist in 2D. See:
// https://stackoverflow.com/questions/243945/calculating-a-2d-vectors-cross-product
//
// The sign of Cross indicates whether the direction between the points described by vectors v and o around the origin
// (0,0) moves clockwise or anti-clockwise. The perspective is from the would-be position of positive Z and the
// direction is from v to o.
//
// Negative = clockwise
// Positive = anti-clockwise
// 0 = vectors are identical.
func (v *Vector) Cross(o Vector) float64 {
	return v.x*o.y - v.y*o.x
}

// Normalize sets the vector length to 1 without changing the direction. The normalized vector may be scaled by the
// float64 return value to restore it's original length.
func (v *Vector) Normalize() float64 {
	if v.x == 0 && v.y == 0 {
		return 0
	}

	multiplier := 1 / v.Length()
	v.Scale(multiplier)

	return 1 / multiplier
}

// Angle computes the unsigned angle in radians from this vector to the given vector. This angle will never exceed half
// a full circle. For angles describing a full circumference use SignedAngle.
func (v *Vector) Angle(o Vector) float64 {
	if v.IsZero() || o.IsZero() {
		return 0
	}

	from := v.Clone()
	from.Normalize()

	to := o.Clone()
	to.Normalize()

	denominator := math.Sqrt(from.Length() * to.Length())
	dotClamped := d2math.ClampFloat64(from.Dot(&to)/denominator, -1, 1)

	return math.Acos(dotClamped)
}

// SignedAngle computes the signed (clockwise) angle in radians from this vector to the given vector.
func (v *Vector) SignedAngle(o Vector) float64 {
	unsigned := v.Angle(o)
	sign := d2math.Sign(v.x*o.y - v.y*o.x)

	if sign > 0 {
		return d2math.RadFull - unsigned
	}

	return unsigned
}

// TODO: Move this to Position and take a target argument, call it "DirectionTo"
// Direction returns the anti-clockwise angle looking from this vector to the given vector. Positive X is zero.
func (v *Vector) Direction() float64 {
	angle := v.SignedAngle(VectorRight())
	radiansPerDirection := d2math.RadFull / entityDirectionCount
	offset := (45 / d2math.RadToDeg) - (radiansPerDirection / 2)
	newDirection := (angle - offset) / radiansPerDirection

	if newDirection >= 64 {
		newDirection = newDirection - 64
	} else if newDirection < 0 {
		newDirection = 64 + newDirection
	}

	return newDirection
}

// Reflect sets this Vector to it's reflection off a line defined by the given normal.
func (v *Vector) Reflect(normal Vector) *Vector {
	normal.Normalize()
	undo := v.Normalize()

	// 1*Dot is the directional (ignoring length) difference between the vector and the normal. Therefore 2*Dot takes
	// us beyond the normal to the angle with the equivalent distance in the other direction i.e. the reflection.
	normal.Scale(two * v.Dot(&normal))
	v.Subtract(&normal)
	v.Scale(undo)

	return v
}

// ReflectSurface does the same thing as Reflect, except the given vector describes,
// the surface line, not it's normal.
func (v *Vector) ReflectSurface(surface Vector) *Vector {
	v.Reflect(surface).Negate()

	return v
}

// Rotate moves the vector around it's origin clockwise, by the given angle in radians.
func (v *Vector) Rotate(angle float64) *Vector {
	a := -angle
	x := v.x*math.Cos(a) - v.y*math.Sin(a)
	y := v.x*math.Sin(a) + v.y*math.Cos(a)
	v.x = x
	v.y = y

	return v
}

// NinetyAnti rotates this vector by 90 degrees anti-clockwise.
func (v *Vector) NinetyAnti() *Vector {
	x := v.x
	v.x = v.y * -1
	v.y = x

	return v
}

// NinetyClock rotates this vector by 90 degrees clockwise.
func (v *Vector) NinetyClock() *Vector {
	y := v.y
	v.y = v.x * -1
	v.x = y

	return v
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector{%.3f, %.3f}", v.x, v.y)
}

// VectorUp returns a new vector (0, 1)
func VectorUp() Vector {
	return NewVector(0, 1)
}

// VectorDown returns a new vector (0, -1)
func VectorDown() Vector {
	return NewVector(0, -1)
}

// VectorRight returns a new vector (1, 0)
func VectorRight() Vector {
	return NewVector(1, 0)
}

// VectorLeft returns a new vector (-1, 0)
func VectorLeft() Vector {
	return NewVector(-1, 0)
}

// VectorOne returns a new vector (1, 1)
func VectorOne() Vector {
	return NewVector(1, 1)
}

// VectorZero returns a new vector (0, 0)
func VectorZero() Vector {
	return NewVector(0, 0)
}
