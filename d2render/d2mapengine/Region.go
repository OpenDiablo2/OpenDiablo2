package d2mapengine

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2render"

	"github.com/OpenDiablo2/OpenDiablo2/d2data"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"

	"github.com/hajimehoshi/ebiten"
)

type TileCacheRecord struct {
	Image   *ebiten.Image
	XOffset int
	YOffset int
}

type Region struct {
	RegionPath        string
	LevelType         d2datadict.LevelTypeRecord
	levelPreset       *d2datadict.LevelPresetRecord
	TileWidth         int32
	TileHeight        int32
	Tiles             []d2data.Tile
	DS1               *d2data.DS1
	Palette           d2datadict.PaletteRec
	FloorCache        map[uint32]*TileCacheRecord
	ShadowCache       map[uint32]*TileCacheRecord
	WallCache         map[uint32]*TileCacheRecord
	AnimationEntities []*d2render.AnimatedEntity
	NPCs              []*d2core.NPC
	StartX            float64
	StartY            float64
}

type RegionLayerType int

const (
	RegionLayerTypeFloors  RegionLayerType = 0
	RegionLayerTypeWalls   RegionLayerType = 1
	RegionLayerTypeShadows RegionLayerType = 2
)

type RegionIdType int

const (
	RegionAct1Town       RegionIdType = 1
	RegionAct1Wilderness RegionIdType = 2
	RegionAct1Cave       RegionIdType = 3
	RegionAct1Crypt      RegionIdType = 4
	RegionAct1Monestary  RegionIdType = 5
	RegionAct1Courtyard  RegionIdType = 6
	RegionAct1Barracks   RegionIdType = 7
	RegionAct1Jail       RegionIdType = 8
	RegionAct1Cathedral  RegionIdType = 9
	RegionAct1Catacombs  RegionIdType = 10
	RegionAct1Tristram   RegionIdType = 11
	RegionAct2Town       RegionIdType = 12
	RegionAct2Sewer      RegionIdType = 13
	RegionAct2Harem      RegionIdType = 14
	RegionAct2Basement   RegionIdType = 15
	RegionAct2Desert     RegionIdType = 16
	RegionAct2Tomb       RegionIdType = 17
	RegionAct2Lair       RegionIdType = 18
	RegionAct2Arcane     RegionIdType = 19
	RegionAct3Town       RegionIdType = 20
	RegionAct3Jungle     RegionIdType = 21
	RegionAct3Kurast     RegionIdType = 22
	RegionAct3Spider     RegionIdType = 23
	RegionAct3Dungeon    RegionIdType = 24
	RegionAct3Sewer      RegionIdType = 25
	RegionAct4Town       RegionIdType = 26
	RegionAct4Mesa       RegionIdType = 27
	RegionAct4Lava       RegionIdType = 28
	RegonAct5Town        RegionIdType = 29
	RegionAct5Siege      RegionIdType = 30
	RegionAct5Barricade  RegionIdType = 31
	RegionAct5Temple     RegionIdType = 32
	RegionAct5IceCaves   RegionIdType = 33
	RegionAct5Baal       RegionIdType = 34
	RegionAct5Lava       RegionIdType = 35
)

func LoadRegion(seed rand.Source, levelType RegionIdType, levelPreset int, fileProvider d2interface.FileProvider) *Region {
	result := &Region{
		LevelType:   d2datadict.LevelTypes[levelType],
		levelPreset: d2datadict.LevelPresets[levelPreset],
		Tiles:       make([]d2data.Tile, 0),
		FloorCache:  make(map[uint32]*TileCacheRecord),
		ShadowCache: make(map[uint32]*TileCacheRecord),
		WallCache:   make(map[uint32]*TileCacheRecord),
	}
	result.Palette = d2datadict.Palettes[d2enum.PaletteType("act"+strconv.Itoa(int(result.LevelType.Act)))]
	//\bm := result.levelPreset.Dt1Mask
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
		dt1 := d2data.LoadDT1("/data/global/tiles/"+levelTypeDt1, fileProvider)
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
	result.DS1 = d2data.LoadDS1("/data/global/tiles/"+levelFile, fileProvider)
	result.TileWidth = result.DS1.Width
	result.TileHeight = result.DS1.Height
	result.loadObjects(fileProvider)
	return result
}

func (v *Region) loadObjects(fileProvider d2interface.FileProvider) {
	var wg sync.WaitGroup
	wg.Add(len(v.DS1.Objects))
	v.AnimationEntities = make([]*d2render.AnimatedEntity, 0)
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
				npc := d2core.CreateNPC(object, fileProvider)
				v.NPCs = append(v.NPCs, npc)
			case d2datadict.ObjectTypeItem:
				if object.ObjectInfo == nil || !object.ObjectInfo.Draw || object.Lookup.Base == "" || object.Lookup.Token == "" {
					return
				}
				entity := d2render.CreateAnimatedEntity(object, fileProvider, d2enum.Units)
				entity.SetMode(object.Lookup.Mode, object.Lookup.Class, 0, fileProvider)
				v.AnimationEntities = append(v.AnimationEntities, entity)
			}
		}(object)
	}
	wg.Wait()
}

func (v *Region) RenderTile(offsetX, offsetY, tileX, tileY int, layerType RegionLayerType, layerIndex int, target *ebiten.Image) {
	offsetX -= 80
	switch layerType {
	case RegionLayerTypeFloors:
		v.renderFloor(v.DS1.Tiles[tileY][tileX].Floors[layerIndex], offsetX, offsetY, target)
	case RegionLayerTypeWalls:
		v.renderWall(v.DS1.Tiles[tileY][tileX].Walls[layerIndex], offsetX, offsetY, target)
	case RegionLayerTypeShadows:
		v.renderShadow(v.DS1.Tiles[tileY][tileX].Shadows[layerIndex], offsetX, offsetY, target)
	}
}

func (v *Region) getTile(mainIndex, subIndex, orientation int32) *d2data.Tile {
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

func (v *Region) renderFloor(tile d2data.FloorShadowRecord, offsetX, offsetY int, target *ebiten.Image) {
	tileCacheIndex := (uint32(tile.MainIndex) << 16) | (uint32(tile.SubIndex) << 8)
	tileCache, exists := v.FloorCache[tileCacheIndex]
	if !exists {
		v.FloorCache[tileCacheIndex] = v.generateFloorCache(tile)
		tileCache = v.FloorCache[tileCacheIndex]
		if tileCache == nil {
			log.Fatal("Could not load floor tile")
		}
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(offsetX+tileCache.XOffset), float64(offsetY+tileCache.YOffset))
	target.DrawImage(tileCache.Image, opts)
}

func (v *Region) renderWall(tile d2data.WallRecord, offsetX, offsetY int, target *ebiten.Image) {
	tileCacheIndex := (uint32(tile.MainIndex) << 16) | (uint32(tile.SubIndex) << 8) | (uint32(tile.Orientation))
	tileCache, exists := v.WallCache[tileCacheIndex]
	if !exists {
		v.WallCache[tileCacheIndex] = v.generateWallCache(tile)
		if v.WallCache[tileCacheIndex] == nil {
			log.Fatal("Could not generate wall")
		}
		tileCache = v.WallCache[tileCacheIndex]
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(offsetX+tileCache.XOffset), float64(offsetY+tileCache.YOffset))
	target.DrawImage(tileCache.Image, opts)
}

func (v *Region) renderShadow(tile d2data.FloorShadowRecord, offsetX, offsetY int, target *ebiten.Image) {
	tileCacheIndex := (uint32(tile.MainIndex) << 16) + (uint32(tile.SubIndex) << 8) + 0
	tileCache, exists := v.ShadowCache[tileCacheIndex]
	if !exists {
		v.ShadowCache[tileCacheIndex] = v.generateShadowCache(tile)
		tileCache = v.ShadowCache[tileCacheIndex]
		if tileCache == nil {
			log.Fatal("Could not load shadow tile")
		}
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(offsetX+tileCache.XOffset), float64(offsetY+tileCache.YOffset))
	opts.ColorM = d2helper.ColorToColorM(color.RGBA{255, 255, 255, 160})
	target.DrawImage(tileCache.Image, opts)
}

func (v *Region) decodeTileGfxData(blocks []d2data.Block, pixels []byte, tileYOffset int32, tileWidth int32) {
	for _, block := range blocks {
		if block.Format == d2data.BlockFormatIsometric {
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

func (v *Region) generateFloorCache(tile d2data.FloorShadowRecord) *TileCacheRecord {
	tileData := v.getTile(int32(tile.MainIndex), int32(tile.SubIndex), 0)
	if tileData == nil {
		log.Fatalf("Could not locate tile Idx:%d, Sub: %d, Ori: %d", tile.MainIndex, tile.SubIndex, 0)
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

func (v *Region) generateShadowCache(tile d2data.FloorShadowRecord) *TileCacheRecord {
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

func (v *Region) generateWallCache(tile d2data.WallRecord) *TileCacheRecord {
	tileData := v.getTile(int32(tile.MainIndex), int32(tile.SubIndex), int32(tile.Orientation))
	if tileData == nil {
		return nil
	}
	var newTileData *d2data.Tile = nil
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

	image.ReplacePixels(pixels)
	return &TileCacheRecord{
		image,
		0,
		yAdjust,
	}
}
