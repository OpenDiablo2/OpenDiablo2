package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type assetManager struct {
	archiveManager          d2interface.ArchiveManager
	archivedFileManager     d2interface.ArchivedFileManager
	paletteManager          d2interface.ArchivedPaletteManager
	paletteTransformManager *paletteTransformManager
	animationManager        d2interface.ArchivedAnimationManager
	fontManager             d2interface.ArchivedFontManager
}

func loadDC6(dc6Path string) (*d2dc6.DC6, error) {
	dc6Data, err := LoadFile(dc6Path)
	if err != nil {
		return nil, err
	}

	dc6, err := d2dc6.Load(dc6Data)
	if err != nil {
		return nil, err
	}

	return dc6, nil
}

func loadDCC(dccPath string) (*d2dcc.DCC, error) {
	dccData, err := LoadFile(dccPath)
	if err != nil {
		return nil, err
	}

	return d2dcc.Load(dccData)
}

func loadCOF(cofPath string) (*d2cof.COF, error) {
	cofData, err := LoadFile(cofPath)
	if err != nil {
		return nil, err
	}

	return d2cof.Load(cofData)
}

func (am *assetManager) BindTerminalCommands(term d2interface.Terminal) error {
	if err := term.BindAction("assetspam", "display verbose asset manager logs", func(verbose bool) {
		if verbose {
			term.OutputInfof("asset manager verbose logging enabled")
		} else {
			term.OutputInfof("asset manager verbose logging disabled")
		}

		am.archiveManager.GetCache().SetVerbose(verbose)
		am.archivedFileManager.GetCache().SetVerbose(verbose)
		am.paletteManager.GetCache().SetVerbose(verbose)
		am.paletteTransformManager.cache.SetVerbose(verbose)
		am.animationManager.GetCache().SetVerbose(verbose)
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetstat", "display asset manager cache statistics", func() {
		var cacheStatistics = func(c d2interface.Cache) float64 {
			const percent = 100.0
			return float64(c.GetWeight()) / float64(c.GetBudget()) * percent
		}

		term.OutputInfof("archive cache: %f", cacheStatistics(am.archiveManager.GetCache()))
		term.OutputInfof("file cache: %f", cacheStatistics(am.archivedFileManager.GetCache()))
		term.OutputInfof("palette cache: %f", cacheStatistics(am.paletteManager.GetCache()))
		term.OutputInfof("palette transform cache: %f", cacheStatistics(am.paletteTransformManager.
			cache))
		term.OutputInfof("animation cache: %f", cacheStatistics(am.animationManager.GetCache()))
		term.OutputInfof("font cache: %f", cacheStatistics(am.fontManager.GetCache()))
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetclear", "clear asset manager cache", func() {
		am.archiveManager.ClearCache()
		am.archivedFileManager.GetCache().Clear()
		am.paletteManager.ClearCache()
		am.paletteTransformManager.cache.Clear()
		am.animationManager.ClearCache()
		am.fontManager.ClearCache()
	}); err != nil {
		return err
	}

	return nil
}
