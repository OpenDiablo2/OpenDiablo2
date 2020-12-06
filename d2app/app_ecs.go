package d2app

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2systems"
	"github.com/gravestench/akara"
)

func Run() {
	cfg := akara.NewWorldConfig().With(&d2systems.AppBootstrap{})
	akara.NewWorld(cfg)
}
