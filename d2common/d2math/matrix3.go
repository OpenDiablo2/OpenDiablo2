package d2math

import "math"

const numMatrix3Values = 3 * 3

// NewMatrix3 creates a new three-dimensional matrix. Argument is optional Matrix3 to copy from.
func NewMatrix3(from *Matrix3) *Matrix3 {
	m := &Matrix3{Values: [numMatrix3Values]float64{}}

	if from != nil {
		return m.Copy(from)
	}

	return m.Identity()
}

// Matrix3 is a three-dimensional matrix
type Matrix3 struct {
	Values [numMatrix3Values]float64
}

// Clone makes a clone of this Matrix3.
func (m *Matrix3) Clone() *Matrix3 {
	return NewMatrix3(m)
}

// Copy the values of a given Matrix into this Matrix.
func (m *Matrix3) Copy(other *Matrix3) *Matrix3 {
	m.Values[0] = other.Values[0]
	m.Values[1] = other.Values[1]
	m.Values[2] = other.Values[2]
	m.Values[3] = other.Values[3]
	m.Values[4] = other.Values[4]
	m.Values[5] = other.Values[5]
	m.Values[6] = other.Values[6]
	m.Values[7] = other.Values[7]
	m.Values[8] = other.Values[8]

	return m
}

// Set is an alias for Matrix3.Copy
func (m *Matrix3) Set(other *Matrix3) *Matrix3 {
	return m.Copy(other)
}

// Identity resets this Matrix to an identity (default) matrix.
func (m *Matrix3) Identity() *Matrix3 {
	m.Values[0] = 1
	m.Values[1] = 0
	m.Values[2] = 0
	m.Values[3] = 0
	m.Values[4] = 1
	m.Values[5] = 0
	m.Values[6] = 0
	m.Values[7] = 0
	m.Values[8] = 1

	return m
}

// FromMatrix4 copies the values of a given Matrix4 into this Matrix3.
func (m *Matrix3) FromMatrix4(m4 *Matrix4) *Matrix3 {
	m.Values[0] = m4.Values[0]
	m.Values[1] = m4.Values[1]
	m.Values[2] = m4.Values[2]
	m.Values[3] = m4.Values[4]
	m.Values[4] = m4.Values[5]
	m.Values[5] = m4.Values[6]
	m.Values[6] = m4.Values[8]
	m.Values[7] = m4.Values[9]
	m.Values[8] = m4.Values[10]

	return m
}

// FromSlice sets the values of this Matrix from the given slice.
func (m *Matrix3) FromSlice(s []float64) *Matrix3 {
	if s == nil {
		return m
	}

	numVals := len(s)
	for idx := 0; idx < numVals && idx < numMatrix3Values; idx++ {
		m.Values[idx] = s[idx]
	}

	return m
}

// Transpose this Matrix.
func (m *Matrix3) Transpose() *Matrix3 {
	v := m.Values

	v[1], v[2], v[3], v[5], v[6], v[7] = v[3], v[6], v[1], v[7], v[2], v[5]

	return m
}

// Invert this Matrix.
func (m *Matrix3) Invert() *Matrix3 {
	a := m.Values

	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]
	a20 := a[6]
	a21 := a[7]
	a22 := a[8]

	b01 := a22*a11 - a12*a21
	b11 := -a22*a10 + a12*a20
	b21 := a21*a10 - a11*a20

	// calculate the determinant
	det := a00*b01 + a01*b11 + a02*b21

	if det < Epsilon {
		return nil
	}

	det = 1 / det

	a[0] = b01 * det
	a[1] = (-a22*a01 + a02*a21) * det
	a[2] = (a12*a01 - a02*a11) * det
	a[3] = b11 * det
	a[4] = (a22*a00 - a02*a20) * det
	a[5] = (-a12*a00 + a02*a10) * det
	a[6] = b21 * det
	a[7] = (-a21*a00 + a01*a20) * det
	a[8] = (a11*a00 - a01*a10) * det

	return m
}

// Adjoint calculates the adjoint, or adjugate, of this Matrix.
func (m *Matrix3) Adjoint() *Matrix3 {
	a := m.Values

	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]
	a20 := a[6]
	a21 := a[7]
	a22 := a[8]

	a[0] = a11*a22 - a12*a21
	a[1] = a02*a21 - a01*a22
	a[2] = a01*a12 - a02*a11
	a[3] = a12*a20 - a10*a22
	a[4] = a00*a22 - a02*a20
	a[5] = a02*a10 - a00*a12
	a[6] = a10*a21 - a11*a20
	a[7] = a01*a20 - a00*a21
	a[8] = a00*a11 - a01*a10

	return m
}

// Determinant calculates the determinant of this Matrix.
func (m *Matrix3) Determinant() float64 {
	a := m.Values

	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]
	a20 := a[6]
	a21 := a[7]
	a22 := a[8]

	return a00*(a22*a11-a12*a21) + a01*(-a22*a10+a12*a20) + a02*(a21*a10-a11*a20)
}

// Multiply this Matrix by the given Matrix.
func (m *Matrix3) Multiply(other *Matrix3) *Matrix3 {
	a, b := m.Values, other.Values

	a00, b00 := a[0], b[0]
	a01, b01 := a[1], b[1]
	a02, b02 := a[2], b[2]
	a10, b10 := a[3], b[3]
	a11, b11 := a[4], b[4]
	a12, b12 := a[5], b[5]
	a20, b20 := a[6], b[6]
	a21, b21 := a[7], b[7]
	a22, b22 := a[8], b[8]

	a[0] = b00*a00 + b01*a10 + b02*a20
	a[1] = b00*a01 + b01*a11 + b02*a21
	a[2] = b00*a02 + b01*a12 + b02*a22
	a[3] = b10*a00 + b11*a10 + b12*a20
	a[4] = b10*a01 + b11*a11 + b12*a21
	a[5] = b10*a02 + b11*a12 + b12*a22
	a[6] = b20*a00 + b21*a10 + b22*a20
	a[7] = b20*a01 + b21*a11 + b22*a21
	a[8] = b20*a02 + b21*a12 + b22*a22

	return m
}

// Translate this Matrix using the given Vector.
func (m *Matrix3) Translate(v Vector2Like) *Matrix3 {
	a := m.Values
	x, y := v.XY()

	a[6] = x*a[0] + y*a[3] + a[6]
	a[7] = x*a[1] + y*a[4] + a[7]
	a[8] = x*a[2] + y*a[5] + a[8]

	return m
}

// Rotate applies a rotation transformation to this Matrix.
func (m *Matrix3) Rotate(radians float64) *Matrix3 {
	a := m.Values

	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a10 := a[3]
	a11 := a[4]
	a12 := a[5]

	s, c := math.Sin(radians), math.Cos(radians)

	a[0] = c*a00 + s*a10
	a[1] = c*a01 + s*a11
	a[2] = c*a02 + s*a12
	a[3] = c*a10 - s*a00
	a[4] = c*a11 - s*a01
	a[5] = c*a12 - s*a02

	return m
}

// Scale applies a scalar transformation to this Matrix.
func (m *Matrix3) Scale(v Vector2Like) *Matrix3 {
	a := m.Values
	x, y := v.XY()

	a[0] = x * a[0]
	a[1] = x * a[1]
	a[2] = x * a[2]

	a[3] = y * a[3]
	a[4] = y * a[4]
	a[5] = y * a[5]

	return m
}

// FromQuaternion sets the values of this Matrix from the given Quaternion.
func (m *Matrix3) FromQuaternion(q *Quaternion) *Matrix3 {
	x, y, z, w := q.X, q.Y, q.Z, q.W

	x2, y2, z2 := x+x, y+y, z+z
	xx, xy, xz := x*x2, x*y2, x*z2
	yy, yz, zz := y*y2, y*z2, z*z2
	wx, wy, wz := w*x2, w*y2, w*z2

	out := m.Values

	out[0] = 1 - (yy + zz)
	out[3] = xy + wz
	out[6] = xz - wy
	out[1] = xy - wz
	out[4] = 1 - (xx + zz)
	out[7] = yz + wx
	out[2] = xz + wy
	out[5] = yz - wx
	out[8] = 1 - (xx + yy)

	return m
}

// NormalFromMatrix4 sets the values of this Matrix3 to be normalized from the given Matrix4.
func (m *Matrix3) NormalFromMatrix4(m4 *Matrix4) *Matrix3 {
	a, out := m4.Values, m.Values

	a00 := a[0]
	a01 := a[1]
	a02 := a[2]
	a03 := a[3]
	a10 := a[4]
	a11 := a[5]
	a12 := a[6]
	a13 := a[7]
	a20 := a[8]
	a21 := a[9]
	a22 := a[10]
	a23 := a[11]
	a30 := a[12]
	a31 := a[13]
	a32 := a[14]
	a33 := a[15]

	b00 := a00*a11 - a01*a10
	b01 := a00*a12 - a02*a10
	b02 := a00*a13 - a03*a10
	b03 := a01*a12 - a02*a11
	b04 := a01*a13 - a03*a11
	b05 := a02*a13 - a03*a12
	b06 := a20*a31 - a21*a30
	b07 := a20*a32 - a22*a30
	b08 := a20*a33 - a23*a30
	b09 := a21*a32 - a22*a31
	b10 := a21*a33 - a23*a31
	b11 := a22*a33 - a23*a32

	// calculate the determinant
	det := b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06

	if det < Epsilon {
		return nil
	}

	det = 1 / det

	out[0] = (a11*b11 - a12*b10 + a13*b09) * det
	out[1] = (a12*b08 - a10*b11 - a13*b07) * det
	out[2] = (a10*b10 - a11*b08 + a13*b06) * det
	out[3] = (a02*b10 - a01*b11 - a03*b09) * det
	out[4] = (a00*b11 - a02*b08 + a03*b07) * det
	out[5] = (a01*b08 - a00*b10 - a03*b06) * det
	out[6] = (a31*b05 - a32*b04 + a33*b03) * det
	out[7] = (a32*b02 - a30*b05 - a33*b01) * det
	out[8] = (a30*b04 - a31*b02 + a33*b00) * det

	return m
}
