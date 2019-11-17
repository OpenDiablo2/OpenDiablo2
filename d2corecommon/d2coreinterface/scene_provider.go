package d2coreinterface

// SceneProvider provides the ability to change scenes
type SceneProvider interface {
	SetNextScene(nextScene Scene)
}
