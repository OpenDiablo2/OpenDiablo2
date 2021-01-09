package d2mapengine

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapstamp"
)

const (
	logPrefix = "Map Engine"
)

// MapEngine loads the tiles which make up the isometric map and the entities
type MapEngine struct {
	asset *d2asset.AssetManager
	*d2mapstamp.StampFactory
	*d2mapentity.MapEntityFactory
	seed          int64                            // The map seed
	entities      map[string]d2interface.MapEntity // Entities on the map
	tiles         []MapTile
	size          d2geom.Size               // Size of the map, in tiles
	levelType     d2records.LevelTypeRecord // Level type of this map
	dt1TileData   []d2dt1.Tile              // DT1 tile data
	startSubTileX int                       // Starting X position
	startSubTileY int                       // Starting Y position
	dt1Files      []string                  // List of DS1 strings

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/789
	IsLoading bool // (temp) Whether we have processed the GenerateMapPacket(only for remote client)

	*d2util.Logger
}

const (
	subtilesPerTile = 5
)

// CreateMapEngine creates a new instance of the map engine and returns a pointer to it.
func CreateMapEngine(l d2util.LogLevel, asset *d2asset.AssetManager) *MapEngine {
	entity, _ := d2mapentity.NewMapEntityFactory(asset)
	stamp := d2mapstamp.NewStampFactory(asset, l, entity)

	engine := &MapEngine{
		asset:            asset,
		MapEntityFactory: entity,
		StampFactory:     stamp,
		// This will be set to true when we are using a remote client connection, and then set to false after we process the GenerateMapPacket
		IsLoading: false,
	}

	engine.Logger = d2util.NewLogger()
	engine.Logger.SetLevel(l)
	engine.Logger.SetPrefix(logPrefix)

	return engine
}

// GetStartingPosition returns the starting position on the map in sub-tiles.
func (m *MapEngine) GetStartingPosition() (x, y int) {
	return m.startSubTileX, m.startSubTileY
}

// ResetMap clears all map and entity data and reloads it from the cached files.
func (m *MapEngine) ResetMap(levelType d2enum.RegionIdType, width, height int) {
	m.entities = make(map[string]d2interface.MapEntity)
	m.levelType = *m.asset.Records.Level.Types[levelType]
	m.size = d2geom.Size{Width: width, Height: height}
	m.tiles = make([]MapTile, width*height)
	m.dt1TileData = make([]d2dt1.Tile, 0)
	m.dt1Files = make([]string, 0)

	for idx := range m.levelType.Files {
		m.addDT1(m.levelType.Files[idx])
	}
}

func (m *MapEngine) addDT1(fileName string) {
	if fileName == "" || fileName == "0" {
		return
	}

	fileName = strings.ToLower(fileName)
	for i := 0; i < len(m.dt1Files); i++ {
		if m.dt1Files[i] == fileName {
			return
		}
	}

	fileData, err := m.asset.LoadFile("/data/global/tiles/" + fileName)
	if err != nil {
		m.Fatalf("Could not load /data/global/tiles/%s", fileName)
		return
	}

	dt1, err := d2dt1.LoadDT1(fileData)
	if err != nil {
		m.Error(err.Error())
	}

	m.dt1TileData = append(m.dt1TileData, dt1.Tiles...)
	m.dt1Files = append(m.dt1Files, fileName)
}

// AddDS1 loads DT1 files and performs string replacements on them. It
// appends the tile data and files to MapEngine.dt1TileData and
// MapEngine.dt1Files.
func (m *MapEngine) AddDS1(fileName string) {
	if fileName == "" || fileName == "0" {
		return
	}

	fileData, err := m.asset.LoadFile("/data/global/tiles/" + fileName)
	if err != nil {
		panic(err)
	}

	ds1, err := d2ds1.LoadDS1(fileData)
	if err != nil {
		m.Error(err.Error())
	}

	for idx := range ds1.Files {
		dt1File := ds1.Files[idx]
		dt1File = strings.ToLower(dt1File)
		dt1File = strings.ReplaceAll(dt1File, "c:", "")       // Yes they did...
		dt1File = strings.ReplaceAll(dt1File, ".tg1", ".dt1") // Yes they did...
		dt1File = strings.ReplaceAll(dt1File, "\\d2\\data\\global\\tiles\\", "")
		m.addDT1(strings.ReplaceAll(dt1File, "\\", "/"))
	}
}

// LevelType returns the level type of this map.
func (m *MapEngine) LevelType() d2records.LevelTypeRecord {
	return m.levelType
}

// SetSeed sets the seed of the map for generation.
func (m *MapEngine) SetSeed(seed int64) {
	m.Infof("Setting map engine seed to %d", seed)
	m.seed = seed
}

// Size returns the size of the map in sub-tiles.
func (m *MapEngine) Size() d2geom.Size {
	return m.size
}

// Tile returns the TileRecord containing the data
// for a single map tile.
func (m *MapEngine) Tile(x, y int) *MapTile {
	return &m.tiles[x+(y*m.size.Width)]
}

// Tiles returns a pointer to a slice contaning all
// map tile data.
func (m *MapEngine) Tiles() *[]MapTile {
	return &m.tiles
}

// PlaceStamp places a map stamp at the specified location, creating both entities
// and tiles. Stamps are pre-defined map areas, see d2mapstamp.
func (m *MapEngine) PlaceStamp(stamp *d2mapstamp.Stamp, tileOffsetX, tileOffsetY int) {
	stampSize := stamp.Size()
	stampW := stampSize.Width
	stampH := stampSize.Height

	mapW := m.size.Width
	mapH := m.size.Height

	xMin := tileOffsetX
	yMin := tileOffsetY
	xMax := xMin + stampSize.Width
	yMax := yMin + stampSize.Height

	if (xMin < 0) || (yMin < 0) || (xMax > mapW) || (yMax > mapH) {
		panic("Tried placing a stamp outside the bounds of the map")
	}

	// Copy over the map tile data
	for y := 0; y < stampH; y++ {
		for x := 0; x < stampW; x++ {
			targetTileIndex := m.tileCoordinateToIndex(x+xMin, y+yMin)
			stampTile := *stamp.Tile(x, y)
			m.tiles[targetTileIndex].RegionType = stamp.RegionID()
			m.tiles[targetTileIndex].Components = stampTile
			m.tiles[targetTileIndex].PrepareTile(x, y, m)
		}
	}

	// Copy over the entities
	stampEntities := stamp.Entities(tileOffsetX, tileOffsetY)
	for idx := range stampEntities {
		e := stampEntities[idx]
		m.entities[e.ID()] = e
	}
}

// converts x,y tile coordinate into index in MapEngine.tiles
func (m *MapEngine) tileCoordinateToIndex(x, y int) int {
	return x + (y * m.size.Width)
}

// SubTileAt gets the flags for the given subtile
func (m *MapEngine) SubTileAt(subX, subY int) *d2dt1.SubTileFlags {
	tile := m.TileAt(subX/subtilesPerTile, subY/subtilesPerTile)

	return tile.GetSubTileFlags(subX%subtilesPerTile, subY%subtilesPerTile)
}

// TileAt returns a pointer to the data for the map tile at the given
// x and y index.
func (m *MapEngine) TileAt(tileX, tileY int) *MapTile {
	idx := m.tileCoordinateToIndex(tileX, tileY)
	if idx < 0 || idx >= len(m.tiles) {
		return nil
	}

	return &m.tiles[idx]
}

// Entities returns a pointer a slice of all map entities.
func (m *MapEngine) Entities() map[string]d2interface.MapEntity {
	return m.entities
}

// Seed returns the map generation seed.
func (m *MapEngine) Seed() int64 {
	return m.seed
}

// AddEntity adds an entity to a slice containing all entities.
func (m *MapEngine) AddEntity(entity d2interface.MapEntity) {
	m.entities[entity.ID()] = entity
}

// RemoveEntity removes an entity from the map engine
func (m *MapEngine) RemoveEntity(entity d2interface.MapEntity) {
	if entity == nil {
		return
	}

	delete(m.entities, entity.ID())
}

// GetTiles returns a slice of all tiles matching the given style,
// sequence and tileType.
func (m *MapEngine) GetTiles(style, sequence int, tileType d2enum.TileType) []d2dt1.Tile {
	tiles := make([]d2dt1.Tile, 0, len(m.dt1TileData))

	for idx := range m.dt1TileData {
		if m.dt1TileData[idx].Style != int32(style) || m.dt1TileData[idx].Sequence != int32(sequence) ||
			m.dt1TileData[idx].Type != int32(tileType) {
			continue
		}

		tiles = append(tiles, m.dt1TileData[idx])
	}

	if len(tiles) == 0 {
		m.Warningf("Unknown tile ID [%d %d %d]", style, sequence, tileType)
		return nil
	}

	return tiles
}

// GetStartPosition returns the spawn point on entering the current map.
func (m *MapEngine) GetStartPosition() (x, y float64) {
	for tileY := 0; tileY < m.size.Height; tileY++ {
		for tileX := 0; tileX < m.size.Width; tileX++ {
			tile := m.tiles[tileX+(tileY*m.size.Width)].Components
			for idx := range tile.Walls {
				if tile.Walls[idx].Type.Special() && tile.Walls[idx].Style == 30 {
					return float64(tileX) + 0.5, float64(tileY) + 0.5
				}
			}
		}
	}

	return m.GetCenterPosition()
}

// GetCenterPosition returns the center point of the map.
func (m *MapEngine) GetCenterPosition() (x, y float64) {
	return float64(m.size.Width) / 2.0, float64(m.size.Height) / 2.0
}

// Advance calls the Advance() method for all entities,
// processing a single tick.
func (m *MapEngine) Advance(tickTime float64) {
	if m.IsLoading {
		// https://github.com/OpenDiablo2/OpenDiablo2/issues/789
		return
	}

	for ID := range m.entities {
		m.entities[ID].Advance(tickTime)
	}
}

// TileExists returns true if the tile at the given coordinates exists.
func (m *MapEngine) TileExists(tileX, tileY int) bool {
	tileIndex := m.tileCoordinateToIndex(tileX, tileY)

	if valid := (tileIndex >= 0) && (tileIndex <= len(m.tiles)); valid {
		tile := m.tiles[tileIndex].Components
		numFeatures := len(tile.Floors)
		numFeatures += len(tile.Shadows)
		numFeatures += len(tile.Walls)
		numFeatures += len(tile.Substitutions)

		return numFeatures > 0
	}

	return false
}

// GenerateMap clears the map and places the specified stamp.
func (m *MapEngine) GenerateMap(regionType d2enum.RegionIdType, levelPreset, fileIndex int) {
	region := m.LoadStamp(regionType, levelPreset, fileIndex)
	regionSize := region.Size()
	m.ResetMap(regionType, regionSize.Width, regionSize.Height)
	m.PlaceStamp(region, 0, 0)
}

// GetTileData returns the tile with the given style, sequence, tileType and index.
func (m *MapEngine) GetTileData(style, sequence int, tileType d2enum.TileType, index byte) *d2dt1.Tile {
	for idx := range m.dt1TileData {
		if m.dt1TileData[idx].Style == int32(style) && m.dt1TileData[idx].Sequence == int32(sequence) &&
			m.dt1TileData[idx].Type == int32(tileType) && m.dt1TileData[idx].RarityFrameIndex == int32(index) {
			return &m.dt1TileData[idx]
		}
	}

	return nil
}
