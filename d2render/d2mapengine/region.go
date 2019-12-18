package d2mapengine

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/OpenDiablo2/D2Shared/d2common"
	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/D2Shared/d2data/d2ds1"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dt1"
	"github.com/OpenDiablo2/D2Shared/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2corehelper"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
)

type MapRegion struct {
	tileRect          d2common.Rectangle
	regionPath        string
	levelType         d2datadict.LevelTypeRecord
	levelPreset       d2datadict.LevelPresetRecord
	tiles             []d2dt1.Tile
	ds1               d2ds1.DS1
	palette           d2datadict.PaletteRec
	startX            float64
	startY            float64
	imageCacheRecords map[uint32]*ebiten.Image
	seed              int64
	currentFrame      byte
	lastFrameTime     float64
}

func loadRegion(seed int64, tileOffsetX, tileOffsetY int, levelType d2enum.RegionIdType, levelPreset int, fileIndex int) (*MapRegion, []MapEntity) {
	region := &MapRegion{
		levelType:         d2datadict.LevelTypes[levelType],
		levelPreset:       d2datadict.LevelPresets[levelPreset],
		imageCacheRecords: map[uint32]*ebiten.Image{},
		seed:              seed,
	}

	region.palette = d2datadict.Palettes[d2enum.PaletteType("act"+strconv.Itoa(int(region.levelType.Act)))]
	if levelType == d2enum.RegionAct5Lava {
		region.palette = d2datadict.Palettes[d2enum.PaletteType("act4")]
	}

	for _, levelTypeDt1 := range region.levelType.Files {
		if len(levelTypeDt1) != 0 && levelTypeDt1 != "" && levelTypeDt1 != "0" {
			dt1 := d2dt1.LoadDT1(d2asset.MustLoadFile("/data/global/tiles/" + levelTypeDt1))
			region.tiles = append(region.tiles, dt1.Tiles...)
		}
	}

	var levelFilesToPick []string
	for _, fileRecord := range region.levelPreset.Files {
		if len(fileRecord) != 0 && fileRecord != "" && fileRecord != "0" {
			levelFilesToPick = append(levelFilesToPick, fileRecord)
		}
	}

	levelIndex := int(math.Round(float64(len(levelFilesToPick)-1) * rand.Float64()))
	if fileIndex >= 0 && fileIndex < len(levelFilesToPick) {
		levelIndex = fileIndex
	}

	region.regionPath = levelFilesToPick[levelIndex]
	region.ds1 = d2ds1.LoadDS1(d2asset.MustLoadFile("/data/global/tiles/" + region.regionPath))
	region.tileRect = d2common.Rectangle{
		Left:   tileOffsetX,
		Top:    tileOffsetY,
		Width:  int(region.ds1.Width),
		Height: int(region.ds1.Height),
	}

	entities := region.loadEntities()
	region.loadSpecials()
	region.generateTileCache()

	return region, entities
}

func (mr *MapRegion) GetTileRect() d2common.Rectangle {
	return mr.tileRect
}

func (mr *MapRegion) GetLevelPreset() d2datadict.LevelPresetRecord {
	return mr.levelPreset
}

func (mr *MapRegion) GetLevelType() d2datadict.LevelTypeRecord {
	return mr.levelType
}

func (mr *MapRegion) GetPath() string {
	return mr.regionPath
}

func (mr *MapRegion) loadSpecials() {
	for tileY := range mr.ds1.Tiles {
		for tileX := range mr.ds1.Tiles[tileY] {
			for _, wall := range mr.ds1.Tiles[tileY][tileX].Walls {
				if wall.Type == 10 && wall.Style == 30 && wall.Sequence == 0 {
					mr.startX, mr.startY = mr.getTileWorldPosition(tileX, tileY)
					mr.startX += 0.5
					mr.startY += 0.5
					return
				}
			}
		}
	}
}

func (mr *MapRegion) loadEntities() []MapEntity {
	var entities []MapEntity

	for _, object := range mr.ds1.Objects {
		worldX, worldY := mr.getTileWorldPosition(int(object.X), int(object.Y))

		switch object.Lookup.Type {
		case d2datadict.ObjectTypeCharacter:
			if object.Lookup.Base != "" && object.Lookup.Token != "" && object.Lookup.TR != "" {
				npc := d2core.CreateNPC(int32(worldX), int32(worldY), object.Lookup, 0)
				npc.SetPaths(object.Paths)
				entities = append(entities, npc)
			}
		case d2datadict.ObjectTypeItem:
			if object.ObjectInfo != nil && object.ObjectInfo.Draw && object.Lookup.Base != "" && object.Lookup.Token != "" {
				entity := d2render.CreateAnimatedEntity(int32(worldX), int32(worldY), object.Lookup, d2enum.Units)
				entity.SetMode(object.Lookup.Mode, object.Lookup.Class, 0)
				entities = append(entities, &entity)
			}
		}
	}

	return entities
}

func (mr *MapRegion) updateAnimations() {
	now := d2helper.Now()
	framesToAdd := math.Floor((now - mr.lastFrameTime) / 0.1)
	if framesToAdd > 0 {
		mr.lastFrameTime += 0.1 * framesToAdd
		mr.currentFrame += byte(math.Floor(framesToAdd))
		if mr.currentFrame > 9 {
			mr.currentFrame = 0
		}
	}
}

func (mr *MapRegion) getStartTilePosition() (float64, float64) {
	return float64(mr.tileRect.Left) + mr.startX, float64(mr.tileRect.Top) + mr.startY
}

func (mr *MapRegion) getRandomTile(tiles []d2dt1.Tile, x, y int, seed int64) byte {
	/* Walker's Alias Method for weighted random selection
	 * with xorshifting for random numbers */

	var tileSeed uint64
	tileSeed = uint64(seed) + uint64(x)
	tileSeed *= uint64(y) + uint64(mr.levelType.Id)

	tileSeed ^= tileSeed << 13
	tileSeed ^= tileSeed >> 17
	tileSeed ^= tileSeed << 5

	weightSum := 0
	for _, tile := range tiles {
		weightSum += int(tile.RarityFrameIndex)
	}

	if weightSum == 0 {
		return 0
	}

	random := (tileSeed % uint64(weightSum))

	sum := 0
	for i, tile := range tiles {
		sum += int(tile.RarityFrameIndex)
		if sum >= int(random) {
			return byte(i)
		}
	}

	// This return shouldn't be hit
	return 0
}

func (mr *MapRegion) getTiles(style, sequence, tileType int32, x, y int, seed int64) []d2dt1.Tile {
	var tiles []d2dt1.Tile
	for _, tile := range mr.tiles {
		if tile.Style != style || tile.Sequence != sequence || tile.Type != tileType {
			continue
		}
		tiles = append(tiles, tile)
	}
	if len(tiles) == 0 {
		log.Printf("Unknown tile ID [%d %d %d]\n", style, sequence, tileType)
		return nil
	}
	return tiles
}

func (mr *MapRegion) isVisbile(viewport *Viewport) bool {
	return viewport.IsTileRectVisible(mr.tileRect)
}

func (mr *MapRegion) advance(tickTime float64) {
	mr.updateAnimations()
}

func (mr *MapRegion) getTileWorldPosition(tileX, tileY int) (float64, float64) {
	return float64(tileX + mr.tileRect.Left), float64(tileY + mr.tileRect.Top)
}

func (mr *MapRegion) renderPass1(viewport *Viewport, target *ebiten.Image) {
	for tileY := range mr.ds1.Tiles {
		for tileX, tile := range mr.ds1.Tiles[tileY] {
			worldX, worldY := mr.getTileWorldPosition(tileX, tileY)
			if viewport.IsTileVisible(worldX, worldY) {
				viewport.PushTranslationWorld(worldX, worldY)
				mr.renderTilePass1(tile, viewport, target)
				viewport.PopTranslation()
			}
		}
	}
}

func (mr *MapRegion) renderPass2(entities []MapEntity, viewport *Viewport, target *ebiten.Image) {
	for tileY := range mr.ds1.Tiles {
		for tileX, tile := range mr.ds1.Tiles[tileY] {
			worldX, worldY := mr.getTileWorldPosition(tileX, tileY)
			if viewport.IsTileVisible(worldX, worldY) {
				viewport.PushTranslationWorld(worldX, worldY)
				mr.renderTilePass2(tile, viewport, target)

				for _, entity := range entities {
					entWorldX, entWorldY := entity.GetPosition()
					if entWorldX == worldX && entWorldY == worldY {
						screenX, screenY := viewport.GetTranslationScreen()
						entity.Render(target, screenX, screenY)
					}
				}

				viewport.PopTranslation()
			}
		}
	}
}

func (mr *MapRegion) renderPass3(viewport *Viewport, target *ebiten.Image) {
	for tileY := range mr.ds1.Tiles {
		for tileX, tile := range mr.ds1.Tiles[tileY] {
			worldX, worldY := mr.getTileWorldPosition(tileX, tileY)
			if viewport.IsTileVisible(worldX, worldY) {
				viewport.PushTranslationWorld(worldX, worldY)
				mr.renderTilePass3(tile, viewport, target)
				viewport.PopTranslation()
			}
		}
	}
}

func (mr *MapRegion) renderTilePass1(tile d2ds1.TileRecord, viewport *Viewport, target *ebiten.Image) {
	for _, wall := range tile.Walls {
		if !wall.Hidden && wall.Prop1 != 0 && wall.Type.LowerWall() {
			mr.renderWall(wall, viewport, target)
		}
	}

	for _, floor := range tile.Floors {
		if !floor.Hidden && floor.Prop1 != 0 {
			mr.renderFloor(floor, viewport, target)
		}
	}

	for _, shadow := range tile.Shadows {
		if !shadow.Hidden && shadow.Prop1 != 0 {
			mr.renderShadow(shadow, viewport, target)
		}
	}
}

func (mr *MapRegion) renderTilePass2(tile d2ds1.TileRecord, viewport *Viewport, target *ebiten.Image) {
	for _, wall := range tile.Walls {
		if !wall.Hidden && wall.Type.UpperWall() {
			mr.renderWall(wall, viewport, target)
		}
	}
}

func (mr *MapRegion) renderTilePass3(tile d2ds1.TileRecord, viewport *Viewport, target *ebiten.Image) {
	for _, wall := range tile.Walls {
		if wall.Type == d2enum.Roof {
			mr.renderWall(wall, viewport, target)
		}
	}
}

func (mr *MapRegion) renderFloor(tile d2ds1.FloorShadowRecord, viewport *Viewport, target *ebiten.Image) {
	var img *ebiten.Image
	if !tile.Animated {
		img = mr.getImageCacheRecord(tile.Style, tile.Sequence, 0, tile.RandomIndex)
	} else {
		img = mr.getImageCacheRecord(tile.Style, tile.Sequence, 0, mr.currentFrame)
	}
	if img == nil {
		log.Printf("Render called on uncached floor {%v,%v}", tile.Style, tile.Sequence)
		return
	}

	viewport.PushTranslationOrtho(-80, float64(tile.YAdjust))
	screenX, screenY := viewport.GetTranslationScreen()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(screenX), float64(screenY))
	target.DrawImage(img, opts)
	viewport.PopTranslation()
}

func (mr *MapRegion) renderWall(tile d2ds1.WallRecord, viewport *Viewport, target *ebiten.Image) {
	img := mr.getImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tile.RandomIndex)
	if img == nil {
		log.Printf("Render called on uncached wall {%v,%v,%v}", tile.Style, tile.Sequence, tile.Type)
		return
	}

	viewport.PushTranslationOrtho(-80, float64(tile.YAdjust))
	screenX, screenY := viewport.GetTranslationScreen()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(screenX), float64(screenY))
	target.DrawImage(img, opts)
	viewport.PopTranslation()
}

func (mr *MapRegion) renderShadow(tile d2ds1.FloorShadowRecord, viewport *Viewport, target *ebiten.Image) {
	img := mr.getImageCacheRecord(tile.Style, tile.Sequence, 13, tile.RandomIndex)
	if img == nil {
		log.Printf("Render called on uncached shadow {%v,%v}", tile.Style, tile.Sequence)
		return
	}

	viewport.PushTranslationOrtho(-80, float64(tile.YAdjust))
	screenX, screenY := viewport.GetTranslationScreen()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(screenX), float64(screenY))
	opts.ColorM = d2corehelper.ColorToColorM(color.RGBA{255, 255, 255, 160})
	target.DrawImage(img, opts)
	viewport.PopTranslation()
}

func (mr *MapRegion) renderDebug(debugVisLevel int, viewport *Viewport, target *ebiten.Image) {
	for tileY := range mr.ds1.Tiles {
		for tileX := range mr.ds1.Tiles[tileY] {
			worldX, worldY := mr.getTileWorldPosition(tileX, tileY)
			if viewport.IsTileVisible(worldX, worldY) {
				mr.renderTileDebug(int(worldX), int(worldY), debugVisLevel, viewport, target)
			}
		}
	}
}

func (mr *MapRegion) renderTileDebug(x, y int, debugVisLevel int, viewport *Viewport, target *ebiten.Image) {
	if debugVisLevel > 0 {
		subtileColor := color.RGBA{80, 80, 255, 100}
		tileColor := color.RGBA{255, 255, 255, 255}

		screenX1, screenY1 := viewport.WorldToScreen(float64(x), float64(y))
		screenX2, screenY2 := viewport.WorldToScreen(float64(x+1), float64(y))
		screenX3, screenY3 := viewport.WorldToScreen(float64(x), float64(y+1))

		ebitenutil.DrawLine(target, float64(screenX1), float64(screenY1), float64(screenX2), float64(screenY2), tileColor)
		ebitenutil.DrawLine(target, float64(screenX1), float64(screenY1), float64(screenX3), float64(screenY3), tileColor)
		ebitenutil.DebugPrintAt(target, fmt.Sprintf("%v,%v", x, y), screenX1-10, screenY1+10)

		if debugVisLevel > 1 {
			for i := 1; i <= 4; i++ {
				x := i * 16
				y := i * 8

				ebitenutil.DrawLine(target, float64(screenX1-x), float64(screenY1+y), float64(screenX1-x+80), float64(screenY1+y+40), subtileColor)
				ebitenutil.DrawLine(target, float64(screenX1+x), float64(screenY1+y), float64(screenX1+x-80), float64(screenY1+y+40), subtileColor)
			}

			tile := mr.ds1.Tiles[y][x]
			for i := range tile.Floors {
				ebitenutil.DebugPrintAt(target, fmt.Sprintf("f: %v-%v", tile.Floors[i].Style, tile.Floors[i].Sequence), screenX1-20, screenY1+10+(i+1)*14)
			}
		}
	}
}

func (mr *MapRegion) generateTileCache() {
	for tileY := range mr.ds1.Tiles {
		for tileX := range mr.ds1.Tiles[tileY] {
			tile := mr.ds1.Tiles[tileY][tileX]

			for i := range tile.Floors {
				if !tile.Floors[i].Hidden && tile.Floors[i].Prop1 != 0 {
					mr.generateFloorCache(&tile.Floors[i], tileX, tileY)
				}
			}

			for i := range tile.Shadows {
				if !tile.Shadows[i].Hidden && tile.Shadows[i].Prop1 != 0 {
					mr.generateShadowCache(&tile.Shadows[i], tileX, tileY)
				}
			}

			for i := range tile.Walls {
				if !tile.Walls[i].Hidden && tile.Walls[i].Prop1 != 0 {
					mr.generateWallCache(&tile.Walls[i], tileX, tileY)
				}
			}
		}
	}
}

func (mr *MapRegion) getImageCacheRecord(style, sequence byte, tileType d2enum.TileType, randomIndex byte) *ebiten.Image {
	lookupIndex := uint32(style)<<24 | uint32(sequence)<<16 | uint32(tileType)<<8 | uint32(randomIndex)
	return mr.imageCacheRecords[lookupIndex]
}

func (mr *MapRegion) setImageCacheRecord(style, sequence byte, tileType d2enum.TileType, randomIndex byte, image *ebiten.Image) {
	lookupIndex := uint32(style)<<24 | uint32(sequence)<<16 | uint32(tileType)<<8 | uint32(randomIndex)
	mr.imageCacheRecords[lookupIndex] = image
}

func (mr *MapRegion) generateFloorCache(tile *d2ds1.FloorShadowRecord, tileX, tileY int) {
	tileOptions := mr.getTiles(int32(tile.Style), int32(tile.Sequence), 0, tileX, tileY, mr.seed)
	var tileData []*d2dt1.Tile
	var tileIndex byte

	if tileOptions == nil {
		log.Printf("Could not locate tile Style:%d, Seq: %d, Type: %d\n", tile.Style, tile.Sequence, 0)
		tileData = append(tileData, &d2dt1.Tile{})
		tileData[0].Width = 10
		tileData[0].Height = 10
	} else {
		if !tileOptions[0].MaterialFlags.Animated {
			tileIndex = mr.getRandomTile(tileOptions, tileX, tileY, mr.seed)
			tileData = append(tileData, &tileOptions[tileIndex])
		} else {
			tile.Animated = true
			for i := range tileOptions {
				tileData = append(tileData, &tileOptions[i])
			}
		}
	}

	for i := range tileData {
		if !tileData[i].MaterialFlags.Animated {
			tile.RandomIndex = tileIndex
		} else {
			tileIndex = byte(tileData[i].RarityFrameIndex)
		}
		cachedImage := mr.getImageCacheRecord(tile.Style, tile.Sequence, 0, tileIndex)
		if cachedImage != nil {
			return
		}
		tileYMinimum := int32(0)
		for _, block := range tileData[i].Blocks {
			tileYMinimum = d2helper.MinInt32(tileYMinimum, int32(block.Y))
		}
		tileYOffset := d2helper.AbsInt32(tileYMinimum)
		tileHeight := d2helper.AbsInt32(tileData[i].Height)
		image, _ := ebiten.NewImage(int(tileData[i].Width), int(tileHeight), ebiten.FilterNearest)
		pixels := make([]byte, 4*tileData[i].Width*tileHeight)
		mr.decodeTileGfxData(tileData[i].Blocks, &pixels, tileYOffset, tileData[i].Width)
		image.ReplacePixels(pixels)
		mr.setImageCacheRecord(tile.Style, tile.Sequence, 0, tileIndex, image)
	}
}

func (mr *MapRegion) generateShadowCache(tile *d2ds1.FloorShadowRecord, tileX, tileY int) {
	tileOptions := mr.getTiles(int32(tile.Style), int32(tile.Sequence), 13, tileX, tileY, mr.seed)
	var tileIndex byte
	var tileData *d2dt1.Tile
	if tileOptions == nil {
		return
	} else {
		tileIndex = mr.getRandomTile(tileOptions, tileX, tileY, mr.seed)
		tileData = &tileOptions[tileIndex]
	}

	tile.RandomIndex = tileIndex
	tileMinY := int32(0)
	tileMaxY := int32(0)
	for _, block := range tileData.Blocks {
		tileMinY = d2helper.MinInt32(tileMinY, int32(block.Y))
		tileMaxY = d2helper.MaxInt32(tileMaxY, int32(block.Y+32))
	}
	tileYOffset := -tileMinY
	tileHeight := int(tileMaxY - tileMinY)
	tile.YAdjust = int(tileMinY + 80)

	cachedImage := mr.getImageCacheRecord(tile.Style, tile.Sequence, 13, tileIndex)
	if cachedImage != nil {
		return
	}

	image, _ := ebiten.NewImage(int(tileData.Width), int(tileHeight), ebiten.FilterNearest)
	pixels := make([]byte, 4*tileData.Width*int32(tileHeight))
	mr.decodeTileGfxData(tileData.Blocks, &pixels, tileYOffset, tileData.Width)
	image.ReplacePixels(pixels)
	mr.setImageCacheRecord(tile.Style, tile.Sequence, 13, tileIndex, image)
}

func (mr *MapRegion) generateWallCache(tile *d2ds1.WallRecord, tileX, tileY int) {
	tileOptions := mr.getTiles(int32(tile.Style), int32(tile.Sequence), int32(tile.Type), tileX, tileY, mr.seed)
	var tileIndex byte
	var tileData *d2dt1.Tile
	if tileOptions == nil {
		return
	} else {
		tileIndex = mr.getRandomTile(tileOptions, tileX, tileY, mr.seed)
		tileData = &tileOptions[tileIndex]
	}

	tile.RandomIndex = tileIndex
	var newTileData *d2dt1.Tile = nil

	if tile.Type == 3 {
		newTileOptions := mr.getTiles(int32(tile.Style), int32(tile.Sequence), int32(4), tileX, tileY, mr.seed)
		newTileIndex := mr.getRandomTile(newTileOptions, tileX, tileY, mr.seed)
		newTileData = &newTileOptions[newTileIndex]
	}

	tileMinY := int32(0)
	tileMaxY := int32(0)

	target := tileData

	if newTileData != nil && newTileData.Height < tileData.Height {
		target = newTileData
	}

	for _, block := range target.Blocks {
		tileMinY = d2helper.MinInt32(tileMinY, int32(block.Y))
		tileMaxY = d2helper.MaxInt32(tileMaxY, int32(block.Y+32))
	}

	realHeight := d2helper.MaxInt32(d2helper.AbsInt32(tileData.Height), tileMaxY-tileMinY)
	tileYOffset := -tileMinY
	//tileHeight := int(tileMaxY - tileMinY)

	if tile.Type == 15 {
		tile.YAdjust = -int(tileData.RoofHeight)
	} else {
		tile.YAdjust = int(tileMinY) + 80
	}

	cachedImage := mr.getImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tileIndex)
	if cachedImage != nil {
		return
	}

	if realHeight == 0 {
		log.Printf("Invalid 0 height for wall tile")
		return
	}

	image, _ := ebiten.NewImage(160, int(realHeight), ebiten.FilterNearest)
	pixels := make([]byte, 4*160*realHeight)
	mr.decodeTileGfxData(tileData.Blocks, &pixels, tileYOffset, 160)

	if newTileData != nil {
		mr.decodeTileGfxData(newTileData.Blocks, &pixels, tileYOffset, 160)
	}

	if err := image.ReplacePixels(pixels); err != nil {
		log.Panicf(err.Error())
	}

	mr.setImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tileIndex, image)
}

func (mr *MapRegion) decodeTileGfxData(blocks []d2dt1.Block, pixels *[]byte, tileYOffset int32, tileWidth int32) {
	for _, block := range blocks {
		if block.Format == d2dt1.BlockFormatIsometric {
			// 3D isometric decoding
			xjump := []int32{14, 12, 10, 8, 6, 4, 2, 0, 2, 4, 6, 8, 10, 12, 14}
			nbpix := []int32{4, 8, 12, 16, 20, 24, 28, 32, 28, 24, 20, 16, 12, 8, 4}
			blockX := int32(block.X)
			blockY := int32(block.Y)
			length := int32(256)
			x := int32(0)
			y := int32(0)
			idx := 0
			for length > 0 {
				x = xjump[y]
				n := nbpix[y]
				length -= n
				for n > 0 {
					colorIndex := block.EncodedData[idx]
					if colorIndex != 0 {
						pixelColor := mr.palette.Colors[colorIndex]
						offset := 4 * (((blockY + y + tileYOffset) * tileWidth) + (blockX + x))
						(*pixels)[offset] = pixelColor.R
						(*pixels)[offset+1] = pixelColor.G
						(*pixels)[offset+2] = pixelColor.B
						(*pixels)[offset+3] = 255
					}
					x++
					n--
					idx++
				}
				y++
			}
		} else {
			// RLE Encoding
			blockX := int32(block.X)
			blockY := int32(block.Y)
			x := int32(0)
			y := int32(0)
			idx := 0
			length := block.Length
			for length > 0 {
				b1 := block.EncodedData[idx]
				b2 := block.EncodedData[idx+1]
				idx += 2
				length -= 2
				if (b1 | b2) == 0 {
					x = 0
					y++
					continue
				}
				x += int32(b1)
				length -= int32(b2)
				for b2 > 0 {
					colorIndex := block.EncodedData[idx]
					if colorIndex != 0 {
						pixelColor := mr.palette.Colors[colorIndex]
						offset := 4 * (((blockY + y + tileYOffset) * tileWidth) + (blockX + x))
						(*pixels)[offset] = pixelColor.R
						(*pixels)[offset+1] = pixelColor.G
						(*pixels)[offset+2] = pixelColor.B
						(*pixels)[offset+3] = 255

					}
					idx++
					x++
					b2--
				}
			}
		}
	}
}
