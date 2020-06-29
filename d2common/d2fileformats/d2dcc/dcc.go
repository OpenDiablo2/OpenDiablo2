// Package d2dcc contains the logic for loading and processing DCC files.
package d2dcc

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const dccFileSignature = 0x74
const directionOffsetMultiplier = 8

// DCC represents a DCC file.
type DCC struct {
	Signature          int
	Version            int
	NumberOfDirections int
	FramesPerDirection int
	Directions         []*DCCDirection
}

// Load loads a DCC file.
func Load(fileData []byte) (*DCC, error) {
	result := &DCC{}

	var bm = d2common.CreateBitMuncher(fileData, 0)

	result.Signature = int(bm.GetByte())

	if result.Signature != dccFileSignature {
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
		result.Directions[i] = CreateDCCDirection(d2common.CreateBitMuncher(
			fileData, directionOffsets[i]*directionOffsetMultiplier), result)
	}

	return result, nil
}
