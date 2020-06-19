package d2map

import (
	"log"
	"math"

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
func CreateMapEngine(seed int64) *MapEngine {
	engine := &MapEngine{
		seed:     seed,
		entities: NewRangeSearcher(),
	}

	return engine
}

// Sets the seed of the map for generation
func (m *MapEngine) SetSeed(seed int64) {
	log.Printf("Setting map engine seed to %d", seed)
	m.seed = seed
}

func (m *MapEngine) GetStartPosition() (float64, float64) {
	var startX, startY float64

	// TODO: Temporary code, only works for starting map
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

// Appends a region to the map
func (m *MapEngine) AppendRegion(region *MapRegion) {
	m.regions = append(m.regions, region)
	// Stitch together the walk map
	for x := 0; x < region.tileRect.Width*5; x++ {
		otherRegion := m.GetRegionAtTile(region.tileRect.Left+(x/5), region.tileRect.Top-1)
		if otherRegion == nil {
			continue
		}
		xDiff := region.tileRect.Left - otherRegion.tileRect.Left

		sourceSubtile := &region.walkableArea[0][x]
		if !sourceSubtile.Walkable {
			continue
		}

		// North West
		otherX := x + xDiff - 1
		otherY := (otherRegion.tileRect.Height * 5) - 1
		if otherX < 0 || otherX >= len(otherRegion.walkableArea[otherY]) {
			continue
		}
		otherRegion.walkableArea[otherY][x+xDiff].DownRight = sourceSubtile
		sourceSubtile.UpLeft = &otherRegion.walkableArea[otherY][x+xDiff]

		// North
		otherX++
		if otherX < 0 || otherX >= len(otherRegion.walkableArea[otherY]) {
			continue
		}
		otherRegion.walkableArea[otherY][x+xDiff].Down = sourceSubtile
		sourceSubtile.Up = &otherRegion.walkableArea[otherY][x+xDiff]

		// NorthEast
		otherX++
		if otherX < 0 || otherX >= len(otherRegion.walkableArea[otherY]) {
			continue
		}
		otherRegion.walkableArea[otherY][x+xDiff].DownLeft = sourceSubtile
		sourceSubtile.UpRight = &otherRegion.walkableArea[otherY][x+xDiff]
	}

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

// Adds an entity to the map engine
func (m *MapEngine) AddEntity(entity MapEntity) {
	m.entities.Add(entity)
}

// Removes an entity from the map engine
func (m *MapEngine) RemoveEntity(entity MapEntity) {
	if entity == nil {
		return
	}

	m.entities.Remove(entity)
}

// Advances time on the map engine
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

// Finds a walkable path between two points
func (m *MapEngine) PathFind(startX, startY, endX, endY float64) (path []astar.Pather, distance float64, found bool) {
	startTileX := int(math.Floor(startX))
	startTileY := int(math.Floor(startY))
	startSubtileX := int((startX - float64(int(startX))) * 5)
	startSubtileY := int((startY - float64(int(startY))) * 5)
	startRegion := m.GetRegionAtTile(startTileX, startTileY)
	if startRegion == nil {
		return
	}
	startNode := &startRegion.walkableArea[startSubtileY+((startTileY-startRegion.tileRect.Top)*5)][startSubtileX+((startTileX-startRegion.tileRect.Left)*5)]

	endTileX := int(math.Floor(endX))
	endTileY := int(math.Floor(endY))
	endSubtileX := int((endX - float64(int(endX))) * 5)
	endSubtileY := int((endY - float64(int(endY))) * 5)
	endRegion := m.GetRegionAtTile(endTileX, endTileY)
	if endRegion == nil {
		return
	}
	endNodeY := endSubtileY + ((endTileY - endRegion.tileRect.Top) * 5)
	endNodeX := endSubtileX + ((endTileX - endRegion.tileRect.Left) * 5)
	if endNodeY < 0 || endNodeY >= len(endRegion.walkableArea) {
		return
	}
	if endNodeX < 0 || endNodeX >= len(endRegion.walkableArea[endNodeY]) {
		return
	}
	endNode := &endRegion.walkableArea[endNodeY][endNodeX]

	path, distance, found = astar.Path(endNode, startNode)
	if path != nil {
		path = path[1:]
	}
	return
}
