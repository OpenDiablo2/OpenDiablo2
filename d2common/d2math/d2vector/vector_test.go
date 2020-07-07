package d2vector

import (
	"fmt"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

func evaluateVectors(description string, want, got Vector, t *testing.T) {
	if !got.Equals(want) {
		t.Errorf("%s: wanted %s: got %s", description, want, got)
	}
}

func evaluateScalar(description string, want, got float64, t *testing.T) {
	if want != got {
		t.Errorf("%s: wanted %f: got %f", description, want, got)
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

	evaluateVectors(fmt.Sprintf("set %s to (2, 3)", v), want, got, t)
}

func TestClone(t *testing.T) {
	want := NewVector(1, 2)
	got := want.Clone()

	evaluateVectors(fmt.Sprintf("clone %s", want), want, got, t)
}

func TestFloor(t *testing.T) {
	v := NewVector(1.6, 1.6)
	want := NewVector(1, 1)
	got := v.Clone()
	got.Floor()

	evaluateVectors(fmt.Sprintf("round %s down", v), want, got, t)
}

func TestAdd(t *testing.T) {
	v := NewVector(1, 1)
	add := NewVector(0.5, 0.5)
	want := NewVector(1.5, 1.5)
	got := v.Clone()
	got.Add(&add)

	evaluateVectors(fmt.Sprintf("add %s to %s", add, v), want, got, t)
}

func TestSubtract(t *testing.T) {
	v := NewVector(1, 1)
	subtract := NewVector(0.6, 0.6)
	want := NewVector(0.4, 0.4)
	got := v.Clone()
	got.Subtract(&subtract)

	evaluateVectors(fmt.Sprintf("subtract %s from %s", subtract, v), want, got, t)
}

func TestMultiply(t *testing.T) {
	v := NewVector(1, 1)
	multiply := NewVector(2, 2)
	want := NewVector(2, 2)
	got := v.Clone()
	got.Multiply(&multiply)

	evaluateVectors(fmt.Sprintf("multiply %s by %s", v, multiply), want, got, t)
}

func TestDivide(t *testing.T) {
	v := NewVector(1, 1)
	divide := NewVector(2, 2)
	want := NewVector(0.5, 0.5)
	got := v.Clone()
	got.Divide(&divide)

	evaluateVectors(fmt.Sprintf("divide %s by %s", v, divide), want, got, t)
}

func TestScale(t *testing.T) {
	v := NewVector(2, 3)
	want := NewVector(4, 6)
	got := v.Clone()
	got.Scale(2)

	evaluateVectors(fmt.Sprintf("scale %s by 2", v), want, got, t)
}

func TestAbs(t *testing.T) {
	v := NewVector(-1, 1)
	want := NewVector(1, 1)
	got := v.Clone()
	got.Abs()

	evaluateVectors(fmt.Sprintf("absolute value of %s", v), want, got, t)
}

func TestNegate(t *testing.T) {
	v := NewVector(-1, 1)
	want := NewVector(1, -1)
	got := v.Clone()
	got.Negate()

	evaluateVectors(fmt.Sprintf("inverse value of %s", v), want, got, t)
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
	want := 2.0
	got := v.Length()

	evaluateScalar(fmt.Sprintf("length of %s", v), want, got, t)
}

func TestDot(t *testing.T) {
	v := NewVector(1, 1)
	want := 2.0
	got := v.Dot(&v)
	evaluateScalar(fmt.Sprintf("dot product of %s", v), want, got, t)
}

func TestAngle(t *testing.T) {
	v := NewVector(-1, 0)
	want := 45.0
	got := v.Angle()

	fmt.Println(want, got)
}
