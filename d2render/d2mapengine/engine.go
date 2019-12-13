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
	GetPosition() (float64, float64)
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

func (me *MapEngine) GetStartPosition() (float64, float64) {
	var startX, startY float64
	if len(me.regions) > 0 {
		region := me.regions[0]
		startX, startY = region.getStartTilePosition()
	}

	return startX, startY
}

func (me *MapEngine) GetCenterPosition() (float64, float64) {
	var centerX, centerY float64
	if len(me.regions) > 0 {
		region := me.regions[0]
		centerX = float64(region.tileRect.Left) + float64(region.tileRect.Width)/2
		centerY = float64(region.tileRect.Top) + float64(region.tileRect.Height)/2
	}

	return centerX, centerY
}

func (me *MapEngine) MoveCameraTo(x, y float64) {
	me.camera.MoveTo(x, y)
}

func (me *MapEngine) MoveCameraBy(x, y float64) {
	me.camera.MoveBy(x, y)
}

func (me *MapEngine) ScreenToWorld(x, y int) (float64, float64) {
	return me.viewport.ScreenToWorld(x, y)
}

func (me *MapEngine) ScreenToOrtho(x, y int) (float64, float64) {
	return me.viewport.ScreenToOrtho(x, y)
}

func (me *MapEngine) WorldToOrtho(x, y float64) (float64, float64) {
	return me.viewport.WorldToOrtho(x, y)
}

func (me *MapEngine) SetDebugVisLevel(debugVisLevel int) {
	me.debugVisLevel = debugVisLevel
}

func (me *MapEngine) GenerateMap(regionType d2enum.RegionIdType, levelPreset int, fileIndex int) {
	region, entities := loadRegion(me.gameState.Seed, 0, 0, regionType, levelPreset, me.fileProvider, fileIndex)
	me.regions = append(me.regions, region)
	me.entities = append(me.entities, entities...)
}

func (me *MapEngine) GenerateAct1Overworld() {
	me.soundManager.PlayBGM("/data/global/music/Act1/town1.wav") // TODO: Temp stuff here

	region, entities := loadRegion(me.gameState.Seed, 0, 0, d2enum.RegionAct1Town, 1, me.fileProvider, -1)
	me.regions = append(me.regions, region)
	me.entities = append(me.entities, entities...)

	if strings.Contains(region.regionPath, "E1") {
		region, entities := loadRegion(me.gameState.Seed, int(region.tileRect.Width-1), 0, d2enum.RegionAct1Town, 2, me.fileProvider, -1)
		me.regions = append(me.regions, region)
		me.entities = append(me.entities, entities...)
	} else if strings.Contains(region.regionPath, "S1") {
		region, entities := loadRegion(me.gameState.Seed, 0, int(region.tileRect.Height-1), d2enum.RegionAct1Town, 3, me.fileProvider, -1)
		me.regions = append(me.regions, region)
		me.entities = append(me.entities, entities...)
	}
}

func (me *MapEngine) GetRegionAtTile(x, y int) *MapRegion {
	for _, region := range me.regions {
		if region.tileRect.IsInRect(x, y) {
			return region
		}
	}

	return nil
}

func (me *MapEngine) AddEntity(entity MapEntity) {
	me.entities = append(me.entities, entity)
}

func (me *MapEngine) Advance(tickTime float64) {
	for _, region := range me.regions {
		if region.isVisbile(me.viewport) {
			region.advance(tickTime)
		}
	}

	for _, entity := range me.entities {
		entity.Advance(tickTime)
	}
}

func (me *MapEngine) Render(target *ebiten.Image) {
	for _, region := range me.regions {
		if region.isVisbile(me.viewport) {
			region.renderPass1(me.viewport, target)
			me.renderDebug(target)
			region.renderPass2(me.entities, me.viewport, target)
			region.renderPass3(me.viewport, target)
		}
	}
}

func (me *MapEngine) renderDebug(target *ebiten.Image) {
	for _, region := range me.regions {
		if region.isVisbile(me.viewport) {
			region.renderDebug(me.debugVisLevel, me.viewport, target)
		}
	}
}
