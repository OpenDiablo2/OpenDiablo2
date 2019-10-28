package Common

import (
	"log"
	"strings"

	"github.com/essial/OpenDiablo2/ResourcePaths"
)

type LevelPresetRecord struct {
	DefinitionId int32
	LevelId      int32
	Populate     bool
	Logicals     bool
	Outdoors     bool
	Animate      bool
	KillEdge     bool
	FillBlanks   bool
	SizeX        int32
	SizeY        int32
	AutoMap      bool
	Scan         bool
	Pops         int32
	PopPad       int32
	Files        [6]string
	Dt1Mask      uint32
}

var LevelPresets []LevelPresetRecord

func LoadLevelPresets(fileProvider FileProvider) {
	LevelPresets = make([]LevelPresetRecord, 0)
	levelTypesData := fileProvider.LoadFile(ResourcePaths.LevelPreset)
	sr := CreateStreamReader(levelTypesData)
	numRecords := sr.GetInt32()
	LevelPresets = make([]LevelPresetRecord, numRecords)
	for i := range LevelPresets {
		LevelPresets[i].DefinitionId = sr.GetInt32()
		LevelPresets[i].LevelId = sr.GetInt32()
		LevelPresets[i].Populate = sr.GetInt32() != 0
		LevelPresets[i].Logicals = sr.GetInt32() != 0
		LevelPresets[i].Outdoors = sr.GetInt32() != 0
		LevelPresets[i].Animate = sr.GetInt32() != 0
		LevelPresets[i].KillEdge = sr.GetInt32() != 0
		LevelPresets[i].FillBlanks = sr.GetInt32() != 0
		LevelPresets[i].SizeX = sr.GetInt32()
		LevelPresets[i].SizeY = sr.GetInt32()
		LevelPresets[i].AutoMap = sr.GetInt32() != 0
		LevelPresets[i].Scan = sr.GetInt32() != 0
		LevelPresets[i].Pops = sr.GetInt32()
		LevelPresets[i].PopPad = sr.GetInt32()
		sr.GetInt32()
		for fileIdx := 0; fileIdx < 6; fileIdx++ {
			strData, _ := sr.ReadBytes(60)
			s := strings.Trim(string(strData), string(0))
			if s == "0" {
				LevelPresets[i].Files[fileIdx] = ""
			} else {
				LevelPresets[i].Files[fileIdx] = s
			}

		}
		LevelPresets[i].Dt1Mask = sr.GetUInt32()
	}
	log.Printf("Loaded %d LevelPreset records")
}
