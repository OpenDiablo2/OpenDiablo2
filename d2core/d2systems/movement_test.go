package d2systems

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

func TestMovementSystem_Init(t *testing.T) {
	cfg := akara.NewWorldConfig()

	cfg.With(NewMovementSystem())

	world := akara.NewWorld(cfg)

	if len(world.Systems) != 1 {
		t.Error("system not added to the world")
	}
}

func TestMovementSystem_Active(t *testing.T) {
	movement := NewMovementSystem()

	if movement.Active() {
		t.Error("system should not be active at creation")
	}
}

func TestMovementSystem_SetActive(t *testing.T) {
	movement := NewMovementSystem()

	movement.SetActive(false)

	if movement.Active() {
		t.Error("system should be inactive after being set inactive")
	}
}

func TestMovementSystem_EntityAdded(t *testing.T) {
	cfg := akara.NewWorldConfig()

	movement := NewMovementSystem()

	cfg.With(movement).
		With(d2components.NewPositionMap()).
		With(d2components.NewVelocityMap())

	world := akara.NewWorld(cfg)

	e := world.NewEntity()

	position := movement.positions.AddPosition(e)
	velocity := movement.velocities.AddVelocity(e)

	px, py := 10., 10.
	vx, vy := 1., 0.

	position.Set(px, py)
	velocity.Set(vx, vy)

	if len(movement.Subscriptions[0].GetEntities()) != 1 {
		t.Error("entity not added to the system")
	}

	if p, found := movement.positions.GetPosition(e); !found {
		t.Error("position component not found")
	} else if p.X() != px || p.Y() != py {
		fmtError := "position component values incorrect:\n\t expected %v, %v but got %v, %v"
		t.Errorf(fmtError, px, py, p.X(), p.Y())
	}

	if v, found := movement.velocities.GetVelocity(e); !found {
		t.Error("position component not found")
	} else if v.X() != vx || v.Y() != vy {
		fmtError := "velocity component values incorrect:\n\t expected %v, %v but got %v, %v"
		t.Errorf(fmtError, px, py, v.X(), v.Y())
	}
}

func TestMovementSystem_Update(t *testing.T) {
	// world bootstrap
	cfg := akara.NewWorldConfig()

	movementSystem := NewMovementSystem()
	positions := d2components.NewPositionMap()
	velocities := d2components.NewVelocityMap()

	cfg.With(movementSystem).With(positions).With(velocities)

	world := akara.NewWorld(cfg)

	// lets make an entity and add some components to it
	e := world.NewEntity()
	position := movementSystem.positions.AddPosition(e)
	velocity := movementSystem.velocities.AddVelocity(e)

	px, py := 10., 10.
	vx, vy := 1., -1.

	// mutate the components a bit
	position.Set(px, py)
	velocity.Set(vx, vy)

	// should apply the velocity to the position
	_ = world.Update(time.Second)

	if position.X() != px+vx || position.Y() != py+vy {
		fmtError := "expected position (%v, %v) but got (%v, %v)"
		t.Errorf(fmtError, px+vx, py+vy, position.X(), position.Y())
	}
}

func bench_N_entities(n int, b *testing.B) {
	cfg := akara.NewWorldConfig()

	movementSystem := NewMovementSystem()

	cfg.With(movementSystem)

	world := akara.NewWorld(cfg)

	for idx := 0; idx < n; idx++ {
		e := world.NewEntity()
		p := movementSystem.positions.AddPosition(e)
		v := movementSystem.velocities.AddVelocity(e)

		p.Set(0, 0)
		v.Set(rand.Float64(), rand.Float64())
	}

	benchName := strconv.Itoa(n) + "_entity update"
	b.Run(benchName, func(b *testing.B) {
		for idx := 0; idx < b.N; idx++ {
			_ = world.Update(time.Millisecond)
		}
	})

	fmt.Println("done!")
}

func BenchmarkMovementSystem_Update(b *testing.B) {
	bench_N_entities(1e1, b)
	bench_N_entities(1e2, b)
	bench_N_entities(1e3, b)
	bench_N_entities(1e4, b)
	bench_N_entities(1e5, b)
	bench_N_entities(1e6, b)
}
