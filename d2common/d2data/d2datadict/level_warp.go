package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// LevelWarpRecord is a representation of a row from lvlwarp.txt
// it describes the warp graphics offsets and dimensions for levels
type LevelWarpRecord struct {
	ID         int32
	SelectX    int32
	SelectY    int32
	SelectDX   int32
	SelectDY   int32
	ExitWalkX  int32
	ExitWalkY  int32
	OffsetX    int32
	OffsetY    int32
	LitVersion bool
	Tiles      int32
	Direction  string
}

// LevelWarps loaded from txt records
//nolint:gochecknoglobals // Currently global by design, only written once
var LevelWarps map[int]*LevelWarpRecord

// LoadLevelWarps loads LevelWarpRecord's from text file data
func LoadLevelWarps(levelWarpData []byte) {
	LevelWarps = make(map[int]*LevelWarpRecord)
	streamReader := d2common.CreateStreamReader(levelWarpData)
	numRecords := int(streamReader.GetInt32())

	for i := 0; i < numRecords; i++ {
		id := int(streamReader.GetInt32())
		LevelWarps[id] = &LevelWarpRecord{}
		LevelWarps[id].ID = int32(id)
		LevelWarps[id].SelectX = streamReader.GetInt32()
		LevelWarps[id].SelectY = streamReader.GetInt32()
		LevelWarps[id].SelectDX = streamReader.GetInt32()
		LevelWarps[id].SelectDY = streamReader.GetInt32()
		LevelWarps[id].ExitWalkX = streamReader.GetInt32()
		LevelWarps[id].ExitWalkY = streamReader.GetInt32()
		LevelWarps[id].OffsetX = streamReader.GetInt32()
		LevelWarps[id].OffsetY = streamReader.GetInt32()
		LevelWarps[id].LitVersion = streamReader.GetInt32() == 1
		LevelWarps[id].Tiles = streamReader.GetInt32()
		LevelWarps[id].Direction = string(streamReader.GetByte())
		streamReader.SkipBytes(3)
	}

	log.Printf("Loaded %d level warps", len(LevelWarps))
}
