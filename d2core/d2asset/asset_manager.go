package d2asset

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
)

var (
	ErrHasInit = errors.New("asset system is already initialized")
	ErrNoInit  = errors.New("asset system is not initialized")
)

type assetManager struct {
	archiveManager   *archiveManager
	fileManager      *fileManager
	paletteManager   *paletteManager
	animationManager *animationManager
}

func loadPalette(palettePath string) (*d2datadict.PaletteRec, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.paletteManager.loadPalette(palettePath)
}

func loadDC6(dc6Path, palettePath string) (*d2dc6.DC6File, error) {
	dc6Data, err := LoadFile(dc6Path)
	if err != nil {
		return nil, err
	}

	paletteData, err := loadPalette(palettePath)
	if err != nil {
		return nil, err
	}

	dc6, err := d2dc6.LoadDC6(dc6Data, *paletteData)
	if err != nil {
		return nil, err
	}

	return &dc6, nil
}

func loadDCC(dccPath string) (*d2dcc.DCC, error) {
	dccData, err := LoadFile(dccPath)
	if err != nil {
		return nil, err
	}

	return d2dcc.LoadDCC(dccData)
}

func loadCOF(cofPath string) (*d2cof.COF, error) {
	cofData, err := LoadFile(cofPath)
	if err != nil {
		return nil, err
	}

	return d2cof.LoadCOF(cofData)
}
