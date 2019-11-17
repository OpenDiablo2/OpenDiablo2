package d2mapengine

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2dt1"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2ds1"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2render"

	"github.com/OpenDiablo2/OpenDiablo2/d2data"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"

	"github.com/hajimehoshi/ebiten"
)

type Region struct {
	RegionPath        string
	LevelType         d2datadict.LevelTypeRecord
	levelPreset       d2datadict.LevelPresetRecord
	TileWidth         int32
	TileHeight        int32
	Tiles             []d2dt1.Tile
	DS1               d2ds1.DS1
	Palette           d2datadict.PaletteRec
	FloorCache        map[uint32]*TileCacheRecord
	ShadowCache       map[uint32]*TileCacheRecord
	WallCache         map[uint32]*TileCacheRecord
	AnimationEntities []d2render.AnimatedEntity
	NPCs              []*d2core.NPC
	StartX            float64
	StartY            float64
}

func LoadRegion(seed rand.Source, levelType d2enum.RegionIdType, levelPreset int, fileProvider d2interface.FileProvider) *Region {
	result := &Region{
		LevelType:   d2datadict.LevelTypes[levelType],
		levelPreset: d2datadict.LevelPresets[levelPreset],
		Tiles:       make([]d2dt1.Tile, 0),
		FloorCache:  make(map[uint32]*TileCacheRecord),
		ShadowCache: make(map[uint32]*TileCacheRecord),
		WallCache:   make(map[uint32]*TileCacheRecord),
	}
	result.Palette = d2datadict.Palettes[d2enum.PaletteType("act"+strconv.Itoa(int(result.LevelType.Act)))]
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
	for _, fileRecord := range result.levelPreset.Files {
		if len(fileRecord) == 0 || fileRecord == "" || fileRecord == "0" {
			continue
		}
		levelFilesToPick = append(levelFilesToPick, fileRecord)
	}
	random := rand.New(seed)
	levelIndex := int(math.Round(float64(len(levelFilesToPick)-1) * random.Float64()))
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
				if wall.Orientation != 10 {
					continue
				}
				if wall.MainIndex == 30 && wall.SubIndex == 0 {
					v.StartX = float64(x) + 0.5
					v.StartY = float64(y) + 0.5
					log.Printf("Starting location: %d, %d", x, y)
				}
			}
		}
	}
}

func (v *Region) loadObjects(fileProvider d2interface.FileProvider) {
	var wg sync.WaitGroup
	wg.Add(len(v.DS1.Objects))
	v.AnimationEntities = make([]d2render.AnimatedEntity, 0)
	v.NPCs = make([]*d2core.NPC, 0)
	for _, object := range v.DS1.Objects {
		go func(object d2data.Object) {
			defer wg.Done()
			switch object.Lookup.Type {
			case d2datadict.ObjectTypeCharacter:
				// Temp code, maybe..
				if object.Lookup.Base == "" || object.Lookup.Token == "" || object.Lookup.TR == "" {
					return
				}
				npc := d2core.CreateNPC(object.X, object.Y, object.Lookup, fileProvider, 1)
				npc.SetPaths(object.Paths)
				v.NPCs = append(v.NPCs, npc)
			case d2datadict.ObjectTypeItem:
				if object.ObjectInfo == nil || !object.ObjectInfo.Draw || object.Lookup.Base == "" || object.Lookup.Token == "" {
					return
				}
				entity := d2render.CreateAnimatedEntity(object.X, object.Y, object.Lookup, fileProvider, d2enum.Units)
				entity.SetMode(object.Lookup.Mode, object.Lookup.Class, 0)
				v.AnimationEntities = append(v.AnimationEntities, entity)
			}
		}(object)
	}
	wg.Wait()
}

func (v *Region) RenderTile(offsetX, offsetY, tileX, tileY int, layerType d2enum.RegionLayerType, layerIndex int, target *ebiten.Image) {
	offsetX -= 80
	switch layerType {
	case d2enum.RegionLayerTypeFloors:
		v.renderFloor(v.DS1.Tiles[tileY][tileX].Floors[layerIndex], offsetX, offsetY, target)
	case d2enum.RegionLayerTypeWalls:
		v.renderWall(v.DS1.Tiles[tileY][tileX].Walls[layerIndex], offsetX, offsetY, target)
	case d2enum.RegionLayerTypeShadows:
		v.renderShadow(v.DS1.Tiles[tileY][tileX].Shadows[layerIndex], offsetX, offsetY, target)
	}
}

func (v *Region) getTile(mainIndex, subIndex, orientation int32) *d2dt1.Tile {
	// TODO: Need to support randomly grabbing tile based on x/y as there can be multiple matches for same main/sub index
	for _, tile := range v.Tiles {
		if tile.MainIndex != mainIndex || tile.SubIndex != subIndex || tile.Orientation != orientation {
			continue
		}
		return &tile
	}
	//log.Fatalf("Unknown tile ID [%d %d %d]", mainIndex, subIndex, orientation)
	return nil
}

func (v *Region) renderFloor(tile d2ds1.FloorShadowRecord, offsetX, offsetY int, target *ebiten.Image) {
	tileCacheIndex := (uint32(tile.MainIndex) << 16) | (uint32(tile.SubIndex) << 8)
	tileCache, exists := v.FloorCache[tileCacheIndex]
	if !exists {
		v.FloorCache[tileCacheIndex] = v.generateFloorCache(tile)
		tileCache = v.FloorCache[tileCacheIndex]
		if tileCache == nil {
			log.Println("Could not load floor tile")
			return
		}
	}
	if tileCache == nil {
		log.Println("Nil tile cache")
		return
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(offsetX+tileCache.XOffset), float64(offsetY+tileCache.YOffset))
	target.DrawImage(tileCache.Image, opts)
}

func (v *Region) renderWall(tile d2ds1.WallRecord, offsetX, offsetY int, target *ebiten.Image) {
	tileCacheIndex := (uint32(tile.MainIndex) << 16) | (uint32(tile.SubIndex) << 8) | (uint32(tile.Orientation))
	tileCache, exists := v.WallCache[tileCacheIndex]
	if !exists {
		v.WallCache[tileCacheIndex] = v.generateWallCache(tile)
		if v.WallCache[tileCacheIndex] == nil {
			log.Println("Could not generate wall")
			return
		}
		tileCache = v.WallCache[tileCacheIndex]
	}
	if tileCache == nil {
		log.Println("Nil tile cache")
		return
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(offsetX+tileCache.XOffset), float64(offsetY+tileCache.YOffset))
	target.DrawImage(tileCache.Image, opts)
}

func (v *Region) renderShadow(tile d2ds1.FloorShadowRecord, offsetX, offsetY int, target *ebiten.Image) {
	tileCacheIndex := (uint32(tile.MainIndex) << 16) + (uint32(tile.SubIndex) << 8) + 0
	tileCache, exists := v.ShadowCache[tileCacheIndex]
	if !exists {
		v.ShadowCache[tileCacheIndex] = v.generateShadowCache(tile)
		tileCache = v.ShadowCache[tileCacheIndex]
		if tileCache == nil {
			log.Println("Could not load shadow tile")
			return
		}
	}
	if tileCache == nil {
		log.Println("Nil tile cache")
		return
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(offsetX+tileCache.XOffset), float64(offsetY+tileCache.YOffset))
	opts.ColorM = d2helper.ColorToColorM(color.RGBA{255, 255, 255, 160})
	target.DrawImage(tileCache.Image, opts)
}

func (v *Region) decodeTileGfxData(blocks []d2dt1.Block, pixels []byte, tileYOffset int32, tileWidth int32) {
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
						pixels[offset] = pixelColor.R
						pixels[offset+1] = pixelColor.G
						pixels[offset+2] = pixelColor.B
						pixels[offset+3] = 255
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
						pixels[offset] = pixelColor.R
						pixels[offset+1] = pixelColor.G
						pixels[offset+2] = pixelColor.B
						pixels[offset+3] = 255

					}
					idx++
					x++
					b2--
				}
			}
		}
	}
}

func (v *Region) generateFloorCache(tile d2ds1.FloorShadowRecord) *TileCacheRecord {
	tileData := v.getTile(int32(tile.MainIndex), int32(tile.SubIndex), 0)
	if tileData == nil {
		log.Printf("Could not locate tile Idx:%d, Sub: %d, Ori: %d\n", tile.MainIndex, tile.SubIndex, 0)
		tileData = &d2dt1.Tile{}
		tileData.Width = 10
		tileData.Height = 10
	}
	tileYMinimum := int32(0)
	for _, block := range tileData.Blocks {
		tileYMinimum = d2helper.MinInt32(tileYMinimum, int32(block.Y))
	}
	tileYOffset := d2helper.AbsInt32(tileYMinimum)
	tileHeight := d2helper.AbsInt32(tileData.Height)
	image, _ := ebiten.NewImage(int(tileData.Width), int(tileHeight), ebiten.FilterNearest)
	pixels := make([]byte, 4*tileData.Width*tileHeight)
	v.decodeTileGfxData(tileData.Blocks, pixels, tileYOffset, tileData.Width)
	image.ReplacePixels(pixels)
	return &TileCacheRecord{image, 0, 0}
}

func (v *Region) generateShadowCache(tile d2ds1.FloorShadowRecord) *TileCacheRecord {
	tileData := v.getTile(int32(tile.MainIndex), int32(tile.SubIndex), 13)
	if tileData == nil {
		return nil
	}
	tileMinY := int32(0)
	tileMaxY := int32(0)
	for _, block := range tileData.Blocks {
		tileMinY = d2helper.MinInt32(tileMinY, int32(block.Y))
		tileMaxY = d2helper.MaxInt32(tileMaxY, int32(block.Y+32))
	}
	tileYOffset := -tileMinY
	tileHeight := int(tileMaxY - tileMinY)
	image, _ := ebiten.NewImage(int(tileData.Width), int(tileHeight), ebiten.FilterNearest)
	pixels := make([]byte, 4*tileData.Width*int32(tileHeight))
	v.decodeTileGfxData(tileData.Blocks, pixels, tileYOffset, tileData.Width)
	image.ReplacePixels(pixels)
	return &TileCacheRecord{image, 0, int(tileMinY) + 80}
}

func (v *Region) generateWallCache(tile d2ds1.WallRecord) *TileCacheRecord {
	tileData := v.getTile(int32(tile.MainIndex), int32(tile.SubIndex), int32(tile.Orientation))
	if tileData == nil {
		return nil
	}
	var newTileData *d2dt1.Tile = nil
	if tile.Orientation == 3 {
		newTileData = v.getTile(int32(tile.MainIndex), int32(tile.SubIndex), int32(4))
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
	image, _ := ebiten.NewImage(160, int(realHeight), ebiten.FilterNearest)
	pixels := make([]byte, 4*160*realHeight)
	v.decodeTileGfxData(tileData.Blocks, pixels, tileYOffset, 160)
	if newTileData != nil {
		v.decodeTileGfxData(newTileData.Blocks, pixels, tileYOffset, 160)
	}
	yAdjust := 0
	if tile.Orientation > 15 {
		// Lower Walls
		yAdjust = 80
	} else if tile.Orientation == 15 {
		// Roof
		yAdjust = -int(tileData.RoofHeight)
	} else {
		// Upper Walls, Special Tiles
		yAdjust = int(tileMinY) + 80
	}

	// TODO: This may also need to be an atlas, but could get pretty large...
	if err := image.ReplacePixels(pixels); err != nil {
		log.Panicf(err.Error())
	}
	return &TileCacheRecord{
		image,
		0,
		yAdjust,
	}
}
