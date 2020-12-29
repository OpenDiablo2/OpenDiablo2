package d2vector

import (
	"fmt"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

//nolint:gochecknoglobals // These variables are assigned to in benchmark functions to avoid compiler optimisations
// lowering the runtime of the benchmark. See: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go (A note
// on compiler optimisations)
var (
	outVector Vector
	outFloat  float64
	outBool   bool
	outInt    int
)

func TestVector_Equals(t *testing.T) {
	a := NewVector(1, 2)
	b := NewVector(1, 2)

	got := a.Equals(b)

	if !got {
		t.Errorf("exact equality %s and %s: wanted true: got %t", a, b, got)
	}

	c := NewVector(3, 4)

	got = a.Equals(c)

	if got {
		t.Errorf("exact equality %s and %s: wanted false: got %t", a, c, got)
	}
}

func BenchmarkVector_Equals(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(2, 2)

	for n := 0; n < b.N; n++ {
		outBool = v.Equals(o)
	}
}

func TestVector_EqualsApprox(t *testing.T) {
	subEpsilon := d2math.Epsilon / 3

	a := NewVector(1, 2)
	b := NewVector(1+subEpsilon, 2+subEpsilon)

	got := a.EqualsApprox(b)

	if !got {
		t.Errorf("approximate equality %s and %s: wanted true: got %t", a, b, got)
	}

	c := NewVector(1+d2math.Epsilon, 2+d2math.Epsilon)

	got = a.EqualsApprox(c)

	if got {
		t.Errorf("approximate equality %s and %s: wanted false: got %t", a, c, got)
	}
}

func BenchmarkVector_EqualsApprox(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(2, 2)

	for n := 0; n < b.N; n++ {
		outBool = v.EqualsApprox(o)
	}
}

func TestVector_CompareApprox(t *testing.T) {
	subEpsilon := d2math.Epsilon / 3

	f := NewVector(1+subEpsilon, 1+subEpsilon)
	c := NewVector(1, 1)
	xWant, yWant := 0, 0
	yGot, xGot := f.CompareApprox(c)

	if xGot != xWant || yGot != yWant {
		t.Errorf("approximate comparison %s and %s: wanted (%d, %d): got (%d, %d)", f, c, xWant, yWant, xGot, yGot)
	}

	f = NewVector(2, 2)
	c = NewVector(-1, 3)
	xWant, yWant = 1, -1
	xGot, yGot = f.CompareApprox(c)

	if xGot != xWant || yGot != yWant {
		t.Errorf("approximate comparison %s and %s: wanted (%d, %d): got (%d, %d)", f, c, xWant, yWant, xGot, yGot)
	}

	f = NewVector(2, 2)
	c = NewVector(3, -1)
	xWant, yWant = -1, 1
	xGot, yGot = f.CompareApprox(c)

	if xGot != xWant || yGot != yWant {
		t.Errorf("approximate comparison %s and %s: wanted (%d, %d): got (%d, %d)", f, c, xWant, yWant, xGot, yGot)
	}
}

func BenchmarkVector_CompareApprox(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(2, 2)

	for n := 0; n < b.N; n++ {
		outInt, outInt = v.CompareApprox(o)
	}
}

func TestVector_IsZero(t *testing.T) {
	testIsZero(*NewVector(0, 0), true, t)
	testIsZero(*NewVector(1, 0), false, t)
	testIsZero(*NewVector(0, 1), false, t)
	testIsZero(*NewVector(1, 1), false, t)
}

func testIsZero(v Vector, want bool, t *testing.T) {
	got := v.IsZero()
	if got != want {
		t.Errorf("%s is zero: want %t: got %t", v, want, got)
	}
}

func BenchmarkVector_IsZero(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		outBool = v.IsZero()
	}
}

func TestVector_Set(t *testing.T) {
	v := NewVector(1, 1)
	want := NewVector(2, 3)
	got := v.Clone()
	got.Set(2, 3)

	if !got.Equals(want) {
		t.Errorf("set %s to (2, 3): want %s: got %s", v, want, got)
	}
}

func BenchmarkVector_Set(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		v.Set(1, 1)
	}
}

func TestVector_Clone(t *testing.T) {
	want := NewVector(1, 2)
	got := want.Clone()

	if !got.Equals(want) {
		t.Errorf("clone %s: want %s: got %s", want, want, got)
	}
}

func BenchmarkVector_Clone(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		outVector = *v.Clone()
	}
}

func TestVector_Copy(t *testing.T) {
	want := NewVector(1, 2)
	got := NewVector(0, 0)
	got.Copy(want)

	if !got.Equals(want) {
		t.Errorf("copy %s to %s: want %s: got %s", got, want, want, got)
	}
}

func BenchmarkVector_Copy(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(2, 2)

	for n := 0; n < b.N; n++ {
		v.Copy(o)
	}
}

func TestVector_Floor(t *testing.T) {
	v := NewVector(1.6, 1.6)
	want := NewVector(1, 1)
	got := v.Clone()
	got.Floor()

	if !got.Equals(want) {
		t.Errorf("round %s down: want %s: got %s", v, want, got)
	}
}

func BenchmarkVector_Floor(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		v.Floor()
	}
}

func TestVector_Clamp(t *testing.T) {
	v := NewVector(-10, 10)
	c := v.Clone()
	a := NewVector(2, 2)
	b := NewVector(7, 7)

	want := NewVector(2, 7)
	got := v.Clamp(a, b)

	if !got.Equals(want) {
		t.Errorf("clamp %s between %s and %s: want %s: got %s", c, a, b, want, *got)
	}
}

func BenchmarkVector_Clamp(b *testing.B) {
	v := NewVector(1, 1)
	min := NewVector(0, 0)
	max := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		v.Clamp(min, max)
	}
}

func TestVector_Add(t *testing.T) {
	v := NewVector(1, 2)
	add := NewVector(0.5, 3)
	want := NewVector(1.5, 5)
	got := v.Clone()
	got.Add(add)

	if !got.Equals(want) {
		t.Errorf("add %s to %s: want %s: got %s", add, v, want, got)
	}
}

func BenchmarkVector_Add(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(0.01, 0.01)

	for n := 0; n < b.N; n++ {
		v.Add(o)
	}
}

func TestVector_AddScalar(t *testing.T) {
	v := NewVector(1, -1)
	add := 0.5
	want := NewVector(1.5, -0.5)
	got := v.Clone()
	got.AddScalar(add)

	if !got.Equals(want) {
		t.Errorf("add %.2f to %s: want %s: got %s", add, v, want, got)
	}
}

func BenchmarkVector_AddScalar(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		v.AddScalar(0.01)
	}
}

func TestVector_Subtract(t *testing.T) {
	v := NewVector(1, 1)
	subtract := NewVector(0.6, 0.6)
	want := NewVector(0.4, 0.4)
	got := v.Clone()
	got.Subtract(subtract)

	if !got.Equals(want) {
		t.Errorf("subtract %s from %s: want %s: got %s", subtract, v, want, got)
	}
}

func BenchmarkVector_Subtract(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(0.01, 0.01)

	for n := 0; n < b.N; n++ {
		v.Subtract(o)
	}
}

func TestVector_Multiply(t *testing.T) {
	v := NewVector(1, 1)
	multiply := NewVector(2, 2)
	want := NewVector(2, 2)
	got := v.Clone()
	got.Multiply(multiply)

	if !got.Equals(want) {
		t.Errorf("multiply %s by %s: want %s: got %s", v, multiply, want, got)
	}
}

func BenchmarkVector_Multiply(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(0.01, 0.01)

	for n := 0; n < b.N; n++ {
		v.Multiply(o)
	}
}

func TestVector_Divide(t *testing.T) {
	v := NewVector(1, 8)
	divide := NewVector(2, 4)
	want := NewVector(0.5, 2)
	got := v.Clone()
	got.Divide(divide)

	if !got.Equals(want) {
		t.Errorf("divide %s by %s: want %s: got %s", v, divide, want, got)
	}
}

func BenchmarkVector_Divide(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(3, 3)

	for n := 0; n < b.N; n++ {
		v.Divide(o)
	}
}

func TestVector_DivideScalar(t *testing.T) {
	v := NewVector(1, 2)
	divide := 2.0
	want := NewVector(0.5, 1.0)
	got := v.Clone()
	got.DivideScalar(divide)

	if !got.Equals(want) {
		t.Errorf("divide %s by %.2f: want %s: got %s", v, divide, want, got)
	}
}

func BenchmarkVector_DivideScalar(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		v.DivideScalar(3)
	}
}

func TestVector_Scale(t *testing.T) {
	v := NewVector(2, 3)
	want := NewVector(4, 6)
	got := v.Clone()
	got.Scale(2)

	if !got.Equals(want) {
		t.Errorf("scale %s by 2: want %s: got %s", v, want, got)
	}
}

func BenchmarkVector_Scale(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		v.Scale(1.01)
	}
}

func TestVector_Abs(t *testing.T) {
	v := NewVector(-1, 1)
	want := NewVector(1, 1)
	got := v.Clone()
	got.Abs()

	if !want.Equals(got) {
		t.Errorf("absolute value of %s: want %s: got %s", v, want, got)
	}
}

func BenchmarkVector_Abs(b *testing.B) {
	v := NewVector(-1, -1)

	for n := 0; n < b.N; n++ {
		outVector = *v.Abs()
	}
}

func TestVector_Negate(t *testing.T) {
	v := NewVector(-1, 1)
	want := NewVector(1, -1)
	got := v.Clone()
	got.Negate()

	if !want.Equals(got) {
		t.Errorf("inverse value of %s: want %s: got %s", v, want, got)
	}
}

func BenchmarkVector_Negate(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		outVector = *v.Negate()
	}
}

func TestVector_Distance(t *testing.T) {
	v := NewVector(1, 3)
	other := NewVector(1, -1)
	want := 4.0
	c := v.Clone()
	got := c.Distance(other)

	if got != want {
		t.Errorf("distance from %s to %s: want %.3f: got %.3f", v, other, want, got)
	}
}

func BenchmarkVector_Distance(b *testing.B) {
	v := NewVector(1, 1)
	d := NewVector(2, 2)

	for n := 0; n < b.N; n++ {
		outFloat = v.Distance(d)
	}
}

func TestVector_Length(t *testing.T) {
	v := NewVector(2, 0)
	c := v.Clone()
	want := 2.0
	got := c.Length()

	d := fmt.Sprintf("length of %s", v)

	if !c.Equals(v) {
		t.Errorf("%s: changed vector %s to %s unexpectedly", d, v, c)
	}

	if got != want {
		t.Errorf("%s: want %.3f: got %.3f", d, want, got)
	}
}

func BenchmarkVector_Length(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		outFloat = v.Length()
	}
}

func TestVector_SetLength(t *testing.T) {
	v := NewVector(1, 1)
	c := v.Clone()
	want := 2.0
	got := c.SetLength(want).Length()

	if !d2math.EqualsApprox(got, want) {
		t.Errorf("set length of %s to %.3f :want %.3f: got %.3f", v, want, want, got)
	}
}

func BenchmarkVector_SetLength(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		v.SetLength(5)
	}
}

func TestVector_Lerp(t *testing.T) {
	a := NewVector(0, 0)
	b := NewVector(-20, 10)

	interp := 0.3

	want := NewVector(-6, 3)
	got := a.Lerp(b, interp)

	if !got.Equals(want) {
		t.Errorf("linear interpolation between %s and %s by %.2f: want %s: got %s", a, b, interp, want, got)
	}
}

func BenchmarkVector_Lerp(b *testing.B) {
	v := NewVector(1, 1)
	t := NewVector(1000, 1000)

	for n := 0; n < b.N; n++ {
		v.Lerp(t, 1.01)
	}
}

func TestVector_Dot(t *testing.T) {
	v := NewVector(1, 1)
	c := v.Clone()
	want := 2.0
	got := c.Dot(c)

	d := fmt.Sprintf("dot product of %s", v)

	if !c.Equals(v) {
		t.Errorf("%s: changed vector %s to %s unexpectedly", d, v, c)
	}

	if got != want {
		t.Errorf("%s: want %.3f: got %.3f", d, want, got)
	}
}

func BenchmarkVector_Dot(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		outFloat = v.Dot(v)
	}
}

func TestVector_Cross(t *testing.T) {
	v := NewVector(1, 1)

	clock := NewVector(1, 0)
	anti := NewVector(0, 1)

	want := -1.0
	got := v.Cross(clock)

	if got != want {
		t.Errorf("cross product of %s and %s: want %.3f: got %.3f", v, clock, want, got)
	}

	want = 1.0
	got = v.Cross(anti)

	if got != want {
		t.Errorf("cross product of %s and %s: want %.3f: got %.3f", v, clock, want, got)
	}
}

func BenchmarkVector_Cross(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(0, 1)

	for n := 0; n < b.N; n++ {
		outFloat = v.Cross(o)
	}
}

func TestVector_Normalize(t *testing.T) {
	v := NewVector(10, 0)
	c := v.Clone()
	want := NewVector(1, 0)

	c.Normalize()

	if !want.Equals(c) {
		t.Errorf("normalize %s: want %s: got %s", v, want, c)
	}

	v = NewVector(0, 10)
	c = v.Clone()
	want = NewVector(0, 1)

	c.Normalize()

	if !want.Equals(c) {
		t.Errorf("normalize %s: want %s: got %s", v, want, c)
	}

	v = NewVector(0, 0)
	c = v.Clone()
	want = NewVector(0, 0)

	c.Normalize()

	if !want.Equals(c) {
		t.Errorf("normalize zero vector %s should do nothing: want %s: got %s", v, want, c)
	}
}

func BenchmarkVector_Normalize(b *testing.B) {
	v := NewVector(1, 1)

	for n := 0; n < b.N; n++ {
		outVector = *v.Normalize()
	}
}

func TestVector_Angle(t *testing.T) {
	v := NewVector(0, 1)
	c := v.Clone()
	other := NewVector(1, 0.3)

	want := 1.2793395323170293
	got := c.Angle(other)

	d := fmt.Sprintf("angle from %s to %s", v, other)

	if got != want {
		t.Errorf("%s: want %g: got %g", d, want, got)
	}

	if !c.Equals(v) {
		t.Errorf("%s: changed vector %s to %s unexpectedly", d, v, c)
	}

	other.Set(-1, 0.3)
	co := other.Clone()
	got = c.Angle(other)

	d = fmt.Sprintf("angle from %s to %s", c, other)

	if got != want {
		t.Errorf("%s: want %g: got %g", d, want, got)
	}

	if !co.Equals(other) {
		t.Errorf("%s: changed vector %s to %s unexpectedly", d, co, other)
	}
}

func BenchmarkVector_Angle(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(0, 1)

	for n := 0; n < b.N; n++ {
		outFloat = v.Angle(o)
	}
}

func TestVector_SignedAngle(t *testing.T) {
	v := NewVector(0, 1)
	c := v.Clone()
	other := NewVector(1, 0.3)
	want := 1.2793395323170293
	got := c.SignedAngle(other)

	d := fmt.Sprintf("angle from %s to %s", v, other)

	if got != want {
		t.Errorf("%s: want %g: got %g", d, want, got)
	}

	if !c.Equals(v) {
		t.Errorf("%s: changed vector %s to %s unexpectedly", d, v, c)
	}

	other.Set(-1, 0.3)
	co := other.Clone()
	want = 5.0038457214660585
	got = c.SignedAngle(other)

	d = fmt.Sprintf("angle from %s to %s", v, other)

	if got != want {
		t.Errorf("%s: want %g: got %g", d, want, got)
	}

	if !co.Equals(other) {
		t.Errorf("%s: changed vector %s to %s unexpectedly", d, co, other)
	}
}

func BenchmarkVector_SignedAngle(b *testing.B) {
	v := NewVector(1, 1)
	o := NewVector(0, 1)

	for n := 0; n < b.N; n++ {
		outFloat = v.SignedAngle(o)
	}
}

func TestVector_Reflect(t *testing.T) {
	v := NewVector(1, -1)
	c := v.Clone()
	up := NewVector(0, 1)

	want := NewVector(1, 1)
	want.Normalize()

	v.Reflect(up)

	if !want.Equals(v) {
		t.Errorf("reflect direction %s off surface with normal %s: want %s: got %s", c, up, want, v)
	}
}

func BenchmarkVector_Reflect(b *testing.B) {
	v := NewVector(1, -1)
	o := NewVector(0, 1)

	for n := 0; n < b.N; n++ {
		v.Reflect(o)
	}
}

func TestVector_ReflectSurface(t *testing.T) {
	v := NewVector(1, -1)
	c := v.Clone()
	up := NewVector(0, 1)

	want := NewVector(-1, -1)
	want.Normalize()

	v.ReflectSurface(up)

	if !want.Equals(v) {
		t.Errorf("reflect direction %s off surface with normal %s: want %s: got %s", c, up, want, v)
	}
}

func BenchmarkVector_ReflectSurface(b *testing.B) {
	v := NewVector(1, -1)
	o := NewVector(0, 1)

	for n := 0; n < b.N; n++ {
		v.ReflectSurface(o)
	}
}

func TestVector_Rotate(t *testing.T) {
	right := NewVector(1, 0)
	c := right.Clone()

	up := NewVector(0, 1)
	angle := -up.SignedAngle(right)
	want := NewVector(0, 1)
	got := right.Rotate(angle)

	if !want.EqualsApprox(got) {
		t.Errorf("rotated %s by %.1f: want %s: got %s", c, angle*d2math.RadToDeg, want, got)
	}

	c = up.Clone()
	angle -= d2math.RadFull
	want = NewVector(-1, 0)
	got = up.Rotate(angle)

	if !want.EqualsApprox(got) {
		t.Errorf("rotated %s by %.1f: want %s: got %s", c, angle*d2math.RadToDeg, want, got)
	}
}

func BenchmarkVector_Rotate(b *testing.B) {
	v := NewVector(1, 0)
	angle := 45.0

	for n := 0; n < b.N; n++ {
		v.Rotate(angle)
	}
}

func TestVector_NinetyAnti(t *testing.T) {
	v := NewVector(0, 1)
	c := v.Clone()

	want := NewVector(-1, 0)
	got := v.NinetyAnti()

	if !want.Equals(got) {
		t.Errorf("rotated %s by 90 degrees clockwise: want %s: got %s", c, want, *got)
	}
}

func BenchmarkVector_NinetyAnti(b *testing.B) {
	v := NewVector(1, 0)

	for n := 0; n < b.N; n++ {
		v.NinetyAnti()
	}
}

func TestVector_NinetyClock(t *testing.T) {
	v := NewVector(0, 1)
	c := v.Clone()

	want := NewVector(1, 0)
	v = c.Clone()
	got := v.NinetyClock()

	if !want.Equals(got) {
		t.Errorf("rotated %s by 90 degrees anti-clockwise: want %s: got %s", c, want, *got)
	}
}

func BenchmarkVector_NinetyClock(b *testing.B) {
	v := NewVector(1, 0)

	for n := 0; n < b.N; n++ {
		v.NinetyClock()
	}
}

func TestVectorUp(t *testing.T) {
	got := VectorUp()
	want := NewVector(0, 1)

	if !want.Equals(got) {
		t.Errorf("create normalized vector with up direction: want %s: got %s", want, got)
	}
}

func TestVectorDown(t *testing.T) {
	got := VectorDown()
	want := NewVector(0, -1)

	if !want.Equals(got) {
		t.Errorf("create normalized vector with down direction: want %s: got %s", want, got)
	}
}

func TestVectorRight(t *testing.T) {
	got := VectorRight()
	want := NewVector(1, 0)

	if !want.Equals(got) {
		t.Errorf("create normalized vector with right direction: want %s: got %s", want, got)
	}
}

func TestVectorLeft(t *testing.T) {
	got := VectorLeft()
	want := NewVector(-1, 0)

	if !want.Equals(got) {
		t.Errorf("create normalized vector with left direction: want %s: got %s", want, got)
	}
}

func TestVectorOne(t *testing.T) {
	got := VectorOne()
	want := NewVector(1, 1)

	if !want.Equals(got) {
		t.Errorf("create vector with X and Y values of 1: want %s: got %s", want, got)
	}
}

func TestVectorZero(t *testing.T) {
	got := VectorZero()
	want := NewVector(0, 0)

	if !want.Equals(got) {
		t.Errorf("create vector with X and Y values of 0: want %s: got %s", want, got)
	}
}
