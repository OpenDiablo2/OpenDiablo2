package d2systems

import (
	"strings"
	"testing"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

func Test_FileHandleResolver_Process(t *testing.T) {
	cfg := akara.NewWorldConfig()

	fileTypeResolver := NewFileTypeResolver()
	fileHandleResolver := NewFileHandleResolver()
	fileSourceResolver := NewFileSourceResolver()

	cfg.With(fileTypeResolver).
		With(fileSourceResolver).
		With(fileHandleResolver)

	world := akara.NewWorld(cfg)

	filepathMap, err := world.GetMap(d2components.FilePath)
	if err != nil {
		t.Error("file path component map not found")
	}

	filePaths := filepathMap.(*d2components.FilePathMap)

	sourceEntity := world.NewEntity()
	sourceFp := filePaths.AddFilePath(sourceEntity)
	sourceFp.Path = "./testdata/"

	//_ = world.Update(0)

	fileEntity := world.NewEntity()
	fileFp := filePaths.AddFilePath(fileEntity)
	fileFp.Path = "testfile_a.txt"

	_ = world.Update(0)

	ft, found := fileTypeResolver.fileTypes.GetFileType(sourceEntity)
	if !found {
		t.Error("file source type not created for entity")
		return
	}

	if ft.Type != d2enum.FileTypeDirectory {
		t.Error("expected file system source type for entity")
		return
	}

	handleMap, err := world.GetMap(d2components.FileHandle)
	if err != nil {
		t.Error("file handle component map is nil")
		return
	}

	fileHandles := handleMap.(*d2components.FileHandleMap)

	handle, found := fileHandles.GetFileHandle(fileEntity)
	if !found {
		t.Error("file handle for entity was not found")
		return
	}

	data, buf := make([]byte, 0), make([]byte, 16)

	for {
		numRead, err := handle.Data.Read(buf)

		data = append(data, buf[:numRead]...)

		if err != nil || numRead == 0 {
			break
		}
	}

	result := strings.Trim(string(data), "\r\n")

	if result != "test a" {
		t.Error("unexpected data read from `./testdata/testfile_a.txt`")
	}
}
