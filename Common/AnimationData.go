package Common

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/ResourcePaths"
)

type AnimationDataRecord struct {
	COFName            string
	FramesPerDirection int
	AnimationSpeed     int
	Flags              []byte
}

var AnimationData map[string][]*AnimationDataRecord

func LoadAnimationData(fileProvider FileProvider) {
	AnimationData = make(map[string][]*AnimationDataRecord)
	rawData := fileProvider.LoadFile(ResourcePaths.AnimationData)
	streamReader := CreateStreamReader(rawData)
	for !streamReader.Eof() {
		dataCount := int(streamReader.GetInt32())
		for i := 0; i < dataCount; i++ {
			cofNameBytes, _ := streamReader.ReadBytes(8)
			data := &AnimationDataRecord{
				COFName:            strings.ReplaceAll(string(cofNameBytes), string(0), ""),
				FramesPerDirection: int(streamReader.GetInt32()),
				AnimationSpeed:     int(streamReader.GetInt32()),
			}
			data.Flags, _ = streamReader.ReadBytes(144)
			cofIndex := strings.ToLower(data.COFName)
			if _, found := AnimationData[cofIndex]; !found {
				AnimationData[cofIndex] = make([]*AnimationDataRecord, 0)
			}
			AnimationData[cofIndex] = append(AnimationData[cofIndex], data)
		}
	}
	log.Printf("Loaded %d animation data records", len(AnimationData))
}
