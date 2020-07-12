package d2math

import (
	"testing"
)

func TestEqualsApprox(t *testing.T) {
	subEpsilon := Epsilon / 3

	a, b := 1+subEpsilon, 1.0
	got := EqualsApprox(a, b)

	if !got {
		t.Errorf("compare %.2f and %.2f: wanted %t: got %t", a, b, true, got)
	}

	a, b = 1+Epsilon, 1.0
	got = EqualsApprox(a, b)

	if !got {
		t.Errorf("compare %.2f and %.2f: wanted %t: got %t", a, b, false, got)
	}
}

func TestCompareFloat64Fuzzy(t *testing.T) {
	subEpsilon := Epsilon / 3

	want := 0
	a, b := 1+subEpsilon, 1.0
	got := CompareFloat64Fuzzy(a, b)

	if got != want {
		t.Errorf("compare %.2f and %.2f: wanted %d: got %d", a, b, want, got)
	}

	want = 1
	a, b = 2, 1.0
	got = CompareFloat64Fuzzy(a, b)

	if got != want {
		t.Errorf("compare %.2f and %.2f: wanted %d: got %d", a, b, want, got)
	}

	want = -1
	a, b = -2, 1.0
	got = CompareFloat64Fuzzy(a, b)

	if got != want {
		t.Errorf("compare %.2f and %.2f: wanted %d: got %d", a, b, want, got)
	}
}

func TestClampFloat64(t *testing.T) {
	want := 0.5
	a := 0.5
	got := ClampFloat64(a, 0, 1)

	if got != want {
		t.Errorf("clamped %.2f between 0 and 1: wanted %.2f: got %.2f", a, want, got)
	}

	want = 0.0
	a = -1.0
	got = ClampFloat64(a, 0, 1)

	if got != want {
		t.Errorf("clamped %.2f between 0 and 1: wanted %.2f: got %.2f", a, want, got)
	}

	want = 1.0
	a = 2.0
	got = ClampFloat64(a, 0, 1)

	if got != want {
		t.Errorf("clamped %.2f between 0 and 1: wanted %.2f: got %.2f", a, want, got)
	}
}

func TestSign(t *testing.T) {
	want := 1
	a := 0.5
	got := Sign(a)

	if got != want {
		t.Errorf("sign of %.2f: wanted %df: got %d", a, want, got)
	}

	want = -1
	a = -3
	got = Sign(a)

	if got != want {
		t.Errorf("sign of %.2f: wanted %df: got %d", a, want, got)
	}

	want = 0
	a = 0.0
	got = Sign(a)

	if got != want {
		t.Errorf("sign of %.2f: wanted %df: got %d", a, want, got)
	}
}

func TestLerp(t *testing.T) {
	want := 3.0
	x := 0.3
	a, b := 0.0, 10.0

	got := Lerp(a, b, x)

	d := "linear interpolation between %.2f and %.2f with interpolator %.2f: wanted %.2f: got %.2f"

	if got != want {
		t.Errorf(d, a, b, x, want, got)
	}
}

func TestUnlerp(t *testing.T) {
	want := 0.3
	x := 3.0
	a, b := 0.0, 10.0

	got := Unlerp(a, b, x)

	d := "find the interpolator of %.2f between %.2f and %.2f: wanted %.2f: got %.2f"

	if got != want {
		t.Errorf(d, x, a, b, want, got)
	}
}

func TestWrapInt(t *testing.T) {
	want := 50
	a, b := 1050, 100
	got := WrapInt(a, b)

	d := "wrap %d between 0 and %d: want %d: got %d"

	if got != want {
		t.Errorf(d, a, b, want, got)
	}

	want = 270
	a, b = -1170, 360
	got = WrapInt(a, b)

	if got != want {
		t.Errorf(d, a, b, want, got)
	}

	want = 270
	a, b = 270, 360
	got = WrapInt(a, b)

	if got != want {
		t.Errorf(d, a, b, want, got)
	}

	want = 90
	a, b = -270, 360
	got = WrapInt(a, b)

	if got != want {
		t.Errorf(d, a, b, want, got)
	}
}
