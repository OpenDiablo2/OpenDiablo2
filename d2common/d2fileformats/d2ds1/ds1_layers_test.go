package d2ds1

import (
	"testing"
)

func Test_ds1Layers_Delete(t *testing.T) {
	t.Run("Floors", func(t *testing.T) {
		ds1LayersDelete(t, floorLayerGroup)
	})
	t.Run("Walls", func(t *testing.T) {
		ds1LayersDelete(t, wallLayerGroup)
	})
	t.Run("Shadows", func(t *testing.T) {
		ds1LayersDelete(t, shadowLayerGroup)
	})
	t.Run("Substitution", func(t *testing.T) {
		ds1LayersDelete(t, substitutionLayerGroup)
	})
}

func ds1LayersDelete(t *testing.T, lt layerGroupType) {
	ds1 := DS1{}

	ds1.ds1Layers = &ds1Layers{
		width:         1,
		height:        1,
		Floors:        make(layerGroup, 1),
		Walls:         make(layerGroup, 1),
		Shadows:       make(layerGroup, 1),
		Substitutions: make(layerGroup, 1),
	}

	var lg layerGroup

	var del func(i int)

	switch lt {
	case floorLayerGroup:
		del = func(i int) { ds1.DeleteFloor(0) }
	case wallLayerGroup:
		del = func(i int) { ds1.DeleteWall(0) }
	case shadowLayerGroup:
		del = func(i int) { ds1.DeleteShadow(0) }
	case substitutionLayerGroup:
		del = func(i int) { ds1.DeleteSubstitution(0) }
	default:
		t.Fatal("unknown layer type given")
		return
	}

	del(0)

	if len(lg) > 0 {
		t.Errorf("unexpected layer present after deletion")
	}
}

func Test_ds1Layers_Get(t *testing.T) {
	t.Run("Floors", func(t *testing.T) {
		ds1LayersGet(t, floorLayerGroup)
	})
	t.Run("Walls", func(t *testing.T) {
		ds1LayersGet(t, wallLayerGroup)
	})
	t.Run("Shadows", func(t *testing.T) {
		ds1LayersGet(t, shadowLayerGroup)
	})
	t.Run("Substitution", func(t *testing.T) {
		ds1LayersGet(t, substitutionLayerGroup)
	})
}

func ds1LayersGet(t *testing.T, lt layerGroupType) {
	ds1 := exampleData()

	var get func(i int) *layer

	switch lt {
	case floorLayerGroup:
		get = func(i int) *layer { return ds1.GetFloor(0) }
	case wallLayerGroup:
		get = func(i int) *layer { return ds1.GetWall(0) }
	case shadowLayerGroup:
		get = func(i int) *layer { return ds1.GetShadow(0) }
	case substitutionLayerGroup:
		get = func(i int) *layer { return ds1.GetSubstitution(0) }
	default:
		t.Fatal("unknown layer type given")
		return
	}

	layer := get(0)

	// example has nil substitution layer, maybe we need another test
	if layer == nil && lt != substitutionLayerGroup {
		t.Errorf("layer expected")
	}
}

func Test_ds1Layers_Insert(t *testing.T) {
	t.Run("Floors", func(t *testing.T) {
		ds1LayersInsert(t, floorLayerGroup)
	})
	t.Run("Walls", func(t *testing.T) {
		ds1LayersInsert(t, wallLayerGroup)
	})
	t.Run("Shadows", func(t *testing.T) {
		ds1LayersInsert(t, shadowLayerGroup)
	})
	t.Run("Substitution", func(t *testing.T) {
		ds1LayersInsert(t, substitutionLayerGroup)
	})
}

func ds1LayersInsert(t *testing.T, lt layerGroupType) {
	ds1 := DS1{}

	layers := make([]*layer, getMaxGroupLen(lt)+1)

	for i := range layers {
		i := i
		layers[i] = &layer{}
		layers[i].tiles = make(tileGrid, 1)
		layers[i].tiles[0] = make(tileRow, 1)
		layers[i].SetSize(3, 3)
		layers[i].tiles[0][0].Prop1 = byte(i)
	}

	ds1.ds1Layers = &ds1Layers{}

	var insert func(i int)

	group := ds1.getLayersGroup(lt)

	switch lt {
	case floorLayerGroup:
		insert = func(i int) { ds1.InsertFloor(0, layers[i]) }
	case wallLayerGroup:
		insert = func(i int) { ds1.InsertWall(0, layers[i]) }
	case shadowLayerGroup:
		insert = func(i int) { ds1.InsertShadow(0, layers[i]) }
	case substitutionLayerGroup:
		insert = func(i int) { ds1.InsertSubstitution(0, layers[i]) }
	default:
		t.Fatal("unknown layer type given")
	}

	for i := range layers {
		insert(i)
	}

	if len(*group) != getMaxGroupLen(lt) {
		t.Fatal("unexpected floor len after setting")
	}

	idx := 0
	for i := len(layers) - 2; i > 0; i-- {
		if (*group)[idx].tiles[0][0].Prop1 != byte(i) {
			t.Fatal("unexpected tile inserted")
		}
		idx++
	}
}

func Test_ds1Layers_Pop(t *testing.T) {
	t.Run("Floor", func(t *testing.T) {
		ds1layerPop(floorLayerGroup, t)
	})

	t.Run("Wall", func(t *testing.T) {
		ds1layerPop(wallLayerGroup, t)
	})

	t.Run("Shadow", func(t *testing.T) {
		ds1layerPop(shadowLayerGroup, t)
	})

	t.Run("Substitution", func(t *testing.T) {
		ds1layerPop(substitutionLayerGroup, t)
	})
}

func ds1layerPop(lt layerGroupType, t *testing.T) {
	ds1 := exampleData()

	var pop func() *layer

	var numBefore, numAfter int

	switch lt {
	case floorLayerGroup:
		numBefore = len(ds1.Floors)
		pop = func() *layer {
			l := ds1.PopFloor()
			numAfter = len(ds1.Floors)

			return l
		}
	case wallLayerGroup:
		numBefore = len(ds1.Walls)
		pop = func() *layer {
			l := ds1.PopWall()
			numAfter = len(ds1.Walls)

			return l
		}
	case shadowLayerGroup:
		numBefore = len(ds1.Shadows)
		pop = func() *layer {
			l := ds1.PopShadow()
			numAfter = len(ds1.Shadows)

			return l
		}
	case substitutionLayerGroup:
		numBefore = len(ds1.Substitutions)
		pop = func() *layer {
			l := ds1.PopSubstitution()
			numAfter = len(ds1.Substitutions)

			return l
		}
	default:
		t.Fatal("unknown layer type given")
		return
	}

	attempts := 10

	for attempts > 0 {
		attempts--

		l := pop()

		if l == nil && numBefore < numAfter {
			t.Fatal("popped nil layer, expected layer count to remain the same")
		}

		if l != nil && numBefore <= numAfter {
			t.Fatal("popped non-nil, expected layer count to be lower")
		}
	}
}

func Test_ds1Layers_Push(t *testing.T) {
	t.Run("Floor", func(t *testing.T) {
		ds1layerPush(floorLayerGroup, t)
	})

	t.Run("Wall", func(t *testing.T) {
		ds1layerPush(wallLayerGroup, t)
	})

	t.Run("Shadow", func(t *testing.T) {
		ds1layerPush(shadowLayerGroup, t)
	})

	t.Run("Substitution", func(t *testing.T) {
		ds1layerPush(substitutionLayerGroup, t)
	})
}

// for all layer types, the test is the same
// when we push a layer, we expect an increment, and when we push a bunch of times,
// we expect to never exceed the max. we also expect to be able to retrieve a non-nil
// layer after we push.
func ds1layerPush(lt layerGroupType, t *testing.T) { //nolint:funlen // no biggie
	layers := &ds1Layers{}

	// we need to set up some shit to handle the test in a generic way
	var push func()

	var get func(idx int) *layer

	var max int

	var group *layerGroup

	check := func(expected int) {
		actual := len(*group)
		got := get(expected - 1)

		if actual != expected {
			t.Fatalf("unexpected number of layers: expected %d, got %d", expected, actual)
		}

		if got == nil {
			t.Fatal("got nil layer")
		}
	}

	switch lt {
	case floorLayerGroup:
		push = func() { layers.PushFloor(&layer{}) }
		get = layers.GetFloor
		max = maxFloorLayers
		group = &layers.Floors
	case wallLayerGroup:
		push = func() { layers.PushWall(&layer{}) }
		get = layers.GetWall
		max = maxWallLayers
		group = &layers.Walls
	case shadowLayerGroup:
		push = func() { layers.PushShadow(&layer{}) }
		get = layers.GetShadow
		max = maxShadowLayers
		group = &layers.Shadows
	case substitutionLayerGroup:
		push = func() { layers.PushSubstitution(&layer{}) }
		get = layers.GetSubstitution
		max = maxSubstitutionLayers
		group = &layers.Substitutions
	default:
		t.Fatal("unknown layer type given")
	}

	// push one time, we expect a single layer to exist
	push()
	check(1)

	// if we push a bunch of times, we expect to not exceed the max
	push()
	push()
	push()
	push()
	push()
	push()
	push()
	push()
	push()
	check(max)
}
