package d2vector

import "testing"

func TestClone(t *testing.T) {
	v := New(1, 1)
	want := New(1, 1)
	got := v.Clone()

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func TestAbs(t *testing.T) {
	v := New(-1, -1)
	want := New(1, 1)
	got := v.Abs()

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}
