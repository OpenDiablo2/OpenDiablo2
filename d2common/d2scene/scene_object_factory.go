package d2scene

import (
	"github.com/gravestench/akara"
)

func NewSceneObjectFactory(world *akara.World, scene *BasicScene) *SceneObjectFactory {
	return &SceneObjectFactory{
		world: world,
		scene: scene,
	}
}

type SceneObjectFactory struct {
	world *akara.World
	scene *BasicScene
}
