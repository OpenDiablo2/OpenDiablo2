package d2map

import (
	"errors"
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/beefsack/go-astar"
)

type PathTile struct {
	Walkable                                                    bool
	Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight *PathTile
	X, Y                                                        float64
}

func (t *PathTile) PathNeighbors() []astar.Pather {
	result := make([]astar.Pather, 0)
	if t.Up != nil {
		result = append(result, t.Up)
	}
	if t.Right != nil {
		result = append(result, t.Right)
	}
	if t.Down != nil {
		result = append(result, t.Down)
	}
	if t.Left != nil {
		result = append(result, t.Left)
	}
	if t.UpLeft != nil {
		result = append(result, t.UpLeft)
	}
	if t.UpRight != nil {
		result = append(result, t.UpRight)
	}
	if t.DownLeft != nil {
		result = append(result, t.DownLeft)
	}
	if t.DownRight != nil {
		result = append(result, t.DownRight)
	}

	return result
}

func (t *PathTile) PathNeighborCost(to astar.Pather) float64 {
	return 1 // No cost specifics currently...
}

func (t *PathTile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*PathTile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	r := absX + absY

	return r
}

type MapRegion struct {
	tileRect          d2common.Rectangle
	regionPath        string
	levelType         d2datadict.LevelTypeRecord
	levelPreset       d2datadict.LevelPresetRecord
	tiles             []d2dt1.Tile
	ds1               *d2ds1.DS1
	palette           *d2dat.DATPalette
	startX            float64
	startY            float64
	imageCacheRecords map[uint32]d2render.Surface
	seed              int64
	currentFrame      int
	lastFrameTime     float64
	walkableArea      [][]PathTile
}

func loadRegion(seed int64, tileOffsetX, tileOffsetY int, levelType d2enum.RegionIdType, levelPreset int, fileIndex int) (*MapRegion, []MapEntity) {
	region := &MapRegion{
		levelType:         d2datadict.LevelTypes[levelType],
		levelPreset:       d2datadict.LevelPresets[levelPreset],
		imageCacheRecords: map[uint32]d2render.Surface{},
		seed:              seed,
	}

	region.palette, _ = loadPaletteForAct(levelType)

	for _, levelTypeDt1 := range region.levelType.Files {
		if len(levelTypeDt1) != 0 && levelTypeDt1 != "" && levelTypeDt1 != "0" {
			fileData, err := d2asset.LoadFile("/data/global/tiles/" + levelTypeDt1)
			if err != nil {
				panic(err)
			}

			dt1, _ := d2dt1.LoadDT1(fileData)
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

	if levelFilesToPick == nil {
		panic("no level files to pick from")
	}

	region.regionPath = levelFilesToPick[levelIndex]
	fileData, err := d2asset.LoadFile("/data/global/tiles/" + region.regionPath)
	if err != nil {
		panic(err)
	}
	region.ds1, _ = d2ds1.LoadDS1(fileData)
	region.tileRect = d2common.Rectangle{
		Left:   tileOffsetX,
		Top:    tileOffsetY,
		Width:  int(region.ds1.Width),
		Height: int(region.ds1.Height),
	}

	entities := region.loadEntities()
	region.loadSpecials()
	region.generateTileCache()
	region.generateWalkableMatrix()
	return region, entities
}

func (mr *MapRegion) generateWalkableMatrix() {
	mr.walkableArea = make([][]PathTile, mr.tileRect.Height*5)
	for y := 0; y < (mr.tileRect.Height-1)*5; y++ {
		mr.walkableArea[y] = make([]PathTile, mr.tileRect.Width*5)
		ty := int(float64(y) / 5.0)
		for x := 0; x < (mr.tileRect.Width-1)*5; x++ {
			tx := int(float64(x) / 5.0)
			tile := mr.GetTile(tx, ty)
			isBlocked := false
			for _, floor := range tile.Floors {
				tileData := mr.GetTileData(int32(floor.Style), int32(floor.Sequence), d2enum.Floor)
				tileSubAttrs := &d2dt1.SubTileFlags{}
				if tileData != nil {
					tileSubAttrs = tileData.GetSubTileFlags(x%5, y%5)
				}
				isBlocked = isBlocked || tileSubAttrs.BlockWalk
				if isBlocked {
					break
				}
			}
			if !isBlocked {
				for _, wall := range tile.Walls {
					tileData := mr.GetTileData(int32(wall.Style), int32(wall.Sequence), d2enum.TileType(wall.Type))
					tileSubAttrs := &d2dt1.SubTileFlags{}
					if tileData != nil {
						tileSubAttrs = tileData.GetSubTileFlags(x%5, y%5)
					}
					isBlocked = isBlocked || tileSubAttrs.BlockWalk
					if isBlocked {
						break
					}
				}
			}
			mr.walkableArea[y][x] = PathTile{
				Walkable: !isBlocked,
				X:        float64(x) / 5.0,
				Y:        float64(y) / 5.0,
			}
			if !isBlocked && y > 0 && mr.walkableArea[y-1][x].Walkable {
				mr.walkableArea[y][x].Up = &mr.walkableArea[y-1][x]
				mr.walkableArea[y-1][x].Down = &mr.walkableArea[y][x]
			}
			if !isBlocked && x > 0 && mr.walkableArea[y][x-1].Walkable {
				mr.walkableArea[y][x].Left = &mr.walkableArea[y][x-1]
				mr.walkableArea[y][x-1].Right = &mr.walkableArea[y][x]
			}
			if !isBlocked && x > 0 && y > 0 && mr.walkableArea[y-1][x-1].Walkable {
				mr.walkableArea[y][x].UpLeft = &mr.walkableArea[y-1][x-1]
				mr.walkableArea[y-1][x-1].DownRight = &mr.walkableArea[y][x]
			}
			if !isBlocked && y > 0 && x < (mr.tileRect.Width*5) && mr.walkableArea[y-1][x+1].Walkable {
				mr.walkableArea[y][x].UpRight = &mr.walkableArea[y-1][x+1]
				mr.walkableArea[y-1][x+1].DownLeft = &mr.walkableArea[y][x]
			}
		}
	}
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

func (mr *MapRegion) GetTile(x, y int) *d2ds1.TileRecord {
	return &mr.ds1.Tiles[y][x]
}

func (mr *MapRegion) GetTileData(style int32, sequence int32, tileType d2enum.TileType) *d2dt1.Tile {
	for _, tile := range mr.tiles {
		if tile.Style == style && tile.Sequence == sequence && tile.Type == int32(tileType) {
			return &tile
		}
	}
	return nil
}

func (mr *MapRegion) GetTileSize() (int, int) {
	return mr.tileRect.Width, mr.tileRect.Height
}

func (mr *MapRegion) loadEntities() []MapEntity {
	var entities []MapEntity

	for _, object := range mr.ds1.Objects {
		worldX, worldY := mr.getTileWorldPosition(object.X, object.Y)

		switch object.Lookup.Type {
		case d2datadict.ObjectTypeCharacter:
			if object.Lookup.Base != "" && object.Lookup.Token != "" && object.Lookup.TR != "" {
				npc := CreateNPC(int(worldX), int(worldY), object.Lookup, 0)
				npc.SetPaths(object.Paths)
				entities = append(entities, npc)
			}
		case d2datadict.ObjectTypeItem:
			if object.ObjectInfo != nil && object.ObjectInfo.Draw && object.Lookup.Base != "" && object.Lookup.Token != "" {
				entity, err := CreateAnimatedComposite(int(worldX), int(worldY), object.Lookup, d2resource.PaletteUnits)
				if err != nil {
					panic(err)
				}
				entity.SetMode(object.Lookup.Mode, object.Lookup.Class, 0)
				entities = append(entities, entity)
			}
		}
	}

	return entities
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

	random := tileSeed % uint64(weightSum)

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

func (mr *MapRegion) getTiles(style, sequence, tileType int32) []d2dt1.Tile {
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

func (mr *MapRegion) advance(elapsed float64) {
	frameLength := 0.1

	mr.lastFrameTime += elapsed
	framesAdvanced := int(mr.lastFrameTime / frameLength)
	mr.lastFrameTime -= float64(framesAdvanced) * frameLength

	mr.currentFrame += framesAdvanced
	if mr.currentFrame > 9 {
		mr.currentFrame = 0
	}
}

func (mr *MapRegion) getTileWorldPosition(tileX, tileY int) (float64, float64) {
	return float64(tileX + mr.tileRect.Left), float64(tileY + mr.tileRect.Top)
}

func (mr *MapRegion) renderPass1(viewport *Viewport, target d2render.Surface) {
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

func (mr *MapRegion) renderPass2(entities MapEntitiesSearcher, viewport *Viewport, target d2render.Surface) {
	tileEntities := entities.SearchByRect(mr.tileRect)
	nextIndex := 0

	for tileY := range mr.ds1.Tiles {
		for tileX, tile := range mr.ds1.Tiles[tileY] {
			worldX, worldY := mr.getTileWorldPosition(tileX, tileY)
			if viewport.IsTileVisible(worldX, worldY) {
				viewport.PushTranslationWorld(worldX, worldY)
				mr.renderTilePass2(tile, viewport, target)

				for nextIndex < len(tileEntities) {
					nextX, nextY := tileEntities[nextIndex].GetPosition()
					if nextX == worldX && nextY == worldY {
						target.PushTranslation(viewport.GetTranslationScreen())
						tileEntities[nextIndex].Render(target)
						target.Pop()
						nextIndex++
					} else if (nextY == worldY && nextX < worldX) || nextY < worldY {
						nextIndex++
					} else {
						break
					}
				}

				viewport.PopTranslation()
			}
		}
	}
}

func (mr *MapRegion) renderPass3(viewport *Viewport, target d2render.Surface) {
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

func (mr *MapRegion) renderTilePass1(tile d2ds1.TileRecord, viewport *Viewport, target d2render.Surface) {
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

func (mr *MapRegion) renderTilePass2(tile d2ds1.TileRecord, viewport *Viewport, target d2render.Surface) {
	for _, wall := range tile.Walls {
		if !wall.Hidden && wall.Type.UpperWall() {
			mr.renderWall(wall, viewport, target)
		}
	}
}

func (mr *MapRegion) renderTilePass3(tile d2ds1.TileRecord, viewport *Viewport, target d2render.Surface) {
	for _, wall := range tile.Walls {
		if wall.Type == d2enum.Roof {
			mr.renderWall(wall, viewport, target)
		}
	}
}

func (mr *MapRegion) renderFloor(tile d2ds1.FloorShadowRecord, viewport *Viewport, target d2render.Surface) {
	var img d2render.Surface
	if !tile.Animated {
		img = mr.getImageCacheRecord(tile.Style, tile.Sequence, 0, tile.RandomIndex)
	} else {
		img = mr.getImageCacheRecord(tile.Style, tile.Sequence, 0, byte(mr.currentFrame))
	}
	if img == nil {
		log.Printf("Render called on uncached floor {%v,%v}", tile.Style, tile.Sequence)
		return
	}

	viewport.PushTranslationOrtho(-80, float64(tile.YAdjust))
	defer viewport.PopTranslation()

	target.PushTranslation(viewport.GetTranslationScreen())
	defer target.Pop()

	target.Render(img)
}

func (mr *MapRegion) renderWall(tile d2ds1.WallRecord, viewport *Viewport, target d2render.Surface) {
	img := mr.getImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tile.RandomIndex)
	if img == nil {
		log.Printf("Render called on uncached wall {%v,%v,%v}", tile.Style, tile.Sequence, tile.Type)
		return
	}

	viewport.PushTranslationOrtho(-80, float64(tile.YAdjust))
	defer viewport.PopTranslation()

	target.PushTranslation(viewport.GetTranslationScreen())
	defer target.Pop()

	target.Render(img)
}

func (mr *MapRegion) renderShadow(tile d2ds1.FloorShadowRecord, viewport *Viewport, target d2render.Surface) {
	img := mr.getImageCacheRecord(tile.Style, tile.Sequence, 13, tile.RandomIndex)
	if img == nil {
		log.Printf("Render called on uncached shadow {%v,%v}", tile.Style, tile.Sequence)
		return
	}

	viewport.PushTranslationOrtho(-80, float64(tile.YAdjust))
	defer viewport.PopTranslation()

	target.PushTranslation(viewport.GetTranslationScreen())
	target.PushColor(color.RGBA{R: 255, G: 255, B: 255, A: 160})
	defer target.PopN(2)

	target.Render(img)
}

func (mr *MapRegion) renderDebug(debugVisLevel int, viewport *Viewport, target d2render.Surface) {
	for tileY := range mr.ds1.Tiles {
		for tileX := range mr.ds1.Tiles[tileY] {
			worldX, worldY := mr.getTileWorldPosition(tileX, tileY)
			if viewport.IsTileVisible(worldX, worldY) {
				mr.renderTileDebug(int(worldX), int(worldY), debugVisLevel, viewport, target)
			}
		}
	}
}

func (mr *MapRegion) renderTileDebug(x, y int, debugVisLevel int, viewport *Viewport, target d2render.Surface) {
	ax := x - mr.tileRect.Left
	ay := y - mr.tileRect.Top

	if debugVisLevel > 0 {
		if ay < 0 || ax < 0 || ay >= len(mr.ds1.Tiles) || x >= len(mr.ds1.Tiles[ay]) {
			return
		}
		subTileColor := color.RGBA{R: 80, G: 80, B: 255, A: 50}
		tileColor := color.RGBA{R: 255, G: 255, B: 255, A: 100}
		tileCollisionColor := color.RGBA{R: 128, G: 0, B: 0, A: 100}

		screenX1, screenY1 := viewport.WorldToScreen(float64(x), float64(y))
		screenX2, screenY2 := viewport.WorldToScreen(float64(x+1), float64(y))
		screenX3, screenY3 := viewport.WorldToScreen(float64(x), float64(y+1))

		target.PushTranslation(screenX1, screenY1)
		defer target.Pop()

		target.DrawLine(screenX2-screenX1, screenY2-screenY1, tileColor)
		target.DrawLine(screenX3-screenX1, screenY3-screenY1, tileColor)
		target.PushTranslation(-10, 10)
		target.DrawText("%v, %v", x, y)
		target.Pop()

		if debugVisLevel > 1 {
			for i := 1; i <= 4; i++ {
				x2 := i * 16
				y2 := i * 8

				target.PushTranslation(-x2, y2)
				target.DrawLine(80, 40, subTileColor)
				target.Pop()

				target.PushTranslation(x2, y2)
				target.DrawLine(-80, 40, subTileColor)
				target.Pop()
			}

			tile := mr.ds1.Tiles[ay][ax]

			for i, floor := range tile.Floors {
				target.PushTranslation(-20, 10+(i+1)*14)
				target.DrawText("f: %v-%v", floor.Style, floor.Sequence)
				target.Pop()
			}

			for yy := 0; yy < 5; yy++ {
				for xx := 0; xx < 5; xx++ {
					isoX := (xx - yy) * 16
					isoY := (xx + yy) * 8
					target.PushTranslation(isoX-3, isoY+4)
					var walkableArea = mr.walkableArea[yy+(ay*5)][xx+(ax*5)]
					if !walkableArea.Walkable {
						target.DrawRect(5, 5, tileCollisionColor)
					}
					target.Pop()
				}
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

func (mr *MapRegion) getImageCacheRecord(style, sequence byte, tileType d2enum.TileType, randomIndex byte) d2render.Surface {
	lookupIndex := uint32(style)<<24 | uint32(sequence)<<16 | uint32(tileType)<<8 | uint32(randomIndex)
	return mr.imageCacheRecords[lookupIndex]
}

func (mr *MapRegion) setImageCacheRecord(style, sequence byte, tileType d2enum.TileType, randomIndex byte, image d2render.Surface) {
	lookupIndex := uint32(style)<<24 | uint32(sequence)<<16 | uint32(tileType)<<8 | uint32(randomIndex)
	mr.imageCacheRecords[lookupIndex] = image
}

func (mr *MapRegion) generateFloorCache(tile *d2ds1.FloorShadowRecord, tileX, tileY int) {
	tileOptions := mr.getTiles(int32(tile.Style), int32(tile.Sequence), 0)
	var tileData []*d2dt1.Tile
	var tileIndex byte

	if tileOptions == nil {
		log.Printf("Could not locate tile Style:%d, Seq: %d, Type: %d\n", tile.Style, tile.Sequence, 0)
		tileData = append(tileData, &d2dt1.Tile{})
		tileData[0].Width = 10
		tileData[0].Height = 10
	} else {
		if !tileOptions[0].MaterialFlags.Lava {
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
		if !tileData[i].MaterialFlags.Lava {
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
			tileYMinimum = d2common.MinInt32(tileYMinimum, int32(block.Y))
		}
		tileYOffset := d2common.AbsInt32(tileYMinimum)
		tileHeight := d2common.AbsInt32(tileData[i].Height)
		image, _ := d2render.NewSurface(int(tileData[i].Width), int(tileHeight), d2render.FilterNearest)
		pixels := make([]byte, 4*tileData[i].Width*tileHeight)
		mr.decodeTileGfxData(tileData[i].Blocks, &pixels, tileYOffset, tileData[i].Width)
		image.ReplacePixels(pixels)
		mr.setImageCacheRecord(tile.Style, tile.Sequence, 0, tileIndex, image)
	}
}

func (mr *MapRegion) generateShadowCache(tile *d2ds1.FloorShadowRecord, tileX, tileY int) {
	tileOptions := mr.getTiles(int32(tile.Style), int32(tile.Sequence), 13)
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
		tileMinY = d2common.MinInt32(tileMinY, int32(block.Y))
		tileMaxY = d2common.MaxInt32(tileMaxY, int32(block.Y+32))
	}
	tileYOffset := -tileMinY
	tileHeight := int(tileMaxY - tileMinY)
	tile.YAdjust = int(tileMinY + 80)

	cachedImage := mr.getImageCacheRecord(tile.Style, tile.Sequence, 13, tileIndex)
	if cachedImage != nil {
		return
	}

	image, _ := d2render.NewSurface(int(tileData.Width), tileHeight, d2render.FilterNearest)
	pixels := make([]byte, 4*tileData.Width*int32(tileHeight))
	mr.decodeTileGfxData(tileData.Blocks, &pixels, tileYOffset, tileData.Width)
	image.ReplacePixels(pixels)
	mr.setImageCacheRecord(tile.Style, tile.Sequence, 13, tileIndex, image)
}

func (mr *MapRegion) generateWallCache(tile *d2ds1.WallRecord, tileX, tileY int) {
	tileOptions := mr.getTiles(int32(tile.Style), int32(tile.Sequence), int32(tile.Type))
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
		newTileOptions := mr.getTiles(int32(tile.Style), int32(tile.Sequence), int32(4))
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
		tileMinY = d2common.MinInt32(tileMinY, int32(block.Y))
		tileMaxY = d2common.MaxInt32(tileMaxY, int32(block.Y+32))
	}

	realHeight := d2common.MaxInt32(d2common.AbsInt32(tileData.Height), tileMaxY-tileMinY)
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

	image, _ := d2render.NewSurface(160, int(realHeight), d2render.FilterNearest)
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

func loadPaletteForAct(levelType d2enum.RegionIdType) (*d2dat.DATPalette, error) {
	var palettePath string
	switch levelType {
	case d2enum.RegionAct1Town, d2enum.RegionAct1Wilderness, d2enum.RegionAct1Cave, d2enum.RegionAct1Crypt,
		d2enum.RegionAct1Monestary, d2enum.RegionAct1Courtyard, d2enum.RegionAct1Barracks,
		d2enum.RegionAct1Jail, d2enum.RegionAct1Cathedral, d2enum.RegionAct1Catacombs, d2enum.RegionAct1Tristram:
		palettePath = d2resource.PaletteAct1
		break
	case d2enum.RegionAct2Town, d2enum.RegionAct2Sewer, d2enum.RegionAct2Harem, d2enum.RegionAct2Basement,
		d2enum.RegionAct2Desert, d2enum.RegionAct2Tomb, d2enum.RegionAct2Lair, d2enum.RegionAct2Arcane:
		palettePath = d2resource.PaletteAct2
		break
	case d2enum.RegionAct3Town, d2enum.RegionAct3Jungle, d2enum.RegionAct3Kurast, d2enum.RegionAct3Spider,
		d2enum.RegionAct3Dungeon, d2enum.RegionAct3Sewer:
		palettePath = d2resource.PaletteAct3
		break
	case d2enum.RegionAct4Town, d2enum.RegionAct4Mesa, d2enum.RegionAct4Lava, d2enum.RegionAct5Lava:
		palettePath = d2resource.PaletteAct4
		break
	case d2enum.RegonAct5Town, d2enum.RegionAct5Siege, d2enum.RegionAct5Barricade, d2enum.RegionAct5Temple,
		d2enum.RegionAct5IceCaves, d2enum.RegionAct5Baal:
		palettePath = d2resource.PaletteAct5
		break
	default:
		return nil, errors.New("failed to find palette for region")
	}

	return d2asset.LoadPalette(palettePath)
}
