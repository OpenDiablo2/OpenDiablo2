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
	d2components.PositionFactory
	d2components.VelocityFactory
	movableEntities *akara.Subscription
}

// Init initializes the system with the given world
func (m *MovementSystem) Init(world *akara.World) {
	m.World = world

	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixMovementSystem)

	m.Info("initializing ...")

	positionID := m.RegisterComponent(&d2components.Position{})
	velocityID := m.RegisterComponent(&d2components.Velocity{})

	m.Position = m.GetComponentFactory(positionID)
	m.Velocity = m.GetComponentFactory(velocityID)

	movable := m.NewComponentFilter().Require(
		&d2components.Position{},
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
	position, found := m.GetPosition(id)
	if !found {
		return
	}

	velocity, found := m.GetVelocity(id)
	if !found {
		return
	}

	s := float64(m.World.TimeDelta) / float64(time.Second)
	position.Add(velocity.Clone().Scale(s))
}
