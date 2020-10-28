package d2scene

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"

type SceneObject interface {
	Destroy()
}

type SceneObjectInteractive interface {
	SetInteractive(bool, *d2geom.Rectangle)
	RemoveInteractive()
}
