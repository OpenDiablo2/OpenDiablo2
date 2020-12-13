package d2math

import "math"

type Vector4Like interface {
	Vector2Like
	Vector3Like
	XYZW() (x, y, z, w float64)
}

// static check that Vector4 is Vector4Like
var _ Vector4Like = &Vector4{}

// NewVector4 creates a new Vector4
func NewVector4(x, y, z, w float64) *Vector4 {
	return &Vector4{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// Vector4 is a representation of a vector in 4D space.
type Vector4 struct {
	X, Y, Z, W float64
}

// XYZ returns the x and y components of the vector
func (v *Vector4) XY() (float64, float64) {
	return v.X, v.Y
}

// XYZ returns the x, y, and z components of the vector
func (v *Vector4) XYZ() (x, y, z float64) {
	return v.X, v.Y, v.Z
}

// XYZW returns the x, y, z, and w components of the vector
func (v *Vector4) XYZW() (x, y, z, w float64) {
	return v.X, v.Y, v.Z, v.W
}

// Clone makes a clone of this vector
func (v *Vector4) Clone() *Vector4 {
	return NewVector4(v.XYZW())
}

// Copy the components of a given Vector into this Vector.
func (v *Vector4) Copy(other *Vector4) *Vector4 {
	return v.Set(other.XYZW())
}

// Equals checks if this vector is equal to the given vector
func (v *Vector4) Equals(other *Vector4) bool {
	return math.Abs(v.X-other.X) < Epsilon &&
		math.Abs(v.Y-other.Y) < Epsilon &&
		math.Abs(v.Z-other.Z) < Epsilon &&
		math.Abs(v.W-other.W) < Epsilon
}

// Set the x, y, z, and w components of this vector
func (v *Vector4) Set(x, y, z, w float64) *Vector4 {
	v.X, v.Y, v.Z, v.W = x, y, z, w

	return v
}

// Add the given vector to this vector
func (v *Vector4) Add(other *Vector4) *Vector4 {
	return v.Set(v.X+other.X, v.Y+other.Y, v.Z+other.Z, v.W+other.W)
}

// Subtract the given vector from this vector
func (v *Vector4) Subtract(other *Vector4) *Vector4 {
	return v.Set(v.X-other.X, v.Y-other.Y, v.Z-other.Z, v.W-other.W)
}

// Scale this vector by a scalar value
func (v *Vector4) Scale(s float64) *Vector4 {
	return v.Set(v.X*s, v.Y*s, v.Z*s, v.W*s)
}

// Length returns the length (magnitude) of this vector
func (v *Vector4) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared returns the length of this vector, squared
func (v *Vector4) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W
}

// Normalize this vector to unit length of 1 in the same direction
func (v *Vector4) Normalize() *Vector4 {
	l := v.LengthSquared()

	if l > 0 {
		v.Scale(1 / math.Sqrt(l))
	}

	return v
}

// Dot calculates the dot product with the given vector
func (v *Vector4) Dot(other *Vector4) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z + v.W*other.W
}

// Lerp linearly interpolates with the given vector using position `t`
func (v *Vector4) Lerp(other *Vector4, t float64) *Vector4 {
	return v.Set(
		v.X+t*(other.X-v.X),
		v.Y+t*(other.Y-v.Y),
		v.Z+t*(other.Z-v.Z),
		v.W+t*(other.W-v.W),
	)
}

// Multiply this vector by the given vector
func (v *Vector4) Multiply(other *Vector4) *Vector4 {
	return v.Set(
		v.X*other.X,
		v.Y*other.Y,
		v.Z*other.Z,
		v.W*other.W,
	)
}

// Divide this vector by the given vector
func (v *Vector4) Divide(other *Vector4) *Vector4 {
	return v.Set(
		v.X/other.X,
		v.Y/other.Y,
		v.Z/other.Z,
		v.W/other.W,
	)
}

// Distance calculates the distance between this vector and the given vector
func (v *Vector4) Distance(other *Vector4) float64 {
	return math.Sqrt(v.DistanceSquared(other))
}

// DistanceSquared calculates the distance between this vector and the given vector, squared
func (v *Vector4) DistanceSquared(other *Vector4) float64 {
	dx := other.X - v.X
	dy := other.Y - v.Y
	dz := other.Z - v.Z
	dw := other.W - v.W

	return dx*dx + dy*dy + dz*dz + dw*dw
}

// Negate the signs of the components of this vector
func (v *Vector4) Negate() *Vector4 {
	return v.Scale(-1)
}

// TransformMatrix4 transforms this vector with the given Matrix4
func (v *Vector4) TransformMatrix4(m4 *Matrix4) *Vector4 {
	m := m4.Values

	return v.Set(
		m[0]*v.X+m[4]*v.Y+m[8]*v.Z+m[12]*v.W,
		m[1]*v.X+m[5]*v.Y+m[9]*v.Z+m[13]*v.W,
		m[2]*v.X+m[6]*v.Y+m[10]*v.Z+m[14]*v.W,
		m[3]*v.X+m[7]*v.Y+m[11]*v.Z+m[15]*v.W,
	)
}

// TransformQuaternion transforms this vector with the given Quaternion
func (v *Vector4) TransformQuaternion(q *Quaternion) *Vector4 {
	ix := q.W*v.X + q.Y*v.Z - q.Z*v.Y
	iy := q.W*v.Y + q.Z*v.X - q.X*v.Z
	iz := q.W*v.Z + q.X*v.Y - q.Y*v.X
	iw := -q.X*v.X - q.Y*v.Y - q.Z*v.Z

	v.X = ix*q.W + iw*-q.X + iy*-q.Z - iz*-q.Y
	v.Y = iy*q.W + iw*-q.Y + iz*-q.X - ix*-q.Z
	v.Z = iz*q.W + iw*-q.Z + ix*-q.Y - iy*-q.X

	return v
}

// Reset all components of this vector to 0's
func (v *Vector4) Reset() *Vector4 {
	return v.Set(0, 0, 0, 0)
}
