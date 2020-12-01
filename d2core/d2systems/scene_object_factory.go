package d2systems

import (
	"path/filepath"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

// responsible for wrapping the object factory calls and assigning the created object entity id's to the scene
type sceneObjectFactory struct {
	*BaseScene
	*d2util.Logger
}

func (s *sceneObjectFactory) addBasicComponenets(id akara.EID) {
	node := s.AddSceneGraphNode(id)
	node.SetParent(s.Graph)

	_ = s.AddAlpha(id)
	_ = s.AddOrigin(id)
}

func (s *sceneObjectFactory) Sprite(x, y float64, imgPath, palPath string) akara.EID {
	s.Infof("creating sprite: %s, %s", filepath.Base(imgPath), palPath)

	eid := s.systems.SpriteFactory.Sprite(x, y, imgPath, palPath)
	s.GameObjects = append(s.GameObjects, eid)

	s.addBasicComponenets(eid)

	return eid
}

func (s *sceneObjectFactory) SegmentedSprite(x, y float64, imgPath, palPath string, xseg, yseg, frame int) akara.EID {
	s.Infof("creating segmented sprite: %s, %s", filepath.Base(imgPath), palPath)

	eid := s.systems.SpriteFactory.SegmentedSprite(x, y, imgPath, palPath, xseg, yseg, frame)
	s.GameObjects = append(s.GameObjects, eid)

	s.addBasicComponenets(eid)

	return eid
}
