package d2mapengine

import (
	"strings"

	"github.com/hajimehoshi/ebiten"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
)

type MapEntity interface {
	Render(target *ebiten.Image, screenX, screenY int)
	GetTilePosition() (float64, float64)
	Advance(tickTime float64)
}

type MapEngine struct {
	soundManager *d2audio.Manager
	gameState    *d2core.GameState
	fileProvider d2interface.FileProvider

	debugVisLevel int

	regions  []*MapRegion
	entities []MapEntity
	viewport *Viewport
	camera   Camera
}

func CreateMapEngine(gameState *d2core.GameState, soundManager *d2audio.Manager, fileProvider d2interface.FileProvider) *MapEngine {
	engine := &MapEngine{
		gameState:    gameState,
		soundManager: soundManager,
		fileProvider: fileProvider,
		viewport:     NewViewport(0, 0, 800, 600),
	}

	engine.viewport.SetCamera(&engine.camera)
	return engine
}

func (e *MapEngine) GetStartTilePosition() (float64, float64) {
	var startX, startY float64
	if len(e.regions) > 0 {
		region := e.regions[0]
		startX, startY = region.getStartTilePosition()
	}

	return startX, startY
}

func (e *MapEngine) MoveCameraTo(x, y float64) {
	e.camera.MoveTo(x, y)
}

func (e *MapEngine) MoveCameraBy(x, y float64) {
	e.camera.MoveBy(x, y)
}

func (e *MapEngine) ScreenToIso(x, y int) (float64, float64) {
	return e.viewport.ScreenToIso(x, y)
}

func (e *MapEngine) ScreenToWorld(x, y int) (float64, float64) {
	return e.viewport.ScreenToWorld(x, y)
}

func (e *MapEngine) IsoToWorld(x, y float64) (float64, float64) {
	return e.viewport.IsoToWorld(x, y)
}

func (e *MapEngine) SetDebugVisLevel(debugVisLevel int) {
	e.debugVisLevel = debugVisLevel
}

func (e *MapEngine) GenerateMap(regionType d2enum.RegionIdType, levelPreset int, fileIndex int) {
	region, entities := loadRegion(e.gameState.Seed, 0, 0, regionType, levelPreset, e.fileProvider, fileIndex)
	e.regions = append(e.regions, region)
	e.entities = append(e.entities, entities...)
}

func (e *MapEngine) GenerateAct1Overworld() {
	e.soundManager.PlayBGM("/data/global/music/Act1/town1.wav") // TODO: Temp stuff here

	region, entities := loadRegion(e.gameState.Seed, 0, 0, d2enum.RegionAct1Town, 1, e.fileProvider, -1)
	e.regions = append(e.regions, region)
	e.entities = append(e.entities, entities...)

	if strings.Contains(region.regionPath, "E1") {
		region, entities := loadRegion(e.gameState.Seed, int(region.tileRect.Width-1), 0, d2enum.RegionAct1Town, 2, e.fileProvider, -1)
		e.regions = append(e.regions, region)
		e.entities = append(e.entities, entities...)
	} else if strings.Contains(region.regionPath, "S1") {
		region, entities := loadRegion(e.gameState.Seed, 0, int(region.tileRect.Height-1), d2enum.RegionAct1Town, 3, e.fileProvider, -1)
		e.regions = append(e.regions, region)
		e.entities = append(e.entities, entities...)
	}
}

func (e *MapEngine) GetRegionAt(x, y int) *MapRegion {
	for _, region := range e.regions {
		if region.tileRect.IsInRect(x, y) {
			return region
		}
	}

	return nil
}

func (e *MapEngine) AddEntity(entity MapEntity) {
	e.entities = append(e.entities, entity)
}

func (e *MapEngine) Advance(tickTime float64) {
	for _, region := range e.regions {
		if region.isVisbile(e.viewport) {
			region.advance(tickTime)
		}
	}

	for _, entity := range e.entities {
		entity.Advance(tickTime)
	}
}

func (e *MapEngine) Render(target *ebiten.Image) {
	for _, region := range e.regions {
		region.renderPass1(e.viewport, target)
		e.renderDebug(target)
		region.renderPass2(e.viewport, target)
		e.renderEntities(target)
		region.renderPass3(e.viewport, target)
	}
}

func (e *MapEngine) renderEntities(target *ebiten.Image) {
	for _, entity := range e.entities {
		e.viewport.PushTranslation(e.viewport.IsoToWorld(entity.GetTilePosition()))
		screenX, screenY := e.viewport.WorldToScreen(e.viewport.GetTranslation())
		entity.Render(target, screenX, screenY)
		e.viewport.PopTranslation()
	}
}

func (e *MapEngine) renderDebug(target *ebiten.Image) {
	for _, region := range e.regions {
		if region.isVisbile(e.viewport) {
			region.renderDebug(e.debugVisLevel, e.viewport, target)
		}
	}
}
