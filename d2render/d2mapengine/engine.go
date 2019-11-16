package d2mapengine

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
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
	regions      []EngineRegion
	OffsetX      float64
	OffsetY      float64
	ShowTiles    int
}

func CreateMapEngine(gameState *d2core.GameState, soundManager *d2audio.Manager, fileProvider d2interface.FileProvider) *Engine {
	result := &Engine{
		gameState:    gameState,
		soundManager: soundManager,
		fileProvider: fileProvider,
		regions:      make([]EngineRegion, 0),
	}
	return result
}

func (v *Engine) GenerateMap(regionType d2enum.RegionIdType, levelPreset int) {
	randomSource := rand.NewSource(v.gameState.Seed)
	region := LoadRegion(randomSource, regionType, levelPreset, v.fileProvider)
	v.regions = append(v.regions, EngineRegion{
		Rect:   d2common.Rectangle{0, 0, int(region.TileWidth), int(region.TileHeight)},
		Region: region,
	})

	for i, _ := range v.regions {
		v.GenTiles(&v.regions[i])
	}
}

func (v *Engine) GenerateAct1Overworld() {
	v.soundManager.PlayBGM("/data/global/music/Act1/town1.wav") // TODO: Temp stuff here
	randomSource := rand.NewSource(v.gameState.Seed)
	region := LoadRegion(randomSource, d2enum.RegionAct1Town, 1, v.fileProvider)
	v.regions = append(v.regions, EngineRegion{
		Rect:   d2common.Rectangle{0, 0, int(region.TileWidth), int(region.TileHeight)},
		Region: region,
	})
	if strings.Contains(region.RegionPath, "E1") {
		region2 := LoadRegion(randomSource, d2enum.RegionAct1Town, 2, v.fileProvider)
		v.regions = append(v.regions, EngineRegion{
			Rect:   d2common.Rectangle{int(region.TileWidth - 1), 0, int(region2.TileWidth), int(region2.TileHeight)},
			Region: region2,
		})
	} else if strings.Contains(region.RegionPath, "S1") {
		region2 := LoadRegion(randomSource, d2enum.RegionAct1Town, 3, v.fileProvider)
		v.regions = append(v.regions, EngineRegion{
			Rect:   d2common.Rectangle{0, int(region.TileHeight - 1), int(region2.TileWidth), int(region2.TileHeight)},
			Region: region2,
		})
	}

	for i, _ := range v.regions {
		v.GenTiles(&v.regions[i])
	}

	sx, sy := d2helper.IsoToScreen(int(region.StartX), int(region.StartY), 0, 0)
	v.OffsetX = float64(sx) - 400
	v.OffsetY = float64(sy) - 300
}

func (v *Engine) GetRegionAt(x, y int) *EngineRegion {
	if v.regions == nil {
		return nil
	}
	for _, region := range v.regions {
		if !region.Rect.IsInRect(x, y) {
			continue
		}
		return &region
	}
	return nil
}

func (v *Engine) Render(target *ebiten.Image) {
	for _, region := range v.regions {
		v.RenderRegion(region, target)
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

func (v *Engine) RenderRegion(region EngineRegion, target *ebiten.Image) {
	//tilesToRender := make([]RegionTile, len(region.Tiles))

	for tileIdx := range region.Tiles {
		sx, sy := d2helper.IsoToScreen(region.Tiles[tileIdx].tileX+region.Rect.Left, region.Tiles[tileIdx].tileY+region.Rect.Top, int(v.OffsetX), int(v.OffsetY))
		if sx > -160 && sy > -160 && sx <= 880 && sy <= 1000 {
			v.RenderTile(region.Region, region.Tiles[tileIdx].offX, region.Tiles[tileIdx].offY, region.Tiles[tileIdx].tileX, region.Tiles[tileIdx].tileY, target)
		}
	}

	//for _, tile := range tilesToRender {
	//	v.RenderTileObjects(region.Region, tile.offX, tile.offY, tile.tileX, tile.tileY, target)
	//}
}

func (v *Engine) RenderTileObjects(region *Region, offX, offY, x, y int, target *ebiten.Image) {
	tile := region.DS1.Tiles[y][x]
	for _, obj := range region.AnimationEntities {
		if int(math.Floor(obj.LocationX)) == x && int(math.Floor(obj.LocationY)) == y {
			obj.Render(target, offX+int(v.OffsetX), offY+int(v.OffsetY))
		}
	}
	for _, npc := range region.NPCs {
		if int(math.Floor(npc.AnimatedEntity.LocationX)) == x && int(math.Floor(npc.AnimatedEntity.LocationY)) == y {
			npc.Render(target, offX+int(v.OffsetX), offY+int(v.OffsetY))
		}
	}
	for i := range tile.Walls {
		if tile.Walls[i].Hidden || tile.Walls[i].Orientation != 15 {
			continue
		}
		region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, d2enum.RegionLayerTypeWalls, i, target)
	}

	if v.ShowTiles > 0 {
		subtileColor := color.RGBA{255, 100, 100, 140}
		tileColor := color.RGBA{255, 255, 255, 255}

		ebitenutil.DrawLine(target, float64(offX)+v.OffsetX, float64(offY)+v.OffsetY, float64(offX)+v.OffsetX+80, float64(offY)+v.OffsetY+40, tileColor)
		ebitenutil.DrawLine(target, float64(offX)+v.OffsetX, float64(offY)+v.OffsetY, float64(offX)+v.OffsetX-80, float64(offY)+v.OffsetY+40, tileColor)

		coords := fmt.Sprintf("%v,%v", x, y)
		ebitenutil.DebugPrintAt(target, coords, offX+int(v.OffsetX)-10, offY+int(v.OffsetY)+10)

		if v.ShowTiles > 1 {
			for i := 1; i <= 4; i++ {
				x := (16 * i)
				y := (8 * i)
				ebitenutil.DrawLine(target, float64(offX-x)+v.OffsetX, float64(offY+y)+v.OffsetY,
					float64(offX-x)+v.OffsetX+80, float64(offY+y)+v.OffsetY+40, subtileColor)
				ebitenutil.DrawLine(target, float64(offX+x)+v.OffsetX, float64(offY+y)+v.OffsetY,
					float64(offX+x)+v.OffsetX-80, float64(offY+y)+v.OffsetY+40, subtileColor)
			}
		}
	}
}

func (v *Engine) RenderTile(region *Region, offX, offY, x, y int, target *ebiten.Image) {
	tile := region.DS1.Tiles[y][x]
	for i := range tile.Floors {
		if tile.Floors[i].Hidden || tile.Floors[i].Prop1 == 0 {
			continue
		}
		region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, d2enum.RegionLayerTypeFloors, i, target)
	}
	for i := range tile.Shadows {
		if tile.Shadows[i].Hidden || tile.Shadows[i].Prop1 == 0 {
			continue
		}
		region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, d2enum.RegionLayerTypeShadows, i, target)
	}
	for i := range tile.Walls {
		if tile.Walls[i].Orientation >= 15 {
			// Upper walls only
			continue
		}
		if tile.Walls[i].Hidden || tile.Walls[i].Orientation == 15 || tile.Walls[i].Orientation == 10 || tile.Walls[i].Orientation == 11 || tile.Walls[i].Orientation == 0 {
			continue
		}
		region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, d2enum.RegionLayerTypeWalls, i, target)
	}
	v.RenderTileObjects(region, offX, offY, x, y, target)
	for i := range tile.Walls {
		if tile.Walls[i].Orientation <= 15 {
			// Lower walls only
			continue
		}
		if tile.Walls[i].Hidden || tile.Walls[i].Orientation == 15 || tile.Walls[i].Orientation == 10 || tile.Walls[i].Orientation == 11 || tile.Walls[i].Orientation == 0 {
			continue
		}
		region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, d2enum.RegionLayerTypeWalls, i, target)
	}
	for i := range tile.Walls {
		if tile.Walls[i].Orientation != 15 {
			// Ceiling only
			continue
		}
		region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, d2enum.RegionLayerTypeWalls, i, target)
	}
}
