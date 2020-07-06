package d2vector

import (
	"fmt"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

func TestEquals(t *testing.T) {
	for _, vector := range typeArray() {
		testEquals(vector, t)
	}
}

func TestEqualsF(t *testing.T) {
	for _, vector := range typeArray() {
		testEqualsF(vector, t)
	}
}

func TestCompareF(t *testing.T) {
	for _, vector := range typeArray() {
		testCompareF(vector, t)
	}
}

func TestSet(t *testing.T) {
	for _, vector := range typeArray() {
		testSet(vector, t)
	}
}

func TestClone(t *testing.T) {
	for _, vector := range typeArray() {
		testClone(vector, t)
	}
}

func TestFloor(t *testing.T) {
	for _, vector := range typeArray() {
		testFloor(vector, t)
	}
}

func TestAdd(t *testing.T) {
	for _, vector := range typeArray() {
		testAdd(vector, t)
	}
}

func TestSubtract(t *testing.T) {
	for _, vector := range typeArray() {
		testSubtract(vector, t)
	}
}

func TestMultiply(t *testing.T) {
	for _, vector := range typeArray() {
		testMultiply(vector, t)
	}
}

func TestDivide(t *testing.T) {
	for _, vector := range typeArray() {
		testDivide(vector, t)
	}
}

func TestDistance(t *testing.T) {
	for _, vector := range typeArray() {
		testDistance(vector, t)
	}
}

func TestScale(t *testing.T) {
	for _, vector := range typeArray() {
		testScale(vector, t)
	}
}

func TestAbs(t *testing.T) {
	for _, vector := range typeArray() {
		testAbs(vector, t)
	}
}

func TestNegate(t *testing.T) {
	for _, vector := range typeArray() {
		testNegate(vector, t)
	}
}

func typeArray() []func(x, y float64) d2interface.Vector {
	return []func(x, y float64) d2interface.Vector{
		NewFloat64,
		NewBigFloat,
	}
}

func evaluateVectors(description string, want, got d2interface.Vector, t *testing.T) {
	if !got.Equals(want) {
		t.Errorf("%s: wanted %s: got %s", description, want, got)
	}
}

func evaluateBool(description string, want, got bool, t *testing.T) {
	if want != got {
		t.Errorf("%s: wanted %t: got %t", description, want, got)
	}
}

func testEquals(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	a := vector(1, 2)
	b := vector(1, 2)

	got := a.Equals(b)

	evaluateBool(fmt.Sprintf("exact equality %s and %s", a, b), true, got, t)

	b.Set(3, 4)

	got = a.Equals(b)

	evaluateBool(fmt.Sprintf("exact equality %s and %s", a, b), false, got, t)
}

func testEqualsF(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	subEpsilon := epsilon / 3

	a := vector(1+subEpsilon, 2+subEpsilon)
	b := vector(1, 2)

	got := a.EqualsF(b)

	evaluateBool(fmt.Sprintf("approximate equality %s and %s", a, b), true, got, t)

	a.Add(vector(epsilon, epsilon))

	got = a.EqualsF(b)

	evaluateBool(fmt.Sprintf("approximate equality %s and %s", a, b), false, got, t)
}

func testCompareF(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	subEpsilon := epsilon / 3

	f := vector(1+subEpsilon, 1+subEpsilon)
	c := vector(1, 1)

	x, y := f.CompareF(c)

	got := x == 0 && y == 0

	evaluateBool(fmt.Sprintf("approximate comparison %s and %s", f, c), true, got, t)

	l := vector(2, 2)
	s := vector(-1, 3)

	x, y = l.CompareF(s)

	got = x == -1 && y == 1

	evaluateBool(fmt.Sprintf("approximate comparison %s and %s", l, s), true, got, t)

	e := vector(2, 2)
	q := vector(3, -1)

	x, y = e.CompareF(q)

	got = x == 1 && y == -1

	evaluateBool(fmt.Sprintf("approximate comparison %s and %s", e, q), true, got, t)
}

func testSet(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 1)
	want := vector(2, 3)
	got := v.Clone().Set(2, 3)

	evaluateVectors(fmt.Sprintf("set %s to (2, 3)", v), want, got, t)
}

func testClone(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	want := vector(1, 2)
	got := want.Clone()

	evaluateVectors(fmt.Sprintf("clone %s", want), want, got, t)
}

func testFloor(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1.6, 1.6)
	want := vector(1, 1)
	got := v.Clone().Floor()

	evaluateVectors(fmt.Sprintf("round %s down", v), want, got, t)
}

func testAdd(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 1)
	add := vector(0.5, 0.5)
	want := vector(1.5, 1.5)
	got := v.Clone().Add(add)

	evaluateVectors(fmt.Sprintf("add %s to %s", add, v), want, got, t)
}

func testSubtract(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 1)
	subtract := vector(0.6, 0.6)
	want := vector(0.4, 0.4)
	got := v.Clone().Subtract(subtract)

	evaluateVectors(fmt.Sprintf("subtract %s from %s", subtract, v), want, got, t)
}

func testMultiply(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 1)
	multiply := vector(2, 2)
	want := vector(2, 2)
	got := v.Clone().Multiply(multiply)

	evaluateVectors(fmt.Sprintf("multiply %s by %s", v, multiply), want, got, t)
}

func testDivide(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 1)
	divide := vector(2, 2)
	want := vector(0.5, 0.5)
	got := v.Clone().Divide(divide)

	evaluateVectors(fmt.Sprintf("divide %s by %s", v, divide), want, got, t)
}

func testDistance(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 3)
	other := vector(1, -1)
	got := v.Clone().Distance(other) == 4.0

	evaluateBool(fmt.Sprintf("distance from %s to %s", v, other), true, got, t)
}

func testScale(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(2, 3)
	want := vector(4, 6)
	got := v.Clone().Scale(2)

	evaluateVectors(fmt.Sprintf("scale %s by 2", v), want, got, t)
}

func testAbs(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(-1, 1)
	want := vector(1, 1)
	got := v.Clone().Abs()

	evaluateVectors(fmt.Sprintf("absolute value of %s", v), want, got, t)
}

func testNegate(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(-1, 1)
	want := vector(1, -1)
	got := v.Clone().Negate()

	evaluateVectors(fmt.Sprintf("inverse value of %s", v), want, got, t)
}
