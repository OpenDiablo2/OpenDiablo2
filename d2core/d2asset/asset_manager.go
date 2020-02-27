package d2asset

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
)

var (
	ErrWasInit = errors.New("asset system is already initialized")
	ErrNotInit = errors.New("asset system is not initialized")
)

type assetManager struct {
	archiveManager			*archiveManager
	fileManager				*fileManager
	paletteManager			*paletteManager
	paletteTransformManager	*paletteTransformManager
	animationManager		*animationManager
	fontManager				*fontManager
}

func loadDC6(dc6Path string) (*d2dc6.DC6File, error) {
	dc6Data, err := LoadFile(dc6Path)
	if err != nil {
		return nil, err
	}

	dc6, err := d2dc6.LoadDC6(dc6Data)
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

	return d2dcc.LoadDCC(dccData)
}

func loadCOF(cofPath string) (*d2cof.COF, error) {
	cofData, err := LoadFile(cofPath)
	if err != nil {
		return nil, err
	}

	return d2cof.LoadCOF(cofData)
}
