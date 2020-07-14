package d2vector

import (
	"fmt"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

func evaluateVector(description string, want, got Vector, t *testing.T) {
	if !got.Equals(want) {
		t.Errorf("%s: wanted %s: got %s", description, want, got)
	}
}

func evaluateVectorApprox(description string, want, got Vector, t *testing.T) {
	if !got.EqualsApprox(want) {
		t.Errorf("%s: wanted %s: got %s", description, want, got)
	}
}

func evaluateScalar(description string, want, got float64, t *testing.T) {
	if want != got {
		t.Errorf("%s: wanted %f: got %f", description, want, got)
	}
}

func evaluateScalarApprox(description string, want, got float64, t *testing.T) {
	if d2math.CompareFloat64Fuzzy(want, got) != 0 {
		t.Errorf("%s: wanted %f: got %f", description, want, got)
	}
}

func evaluateChanged(description string, original, clone Vector, t *testing.T) {
	if !original.Equals(clone) {
		t.Errorf("%s: changed vector %s to %s unexpectedly", description, clone, original)
	}
}

func TestEquals(t *testing.T) {
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

func TestEqualsF(t *testing.T) {
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

func TestCompareApprox(t *testing.T) {
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

func TestIsZero(t *testing.T) {
	testIsZero(NewVector(0, 0), true, t)
	testIsZero(NewVector(1, 0), false, t)
	testIsZero(NewVector(0, 1), false, t)
	testIsZero(NewVector(1, 1), false, t)
}

func testIsZero(v Vector, want bool, t *testing.T) {
	got := v.IsZero()
	if got != want {
		t.Errorf("%s is zero: want %t: got %t", v, want, got)
	}
}

func TestSet(t *testing.T) {
	v := NewVector(1, 1)
	want := NewVector(2, 3)
	got := v.Clone()
	got.Set(2, 3)

	evaluateVector(fmt.Sprintf("set %s to (2, 3)", v), want, got, t)
}

func TestClone(t *testing.T) {
	want := NewVector(1, 2)
	got := want.Clone()

	evaluateVector(fmt.Sprintf("clone %s", want), want, got, t)
}

func TestCopy(t *testing.T) {
	want := NewVector(1, 2)
	got := NewVector(0, 0)
	got.Copy(&want)

	evaluateVector(fmt.Sprintf("copy %s to %s", got, want), want, got, t)
}

func TestFloor(t *testing.T) {
	v := NewVector(1.6, 1.6)
	want := NewVector(1, 1)
	got := v.Clone()
	got.Floor()

	evaluateVector(fmt.Sprintf("round %s down", v), want, got, t)
}

func TestClamp(t *testing.T) {
	v := NewVector(-10, 10)
	c := v.Clone()
	a := NewVector(2, 2)
	b := NewVector(7, 7)

	want := NewVector(2, 7)
	got := v.Clamp(&a, &b)

	evaluateVector(fmt.Sprintf("clamp %s between %s and %s", c, a, b), want, *got, t)
}

func TestAdd(t *testing.T) {
	v := NewVector(1, 2)
	add := NewVector(0.5, 3)
	want := NewVector(1.5, 5)
	got := v.Clone()
	got.Add(&add)

	evaluateVector(fmt.Sprintf("add %s to %s", add, v), want, got, t)
}

func TestAddScalar(t *testing.T) {
	v := NewVector(1, -1)
	add := 0.5
	want := NewVector(1.5, -0.5)
	got := v.Clone()
	got.AddScalar(add)

	evaluateVector(fmt.Sprintf("add %.2f to %s", add, v), want, got, t)
}

func TestSubtract(t *testing.T) {
	v := NewVector(1, 1)
	subtract := NewVector(0.6, 0.6)
	want := NewVector(0.4, 0.4)
	got := v.Clone()
	got.Subtract(&subtract)

	evaluateVector(fmt.Sprintf("subtract %s from %s", subtract, v), want, got, t)
}

func TestMultiply(t *testing.T) {
	v := NewVector(1, 1)
	multiply := NewVector(2, 2)
	want := NewVector(2, 2)
	got := v.Clone()
	got.Multiply(&multiply)

	evaluateVector(fmt.Sprintf("multiply %s by %s", v, multiply), want, got, t)
}

func TestDivide(t *testing.T) {
	v := NewVector(1, 8)
	divide := NewVector(2, 4)
	want := NewVector(0.5, 2)
	got := v.Clone()
	got.Divide(&divide)

	evaluateVector(fmt.Sprintf("divide %s by %s", v, divide), want, got, t)
}

func TestDivideScalar(t *testing.T) {
	v := NewVector(1, 2)
	divide := 2.0
	want := NewVector(0.5, 1.0)
	got := v.Clone()
	got.DivideScalar(divide)

	evaluateVector(fmt.Sprintf("divide %s by %.2f", v, divide), want, got, t)
}

func TestScale(t *testing.T) {
	v := NewVector(2, 3)
	want := NewVector(4, 6)
	got := v.Clone()
	got.Scale(2)

	evaluateVector(fmt.Sprintf("scale %s by 2", v), want, got, t)
}

func TestAbs(t *testing.T) {
	v := NewVector(-1, 1)
	want := NewVector(1, 1)
	got := v.Clone()
	got.Abs()

	evaluateVector(fmt.Sprintf("absolute value of %s", v), want, got, t)
}

func TestNegate(t *testing.T) {
	v := NewVector(-1, 1)
	want := NewVector(1, -1)
	got := v.Clone()
	got.Negate()

	evaluateVector(fmt.Sprintf("inverse value of %s", v), want, got, t)
}

func TestDistance(t *testing.T) {
	v := NewVector(1, 3)
	other := NewVector(1, -1)
	want := 4.0
	c := v.Clone()
	got := c.Distance(other)

	evaluateScalar(fmt.Sprintf("distance from %s to %s", v, other), want, got, t)
}

func TestLength(t *testing.T) {
	v := NewVector(2, 0)
	c := v.Clone()
	want := 2.0
	got := v.Length()

	d := fmt.Sprintf("length of %s", c)

	evaluateChanged(d, v, c, t)

	evaluateScalar(d, want, got, t)
}

func TestSetLength(t *testing.T) {
	v := NewVector(1, 1)
	c := v.Clone()
	want := 2.0
	got := v.SetLength(want).Length()

	d := fmt.Sprintf("length of %s", c)

	evaluateScalarApprox(d, want, got, t)
}

func TestLerp(t *testing.T) {
	a := NewVector(0, 0)
	b := NewVector(-20, 10)

	x := 0.3

	want := NewVector(-6, 3)
	got := a.Lerp(&b, x)

	evaluateVector(fmt.Sprintf("linear interpolation between %s and %s by %.2f", a, b, x), want, *got, t)
}

func TestDot(t *testing.T) {
	v := NewVector(1, 1)
	c := v.Clone()
	want := 2.0
	got := v.Dot(&v)

	d := fmt.Sprintf("dot product of %s", c)

	evaluateChanged(d, v, c, t)

	evaluateScalar(d, want, got, t)
}

func TestCross(t *testing.T) {
	v := NewVector(1, 1)

	clock := NewVector(1, 0)
	anti := NewVector(0, 1)

	want := -1.0
	got := v.Cross(clock)

	evaluateScalar(fmt.Sprintf("cross product of %s and %s", v, clock), want, got, t)

	want = 1.0
	got = v.Cross(anti)

	evaluateScalar(fmt.Sprintf("cross product of %s and %s", v, anti), want, got, t)
}

func TestNormalize(t *testing.T) {
	v := NewVector(10, 0)
	c := v.Clone()
	want := NewVector(1, 0)

	v.Normalize()

	evaluateVector(fmt.Sprintf("normalize %s", c), want, v, t)

	v = NewVector(0, 10)
	c = v.Clone()
	want = NewVector(0, 1)
	reverse := v.Normalize()

	evaluateVector(fmt.Sprintf("normalize %s", c), want, v, t)

	want = NewVector(0, 10)

	v.Scale(reverse)

	evaluateVector(fmt.Sprintf("reverse normalizing of %s", c), want, v, t)

	v = NewVector(0, 0)
	c = v.Clone()
	want = NewVector(0, 0)

	v.Normalize()

	evaluateVector(fmt.Sprintf("normalize zero vector should do nothing %s", c), want, v, t)
}

func TestAngle(t *testing.T) {
	v := NewVector(0, 1)
	c := v.Clone()
	other := NewVector(1, 0.3)

	d := fmt.Sprintf("angle from %s to %s", c, other)

	want := 1.2793395323170293
	got := v.Angle(other)

	evaluateScalar(d, want, got, t)
	evaluateChanged(d, v, c, t)

	other.Set(-1, 0.3)
	c = other.Clone()

	d = fmt.Sprintf("angle from %s to %s", c, other)

	got = v.Angle(other)

	evaluateScalar(d, want, got, t)
	evaluateChanged(d, other, c, t)
}

func TestSignedAngle(t *testing.T) {
	v := NewVector(0, 1)
	c := v.Clone()
	other := NewVector(1, 0.3)
	want := 1.2793395323170293
	got := v.SignedAngle(other)

	d := fmt.Sprintf("angle from %s to %s", v, other)
	evaluateScalar(d, want, got, t)
	evaluateChanged(d, v, c, t)

	other.Set(-1, 0.3)
	c = other.Clone()
	want = 5.0038457214660585
	got = v.SignedAngle(other)

	d = fmt.Sprintf("angle from %s to %s", v, other)
	evaluateScalar(d, want, got, t)
	evaluateChanged(d, other, c, t)
}

func TestReflect(t *testing.T) {
	rightDown := NewVector(1, -1)
	up := NewVector(0, 1)

	want := NewVector(1, 1)
	got := rightDown.Reflect(up)

	evaluateVector(fmt.Sprintf("reflect direction %s off surface with normal %s", rightDown, up), want, *got, t)
}

func TestReflectSurface(t *testing.T) {
	rightDown := NewVector(1, -1)
	up := NewVector(0, 1)

	want := NewVector(-1, -1)
	got := rightDown.ReflectSurface(up)

	evaluateVector(fmt.Sprintf("reflect direction %s off surface with normal %s", rightDown, up), want, *got, t)
}

func TestRotate(t *testing.T) {
	up := NewVector(0, 1)
	right := NewVector(1, 0)

	c := right.Clone()
	angle := -up.SignedAngle(right)
	want := NewVector(0, 1)
	got := right.Rotate(angle)

	evaluateVectorApprox(fmt.Sprintf("rotated %s by %.1f", c, angle*d2math.RadToDeg), want, *got, t)

	c = up.Clone()
	angle -= d2math.RadFull
	want = NewVector(-1, 0)
	got = up.Rotate(angle)

	evaluateVectorApprox(fmt.Sprintf("rotated %s by %.1f", c, angle*d2math.RadToDeg), want, *got, t)
}

func TestNinetyAnti(t *testing.T) {
	v := NewVector(0, 1)
	c := v.Clone()

	want := NewVector(-1, 0)
	got := v.NinetyAnti()

	evaluateVector(fmt.Sprintf("rotated %s by 90 degrees clockwise", c), want, *got, t)
}

func TestNinetyClock(t *testing.T) {
	v := NewVector(0, 1)
	c := v.Clone()

	want := NewVector(1, 0)
	v = c.Clone()
	got := v.NinetyClock()

	evaluateVector(fmt.Sprintf("rotated %s by 90 degrees anti-clockwise", c), want, *got, t)
}

func TestVectorUp(t *testing.T) {
	got := VectorUp()
	want := NewVector(0, 1)

	evaluateVector("create normalized vector with up direction", want, got, t)
}

func TestVectorDown(t *testing.T) {
	got := VectorDown()
	want := NewVector(0, -1)

	evaluateVector("create normalized vector with down direction", want, got, t)
}

func TestVectorRight(t *testing.T) {
	got := VectorRight()
	want := NewVector(1, 0)

	evaluateVector("create normalized vector with right direction", want, got, t)
}

func TestVectorLeft(t *testing.T) {
	got := VectorLeft()
	want := NewVector(-1, 0)

	evaluateVector("create normalized vector with left direction", want, got, t)
}

func TestVectorOne(t *testing.T) {
	got := VectorOne()
	want := NewVector(1, 1)

	evaluateVector("create vector with X and Y values of 1", want, got, t)
}

func TestVectorZero(t *testing.T) {
	got := VectorZero()
	want := NewVector(0, 0)

	evaluateVector("create vector with X and Y values of 0", want, got, t)
}
