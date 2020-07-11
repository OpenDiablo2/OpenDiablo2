package d2mapentity

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2astar"
)

// mapEntity represents an entity on the map that can be animated
type mapEntity struct {
	Position d2vector.Position
	Target   d2vector.Position

	Speed     float64
	path      []d2astar.Pather
	drawLayer int

	done        func()
	directioner func(direction int)
}

// newMapEntity creates an instance of mapEntity
func newMapEntity(x, y int) mapEntity {
	locX, locY := float64(x), float64(y)

	return mapEntity{
		Position:  d2vector.NewPosition(locX, locY),
		Target:    d2vector.NewPosition(locX, locY),
		Speed:     6,
		drawLayer: 0,
		path:      []d2astar.Pather{},
	}
}

// GetLayer returns the draw layer for this entity.
func (m *mapEntity) GetLayer() int {
	return m.drawLayer
}

// SetPath sets the entity movement path. done() is called when the entity reaches it's path destination. For example,
// when the player entity reaches the point a player clicked.
func (m *mapEntity) SetPath(path []d2astar.Pather, done func()) {
	m.path = path
	m.done = done
}

// ClearPath clears the entity movement path.
func (m *mapEntity) ClearPath() {
	m.path = nil
}

// SetSpeed sets the entity movement speed.
func (m *mapEntity) SetSpeed(speed float64) {
	m.Speed = speed
}

// GetSpeed returns the entity movement speed.
func (m *mapEntity) GetSpeed() float64 {
	return m.Speed
}

// atTarget returns true if the distance between entity and target is almost zero.
// TODO: !m.hasPath() should be called separately by the caller, not hidden in here.
func (m *mapEntity) atTarget() bool {
	return m.Position.EqualsApprox(m.Target.Vector) && !m.hasPath()
}

// Step moves the entity along it's path by one tick. If the path is complete it calls entity.done() then returns.
func (m *mapEntity) Step(tickTime float64) {
	if m.atTarget() {
		if m.done != nil {
			m.done()
			m.done = nil
		}

		return
	}

	velocity := m.velocity(tickTime)

	for {
		// Add the velocity to the position and set new velocity to remainder
		applyVelocity(&m.Position.Vector, &velocity, &m.Target.Vector)

		// New position is at target
		if m.Position.EqualsApprox(m.Target.Vector) {
			if m.hasPath() {
				// Set next path node
				m.setTarget(m.path[0].(*d2common.PathTile).X*5, m.path[0].(*d2common.PathTile).Y*5, m.done)

				if len(m.path) > 1 {
					m.path = m.path[1:]
				} else {
					m.path = []d2astar.Pather{}
				}
			} else {
				// End of path.
				m.Position.Copy(&m.Target.Vector)
			}
		}

		if velocity.IsZero() {
			break
		}
	}
}

// velocity returns a vector describing the change in position this tick.
func (m *mapEntity) velocity(tickTime float64) d2vector.Vector {
	length := tickTime * m.Speed
	v := m.Target.Vector.Clone()
	v.Subtract(&m.Position.Vector)
	v.SetLength(length)

	return v
}

// applyVelocity adds velocity to position. If the new position extends beyond target from the original position, the
// position is set to the target and velocity is set to the overshot amount.
func applyVelocity(position, velocity, target *d2vector.Vector) {
	// Set velocity values to zero if almost zero
	x, y := position.CompareApprox(*target)
	vx, vy := velocity.X(), velocity.Y()

	if x == 0 {
		vx = 0
	}

	if y == 0 {
		vy = 0
	}

	velocity.Set(vx, vy)

	dest := position.Clone()
	dest.Add(velocity)

	destDistance := position.Distance(dest)
	targetDistance := position.Distance(*target)

	if destDistance > targetDistance {
		// Destination overshot target. Set position to target and velocity to overshot amount.
		position.Copy(target)
		velocity.Copy(dest.Subtract(target))
	} else {
		// At or before target, set position to destination and velocity to zero.
		position.Copy(&dest)
		velocity.Set(0, 0)
	}
}

// hasPath returns false if the length of the entity movement path is 0.
func (m *mapEntity) hasPath() bool {
	return len(m.path) > 0
}

// setTarget sets target coordinates and changes animation based on proximity and direction.
func (m *mapEntity) setTarget(tx, ty float64, done func()) {
	m.Target.Set(tx, ty)
	m.done = done

	if m.directioner != nil {
		d := m.Position.DirectionTo(m.Target.Vector)

		m.directioner(d)
	}
}

// GetPosition returns the entity's current tile position, always a whole number.
func (m *mapEntity) GetPosition() (x, y float64) {
	t := m.Position.Tile()
	return t.X(), t.Y()
}

// GetPositionF returns the entity's current tile position where 0.2 is one sub tile.
func (m *mapEntity) GetPositionF() (x, y float64) {
	w := m.Position.World()
	return w.X(), w.Y()
}

// Name returns the NPC's in-game name (e.g. "Deckard Cain") or an empty string if it does not have a name
func (m *mapEntity) Name() string {
	return ""
}

// Highlight is not currently implemented.
func (m *mapEntity) Highlight() {
}

// Selectable returns true if the object can be highlighted/selected.
func (m *mapEntity) Selectable() bool {
	return false
}
