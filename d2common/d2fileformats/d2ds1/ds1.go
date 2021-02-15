package d2ds1

import (
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
	Files                      []string            // FilePtr table of file string pointers
	Objects                    []Object            // Objects
	Tiles                      [][]TileRecord      // The tile data for the DS1
	SubstitutionGroups         []SubstitutionGroup // Substitution groups for the DS1
	Version                    int32               // The version of the DS1
	Width                      int32               // Width of map, in # of tiles
	Height                     int32               // Height of map, in # of tiles
	Act                        int32               // Act, from 1 to 5. This tells which act table to use for the Objects list
	SubstitutionType           int32               // SubstitutionType (layer type): 0 if no layer, else type 1 or type 2
	NumberOfWalls              int32               // WallNum number of wall & orientation layers used
	NumberOfFloors             int32               // number of floor layers used
	NumberOfShadowLayers       int32               // ShadowNum number of shadow layer used
	NumberOfSubstitutionLayers int32               // SubstitutionNum number of substitution layer used
	SubstitutionGroupsNum      int32               // SubstitutionGroupsNum number of substitution groups, datas between objects & NPC paths
	unknown1                   []byte
	LayerStreamTypes           []d2enum.LayerStreamType
	unknown2                   uint32
	NpcIndexes                 []int
}

// LoadDS1 loads the specified DS1 file
func LoadDS1(fileData []byte) (*DS1, error) {
	ds1 := &DS1{
		Act:                        1,
		NumberOfFloors:             0,
		NumberOfWalls:              0,
		NumberOfShadowLayers:       1,
		NumberOfSubstitutionLayers: 0,
	}

	br := d2datautils.CreateStreamReader(fileData)

	var err error

	err = ds1.loadHeader(br)
	if err != nil {
		return nil, err
	}

	if ds1.Version >= v9 && ds1.Version <= v13 {
		// Skipping two dwords because they are "meaningless"?
		ds1.unknown1, err = br.ReadBytes(unknown1BytesCount)
		if err != nil {
			return nil, err
		}
	}

	if ds1.Version >= v4 {
		ds1.NumberOfWalls, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		if ds1.Version >= v16 {
			ds1.NumberOfFloors, err = br.ReadInt32()
			if err != nil {
				return nil, err
			}
		} else {
			ds1.NumberOfFloors = 1
		}
	}

	ds1.LayerStreamTypes = ds1.setupStreamLayerTypes()

	ds1.Tiles = make([][]TileRecord, ds1.Height)

	for y := range ds1.Tiles {
		ds1.Tiles[y] = make([]TileRecord, ds1.Width)
		for x := 0; x < int(ds1.Width); x++ {
			ds1.Tiles[y][x].Walls = make([]WallRecord, ds1.NumberOfWalls)
			ds1.Tiles[y][x].Floors = make([]FloorShadowRecord, ds1.NumberOfFloors)
			ds1.Tiles[y][x].Shadows = make([]FloorShadowRecord, ds1.NumberOfShadowLayers)
			ds1.Tiles[y][x].Substitutions = make([]SubstitutionRecord, ds1.NumberOfSubstitutionLayers)
		}
	}

	err = ds1.loadLayerStreams(br)
	if err != nil {
		return nil, err
	}

	err = ds1.loadObjects(br)
	if err != nil {
		return nil, err
	}

	err = ds1.loadSubstitutions(br)
	if err != nil {
		return nil, err
	}

	err = ds1.loadNPCs(br)
	if err != nil {
		return nil, err
	}

	return ds1, nil
}

func (ds1 *DS1) loadHeader(br *d2datautils.StreamReader) error {
	var err error

	ds1.Version, err = br.ReadInt32()
	if err != nil {
		return err
	}

	ds1.Width, err = br.ReadInt32()
	if err != nil {
		return err
	}

	ds1.Height, err = br.ReadInt32()
	if err != nil {
		return err
	}

	ds1.Width++
	ds1.Height++

	if ds1.Version >= v8 {
		ds1.Act, err = br.ReadInt32()
		if err != nil {
			return err
		}

		ds1.Act = d2math.MinInt32(d2enum.ActsNumber, ds1.Act+1)
	}

	if ds1.Version >= v10 {
		ds1.SubstitutionType, err = br.ReadInt32()
		if err != nil {
			return err
		}

		if ds1.SubstitutionType == 1 || ds1.SubstitutionType == 2 {
			ds1.NumberOfSubstitutionLayers = 1
		}
	}

	err = ds1.loadFileList(br)
	if err != nil {
		return err
	}

	return nil
}

func (ds1 *DS1) loadFileList(br *d2datautils.StreamReader) error {
	if ds1.Version >= v3 {
		// These files reference things that don't exist anymore :-?
		numberOfFiles, err := br.ReadInt32()
		if err != nil {
			return err
		}

		ds1.Files = make([]string, numberOfFiles)

		for i := 0; i < int(numberOfFiles); i++ {
			ds1.Files[i] = ""

			for {
				ch, err := br.ReadByte()
				if err != nil {
					return err
				}

				if ch == 0 {
					break
				}

				ds1.Files[i] += string(ch)
			}
		}
	}

	return nil
}

func (ds1 *DS1) loadObjects(br *d2datautils.StreamReader) error {
	if ds1.Version < v2 {
		ds1.Objects = make([]Object, 0)
	} else {
		numberOfObjects, err := br.ReadInt32()
		if err != nil {
			return err
		}

		ds1.Objects = make([]Object, numberOfObjects)

		for objIdx := 0; objIdx < int(numberOfObjects); objIdx++ {
			obj := Object{}
			objType, err := br.ReadInt32()
			if err != nil {
				return err
			}

			objID, err := br.ReadInt32()
			if err != nil {
				return err
			}

			objX, err := br.ReadInt32()
			if err != nil {
				return err
			}

			objY, err := br.ReadInt32()
			if err != nil {
				return err
			}

			objFlags, err := br.ReadInt32()
			if err != nil {
				return err
			}

			obj.Type = int(objType)
			obj.ID = int(objID)
			obj.X = int(objX)
			obj.Y = int(objY)
			obj.Flags = int(objFlags)

			ds1.Objects[objIdx] = obj
		}
	}

	return nil
}

func (ds1 *DS1) loadSubstitutions(br *d2datautils.StreamReader) error {
	var err error

	hasSubstitutions := ds1.Version >= v12 && (ds1.SubstitutionType == subType1 || ds1.SubstitutionType == subType2)

	if !hasSubstitutions {
		ds1.SubstitutionGroups = make([]SubstitutionGroup, 0)
		return nil
	}

	if ds1.Version >= v18 {
		ds1.unknown2, err = br.ReadUInt32()
		if err != nil {
			return err
		}
	}

	numberOfSubGroups, err := br.ReadInt32()
	if err != nil {
		return err
	}

	ds1.SubstitutionGroups = make([]SubstitutionGroup, numberOfSubGroups)

	for subIdx := 0; subIdx < int(numberOfSubGroups); subIdx++ {
		newSub := SubstitutionGroup{}

		newSub.TileX, err = br.ReadInt32()
		if err != nil {
			return err
		}

		newSub.TileY, err = br.ReadInt32()
		if err != nil {
			return err
		}

		newSub.WidthInTiles, err = br.ReadInt32()
		if err != nil {
			return err
		}

		newSub.HeightInTiles, err = br.ReadInt32()
		if err != nil {
			return err
		}

		newSub.Unknown, err = br.ReadInt32()
		if err != nil {
			return err
		}

		ds1.SubstitutionGroups[subIdx] = newSub
	}

	return err
}

func (ds1 *DS1) setupStreamLayerTypes() []d2enum.LayerStreamType {
	var layerStream []d2enum.LayerStreamType

	if ds1.Version < v4 {
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
			(ds1.NumberOfWalls*2)+ds1.NumberOfFloors+ds1.NumberOfShadowLayers+ds1.NumberOfSubstitutionLayers)

		layerIdx := 0
		for i := 0; i < int(ds1.NumberOfWalls); i++ {
			layerStream[layerIdx] = d2enum.LayerStreamType(int(d2enum.LayerStreamWall1) + i)
			layerStream[layerIdx+1] = d2enum.LayerStreamType(int(d2enum.LayerStreamOrientation1) + i)
			layerIdx += 2
		}
		for i := 0; i < int(ds1.NumberOfFloors); i++ {
			layerStream[layerIdx] = d2enum.LayerStreamType(int(d2enum.LayerStreamFloor1) + i)
			layerIdx++
		}
		if ds1.NumberOfShadowLayers > 0 {
			layerStream[layerIdx] = d2enum.LayerStreamShadow
			layerIdx++
		}
		if ds1.NumberOfSubstitutionLayers > 0 {
			layerStream[layerIdx] = d2enum.LayerStreamSubstitute
		}
	}

	return layerStream
}

func (ds1 *DS1) loadNPCs(br *d2datautils.StreamReader) error {
	var err error

	if ds1.Version < v14 {
		return err
	}

	numberOfNpcs, err := br.ReadInt32()
	if err != nil {
		return err
	}

	for npcIdx := 0; npcIdx < int(numberOfNpcs); npcIdx++ {
		numPaths, err := br.ReadInt32() // nolint:govet // I want to re-use this error variable
		if err != nil {
			return err
		}

		npcX, err := br.ReadInt32()
		if err != nil {
			return err
		}

		npcY, err := br.ReadInt32()
		if err != nil {
			return err
		}

		objIdx := -1

		for idx, ds1Obj := range ds1.Objects {
			if ds1Obj.X == int(npcX) && ds1Obj.Y == int(npcY) {
				objIdx = idx
				ds1.NpcIndexes = append(ds1.NpcIndexes, idx)

				break
			}
		}

		if objIdx > -1 {
			err = ds1.loadNpcPaths(br, objIdx, int(numPaths))
			if err != nil {
				return err
			}
		} else {
			if ds1.Version >= v15 {
				br.SkipBytes(int(numPaths) * 3) //nolint:gomnd // Unknown data
			} else {
				br.SkipBytes(int(numPaths) * 2) //nolint:gomnd // Unknown data
			}
		}
	}

	return err
}

func (ds1 *DS1) loadNpcPaths(br *d2datautils.StreamReader, objIdx, numPaths int) error {
	var err error

	if ds1.Objects[objIdx].Paths == nil {
		ds1.Objects[objIdx].Paths = make([]d2path.Path, numPaths)
	}

	for pathIdx := 0; pathIdx < numPaths; pathIdx++ {
		newPath := d2path.Path{}

		px, err := br.ReadInt32() //nolint:govet // i want to re-use the err variable...
		if err != nil {
			return err
		}

		py, err := br.ReadInt32() //nolint:govet // i want to re-use the err variable...
		if err != nil {
			return err
		}

		newPath.Position = d2vector.NewPosition(float64(px), float64(py))

		if ds1.Version >= v15 {
			action, err := br.ReadInt32()
			if err != nil {
				return err
			}

			newPath.Action = int(action)
		}

		ds1.Objects[objIdx].Paths[pathIdx] = newPath
	}

	return err
}

func (ds1 *DS1) loadLayerStreams(br *d2datautils.StreamReader) error {
	var err error

	var dirLookup = []int32{
		0x00, 0x01, 0x02, 0x01, 0x02, 0x03, 0x03, 0x05, 0x05, 0x06,
		0x06, 0x07, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
		0x0F, 0x10, 0x11, 0x12, 0x14,
	}

	for lIdx := range ds1.LayerStreamTypes {
		layerStreamType := ds1.LayerStreamTypes[lIdx]

		for y := 0; y < int(ds1.Height); y++ {
			for x := 0; x < int(ds1.Width); x++ {
				dw, err := br.ReadUInt32() //nolint:govet // i want to re-use the err variable...
				if err != nil {
					return err
				}

				switch layerStreamType {
				case d2enum.LayerStreamWall1, d2enum.LayerStreamWall2, d2enum.LayerStreamWall3, d2enum.LayerStreamWall4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamWall1)
					ds1.Tiles[y][x].Walls[wallIndex].Decode(dw)
				case d2enum.LayerStreamOrientation1, d2enum.LayerStreamOrientation2,
					d2enum.LayerStreamOrientation3, d2enum.LayerStreamOrientation4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamOrientation1)
					c := int32(dw & wallTypeBitmask)

					if ds1.Version < v7 {
						if c < int32(len(dirLookup)) {
							c = dirLookup[c]
						}
					}

					ds1.Tiles[y][x].Walls[wallIndex].Type = d2enum.TileType(c)
					ds1.Tiles[y][x].Walls[wallIndex].Zero = byte((dw & wallZeroBitmask) >> wallZeroOffset)
				case d2enum.LayerStreamFloor1, d2enum.LayerStreamFloor2:
					floorIndex := int(layerStreamType) - int(d2enum.LayerStreamFloor1)
					ds1.Tiles[y][x].Floors[floorIndex].Decode(dw)
				case d2enum.LayerStreamShadow:
					ds1.Tiles[y][x].Shadows[0].Decode(dw)
				case d2enum.LayerStreamSubstitute:
					ds1.Tiles[y][x].Substitutions[0].Unknown = dw
				}
			}
		}
	}

	return err
}

// Marshal encodes ds1 back to byte slice
func (ds1 *DS1) Marshal() []byte {
	// create stream writer
	sw := d2datautils.CreateStreamWriter()

	// Step 1 - encode header
	sw.PushInt32(ds1.Version)
	sw.PushInt32(ds1.Width - 1)
	sw.PushInt32(ds1.Height - 1)

	if ds1.Version >= v8 {
		sw.PushInt32(ds1.Act - 1)
	}

	if ds1.Version >= v10 {
		sw.PushInt32(ds1.SubstitutionType)
	}

	if ds1.Version >= v3 {
		sw.PushInt32(int32(len(ds1.Files)))

		for _, i := range ds1.Files {
			sw.PushBytes([]byte(i)...)

			// separator
			sw.PushBytes(0)
		}
	}

	if ds1.Version >= v9 && ds1.Version <= v13 {
		sw.PushBytes(ds1.unknown1...)
	}

	if ds1.Version >= v4 {
		sw.PushInt32(ds1.NumberOfWalls)

		if ds1.Version >= v16 {
			sw.PushInt32(ds1.NumberOfFloors)
		}
	}

	// Step 2 - encode layers
	ds1.encodeLayers(sw)

	// Step 3 - encode objects
	if !(ds1.Version < v2) {
		sw.PushInt32(int32(len(ds1.Objects)))

		for _, i := range ds1.Objects {
			sw.PushUint32(uint32(i.Type))
			sw.PushUint32(uint32(i.ID))
			sw.PushUint32(uint32(i.X))
			sw.PushUint32(uint32(i.Y))
			sw.PushUint32(uint32(i.Flags))
		}
	}

	// Step 4 - encode substitutions
	if ds1.Version >= v12 && (ds1.SubstitutionType == subType1 || ds1.SubstitutionType == subType2) {
		sw.PushUint32(ds1.unknown2)

		sw.PushUint32(uint32(len(ds1.SubstitutionGroups)))

		for _, i := range ds1.SubstitutionGroups {
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
	for lIdx := range ds1.LayerStreamTypes {
		layerStreamType := ds1.LayerStreamTypes[lIdx]

		for y := 0; y < int(ds1.Height); y++ {
			for x := 0; x < int(ds1.Width); x++ {
				dw := uint32(0)

				switch layerStreamType {
				case d2enum.LayerStreamWall1, d2enum.LayerStreamWall2, d2enum.LayerStreamWall3, d2enum.LayerStreamWall4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamWall1)
					ds1.Tiles[y][x].Walls[wallIndex].Encode(sw)
				case d2enum.LayerStreamOrientation1, d2enum.LayerStreamOrientation2,
					d2enum.LayerStreamOrientation3, d2enum.LayerStreamOrientation4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamOrientation1)
					dw |= uint32(ds1.Tiles[y][x].Walls[wallIndex].Type)
					dw |= (uint32(ds1.Tiles[y][x].Walls[wallIndex].Zero) & wallZeroBitmask) << wallZeroOffset

					sw.PushUint32(dw)
				case d2enum.LayerStreamFloor1, d2enum.LayerStreamFloor2:
					floorIndex := int(layerStreamType) - int(d2enum.LayerStreamFloor1)
					ds1.Tiles[y][x].Floors[floorIndex].Encode(sw)
				case d2enum.LayerStreamShadow:
					ds1.Tiles[y][x].Shadows[0].Encode(sw)
				case d2enum.LayerStreamSubstitute:
					sw.PushUint32(ds1.Tiles[y][x].Substitutions[0].Unknown)
				}
			}
		}
	}
}

func (ds1 *DS1) encodeNPCs(sw *d2datautils.StreamWriter) {
	// Step 5.1 - encode npc's
	sw.PushUint32(uint32(len(ds1.NpcIndexes)))

	// Step 5.2 - enoce npcs' paths
	for _, i := range ds1.NpcIndexes {
		sw.PushUint32(uint32(len(ds1.Objects[i].Paths)))
		sw.PushUint32(uint32(ds1.Objects[i].X))
		sw.PushUint32(uint32(ds1.Objects[i].Y))

		for _, j := range ds1.Objects[i].Paths {
			sw.PushUint32(uint32(j.Position.X()))
			sw.PushUint32(uint32(j.Position.Y()))

			if ds1.Version >= v15 {
				sw.PushUint32(uint32(j.Action))
			}
		}
	}
}
