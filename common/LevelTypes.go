package common

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/resourcepaths"
)

type LevelTypeRecord struct {
	Name      string
	Id        int
	Files     [32]string
	Beta      bool
	Act       int
	Expansion bool
}

var LevelTypes []LevelTypeRecord

func LoadLevelTypes(fileProvider FileProvider) {
	data := strings.Split(string(fileProvider.LoadFile(resourcepaths.LevelType)), "\r\n")[1:]
	LevelTypes = make([]LevelTypeRecord, len(data))
	for i, line := range data {
		idx := -1
		inc := func() int {
			idx++
			return idx
		}
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "\t")
		if parts[0] == "Expansion" {
			continue
		}
		LevelTypes[i].Name = parts[inc()]
		LevelTypes[i].Id = StringToInt(parts[inc()])
		for fileIdx := range LevelTypes[i].Files {
			LevelTypes[i].Files[fileIdx] = parts[inc()]
			if LevelTypes[i].Files[fileIdx] == "0" {
				LevelTypes[i].Files[fileIdx] = ""
			}

		}
		LevelTypes[i].Beta = parts[inc()] != "1"
		LevelTypes[i].Act = StringToInt(parts[inc()])
		LevelTypes[i].Expansion = parts[inc()] != "1"
	}
	log.Printf("Loaded %d LevelType records", len(LevelTypes))
}
