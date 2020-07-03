package d2maprenderer

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

func (mr *MapRenderer) generateTileCache() {
	mr.palette, _ = loadPaletteForAct(d2enum.RegionIdType(mr.mapEngine.LevelType().ID))
	mapEngineSize := mr.mapEngine.Size()

	for idx, tile := range *mr.mapEngine.Tiles() {
		tileX := idx % mapEngineSize.Width
		tileY := (idx - tileX) / mapEngineSize.Width
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

func (mr *MapRenderer) generateFloorCache(tile *d2ds1.FloorShadowRecord, tileX, tileY int) {
	tileOptions := mr.mapEngine.GetTiles(int32(tile.Style), int32(tile.Sequence), 0)
	var tileData []*d2dt1.Tile
	var tileIndex byte

	if tileOptions == nil {
		log.Printf("Could not locate tile Style:%d, Seq: %d, Type: %d\n", tile.Style, tile.Sequence, 0)
		tileData = append(tileData, &d2dt1.Tile{})
		tileData[0].Width = 10
		tileData[0].Height = 10
	} else {
		if !tileOptions[0].MaterialFlags.Lava {
			tileIndex = mr.getRandomTile(tileOptions, tileX, tileY, mr.mapEngine.Seed())
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
		image, _ := mr.renderer.NewSurface(int(tileData[i].Width), int(tileHeight), d2interface.FilterNearest)
		pixels := make([]byte, 4*tileData[i].Width*tileHeight)
		mr.decodeTileGfxData(tileData[i].Blocks, &pixels, tileYOffset, tileData[i].Width)
		_ = image.ReplacePixels(pixels)
		mr.setImageCacheRecord(tile.Style, tile.Sequence, 0, tileIndex, image)
	}
}

func (mr *MapRenderer) generateShadowCache(tile *d2ds1.FloorShadowRecord, tileX, tileY int) {
	tileOptions := mr.mapEngine.GetTiles(int32(tile.Style), int32(tile.Sequence), 13)
	var tileIndex byte
	var tileData *d2dt1.Tile
	if tileOptions == nil {
		return
	} else {
		tileIndex = mr.getRandomTile(tileOptions, tileX, tileY, mr.mapEngine.Seed())
		tileData = &tileOptions[tileIndex]
	}

	if tileData.Width == 0 || tileData.Height == 0 {
		return
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

	image, _ := mr.renderer.NewSurface(int(tileData.Width), tileHeight, d2interface.FilterNearest)
	pixels := make([]byte, 4*tileData.Width*int32(tileHeight))
	mr.decodeTileGfxData(tileData.Blocks, &pixels, tileYOffset, tileData.Width)
	_ = image.ReplacePixels(pixels)
	mr.setImageCacheRecord(tile.Style, tile.Sequence, 13, tileIndex, image)
}

func (mr *MapRenderer) generateWallCache(tile *d2ds1.WallRecord, tileX, tileY int) {
	tileOptions := mr.mapEngine.GetTiles(int32(tile.Style), int32(tile.Sequence), int32(tile.Type))
	var tileIndex byte
	var tileData *d2dt1.Tile
	if tileOptions == nil {
		return
	} else {
		tileIndex = mr.getRandomTile(tileOptions, tileX, tileY, mr.mapEngine.Seed())
		tileData = &tileOptions[tileIndex]
	}

	tile.RandomIndex = tileIndex
	var newTileData *d2dt1.Tile = nil

	if tile.Type == 3 {
		newTileOptions := mr.mapEngine.GetTiles(int32(tile.Style), int32(tile.Sequence), int32(4))
		newTileIndex := mr.getRandomTile(newTileOptions, tileX, tileY, mr.mapEngine.Seed())
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

	image, _ := mr.renderer.NewSurface(160, int(realHeight), d2interface.FilterNearest)
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

func (mr *MapRenderer) getRandomTile(tiles []d2dt1.Tile, x, y int, seed int64) byte {
	/* Walker's Alias Method for weighted random selection
	 * with xorshifting for random numbers */

	var tileSeed uint64
	tileSeed = uint64(seed) + uint64(x)
	tileSeed *= uint64(y) + uint64(mr.mapEngine.LevelType().ID)

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
