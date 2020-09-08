package d2animdata

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const (
	numBlocks             = 256
	maxRecordsPerBlock    = 67
	byteCountName         = 8
	byteCountSpeedPadding = 2
	numActions            = 144
)

type record struct {
	name               string
	framesPerDirection uint32
	speed              uint16
	actions            map[int]byte
}

type hashTable [numBlocks]byte

type block struct {
	recordCount uint32
	records     []*record
}

type AnimationDataSet struct {
	hashTable
	blocks [numBlocks]*block
}

// Load loads the data into an AnimationDataSet struct
func Load(data []byte) (*AnimationDataSet, error) {
	reader := d2common.CreateStreamReader(data)
	animdata := &AnimationDataSet{}
	hashIdx := 0

	for blockIdx := range animdata.blocks {
		recordCount := reader.GetUInt32()
		if recordCount > maxRecordsPerBlock {
			return nil, fmt.Errorf("more than %d records in block", maxRecordsPerBlock)
		}

		records := make([]*record, recordCount)

		for recordIdx := uint32(0); recordIdx < recordCount; recordIdx++ {
			nameBytes := reader.ReadBytes(byteCountName)

			if nameBytes[7] != byte(0) {
				return nil, errors.New("animdata record name missing null terminator byte")
			}

			name := string(nameBytes)
			name = strings.ReplaceAll(name, string(byte(0)), "")

			animdata.hashTable[hashIdx] = hashName(name)

			frames := reader.GetUInt32()
			speed := reader.GetUInt16()

			reader.SkipBytes(byteCountSpeedPadding)

			actions := make(map[int]byte)

			for actionIdx := 0; actionIdx < numActions; actionIdx++ {
				actionByte := reader.GetByte()
				if actionByte != 0 {
					actions[actionIdx] = actionByte
				}
			}

			r := &record{
				name,
				frames,
				speed,
				actions,
			}

			records[recordIdx] = r
		}

		b := &block{
			recordCount,
			records,
		}

		animdata.blocks[blockIdx] = b
	}

	if reader.GetPosition() != uint64(len(data)) {
		return nil, errors.New("unable to parse animation data")
	}

	return animdata, nil
}

func hashName(name string) byte {
	hashBytes := []byte(strings.ToUpper(name))

	var hash uint32
	for hashByteIdx := range hashBytes {
		hash += uint32(hashBytes[hashByteIdx])
	}

	hash %= numBlocks
	return byte(hash)
}
