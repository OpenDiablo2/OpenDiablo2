package d2enum

// SceneState enumerates the different states a scene can be in
type SceneState int

// Scene states
const (
	SceneStateUninitialized SceneState = iota
	SceneStateBooting
	SceneStateBooted
)
