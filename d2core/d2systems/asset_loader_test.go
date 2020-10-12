package d2systems

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/gravestench/akara"
)

func Test_AssetLoader(t *testing.T) {
	cfg := akara.NewWorldConfig()

	fileTypeResolver := NewFileTypeResolver()
	fileHandleResolver := NewFileHandleResolver()
	fileSourceResolver := NewFileSourceResolver()
	gameConfig := NewGameConfigSystem()
	assetLoader := NewAssetLoader()

	cfg.With(fileTypeResolver).
		With(fileSourceResolver).
		With(fileHandleResolver).
		With(gameConfig).
		With(assetLoader)

	world := akara.NewWorld(cfg)

	toLoad := []struct {
		path string
	}{
		{d2resource.CharacterSelectBarbarianUnselected},
		{d2resource.CharacterSelectBarbarianUnselectedH},
		{d2resource.CharacterSelectBarbarianSelected},
		{d2resource.CharacterSelectBarbarianForwardWalk},
		{d2resource.CharacterSelectBarbarianForwardWalkOverlay},
		{d2resource.CharacterSelectBarbarianBackWalk},
	}

	filePathsAbstract, err := world.GetMap(d2components.FilePath)
	if err != nil {
		t.Error("file path component map not found")
	}

	filePaths := filePathsAbstract.(*d2components.FilePathMap)

	fpComponents := make(map[akara.EID]*d2components.FilePathComponent)

	for idx := range toLoad {
		eid := world.NewEntity()
		fp := filePaths.AddFilePath(eid)
		fp.Path = toLoad[idx].path
		fpComponents[eid] = fp
	}

	_ = world.Update(0)
	_ = world.Update(0)

}
