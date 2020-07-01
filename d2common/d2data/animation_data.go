package d2data

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
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
var AnimationData map[string][]*AnimationDataRecord //nolint:gochecknoglobals // Currently global by design

// LoadAnimationData loads the animation data table into the global AnimationData dictionary
func LoadAnimationData(rawData []byte) {
	AnimationData = make(map[string][]*AnimationDataRecord)
	streamReader := d2common.CreateStreamReader(rawData)

	for !streamReader.Eof() {
		dataCount := int(streamReader.GetInt32())
		for i := 0; i < dataCount; i++ {
			cofNameBytes := streamReader.ReadBytes(8)
			data := &AnimationDataRecord{
				COFName:            strings.ReplaceAll(string(cofNameBytes), string(0), ""),
				FramesPerDirection: int(streamReader.GetInt32()),
				AnimationSpeed:     int(streamReader.GetInt32()),
			}
			data.Flags = streamReader.ReadBytes(144)
			cofIndex := strings.ToLower(data.COFName)

			if _, found := AnimationData[cofIndex]; !found {
				AnimationData[cofIndex] = make([]*AnimationDataRecord, 0)
			}

			AnimationData[cofIndex] = append(AnimationData[cofIndex], data)
		}
	}

	log.Printf("Loaded %d animation data records", len(AnimationData))
}
