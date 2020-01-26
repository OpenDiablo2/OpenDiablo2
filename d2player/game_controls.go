package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"
)

type Panel interface {
	IsOpen() bool
	Toggle()
	Open()
	Close()
}

type GameControls struct {
	hero      *d2core.Hero
	mapEngine *d2mapengine.MapEngine
	inventory *Inventory

	// UI
	globeSprite *d2render.Sprite
	mainPanel   *d2render.Sprite
	menuButton  *d2render.Sprite
	skillIcon   *d2render.Sprite
}

func NewGameControls(hero *d2core.Hero, mapEngine *d2mapengine.MapEngine) *GameControls {
	return &GameControls{
		hero:      hero,
		mapEngine: mapEngine,
		inventory: NewInventory(),
	}
}

func (g *GameControls) OnKeyDown(event d2input.KeyEvent) bool {
	if event.Key == d2input.KeyI {
		g.inventory.Toggle()
		return true
	}

	return false
}

func (g *GameControls) OnMouseButtonDown(event d2input.MouseEvent) bool {
	if event.Button == d2input.MouseButtonLeft {
		px, py := g.mapEngine.ScreenToWorld(event.X, event.Y)
		g.hero.AnimatedEntity.SetTarget(px*5, py*5, 1)
		return true
	}

	return false
}

func (g *GameControls) Load() {
	g.globeSprite, _ = d2render.LoadSprite(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	g.mainPanel, _ = d2render.LoadSprite(d2resource.GamePanels, d2resource.PaletteSky)
	g.menuButton, _ = d2render.LoadSprite(d2resource.MenuButton, d2resource.PaletteSky)
	g.skillIcon, _ = d2render.LoadSprite(d2resource.GenericSkills, d2resource.PaletteSky)
	g.inventory.Load()
}

// TODO: consider caching the panels to single image that is reused.
func (g *GameControls) Render(target *d2surface.Surface) {
	g.inventory.Render(target)

	width, height := target.GetSize()
	offset := int(0)

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
