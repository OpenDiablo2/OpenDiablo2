package d2math

import "math"

func qNoop(_ *Quaternion) { /* no operation, the default OnChangeCallback */ }

func NewQuaternion(x, y, z, w float64) *Quaternion {
	return &Quaternion{
		X:                x,
		Y:                y,
		Z:                z,
		W:                w,
		OnChangeCallback: qNoop,
	}
}

// Quaternion is a quaternion. :)
type Quaternion struct {
	X, Y, Z, W float64

	// This callback is invoked, if set, each time a value in this quaternion is changed.
	// The callback is passed one argument, a reference to this quaternion.
	OnChangeCallback func(*Quaternion)
}

// XY returns the x and y components of the quaternion
func (q *Quaternion) XY() (float64, float64) {
	return q.X, q.Y
}

// XYZ returns the x, y and z components of the quaternion
func (q *Quaternion) XYZ() (x, y, z float64) {
	return q.X, q.Y, q.Z
}

// XYZW returns the x, y, z, and w components of the quaternion
func (q *Quaternion) XYZW() (x, y, z, w float64) {
	return q.X, q.Y, q.Z, q.W
}

// SetX sets the x component and calls the OnChangeCallback function
func (q *Quaternion) SetX(v float64) *Quaternion {
	return q.Set(v, q.Y, q.Z, q.W)
}

// SetY sets the y component and calls the OnChangeCallback function
func (q *Quaternion) SetY(v float64) *Quaternion {
	return q.Set(q.X, v, q.Z, q.W)
}

// SetZ sets the z component and calls the OnChangeCallback function
func (q *Quaternion) SetZ(v float64) *Quaternion {
	return q.Set(q.X, q.Y, v, q.W)
}

// SetW sets the w component and calls the OnChangeCallback function
func (q *Quaternion) SetW(v float64) *Quaternion {
	return q.Set(q.X, q.Y, q.Z, v)
}

// Set the x, y, z, and w components of this quaternion and call the OnChangeCallback function
func (q *Quaternion) Set(x, y, z, w float64) *Quaternion {
	q.X, q.Y, q.Z, q.W = x, y, z, w

	q.OnChangeCallback(q)

	return q
}

// Clone creates a clone of this quaternion
func (q *Quaternion) Clone() *Quaternion {
	return NewQuaternion(q.XYZW())
}

// Copy the values of the given quaternion
func (q *Quaternion) Copy(other *Quaternion) *Quaternion {
	return q.Set(other.XYZW())
}

// Add (sum) the given quaternion components to this quaternion
func (q *Quaternion) Add(other *Quaternion) *Quaternion {
	return q.Set(
		q.X+other.X,
		q.Y+other.Y,
		q.Z+other.Z,
		q.W+other.W,
	)
}

// Subtract the given quaternion components from this quaternion
func (q *Quaternion) Subtract(other *Quaternion) *Quaternion {
	return q.Set(
		q.X-other.X,
		q.Y-other.Y,
		q.Z-other.Z,
		q.W-other.W,
	)
}

// Scale this quaternion's component values by a scalar
func (q *Quaternion) Scale(s float64) *Quaternion {
	return q.Set(
		q.X*s,
		q.Y*s,
		q.Z*s,
		q.W*s,
	)
}

// Length returns the length, or magnitude, of this quaternion
func (q *Quaternion) Length() float64 {
	return math.Sqrt(q.LengthSquared())
}

// LengthSquared returns the length, or magnitude, of this quaternion, squared
func (q *Quaternion) LengthSquared() float64 {
	return q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W
}

// Normalize this quaternion to length of 1
func (q *Quaternion) Normalize() *Quaternion {
	l := q.LengthSquared()

	if l > 0 {
		q.Scale(1 / math.Sqrt(l))
	}

	return q
}

// Dot calculates the dot product with the given quaternion
func (q *Quaternion) Dot(other *Quaternion) float64 {
	return q.X*other.X + q.Y*other.Y + q.Z*other.Z + q.W*other.W
}

// Lerp linearly interpolates to the given quaternion
func (q *Quaternion) Lerp(other *Quaternion, t float64) *Quaternion {
	return q.Set(
		q.X+t*(other.X-q.X),
		q.Y+t*(other.Y-q.Y),
		q.Z+t*(other.Z-q.Z),
		q.W+t*(other.W-q.W),
	)
}

// RotationTo rotates this Quaternion based on the two given vectors.
func (q *Quaternion) RotationTo(a, b *Vector3) *Quaternion {
	dot := a.Dot(b)

	if dot < (-1 + Epsilon) {
		tmpVec, xunit, yunit := NewVector3(0, 0, 0), NewVector3Right(), NewVector3Down()

		if tmpVec.Copy(xunit).Cross(a).Length() < Epsilon {
			tmpVec.Copy(yunit).Cross(a)
		}

		tmpVec.Normalize()

		return q.SetAxisAngle(tmpVec, PI)
	} else if dot > (1 - Epsilon) {
		return q.Identity()
	} else {
		tmpVec := NewVector3(0, 0, 0).Copy(a).Cross(b)

		q.Set(
			tmpVec.X,
			tmpVec.Y,
			tmpVec.Z,
			1+dot,
		)

		return q.Normalize()
	}
}

// SetAxes sets the axes of this Quaternion.
func (q *Quaternion) SetAxes(view, right, up *Vector3) *Quaternion {
	tmpMat3 := NewMatrix3(nil)

	m := tmpMat3.Values

	m[0] = right.X
	m[3] = right.Y
	m[6] = right.Z

	m[1] = up.X
	m[4] = up.Y
	m[7] = up.Z

	m[2] = -view.X
	m[5] = -view.Y
	m[8] = -view.Z

	return q.FromMatrix3(tmpMat3).Normalize()
}

// Identity returns the identity quaternion
func (q *Quaternion) Identity() *Quaternion {
	return q.Set(0, 0, 0, 1)
}

// SetAxisAngle sets the axis angle of this Quaternion.
func (q *Quaternion) SetAxisAngle(axis *Vector3, radians float64) *Quaternion {
	radians = radians / 2
	s := math.Sin(radians)

	return q.Set(
		s*axis.X,
		s*axis.Y,
		s*axis.Z,
		math.Cos(radians),
	)
}

// Multiply this quaternion by the given quaternion
func (q *Quaternion) Multiply(other *Quaternion) *Quaternion {
	return q.Set(
		q.X*other.W+q.W*other.X+q.Y*other.Z-q.Z*other.Y,
		q.Y*other.W+q.W*other.Y+q.Z*other.X-q.X*other.Z,
		q.Z*other.W+q.W*other.Z+q.X*other.Y-q.Y*other.X,
		q.W*other.W-q.X*other.X-q.Y*other.Y-q.Z*other.Z,
	)
}

// Slerp smoothly linaerly interpolate this Quaternion towards the given Quaternion or Vector.
func (q *Quaternion) Slerp(other Vector4Like, t float64) *Quaternion {
	ax, ay, az, aw := q.XYZW()
	bx, by, bz, bw := other.XYZW()

	// calc cosine
	cosom := ax*bx + ay*by + az*bz + aw*bw

	// adjust signs (if necessary)
	if cosom < 0 {
		cosom = -cosom
		bx = -bx
		by = -by
		bz = -bz
		bw = -bw
	}

	// "from" and "to" quaternions are very close ... so we can do a linear interpolation
	scale0 := 1 - t
	scale1 := t

	// calculate coefficients
	if (1 - cosom) > Epsilon {
		// standard case (slerp)
		var omega = math.Acos(cosom)
		var sinom = math.Sin(omega)

		scale0 = math.Sin((1.0-t)*omega) / sinom
		scale1 = math.Sin(t*omega) / sinom
	}

	return q.Set(
		scale0*ax+scale1*bx,
		scale0*ay+scale1*by,
		scale0*az+scale1*bz,
		scale0*aw+scale1*bw,
	)
}

// Invert this quaternion
func (q *Quaternion) Invert() *Quaternion {
	dot := q.Clone().Dot(q)

	if dot != 0 {
		dot = 1 / dot
	}

	return q.Conjugate().Scale(dot)
}

// Conjugate converts this Quaternion into its conjugate.
func (q *Quaternion) Conjugate() *Quaternion {
	return q.Set(-q.X, -q.Y, -q.Z, q.W)
}

// RotateX rotates this quaternion on the x-axis.
func (q *Quaternion) RotateX(radians float64) *Quaternion {
	radians /= 2
	bx, bw := math.Sin(radians), math.Cos(radians)

	return q.Set(
		q.X*bw+q.W*bx,
		q.Y*bw+q.Z*bx,
		q.Z*bw-q.Y*bx,
		q.W*bw-q.X*bx,
	)
}

// RotateY rotates this quaternion on the y-axis.
func (q *Quaternion) RotateY(radians float64) *Quaternion {
	radians /= 2
	by, bw := math.Sin(radians), math.Cos(radians)

	return q.Set(
		q.X*bw-q.Z*by,
		q.Y*bw+q.W*by,
		q.Z*bw+q.X*by,
		q.W*bw-q.Y*by,
	)
}

// RotateZ rotates this quaternion on the z-axis.
func (q *Quaternion) RotateZ(radians float64) *Quaternion {
	radians /= 2
	bz, bw := math.Sin(radians), math.Cos(radians)

	return q.Set(
		q.X*bw+q.Y*bz,
		q.Y*bw-q.X*bz,
		q.Z*bw+q.W*bz,
		q.W*bw-q.Z*bz,
	)
}

// CalculateW creates a unit (or rotation) Quaternion from its x, y, and z components.
func (q *Quaternion) CalculateW() *Quaternion {
	x, y, z := q.XYZ()
	return q.SetW(-math.Sqrt(1.0 - x*x - y*y - z*z))
}

// SetFromEuler sets this Quaternion from the given Euler, based on Euler order.
func (q *Quaternion) SetFromEuler(e *Euler) *Quaternion {
	x, y, z := e.X/2, e.Y/2, e.Z/2
	c1, c2, c3 := math.Cos(x), math.Cos(y), math.Cos(z)
	s1, s2, s3 := math.Sin(x), math.Sin(y), math.Sin(z)

	switch e.Order {
	case EulerOrderYXZ:
		q.Set(
			s1*c2*c3+c1*s2*s3,
			c1*s2*c3-s1*c2*s3,
			c1*c2*s3-s1*s2*c3,
			c1*c2*c3+s1*s2*s3,
		)
	case EulerOrderZXY:
		q.Set(
			s1*c2*c3-c1*s2*s3,
			c1*s2*c3+s1*c2*s3,
			c1*c2*s3+s1*s2*c3,
			c1*c2*c3-s1*s2*s3,
		)
	case EulerOrderZYX:
		q.Set(
			s1*c2*c3-c1*s2*s3,
			c1*s2*c3+s1*c2*s3,
			c1*c2*s3-s1*s2*c3,
			c1*c2*c3+s1*s2*s3,
		)
	case EulerOrderYZX:
		q.Set(
			s1*c2*c3+c1*s2*s3,
			c1*s2*c3+s1*c2*s3,
			c1*c2*s3-s1*s2*c3,
			c1*c2*c3-s1*s2*s3,
		)
	case EulerOrderXZY:
		q.Set(
			s1*c2*c3-c1*s2*s3,
			c1*s2*c3-s1*c2*s3,
			c1*c2*s3+s1*s2*c3,
			c1*c2*c3+s1*s2*s3,
		)
	case EulerOrderXYZ:
		fallthrough
	default:
		q.Set(
			s1*c2*c3+c1*s2*s3,
			c1*s2*c3-s1*c2*s3,
			c1*c2*s3+s1*s2*c3,
			c1*c2*c3-s1*s2*s3,
		)
	}

	return q
}

// SetFromRotationMatrix sets the rotation of this Quaternion from the given Matrix4.
func (q *Quaternion) SetFromRotationMatrix(m4 *Matrix4) *Quaternion {
	m11 := m4.Values[0]
	m12 := m4.Values[4]
	m13 := m4.Values[8]
	m21 := m4.Values[1]
	m22 := m4.Values[5]
	m23 := m4.Values[9]
	m31 := m4.Values[2]
	m32 := m4.Values[6]
	m33 := m4.Values[10]

	trace := m11 + m22 + m33
	var s float64

	if trace > 0 {
		s = 0.5 / math.Sqrt(trace+1.0)

		return q.Set(
			(m32-m23)*s,
			(m13-m31)*s,
			(m21-m12)*s,
			0.25/s,
		)
	} else if m11 > m22 && m11 > m33 {
		s = 2.0 * math.Sqrt(1.0+m11-m22-m33)

		return q.Set(
			0.25*s,
			(m12+m21)/s,
			(m13+m31)/s,
			(m32-m23)/s,
		)
	} else if m22 > m33 {
		s = 2.0 * math.Sqrt(1.0+m22-m11-m33)

		return q.Set(
			(m12+m21)/s,
			0.25*s,
			(m23+m32)/s,
			(m13-m31)/s,
		)
	}

	s = 2.0 * math.Sqrt(1.0+m33-m11-m22)

	return q.Set(
		(m13+m31)/s,
		(m23+m32)/s,
		0.25*s,
		(m21-m12)/s,
	)

}

// FromMatrix3 converts the given Matrix into this Quaternion.
func (q *Quaternion) FromMatrix3(m3 *Matrix3) *Quaternion {
	m := m3.Values
	fTrace := m[0] + m[4] + m[8]
	var fRoot float64

	siNext, tmp := []int{1, 2, 0}, []float64{0, 0, 0}

	if fTrace > 0 {
		// |w| > 1/2, may as well choose w > 1/2
		fRoot = math.Sqrt(fTrace + 1.0) // 2w
		q.W = 0.5 * fRoot
		fRoot = 0.5 / fRoot // 1/(4w)
		q.X = (m[7] - m[5]) * fRoot
		q.Y = (m[2] - m[6]) * fRoot
		q.Z = (m[3] - m[1]) * fRoot
	} else {
		// |w| <= 1/2
		var i = 0

		if m[4] > m[0] {
			i = 1
		}

		if m[8] > m[i*3+i] {
			i = 2
		}

		var j = siNext[i]
		var k = siNext[j]

		//  This isn't quite as clean without array access
		fRoot = math.Sqrt(m[i*3+i] - m[j*3+j] - m[k*3+k] + 1)

		tmp[i] = 0.5 * fRoot
		fRoot = 0.5 / fRoot

		tmp[j] = (m[j*3+i] + m[i*3+j]) * fRoot
		tmp[k] = (m[k*3+i] + m[i*3+k]) * fRoot

		q.X = tmp[0]
		q.Y = tmp[1]
		q.Z = tmp[2]
		q.W = (m[k*3+j] - m[j*3+k]) * fRoot
	}

	q.OnChangeCallback(q)

	return q
}
