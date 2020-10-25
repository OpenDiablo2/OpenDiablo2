package d2player

import (
	"fmt"
	"image"
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
	staminaBarHeight         = 19.0
	globeHeight              = 80
	globeWidth               = 80
	hoverLabelOuterPad       = 5
	mouseBtnActionsThreshold = 0.25
	percentStaminaBarLow     = 0.25
)

const (
	// Since they require special handling, not considering (1) globes, (2) content of the mini panel, (3) belt
	leftSkill actionableType = iota
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

const (
	leftSkillX,
	leftSkillY,
	leftSkillWidth,
	leftSkillHeight = 115, 550, 50, 50

	newStatsX,
	newStatsY,
	newStatsWidth,
	newStatsHeight = 206, 563, 30, 30

	xpX,
	xpY,
	xpWidth,
	xpHeight = 253, 560, 125, 5

	walkRunX,
	walkRunY,
	walkRunWidth,
	walkRunHeight = 255, 573, 17, 20

	staminaX,
	staminaY,
	staminaWidth,
	staminaHeight = 273, 573, 105, 20

	miniPnlX,
	miniPnlY,
	miniPnlWidth,
	miniPnlHeight = 393, 563, 12, 23

	newSkillsX,
	newSkillsY,
	newSkillsWidth,
	newSkillsHeight = 562, 563, 30, 30

	rightSkillX,
	rightSkillY,
	rightSkillWidth,
	rightSkillHeight = 634, 550, 50, 50

	hpGlobeX,
	hpGlobeY,
	hpGlobeWidth,
	hpGlobeHeight = 30, 525, 80, 60

	manaGlobeX,
	manaGlobeY,
	manaGlobeWidth,
	manaGlobeHeight = 695, 525, 80, 60

	miniPanelCharacterX,
	miniPanelCharacterY,
	miniPanelCharacterWidth,
	miniPanelCharacterHeight = 324, 528, 22, 26

	miniPanelInventoryX,
	miniPanelInventoryY,
	miniPanelInventoryWidth,
	miniPanelInventoryHeight = 346, 528, 22, 26

	miniPanelSkillTreeX,
	miniPanelSkillTreeY,
	miniPanelSkillTreeWidth,
	miniPanelSkillTreeHeight = 368, 528, 22, 26

	miniPanelAutomapX,
	miniPanelAutomapY,
	miniPanelAutomapWidth,
	miniPanelAutomapHeight = 390, 528, 22, 26

	miniPanelMessageLogX,
	miniPanelMessageLogY,
	miniPanelMessageLogWidth,
	miniPanelMessageLogHeight = 412, 528, 22, 26

	miniPanelQuestLogX,
	miniPanelQuestLogY,
	miniPanelQuestLogWidth,
	miniPanelQuestLogHeight = 434, 528, 22, 26

	miniPanelGameMenuX,
	miniPanelGameMenuY,
	miniPanelGameMenuWidth,
	miniPanelGameMenuHeight = 456, 528, 22, 26
)

const (
	hpLabelX = 15
	hpLabelY = 487

	manaLabelX = 785
	manaLabelY = 487
)

const (
	runButtonX = 255
	runButtonY = 570
)

const (
	frameMenuButton        = 2
	frameHealthStatus      = 0
	frameManaStatus        = 1
	frameNewStatsSelector  = 1
	frameStamina           = 2
	framePotions           = 3
	frameNewSkillsSelector = 4
	frameRightGlobeHolder  = 5
	frameRightGlobe        = 1
)

const (
	manaStatusOffsetX = 7
	manaStatusOffsetY = -12

	healthStatusOffsetX = 30
	healthStatusOffsetY = -13

	globeSpriteOffsetX = 28
	globeSpriteOffsetY = -5

	staminaBarOffsetX = 273
	staminaBarOffsetY = 572

	experienceBarOffsetX = 256
	experienceBarOffsetY = 561

	rightGlobeOffsetX = 8
	rightGlobeOffsetY = -8

	miniPanelButtonOffsetX = -8
	miniPanelButtonOffsetY = -16
)

const (
	lightBrownAlpha72 = 0xaf8848c8
	redAlpha72        = 0xff0000c8
	whiteAlpha100     = 0xffffffff
)

const (
	zoneChangeTextX = screenWidth / 2
	zoneChangeTextY = screenHeight / 4
)

const (
	menuBottomRectX,
	menuBottomRectY,
	menuBottomRectW,
	menuBottomRectH = 0, 550, 800, 50

	menuLeftRectX,
	menuLeftRectY,
	menuLeftRectW,
	menuLeftRectH = 0, 0, 400, 600

	menuRightRectX,
	menuRightRectY,
	menuRightRectW,
	menuRightRectH = 400, 0, 400, 600
)

// GameControls represents the game's controls on the screen
type GameControls struct {
	actionableRegions      []actionableRegion
	asset                  *d2asset.AssetManager
	renderer               d2interface.Renderer // https://github.com/OpenDiablo2/OpenDiablo2/issues/798
	inputListener          inputCallbackListener
	hero                   *d2mapentity.Player
	heroState              *d2hero.HeroStateFactory
	mapEngine              *d2mapengine.MapEngine
	mapRenderer            *d2maprenderer.MapRenderer
	escapeMenu             *EscapeMenu
	ui                     *d2ui.UIManager
	inventory              *Inventory
	skilltree              *skillTree
	heroStatsPanel         *HeroStatsPanel
	HelpOverlay            *help.Overlay
	miniPanel              *miniPanel
	bottomMenuRect         *d2geom.Rectangle
	leftMenuRect           *d2geom.Rectangle
	rightMenuRect          *d2geom.Rectangle
	lastMouseX             int
	lastMouseY             int
	globeSprite            *d2ui.Sprite
	hpManaStatusSprite     *d2ui.Sprite
	mainPanel              *d2ui.Sprite
	menuButton             *d2ui.Sprite
	skillSelectMenu        *SkillSelectMenu
	leftSkillResource      *SkillResource
	rightSkillResource     *SkillResource
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

type actionableType int

type actionableRegion struct {
	actionableTypeID actionableType
	rect             d2geom.Rectangle
}

// SkillResource represents a Skill with its corresponding icon sprite, path to DC6 file and icon number.
// SkillResourcePath points to a DC6 resource which contains the icons of multiple skills as frames.
// The IconNumber is the frame at which we can find our skill sprite in the DC6 file.
type SkillResource struct {
	SkillResourcePath string // path to a skills DC6 file(see getSkillResourceByClass)
	IconNumber        int    // the index of the frame in the DC6 file
	SkillIcon         *d2ui.Sprite
}

// NewGameControls creates a GameControls instance and returns a pointer to it
// nolint:funlen // doesn't make sense to split this up
func NewGameControls(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	hero *d2mapentity.Player,
	mapEngine *d2mapengine.MapEngine,
	escapeMenu *EscapeMenu,
	mapRenderer *d2maprenderer.MapRenderer,
	inputListener inputCallbackListener,
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

	var inventoryRecordKey string

	switch hero.Class {
	case d2enum.HeroAssassin:
		inventoryRecordKey = "Assassin2"
	case d2enum.HeroAmazon:
		inventoryRecordKey = "Amazon2"
	case d2enum.HeroBarbarian:
		inventoryRecordKey = "Barbarian2"
	case d2enum.HeroDruid:
		inventoryRecordKey = "Druid2"
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

	const blackAlpha50percent = 0x0000007f

	hoverLabel := nameLabel
	hoverLabel.SetBackgroundColor(d2util.Color(blackAlpha50percent))

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
		escapeMenu:       escapeMenu,
		inputListener:    inputListener,
		mapRenderer:      mapRenderer,
		inventory:        NewInventory(asset, ui, inventoryRecord),
		skillSelectMenu:  NewSkillSelectMenu(asset, ui, hero),
		skilltree:        newSkillTree(hero.Skills, hero.Class, asset, renderer, ui, guiManager),
		heroStatsPanel:   NewHeroStatsPanel(asset, ui, hero.Name(), hero.Class, hero.Stats),
		HelpOverlay:      help.NewHelpOverlay(asset, renderer, ui, guiManager),
		miniPanel:        newMiniPanel(asset, ui, isSinglePlayer),
		nameLabel:        hoverLabel,
		zoneChangeText:   zoneLabel,
		hpManaStatsLabel: globeStatsLabel,
		bottomMenuRect: &d2geom.Rectangle{
			Left:   menuBottomRectX,
			Top:    menuBottomRectY,
			Width:  menuBottomRectW,
			Height: menuBottomRectH,
		},
		leftMenuRect: &d2geom.Rectangle{
			Left:   menuLeftRectX,
			Top:    menuLeftRectY,
			Width:  menuLeftRectW,
			Height: menuLeftRectH,
		},
		rightMenuRect: &d2geom.Rectangle{
			Left:   menuRightRectX,
			Top:    menuRightRectY,
			Width:  menuRightRectW,
			Height: menuRightRectH,
		},
		actionableRegions: []actionableRegion{
			{leftSkill, d2geom.Rectangle{
				Left:   leftSkillX,
				Top:    leftSkillY,
				Width:  leftSkillWidth,
				Height: leftSkillHeight,
			}},
			{newStats, d2geom.Rectangle{
				Left:   newStatsX,
				Top:    newStatsY,
				Width:  newStatsWidth,
				Height: newStatsHeight,
			}},
			{xp, d2geom.Rectangle{
				Left:   xpX,
				Top:    xpY,
				Width:  xpWidth,
				Height: xpHeight,
			}},
			{walkRun, d2geom.Rectangle{
				Left:   walkRunX,
				Top:    walkRunY,
				Width:  walkRunWidth,
				Height: walkRunHeight,
			}},
			{stamina, d2geom.Rectangle{
				Left:   staminaX,
				Top:    staminaY,
				Width:  staminaWidth,
				Height: staminaHeight,
			}},
			{miniPnl, d2geom.Rectangle{
				Left:   miniPnlX,
				Top:    miniPnlY,
				Width:  miniPnlWidth,
				Height: miniPnlHeight,
			}},
			{newSkills, d2geom.Rectangle{
				Left:   newSkillsX,
				Top:    newSkillsY,
				Width:  newSkillsWidth,
				Height: newSkillsHeight,
			}},
			{rightSkill, d2geom.Rectangle{
				Left:   rightSkillX,
				Top:    rightSkillY,
				Width:  rightSkillWidth,
				Height: rightSkillHeight,
			}},
			{hpGlobe, d2geom.Rectangle{
				Left:   hpGlobeX,
				Top:    hpGlobeY,
				Width:  hpGlobeWidth,
				Height: hpGlobeHeight,
			}},
			{manaGlobe, d2geom.Rectangle{
				Left:   manaGlobeX,
				Top:    manaGlobeY,
				Width:  manaGlobeWidth,
				Height: manaGlobeHeight,
			}},
			{miniPanelCharacter, d2geom.Rectangle{
				Left:   miniPanelCharacterX,
				Top:    miniPanelCharacterY,
				Width:  miniPanelCharacterWidth,
				Height: miniPanelCharacterHeight,
			}},
			{miniPanelInventory, d2geom.Rectangle{
				Left:   miniPanelInventoryX,
				Top:    miniPanelInventoryY,
				Width:  miniPanelInventoryWidth,
				Height: miniPanelInventoryHeight,
			}},
			{miniPanelSkillTree, d2geom.Rectangle{
				Left:   miniPanelSkillTreeX,
				Top:    miniPanelSkillTreeY,
				Width:  miniPanelSkillTreeWidth,
				Height: miniPanelSkillTreeHeight,
			}},
			{miniPanelAutomap, d2geom.Rectangle{
				Left:   miniPanelAutomapX,
				Top:    miniPanelAutomapY,
				Width:  miniPanelAutomapWidth,
				Height: miniPanelAutomapHeight,
			}},
			{miniPanelMessageLog, d2geom.Rectangle{
				Left:   miniPanelMessageLogX,
				Top:    miniPanelMessageLogY,
				Width:  miniPanelMessageLogWidth,
				Height: miniPanelMessageLogHeight,
			}},
			{miniPanelQuestLog, d2geom.Rectangle{
				Left:   miniPanelQuestLogX,
				Top:    miniPanelQuestLogY,
				Width:  miniPanelQuestLogWidth,
				Height: miniPanelQuestLogHeight,
			}},
			{miniPanelGameMenu, d2geom.Rectangle{
				Left:   miniPanelGameMenuX,
				Top:    miniPanelGameMenuY,
				Width:  miniPanelGameMenuWidth,
				Height: miniPanelGameMenuHeight,
			}},
		},
		lastLeftBtnActionTime:  0,
		lastRightBtnActionTime: 0,
		isSinglePlayer:         isSinglePlayer,
	}

	closeCb := func() { gc.updateLayout() }
	gc.heroStatsPanel.SetOnCloseCb(closeCb)
	gc.inventory.SetOnCloseCb(closeCb)
	gc.skilltree.SetOnCloseCb(closeCb)

	err = gc.bindTerminalCommands(term)
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
		g.onEscKey()
	case d2enum.KeyI:
		g.inventory.Toggle()
		g.updateLayout()
	case d2enum.KeyT:
		g.skilltree.Toggle()
		g.updateLayout()
	case d2enum.KeyC:
		g.heroStatsPanel.Toggle()
		g.updateLayout()
	case d2enum.KeyR, d2enum.KeyControl:
		g.onToggleRunButton()
	case d2enum.KeyH:
		g.HelpOverlay.Toggle()
		g.updateLayout()
	default:
		return false
	}

	return false
}

// OnKeyUp handles key release
func (g *GameControls) OnKeyUp(event d2interface.KeyEvent) bool {
	switch event.Key() {
	case d2enum.KeyControl:
		g.onToggleRunButton()
	default:
		return false
	}

	return false
}

// When escape is pressed:
// 1. If there was some overlay or panel open, close it
// 2. Otherwise, if the Escape Menu was open, let the Escape Menu handle it
// 3. If nothing was open, open the Escape Menu
func (g *GameControls) onEscKey() {
	escHandled := false

	if g.skillSelectMenu.IsOpen() {
		g.skillSelectMenu.ClosePanels()

		escHandled = true
	}

	if g.inventory.IsOpen() {
		g.inventory.Close()

		escHandled = true
	}

	if g.skilltree.IsOpen() {
		g.skilltree.Close()

		escHandled = true
	}

	if g.heroStatsPanel.IsOpen() {
		g.heroStatsPanel.Close()

		escHandled = true
	}

	if g.HelpOverlay.IsOpen() {
		g.HelpOverlay.Toggle()

		escHandled = true
	}

	switch escHandled {
	case true:
		g.updateLayout()
	case false:
		if g.escapeMenu.IsOpen() {
			g.escapeMenu.OnEscKey()
		} else {
			g.escapeMenu.open()
		}
	}
}

func truncateFloat64(n float64) float64 {
	const ten = 10.0
	return float64(int(n*ten)) / ten
}

// OnMouseButtonRepeat handles repeated mouse clicks
func (g *GameControls) OnMouseButtonRepeat(event d2interface.MouseEvent) bool {
	const (
		screenWidth, screenHeight         = 800, 600
		halfScreenWidth, halfScreenHeight = screenWidth / 2, screenHeight / 2
		subtilesPerTile                   = 5
	)

	px, py := g.mapRenderer.ScreenToWorld(event.X(), event.Y())
	px = truncateFloat64(px)
	py = truncateFloat64(py)

	now := d2util.Now()
	button := event.Button()
	isLeft := button == d2enum.MouseButtonLeft
	isRight := button == d2enum.MouseButtonRight
	lastLeft := now - g.lastLeftBtnActionTime
	lastRight := now - g.lastRightBtnActionTime
	inRect := !g.isInActiveMenusRect(event.X(), event.Y())
	shouldDoLeft := lastLeft >= mouseBtnActionsThreshold
	shouldDoRight := lastRight >= mouseBtnActionsThreshold

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

				x := float64(halfScreenWidth) / subtilesPerTile
				y := float64(halfScreenHeight) / subtilesPerTile

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
		if g.actionableRegions[i].rect.IsInRect(mx, my) {
			g.onHoverActionable(g.actionableRegions[i].actionableTypeID)
		}
	}

	g.skillSelectMenu.LeftPanel.HandleMouseMove(mx, my)
	g.skillSelectMenu.RightPanel.HandleMouseMove(mx, my)

	return false
}

// OnMouseButtonDown handles mouse button presses
func (g *GameControls) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	mx, my := event.X(), event.Y()

	for i := range g.actionableRegions {
		// If click is on a game control element
		if g.actionableRegions[i].rect.IsInRect(mx, my) {
			g.onClickActionable(g.actionableRegions[i].actionableTypeID)
			return false
		}
	}

	if g.skillSelectMenu.IsOpen() && event.Button() == d2enum.MouseButtonLeft {
		g.lastLeftBtnActionTime = d2util.Now()
		g.skillSelectMenu.HandleClick(mx, my)
		g.skillSelectMenu.ClosePanels()

		return false
	}

	px, py := g.mapRenderer.ScreenToWorld(mx, my)
	px = truncateFloat64(px)
	py = truncateFloat64(py)

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

	err = g.menuButton.SetCurrentFrame(frameMenuButton)
	if err != nil {
		log.Print(err)
	}

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/799
	genericSkillsSprite, err := g.ui.NewSprite(d2resource.GenericSkills, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	attackIconID := 2

	g.leftSkillResource = &SkillResource{
		SkillIcon:         genericSkillsSprite,
		IconNumber:        attackIconID,
		SkillResourcePath: d2resource.GenericSkills,
	}

	g.rightSkillResource = &SkillResource{
		SkillIcon:         genericSkillsSprite,
		IconNumber:        attackIconID,
		SkillResourcePath: d2resource.GenericSkills,
	}

	g.loadUIButtons()

	g.inventory.Load()
	g.skilltree.load()
	g.heroStatsPanel.Load()
	g.HelpOverlay.Load()
}

func (g *GameControls) loadUIButtons() {
	// Run button
	g.runButton = g.ui.NewButton(d2ui.ButtonTypeRun, "")

	g.runButton.SetPosition(runButtonX, runButtonY)
	g.runButton.OnActivated(func() { g.onToggleRunButton() })

	if g.hero.IsRunToggled() {
		g.runButton.Toggle()
	}
}

func (g *GameControls) onToggleRunButton() {
	g.runButton.Toggle()
	g.hero.ToggleRunWalk()

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/800
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
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/801
	return g.heroStatsPanel.IsOpen()
}

func (g *GameControls) isRightPanelOpen() bool {
	return g.inventory.IsOpen() || g.skilltree.IsOpen()
}

func (g *GameControls) isInActiveMenusRect(px, py int) bool {
	if g.bottomMenuRect.IsInRect(px, py) {
		return true
	}

	if g.isLeftPanelOpen() && g.leftMenuRect.IsInRect(px, py) {
		return true
	}

	if g.isRightPanelOpen() && g.rightMenuRect.IsInRect(px, py) {
		return true
	}

	if g.miniPanel.IsOpen() && g.miniPanel.isInRect(px, py) {
		return true
	}

	if g.escapeMenu.IsOpen() {
		return true
	}

	if g.HelpOverlay.IsOpen() && g.HelpOverlay.IsInRect(px, py) {
		return true
	}

	if g.skillSelectMenu.IsOpen() {
		return true
	}

	return false
}

// Render draws the GameControls onto the target
func (g *GameControls) Render(target d2interface.Surface) error {
	g.renderForSelectableEntitiesHovered(target)

	if err := g.renderPanels(target); err != nil {
		return err
	}

	if err := g.renderHUD(target); err != nil {
		return err
	}

	return nil
}

func (g *GameControls) renderForSelectableEntitiesHovered(target d2interface.Surface) {
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
		halfWidth, halfHeight := entityWidth>>1, entityHeight>>1
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
}

func (g *GameControls) renderPanels(target d2interface.Surface) error {
	if err := g.heroStatsPanel.Render(target); err != nil {
		return err
	}

	if err := g.inventory.Render(target); err != nil {
		return err
	}

	if err := g.skilltree.Render(target); err != nil {
		return err
	}

	return nil
}

func (g *GameControls) renderHUD(target d2interface.Surface) error {
	mx, my := g.lastMouseX, g.lastMouseY
	width, height := target.GetSize()
	offsetX := 0

	// Left globe holder
	if err := g.mainPanel.SetCurrentFrame(0); err != nil {
		return err
	}

	w, _ := g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offsetX, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	// Health status bar
	healthPercent := float64(g.hero.Stats.Health) / float64(g.hero.Stats.MaxHealth)
	hpBarHeight := int(healthPercent * float64(globeHeight))

	if err := g.hpManaStatusSprite.SetCurrentFrame(0); err != nil {
		return err
	}

	g.hpManaStatusSprite.SetPosition(offsetX+healthStatusOffsetX, height+healthStatusOffsetY)

	healthMaskRect := image.Rect(0, globeHeight-hpBarHeight, globeWidth, globeHeight)
	if err := g.hpManaStatusSprite.RenderSection(target, healthMaskRect); err != nil {
		return err
	}

	// Left globe
	if err := g.globeSprite.SetCurrentFrame(frameHealthStatus); err != nil {
		return err
	}

	g.globeSprite.SetPosition(offsetX+globeSpriteOffsetX, height+globeSpriteOffsetY)

	if err := g.globeSprite.Render(target); err != nil {
		return err
	}

	offsetX += w

	// Left skill
	newSkillResourcePath := g.getSkillResourceByClass(g.hero.LeftSkill.Charclass)
	if newSkillResourcePath != g.leftSkillResource.SkillResourcePath {
		g.leftSkillResource.SkillResourcePath = newSkillResourcePath
		g.leftSkillResource.SkillIcon, _ = g.ui.NewSprite(newSkillResourcePath, d2resource.PaletteSky)
	}

	if err := g.leftSkillResource.SkillIcon.SetCurrentFrame(g.hero.LeftSkill.IconCel); err != nil {
		return err
	}

	w, _ = g.leftSkillResource.SkillIcon.GetCurrentFrameSize()

	g.leftSkillResource.SkillIcon.SetPosition(offsetX, height)

	if err := g.leftSkillResource.SkillIcon.Render(target); err != nil {
		return err
	}

	offsetX += w

	// New Stats Selector
	if err := g.mainPanel.SetCurrentFrame(frameNewStatsSelector); err != nil {
		return err
	}

	w, _ = g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offsetX, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	offsetX += w

	// Stamina
	if err := g.mainPanel.SetCurrentFrame(frameStamina); err != nil {
		return err
	}

	w, _ = g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offsetX, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	offsetX += w

	// Stamina status bar
	target.PushTranslation(staminaBarOffsetX, staminaBarOffsetY)
	target.PushEffect(d2enum.DrawEffectModulate)

	staminaPercent := g.hero.Stats.Stamina / float64(g.hero.Stats.MaxStamina)

	staminaBarColor := d2util.Color(lightBrownAlpha72)
	if staminaPercent < percentStaminaBarLow {
		staminaBarColor = d2util.Color(redAlpha72)
	}

	target.DrawRect(int(staminaPercent*staminaBarWidth), staminaBarHeight, staminaBarColor)
	target.Pop()
	target.Pop()

	// Experience status bar
	target.PushTranslation(experienceBarOffsetX, experienceBarOffsetY)

	expPercent := float64(g.hero.Stats.Experience) / float64(g.hero.Stats.NextLevelExp)

	target.DrawRect(int(expPercent*expBarWidth), 2, d2util.Color(whiteAlpha100))
	target.Pop()

	// Center menu button
	menuButtonFrameIndex := 0
	if g.miniPanel.isOpen {
		menuButtonFrameIndex = 2
	}

	if err := g.menuButton.SetCurrentFrame(menuButtonFrameIndex); err != nil {
		return err
	}

	buttonX, buttonY := (width>>1)+miniPanelButtonOffsetX, height+miniPanelButtonOffsetY

	g.menuButton.SetPosition(buttonX, buttonY)

	if err := g.menuButton.Render(target); err != nil {
		return err
	}

	if err := g.miniPanel.Render(target); err != nil {
		return err
	}

	// Potions
	if err := g.mainPanel.SetCurrentFrame(framePotions); err != nil {
		return err
	}

	w, _ = g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offsetX, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	offsetX += w

	// New Skills Selector
	if err := g.mainPanel.SetCurrentFrame(frameNewSkillsSelector); err != nil {
		return err
	}

	w, _ = g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offsetX, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	offsetX += w

	// Right skill
	newSkillResourcePath = g.getSkillResourceByClass(g.hero.RightSkill.Charclass)
	if newSkillResourcePath != g.rightSkillResource.SkillResourcePath {
		g.rightSkillResource.SkillIcon, _ = g.ui.NewSprite(newSkillResourcePath, d2resource.PaletteSky)
		g.rightSkillResource.SkillResourcePath = newSkillResourcePath
	}

	if err := g.rightSkillResource.SkillIcon.SetCurrentFrame(g.hero.RightSkill.IconCel); err != nil {
		return err
	}

	w, _ = g.rightSkillResource.SkillIcon.GetCurrentFrameSize()

	g.rightSkillResource.SkillIcon.SetPosition(offsetX, height)

	if err := g.rightSkillResource.SkillIcon.Render(target); err != nil {
		return err
	}

	offsetX += w

	// Right globe holder
	if err := g.mainPanel.SetCurrentFrame(frameRightGlobeHolder); err != nil {
		return err
	}

	g.mainPanel.GetCurrentFrameSize()

	g.mainPanel.SetPosition(offsetX, height)

	if err := g.mainPanel.Render(target); err != nil {
		return err
	}

	// Mana status bar
	manaPercent := float64(g.hero.Stats.Mana) / float64(g.hero.Stats.MaxMana)
	manaBarHeight := int(manaPercent * float64(globeHeight))

	if err := g.hpManaStatusSprite.SetCurrentFrame(frameManaStatus); err != nil {
		return err
	}

	g.hpManaStatusSprite.SetPosition(offsetX+manaStatusOffsetX, height+manaStatusOffsetY)

	manaMaskRect := image.Rect(0, globeHeight-manaBarHeight, globeWidth, globeHeight)
	if err := g.hpManaStatusSprite.RenderSection(target, manaMaskRect); err != nil {
		return err
	}

	// Right globe
	if err := g.globeSprite.SetCurrentFrame(frameRightGlobe); err != nil {
		return err
	}

	g.globeSprite.SetPosition(offsetX+rightGlobeOffsetX, height+rightGlobeOffsetY)

	if err := g.globeSprite.Render(target); err != nil {
		return err
	}

	if err := g.globeSprite.Render(target); err != nil {
		return err
	}

	if g.isZoneTextShown {
		g.zoneChangeText.SetPosition(zoneChangeTextX, zoneChangeTextY)
		g.zoneChangeText.Render(target)
	}

	// Create and format Health string from string lookup table.
	fmtHealth := d2tbl.TranslateString("panelhealth")
	healthCurr, healthMax := g.hero.Stats.Health, g.hero.Stats.MaxHealth
	strPanelHealth := fmt.Sprintf(fmtHealth, healthCurr, healthMax)

	// Display current hp and mana stats hpGlobe or manaGlobe region is clicked
	if g.actionableRegions[hpGlobe].rect.IsInRect(mx, my) || g.hpStatsIsVisible {
		g.hpManaStatsLabel.SetText(strPanelHealth)
		g.hpManaStatsLabel.SetPosition(hpLabelX, hpLabelY)
		g.hpManaStatsLabel.Render(target)
	}

	// Create and format Mana string from string lookup table.
	fmtMana := d2tbl.TranslateString("panelmana")
	manaCurr, manaMax := g.hero.Stats.Mana, g.hero.Stats.MaxMana
	strPanelMana := fmt.Sprintf(fmtMana, manaCurr, manaMax)

	if g.actionableRegions[manaGlobe].rect.IsInRect(mx, my) || g.manaStatsIsVisible {
		g.hpManaStatsLabel.SetText(strPanelMana)
		// In case if the mana value gets higher, we need to shift the label to the left a little, hence widthManaLabel.
		widthManaLabel, _ := g.hpManaStatsLabel.GetSize()
		xManaLabel := manaLabelX - widthManaLabel
		g.hpManaStatsLabel.SetPosition(xManaLabel, manaLabelY)
		g.hpManaStatsLabel.Render(target)
	}

	if err := g.HelpOverlay.Render(target); err != nil {
		return err
	}

	miniPanelButtons := map[actionableType]string{
		miniPanelCharacter:  "minipanelchar",
		miniPanelInventory:  "minipanelinv",
		miniPanelSkillTree:  "minipaneltree",
		miniPanelAutomap:    "minipanelautomap",
		miniPanelMessageLog: "minipanelmessage",
		miniPanelQuestLog:   "minipanelquest",
		miniPanelGameMenu:   "minipanelmenubtn",
	}

	for miniPanelButton, stringTableKey := range miniPanelButtons {
		if !g.miniPanel.IsOpen() {
			continue
		}

		if g.actionableRegions[miniPanelButton].rect.IsInRect(mx, my) {
			rect := &g.actionableRegions[miniPanelButton].rect
			g.nameLabel.SetText(d2tbl.TranslateString(stringTableKey))

			halfButtonWidth := rect.Width >> 1
			halfButtonHeight := rect.Height >> 1

			centerX := rect.Left + halfButtonWidth
			centerY := rect.Top + halfButtonHeight

			_, labelHeight := g.nameLabel.GetSize()

			labelX := centerX
			labelY := centerY - halfButtonHeight - labelHeight

			g.nameLabel.SetPosition(labelX, labelY)
			g.nameLabel.Render(target)
		}
	}

	// Display run/walk tooltip when hovered.
	// Note that whether the player is walking or running, the tooltip is the same in Diablo 2.
	if g.actionableRegions[walkRun].rect.IsInRect(mx, my) {
		var stringTableKey string

		if g.hero.IsRunToggled() {
			stringTableKey = "RunOff"
		} else {
			stringTableKey = "RunOn"
		}

		g.nameLabel.SetText(d2tbl.TranslateString(stringTableKey))

		rect := &g.actionableRegions[walkRun].rect

		halfButtonWidth := rect.Width >> 1
		halfButtonHeight := rect.Height >> 1

		centerX := rect.Left + halfButtonWidth
		centerY := rect.Top + halfButtonHeight

		_, labelHeight := g.nameLabel.GetSize()

		labelX := centerX
		labelY := centerY - halfButtonHeight - labelHeight

		g.nameLabel.SetPosition(labelX, labelY)
		g.nameLabel.Render(target)
	}

	const (
		staminaExperienceY = 535
	)

	// Display stamina tooltip when hovered.
	if g.actionableRegions[stamina].rect.IsInRect(mx, my) {
		// Create and format Stamina string from string lookup table.
		fmtStamina := d2tbl.TranslateString("panelstamina")
		staminaCurr, staminaMax := int(g.hero.Stats.Stamina), g.hero.Stats.MaxStamina
		strPanelStamina := fmt.Sprintf(fmtStamina, staminaCurr, staminaMax)

		g.nameLabel.SetText(strPanelStamina)

		rect := &g.actionableRegions[stamina].rect

		halfButtonWidth := rect.Width >> 1
		centerX := rect.Left + halfButtonWidth

		_, labelHeight := g.nameLabel.GetSize()
		halfLabelHeight := labelHeight >> 1

		labelX := centerX
		labelY := staminaExperienceY - halfLabelHeight

		g.nameLabel.SetPosition(labelX, labelY)
		g.nameLabel.Render(target)
	}

	// Display experience tooltip when hovered.
	if g.actionableRegions[xp].rect.IsInRect(mx, my) {
		// Create and format Experience string from string lookup table.
		fmtExp := d2tbl.TranslateString("panelexp")

		// The English string for "panelexp" is "Experience: %u / %u", however %u doesn't
		// translate well. So we need to rewrite %u into a formatable Go verb. %d is used in other
		// strings, so we go with that, keeping in mind that %u likely referred to
		// an unsigned integer.
		fmtExp = strings.ReplaceAll(fmtExp, "%u", "%d")

		expCurr, expMax := uint(g.hero.Stats.Experience), uint(g.hero.Stats.NextLevelExp)
		strPanelExp := fmt.Sprintf(fmtExp, expCurr, expMax)

		g.nameLabel.SetText(strPanelExp)
		rect := &g.actionableRegions[stamina].rect

		halfButtonWidth := rect.Width >> 1
		centerX := rect.Left + halfButtonWidth

		_, labelHeight := g.nameLabel.GetSize()
		halfLabelHeight := labelHeight >> 1

		labelX := centerX
		labelY := staminaExperienceY - halfLabelHeight

		g.nameLabel.SetPosition(labelX, labelY)
		g.nameLabel.Render(target)
	}

	if g.skillSelectMenu.IsOpen() {
		g.skillSelectMenu.Render(target)
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
func (g *GameControls) onHoverActionable(item actionableType) {
	hoverMap := map[actionableType]func(){
		leftSkill:           func() {},
		newStats:            func() {},
		xp:                  func() {},
		walkRun:             func() {},
		stamina:             func() {},
		miniPnl:             func() {},
		newSkills:           func() {},
		rightSkill:          func() {},
		hpGlobe:             func() {},
		manaGlobe:           func() {},
		miniPanelCharacter:  func() {},
		miniPanelInventory:  func() {},
		miniPanelSkillTree:  func() {},
		miniPanelAutomap:    func() {},
		miniPanelMessageLog: func() {},
		miniPanelQuestLog:   func() {},
		miniPanelGameMenu:   func() {},
	}

	onHover, found := hoverMap[item]
	if !found {
		log.Printf("Unrecognized actionableType(%d) being hovered\n", item)
		return
	}

	onHover()
}

// Handles what to do when an actionable is clicked
func (g *GameControls) onClickActionable(item actionableType) {
	actionMap := map[actionableType]func(){
		leftSkill: func() {
			g.skillSelectMenu.ToggleLeftPanel()
		},

		newStats: func() {
			log.Println("New Stats Selector Action Pressed")
		},

		xp: func() {
			log.Println("XP Action Pressed")
		},

		walkRun: func() {
			log.Println("Walk/Run Action Pressed")
		},

		stamina: func() {
			log.Println("Stamina Action Pressed")
		},

		miniPnl: func() {
			log.Println("Mini Panel Action Pressed")

			g.miniPanel.Toggle()
		},

		newSkills: func() {
			log.Println("New Skills Selector Action Pressed")
		},

		rightSkill: func() {
			g.skillSelectMenu.ToggleRightPanel()
		},

		hpGlobe: func() {
			g.ToggleHpStats()
			log.Println("HP Globe Pressed")
		},

		manaGlobe: func() {
			g.ToggleManaStats()
			log.Println("Mana Globe Pressed")
		},

		miniPanelCharacter: func() {
			log.Println("Character button on mini panel is pressed")

			g.heroStatsPanel.Toggle()
			g.updateLayout()
		},

		miniPanelInventory: func() {
			log.Println("Inventory button on mini panel is pressed")

			g.inventory.Toggle()
			g.updateLayout()
		},

		miniPanelSkillTree: func() {
			log.Println("Skilltree button on mini panel is pressed")

			g.skilltree.Toggle()
			g.updateLayout()
		},

		miniPanelGameMenu: func() {
			g.miniPanel.Close()
			g.escapeMenu.open()
		},
	}

	action, found := actionMap[item]
	if !found {
		log.Printf("Unrecognized actionableType(%d) being clicked\n", item)
		return
	}

	action()
}

func (g *GameControls) bindFreeCamCommand(term d2interface.Terminal) error {
	return term.BindAction("freecam", "toggle free camera movement", func() {
		g.FreeCam = !g.FreeCam
	})
}

func (g *GameControls) bindSetLeftSkillCommand(term d2interface.Terminal) error {
	setLeftSkill := func(id int) {
		skillRecord := g.asset.Records.Skill.Details[id]
		skill, err := g.heroState.CreateHeroSkill(1, skillRecord.Skill)

		if err != nil {
			term.OutputErrorf("cannot create skill with ID of %d, error: %s", id, err)
			return
		}

		g.hero.LeftSkill = skill
	}

	return term.BindAction(
		"setleftskill",
		"set skill to fire on left click",
		setLeftSkill,
	)
}

func (g *GameControls) bindSetRightSkillCommand(term d2interface.Terminal) error {
	setRightSkill := func(id int) {
		skillRecord := g.asset.Records.Skill.Details[id]
		skill, err := g.heroState.CreateHeroSkill(0, skillRecord.Skill)

		if err != nil {
			term.OutputErrorf("cannot create skill with ID of %d, error: %s", id, err)
			return
		}

		g.hero.RightSkill = skill
	}

	return term.BindAction(
		"setrightskill",
		"set skill to fire on right click",
		setRightSkill,
	)
}

const classTokenLength = 3

func (g *GameControls) bindLearnSkillsCommand(term d2interface.Terminal) error {
	learnSkills := func(token string) {
		if len(token) < classTokenLength {
			term.OutputErrorf("The given class token should be at least 3 characters")
			return
		}

		validPrefixes := []string{"ama", "ass", "nec", "bar", "sor", "dru", "pal"}
		classToken := strings.ToLower(token)
		tokenPrefix := classToken[0:3]
		isValidToken := false

		for idx := range validPrefixes {
			if strings.Compare(tokenPrefix, validPrefixes[idx]) == 0 {
				isValidToken = true
			}
		}

		if !isValidToken {
			fmtInvalid := "Invalid class, must be a value starting with(case insensitive): %s"
			term.OutputErrorf(fmtInvalid, strings.Join(validPrefixes, ", "))

			return
		}

		var err error

		learnedSkillsCount := 0

		for _, skillDetailRecord := range g.asset.Records.Skill.Details {
			if skillDetailRecord.Charclass != classToken && skillDetailRecord.Charclass != "" {
				continue
			}

			skill, skillErr := g.heroState.CreateHeroSkill(1, skillDetailRecord.Skill)
			if skill == nil {
				continue
			}

			learnedSkillsCount++

			g.hero.Skills[skill.ID] = skill

			if skillErr != nil {
				err = skillErr
				break
			}
		}

		g.skillSelectMenu.RegenerateImageCache()
		log.Printf("Learned %d skills", learnedSkillsCount)

		if err != nil {
			term.OutputErrorf("cannot learn skill for class, error: %s", err)
			return
		}
	}

	return term.BindAction(
		"learnskills",
		"learn all skills for the a given class",
		learnSkills,
	)
}

func (g *GameControls) bindLearnSkillByIDCommand(term d2interface.Terminal) error {
	learnByID := func(id int) {
		skillRecord := g.asset.Records.Skill.Details[id]
		if skillRecord == nil {
			term.OutputErrorf("cannot find a skill record for ID: %d", id)
			return
		}

		skill, err := g.heroState.CreateHeroSkill(1, skillRecord.Skill)
		if skill == nil {
			term.OutputErrorf("cannot create skill: %s", skillRecord.Skill)
			return
		}

		g.hero.Skills[skill.ID] = skill

		if err != nil {
			term.OutputErrorf("cannot learn skill for class, error: %s", err)
			return
		}

		g.skillSelectMenu.RegenerateImageCache()
		log.Println("Learned skill: ", skill.Skill)
	}

	return term.BindAction(
		"learnskillid",
		"learn a skill by a given ID",
		learnByID,
	)
}

func (g *GameControls) bindTerminalCommands(term d2interface.Terminal) error {
	if err := g.bindFreeCamCommand(term); err != nil {
		return err
	}

	if err := g.bindSetLeftSkillCommand(term); err != nil {
		return err
	}

	if err := g.bindSetRightSkillCommand(term); err != nil {
		return err
	}

	if err := g.bindLearnSkillsCommand(term); err != nil {
		return err
	}

	if err := g.bindLearnSkillByIDCommand(term); err != nil {
		return err
	}

	return nil
}

func (g *GameControls) getSkillResourceByClass(class string) string {
	resourceMap := map[string]string{
		"":    d2resource.GenericSkills,
		"bar": d2resource.BarbarianSkills,
		"nec": d2resource.NecromancerSkills,
		"pal": d2resource.PaladinSkills,
		"ass": d2resource.AssassinSkills,
		"sor": d2resource.SorcererSkills,
		"ama": d2resource.AmazonSkills,
		"dru": d2resource.DruidSkills,
	}

	entry, found := resourceMap[class]
	if !found {
		log.Fatalf("Unknown class token: '%s'", class)
	}

	return entry
}
