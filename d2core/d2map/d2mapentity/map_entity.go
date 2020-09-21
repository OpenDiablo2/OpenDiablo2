package d2mapentity

import (
	"github.com/google/uuid"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

const (
	minHitboxSize = 30
)

// mapEntity represents an entity on the map that can be animated
type mapEntity struct {
	uuid     string
	Position d2vector.Position
	Target   d2vector.Position
	velocity d2vector.Vector

	Speed     float64
	path      []d2vector.Position
	drawLayer int

	done        func()
	directioner func(direction int)

	highlight bool
}

// newMapEntity creates an instance of mapEntity
func newMapEntity(x, y int) mapEntity {
	pos := d2vector.NewPosition(float64(x), float64(y))

	return mapEntity{
		uuid:     uuid.New().String(),
		Position: pos,
		Target:   pos,
		velocity: *d2vector.VectorZero(),
	}
}

// GetLayer returns the draw layer for this entity.
func (m *mapEntity) GetLayer() int {
	return m.drawLayer
}

// SetPath sets the entity movement path. done() is called when the entity reaches it's path destination. For example,
// when the player entity reaches the point a player clicked.
func (m *mapEntity) SetPath(path []d2vector.Position, done func()) {
	m.path = path
	m.done = done
	m.nextPath()
}

// ClearPath clears the entity movement path.
func (m *mapEntity) ClearPath() {
	m.path = nil
}

// StopMoving  will clear the path and target of the entity.
func (m *mapEntity) StopMoving() {
	m.ClearPath()
	m.setTarget(m.Position, nil)
}

// SetSpeed sets the entity movement speed.
func (m *mapEntity) SetSpeed(speed float64) {
	m.Speed = speed
}

// GetSpeed returns the entity movement speed.
func (m *mapEntity) GetSpeed() float64 {
	return m.Speed
}

// Step moves the entity along it's path by one tick. If the path is complete it calls entity.done() then returns.
func (m *mapEntity) Step(tickTime float64) {
	if m.atTarget() && !m.hasPath() {
		if m.done != nil {
			m.done()
			m.done = nil
		}

		m.velocity.SetLength(0)

		return
	}

	m.setVelocity(tickTime * m.Speed)

	v := m.velocity.Clone() // Create a new vector

	for {
		applyVelocity(&m.Position.Vector, v, &m.Target.Vector) // Pass the new vector to the function which alters it

		if m.atTarget() {
			m.nextPath()
		}

		if v.IsZero() { // Check if the new vector is zero (keeping this as m.velocity.IsZero() would break the test (infinite loop)
			break
		}
	}
}

// atTarget returns true if the distance between entity and target is almost zero.
func (m *mapEntity) atTarget() bool {
	return m.Position.EqualsApprox(&m.Target.Vector)
}

// setVelocity returns a vector describing the given length and the direction to the current target.
func (m *mapEntity) setVelocity(length float64) {
	m.velocity.Copy(&m.Target.Vector)
	m.velocity.Subtract(&m.Position.Vector)
	m.velocity.SetLength(length)
}

// applyVelocity adds velocity to position. If the new position extends beyond the target: Target is set to the next
// path node, Position is set to target and velocity is set to the over-extended length with the direction of to the
// next node.
func applyVelocity(position, velocity, target *d2vector.Vector) {
	// Set velocity values to zero if almost zero
	x, y := position.CompareApprox(target)
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
	targetDistance := position.Distance(target)

	if destDistance > targetDistance {
		// Destination overshot target. Set position to target and velocity to overshot amount.
		position.Copy(target)
		velocity.Copy(dest.Subtract(target))
	} else {
		// At or before target, set position to destination and velocity to zero.
		position.Copy(dest)
		velocity.Set(0, 0)
	}
}

// Returns false if the path has ended.
func (m *mapEntity) nextPath() {
	if m.hasPath() {
		// Set next path node
		m.setTarget(
			m.path[0],
			m.done,
		)

		if len(m.path) > 1 {
			m.path = m.path[1:]
		} else {
			m.path = []d2vector.Position{}
		}
	} else {
		// End of path.
		m.Position.Copy(&m.Target.Vector)
	}
}

// hasPath returns false if the length of the entity movement path is 0.
func (m *mapEntity) hasPath() bool {
	return len(m.path) > 0
}

// setTarget sets target coordinates and changes animation based on proximity and direction.
func (m *mapEntity) setTarget(p d2vector.Position, done func()) {
	// Set the target
	m.Target.Copy(&p.Vector)
	m.done = done

	// Update the direction
	if m.directioner != nil {
		d := m.Position.DirectionTo(m.Target.Vector)

		m.directioner(d)
	}

	// Update the velocity direction
	if !m.velocity.IsZero() {
		m.setVelocity(m.velocity.Length())
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

// Label returns the NPC's in-game name (e.g. "Deckard Cain") or an empty string if it does not have a name
func (m *mapEntity) Label() string {
	return ""
}

// Highlight is not currently implemented.
func (m *mapEntity) Highlight() {
	m.highlight = true
}

// Selectable returns true if the object can be highlighted/selected.
func (m *mapEntity) Selectable() bool {
	return false
}

// GetSize returns the current frame size
func (m *mapEntity) GetSize() (width, height int) {
	return minHitboxSize, minHitboxSize
}
