package d2vector

import "testing"

func TestClone(t *testing.T) {
	v := NewBigFloat(1, 1)
	want := NewBigFloat(1, 1)
	got := v.Clone()

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func TestAbs(t *testing.T) {
	v := NewBigFloat(-1, -1)
	want := NewBigFloat(1, 1)
	got := v.Abs()

	if !got.Equals(want) {
		t.Errorf("wanted %s: got %s", want, got)
	}
}

func TestFloor(t *testing.T) {
	v := NewBigFloat(1.6, 1.6)

	want := NewBigFloat(1, 1)

	if !v.Floor().Equals(want) {
		t.Errorf("want %s: got %s", want, v)
	}
}
