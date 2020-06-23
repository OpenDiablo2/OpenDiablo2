package d2mapgen

import (
	"log"
	"math/rand"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapstamp"
)

func GenerateAct1Overworld(mapEngine *d2mapengine.MapEngine) {
	wildernessGrass := d2ds1.FloorShadowRecord{Prop1: 195, Style: 0}

	rand.Seed(mapEngine.Seed())

	levelDetails := d2datadict.GetLevelDetails(1)
	if levelDetails == nil {
		panic("Could not find level info for act 1!")
	}

	townStamp := d2mapstamp.LoadStamp(d2enum.RegionAct1Town, 1, -1)
	townStamp.RegionPath()
	townSize := townStamp.Size()
	mapEngine.ResetMap(d2enum.RegionAct1Town, levelDetails.SizeXNormal+townSize.Width, levelDetails.SizeYNormal+townSize.Height)
	mapWidth := mapEngine.Size().Width
	mapHeight := mapEngine.Size().Height
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			(*mapEngine.Tiles())[x+(y*mapWidth)].RegionType = d2enum.RegionIdType(levelDetails.LevelType)
			(*mapEngine.Tiles())[x+(y*mapWidth)].Floors = []d2ds1.FloorShadowRecord{wildernessGrass}
		}
	}

	log.Printf("Region Path: %s", townStamp.RegionPath())
	if strings.Contains(townStamp.RegionPath(), "E1") {
		// East Exit
		mapEngine.PlaceStamp(townStamp, 0, 0)

	} else if strings.Contains(townStamp.RegionPath(), "S1") {
		// South Exit
		mapEngine.PlaceStamp(townStamp, mapWidth-townSize.Width, 0)

	} else if strings.Contains(townStamp.RegionPath(), "W1") {
		// West Exit
		mapEngine.PlaceStamp(townStamp, mapWidth-townSize.Width, mapHeight-townSize.Height)

	} else {
		// North Exit
		mapEngine.PlaceStamp(townStamp, mapWidth-townSize.Width, mapHeight-townSize.Height)

	}

	mapEngine.RegenerateWalkPaths()
}
