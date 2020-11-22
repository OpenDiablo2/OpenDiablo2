package d2systems

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"

	"github.com/gravestench/akara"
)

func Test_GameConfigSystem_Bootstrap(t *testing.T) {
	const testDataPath = "testdata"

	cfg := akara.NewWorldConfig()

	typeSys := NewFileTypeResolver()
	handleSys := NewFileHandleResolver()
	srcSys := NewFileSourceResolver()
	cfgSys := NewGameConfigSystem()

	cfg.With(typeSys).
		With(srcSys).
		With(handleSys).
		With(cfgSys)

	world := akara.NewWorld(cfg)

	// for the purpose of this test, we want to add the testdata directory, so that
	// when the game looks for the config file it gets pulled from there
	filePathsAbstract, err := world.GetMap(d2components.FilePath)
	if err != nil {
		t.Error("file path component map not found")
		return
	}

	filePaths := filePathsAbstract.(*d2components.FilePathMap)

	cfgDir := filePaths.AddFilePath(world.NewEntity())
	cfgDir.Path = testDataPath

	cfgFile := filePaths.AddFilePath(world.NewEntity())
	cfgFile.Path = "config.json"

	// at this point the world has initialized the systems. when the world
	// updates it should process the config dir to a source and then
	// use the source to resolve a file handle, and finally the config file
	// will get loaded by the config system.
	_ = world.Update(0)

	if len(cfgSys.gameConfigs.GetEntities()) < 1 {
		t.Error("no game configs were loaded")
	}
}
