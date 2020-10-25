package d2mapentity

import (
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

// CastOverlay is an animated entity representing a projectile that is a result of a skill cast.
type CastOverlay struct {
	*AnimatedEntity
	record     *d2records.OverlayRecord
	playLoop   bool
	onDoneFunc func()
}

// ID returns the overlay uuid
func (co *CastOverlay) ID() string {
	return co.AnimatedEntity.uuid
}

// GetPosition returns the position of the overlay
func (co *CastOverlay) GetPosition() d2vector.Position {
	return co.AnimatedEntity.Position
}

// GetVelocity returns the velocity vector of the overlay
func (co *CastOverlay) GetVelocity() d2vector.Vector {
	return co.AnimatedEntity.velocity
}

// SetRadians adjusts the entity target based on it's range, rotating it's
// current destination by the value of angle in radians.
func (co *CastOverlay) SetRadians(angle float64, done func()) {
	rads := float64(co.record.Height2)

	x := co.Position.X() + (rads * math.Cos(angle))
	y := co.Position.Y() + (rads * math.Sin(angle))

	co.setTarget(d2vector.NewPosition(x, y), done)
}

// SetOnDoneFunc changes the handler func that gets called when the overlay finishes playing.
func (co *CastOverlay) SetOnDoneFunc(onDoneFunc func()) {
	co.onDoneFunc = onDoneFunc
}

// Advance is called once per frame and processes a single game tick.
func (co *CastOverlay) Advance(tickTime float64) {
	co.Step(tickTime)
	co.AnimatedEntity.Advance(tickTime)

	if !co.playLoop && co.AnimatedEntity.animation.GetPlayedCount() >= 1 {
		co.onDoneFunc()
	}
}
