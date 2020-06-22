package d2mapgen

import (
	"log"
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapstamp"
)

func GenerateAct1Overworld(mapEngine *d2mapengine.MapEngine) {
	log.Printf("Map seed: %d", mapEngine.Seed())
	rand.Seed(mapEngine.Seed())
	townStamp := d2mapstamp.LoadStamp(d2enum.RegionAct1Town, 1, -1)
	townSize := townStamp.Size()
	mapEngine.ResetMap(d2enum.RegionAct1Town, 100, 100) // TODO: Mapgen - Needs levels.txt stuff

	wildernessGrass := d2ds1.FloorShadowRecord{
		Prop1: 195,
		Style: 0,
	}

	mapWidth := mapEngine.Size().Width
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			(*mapEngine.Tiles())[x+(y*mapWidth)].RegionType = d2enum.RegionAct1Wilderness
			(*mapEngine.Tiles())[x+(y*mapWidth)].Floors = []d2ds1.FloorShadowRecord{wildernessGrass}
		}
	}

	mapEngine.PlaceStamp(townStamp, 50-(townSize.Width/2), 50-(townSize.Height/2))

	mapEngine.RegenerateWalkPaths()
}
