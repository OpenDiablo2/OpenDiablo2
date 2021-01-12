package d2data

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

const (
	numCofNameBytes = 8
	numFlagBytes    = 144
)

// AnimationDataRecord represents a single entry in the animation data dictionary file
type AnimationDataRecord struct {
	// COFName is the name of the COF file used for this animation
	COFName string
	// FramesPerDirection specifies how many frames are in each direction
	FramesPerDirection int
	// AnimationSpeed represents a value of X where the rate is a ration of (x/255) at 25FPS
	AnimationSpeed int
	// Flags are used in keyframe triggers
	Flags []byte
}

// AnimationData represents all of the animation data records, mapped by the COF index
type AnimationData map[string][]*AnimationDataRecord

// LoadAnimationData loads the animation data table into the global AnimationData dictionary
func LoadAnimationData(rawData []byte) (AnimationData, error) {
	animdata := make(AnimationData)
	streamReader := d2datautils.CreateStreamReader(rawData)

	for !streamReader.EOF() {
		var dataCount int

		b, err := streamReader.ReadInt32()
		if err != nil {
			return nil, err
		}

		dataCount = int(b)

		for i := 0; i < dataCount; i++ {
			cofNameBytes, err := streamReader.ReadBytes(numCofNameBytes)
			if err != nil {
				return nil, err
			}

			fpd, err := streamReader.ReadInt32()
			if err != nil {
				return nil, err
			}

			animSpeed, err := streamReader.ReadInt32()
			if err != nil {
				return nil, err
			}

			data := &AnimationDataRecord{
				COFName:            strings.ReplaceAll(string(cofNameBytes), string(byte(0)), ""),
				FramesPerDirection: int(fpd),
				AnimationSpeed:     int(animSpeed),
			}

			flagBytes, err := streamReader.ReadBytes(numFlagBytes)
			if err != nil {
				return nil, err
			}

			data.Flags = flagBytes
			cofIndex := strings.ToLower(data.COFName)

			if _, found := animdata[cofIndex]; !found {
				animdata[cofIndex] = make([]*AnimationDataRecord, 0)
			}

			animdata[cofIndex] = append(animdata[cofIndex], data)
		}
	}

	return animdata, nil
}
