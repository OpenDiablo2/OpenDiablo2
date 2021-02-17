package d2ds1

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"testing"
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
		layerStreamTypes:           []d2enum.LayerStreamType{
			d2enum.LayerStreamWall1,
			d2enum.LayerStreamOrientation1,
			d2enum.LayerStreamFloor1,
			d2enum.LayerStreamShadow,
		},
		npcIndexes:                 []int{},
	}
}

func TestDS1_Act(t *testing.T) {
	ds1 := exampleDS1()

	if ds1.Act() != int(ds1.act) {
		t.Error("unexpected value in example ds1")
	}
}

func TestDS1_AddFile(t *testing.T) {
	ds1 := exampleDS1()

	numBefore := len(ds1.files)

	ds1.AddFile("other.ds1")

	numAfter := len(ds1.files)

	if (numBefore+1) != numAfter {
		t.Error("unexpected number of files in ds1")
	}
}

func TestDS1_AddObject(t *testing.T) {
	ds1 := exampleDS1()

	numBefore := len(ds1.objects)

	ds1.AddObject(Object{})

	numAfter := len(ds1.objects)

	if (numBefore+1) != numAfter {
		t.Error("unexpected number of objects in ds1")
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

func TestDS1_Height(t *testing.T) {
	ds1 := exampleDS1()

	if int(ds1.height) != ds1.Height(){
		t.Error("unexpected height")
	}
}

func TestDS1_Marshal(t *testing.T) {
	a := exampleDS1()

	bytes := a.Marshal()

	b, err := LoadDS1(bytes)
	if err != nil {
		t.Error("could not load new ds1 from marshalled ds1 data")
		return
	}

	if b.width != a.width {
		t.Error("new ds1 does not match original")
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

func TestDS1_NumberOfWalls(t *testing.T) {
	ds1 := exampleDS1()

	if ds1.NumberOfWallLayers() != int(ds1.numberOfWallLayers) {
		t.Error("unexpected number of wall layers")
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

func TestDS1_RemoveFile(t *testing.T) {
	ds1 := exampleDS1()

	numBefore := len(ds1.files)

	ds1.RemoveFile("nonexistant file")

	if len(ds1.files) != numBefore {
		t.Error("file removed when it should not have been")
	}

	filename := "c.ds1"

	ds1.AddFile(filename)

	if len(ds1.files) == numBefore {
		t.Error("file not added when it should have been")
	}

	ds1.RemoveFile(filename)

	if len(ds1.files) != numBefore {
		t.Error("file not removed when it should have been")
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
}

func TestDS1_SetHeight(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_SetSize(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_SetSubstitutionGroups(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_SetSubstitutionType(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_SetTile(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_SetTiles(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_SetVersion(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_SetWidth(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_Size(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_SubstitutionGroups(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_SubstitutionType(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_Tile(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_Tiles(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_Version(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestDS1_Width(t *testing.T) {
	//ds1 := exampleDS1()
}

func TestLoadDS1(t *testing.T) {
	//ds1 := exampleDS1()
}
