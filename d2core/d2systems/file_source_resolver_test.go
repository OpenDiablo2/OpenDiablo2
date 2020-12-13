package d2systems

import (
	"testing"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func Test_FileSourceResolution(t *testing.T) {
	const testDataPath = "testdata"

	cfg := akara.NewWorldConfig()

	sourceSys := &FileSourceResolver{}
	typeSys := &FileTypeResolver{}

	cfg.With(typeSys).
		With(sourceSys)

	world := akara.NewWorld(cfg)

	sourceEntity := world.NewEntity()
	sourceFp := typeSys.Components.File.Add(sourceEntity)
	sourceFp.Path = testDataPath

	_ = world.Update(0)

	ft, found := typeSys.Components.FileType.Get(sourceEntity)
	if !found {
		t.Error("file source type not created for entity")
	}

	if ft.Type != d2enum.FileTypeDirectory {
		t.Error("expected file system source type for entity")
	}

	fs, found := sourceSys.Components.FileSource.Get(sourceEntity)
	if !found {
		t.Error("file source not created for entity")
	}

	if fs.AbstractSource == nil {
		t.Error("nil file AbstractSource interface inside of file source component")
	}
}
