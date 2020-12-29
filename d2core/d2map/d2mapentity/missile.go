package d2mapentity

import (
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

// Missile is a simple animated entity representing a projectile,
// such as a spell or arrow.
type Missile struct {
	*AnimatedEntity
	record *d2records.MissileRecord
}

// ID returns the missile uuid
func (m *Missile) ID() string {
	return m.AnimatedEntity.uuid
}

// GetPosition returns the position of the missile
func (m *Missile) GetPosition() d2vector.Position {
	return m.AnimatedEntity.Position
}

// GetVelocity returns the velocity vector of the missile
func (m *Missile) GetVelocity() d2vector.Vector {
	return m.AnimatedEntity.velocity
}

// SetRadians adjusts the entity target based on it's range, rotating it's
// current destination by the value of angle in radians.
func (m *Missile) SetRadians(angle float64, done func()) {
	r := float64(m.record.Range)

	x := m.Position.X() + (r * math.Cos(angle))
	y := m.Position.Y() + (r * math.Sin(angle))

	m.setTarget(d2vector.NewPosition(x, y), done)
}

// Advance is called once per frame and processes a
// single game tick.
func (m *Missile) Advance(tickTime float64) {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/819
	m.Step(tickTime)
	m.AnimatedEntity.Advance(tickTime)
}
