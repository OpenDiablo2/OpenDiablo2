package Map

import (
	"image"
	"math/rand"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Sound"
	"github.com/hajimehoshi/ebiten"
)

type EngineRegion struct {
	Rect   image.Rectangle
	Region *Region
}

type Engine struct {
	soundManager *Sound.Manager
	gameState    *Common.GameState
	fileProvider Common.FileProvider
	regions      []EngineRegion
	OffsetX      float64
	OffsetY      float64
}

func CreateMapEngine(gameState *Common.GameState, soundManager *Sound.Manager, fileProvider Common.FileProvider) *Engine {
	result := &Engine{
		gameState:    gameState,
		soundManager: soundManager,
		fileProvider: fileProvider,
		regions:      make([]EngineRegion, 0),
	}
	return result
}

func (v *Engine) GenerateMap(regionType RegionIdType, levelPreset int) {
	randomSource := rand.NewSource(v.gameState.Seed)
	region := LoadRegion(randomSource, regionType, levelPreset, v.fileProvider)
	v.regions = append(v.regions, EngineRegion{
		Rect:   image.Rectangle{image.Point{0, 0}, image.Point{int(region.TileWidth), int(region.TileHeight)}},
		Region: region,
	})
	v.soundManager.PlayBGM("/data/global/music/Act1/tristram.wav") // TODO: Temp stuff here
}

func (v *Engine) Render(target *ebiten.Image) {
	// TODO: Temporary hack for testing
	for y := 0; y < int(v.regions[0].Region.TileHeight); y++ {
		offX := -(y * 80)
		offY := y * 40
		for x := 0; x < int(v.regions[0].Region.TileWidth); x++ {
			sx, sy := Common.IsoToScreen(x, y, int(v.OffsetX), int(v.OffsetY))
			if sx > -160 && sy > -160 && sx <= 800 && sy <= 1000 {
				tile := v.regions[0].Region.DS1.Tiles[y][x]
				for i := range tile.Floors {
					if !tile.Floors[i].Hidden && tile.Floors[i].Prop1 != 0 {
						v.regions[0].Region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, RegionLayerTypeFloors, i, target)
					}
				}
			}
			offX += 80
			offY += 40
		}
	}
	for y := 0; y < int(v.regions[0].Region.TileHeight); y++ {
		offX := -(y * 80)
		offY := y * 40
		for x := 0; x < int(v.regions[0].Region.TileWidth); x++ {
			sx, sy := Common.IsoToScreen(x, y, int(v.OffsetX), int(v.OffsetY))
			if sx > -160 && sy > -160 && sx <= 800 && sy <= 1000 {
				tile := v.regions[0].Region.DS1.Tiles[y][x]
				for i := range tile.Shadows {
					if tile.Shadows[i].Hidden || tile.Shadows[i].Prop1 == 0 {
						continue
					}
					v.regions[0].Region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, RegionLayerTypeShadows, i, target)
				}
			}
			offX += 80
			offY += 40
		}
	}
	for y := 0; y < int(v.regions[0].Region.TileHeight); y++ {
		offX := -(y * 80)
		offY := y * 40
		for x := 0; x < int(v.regions[0].Region.TileWidth); x++ {
			sx, sy := Common.IsoToScreen(x, y, int(v.OffsetX), int(v.OffsetY))
			if sx > -160 && sy > -160 && sx <= 800 && sy <= 1000 {
				tile := v.regions[0].Region.DS1.Tiles[y][x]
				for i := range tile.Walls {
					if tile.Walls[i].Hidden || tile.Walls[i].Orientation == 15 || tile.Walls[i].Orientation == 10 || tile.Walls[i].Orientation == 11 || tile.Walls[i].Orientation == 0 {
						continue
					}
					v.regions[0].Region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, RegionLayerTypeWalls, i, target)
				}
			}
			offX += 80
			offY += 40
		}
	}
	for y := 0; y < int(v.regions[0].Region.TileHeight); y++ {
		offX := -(y * 80)
		offY := y * 40
		for x := 0; x < int(v.regions[0].Region.TileWidth); x++ {
			sx, sy := Common.IsoToScreen(x, y, int(v.OffsetX), int(v.OffsetY))
			if sx > -160 && sy > -160 && sx <= 800 && sy <= 1000 {
				tile := v.regions[0].Region.DS1.Tiles[y][x]
				for i := range tile.Walls {
					if tile.Walls[i].Hidden || tile.Walls[i].Orientation != 15 {
						continue
					}
					v.regions[0].Region.RenderTile(offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, RegionLayerTypeWalls, i, target)
				}
			}
			offX += 80
			offY += 40
		}
	}
}
