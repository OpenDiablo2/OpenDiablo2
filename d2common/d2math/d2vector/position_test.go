package d2vector

import (
	"math"
	"math/rand"
	"testing"
)

func TestNewPosition(t *testing.T) {
	const maxXY = 1000

	x, y := rand.Intn(maxXY), rand.Intn(maxXY) // nolint:gosec // just a test
	locX, locY := float64(x), float64(y)
	pos := NewPosition(locX, locY)

	// old coordinate values			Position equivalent
	locationX := locX                 // .SubWord().X()
	locationY := locY                 // .SubWord().Y()
	tileX := float64(x / 5)           // .Tile().X()
	tileY := float64(y / 5)           // .Tile().Y()
	subcellX := 1 + math.Mod(locX, 5) // .RenderOffset().X()
	subcellY := 1 + math.Mod(locY, 5) // .RenderOffset().Y()

	want := NewVector(tileX, tileY)
	got := pos.Tile()

	if !got.Equals(want) {
		t.Errorf("world position should match old value: got %s: want %s", got, want)
	}

	want = NewVector(subcellX, subcellY)
	got = pos.RenderOffset()

	if !got.Equals(want) {
		t.Errorf("render offset position should match old value: got %s: want %s", got, want)
	}

	want = NewVector(locationX, locationY)
	got = &pos.Vector

	if !got.Equals(want) {
		t.Errorf("sub tile position should match old value: got %s: want %s", got, want)
	}
}

func validate(description string, t *testing.T, original, got, want, unchanged Vector) {
	if !got.EqualsApprox(&want) {
		t.Errorf("%s: want %s: got %s", description, want, got)
	}

	if !original.EqualsApprox(&unchanged) {
		t.Errorf("Position value %s was incorrectly changed to %s when calling this method", unchanged, original)
	}
}

func TestPosition_World(t *testing.T) {
	p := NewPosition(5, 10)
	unchanged := p.Clone()
	got := p.World()
	want := NewVector(1, 2)

	validate("world position", t, p.Vector, *got, *want, *unchanged)
}

func TestPosition_Tile(t *testing.T) {
	p := NewPosition(23, 24)
	unchanged := p.Clone()
	got := p.Tile()
	want := NewVector(4, 4)

	validate("tile position", t, p.Vector, *got, *want, *unchanged)
}

func TestPosition_RenderOffset(t *testing.T) {
	p := NewPosition(12.1, 14.2)
	unchanged := p.Clone()
	got := p.RenderOffset()
	want := NewVector(3.1, 5.2)

	validate("offset from sub tile", t, p.Vector, *got, *want, *unchanged)
}
