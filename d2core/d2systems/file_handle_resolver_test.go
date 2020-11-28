package d2systems

import (
	"strings"
	"testing"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func Test_FileHandleResolver_Process(t *testing.T) {
	const testDataPath = "testdata"

	cfg := akara.NewWorldConfig()

	typeSys := &FileTypeResolver{}
	handleSys := &FileHandleResolver{}
	sourceSys := &FileSourceResolver{}

	cfg.With(typeSys).
		With(sourceSys).
		With(handleSys)

	world := akara.NewWorld(cfg)

	filePaths := typeSys.FilePathFactory
	fileHandles := handleSys.FileHandleFactory

	sourceEntity := world.NewEntity()
	source := filePaths.AddFilePath(sourceEntity)
	source.Path = testDataPath

	fileEntity := world.NewEntity()
	file := filePaths.AddFilePath(fileEntity)
	file.Path = "testfile_a.txt"

	_ = world.Update(0)

	ft, found := typeSys.GetFileType(sourceEntity)
	if !found {
		t.Error("file source type not created for entity")
		return
	}

	if ft.Type != d2enum.FileTypeDirectory {
		t.Error("expected file system source type for entity")
		return
	}

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
