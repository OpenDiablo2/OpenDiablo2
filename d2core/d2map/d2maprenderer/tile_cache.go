package d2maprenderer

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

func (mr *MapRenderer) generateTileCache() {
	mr.palette, _ = loadPaletteForAct(d2enum.RegionIdType(mr.mapEngine.LevelType().ID))

	tiles := *mr.mapEngine.Tiles()
	for idx := range tiles {
		tile := &tiles[idx]

		for i := range tile.Components.Floors {
			if !tile.Components.Floors[i].Hidden && tile.Components.Floors[i].Prop1 != 0 {
				mr.generateFloorCache(&tile.Components.Floors[i])
			}
		}

		for i := range tile.Components.Shadows {
			if !tile.Components.Shadows[i].Hidden && tile.Components.Shadows[i].Prop1 != 0 {
				mr.generateShadowCache(&tile.Components.Shadows[i])
			}
		}

		for i := range tile.Components.Walls {
			if !tile.Components.Walls[i].Hidden && tile.Components.Walls[i].Prop1 != 0 {
				mr.generateWallCache(&tile.Components.Walls[i])
			}
		}
	}
}

func (mr *MapRenderer) generateFloorCache(tile *d2ds1.FloorShadowRecord) {
	tileOptions := mr.mapEngine.GetTiles(int(tile.Style), int(tile.Sequence), 0)

	var tileData []*d2dt1.Tile

	if tileOptions == nil {
		log.Printf("Could not locate tile Style:%d, Seq: %d, Type: %d\n", tile.Style, tile.Sequence, 0)

		tileData = append(tileData, &d2dt1.Tile{})
		tileData[0].Width = 10
		tileData[0].Height = 10
	} else {
		if !tileOptions[0].MaterialFlags.Lava {
			tileData = append(tileData, &tileOptions[tile.RandomIndex])
		} else {
			tile.Animated = true
			for i := range tileOptions {
				tileData = append(tileData, &tileOptions[i])
			}
		}
	}

	var tileIndex byte

	for i := range tileData {
		if tileData[i].MaterialFlags.Lava {
			tileIndex = byte(tileData[i].RarityFrameIndex)
		} else {
			tileIndex = tile.RandomIndex
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
		image, _ := mr.renderer.NewSurface(int(tileData[i].Width), int(tileHeight), d2enum.FilterNearest)
		indexData := make([]byte, tileData[i].Width*tileHeight)
		d2dt1.DecodeTileGfxData(tileData[i].Blocks, &indexData, tileYOffset, tileData[i].Width)
		pixels := d2asset.ImgIndexToRGBA(indexData, mr.palette)

		_ = image.ReplacePixels(pixels)
		mr.setImageCacheRecord(tile.Style, tile.Sequence, 0, tileIndex, image)
	}
}

func (mr *MapRenderer) generateShadowCache(tile *d2ds1.FloorShadowRecord) {
	tileOptions := mr.mapEngine.GetTiles(int(tile.Style), int(tile.Sequence), 13)

	var tileData *d2dt1.Tile

	if tileOptions == nil {
		return
	}

	tileData = &tileOptions[tile.RandomIndex]

	if tileData.Width == 0 || tileData.Height == 0 {
		return
	}

	tileMinY := int32(0)
	tileMaxY := int32(0)

	for _, block := range tileData.Blocks {
		tileMinY = d2common.MinInt32(tileMinY, int32(block.Y))
		tileMaxY = d2common.MaxInt32(tileMaxY, int32(block.Y+32))
	}

	tileYOffset := -tileMinY
	tileHeight := int(tileMaxY - tileMinY)
	tile.YAdjust = int(tileMinY + 80)

	cachedImage := mr.getImageCacheRecord(tile.Style, tile.Sequence, 13, tile.RandomIndex)
	if cachedImage != nil {
		return
	}

	image, _ := mr.renderer.NewSurface(int(tileData.Width), tileHeight, d2enum.FilterNearest)
	indexData := make([]byte, tileData.Width*int32(tileHeight))
	d2dt1.DecodeTileGfxData(tileData.Blocks, &indexData, tileYOffset, tileData.Width)
	pixels := d2asset.ImgIndexToRGBA(indexData, mr.palette)
	_ = image.ReplacePixels(pixels)
	mr.setImageCacheRecord(tile.Style, tile.Sequence, 13, tile.RandomIndex, image)
}

func (mr *MapRenderer) generateWallCache(tile *d2ds1.WallRecord) {
	tileOptions := mr.mapEngine.GetTiles(int(tile.Style), int(tile.Sequence), int(tile.Type))

	var tileData *d2dt1.Tile

	if tileOptions == nil {
		return
	}

	tileData = &tileOptions[tile.RandomIndex]

	var newTileData *d2dt1.Tile = nil

	if tile.Type == 3 {
		newTileOptions := mr.mapEngine.GetTiles(int(tile.Style), int(tile.Sequence), int(4))
		newTileData = &newTileOptions[tile.RandomIndex]
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

	if tile.Type == 15 {
		tile.YAdjust = -int(tileData.RoofHeight)
	} else {
		tile.YAdjust = int(tileMinY) + 80
	}

	cachedImage := mr.getImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tile.RandomIndex)
	if cachedImage != nil {
		return
	}

	if realHeight == 0 {
		log.Printf("Invalid 0 height for wall tile")
		return
	}

	image, _ := mr.renderer.NewSurface(160, int(realHeight), d2enum.FilterNearest)
	indexData := make([]byte, 160*realHeight)

	d2dt1.DecodeTileGfxData(tileData.Blocks, &indexData, tileYOffset, 160)

	if newTileData != nil {
		d2dt1.DecodeTileGfxData(newTileData.Blocks, &indexData, tileYOffset, 160)
	}

	pixels := d2asset.ImgIndexToRGBA(indexData, mr.palette)

	if err := image.ReplacePixels(pixels); err != nil {
		log.Panicf(err.Error())
	}

	mr.setImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tile.RandomIndex, image)
}
