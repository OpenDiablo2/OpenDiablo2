package d2mapstamp

import (
	"math"
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// Represents a pre-fabricated map stamp that can be placed on a map
type Stamp struct {
	regionPath  string                       // The file path of the region
	levelType   d2datadict.LevelTypeRecord   // The level type id for this stamp
	levelPreset d2datadict.LevelPresetRecord // The level preset id for this stamp
	tiles       []d2dt1.Tile                 // The tiles contained on this stamp
	ds1         *d2ds1.DS1                   // The backing DS1 file for this stamp
}

// Loads a stamp based on the supplied parameters
func LoadStamp(seed int64, levelType d2enum.RegionIdType, levelPreset int, fileIndex int) *Stamp {
	stamp := &Stamp{
		levelType:   d2datadict.LevelTypes[levelType],
		levelPreset: d2datadict.LevelPresets[levelPreset],
	}

	//stamp.palette, _ = loadPaletteForAct(levelType)

	for _, levelTypeDt1 := range stamp.levelType.Files {
		if len(levelTypeDt1) != 0 && levelTypeDt1 != "" && levelTypeDt1 != "0" {
			fileData, err := d2asset.LoadFile("/data/global/tiles/" + levelTypeDt1)
			if err != nil {
				panic(err)
			}

			dt1, _ := d2dt1.LoadDT1(fileData)

			stamp.tiles = append(stamp.tiles, dt1.Tiles...)
		}
	}

	var levelFilesToPick []string
	for _, fileRecord := range stamp.levelPreset.Files {
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

	stamp.regionPath = levelFilesToPick[levelIndex]
	fileData, err := d2asset.LoadFile("/data/global/tiles/" + stamp.regionPath)
	if err != nil {
		panic(err)
	}
	stamp.ds1, _ = d2ds1.LoadDS1(fileData)

	// Update the region info for the tiles
	for rx := 0; rx < len(stamp.ds1.Tiles); rx++ {
		for x := 0; x < len(stamp.ds1.Tiles[rx]); x++ {
			stamp.ds1.Tiles[rx][x].RegionType = levelType
		}
	}

	//entities := stamp.loadEntities()
	//stamp.loadSpecials()

	return stamp
}

// Returns the size of the stamp, in tiles
func (mr *Stamp) Size() d2common.Size {
	return d2common.Size{int(mr.ds1.Width), int(mr.ds1.Height)}
}

// Gets the level preset id
func (mr *Stamp) LevelPreset() d2datadict.LevelPresetRecord {
	return mr.levelPreset
}

// Returns the level type id
func (mr *Stamp) LevelType() d2datadict.LevelTypeRecord {
	return mr.levelType
}

// Gets the file path of the region
func (mr *Stamp) RegionPath() string {
	return mr.regionPath
}

// Returns the specified tile
func (mr *Stamp) Tile(x, y int) *d2ds1.TileRecord {
	return &mr.ds1.Tiles[y][x]
}

// Returns tile data based on the supplied paramters
func (mr *Stamp) TileData(style int32, sequence int32, tileType d2enum.TileType) *d2dt1.Tile {
	for _, tile := range mr.tiles {
		if tile.Style == style && tile.Sequence == sequence && tile.Type == int32(tileType) {
			return &tile
		}
	}
	return nil
}

func (mr *Stamp) Entities() []d2mapentity.MapEntity {
	entities := make([]d2mapentity.MapEntity, 0)

	for _, object := range mr.ds1.Objects {

		switch object.Lookup.Type {
		case d2datadict.ObjectTypeCharacter:
			if object.Lookup.Base != "" && object.Lookup.Token != "" && object.Lookup.TR != "" {
				npc := d2mapentity.CreateNPC(object.X, object.Y, object.Lookup, 0)
				npc.SetPaths(object.Paths)
				entities = append(entities, npc)
			}
		case d2datadict.ObjectTypeItem:
			if object.ObjectInfo != nil && object.ObjectInfo.Draw && object.Lookup.Base != "" && object.Lookup.Token != "" {
				entity, err := d2mapentity.CreateAnimatedComposite(object.X, object.Y, object.Lookup, d2resource.PaletteUnits)
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

//
//func (mr *Stamp) loadSpecials() {
//	for tileY := range mr.ds1.Tiles {
//		for tileX := range mr.ds1.Tiles[tileY] {
//			for _, wall := range mr.ds1.Tiles[tileY][tileX].Walls {
//				if wall.Type == 10 && wall.Style == 30 && wall.Sequence == 0 && mr.startX == 0 && mr.startY == 0 {
//					mr.startX, mr.startY = mr.getTileWorldPosition(tileX, tileY)
//					mr.startX += 0.5
//					mr.startY += 0.5
//					return
//				}
//			}
//		}
//	}
//}
//
