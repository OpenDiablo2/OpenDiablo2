package d2scene

import (
	"time"

	"github.com/gravestench/akara"
)

type Scene interface {
	Key() string
	Init(world *akara.World)
	Create()
	Update(time.Duration)
	Destroy()
}
