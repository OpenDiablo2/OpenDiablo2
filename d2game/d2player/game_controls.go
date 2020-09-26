package d2player

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"strings"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player/help"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

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
	heroState              *d2hero.HeroStateFactory
	mapEngine              *d2mapengine.MapEngine
	mapRenderer            *d2maprenderer.MapRenderer
	ui                     *d2ui.UIManager
	inventory              *Inventory
	heroStatsPanel         *HeroStatsPanel
	HelpOverlay            *help.Overlay
	miniPanel              *miniPanel
	lastMouseX             int
	lastMouseY             int
	missileID              int
	globeSprite            *d2ui.Sprite
	hpManaStatusSprite     *d2ui.Sprite
	mainPanel              *d2ui.Sprite
	menuButton             *d2ui.Sprite
	leftSkill              *SkillResource
	rightSkill             *SkillResource
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

type SkillResource struct {
	SkillResourcePath string
	IconNumber        int
	SkillIcon         *d2ui.Sprite
}

const (
	// Since they require special handling, not considering (1) globes, (2) content of the mini panel, (3) belt
	leftSkill ActionableType = iota
	newStats
	xp
	walkRun
	stamina
	miniPnl
	newSkills
	rightSkill
	hpGlobe
	manaGlobe
	miniPanelCharacter
	miniPanelInventory
	miniPanelSkillTree
	miniPanelAutomap
	miniPanelMessageLog
	miniPanelQuestLog
	miniPanelGameMenu
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

	zoneLabel := ui.NewLabel(d2resource.Font30, d2resource.PaletteUnits)
	zoneLabel.Alignment = d2gui.HorizontalAlignCenter

	nameLabel := ui.NewLabel(d2resource.Font16, d2resource.PaletteStatic)
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

	inventoryRecord := asset.Records.Layout.Inventory[inventoryRecordKey]

	hoverLabel := nameLabel
	hoverLabel.SetBackgroundColor(color.RGBA{0, 0, 0, uint8(128)})

	globeStatsLabel := hpManaStatsLabel

	heroState, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	gc := &GameControls{
		asset:            asset,
		ui:               ui,
		renderer:         renderer,
		hero:             hero,
		heroState:        heroState,
		mapEngine:        mapEngine,
		inputListener:    inputListener,
		mapRenderer:      mapRenderer,
		inventory:        NewInventory(asset, ui, inventoryRecord),
		heroStatsPanel:   NewHeroStatsPanel(asset, ui, hero.Name(), hero.Class, hero.Stats),
		HelpOverlay:      help.NewHelpOverlay(asset, renderer, ui, guiManager),
		miniPanel:        newMiniPanel(asset, ui, isSinglePlayer),
		nameLabel:        hoverLabel,
		zoneChangeText:   zoneLabel,
		hpManaStatsLabel: globeStatsLabel,
		actionableRegions: []ActionableRegion{
			{leftSkill, d2geom.Rectangle{Left: 115, Top: 550, Width: 50, Height: 50}},
			{newStats, d2geom.Rectangle{Left: 206, Top: 563, Width: 30, Height: 30}},
			{xp, d2geom.Rectangle{Left: 253, Top: 560, Width: 125, Height: 5}},
			{walkRun, d2geom.Rectangle{Left: 255, Top: 573, Width: 17, Height: 20}},
			{stamina, d2geom.Rectangle{Left: 273, Top: 573, Width: 105, Height: 20}},
			{miniPnl, d2geom.Rectangle{Left: 393, Top: 563, Width: 12, Height: 23}},
			{newSkills, d2geom.Rectangle{Left: 562, Top: 563, Width: 30, Height: 30}},
			{rightSkill, d2geom.Rectangle{Left: 634, Top: 550, Width: 50, Height: 50}},
			{hpGlobe, d2geom.Rectangle{Left: 30, Top: 525, Width: 80, Height: 60}},
			{manaGlobe, d2geom.Rectangle{Left: 695, Top: 525, Width: 80, Height: 60}},
			{miniPanelCharacter, d2geom.Rectangle{Left: 324, Top: 528, Width: 22, Height: 26}},
			{miniPanelInventory, d2geom.Rectangle{Left: 346, Top: 528, Width: 22, Height: 26}},
			{miniPanelSkillTree, d2geom.Rectangle{Left: 368, Top: 528, Width: 22, Height: 26}},
			{miniPanelAutomap, d2geom.Rectangle{Left: 390, Top: 528, Width: 22, Height: 26}},
			{miniPanelMessageLog, d2geom.Rectangle{Left: 412, Top: 528, Width: 22, Height: 26}},
			{miniPanelQuestLog, d2geom.Rectangle{Left: 434, Top: 528, Width: 22, Height: 26}},
			{miniPanelGameMenu, d2geom.Rectangle{Left: 456, Top: 528, Width: 22, Height: 26}},
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

	err = term.BindAction("setleftskill", "set skill to fire on left click", func(id int) {
		skillRecord := gc.asset.Records.Skill.Details[id]
		skill, err := heroState.CreateHeroSkill(0, skillRecord.Skill)
		if err != nil {
			term.OutputErrorf("cannot create skill with ID of %d", id)
		}

		gc.hero.LeftSkill = skill
	})

	err = term.BindAction("setrightskill", "set skill to fire on right click", func(id int) {
		skillRecord := gc.asset.Records.Skill.Details[id]
		skill, err := heroState.CreateHeroSkill(0, skillRecord.Skill)
		if err != nil {
			term.OutputErrorf("cannot create skill with ID of %d", id)
		}

		gc.hero.RightSkill = skill
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
		g.HelpOverlay.Toggle()
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

		if event.KeyMod() == d2enum.KeyModShift {
			g.inputListener.OnPlayerCast(g.hero.LeftSkill.ID, px, py)
		} else {
			g.inputListener.OnPlayerMove(px, py)
		}

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

		g.inputListener.OnPlayerCast(g.hero.RightSkill.ID, px, py)

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

		if event.KeyMod() == d2enum.KeyModShift {
			g.inputListener.OnPlayerCast(g.hero.LeftSkill.ID, px, py)
		} else {
			g.inputListener.OnPlayerMove(px, py)
		}

		return true
	}

	if event.Button() == d2enum.MouseButtonRight && !g.isInActiveMenusRect(mx, my) && !g.hero.IsCasting() {
		g.lastRightBtnActionTime = d2util.Now()

		g.inputListener.OnPlayerCast(g.hero.RightSkill.ID, px, py)

		return true
	}

	return false
}

// Load the resources required for the GameControls
func (g *GameControls) Load() {
	var err error
	g.globeSprite, err = g.ui.NewSprite(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	g.hpManaStatusSprite, err = g.ui.NewSprite(d2resource.HealthManaIndicator, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	g.mainPanel, err = g.ui.NewSprite(d2resource.GamePanels, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	g.menuButton, err = g.ui.NewSprite(d2resource.MenuButton, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}
	err = g.menuButton.SetCurrentFrame(2)
	if err != nil {
		log.Print(err)
	}

	// TODO: temporarily hardcoded to Attack, should come from saved state for hero
	genericSkillsSprite, err := g.ui.NewSprite(d2resource.GenericSkills, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	attackIconID := 2

	g.leftSkill = &SkillResource{SkillIcon: genericSkillsSprite, IconNumber: attackIconID, SkillResourcePath: d2resource.GenericSkills}
	g.rightSkill = &SkillResource{SkillIcon: genericSkillsSprite, IconNumber: attackIconID, SkillResourcePath: d2resource.GenericSkills}

	g.loadUIButtons()

	g.inventory.Load()
	g.heroStatsPanel.Load()
	g.HelpOverlay.Load()
}

func (g *GameControls) loadUIButtons() {
	// Run button
	g.runButton = g.ui.NewButton(d2ui.ButtonTypeRun, "")

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

	if g.HelpOverlay.IsOpen() && g.HelpOverlay.IsInRect(px, py) {
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
	skillResourcePath := g.getSkillResourceByClass(g.hero.LeftSkill.Charclass)
	if skillResourcePath != g.leftSkill.SkillResourcePath {
		g.leftSkill.SkillIcon, _ = g.ui.NewSprite(skillResourcePath, d2resource.PaletteSky)
	}

	if err := g.leftSkill.SkillIcon.SetCurrentFrame(g.hero.LeftSkill.IconCel); err != nil {
		return err
	}

	w, _ = g.leftSkill.SkillIcon.GetCurrentFrameSize()

	g.leftSkill.SkillIcon.SetPosition(offset, height)

	if err := g.leftSkill.SkillIcon.Render(target); err != nil {
		return err
	}

	offset += w

	// New Stats Selector
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

	// New Skills Selector
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
	skillResourcePath = g.getSkillResourceByClass(g.hero.RightSkill.Charclass)
	if skillResourcePath != g.rightSkill.SkillResourcePath {
		g.rightSkill.SkillIcon, _ = g.ui.NewSprite(skillResourcePath, d2resource.PaletteSky)
	}

	if err := g.rightSkill.SkillIcon.SetCurrentFrame(g.hero.RightSkill.IconCel); err != nil {
		return err
	}

	w, _ = g.rightSkill.SkillIcon.GetCurrentFrameSize()

	g.rightSkill.SkillIcon.SetPosition(offset, height)

	if err := g.rightSkill.SkillIcon.Render(target); err != nil {
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

	// Create and format Health string from string lookup table.
	fmtHealth := d2tbl.TranslateString("panelhealth")
	healthCurr, healthMax := int(g.hero.Stats.Health), int(g.hero.Stats.MaxHealth)
	strPanelHealth := fmt.Sprintf(fmtHealth, healthCurr, healthMax)

	// Display current hp and mana stats hpGlobe or manaGlobe region is clicked
	if g.actionableRegions[hpGlobe].Rect.IsInRect(mx, my) || g.hpStatsIsVisible {
		g.hpManaStatsLabel.SetText(strPanelHealth)
		g.hpManaStatsLabel.SetPosition(15, 487)
		g.hpManaStatsLabel.Render(target)
	}

	// Create and format Mana string from string lookup table.
	fmtMana := d2tbl.TranslateString("panelmana")
	manaCurr, manaMax := int(g.hero.Stats.Mana), int(g.hero.Stats.MaxMana)
	strPanelMana := fmt.Sprintf(fmtMana, manaCurr, manaMax)

	if g.actionableRegions[manaGlobe].Rect.IsInRect(mx, my) || g.manaStatsIsVisible {
		g.hpManaStatsLabel.SetText(strPanelMana)
		// In case if the mana value gets higher, we need to shift the label to the left a little, hence widthManaLabel.
		widthManaLabel, _ := g.hpManaStatsLabel.GetSize()
		xManaLabel := 785 - widthManaLabel
		g.hpManaStatsLabel.SetPosition(xManaLabel, 487)
		g.hpManaStatsLabel.Render(target)
	}

	if err := g.HelpOverlay.Render(target); err != nil {
		return err
	}

	// Minipanel is closed and minipanel button is hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[miniPnl].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(d2tbl.TranslateString("panelcmini")) //"Close Mini Panel"
		g.nameLabel.SetPosition(399, 544)
		g.nameLabel.Render(target)
	}

	// Minipanel is open and minipanel button is hovered.
	if !g.miniPanel.IsOpen() && g.actionableRegions[miniPnl].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(d2tbl.TranslateString("panelmini")) //"Open Mini Panel"
		g.nameLabel.SetPosition(399, 544)
		g.nameLabel.Render(target)
	}

	// Display character tooltip when hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[miniPanelCharacter].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(d2tbl.TranslateString("minipanelchar")) //"Character" no hotkey
		g.nameLabel.SetPosition(340, 510)
		g.nameLabel.Render(target)
	}

	// Display inventory tooltip when hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[miniPanelInventory].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(d2tbl.TranslateString("minipanelinv")) //"Inventory" no hotkey
		g.nameLabel.SetPosition(360, 510)
		g.nameLabel.Render(target)
	}

	// Display skill tree tooltip when hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[miniPanelSkillTree].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(d2tbl.TranslateString("minipaneltree")) //"Skill Treee" no hotkey
		g.nameLabel.SetPosition(380, 510)
		g.nameLabel.Render(target)
	}

	// Display automap tooltip when hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[miniPanelAutomap].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(d2tbl.TranslateString("minipanelautomap")) //"Automap" no hotkey
		g.nameLabel.SetPosition(400, 510)
		g.nameLabel.Render(target)
	}

	// Display message log tooltip when hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[miniPanelMessageLog].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(d2tbl.TranslateString("minipanelmessage")) //"Message Log" no hotkey
		g.nameLabel.SetPosition(420, 510)
		g.nameLabel.Render(target)
	}

	// Display quest log tooltip when hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[miniPanelQuestLog].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(d2tbl.TranslateString("minipanelquest")) //"Quest Log" no hotkey
		g.nameLabel.SetPosition(440, 510)
		g.nameLabel.Render(target)
	}

	// Display game menu tooltip when hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[miniPanelGameMenu].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(d2tbl.TranslateString("minipanelmenubtn")) //"Game Menu (Esc)" // the (Esc) is hardcoded in.
		g.nameLabel.SetPosition(460, 510)
		g.nameLabel.Render(target)
	}

	// Create and format Stamina string from string lookup table.
	fmtStamina := d2tbl.TranslateString("panelstamina")
	staminaCurr, staminaMax := int(g.hero.Stats.Stamina), int(g.hero.Stats.MaxStamina)
	strPanelStamina := fmt.Sprintf(fmtStamina, staminaCurr, staminaMax)

	// Display stamina tooltip when hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[stamina].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(strPanelStamina)
		g.nameLabel.SetPosition(320, 535)
		g.nameLabel.Render(target)
	}

	// Display run/walk tooltip when hovered.  Note that whether the player is walking or running, the tooltip is the same in Diablo 2.
	if g.actionableRegions[walkRun].Rect.IsInRect(mx, my) && !g.hero.IsRunToggled() {
		g.nameLabel.SetText(d2tbl.TranslateString("RunOn")) //"Run" no hotkeys
		g.nameLabel.SetPosition(263, 563)
		g.nameLabel.Render(target)
	}

	if g.actionableRegions[walkRun].Rect.IsInRect(mx, my) && g.hero.IsRunToggled() {
		g.nameLabel.SetText(d2tbl.TranslateString("RunOff")) //"Walk" no hotkeys
		g.nameLabel.SetPosition(263, 563)
		g.nameLabel.Render(target)
	}

	// Create and format Experience string from string lookup table.
	fmtExp := d2tbl.TranslateString("panelexp")
	// The English string for "panelexp" is "Experience: %u / %u", however %u doesn't translate well. So
	// we need to rewrite %u into a formatable Go verb. %d is used in other strings, so we go with that,
	// keeping in mind that %u likely referred to an unsigned integer.
	fmtExp = strings.ReplaceAll(fmtExp, "%u", "%d")
	expCurr, expMax := uint(g.hero.Stats.Experience), uint(g.hero.Stats.NextLevelExp)
	strPanelExp := fmt.Sprintf(fmtExp, expCurr, expMax)

	// Display experience tooltip when hovered.
	if g.miniPanel.IsOpen() && g.actionableRegions[xp].Rect.IsInRect(mx, my) {
		g.nameLabel.SetText(strPanelExp)
		g.nameLabel.SetPosition(255, 535)
		g.nameLabel.Render(target)
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

// ToggleHpStats toggles the visibility of the hp and mana stats placed above their respective globe and load only if they do not match
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
	case newStats:
		return
	case xp:
		return
	case walkRun:
		return
	case stamina:
		return
	case miniPnl:
		return
	case newSkills:
		return
	case rightSkill:
		return
	case hpGlobe:
		return
	case manaGlobe:
		return
	case miniPanelCharacter:
		return
	case miniPanelInventory:
		return
	case miniPanelSkillTree:
		return
	case miniPanelAutomap:
		return
	case miniPanelMessageLog:
		return
	case miniPanelQuestLog:
		return
	case miniPanelGameMenu:
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
	case newStats:
		log.Println("New Stats Selector Action Pressed")
	case xp:
		log.Println("XP Action Pressed")
	case walkRun:
		log.Println("Walk/Run Action Pressed")
	case stamina:
		log.Println("Stamina Action Pressed")
	case miniPnl:
		log.Println("Mini Panel Action Pressed")

		g.miniPanel.Toggle()
	case newSkills:
		log.Println("New Skills Selector Action Pressed")
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

func (g *GameControls) getSkillResourceByClass(class string) string {
	resource := ""

	switch class {
	case "":
		resource = d2resource.GenericSkills
	case "bar":
		resource = d2resource.BarbarianSkills
	case "nec":
		resource = d2resource.NecromancerSkills
	case "pal":
		resource = d2resource.PaladinSkills
	case "ass":
		resource = d2resource.AssassinSkills
	case "sor":
		resource = d2resource.SorcererSkills
	case "ama":
		resource = d2resource.AmazonSkills
	case "dru":
		resource = d2resource.DruidSkills
	default:
		log.Fatalf("Unknown class token: '%s'", class)
	}

	return resource
}
