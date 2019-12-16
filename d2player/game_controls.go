package d2player

import (
	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"
	"github.com/hajimehoshi/ebiten"
	"log"
)

type GameControls struct {
	fileProvider d2interface.FileProvider
	hero         *d2core.Hero
	mapEngine    *d2mapengine.MapEngine

	// UI
	globeSprite *d2render.Sprite
	mainPanel   *d2render.Sprite
	menuButton  *d2render.Sprite
	skillIcon   *d2render.Sprite
}

func NewGameControls(fileProvider d2interface.FileProvider, hero *d2core.Hero, mapEngine *d2mapengine.MapEngine) *GameControls {
	return &GameControls{
		fileProvider: fileProvider,
		hero:         hero,
		mapEngine:    mapEngine,
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
	dc6, err := d2dc6.LoadDC6(g.fileProvider.LoadFile(d2resource.GameGlobeOverlap), d2datadict.Palettes[d2enum.Sky])
	if err != nil {
		log.Panicf("failed to load %s: %v", d2resource.GameGlobeOverlap, err)
	}
	globeSprite := d2render.CreateSpriteFromDC6(dc6)
	g.globeSprite = &globeSprite

	dc6, err = d2dc6.LoadDC6(g.fileProvider.LoadFile(d2resource.GamePanels), d2datadict.Palettes[d2enum.Sky])
	if err != nil {
		log.Panicf("failed to load %s: %v", d2resource.GamePanels, err)
	}
	mainPanel := d2render.CreateSpriteFromDC6(dc6)
	g.mainPanel = &mainPanel

	dc6, err = d2dc6.LoadDC6(g.fileProvider.LoadFile(d2resource.MenuButton), d2datadict.Palettes[d2enum.Sky])
	if err != nil {
		log.Panicf("failed to load %s: %v", d2resource.MenuButton, err)
	}
	menuButton := d2render.CreateSpriteFromDC6(dc6)
	g.menuButton = &menuButton

	dc6, err = d2dc6.LoadDC6(g.fileProvider.LoadFile(d2resource.GenericSkills), d2datadict.Palettes[d2enum.Sky])
	if err != nil {
		log.Panicf("failed to load %s: %v", d2resource.GenericSkills, err)
	}
	skillIcon := d2render.CreateSpriteFromDC6(dc6)
	g.skillIcon = &skillIcon

}


// TODO: consider caching the panels to single image that is reused.
func (g *GameControls) Render(target *ebiten.Image) {
	width, height := target.Size()
	offset := uint32(0)

	// Left globe holder
	g.mainPanel.Frame = 0
	w, _ := g.mainPanel.GetSize()
	g.mainPanel.MoveTo(int(offset), height)
	g.mainPanel.Draw(target)

	// Left globe
	g.globeSprite.Frame = 0
	g.globeSprite.MoveTo(int(offset+28), height - 5)
	g.globeSprite.Draw(target)
	offset += w

	// Left skill
	g.skillIcon.Frame = 2
	w, _ = g.skillIcon.GetSize()
	g.skillIcon.MoveTo(int(offset), height)
	g.skillIcon.Draw(target)
	offset += w

	// Left skill selector
	g.mainPanel.Frame = 1
	w, _ = g.mainPanel.GetSize()
	g.mainPanel.MoveTo(int(offset), height)
	g.mainPanel.Draw(target)
	offset += w

	// Stamina
	g.mainPanel.Frame = 2
	w, _ = g.mainPanel.GetSize()
	g.mainPanel.MoveTo(int(offset), height)
	g.mainPanel.Draw(target)
	offset += w

	// Center menu button
	g.menuButton.Frame = 0
	w, _ = g.mainPanel.GetSize()
	g.menuButton.MoveTo((width / 2) - 8 , height - 16)
	g.menuButton.Draw(target)

	// Potions
	g.mainPanel.Frame = 3
	w, _ = g.mainPanel.GetSize()
	g.mainPanel.MoveTo(int(offset), height)
	g.mainPanel.Draw(target)
	offset += w

	// Right skill selector
	g.mainPanel.Frame = 4
	w, _ = g.mainPanel.GetSize()
	g.mainPanel.MoveTo(int(offset), height)
	g.mainPanel.Draw(target)
	offset += w

	// Right skill
	g.skillIcon.Frame = 10
	w, _ = g.skillIcon.GetSize()
	g.skillIcon.MoveTo(int(offset), height)
	g.skillIcon.Draw(target)
	offset += w

	// Right globe holder
	g.mainPanel.Frame = 5
	w, _ = g.mainPanel.GetSize()
	g.mainPanel.MoveTo(int(offset), height)
	g.mainPanel.Draw(target)

	// Right globe
	g.globeSprite.Frame = 1
	g.globeSprite.MoveTo(int(offset) + 8, height - 8)
	g.globeSprite.Draw(target)

}
