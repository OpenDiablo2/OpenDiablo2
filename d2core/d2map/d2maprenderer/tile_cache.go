package d2maprenderer

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const (
	tileMinHeight          int16 = 32
	shadowAdjustY          int32 = 80
	defaultFloorTileWidth        = 10
	defaultFloorTileHeight       = 10
)

const (
	blockOffsetY      = 32
	tileSurfaceWidth  = 160
	tileSurfaceHeight = 80
)

func (mr *MapRenderer) generateTileCache() {
	var err error
	mr.palette, err = mr.loadPaletteForAct(d2enum.RegionIdType(mr.mapEngine.LevelType().ID))

	if err != nil {
		mr.Error(err.Error())
	}

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
		mr.Errorf("Could not locate tile Style:%d, Seq: %d, Type: %d", tile.Style, tile.Sequence, 0)

		tileData = append(tileData, &d2dt1.Tile{})
		tileData[0].Width = defaultFloorTileWidth
		tileData[0].Height = defaultFloorTileHeight
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
			tileYMinimum = d2math.MinInt32(tileYMinimum, int32(block.Y))
		}

		tileYOffset := d2math.AbsInt32(tileYMinimum)
		tileHeight := d2math.AbsInt32(tileData[i].Height)
		image := mr.renderer.NewSurface(int(tileData[i].Width), int(tileHeight))

		indexData := make([]byte, tileData[i].Width*tileHeight)
		d2dt1.DecodeTileGfxData(tileData[i].Blocks, &indexData, tileYOffset, tileData[i].Width)
		pixels := d2util.ImgIndexToRGBA(indexData, mr.palette)

		image.ReplacePixels(pixels)

		mr.setImageCacheRecord(tile.Style, tile.Sequence, 0, tileIndex, image)
	}
}

func (mr *MapRenderer) generateShadowCache(tile *d2ds1.FloorShadowRecord) {
	tileOptions := mr.mapEngine.GetTiles(int(tile.Style), int(tile.Sequence), d2enum.TileShadow)

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
		tileMinY = d2math.MinInt32(tileMinY, int32(block.Y))
		tileMaxY = d2math.MaxInt32(tileMaxY, int32(block.Y+tileMinHeight))
	}

	tileYOffset := -tileMinY
	tileHeight := int(tileMaxY - tileMinY)
	tile.YAdjust = int(tileMinY + shadowAdjustY)

	cachedImage := mr.getImageCacheRecord(tile.Style, tile.Sequence, d2enum.TileShadow, tile.RandomIndex)
	if cachedImage != nil {
		return
	}

	image := mr.renderer.NewSurface(int(tileData.Width), tileHeight)

	indexData := make([]byte, tileData.Width*int32(tileHeight))
	d2dt1.DecodeTileGfxData(tileData.Blocks, &indexData, tileYOffset, tileData.Width)
	pixels := d2util.ImgIndexToRGBA(indexData, mr.palette)

	image.ReplacePixels(pixels)

	mr.setImageCacheRecord(tile.Style, tile.Sequence, d2enum.TileShadow, tile.RandomIndex, image)
}

func (mr *MapRenderer) generateWallCache(tile *d2ds1.WallRecord) {
	tileOptions := mr.mapEngine.GetTiles(int(tile.Style), int(tile.Sequence), tile.Type)

	var tileData *d2dt1.Tile

	if tileOptions == nil {
		return
	}

	tileData = &tileOptions[tile.RandomIndex]

	var newTileData *d2dt1.Tile = nil

	if tile.Type == d2enum.TileRightPartOfNorthCornerWall {
		newTileOptions := mr.mapEngine.GetTiles(
			int(tile.Style), int(tile.Sequence),
			d2enum.TileLeftPartOfNorthCornerWall,
		)
		newTileData = &newTileOptions[tile.RandomIndex]
	}

	tileMinY := int32(0)
	tileMaxY := int32(0)

	target := tileData

	if newTileData != nil && newTileData.Height < tileData.Height {
		target = newTileData
	}

	for _, block := range target.Blocks {
		tileMinY = d2math.MinInt32(tileMinY, int32(block.Y))
		tileMaxY = d2math.MaxInt32(tileMaxY, int32(block.Y+blockOffsetY))
	}

	realHeight := d2math.MaxInt32(d2math.AbsInt32(tileData.Height), tileMaxY-tileMinY)
	tileYOffset := -tileMinY

	if tile.Type == d2enum.TileRoof {
		tile.YAdjust = -int(tileData.RoofHeight)
	} else {
		tile.YAdjust = int(tileMinY) + tileSurfaceHeight
	}

	cachedImage := mr.getImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tile.RandomIndex)
	if cachedImage != nil {
		return
	}

	if realHeight == 0 {
		mr.Error("Invalid 0 height for wall tile")
		return
	}

	image := mr.renderer.NewSurface(tileSurfaceWidth, int(realHeight))

	indexData := make([]byte, tileSurfaceWidth*realHeight)

	d2dt1.DecodeTileGfxData(tileData.Blocks, &indexData, tileYOffset, tileSurfaceWidth)

	if newTileData != nil {
		d2dt1.DecodeTileGfxData(newTileData.Blocks, &indexData, tileYOffset, tileSurfaceWidth)
	}

	pixels := d2util.ImgIndexToRGBA(indexData, mr.palette)

	image.ReplacePixels(pixels)

	mr.setImageCacheRecord(tile.Style, tile.Sequence, tile.Type, tile.RandomIndex, image)
}
