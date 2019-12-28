package d2player

import (
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"
	"github.com/hajimehoshi/ebiten"
)

type GameControls struct {
	hero      *d2core.Hero
	mapEngine *d2mapengine.MapEngine

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
	}
}

func (g *GameControls) Move(tickTime float64) {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		px, py := g.mapEngine.ScreenToWorld(ebiten.CursorPosition())
		g.hero.AnimatedEntity.SetTarget(px*5, py*5, 1)
	}

	arrowDistance := 1.0
	moveX := 0.0
	moveY := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		moveY -= arrowDistance
		moveX -= arrowDistance
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		moveY += arrowDistance
		moveX += arrowDistance
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		moveY += arrowDistance
		moveX -= arrowDistance
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		moveY -= arrowDistance
		moveX += arrowDistance
	}

	if moveY != 0 || moveX != 0 {
		g.hero.AnimatedEntity.SetTarget(g.hero.AnimatedEntity.LocationX+moveX, g.hero.AnimatedEntity.LocationY+moveY, 1)
	}

}

func (g *GameControls) Load() {
	g.globeSprite, _ = d2render.LoadSprite(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	g.mainPanel, _ = d2render.LoadSprite(d2resource.GamePanels, d2resource.PaletteSky)
	g.menuButton, _ = d2render.LoadSprite(d2resource.MenuButton, d2resource.PaletteSky)
	g.skillIcon, _ = d2render.LoadSprite(d2resource.GenericSkills, d2resource.PaletteSky)
}

// TODO: consider caching the panels to single image that is reused.
func (g *GameControls) Render(target *d2surface.Surface) {
	width, height := target.GetSize()
	offset := int(0)

	// Left globe holder
	g.mainPanel.SetCurrentFrame(0)
	w, _ := g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(int(offset), height)
	g.mainPanel.Render(target)

	// Left globe
	g.globeSprite.SetCurrentFrame(0)
	g.globeSprite.SetPosition(int(offset+28), height-5)
	g.globeSprite.Render(target)
	offset += w

	// Left skill
	g.skillIcon.SetCurrentFrame(2)
	w, _ = g.skillIcon.GetCurrentFrameSize()
	g.skillIcon.SetPosition(int(offset), height)
	g.skillIcon.Render(target)
	offset += w

	// Left skill selector
	g.mainPanel.SetCurrentFrame(1)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(int(offset), height)
	g.mainPanel.Render(target)
	offset += w

	// Stamina
	g.mainPanel.SetCurrentFrame(2)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(int(offset), height)
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
	g.mainPanel.SetPosition(int(offset), height)
	g.mainPanel.Render(target)
	offset += w

	// Right skill selector
	g.mainPanel.SetCurrentFrame(4)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(int(offset), height)
	g.mainPanel.Render(target)
	offset += w

	// Right skill
	g.skillIcon.SetCurrentFrame(10)
	w, _ = g.skillIcon.GetCurrentFrameSize()
	g.skillIcon.SetPosition(int(offset), height)
	g.skillIcon.Render(target)
	offset += w

	// Right globe holder
	g.mainPanel.SetCurrentFrame(5)
	w, _ = g.mainPanel.GetCurrentFrameSize()
	g.mainPanel.SetPosition(int(offset), height)
	g.mainPanel.Render(target)

	// Right globe
	g.globeSprite.SetCurrentFrame(1)
	g.globeSprite.SetPosition(int(offset)+8, height-8)
	g.globeSprite.Render(target)

}
