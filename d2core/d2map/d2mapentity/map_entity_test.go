package d2mapentity

import (
	"os"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

const (
	normalTickTime float64 = 0.05
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func entity() mapEntity {
	return newMapEntity(10, 10)
}

func movingEntity() mapEntity {
	e := entity()
	e.SetSpeed(9)
	newPath := path(10, e.Position)
	e.SetPath(newPath, func() {})

	return e
}

func path(length int, origin d2vector.Position) []d2vector.Position {
	path := make([]d2vector.Position, length)

	for i := 0; i < length; i++ {
		origin.AddScalar(float64(i+1) / 5)
		tile := origin.World()
		path[i] = pathTile(tile.X(), tile.Y())
	}

	return path
}

func pathTile(x, y float64) d2vector.Position {
	return d2vector.NewPositionTile(x, y)
}

func TestMapEntity_Step(t *testing.T) {
	stepCount := 10
	e := movingEntity()
	start := e.Position.Clone()

	for i := 0; i < stepCount; i++ {
		e.Step(normalTickTime)
	}

	// velocity
	change := d2vector.NewVector(0, 0)
	change.Copy(&e.Target.Vector)
	change.Subtract(&e.Position.Vector)
	change.SetLength(e.Speed * normalTickTime)
	// change in position
	change.Scale(float64(stepCount))

	want := change.Add(start)

	if !e.Position.EqualsApprox(want) {
		t.Errorf("entity position after %d steps: want %s: got %s", stepCount, want, e.Position.Vector)
	}

	if e.Position.Equals(start) {
		t.Errorf("entity did not move, still at position %s", start)
	}
}

func BenchmarkMapEntity_Step(b *testing.B) {
	stepEntity := movingEntity()

	for n := 0; n < b.N; n++ {
		stepEntity.Step(normalTickTime)
	}
}
