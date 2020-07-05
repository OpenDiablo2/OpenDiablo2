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

func testEquals(fn func(float64, float64) d2interface.Vector, t *testing.T) {
	a := fn(1, 2)
	b := fn(1, 2)

	if !a.Equals(b) {
		t.Errorf("vectors %s and %s were not exactly equal", a, b)
	}

	b.Set(3, 4)

	if a.Equals(b) {
		t.Errorf("vectors %s and %s were exactly equal", a, b)
	}
}

func testEqualsF(fn func(float64, float64) d2interface.Vector, t *testing.T) {
	subEpsilon := epsilon / 3

	a := fn(1+subEpsilon, 2+subEpsilon)
	b := fn(1, 2)

	if !a.EqualsF(b) {
		t.Errorf("vectors %s and %s were not almost equal", a, b)
	}

	a.Add(fn(epsilon, epsilon))

	if a.Equals(b) {
		t.Errorf("vectors %s and %s were almost equal", a, b)
	}
}

func testCompareF(fn func(float64, float64) d2interface.Vector, t *testing.T) {
	subEpsilon := epsilon / 3

	f := fn(1+subEpsilon, 1+subEpsilon)
	c := fn(1, 1)

	if x, y := f.CompareF(c); (x + y) != 0 {
		t.Errorf("call to %s.Compare(%s) returned comparison (%d, %d)", f, c, x, y)
	}

	l := fn(2, 2)
	s := fn(-1, -1)

	if x, y := l.CompareF(s); x != 1 || y != 1 {
		t.Errorf("call to %s.Compare(%s) returned comparison (%d, %d)", l, s, x, y)
	}

	e := fn(2, 2)
	q := fn(2, 2)

	if x, y := e.CompareF(q); x != 0 || y != 0 {
		t.Errorf("call to %s.Compare(%s) returned comparison (%d, %d)", e, q, x, y)
	}
}

func testSet(fn func(float64, float64) d2interface.Vector, t *testing.T) {
	got := fn(1, 1)
	want := fn(2, 3)
	got.Add(fn(1, 2))

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func testClone(fn func(float64, float64) d2interface.Vector, t *testing.T) {
	want := fn(1, 2)
	got := want.Clone()

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func testFloor(fn func(float64, float64) d2interface.Vector, t *testing.T) {
	v := fn(1.6, 1.6)

	want := fn(1, 1)

	if !v.Floor().Equals(want) {
		t.Errorf("want %s: got %s", want, v)
	}
}
