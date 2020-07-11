package d2mapentity

import (
	"fmt"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// Missile is a simple animated entity representing a projectile,
// such as a spell or arrow.
type Missile struct {
	*AnimatedEntity
	record *d2datadict.MissileRecord
}

// CreateMissile creates a new Missile and initializes it's animation.
func CreateMissile(x, y int, record *d2datadict.MissileRecord) (*Missile, error) {
	animation, err := d2asset.LoadAnimation(
		fmt.Sprintf("%s/%s.dcc", d2resource.MissileData, record.Animation.CelFileName),
		d2resource.PaletteUnits,
	)
	if err != nil {
		return nil, err
	}

	if record.Animation.HasSubLoop {
		animation.SetSubLoop(record.Animation.SubStartingFrame, record.Animation.SubEndingFrame)
	}

	animation.SetEffect(d2enum.DrawEffectModulate)
	// animation.SetPlaySpeed(float64(record.Animation.AnimationSpeed))
	animation.SetPlayLoop(record.Animation.LoopAnimation)
	animation.PlayForward()
	entity := CreateAnimatedEntity(x, y, animation)

	result := &Missile{
		AnimatedEntity: entity,
		record:         record,
	}
	result.Speed = float64(record.Velocity)

	return result, nil
}

// SetRadians adjusts the entity target based on it's range, rotating it's
// current destination by the value of angle in radians.
func (m *Missile) SetRadians(angle float64, done func()) {
	r := float64(m.record.Range)

	x := m.Position.X() + (r * math.Cos(angle))
	y := m.Position.Y() + (r * math.Sin(angle))

	m.setTarget(x, y, done)
}

// Advance is called once per frame and processes a
// single game tick.
func (m *Missile) Advance(tickTime float64) {
	// TODO: collision detection
	m.Step(tickTime)
	m.AnimatedEntity.Advance(tickTime)
}
