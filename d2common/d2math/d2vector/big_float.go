// Package d2vector is an Implementation of 2-dimensional vectors with big.Float components
package d2vector

import (
	"fmt"
	"math"
	"math/big"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	// Epsilon is the threshold for what is `smol enough`
	epsilon float64 = 0.0001

	// d2precision is how much precision we want from big.Float
	d2precision uint = 64 // was chosen arbitrarily

	// for convenience in negating sign
	negative1 float64 = -1.0

	// for convenience
	zero float64 = 0.0
)

// BigFloat is the implementation of Vector using float64
// to store x and y.
type BigFloat struct {
	x *big.Float
	y *big.Float
}

// NewBigFloat creates a new BigFloat and returns a pointer to it.
func NewBigFloat(x, y float64) *BigFloat {
	xbf, ybf := big.NewFloat(x), big.NewFloat(y)
	xbf.SetPrec(d2precision)
	ybf.SetPrec(d2precision)
	result := &BigFloat{xbf, ybf}

	return result
}

// XBig returns the big.Float value of x.
func (v *BigFloat) XBig() *big.Float {
	return v.x
}

// YBig returns the big.Float value of y.
func (v *BigFloat) YBig() *big.Float {
	return v.y
}

// X64 returns the float64 value of x.
func (v *BigFloat) X64() (float64, big.Accuracy) {
	return v.x.Float64()
}

// Y64 returns the float64 value of y.
func (v *BigFloat) Y64() (float64, big.Accuracy) {
	return v.y.Float64()
}

// CopyFloat64 copies the values from a Float64 to f.
func (v *BigFloat) CopyFloat64(sf Float64) {
	v.x.Set(sf.XBig())
	v.x.Set(sf.YBig())
}

// AsFloat64 returns a pointer to a float64 base on
// the values of v.
func (v *BigFloat) AsFloat64() *Float64 {
	sf := new(Float64)
	sf.x, _ = v.X64()
	sf.y, _ = v.Y64()
	return sf
}

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

// Clone creates a copy of this Vector
func (v *BigFloat) Clone() d2interface.Vector {
	result := NewBigFloat(0, 0)
	result.Copy(v)

	return result
}

// Copy copies the src x/y members to this Vector x/y members
func (v *BigFloat) Copy(src d2interface.Vector) d2interface.Vector {
	v.x.Copy(src.XBig())
	v.y.Copy(src.YBig())

	return v
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

// Equals check whether this Vector is equal to a given Vector.
func (v *BigFloat) Equals(src d2interface.Vector) bool {
	return v.x.Cmp(src.XBig()) == 0 && v.y.Cmp(src.YBig()) == 0
}

// FuzzyEquals checks if the Vector is approximately equal
// to the given Vector. epsilon is what we consider `smol enough`
func (v *BigFloat) FuzzyEquals(src d2interface.Vector) bool {
	smol := big.NewFloat(epsilon)
	d := v.Distance(src)
	d.Abs(d)

	return d.Cmp(smol) < 1 || d.Cmp(smol) < 1
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

// Add to this Vector the components of the given Vector
func (v *BigFloat) Add(src d2interface.Vector) d2interface.Vector {
	v.x.Add(v.x, src.XBig())
	v.y.Add(v.y, src.YBig())

	return v
}

// Subtract from this Vector the components of the given Vector
func (v *BigFloat) Subtract(src d2interface.Vector) d2interface.Vector {
	v.x.Sub(v.x, src.XBig())
	v.y.Sub(v.y, src.YBig())

	return v
}

// Multiply this Vector with the components of the given Vector
func (v *BigFloat) Multiply(src d2interface.Vector) d2interface.Vector {
	v.x.Mul(v.x, src.XBig())
	v.y.Mul(v.y, src.YBig())

	return v
}

// Scale this Vector by the given value
func (v *BigFloat) Scale(s *big.Float) d2interface.Vector {
	v.x.Sub(v.x, s)
	v.y.Sub(v.y, s)

	return v
}

// Divide this Vector by the given Vector
func (v *BigFloat) Divide(src d2interface.Vector) d2interface.Vector {
	v.x.Quo(v.x, src.XBig())
	v.y.Quo(v.y, src.YBig())

	return v
}

// Negate thex and y components of this Vector
func (v *BigFloat) Negate() d2interface.Vector {
	return v.Scale(big.NewFloat(negative1))
}

// Distance calculate the distance between this Vector and the given Vector
func (v *BigFloat) Distance(src d2interface.Vector) *big.Float {
	dist := v.DistanceSq(src)

	return dist.Sqrt(dist)
}

// DistanceSq calculate the distance suared between this Vector and the given
// Vector
func (v *BigFloat) DistanceSq(src d2interface.Vector) *big.Float {
	delta := src.Clone().Subtract(v)
	deltaSq := delta.Multiply(delta)

	return big.NewFloat(zero).Add(deltaSq.XBig(), deltaSq.YBig())
}

// Length returns the length of this Vector
func (v *BigFloat) Length() *big.Float {
	xsq, ysq := v.LengthSq()

	return xsq.Add(xsq, ysq)
}

// LengthSq returns the x and y values squared
func (v *BigFloat) LengthSq() (*big.Float, *big.Float) {
	clone := v.Clone()
	x, y := clone.XBig(), clone.YBig()

	return x.Mul(x, x), y.Mul(y, y)
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

// Floor rounds the vector down to the nearest whole numbers.
func (v *BigFloat) Floor() d2interface.Vector {
	var xi, yi big.Int

	v.x.Int(&xi)
	v.y.Int(&yi)
	v.XBig().SetInt(&xi)
	v.YBig().SetInt(&yi)

	return v
}

func (v *BigFloat) String() string {
	return fmt.Sprintf("BigFloat{%s, %s}", v.x.Text('v', 5), v.y.Text('v', 5))
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
}
