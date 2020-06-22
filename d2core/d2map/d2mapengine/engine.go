package d2mapengine

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapstamp"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
)

// Represents the map data for a specific location
type MapEngine struct {
	seed          int64                      // The map seed
	entities      []d2mapentity.MapEntity    // Entities on the map
	tiles         []d2ds1.TileRecord         // The map tiles
	size          d2common.Size              // The size of the map, in tiles
	levelType     d2datadict.LevelTypeRecord // The level type of this map
	dt1TileData   []d2dt1.Tile               // The DT1 tile data
	walkMesh      []d2common.PathTile        // The walk mesh
	startSubTileX int                        // The starting X position
	startSubTileY int                        // The starting Y position
}

// Creates a new instance of the map engine
func CreateMapEngine() *MapEngine {
	engine := &MapEngine{}
	return engine
}

func (m *MapEngine) WalkMesh() *[]d2common.PathTile {
	return &m.walkMesh
}

// Returns the starting position on the map in sub-tiles
func (m *MapEngine) GetStartingPosition() (int, int) {
	return m.startSubTileX, m.startSubTileY
}

func (m *MapEngine) ResetMap(seed int64, levelType d2enum.RegionIdType, width, height int) {
	m.seed = seed
	m.entities = make([]d2mapentity.MapEntity, 0)
	m.levelType = d2datadict.LevelTypes[levelType]
	m.size = d2common.Size{Width: width, Height: height}
	m.tiles = make([]d2ds1.TileRecord, width*height)
	m.dt1TileData = make([]d2dt1.Tile, 0)
	m.walkMesh = make([]d2common.PathTile, width*height*25)

	for _, dtFileName := range m.levelType.Files {
		if len(dtFileName) == 0 || dtFileName == "0" {
			continue
		}
		fileData, err := d2asset.LoadFile("/data/global/tiles/" + dtFileName)
		if err != nil {
			panic(err)
		}
		dt1, _ := d2dt1.LoadDT1(fileData)
		m.dt1TileData = append(m.dt1TileData, dt1.Tiles...)
	}
}

// Returns the level type of this map
func (m *MapEngine) LevelType() d2datadict.LevelTypeRecord {
	return m.levelType
}

// Sets the seed of the map for generation
func (m *MapEngine) SetSeed(seed int64) {
	log.Printf("Setting map engine seed to %d", seed)
	m.seed = seed
}

// Returns the size of the map (in sub-tiles)
func (m *MapEngine) Size() d2common.Size {
	return m.size
}

// Returns the map's tiles
func (m *MapEngine) Tiles() *[]d2ds1.TileRecord {
	return &m.tiles
}

// Places a stamp at the specified location. Also adds any entities from the stamp to the map engine
func (m *MapEngine) PlaceStamp(stamp *d2mapstamp.Stamp, tileOffsetX, tileOffsetY int) {
	stampSize := stamp.Size()
	if (tileOffsetX < 0) || (tileOffsetY < 0) || ((tileOffsetX + stampSize.Width) > m.size.Width) || ((tileOffsetY + stampSize.Height) > m.size.Height) {
		panic("Tried placing a stamp outside the bounds of the map")
	}

	// Copy over the map tile data
	for y := 0; y < stampSize.Height; y++ {
		for x := 0; x < stampSize.Width; x++ {
			mapTileIdx := x + tileOffsetX + ((y + tileOffsetY) * stampSize.Width)
			m.tiles[mapTileIdx] = *stamp.Tile(x, y)
		}
	}

	// Copy over the entities
	m.entities = append(m.entities, stamp.Entities()...)
}

// Returns a reference to a map tile based on the specified tile X and Y coordinate
func (m *MapEngine) TileAt(tileX, tileY int) *d2ds1.TileRecord {
	idx := tileX + (tileY * m.size.Width)
	if idx < 0 || idx >= len(m.tiles) {
		return nil
	}
	return &m.tiles[idx]
}

// Returns a reference to the map entities
func (m *MapEngine) Entities() *[]d2mapentity.MapEntity {
	return &m.entities
}

// Returns the map engine's seed
func (m *MapEngine) Seed() int64 {
	return m.seed
}

// Adds an entity to the map engine
func (m *MapEngine) AddEntity(entity d2mapentity.MapEntity) {
	m.entities = append(m.entities, entity)
}

// Removes an entity from the map engine
func (m *MapEngine) RemoveEntity(entity d2mapentity.MapEntity) {
	if entity == nil {
		return
	}
	panic("Removing entities is not currently implemented")
	//m.entities.Remove(entity)
}

func (m *MapEngine) GetTiles(style, sequence, tileType int32) []d2dt1.Tile {
	var tiles []d2dt1.Tile
	for _, tile := range m.dt1TileData {
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

func (m *MapEngine) GetStartPosition() (float64, float64) {
	for tileY := 0; tileY < m.size.Height; tileY++ {
		for tileX := 0; tileX < m.size.Width; tileX++ {
			tile := m.tiles[tileX+(tileY*m.size.Width)]
			for _, wall := range tile.Walls {
				if wall.Type.Special() && wall.Style == 30 {
					return float64(tileX) + 0.5, float64(tileY) + 0.5
				}
			}
		}
	}

	return m.GetCenterPosition()
}

// Returns the center of the map
func (m *MapEngine) GetCenterPosition() (float64, float64) {
	return float64(m.size.Width) / 2.0, float64(m.size.Height) / 2.0
}

// Advances time on the map engine
func (m *MapEngine) Advance(tickTime float64) {
	for _, entity := range m.entities {
		entity.Advance(tickTime)
	}
}

func (m *MapEngine) TileExists(tileX, tileY int) bool {
	if tileX < 0 || tileX >= m.size.Width || tileY < 0 || tileY >= m.size.Height {
		return false
	}
	tile := m.tiles[tileX+(tileY*m.size.Width)]
	return len(tile.Floors) > 0 || len(tile.Shadows) > 0 || len(tile.Walls) > 0 || len(tile.Substitutions) > 0
}

func (m *MapEngine) GenerateMap(regionType d2enum.RegionIdType, levelPreset int, fileIndex int, cacheTiles bool) {
	region := d2mapstamp.LoadStamp(m.seed, regionType, levelPreset, fileIndex)
	regionSize := region.Size()
	m.ResetMap(0, regionType, regionSize.Width, regionSize.Height)
	m.PlaceStamp(region, 0, 0)
}

func (m *MapEngine) GetTileData(style int32, sequence int32, tileType d2enum.TileType) *d2dt1.Tile {
	for _, tile := range m.dt1TileData {
		if tile.Style == style && tile.Sequence == sequence && tile.Type == int32(tileType) {
			return &tile
		}
	}
	return nil
}
