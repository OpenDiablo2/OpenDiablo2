package Scenes

// SceneProvider provides the ability to change scenes
type SceneProvider interface {
	SetNextScene(nextScene Scene)
}
