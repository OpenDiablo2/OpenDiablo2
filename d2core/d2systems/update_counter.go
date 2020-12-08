package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const (
	logPrefixUpdateCounter = "Update Counter"
)

var _ akara.System = &UpdateCounter{}

// UpdateCounter is a utility system that logs the number of updates per second
type UpdateCounter struct {
	akara.BaseSystem
	*d2util.Logger
	secondsElapsed float64
	count          int
}

// Init initializes the update counter
func (u *UpdateCounter) Init(world *akara.World) {
	u.World = world

	u.setupLogger()

	if u.World == nil {
		u.SetActive(false)
	}

	u.Debug("initializing")
}

func (u *UpdateCounter) setupLogger() {
	u.Logger = d2util.NewLogger()
	u.SetPrefix(logPrefixUpdateCounter)
}

// Update the world update count in 1 second intervals
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
