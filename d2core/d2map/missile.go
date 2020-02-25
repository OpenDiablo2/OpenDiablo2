package d2map

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"math"
)

type Missile struct {
	*AnimatedEntity
	record *d2datadict.MissileRecord
}

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

	animation.SetBlend(true)
	//animation.SetPlaySpeed(float64(record.Animation.AnimationSpeed))
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

func (m *Missile) SetRadians(angle float64, done func()) {
	r := float64(m.record.Range)

	x := m.LocationX + (r * math.Cos(angle))
	y := m.LocationY + (r * math.Sin(angle))

	m.SetTarget(x, y, done)
}

func (m *Missile) Advance(tickTime float64) {
	// TODO: collision detection
	m.Step(tickTime)
	m.AnimatedEntity.Advance(tickTime)
}
