package d2dcc

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type DCC struct {
	Signature          int
	Version            int
	NumberOfDirections int
	FramesPerDirection int
	Directions         []*DCCDirection
}

func LoadDCC(fileData []byte) (*DCC, error) {
	result := &DCC{}
	var bm = d2common.CreateBitMuncher(fileData, 0)
	result.Signature = int(bm.GetByte())
	if result.Signature != 0x74 {
		return nil, errors.New("signature expected to be 0x74 but it is not")
	}
	result.Version = int(bm.GetByte())
	result.NumberOfDirections = int(bm.GetByte())
	result.FramesPerDirection = int(bm.GetInt32())
	if bm.GetInt32() != 1 {
		return nil, errors.New("this value isn't 1. It has to be 1")
	}
	bm.GetInt32() // TotalSizeCoded
	directionOffsets := make([]int, result.NumberOfDirections)
	for i := 0; i < result.NumberOfDirections; i++ {
		directionOffsets[i] = int(bm.GetInt32())
	}
	result.Directions = make([]*DCCDirection, result.NumberOfDirections)
	for i := 0; i < result.NumberOfDirections; i++ {
		result.Directions[i] = CreateDCCDirection(d2common.CreateBitMuncher(fileData, directionOffsets[i]*8), *result)
	}
	return result, nil
}
