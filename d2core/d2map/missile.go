package d2map

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Missile struct {
	AnimatedEntity *AnimatedEntity
	record         *d2datadict.MissileRecord
	done           func()
}

func CreateMissile(x, y int32, targetX, targetY float64, record *d2datadict.MissileRecord) *Missile {
	animation, err := d2asset.LoadAnimation(
		fmt.Sprintf("%s/%s.dcc", d2resource.MissileData, record.Animation.CelFileName),
		d2resource.PaletteUnits,
	)
	if err != nil {
		panic(err)
	}

	animation.PlayForward()
	entity, err := CreateAnimatedEntity(x, y, animation)
	if err != nil {
		panic(err)
	}

	result := &Missile{
		AnimatedEntity: entity,
		record:         record,
	}
	result.AnimatedEntity.SetTarget(targetX, targetY)
	return result
}

func (v *Missile) SetDone(done func()) {
	v.done = done
}

func (v *Missile) Render(target d2render.Surface) {
	v.AnimatedEntity.Render(target)
}

func (v *Missile) GetPosition() (float64, float64) {
	return v.AnimatedEntity.GetPosition()
}

func (v *Missile) Advance(tickTime float64) {
	if v.AnimatedEntity.LocationX != v.AnimatedEntity.TargetX ||
		v.AnimatedEntity.LocationY != v.AnimatedEntity.TargetY {

		v.AnimatedEntity.Step(tickTime)
	} else if v.done != nil {
		v.done()
	}

	// TODO: collision detection
	v.AnimatedEntity.Advance(tickTime)
}
