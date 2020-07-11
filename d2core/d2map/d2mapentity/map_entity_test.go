package d2mapentity

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2astar"
)

var stepEntity mapEntity

const (
	normalTickTime float64 = 0.05
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	stepEntity = setupMovement()
}

func tickTime() float64 {
	max := 0.06
	min := 0.04
	return rand.Float64() * (max - min)
}

func entity() mapEntity {
	return newMapEntity(10, 10)
}

func path(length int, origin d2vector.Vector) []d2astar.Pather {
	path := make([]d2astar.Pather, length)

	x, y := origin.X(), origin.Y()

	for i := 0; i < length; i++ {
		path[i] = pathTile(x/5, (float64(i)+y+1)/5)
	}

	return path
}

func pathTile(x, y float64) *d2common.PathTile {
	return &d2common.PathTile{X: x, Y: y}
}

func setupMovement() mapEntity {
	e := entity()
	e.SetSpeed(12)
	newPath := path(100, e.Position.Vector)
	e.SetPath(newPath, func() {})

	return e
}

func TestMapEntity_Step(t *testing.T) {
	stepCount := 5
	e := setupMovement()
	start := e.Position.Clone()
	//fmt.Printf("start\t%s\t%d\t%s\n", e.Position, len(e.path), e.Target)

	for i := 0; i < stepCount; i++ {
		e.Step(normalTickTime)
		//fmt.Printf("move\t%s\t%d\t%s\n", e.Position, len(e.path), e.Target)
	}

	change := e.velocity(0.05)
	fmt.Println(change)
	change.Scale(float64(stepCount - 1))

	want := change.Add(&start)

	if !e.Position.EqualsApprox(*want) {
		t.Errorf("entity position after 5 steps: want %s: got %s", want, e.Position.Vector)
	}
}

func BenchmarkMapEntity_Step(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stepEntity.Step(0.05)
	}
}
