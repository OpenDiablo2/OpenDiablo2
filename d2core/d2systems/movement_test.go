package d2systems

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gravestench/akara"
)

func TestMovementSystem_Init(t *testing.T) {
	cfg := akara.NewWorldConfig()

	cfg.With(&MovementSystem{})

	world := akara.NewWorld(cfg)

	if len(world.Systems) != 1 {
		t.Error("system not added to the world")
	}
}

func TestMovementSystem_Active(t *testing.T) {
	sys := &MovementSystem{}

	if sys.Active() {
		t.Error("system should not be active at creation")
	}
}

func TestMovementSystem_SetActive(t *testing.T) {
	sys := &MovementSystem{}

	sys.SetActive(false)

	if sys.Active() {
		t.Error("system should be inactive after being set inactive")
	}
}

func TestMovementSystem_EntityAdded(t *testing.T) {
	moveSys := &MovementSystem{}
	cfg := akara.NewWorldConfig().With(moveSys)
	world := akara.NewWorld(cfg)

	e := world.NewEntity()
	trs := moveSys.Components.Transform.Add(e)
	velocity := moveSys.Components.Velocity.Add(e)

	px, py := 10., 10.
	vx, vy := 1., 0.

	trs.Translation.X, trs.Translation.Y = px, py
	velocity.X, velocity.Y = vx, vy

	if len(moveSys.movableEntities.GetEntities()) != 1 {
		t.Error("entity not added to the system")
	}

	if trsGot, found := moveSys.Components.Transform.Get(e); !found {
		t.Error("position component not found")
	} else if trsGot.Translation.X != px || trsGot.Translation.Y != py {
		fmtError := "position component values incorrect:\n\t expected %v, %v but got %v, %v"
		t.Errorf(fmtError, px, py, trsGot.Translation.X, trsGot.Translation.Y)
	}

	if v, found := moveSys.Components.Velocity.Get(e); !found {
		t.Error("position component not found")
	} else if v.X != vx || v.Y != vy {
		fmtError := "velocity component values incorrect:\n\t expected %v, %v but got %v, %v"
		t.Errorf(fmtError, px, py, v.X, v.Y)
	}
}

func TestMovementSystem_Update(t *testing.T) {
	// world configFileBootstrap
	cfg := akara.NewWorldConfig()

	movementSystem := &MovementSystem{}

	cfg.With(movementSystem)
	world := akara.NewWorld(cfg)

	// lets make an entity and add some components to it
	e := world.NewEntity()
	trs := movementSystem.Components.Transform.Add(e)
	velocity := movementSystem.Components.Velocity.Add(e)

	px, py := 10., 10.
	vx, vy := 1., -1.

	// mutate the components a bit
	trs.Translation.X, trs.Translation.Y = px, py
	velocity.X, velocity.Y = vx, vy

	// should apply the velocity to the position
	_ = world.Update(time.Second)

	if trs.Translation.X != px+vx || trs.Translation.Y != py+vy {
		fmtError := "expected position (%v, %v) but got (%v, %v)"
		t.Errorf(fmtError, px+vx, py+vy, trs.Translation.X, trs.Translation.Y)
	}
}

func benchN(n int, b *testing.B) {
	cfg := akara.NewWorldConfig()

	movementSystem := &MovementSystem{}

	cfg.With(movementSystem)

	world := akara.NewWorld(cfg)

	for idx := 0; idx < n; idx++ {
		e := world.NewEntity()
		trs := movementSystem.Components.Transform.Add(e)
		v := movementSystem.Components.Velocity.Add(e)

		trs.Translation.X, trs.Translation.Y = 0, 0
		v.X, v.Y = rand.Float64(), rand.Float64() //nolint:gosec // it's just a test
	}

	benchName := strconv.Itoa(n) + "_entity update"
	b.Run(benchName, func(b *testing.B) {
		for idx := 0; idx < b.N; idx++ {
			_ = world.Update(time.Millisecond)
		}
	})
}

func BenchmarkMovementSystem_Update(b *testing.B) {
	benchN(1e1, b)
	benchN(1e2, b)
	benchN(1e3, b)
	benchN(1e4, b)
	benchN(1e5, b)
	benchN(1e6, b)
}
