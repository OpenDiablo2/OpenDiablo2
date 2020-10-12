package d2systems

import (
	"testing"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func TestNewFileTypeResolver_KnownType(t *testing.T) {
	cfg := akara.NewWorldConfig()

	resolver := NewFileTypeResolver()

	cfg.With(resolver)

	world := akara.NewWorld(cfg)

	e := world.NewEntity()

	fp := resolver.filePaths.AddFilePath(e)
	fp.Path = "/some/path/to/a/file.dcc"

	if len(resolver.Subscriptions[0].GetEntities()) != 1 {
		t.Error("entity with file path not added to file type resolver subscription")
	}

	_ = world.Update(0)

	if len(resolver.Subscriptions[0].GetEntities()) != 0 {
		t.Error("entity with existing file type not removed from file type resolver subscription")
	}

	ft, found := resolver.fileTypes.GetFileType(e)
	if !found {
		t.Error("file type component not added to entity with file path component")
	}

	if ft.Type != d2enum.FileTypeDCC {
		t.Error("unexpected file type")
	}
}

func TestNewFileTypeResolver_UnknownType(t *testing.T) {
	cfg := akara.NewWorldConfig()

	resolver := NewFileTypeResolver()

	cfg.With(resolver)

	world := akara.NewWorld(cfg)

	e := world.NewEntity()

	fp := resolver.filePaths.AddFilePath(e)
	fp.Path = "/some/path/to/a/file.XYZ"

	_ = world.Update(0)

	ft, _ := resolver.fileTypes.GetFileType(e)

	if ft.Type != d2enum.FileTypeUnknown {
		t.Error("unexpected file type")
	}
}
