package d2app

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2systems"
)

// Run initializes the ECS framework
func Run() {
	cfg := akara.NewWorldConfig().With(&d2systems.AppBootstrap{})
	akara.NewWorld(cfg)
}
