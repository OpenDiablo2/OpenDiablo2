package d2systems

import (
	"image/color"
	"path/filepath"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2button"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2checkbox"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

// responsible for wrapping the object factory calls and assigning the created object entity id's to the scene
type sceneObjectFactory struct {
	*BaseScene
	*d2util.Logger
}

func (s *sceneObjectFactory) addBasicComponents(id akara.EID) {
	node := s.Components.SceneGraphNode.Add(id)
	node.SetParent(s.Graph)

	_ = s.Components.Transform.Add(id)
	_ = s.Components.Origin.Add(id)
	_ = s.Components.Alpha.Add(id)
}

func (s *sceneObjectFactory) Sprite(x, y float64, imgPath, palPath string) akara.EID {
	s.Debugf("creating sprite: %s, %s", filepath.Base(imgPath), palPath)

	eid := s.sceneSystems.Sprites.Sprite(x, y, imgPath, palPath)
	s.SceneObjects = append(s.SceneObjects, eid)

	s.addBasicComponents(eid)

	return eid
}

func (s *sceneObjectFactory) SegmentedSprite(x, y float64, imgPath, palPath string, xseg, yseg, frame int) akara.EID {
	s.Debugf("creating segmented sprite: %s, %s", filepath.Base(imgPath), palPath)

	eid := s.sceneSystems.Sprites.SegmentedSprite(x, y, imgPath, palPath, xseg, yseg, frame)
	s.SceneObjects = append(s.SceneObjects, eid)

	s.addBasicComponents(eid)

	return eid
}

func (s *sceneObjectFactory) Viewport(priority, width, height int) akara.EID {
	s.Debugf("creating viewport #%d", priority)

	eid := s.NewEntity()
	s.Components.Viewport.Add(eid)

	s.Components.Priority.Add(eid).Priority = priority

	if priority == mainViewport {
		s.Components.MainViewport.Add(eid)
	}

	camera := s.Components.Camera.Add(eid)
	camera.Size.X = float64(width)
	camera.Size.Y = float64(height)

	sfc := s.sceneSystems.Render.renderer.NewSurface(width, height)

	sfc.Clear(color.Transparent)

	s.Components.Texture.Add(eid).Texture = sfc

	s.Viewports = append(s.Viewports, eid)

	s.addBasicComponents(eid)

	return eid
}

func (s *sceneObjectFactory) Rectangle(x, y, width, height int, c color.Color) akara.EID {
	s.Debug("creating rectangle")

	eid := s.sceneSystems.Shapes.Rectangle(x, y, width, height, c)

	s.addBasicComponents(eid)

	transform := s.Components.Transform.Add(eid)
	transform.Translation.X, transform.Translation.Y = float64(x), float64(y)

	s.SceneObjects = append(s.SceneObjects, eid)

	return eid
}

func (s *sceneObjectFactory) Button(x, y float64, btnType d2button.ButtonType, text string) akara.EID {
	s.Debug("creating button")

	buttonEID := s.sceneSystems.UI.Button(x, y, btnType, "")

	s.SceneObjects = append(s.SceneObjects, buttonEID)

	layout := d2button.GetLayout(btnType)

	s.addBasicComponents(buttonEID)

	btnTRS := s.Components.Transform.Add(buttonEID)
	btnTRS.Translation.X, btnTRS.Translation.Y = x, y

	btnNode := s.Components.SceneGraphNode.Add(buttonEID)

	if text != "" {
		labelEID := s.Label(text, layout.FontPath, layout.PalettePath)
		labelNode := s.Components.SceneGraphNode.Add(labelEID)
		labelNode.SetParent(btnNode.Node)
	}

	return buttonEID
}

func (s *sceneObjectFactory) Label(text, fontSpritePath, palettePath string) akara.EID {
	s.Debug("creating label")

	eid := s.sceneSystems.UI.Label(text, fontSpritePath, palettePath)

	s.addBasicComponents(eid)

	s.SceneObjects = append(s.SceneObjects, eid)

	return eid
}

// Checkbox creates a Checkbox in the scene, with an attached Label
func (s *sceneObjectFactory) Checkbox(x, y float64, checkedState, enabled bool,
	text string, callback func(akara.Component) bool) akara.EID {
	checkboxEID := s.sceneSystems.UI.Checkbox(x, y, checkedState, enabled, callback)
	s.SceneObjects = append(s.SceneObjects, checkboxEID)

	s.addBasicComponents(checkboxEID)

	checkboxNode := s.Components.SceneGraphNode.Add(checkboxEID)

	// create a Label as a child of the Checkbox if text was given
	if text != "" {
		layout := d2checkbox.GetDefaultLayout()
		labelEID := s.Label(text, layout.FontPath, layout.PalettePath)
		labelNode := s.Components.SceneGraphNode.Add(labelEID)
		labelNode.SetParent(checkboxNode.Node)

		labelTrs := s.Components.Transform.Add(labelEID)
		labelTrs.Translation.X = layout.TextOffset

		label, _ := s.Components.Label.Get(labelEID)
		checkbox, _ := s.Components.Checkbox.Get(checkboxEID)
		checkbox.Label = label.Label
	}

	return checkboxEID
}
