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
		Position:  d2vector.EntityPosition(locX, locY),
		Target:    d2vector.EntityPosition(locX, locY),
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

func (m *mapEntity) getStepLength(tickTime float64) (v *d2vector.Vector) {
	length := tickTime * m.Speed
	v = m.Target.WorldSubTile()
	v.Subtract(m.Position.WorldSubTile())
	v.SetLength(length)
	return
}

// IsAtTarget returns true if the entity is within a 0.0002 square of it's target and has a path.
func (m *mapEntity) IsAtTarget() bool {
	return m.Position.EqualsApprox(m.Target.Vector) && !m.HasPathFinding()
}

// Step moves the entity along it's path by one tick. If the path is complete it calls entity.done() then returns.
func (m *mapEntity) Step(tickTime float64) {
	// no movement needed
	if m.IsAtTarget() {
		if m.done != nil {
			m.done()
			m.done = nil
		}

		return
	}

	// per tick velocity
	step := m.getStepLength(tickTime)

	// endless loop - why?

	for {
		stepX, stepY := step.X(), step.Y()

		position := m.Position.WorldSubTile()
		targetPosition := m.Target.WorldSubTile()

		// zero small values
		if d2math.CompareFloat64Fuzzy(position.X(), targetPosition.X()) == 0 {
			stepX = 0
		}

		if d2math.CompareFloat64Fuzzy(position.Y(), targetPosition.Y()) == 0 {
			stepY = 0
		}

		step.Set(stepX, stepY)

		// add velocity to current position, or return target if new position exceeds it

		newPositionX, newStepX := d2common.AdjustWithRemainder(position.X(), step.X(), targetPosition.X())
		newPositionY, newStepY := d2common.AdjustWithRemainder(position.Y(), step.Y(), targetPosition.Y())

		m.Position.SetSubWorld(newPositionX, newPositionY)
		step.Set(newStepX, newStepY)

		position = m.Position.WorldSubTile()
		// position is close to target
		if d2common.AlmostEqual(position.X(), targetPosition.X(), 0.01) && d2common.AlmostEqual(position.Y(), targetPosition.Y(), 0.01) {
			// entity has a path
			if len(m.path) > 0 {
				// set target as next node in path
				m.SetTarget(m.path[0].(*d2common.PathTile).X*5, m.path[0].(*d2common.PathTile).Y*5, m.done)

				// remove path node or set to empty slice if path is empty
				if len(m.path) > 1 {
					m.path = m.path[1:]
				} else {
					m.path = []d2astar.Pather{}
				}
				// entity had no path
				// set location to target
			} else {
				m.Position.Copy(&m.Target.Vector)
				position = m.Position.WorldSubTile()
			}
		}

		if step.IsZero() {
			break
		}
	}
}

// HasPathFinding returns false if the length of the entity movement path is 0.
func (m *mapEntity) HasPathFinding() bool {
	return len(m.path) > 0
}

// SetTarget sets target coordinates and changes animation based on proximity and direction.
func (m *mapEntity) SetTarget(tx, ty float64, done func()) {
	m.Target.SetSubWorld(tx, ty)
	m.done = done

	if m.directioner != nil {
		direction := m.Target.Clone()
		direction.Subtract(&m.Position.Vector)

		angle := direction.Direction()

		m.directioner(int(angle))
	}
}

// GetPosition returns the entity's current tile position.
func (m *mapEntity) GetPosition() (float64, float64) {
	t := m.Position.Tile()
	return t.X(), t.Y()
}

// GetPositionF returns the entity's current sub tile position.
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
