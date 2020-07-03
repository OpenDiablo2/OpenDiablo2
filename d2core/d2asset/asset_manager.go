package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type assetManager struct {
	archiveManager          *archiveManager
	fileManager             *fileManager
	paletteManager          *paletteManager
	paletteTransformManager *paletteTransformManager
	animationManager        *animationManager
	fontManager             *fontManager
}

func (a assetManager) Advance(elapsedTime, currentTime float64) error {
	panic("implement me")
}

func (a assetManager) Render(surface d2interface.Surface) error {
	panic("implement me")
}

func (a assetManager) Initialize() error {
	panic("implement me")
}

func (a assetManager) BindApp(app d2interface.App) error {
	panic("implement me")
}

func (a assetManager) UnbindApp(app d2interface.App) error {
	panic("implement me")
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
