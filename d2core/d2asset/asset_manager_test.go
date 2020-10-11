package d2asset

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader"
)

func TestAssetManager_LoadFile_NoSource(t *testing.T) {
	am := &AssetManager{
		loader:     d2loader.NewLoader(nil),
		tables:     d2cache.CreateCache(tableBudget),
		animations: d2cache.CreateCache(animationBudget),
		fonts:      d2cache.CreateCache(fontBudget),
		palettes:   d2cache.CreateCache(paletteBudget),
		transforms: d2cache.CreateCache(paletteTransformBudget),
	}

	_, err := am.LoadFile("an/invalid/path")
	if err == nil {
		t.Error("asset manager loaded a file for which there is no source")
	}
}

func BenchmarkAssetManager_LoadFile_NoSource(b *testing.B) {
	am := &AssetManager{
		loader:     d2loader.NewLoader(nil),
		tables:     d2cache.CreateCache(tableBudget),
		animations: d2cache.CreateCache(animationBudget),
		fonts:      d2cache.CreateCache(fontBudget),
		palettes:   d2cache.CreateCache(paletteBudget),
		transforms: d2cache.CreateCache(paletteTransformBudget),
	}

	for idx := 0; idx < b.N; idx++ {
		_, _ = am.LoadFile("an/invalid/path")
	}
}
