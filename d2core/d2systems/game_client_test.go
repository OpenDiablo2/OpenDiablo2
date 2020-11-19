package d2systems

import (
	"github.com/gravestench/akara"
	"testing"
)

func Test_game_client(t *testing.T) {
	cfg := akara.NewWorldConfig()

	cfg.
		With(NewAppBootstrapSystem()).
		With(NewGameClientBootstrapSystem())

	akara.NewWorld(cfg)
}
