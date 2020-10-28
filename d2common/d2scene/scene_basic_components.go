package d2scene

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
	"github.com/gravestench/akara"
)

func NewBasicSceneComponents(w *akara.World) *BasicSceneComponents {
	return &BasicSceneComponents{
		FilePathMap: w.InjectMap(d2components.Camera).(*d2components.FilePathMap),
		CameraMap:   w.InjectMap(d2components.Camera).(*d2components.CameraMap),
		ViewPortMap: w.InjectMap(d2components.ViewPort).(*d2components.ViewPortMap),
		SurfaceMap:  w.InjectMap(d2components.Surface).(*d2components.SurfaceMap),
		PositionMap: w.InjectMap(d2components.Position).(*d2components.PositionMap),
		SizeMap:     w.InjectMap(d2components.Size).(*d2components.SizeMap),
		OriginMap:   w.InjectMap(d2components.Origin).(*d2components.OriginMap),
	}
}

type BasicSceneComponents struct {
	*d2components.CameraMap
	*d2components.ViewPortMap
	*d2components.SurfaceMap
	*d2components.FilePathMap
	*d2components.FileTypeMap
	*d2components.FileHandleMap
	*d2components.PositionMap
	*d2components.SizeMap
	*d2components.OriginMap
}
