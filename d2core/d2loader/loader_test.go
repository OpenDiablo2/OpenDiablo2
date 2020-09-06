package d2loader

import (
	"fmt"
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2loader/asset"
)

const (
	sourceA     = "testdata/A"
	sourceB     = "testdata/B"
	sourceC     = "testdata/C"
	sourceD     = "testdata/D.mpq"
	commonFile  = "common.txt"
	exclusiveA  = "exclusive_a.txt"
	exclusiveB  = "exclusive_b.txt"
	exclusiveC  = "exclusive_c.txt"
	exclusiveD  = "exclusive_d.txt"
	badFilePath = "a/bad/file/path.txt"
)

func TestLoader_NewLoader(t *testing.T) {
	loader := NewLoader()

	if loader.Cache == nil {
		t.Error("loader should not be nil")
	}

	if loader.Cache.entries == nil {
		t.Error("loader cache should not be nil")
	}

	if &loader.entries != &loader.Cache.entries {
		t.Error("loader cache should be embedded")
	}
}

func TestLoader_AddSource(t *testing.T) {
	loader := NewLoader()

	loader.AddSource(sourceA)
	loader.AddSource(sourceB)
	loader.AddSource(sourceC)
	loader.AddSource(sourceD)
	loader.AddSource("bad/path")

	if loader.sources[0].String() != sourceA {
		t.Error("source path not the same as what we added")
	}

	if loader.sources[1].String() != sourceB {
		t.Error("source path not the same as what we added")
	}

	if loader.sources[2].String() != sourceC {
		t.Error("source path not the same as what we added")
	}

	if loader.sources[3].String() != sourceD {
		t.Error("source path not the same as what we added")
	}
}

func TestLoader_Load(t *testing.T) {
	loader := NewLoader()

	loader.AddSource(sourceB) // we expect files common to any source to come from here
	loader.AddSource(sourceD)
	loader.AddSource(sourceA)
	loader.AddSource(sourceC)

	entryCommon, errCommon := loader.Load(commonFile) // common file exists in all three sources

	entryA, errA := loader.Load(exclusiveA) // each source has a file exclusive to itself
	entryB, errB := loader.Load(exclusiveB)
	entryC, errC := loader.Load(exclusiveC)
	entryD, errD := loader.Load(exclusiveD)

	_, expectedError := loader.Load(badFilePath) // we expect an Error for this bad file path

	if entryCommon == nil || errCommon != nil {
		t.Error("common entry should exist")
	} else if entryCommon.Source() != loader.sources[0] {
		t.Error("common entry should come from the first loader source")
	}

	if errA != nil || errB != nil || errC != nil || errD != nil {
		t.Error("files exclusive to each source don't exist")
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
		{entryCommon, "b"}, // sourceB is loaded first, we expect a "b"
		{entryA, "a"},
		{entryB, "b"},
		{entryC, "c"},
		{entryD, "d"},
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
