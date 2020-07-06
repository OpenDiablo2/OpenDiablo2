// Package d2vector is an Implementation of 2-dimensional vectors with big.Float components
package d2vector

import (
	"fmt"
	"math/big"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

const (
	// Epsilon is the threshold for what is `smol enough`
	epsilon float64 = 0.0001

	// d2precision is how much precision we want from big.Float
	d2precision uint = 64 // was chosen arbitrarily
)

// BigFloat is the implementation of Vector using big.Float.
type BigFloat struct {
	x, y *big.Float
}

// NewBigFloat creates a new BigFloat and returns a pointer to it.
func NewBigFloat(x, y float64) d2interface.Vector {
	xbf, ybf := big.NewFloat(x), big.NewFloat(y)
	xbf.SetPrec(d2precision)
	ybf.SetPrec(d2precision)
	result := &BigFloat{xbf, ybf}

	return result
}

// XYBigFloat returns the values as big.Float.
func (f *BigFloat) XYBigFloat() (x, y *big.Float) {
	return f.x, f.y
}

// XYFloat64 returns the values as float64.
func (f *BigFloat) XYFloat64() (x, y *float64) {
	xf, _ := f.x.Float64()
	yf, _ := f.y.Float64()

	return &xf, &yf
}

// Equals check whether this Vector is equal to a given Vector.
func (f *BigFloat) Equals(v d2interface.Vector) bool {
	vx, vy := v.XYBigFloat()
	return f.x.Cmp(vx) == 0 && f.y.Cmp(vy) == 0
}

// EqualsApprox checks if the Vector is approximately equal
// to the given Vector.
func (f *BigFloat) EqualsApprox(v d2interface.Vector) bool {
	x, y := f.CompareApprox(v)
	return x == 0 && y == 0
}

// CompareApprox performs a fuzzy comparison and returns 2
// ints represending the -1 to 1 comparison of x and y.
func (f *BigFloat) CompareApprox(v d2interface.Vector) (x, y int) {
	vx, vy := v.XYFloat64()
	fx, fy := f.XYFloat64()

	return d2math.CompareFloat64Fuzzy(fx, vx),
		d2math.CompareFloat64Fuzzy(fy, vy)
}

// Set sets the vector values to the given float64 values.
func (f *BigFloat) Set(x, y float64) d2interface.Vector {
	f.x.SetFloat64(x)
	f.y.SetFloat64(y)

	return f
}

// Clone creates a copy of this Vector.
func (f *BigFloat) Clone() d2interface.Vector {
	result := NewBigFloat(0, 0)
	x, y := result.XYBigFloat()
	x.Copy(f.x)
	y.Copy(f.y)

	return result
}

// Floor rounds the vector down to the nearest whole numbers.
func (f *BigFloat) Floor() d2interface.Vector {
	var xi, yi big.Int

	f.x.Int(&xi)
	f.y.Int(&yi)
	f.x.SetInt(&xi)
	f.y.SetInt(&yi)

	return f
}

// Add to this Vector the components of the given Vector.
func (f *BigFloat) Add(v d2interface.Vector) d2interface.Vector {
	vx, vy := v.XYBigFloat()
	f.x.Add(f.x, vx)
	f.y.Add(f.y, vy)

	return f
}

// Subtract from this Vector from the components of the given Vector.
func (f *BigFloat) Subtract(v d2interface.Vector) d2interface.Vector {
	vx, vy := v.XYBigFloat()
	f.x.Sub(f.x, vx)
	f.y.Sub(f.y, vy)

	return f
}

// Multiply this Vector by the components of the given Vector.
func (f *BigFloat) Multiply(v d2interface.Vector) d2interface.Vector {
	vx, vy := v.XYBigFloat()
	f.x.Mul(f.x, vx)
	f.y.Mul(f.y, vy)

	return f
}

// Scale multiplies this vector by a single value.
func (f *BigFloat) Scale(s float64) d2interface.Vector {
	sb := big.NewFloat(s)
	f.x.Mul(f.x, sb)
	f.y.Mul(f.y, sb)

	return f
}

// Divide divides this vector by the components of the given vector.
func (f *BigFloat) Divide(v d2interface.Vector) d2interface.Vector {
	vx, vy := v.XYBigFloat()
	f.x.Quo(f.x, vx)
	f.y.Quo(f.y, vy)

	return f
}

// Abs sets the vector to it's absolute (positive) equivalent.
func (f *BigFloat) Abs() d2interface.Vector {
	xm, ym := 1.0, 1.0
	if f.x.Sign() == -1 {
		xm = -1
	}

	if f.y.Sign() == -1 {
		ym = -1
	}

	m := NewBigFloat(xm, ym)
	f.Multiply(m)

	return f
}

// Negate multiplies the vector by -1.
func (f *BigFloat) Negate() d2interface.Vector {
	return f.Scale(-1)
}

// Distance is the distance between this Vector and the given Vector.
func (f *BigFloat) Distance(v d2interface.Vector) float64 {
	delta := v.Clone().Subtract(f)
	return delta.Length()
}

// Length is the length/magnitude/quantity of this Vector.
func (f *BigFloat) Length() float64 {
	sqx, sqy := f.Clone().Multiply(f).XYBigFloat()
	sum := big.NewFloat(0).Add(sqx, sqy)
	r, _ := sum.Sqrt(sum).Float64()

	return r
}

// TODO: Improve this based on float64 standard library Stringer implementation.
func (f *BigFloat) String() string {
	return fmt.Sprintf("BigFloat{%s, %s}", f.x.Text('f', 5), f.y.Text('f', 5))
}

/*


// Marshal converts the Vector into a slice of bytes
func (v *BigFloat) Marshal() ([]byte, error) {
	// TODO not sure how to do this properly
	return nil, nil
}

// Unmarshal converts a slice of bytes to x/y *big.Float
// and assigns them to itself
func (v *BigFloat) Unmarshal(buf []byte) error {
	// TODO not sure how to do this properly
	return nil
}

// SetFromEntity copies the vector of a world entity
// func (v *BigFloat) SetFromEntity(entity d2interface.WorldEntity) d2interface.Vector {
// 	return v.Copy(entity.Position())
// }

// Set the x,y members of the Vector
func (v *BigFloat) Set(x, y *big.Float) d2interface.Vector {
	v.x = x
	v.y = y

	return v
}

// SetToPolar sets the `x` and `y` values of this object
// from a given polar coordinate.
func (v *BigFloat) SetToPolar(azimuth, radius *big.Float) d2interface.Vector {
	// HACK we should do this better, with the big.Float
	a, _ := azimuth.Float64()
	r, _ := radius.Float64()
	v.x.SetFloat64(math.Cos(a) * r)
	v.y.SetFloat64(math.Sin(a) * r)

	return v
}

// Abs returns a clone that is positive
func (v *BigFloat) Abs() d2interface.Vector {
	clone := v.Clone()
	neg1 := big.NewFloat(-1.0)

	if clone.XBig().Sign() == -1 { // is negative1
		clone.XBig().Mul(clone.XBig(), neg1)
	}

	if v.YBig().Sign() == -1 { // is negative1
		clone.YBig().Mul(clone.YBig(), neg1)
	}

	return clone
}

// Angle computes the angle in radians with respect
// to the positive x-axis
func (v *BigFloat) Angle() *big.Float {
	// HACK we should find a way to do this purely
	// with big.Float
	floatX, _ := v.XBig().Float64()
	floatY, _ := v.YBig().Float64()
	floatAngle := math.Atan2(floatY, floatX)

	if floatAngle < 0 {
		floatAngle += 2.0 * math.Pi
	}

	return big.NewFloat(floatAngle)
}

// SetAngle sets the angle of this Vector
func (v *BigFloat) SetAngle(angle *big.Float) d2interface.Vector {
	return v.SetToPolar(angle, v.Length())
}



// Scale this Vector by the given value
func (v *BigFloat) Scale(s *big.Float) d2interface.Vector {
	v.x.Sub(v.x, s)
	v.y.Sub(v.y, s)

	return v
}


// Negate thex and y components of this Vector
func (v *BigFloat) Negate() d2interface.Vector {
	return v.Scale(big.NewFloat(negative1))
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

// Dot returns the dot product of this Vector and the given Vector.
func (v *BigFloat) Dot(src d2interface.Vector) *big.Float {
	c := v.Clone()
	c.XBig().Mul(c.XBig(), src.XBig())
	c.YBig().Mul(c.YBig(), src.YBig())

	return c.XBig().Add(c.XBig(), c.YBig())
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
