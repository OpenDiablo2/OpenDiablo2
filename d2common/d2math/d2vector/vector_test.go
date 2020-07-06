package d2vector

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

func TestEquals(t *testing.T) {
	testEquals(NewFloat64, t)
	testEquals(NewBigFloat, t)
}

func TestEqualsF(t *testing.T) {
	testEqualsF(NewFloat64, t)
	testEqualsF(NewBigFloat, t)
}

func TestCompareF(t *testing.T) {
	testCompareF(NewFloat64, t)
	testCompareF(NewBigFloat, t)
}

func TestSet(t *testing.T) {
	testSet(NewFloat64, t)
	testSet(NewBigFloat, t)
}

func TestClone(t *testing.T) {
	testClone(NewFloat64, t)
	testClone(NewBigFloat, t)
}

func TestFloor(t *testing.T) {
	testFloor(NewFloat64, t)
	testFloor(NewBigFloat, t)
}

func TestAdd(t *testing.T) {
	testAdd(NewFloat64, t)
	testAdd(NewBigFloat, t)
}

func TestSubtract(t *testing.T) {
	testSubtract(NewFloat64, t)
	testSubtract(NewBigFloat, t)
}

func TestMultiply(t *testing.T) {
	testMultiply(NewFloat64, t)
	testMultiply(NewBigFloat, t)
}

func TestDivide(t *testing.T) {
	testDivide(NewFloat64, t)
	testDivide(NewBigFloat, t)
}

func TestDistance(t *testing.T) {
	testDistance(NewFloat64, t)
	testDistance(NewBigFloat, t)
}

func TestScale(t *testing.T) {
	testScale(NewFloat64, t)
	testScale(NewBigFloat, t)
}

func TestAbs(t *testing.T) {
	testAbs(NewFloat64, t)
	testAbs(NewBigFloat, t)
}

func TestNegate(t *testing.T) {
	testNegate(NewFloat64, t)
	testNegate(NewBigFloat, t)
}

func testEquals(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	a := vector(1, 2)
	b := vector(1, 2)

	if !a.Equals(b) {
		t.Errorf("vectors %s and %s were not exactly equal", a, b)
	}

	b.Set(3, 4)

	if a.Equals(b) {
		t.Errorf("vectors %s and %s were exactly equal", a, b)
	}
}

func testEqualsF(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	subEpsilon := epsilon / 3

	a := vector(1+subEpsilon, 2+subEpsilon)
	b := vector(1, 2)

	if !a.EqualsF(b) {
		t.Errorf("vectors %s and %s were not almost equal", a, b)
	}

	a.Add(vector(epsilon, epsilon))

	if a.Equals(b) {
		t.Errorf("vectors %s and %s were almost equal", a, b)
	}
}

func testCompareF(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	subEpsilon := epsilon / 3

	f := vector(1+subEpsilon, 1+subEpsilon)
	c := vector(1, 1)

	if x, y := f.CompareF(c); (x + y) != 0 {
		t.Errorf("call to %s.Compare(%s) returned comparison (%d, %d)", f, c, x, y)
	}

	l := vector(2, 2)
	s := vector(-1, -1)

	if x, y := l.CompareF(s); x != 1 || y != 1 {
		t.Errorf("call to %s.Compare(%s) returned comparison (%d, %d)", l, s, x, y)
	}

	e := vector(2, 2)
	q := vector(2, 2)

	if x, y := e.CompareF(q); x != 0 || y != 0 {
		t.Errorf("call to %s.Compare(%s) returned comparison (%d, %d)", e, q, x, y)
	}
}

func testSet(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	got := vector(1, 1)
	want := vector(2, 3)
	got.Add(vector(1, 2))

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func testClone(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	want := vector(1, 2)
	got := want.Clone()

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func testFloor(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1.6, 1.6)
	want := vector(1, 1)

	if !v.Floor().Equals(want) {
		t.Errorf("want %s: got %s", want, v)
	}
}

func testAdd(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 1)
	want := vector(1.5, 1.5)
	got := v.Add(vector(0.5, 0.5))

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func testSubtract(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 1)
	want := vector(0.4, 0.4)
	got := v.Subtract(vector(0.6, 0.6))

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func testMultiply(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 1)
	want := vector(2, 2)
	got := v.Multiply(vector(2, 2))

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func testDivide(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 1)
	want := vector(0.5, 0.5)
	got := v.Divide(vector(2, 2))

	if !got.Equals(want) {
		t.Errorf("divide %s by (2,2): wanted %s: got %s", v, want, got)
	}
}

func testDistance(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(1, 3)
	d := vector(1, -1)
	want := 4.0
	got := v.Distance(d)

	if got != want {
		t.Errorf("distance from %s to %s: wanted %f: got %f", v, d, want, got)
	}
}

func testScale(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(2, 3)
	want := vector(4, 6)
	got := v.Scale(2)

	if !got.Equals(want) {
		t.Errorf("scale (2, 3) by 2: wanted %s: got %s", want, got)
	}
}

func testAbs(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(-1, 1)
	want := vector(1, 1)
	got := v.Abs()

	if !got.Equals(want) {
		t.Errorf("absolute value of (-1, 1): wanted %s: got %s", want, got)
	}
}

func testNegate(vector func(float64, float64) d2interface.Vector, t *testing.T) {
	v := vector(-1, 1)
	want := vector(1, -1)
	got := v.Negate()

	if !got.Equals(want) {
		t.Errorf("negated value of (-1, 1): wanted %s: got %s", want, got)
	}
}
