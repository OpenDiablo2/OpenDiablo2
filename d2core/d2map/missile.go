package d2map

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type Missile struct {
	*AnimatedEntity
	record         *d2datadict.MissileRecord
}

func CreateMissile(x, y int, record *d2datadict.MissileRecord) (*Missile, error) {
	animation, err := d2asset.LoadAnimation(
		fmt.Sprintf("%s/%s.dcc", d2resource.MissileData, record.Animation.CelFileName),
		d2resource.PaletteUnits,
	)
	if err != nil {
		return nil, err
	}

	animation.PlayForward()
	entity := CreateAnimatedEntity(x, y, animation)

	result := &Missile{
		AnimatedEntity: entity,
		record:         record,
	}
	result.Speed = float64(record.Velocity)
	return result, nil
}

func (m *Missile) Advance(tickTime float64) {
	// TODO: collision detection
	m.Step(tickTime)
	m.AnimatedEntity.Advance(tickTime)
}
