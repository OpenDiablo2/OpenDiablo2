package d2interface

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Scene is an extension of akara.System
type Scene interface {
	akara.SystemInitializer
	State() d2enum.SceneState
	Key() string
	Booted() bool
	Paused() bool
}
