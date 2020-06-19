package d2map

import (
	"log"
	"math"
	"strings"

	"github.com/beefsack/go-astar"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type MapEntity interface {
	Render(target d2render.Surface)
	Advance(tickTime float64)
	GetPosition() (float64, float64)
}

// Represents the map data for a specific location
type MapEngine struct {
	seed     int64
	regions  []*MapRegion
	entities MapEntitiesSearcher
}

// Creates a new instance of the map engine
func CreateMapEngine() *MapEngine {
	engine := &MapEngine{
		seed:     0,
		entities: NewRangeSearcher(),
	}

	return engine
}

func (m *MapEngine) SetSeed(seed int64) {
	log.Printf("Setting map engine seed to %d", seed)
	m.seed = seed
}

func (m *MapEngine) GetStartPosition() (float64, float64) {
	var startX, startY float64
	if len(m.regions) > 0 {
		region := m.regions[0]
		startX, startY = region.getStartTilePosition()
	}

	return startX, startY
}

// Returns the center of the map
func (m *MapEngine) GetCenterPosition() (float64, float64) {
	var centerX, centerY float64
	if len(m.regions) > 0 {
		region := m.regions[0]
		centerX = float64(region.tileRect.Left) + float64(region.tileRect.Width)/2
		centerY = float64(region.tileRect.Top) + float64(region.tileRect.Height)/2
	}

	return centerX, centerY
}

func (m *MapEngine) GenerateMap(regionType d2enum.RegionIdType, levelPreset int, fileIndex int, cacheTiles bool) {
	region, entities := loadRegion(m.seed, 0, 0, regionType, levelPreset, fileIndex, cacheTiles)
	m.regions = append(m.regions, region)
	m.entities.Add(entities...)
}

func (m *MapEngine) GenerateAct1Overworld(cacheTiles bool) {
	//d2audio.PlayBGM("/data/global/music/Act1/town1.wav") // TODO: Temp stuff here

	region, entities := loadRegion(m.seed, 0, 0, d2enum.RegionAct1Town, 1, -1, cacheTiles)
	m.regions = append(m.regions, region)
	m.entities.Add(entities...)

	if strings.Contains(region.regionPath, "E1") {
		region, entities := loadRegion(m.seed, region.tileRect.Width-1, 0, d2enum.RegionAct1Town, 2, -1, cacheTiles)
		m.AppendRegion(region)
		m.entities.Add(entities...)
	} else if strings.Contains(region.regionPath, "S1") {
		region, entities := loadRegion(m.seed, 0, region.tileRect.Height-1, d2enum.RegionAct1Town, 3, -1, cacheTiles)
		m.AppendRegion(region)
		m.entities.Add(entities...)
	}
}

// Appends a region to the map
func (m *MapEngine) AppendRegion(region *MapRegion) {
	// TODO: Stitch together region.walkableArea
	log.Printf("Warning: Walkable areas are not currently implemented")
	m.regions = append(m.regions, region)
}

// Returns the region located at the specified tile location
func (m *MapEngine) GetRegionAtTile(x, y int) *MapRegion {
	for _, region := range m.regions {
		if region.tileRect.IsInRect(x, y) {
			return region
		}
	}

	return nil
}

func (m *MapEngine) AddEntity(entity MapEntity) {
	m.entities.Add(entity)
}

func (m *MapEngine) RemoveEntity(entity MapEntity) {
	if entity == nil {
		return
	}

	m.entities.Remove(entity)
}

func (m *MapEngine) Advance(tickTime float64) {
	for _, region := range m.regions {
		//if region.isVisbile(m.viewport) {
		region.advance(tickTime)
		//}
	}

	for _, entity := range m.entities.All() {
		entity.Advance(tickTime)
	}

	m.entities.Update()
}

func (m *MapEngine) PathFind(startX, startY, endX, endY float64) (path []astar.Pather, distance float64, found bool) {
	startTileX := int(math.Floor(startX))
	startTileY := int(math.Floor(startY))
	startSubtileX := int((startX - float64(int(startX))) * 5)
	startSubtileY := int((startY - float64(int(startY))) * 5)
	startRegion := m.GetRegionAtTile(startTileX, startTileY)
	startNode := &startRegion.walkableArea[startSubtileY+((startTileY-startRegion.tileRect.Top)*5)][startSubtileX+((startTileX-startRegion.tileRect.Left)*5)]

	endTileX := int(math.Floor(endX))
	endTileY := int(math.Floor(endY))
	endSubtileX := int((endX - float64(int(endX))) * 5)
	endSubtileY := int((endY - float64(int(endY))) * 5)
	endRegion := m.GetRegionAtTile(endTileX, endTileY)
	endNode := &endRegion.walkableArea[endSubtileY+((endTileY-endRegion.tileRect.Top)*5)][endSubtileX+((endTileX-endRegion.tileRect.Left)*5)]

	path, distance, found = astar.Path(endNode, startNode)
	if path != nil {
		path = path[1:]
	}
	return
}
