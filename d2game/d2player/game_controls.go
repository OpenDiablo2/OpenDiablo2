package d2player

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player/help"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// Panel represents the panel at the bottom of the game screen
type Panel interface {
	IsOpen() bool
	Toggle()
	Open()
	Close()
}

const (
	initialMissileID         = 59
	expBarWidth              = 120.0
	staminaBarWidth          = 102.0
	globeHeight              = 80
	globeWidth               = 80
	hoverLabelOuterPad       = 5
	mouseBtnActionsTreshhold = 0.25
)

// GameControls represents the game's controls on the screen
type GameControls struct {
	actionableRegions      []ActionableRegion
	asset                  *d2asset.AssetManager
	renderer               d2interface.Renderer // TODO: This shouldn't be a dependency
	inputListener          InputCallbackListener
	hero                   *d2mapentity.Player
	mapEngine              *d2mapengine.MapEngine
	mapRenderer            *d2maprenderer.MapRenderer
	uiManager              *d2ui.UIManager
	inventory              *Inventory
	heroStatsPanel         *HeroStatsPanel
	helpOverlay            *help.Overlay
	miniPanel              *miniPanel
	lastMouseX             int
	lastMouseY             int
	missileID              int
	globeSprite            *d2ui.Sprite
	hpManaStatusSprite     *d2ui.Sprite
	mainPanel              *d2ui.Sprite
	menuButton             *d2ui.Sprite
	skillIcon              *d2ui.Sprite
	zoneChangeText         *d2ui.Label
	nameLabel              *d2ui.Label
	hpManaStatsLabel       *d2ui.Label
	runButton              *d2ui.Button
	lastLeftBtnActionTime  float64
	lastRightBtnActionTime float64
	FreeCam                bool
	isZoneTextShown        bool
	hpStatsIsVisible       bool
	manaStatsIsVisible     bool
	isSinglePlayer         bool
}

type ActionableType int

type ActionableRegion struct {
	ActionableTypeID ActionableType
	Rect             d2geom.Rectangle
}

const (
	// Since they require special handling, not considering (1) globes, (2) content of the mini panel, (3) belt
	leftSkill ActionableType = iota
	leftSelect
	xp
	walkRun
	stamina
	miniPnl
	rightSelect
	rightSkill
	hpGlobe
	manaGlobe
	miniPanelCharacter
	miniPanelInventory
)

// NewGameControls creates a GameControls instance and returns a pointer to it
func NewGameControls(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	hero *d2mapentity.Player,
	mapEngine *d2mapengine.MapEngine,
	mapRenderer *d2maprenderer.MapRenderer,
	inputListener InputCallbackListener,
	term d2interface.Terminal,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
	isSinglePlayer bool,
) (*GameControls, error) {
	missileID := initialMissileID

	err := term.BindAction("setmissile", "set missile id to summon on right click", func(id int) {
		missileID = id
	})
	if err != nil {
		return nil, err
	}

	zoneLabel := ui.NewLabel(d2resource.Font30, d2resource.PaletteUnits)
	zoneLabel.Alignment = d2gui.HorizontalAlignCenter

	nameLabel := ui.NewLabel(d2resource.FontFormal11, d2resource.PaletteStatic)
	nameLabel.Alignment = d2gui.HorizontalAlignCenter
	nameLabel.SetText(d2ui.ColorTokenize("", d2ui.ColorTokenServer))

	hpManaStatsLabel := ui.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	hpManaStatsLabel.Alignment = d2gui.HorizontalAlignLeft

	// TODO make this depend on the hero type to respect inventory.txt
	var inventoryRecordKey string

	switch hero.Class {
	case d2enum.HeroAssassin:
		inventoryRecordKey = "Assassin"
	case d2enum.HeroAmazon:
		inventoryRecordKey = "Amazon2"
	case d2enum.HeroBarbarian:
		inventoryRecordKey = "Barbarian2"
	case d2enum.HeroDruid:
		inventoryRecordKey = "Druid"
	case d2enum.HeroNecromancer:
		inventoryRecordKey = "Necromancer2"
	case d2enum.HeroPaladin:
		inventoryRecordKey = "Paladin2"
	case d2enum.HeroSorceress:
		inventoryRecordKey = "Sorceress2"
	default:
		return nil, fmt.Errorf("unknown hero class: %d", hero.Class)
	}

	inventoryRecord := d2datadict.Inventory[inventoryRecordKey]

	hoverLabel := nameLabel
	hoverLabel.SetBackgroundColor(color.RGBA{0, 0, 0, uint8(128)})

	globeStatsLabel := hpManaStatsLabel

	gc := &GameControls{
		asset:            asset,
		uiManager:        ui,
		renderer:         renderer,
		hero:             hero,
		mapEngine:        mapEngine,
		inputListener:    inputListener,
		mapRenderer:      mapRenderer,
		inventory:        NewInventory(asset, ui, inventoryRecord),
		heroStatsPanel:   NewHeroStatsPanel(asset, ui, hero.Name(), hero.Class, hero.Stats),
		helpOverlay:      help.NewHelpOverlay(asset, renderer, ui, guiManager),
		miniPanel:        newMiniPanel(asset, ui, isSinglePlayer),
		missileID:        missileID,
		nameLabel:        hoverLabel,
		zoneChangeText:   zoneLabel,
		hpManaStatsLabel: globeStatsLabel,
		actionableRegions: []ActionableRegion{
			{leftSkill, d2geom.Rectangle{Left: 115, Top: 550, Width: 50, Height: 50}},
			{leftSelect, d2geom.Rectangle{Left: 206, Top: 563, Width: 30, Height: 30}},
			{xp, d2geom.Rectangle{Left: 253, Top: 560, Width: 125, Height: 5}},
			{walkRun, d2geom.Rectangle{Left: 255, Top: 573, Width: 17, Height: 20}},
			{stamina, d2geom.Rectangle{Left: 273, Top: 573, Width: 105, Height: 20}},
			{miniPnl, d2geom.Rectangle{Left: 393, Top: 563, Width: 12, Height: 23}},
			{rightSelect, d2geom.Rectangle{Left: 562, Top: 563, Width: 30, Height: 30}},
			{rightSkill, d2geom.Rectangle{Left: 634, Top: 550, Width: 50, Height: 50}},
			{hpGlobe, d2geom.Rectangle{Left: 30, Top: 525, Width: 65, Height: 50}},
			{manaGlobe, d2geom.Rectangle{Left: 700, Top: 525, Width: 65, Height: 50}},
			{miniPanelCharacter, d2geom.Rectangle{Left: 325, Top: 526, Width: 26, Height: 26}},
			{miniPanelInventory, d2geom.Rectangle{Left: 351, Top: 526, Width: 26, Height: 26}},
		},
		lastLeftBtnActionTime:  0,
		lastRightBtnActionTime: 0,
		isSinglePlayer:         isSinglePlayer,
	}

	err = term.BindAction("freecam", "toggle free camera movement", func() {
		gc.FreeCam = !gc.FreeCam
	})

	if err != nil {
		return nil, err
	}

	return gc, nil
}

// OnKeyRepeat is called to handle repeated key presses
func (g *GameControls) OnKeyRepeat(event d2interface.KeyEvent) bool {
	if g.FreeCam {
		var moveSpeed float64 = 8
		if event.KeyMod() == d2enum.KeyModShift {
			moveSpeed *= 2
		}

		if event.Key() == d2enum.KeyDown {
			v := d2vector.NewVector(0, moveSpeed)
			g.mapRenderer.MoveCameraTargetBy(v)

			return true
		}

		if event.Key() == d2enum.KeyUp {
			v := d2vector.NewVector(0, -moveSpeed)
			g.mapRenderer.MoveCameraTargetBy(v)

			return true
		}

		if event.Key() == d2enum.KeyRight {
			v := d2vector.NewVector(moveSpeed, 0)
			g.mapRenderer.MoveCameraTargetBy(v)

			return true
		}

		if event.Key() == d2enum.KeyLeft {
			v := d2vector.NewVector(-moveSpeed, 0)
			g.mapRenderer.MoveCameraTargetBy(v)

			return true
		}
	}

	return false
}

// OnKeyDown handles key presses
func (g *GameControls) OnKeyDown(event d2interface.KeyEvent) bool {
	switch event.Key() {
	case d2enum.KeyEscape:
		if g.inventory.IsOpen() || g.heroStatsPanel.IsOpen() {
			g.inventory.Close()
			g.heroStatsPanel.Close()
			g.updateLayout()

			break
		}
	case d2enum.KeyI:
		g.inventory.Toggle()
		g.updateLayout()
	case d2enum.KeyC:
		g.heroStatsPanel.Toggle()
		g.updateLayout()
	case d2enum.KeyR:
		g.onToggleRunButton()
	case d2enum.KeyH:
		g.helpOverlay.Toggle()
		g.updateLayout()
	default:
		return false
	}

	return false
}

// OnMouseButtonRepeat handles repeated mouse clicks
func (g *GameControls) OnMouseButtonRepeat(event d2interface.MouseEvent) bool {
	px, py := g.mapRenderer.ScreenToWorld(event.X(), event.Y())
	px = float64(int(px*10)) / 10.0
	py = float64(int(py*10)) / 10.0

	now := d2util.Now()
	button := event.Button()
	isLeft := button == d2enum.MouseButtonLeft
	isRight := button == d2enum.MouseButtonRight
	lastLeft := now - g.lastLeftBtnActionTime
	lastRight := now - g.lastRightBtnActionTime
	inRect := !g.isInActiveMenusRect(event.X(), event.Y())
	shouldDoLeft := lastLeft >= mouseBtnActionsTreshhold
	shouldDoRight := lastRight >= mouseBtnActionsTreshhold

	if isLeft && shouldDoLeft && inRect && !g.hero.IsCasting() {
		g.lastLeftBtnActionTime = now

		g.inputListener.OnPlayerMove(px, py)

		if g.FreeCam {
			if event.Button() == d2enum.MouseButtonLeft {
				camVect := g.mapRenderer.Camera.GetPosition().Vector

				x, y := float64(g.lastMouseX-400)/5, float64(g.lastMouseY-300)/5
				targetPosition := d2vector.NewPositionTile(x, y)
				targetPosition.Add(&camVect)

				g.mapRenderer.SetCameraTarget(&targetPosition)

				return true
			}
		}

		return true
	}

	if isRight && shouldDoRight && inRect && !g.hero.IsCasting() {
		g.lastRightBtnActionTime = now

		g.inputListener.OnPlayerCast(g.missileID, px, py)

		return true
	}

	return true
}

// OnMouseMove handles mouse movement events
func (g *GameControls) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	mx, my := event.X(), event.Y()
	g.lastMouseX = mx
	g.lastMouseY = my
	g.inventory.lastMouseX = mx
	g.inventory.lastMouseY = my

	for i := range g.actionableRegions {
		// Mouse over a game control element
		if g.actionableRegions[i].Rect.IsInRect(mx, my) {
			g.onHoverActionable(g.actionableRegions[i].ActionableTypeID)
		}
	}

	return false
}

// OnMouseButtonDown handles mouse button presses
func (g *GameControls) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	mx, my := event.X(), event.Y()

	for i := range g.actionableRegions {
		// If click is on a game control element
		if g.actionableRegions[i].Rect.IsInRect(mx, my) {
			g.onClickActionable(g.actionableRegions[i].ActionableTypeID)
			return false
		}
	}

	px, py := g.mapRenderer.ScreenToWorld(mx, my)
	px = float64(int(px*10)) / 10.0
	py = float64(int(py*10)) / 10.0

	if event.Button() == d2enum.MouseButtonLeft && !g.isInActiveMenusRect(mx, my) && !g.hero.IsCasting() {
		g.lastLeftBtnActionTime = d2util.Now()

		g.inputListener.OnPlayerMove(px, py)

		return true
	}

	if event.Button() == d2enum.MouseButtonRight && !g.isInActiveMenusRect(mx, my) && !g.hero.IsCasting() {
		g.lastRightBtnActionTime = d2util.Now()

		g.inputListener.OnPlayerCast(g.missileID, px, py)

		return true
	}

	return false
}

// Load the resources required for the GameControls
func (g *GameControls) Load() {
	var err error
	g.globeSprite, err = g.uiManager.NewSprite(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	g.hpManaStatusSprite, err = g.uiManager.NewSprite(d2resource.HealthManaIndicator, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	g.mainPanel, err = g.uiManager.NewSprite(d2resource.GamePanels, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	g.menuButton, err = g.uiManager.NewSprite(d2resource.MenuButton, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}
	err = g.menuButton.SetCurrentFrame(2)
	if err != nil {
		log.Print(err)
	}

	g.skillIcon, err = g.uiManager.NewSprite(d2resource.GenericSkills, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	g.loadUIButtons()

	g.inventory.Load()
	g.heroStatsPanel.Load()
	g.helpOverlay.Load()
}

func (g *GameControls) loadUIButtons() {
	// Run button
	g.runButton = g.uiManager.NewButton(d2ui.ButtonTypeRun, "")

	g.runButton.SetPosition(255, 570)
	g.runButton.OnActivated(func() { g.onToggleRunButton() })

	if g.hero.IsRunToggled() {
		g.runButton.Toggle()
	}
}

func (g *GameControls) onToggleRunButton() {
	g.runButton.Toggle()
	g.hero.ToggleRunWalk()
	// TODO: change the running menu icon
	g.hero.SetIsRunning(g.hero.IsRunToggled())
}

// Advance advances the state of the GameControls
func (g *GameControls) Advance(elapsed float64) error {
	g.mapRenderer.Advance(elapsed)
	return nil
}

func (g *GameControls) updateLayout() {
	isRightPanelOpen := g.isLeftPanelOpen()
	isLeftPanelOpen := g.isRightPanelOpen()

	switch {
	case isRightPanelOpen == isLeftPanelOpen:
		g.mapRenderer.ViewportDefault()
	case isRightPanelOpen:
		g.mapRenderer.ViewportToLeft()
	default:
		g.mapRenderer.ViewportToRight()
	}
}

func (g *GameControls) isLeftPanelOpen() bool {
	// TODO: add quest log panel
	return g.heroStatsPanel.IsOpen()
}

func (g *GameControls) isRightPanelOpen() bool {
	// TODO: add skills tree panel
	return g.inventory.IsOpen()
}

func (g *GameControls) isInActiveMenusRect(px, py int) bool {
	var bottomMenuRect = d2geom.Rectangle{Left: 0, Top: 550, Width: 800, Height: 50}

	var leftMenuRect = d2geom.Rectangle{Left: 0, Top: 0, Width: 400, Height: 600}

	var rightMenuRect = d2geom.Rectangle{Left: 400, Top: 0, Width: 400, Height: 600}

	if bottomMenuRect.IsInRect(px, py) {
		return true
	}

	if g.isLeftPanelOpen() && leftMenuRect.IsInRect(px, py) {
		return true
	}

	if g.isRightPanelOpen() && rightMenuRect.IsInRect(px, py) {
		return true
	}

	if g.miniPanel.IsOpen() && g.miniPanel.isInRect(px, py) {
		return true
	}

	if g.helpOverlay.IsOpen() && g.helpOverlay.IsInRect(px, py) {
		return true
	}

	return false
}

// TODO: consider caching the panels to single image that is reused.
// Render draws the GameControls onto the target
func (g *GameControls) Render(target d2interface.Surface) error {
	mx, my := g.lastMouseX, g.lastMouseY

	for entityIdx := range g.mapEngine.Entities() {
		entity := (g.mapEngine.Entities())[entityIdx]
		if !entity.Selectable() {
			continue
		}

		entPos := entity.GetPosition()
		entOffset := entPos.RenderOffset()
		entScreenXf, entScreenYf := g.mapRenderer.WorldToScreenF(entity.GetPositionF())
		entScreenX := int(math.Floor(entScreenXf))
		entScreenY := int(math.Floor(entScreenYf))
		entityWidth, entityHeight := entity.GetSize()
		halfWidth, halfHeight := entityWidth/2, entityHeight/2
		l, r := entScreenX-halfWidth-hoverLabelOuterPad, entScreenX+halfWidth+hoverLabelOuterPad
		t, b := entScreenY-halfHeight-hoverLabelOuterPad, entScreenY+halfHeight-hoverLabelOuterPad
		xWithin := (l <= mx) && (r >= mx)
		yWithin := (t <= my) && (b >= my)
		within := xWithin && yWithin

		if within {
			xOff, yOff := int(entOffset.X()), int(entOffset.Y())

			g.nameLabel.SetText(entity.Label())

			xLabel, yLabel := entScreenX-xOff, entScreenY-yOff-entityHeight-hoverLabelOuterPad
			g.nameLabel.SetPosition(xLabel, yLabel)

			g.nameLabel.Render(target)
			entity.Highlight()

			break
		}
	}

	if err := g.heroStatsPanel.Render(target); err != nil {
		return err
	}

	if err := g.inventory.Render(target); err != nil {
		return err
	}

	width, height := target.GetSize()
	offset := 0

	// Left globe holder
	if err := g.mainPanel.SetCurrentFrame(0); err != nil {
		return err
	}

	w, _ := g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offset, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	// Health status bar
	healthPercent := float64(g.hero.Stats.Health) / float64(g.hero.Stats.MaxHealth)
	hpBarHeight := int(healthPercent * float64(globeHeight))

	if err := g.hpManaStatusSprite.SetCurrentFrame(0); err != nil {
		return err
	}

	g.hpManaStatusSprite.SetPosition(offset+30, height-13)

	if err := g.hpManaStatusSprite.RenderSection(target, image.Rect(0, globeHeight-hpBarHeight, globeWidth, globeHeight)); err != nil {
		return err
	}

	// Left globe
	if err := g.globeSprite.SetCurrentFrame(0); err != nil {
		return err
	}

	g.globeSprite.SetPosition(offset+28, height-5)

	if err := g.globeSprite.Render(target); err != nil {
		return err
	}

	offset += w

	// Left skill
	if err := g.skillIcon.SetCurrentFrame(2); err != nil {
		return err
	}

	w, _ = g.skillIcon.GetCurrentFrameSize()

	g.skillIcon.SetPosition(offset, height)

	if err := g.skillIcon.Render(target); err != nil {
		return err
	}

	offset += w

	// Left skill selector
	if err := g.mainPanel.SetCurrentFrame(1); err != nil {
		return err
	}

	w, _ = g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offset, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	offset += w

	// Stamina
	if err := g.mainPanel.SetCurrentFrame(2); err != nil {
		return err
	}

	w, _ = g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offset, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	offset += w

	// Stamina status bar
	target.PushTranslation(273, 572)
	target.PushEffect(d2enum.DrawEffectModulate)

	staminaPercent := float64(g.hero.Stats.Stamina) / float64(g.hero.Stats.MaxStamina)

	target.DrawRect(int(staminaPercent*staminaBarWidth), 19, color.RGBA{R: 175, G: 136, B: 72, A: 200})
	target.PopN(2)

	// Experience status bar
	target.PushTranslation(256, 561)

	expPercent := float64(g.hero.Stats.Experience) / float64(g.hero.Stats.NextLevelExp)

	target.DrawRect(int(expPercent*expBarWidth), 2, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	target.Pop()

	// Center menu button
	menuButtonFrameIndex := 0
	if g.miniPanel.isOpen {
		menuButtonFrameIndex = 2
	}

	if err := g.menuButton.SetCurrentFrame(menuButtonFrameIndex); err != nil {
		return err
	}

	g.mainPanel.GetCurrentFrameSize()

	g.menuButton.SetPosition((width/2)-8, height-16)

	if err := g.menuButton.Render(target); err != nil {
		return err
	}

	if err := g.miniPanel.Render(target); err != nil {
		return err
	}

	// Potions
	if err := g.mainPanel.SetCurrentFrame(3); err != nil {
		return err
	}

	w, _ = g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offset, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	offset += w

	// Right skill selector
	if err := g.mainPanel.SetCurrentFrame(4); err != nil {
		return err
	}

	w, _ = g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offset, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	offset += w

	// Right skill
	if err := g.skillIcon.SetCurrentFrame(2); err != nil {
		return err
	}

	w, _ = g.skillIcon.GetCurrentFrameSize()

	g.skillIcon.SetPosition(offset, height)

	if err := g.skillIcon.Render(target); err != nil {
		return err
	}

	offset += w

	// Right globe holder
	if err := g.mainPanel.SetCurrentFrame(5); err != nil {
		return err
	}

	g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offset, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	// Mana status bar
	manaPercent := float64(g.hero.Stats.Mana) / float64(g.hero.Stats.MaxMana)
	manaBarHeight := int(manaPercent * float64(globeHeight))

	if err := g.hpManaStatusSprite.SetCurrentFrame(1); err != nil {
		return err
	}

	g.hpManaStatusSprite.SetPosition(offset+7, height-12)

	if err := g.hpManaStatusSprite.RenderSection(target, image.Rect(0, globeHeight-manaBarHeight, globeWidth, globeHeight)); err != nil {
		return err
	}

	// Right globe
	if err := g.globeSprite.SetCurrentFrame(1); err != nil {
		return err
	}

	g.globeSprite.SetPosition(offset+8, height-8)

	if err := g.globeSprite.Render(target); err != nil {
		return err
	}

	if err := g.globeSprite.Render(target); err != nil {
		return err
	}

	if g.isZoneTextShown {
		g.zoneChangeText.SetPosition(width/2, height/4)
		g.zoneChangeText.Render(target)
	}

	hpWithin := (mx <= 95) && (mx >= 30) && (my <= 575) && (my >= 525)
	manaWithin := (mx <= 765) && (mx >= 700) && (my <= 575) && (my >= 525)

	// Display current hp and mana stats hpGlobe or manaGlobe region is clicked
	if hpWithin || g.hpStatsIsVisible {
		g.hpManaStatsLabel.SetText(d2ui.ColorTokenize(
			fmt.Sprintf("LIFE: %v / %v", float64(g.hero.Stats.Health), float64(g.hero.Stats.MaxHealth)),
			d2ui.ColorTokenWhite),
		)
		g.hpManaStatsLabel.SetPosition(15, 485)
		g.hpManaStatsLabel.Render(target)
	}

	if manaWithin || g.manaStatsIsVisible {
		g.hpManaStatsLabel.SetText(fmt.Sprintf("MANA: %v / %v", float64(g.hero.Stats.Mana), float64(g.hero.Stats.MaxMana)))
		widthManaLabel, _ := g.hpManaStatsLabel.GetSize()
		xManaLabel := 785 - widthManaLabel
		g.hpManaStatsLabel.SetPosition(xManaLabel, 485)
		g.hpManaStatsLabel.Render(target)
	}

	if err := g.helpOverlay.Render(target); err != nil {
		return err
	}

	return nil
}

// SetZoneChangeText sets the zoneChangeText
func (g *GameControls) SetZoneChangeText(text string) {
	g.zoneChangeText.SetText(text)
}

// ShowZoneChangeText shows the zoneChangeText
func (g *GameControls) ShowZoneChangeText() {
	g.isZoneTextShown = true
}

// HideZoneChangeTextAfter hides the zoneChangeText after the given amount of seconds
func (g *GameControls) HideZoneChangeTextAfter(delay float64) {
	time.AfterFunc(time.Duration(delay)*time.Second, func() {
		g.isZoneTextShown = false
	})
}

// HpStatsIsVisible returns true if the hp and mana stats are visible to the player
func (g *GameControls) HpStatsIsVisible() bool {
	return g.hpStatsIsVisible
}

// ManaStatsIsVisible returns true if the hp and mana stats are visible to the player
func (g *GameControls) ManaStatsIsVisible() bool {
	return g.manaStatsIsVisible
}

// ToggleHpStats toggles the visibility of the hp and mana stats placed above their respective globe
func (g *GameControls) ToggleHpStats() {
	g.hpStatsIsVisible = !g.hpStatsIsVisible
}

// ToggleManaStats toggles the visibility of the hp and mana stats placed above their respective globe
func (g *GameControls) ToggleManaStats() {
	g.manaStatsIsVisible = !g.manaStatsIsVisible
}

// Handles what to do when an actionable is hovered
func (g *GameControls) onHoverActionable(item ActionableType) {
	switch item {
	case leftSkill:
		return
	case leftSelect:
		return
	case xp:
		return
	case walkRun:
		return
	case stamina:
		return
	case miniPnl:
		return
	case rightSelect:
		return
	case rightSkill:
		return
	case hpGlobe:
		return
	case manaGlobe:
		return
	default:
		log.Printf("Unrecognized ActionableType(%d) being hovered\n", item)
	}
}

// Handles what to do when an actionable is clicked
func (g *GameControls) onClickActionable(item ActionableType) {
	switch item {
	case leftSkill:
		log.Println("Left Skill Action Pressed")
	case leftSelect:
		log.Println("Left Skill Selector Action Pressed")
	case xp:
		log.Println("XP Action Pressed")
	case walkRun:
		log.Println("Walk/Run Action Pressed")
	case stamina:
		log.Println("Stamina Action Pressed")
	case miniPnl:
		log.Println("Mini Panel Action Pressed")

		g.miniPanel.Toggle()
	case rightSelect:
		log.Println("Right Skill Selector Action Pressed")
	case rightSkill:
		log.Println("Right Skill Action Pressed")
	case hpGlobe:
		g.ToggleHpStats()
		log.Println("HP Globe Pressed")
	case manaGlobe:
		g.ToggleManaStats()
		log.Println("Mana Globe Pressed")
	case miniPanelCharacter:
		log.Println("Character button on mini panel is pressed")

		g.heroStatsPanel.Toggle()
		g.updateLayout()
	case miniPanelInventory:
		log.Println("Inventory button on mini panel is pressed")

		g.inventory.Toggle()
		g.updateLayout()
	default:
		log.Printf("Unrecognized ActionableType(%d) being clicked\n", item)
	}
}
