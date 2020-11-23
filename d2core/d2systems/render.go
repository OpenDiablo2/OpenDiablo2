package d2systems

import (
	"errors"
	"os"
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

// NewRenderSystem creates a new render system
func NewRenderSystem() *RenderSystem {
	viewports := akara.NewFilter().
		Require(d2components.Viewport).
		Require(d2components.MainViewport).
		Require(d2components.Renderable).
		Build()

	gameConfigs := akara.NewFilter().Require(d2components.GameConfig).Build()

	r := &RenderSystem{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(viewports, gameConfigs),
		Logger:               d2util.NewLogger(),
	}

	r.SetPrefix(logPrefixRenderSystem)

	return r
}

// static check that RenderSystem implements the System interface
var _ akara.System = &RenderSystem{}

// RenderSystem is responsible for rendering the main viewports of scenes
// to the game screen.
type RenderSystem struct {
	*akara.BaseSubscriberSystem
	*d2util.Logger
	renderer  d2interface.Renderer
	viewports *akara.Subscription
	configs   *akara.Subscription
	*d2components.GameConfigMap
	*d2components.ViewportMap
	*d2components.MainViewportMap
	*d2components.RenderableMap
	lastUpdate        time.Time
	gameLoopInitDelay time.Duration // there is a race condition, this is a hack
}

// Init initializes the system with the given world, injecting the necessary components
func (m *RenderSystem) Init(_ *akara.World) {
	m.Info("initializing ...")

	m.gameLoopInitDelay = time.Millisecond

	m.viewports = m.Subscriptions[0]
	m.configs = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.GameConfigMap = m.InjectMap(d2components.GameConfig).(*d2components.GameConfigMap)
	m.ViewportMap = m.InjectMap(d2components.Viewport).(*d2components.ViewportMap)
	m.MainViewportMap = m.InjectMap(d2components.MainViewport).(*d2components.MainViewportMap)
	m.RenderableMap = m.InjectMap(d2components.Renderable).(*d2components.RenderableMap)
}

// Update will initialize the renderer, start the game loop, and
// disable the system (to prevent it from being called during the game loop).
//
// The reason why this isn't in the init step is because we use other systems
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

	err := m.startGameLoop()
	if err != nil {
		m.Fatal(err.Error())
	}

	os.Exit(0)
}

func (m *RenderSystem) createRenderer() {
	m.Info("creating renderer instance")

	configs := m.configs.GetEntities()
	if len(configs) < 1 {
		return
	}

	config, found := m.GetGameConfig(configs[0])
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
	if m.gameLoopInitDelay > 0 {
		return nil
	}

	for _, id := range m.viewports.GetEntities() {
		vp, found := m.GetViewport(id)
		if !found {
			return errors.New("main viewport not found")
		}

		renderable, found := m.GetRenderable(id)
		if !found {
			return errors.New("main viewport doesn't have a surface")
		}

		if renderable.Surface == nil {
			renderable.Surface = m.renderer.NewSurface(vp.Width, vp.Height)
		}

		screen.PushTranslation(vp.Left, vp.Top)
		screen.Render(renderable.Surface)
		screen.Pop()
	}

	return nil
}

func (m *RenderSystem) updateWorld() error {
	currentTime := time.Now()
	elapsed := currentTime.Sub(m.lastUpdate)
	m.lastUpdate = currentTime

	if m.gameLoopInitDelay > 0 {
		m.gameLoopInitDelay -= elapsed
		return nil
	}

	return m.World.Update(elapsed)
}

func (m *RenderSystem) startGameLoop() error {
	m.Infof("starting game loop ...")

	return m.renderer.Run(m.render, m.updateWorld, 800, 600, gameTitle)
}
