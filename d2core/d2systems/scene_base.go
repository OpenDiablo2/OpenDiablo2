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
	Render *RenderSystem
	Input *InputSystem
	*GameObjectFactory
}

type sceneComponents struct {
	SceneGraphNode d2components.SceneGraphNodeFactory
	Viewport d2components.ViewportFactory
	MainViewport d2components.MainViewportFactory
	ViewportFilter d2components.ViewportFilterFactory
	Priority d2components.PriorityFactory
	Camera d2components.CameraFactory
	Texture d2components.TextureFactory
	Interactive d2components.InteractiveFactory
	Transform d2components.TransformFactory
	Sprite d2components.SpriteFactory
	SegmentedSprite d2components.SegmentedSpriteFactory
	Origin d2components.OriginFactory
	Alpha d2components.AlphaFactory
	DrawEffect d2components.DrawEffectFactory
	Rectangle d2components.RectangleFactory
	Color d2components.ColorFactory
	CommandRegistration d2components.CommandRegistrationFactory
	Dirty d2components.DirtyFactory
}

// BaseScene encapsulates common behaviors for systems that are considered "scenes",
// such as the main menu, the in-game map, the console, etc.
//
// The base scene is responsible for generic behaviors common to all scenes,
// like initializing the default viewport, or rendering game objects to the viewports.
type BaseScene struct {
	*akara.BaseSystem
	sceneSystems
	Components sceneComponents
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
		noRenderSys := s.sceneSystems.Render == nil
		noInputSys := s.sceneSystems.Input == nil
		noObjectFactory := s.sceneSystems.GameObjectFactory == nil

		sys := s.Systems[idx]

		if rendersys, found := sys.(*RenderSystem); found && noRenderSys {
			s.sceneSystems.Render = rendersys
			continue
		}

		if inputSys, found := sys.(*InputSystem); found && noInputSys {
			s.sceneSystems.Input = inputSys
			continue
		}

		if objFactory, found := sys.(*GameObjectFactory); found && noObjectFactory {
			s.sceneSystems.GameObjectFactory = objFactory
			continue
		}
	}
}

func (s *BaseScene) requiredSystemsPresent() bool {
	if s.sceneSystems.Render == nil {
		s.Debug("waiting for render system ...")
		return false
	}

	if s.sceneSystems.Render.renderer == nil {
		s.Debug("waiting for renderer instance ...")
		return false
	}

	if s.sceneSystems.Input == nil {
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
	s.sceneSystems.GameObjectFactory.Sprites.RenderSystem = s.sceneSystems.Render
	s.sceneSystems.GameObjectFactory.Shapes.RenderSystem = s.sceneSystems.Render
	s.sceneSystems.GameObjectFactory.UI.RenderSystem = s.sceneSystems.Render
}

func (s *BaseScene) setupFactories() {
	s.Debug("setting up component factories")

	s.InjectComponent(&d2components.MainViewport{}, &s.Components.MainViewport.ComponentFactory)
	s.InjectComponent(&d2components.Viewport{}, &s.Components.Viewport.ComponentFactory)
	s.InjectComponent(&d2components.ViewportFilter{}, &s.Components.ViewportFilter.ComponentFactory)
	s.InjectComponent(&d2components.Camera{}, &s.Components.Camera.ComponentFactory)
	s.InjectComponent(&d2components.Priority{}, &s.Components.Priority.ComponentFactory)
	s.InjectComponent(&d2components.Texture{}, &s.Components.Texture.ComponentFactory)
	s.InjectComponent(&d2components.Interactive{}, &s.Components.Interactive.ComponentFactory)
	s.InjectComponent(&d2components.Transform{}, &s.Components.Transform.ComponentFactory)
	s.InjectComponent(&d2components.Origin{}, &s.Components.Origin.ComponentFactory)
	s.InjectComponent(&d2components.Alpha{}, &s.Components.Alpha.ComponentFactory)
	s.InjectComponent(&d2components.SceneGraphNode{}, &s.Components.SceneGraphNode.ComponentFactory)
	s.InjectComponent(&d2components.DrawEffect{}, &s.Components.DrawEffect.ComponentFactory)
	s.InjectComponent(&d2components.Sprite{}, &s.Components.Sprite.ComponentFactory)
	s.InjectComponent(&d2components.SegmentedSprite{}, &s.Components.SegmentedSprite.ComponentFactory)
	s.InjectComponent(&d2components.Rectangle{}, &s.Components.Rectangle.ComponentFactory)
	s.InjectComponent(&d2components.Color{}, &s.Components.Color.ComponentFactory)
	s.InjectComponent(&d2components.CommandRegistration{}, &s.Components.CommandRegistration.ComponentFactory)
	s.InjectComponent(&d2components.Dirty{}, &s.Components.Dirty.ComponentFactory)
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
		node, found := s.Components.SceneGraphNode.Get(eid)
		if !found {
			continue
		}

		transform, found := s.Components.Transform.Get(eid)
		if !found {
			continue
		}

		node.Local = transform.GetMatrix()
	}

	s.Graph.UpdateWorldMatrix()
}

func (s *BaseScene) renderViewports() {
	if s.sceneSystems.Render == nil {
		s.Warning("render system not present")
		return
	}

	if s.sceneSystems.Render.renderer == nil {
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
		vpfilter, found := s.Components.ViewportFilter.Get(eid)
		if !found {
			vpfilter = s.Components.ViewportFilter.Add(eid)
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
		s.Components.MainViewport.Add(viewportEID)
	} else {
		s.Components.MainViewport.Remove(viewportEID)
	}

	sfc, found := s.Components.Texture.Get(viewportEID)
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
	vpTexture, found := s.Components.Texture.Get(viewportEID)
	if !found || vpTexture.Texture == nil {
		return
	}

	vpNode, found := s.Components.SceneGraphNode.Get(viewportEID)
	if !found {
		vpNode = s.Components.SceneGraphNode.Add(viewportEID)
	}

	// translation, rotation, and scale vec3's
	vpTrans, vpRot, vpScale := vpNode.Local.Invert().Decompose()

	objTexture, found := s.Components.Texture.Get(objectEID)
	if !found {
		return
	}


	alpha, found := s.Components.Alpha.Get(objectEID)
	if !found {
		alpha = s.Components.Alpha.Add(objectEID)
	}

	origin, found := s.Components.Origin.Get(objectEID)
	if !found {
		origin = s.Components.Origin.Add(objectEID)
	}

	drawEffect, found := s.Components.DrawEffect.Get(objectEID)
	if found {
		vpTexture.Texture.PushEffect(drawEffect.DrawEffect)
		defer vpTexture.Texture.Pop()
	}

	objNode, found := s.Components.SceneGraphNode.Get(objectEID)
	if !found {
		objNode = s.Components.SceneGraphNode.Add(objectEID)
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

	segment, found := s.Components.SegmentedSprite.Get(objectEID)
	if found {
		s.renderSegmentedSprite(vpTexture.Texture, objectEID, segment)
		return
	}

	vpTexture.Texture.Render(objTexture.Texture)
}

func (s *BaseScene) renderSegmentedSprite(screen d2interface.Surface, id akara.EID, seg *d2components.SegmentedSprite) {
	target := screen.Renderer().NewSurface(screen.GetSize())

	sprite, found := s.Components.Sprite.Get(id)
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
			if err := sprite.SetCurrentFrame(idx); err != nil {
				s.Error("SetCurrentFrame error " + err.Error())
			}

			target.PushTranslation(x+offsetX, y+offsetY)
			target.Render(sprite.GetCurrentFrameSurface())
			target.Pop()

			frameWidth, frameHeight := sprite.GetCurrentFrameSize()
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
		p1, found := s.Components.Priority.Get(otherIDs[i])
		if !found {
			return false
		}

		p2, found := s.Components.Priority.Get(otherIDs[j])
		if !found {
			return false
		}

		return p1.Priority > p2.Priority
	})

	main, found := s.Components.Texture.Get(mainID)
	if !found {
		return
	}

	for _, id := range otherIDs {
		other, found := s.Components.Texture.Get(id)
		if !found {
			continue
		}

		main.Texture.Render(other.Texture)
	}
}

func (s *BaseScene) RegisterTerminalCommand(name, desc string, args []string, fn func(args []string) error) {
	regID := s.NewEntity()
	reg := s.Components.CommandRegistration.Add(regID)
	s.Components.Dirty.Add(regID)

	reg.Name = name
	reg.Description = desc
	reg.Arguments = args
	reg.Callback = fn
}
