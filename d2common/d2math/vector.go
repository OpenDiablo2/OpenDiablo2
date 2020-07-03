package d2math

import (
	"math"
	"math/big"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// I want to put these somewhere convenient...
// ZERO *Vector = &Vector{}
// ONE *Vector = &Vector{1, 1}
// RIGHT *Vector = &Vector{1, 0}
// LEFT *Vector = &Vector{-1, 0}
// UP *Vector = &Vector{0, -1}
// DOWN *Vector = &Vector{0, 1}

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

// NewVector2 creates a new Vector
func NewVector2(x, y float64) *Vector2 {
	xbf, ybf := big.NewFloat(x), big.NewFloat(y)
	xbf.SetPrec(d2precision)
	ybf.SetPrec(d2precision)
	result := &Vector2{xbf, ybf}

	return result
}

// Vector is an Implementation of 2-dimensional vectors with
// big.Float components
type Vector2 struct {
	x *big.Float
	y *big.Float
}

// X returns the x member of the Vector
func (v *Vector2) X() *big.Float {
	return v.x
}

// Y returns the y member of the Vector
func (v *Vector2) Y() *big.Float {
	return v.y
}

// Marshal converts the Vector into a slice of bytes
func (v *Vector2) Marshal() ([]byte, error) {
	// TODO not sure how to do this properly
	return nil, nil
}

// Unmarshal converts a slice of bytes to x/y *big.Float
// and assigns them to itself
func (v *Vector2) Unmarshal(buf []byte) error {
	// TODO not sure how to do this properly
	return nil
}

// Clone creates a copy of this Vector
func (v *Vector2) Clone() d2interface.Vector {
	result := &Vector2{}
	result.Copy(v)

	return result
}

// Copy copies the src x/y members to this Vector x/y members
func (v *Vector2) Copy(src d2interface.Vector) d2interface.Vector {
	v.x.Copy(src.X())
	v.y.Copy(src.Y())

	return v
}

// SetFromEntity copies the vector of a world entity
// func (v *Vector2) SetFromEntity(entity d2interface.WorldEntity) d2interface.Vector {
// 	return v.Copy(entity.Position())
// }

// Set the x,y members of the Vector
func (v *Vector2) Set(x, y *big.Float) d2interface.Vector {
	v.x = x
	v.y = y

	return v
}

// SetToPolar sets the `x` and `y` values of this object
// from a given polar coordinate.
func (v *Vector2) SetToPolar(azimuth, radius *big.Float) d2interface.Vector {
	// HACK we should do this better, with the big.Float
	a, _ := azimuth.Float64()
	r, _ := radius.Float64()
	v.x.SetFloat64(math.Cos(a) * r)
	v.y.SetFloat64(math.Sin(a) * r)

	return v
}

// Equals check whether this Vector is equal to a given Vector.
func (v *Vector2) Equals(src d2interface.Vector) bool {
	return v.x.Cmp(src.X()) == 0 && v.y.Cmp(src.Y()) == 0
}

// FuzzyEquals checks if the Vector is approximately equal
// to the given Vector. epsilon is what we consider `smol enough`
func (v *Vector2) FuzzyEquals(src d2interface.Vector) bool {
	smol := big.NewFloat(epsilon)
	d := v.Distance(src)
	d.Abs(d)

	return d.Cmp(smol) < 1 || d.Cmp(smol) < 1
}

// Abs returns a clone that is positive
func (v *Vector2) Abs() d2interface.Vector {
	clone := v.Clone()
	neg1 := big.NewFloat(-1.0)

	if clone.X().Sign() == -1 { // is negative1
		clone.X().Mul(clone.X(), neg1)
	}

	if v.Y().Sign() == -1 { // is negative1
		clone.Y().Mul(clone.Y(), neg1)
	}

	return clone
}

// Angle computes the angle in radians with respect
// to the positive x-axis
func (v *Vector2) Angle() *big.Float {
	// HACK we should find a way to do this purely
	// with big.Float
	floatX, _ := v.X().Float64()
	floatY, _ := v.Y().Float64()
	floatAngle := math.Atan2(floatY, floatX)

	if floatAngle < 0 {
		floatAngle += 2.0 * math.Pi
	}

	return big.NewFloat(floatAngle)
}

// SetAngle sets the angle of this Vector
func (v *Vector2) SetAngle(angle *big.Float) d2interface.Vector {
	return v.SetToPolar(angle, v.Length())
}

// Add to this Vector the components of the given Vector
func (v *Vector2) Add(src d2interface.Vector) d2interface.Vector {
	v.x.Add(v.x, src.X())
	v.y.Add(v.y, src.Y())

	return v
}

// Subtract from this Vector the components of the given Vector
func (v *Vector2) Subtract(src d2interface.Vector) d2interface.Vector {
	v.x.Sub(v.x, src.X())
	v.y.Sub(v.y, src.Y())

	return v
}

// Multiply this Vector with the components of the given Vector
func (v *Vector2) Multiply(src d2interface.Vector) d2interface.Vector {
	v.x.Mul(v.x, src.X())
	v.y.Mul(v.y, src.Y())

	return v
}

// Scale this Vector by the given value
func (v *Vector2) Scale(s *big.Float) d2interface.Vector {
	v.x.Sub(v.x, s)
	v.y.Sub(v.y, s)

	return v
}

// Divide this Vector by the given Vector
func (v *Vector2) Divide(src d2interface.Vector) d2interface.Vector {
	v.x.Quo(v.x, src.X())
	v.y.Quo(v.y, src.Y())

	return v
}

// Negate thex and y components of this Vector
func (v *Vector2) Negate() d2interface.Vector {
	return v.Scale(big.NewFloat(negative1))
}

// Distance calculate the distance between this Vector and the given Vector
func (v *Vector2) Distance(src d2interface.Vector) *big.Float {
	dist := v.DistanceSq(src)

	return dist.Sqrt(dist)
}

// DistanceSq calculate the distance suared between this Vector and the given
// Vector
func (v *Vector2) DistanceSq(src d2interface.Vector) *big.Float {
	delta := src.Clone().Subtract(v)
	deltaSq := delta.Multiply(delta)

	return big.NewFloat(zero).Add(deltaSq.X(), deltaSq.Y())
}

// Length returns the length of this Vector
func (v *Vector2) Length() *big.Float {
	xsq, ysq := v.LengthSq()

	return xsq.Add(xsq, ysq)
}

func (v *Vector2) LengthSq() (*big.Float, *big.Float) {
	clone := v.Clone()
	x, y := clone.X(), clone.Y()

	return x.Mul(x, x), y.Mul(y, y)
}

// SetLength sets the length of this Vector
func (v *Vector2) SetLength(length *big.Float) d2interface.Vector {
	return v.Normalize().Scale(length)
}

// Normalize Makes the vector a unit length vector (magnitude of 1) in the same
// direction.
func (v *Vector2) Normalize() d2interface.Vector {
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
//in the positive direction.
func (v *Vector2) NormalizeRightHand() d2interface.Vector {
	x := v.x
	v.x = v.y.Mul(v.y, big.NewFloat(negative1))
	v.y = x

	return v
}

// NormalizeLeftHand rotate this Vector to its perpendicular,
//in the negative1 direction.
func (v *Vector2) NormalizeLeftHand() d2interface.Vector {
	x := v.x
	v.x = v.y
	v.y = x.Mul(x, big.NewFloat(negative1))

	return v
}

// Calculate the dot product of this Vector and the given Vector
func (v *Vector2) Dot(src d2interface.Vector) *big.Float {
	c := v.Clone()
	c.X().Mul(c.X(), src.X())
	c.Y().Mul(c.Y(), src.Y())

	return c.X().Add(c.X(), c.Y())
}

// Cross Calculate the cross product of this Vector and the given Vector.
func (v *Vector2) Cross(src d2interface.Vector) *big.Float {
	c := v.Clone()
	c.X().Mul(c.X(), src.X())
	c.Y().Mul(c.Y(), src.Y())

	return c.X().Sub(c.X(), c.Y())
}

// Lerp Linearly interpolate between this Vector and the given Vector.
func (v *Vector2) Lerp(
	src d2interface.Vector,
	t *big.Float,
) d2interface.Vector {
	vc, sc := v.Clone(), src.Clone()
	x, y := vc.X(), vc.Y()
	v.x.Set(x.Add(x, t.Mul(t, sc.X().Sub(sc.X(), x))))
	v.y.Set(y.Add(y, t.Mul(t, sc.Y().Sub(sc.Y(), y))))

	return v
}

// Reset this Vector the zero vector (0, 0).
func (v *Vector2) Reset() d2interface.Vector {

	v.x.SetFloat64(zero)
	v.y.SetFloat64(zero)

	return v
}

// Limit the length (or magnitude) of this Vector
func (v *Vector2) Limit(max *big.Float) d2interface.Vector {
	length := v.Length()

	if max.Cmp(length) < 0 {
		v.Scale(length.Quo(max, length))
	}

	return v
}

// Reflect this Vector off a line defined by a normal.
func (v *Vector2) Reflect(normal d2interface.Vector) d2interface.Vector {
	clone := v.Clone()
	clone.Normalize()

	two := big.NewFloat(2.0) // there's some matrix algebra magic here
	dot := v.Clone().Dot(normal)
	normal.Scale(two.Mul(two, dot))

	return v.Subtract(normal)
}

// Mirror reflect this Vector across another.
func (v *Vector2) Mirror(axis d2interface.Vector) d2interface.Vector {
	return v.Reflect(axis).Negate()
}

// Rotate this Vector by an angle amount.
func (v *Vector2) Rotate(angle *big.Float) d2interface.Vector {
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
