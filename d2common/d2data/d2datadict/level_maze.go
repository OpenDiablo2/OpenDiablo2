package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type LevelMazeDetailsRecord struct {
	// descriptive, not loaded in game. Corresponds with Name field in
	// Levels.txt
	Name string // Name

	// ID from Levels.txt
	// NOTE: Cave 1 is the Den of Evil, its associated treasure level is quest
	// only.
	LevelId int // Level

	// the minimum number of .ds1 map sections that will make up the maze in
	// Normal, Nightmare and Hell difficulties.
	NumRoomsNormal    int // Rooms
	NumRoomsNightmare int // Rooms(N)
	NumRoomsHell      int // Rooms(H)

	// the size in the X\Y direction of any component ds1 map section.
	SizeX int // SizeX
	SizeY int // SizeY

	// Possibly related to how adjacent .ds1s are connected with each other,
	// but what the different values are for is unknown.
	// Merge int // Merge

	// Included in the original Diablo II beta tests and in the demo version.
	// Beta
}

var LevelMazeDetails map[int]*LevelMazeDetailsRecord

func LoadLevelMazeDetails(file []byte) {
	dict := d2common.LoadDataDictionary(string(file))
	numRecords := len(dict.Data)
	LevelMazeDetails = make(map[int]*LevelMazeDetailsRecord, numRecords)
	for idx, _ := range dict.Data {
		record := &LevelMazeDetailsRecord{
			Name:              dict.GetString("Name", idx),
			LevelId:           dict.GetNumber("Level", idx),
			NumRoomsNormal:    dict.GetNumber("Rooms", idx),
			NumRoomsNightmare: dict.GetNumber("Rooms(N)", idx),
			NumRoomsHell:      dict.GetNumber("Rooms(H)", idx),
			SizeX:             dict.GetNumber("SizeX", idx),
			SizeY:             dict.GetNumber("SizeY", idx),
		}
		LevelMazeDetails[record.LevelId] = record
	}
	log.Printf("Loaded %d LevelMazeDetails records", len(LevelMazeDetails))
}
