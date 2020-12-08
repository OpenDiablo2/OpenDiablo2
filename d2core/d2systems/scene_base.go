package d2systems

import (
	"fmt"
	"image/color"
	"sort"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom/rectangle"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2scene"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	mainViewport int = 0
)

// NewBaseScene creates a new base scene instance
func NewBaseScene(key string) *BaseScene {
	base := &BaseScene{
		Graph:           d2scene.NewNode(),
		BaseSystem:      akara.NewBaseSystem(),
		Logger:          d2util.NewLogger(),
		key:             key,
		Viewports:       make([]akara.EID, 0),
		SceneObjects:    make([]akara.EID, 0),
		backgroundColor: color.Transparent,
	}

	base.SetPrefix(key)

	return base
}

var _ akara.System = &BaseScene{}

type sceneSystems struct {
	*RenderSystem
	*InputSystem
	*GameObjectFactory
}

type sceneComponents struct {
	d2components.SceneGraphNodeFactory
	d2components.ViewportFactory
	d2components.MainViewportFactory
	d2components.ViewportFilterFactory
	d2components.PriorityFactory
	d2components.CameraFactory
	d2components.TextureFactory
	d2components.InteractiveFactory
	d2components.TransformFactory
	d2components.SpriteFactory
	d2components.OriginFactory
	d2components.AlphaFactory
	d2components.DrawEffectFactory
	d2components.RectangleFactory
	d2components.ColorFactory
	d2components.CommandRegistrationFactory
	d2components.DirtyFactory
}

// BaseScene encapsulates common behaviors for systems that are considered "scenes",
// such as the main menu, the in-game map, the console, etc.
//
// The base scene is responsible for generic behaviors common to all scenes,
// like initializing the default viewport, or rendering game objects to the viewports.
type BaseScene struct {
	*akara.BaseSystem
	sceneSystems
	sceneComponents
	Geom struct {
		Rectangle rectangle.Namespace
	}
	*d2util.Logger
	key             string
	booted          bool
	paused          bool
	Add             *sceneObjectFactory
	Viewports       []akara.EID
	SceneObjects    []akara.EID
	Graph           *d2scene.Node // the root node
	backgroundColor color.Color
}

// Booted returns whether or not the scene has booted
func (s *BaseScene) Booted() bool {
	return s.booted
}

// Paused returns whether or not the scene is paused
func (s *BaseScene) Paused() bool {
	return s.paused
}

// Init the base scene
func (s *BaseScene) Init(world *akara.World) {
	s.World = world

	if s.World == nil {
		s.SetActive(false)
		return
	}
}

func (s *BaseScene) boot() {
	s.Debug("base scene booting ...")

	s.Add = &sceneObjectFactory{
		BaseScene: s,
		Logger:    d2util.NewLogger(),
	}

	s.Add.SetPrefix(fmt.Sprintf("%s -> %s", s.key, "Object Factory"))

	s.bindRequiredSystems()

	if !s.requiredSystemsPresent() {
		return
	}

	s.setupFactories()

	s.setupSceneObjectFactories()

	const (
		defaultWidth  = 800
		defaultHeight = 600
	)

	s.Add.Viewport(mainViewport, defaultWidth, defaultHeight)

	s.Debug("base scene booted!")
	s.booted = true
}

func (s *BaseScene) bindRequiredSystems() {
	for idx := range s.Systems {
		noRenderSys := s.sceneSystems.RenderSystem == nil
		noInputSys := s.sceneSystems.InputSystem == nil
		noObjectFactory := s.sceneSystems.GameObjectFactory == nil

		sys := s.Systems[idx]

		if rendersys, found := sys.(*RenderSystem); found && noRenderSys {
			s.sceneSystems.RenderSystem = rendersys
			continue
		}

		if inputSys, found := sys.(*InputSystem); found && noInputSys {
			s.sceneSystems.InputSystem = inputSys
			continue
		}

		if objFactory, found := sys.(*GameObjectFactory); found && noObjectFactory {
			s.sceneSystems.GameObjectFactory = objFactory
			continue
		}
	}
}

func (s *BaseScene) requiredSystemsPresent() bool {
	if s.sceneSystems.RenderSystem == nil {
		s.Debug("waiting for render system ...")
		return false
	}

	if s.sceneSystems.RenderSystem.renderer == nil {
		s.Debug("waiting for renderer instance ...")
		return false
	}

	if s.sceneSystems.InputSystem == nil {
		s.Debug("waiting for input system")
		return false
	}

	if s.sceneSystems.GameObjectFactory == nil {
		s.Debug("waiting for game object factory ...")
		return false
	}

	return true
}

func (s *BaseScene) setupSceneObjectFactories() {
	s.sceneSystems.SpriteFactory.RenderSystem = s.sceneSystems.RenderSystem
	s.sceneSystems.ShapeSystem.RenderSystem = s.sceneSystems.RenderSystem
	s.sceneSystems.UIWidgetFactory.RenderSystem = s.sceneSystems.RenderSystem
}

func (s *BaseScene) setupFactories() {
	s.Debug("setting up component factories")

	s.InjectComponent(&d2components.MainViewport{}, &s.MainViewport)
	s.InjectComponent(&d2components.Viewport{}, &s.Viewport)
	s.InjectComponent(&d2components.ViewportFilter{}, &s.ViewportFilter)
	s.InjectComponent(&d2components.Camera{}, &s.Camera)
	s.InjectComponent(&d2components.Priority{}, &s.Priority)
	s.InjectComponent(&d2components.Texture{}, &s.Texture)
	s.InjectComponent(&d2components.Interactive{}, &s.Interactive)
	s.InjectComponent(&d2components.Transform{}, &s.Transform)
	s.InjectComponent(&d2components.Origin{}, &s.Origin)
	s.InjectComponent(&d2components.Alpha{}, &s.Alpha)
	s.InjectComponent(&d2components.SceneGraphNode{}, &s.SceneGraphNode)
	s.InjectComponent(&d2components.DrawEffect{}, &s.DrawEffect)
	s.InjectComponent(&d2components.Sprite{}, &s.SpriteFactory.Sprite)
	s.InjectComponent(&d2components.Rectangle{}, &s.RectangleFactory.Rectangle)
	s.InjectComponent(&d2components.Color{}, &s.Color)
	s.InjectComponent(&d2components.CommandRegistration{}, &s.CommandRegistration)
	s.InjectComponent(&d2components.Dirty{}, &s.Dirty)
}

// Key returns the scene's key
func (s *BaseScene) Key() string {
	return s.key
}

// Update performs scene boot and renders the scene viewports
func (s *BaseScene) Update() {
	if !s.booted {
		s.boot()
	}

	if !s.booted {
		return
	}

	s.updateSceneGraph()
	s.renderViewports()
}

func (s *BaseScene) updateSceneGraph() {
	for _, eid := range s.SceneObjects {
		node, found := s.GetSceneGraphNode(eid)
		if !found {
			continue
		}

		transform, found := s.GetTransform(eid)
		if !found {
			continue
		}

		node.Local = transform.GetMatrix()
	}

	s.Graph.UpdateWorldMatrix()
}

func (s *BaseScene) renderViewports() {
	if s.sceneSystems.RenderSystem == nil {
		s.Warning("render system not present")
		return
	}

	if s.sceneSystems.RenderSystem.renderer == nil {
		s.Warning("render system doesn't have a renderer instance")
		return
	}

	numViewports := len(s.Viewports)

	if numViewports < 1 {
		s.Warning("scene does not have a main viewport")
		return
	}

	viewportObjects := s.binGameObjectsByViewport()

	for idx := numViewports - 1; idx >= 0; idx-- {
		s.renderViewport(idx, viewportObjects[idx])
	}

	if len(s.Viewports) > 1 {
		s.renderViewportsToMainViewport()
	}
}

func (s *BaseScene) binGameObjectsByViewport() map[int][]akara.EID {
	bins := make(map[int][]akara.EID)

	for _, eid := range s.SceneObjects {
		vpfilter, found := s.GetViewportFilter(eid)
		if !found {
			vpfilter = s.AddViewportFilter(eid)
			vpfilter.Set(mainViewport, true)
		}

		for _, vpidx64 := range vpfilter.ToIntArray() {
			vpidx := int(vpidx64)

			_, found := bins[vpidx]
			if !found {
				bins[vpidx] = make([]akara.EID, 0)
			}

			bins[vpidx] = append(bins[vpidx], eid)
		}
	}

	return bins
}

func (s *BaseScene) renderViewport(idx int, objects []akara.EID) {
	viewportEID := s.Viewports[idx]

	// the first viewport is always the main viewport
	if idx == mainViewport {
		s.AddMainViewport(viewportEID)
	} else {
		s.MainViewport.Remove(viewportEID)
	}

	sfc, found := s.GetTexture(viewportEID)
	if !found || sfc.Texture == nil {
		return
	}

	if idx == mainViewport {
		sfc.Texture.Clear(s.backgroundColor)
	}

	for _, objectEID := range objects {
		s.renderObject(viewportEID, objectEID)
	}
}

func (s *BaseScene) renderObject(viewportEID, objectEID akara.EID) {
	vpTexture, found := s.GetTexture(viewportEID)
	if !found || vpTexture.Texture == nil {
		return
	}

	vpNode, found := s.GetSceneGraphNode(viewportEID)
	if !found {
		vpNode = s.AddSceneGraphNode(viewportEID)
	}

	// translation, rotation, and scale vec3's
	vpTrans, vpRot, vpScale := vpNode.Local.Invert().Decompose()

	objTexture, found := s.GetTexture(objectEID)
	if !found {
		return
	}


	alpha, found := s.GetAlpha(objectEID)
	if !found {
		alpha = s.AddAlpha(objectEID)
	}

	origin, found := s.GetOrigin(objectEID)
	if !found {
		origin = s.AddOrigin(objectEID)
	}

	node, found := s.GetSceneGraphNode(objectEID)
	if !found {
		node = s.AddSceneGraphNode(objectEID)
		node.SetParent(s.Graph)
	}

	drawEffect, found := s.GetDrawEffect(objectEID)
	if found {
		vpTexture.Texture.PushEffect(drawEffect.DrawEffect)
		defer vpTexture.Texture.Pop()
	}

	objNode, found := s.GetSceneGraphNode(objectEID)
	if !found {
		objNode = s.AddSceneGraphNode(objectEID)
	}

	// translation, rotation, and scale vec3's
	objTrans, objRot, objScale := objNode.Local.Decompose()

	ox, oy := origin.X, origin.Y
	tx, ty := objTrans.Add(vpTrans).XY()

	vpTexture.Texture.PushTranslation(int(tx+ox), int(ty+oy))
	defer vpTexture.Texture.Pop()

	vpTexture.Texture.PushScale(objScale.Multiply(vpScale).XY())
	defer vpTexture.Texture.Pop()

	vpTexture.Texture.PushRotate(objRot.Add(vpRot).Z)
	defer vpTexture.Texture.Pop()

	const maxAlpha = 255

	vpTexture.Texture.PushColor(color.Alpha{A: uint8(alpha.Alpha * maxAlpha)})
	defer vpTexture.Texture.Pop()

	segment, found := s.sceneSystems.SpriteFactory.GetSegmentedSprite(objectEID)
	if found {
		s.renderSegmentedSprite(vpTexture.Texture, objectEID, segment)
		return
	}

	vpTexture.Texture.Render(objTexture.Texture)
}

func (s *BaseScene) renderSegmentedSprite(screen d2interface.Surface, id akara.EID, seg *d2components.SegmentedSprite) {
	target := screen.Renderer().NewSurface(screen.GetSize())

	animation, found := s.GetSprite(id)
	if !found {
		return
	}

	var offsetY int

	segmentsX, segmentsY := seg.Xsegments, seg.Ysegments
	frameOffset := seg.FrameOffset

	for y := 0; y < segmentsY; y++ {
		var offsetX, maxFrameHeight int

		for x := 0; x < segmentsX; x++ {
			idx := x + y*segmentsX + frameOffset*segmentsX*segmentsY
			if err := animation.SetCurrentFrame(idx); err != nil {
				s.Error("SetCurrentFrame error " + err.Error())
			}

			target.PushTranslation(x+offsetX, y+offsetY)
			target.Render(animation.GetCurrentFrameSurface())
			target.Pop()

			frameWidth, frameHeight := animation.GetCurrentFrameSize()
			maxFrameHeight = d2math.MaxInt(maxFrameHeight, frameHeight)
			offsetX += frameWidth - 1
		}

		offsetY += maxFrameHeight - 1
	}

	screen.Render(target)
}

func (s *BaseScene) renderViewportsToMainViewport() {
	mainID := s.Viewports[mainViewport]
	otherIDs := s.Viewports[mainViewport+1:]

	sort.Slice(otherIDs, func(i, j int) bool {
		p1, found := s.GetPriority(otherIDs[i])
		if !found {
			return false
		}

		p2, found := s.GetPriority(otherIDs[j])
		if !found {
			return false
		}

		return p1.Priority > p2.Priority
	})

	main, found := s.GetTexture(mainID)
	if !found {
		return
	}

	for _, id := range otherIDs {
		other, found := s.GetTexture(id)
		if !found {
			continue
		}

		main.Texture.Render(other.Texture)
	}
}

func (s *BaseScene) RegisterTerminalCommand(name, desc string, fn interface{}) {
	regID := s.NewEntity()
	reg := s.AddCommandRegistration(regID)
	s.AddDirty(regID)

	reg.Name = name
	reg.Description = desc
	reg.Callback = fn
}
