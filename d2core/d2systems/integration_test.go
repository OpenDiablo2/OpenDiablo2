package d2systems

import (
	"testing"
	"time"

	"github.com/gravestench/ecs"
)

func Test_SystemIntegrationTest(t *testing.T) {
	cfg := ecs.NewWorldConfig()

	scale := NewTimeScaleSystem()
	movement := NewMovementSystem()

	cfg.With(scale)
	cfg.With(movement)

	world := ecs.NewWorld(cfg)

	e := world.NewEntity()
	pos := movement.positions.AddPosition(e)
	vel := movement.velocities.AddVelocity(e)

	vel.Set(1, 2)

	// first test without time scaling active
	scale.scale = 0.001
	scale.SetActive(false)

	timeDelta := time.Millisecond

	expectX, expectY := pos.X()+vel.X(), pos.Y()+vel.Y()

	for idx := 0; idx < 1000; idx++ {
		_ = world.Update(timeDelta)
	}

	if !pos.EqualsApprox(vel.Vector) {
		fmtStr := "position component not updated, expected (%v,%v) but got (%v,%v)"
		t.Errorf(fmtStr, expectX, expectY, pos.X(), pos.Y())
	}

	// now enable time scaling
	scale.SetActive(true)

	expectX, expectY = pos.X()+vel.X(), pos.Y()+vel.Y()

	for idx := 0; idx < 1000000; idx++ {
		_ = world.Update(timeDelta)
	}

	if pos.EqualsApprox(vel.Vector.Clone().Scale(2)) {
		fmtStr := "position component not updated, expected (%v,%v) but got (%v,%v)"
		t.Errorf(fmtStr, expectX, expectY, pos.X(), pos.Y())
	}
}
