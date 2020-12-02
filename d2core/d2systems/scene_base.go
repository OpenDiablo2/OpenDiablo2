package d2systems

import (
	"fmt"
	"image/color"

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
		BaseSystem:      &akara.BaseSystem{},
		Logger:          d2util.NewLogger(),
		key:             key,
		Viewports:       make([]akara.EID, 0),
		GameObjects:     make([]akara.EID, 0),
		baseSystems:     &baseSystems{},
		backgroundColor: color.Transparent,
	}

	base.SetPrefix(key)

	return base
}

var _ akara.System = &BaseScene{}

type baseSystems struct {
	*RenderSystem
	*InputSystem
	*GameObjectFactory
}

// BaseScene encapsulates common behaviors for baseSystems that are considered "scenes",
// such as the main menu, the in-game map, the console, etc.
//
// The base scene is responsible for generic behaviors common to all scenes,
// like initializing the default viewport, or rendering game objects to the viewports.
type BaseScene struct {
	*akara.BaseSystem
	*baseSystems
	Geom struct {
		Rectangle rectangle.Namespace
	}
	*d2util.Logger
	key             string
	booted          bool
	paused          bool
	Add             *sceneObjectFactory
	Viewports       []akara.EID
	GameObjects     []akara.EID
	Graph           *d2scene.Node // the root node
	backgroundColor color.Color
	d2components.SceneGraphNodeFactory
	d2components.ViewportFactory
	d2components.MainViewportFactory
	d2components.ViewportFilterFactory
	d2components.PriorityFactory
	d2components.CameraFactory
	d2components.TextureFactory
	d2components.InteractiveFactory
	d2components.PositionFactory
	d2components.ScaleFactory
	d2components.SpriteFactory
	d2components.OriginFactory
	d2components.AlphaFactory
	d2components.DrawEffectFactory
	d2components.RectangleFactory
	d2components.ColorFactory
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
	s.Info("base scene booting ...")

	s.Add = &sceneObjectFactory{
		BaseScene: s,
		Logger:    d2util.NewLogger(),
	}

	s.Add.SetPrefix(fmt.Sprintf("%s -> %s", s.key, "Object Factory"))

	for idx := range s.Systems {
		if rendersys, ok := s.Systems[idx].(*RenderSystem); ok && s.baseSystems.RenderSystem == nil {
			s.baseSystems.RenderSystem = rendersys
			continue
		}

		if inputSys, ok := s.Systems[idx].(*InputSystem); ok && s.baseSystems.InputSystem == nil {
			s.baseSystems.InputSystem = inputSys
			continue
		}

		if objFactory, ok := s.Systems[idx].(*GameObjectFactory); ok && s.baseSystems.GameObjectFactory == nil {
			s.baseSystems.GameObjectFactory = objFactory
			continue
		}
	}

	if s.baseSystems.RenderSystem == nil {
		s.Info("waiting for render system ...")
		return
	}

	if s.baseSystems.RenderSystem.renderer == nil {
		s.Info("waiting for renderer instance ...")
		return
	}

	if s.baseSystems.InputSystem == nil {
		s.Info("waiting for input system")
		return
	}

	if s.baseSystems.GameObjectFactory == nil {
		s.Info("waiting for game object factory ...")
		return
	}

	s.setupFactories()

	s.baseSystems.SpriteFactory.RenderSystem = s.baseSystems.RenderSystem

	const (
		defaultWidth  = 800
		defaultHeight = 600
	)

	s.Add.Viewport(mainViewport, defaultWidth, defaultHeight)

	s.Info("base scene booted!")
	s.booted = true
}

func (s *BaseScene) setupFactories() {
	s.Info("setting up component factories")

	s.InjectComponent(&d2components.MainViewport{}, &s.MainViewport)
	s.InjectComponent(&d2components.Viewport{}, &s.Viewport)
	s.InjectComponent(&d2components.ViewportFilter{}, &s.ViewportFilter)
	s.InjectComponent(&d2components.Camera{}, &s.Camera)
	s.InjectComponent(&d2components.Priority{}, &s.Priority)
	s.InjectComponent(&d2components.Texture{}, &s.Texture)
	s.InjectComponent(&d2components.Interactive{}, &s.Interactive)
	s.InjectComponent(&d2components.Position{}, &s.Position)
	s.InjectComponent(&d2components.Scale{}, &s.Scale)
	s.InjectComponent(&d2components.Origin{}, &s.Origin)
	s.InjectComponent(&d2components.Alpha{}, &s.Alpha)
	s.InjectComponent(&d2components.SceneGraphNode{}, &s.SceneGraphNode)
	s.InjectComponent(&d2components.DrawEffect{}, &s.DrawEffect)
	s.InjectComponent(&d2components.Sprite{}, &s.SpriteFactory.Sprite)
	s.InjectComponent(&d2components.Rectangle{}, &s.RectangleFactory.Rectangle)
	s.InjectComponent(&d2components.Color{}, &s.Color)
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

	s.Graph.UpdateWorldMatrix()
	s.renderViewports()
}

func (s *BaseScene) renderViewports() {
	if s.baseSystems.RenderSystem == nil {
		s.Warning("render system not present")
		return
	}

	if s.baseSystems.RenderSystem.renderer == nil {
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
}

func (s *BaseScene) binGameObjectsByViewport() map[int][]akara.EID {
	bins := make(map[int][]akara.EID)

	for _, eid := range s.GameObjects {
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
	id := s.Viewports[idx]

	// the first viewport is always the main viewport
	if idx == mainViewport {
		s.AddMainViewport(id)
	} else {
		s.MainViewport.Remove(id)
	}

	camera, found := s.GetCamera(id)
	if !found {
		return
	}

	node, found := s.GetSceneGraphNode(id)
	if !found {
		node = s.AddSceneGraphNode(id)
	}

	// translate the camera position using the camera's scene graph node
	cx, cy := camera.Position.Clone().ApplyMatrix4(node.Local).XY()
	cw, ch := camera.Size.XY()

	sfc, found := s.GetTexture(id)
	if !found {
		return
	}

	if sfc.Texture == nil {
		sfc.Texture = s.baseSystems.RenderSystem.renderer.NewSurface(int(cw), int(ch))
	}

	if idx == mainViewport {
		sfc.Texture.Clear(s.backgroundColor)
	}

	sfc.Texture.PushTranslation(int(-cx), int(-cy)) // negative because we're offsetting everything that gets rendered

	for _, object := range objects {
		s.renderObject(sfc.Texture, object)
	}

	sfc.Texture.Pop()
}

func (s *BaseScene) renderObject(target d2interface.Surface, id akara.EID) {
	texture, found := s.GetTexture(id)
	if !found {
		return
	}

	position, found := s.GetPosition(id)
	if !found {
		position = s.AddPosition(id)
	}

	scale, found := s.GetScale(id)
	if !found {
		scale = s.AddScale(id)
	}

	alpha, found := s.GetAlpha(id)
	if !found {
		alpha = s.AddAlpha(id)
	}

	origin, found := s.GetOrigin(id)
	if !found {
		origin = s.AddOrigin(id)
	}

	node, found := s.GetSceneGraphNode(id)
	if !found {
		node = s.AddSceneGraphNode(id)
		node.SetParent(s.Graph)
	}

	drawEffect, found := s.GetDrawEffect(id)
	if found {
		target.PushEffect(drawEffect.DrawEffect)
		defer target.Pop()
	}

	// translate the entity position using the scene graph node
	x, y := position.Clone().
		Add(origin.Vector3).
		ApplyMatrix4(node.Local).
		XY()

	target.PushTranslation(int(x), int(y))
	defer target.Pop()

	target.PushScale(scale.X(), scale.Y())
	defer target.Pop()

	const maxAlpha = 255

	target.PushColor(color.Alpha{A: uint8(alpha.Alpha * maxAlpha)})
	defer target.Pop()

	segment, found := s.baseSystems.SpriteFactory.GetSegmentedSprite(id)
	if found {
		s.renderSegmentedSprite(target, id, segment)
		return
	}

	target.Render(texture.Texture)
}

func (s *BaseScene) renderSegmentedSprite(target d2interface.Surface, id akara.EID, seg *d2components.SegmentedSprite) {
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
				s.Error("SetCurrentFrame error" + err.Error())
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
}
