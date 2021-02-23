package d2ds1

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2path"
)

const (
	subType1 = 1
	subType2 = 2
	v2       = 2
	v3       = 3
	v4       = 4
	v7       = 7
	v8       = 8
	v9       = 9
	v10      = 10
	v12      = 12
	v13      = 13
	v14      = 14
	v15      = 15
	v16      = 16
	v18      = 18
)

const (
	wallZeroBitmask = 0xFFFFFF00
	wallZeroOffset  = 8

	wallTypeBitmask = 0x000000FF
)

const (
	unknown1BytesCount = 8
)

// DS1 represents the "stamp" data that is used to build up maps.
type DS1 struct {
	files                      []string            // FilePtr table of file string pointers
	objects                    []Object            // objects
	tiles                      [][]Tile            // The tile data for the DS1
	substitutionGroups         []SubstitutionGroup // Substitution groups for the DS1
	version                    int32               // The version of the DS1
	width                      int32               // Width of map, in # of tiles
	height                     int32               // Height of map, in # of tiles
	act                        int32               // Act, from 1 to 5. This tells which act table to use for the objects list
	substitutionType           int32               // SubstitutionType (layer type): 0 if no layer, else type 1 or type 2
	numberOfWallLayers         int32               // WallNum number of wall & orientation layers used
	numberOfFloorLayers        int32               // number of floor layers used
	numberOfShadowLayers       int32               // ShadowNum number of shadow layer used
	numberOfSubstitutionLayers int32               // SubstitutionNum number of substitution layer used
	// substitutionGroupsNum      int32               // SubstitutionGroupsNum number of substitution groups, datas between objects & NPC paths

	dirty    bool // when modifying tiles, need to perform upkeep on ds1 state
	unknown1 []byte
	unknown2 uint32
}

// Files returns a list of file path strings.
// These correspond to DT1 paths that this DS1 will use.
func (ds1 *DS1) Files() []string {
	return ds1.files
}

// AddFile adds a file path to the list of file paths
func (ds1 *DS1) AddFile(file string) {
	if ds1.files == nil {
		ds1.files = make([]string, 0)
	}

	ds1.files = append(ds1.files, file)
}

// RemoveFile removes a file from the files slice
func (ds1 *DS1) RemoveFile(file string) error {
	for idx := range ds1.files {
		if ds1.files[idx] == file {
			ds1.files = append(ds1.files[:idx], ds1.files[idx+1:]...)

			return nil
		}
	}

	return fmt.Errorf("file %s not found", file)
}

// Objects returns the slice of objects found in this ds1
func (ds1 *DS1) Objects() []Object {
	return ds1.objects
}

// AddObject adds an object to this ds1
func (ds1 *DS1) AddObject(obj Object) {
	if ds1.objects == nil {
		ds1.objects = make([]Object, 0)
	}

	ds1.objects = append(ds1.objects, obj)
}

// RemoveObject removes the first equivalent object found in this ds1's object list
func (ds1 *DS1) RemoveObject(obj Object) {
	for idx := range ds1.objects {
		if ds1.objects[idx].Equals(&obj) {
			ds1.objects = append(ds1.objects[:idx], ds1.objects[idx+1:]...)
			break
		}
	}
}

func defaultTiles() [][]Tile {
	return [][]Tile{{makeDefaultTile()}}
}

// Tiles returns the 2-dimensional (y,x) slice of tiles
func (ds1 *DS1) Tiles() [][]Tile {
	if ds1.tiles == nil {
		ds1.SetTiles(defaultTiles())
	}

	return ds1.tiles
}

// SetTiles sets the 2-dimensional (y,x) slice of tiles for this ds1
func (ds1 *DS1) SetTiles(tiles [][]Tile) {
	if len(tiles) == 0 {
		tiles = defaultTiles()
	}

	ds1.tiles = tiles
	ds1.dirty = true
	ds1.update()
}

// Tile returns the tile at the given x,y tile coordinate (nil if x,y is out of bounds)
func (ds1 *DS1) Tile(x, y int) *Tile {
	if ds1.dirty {
		ds1.update()
	}

	if y >= len(ds1.tiles) {
		return nil
	}

	if x >= len(ds1.tiles[y]) {
		return nil
	}

	return &ds1.tiles[y][x]
}

// SetTile sets the tile at the given tile x,y coordinates
func (ds1 *DS1) SetTile(x, y int, t *Tile) {
	if ds1.Tile(x, y) == nil {
		return
	}

	ds1.tiles[y][x] = *t
	ds1.dirty = true
	ds1.update()
}

// Version returns the ds1's version
func (ds1 *DS1) Version() int {
	return int(ds1.version)
}

// SetVersion sets the ds1's version
func (ds1 *DS1) SetVersion(v int) {
	ds1.version = int32(v)
}

// Width returns te ds1's width
func (ds1 *DS1) Width() int {
	if ds1.dirty {
		ds1.update()
	}

	return int(ds1.width)
}

// SetWidth sets the ds1's width
func (ds1 *DS1) SetWidth(w int) {
	if w <= 1 {
		w = 1
	}

	if int(ds1.width) == w {
		return
	}

	ds1.dirty = true // we know we're about to edit this ds1
	defer ds1.update()

	for rowIdx := range ds1.tiles {
		// if the row has too many tiles
		if len(ds1.tiles[rowIdx]) > w {
			// remove the extras
			ds1.tiles[rowIdx] = ds1.tiles[rowIdx][:w]
		}

		// if the row doesn't have enough tiles
		if len(ds1.tiles[rowIdx]) < w {
			// figure out how many more we need
			numNeeded := w - len(ds1.tiles[rowIdx])
			newTiles := make([]Tile, numNeeded)

			// make new default tiles
			for idx := range newTiles {
				newTiles[idx] = makeDefaultTile()
			}

			// add them to this ds1
			ds1.tiles[rowIdx] = append(ds1.tiles[rowIdx], newTiles...)
		}
	}

	ds1.width = int32(w)
}

// Height returns te ds1's height
func (ds1 *DS1) Height() int {
	if ds1.dirty {
		ds1.update()
	}

	return int(ds1.height)
}

// SetHeight sets the ds1's height
func (ds1 *DS1) SetHeight(h int) {
	if h <= 1 {
		h = 1
	}

	if int(ds1.height) == h {
		return
	}

	if len(ds1.tiles) < h {
		ds1.dirty = true // we know we're about to edit this ds1
		defer ds1.update()

		// figure out how many more rows we need
		numRowsNeeded := h - len(ds1.tiles)
		newRows := make([][]Tile, numRowsNeeded)

		// populate the new rows with tiles
		for rowIdx := range newRows {
			newRows[rowIdx] = make([]Tile, ds1.width)

			for colIdx := range newRows[rowIdx] {
				newRows[rowIdx][colIdx] = makeDefaultTile()
			}
		}

		ds1.tiles = append(ds1.tiles, newRows...)
	}

	// if the ds1 has too many rows
	if len(ds1.tiles) > h {
		ds1.dirty = true // we know we're about to edit this ds1
		defer ds1.update()

		// remove the extras
		ds1.tiles = ds1.tiles[:h]
	}
}

// Size returns te ds1's size (width, height)
func (ds1 *DS1) Size() (w, h int) {
	if ds1.dirty {
		ds1.update()
	}

	return int(ds1.width), int(ds1.height)
}

// SetSize force sets the ds1's size (width,height)
func (ds1 *DS1) SetSize(w, h int) {
	ds1.SetWidth(w)
	ds1.SetHeight(h)
	ds1.width, ds1.height = int32(w), int32(h)
}

// Act returns the ds1's act
func (ds1 *DS1) Act() int {
	return int(ds1.act)
}

// SetAct sets the ds1's act
func (ds1 *DS1) SetAct(act int) {
	if act < 0 {
		act = 0
	}

	ds1.act = int32(act)
}

// SubstitutionType returns the ds1's subtitution type
func (ds1 *DS1) SubstitutionType() int {
	return int(ds1.substitutionType)
}

// SetSubstitutionType sets the ds1's subtitution type
func (ds1 *DS1) SetSubstitutionType(t int) {
	ds1.substitutionType = int32(t)
}

// NumberOfWallLayers returns the number of wall layers per tile
func (ds1 *DS1) NumberOfWallLayers() int {
	if ds1.dirty {
		ds1.update()
	}

	return int(ds1.numberOfWallLayers)
}

// SetNumberOfWallLayers sets new number of tiles' walls
func (ds1 *DS1) SetNumberOfWallLayers(n int32) error {
	if n > d2enum.MaxNumberOfWalls {
		return fmt.Errorf("cannot set number of walls to %d: number of walls is greater than %d", n, d2enum.MaxNumberOfWalls)
	}

	for y := range ds1.tiles {
		for x := range ds1.tiles[y] {
			// ugh, I don't know, WHY do I nned to use
			// helper variable, but other way
			// simply doesn't work
			newWalls := ds1.tiles[y][x].Walls
			for v := int32(0); v < (n - int32(len(ds1.tiles[y][x].Walls))); v++ {
				newWalls = append(newWalls, Wall{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
			}

			ds1.tiles[y][x].Walls = newWalls
		}
	}

	// if n = number of walls, do nothing
	if n == ds1.numberOfWallLayers {
		return nil
	}

	for y := range ds1.tiles {
		for x := range ds1.tiles[y] {
			ds1.tiles[y][x].Walls = ds1.tiles[y][x].Walls[:n]
		}
	}

	ds1.numberOfWallLayers = n

	ds1.dirty = true
	ds1.update()

	return nil
}

// NumberOfFloorLayers returns the number of floor layers per tile
func (ds1 *DS1) NumberOfFloorLayers() int {
	if ds1.dirty {
		ds1.update()
	}

	return int(ds1.numberOfFloorLayers)
}

// SetNumberOfFloorLayers sets new number of tiles' floors
func (ds1 *DS1) SetNumberOfFloorLayers(n int32) error {
	if n > d2enum.MaxNumberOfFloors {
		return fmt.Errorf("cannot set number of floors to %d: number is greater than %d", n, d2enum.MaxNumberOfFloors)
	}

	for y := range ds1.tiles {
		for x := range ds1.tiles[y] {
			newFloors := ds1.tiles[y][x].Floors
			for v := int32(0); v < (n - int32(len(ds1.tiles[y][x].Floors))); v++ {
				newFloors = append(newFloors, Floor{})
			}

			ds1.tiles[y][x].Floors = newFloors
		}
	}

	// if n = number of walls, do nothing
	if n == ds1.numberOfFloorLayers {
		return nil
	}

	ds1.dirty = true
	defer ds1.update()

	for y := range ds1.tiles {
		for x := range ds1.tiles[y] {
			newFloors := make([]Floor, n)
			for v := int32(0); v < n; v++ {
				newFloors[v] = ds1.tiles[y][x].Floors[v]
			}

			ds1.tiles[y][x].Floors = newFloors
		}
	}

	ds1.numberOfFloorLayers = n

	return nil
}

// NumberOfShadowLayers returns the number of shadow layers per tile
func (ds1 *DS1) NumberOfShadowLayers() int {
	if ds1.dirty {
		ds1.update()
	}

	return int(ds1.numberOfShadowLayers)
}

// NumberOfSubstitutionLayers returns the number of substitution layers per tile
func (ds1 *DS1) NumberOfSubstitutionLayers() int {
	if ds1.dirty {
		ds1.update()
	}

	return int(ds1.numberOfSubstitutionLayers)
}

// SubstitutionGroups returns the number of wall layers per tile
func (ds1 *DS1) SubstitutionGroups() []SubstitutionGroup {
	return ds1.substitutionGroups
}

// SetSubstitutionGroups sets the substitution groups for the ds1
func (ds1 *DS1) SetSubstitutionGroups(groups []SubstitutionGroup) {
	ds1.substitutionGroups = groups
}

func (ds1 *DS1) update() {
	ds1.ensureAtLeastOneTile()
	ds1.enforceAllTileLayersMatch()
	ds1.updateLayerCounts()

	ds1.SetSize(len(ds1.tiles[0]), len(ds1.tiles))

	maxWalls := ds1.numberOfWallLayers

	for y := range ds1.tiles {
		for x := range ds1.tiles[y] {
			if len(ds1.tiles[y][x].Walls) > int(maxWalls) {
				maxWalls = int32(len(ds1.tiles[y][x].Walls))
			}
		}
	}

	err := ds1.SetNumberOfWallLayers(maxWalls)
	if err != nil {
		log.Print(err)
	}

	maxFloors := ds1.numberOfFloorLayers

	for y := range ds1.tiles {
		for x := range ds1.tiles[y] {
			if len(ds1.tiles[y][x].Floors) > int(maxFloors) {
				maxFloors = int32(len(ds1.tiles[y][x].Floors))
			}
		}
	}

	err = ds1.SetNumberOfFloorLayers(maxFloors)
	if err != nil {
		log.Print(err)
	}

	ds1.dirty = false
}

func (ds1 *DS1) ensureAtLeastOneTile() {
	// guarantee at least one tile exists
	if len(ds1.tiles) == 0 {
		ds1.tiles = [][]Tile{{makeDefaultTile()}}
	}
}

func (ds1 *DS1) enforceAllTileLayersMatch() {

}

func (ds1 *DS1) updateLayerCounts() {
	t := ds1.tiles[0][0] // the one tile that is guaranteed to exist
	ds1.numberOfFloorLayers = int32(len(t.Floors))
	ds1.numberOfShadowLayers = int32(len(t.Shadows))
	ds1.numberOfWallLayers = int32(len(t.Walls))
	ds1.numberOfSubstitutionLayers = int32(len(t.Substitutions))
}

// LoadDS1 loads the specified DS1 file
func LoadDS1(fileData []byte) (*DS1, error) {
	ds1 := &DS1{
		act:                        1,
		numberOfFloorLayers:        0,
		numberOfWallLayers:         0,
		numberOfShadowLayers:       1,
		numberOfSubstitutionLayers: 0,
	}

	br := d2datautils.CreateStreamReader(fileData)

	var err error

	err = ds1.loadHeader(br)
	if err != nil {
		return nil, fmt.Errorf("loading header: %v", err)
	}

	if ds1.version >= v9 && ds1.version <= v13 {
		// Skipping two dwords because they are "meaningless"?
		ds1.unknown1, err = br.ReadBytes(unknown1BytesCount)
		if err != nil {
			return nil, fmt.Errorf("reading unknown1: %v", err)
		}
	}

	if ds1.version >= v4 {
		ds1.numberOfWallLayers, err = br.ReadInt32()
		if err != nil {
			return nil, fmt.Errorf("reading wall number: %v", err)
		}

		if ds1.version >= v16 {
			ds1.numberOfFloorLayers, err = br.ReadInt32()
			if err != nil {
				return nil, fmt.Errorf("reading number of floors: %v", err)
			}
		} else {
			ds1.numberOfFloorLayers = 1
		}
	}

	ds1.tiles = make([][]Tile, ds1.height)

	for y := range ds1.tiles {
		ds1.tiles[y] = make([]Tile, ds1.width)
		for x := 0; x < int(ds1.width); x++ {
			ds1.tiles[y][x].Walls = make([]Wall, ds1.numberOfWallLayers)
			ds1.tiles[y][x].Floors = make([]Floor, ds1.numberOfFloorLayers)
			ds1.tiles[y][x].Shadows = make([]Shadow, ds1.numberOfShadowLayers)
			ds1.tiles[y][x].Substitutions = make([]Substitution, ds1.numberOfSubstitutionLayers)
		}
	}

	err = ds1.loadLayerStreams(br)
	if err != nil {
		return nil, fmt.Errorf("loading layer streams: %v", err)
	}

	err = ds1.loadObjects(br)
	if err != nil {
		return nil, fmt.Errorf("loading objects: %v", err)
	}

	err = ds1.loadSubstitutions(br)
	if err != nil {
		return nil, fmt.Errorf("loading substitutions: %v", err)
	}

	err = ds1.loadNPCs(br)
	if err != nil {
		return nil, fmt.Errorf("loading npc's: %v", err)
	}

	return ds1, nil
}

func (ds1 *DS1) loadHeader(br *d2datautils.StreamReader) error {
	var err error

	ds1.version, err = br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading version: %v", err)
	}

	ds1.width, err = br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading width: %v", err)
	}

	ds1.height, err = br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading height: %v", err)
	}

	ds1.width++
	ds1.height++

	if ds1.version >= v8 {
		ds1.act, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading act: %v", err)
		}

		ds1.act = d2math.MinInt32(d2enum.ActsNumber, ds1.act+1)
	}

	if ds1.version >= v10 {
		ds1.substitutionType, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution type: %v", err)
		}

		if ds1.substitutionType == 1 || ds1.substitutionType == 2 {
			ds1.numberOfSubstitutionLayers = 1
		}
	}

	err = ds1.loadFileList(br)
	if err != nil {
		return fmt.Errorf("loading file list: %v", err)
	}

	return nil
}

func (ds1 *DS1) loadFileList(br *d2datautils.StreamReader) error {
	if ds1.version >= v3 {
		// These files reference things that don't exist anymore :-?
		numberOfFiles, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading number of files: %v", err)
		}

		ds1.files = make([]string, numberOfFiles)

		for i := 0; i < int(numberOfFiles); i++ {
			ds1.files[i] = ""

			for {
				ch, err := br.ReadByte()
				if err != nil {
					return fmt.Errorf("reading file character: %v", err)
				}

				if ch == 0 {
					break
				}

				ds1.files[i] += string(ch)
			}
		}
	}

	return nil
}

func (ds1 *DS1) loadObjects(br *d2datautils.StreamReader) error {
	if ds1.version < v2 {
		ds1.objects = make([]Object, 0)
	} else {
		numberOfobjects, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading number of objects: %v", err)
		}

		ds1.objects = make([]Object, numberOfobjects)

		for objIdx := 0; objIdx < int(numberOfobjects); objIdx++ {
			obj := Object{}
			objType, err := br.ReadInt32()
			if err != nil {
				return fmt.Errorf("reading object's %d type: %v", objIdx, err)
			}

			objID, err := br.ReadInt32()
			if err != nil {
				return fmt.Errorf("reading object's %d ID: %v", objIdx, err)
			}

			objX, err := br.ReadInt32()
			if err != nil {
				return fmt.Errorf("reading object's %d X: %v", objIdx, err)
			}

			objY, err := br.ReadInt32()
			if err != nil {
				return fmt.Errorf("reading object's %d Y: %v", objY, err)
			}

			objFlags, err := br.ReadInt32()
			if err != nil {
				return fmt.Errorf("reading object's %d flags: %v", objIdx, err)
			}

			obj.Type = int(objType)
			obj.ID = int(objID)
			obj.X = int(objX)
			obj.Y = int(objY)
			obj.Flags = int(objFlags)

			ds1.objects[objIdx] = obj
		}
	}

	return nil
}

func (ds1 *DS1) loadSubstitutions(br *d2datautils.StreamReader) error {
	var err error

	hasSubstitutions := ds1.version >= v12 && (ds1.substitutionType == subType1 || ds1.substitutionType == subType2)

	if !hasSubstitutions {
		ds1.substitutionGroups = make([]SubstitutionGroup, 0)
		return nil
	}

	if ds1.version >= v18 {
		ds1.unknown2, err = br.ReadUInt32()
		if err != nil {
			return fmt.Errorf("reading unknown 2: %v", err)
		}
	}

	numberOfSubGroups, err := br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading number of sub groups: %v", err)
	}

	ds1.substitutionGroups = make([]SubstitutionGroup, numberOfSubGroups)

	for subIdx := 0; subIdx < int(numberOfSubGroups); subIdx++ {
		newSub := SubstitutionGroup{}

		newSub.TileX, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d X: %v", subIdx, err)
		}

		newSub.TileY, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d Y: %v", subIdx, err)
		}

		newSub.WidthInTiles, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d W: %v", subIdx, err)
		}

		newSub.HeightInTiles, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d H: %v", subIdx, err)
		}

		newSub.Unknown, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d unknown: %v", subIdx, err)
		}

		ds1.substitutionGroups[subIdx] = newSub
	}

	return err
}

// GetStreamLayerTypes returns layers used in ds1
func (ds1 *DS1) GetStreamLayerTypes() []d2enum.LayerStreamType {
	var layerStream []d2enum.LayerStreamType

	if ds1.version < v4 {
		layerStream = []d2enum.LayerStreamType{
			d2enum.LayerStreamWall1,
			d2enum.LayerStreamFloor1,
			d2enum.LayerStreamOrientation1,
			d2enum.LayerStreamSubstitute,
			d2enum.LayerStreamShadow,
		}
	} else {
		// nolint:gomnd // constant
		layerStream = make([]d2enum.LayerStreamType,
			(ds1.numberOfWallLayers*2)+ds1.numberOfFloorLayers+ds1.numberOfShadowLayers+ds1.numberOfSubstitutionLayers)

		layerIdx := 0
		for i := 0; i < int(ds1.numberOfWallLayers); i++ {
			layerStream[layerIdx] = d2enum.LayerStreamType(int(d2enum.LayerStreamWall1) + i)
			layerStream[layerIdx+1] = d2enum.LayerStreamType(int(d2enum.LayerStreamOrientation1) + i)
			layerIdx += 2
		}
		for i := 0; i < int(ds1.numberOfFloorLayers); i++ {
			layerStream[layerIdx] = d2enum.LayerStreamType(int(d2enum.LayerStreamFloor1) + i)
			layerIdx++
		}
		if ds1.numberOfShadowLayers > 0 {
			layerStream[layerIdx] = d2enum.LayerStreamShadow
			layerIdx++
		}
		if ds1.numberOfSubstitutionLayers > 0 {
			layerStream[layerIdx] = d2enum.LayerStreamSubstitute
		}
	}

	return layerStream
}

func (ds1 *DS1) loadNPCs(br *d2datautils.StreamReader) error {
	var err error

	if ds1.version < v14 {
		return nil
	}

	numberOfNpcs, err := br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading number of npcs: %v", err)
	}

	for npcIdx := 0; npcIdx < int(numberOfNpcs); npcIdx++ {
		numPaths, err := br.ReadInt32() // nolint:govet // I want to re-use this error variable
		if err != nil {
			return fmt.Errorf("reading number of paths for npc %d: %v", npcIdx, err)
		}

		npcX, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading X pos for NPC %d: %v", npcIdx, err)
		}

		npcY, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading Y pos for NPC %d: %v", npcIdx, err)
		}

		objIdx := -1

		for idx, ds1Obj := range ds1.objects {
			if ds1Obj.X == int(npcX) && ds1Obj.Y == int(npcY) {
				objIdx = idx

				break
			}
		}

		if objIdx > -1 {
			err = ds1.loadNpcPaths(br, objIdx, int(numPaths))
			if err != nil {
				return fmt.Errorf("loading paths for NPC %d: %v", npcIdx, err)
			}
		} else {
			if ds1.version >= v15 {
				br.SkipBytes(int(numPaths) * 3) //nolint:gomnd // Unknown data
			} else {
				br.SkipBytes(int(numPaths) * 2) //nolint:gomnd // Unknown data
			}
		}
	}

	return err
}

func (ds1 *DS1) loadNpcPaths(br *d2datautils.StreamReader, objIdx, numPaths int) error {
	if ds1.objects[objIdx].Paths == nil {
		ds1.objects[objIdx].Paths = make([]d2path.Path, numPaths)
	}

	for pathIdx := 0; pathIdx < numPaths; pathIdx++ {
		newPath := d2path.Path{}

		px, err := br.ReadInt32() //nolint:govet // i want to re-use the err variable...
		if err != nil {
			return fmt.Errorf("reading X point for path %d: %v", pathIdx, err)
		}

		py, err := br.ReadInt32() //nolint:govet // i want to re-use the err variable...
		if err != nil {
			return fmt.Errorf("reading Y point for path %d: %v", pathIdx, err)
		}

		newPath.Position = d2vector.NewPosition(float64(px), float64(py))

		if ds1.version >= v15 {
			action, err := br.ReadInt32()
			if err != nil {
				return fmt.Errorf("reading action for path %d: %v", pathIdx, err)
			}

			newPath.Action = int(action)
		}

		ds1.objects[objIdx].Paths[pathIdx] = newPath
	}

	return nil
}

func (ds1 *DS1) loadLayerStreams(br *d2datautils.StreamReader) error {
	var dirLookup = []int32{
		0x00, 0x01, 0x02, 0x01, 0x02, 0x03, 0x03, 0x05, 0x05, 0x06,
		0x06, 0x07, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
		0x0F, 0x10, 0x11, 0x12, 0x14,
	}

	layerStreamTypes := ds1.GetStreamLayerTypes()

	for _, layerStreamType := range layerStreamTypes {
		for y := 0; y < int(ds1.height); y++ {
			for x := 0; x < int(ds1.width); x++ {
				dw, err := br.ReadUInt32()
				if err != nil {
					return fmt.Errorf("reading layer's dword: %v", err)
				}

				switch layerStreamType {
				case d2enum.LayerStreamWall1, d2enum.LayerStreamWall2, d2enum.LayerStreamWall3, d2enum.LayerStreamWall4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamWall1)
					ds1.tiles[y][x].Walls[wallIndex].Decode(dw)
				case d2enum.LayerStreamOrientation1, d2enum.LayerStreamOrientation2,
					d2enum.LayerStreamOrientation3, d2enum.LayerStreamOrientation4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamOrientation1)
					c := int32(dw & wallTypeBitmask)

					if ds1.version < v7 {
						if c < int32(len(dirLookup)) {
							c = dirLookup[c]
						}
					}

					ds1.tiles[y][x].Walls[wallIndex].Type = d2enum.TileType(c)
					ds1.tiles[y][x].Walls[wallIndex].Zero = byte((dw & wallZeroBitmask) >> wallZeroOffset)
				case d2enum.LayerStreamFloor1, d2enum.LayerStreamFloor2:
					floorIndex := int(layerStreamType) - int(d2enum.LayerStreamFloor1)
					ds1.tiles[y][x].Floors[floorIndex].Decode(dw)
				case d2enum.LayerStreamShadow:
					ds1.tiles[y][x].Shadows[0].Decode(dw)
				case d2enum.LayerStreamSubstitute:
					ds1.tiles[y][x].Substitutions[0].Unknown = dw
				}
			}
		}
	}

	return nil
}

// Marshal encodes ds1 back to byte slice
func (ds1 *DS1) Marshal() []byte {
	// create stream writer
	sw := d2datautils.CreateStreamWriter()

	// Step 1 - encode header
	sw.PushInt32(ds1.version)
	sw.PushInt32(ds1.width - 1)
	sw.PushInt32(ds1.height - 1)

	if ds1.version >= v8 {
		sw.PushInt32(ds1.act - 1)
	}

	if ds1.version >= v10 {
		sw.PushInt32(ds1.substitutionType)
	}

	if ds1.version >= v3 {
		sw.PushInt32(int32(len(ds1.files)))

		for _, i := range ds1.files {
			sw.PushBytes([]byte(i)...)

			// separator
			sw.PushBytes(0)
		}
	}

	if ds1.version >= v9 && ds1.version <= v13 {
		sw.PushBytes(ds1.unknown1...)
	}

	if ds1.version >= v4 {
		sw.PushInt32(ds1.numberOfWallLayers)

		if ds1.version >= v16 {
			sw.PushInt32(ds1.numberOfFloorLayers)
		}
	}

	// Step 2 - encode layers
	ds1.encodeLayers(sw)

	// Step 3 - encode objects
	if !(ds1.version < v2) {
		sw.PushInt32(int32(len(ds1.objects)))

		for _, i := range ds1.objects {
			sw.PushUint32(uint32(i.Type))
			sw.PushUint32(uint32(i.ID))
			sw.PushUint32(uint32(i.X))
			sw.PushUint32(uint32(i.Y))
			sw.PushUint32(uint32(i.Flags))
		}
	}

	// Step 4 - encode substitutions
	if ds1.version >= v12 && (ds1.substitutionType == subType1 || ds1.substitutionType == subType2) {
		sw.PushUint32(ds1.unknown2)

		sw.PushUint32(uint32(len(ds1.substitutionGroups)))

		for _, i := range ds1.substitutionGroups {
			sw.PushInt32(i.TileX)
			sw.PushInt32(i.TileY)
			sw.PushInt32(i.WidthInTiles)
			sw.PushInt32(i.HeightInTiles)
			sw.PushInt32(i.Unknown)
		}
	}

	// Step 5 - encode NPC's and its paths
	ds1.encodeNPCs(sw)

	return sw.GetBytes()
}

func (ds1 *DS1) encodeLayers(sw *d2datautils.StreamWriter) {
	layerStreamTypes := ds1.GetStreamLayerTypes()

	for _, layerStreamType := range layerStreamTypes {
		for y := 0; y < int(ds1.height); y++ {
			for x := 0; x < int(ds1.width); x++ {
				dw := uint32(0)

				switch layerStreamType {
				case d2enum.LayerStreamWall1, d2enum.LayerStreamWall2, d2enum.LayerStreamWall3, d2enum.LayerStreamWall4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamWall1)
					ds1.tiles[y][x].Walls[wallIndex].Encode(sw)
				case d2enum.LayerStreamOrientation1, d2enum.LayerStreamOrientation2,
					d2enum.LayerStreamOrientation3, d2enum.LayerStreamOrientation4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamOrientation1)
					dw |= uint32(ds1.tiles[y][x].Walls[wallIndex].Type)
					dw |= (uint32(ds1.tiles[y][x].Walls[wallIndex].Zero) & wallZeroBitmask) << wallZeroOffset

					sw.PushUint32(dw)
				case d2enum.LayerStreamFloor1, d2enum.LayerStreamFloor2:
					floorIndex := int(layerStreamType) - int(d2enum.LayerStreamFloor1)
					ds1.tiles[y][x].Floors[floorIndex].Encode(sw)
				case d2enum.LayerStreamShadow:
					ds1.tiles[y][x].Shadows[0].Encode(sw)
				case d2enum.LayerStreamSubstitute:
					sw.PushUint32(ds1.tiles[y][x].Substitutions[0].Unknown)
				}
			}
		}
	}
}

func (ds1 *DS1) encodeNPCs(sw *d2datautils.StreamWriter) {
	objectsWithPaths := make([]int, 0)

	for n, obj := range ds1.objects {
		if len(obj.Paths) != 0 {
			objectsWithPaths = append(objectsWithPaths, n)
		}
	}

	// Step 5.1 - encode npc's
	sw.PushUint32(uint32(len(objectsWithPaths)))

	// Step 5.2 - enoce npcs' paths
	for objectIdx := range objectsWithPaths {
		sw.PushUint32(uint32(len(ds1.objects[objectIdx].Paths)))
		sw.PushUint32(uint32(ds1.objects[objectIdx].X))
		sw.PushUint32(uint32(ds1.objects[objectIdx].Y))

		for _, path := range ds1.objects[objectIdx].Paths {
			sw.PushUint32(uint32(path.Position.X()))
			sw.PushUint32(uint32(path.Position.Y()))

			if ds1.version >= v15 {
				sw.PushUint32(uint32(path.Action))
			}
		}
	}
}
