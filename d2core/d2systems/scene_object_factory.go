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
	s.Infof("creating sprite: %s, %s", filepath.Base(imgPath), palPath)

	eid := s.baseSystems.SpriteFactory.Sprite(x, y, imgPath, palPath)
	s.GameObjects = append(s.GameObjects, eid)

	s.addBasicComponents(eid)

	return eid
}

func (s *sceneObjectFactory) SegmentedSprite(x, y float64, imgPath, palPath string, xseg, yseg, frame int) akara.EID {
	s.Infof("creating segmented sprite: %s, %s", filepath.Base(imgPath), palPath)

	eid := s.baseSystems.SpriteFactory.SegmentedSprite(x, y, imgPath, palPath, xseg, yseg, frame)
	s.GameObjects = append(s.GameObjects, eid)

	s.addBasicComponents(eid)

	return eid
}

func (s *sceneObjectFactory) Viewport(priority, width, height int) akara.EID {
	s.Infof("creating viewport #%d", priority)

	eid := s.NewEntity()
	s.AddViewport(eid)
	s.AddPriority(eid).Priority = priority

	if priority == mainViewport {
		s.AddMainViewport(eid)
	}

	camera := s.AddCamera(eid)
	camera.Size.X = float64(width)
	camera.Size.Y = float64(height)

	sfc := s.baseSystems.RenderSystem.renderer.NewSurface(width, height)

	sfc.Clear(color.Transparent)

	s.AddTexture(eid).Texture = sfc

	s.Viewports = append(s.Viewports, eid)

	s.addBasicComponents(eid)

	return eid
}

func (s *sceneObjectFactory) Rectangle(x, y, width, height int, color color.Color) akara.EID {
	s.Info("creating rectangle")

	eid := s.baseSystems.ShapeSystem.Rectangle(x, y, width, height, color)

	s.addBasicComponents(eid)

	s.GameObjects = append(s.GameObjects, eid)

	return eid
}
