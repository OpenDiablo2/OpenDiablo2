package d2systems

import (
	"fmt"
	"image/color"
	"path/filepath"

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
		BaseSystem:  &akara.BaseSystem{},
		Logger:      d2util.NewLogger(),
		key:         key,
		Viewports:   make([]akara.EID, 0),
		GameObjects: make([]akara.EID, 0),
		systems:     &baseSystems{},
	}

	base.SetPrefix(key)

	return base
}

var _ akara.System = &BaseScene{}

type baseSystems struct {
	*RenderSystem
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
	key         string
	booted      bool
	paused      bool
	systems     *baseSystems
	Add         *sceneObjectFactory
	Viewports   []akara.EID
	GameObjects []akara.EID
	d2components.ViewportFactory
	d2components.MainViewportFactory
	d2components.ViewportFilterFactory
	d2components.CameraFactory
	d2components.RenderableFactory
	d2components.PositionFactory
	d2components.ScaleFactory
	d2components.AnimationFactory
	d2components.OriginFactory
	d2components.AlphaFactory
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
		if rendersys, ok := s.Systems[idx].(*RenderSystem); ok {
			s.systems.RenderSystem = rendersys
			continue
		}

		if objFactory, ok := s.Systems[idx].(*GameObjectFactory); ok {
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
	renderableID := s.RegisterComponent(&d2components.Renderable{})
	positionID := s.RegisterComponent(&d2components.Position{})
	scaleID := s.RegisterComponent(&d2components.Scale{})
	animationID := s.RegisterComponent(&d2components.Animation{})
	originID := s.RegisterComponent(&d2components.Origin{})
	alphaID := s.RegisterComponent(&d2components.Alpha{})

	s.MainViewport = s.GetComponentFactory(mainViewportID)
	s.Viewport = s.GetComponentFactory(viewportID)
	s.ViewportFilter = s.GetComponentFactory(viewportFilterID)
	s.Camera = s.GetComponentFactory(cameraID)
	s.Renderable = s.GetComponentFactory(renderableID)
	s.Position = s.GetComponentFactory(positionID)
	s.Scale = s.GetComponentFactory(scaleID)
	s.Animation = s.GetComponentFactory(animationID)
	s.Origin = s.GetComponentFactory(originID)
	s.Alpha = s.GetComponentFactory(alphaID)
}

func (s *BaseScene) createDefaultViewport() {
	s.Info("creating default viewport")
	viewportID := s.NewEntity()
	s.AddViewport(viewportID)

	camera := s.AddCamera(viewportID)
	camera.Width = 800
	camera.Height = 600
	camera.Zoom = 1

	sfc := s.systems.renderer.NewSurface(camera.Width, camera.Height)

	sfc.Clear(color.Transparent)

	s.AddRenderable(viewportID).Surface = sfc
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

	if idx == mainViewport {
		s.AddMainViewport(id)
	} else {
		s.MainViewport.Remove(id)
	}

	camera, found := s.GetCamera(id)
	if !found {
		return
	}

	sfc, found := s.GetRenderable(id)
	if !found {
		return
	}

	if sfc.Surface == nil {
		sfc.Surface = s.systems.renderer.NewSurface(camera.Width, camera.Height)
	}

	cx, cy := int(camera.X()), int(camera.Y())

	sfc.Surface.PushTranslation(cx, cy) // negative because we're offsetting everything that gets rendered
	sfc.Surface.PushScale(camera.Zoom, camera.Zoom)

	for _, object := range objects {
		s.renderObject(sfc.Surface, object)
	}

	sfc.Pop()
	sfc.Pop()
}

func (s *BaseScene) renderObject(target d2interface.Surface, id akara.EID) {
	renderable, found := s.GetRenderable(id)
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

	x, y := int(position.X()), int(position.Y())

	target.PushTranslation(x, y)
	defer target.Pop()

	target.PushScale(scale.X(), scale.Y())
	defer target.Pop()

	segment, found := s.systems.SpriteFactory.GetSegmentedSprite(id)
	if found {
		animation, found := s.GetAnimation(id)
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

	target.Render(renderable.Surface)
}

// responsible for wrapping the object factory calls and assigning the created object entity id's to the scene
type sceneObjectFactory struct {
	*BaseScene
	*d2util.Logger
}

func (s *sceneObjectFactory) addBasicComponenets(id akara.EID) {
	_ = s.AddScale(id)
	_ = s.AddOrigin(id)
	_ = s.AddPosition(id)
	_ = s.AddAlpha(id)
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
