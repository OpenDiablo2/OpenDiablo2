package d2player

import (
	"image/color"
	"time"

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
	inputListener InputCallbackListener

	// UI
	globeSprite     *d2ui.Sprite
	mainPanel       *d2ui.Sprite
	menuButton      *d2ui.Sprite
	skillIcon       *d2ui.Sprite
	zoneChangeText  *d2ui.Label
	isZoneTextShown bool
}

func NewGameControls(hero *d2map.Player, mapEngine *d2map.MapEngine, mapRenderer *d2map.MapRenderer, inputListener InputCallbackListener) *GameControls {
	d2term.BindAction("setmissile", "set missile id to summon on right click", func(id int) {
		missileID = id
	})

	label := d2ui.CreateLabel(d2resource.Font30, d2resource.PaletteUnits)
	label.Color = color.RGBA{R: 255, G: 88, B: 82, A: 255}
	label.Alignment = d2ui.LabelAlignCenter

	return &GameControls{
		hero:           hero,
		mapEngine:      mapEngine,
		inputListener:  inputListener,
		mapRenderer:    mapRenderer,
		inventory:      NewInventory(),
		heroStats:      NewHeroStats(),
		zoneChangeText: &label,
	}
}

func (g *GameControls) OnKeyDown(event d2input.KeyEvent) bool {
	if event.Key == d2input.KeyI {
		g.inventory.Toggle()
		g.updateLayout()
		return true
	}
	if event.Key == d2input.KeyC {
		g.heroStats.Toggle()
		g.updateLayout()
		return true
	}
	if event.Key == d2input.KeyR {
		g.hero.ToggleIsRunning()
		// TODO: change the running menu icon
		if g.hero.IsRunToggled() {
			g.hero.SetIsRunning(true)
		} else {
			g.hero.SetIsRunning(false)
		}
		return true
	}

	return false
}

var lastLeftBtnActionTime float64 = 0
var lastRightBtnActionTime float64 = 0
var mouseBtnActionsTreshhold = 0.25

func (g *GameControls) OnMouseButtonRepeat(event d2input.MouseEvent) bool {
	px, py := g.mapRenderer.ScreenToWorld(event.X, event.Y)
	px = float64(int(px*10)) / 10.0
	py = float64(int(py*10)) / 10.0

	now := d2common.Now()
	if event.Button == d2input.MouseButtonLeft && now-lastLeftBtnActionTime >= mouseBtnActionsTreshhold {
		lastLeftBtnActionTime = now
		g.inputListener.OnPlayerMove(px, py)
		return true
	}

	if event.Button == d2input.MouseButtonRight && now-lastRightBtnActionTime >= mouseBtnActionsTreshhold {
		lastRightBtnActionTime = now
		g.ShootMissile(px, py)
		return true
	}

	return false
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

func (g *GameControls) OnMouseButtonDown(event d2input.MouseEvent) bool {
	px, py := g.mapRenderer.ScreenToWorld(event.X, event.Y)
	px = float64(int(px*10)) / 10.0
	py = float64(int(py*10)) / 10.0

	if event.Button == d2input.MouseButtonLeft {
		lastLeftBtnActionTime = d2common.Now()
		g.inputListener.OnPlayerMove(px, py)
		return true
	}

	if event.Button == d2input.MouseButtonRight {
		return g.ShootMissile(px, py)
	}

	return false
}

func (g *GameControls) ShootMissile(px float64, py float64) bool {
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
}

// TODO: consider caching the panels to single image that is reused.
func (g *GameControls) Render(target d2render.Surface) {
	g.inventory.Render(target)
	g.heroStats.Render(target)

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

	if g.isZoneTextShown {
		g.zoneChangeText.SetPosition(width / 2, height/4)
		g.zoneChangeText.Render(target)
	}
}

func (g *GameControls) SetZoneChangeText(text string) {
	g.zoneChangeText.SetText(text)
}

func (g *GameControls) ShowZoneChangeText() {
	g.isZoneTextShown = true
}

func (g *GameControls) HideZoneChangeTextAfter(delay float64) {
	time.AfterFunc(time.Duration(delay)*time.Second, func() {
		g.isZoneTextShown = false
	})
}
