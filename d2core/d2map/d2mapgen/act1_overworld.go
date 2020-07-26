//nolint:gomnd
package d2mapgen

import (
	"log"
	"math/rand"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen/d2wilderness"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapstamp"
)

func loadPreset(mapEngine *d2mapengine.MapEngine, id, index int) *d2mapstamp.Stamp {
	for _, file := range d2datadict.LevelPreset(id).Files {
		mapEngine.AddDS1(file)
	}

	return d2mapstamp.LoadStamp(d2enum.RegionAct1Wilderness, id, index)
}

// GenerateAct1Overworld generates the map and entities for the first town and surrounding area.
func GenerateAct1Overworld(mapEngine *d2mapengine.MapEngine) {
	rand.Seed(mapEngine.Seed())

	wilderness1Details := d2datadict.GetLevelDetails(2)

	mapEngine.ResetMap(d2enum.RegionAct1Town, 150, 150)
	mapWidth := mapEngine.Size().Width
	mapHeight := mapEngine.Size().Height

	townStamp := d2mapstamp.LoadStamp(d2enum.RegionAct1Town, 1, -1)
	townStamp.RegionPath()
	townSize := townStamp.Size()

	log.Printf("Region Path: %s", townStamp.RegionPath())

	if strings.Contains(townStamp.RegionPath(), "E1") {
		// East Exit
		mapEngine.PlaceStamp(townStamp, 0, 0)
		generateWilderness1TownEast(mapEngine, townSize.Width, 0)
	} else if strings.Contains(townStamp.RegionPath(), "S1") {
		// South Exit
		mapEngine.PlaceStamp(townStamp, mapWidth-townSize.Width, 0)

		// Generate the river running along the edge of the map
		rightWaterBorderStamp := loadPreset(mapEngine, d2wilderness.WaterBorderEast, 0)
		rightWaterBorderStamp2 := loadPreset(mapEngine, d2wilderness.WaterBorderWest, 0)
		// Place the water on the right side of the map
		for y := townSize.Height; y < mapHeight-9; y += 9 {
			mapEngine.PlaceStamp(rightWaterBorderStamp, mapWidth-17, y)
			mapEngine.PlaceStamp(rightWaterBorderStamp2, mapWidth-9, y)
		}
		generateWilderness1TownSouth(mapEngine, mapWidth-wilderness1Details.SizeXNormal-14, townSize.Height)
	} else if strings.Contains(townStamp.RegionPath(), "W1") {
		// West Exit
		mapEngine.PlaceStamp(townStamp, mapWidth-townSize.Width, mapHeight-townSize.Height)

		generateWilderness1TownWest(mapEngine, mapWidth-townSize.Width-wilderness1Details.SizeXNormal, mapHeight-wilderness1Details.SizeYNormal)
	} else {
		// North Exit
		mapEngine.PlaceStamp(townStamp, mapWidth-townSize.Width, mapHeight-townSize.Height)
	}
}

func generateWilderness1TownEast(mapEngine *d2mapengine.MapEngine, startX, startY int) {
	levelDetails := d2datadict.GetLevelDetails(2)

	fenceNorthStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderNorth, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderNorth, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderNorth, 2),
	}

	fenceSouthStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderSouth, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderSouth, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderSouth, 2),
	}

	fenceWestStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderWest, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderWest, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderWest, 2),
	}

	fenceEastStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderEast, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderEast, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderEast, 2),
	}

	fenceSouthWestStamp := loadPreset(mapEngine, d2wilderness.TreeBorderSouthWest, 0)
	fenceNorthEastStamp := loadPreset(mapEngine, d2wilderness.TreeBorderNorthEast, 0)
	fenceSouthEastStamp := loadPreset(mapEngine, d2wilderness.TreeBorderSouthEast, 0)
	fenceWestEdge := loadPreset(mapEngine, d2wilderness.TreeBoxNorthEast, 0)

	areaRect := d2common.Rectangle{
		Left:   startX,
		Top:    startY + 9,
		Width:  levelDetails.SizeXNormal,
		Height: levelDetails.SizeYNormal - 3,
	}
	generateWilderness1Contents(mapEngine, areaRect)

	// Draw the north and south fence
	for i := 0; i < 9; i++ {
		mapEngine.PlaceStamp(fenceNorthStamp[rand.Intn(3)], startX+(i*9), startY)
		mapEngine.PlaceStamp(fenceSouthStamp[rand.Intn(3)], startX+(i*9), startY+(levelDetails.SizeYNormal+6))
	}

	// West fence
	for i := 1; i < 6; i++ {
		mapEngine.PlaceStamp(fenceWestStamp[rand.Intn(3)], startX, startY+(levelDetails.SizeYNormal+6)-(i*9))
	}

	// East Fence
	for i := 1; i < 10; i++ {
		mapEngine.PlaceStamp(fenceEastStamp[rand.Intn(3)], startX+levelDetails.SizeXNormal, startY+(i*9))
	}

	mapEngine.PlaceStamp(fenceSouthWestStamp, startX, startY+levelDetails.SizeYNormal+6)
	mapEngine.PlaceStamp(fenceWestEdge, startX, startY+(levelDetails.SizeYNormal-3)-45)
	mapEngine.PlaceStamp(fenceNorthEastStamp, startX+levelDetails.SizeXNormal, startY)
	mapEngine.PlaceStamp(fenceSouthEastStamp, startX+levelDetails.SizeXNormal, startY+levelDetails.SizeYNormal+6)
}

func generateWilderness1TownSouth(mapEngine *d2mapengine.MapEngine, startX, startY int) {
	levelDetails := d2datadict.GetLevelDetails(2)

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

	areaRect := d2common.Rectangle{
		Left:   startX + 2,
		Top:    startY,
		Width:  levelDetails.SizeXNormal - 2,
		Height: levelDetails.SizeYNormal - 3,
	}
	generateWilderness1Contents(mapEngine, areaRect)

	// Draw the north fence
	for i := 0; i < 4; i++ {
		mapEngine.PlaceStamp(fenceNorthStamp[rand.Intn(3)], startX+(i*9)+5, startY-6)
	}

	// Draw the west fence
	for i := 0; i < 8; i++ {
		mapEngine.PlaceStamp(fenceWestStamp[rand.Intn(3)], startX, startY+(i*9)+3)
	}

	// Draw the south fence
	for i := 1; i < 9; i++ {
		mapEngine.PlaceStamp(fenceSouthStamp[rand.Intn(3)], startX+(i*9), startY+(8*9)+3)
	}

	mapEngine.PlaceStamp(fenceNorthWestStamp, startX, startY-6)
	mapEngine.PlaceStamp(fenceSouthWestStamp, startX, startY+(8*9)+3)
	mapEngine.PlaceStamp(fenceWaterBorderSouthEast, startX+(9*9)-4, startY+(8*9)+1)
}

func generateWilderness1TownWest(mapEngine *d2mapengine.MapEngine, startX, startY int) {
	levelDetails := d2datadict.GetLevelDetails(2)

	fenceEastEdge := loadPreset(mapEngine, d2wilderness.TreeBoxSouthWest, 0)
	fenceNorthWestStamp := loadPreset(mapEngine, d2wilderness.TreeBorderNorthWest, 0)
	fenceNorthEastStamp := loadPreset(mapEngine, d2wilderness.TreeBorderNorthEast, 0)
	fenceSouthWestStamp := loadPreset(mapEngine, d2wilderness.TreeBorderSouthWest, 0)

	fenceSouthStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderSouth, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderSouth, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderSouth, 2),
	}

	fenceNorthStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderNorth, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderNorth, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderNorth, 2),
	}

	fenceEastStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderEast, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderEast, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderEast, 2),
	}

	fenceWestStamp := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.TreeBorderWest, 0),
		loadPreset(mapEngine, d2wilderness.TreeBorderWest, 1),
		loadPreset(mapEngine, d2wilderness.TreeBorderWest, 2),
	}

	// Draw the north and south fences
	for i := 0; i < 9; i++ {
		if i > 0 && i < 8 {
			mapEngine.PlaceStamp(fenceNorthStamp[rand.Intn(3)], startX+(i*9)-1, startY-15)
		}

		mapEngine.PlaceStamp(fenceSouthStamp[rand.Intn(3)], startX+(i*9)-1, startY+levelDetails.SizeYNormal-12)
	}

	// Draw the east fence
	for i := 0; i < 6; i++ {
		mapEngine.PlaceStamp(fenceEastStamp[rand.Intn(3)], startX+levelDetails.SizeXNormal-9, startY+(i*9)-6)
	}

	// Draw the west fence
	for i := 0; i < 9; i++ {
		mapEngine.PlaceStamp(fenceWestStamp[rand.Intn(3)], startX, startY+(i*9)-6)
	}

	// Draw the west fence
	mapEngine.PlaceStamp(fenceEastEdge, startX+levelDetails.SizeXNormal-9, startY+39)
	mapEngine.PlaceStamp(fenceNorthWestStamp, startX, startY-15)
	mapEngine.PlaceStamp(fenceSouthWestStamp, startX, startY+levelDetails.SizeYNormal-12)
	mapEngine.PlaceStamp(fenceNorthEastStamp, startX+levelDetails.SizeXNormal-9, startY-15)

	areaRect := d2common.Rectangle{
		Left:   startX + 9,
		Top:    startY - 10,
		Width:  levelDetails.SizeXNormal - 9,
		Height: levelDetails.SizeYNormal - 2,
	}
	generateWilderness1Contents(mapEngine, areaRect)
}

func generateWilderness1Contents(mapEngine *d2mapengine.MapEngine, rect d2common.Rectangle) {
	levelDetails := d2datadict.GetLevelDetails(2)

	denOfEvil := loadPreset(mapEngine, d2wilderness.DenOfEvilEntrance, 0)
	denOfEvilLoc := d2common.Point{
		X: rect.Left + (rect.Width / 2) + rand.Intn(10),
		Y: rect.Top + (rect.Height / 2) + rand.Intn(10),
	}

	// Fill in the grass
	for y := 0; y < rect.Height; y++ {
		for x := 0; x < rect.Width; x++ {
			tile := mapEngine.Tile(rect.Left+x, rect.Top+y)
			tile.RegionType = d2enum.RegionIdType(levelDetails.LevelType)
			tile.Components.Floors = []d2ds1.FloorShadowRecord{{Prop1: 1, Style: 0, Sequence: 0}} // wildernessGrass
			tile.PrepareTile(x, y, mapEngine)
		}
	}

	stuff := []*d2mapstamp.Stamp{
		loadPreset(mapEngine, d2wilderness.StoneFill1, 0),
		loadPreset(mapEngine, d2wilderness.StoneFill1, 1),
		loadPreset(mapEngine, d2wilderness.StoneFill1, 2),
		loadPreset(mapEngine, d2wilderness.StoneFill2, 0),
		loadPreset(mapEngine, d2wilderness.StoneFill2, 1),
		loadPreset(mapEngine, d2wilderness.StoneFill2, 2),
		loadPreset(mapEngine, d2wilderness.Cottages1, 0),
		loadPreset(mapEngine, d2wilderness.Cottages1, 1),
		loadPreset(mapEngine, d2wilderness.Cottages1, 2),
		loadPreset(mapEngine, d2wilderness.Cottages1, 3),
		loadPreset(mapEngine, d2wilderness.Cottages1, 4),
		loadPreset(mapEngine, d2wilderness.Cottages1, 5),
		loadPreset(mapEngine, d2wilderness.FallenCamp1, 0),
		loadPreset(mapEngine, d2wilderness.FallenCamp1, 1),
		loadPreset(mapEngine, d2wilderness.FallenCamp1, 2),
		loadPreset(mapEngine, d2wilderness.FallenCamp1, 3),
		loadPreset(mapEngine, d2wilderness.Pond, 0),
		loadPreset(mapEngine, d2wilderness.SwampFill1, 0),
		loadPreset(mapEngine, d2wilderness.SwampFill2, 0),
	}

	mapEngine.PlaceStamp(denOfEvil, denOfEvilLoc.X, denOfEvilLoc.Y)

	numPlaced := 0
	for numPlaced < 25 {
		stamp := stuff[rand.Intn(len(stuff))]

		stampRect := d2common.Rectangle{
			Left:   rect.Left + rand.Intn(rect.Width) - stamp.Size().Width,
			Top:    rect.Top + rand.Intn(rect.Height) - stamp.Size().Height,
			Width:  stamp.Size().Width,
			Height: stamp.Size().Height,
		}

		if areaEmpty(mapEngine, stampRect) {
			mapEngine.PlaceStamp(stamp, stampRect.Left, stampRect.Top)
			numPlaced++
		}
	}
}

func areaEmpty(mapEngine *d2mapengine.MapEngine, rect d2common.Rectangle) bool {
	mapHeight := mapEngine.Size().Height
	mapWidth := mapEngine.Size().Width

	if rect.Top < 0 || rect.Left < 0 || rect.Bottom() >= mapHeight || rect.Right() >= mapWidth {
		return false
	}

	for y := rect.Top; y <= rect.Bottom(); y++ {
		for x := rect.Left; x <= rect.Right(); x++ {
			if len(mapEngine.Tile(x, y).Components.Floors) == 0 {
				continue
			}

			floor := mapEngine.Tile(x, y).Components.Floors[0]

			if floor.Style != 0 || floor.Sequence != 0 || floor.Prop1 != 1 {
				return false
			}
		}
	}

	return true
}
