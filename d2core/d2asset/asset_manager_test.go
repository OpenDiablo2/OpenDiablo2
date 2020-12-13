package d2asset

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

func TestAssetManager_LoadFile_NoSource(t *testing.T) {
	loader, err := d2loader.NewLoader(d2util.LogLevelDefault)
	if err != nil {
		t.Error(err)
	}

	records, err := d2records.NewRecordManager(d2util.LogLevelDebug)
	if err != nil {
		t.Error(err)
	}

	am := &AssetManager{
		Logger:     d2util.NewLogger(),
		Loader:     loader,
		tables:     make([]d2tbl.TextDictionary, 0),
		animations: d2cache.CreateCache(animationBudget),
		fonts:      d2cache.CreateCache(fontBudget),
		palettes:   d2cache.CreateCache(paletteBudget),
		transforms: d2cache.CreateCache(paletteTransformBudget),
		Records:    records,
	}

	_, err = am.LoadFile("an/invalid/path")
	if err == nil {
		t.Error("asset manager loaded a file for which there is no source")
	}
}

func BenchmarkAssetManager_LoadFile_NoSource(b *testing.B) {
	loader, err := d2loader.NewLoader(d2util.LogLevelDefault)
	if err != nil {
		b.Error(err)
	}

	records, err := d2records.NewRecordManager(d2util.LogLevelDebug)
	if err != nil {
		b.Error(err)
	}

	am := &AssetManager{
		Logger:     d2util.NewLogger(),
		Loader:     loader,
		tables:     make([]d2tbl.TextDictionary, 0),
		animations: d2cache.CreateCache(animationBudget),
		fonts:      d2cache.CreateCache(fontBudget),
		palettes:   d2cache.CreateCache(paletteBudget),
		transforms: d2cache.CreateCache(paletteTransformBudget),
		Records:    records,
	}

	for idx := 0; idx < b.N; idx++ {
		_, _ = am.LoadFile("an/invalid/path")
	}
}
