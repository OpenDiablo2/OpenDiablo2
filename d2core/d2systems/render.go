package d2systems

import (
	"errors"
	"image/color"
	"os"
	"sort"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"

	"github.com/gravestench/akara"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
	d2render "github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"
)

const (
	gameTitle             = "Open Diablo 2"
	logPrefixRenderSystem = "Render System"
)

// static check that RenderSystem implements the System interface
var _ akara.System = &RenderSystem{}

// RenderSystem is responsible for rendering the main viewports of scenes
// to the game screen.
type RenderSystem struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	renderer   d2interface.Renderer
	viewports  *akara.Subscription
	configs    *akara.Subscription
	lastUpdate time.Time
	Components struct {
		GameConfig   d2components.GameConfigFactory
		Viewport     d2components.ViewportFactory
		MainViewport d2components.MainViewportFactory
		Texture      d2components.TextureFactory
		Priority     d2components.PriorityFactory
		Alpha        d2components.AlphaFactory
		Camera       d2components.CameraFactory
	}
}

// Init initializes the system with the given world, injecting the necessary components
func (m *RenderSystem) Init(world *akara.World) {
	m.World = world

	m.lastUpdate = time.Now()

	m.setupLogger()

	m.Debug("initializing ...")

	m.setupFactories()
	m.setupSubscriptions()
}

func (m *RenderSystem) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixRenderSystem)
}

func (m *RenderSystem) setupFactories() {
	m.InjectComponent(&d2components.GameConfig{}, &m.Components.GameConfig.ComponentFactory)
	m.InjectComponent(&d2components.Viewport{}, &m.Components.Viewport.ComponentFactory)
	m.InjectComponent(&d2components.MainViewport{}, &m.Components.MainViewport.ComponentFactory)
	m.InjectComponent(&d2components.Texture{}, &m.Components.Texture.ComponentFactory)
	m.InjectComponent(&d2components.Priority{}, &m.Components.Priority.ComponentFactory)
	m.InjectComponent(&d2components.Alpha{}, &m.Components.Alpha.ComponentFactory)
	m.InjectComponent(&d2components.Camera{}, &m.Components.Camera.ComponentFactory)
}

func (m *RenderSystem) setupSubscriptions() {
	viewports := m.NewComponentFilter().
		Require(
			&d2components.Viewport{},
			&d2components.MainViewport{},
			&d2components.Texture{},
			&d2components.Camera{},
		).
		Build()

	gameConfigs := m.NewComponentFilter().
		Require(&d2components.GameConfig{}).
		Build()

	m.viewports = m.AddSubscription(viewports)
	m.configs = m.AddSubscription(gameConfigs)
}

// Update will initialize the renderer, start the game loop, and
// disable the system (to prevent it from being called during the game loop).
//
// The reason why this isn't in the init step is because we use other sceneSystems
// for loading the config file, and it may take more than one iteration
func (m *RenderSystem) Update() {
	if m.renderer != nil {
		return // we already created the renderer
	}

	m.createRenderer()

	if m.renderer == nil {
		return // the renderer has not yet been created!
	}

	// if we have created the renderer, we can safely disable
	// this system and start the run loop.
	m.SetActive(false)

	err := m.StartGameLoop()
	if err != nil {
		m.Fatal(err.Error())
	}

	os.Exit(0)
}

func (m *RenderSystem) createRenderer() {
	m.Debug("creating renderer instance")

	configs := m.configs.GetEntities()
	if len(configs) < 1 {
		return
	}

	config, found := m.Components.GameConfig.Get(configs[0])
	if !found {
		return
	}

	// we should get rid of d2config.Configuration and use components instead...
	oldStyleConfig := &d2config.Configuration{
		MpqLoadOrder:    config.MpqLoadOrder,
		MpqPath:         config.MpqPath,
		TicksPerSecond:  config.TicksPerSecond,
		FpsCap:          config.FpsCap,
		SfxVolume:       config.SfxVolume,
		BgmVolume:       config.BgmVolume,
		FullScreen:      config.FullScreen,
		RunInBackground: config.RunInBackground,
		VsyncEnabled:    config.VsyncEnabled,
		Backend:         config.Backend,
		LogLevel:        config.LogLevel,
	}

	renderer, err := d2render.CreateRenderer(oldStyleConfig)
	if err != nil {
		m.Fatal(err.Error())
	}

	// HACK: hardcoded with ebiten for now
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetFullscreen(config.FullScreen)
	ebiten.SetRunnableOnUnfocused(config.RunInBackground)
	ebiten.SetVsyncEnabled(config.VsyncEnabled)
	ebiten.SetMaxTPS(config.TicksPerSecond)

	m.renderer = renderer
}

func (m *RenderSystem) render(screen d2interface.Surface) error {
	entities := m.viewports.GetEntities()

	sort.Slice(entities, func(i, j int) bool {
		pi, pj := m.Components.Priority.Add(entities[i]), m.Components.Priority.Add(entities[j])
		return pi.Priority < pj.Priority
	})

	for _, id := range entities {
		vp, found := m.Components.Viewport.Get(id)
		if !found {
			return errors.New("main viewport not found")
		}

		cam, found := m.Components.Camera.Get(id)
		if !found {
			return errors.New("main viewport camera not found")
		}

		texture, found := m.Components.Texture.Get(id)
		if !found {
			return errors.New("main viewport doesn't have a surface")
		}

		if texture.Texture == nil {
			texture.Texture = m.renderer.NewSurface(vp.Width, vp.Height)
		}

		alpha, found := m.Components.Alpha.Get(id)
		if !found {
			alpha = m.Components.Alpha.Add(id)
		}

		const maxAlpha = 255

		screen.PushColor(color.Alpha{A: uint8(alpha.Alpha * maxAlpha)})
		screen.PushTranslation(vp.Left, vp.Top)
		screen.PushScale(float64(vp.Width)/cam.Size.X, float64(vp.Height)/cam.Size.Y)

		screen.Render(texture.Texture)

		screen.Pop()
		screen.Pop()
		screen.Pop()
	}

	return nil
}

func (m *RenderSystem) updateWorld() error {
	currentTime := time.Now()
	elapsed := currentTime.Sub(m.lastUpdate)
	m.lastUpdate = currentTime

	return m.World.Update(elapsed)
}

func (m *RenderSystem) StartGameLoop() error {
	m.Info("starting game loop ...")

	return m.renderer.Run(m.render, m.updateWorld, 800, 600, gameTitle)
}
