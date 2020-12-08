package d2systems

import (
	"testing"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func TestNewFileTypeResolver_KnownType(t *testing.T) {
	typeSys := &FileTypeResolver{}
	world := akara.NewWorld(akara.NewWorldConfig().With(typeSys))

	e := world.NewEntity()
	typeSys.Components.File.Add(e).Path = "/some/path/to/a/file.dcc"

	if len(typeSys.filesToCheck.GetEntities()) != 1 {
		t.Error("entity with file path not added to file type typeSys subscription")
	}

	_ = world.Update(0)

	if len(typeSys.filesToCheck.GetEntities()) != 0 {
		t.Error("entity with existing file type not removed from file type typeSys subscription")
	}

	ft, found := typeSys.Components.FileType.Get(e)
	if !found {
		t.Error("file type component not added to entity with file path component")
	}

	if ft.Type != d2enum.FileTypeDCC {
		t.Error("unexpected file type")
	}
}

func TestNewFileTypeResolver_UnknownType(t *testing.T) {
	typeSys := &FileTypeResolver{}
	world := akara.NewWorld(akara.NewWorldConfig().With(typeSys))

	e := world.NewEntity()

	fp := typeSys.Components.File.Add(e)
	fp.Path = "/some/path/to/a/file.XYZ"

	_ = world.Update(0)

	ft, _ := typeSys.Components.FileType.Get(e)

	if ft.Type != d2enum.FileTypeUnknown {
		t.Error("unexpected file type")
	}
}
