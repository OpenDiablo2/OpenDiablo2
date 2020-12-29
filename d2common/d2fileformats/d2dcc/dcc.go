package d2dcc

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
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
	directionOffsets   []int
	fileData           []byte
}

// Load loads a DCC file.
func Load(fileData []byte) (*DCC, error) {
	result := &DCC{
		fileData: fileData,
	}

	var bm = d2datautils.CreateBitMuncher(fileData, 0)

	result.Signature = int(bm.GetByte())

	if result.Signature != dccFileSignature {
		return nil, errors.New("signature expected to be 0x74 but it is not")
	}

	result.Version = int(bm.GetByte())
	result.NumberOfDirections = int(bm.GetByte())
	result.FramesPerDirection = int(bm.GetInt32())

	result.Directions = make([]*DCCDirection, result.NumberOfDirections)

	if bm.GetInt32() != 1 {
		return nil, errors.New("this value isn't 1. It has to be 1")
	}

	bm.GetInt32() // TotalSizeCoded

	result.directionOffsets = make([]int, result.NumberOfDirections)

	for i := 0; i < result.NumberOfDirections; i++ {
		result.directionOffsets[i] = int(bm.GetInt32())
		result.Directions[i] = result.decodeDirection(i)
	}

	return result, nil
}

// decodeDirection decodes and returns the given direction
func (d *DCC) decodeDirection(direction int) *DCCDirection {
	return CreateDCCDirection(d2datautils.CreateBitMuncher(d.fileData,
		d.directionOffsets[direction]*directionOffsetMultiplier), d)
}

// Clone creates a copy of the DCC
func (d *DCC) Clone() *DCC {
	clone := *d
	copy(clone.directionOffsets, d.directionOffsets)
	copy(clone.fileData, d.fileData)
	clone.Directions = make([]*DCCDirection, len(d.Directions))

	for i := range d.Directions {
		cloneDirection := *d.Directions[i]
		clone.Directions = append(clone.Directions, &cloneDirection)
	}

	return &clone
}
