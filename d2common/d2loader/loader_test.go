package d2loader

import (
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const (
	sourcePathA   = "testdata/A"
	sourcePathB   = "testdata/B"
	sourcePathC   = "testdata/C"
	sourcePathD   = "testdata/D.mpq"
	commonFile    = "common.txt"
	exclusiveA    = "exclusive_a.txt"
	exclusiveB    = "exclusive_b.txt"
	exclusiveC    = "exclusive_c.txt"
	exclusiveD    = "exclusive_d.txt"
	subdirCommonD = "dir\\common.txt"
	badSourcePath = "/x/y/z.mpq"
	badFilePath   = "a/bad/file/path.txt"
)

func TestLoader_NewLoader(t *testing.T) {
	loader, _ := NewLoader(d2util.LogLevelDefault)

	if loader.Cache == nil {
		t.Error("loader should not be nil")
	}
}

func TestLoader_AddSource(t *testing.T) {
	loader, _ := NewLoader(d2util.LogLevelDefault)

	errA := loader.AddSource(sourcePathA, types.AssetSourceFileSystem)
	errB := loader.AddSource(sourcePathB, types.AssetSourceFileSystem)
	errC := loader.AddSource(sourcePathC, types.AssetSourceFileSystem)
	errD := loader.AddSource(sourcePathD, types.AssetSourceFileSystem)
	errE := loader.AddSource(badSourcePath, types.AssetSourceMPQ)

	if errA != nil {
		t.Error(errA)
	}

	if errB != nil {
		t.Error(errB)
	}

	if errC != nil {
		t.Error(errC)
	}

	if errD != nil {
		t.Error(errD)
	}

	if errE == nil {
		t.Error("expecting error on bad file path")
	}

}

// nolint:gocyclo // this is just a test, not a big deal if we ignore linter here
func TestLoader_Load(t *testing.T) {
	loader, _ := NewLoader(d2util.LogLevelDefault)

	// we expect files common to any source to come from here
	err := loader.AddSource(sourcePathB, types.AssetSourceFileSystem)
	if err != nil {
		t.Fail()
		log.Print(err)
	}

	err = loader.AddSource(sourcePathD, types.AssetSourceMPQ)
	if err != nil {
		t.Fail()
		log.Print(err)
	}

	err = loader.AddSource(sourcePathA, types.AssetSourceFileSystem)
	if err != nil {
		t.Fail()
		log.Print(err)
	}

	err = loader.AddSource(sourcePathC, types.AssetSourceFileSystem)
	if err != nil {
		t.Fail()
		log.Print(err)
	}

	entryCommon, errCommon := loader.Load(commonFile) // common file exists in all three Sources

	entryA, errA := loader.Load(exclusiveA) // each source has a file exclusive to itself
	entryB, errB := loader.Load(exclusiveB)
	entryC, errC := loader.Load(exclusiveC)
	entryD, errD := loader.Load(exclusiveD)
	entryDsubdir, errDsubdir := loader.Load(subdirCommonD)

	_, expectedError := loader.Load(badFilePath) // we expect an Error for this bad file path

	if entryCommon == nil || errCommon != nil {
		t.Error("common entry should exist")
	}

	if errA != nil || errB != nil || errC != nil || errD != nil {
		t.Error("files exclusive to each source don't exist")
	}

	if errDsubdir != nil {
		t.Error("mpq subdir entry not found")
	}

	if expectedError == nil {
		t.Error("expected Error for nonexistant file path")
	}

	var result []byte

	buffer := make([]byte, 1)

	tests := []struct {
		entry io.ReadSeeker
		data  string
	}{
		{entryCommon, "b"}, // sourcePathB is loaded first, we expect a "b"
		{entryA, "a"},
		{entryB, "b"},
		{entryC, "c"},
		{entryD, "d"},
		{entryDsubdir, "d"},
	}

	for idx := range tests {
		entry, expected := tests[idx].entry, tests[idx].data

		result = make([]byte, 0)

		for {
			if bytesRead, err := entry.Read(buffer); err != nil || bytesRead == 0 {
				break
			}

			result = append(result, buffer...)
		}

		got := string(result[0])

		if got != expected {
			fmtStr := "unexpected data in file, expected %q, got %q"
			msg := fmt.Sprintf(fmtStr, expected, got)
			t.Error(msg)
		}
	}
}
