package d2scenemanager

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// Scene defines the function necessary for scene management
type Scene interface {
	Load() []func()
	Unload()
	Render(target d2common.Surface)
	Advance(tickTime float64)
}
