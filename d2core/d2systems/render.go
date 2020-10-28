package d2systems

import (
	"time"

	"github.com/gravestench/akara"
	"github.com/hajimehoshi/ebiten"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
	d2render "github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"
)

const (
	gameTitle = "Open Diablo 2"
)

// NewRenderSystem creates a movement system
func NewRenderSystem() *RenderSystem {
	viewports := akara.NewFilter().Require(d2components.ViewPort).Build()
	gameConfigs := akara.NewFilter().Require(d2components.GameConfig).Build()

	return &RenderSystem{
		SubscriberSystem: akara.NewSubscriberSystem(viewports, gameConfigs),
	}
}

// static check that RenderSystem implements the System interface
var _ akara.System = &RenderSystem{}

// RenderSystem handles entity movement based on velocity and position components
type RenderSystem struct {
	*akara.SubscriberSystem
	renderer      d2interface.Renderer
	screenSurface d2interface.Surface
	viewports     *akara.Subscription
	configs       *akara.Subscription
	configMap     *d2components.GameConfigMap
	viewportMap   *d2components.ViewPortMap
	lastUpdate    time.Time
}

// Init initializes the system with the given world
func (m *RenderSystem) Init(world *akara.World) {
	m.World = world

	if world == nil {
		m.SetActive(false)
		return
	}

	for subIdx := range m.Subscriptions {
		m.Subscriptions[subIdx] = m.AddSubscription(m.Subscriptions[subIdx].Filter)
	}

	m.viewports = m.Subscriptions[0]
	m.configs = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.configMap = m.InjectMap(d2components.GameConfig).(*d2components.GameConfigMap)
	m.viewportMap = m.InjectMap(d2components.ViewPort).(*d2components.ViewPortMap)
}

// Process will create a renderer if it doesnt exist yet,
// and then
func (m *RenderSystem) Process() {
	if m.renderer == nil {
		m.createRenderer()
	}

	if m.screenSurface == nil {
		return
	}

	for _, eid := range m.viewports.GetEntities() {
		m.render(eid)
	}
}

func (m *RenderSystem) createRenderer() {
	configs := m.configs.GetEntities()
	if len(configs) < 1 {
		return
	}

	config, found := m.configMap.GetGameConfig(configs[0])
	if !found {
		return
	}

	renderer, err := d2render.CreateRenderer()
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

	_ = m.renderer.Run(m.wrapWorldUpdate, 800, 600, gameTitle)
}

func (m *RenderSystem) render(id akara.EID) {
	vp, found := m.viewportMap.GetViewPort(id)
	if !found {
		return
	}

	if m.screenSurface != nil {
		_ = m.screenSurface.Render(vp.Surface)
	}
}

func (m *RenderSystem) wrapWorldUpdate(s d2interface.Surface) error {
	m.screenSurface = s

	currentTime := time.Now()
	elapsed := currentTime.Sub(m.lastUpdate)

	return m.World.Update(elapsed)
}
