package d2interface

import (
	"github.com/gravestench/akara"
)

type Scene interface {
	akara.SystemInitializer
	Key() string
	Booted() bool
	Paused() bool
}
