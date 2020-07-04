package d2vector

import (
	"testing"
)

func TestTile(t *testing.T) {
	p := Position{New(1.6, 1.6)}

	got := p.Tile()

	want := New(1, 1)

	if !got.FuzzyEquals(want) {
		t.Errorf("want %s: got %s", want, got)
	}

	if !p.FuzzyEquals(New(1.6, 1.6)) {
		t.Errorf("value of p was changed to %s when calling p.Tile()", p)
	}
}

func TestTileOffset(t *testing.T) {
	p := Position{New(1.6, 1.6)}

	got := p.TileOffset()

	want := New(0.6, 0.6)

	if !got.FuzzyEquals(want) {
		t.Errorf("want %s: got %s", want, got)
	}

	if !p.FuzzyEquals(New(1.6, 1.6)) {
		t.Errorf("value of p was changed to %s when calling p.Tile()", p)
	}
}

func TestSubTile(t *testing.T) {
	p := Position{New(1, 1)}

	got := p.SubTile()

	want := New(5, 5)

	if !got.FuzzyEquals(want) {
		t.Errorf("want %s: got %s", want, got)
	}

	if !p.FuzzyEquals(New(1, 1)) {
		t.Errorf("value of p was changed to %s when calling p.Tile()", p)
	}
}

func TestSubTileOffset(t *testing.T) {
	p := Position{New(1.1, 1.1)}

	got := p.SubTileOffset()

	want := New(0.5, 0.5)

	if !got.FuzzyEquals(want) {
		t.Errorf("want %s: got %s", want, got)
	}

	if !p.FuzzyEquals(New(1.1, 1.1)) {
		t.Errorf("value of p was changed to %s when calling p.Tile()", p)
	}
}
