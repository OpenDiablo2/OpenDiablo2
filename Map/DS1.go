package Map

import (
	"github.com/essial/OpenDiablo2/Common"
)

var dirLookup = []int32{
	0x00, 0x01, 0x02, 0x01, 0x02, 0x03, 0x03, 0x05, 0x05, 0x06,
	0x06, 0x07, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
	0x0F, 0x10, 0x11, 0x12, 0x14,
}

type LayerStreamType int

const (
	LayerStreamWall1        LayerStreamType = 0
	LayerStreamWall2        LayerStreamType = 1
	LayerStreamWall3        LayerStreamType = 2
	LayerStreamWall4        LayerStreamType = 3
	LayerStreamOrientation1 LayerStreamType = 4
	LayerStreamOrientation2 LayerStreamType = 5
	LayerStreamOrientation3 LayerStreamType = 6
	LayerStreamOrientation4 LayerStreamType = 7
	LayerStreamFloor1       LayerStreamType = 8
	LayerStreamFloor2       LayerStreamType = 9
	LayerStreamShadow       LayerStreamType = 10
	LayerStreamSubstitute   LayerStreamType = 11
)

type FloorShadowRecord struct {
	Prop1     byte
	SubIndex  byte
	Unknown1  byte
	MainIndex byte
	Unknown2  byte
	Hidden    bool
}

type WallRecord struct {
	Orientation byte
	Zero        byte
	Prop1       byte
	SubIndex    byte
	Unknown1    byte
	MainIndex   byte
	Unknown2    byte
	Hidden      bool
}

type SubstitutionRecord struct {
	Unknown uint32
}

type TileRecord struct {
	Floors        []FloorShadowRecord
	Walls         []WallRecord
	Shadows       []FloorShadowRecord
	Substitutions []SubstitutionRecord
}

type SubstitutionGroup struct {
	TileX         int32
	TileY         int32
	WidthInTiles  int32
	HeightInTiles int32
	Unknown       int32
}

type Path struct {
	X      int32
	Y      int32
	Action int32
}

type Object struct {
	Type  int32
	Id    int32
	X     int32
	Y     int32
	Flags int32
	Paths []Path
}

type DS1 struct {
	Version                    int32    // The version of the DS1
	Width                      int32    // Width of map, in # of tiles
	Height                     int32    // Height of map, in # of tiles
	Act                        int32    // Act, from 1 to 5. This tells which act table to use for the Objects list
	SubstitutionType           int32    // SubstitutionType (layer type): 0 if no layer, else type 1 or type 2
	Files                      []string // FilePtr table of file string pointers
	NumberOfWalls              int32    // WallNum number of wall & orientation layers used
	NumberOfFloors             int32    // number of floor layers used
	NumberOfShadowLayers       int32    // ShadowNum number of shadow layer used
	NumberOfSubstitutionLayers int32    // SubstitutionNum number of substitution layer used
	SubstitutionGroupsNum      int32    // SubstitutionGroupsNum number of substitution groups, datas between objects & NPC paths
	Objects                    []Object // Objects
	Tiles                      [][]TileRecord
	SubstitutionGroups         []SubstitutionGroup
}

func LoadDS1(path string, fileProvider Common.FileProvider) *DS1 {
	ds1 := &DS1{
		NumberOfFloors:             1,
		NumberOfWalls:              1,
		NumberOfShadowLayers:       1,
		NumberOfSubstitutionLayers: 0,
	}
	fileData := fileProvider.LoadFile(path)
	br := Common.CreateStreamReader(fileData)
	ds1.Version = br.GetInt32()
	ds1.Width = br.GetInt32() + 1
	ds1.Height = br.GetInt32() + 1
	if ds1.Version >= 8 {
		ds1.Act = Common.MinInt32(5, br.GetInt32()+1)
	}
	if ds1.Version >= 10 {
		ds1.SubstitutionType = br.GetInt32()
		if ds1.SubstitutionType == 1 || ds1.SubstitutionType == 2 {
			ds1.NumberOfSubstitutionLayers = 1
		}
	}
	if ds1.Version >= 3 {
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
		br.SkipBytes(16)
	}
	if ds1.Version >= 4 {
		ds1.NumberOfWalls = br.GetInt32()
		if ds1.Version >= 16 {
			ds1.NumberOfFloors = br.GetInt32()
		} else {
			ds1.NumberOfFloors = 1
		}
	}
	var layerStream []LayerStreamType
	if ds1.Version < 4 {
		layerStream = []LayerStreamType{
			LayerStreamWall1,
			LayerStreamFloor1,
			LayerStreamOrientation1,
			LayerStreamSubstitute,
			LayerStreamShadow,
		}
	} else {
		layerStream = make([]LayerStreamType, 0)
		for i := 0; i < int(ds1.NumberOfWalls); i++ {
			layerStream = append(layerStream, LayerStreamType(int(LayerStreamWall1)+i))
			layerStream = append(layerStream, LayerStreamType(int(LayerStreamOrientation1)+i))
		}
		for i := 0; i < int(ds1.NumberOfFloors); i++ {
			layerStream = append(layerStream, LayerStreamType(int(LayerStreamFloor1)+i))
		}
		if ds1.NumberOfShadowLayers > 0 {
			layerStream = append(layerStream, LayerStreamShadow)
		}
		if ds1.NumberOfSubstitutionLayers > 0 {
			layerStream = append(layerStream, LayerStreamSubstitute)
		}
	}
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
	for _, layerStreamType := range layerStream {
		for y := 0; y < int(ds1.Height); y++ {
			for x := 0; x < int(ds1.Width); x++ {
				dw := br.GetUInt32()
				switch layerStreamType {
				case LayerStreamWall1:
					fallthrough
				case LayerStreamWall2:
					fallthrough
				case LayerStreamWall3:
					fallthrough
				case LayerStreamWall4:
					wallIndex := int(layerStreamType) - int(LayerStreamWall1)
					ds1.Tiles[y][x].Walls[wallIndex].Prop1 = byte(dw & 0x000000FF)
					ds1.Tiles[y][x].Walls[wallIndex].SubIndex = byte((dw & 0x00003F00) >> 8)
					ds1.Tiles[y][x].Walls[wallIndex].Unknown1 = byte((dw & 0x000FC000) >> 14)
					ds1.Tiles[y][x].Walls[wallIndex].MainIndex = byte((dw & 0x03F00000) >> 20)
					ds1.Tiles[y][x].Walls[wallIndex].Unknown2 = byte((dw & 0x7C000000) >> 26)
					ds1.Tiles[y][x].Walls[wallIndex].Hidden = byte((dw&0x80000000)>>31) > 0
				case LayerStreamOrientation1:
					fallthrough
				case LayerStreamOrientation2:
					fallthrough
				case LayerStreamOrientation3:
					fallthrough
				case LayerStreamOrientation4:
					wallIndex := int(layerStreamType) - int(LayerStreamOrientation1)
					c := int32(dw & 0x000000FF)
					if ds1.Version < 7 {
						if c < 25 {
							c = dirLookup[c]
						}
					}
					ds1.Tiles[y][x].Walls[wallIndex].Orientation = byte(c)
					ds1.Tiles[y][x].Walls[wallIndex].Zero = byte((dw & 0xFFFFFF00) >> 8)
				case LayerStreamFloor1:
					fallthrough
				case LayerStreamFloor2:
					floorIndex := int(layerStreamType) - int(LayerStreamFloor1)
					ds1.Tiles[y][x].Floors[floorIndex].Prop1 = byte(dw & 0x000000FF)
					ds1.Tiles[y][x].Floors[floorIndex].SubIndex = byte((dw & 0x00003F00) >> 8)
					ds1.Tiles[y][x].Floors[floorIndex].Unknown1 = byte((dw & 0x000FC000) >> 14)
					ds1.Tiles[y][x].Floors[floorIndex].MainIndex = byte((dw & 0x03F00000) >> 20)
					ds1.Tiles[y][x].Floors[floorIndex].Unknown2 = byte((dw & 0x7C000000) >> 26)
					ds1.Tiles[y][x].Floors[floorIndex].Hidden = byte((dw&0x80000000)>>31) > 0
				case LayerStreamShadow:
					ds1.Tiles[y][x].Shadows[0].Prop1 = byte(dw & 0x000000FF)
					ds1.Tiles[y][x].Shadows[0].SubIndex = byte((dw & 0x00003F00) >> 8)
					ds1.Tiles[y][x].Shadows[0].Unknown1 = byte((dw & 0x000FC000) >> 14)
					ds1.Tiles[y][x].Shadows[0].MainIndex = byte((dw & 0x03F00000) >> 20)
					ds1.Tiles[y][x].Shadows[0].Unknown2 = byte((dw & 0x7C000000) >> 26)
					ds1.Tiles[y][x].Shadows[0].Hidden = byte((dw&0x80000000)>>31) > 0
				case LayerStreamSubstitute:
					ds1.Tiles[y][x].Substitutions[0].Unknown = dw
				}
			}
		}
	}
	ds1.Objects = make([]Object, 0)
	if ds1.Version >= 2 {
		numberOfObjects := br.GetInt32()
		for objIdx := 0; objIdx < int(numberOfObjects); objIdx++ {
			newObject := Object{}
			newObject.Type = br.GetInt32()
			newObject.Id = br.GetInt32()
			newObject.X = br.GetInt32()
			newObject.Y = br.GetInt32()
			newObject.Flags = br.GetInt32()
			ds1.Objects = append(ds1.Objects, newObject)
		}
	}
	ds1.SubstitutionGroups = make([]SubstitutionGroup, 0)
	if ds1.Version >= 12 && (ds1.SubstitutionType == 1 || ds1.SubstitutionType == 2) {
		if ds1.Version >= 18 {
			br.GetUInt32()
		}
		numberOfSubGroups := br.GetInt32()
		for subIdx := 0; subIdx < int(numberOfSubGroups); subIdx++ {
			newSub := SubstitutionGroup{}
			newSub.TileX = br.GetInt32()
			newSub.TileY = br.GetInt32()
			newSub.WidthInTiles = br.GetInt32()
			newSub.HeightInTiles = br.GetInt32()
			newSub.Unknown = br.GetInt32()

			ds1.SubstitutionGroups = append(ds1.SubstitutionGroups, newSub)
		}
	}
	if ds1.Version >= 14 {
		numberOfNpcs := br.GetInt32()
		for npcIdx := 0; npcIdx < int(numberOfNpcs); npcIdx++ {
			numPaths := br.GetInt32()
			npcX := br.GetInt32()
			npcY := br.GetInt32()
			objIdx := -1
			for idx, ds1Obj := range ds1.Objects {
				if ds1Obj.X == npcX && ds1Obj.Y == npcY {
					objIdx = idx
					break
				}
			}
			if objIdx > -1 {
				if ds1.Objects[objIdx].Paths == nil {
					ds1.Objects[objIdx].Paths = make([]Path, numPaths)
				}
				for pathIdx := 0; pathIdx < int(numPaths); pathIdx++ {
					newPath := Path{}
					newPath.X = br.GetInt32()
					newPath.Y = br.GetInt32()
					if ds1.Version >= 15 {
						newPath.Action = br.GetInt32()
					}
					ds1.Objects[objIdx].Paths[pathIdx] = newPath
				}
			} else {
				if ds1.Version >= 15 {
					br.SkipBytes(int(numPaths) * 3)
				} else {
					br.SkipBytes(int(numPaths) * 2)
				}
			}
		}
	}
	return ds1
}
