package d2systems

import (
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/gravestench/akara"
)

const (
	defaultScale float64 = 1
)

const (
	logPrefixTimeScaleSystem = "Time Scale"
)

// NewTimeScaleSystem creates a timescale system
func NewTimeScaleSystem() *TimeScaleSystem {
	m := &TimeScaleSystem{
		BaseSystem: &akara.BaseSystem{},
		Logger:     d2util.NewLogger(),
	}

	m.SetPrefix(logPrefixTimeScaleSystem)

	return m
}

// static check that TimeScaleSystem implements the System interface
var _ akara.System = &TimeScaleSystem{}

// TimeScaleSystem should be the first system added to the world, and whose only job is to
// apply a scalar the world's TimeDelta between frames. It's useful for slowing down or speeding
// up the game time without affecting the render rate.
type TimeScaleSystem struct {
	*akara.BaseSystem
	*d2util.Logger
	scale     float64
	lastScale float64
}

// Init will initialize the TimeScale system
func (t *TimeScaleSystem) Init(world *akara.World) {
	t.World = world

	t.Info("initializing ...")

	t.scale = defaultScale
}

// Update scales the worlds time delta for this frame
func (t *TimeScaleSystem) Update() {
	if !t.Active() || t.scale == t.lastScale {
		return
	}

	t.Infof("setting time scale to %.1f", t.scale)
	t.lastScale = t.scale

	t.World.TimeDelta *= time.Duration(t.scale)
}
