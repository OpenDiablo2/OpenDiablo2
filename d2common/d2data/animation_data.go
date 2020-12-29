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
func LoadAnimationData(rawData []byte) AnimationData {
	animdata := make(AnimationData)
	streamReader := d2datautils.CreateStreamReader(rawData)

	for !streamReader.EOF() {
		dataCount := int(streamReader.GetInt32())
		for i := 0; i < dataCount; i++ {
			cofNameBytes := streamReader.ReadBytes(numCofNameBytes)
			data := &AnimationDataRecord{
				COFName:            strings.ReplaceAll(string(cofNameBytes), string(byte(0)), ""),
				FramesPerDirection: int(streamReader.GetInt32()),
				AnimationSpeed:     int(streamReader.GetInt32()),
			}
			data.Flags = streamReader.ReadBytes(numFlagBytes)
			cofIndex := strings.ToLower(data.COFName)

			if _, found := animdata[cofIndex]; !found {
				animdata[cofIndex] = make([]*AnimationDataRecord, 0)
			}

			animdata[cofIndex] = append(animdata[cofIndex], data)
		}
	}

	return animdata
}
