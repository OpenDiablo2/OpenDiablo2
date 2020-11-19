package d2systems

import (
	"testing"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

func Test_FileSourceResolution(t *testing.T) {
	cfg := akara.NewWorldConfig()

	srcResolver := NewFileSourceResolver()
	fileTypeResolver := NewFileTypeResolver()

	cfg.With(fileTypeResolver).
		With(srcResolver)

	world := akara.NewWorld(cfg)

	filepathMap, err := world.GetMap(d2components.FilePath)
	if err != nil {
		t.Error("file path component map not found")
	}

	filePaths := filepathMap.(*d2components.FilePathMap)

	sourceEntity := world.NewEntity()
	sourceFp := filePaths.AddFilePath(sourceEntity)
	sourceFp.Path = "./testdata/"

	_ = world.Update(0)

	ft, found := fileTypeResolver.GetFileType(sourceEntity)
	if !found {
		t.Error("file source type not created for entity")
	}

	if ft.Type != d2enum.FileTypeDirectory {
		t.Error("expected file system source type for entity")
	}

	fs, found := srcResolver.fileSources.GetFileSource(sourceEntity)
	if !found {
		t.Error("file source not created for entity")
	}

	if fs.AbstractSource == nil {
		t.Error("nil file AbstractSource interface inside of file source component")
	}
}
