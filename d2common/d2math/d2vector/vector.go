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

// Dot returns the dot product of this Vector and the given Vector.
func (v *Vector) Dot(o *Vector) float64 {
	return v.x*o.x + v.y*o.y
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

func (v Vector) String() string {
	return fmt.Sprintf("Vector{%g, %g}", v.x, v.y)
}

/*
// SetLength sets the length of this Vector
func (v *BigFloat) SetLength(length *big.Float) d2interface.Vector {
	return v.Normalize().Scale(length)
}

// NormalizeRightHand rotate this Vector to its perpendicular,
// in the positive direction.
func (v *BigFloat) NormalizeRightHand() d2interface.Vector {
	x := v.x
	v.x = v.y.Mul(v.y, big.NewFloat(negative1))
	v.y = x

	return v
}

// NormalizeLeftHand rotate this Vector to its perpendicular,
// in the negative1 direction.
func (v *BigFloat) NormalizeLeftHand() d2interface.Vector {
	x := v.x
	v.x = v.y
	v.y = x.Mul(x, big.NewFloat(negative1))

	return v
}



// Cross Calculate the cross product of this Vector and the given Vector.
func (v *BigFloat) Cross(src d2interface.Vector) *big.Float {
	c := v.Clone()
	c.XBig().Mul(c.XBig(), src.XBig())
	c.YBig().Mul(c.YBig(), src.YBig())

	return c.XBig().Sub(c.XBig(), c.YBig())
}

// Lerp Linearly interpolate between this Vector and the given Vector.
func (v *BigFloat) Lerp(
	src d2interface.Vector,
	t *big.Float,
) d2interface.Vector {
	vc, sc := v.Clone(), src.Clone()
	x, y := vc.XBig(), vc.YBig()
	v.x.Set(x.Add(x, t.Mul(t, sc.XBig().Sub(sc.XBig(), x))))
	v.y.Set(y.Add(y, t.Mul(t, sc.YBig().Sub(sc.YBig(), y))))

	return v
}

// Reset this Vector the zero vector (0, 0).
func (v *BigFloat) Reset() d2interface.Vector {
	v.x.SetFloat64(zero)
	v.y.SetFloat64(zero)

	return v
}

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

// Rotate this Vector by an angle amount.
func (v *BigFloat) Rotate(angle *big.Float) d2interface.Vector {
	// HACK we should do this only with big.Float, not float64
	// we are throwing away the precision here
	floatAngle, _ := angle.Float64()
	cos := math.Cos(floatAngle)
	sin := math.Sin(floatAngle)

	oldX, _ := v.x.Float64()
	oldY, _ := v.y.Float64()

	newX := big.NewFloat(cos*oldX - sin*oldY)
	newY := big.NewFloat(sin*oldX + cos*oldY)

	v.Set(newX, newY)

	return v
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
