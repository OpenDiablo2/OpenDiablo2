package d2map

import (
	"github.com/beefsack/go-astar"
	"math"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
)

type MapEntity interface {
	Render(target d2render.Surface)
	Advance(tickTime float64)
	GetPosition() (float64, float64)
}

type MapEngine struct {
	gameState *d2gamestate.GameState

	debugVisLevel int

	regions  []*MapRegion
	entities []MapEntity
	viewport *Viewport
	camera   Camera
}

func CreateMapEngine(gameState *d2gamestate.GameState) *MapEngine {
	engine := &MapEngine{
		gameState: gameState,
		viewport:  NewViewport(0, 0, 800, 600),
	}

	d2term.BindAction("mapdebugvis", "set map debug visualization level", func(level int) {
		engine.debugVisLevel = level
	})

	engine.viewport.SetCamera(&engine.camera)
	return engine
}

func (m *MapEngine) GetStartPosition() (float64, float64) {
	var startX, startY float64
	if len(m.regions) > 0 {
		region := m.regions[0]
		startX, startY = region.getStartTilePosition()
	}

	return startX, startY
}

func (m *MapEngine) GetCenterPosition() (float64, float64) {
	var centerX, centerY float64
	if len(m.regions) > 0 {
		region := m.regions[0]
		centerX = float64(region.tileRect.Left) + float64(region.tileRect.Width)/2
		centerY = float64(region.tileRect.Top) + float64(region.tileRect.Height)/2
	}

	return centerX, centerY
}

func (m *MapEngine) MoveCameraTo(x, y float64) {
	m.camera.MoveTo(x, y)
}

func (m *MapEngine) MoveCameraBy(x, y float64) {
	m.camera.MoveBy(x, y)
}

func (m *MapEngine) ScreenToWorld(x, y int) (float64, float64) {
	return m.viewport.ScreenToWorld(x, y)
}

func (m *MapEngine) ScreenToOrtho(x, y int) (float64, float64) {
	return m.viewport.ScreenToOrtho(x, y)
}

func (m *MapEngine) WorldToOrtho(x, y float64) (float64, float64) {
	return m.viewport.WorldToOrtho(x, y)
}

func (m *MapEngine) GenerateMap(regionType d2enum.RegionIdType, levelPreset int, fileIndex int) {
	region, entities := loadRegion(m.gameState.Seed, 0, 0, regionType, levelPreset, fileIndex)
	m.regions = append(m.regions, region)
	m.entities = append(m.entities, entities...)
}

func (m *MapEngine) GenerateAct1Overworld() {
	d2audio.PlayBGM("/data/global/music/Act1/town1.wav") // TODO: Temp stuff here

	region, entities := loadRegion(m.gameState.Seed, 0, 0, d2enum.RegionAct1Town, 1, -1)
	m.regions = append(m.regions, region)
	m.entities = append(m.entities, entities...)

	if strings.Contains(region.regionPath, "E1") {
		region, entities := loadRegion(m.gameState.Seed, region.tileRect.Width-1, 0, d2enum.RegionAct1Town, 2, -1)
		m.AppendRegion(region)
		m.entities = append(m.entities, entities...)
	} else if strings.Contains(region.regionPath, "S1") {
		region, entities := loadRegion(m.gameState.Seed, 0, region.tileRect.Height-1, d2enum.RegionAct1Town, 3, -1)
		m.AppendRegion(region)
		m.entities = append(m.entities, entities...)
	}
}

func (m *MapEngine) AppendRegion(region *MapRegion) {
	// TODO: Stitch together region.walkableArea
	m.regions = append(m.regions, region)
}

func (m *MapEngine) GetRegionAtTile(x, y int) *MapRegion {
	for _, region := range m.regions {
		if region.tileRect.IsInRect(x, y) {
			return region
		}
	}

	return nil
}

func (m *MapEngine) AddEntity(entity MapEntity) {
	m.entities = append(m.entities, entity)
}

func (m *MapEngine) Advance(tickTime float64) {
	for _, region := range m.regions {
		if region.isVisbile(m.viewport) {
			region.advance(tickTime)
		}
	}

	for _, entity := range m.entities {
		entity.Advance(tickTime)
	}
}

func (m *MapEngine) Render(target d2render.Surface) {
	for _, region := range m.regions {
		if region.isVisbile(m.viewport) {
			region.renderPass1(m.viewport, target)
			region.renderDebug(m.debugVisLevel, m.viewport, target)
			region.renderPass2(m.entities, m.viewport, target)
			region.renderPass3(m.viewport, target)
		}
	}
}

func (m *MapEngine) PathFind(startX, startY, endX, endY float64) (path []astar.Pather, distance float64, found bool){
	startTileX := int(math.Floor(startX))
	startTileY := int(math.Floor(startY))
	startSubtileX := int((startX - float64(int(startX))) * 5)
	startSubtileY := int((startY - float64(int(startY))) * 5)
	startRegion := m.GetRegionAtTile(startTileX, startTileY)
	startNode := &startRegion.walkableArea[startSubtileY + ((startTileY - startRegion.tileRect.Top) * 5)][startSubtileX + ((startTileX - startRegion.tileRect.Left) * 5)]

	endTileX := int(math.Floor(endX))
	endTileY := int(math.Floor(endY))
	endSubtileX := int((endX - float64(int(endX))) * 5)
	endSubtileY := int((endY - float64(int(endY))) * 5)
	endRegion := m.GetRegionAtTile(endTileX, endTileY)
	endNode := &endRegion.walkableArea[endSubtileY + ((endTileY - endRegion.tileRect.Top) * 5)][endSubtileX + ((endTileX - endRegion.tileRect.Left) * 5)]

	path, distance, found = astar.Path(endNode, startNode)
	if path != nil {
		path = path[1:]
	}
	return
}
