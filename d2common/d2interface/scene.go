package d2interface

import (
	"github.com/gravestench/akara"
)

// Scene is an extension of akara.System
type Scene interface {
	akara.SystemInitializer
	Key() string
	Booted() bool
	Paused() bool
}
