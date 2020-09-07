package d2ds1

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2path"
)

const maxActNumber = 5

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
	ds1.Version = br.GetInt32()
	ds1.Width = br.GetInt32() + 1
	ds1.Height = br.GetInt32() + 1

	if ds1.Version >= 8 { //nolint:gomnd // Version number
		ds1.Act = d2math.MinInt32(maxActNumber, br.GetInt32()+1)
	}

	if ds1.Version >= 10 { //nolint:gomnd // Version number
		ds1.SubstitutionType = br.GetInt32()
		if ds1.SubstitutionType == 1 || ds1.SubstitutionType == 2 {
			ds1.NumberOfSubstitutionLayers = 1
		}
	}

	if ds1.Version >= 3 { //nolint:gomnd // Version number
		// These files reference things that don't exist anymore :-?
		numberOfFiles := br.GetInt32()
		ds1.Files = make([]string, numberOfFiles)

		for i := 0; i < int(numberOfFiles); i++ {
			ds1.Files[i] = ""

			for {
				ch := br.GetByte()
				if ch == 0 {
					break
				}

				ds1.Files[i] += string(ch)
			}
		}
	}

	if ds1.Version >= 9 && ds1.Version <= 13 {
		// Skipping two dwords because they are "meaningless"?
		br.SkipBytes(8) //nolint:gomnd // We don't know what's here
	}

	if ds1.Version >= 4 { //nolint:gomnd // Version number
		ds1.NumberOfWalls = br.GetInt32()
		if ds1.Version >= 16 { //nolint:gomnd // Version number
			ds1.NumberOfFloors = br.GetInt32()
		} else {
			ds1.NumberOfFloors = 1
		}
	}

	layerStream := ds1.setupStreamLayerTypes()

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

	ds1.loadLayerStreams(br, layerStream)
	ds1.loadObjects(br)
	ds1.loadSubstitutions(br)
	ds1.loadNPCs(br)

	return ds1, nil
}

func (ds1 *DS1) loadObjects(br *d2datautils.StreamReader) {
	if ds1.Version >= 2 { //nolint:gomnd // Version number
		numberOfObjects := br.GetInt32()
		ds1.Objects = make([]Object, numberOfObjects)

		for objIdx := 0; objIdx < int(numberOfObjects); objIdx++ {
			newObject := Object{}
			newObject.Type = int(br.GetInt32())
			newObject.ID = int(br.GetInt32())
			newObject.X = int(br.GetInt32())
			newObject.Y = int(br.GetInt32())
			newObject.Flags = int(br.GetInt32())

			ds1.Objects[objIdx] = newObject
		}
	} else {
		ds1.Objects = make([]Object, 0)
	}
}

func (ds1 *DS1) loadSubstitutions(br *d2datautils.StreamReader) {
	if ds1.Version >= 12 && (ds1.SubstitutionType == 1 || ds1.SubstitutionType == 2) {
		if ds1.Version >= 18 { //nolint:gomnd // Version number
			br.GetUInt32()
		}

		numberOfSubGroups := br.GetInt32()
		ds1.SubstitutionGroups = make([]SubstitutionGroup, numberOfSubGroups)

		for subIdx := 0; subIdx < int(numberOfSubGroups); subIdx++ {
			newSub := SubstitutionGroup{}
			newSub.TileX = br.GetInt32()
			newSub.TileY = br.GetInt32()
			newSub.WidthInTiles = br.GetInt32()
			newSub.HeightInTiles = br.GetInt32()
			newSub.Unknown = br.GetInt32()

			ds1.SubstitutionGroups[subIdx] = newSub
		}
	} else {
		ds1.SubstitutionGroups = make([]SubstitutionGroup, 0)
	}
}

func (ds1 *DS1) setupStreamLayerTypes() []d2enum.LayerStreamType {
	var layerStream []d2enum.LayerStreamType

	if ds1.Version < 4 { //nolint:gomnd // Version number
		layerStream = []d2enum.LayerStreamType{
			d2enum.LayerStreamWall1,
			d2enum.LayerStreamFloor1,
			d2enum.LayerStreamOrientation1,
			d2enum.LayerStreamSubstitute,
			d2enum.LayerStreamShadow,
		}
	} else {
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

func (ds1 *DS1) loadNPCs(br *d2datautils.StreamReader) {
	if ds1.Version >= 14 { //nolint:gomnd // Version number
		numberOfNpcs := br.GetInt32()
		for npcIdx := 0; npcIdx < int(numberOfNpcs); npcIdx++ {
			numPaths := br.GetInt32()
			npcX := int(br.GetInt32())
			npcY := int(br.GetInt32())
			objIdx := -1

			for idx, ds1Obj := range ds1.Objects {
				if ds1Obj.X == npcX && ds1Obj.Y == npcY {
					objIdx = idx
					break
				}
			}

			if objIdx > -1 {
				ds1.loadNpcPaths(br, objIdx, int(numPaths))
			} else {
				if ds1.Version >= 15 { //nolint:gomnd // Version number
					br.SkipBytes(int(numPaths) * 3) //nolint:gomnd // Unknown data
				} else {
					br.SkipBytes(int(numPaths) * 2) //nolint:gomnd // Unknown data
				}
			}
		}
	}
}

func (ds1 *DS1) loadNpcPaths(br *d2datautils.StreamReader, objIdx, numPaths int) {
	if ds1.Objects[objIdx].Paths == nil {
		ds1.Objects[objIdx].Paths = make([]d2path.Path, numPaths)
	}

	for pathIdx := 0; pathIdx < numPaths; pathIdx++ {
		newPath := d2path.Path{}
		newPath.Position = d2vector.NewPosition(
			float64(br.GetInt32()),
			float64(br.GetInt32()))

		if ds1.Version >= 15 { //nolint:gomnd // Version number
			newPath.Action = int(br.GetInt32())
		}

		ds1.Objects[objIdx].Paths[pathIdx] = newPath
	}
}

func (ds1 *DS1) loadLayerStreams(br *d2datautils.StreamReader, layerStream []d2enum.LayerStreamType) {
	var dirLookup = []int32{
		0x00, 0x01, 0x02, 0x01, 0x02, 0x03, 0x03, 0x05, 0x05, 0x06,
		0x06, 0x07, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
		0x0F, 0x10, 0x11, 0x12, 0x14,
	}

	for lIdx := range layerStream {
		layerStreamType := layerStream[lIdx]

		for y := 0; y < int(ds1.Height); y++ {
			for x := 0; x < int(ds1.Width); x++ {
				dw := br.GetUInt32()

				switch layerStreamType {
				case d2enum.LayerStreamWall1, d2enum.LayerStreamWall2, d2enum.LayerStreamWall3, d2enum.LayerStreamWall4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamWall1)
					ds1.Tiles[y][x].Walls[wallIndex].Prop1 = byte(dw & 0x000000FF)            //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Walls[wallIndex].Sequence = byte((dw & 0x00003F00) >> 8)  //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Walls[wallIndex].Unknown1 = byte((dw & 0x000FC000) >> 14) //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Walls[wallIndex].Style = byte((dw & 0x03F00000) >> 20)    //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Walls[wallIndex].Unknown2 = byte((dw & 0x7C000000) >> 26) //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Walls[wallIndex].Hidden = byte((dw&0x80000000)>>31) > 0   //nolint:gomnd // Bitmask
				case d2enum.LayerStreamOrientation1, d2enum.LayerStreamOrientation2,
					d2enum.LayerStreamOrientation3, d2enum.LayerStreamOrientation4:
					wallIndex := int(layerStreamType) - int(d2enum.LayerStreamOrientation1)
					c := int32(dw & 0x000000FF) //nolint:gomnd // Bitmask

					if ds1.Version < 7 { //nolint:gomnd // Version number
						if c < int32(len(dirLookup)) {
							c = dirLookup[c]
						}
					}

					ds1.Tiles[y][x].Walls[wallIndex].Type = d2enum.TileType(c)
					ds1.Tiles[y][x].Walls[wallIndex].Zero = byte((dw & 0xFFFFFF00) >> 8) //nolint:gomnd // Bitmask
				case d2enum.LayerStreamFloor1, d2enum.LayerStreamFloor2:
					floorIndex := int(layerStreamType) - int(d2enum.LayerStreamFloor1)
					ds1.Tiles[y][x].Floors[floorIndex].Prop1 = byte(dw & 0x000000FF)            //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Floors[floorIndex].Sequence = byte((dw & 0x00003F00) >> 8)  //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Floors[floorIndex].Unknown1 = byte((dw & 0x000FC000) >> 14) //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Floors[floorIndex].Style = byte((dw & 0x03F00000) >> 20)    //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Floors[floorIndex].Unknown2 = byte((dw & 0x7C000000) >> 26) //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Floors[floorIndex].Hidden = byte((dw&0x80000000)>>31) > 0   //nolint:gomnd // Bitmask
				case d2enum.LayerStreamShadow:
					ds1.Tiles[y][x].Shadows[0].Prop1 = byte(dw & 0x000000FF)            //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Shadows[0].Sequence = byte((dw & 0x00003F00) >> 8)  //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Shadows[0].Unknown1 = byte((dw & 0x000FC000) >> 14) //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Shadows[0].Style = byte((dw & 0x03F00000) >> 20)    //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Shadows[0].Unknown2 = byte((dw & 0x7C000000) >> 26) //nolint:gomnd // Bitmask
					ds1.Tiles[y][x].Shadows[0].Hidden = byte((dw&0x80000000)>>31) > 0   //nolint:gomnd // Bitmask
				case d2enum.LayerStreamSubstitute:
					ds1.Tiles[y][x].Substitutions[0].Unknown = dw
				}
			}
		}
	}
}
