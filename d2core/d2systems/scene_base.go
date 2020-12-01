package d2systems

import (
	"fmt"
	"image/color"

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
		systems:         &baseSystems{},
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

// BaseScene encapsulates common behaviors for systems that are considered "scenes",
// such as the main menu, the in-game map, the console, etc.
//
// The base scene is responsible for generic behaviors common to all scenes,
// like initializing the default viewport, or rendering game objects to the viewports.
type BaseScene struct {
	*akara.BaseSystem
	*d2util.Logger
	key             string
	booted          bool
	paused          bool
	systems         *baseSystems
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

	s.setupFactories()

	s.Add = &sceneObjectFactory{
		BaseScene: s,
		Logger:    d2util.NewLogger(),
	}

	s.Add.SetPrefix(fmt.Sprintf("%s -> %s", s.key, "Object Factory"))

	for idx := range s.Systems {
		if rendersys, ok := s.Systems[idx].(*RenderSystem); ok && s.systems.RenderSystem == nil {
			s.systems.RenderSystem = rendersys
			continue
		}

		if inputSys, ok := s.Systems[idx].(*InputSystem); ok && s.systems.InputSystem == nil {
			s.systems.InputSystem = inputSys
			continue
		}

		if objFactory, ok := s.Systems[idx].(*GameObjectFactory); ok && s.systems.GameObjectFactory == nil {
			s.systems.GameObjectFactory = objFactory
			continue
		}
	}

	if s.systems.RenderSystem == nil {
		s.Info("waiting for render system ...")
		return
	}

	if s.systems.RenderSystem.renderer == nil {
		s.Info("waiting for renderer instance ...")
		return
	}

	if s.systems.InputSystem == nil {
		s.Info("waiting for input system")
		return
	}

	if s.systems.GameObjectFactory == nil {
		s.Info("waiting for game object factory ...")
		return
	}

	s.systems.SpriteFactory.RenderSystem = s.systems.RenderSystem

	s.createDefaultViewport()

	s.Info("base scene booted!")
	s.booted = true
}

func (s *BaseScene) setupFactories() {
	s.Info("setting up component factories")

	mainViewportID := s.RegisterComponent(&d2components.MainViewport{})
	viewportID := s.RegisterComponent(&d2components.Viewport{})
	viewportFilterID := s.RegisterComponent(&d2components.ViewportFilter{})
	cameraID := s.RegisterComponent(&d2components.Camera{})
	priorityID := s.RegisterComponent(&d2components.Priority{})
	renderableID := s.RegisterComponent(&d2components.Texture{})
	interactiveID := s.RegisterComponent(&d2components.Interactive{})
	positionID := s.RegisterComponent(&d2components.Position{})
	scaleID := s.RegisterComponent(&d2components.Scale{})
	animationID := s.RegisterComponent(&d2components.Sprite{})
	originID := s.RegisterComponent(&d2components.Origin{})
	alphaID := s.RegisterComponent(&d2components.Alpha{})
	sceneGraphNodeID := s.RegisterComponent(&d2components.SceneGraphNode{})
	drawEffectID := s.RegisterComponent(&d2components.DrawEffect{})

	s.MainViewport = s.GetComponentFactory(mainViewportID)
	s.Viewport = s.GetComponentFactory(viewportID)
	s.ViewportFilter = s.GetComponentFactory(viewportFilterID)
	s.Camera = s.GetComponentFactory(cameraID)
	s.Priority = s.GetComponentFactory(priorityID)
	s.Texture = s.GetComponentFactory(renderableID)
	s.Interactive = s.GetComponentFactory(interactiveID)
	s.Position = s.GetComponentFactory(positionID)
	s.Scale = s.GetComponentFactory(scaleID)
	s.Sprite = s.GetComponentFactory(animationID)
	s.Origin = s.GetComponentFactory(originID)
	s.Alpha = s.GetComponentFactory(alphaID)
	s.SceneGraphNode = s.GetComponentFactory(sceneGraphNodeID)
	s.DrawEffect = s.GetComponentFactory(drawEffectID)
}

func (s *BaseScene) createDefaultViewport() {
	s.Info("creating default viewport")
	viewportID := s.NewEntity()
	s.AddViewport(viewportID)
	s.AddPriority(viewportID)

	camera := s.AddCamera(viewportID)
	width, height := camera.Size.XY()

	s.AddSceneGraphNode(viewportID).SetParent(s.Graph)

	sfc := s.systems.RenderSystem.renderer.NewSurface(int(width), int(height))

	sfc.Clear(color.Transparent)

	s.AddTexture(viewportID).Texture = sfc
	s.AddMainViewport(viewportID)

	s.Viewports = append(s.Viewports, viewportID)
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
	if s.systems.RenderSystem == nil {
		s.Warning("render system not present")
		return
	}

	if s.systems.RenderSystem.renderer == nil {
		s.Warning("render system doesn't have a renderer instance")
		return
	}

	numViewports := len(s.Viewports)

	if numViewports < 1 {
		s.createDefaultViewport()
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
		sfc.Texture = s.systems.RenderSystem.renderer.NewSurface(int(cw), int(ch))
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

	segment, found := s.systems.SpriteFactory.GetSegmentedSprite(id)
	if found {
		animation, found := s.GetSprite(id)
		if !found {
			return
		}

		var offsetY int

		segmentsX, segmentsY := segment.Xsegments, segment.Ysegments
		frameOffset := segment.FrameOffset

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

		return
	}

	target.Render(texture.Texture)
}
