package d2loader

import (
	"fmt"
	"log"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
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

	sourceA, errA := loader.AddSource(sourcePathA)
	sourceB, errB := loader.AddSource(sourcePathB)
	sourceC, errC := loader.AddSource(sourcePathC)
	sourceD, errD := loader.AddSource(sourcePathD)
	sourceE, errE := loader.AddSource(badSourcePath)

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

	if sourceA.String() != sourcePathA {
		t.Error("source path not the same as what we added")
	}

	if sourceB.String() != sourcePathB {
		t.Error("source path not the same as what we added")
	}

	if sourceC.String() != sourcePathC {
		t.Error("source path not the same as what we added")
	}

	if sourceD.String() != sourcePathD {
		t.Error("source path not the same as what we added")
	}

	if sourceE != nil {
		t.Error("source for bad path should be nil")
	}
}

// nolint:gocyclo // this is just a test, not a big deal if we ignore linter here
func TestLoader_Load(t *testing.T) {
	loader, _ := NewLoader(d2util.LogLevelDefault)

	// we expect files common to any source to come from here
	commonSource, err := loader.AddSource(sourcePathB)
	if err != nil {
		t.Fail()
		log.Print(err)
	}

	_, err = loader.AddSource(sourcePathD)
	if err != nil {
		t.Fail()
		log.Print(err)
	}

	_, err = loader.AddSource(sourcePathA)
	if err != nil {
		t.Fail()
		log.Print(err)
	}

	_, err = loader.AddSource(sourcePathC)
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
	} else if entryCommon.Source() != commonSource {
		t.Error("common entry should come from the first loader source")
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
		entry asset.Asset
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
			fmtStr := "unexpected data in file %s, loaded from source `%s`: expected `%s`, got `%s`"
			msg := fmt.Sprintf(fmtStr, entry.Path(), entry.Source(), expected, got)
			t.Error(msg)
		}
	}
}
