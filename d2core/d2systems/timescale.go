package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
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

// static check that TimeScaleSystem implements the System interface
var _ akara.System = &TimeScaleSystem{}

// TimeScaleSystem should be the first system added to the world, and whose only job is to
// apply a scalar the world's TimeDelta between frames. It's useful for slowing down or speeding
// up the game time without affecting the render rate.
type TimeScaleSystem struct {
	akara.BaseSystem
	*d2util.Logger
	scale     float64
	Components struct {
		Dirty d2components.DirtyFactory
		CommandRegistration d2components.CommandRegistrationFactory
	}
}

// Init will initialize the TimeScale system
func (t *TimeScaleSystem) Init(world *akara.World) {
	t.World = world

	t.Logger = d2util.NewLogger()
	t.SetPrefix(logPrefixTimeScaleSystem)

	t.Debug("initializing ...")

	t.InjectComponent(&d2components.CommandRegistration{}, &t.Components.CommandRegistration.ComponentFactory)
	t.InjectComponent(&d2components.Dirty{}, &t.Components.Dirty.ComponentFactory)

	t.registerCommands()

	t.scale = defaultScale
}

// Update scales the worlds time delta for this frame
func (t *TimeScaleSystem) Update() {
	if !t.Active() {
		return
	}

	t.World.TimeDelta = time.Duration(float64(t.World.TimeDelta) * t.scale)
}

func (t *TimeScaleSystem) registerCommands() {
	e := t.NewEntity()

	reg := t.Components.CommandRegistration.Add(e)

	t.Components.Dirty.Add(e)

	reg.Name = "timescale"
	reg.Description = "set the time scale of the game (default is 1.0)"
	reg.Callback = func(scale float64) {
		t.Infof("setting time scale to %.1f", scale)
		t.scale = scale
	}
}
