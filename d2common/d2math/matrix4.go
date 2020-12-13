package d2math

import "math"

const numMatrix4Values = 4 * 4

// NewMatrix4 creates a new four-dimensional matrix. Argument is optional Matrix4 to copy from.
func NewMatrix4(from *Matrix4) *Matrix4 {
	m := &Matrix4{Values: [numMatrix4Values]float64{}}

	if from != nil {
		return m.Copy(from)
	}

	return m.Identity()
}

// Matrix4 is a four-dimensional matrix
type Matrix4 struct {
	Values [numMatrix4Values]float64
}

// Decompose this matrix into vec3 for translation, rotation, and scale
func (m *Matrix4) Decompose() (t, r, s *Vector3) {
	t = NewVector3(m.Values[12], m.Values[13], m.Values[14])

	s = NewVector3(
		NewVector3(m.Values[0], m.Values[1], m.Values[2]).Length(),
		NewVector3(m.Values[4], m.Values[5], m.Values[6]).Length(),
		NewVector3(m.Values[8], m.Values[9], m.Values[10]).Length(),
	)

	e := NewEuler(0, 0, 0, EulerOrderDefault)

	rotmat := m.Clone()
	rotmat.Values[0] /= s.X
	rotmat.Values[1] /= s.Y
	rotmat.Values[2] /= s.Z

	rotmat.Values[4] /= s.X
	rotmat.Values[5] /= s.Y
	rotmat.Values[6] /= s.Z

	rotmat.Values[8] /= s.X
	rotmat.Values[9] /= s.Y
	rotmat.Values[10] /= s.Z

	e.SetFromRotationMatrix(rotmat, EulerOrderXYZ)

	r = NewVector3(e.XYZ())

	return t, r, s
}

// Clone makes a clone of this Matrix4.
func (m *Matrix4) Clone() *Matrix4 {
	return NewMatrix4(m)
}

// Copy the values of a given Matrix into this Matrix.
func (m *Matrix4) Copy(other *Matrix4) *Matrix4 {
	if other == nil {
		return m.Zero()
	}

	a := other.Values
	return m.SetValues(
		a[0], a[1], a[2], a[3],
		a[4], a[5], a[6], a[7],
		a[8], a[9], a[10], a[11],
		a[12], a[13], a[14], a[15],
	)
}

// Set is an alias for Matrix4.Copy
func (m *Matrix4) Set(other *Matrix4) *Matrix4 {
	return m.Copy(other)
}

// SetValues sets the values of this Matrix4.
func (m *Matrix4) SetValues(a, b, c, d, e, f, g, h, i, j, k, l, mm, n, o, p float64) *Matrix4 {
	m.Values[0], m.Values[1], m.Values[2], m.Values[3] = a, b, c, d
	m.Values[4], m.Values[5], m.Values[6], m.Values[7] = e, f, g, h
	m.Values[8], m.Values[9], m.Values[10], m.Values[11] = i, j, k, l
	m.Values[12], m.Values[13], m.Values[14], m.Values[15] = mm, n, o, p

	return m
}

// Identity resets this Matrix to an identity (default) matrix.
func (m *Matrix4) Identity() *Matrix4 {
	return m.SetValues(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

// FromSlice sets the values of this Matrix from the given slice.
func (m *Matrix4) FromSlice(s []float64) *Matrix4 {
	if s != nil {
		numVals := len(s)
		for idx := 0; idx < numVals && idx < numMatrix4Values; idx++ {
			m.Values[idx] = s[idx]
		}
	}

	return m
}

// Zero resets this matrix, setting all values to 0
func (m *Matrix4) Zero() *Matrix4 {
	return m.SetValues(
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	)
}

// Transform generates a transform matrix based on the given position, scale and rotation.
func (m *Matrix4) Transform(position, scale Vector3Like, rotation *Quaternion) *Matrix4 {
	rotMat := NewMatrix4(nil).FromQuaternion(rotation)
	rm := rotMat.Values

	sx, sy, sz := scale.XYZ()
	px, py, pz := position.XYZ()

	return m.SetValues(
		rm[0]*sx, rm[1]*sx, rm[2]*sx, 0,
		rm[4]*sy, rm[5]*sy, rm[6]*sy, 0,
		rm[8]*sz, rm[9]*sz, rm[10]*sz, 0,
		px, py, pz, 1,
	)
}

// SetXYZ sets the x, y, z values of this matrix
func (m *Matrix4) SetXYZ(x, y, z float64) *Matrix4 {
	m.Identity()

	m.Values[12] = x
	m.Values[13] = y
	m.Values[14] = z

	return m
}

// SetScaling sets the scaling values of this Matrix.
func (m *Matrix4) SetScaling(x, y, z float64) *Matrix4 {
	m.Zero()

	m.Values[0] = x
	m.Values[5] = y
	m.Values[10] = z
	m.Values[15] = 1

	return m
}

// Transpose this Matrix.
func (m *Matrix4) Transpose() *Matrix4 {
	a01 := m.Values[1]
	a02 := m.Values[2]
	a03 := m.Values[3]
	a12 := m.Values[6]
	a13 := m.Values[7]
	a23 := m.Values[11]

	m.Values[1] = m.Values[4]
	m.Values[2] = m.Values[8]
	m.Values[3] = m.Values[12]
	m.Values[4] = a01
	m.Values[6] = m.Values[9]
	m.Values[7] = m.Values[13]
	m.Values[8] = a02
	m.Values[9] = a12
	m.Values[11] = m.Values[14]
	m.Values[12] = a03
	m.Values[13] = a13
	m.Values[14] = a23

	return m
}

// GetInverse copies the given Matrix4 into this Matrix and then inverses it.
func (m *Matrix4) GetInverse(other *Matrix4) *Matrix4 {
	return m.Copy(other).Invert()
}

// Invert this Matrix.
func (m *Matrix4) Invert() *Matrix4 {
	a00, a01, a02, a03 := m.Values[0], m.Values[1], m.Values[2], m.Values[3]
	a10, a11, a12, a13 := m.Values[4], m.Values[5], m.Values[6], m.Values[7]
	a20, a21, a22, a23 := m.Values[8], m.Values[9], m.Values[10], m.Values[11]
	a30, a31, a32, a33 := m.Values[12], m.Values[13], m.Values[14], m.Values[15]

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
		return m
	}

	det = 1 / det

	return m.SetValues(
		(a11*b11-a12*b10+a13*b09)*det,
		(a02*b10-a01*b11-a03*b09)*det,
		(a31*b05-a32*b04+a33*b03)*det,
		(a22*b04-a21*b05-a23*b03)*det,
		(a12*b08-a10*b11-a13*b07)*det,
		(a00*b11-a02*b08+a03*b07)*det,
		(a32*b02-a30*b05-a33*b01)*det,
		(a20*b05-a22*b02+a23*b01)*det,
		(a10*b10-a11*b08+a13*b06)*det,
		(a01*b08-a00*b10-a03*b06)*det,
		(a30*b04-a31*b02+a33*b00)*det,
		(a21*b02-a20*b04-a23*b00)*det,
		(a11*b07-a10*b09-a12*b06)*det,
		(a00*b09-a01*b07+a02*b06)*det,
		(a31*b01-a30*b03-a32*b00)*det,
		(a20*b03-a21*b01+a22*b00)*det,
	)
}

// Adjoint calculates the adjoint, or adjugate, of this Matrix.
func (m *Matrix4) Adjoint() *Matrix4 {
	a00, a01, a02, a03 := m.Values[0], m.Values[1], m.Values[2], m.Values[3]
	a10, a11, a12, a13 := m.Values[4], m.Values[5], m.Values[6], m.Values[7]
	a20, a21, a22, a23 := m.Values[8], m.Values[9], m.Values[10], m.Values[11]
	a30, a31, a32, a33 := m.Values[12], m.Values[13], m.Values[14], m.Values[15]

	return m.SetValues(
		a11*(a22*a33-a23*a32)-a21*(a12*a33-a13*a32)+a31*(a12*a23-a13*a22),
		-(a01*(a22*a33-a23*a32) - a21*(a02*a33-a03*a32) + a31*(a02*a23-a03*a22)),
		a01*(a12*a33-a13*a32)-a11*(a02*a33-a03*a32)+a31*(a02*a13-a03*a12),
		-(a01*(a12*a23-a13*a22) - a11*(a02*a23-a03*a22) + a21*(a02*a13-a03*a12)),
		-(a10*(a22*a33-a23*a32) - a20*(a12*a33-a13*a32) + a30*(a12*a23-a13*a22)),
		a00*(a22*a33-a23*a32)-a20*(a02*a33-a03*a32)+a30*(a02*a23-a03*a22),
		-(a00*(a12*a33-a13*a32) - a10*(a02*a33-a03*a32) + a30*(a02*a13-a03*a12)),
		a00*(a12*a23-a13*a22)-a10*(a02*a23-a03*a22)+a20*(a02*a13-a03*a12),
		a10*(a21*a33-a23*a31)-a20*(a11*a33-a13*a31)+a30*(a11*a23-a13*a21),
		-(a00*(a21*a33-a23*a31) - a20*(a01*a33-a03*a31) + a30*(a01*a23-a03*a21)),
		a00*(a11*a33-a13*a31)-a10*(a01*a33-a03*a31)+a30*(a01*a13-a03*a11),
		-(a00*(a11*a23-a13*a21) - a10*(a01*a23-a03*a21) + a20*(a01*a13-a03*a11)),
		-(a10*(a21*a32-a22*a31) - a20*(a11*a32-a12*a31) + a30*(a11*a22-a12*a21)),
		a00*(a21*a32-a22*a31)-a20*(a01*a32-a02*a31)+a30*(a01*a22-a02*a21),
		-(a00*(a11*a32-a12*a31) - a10*(a01*a32-a02*a31) + a30*(a01*a12-a02*a11)),
		a00*(a11*a22-a12*a21)-a10*(a01*a22-a02*a21)+a20*(a01*a12-a02*a11),
	)
}

// Determinant calculates the determinant of this Matrix.
func (m *Matrix4) Determinant() float64 {
	a00, a01, a02, a03 := m.Values[0], m.Values[1], m.Values[2], m.Values[3]
	a10, a11, a12, a13 := m.Values[4], m.Values[5], m.Values[6], m.Values[7]
	a20, a21, a22, a23 := m.Values[8], m.Values[9], m.Values[10], m.Values[11]
	a30, a31, a32, a33 := m.Values[12], m.Values[13], m.Values[14], m.Values[15]

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

	return b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06
}

// Multiply this Matrix4 by the given Matrix4.
func (m *Matrix4) Multiply(other *Matrix4) *Matrix4 {
	a00, a01, a02, a03 := m.Values[0], m.Values[1], m.Values[2], m.Values[3]
	a10, a11, a12, a13 := m.Values[4], m.Values[5], m.Values[6], m.Values[7]
	a20, a21, a22, a23 := m.Values[8], m.Values[9], m.Values[10], m.Values[11]
	a30, a31, a32, a33 := m.Values[12], m.Values[13], m.Values[14], m.Values[15]

	// Cache only the current line of the second matrix
	b0, b1, b2, b3 := other.Values[0], other.Values[1], other.Values[2], other.Values[3]
	m.Values[0] = b0*a00 + b1*a10 + b2*a20 + b3*a30
	m.Values[1] = b0*a01 + b1*a11 + b2*a21 + b3*a31
	m.Values[2] = b0*a02 + b1*a12 + b2*a22 + b3*a32
	m.Values[3] = b0*a03 + b1*a13 + b2*a23 + b3*a33

	b0, b1, b2, b3 = other.Values[4], other.Values[5], other.Values[6], other.Values[7]
	m.Values[4] = b0*a00 + b1*a10 + b2*a20 + b3*a30
	m.Values[5] = b0*a01 + b1*a11 + b2*a21 + b3*a31
	m.Values[6] = b0*a02 + b1*a12 + b2*a22 + b3*a32
	m.Values[7] = b0*a03 + b1*a13 + b2*a23 + b3*a33

	b0, b1, b2, b3 = other.Values[8], other.Values[9], other.Values[10], other.Values[11]
	m.Values[8] = b0*a00 + b1*a10 + b2*a20 + b3*a30
	m.Values[9] = b0*a01 + b1*a11 + b2*a21 + b3*a31
	m.Values[10] = b0*a02 + b1*a12 + b2*a22 + b3*a32
	m.Values[11] = b0*a03 + b1*a13 + b2*a23 + b3*a33

	b0, b1, b2, b3 = other.Values[12], other.Values[13], other.Values[14], other.Values[15]
	m.Values[12] = b0*a00 + b1*a10 + b2*a20 + b3*a30
	m.Values[13] = b0*a01 + b1*a11 + b2*a21 + b3*a31
	m.Values[14] = b0*a02 + b1*a12 + b2*a22 + b3*a32
	m.Values[15] = b0*a03 + b1*a13 + b2*a23 + b3*a33

	return m
}

// MultiplyLocal multiplies the values of this Matrix4 by those given in the `other` argument.
func (m *Matrix4) MultiplyLocal(other *Matrix4) *Matrix4 {
	return m.SetValues(
		m.Values[0]*other.Values[0]+m.Values[1]*other.Values[4]+m.Values[2]*other.Values[8]+m.Values[3]*other.Values[12],
		m.Values[0]*other.Values[1]+m.Values[1]*other.Values[5]+m.Values[2]*other.Values[9]+m.Values[3]*other.Values[13],
		m.Values[0]*other.Values[2]+m.Values[1]*other.Values[6]+m.Values[2]*other.Values[10]+m.Values[3]*other.Values[14],
		m.Values[0]*other.Values[3]+m.Values[1]*other.Values[7]+m.Values[2]*other.Values[11]+m.Values[3]*other.Values[15],
		m.Values[4]*other.Values[0]+m.Values[5]*other.Values[4]+m.Values[6]*other.Values[8]+m.Values[7]*other.Values[12],
		m.Values[4]*other.Values[1]+m.Values[5]*other.Values[5]+m.Values[6]*other.Values[9]+m.Values[7]*other.Values[13],
		m.Values[4]*other.Values[2]+m.Values[5]*other.Values[6]+m.Values[6]*other.Values[10]+m.Values[7]*other.Values[14],
		m.Values[4]*other.Values[3]+m.Values[5]*other.Values[7]+m.Values[6]*other.Values[11]+m.Values[7]*other.Values[15],
		m.Values[8]*other.Values[0]+m.Values[9]*other.Values[4]+m.Values[10]*other.Values[8]+m.Values[11]*other.Values[12],
		m.Values[8]*other.Values[1]+m.Values[9]*other.Values[5]+m.Values[10]*other.Values[9]+m.Values[11]*other.Values[13],
		m.Values[8]*other.Values[2]+m.Values[9]*other.Values[6]+m.Values[10]*other.Values[10]+m.Values[11]*other.Values[14],
		m.Values[8]*other.Values[3]+m.Values[9]*other.Values[7]+m.Values[10]*other.Values[11]+m.Values[11]*other.Values[15],
		m.Values[12]*other.Values[0]+m.Values[13]*other.Values[4]+m.Values[14]*other.Values[8]+m.Values[15]*other.Values[12],
		m.Values[12]*other.Values[1]+m.Values[13]*other.Values[5]+m.Values[14]*other.Values[9]+m.Values[15]*other.Values[13],
		m.Values[12]*other.Values[2]+m.Values[13]*other.Values[6]+m.Values[14]*other.Values[10]+m.Values[15]*other.Values[14],
		m.Values[12]*other.Values[3]+m.Values[13]*other.Values[7]+m.Values[14]*other.Values[11]+m.Values[15]*other.Values[15],
	)
}

// PreMultiply multiplies the given Matrix4 object with this Matrix.
func (m *Matrix4) PreMultiply(other *Matrix4) *Matrix4 {
	return m.MultiplyMatrices(other, m)
}

// MultiplyMatrices multiplies the two given Matrix4 objects and stores the results in this Matrix.
func (m *Matrix4) MultiplyMatrices(a, b *Matrix4) *Matrix4 {
	a11 := a.Values[0]
	a12 := a.Values[4]
	a13 := a.Values[8]
	a14 := a.Values[12]
	a21 := a.Values[1]
	a22 := a.Values[5]
	a23 := a.Values[9]
	a24 := a.Values[13]
	a31 := a.Values[2]
	a32 := a.Values[6]
	a33 := a.Values[10]
	a34 := a.Values[14]
	a41 := a.Values[3]
	a42 := a.Values[7]
	a43 := a.Values[11]
	a44 := a.Values[15]

	b11 := b.Values[0]
	b12 := b.Values[4]
	b13 := b.Values[8]
	b14 := b.Values[12]
	b21 := b.Values[1]
	b22 := b.Values[5]
	b23 := b.Values[9]
	b24 := b.Values[13]
	b31 := b.Values[2]
	b32 := b.Values[6]
	b33 := b.Values[10]
	b34 := b.Values[14]
	b41 := b.Values[3]
	b42 := b.Values[7]
	b43 := b.Values[11]
	b44 := b.Values[15]

	return m.SetValues(
		a11*b11+a12*b21+a13*b31+a14*b41,
		a21*b11+a22*b21+a23*b31+a24*b41,
		a31*b11+a32*b21+a33*b31+a34*b41,
		a41*b11+a42*b21+a43*b31+a44*b41,
		a11*b12+a12*b22+a13*b32+a14*b42,
		a21*b12+a22*b22+a23*b32+a24*b42,
		a31*b12+a32*b22+a33*b32+a34*b42,
		a41*b12+a42*b22+a43*b32+a44*b42,
		a11*b13+a12*b23+a13*b33+a14*b43,
		a21*b13+a22*b23+a23*b33+a24*b43,
		a31*b13+a32*b23+a33*b33+a34*b43,
		a41*b13+a42*b23+a43*b33+a44*b43,
		a11*b14+a12*b24+a13*b34+a14*b44,
		a21*b14+a22*b24+a23*b34+a24*b44,
		a31*b14+a32*b24+a33*b34+a34*b44,
		a41*b14+a42*b24+a43*b34+a44*b44,
	)
}

// Translate this Matrix using the given Vector.
func (m *Matrix4) Translate(v Vector3Like) *Matrix4 {
	return m.TranslateXYZ(v.XYZ())
}

// TranslateXYZ translates this Matrix using the given values.
func (m *Matrix4) TranslateXYZ(x, y, z float64) *Matrix4 {
	m.Values[12] = m.Values[0]*x + m.Values[4]*y + m.Values[8]*z + m.Values[12]
	m.Values[13] = m.Values[1]*x + m.Values[5]*y + m.Values[9]*z + m.Values[13]
	m.Values[14] = m.Values[2]*x + m.Values[6]*y + m.Values[10]*z + m.Values[14]
	m.Values[15] = m.Values[3]*x + m.Values[7]*y + m.Values[11]*z + m.Values[15]

	return m
}

// Scale applies a scale transformation to this Matrix.
func (m *Matrix4) Scale(v Vector3Like) *Matrix4 {
	return m.ScaleXYZ(v.XYZ())
}

// ScaleXYZ applies a scale transformation to this Matrix.
func (m *Matrix4) ScaleXYZ(x, y, z float64) *Matrix4 {
	m.Values[0], m.Values[1], m.Values[2], m.Values[3] = m.Values[0]*x, m.Values[1]*x, m.Values[2]*x, m.Values[3]*x
	m.Values[4], m.Values[5], m.Values[6], m.Values[7] = m.Values[4]*y, m.Values[5]*y, m.Values[6]*y, m.Values[7]*y
	m.Values[8], m.Values[9], m.Values[10], m.Values[11] = m.Values[8]*z, m.Values[9]*z, m.Values[10]*z, m.Values[11]*z

	return m
}

// MakeRotationAxis derives a rotation matrix around the given axis.
func (m *Matrix4) MakeRotationAxis(axis Vector3Like, radians float64) *Matrix4 {
	c, s := math.Cos(radians), math.Sin(radians)
	t := 1 - c
	x, y, z := axis.XYZ()
	tx, ty := t*x, t*y

	return m.SetValues(
		tx*x+c, tx*y-s*z, tx*z+s*y, 0,
		tx*y+s*z, ty*y+c, ty*z-s*x, 0,
		tx*z-s*y, ty*z+s*x, t*z*z+c, 0,
		0, 0, 0, 1,
	)
}

// Rotate applies a rotation transformation to this Matrix.
func (m *Matrix4) Rotate(radians float64, axis Vector3Like) *Matrix4 {
	x, y, z := axis.XYZ()
	length := math.Sqrt(x*x + y*y + z*z)

	if math.Abs(length) < Epsilon {
		return m
	}

	length = 1 / length
	x *= length
	y *= length
	z *= length

	c, s := math.Cos(radians), math.Sin(radians)
	t := 1 - c

	a00, a01, a02, a03 := m.Values[0], m.Values[1], m.Values[2], m.Values[3]
	a10, a11, a12, a13 := m.Values[4], m.Values[5], m.Values[6], m.Values[7]
	a20, a21, a22, a23 := m.Values[8], m.Values[9], m.Values[10], m.Values[11]
	a30, a31, a32, a33 := m.Values[12], m.Values[13], m.Values[14], m.Values[15]

	b00, b01, b02 := x*x*t+c, y*x*t+z*s, z*x*t-y*s
	b10, b11, b12 := x*y*t-z*s, y*y*t+c, z*y*t+x*s
	b20, b21, b22 := x*z*t+y*s, y*z*t-x*s, z*z*t+c

	return m.SetValues(
		a00*b00+a10*b01+a20*b02,
		a01*b00+a11*b01+a21*b02,
		a02*b00+a12*b01+a22*b02,
		a03*b00+a13*b01+a23*b02,
		a00*b10+a10*b11+a20*b12,
		a01*b10+a11*b11+a21*b12,
		a02*b10+a12*b11+a22*b12,
		a03*b10+a13*b11+a23*b12,
		a00*b20+a10*b21+a20*b22,
		a01*b20+a11*b21+a21*b22,
		a02*b20+a12*b21+a22*b22,
		a03*b20+a13*b21+a23*b22,
		a30, a31, a32, a33,
	)
}

// RotateX rotates this matrix on its X axis.
func (m *Matrix4) RotateX(radians float64) *Matrix4 {
	c, s := math.Cos(radians), math.Sin(radians)

	a10, a11, a12, a13 := m.Values[4], m.Values[5], m.Values[6], m.Values[7]
	a20, a21, a22, a23 := m.Values[8], m.Values[9], m.Values[10], m.Values[11]

	//  Perform axis-specific matrix multiplication
	m.Values[4] = a10*c + a20*s
	m.Values[5] = a11*c + a21*s
	m.Values[6] = a12*c + a22*s
	m.Values[7] = a13*c + a23*s
	m.Values[8] = a20*c - a10*s
	m.Values[9] = a21*c - a11*s
	m.Values[10] = a22*c - a12*s
	m.Values[11] = a23*c - a13*s

	return m
}

// RotateY rotates this matrix on its X axis.
func (m *Matrix4) RotateY(radians float64) *Matrix4 {
	c, s := math.Cos(radians), math.Sin(radians)

	a00, a01, a02, a03 := m.Values[0], m.Values[1], m.Values[2], m.Values[3]
	a20, a21, a22, a23 := m.Values[8], m.Values[9], m.Values[10], m.Values[11]

	//  Perform axis-specific matrix multiplication
	m.Values[0] = a00*c + a20*s
	m.Values[1] = a01*c + a21*s
	m.Values[2] = a02*c + a22*s
	m.Values[3] = a03*c + a23*s
	m.Values[4] = a20*c - a00*s
	m.Values[5] = a21*c - a01*s
	m.Values[6] = a22*c - a02*s
	m.Values[7] = a23*c - a03*s

	return m
}

// RotateZ rotates this matrix on its X axis.
func (m *Matrix4) RotateZ(radians float64) *Matrix4 {
	c, s := math.Cos(radians), math.Sin(radians)

	a00, a01, a02, a03 := m.Values[0], m.Values[1], m.Values[2], m.Values[3]
	a10, a11, a12, a13 := m.Values[4], m.Values[5], m.Values[6], m.Values[7]

	//  Perform axis-specific matrix multiplication
	m.Values[0] = a00*c + a10*s
	m.Values[1] = a01*c + a11*s
	m.Values[2] = a02*c + a12*s
	m.Values[3] = a03*c + a13*s
	m.Values[4] = a10*c - a00*s
	m.Values[5] = a11*c - a01*s
	m.Values[6] = a12*c - a02*s
	m.Values[7] = a13*c - a03*s

	return m
}

// FromRotationTranslation sets the values of this Matrix from the given rotation Quaternion and
// translation Vector.
func (m *Matrix4) FromRotationTranslation(q *Quaternion, v *Vector3) *Matrix4 {
	x, y, z, w := q.X, q.Y, q.Z, q.W

	x2, y2, z2 := x+x, y+y, z+z
	xx, xy, xz := x*x2, x*y2, x*z2
	yy, yz, zz := y*y2, y*z2, z*z2
	wx, wy, wz := w*x2, w*y2, w*z2

	return m.SetValues(
		1-(yy+zz), xy+wz, xz-wy, 0,
		xy-wz, 1-(xx+zz), yz+wx, 0,
		xz+wy, yz-wx, 1-(xx+yy), 0,
		v.X, v.Y, v.Z, 1,
	)
}

// FromQuaternion sets the values of this Matrix from the given Quaternion.
func (m *Matrix4) FromQuaternion(q *Quaternion) *Matrix4 {
	x, y, z, w := q.X, q.Y, q.Z, q.W

	x2, y2, z2 := x+x, y+y, z+z
	xx, xy, xz := x*x2, x*y2, x*z2
	yy, yz, zz := y*y2, y*z2, z*z2
	wx, wy, wz := w*x2, w*y2, w*z2

	return m.SetValues(
		1-(yy+zz), xy+wz, xz-wy, 0,
		xy-wz, 1-(xx+zz), yz+wx, 0,
		xz+wy, yz-wx, 1-(xx+yy), 0,
		0, 0, 0, 1,
	)
}

// Frustum generates a frustum matrix with the given bounds.
func (m *Matrix4) Frustum(left, right, bottom, top, near, far float64) *Matrix4 {
	rl, tb, nf := 1/(right-left), 1/(top-bottom), 1/(near-far)

	return m.SetValues(
		(near*2)*rl, 0, 0, 0,
		0, (near*2)*tb, 0, 0,
		(right+left)*rl, (top+bottom)*tb, (far+near)*nf, -1,
		0, 0, (far*near*2)*nf, 0,
	)
}

// Perspective generates a perspective projection matrix with the given bounds.
func (m *Matrix4) Perspective(fovy, aspect, near, far float64) *Matrix4 {
	f, nf := 1/math.Tan(fovy/2), 1/(near-far)

	return m.SetValues(
		f/aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (far+near)*nf, -1,
		0, 0, (2*far*near)*nf, 0,
	)
}

// PerspectiveLH generates a perspective projection matrix with the given bounds.
func (m *Matrix4) PerspectiveLH(width, height, near, far float64) *Matrix4 {
	return m.SetValues(
		(2*near)/width, 0, 0, 0,
		0, (2*near)/height, 0, 0,
		0, 0, -far/(near-far), 1,
		0, 0, (near*far)/(near-far), 0,
	)
}

// Ortho generates an orthogonal projection matrix with the given bounds.
func (m *Matrix4) Ortho(left, right, bottom, top, near, far float64) *Matrix4 {
	lr, bt, nf := left-right, bottom-top, near-far

	// Avoid division by zero
	if lr != 0 {
		lr = 1 / lr
	}

	if bt != 0 {
		bt = 1 / bt
	}

	if nf != 0 {
		nf = 1 / nf
	}

	return m.SetValues(
		-2*lr, 0, 0, 0,
		0, -2*bt, 0, 0,
		0, 0, 2*nf, 0,
		(left+right)*lr, (top+bottom)*bt, (far+near)*nf, 1,
	)
}

// LookAtRightHanded generates a right-handed look-at matrix with the given eye position,
// target and up axis.
func (m *Matrix4) LookAtRightHanded(eye, target, up *Vector3) *Matrix4 {
	vz := eye.Clone().Subtract(target)
	vx := NewVector3(0, 0, 0)
	vy := NewVector3(0, 0, 0)

	if vz.LengthSquared() == 0 {
		// eye and target are in the same position
		vz.Z = 1
	}

	vz.Normalize()
	vx.CrossVectors(up, vz)

	if vx.LengthSquared() == 0 {
		if math.Abs(up.Z) == 1 {
			vz.X += Epsilon
		} else {
			vz.Z += Epsilon
		}

		vz.Normalize()
		vx.CrossVectors(vz, vx)
	}

	vx.Normalize()
	vy.CrossVectors(vz, vx)

	m.Values[0] = vx.X
	m.Values[1] = vx.Y
	m.Values[2] = vx.Z
	m.Values[4] = vy.X
	m.Values[5] = vy.Y
	m.Values[6] = vy.Z
	m.Values[8] = vz.X
	m.Values[9] = vz.Y
	m.Values[10] = vz.Z

	return m
}

//  LookAt generates a look-at matrix with the given eye position, target, and up axis.
func (m *Matrix4) LookAt(eye, target, up *Vector3) *Matrix4 {
	ex, ey, ez := eye.XYZ()
	tx, ty, tz := target.XYZ()
	ux, uy, uz := up.XYZ()

	if math.Abs(ex-tx) < Epsilon && math.Abs(ey-ty) < Epsilon && math.Abs(ez-tz) < Epsilon {
		return m.Identity()
	}

	z0, z1, z2 := ex-tx, ey-ty, ez-tz
	length := 1 / math.Sqrt(z0*z0+z1*z1+z2*z2)

	x0, x1, x2 := uy*z2-uz*z1, uz*z0-ux*z2, ux*z1-uy*z0

	if length == 0 {
		x0, x1, x2 = 0, 0, 0
	} else {
		length = 1 / length
		x0 *= length
		x1 *= length
		x2 *= length
	}

	y0, y1, y2 := z1*x2-z2*x1, z2*x0-z0*x2, z0*x1-z1*x0

	length = math.Sqrt(y0*y0 + y1*y1 + y2*y2)

	if length == 0 {
		y0, y1, y2 = 0, 0, 0
	} else {
		length = 1 / length
		y0 *= length
		y1 *= length
		y2 *= length
	}

	return m.SetValues(
		x0, y0, z0, 0,
		x1, y1, z1, 0,
		x2, y2, z2, 0,
		-(x0*ex + x1*ey + x2*ez), -(y0*ex + y1*ey + y2*ez), -(z0*ex + z1*ey + z2*ez), 1,
	)
}

// SetYawPitchRoll sets the values of this matrix from the given `yaw`, `pitch` and `roll` values.
func (m *Matrix4) SetYawPitchRoll(y, p, r float64) *Matrix4 {
	m.Zero()
	a, b := m.Clone(), m.Clone()
	vm, va, vb := m.Values, a.Values, b.Values

	// Rotate Z
	s, c := math.Sin(r), math.Cos(r)
	vm[10], vm[15], vm[0], vm[1], vm[4], vm[5] = 1, 1, c, s, -s, c

	// Rotate X
	s, c = math.Sin(p), math.Cos(p)
	va[0], va[15], va[5], va[10], va[9], va[6] = 1, 1, c, c, -s, s

	// Rotate Y
	s, c = math.Sin(y), math.Cos(y)
	vb[5], vb[15], vb[0], vb[2], vb[8], vb[10] = 1, 1, c, -s, s, c

	return m.MultiplyLocal(a).MultiplyLocal(b)
}

// SetWorldMatrix generates a world matrix from the given (rotation, position, scale vector3),
// and (view, projection matrix4).
func (m *Matrix4) SetWorldMatrix(rot, pos, scale *Vector3, view, proj *Matrix4) *Matrix4 {
	m.SetYawPitchRoll(rot.XYZ())

	a := NewMatrix4(nil).SetScaling(scale.XYZ())
	b := NewMatrix4(nil).SetXYZ(pos.XYZ())

	m.MultiplyLocal(a).MultiplyLocal(b)

	if view != nil {
		m.MultiplyLocal(view)
	}

	if proj != nil {
		m.MultiplyLocal(proj)
	}

	return m
}

// MultiplyToMatrix4 multiplies this Matrix4 by the given `src` Matrix4 and stores the results in
// the `out` Matrix4.
func (m *Matrix4) MultiplyToMatrix4(src, out *Matrix4) *Matrix4 {
	if out == nil {
		out = NewMatrix4(nil)
	}

	mv, sv := m.Values, src.Values

	a00, a01, a02, a03,
		a10, a11, a12, a13,
		a20, a21, a22, a23,
		a30, a31, a32, a33 :=
		mv[0], mv[1], mv[2], mv[3],
		mv[4], mv[5], mv[6], mv[7],
		mv[8], mv[9], mv[10], mv[11],
		mv[12], mv[13], mv[14], mv[15]

	b00, b01, b02, b03,
		b10, b11, b12, b13,
		b20, b21, b22, b23,
		b30, b31, b32, b33 :=
		sv[0], sv[1], sv[2], sv[3],
		sv[4], sv[5], sv[6], sv[7],
		sv[8], sv[9], sv[10], sv[11],
		sv[12], sv[13], sv[14], sv[15]

	return out.SetValues(
		b00*a00+b01*a10+b02*a20+b03*a30,
		b01*a01+b01*a11+b02*a21+b03*a31,
		b02*a02+b01*a12+b02*a22+b03*a32,
		b03*a03+b01*a13+b02*a23+b03*a33,
		b10*a00+b11*a10+b12*a20+b13*a30,
		b10*a01+b11*a11+b12*a21+b13*a31,
		b10*a02+b11*a12+b12*a22+b13*a32,
		b10*a03+b11*a13+b12*a23+b13*a33,
		b20*a00+b21*a10+b22*a20+b23*a30,
		b20*a01+b21*a11+b22*a21+b23*a31,
		b20*a02+b21*a12+b22*a22+b23*a32,
		b20*a03+b21*a13+b22*a23+b23*a33,
		b30*a00+b31*a10+b32*a20+b33*a30,
		b30*a01+b31*a11+b32*a21+b33*a31,
		b30*a02+b31*a12+b32*a22+b33*a32,
		b30*a03+b31*a13+b32*a23+b33*a33,
	)
}

// FromRotationXYTranslation takes the rotation, position vectors and builds this Matrix4 from them.
func (m *Matrix4) FromRotationXYTranslation(rot, pos *Vector3, translateFirst bool) *Matrix4 {
	x, y, z := pos.XYZ()

	sx, cx := math.Sin(rot.X), math.Cos(rot.X)
	sy, cy := math.Sin(rot.Y), math.Cos(rot.Y)

	a30, a31, a32 := x, y, z

	// rotate x
	b21 := -sx

	// rotate y
	c01, c02, c21, c22 := 0-b21*sy, 0-cx*sy, b21*cy, cx*cy

	// translate
	if !translateFirst {
		a30, a31, a32 = cy*x+sy*z, c01*x+cx*y+c21*z, c02*x+sx*y+c22*z
	}

	return m.SetValues(
		cy, c01, c02, 0,
		0, cx, sx, 0,
		sy, c21, c22, 0,
		a30, a31, a32, 1,
	)
}

// GetMaxScaleOnAxis returns the maximum axis scale from this Matrix4.
func (m *Matrix4) GetMaxScaleOnAxis() float64 {
	v := m.Values[:]

	sx2 := v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
	sy2 := v[4]*v[4] + v[5]*v[5] + v[6]*v[6]
	sz2 := v[8]*v[8] + v[9]*v[9] + v[10]*v[10]

	return math.Sqrt(math.Max(math.Max(sx2, sy2), sz2))
}
