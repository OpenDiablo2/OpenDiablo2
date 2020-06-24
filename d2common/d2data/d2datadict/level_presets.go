package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type LevelPresetRecord struct {
	Name         string
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
	Beta         bool
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
		Name:         props[inc()],
		DefinitionId: d2common.StringToInt(props[inc()]),
		LevelId:      d2common.StringToInt(props[inc()]),
		Populate:     d2common.StringToUint8(props[inc()]) == 1,
		Logicals:     d2common.StringToUint8(props[inc()]) == 1,
		Outdoors:     d2common.StringToUint8(props[inc()]) == 1,
		Animate:      d2common.StringToUint8(props[inc()]) == 1,
		KillEdge:     d2common.StringToUint8(props[inc()]) == 1,
		FillBlanks:   d2common.StringToUint8(props[inc()]) == 1,
		SizeX:        d2common.StringToInt(props[inc()]),
		SizeY:        d2common.StringToInt(props[inc()]),
		AutoMap:      d2common.StringToUint8(props[inc()]) == 1,
		Scan:         d2common.StringToUint8(props[inc()]) == 1,
		Pops:         d2common.StringToInt(props[inc()]),
		PopPad:       d2common.StringToInt(props[inc()]),
		FileCount:    d2common.StringToInt(props[inc()]),
		Files: [6]string{
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
		},
		Dt1Mask:   d2common.StringToUint(props[inc()]),
		Beta:      d2common.StringToUint8(props[inc()]) == 1,
		Expansion: d2common.StringToUint8(props[inc()]) == 1,
	}
	return result
}

var LevelPresets map[int]LevelPresetRecord

func LoadLevelPresets(file []byte) {
	LevelPresets = make(map[int]LevelPresetRecord)
	data := strings.Split(string(file), "\r\n")[1:]
	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		props := strings.Split(line, "\t")
		if props[1] == "" {
			continue // any line without a definition id is skipped (e.g. the "Expansion" line)
		}
		rec := createLevelPresetRecord(props)
		LevelPresets[rec.DefinitionId] = rec
	}
	log.Printf("Loaded %d level presets", len(LevelPresets))
}

func LevelPreset(id int) LevelPresetRecord {
	for i := 0; i < len(LevelPresets); i++ {
		if LevelPresets[i].DefinitionId == id {
			return LevelPresets[i]
		}
	}
	panic("Unknown level preset")
}
