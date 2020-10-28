package d2scene

import "github.com/gravestench/akara"

func NewBasicSceneObject(scene *BasicScene) *BasicSceneObject {
	object := &BasicSceneObject{
		Scene:                 scene,
		World:                 scene.World,
		ID:                    scene.World.NewEntity(),
		CameraVisibilityFlags: akara.NewBitSet(),
	}

	object.CameraVisibilityFlags.Set(0, true) // camera index 0 is the main camera

	return object
}

type BasicSceneObject struct {
	Scene                 *BasicScene
	ID                    akara.EID
	World                 *akara.World
	CameraVisibilityFlags *akara.BitSet
}
