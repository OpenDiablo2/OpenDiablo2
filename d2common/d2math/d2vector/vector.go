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

// NewVector creates a new Vector and returns a pointer to it.
func NewVector(x, y float64) Vector {
	return Vector{x, y}
}

// Equals check whether this Vector is equal to a given Vector.
func (f *Vector) Equals(v Vector) bool {
	return f.x == v.x && f.y == v.y
}

// EqualsApprox checks if the Vector is approximately equal
// to the given Vector.
func (f *Vector) EqualsApprox(v Vector) bool {
	x, y := f.CompareApprox(v)
	return x == 0 && y == 0
}

// CompareApprox performs a fuzzy comparison and returns 2
// ints represending the -1 to 1 comparison of x and y.
func (f *Vector) CompareApprox(v Vector) (x, y int) {
	return d2math.CompareFloat64Fuzzy(f.x, v.x),
		d2math.CompareFloat64Fuzzy(f.y, v.y)
}

// Set sets the vector values to the given float64 values.
func (f *Vector) Set(x, y float64) *Vector {
	f.x = x
	f.y = y

	return f
}

// Clone creates a copy of this Vector.
func (f *Vector) Clone() Vector {
	return NewVector(f.x, f.y)
}

// Floor rounds the vector down to the nearest whole numbers.
func (f *Vector) Floor() *Vector {
	f.x = math.Floor(f.x)
	f.y = math.Floor(f.y)

	return f
}

// Add to this Vector the components of the given Vector.
func (f *Vector) Add(v *Vector) *Vector {
	f.x += v.x
	f.y += v.y

	return f
}

// Subtract from this Vector from the components of the given Vector.
func (f *Vector) Subtract(v *Vector) *Vector {
	f.x -= v.x
	f.y -= v.y

	return f
}

// Multiply this Vector by the components of the given Vector.
func (f *Vector) Multiply(v *Vector) *Vector {
	f.x *= v.x
	f.y *= v.y

	return f
}

// Scale multiplies this vector by a single value.
func (f *Vector) Scale(s float64) *Vector {
	f.x *= s
	f.y *= s

	return f
}

// Divide divides this vector by the components of the given vector.
func (f *Vector) Divide(v *Vector) *Vector {
	f.x /= v.x
	f.y /= v.y

	return f
}

// Abs sets the vector to it's absolute (positive) equivalent.
func (f *Vector) Abs() *Vector {
	xm, ym := 1.0, 1.0
	if f.x < 0 {
		xm = -1
	}

	if f.y < 0 {
		ym = -1
	}

	f.x *= xm
	f.y *= ym

	return f
}

// Negate multiplies the vector by -1.
func (f *Vector) Negate() *Vector {
	return f.Scale(-1)
}

// Distance calculate the distance between this Vector and the given Vector.
func (f *Vector) Distance(v Vector) float64 {
	delta := v.Clone()
	delta.Subtract(f)

	return delta.Length()
}

// Length returns the length of this Vector.
func (f *Vector) Length() float64 {
	return math.Sqrt(f.Dot(f))
}

// Dot returns the dot product of this Vector and the given Vector.
func (f *Vector) Dot(v *Vector) float64 {
	return f.x*v.x + f.y*v.y
}

// Angle computes the angle in radians with respect to the
// positive x-axis
func (f *Vector) Angle() float64 {
	floatAngle := math.Atan2(f.y, f.x)

	if floatAngle < 0 {
		floatAngle += 2.0 * math.Pi
	}

	return floatAngle
}

// SetAngle sets the angle of this Vector
func (f *Vector) SetAngle(angle float64) *Vector {
	return f.SetToPolar(angle, f.Length())
}

// SetToPolar sets the `x` and `y` values of this object from a
// given polar coordinate.
// TODO: How does this work?
func (f *Vector) SetToPolar(azimuth, radius float64) *Vector {
	// HACK we should do this better, with the big.Float
	f.x = math.Cos(azimuth) * radius
	f.y = math.Sin(azimuth) * radius

	return f
}

func (f Vector) String() string {
	return fmt.Sprintf("Vector{%g, %g}", f.x, f.y)
}

/*


// SetAngle sets the angle of this Vector
func (v *BigFloat) SetAngle(angle *big.Float) d2interface.Vector {
	return v.SetToPolar(angle, v.Length())
}


// SetLength sets the length of this Vector
func (v *BigFloat) SetLength(length *big.Float) d2interface.Vector {
	return v.Normalize().Scale(length)
}

// Normalize Makes the vector a unit length vector (magnitude of 1) in the same
// direction.
func (v *BigFloat) Normalize() d2interface.Vector {
	xsq, ysq := v.LengthSq()
	length := big.NewFloat(zero).Add(xsq, ysq)
	one := big.NewFloat(1.0)

	if length.Cmp(one) > 0 {
		length.Quo(one, length.Sqrt(length))

		v.x.Mul(v.x, length)
		v.y.Mul(v.y, length)
	}

	return v
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
