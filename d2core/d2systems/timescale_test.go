package d2systems

import (
	"testing"
	"time"

	"github.com/gravestench/ecs"
)

func TestTimeScaleSystem_Init(t *testing.T) {
	cfg := ecs.NewWorldConfig()

	cfg.With(NewTimeScaleSystem())

	world := ecs.NewWorld(cfg)

	if len(world.Systems) != 1 {
		t.Error("system not added to the world")
	}
}

func TestTimeScaleSystem_Process(t *testing.T) {
	cfg := ecs.NewWorldConfig()

	timescaleSystem := NewTimeScaleSystem()

	cfg.With(timescaleSystem)

	timescaleSystem.scale = 0.01

	world := ecs.NewWorld(cfg)

	actual := time.Second
	expected := time.Duration(timescaleSystem.scale) * actual

	world.Update(actual)

	if world.TimeDelta != expected {
		t.Error("world time delta not scaled")
	}
}
