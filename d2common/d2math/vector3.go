package d2math

import "math"

// Vector3Like is something that has an XYZ method that returns x, y, and z coordinate values
type Vector3Like interface {
	Vector2Like
	XYZ() (x, y, z float64)
}

// static check that Vector3 is Vector3Like
var _ Vector3Like = &Vector3{}

// NewVector3 creates a new Vector3
func NewVector3(x, y, z float64) *Vector3 {
	return &Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

// NewVector3Zero creates a zero Vector3 for use in comparisons and stuff
func NewVector3Zero() *Vector3 {
	return NewVector3(0, 0, 0)
}

// NewVector3One creates a one Vector3 for use in comparisons and stuff
func NewVector3One() *Vector3 {
	return NewVector3(1, 1, 1)
}

// NewVector3Right creates a right Vector3 for use in comparisons and stuff
func NewVector3Right() *Vector3 {
	return NewVector3(1, 0, 0)
}

// NewVector3Left creates a left Vector3 for use in comparisons and stuff
func NewVector3Left() *Vector3 {
	return NewVector3(-1, 0, 0)
}

// NewVector3Up creates a up Vector3 for use in comparisons and stuff
func NewVector3Up() *Vector3 {
	return NewVector3(0, -1, 0)
}

// NewVector3Down creates a down Vector3 for use in comparisons and stuff
func NewVector3Down() *Vector3 {
	return NewVector3(0, 1, 0)
}

// NewVector3Forward creates a forward Vector3 for use in comparisons and stuff
func NewVector3Forward() *Vector3 {
	return NewVector3(0, 0, 1)
}

// NewVector3Back creates a back Vector3 for use in comparisons and stuff
func NewVector3Back() *Vector3 {
	return NewVector3(0, 0, -1)
}

// Vector3 is a representation of a vector in 3D space.
type Vector3 struct {
	X, Y, Z float64
}

// XY returns the x and y components of the vector
func (v *Vector3) XY() (x, y float64) {
	return v.X, v.Y
}

// XYZ returns the x, y, and z components of the vector
func (v *Vector3) XYZ() (x, y, z float64) {
	return v.X, v.Y, v.Z
}

// Up sets this Vector to point up.
func (v *Vector3) Up() *Vector3 {
	return v.Set(0, 1, 0)
}

// Min sets the components of this Vector to be the `math.Min` result from the given vector.
func (v *Vector3) Min(other *Vector3) *Vector3 {
	if other == nil {
		return v
	}

	v.X = math.Min(v.X, other.X)
	v.Y = math.Min(v.Y, other.Y)
	v.Z = math.Min(v.Z, other.Z)

	return v
}

// Max sets the components of this Vector to be the `math.Max` result from the given vector.
func (v *Vector3) Max(other *Vector3) *Vector3 {
	if other == nil {
		return v
	}

	v.X = math.Max(v.X, other.X)
	v.Y = math.Max(v.Y, other.Y)
	v.Z = math.Max(v.Z, other.Z)

	return v
}

// Clone makes a clone of this vector
func (v *Vector3) Clone() *Vector3 {
	return NewVector3(v.XYZ())
}

// AddVectors adds the two given Vector3s and sets the results into this Vector3.
func (v *Vector3) AddVectors(a, b *Vector3) *Vector3 {
	if a == nil {
		a = NewVector3(0, 0, 0)
	}

	if b == nil {
		b = NewVector3(0, 0, 0)
	}

	return v.Set(a.X+b.X, a.Y+b.Y, a.Z+b.Z)
}

// CrossVectors calulcates the cross of the two given Vector3s and sets the
// results into this Vector3.
func (v *Vector3) CrossVectors(a, b *Vector3) *Vector3 {
	if a == nil {
		a = NewVector3(0, 0, 0)
	}

	if b == nil {
		b = NewVector3(0, 0, 0)
	}

	return v.Set(a.Y*b.Z-a.Z*b.Y, a.Z*b.X-a.X*b.Z, a.X*b.Y-a.Y*b.X)
}

// Equals checks if this vector is equal to another vector
func (v *Vector3) Equals(other *Vector3) bool {
	return math.Abs(v.X-other.X) < Epsilon &&
		math.Abs(v.Y-other.Y) < Epsilon &&
		math.Abs(v.Z-other.Z) < Epsilon
}

// Copy copies the values from the given vector to this vector
func (v *Vector3) Copy(other *Vector3) *Vector3 {
	return v.Set(other.X, other.Y, other.Z)
}

// Set the x, y, and z components for this vector
func (v *Vector3) Set(x, y, z float64) *Vector3 {
	v.X, v.Y, v.Z = x, y, z
	return v
}

// SetFromMatrix4 sets the components of this Vector3 from the position of the given Matrix4.
func (v *Vector3) SetFromMatrix4(m *Matrix4) *Vector3 {
	const m4startIdx = 12

	return v.SetFromSlice(m.Values[:], m4startIdx)
}

// SetFromMatrix4Column sets the components of this Vector3 from the column of the given Matrix4.
func (v *Vector3) SetFromMatrix4Column(m *Matrix4, column int) *Vector3 {
	const m4order = 4
	column = int(Clamp(float64(column), 0, m4order-1))
	return v.SetFromSlice(m.Values[:], column*m4order)
}

// SetFromSlice sets the components of this Vector3 from the given array, based on the offset.
func (v *Vector3) SetFromSlice(s []float64, offset int) *Vector3 {
	offset = int(Clamp(float64(offset), 0, float64(len(s))))
	return v.Set(s[offset], s[offset+1], s[offset+2])
}

// Add a given Vector to this Vector. Addition is component-wise.
func (v *Vector3) Add(other *Vector3) *Vector3 {
	return v.Set(v.X+other.X, v.Y+other.Y, v.Z+other.Z)
}

// Subtract a given Vector from this Vector. Addition is component-wise.
func (v *Vector3) Subtract(other *Vector3) *Vector3 {
	return v.Set(v.X-other.X, v.Y-other.Y, v.Z-other.Z)
}

// AddScalar adds the given value to each component of this Vector.
func (v *Vector3) AddScalar(s float64) *Vector3 {
	return v.Set(v.X+s, v.Y+s, v.Z+s)
}

// AddAndScale adds and scales a given Vector and scale to this Vector. Addition is component-wise.
func (v *Vector3) AddAndScale(other *Vector3, s float64) *Vector3 {
	return v.Set(s*other.X, s*other.Y, s*other.Z)
}

// Multiply performs a component-wise multiplication between this Vector and the given Vector.
func (v *Vector3) Multiply(other *Vector3) *Vector3 {
	return v.Set(v.X*other.X, v.Y*other.Y, v.Z*other.Z)
}

// Divide performs a component-wise division between this Vector and the given Vector.
func (v *Vector3) Divide(other *Vector3) *Vector3 {
	return v.Set(v.X/other.X, v.Y/other.Y, v.Z/other.Z)
}

// Negate the x, y, and z components of this vector
func (v *Vector3) Negate() *Vector3 {
	return v.Scale(-1)
}

// Scale this Vector by the given value.
func (v *Vector3) Scale(s float64) *Vector3 {
	return v.Set(v.X*s, v.Y*s, v.Z*s)
}

// DistanceSquared calculates the distance between this Vector and the given Vector, squared.
func (v *Vector3) DistanceSquared(other *Vector3) float64 {
	dx, dy, dz := other.X-v.X, other.Y-v.Y, other.Z-v.Z
	return dx*dx + dy*dy + dz*dz
}

// Distance calculates the distance between this Vector and the given Vector.
func (v *Vector3) Distance(other *Vector3) float64 {
	return math.Sqrt(v.DistanceSquared(other))
}

// Length calculates the length (or magnitude) of this Vector, squared.
func (v *Vector3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Length calculates the length (or magnitude) of this Vector.
func (v *Vector3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// Normalize this Vector.
// Makes the vector a unit length vector (magnitude of 1) in the same direction.
func (v *Vector3) Normalize() *Vector3 {
	l := v.LengthSquared()

	if l > 0 {
		v.Scale(1 / math.Sqrt(l))
	}

	return v
}

// Dot calculates the dot product of this Vector and the given Vector.
func (v *Vector3) Dot(other *Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross calculates the cross (vector) product of this Vector ( which will be modified) and the
// given Vector.
func (v *Vector3) Cross(other *Vector3) *Vector3 {
	return v.Set(
		v.Y*other.Z-v.Z*other.Y,
		v.Z*other.X-v.X*other.Z,
		v.X*other.Y-v.Y*other.X,
	)
}

// Lerp Linearly interpolates between this Vector and the given Vector.
// Interpolates this Vector towards the given Vector.
func (v *Vector3) Lerp(other *Vector3, t float64) *Vector3 {
	t = Clamp(t, 0, 1)

	return v.Set(
		v.X+t*(other.X-v.X),
		v.Y+t*(other.Y-v.Y),
		v.Z+t*(other.Z-v.Z),
	)
}

// ApplyMatrix3 takes a Matrix3 and applies it to this Vector3.
func (v *Vector3) ApplyMatrix3(m3 *Matrix3) *Vector3 {
	m := m3.Values

	return v.Set(
		m[0]*v.X+m[3]*v.Y+m[6]*v.Z,
		m[1]*v.X+m[4]*v.Y+m[7]*v.Z,
		m[2]*v.X+m[5]*v.Y+m[8]*v.Z,
	)
}

// ApplyMatrix4 takes a Matrix4 and applies it to this Vector3.
func (v *Vector3) ApplyMatrix4(m4 *Matrix4) *Vector3 {
	m := m4.Values
	w := 1 / (m[3]*v.X + m[7]*v.Y + m[11]*v.Z + m[15])

	return v.Set(
		(m[0]*v.X+m[4]*v.Y+m[8]*v.Z+m[12])*w,
		(m[1]*v.X+m[5]*v.Y+m[9]*v.Z+m[13])*w,
		(m[2]*v.X+m[6]*v.Y+m[10]*v.Z+m[14])*w,
	)
}

// TransformMatrix3 transform this Vector with the given Matrix3.
func (v *Vector3) TransformMatrix3(m3 *Matrix3) *Vector3 {
	m := m3.Values

	return v.Set(
		v.X*m[0]+v.Y*m[3]+v.Z*m[6],
		v.X*m[1]+v.Y*m[4]+v.Z*m[7],
		v.X*m[2]+v.Y*m[5]+v.Z*m[8],
	)
}

// TransformMatrix4 transform this Vector with the given Matrix4.
func (v *Vector3) TransformMatrix4(m4 *Matrix4) *Vector3 {
	m := m4.Values

	return v.Set(
		m[0]*v.X+m[4]*v.Y+m[8]*v.Z+m[12],
		m[1]*v.X+m[5]*v.Y+m[9]*v.Z+m[13],
		m[2]*v.X+m[6]*v.Y+m[10]*v.Z+m[14],
	)
}

// TransformCoordinates transforms the coordinates of this Vector3 with the given Matrix4.
func (v *Vector3) TransformCoordinates(m4 *Matrix4) *Vector3 {
	m := m4.Values

	tx := (v.X * m[0]) + (v.Y * m[4]) + (v.Z * m[8]) + m[12]
	ty := (v.X * m[1]) + (v.Y * m[5]) + (v.Z * m[9]) + m[13]
	tz := (v.X * m[2]) + (v.Y * m[6]) + (v.Z * m[10]) + m[14]
	tw := (v.X * m[3]) + (v.Y * m[7]) + (v.Z * m[11]) + m[15]

	return v.Set(tx/tw, ty/tw, tz/tw)
}

// TransformQuaternion transform this Vector with the given Quaternion.
func (v *Vector3) TransformQuaternion(q *Quaternion) *Vector3 {
	// calculate quat * vec
	ix := q.W*v.X + q.Y*v.Z - q.Z*v.Y
	iy := q.W*v.Y + q.Z*v.X - q.X*v.Z
	iz := q.W*v.Z + q.X*v.Y - q.Y*v.X
	iw := -q.X*v.X - q.Y*v.Y - q.Z*v.Z

	// calculate result * inverse quat
	return v.Set(
		ix*q.W+iw*-q.X+iy*-q.Z-iz*-q.Y,
		iy*q.W+iw*-q.Y+iz*-q.X-ix*-q.Z,
		iz*q.W+iw*-q.Z+ix*-q.Y-iy*-q.X,
	)
}

// Project multiplies this Vector3 by the specified matrix, applying a W divide.
// This is useful for projection, e.g. un-projecting a 2D point into 3D space.
func (v *Vector3) Project(m4 *Matrix4) *Vector3 {
	m := m4.Values

	a00, a01, a02, a03,
		a10, a11, a12, a13,
		a20, a21, a22, a23,
		a30, a31, a32, a33 :=
		m[0], m[1], m[2], m[3],
		m[4], m[5], m[6], m[7],
		m[8], m[9], m[10], m[11],
		m[12], m[13], m[14], m[15]

	lw := 1 / (v.X*a03 + v.Y*a13 + v.Z*a23 + a33)

	// calculate result * inverse quat
	return v.Set(
		(v.X*a00+v.Y*a10+v.Z*a20+a30)*lw,
		(v.X*a01+v.Y*a11+v.Z*a21+a31)*lw,
		(v.X*a02+v.Y*a12+v.Z*a22+a32)*lw,
	)
}

// ProjectViewMatrix multiplies this Vector3 by the given view and projection matrices.
func (v *Vector3) ProjectViewMatrix(view, projection *Matrix4) *Vector3 {
	return v.ApplyMatrix4(view).ApplyMatrix4(projection)
}

// UnprojectViewMatrix multiplies this Vector3 by the given inversed projection matrix and world
// matrix.
func (v *Vector3) UnprojectViewMatrix(projection, world *Matrix4) *Vector3 {
	return v.ApplyMatrix4(projection).ApplyMatrix4(world)
}

// Unproject this point from 2D space to 3D space.
// The point should have its x and y properties set to
// 2D screen space, and the z either at 0 (near plane)
// or 1 (far plane). The provided matrix is assumed to already
// be combined, i.e. projection * view * model.
// After this operation, this vector's (x, y, z) components will
// represent the unprojected 3D coordinate.
func (v *Vector3) Unproject(viewport *Vector4, invProjectionView *Matrix4) *Vector3 {
	viewX := viewport.X
	viewY := viewport.Y
	viewWidth := viewport.Z
	viewHeight := viewport.W

	x := v.X - viewX
	y := (viewHeight - v.Y - 1) - viewY
	z := v.Z

	v.X = (2*x)/viewWidth - 1
	v.Y = (2*y)/viewHeight - 1
	v.Z = 2*z - 1

	return v.Project(invProjectionView)
}

// Reset this vectors components to 0
func (v *Vector3) Reset() *Vector3 {
	return v.Set(0, 0, 0)
}
