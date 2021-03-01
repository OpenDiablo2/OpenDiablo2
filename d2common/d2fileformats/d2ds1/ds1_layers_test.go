package d2ds1

import (
	"testing"
)

func Test_ds1Layers_DeleteFloor(t *testing.T) {}

func Test_ds1Layers_DeleteOrientation(t *testing.T) {}

func Test_ds1Layers_DeleteShadow(t *testing.T) {}

func Test_ds1Layers_DeleteSubstitution(t *testing.T) {}

func Test_ds1Layers_DeleteWall(t *testing.T) {}

func Test_ds1Layers_GetFloor(t *testing.T) {}

func Test_ds1Layers_GetOrientation(t *testing.T) {}

func Test_ds1Layers_GetShadow(t *testing.T) {}

func Test_ds1Layers_GetSubstitution(t *testing.T) {}

func Test_ds1Layers_GetWall(t *testing.T) {}

func Test_ds1Layers_InsertFloor(t *testing.T) {}

func Test_ds1Layers_InsertOrientation(t *testing.T) {}

func Test_ds1Layers_InsertShadow(t *testing.T) {}

func Test_ds1Layers_InsertSubstitution(t *testing.T) {}

func Test_ds1Layers_InsertWall(t *testing.T) {}

func Test_ds1Layers_PopFloor(t *testing.T) {}

func Test_ds1Layers_PopOrientation(t *testing.T) {}

func Test_ds1Layers_PopShadow(t *testing.T) {}

func Test_ds1Layers_PopSubstitution(t *testing.T) {}

func Test_ds1Layers_PopWall(t *testing.T) {}

func Test_ds1Layers_Push(t *testing.T) {
	t.Run("Floor", func(t *testing.T) {
		test_ds1Layers_PushLayer(floorLayerGroup, t)
	})

	t.Run("Wall", func(t *testing.T) {
		test_ds1Layers_PushLayer(wallLayerGroup, t)
	})

	t.Run("Orientation", func(t *testing.T) {
		test_ds1Layers_PushLayer(orientationLayerGroup, t)
	})

	t.Run("Shadow", func(t *testing.T) {
		test_ds1Layers_PushLayer(shadowLayerGroup, t)
	})

	t.Run("Substitution", func(t *testing.T) {
		test_ds1Layers_PushLayer(substitutionLayerGroup, t)
	})
}

// for all layer types, the test is the same
// when we push a layer, we expect an increment, and when we push a bunch of times,
// we expect to never exceed the max. we also expect to be able to retrieve a non-nil
// layer after we push.
func test_ds1Layers_PushLayer(lt layerGroupType, t *testing.T) {
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
	case orientationLayerGroup:
		push = func() { layers.PushOrientation(&layer{}) }
		get = layers.GetOrientation
		max = maxOrientationLayers
		group = &layers.Orientations
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
	push()
	check(max)
}

func test_ds1Layers_PopLayer(lt layerGroupType, t *testing.T) {}

func test_ds1Layers_InsertLayer(lt layerGroupType, t *testing.T) {}

func test_ds1Layers_DeleteLayer(lt layerGroupType, t *testing.T) {}

func test_ds1Layers_GetLayer(lt layerGroupType, t *testing.T) {}
