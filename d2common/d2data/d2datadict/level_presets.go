package d2datadict

import (
	"log"
	"strings"

	dh "github.com/OpenDiablo2/OpenDiablo2/d2common"
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
		DefinitionId: dh.StringToInt(props[inc()]),
		LevelId:      dh.StringToInt(props[inc()]),
		Populate:     dh.StringToUint8(props[inc()]) == 1,
		Logicals:     dh.StringToUint8(props[inc()]) == 1,
		Outdoors:     dh.StringToUint8(props[inc()]) == 1,
		Animate:      dh.StringToUint8(props[inc()]) == 1,
		KillEdge:     dh.StringToUint8(props[inc()]) == 1,
		FillBlanks:   dh.StringToUint8(props[inc()]) == 1,
		SizeX:        dh.StringToInt(props[inc()]),
		SizeY:        dh.StringToInt(props[inc()]),
		AutoMap:      dh.StringToUint8(props[inc()]) == 1,
		Scan:         dh.StringToUint8(props[inc()]) == 1,
		Pops:         dh.StringToInt(props[inc()]),
		PopPad:       dh.StringToInt(props[inc()]),
		FileCount:    dh.StringToInt(props[inc()]),
		Files: [6]string{
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
		},
		Dt1Mask:   dh.StringToUint(props[inc()]),
		Beta:      dh.StringToUint8(props[inc()]) == 1,
		Expansion: dh.StringToUint8(props[inc()]) == 1,
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
