package Common

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/ResourcePaths"
)

type LevelPresetRecord struct {
	Name	     string
	DefinitionId int
	LevelId      int
	Populate     bool
	Logicals     bool
	Outdoors     bool
	Animate      bool
	KillEdge     bool
	FillBlanks   bool
	SizeX        int
	SizeY        int
	AutoMap      bool
	Scan         bool
	Pops         int
	PopPad       int
	FileCount    int
	Files        [6]string
	Dt1Mask      uint
	Beta		 bool
	Expansion    bool
}

// CreateLevelPresetRecord parses a row from lvlprest.txt into a LevelPresetRecord
func createLevelPresetRecord(props []string) LevelPresetRecord {
	i := -1
	inc := func() int {
		i++
		return i
	}
	result := LevelPresetRecord{
		Name: props[inc()],
		DefinitionId: StringToInt(props[inc()]),
		LevelId: StringToInt(props[inc()]),
		Populate: StringToUint8(props[inc()]) == 1,
		Logicals: StringToUint8(props[inc()]) == 1,
		Outdoors: StringToUint8(props[inc()]) == 1,
		Animate: StringToUint8(props[inc()]) == 1,
		KillEdge: StringToUint8(props[inc()]) == 1,
		FillBlanks: StringToUint8(props[inc()]) == 1,
		SizeX: StringToInt(props[inc()]),
		SizeY: StringToInt(props[inc()]),
		AutoMap: StringToUint8(props[inc()]) == 1,
		Scan: StringToUint8(props[inc()]) == 1,
		Pops: StringToInt(props[inc()]),
		PopPad: StringToInt(props[inc()]),
		FileCount: StringToInt(props[inc()]),
		Files: [6]string{
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
		},
		Dt1Mask: StringToUint(props[inc()]),
		Beta: StringToUint8(props[inc()]) == 1,
		Expansion: StringToUint8(props[inc()]) == 1,
	}
	return result
}

var LevelPresets map[int]*LevelPresetRecord

func LoadLevelPresets(fileProvider FileProvider) {
	LevelPresets = make(map[int]*LevelPresetRecord)
	data := strings.Split(string(fileProvider.LoadFile(ResourcePaths.LevelPreset)), "\r\n")[1:]
	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		props := strings.Split(line, "\t")
		if(props[1] == "") {
			continue // any line without a definition id is skipped (e.g. the "Expansion" line)
		}
		rec := createLevelPresetRecord(props)
		LevelPresets[rec.DefinitionId] = &rec
	}
	log.Printf("Loaded %d level presets", len(LevelPresets))
}