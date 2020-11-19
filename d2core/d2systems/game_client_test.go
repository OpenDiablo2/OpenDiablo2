package d2systems

import (
	"github.com/gravestench/akara"
	"testing"
)

func Test_game_client(t *testing.T) {
	cfg := akara.NewWorldConfig()

	renderSys := NewRenderSystem()

	cfg.
		With(NewAppBootstrapSystem()).
		With(renderSys).
		With(NewGameClientBootstrapSystem()).
		With(NewUpdateCounterSystem())

	akara.NewWorld(cfg)

	err := renderSys.Loop()
	if err != nil {
		panic(err)
	}
}
