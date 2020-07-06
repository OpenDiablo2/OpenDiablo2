package d2vector

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

func validate(t *testing.T, original, got, want, unchanged d2interface.Vector) {
	if !got.EqualsApprox(want) {
		t.Errorf("want %s: got %s", want, got)
	}

	if !original.EqualsApprox(unchanged) {
		t.Errorf("Position value %s was incorrectly changed to %s when calling this method", unchanged, original)
	}
}

func TestTile(t *testing.T) {
	p := NewPosition(1.6, 1.6)
	got := p.Tile()
	want := NewBigFloat(1, 1)
	unchanged := NewBigFloat(1.6, 1.6)

	validate(t, p, got, want, unchanged)
}

func TestTileOffset(t *testing.T) {
	p := NewPosition(1.6, 1.6)
	got := p.TileOffset()
	want := NewBigFloat(0.6, 0.6)
	unchanged := NewBigFloat(1.6, 1.6)

	validate(t, p, got, want, unchanged)
}

func TestSubWorld(t *testing.T) {
	p := NewPosition(1, 1)
	got := p.SubWorld()
	want := NewBigFloat(5, 5)
	unchanged := NewBigFloat(1, 1)

	validate(t, p, got, want, unchanged)
}

func TestSubTile(t *testing.T) {
	p := NewPosition(1, 1)
	got := p.SubTile()
	want := NewBigFloat(5, 5)
	unchanged := NewBigFloat(1, 1)

	validate(t, p, got, want, unchanged)
}

func TestSubTileOffset(t *testing.T) {
	p := NewPosition(1.1, 1.1)
	got := p.SubTileOffset()
	want := NewBigFloat(0.5, 0.5)
	unchanged := NewBigFloat(1.1, 1.1)

	validate(t, p, got, want, unchanged)
}
