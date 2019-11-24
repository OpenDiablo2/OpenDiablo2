package d2mapengine

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"

	"github.com/OpenDiablo2/D2Shared/d2data/d2dt1"

	"github.com/OpenDiablo2/D2Shared/d2data/d2ds1"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/D2Shared/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2corehelper"

	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2render"

	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"

	"github.com/hajimehoshi/ebiten"
)

type Region struct {
	RegionPath        string
	LevelType         d2datadict.LevelTypeRecord
	LevelPreset       d2datadict.LevelPresetRecord
	TileWidth         int32
	TileHeight        int32
	Tiles             []d2dt1.Tile
	DS1               d2ds1.DS1
	Palette           d2datadict.PaletteRec
	AnimationEntities []d2render.AnimatedEntity
	NPCs              []*d2core.NPC
	StartX            float64
	StartY            float64
	imageCacheRecords map[uint32]*ebiten.Image
	seed              int64
}

func LoadRegion(seed int64, levelType d2enum.RegionIdType, levelPreset int, fileProvider d2interface.FileProvider, fileIndex int) *Region {
	result := &Region{
		LevelType:         d2datadict.LevelTypes[levelType],
		LevelPreset:       d2datadict.LevelPresets[levelPreset],
		Tiles:             make([]d2dt1.Tile, 0),
		imageCacheRecords: map[uint32]*ebiten.Image{},
		seed:              seed,
	}
	result.Palette = d2datadict.Palettes[d2enum.PaletteType("act"+strconv.Itoa(int(result.LevelType.Act)))]
	// Temp hack
	if levelType == d2enum.RegionAct5Lava {
		result.Palette = d2datadict.Palettes[d2enum.PaletteType("act4")]
	}
	//bm := result.levelPreset.Dt1Mask
	for _, levelTypeDt1 := range result.LevelType.Files {
		/*
			if bm&1 == 0 {
				bm >>= 1
				continue
			}
			bm >>= 1
		*/
		if len(levelTypeDt1) == 0 || levelTypeDt1 == "" || levelTypeDt1 == "0" {
			continue
		}
		dt1 := d2dt1.LoadDT1("/data/global/tiles/"+levelTypeDt1, fileProvider)
		result.Tiles = append(result.Tiles, dt1.Tiles...)
	}
	levelFilesToPick := make([]string, 0)
	for _, fileRecord := range result.LevelPreset.Files {
		if len(fileRecord) == 0 || fileRecord == "" || fileRecord == "0" {
			continue
		}
		levelFilesToPick = append(levelFilesToPick, fileRecord)
	}
	levelIndex := int(math.Round(float64(len(levelFilesToPick)-1) * rand.Float64()))
	if fileIndex >= 0 && fileIndex < len(levelFilesToPick) {
		levelIndex = fileIndex
	}
	levelFile := levelFilesToPick[levelIndex]

	result.RegionPath = levelFile
	result.DS1 = d2ds1.LoadDS1("/data/global/tiles/"+levelFile, fileProvider)
	result.TileWidth = result.DS1.Width
	result.TileHeight = result.DS1.Height
	result.loadObjects(fileProvider)
	result.loadSpecials()
	return result
}

func (v *Region) loadSpecials() {
	for y := range v.DS1.Tiles {
		for x := range v.DS1.Tiles[y] {
			for _, wall := range v.DS1.Tiles[y][x].Walls {
				if wall.Type != 10 {
					continue
				}
				if wall.Style == 30 && wall.Sequence == 0 {
					v.StartX = float64(x) + 0.5
					v.StartY = float64(y) + 0.5
					log.Printf("Starting location: %d, %d", x, y)
				}
			}
		}
	}
}

func (v *Region) loadObjects(fileProvider d2interface.FileProvider) {
	v.AnimationEntities = make([]d2render.AnimatedEntity, 0)
	v.NPCs = make([]*d2core.NPC, 0)
	for _, object := range v.DS1.Objects {
		switch object.Lookup.Type {
		case d2datadict.ObjectTypeCharacter:
			// Temp code, maybe..
			if object.Lookup.Base == "" || object.Lookup.Token == "" || object.Lookup.TR == "" {
				continue
			}
			npc := d2core.CreateNPC(object.X, object.Y, object.Lookup, fileProvider, 1)
			npc.SetPaths(object.Paths)
			v.NPCs = append(v.NPCs, npc)
		case d2datadict.ObjectTypeItem:
			if object.ObjectInfo == nil || !object.ObjectInfo.Draw || object.Lookup.Base == "" || object.Lookup.Token == "" {
				continue
			}
			entity := d2render.CreateAnimatedEntity(object.X, object.Y, object.Lookup, fileProvider, d2enum.Units)
			entity.SetMode(object.Lookup.Mode, object.Lookup.Class, 0)
			v.AnimationEntities = append(v.AnimationEntities, entity)
		}
	}
}

func (v *Region) RenderTile(offsetX, offsetY, tileX, tileY int, layerType d2enum.RegionLayerType, layerIndex int, target *ebiten.Image) {
	offsetX -= 80
	switch layerType {
	case d2enum.RegionLayerTypeFloors:
		v.renderFloor(v.DS1.Tiles[tileY][tileX].Floors[layerIndex], offsetX, offsetY, target, tileX, tileY)
	case d2enum.RegionLayerTypeWalls:
		v.renderWall(v.DS1.Tiles[tileY][tileX].Walls[layerIndex], offsetX, offsetY, target, tileX, tileY)
	case d2enum.RegionLayerTypeShadows:
		v.renderShadow(v.DS1.Tiles[tileY][tileX].Shadows[layerIndex], offsetX, offsetY, target, tileX, tileY)
	}
}

func (v *Region) getRandomTile(tiles []d2dt1.Tile, x, y int, seed int64) byte {
	/* Walker's Alias Method for weighted random selection
	 * with xorshifting for random numbers */

	var tileSeed uint64
	tileSeed = uint64(seed) + uint64(x)
	tileSeed *= uint64(y) + uint64(v.LevelType.Id)

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

func (v *Region) getTiles(style, sequence, tileType int32, x, y int, seed int64) []d2dt1.Tile {
	var tiles []d2dt1.Tile
	for _, tile := range v.Tiles {
		if tile.Style != style || tile.Sequence != sequence || tile.Type != tileType {
			continue
		}
		tiles = append(tiles, tile)
	}
	if len(tiles) == 0 {
		log.Printf("Unknown tile ID [%d %d %d]\n", style, sequence, tileType)
		return nil
	}
	//	index := v.getRandomTile(tiles, x, y, seed)
	//	return &tiles[index], index
	return tiles
}

func (v *Region) renderFloor(tile d2ds1.FloorShadowRecord, offsetX, offsetY int, target *ebiten.Image, tileX, tileY int) {
	img := v.GetImageCacheRecord(tile.Style, tile.Sequence, 0, tile.RandomIndex)
	if img == nil {
		log.Printf("Render called on uncached floor {%v,%v}", tile.Style, tile.Sequence)
		return
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(offsetX), float64(offsetY))
	_ = target.DrawImage(img, opts)
	return
}

func (v *Region) renderWall(tile d2ds1.WallRecord, offsetX, offsetY int, target *ebiten.Image, tileX, tileY int) {
	img := v.GetImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tile.RandomIndex)
	if img == nil {
		log.Printf("Render called on uncached wall {%v,%v,%v}", tile.Style, tile.Sequence, tile.Type)
		return
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(offsetX), float64(offsetY+tile.YAdjust))
	target.DrawImage(img, opts)
}

func (v *Region) renderShadow(tile d2ds1.FloorShadowRecord, offsetX, offsetY int, target *ebiten.Image, tileX, tileY int) {
	img := v.GetImageCacheRecord(tile.Style, tile.Sequence, 13, tile.RandomIndex)
	if img == nil {
		log.Printf("Render called on uncached shadow {%v,%v}", tile.Style, tile.Sequence)
		return
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(offsetX), float64(offsetY+tile.YAdjust))
	opts.ColorM = d2corehelper.ColorToColorM(color.RGBA{255, 255, 255, 160})
	target.DrawImage(img, opts)
}

func (v *Region) decodeTileGfxData(blocks []d2dt1.Block, pixels *[]byte, tileYOffset int32, tileWidth int32) {
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
						pixelColor := v.Palette.Colors[colorIndex]
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
						pixelColor := v.Palette.Colors[colorIndex]
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

func (v *Region) generateFloorCache(tile *d2ds1.FloorShadowRecord, tileX, tileY int) {
	tileOptions := v.getTiles(int32(tile.Style), int32(tile.Sequence), 0, tileX, tileY, v.seed)
	tileIndex := v.getRandomTile(tileOptions, tileX, tileY, v.seed)
	tileData := &tileOptions[tileIndex]

	if tileData == nil {
		log.Printf("Could not locate tile Style:%d, Seq: %d, Type: %d\n", tile.Style, tile.Sequence, 0)
		tileData = &d2dt1.Tile{}
		tileData.Width = 10
		tileData.Height = 10
	}
	tile.RandomIndex = tileIndex
	cachedImage := v.GetImageCacheRecord(tile.Style, tile.Sequence, 0, tileIndex)
	if cachedImage != nil {
		return
	}
	tileYMinimum := int32(0)
	for _, block := range tileData.Blocks {
		tileYMinimum = d2helper.MinInt32(tileYMinimum, int32(block.Y))
	}
	tileYOffset := d2helper.AbsInt32(tileYMinimum)
	tileHeight := d2helper.AbsInt32(tileData.Height)
	image, _ := ebiten.NewImage(int(tileData.Width), int(tileHeight), ebiten.FilterNearest)
	pixels := make([]byte, 4*tileData.Width*tileHeight)
	v.decodeTileGfxData(tileData.Blocks, &pixels, tileYOffset, tileData.Width)
	image.ReplacePixels(pixels)
	v.SetImageCacheRecord(tile.Style, tile.Sequence, 0, tileIndex, image)
}

func (v *Region) generateShadowCache(tile *d2ds1.FloorShadowRecord, tileX, tileY int) {
	tileOptions := v.getTiles(int32(tile.Style), int32(tile.Sequence), 13, tileX, tileY, v.seed)
	tileIndex := v.getRandomTile(tileOptions, tileX, tileY, v.seed)
	tileData := &tileOptions[tileIndex]

	if tileData == nil {
		return
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

	cachedImage := v.GetImageCacheRecord(tile.Style, tile.Sequence, 13, tileIndex)
	if cachedImage != nil {
		return
	}

	image, _ := ebiten.NewImage(int(tileData.Width), int(tileHeight), ebiten.FilterNearest)
	pixels := make([]byte, 4*tileData.Width*int32(tileHeight))
	v.decodeTileGfxData(tileData.Blocks, &pixels, tileYOffset, tileData.Width)
	image.ReplacePixels(pixels)
	v.SetImageCacheRecord(tile.Style, tile.Sequence, 13, tileIndex, image)
}

func (v *Region) generateWallCache(tile *d2ds1.WallRecord, tileX, tileY int) {
	tileOptions := v.getTiles(int32(tile.Style), int32(tile.Sequence), int32(tile.Type), tileX, tileY, v.seed)
	tileIndex := v.getRandomTile(tileOptions, tileX, tileY, v.seed)
	tileData := &tileOptions[tileIndex]
	if tileData == nil {
		return
	}

	tile.RandomIndex = tileIndex
	var newTileData *d2dt1.Tile = nil

	if tile.Type == 3 {
		newTileOptions := v.getTiles(int32(tile.Style), int32(tile.Sequence), int32(4), tileX, tileY, v.seed)
		newTileIndex := v.getRandomTile(newTileOptions, tileX, tileY, v.seed)
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

	cachedImage := v.GetImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tileIndex)
	if cachedImage != nil {
		return
	}

	if realHeight == 0 {
		log.Printf("Invalid 0 height for wall tile")
		return
	}

	image, _ := ebiten.NewImage(160, int(realHeight), ebiten.FilterNearest)
	pixels := make([]byte, 4*160*realHeight)
	v.decodeTileGfxData(tileData.Blocks, &pixels, tileYOffset, 160)

	if newTileData != nil {
		v.decodeTileGfxData(newTileData.Blocks, &pixels, tileYOffset, 160)
	}

	if err := image.ReplacePixels(pixels); err != nil {
		log.Panicf(err.Error())
	}

	v.SetImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tileIndex, image)
}

func (v *Region) GetImageCacheRecord(style, sequence byte, tileType d2enum.TileType, randomIndex byte) *ebiten.Image {
	lookupIndex := uint32(style)<<24 | uint32(sequence)<<16 | uint32(tileType)<<8 | uint32(randomIndex)
	return v.imageCacheRecords[lookupIndex]
}

func (v *Region) SetImageCacheRecord(style, sequence byte, tileType d2enum.TileType, randomIndex byte, image *ebiten.Image) {
	lookupIndex := uint32(style)<<24 | uint32(sequence)<<16 | uint32(tileType)<<8 | uint32(randomIndex)
	v.imageCacheRecords[lookupIndex] = image
}
