package d2systems

import (
	"testing"
	"time"

	"github.com/gravestench/akara"
)

func TestTimeScaleSystem_Init(t *testing.T) {
	cfg := akara.NewWorldConfig()

	cfg.With(&TimeScaleSystem{})

	world := akara.NewWorld(cfg)

	if len(world.Systems) != 1 {
		t.Error("system not added to the world")
	}
}

func TestTimeScaleSystem_Process(t *testing.T) {
	cfg := akara.NewWorldConfig()

	timescaleSystem := &TimeScaleSystem{}

	cfg.With(timescaleSystem)

	timescaleSystem.scale = 0.01

	world := akara.NewWorld(cfg)

	actual := time.Second
	expected := time.Duration(timescaleSystem.scale) * actual

	if err := world.Update(actual); err != nil {
		timescaleSystem.Error(err.Error())
	}

	if world.TimeDelta != expected {
		t.Error("world time delta not scaled")
	}
}
