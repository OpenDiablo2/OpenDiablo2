package d2systems

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"

	"github.com/gravestench/akara"
)

func Test_GameConfigSystem_Bootstrap(t *testing.T) {
	cfg := akara.NewWorldConfig()

	fileTypeResolver := NewFileTypeResolver()
	fileHandleResolver := NewFileHandleResolver()
	fileSourceResolver := NewFileSourceResolver()
	gameConfig := NewGameConfigSystem()

	cfg.With(fileTypeResolver).
		With(fileSourceResolver).
		With(fileHandleResolver).
		With(gameConfig)

	world := akara.NewWorld(cfg)

	// for the purpose of this test, we want to add the testdata directory, so that
	// when the game looks for the config file it gets pulled from there
	filePathsAbstract, err := world.GetMap(d2components.FilePath)
	if err != nil {
		t.Error("file path component map not found")
		return
	}

	filePaths := filePathsAbstract.(*d2components.FilePathMap)
	testDir := filePaths.AddFilePath(world.NewEntity())
	testDir.Path = "./testdata/"

	// at this point, bootstrap has been run (NewWorld called the system init methods)
	// but the entities for the sources and config file have not been evaluated by the systems
	// so after this update, we wont yet have a config file
	_ = world.Update(0)

	// This actually creates the config file because a file handle for the
	// config file has been resolved
	_ = world.Update(0)

	if len(gameConfig.gameConfigs.GetEntities()) < 1 {
		t.Error("no game configs were loaded")
	}
}
