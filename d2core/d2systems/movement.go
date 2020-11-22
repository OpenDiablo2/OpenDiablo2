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

// NewMovementSystem creates a movement system
func NewMovementSystem() *MovementSystem {
	cfg := akara.NewFilter().Require(d2components.Position, d2components.Velocity)

	filter := cfg.Build()

	sys := &MovementSystem{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(filter),
		Logger:               d2util.NewLogger(),
	}

	sys.SetPrefix(logPrefixMovementSystem)

	return sys
}

// static check that MovementSystem implements the System interface
var _ akara.System = &MovementSystem{}

// MovementSystem handles entity movement based on velocity and position components
type MovementSystem struct {
	*akara.BaseSubscriberSystem
	*d2util.Logger
	*d2components.PositionMap
	*d2components.VelocityMap
}

// Init initializes the system with the given world
func (m *MovementSystem) Init(_ *akara.World) {
	m.Info("initializing ...")

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.PositionMap = m.InjectMap(d2components.Position).(*d2components.PositionMap)
	m.VelocityMap = m.InjectMap(d2components.Velocity).(*d2components.VelocityMap)
}

// Update positions of all entities with their velocities
func (m *MovementSystem) Update() {
	for subIdx := range m.Subscriptions {
		entities := m.Subscriptions[subIdx].GetEntities()

		m.Infof("Processing movement for %d entities ...", len(entities))

		for entIdx := range entities {
			m.ProcessEntity(entities[entIdx])
		}
	}
}

// ProcessEntity updates an individual entity in the movement system
func (m *MovementSystem) ProcessEntity(id akara.EID) {
	position, found := m.GetPosition(id)
	if !found {
		return
	}

	velocity, found := m.GetVelocity(id)
	if !found {
		return
	}

	s := float64(m.World.TimeDelta) / float64(time.Second)
	position.Vector = *position.Vector.Add(velocity.Vector.Clone().Scale(s))
}
