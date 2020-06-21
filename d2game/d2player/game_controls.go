package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Panel interface {
	IsOpen() bool
	Toggle()
	Open()
	Close()
}

// ID of missile to create when user right clicks.
var missileID = 59

type GameControls struct {
	hero          *d2map.Player
	mapEngine     *d2map.MapEngine
	mapRenderer   *d2map.MapRenderer
	inventory     *Inventory
	heroStats     *HeroStats
	escapeMenu    *EscapeMenu
	inputListener InputCallbackListener

	// UI
	globeSprite *d2ui.Sprite
	mainPanel   *d2ui.Sprite
	menuButton  *d2ui.Sprite
	skillIcon   *d2ui.Sprite
}

func NewGameControls(hero *d2map.Player, mapEngine *d2map.MapEngine, mapRenderer *d2map.MapRenderer, inputListener InputCallbackListener) *GameControls {
	d2term.BindAction("setmissile", "set missile id to summon on right click", func(id int) {
		missileID = id
	})

	return &GameControls{
		hero:          hero,
		mapEngine:     mapEngine,
		inputListener: inputListener,
		mapRenderer:   mapRenderer,
		inventory:     NewInventory(),
		heroStats:     NewHeroStats(),
		escapeMenu:    NewEscapeMenu(),
	}
}

func (g *GameControls) OnKeyDown(event d2input.KeyEvent) bool {
	if event.Key == d2input.KeyEscape {
		g.escapeMenu.Toggle()
		return true
	}
	if event.Key == d2input.KeyI {
		g.inventory.Toggle()
		return true
	}
	if event.Key == d2input.KeyC {
		g.heroStats.Toggle()
		return true
	}

	return false
}

func (g *GameControls) OnMouseMove(event d2input.MouseMoveEvent) bool {
	g.escapeMenu.OnMouseMove(event)
	return false
}

func (g *GameControls) OnMouseButtonDown(event d2input.MouseEvent) bool {
	if g.escapeMenu.IsOpen() {
		return g.escapeMenu.OnMouseButtonDown(event)
	}

	px, py := g.mapRenderer.ScreenToWorld(event.X, event.Y)
	px = float64(int(px*10)) / 10.0
	py = float64(int(py*10)) / 10.0

	if event.Button == d2input.MouseButtonLeft {
		g.inputListener.OnPlayerMove(px, py)
		return true
	}

	if event.Button == d2input.MouseButtonRight {
		missile, err := d2map.CreateMissile(
			int(g.hero.AnimatedComposite.LocationX),
			int(g.hero.AnimatedComposite.LocationY),
			d2datadict.Missiles[missileID],
		)
		if err != nil {
			return false
		}

		rads := d2common.GetRadiansBetween(
			g.hero.AnimatedComposite.LocationX,
			g.hero.AnimatedComposite.LocationY,
			px*5,
			py*5,
		)
		missile.SetRadians(rads, func() {
			g.mapEngine.RemoveEntity(missile)
		})

		g.mapEngine.AddEntity(missile)
		return true
	}

	return false
}

func (g *GameControls) Load() {
	animation, _ := d2asset.LoadAnimation(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	g.globeSprite, _ = d2ui.LoadSprite(animation)

	animation, _ = d2asset.LoadAnimation(d2resource.GamePanels, d2resource.PaletteSky)
	g.mainPanel, _ = d2ui.LoadSprite(animation)

	animation, _ = d2asset.LoadAnimation(d2resource.MenuButton, d2resource.PaletteSky)
	g.menuButton, _ = d2ui.LoadSprite(animation)

	animation, _ = d2asset.LoadAnimation(d2resource.GenericSkills, d2resource.PaletteSky)
	g.skillIcon, _ = d2ui.LoadSprite(animation)

	g.inventory.Load()
	g.heroStats.Load()
	g.escapeMenu.OnLoad()
}

// ScreenAdvanceHandler
func (g *GameControls) Advance(elapsed float64) error {
	g.escapeMenu.Advance(elapsed)
	return nil
}

// TODO: consider caching the panels to single image that is reused.
func (g *GameControls) Render(target d2render.Surface) {
	g.inventory.Render(target)
	g.heroStats.Render(target)
	g.escapeMenu.Render(target)

	width, height := target.GetSize()
	offset := 0

	// Left globe holder
	g.mainPanel.SetCurrentFrame(0)
	w, _ := g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(offset, height)
	g.mainPanel.Render(target)

	// Left globe
	g.globeSprite.SetCurrentFrame(0)
	g.globeSprite.SetPosition(offset+28, height-5)
	g.globeSprite.Render(target)
	offset += w

	// Left skill
	g.skillIcon.SetCurrentFrame(2)
	w, _ = g.skillIcon.GetCurrentFrameSize()
	g.skillIcon.SetPosition(offset, height)
	g.skillIcon.Render(target)
	offset += w

	// Left skill selector
	g.mainPanel.SetCurrentFrame(1)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(offset, height)
	g.mainPanel.Render(target)
	offset += w

	// Stamina
	g.mainPanel.SetCurrentFrame(2)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(offset, height)
	g.mainPanel.Render(target)
	offset += w

	// Center menu button
	g.menuButton.SetCurrentFrame(0)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.menuButton.SetPosition((width/2)-8, height-16)
	g.menuButton.Render(target)

	// Potions
	g.mainPanel.SetCurrentFrame(3)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(offset, height)
	g.mainPanel.Render(target)
	offset += w

	// Right skill selector
	g.mainPanel.SetCurrentFrame(4)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(offset, height)
	g.mainPanel.Render(target)
	offset += w

	// Right skill
	g.skillIcon.SetCurrentFrame(10)
	w, _ = g.skillIcon.GetCurrentFrameSize()
	g.skillIcon.SetPosition(offset, height)
	g.skillIcon.Render(target)
	offset += w

	// Right globe holder
	g.mainPanel.SetCurrentFrame(5)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(offset, height)
	g.mainPanel.Render(target)

	// Right globe
	g.globeSprite.SetCurrentFrame(1)
	g.globeSprite.SetPosition(offset+8, height-8)
	g.globeSprite.Render(target)

}

func (g *GameControls) InEscapeMenu() bool {
	return g != nil && g.escapeMenu != nil && g.escapeMenu.IsOpen()
}
