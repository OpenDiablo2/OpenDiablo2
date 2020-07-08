package d2vector

import (
	"fmt"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// Vector is the implementation of Vector using float64.
type Vector struct {
	x, y float64
}

// NewVector creates a new Vector with the given x and y values.
func NewVector(x, y float64) Vector {
	return Vector{x, y}
}

// Equals check whether this Vector is equal to a given Vector.
func (v *Vector) Equals(o Vector) bool {
	return v.x == o.x && v.y == o.y
}

// EqualsApprox checks if the Vector is approximately equal
// to the given Vector.
func (v *Vector) EqualsApprox(o Vector) bool {
	x, y := v.CompareApprox(o)
	return x == 0 && y == 0
}

// CompareApprox performs a fuzzy comparison and returns 2
// ints represending the -1 to 1 comparison of x and y.
func (v *Vector) CompareApprox(o Vector) (x, y int) {
	return d2math.CompareFloat64Fuzzy(v.x, o.x),
		d2math.CompareFloat64Fuzzy(v.y, o.y)
}

// Set sets the vector values to the given float64 values.
func (v *Vector) Set(x, y float64) *Vector {
	v.x = x
	v.y = y

	return v
}

// Clone creates a copy of this Vector.
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

// Clamp limits the values of v to those of a and b. If the
// values of v are between those of a and b they will be
// unchanged.
func (v *Vector) Clamp(a, b *Vector) *Vector {
	v.x = d2math.ClampFloat64(v.x, a.x, b.x)
	v.y = d2math.ClampFloat64(v.y, a.y, b.y)

	return v
}

// Add to this Vector the components of the given Vector.
func (v *Vector) Add(o *Vector) *Vector {
	v.x += o.x
	v.y += o.y

	return v
}

// Subtract from this Vector from the components of the given Vector.
func (v *Vector) Subtract(o *Vector) *Vector {
	v.x -= o.x
	v.y -= o.y

	return v
}

// Multiply this Vector by the components of the given Vector.
func (v *Vector) Multiply(o *Vector) *Vector {
	v.x *= o.x
	v.y *= o.y

	return v
}

// Scale multiplies this vector by a single value.
func (v *Vector) Scale(s float64) *Vector {
	v.x *= s
	v.y *= s

	return v
}

// Divide divides this vector by the components of the given vector.
func (v *Vector) Divide(o *Vector) *Vector {
	v.x /= o.x
	v.y /= o.y

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

// Negate multiplies the vector by -1.
func (v *Vector) Negate() *Vector {
	return v.Scale(-1)
}

// Distance calculate the distance between this Vector and the given Vector.
func (v *Vector) Distance(o Vector) float64 {
	delta := o.Clone()
	delta.Subtract(v)

	return delta.Length()
}

// Length returns the length of this Vector.
func (v *Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

// SetLength sets the length of this Vector
func (v *Vector) SetLength(length float64) *Vector {
	v.Normalize()
	v.Scale(length)

	return v
}

// Lerp linearly interpolates this vector toward the given vector.
// The interp argument describes the distance between v and o where
// 0 is v and 1 is o.
func (v *Vector) Lerp(o *Vector, interp float64) *Vector {
	v.x = d2math.Lerp(v.x, o.x, interp)
	v.y = d2math.Lerp(v.y, o.y, interp)

	return v
}

// Dot returns the dot product of this Vector and the given Vector.
func (v *Vector) Dot(o *Vector) float64 {
	return v.x*o.x + v.y*o.y
}

// Cross returns the cross product of this Vector and the given Vector.
// Note: Cross product is specific to 3D space. This a not cross product.
// It is the Z component of a 3D vector cross product calculation. The X
// and Y components use the value of z which doesn't exist in 2D.
// See: https://stackoverflow.com/questions/243945/calculating-a-2d-vectors-cross-product
//
// The sign of Cross indicates whether the direction between the points
// described by vectors v and o around the origin (0,0) moves clockwise
// or anti-clockwise. The perspective is from the would-be position of
// positive Z and the direction is from v to o.
//
// Negative = clockwise
// Positive = anti-clockwise
// 0 = vectors are identical.
func (v *Vector) Cross(o Vector) float64 {
	return v.x*o.y - v.y*o.x
}

// Normalize sets the vector length to 1 without changing the direction.
func (v *Vector) Normalize() *Vector {
	v.Scale(1 / v.Length())
	return v
}

// Angle computes the unsigned angle in radians from this vector to the given vector.
func (v *Vector) Angle(o Vector) float64 {
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

/*// Iterates through the elements of this vector and for each element invokes
// the function.
func (self Vector) Do(applyFn func(float64) float64) {
	for i, e := range self {
		self[i] = applyFn(e)
	}
}*/

/*







// Limit the length (or magnitude) of this Vector
func (v *BigFloat) Limit(max *big.Float) d2interface.Vector {
	length := v.Length()

	if max.Cmp(length) < 0 {
		v.Scale(length.Quo(max, length))
	}

	return v
}

// Reflect this Vector off a line defined by a normal.
func (v *BigFloat) Reflect(normal d2interface.Vector) d2interface.Vector {
	clone := v.Clone()
	clone.Normalize()

	two := big.NewFloat(2.0) // there's some matrix algebra magic here
	dot := v.Clone().Dot(normal)
	normal.Scale(two.Mul(two, dot))

	return v.Subtract(normal)
}

// Mirror reflect this Vector across another.
func (v *BigFloat) Mirror(axis d2interface.Vector) d2interface.Vector {
	return v.Reflect(axis).Negate()
}




// BigFloatUp returns a new vector (0, 1)
func BigFloatUp() d2interface.Vector {
	return NewBigFloat(0, 1)
}

// BigFloatDown returns a new vector (0, -1)
func BigFloatDown() d2interface.Vector {
	return NewBigFloat(0, -1)
}

// BigFloatRight returns a new vector (1, 0)
func BigFloatRight() d2interface.Vector {
	return NewBigFloat(1, 0)
}

// BigFloatLeft returns a new vector (-1, 0)
func BigFloatLeft() d2interface.Vector {
	return NewBigFloat(-1, 0)
}

// BigFloatOne returns a new vector (1, 1)
func BigFloatOne() d2interface.Vector {
	return NewBigFloat(1, 1)
}

// BigFloatZero returns a new vector (0, 0)
func BigFloatZero() d2interface.Vector {
	return NewBigFloat(0, 0)
}*/
