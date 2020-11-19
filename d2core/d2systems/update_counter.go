package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/gravestench/akara"
)

const (
	logPrefixUpdateCounter = "Update Counter"
)

func NewUpdateCounterSystem() *UpdateCounter {
	uc := &UpdateCounter{
		BaseSystem: &akara.BaseSystem{},
		Logger:     d2util.NewLogger(),
	}

	uc.SetPrefix(logPrefixUpdateCounter)

	return uc
}

var _ akara.System = &UpdateCounter{}

type UpdateCounter struct {
	*akara.BaseSystem
	*d2util.Logger
	secondsElapsed float64
	count          int
}

func (u *UpdateCounter) Init(world *akara.World) {
	u.World = world

	if u.World == nil {
		u.SetActive(false)
	}

	u.Info("initializing")
}

func (u *UpdateCounter) Update() {
	u.count++
	u.secondsElapsed += u.World.TimeDelta.Seconds()

	if u.secondsElapsed < 1 {
		return
	}

	u.Infof("%d updates per second", u.count)
	u.secondsElapsed = 0
	u.count = 0
}

