package d2ds1

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2path"
)

func exampleData() *DS1 { //nolint:funlen // not a big deal if this is long func
	exampleFloor1 := Tile{
		// common fields
		tileCommonFields: tileCommonFields{
			Prop1:       2,
			Sequence:    89,
			Unknown1:    123,
			Style:       20,
			Unknown2:    53,
			HiddenBytes: 1,
			RandomIndex: 2,
			YAdjust:     21,
		},
		tileFloorShadowFields: tileFloorShadowFields{
			Animated: false,
		},
	}

	exampleFloor2 := Tile{
		// common fields
		tileCommonFields: tileCommonFields{
			Prop1:       3,
			Sequence:    89,
			Unknown1:    213,
			Style:       28,
			Unknown2:    53,
			HiddenBytes: 7,
			RandomIndex: 3,
			YAdjust:     28,
		},
		tileFloorShadowFields: tileFloorShadowFields{
			Animated: true,
		},
	}

	exampleWall1 := Tile{
		// common fields
		tileCommonFields: tileCommonFields{
			Prop1:       3,
			Sequence:    89,
			Unknown1:    213,
			Style:       28,
			Unknown2:    53,
			HiddenBytes: 7,
			RandomIndex: 3,
			YAdjust:     28,
		},
		tileWallFields: tileWallFields{
			Type: d2enum.TileRightWall,
		},
	}

	exampleWall2 := Tile{
		// common fields
		tileCommonFields: tileCommonFields{
			Prop1:       3,
			Sequence:    93,
			Unknown1:    193,
			Style:       17,
			Unknown2:    13,
			HiddenBytes: 1,
			RandomIndex: 1,
			YAdjust:     22,
		},
		tileWallFields: tileWallFields{
			Type: d2enum.TileLeftWall,
		},
	}

	exampleShadow := Tile{
		// common fields
		tileCommonFields: tileCommonFields{
			Prop1:       3,
			Sequence:    93,
			Unknown1:    173,
			Style:       17,
			Unknown2:    12,
			HiddenBytes: 1,
			RandomIndex: 1,
			YAdjust:     22,
		},
		tileFloorShadowFields: tileFloorShadowFields{
			Animated: false,
		},
	}

	result := &DS1{
		ds1Layers: &ds1Layers{
			width:  2,
			height: 2,
			Floors: layerGroup{
				// number of floors (one floor)
				{
					// tile grid = []tileRow
					tiles: tileGrid{
						// tile rows = []Tile
						// 2x2 tiles
						{
							exampleFloor1,
							exampleFloor2,
						},
						{
							exampleFloor2,
							exampleFloor1,
						},
					},
				},
			},
			Walls: layerGroup{
				// number of walls (two floors)
				{
					// tile grid = []tileRow
					tiles: tileGrid{
						// tile rows = []Tile
						// 2x2 tiles
						{
							exampleWall1,
							exampleWall2,
						},
						{
							exampleWall2,
							exampleWall1,
						},
					},
				},
				{
					// tile grid = []tileRow
					tiles: tileGrid{
						// tile rows = []Tile
						// 2x2 tiles
						{
							exampleWall1,
							exampleWall2,
						},
						{
							exampleWall2,
							exampleWall1,
						},
					},
				},
			},
			Shadows: layerGroup{
				// number of shadows (always 1)
				{
					// tile grid = []tileRow
					tiles: tileGrid{
						// tile rows = []Tile
						// 2x2 tiles
						{
							exampleShadow,
							exampleShadow,
						},
						{
							exampleShadow,
							exampleShadow,
						},
					},
				},
			},
		},
		Files: []string{"a.dt1", "bfile.dt1"},
		Objects: []Object{
			{0, 0, 0, 0, 0, nil},
			{0, 1, 0, 0, 0, []d2path.Path{{}}},
			{0, 2, 0, 0, 0, nil},
			{0, 3, 0, 0, 0, nil},
		},
		SubstitutionGroups: nil,
		version:            17,
		Act:                1,
		SubstitutionType:   0,
		unknown2:           20,
	}

	return result
}

func TestDS1_MarshalUnmarshal(t *testing.T) {
	ds1 := exampleData()

	data := ds1.Marshal()

	_, loadErr := Unmarshal(data)
	if loadErr != nil {
		t.Error(loadErr)
	}
}

func TestDS1_Version(t *testing.T) {
	ds1 := exampleData()

	v := ds1.Version()

	ds1.SetVersion(v + 1)

	if ds1.Version() == v {
		t.Fatal("expected different ds1 version")
	}
}

func TestDS1_SetSize(t *testing.T) {
	ds1 := exampleData()

	w, h := ds1.Size()

	ds1.SetSize(w+1, h-1)

	w2, h2 := ds1.Size()

	if w2 != (w+1) || h2 != (h-1) {
		t.Fatal("unexpected width/height after setting size")
	}
}

func Test_getLayerSchema(t *testing.T) {
	ds1 := exampleData()

	expected := map[int]layerStreamType{
		0: layerStreamWall1,
		1: layerStreamOrientation1,
		2: layerStreamWall2,
		3: layerStreamOrientation2,
		4: layerStreamFloor1,
		5: layerStreamShadow1,
	}

	schema := ds1.getLayerSchema()

	if len(schema) != len(expected) {
		t.Fatal("unexpected schema length")
	}

	for idx := range expected {
		if schema[idx] != expected[idx] {
			t.Fatal("unexpected layer type in schema")
		}
	}
}
