package Common

import (
	"log"
	"strings"

	"github.com/essial/OpenDiablo2/ResourcePaths"
)

type LevelTypeRecord struct {
	Files [32]string
	Act   int32
}

var LevelTypes []LevelTypeRecord

func LoadLevelTypes(fileProvider FileProvider) {
	levelTypesData := fileProvider.LoadFile(ResourcePaths.LevelType)
	sr := CreateStreamReader(levelTypesData)
	numRecords := sr.GetInt32()
	LevelTypes = make([]LevelTypeRecord, numRecords)
	for i := range LevelTypes {
		for fileIdx := 0; fileIdx < 32; fileIdx++ {
			strData, _ := sr.ReadBytes(60)
			s := strings.Trim(string(strData), string(0))
			if s == "0" {
				LevelTypes[i].Files[fileIdx] = ""
			} else {
				LevelTypes[i].Files[fileIdx] = s
			}

		}
		LevelTypes[i].Act = int32(sr.GetByte())

	}
	log.Printf("Loaded %d LevelType records", numRecords)
}
