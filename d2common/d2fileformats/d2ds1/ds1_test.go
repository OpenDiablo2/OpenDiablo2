package d2ds1

import (
	"fmt"

	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func exampleDS1() *DS1 {
	return &DS1{
		files: []string{"a.dt1", "b.dt1"},
		objects: []Object{
			{0, 0, 0, 0, 0, nil},
			{0, 1, 0, 0, 0, nil},
			{0, 2, 0, 0, 0, nil},
			{0, 3, 0, 0, 0, nil},
		},
		tiles: [][]Tile{ // 2x2
			{
				Tile{[]Floor{{}}, []Wall{{}}, []Shadow{{}}, []Substitution{}},
				Tile{[]Floor{{}}, []Wall{{}}, []Shadow{{}}, []Substitution{}},
			},
			{
				Tile{[]Floor{{}}, []Wall{{}}, []Shadow{{}}, []Substitution{}},
				Tile{[]Floor{{}}, []Wall{{}}, []Shadow{{}}, []Substitution{}},
			},
		},
		substitutionGroups:         nil,
		version:                    17,
		width:                      2,
		height:                     2,
		act:                        1,
		substitutionType:           0,
		numberOfWallLayers:         1,
		numberOfFloorLayers:        1,
		numberOfShadowLayers:       1,
		numberOfSubstitutionLayers: 1,
		layerStreamTypes: []d2enum.LayerStreamType{
			d2enum.LayerStreamWall1,
			d2enum.LayerStreamOrientation1,
			d2enum.LayerStreamFloor1,
			d2enum.LayerStreamShadow,
		},
		npcIndexes: []int{},
	}
}

// checks, if DS1 structure could be marshaled and unmarshaled
func testIfRestorable(ds1 *DS1) error {
	var err error

	data := ds1.Marshal()
	_, err = LoadDS1(data)

	return err
}

func TestDS1_Marshal(t *testing.T) {
	a := exampleDS1()

	bytes := a.Marshal()

	b, err := LoadDS1(bytes)
	if err != nil {
		t.Error("could not load new ds1 from marshaled ds1 data")
		return
	}

	if b.width != a.width {
		t.Error("new ds1 does not match original")
	}
}

func TestDS1_Files(t *testing.T) {
	ds1 := exampleDS1()

	files := ds1.Files()

	for idx := range files {
		if ds1.files[idx] != files[idx] {
			t.Error("unexpected files from ds1")
		}
	}
}

func TestDS1_AddFile(t *testing.T) {
	ds1 := exampleDS1()

	numBefore := len(ds1.files)

	ds1.AddFile("other.ds1")

	numAfter := len(ds1.files)

	if (numBefore + 1) != numAfter {
		t.Error("unexpected number of files in ds1")
	}

	if err := testIfRestorable(ds1); err != nil {
		t.Errorf("unable to restore: %v", err)
	}
}

func TestDS1_RemoveFile(t *testing.T) {
	ds1 := exampleDS1()

	numBefore := len(ds1.files)

	err := ds1.RemoveFile("nonexistant file")
	if err == nil {
		t.Fatal("file 'nonexistant file' doesn't exist but ds1.RemoveFile doesn't return error")
	}

	if len(ds1.files) != numBefore {
		t.Error("file removed when it should not have been")
	}

	filename := "c.ds1"

	ds1.AddFile(filename)

	if len(ds1.files) == numBefore {
		t.Error("file not added when it should have been")
	}

	err = ds1.RemoveFile(filename)
	if err != nil {
		t.Error(err)
	}

	if len(ds1.files) != numBefore {
		t.Error("file not removed when it should have been")
	}

	if err := testIfRestorable(ds1); err != nil {
		t.Errorf("unable to restore: %v", err)
	}
}

func TestDS1_Objects(t *testing.T) {
	ds1 := exampleDS1()

	objects := ds1.Objects()

	for idx := range ds1.objects {
		if !ds1.objects[idx].Equals(&objects[idx]) {
			t.Error("unexpected object")
		}
	}
}

func TestDS1_AddObject(t *testing.T) {
	ds1 := exampleDS1()

	numBefore := len(ds1.objects)

	ds1.AddObject(Object{})

	numAfter := len(ds1.objects)

	if (numBefore + 1) != numAfter {
		t.Error("unexpected number of objects in ds1")
	}

	if err := testIfRestorable(ds1); err != nil {
		t.Errorf("unable to restore: %v", err)
	}
}

func TestDS1_RemoveObject(t *testing.T) {
	ds1 := exampleDS1()

	nice := 69420

	obj := Object{
		ID: nice,
	}

	ds1.AddObject(obj)

	numBefore := len(ds1.objects)

	ds1.RemoveObject(obj)

	if len(ds1.objects) == numBefore {
		t.Error("did not remove object when expected")
	}

	if err := testIfRestorable(ds1); err != nil {
		t.Errorf("unable to restore: %v", err)
	}
}

func TestDS1_Tiles(t *testing.T) {
	ds1 := exampleDS1()

	tiles := ds1.Tiles()

	for y := range tiles {
		for x := range tiles[y] {
			if len(ds1.tiles[y][x].Floors) != len(tiles[y][x].Floors) {
				t.Fatal("number of tile's floors returned by ds1.Tiles() isn't same as real number")
			}

			if ds1.tiles[y][x].Walls[0] != tiles[y][x].Walls[0] {
				t.Fatal("wall record returned by ds1.Tiles isn't equal to real")
			}
		}
	}
}

func TestDS1_SetTiles(t *testing.T) {
	ds1 := exampleDS1()

	exampleTile1 := Tile{
		Floors: []FloorShadow{
			{0, 0, 2, 3, 4, 55, 33, true, 999},
		},
		Shadows: []FloorShadow{
			{2, 4, 5, 33, 6, 7, 0, false, 1024},
		},
	}

	exampleTile2 := Tile{
		Walls: []Wall{
			{2, 3, 4, 5, 3, 2, 3, 0, 33, 99},
		},
		Shadows: []FloorShadow{
			{2, 4, 5, 33, 6, 7, 0, false, 1024},
		},
	}

	tiles := [][]Tile{{exampleTile1, exampleTile2}}

	ds1.SetTiles(tiles)

	if ds1.tiles[0][0].Floors[0] != exampleTile1.Floors[0] {
		t.Fatal("unexpected tile was set")
	}

	if len(ds1.tiles[0][0].Walls) != len(exampleTile1.Walls) {
		t.Fatal("unexpected tile was set")
	}

	if ds1.tiles[0][1].Walls[0] != exampleTile2.Walls[0] {
		t.Fatal("unexpected tile was set")
	}

	if len(ds1.tiles[0][1].Walls) != len(exampleTile2.Walls) {
		t.Fatal("unexpected tile was set")
	}
}

func TestDS1_Tile(t *testing.T) {
	ds1 := exampleDS1()

	x, y := 1, 0

	if ds1.tiles[y][x].Floors[0] != ds1.Tile(x, y).Floors[0] {
		t.Fatal("ds1.Tile returned invalid value")
	}
}

func TestDS1_SetTile(t *testing.T) {
	ds1 := exampleDS1()

	exampleTile := Tile{
		Floors: []FloorShadow{
			{5, 8, 9, 4, 3, 4, 2, true, 1024},
			{8, 22, 7, 9, 6, 3, 0, false, 1024},
		},
		Walls: []Wall{
			{2, 3, 4, 5, 3, 2, 3, 0, 33, 99},
		},
		Shadows: []FloorShadow{
			{2, 44, 99, 2, 4, 3, 2, true, 933},
		},
	}

	ds1.SetTile(0, 0, &exampleTile)

	if ds1.tiles[0][0].Floors[0] != exampleTile.Floors[0] {
		t.Fatal("unexpected tile was set")
	}

	if len(ds1.tiles[0][0].Walls) != len(exampleTile.Walls) {
		t.Fatal("unexpected tile was set")
	}

	if err := testIfRestorable(ds1); err != nil {
		t.Errorf("unable to restore: %v", err)
	}
}

func TestDS1_Version(t *testing.T) {
	ds1 := exampleDS1()

	version := ds1.version

	if version != int32(ds1.Version()) {
		t.Fatal("version returned by ds1.Version() and real aren't equal")
	}
}

func TestDS1_SetVersion(t *testing.T) {
	ds1 := exampleDS1()

	newVersion := 8

	ds1.SetVersion(newVersion)

	if newVersion != int(ds1.version) {
		t.Fatal("ds1.SetVersion set version incorrectly")
	}

	if err := testIfRestorable(ds1); err != nil {
		t.Errorf("unable to restore: %v", err)
	}
}

func TestDS1_Width(t *testing.T) {
	ds1 := exampleDS1()

	if int(ds1.width) != ds1.Width() {
		t.Error("unexpected width")
	}
}

func TestDS1_SetWidth(t *testing.T) {
	ds1 := exampleDS1()

	var newWidth int32 = 4

	ds1.SetWidth(int(newWidth))

	if newWidth != ds1.width {
		t.Fatal("unexpected width after set")
	}

	if err := testIfRestorable(ds1); err != nil {
		t.Errorf("unable to restore: %v", err)
	}
}

func TestDS1_Height(t *testing.T) {
	ds1 := exampleDS1()

	if int(ds1.height) != ds1.Height() {
		t.Error("unexpected height")
	}
}

func TestDS1_SetHeight(t *testing.T) {
	ds1 := exampleDS1()

	var newHeight int32 = 5

	ds1.SetHeight(int(newHeight))

	if newHeight != ds1.height {
		fmt.Println(newHeight, ds1.height)
		t.Fatal("unexpected heigth after set")
	}

	if err := testIfRestorable(ds1); err != nil {
		t.Errorf("unable to restore: %v", err)
	}
}

func TestDS1_Act(t *testing.T) {
	ds1 := exampleDS1()

	if ds1.Act() != int(ds1.act) {
		t.Error("unexpected value in example ds1")
	}
}

func TestDS1_SetAct(t *testing.T) {
	ds1 := exampleDS1()

	ds1.SetAct(-1)

	if ds1.Act() < 0 {
		t.Error("act cannot be less than 0")
	}

	nice := 69420

	ds1.SetAct(nice)

	if int(ds1.act) != nice {
		t.Error("unexpected value for act")
	}

	if err := testIfRestorable(ds1); err != nil {
		t.Errorf("unable to restore: %v", err)
	}
}

func TestDS1_SubstitutionType(t *testing.T) {
	ds1 := exampleDS1()

	st := ds1.substitutionType

	if int(st) != ds1.SubstitutionType() {
		t.Fatal("unexpected substitution type returned")
	}
}

func TestDS1_SetSubstitutionType(t *testing.T) {
	ds1 := exampleDS1()

	newST := 5

	ds1.SetSubstitutionType(newST)

	if ds1.substitutionType != int32(newST) {
		t.Fatal("unexpected substitutionType was set")
	}
}

func TestDS1_NumberOfWalls(t *testing.T) {
	ds1 := exampleDS1()

	if ds1.NumberOfWallLayers() != int(ds1.numberOfWallLayers) {
		t.Error("unexpected number of wall layers")
	}
}

func TestDS1_NumberOfFloors(t *testing.T) {
	ds1 := exampleDS1()

	if ds1.NumberOfFloorLayers() != int(ds1.numberOfFloorLayers) {
		t.Error("unexpected number of floor layers")
	}
}

func TestDS1_NumberOfShadowLayers(t *testing.T) {
	ds1 := exampleDS1()

	if ds1.NumberOfShadowLayers() != int(ds1.numberOfShadowLayers) {
		t.Error("unexpected number of shadow layers")
	}
}

func TestDS1_NumberOfSubstitutionLayers(t *testing.T) {
	ds1 := exampleDS1()

	if ds1.NumberOfSubstitutionLayers() != int(ds1.numberOfSubstitutionLayers) {
		t.Error("unexpected number of substitution layers")
	}
}

func TestDS1_SubstitutionGroups(t *testing.T) {
	ds1 := exampleDS1()

	sg := ds1.SubstitutionGroups()

	for i := 0; i < len(ds1.substitutionGroups); i++ {
		if sg[i] != ds1.substitutionGroups[i] {
			t.Fatal("unexpected substitution group returned")
		}
	}
}

func TestDS1_SetSubstitutionGroups(t *testing.T) {
	ds1 := exampleDS1()

	newGroup := []SubstitutionGroup{
		{
			TileX:         20,
			TileY:         12,
			WidthInTiles:  212,
			HeightInTiles: 334,
			Unknown:       1024,
		},
	}

	ds1.SetSubstitutionGroups(newGroup)

	if ds1.substitutionGroups[0] != newGroup[0] {
		t.Fatal("unexpected substitution group added")
	}
}
