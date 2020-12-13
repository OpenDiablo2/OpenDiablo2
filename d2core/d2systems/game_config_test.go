package d2systems

import (
	"testing"

	"github.com/gravestench/akara"
)

func Test_GameConfigSystem_Bootstrap(t *testing.T) {
	const testDataPath = "testdata"

	cfg := akara.NewWorldConfig()

	typeSys := &FileTypeResolver{}
	handleSys := &FileHandleResolver{}
	srcSys := &FileSourceResolver{}
	cfgSys := &GameConfigSystem{}

	cfg.With(typeSys).
		With(srcSys).
		With(handleSys).
		With(cfgSys)

	world := akara.NewWorld(cfg)

	cfgSys.Components.File.Add(world.NewEntity()).Path = testDataPath
	cfgSys.Components.File.Add(world.NewEntity()).Path = "config.json"

	// at this point the world has initialized the sceneSystems. when the world
	// updates it should process the config dir to a source and then
	// use the source to resolve a file handle, and finally the config file
	// will get loaded by the config system.
	_ = world.Update(0)

	if len(cfgSys.gameConfigs.GetEntities()) < 1 {
		t.Error("no game configs were loaded")
	}
}
