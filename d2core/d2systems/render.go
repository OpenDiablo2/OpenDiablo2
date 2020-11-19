package d2systems

import (
	"errors"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"time"

	"github.com/gravestench/akara"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
	d2render "github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"
)

const (
	gameTitle = "Open Diablo 2"
	logPrefixRenderSystem = "Render System"
)

// NewRenderSystem creates a movement system
func NewRenderSystem() *RenderSystem {
	viewports := akara.NewFilter().
		Require(d2components.Viewport).
		Require(d2components.MainViewport).
		Require(d2components.Surface).
		Build()

	gameConfigs := akara.NewFilter().Require(d2components.GameConfig).Build()

	r := &RenderSystem{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(viewports, gameConfigs),
		Logger: d2util.NewLogger(),
	}

	r.SetPrefix(logPrefixRenderSystem)

	return r
}

// static check that RenderSystem implements the System interface
var _ akara.System = &RenderSystem{}

// RenderSystem handles entity movement based on velocity and position components
type RenderSystem struct {
	*akara.BaseSubscriberSystem
	*d2util.Logger
	renderer      d2interface.Renderer
	screenSurface d2interface.Surface
	viewports     *akara.Subscription
	configs       *akara.Subscription
	*d2components.GameConfigMap
	*d2components.ViewportMap
	*d2components.MainViewportMap
	*d2components.SurfaceMap
	lastUpdate    time.Time
	gameLoopInitDelay time.Duration // there is a race condition, this is a hack
}

// Init initializes the system with the given world
func (m *RenderSystem) Init(world *akara.World) {
	m.Info("initializing ...")

	m.gameLoopInitDelay = time.Millisecond

	m.viewports = m.Subscriptions[0]
	m.configs = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.GameConfigMap = m.InjectMap(d2components.GameConfig).(*d2components.GameConfigMap)
	m.ViewportMap = m.InjectMap(d2components.Viewport).(*d2components.ViewportMap)
	m.MainViewportMap = m.InjectMap(d2components.MainViewport).(*d2components.MainViewportMap)
	m.SurfaceMap = m.InjectMap(d2components.Surface).(*d2components.SurfaceMap)
}

// Process will create a renderer if it doesnt exist yet,
// and then
func (m *RenderSystem) Update() {
	if m.renderer != nil {
		return
	}

	m.createRenderer()
	m.SetActive(false)
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

	// d2render.CreateRenderer should use a GameConfigComponent instead ...
	oldStyleConfig := &d2config.Configuration{
		MpqLoadOrder:    config.MpqLoadOrder,
		Language:        config.Language,
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
		m.Error(err.Error())
		panic(err)
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

		sfc, found := m.GetSurface(id)
		if !found {
			return errors.New("main viewport doesn't have a surface")
		}

		screen.PushTranslation(vp.Left, vp.Top)
		screen.Render(sfc.Surface)
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

func (m *RenderSystem) Loop() error {
	m.Infof("entering game run loop ...")

	return m.renderer.Run(m.render, m.updateWorld, 800, 600, gameTitle)
}
