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

func TestCompareF(t *testing.T) {
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

func TestAdd(t *testing.T) {
	v := NewVector(1, 1)
	add := NewVector(0.5, 0.5)
	want := NewVector(1.5, 1.5)
	got := v.Clone()
	got.Add(&add)

	evaluateVector(fmt.Sprintf("add %s to %s", add, v), want, got, t)
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
	v := NewVector(1, 1)
	divide := NewVector(2, 2)
	want := NewVector(0.5, 0.5)
	got := v.Clone()
	got.Divide(&divide)

	evaluateVector(fmt.Sprintf("divide %s by %s", v, divide), want, got, t)
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

func TestDot(t *testing.T) {
	v := NewVector(1, 1)
	c := v.Clone()
	want := 2.0
	got := v.Dot(&v)

	d := fmt.Sprintf("dot product of %s", c)

	evaluateChanged(d, v, c, t)

	evaluateScalar(d, want, got, t)
}

func TestNormalize(t *testing.T) {
	v := NewVector(10, 0)
	c := v.Clone()
	want := NewVector(1, 0)
	got := v.Normalize()

	evaluateVector(fmt.Sprintf("normalize %s", c), want, *got, t)

	v = NewVector(0, 10)
	c = v.Clone()
	want = NewVector(0, 1)
	got = v.Normalize()

	evaluateVector(fmt.Sprintf("normalize %s", c), want, *got, t)
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
