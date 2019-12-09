package d2mapengine

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"

	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/D2Shared/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type RegionTile struct {
	tileX, tileY int // tile coordinates
	offX, offY   int // world space coordinates of tile origin
}

type EngineRegion struct {
	Rect   d2common.Rectangle
	Region *Region
	Tiles  []RegionTile
}

type Engine struct {
	soundManager *d2audio.Manager
	gameState    *d2core.GameState
	fileProvider d2interface.FileProvider
	region       int
	regions      []EngineRegion
	ShowTiles    int
	Hero         *d2core.Hero

	viewport *Viewport
	camera   Camera
}

func CreateMapEngine(gameState *d2core.GameState, soundManager *d2audio.Manager, fileProvider d2interface.FileProvider) *Engine {
	engine := &Engine{
		gameState:    gameState,
		soundManager: soundManager,
		fileProvider: fileProvider,
		regions:      make([]EngineRegion, 0),
		viewport:     NewViewport(0, 0, 800, 600),
	}

	engine.viewport.SetCamera(&engine.camera)
	return engine
}

func (v *Engine) Region() *EngineRegion {
	return &v.regions[v.region]
}

func (v *Engine) SetRegion(region int) {
	v.region = region
}

func (v *Engine) GetRegion(regionIndex int) *EngineRegion {
	return &v.regions[regionIndex]
}

func (v *Engine) CenterCameraOn(x, y float64) {
	v.camera.MoveTo(x, y)
}

func (v *Engine) MoveCameraBy(x, y float64) {
	v.camera.MoveBy(x, y)
}

func (v *Engine) ScreenToIso(x, y int) (float64, float64) {
	return v.viewport.ScreenToIso(x, y)
}

func (v *Engine) ScreenToWorld(x, y int) (float64, float64) {
	return v.viewport.ScreenToWorld(x, y)
}

func (v *Engine) GenerateMap(regionType d2enum.RegionIdType, levelPreset int, fileIndex int) {
	region := LoadRegion(v.gameState.Seed, regionType, levelPreset, v.fileProvider, fileIndex)
	fmt.Printf("Loading region: %v\n", region.RegionPath)
	v.regions = append(v.regions, EngineRegion{
		Rect:   d2common.Rectangle{0, 0, int(region.TileWidth), int(region.TileHeight)},
		Region: region,
	})

	for i, _ := range v.regions {
		v.GenTiles(&v.regions[i])
		v.GenTilesCache(&v.regions[i])
	}
}

func (v *Engine) GenerateAct1Overworld() {
	v.soundManager.PlayBGM("/data/global/music/Act1/town1.wav") // TODO: Temp stuff here
	region := LoadRegion(v.gameState.Seed, d2enum.RegionAct1Town, 1, v.fileProvider, -1)
	v.regions = append(v.regions, EngineRegion{
		Rect:   d2common.Rectangle{0, 0, int(region.TileWidth), int(region.TileHeight)},
		Region: region,
	})
	if strings.Contains(region.RegionPath, "E1") {
		region2 := LoadRegion(v.gameState.Seed, d2enum.RegionAct1Town, 2, v.fileProvider, -1)
		v.regions = append(v.regions, EngineRegion{
			Rect:   d2common.Rectangle{int(region.TileWidth - 1), 0, int(region2.TileWidth), int(region2.TileHeight)},
			Region: region2,
		})
	} else if strings.Contains(region.RegionPath, "S1") {
		region2 := LoadRegion(v.gameState.Seed, d2enum.RegionAct1Town, 3, v.fileProvider, -1)
		v.regions = append(v.regions, EngineRegion{
			Rect:   d2common.Rectangle{0, int(region.TileHeight - 1), int(region2.TileWidth), int(region2.TileHeight)},
			Region: region2,
		})
	}

	for i, _ := range v.regions {
		v.GenTiles(&v.regions[i])
		v.GenTilesCache(&v.regions[i])
	}

	v.camera.MoveTo(v.viewport.IsoToWorld(region.StartX, region.StartY))
}

func (v *Engine) GetRegionAt(x, y int) *EngineRegion {
	for _, region := range v.regions {
		if region.Rect.IsInRect(x, y) {
			return &region
		}
	}

	return nil
}

func (v *Engine) Render(target *ebiten.Image) {
	for _, region := range v.regions {
		// X position of leftmost point of region
		left := float64((region.Rect.Left - region.Rect.Bottom()) * 80)
		// Y position of top of region
		top := float64((region.Rect.Left + region.Rect.Top) * 40)
		// X of right
		right := float64((region.Rect.Right() - region.Rect.Top) * 80)
		// Y of bottom
		bottom := float64((region.Rect.Right() + region.Rect.Bottom()) * 40)

		if v.viewport.IsWorldRectVisible(left, top, right, bottom) {
			v.RenderRegion(region, target)
		}
	}
}

func (v *Engine) GenTiles(region *EngineRegion) {
	for y := 0; y < int(region.Region.TileHeight); y++ {
		offX := -((y + region.Rect.Top) * 80) + (region.Rect.Left * 80)
		offY := ((y + region.Rect.Top) * 40) + (region.Rect.Left * 40)
		for x := 0; x < int(region.Region.TileWidth); x++ {
			region.Tiles = append(region.Tiles, RegionTile{
				tileX: x,
				tileY: y,
				offX:  offX,
				offY:  offY,
			})

			offX += 80
			offY += 40
		}
	}
}

func (v *Engine) GenTilesCache(region *EngineRegion) {
	for tileIdx := range region.Tiles {
		t := &region.Tiles[tileIdx]
		if t.tileY < len(region.Region.DS1.Tiles) && t.tileX < len(region.Region.DS1.Tiles[t.tileY]) {
			tile := &region.Region.DS1.Tiles[t.tileY][t.tileX]
			for i := range tile.Floors {
				if tile.Floors[i].Hidden || tile.Floors[i].Prop1 == 0 {
					continue
				}
				region.Region.generateFloorCache(&tile.Floors[i], t.tileX, t.tileY)
			}
			for i := range tile.Shadows {
				if tile.Shadows[i].Hidden || tile.Shadows[i].Prop1 == 0 {
					continue
				}
				region.Region.generateShadowCache(&tile.Shadows[i], t.tileX, t.tileY)
			}
			for i := range tile.Walls {
				if tile.Walls[i].Hidden || tile.Walls[i].Prop1 == 0 {
					continue
				}
				region.Region.generateWallCache(&tile.Walls[i], t.tileX, t.tileY)
			}
		}
	}
}

func (v *Engine) RenderRegion(region EngineRegion, target *ebiten.Image) {
	for tileIdx := range region.Tiles {
		if v.viewport.IsWorldTileVisbile(float64(region.Tiles[tileIdx].tileX+region.Rect.Left), float64(region.Tiles[tileIdx].tileY+region.Rect.Top)) {
			region.Region.UpdateAnimations()

			v.viewport.PushTranslation(float64(region.Tiles[tileIdx].offX), float64(region.Tiles[tileIdx].offY))
			v.RenderPass1(region.Region, region.Tiles[tileIdx].tileX, region.Tiles[tileIdx].tileY, target)
			v.DrawTileLines(region.Region, region.Tiles[tileIdx].tileX, region.Tiles[tileIdx].tileY, target)
			v.viewport.PopTranslation()
		}
	}

	for tileIdx := range region.Tiles {
		if v.viewport.IsWorldTileVisbile(float64(region.Tiles[tileIdx].tileX+region.Rect.Left), float64(region.Tiles[tileIdx].tileY+region.Rect.Top)) {
			v.viewport.PushTranslation(float64(region.Tiles[tileIdx].offX), float64(region.Tiles[tileIdx].offY))
			v.RenderPass2(region.Region, region.Tiles[tileIdx].tileX, region.Tiles[tileIdx].tileY, target)
			v.viewport.PopTranslation()
		}
	}

	for tileIdx := range region.Tiles {
		if v.viewport.IsWorldTileVisbile(float64(region.Tiles[tileIdx].tileX+region.Rect.Left), float64(region.Tiles[tileIdx].tileY+region.Rect.Top)) {
			v.viewport.PushTranslation(float64(region.Tiles[tileIdx].offX), float64(region.Tiles[tileIdx].offY))
			v.RenderPass3(region.Region, region.Tiles[tileIdx].tileX, region.Tiles[tileIdx].tileY, target)
			v.viewport.PopTranslation()
		}
	}
}

func (v *Engine) RenderPass1(region *Region, x, y int, target *ebiten.Image) {
	tile := region.DS1.Tiles[y][x]
	// Draw lower walls
	for i := range tile.Walls {
		if !tile.Walls[i].Type.LowerWall() || tile.Walls[i].Prop1 == 0 || tile.Walls[i].Hidden {
			continue
		}

		region.RenderTile(v.viewport, x, y, d2enum.RegionLayerTypeWalls, i, target)
	}

	for i := range tile.Floors {
		if tile.Floors[i].Hidden || tile.Floors[i].Prop1 == 0 {
			continue
		}

		region.RenderTile(v.viewport, x, y, d2enum.RegionLayerTypeFloors, i, target)
	}
	for i := range tile.Shadows {
		if tile.Shadows[i].Hidden || tile.Shadows[i].Prop1 == 0 {
			continue
		}

		region.RenderTile(v.viewport, x, y, d2enum.RegionLayerTypeShadows, i, target)
	}
}

func (v *Engine) RenderPass2(region *Region, x, y int, target *ebiten.Image) {
	tile := region.DS1.Tiles[y][x]

	// Draw upper walls
	for i := range tile.Walls {
		if !tile.Walls[i].Type.UpperWall() || tile.Walls[i].Hidden {
			continue
		}

		region.RenderTile(v.viewport, x, y, d2enum.RegionLayerTypeWalls, i, target)
	}

	screenX, screenY := v.viewport.WorldToScreen(v.viewport.GetTranslation())

	for _, obj := range region.AnimationEntities {
		if obj.TileX == x && obj.TileY == y {
			obj.Render(target, screenX, screenY)
		}
	}

	for _, npc := range region.NPCs {
		if npc.AnimatedEntity.TileX == x && npc.AnimatedEntity.TileY == y {
			npc.Render(target, screenX, screenY)
		}
	}

	if v.Hero != nil && v.Hero.AnimatedEntity.TileX == x && v.Hero.AnimatedEntity.TileY == y {
		v.Hero.Render(target, screenX, screenY)
	}
}

func (v *Engine) RenderPass3(region *Region, x, y int, target *ebiten.Image) {
	tile := region.DS1.Tiles[y][x]
	// Draw ceilings
	for i := range tile.Walls {
		if tile.Walls[i].Type != d2enum.Roof {
			continue
		}

		region.RenderTile(v.viewport, x, y, d2enum.RegionLayerTypeWalls, i, target)
	}
}

func (v *Engine) DrawTileLines(region *Region, x, y int, target *ebiten.Image) {
	if v.ShowTiles > 0 {
		subtileColor := color.RGBA{80, 80, 255, 100}
		tileColor := color.RGBA{255, 255, 255, 255}

		screenX1, screenY1 := v.viewport.IsoToScreen(float64(x), float64(y))
		screenX2, screenY2 := v.viewport.IsoToScreen(float64(x+1), float64(y))
		screenX3, screenY3 := v.viewport.IsoToScreen(float64(x), float64(y+1))

		ebitenutil.DrawLine(
			target,
			float64(screenX1),
			float64(screenY1),
			float64(screenX2),
			float64(screenY2),
			tileColor,
		)

		ebitenutil.DrawLine(
			target,
			float64(screenX1),
			float64(screenY1),
			float64(screenX3),
			float64(screenY3),
			tileColor,
		)

		ebitenutil.DebugPrintAt(
			target,
			fmt.Sprintf("%v,%v", x, y),
			screenX1-10,
			screenY1+10,
		)

		if v.ShowTiles > 1 {
			for i := 1; i <= 4; i++ {
				x := i * 16
				y := i * 8

				ebitenutil.DrawLine(
					target,
					float64(screenX1-x),
					float64(screenY1+y),
					float64(screenX1-x+80),
					float64(screenY1+y+40),
					subtileColor,
				)

				ebitenutil.DrawLine(
					target,
					float64(screenX1+x),
					float64(screenY1+y),
					float64(screenX1+x-80),
					float64(screenY1+y+40),
					subtileColor,
				)
			}

			tile := region.DS1.Tiles[y][x]
			for i := range tile.Floors {
				ebitenutil.DebugPrintAt(
					target,
					fmt.Sprintf("f: %v-%v", tile.Floors[i].Style, tile.Floors[i].Sequence),
					screenX1-20,
					screenY1+10+(i+1)*14,
				)
			}
		}
	}
}
