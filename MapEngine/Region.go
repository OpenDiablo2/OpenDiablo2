package MapEngine

import (
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"

	"github.com/essial/OpenDiablo2/Common"
)

type Region struct {
	levelType   Common.LevelTypeRecord
	levelPreset Common.LevelPresetRecord
	TileWidth   int32
	TileHeight  int32
	Tiles       []Tile
	DS1         *DS1
}

type RegionLayerType int

const (
	RegionLayerTypeFloors RegionLayerType = 0
	RegionLayerTypeWalls  RegionLayerType = 1
)

type RegionIdType int

const (
	RegionNoneRegionAct1Town              = 1
	RegionAct1Wilderness     RegionIdType = 2
	RegionAct1Cave           RegionIdType = 3
	RegionAct1Crypt          RegionIdType = 4
	RegionAct1Monestary      RegionIdType = 5
	RegionAct1Courtyard      RegionIdType = 6
	RegionAct1Barracks       RegionIdType = 7
	RegionAct1Jail           RegionIdType = 8
	RegionAct1Cathedral      RegionIdType = 9
	RegionAct1Catacombs      RegionIdType = 10
	RegionAct1Tristram       RegionIdType = 11
	RegionAct2Town           RegionIdType = 12
	RegionAct2Sewer          RegionIdType = 13
	RegionAct2Harem          RegionIdType = 14
	RegionAct2Basement       RegionIdType = 15
	RegionAct2Desert         RegionIdType = 16
	RegionAct2Tomb           RegionIdType = 17
	RegionAct2Lair           RegionIdType = 18
	RegionAct2Arcane         RegionIdType = 19
	RegionAct3Town           RegionIdType = 20
	RegionAct3Jungle         RegionIdType = 21
	RegionAct3Kurast         RegionIdType = 22
	RegionAct3Spider         RegionIdType = 23
	RegionAct3Dungeon        RegionIdType = 24
	RegionAct3Sewer          RegionIdType = 25
	RegionAct4Town           RegionIdType = 26
	RegionAct4Mesa           RegionIdType = 27
	RegionAct4Lava           RegionIdType = 28
	RegonAct5Town            RegionIdType = 29
	RegionAct5Siege          RegionIdType = 30
	RegionAct5Barricade      RegionIdType = 31
	RegionAct5Temple         RegionIdType = 32
	RegionAct5IceCaves       RegionIdType = 33
	RegionAct5Baal           RegionIdType = 34
	RegionAct5Lava           RegionIdType = 35
)

func LoadRegion(seed rand.Source, levelType RegionIdType, levelPreset int, fileProvider Common.FileProvider) *Region {
	result := &Region{
		levelType:   Common.LevelTypes[levelType],
		levelPreset: Common.LevelPresets[levelPreset],
		Tiles:       make([]Tile, 0),
	}
	for _, levelTypeDt1 := range result.levelType.Files {
		if len(levelTypeDt1) == 0 || levelTypeDt1 == "" || levelTypeDt1 == "0" {
			continue
		}
		dt1 := LoadDT1("/data/global/tiles/"+levelTypeDt1, fileProvider)
		for _, tileData := range dt1.Tiles {
			result.Tiles = append(result.Tiles, tileData)
		}
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
	result.DS1 = LoadDS1("/data/global/tiles/"+levelFile, fileProvider)
	result.TileWidth = result.DS1.Width
	result.TileHeight = result.DS1.Height

	return result
}

func (v *Region) RenderTile(offsetX, offsetY, tileX, tileY int, layerType RegionLayerType, layerIndex int, target *ebiten.Image) {
	switch layerType {
	case RegionLayerTypeFloors:
		v.renderFloor(v.DS1.Tiles[tileY][tileX].Floors[layerIndex], offsetX, offsetY, target)
	case RegionLayerTypeWalls:
		v.renderWall(v.DS1.Tiles[tileY][tileX].Walls[layerIndex], offsetX, offsetY, target)
	}
}

func (v *Region) getTile(mainIndex, subIndex int32) Tile {
	// TODO: Need to support randomly grabbing tile based on x/y as there can be multiple matches for same main/sub index
	for _, tile := range v.Tiles {
		if tile.MainIndex != mainIndex || tile.SubIndex != subIndex {
			continue
		}
		return tile
	}
	log.Fatalf("Unknown tile ID [%d %d]", mainIndex, subIndex)
	return Tile{}
}

func (v *Region) renderFloor(tile FloorShadowRecord, offsetX, offsetY int, target *ebiten.Image) {
	tileData := v.getTile(int32(tile.MainIndex), int32(tile.SubIndex))
	log.Printf("Pro1: %d", tileData.Direction)
}

func (v *Region) renderWall(tile WallRecord, offsetX, offsetY int, target *ebiten.Image) {

}
