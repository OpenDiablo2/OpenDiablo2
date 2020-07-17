package d2math

import (
	"testing"
)

//nolint:gochecknoglobals // These variables are assigned to in benchmark functions to avoid compiler optimisations
// lowering the runtime of the benchmark. See: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go (A note
// on compiler optimisations)
var (
	outFloat float64
	outBool  bool
	outInt   int
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

func BenchmarkEqualsApprox(b *testing.B) {
	x := 1.0
	y := 2.0

	for n := 0; n < b.N; n++ {
		outBool = EqualsApprox(x, y)
	}
}

func TestCompareApprox(t *testing.T) {
	subEpsilon := Epsilon / 3

	want := 0
	a, b := 1+subEpsilon, 1.0
	got := CompareApprox(a, b)

	if got != want {
		t.Errorf("compare %.2f and %.2f: wanted %d: got %d", a, b, want, got)
	}

	want = 1
	a, b = 2, 1.0
	got = CompareApprox(a, b)

	if got != want {
		t.Errorf("compare %.2f and %.2f: wanted %d: got %d", a, b, want, got)
	}

	want = -1
	a, b = -2, 1.0
	got = CompareApprox(a, b)

	if got != want {
		t.Errorf("compare %.2f and %.2f: wanted %d: got %d", a, b, want, got)
	}
}

func BenchmarkCompareApprox(b *testing.B) {
	x := 1.0
	y := 2.0

	for n := 0; n < b.N; n++ {
		outInt = CompareApprox(x, y)
	}
}

func TestAbs(t *testing.T) {
	want := 1.0
	x := -1.0
	got := Abs(x)

	if got != want {
		t.Errorf("absolute value of %.2f: want %.2f: got %.2f", x, want, got)
	}

	want = 1.0
	x = 1.0
	got = Abs(x)

	if got != want {
		t.Errorf("absolute value of %.2f: want %.2f: got %.2f", x, want, got)
	}
}

func BenchmarkAbs(b *testing.B) {
	x := -1.0

	for n := 0; n < b.N; n++ {
		outFloat = Abs(x)
	}
}

func TestClamp(t *testing.T) {
	want := 0.5
	a := 0.5
	got := Clamp(a, 0, 1)

	if got != want {
		t.Errorf("clamped %.2f between 0 and 1: wanted %.2f: got %.2f", a, want, got)
	}

	want = 0.0
	a = -1.0
	got = Clamp(a, 0, 1)

	if got != want {
		t.Errorf("clamped %.2f between 0 and 1: wanted %.2f: got %.2f", a, want, got)
	}

	want = 1.0
	a = 2.0
	got = Clamp(a, 0, 1)

	if got != want {
		t.Errorf("clamped %.2f between 0 and 1: wanted %.2f: got %.2f", a, want, got)
	}
}

func BenchmarkClamp(b *testing.B) {
	f := 0.5

	for n := 0; n < b.N; n++ {
		outFloat = Clamp(f, 0, 1)
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

func BenchmarkSign(b *testing.B) {
	f := 0.5

	for n := 0; n < b.N; n++ {
		outInt = Sign(f)
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

func BenchmarkLerp(b *testing.B) {
	x := 1.0
	y := 1000.0
	interp := 1.01

	for n := 0; n < b.N; n++ {
		outFloat = Lerp(x, y, interp)
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

func BenchmarkUnlerp(b *testing.B) {
	x := 1.0
	y := 2.0
	lerp := 1.5

	for n := 0; n < b.N; n++ {
		outFloat = Unlerp(x, y, lerp)
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

func BenchmarkWrapInt(b *testing.B) {
	x := 10
	y := 2

	for n := 0; n < b.N; n++ {
		outInt = WrapInt(x, y)
	}
}
