package d2systems

import (
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	logPrefixMovementSystem = "Movement System"
)

// static check that MovementSystem implements the System interface
var _ akara.System = &MovementSystem{}

// MovementSystem handles entity movement based on velocity and position components
type MovementSystem struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	d2components.TransformFactory
	d2components.VelocityFactory
	movableEntities *akara.Subscription
}

// Init initializes the system with the given world
func (m *MovementSystem) Init(world *akara.World) {
	m.World = world

	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixMovementSystem)

	m.Info("initializing ...")

	m.InjectComponent(&d2components.Transform{}, &m.Transform)
	m.InjectComponent(&d2components.Velocity{}, &m.Velocity)

	movable := m.NewComponentFilter().Require(
		&d2components.Transform{},
		&d2components.Velocity{},
	).Build()

	m.movableEntities = m.AddSubscription(movable)
}

// Update positions of all entities with their velocities
func (m *MovementSystem) Update() {
	entities := m.movableEntities.GetEntities()

	for entIdx := range entities {
		m.move(entities[entIdx])
	}
}

func (m *MovementSystem) move(id akara.EID) {
	transform, found := m.GetTransform(id)
	if !found {
		return
	}

	velocity, found := m.GetVelocity(id)
	if !found {
		return
	}

	s := float64(m.World.TimeDelta) / float64(time.Second)
	transform.Translation.Add(velocity.Clone().Scale(s))
}
