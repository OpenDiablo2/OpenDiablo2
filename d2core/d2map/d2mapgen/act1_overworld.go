package d2mapgen

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapstamp"
)

func GenerateAct1Overworld(mapEngine *d2mapengine.MapEngine) {
	rand.Seed(mapEngine.Seed())
	townStamp := d2mapstamp.LoadStamp(mapEngine.Seed(), d2enum.RegionAct1Town, 1, -1)
	townSize := townStamp.Size()
	mapEngine.ResetMap(0, d2enum.RegionAct1Town, townSize.Width, townSize.Height) // TODO: Mapgen - Needs levels.txt stuff
	mapEngine.PlaceStamp(townStamp, 0, 0)

	mapEngine.RegenerateWalkPaths()

	//region, entities := LoadStamp(m.seed, 0, 0, d2enum.RegionAct1Town, 1, -1, cacheTiles)
	//m.regions = append(m.regions, region)
	//m.entities.Add(entities...)
	//if strings.Contains(region.regionPath, "E1") {
	//	region, entities := LoadStamp(m.seed, region.tileRect.Width, 0, d2enum.RegionAct1Town, 2, -1, cacheTiles)
	//	m.AppendRegion(region)
	//	m.entities.Add(entities...)
	//} else if strings.Contains(region.regionPath, "S1") {
	//	region.tileRect.Height -= 1 // For some reason, this has a duplciate wall tile strip...
	//	mapWidthTiles := ((region.tileRect.Width - 18) / 9)
	//	yOffset := region.tileRect.Height
	//	waterXOffset := region.tileRect.Width - 17
	//	region, entities := LoadStamp(m.seed, 0, yOffset, d2enum.RegionAct1Town, 3, -1, cacheTiles)
	//	m.AppendRegion(region)
	//	m.entities.Add(entities...)
	//	yOffset += region.tileRect.Height
	//
	//	var choices = [...]int{
	//		d2wilderness.StoneFill1,
	//		d2wilderness.StoneFill2,
	//		d2wilderness.SwampFill1,
	//		d2wilderness.Cottages1,
	//		d2wilderness.Cottages2,
	//		d2wilderness.Cottages3,
	//		d2wilderness.CorralFill,
	//		d2wilderness.FallenCamp1,
	//		d2wilderness.FallenCamp2,
	//		d2wilderness.Pond,
	//	}
	//
	//	for i := 0; i < 6; i++ {
	//		// West Border
	//		region, entities = LoadStamp(m.seed, 0, yOffset, d2enum.RegionAct1Wilderness, d2wilderness.TreeBorderWest, 0, cacheTiles)
	//		m.AppendRegion(region)
	//		m.entities.Add(entities...)
	//
	//		// East Border
	//		region, entities = LoadStamp(m.seed, waterXOffset, yOffset, d2enum.RegionAct1Wilderness, d2wilderness.WaterBorderEast, 0, cacheTiles)
	//		m.AppendRegion(region)
	//		m.entities.Add(entities...)
	//
	//		// Grass
	//		for ix := 0; ix < mapWidthTiles; ix++ {
	//			region, entities = LoadStamp(m.seed, ((ix)*9)+7, yOffset, d2enum.RegionAct1Wilderness, choices[rand.Intn(len(choices))], 0, cacheTiles)
	//			m.AppendRegion(region)
	//			m.entities.Add(entities...)
	//		}
	//
	//		yOffset += 9
	//	}
	//
	//}
}
