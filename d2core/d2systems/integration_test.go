package d2systems

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"

	"github.com/gravestench/akara"
)

func Test_integration(t *testing.T) {
	cfg := akara.NewWorldConfig()

	cfg.With(NewFileTypeResolver()).
		With(NewFileSourceResolver()).
		With(NewFileHandleResolver()).
		With(NewGameConfigSystem()).
		With(NewAssetLoader()).
		With(NewRenderSystem()).
		With(NewGameBootstrapSystem())

	world := akara.NewWorld(cfg)

	e1 := world.NewEntity()
	m, err := world.GetMap(d2components.FilePath)
	if err != nil {
		t.Error("cannot find file path component map")
		return
	}

	filepaths := m.(*d2components.FilePathMap)

	filepaths.AddFilePath(e1).Path = "Data/Global/Monsters/DI/LA/DILALITDTHTH.DC6"

	mm, _ := world.ComponentManager.GetMap(d2components.Dc6)
	dc6map := mm.(*d2components.Dc6Map)

	for {
		world.Update(0)
		_, found := dc6map.GetDc6(e1)

		if found {
			break
		}
	}
}
