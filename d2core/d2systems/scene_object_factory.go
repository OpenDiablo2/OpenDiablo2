package d2systems

import (
	"image/color"
	"path/filepath"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

// responsible for wrapping the object factory calls and assigning the created object entity id's to the scene
type sceneObjectFactory struct {
	*BaseScene
	*d2util.Logger
}

func (s *sceneObjectFactory) addBasicComponents(id akara.EID) {
	node := s.AddSceneGraphNode(id)
	node.SetParent(s.Graph)

	_ = s.AddAlpha(id)
	_ = s.AddOrigin(id)
}

func (s *sceneObjectFactory) Sprite(x, y float64, imgPath, palPath string) akara.EID {
	s.Debugf("creating sprite: %s, %s", filepath.Base(imgPath), palPath)

	eid := s.sceneSystems.SpriteFactory.Sprite(x, y, imgPath, palPath)
	s.SceneObjects = append(s.SceneObjects, eid)

	s.addBasicComponents(eid)

	return eid
}

func (s *sceneObjectFactory) SegmentedSprite(x, y float64, imgPath, palPath string, xseg, yseg, frame int) akara.EID {
	s.Debugf("creating segmented sprite: %s, %s", filepath.Base(imgPath), palPath)

	eid := s.sceneSystems.SpriteFactory.SegmentedSprite(x, y, imgPath, palPath, xseg, yseg, frame)
	s.SceneObjects = append(s.SceneObjects, eid)

	s.addBasicComponents(eid)

	return eid
}

func (s *sceneObjectFactory) Viewport(priority, width, height int) akara.EID {
	s.Debugf("creating viewport #%d", priority)

	eid := s.NewEntity()
	s.AddViewport(eid)
	s.AddPriority(eid).Priority = priority

	if priority == mainViewport {
		s.AddMainViewport(eid)
	}

	camera := s.AddCamera(eid)
	camera.Size.X = float64(width)
	camera.Size.Y = float64(height)

	sfc := s.sceneSystems.RenderSystem.renderer.NewSurface(width, height)

	sfc.Clear(color.Transparent)

	s.AddTexture(eid).Texture = sfc

	s.Viewports = append(s.Viewports, eid)

	s.addBasicComponents(eid)

	return eid
}

func (s *sceneObjectFactory) Rectangle(x, y, width, height int, c color.Color) akara.EID {
	s.Debug("creating rectangle")

	eid := s.sceneSystems.ShapeSystem.Rectangle(x, y, width, height, c)

	s.addBasicComponents(eid)

	transform := s.AddTransform(eid)
	transform.Translation.X, transform.Translation.Y = float64(x), float64(y)

	s.SceneObjects = append(s.SceneObjects, eid)

	return eid
}

func (s *sceneObjectFactory) Button(x, y float64, imgPath, palPath string) akara.EID {
	s.Debug("creating button")

	eid := s.sceneSystems.UIWidgetFactory.Button(x, y, imgPath, palPath)

	s.addBasicComponents(eid)

	transform := s.AddTransform(eid)
	transform.Translation.X, transform.Translation.Y = float64(x), float64(y)

	s.SceneObjects = append(s.SceneObjects, eid)

	return eid
}

func (s *sceneObjectFactory) Label(fontPath, spritePath, palettePath string) akara.EID {
	s.Debug("creating label")

	eid := s.sceneSystems.UIWidgetFactory.Label(fontPath, spritePath, palettePath)

	s.addBasicComponents(eid)

	s.SceneObjects = append(s.SceneObjects, eid)

	return eid
}
