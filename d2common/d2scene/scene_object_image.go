package d2scene

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
	"github.com/gravestench/akara"
)

// Image creates an image object in the scene
func (factory *SceneObjectFactory) Image(x, y float64, imagePath, palettePath string) *Image {
	img := &Image{
		imagePath:        imagePath,
		palettePath:      palettePath,
		BasicSceneObject: NewBasicSceneObject(factory.scene),
	}

	img.image = factory.world.NewEntity()
	img.palette = factory.world.NewEntity()

	img.Scene.Components.AddFilePath(img.image).Path = imagePath
	img.Scene.Components.AddFilePath(img.palette).Path = palettePath

	img.position = img.Scene.Components.AddPosition(img.ID)
	img.position.Set(x, y)

	img.origin = img.Scene.Components.AddOrigin(img.ID)
	img.SetOrigin(0.5, 0.5)

	return img
}

type Image struct {
	*BasicSceneObject
	imagePath   string
	palettePath string
	image       akara.EID
	palette     akara.EID
	position    *d2components.PositionComponent
	origin      *d2components.OriginComponent
	displaySize *d2components.SizeComponent
	surface     *d2components.SurfaceComponent
}

func (i *Image) SetOrigin(ox, oy float64) *Image {
	i.origin.X, i.origin.Y = ox, oy
	return i
}

func (i *Image) SetDisplaySize(width, height uint) *Image {
	i.displaySize.Width, i.displaySize.Height = width, height
	return i
}

func (i *Image) SetInteractive(set bool, shape *d2geom.Rectangle) *Image {
	return i
}

func (i *Image) Destroy() {
	i.Scene.World.RemoveEntity(i.image)
	i.Scene.World.RemoveEntity(i.palette)
	i.Scene.World.RemoveEntity(i.ID)
}
