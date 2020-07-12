package d2mapentity

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2astar"
)

const (
	directionCount = 64
)

// mapEntity represents an entity on the map that can be animated
// TODO: Has a coordinate (issue #456)
type mapEntity struct {
	Position d2vector.Position
	Target   d2vector.Position

	Speed     float64
	path      []d2astar.Pather
	drawLayer int

	done        func()
	directioner func(direction int)
}

// createMapEntity creates an instance of mapEntity
func createMapEntity(x, y int) mapEntity {
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

func (m *mapEntity) getStepLength(tickTime float64) d2vector.Vector {
	length := tickTime * m.Speed
	v := m.Target.Vector.Clone()
	v.Subtract(&m.Position.Vector)
	v.SetLength(length)
	return v
}

// IsAtTarget returns true if the entity is within a 0.0002 square of it's target and has a path.
func (m *mapEntity) IsAtTarget() bool {
	return m.Position.EqualsApprox(m.Target.Vector) && !m.HasPathFinding()
}

// Step moves the entity along it's path by one tick. If the path is complete it calls entity.done() then returns.
func (m *mapEntity) Step(tickTime float64) {
	if m.IsAtTarget() {
		if m.done != nil {
			m.done()
			m.done = nil
		}

		return
	}

	step := m.getStepLength(tickTime)

	for {
		stepX, stepY := step.X(), step.Y()

		if d2math.EqualsApprox(m.Position.X(), m.Target.X()) {
			stepX = 0
		}

		if d2math.EqualsApprox(m.Position.Y(), m.Target.Y()) {
			stepY = 0
		}

		step.Set(stepX, stepY)
		ApplyVelocity(&m.Position.Vector, &step, &m.Target.Vector)

		if m.Position.EqualsApprox(m.Target.Vector) {
			if len(m.path) > 0 {
				m.SetTarget(m.path[0].(*d2common.PathTile).X*5, m.path[0].(*d2common.PathTile).Y*5, m.done)

				if len(m.path) > 1 {
					m.path = m.path[1:]
				} else {
					m.path = []d2astar.Pather{}
				}
			} else {
				m.Position.Copy(&m.Target.Vector)
			}
		}

		if step.IsZero() {
			break
		}
	}
}

func ApplyVelocity(position, velocity, target *d2vector.Vector) {
	dest := position.Clone()

	dest.Add(velocity)

	destDistance := position.Distance(dest)
	targetDistance := position.Distance(*target)

	if destDistance > targetDistance {
		position.Copy(target)
		velocity.Copy(dest.Subtract(target))
	} else {

		position.Copy(&dest)
		velocity.Set(0, 0)
	}
}

// HasPathFinding returns false if the length of the entity movement path is 0.
func (m *mapEntity) HasPathFinding() bool {
	return len(m.path) > 0
}

// SetTarget sets target coordinates and changes animation based on proximity and direction.
func (m *mapEntity) SetTarget(tx, ty float64, done func()) {
	m.Target.Set(tx, ty)
	m.done = done

	if m.directioner != nil {
		angle := m.Position.DirectionTo(m.Target.Vector)

		m.directioner(angle)
	}
}

// GetPosition returns the entity's current tile position, always a whole number.
func (m *mapEntity) GetPosition() (float64, float64) {
	t := m.Position.Tile()
	return t.X(), t.Y()
}

// GetPositionF returns the entity's current tile position where 0.2 is one sub tile.
func (m *mapEntity) GetPositionF() (float64, float64) {
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
