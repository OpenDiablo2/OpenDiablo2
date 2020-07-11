package d2vector

import (
	"math"
	"math/rand"
	"testing"
)

func TestEntityPosition(t *testing.T) {
	x, y := rand.Intn(1000), rand.Intn(1000)
	pos := EntityPosition(x, y)
	locX, locY := float64(x), float64(y)

	// old coordinate values			Position equivalent
	locationX := locX                 // .SubWord().X()
	locationY := locY                 // .SubWord().Y()
	tileX := float64(x / 5)           // .Tile().X()
	tileY := float64(y / 5)           // .Tile().Y()
	subcellX := 1 + math.Mod(locX, 5) // .SubCell().X()
	subcellY := 1 + math.Mod(locY, 5) // .SubCell().Y()

	want := NewVector(tileX, tileY)
	got := pos.Tile()

	if !got.Equals(want) {
		t.Errorf("world position should match old value: got %s: want %s", got, want)
	}

	want = NewVector(subcellX, subcellY)
	got = pos.SubCell()

	if !got.Equals(want) {
		t.Errorf("sub cell position should match old value: got %s: want %s", got, want)
	}

	want = NewVector(locationX, locationY)
	got = pos.SubWorld()

	if !got.Equals(want) {
		t.Errorf("sub tile position should match old value: got %s: want %s", got, want)
	}
}

func validate(description string, t *testing.T, original, got, want, unchanged Vector) {
	if !got.EqualsApprox(want) {
		t.Errorf("%s: want %s: got %s", description, want, got)
	}

	if !original.EqualsApprox(unchanged) {
		t.Errorf("Position value %s was incorrectly changed to %s when calling this method", unchanged, original)
	}
}

func TestTile(t *testing.T) {
	p := NewPosition(1.6, 1.6)
	got := p.Tile()
	want := NewVector(1, 1)
	unchanged := NewVector(1.6, 1.6)

	validate("tile position", t, p.Vector, *got, want, unchanged)
}

func TestTileOffset(t *testing.T) {
	p := NewPosition(1.6, 1.6)
	got := p.TileOffset()
	want := NewVector(0.6, 0.6)
	unchanged := NewVector(1.6, 1.6)

	validate("tile offset", t, p.Vector, *got, want, unchanged)
}

func TestSubWorld(t *testing.T) {
	p := NewPosition(1, 1)
	got := p.SubWorld()
	want := NewVector(5, 5)
	unchanged := NewVector(1, 1)

	validate("sub tile world position", t, p.Vector, *got, want, unchanged)
}

func TestSubTile(t *testing.T) {
	p := NewPosition(1, 1)
	got := p.SubTile()
	want := NewVector(5, 5)
	unchanged := NewVector(1, 1)

	validate("sub tile with offset", t, p.Vector, *got, want, unchanged)
}

func TestSubTileOffset(t *testing.T) {
	p := NewPosition(1.1, 1.1)
	got := p.SubTileOffset()
	want := NewVector(0.5, 0.5)
	unchanged := NewVector(1.1, 1.1)

	validate("offset from sub tile", t, p.Vector, *got, want, unchanged)
}
