package d2mapgen

import (
	"log"
	"math/rand"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen/d2wilderness"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapstamp"
)

var wildernessGrass = d2ds1.FloorShadowRecord{Prop1: 1, Style: 0, Sequence: 0}

type TownDirection int

const (
	TownDirectionNorth TownDirection = iota
	TownDirectionSouth
	TownDirectionEast
	TownDirectionWest
)

func loadPreset(mapEngine *d2mapengine.MapEngine, id, index int) *d2mapstamp.Stamp {
	for _, file := range d2datadict.LevelPresets[id].Files {
		mapEngine.AddDS1(file)
	}
	return d2mapstamp.LoadStamp(d2enum.RegionAct1Wilderness, id, index)
}

func GenerateAct1Overworld(mapEngine *d2mapengine.MapEngine) {

	rand.Seed(mapEngine.Seed())
	wilderness1Details := d2datadict.GetLevelDetails(2)
	mapEngine.ResetMap(d2enum.RegionAct1Town, 150, 150)
	mapWidth := mapEngine.Size().Width
	mapHeight := mapEngine.Size().Height

	townStamp := d2mapstamp.LoadStamp(d2enum.RegionAct1Town, 1, -1)
	townStamp.RegionPath()
	townSize := townStamp.Size()
	var townDirection TownDirection

	log.Printf("Region Path: %s", townStamp.RegionPath())
	if strings.Contains(townStamp.RegionPath(), "E1") {
		// East Exit
		townDirection = TownDirectionEast
		mapEngine.PlaceStamp(townStamp, 0, 0)

	} else if strings.Contains(townStamp.RegionPath(), "S1") {
		// South Exit
		townDirection = TownDirectionSouth
		mapEngine.PlaceStamp(townStamp, mapWidth-townSize.Width, 0)
		rightWaterBorderStamp := d2mapstamp.LoadStamp(d2enum.RegionAct1Wilderness, d2wilderness.WaterBorderEast, 0)
		rightWaterBorderStamp2 := d2mapstamp.LoadStamp(d2enum.RegionAct1Wilderness, d2wilderness.WaterBorderWest, 0)
		// Place the water on the right side of the map
		for y := townSize.Height; y < mapHeight-9; y += 9 {
			mapEngine.PlaceStamp(rightWaterBorderStamp, mapWidth-17, y)
			mapEngine.PlaceStamp(rightWaterBorderStamp2, mapWidth-9, y)
		}
		generateWilderness1(mapEngine, mapWidth-wilderness1Details.SizeXNormal-14, townSize.Height, townDirection)

	} else if strings.Contains(townStamp.RegionPath(), "W1") {
		// West Exit
		townDirection = TownDirectionWest
		mapEngine.PlaceStamp(townStamp, mapWidth-townSize.Width, mapHeight-townSize.Height)

	} else {
		// North Exit
		townDirection = TownDirectionNorth
		mapEngine.PlaceStamp(townStamp, mapWidth-townSize.Width, mapHeight-townSize.Height)

	}

	mapEngine.RegenerateWalkPaths()
}

func generateWilderness1(mapEngine *d2mapengine.MapEngine, startX, startY int, townDirection TownDirection) {
	levelDetails := d2datadict.GetLevelDetails(2)
	mapWidth := mapEngine.Size().Width

	fenceNorthStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderNorth, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderNorth, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderNorth, 2),
	}

	fenceWestStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderWest, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderWest, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderWest, 2),
	}

	fenceSouthStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderSouth, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderSouth, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderSouth, 2),
	}

	fenceNorthWestStamp := loadPreset(mapEngine, d2wilderness.TreeBorderNorthWest, 0)
	fenceSouthWestStamp := loadPreset(mapEngine, d2wilderness.TreeBorderSouthWest, 0)
	fenceWaterBorderSouthEast := loadPreset(mapEngine, d2wilderness.WaterBorderEast, 1)

	// Fill in the grass
	for y := 0; y < levelDetails.SizeYNormal; y++ {
		for x := 0; x < levelDetails.SizeXNormal; x++ {
			(*mapEngine.Tiles())[startX+x+((startY+y)*mapWidth)].RegionType = d2enum.RegionIdType(levelDetails.LevelType)
			(*mapEngine.Tiles())[startX+x+((startY+y)*mapWidth)].Floors = []d2ds1.FloorShadowRecord{wildernessGrass}
		}
	}

	// Draw the north fence
	if townDirection == TownDirectionSouth {
		for i := 0; i < 4; i++ {
			mapEngine.PlaceStamp(fenceNorthStamp[rand.Intn(3)], startX+(i*9)+5, startY-6)
		}
	}

	// Draw the west fence
	if townDirection == TownDirectionSouth {
		for i := 0; i < 8; i++ {
			mapEngine.PlaceStamp(fenceWestStamp[rand.Intn(3)], startX, startY+(i*9)+3)
		}
	}

	// Draw the south fence
	if townDirection == TownDirectionSouth {
		for i := 1; i < 9; i++ {
			mapEngine.PlaceStamp(fenceSouthStamp[rand.Intn(3)], startX+(i*9), startY+(8*9)+3)
		}
	}

	mapEngine.PlaceStamp(fenceNorthWestStamp, startX, startY-6)
	mapEngine.PlaceStamp(fenceSouthWestStamp, startX, startY+(8*9)+3)
	mapEngine.PlaceStamp(fenceWaterBorderSouthEast, startX+(9*9)-4, startY+(8*9)+1)
}
