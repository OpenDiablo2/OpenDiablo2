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
	viewports := akara.NewFilter().Require(d2components.ViewPort).Build()
	gameConfigs := akara.NewFilter().Require(d2components.GameConfig).Build()

	r := &RenderSystem{
		SubscriberSystem: akara.NewSubscriberSystem(viewports, gameConfigs),
		Logger: d2util.NewLogger(),
	}

	r.SetPrefix(logPrefixRenderSystem)

	return r
}

// static check that RenderSystem implements the System interface
var _ akara.System = &RenderSystem{}

// RenderSystem handles entity movement based on velocity and position components
type RenderSystem struct {
	*akara.SubscriberSystem
	*d2util.Logger
	renderer      d2interface.Renderer
	screenSurface d2interface.Surface
	viewports     *akara.Subscription
	configs       *akara.Subscription
	*d2components.GameConfigMap
	*d2components.ViewPortMap
	lastUpdate    time.Time
}

// Init initializes the system with the given world
func (m *RenderSystem) Init(world *akara.World) {
	m.World = world

	if world == nil {
		m.SetActive(false)
		return
	}

	m.Info("initializing ...")

	for subIdx := range m.Subscriptions {
		m.Subscriptions[subIdx] = m.AddSubscription(m.Subscriptions[subIdx].Filter)
	}

	m.viewports = m.Subscriptions[0]
	m.configs = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.GameConfigMap = m.InjectMap(d2components.GameConfig).(*d2components.GameConfigMap)
	m.ViewPortMap = m.InjectMap(d2components.ViewPort).(*d2components.ViewPortMap)
}

// Process will create a renderer if it doesnt exist yet,
// and then
func (m *RenderSystem) Process() {
	if m.renderer == nil {
		m.createRenderer()
	}
}

func (m *RenderSystem) createRenderer() {
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
		panic(err)
	}

	// HACK: hardcoded with ebiten for now
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetFullscreen(config.FullScreen)
	ebiten.SetRunnableOnUnfocused(config.RunInBackground)
	ebiten.SetVsyncEnabled(config.VsyncEnabled)
	ebiten.SetMaxTPS(config.TicksPerSecond)

	m.renderer = renderer

	m.lastUpdate = time.Now()

	_ = m.renderer.Run(m.render, m.updateWorld, 800, 600, gameTitle)
}

func (m *RenderSystem) render(screen d2interface.Surface) error {
	m.screenSurface = screen

	for _, id := range m.viewports.GetEntities() {
		vp, found := m.GetViewPort(id)
		if !found {
			return errors.New("viewport not found")
		}

		if m.screenSurface != nil {
			m.screenSurface.Render(vp.Surface)
		}
	}

	return nil
}

func (m *RenderSystem) updateWorld() error {
	currentTime := time.Now()
	elapsed := currentTime.Sub(m.lastUpdate)

	return m.World.Update(elapsed)
}
