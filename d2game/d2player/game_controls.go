package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"
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
	hero          *d2mapentity.Player
	mapEngine     *d2mapengine.MapEngine
	mapRenderer   *d2maprenderer.MapRenderer
	inventory     *Inventory
	heroStats     *HeroStats
	escapeMenu    *EscapeMenu
	inputListener InputCallbackListener
	FreeCam       bool

	// UI
	globeSprite *d2ui.Sprite
	mainPanel   *d2ui.Sprite
	menuButton  *d2ui.Sprite
	skillIcon   *d2ui.Sprite
}

func NewGameControls(hero *d2mapentity.Player, mapEngine *d2mapengine.MapEngine, mapRenderer *d2maprenderer.MapRenderer, inputListener InputCallbackListener) *GameControls {
	d2term.BindAction("setmissile", "set missile id to summon on right click", func(id int) {
		missileID = id
	})

	gc := &GameControls{
		hero:          hero,
		mapEngine:     mapEngine,
		inputListener: inputListener,
		mapRenderer:   mapRenderer,
		inventory:     NewInventory(),
		heroStats:     NewHeroStats(),
		escapeMenu:    NewEscapeMenu(),
	}

	d2term.BindAction("freecam", "toggle free camera movement", func() {
		gc.FreeCam = !gc.FreeCam
	})

	return gc
}

func (g *GameControls) OnKeyRepeat(event d2input.KeyEvent) bool {
	if g.FreeCam {
		var moveSpeed float64 = 8
		if event.KeyMod == d2input.KeyModShift {
			moveSpeed *= 2
		}

		if event.Key == d2input.KeyDown {
			g.mapRenderer.MoveCameraBy(0, moveSpeed)
			return true
		}

		if event.Key == d2input.KeyUp {
			g.mapRenderer.MoveCameraBy(0, -moveSpeed)
			return true
		}

		if event.Key == d2input.KeyRight {
			g.mapRenderer.MoveCameraBy(moveSpeed, 0)
			return true
		}

		if event.Key == d2input.KeyLeft {
			g.mapRenderer.MoveCameraBy(-moveSpeed, 0)
			return true
		}
	}

	return false
}

func (g *GameControls) OnKeyDown(event d2input.KeyEvent) bool {
	switch event.Key {
	case d2input.KeyEscape:
		if g.inventory.IsOpen() || g.heroStats.IsOpen() {
			g.inventory.Close()
			g.heroStats.Close()
			g.updateLayout()
			break
		}
		g.escapeMenu.Toggle()
	case d2input.KeyUp:
		g.escapeMenu.OnUpKey()
	case d2input.KeyDown:
		g.escapeMenu.OnDownKey()
	case d2input.KeyEnter:
		g.escapeMenu.OnEnterKey()
	case d2input.KeyI:
		g.inventory.Toggle()
		g.updateLayout()
	case d2input.KeyC:
		g.heroStats.Toggle()
		g.updateLayout()
	default:
		return false
	}

	return true
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
		missile, err := d2mapentity.CreateMissile(
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

func (g *GameControls) updateLayout() {
	isRightPanelOpen := false
	isLeftPanelOpen := false

	// todo : add same logic when adding quest log and skill tree
	isRightPanelOpen = g.inventory.isOpen || isRightPanelOpen
	isLeftPanelOpen = g.heroStats.isOpen || isLeftPanelOpen

	if isRightPanelOpen == isLeftPanelOpen {
		g.mapRenderer.ViewportDefault()
	} else if isRightPanelOpen == true {
		g.mapRenderer.ViewportToLeft()
	} else {
		g.mapRenderer.ViewportToRight()
	}
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
