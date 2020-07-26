package d2player

import (
	"image"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
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
	initialMissileID = 59
	expBarWidth      = 120.0
	staminaBarWidth  = 102.0
	globeHeight      = 80
	globeWidth       = 80
)

// GameControls represents the game's controls on the screen
type GameControls struct {
	renderer       d2interface.Renderer // TODO: This shouldn't be a dependency
	hero           *d2mapentity.Player
	mapEngine      *d2mapengine.MapEngine
	mapRenderer    *d2maprenderer.MapRenderer
	inventory      *Inventory
	heroStatsPanel *HeroStatsPanel
	inputListener  InputCallbackListener
	FreeCam        bool
	lastMouseX     int
	lastMouseY     int
	missileID      int // ID of missile to create when user right clicks.

	// UI
	globeSprite        *d2ui.Sprite
	hpManaStatusSprite *d2ui.Sprite
	mainPanel          *d2ui.Sprite
	menuButton         *d2ui.Sprite
	skillIcon          *d2ui.Sprite
	zoneChangeText     *d2ui.Label
	nameLabel          *d2ui.Label
	runButton          d2ui.Button
	isZoneTextShown    bool
	actionableRegions  []ActionableRegion
}

type ActionableType int

type ActionableRegion struct {
	ActionableTypeId ActionableType
	Rect             d2common.Rectangle
}

const (
	// Since they require special handling, not considering (1) globes, (2) content of the mini panel, (3) belt
	leftSkill  = ActionableType(iota)
	leftSelec  = ActionableType(iota)
	xp         = ActionableType(iota)
	walkRun    = ActionableType(iota)
	stamina    = ActionableType(iota)
	miniPanel  = ActionableType(iota)
	rightSelec = ActionableType(iota)
	rightSkill = ActionableType(iota)
)

func NewGameControls(renderer d2interface.Renderer, hero *d2mapentity.Player, mapEngine *d2mapengine.MapEngine,
	mapRenderer *d2maprenderer.MapRenderer, inputListener InputCallbackListener, term d2interface.Terminal) (*GameControls, error) {
	missileID := initialMissileID
	term.BindAction("setmissile", "set missile id to summon on right click", func(id int) {
		missileID = id
	})

	zoneLabel := d2ui.CreateLabel(d2resource.Font30, d2resource.PaletteUnits)
	zoneLabel.Color = color.RGBA{R: 255, G: 88, B: 82, A: 255}
	zoneLabel.Alignment = d2gui.HorizontalAlignCenter

	nameLabel := d2ui.CreateLabel(d2resource.FontFormal11, d2resource.PaletteStatic)
	nameLabel.Alignment = d2gui.HorizontalAlignCenter
	nameLabel.SetText("")
	nameLabel.Color = color.White

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
		inventoryRecordKey = "Amazon2"
	}

	inventoryRecord := d2datadict.Inventory[inventoryRecordKey]

	gc := &GameControls{
		renderer:       renderer,
		hero:           hero,
		mapEngine:      mapEngine,
		inputListener:  inputListener,
		mapRenderer:    mapRenderer,
		inventory:      NewInventory(inventoryRecord),
		heroStatsPanel: NewHeroStatsPanel(renderer, hero.Name(), hero.Class, hero.Stats),
		missileID:      missileID,
		nameLabel:      &nameLabel,
		zoneChangeText: &zoneLabel,
		actionableRegions: []ActionableRegion{
			{leftSkill, d2common.Rectangle{Left: 115, Top: 550, Width: 50, Height: 50}},
			{leftSelec, d2common.Rectangle{Left: 206, Top: 563, Width: 30, Height: 30}},
			{xp, d2common.Rectangle{Left: 253, Top: 560, Width: 125, Height: 5}},
			{walkRun, d2common.Rectangle{Left: 255, Top: 573, Width: 17, Height: 20}},
			{stamina, d2common.Rectangle{Left: 273, Top: 573, Width: 105, Height: 20}},
			{miniPanel, d2common.Rectangle{Left: 393, Top: 563, Width: 12, Height: 23}},
			{rightSelec, d2common.Rectangle{Left: 562, Top: 563, Width: 30, Height: 30}},
			{rightSkill, d2common.Rectangle{Left: 634, Top: 550, Width: 50, Height: 50}},
		},
	}

	err := term.BindAction("freecam", "toggle free camera movement", func() {
		gc.FreeCam = !gc.FreeCam
	})

	if err != nil {
		return nil, err
	}

	return gc, nil
}

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
	default:
		return false
	}
	return false
}

var lastLeftBtnActionTime float64 = 0
var lastRightBtnActionTime float64 = 0
var mouseBtnActionsTreshhold = 0.25

func (g *GameControls) OnMouseButtonRepeat(event d2interface.MouseEvent) bool {
	px, py := g.mapRenderer.ScreenToWorld(event.X(), event.Y())
	px = float64(int(px*10)) / 10.0
	py = float64(int(py*10)) / 10.0

	now := d2common.Now()
	button := event.Button()
	isLeft := button == d2enum.MouseButtonLeft
	isRight := button == d2enum.MouseButtonRight
	lastLeft := now - lastLeftBtnActionTime
	lastRight := now - lastRightBtnActionTime
	inRect := !g.isInActiveMenusRect(event.X(), event.Y())
	shouldDoLeft := lastLeft >= mouseBtnActionsTreshhold
	shouldDoRight := lastRight >= mouseBtnActionsTreshhold

	if isLeft && shouldDoLeft && inRect {
		lastLeftBtnActionTime = now

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

	if isRight && shouldDoRight && inRect {
		lastRightBtnActionTime = now

		g.inputListener.OnPlayerCast(g.missileID, px, py)

		return true
	}

	return true
}

func (g *GameControls) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	mx, my := event.X(), event.Y()
	g.lastMouseX = mx
	g.lastMouseY = my

	for i := range g.actionableRegions {
		// Mouse over a game control element
		if g.actionableRegions[i].Rect.IsInRect(mx, my) {
			g.onHoverActionable(g.actionableRegions[i].ActionableTypeId)
		}
	}

	return false
}

func (g *GameControls) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	mx, my := event.X(), event.Y()

	for i := range g.actionableRegions {
		// If click is on a game control element
		if g.actionableRegions[i].Rect.IsInRect(mx, my) {
			g.onClickActionable(g.actionableRegions[i].ActionableTypeId)
			return false
		}
	}

	px, py := g.mapRenderer.ScreenToWorld(mx, my)
	px = float64(int(px*10)) / 10.0
	py = float64(int(py*10)) / 10.0

	if event.Button() == d2enum.MouseButtonLeft && !g.isInActiveMenusRect(mx, my) {
		lastLeftBtnActionTime = d2common.Now()

		g.inputListener.OnPlayerMove(px, py)

		return true
	}

	if event.Button() == d2enum.MouseButtonRight && !g.isInActiveMenusRect(mx, my) {
		lastRightBtnActionTime = d2common.Now()

		g.inputListener.OnPlayerCast(g.missileID, px, py)

		return true
	}

	return false
}

func (g *GameControls) Load() {
	animation, _ := d2asset.LoadAnimation(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	g.globeSprite, _ = d2ui.LoadSprite(animation)

	animation, _ = d2asset.LoadAnimation(d2resource.HealthManaIndicator, d2resource.PaletteSky)
	g.hpManaStatusSprite, _ = d2ui.LoadSprite(animation)

	animation, _ = d2asset.LoadAnimation(d2resource.GamePanels, d2resource.PaletteSky)
	g.mainPanel, _ = d2ui.LoadSprite(animation)

	animation, _ = d2asset.LoadAnimation(d2resource.MenuButton, d2resource.PaletteSky)
	g.menuButton, _ = d2ui.LoadSprite(animation)

	animation, _ = d2asset.LoadAnimation(d2resource.GenericSkills, d2resource.PaletteSky)
	g.skillIcon, _ = d2ui.LoadSprite(animation)

	g.loadUIButtons()

	g.inventory.Load()
	g.heroStatsPanel.Load()
}

func (g *GameControls) loadUIButtons() {
	// Run button
	g.runButton = d2ui.CreateButton(g.renderer, d2ui.ButtonTypeRun, "")

	g.runButton.SetPosition(255, 570)
	g.runButton.OnActivated(func() { g.onToggleRunButton() })

	if g.hero.IsRunToggled() {
		g.runButton.Toggle()
	}

	d2ui.AddWidget(&g.runButton)
}

func (g *GameControls) onToggleRunButton() {
	g.runButton.Toggle()
	g.hero.ToggleRunWalk()
	// TODO: change the running menu icon
	g.hero.SetIsRunning(g.hero.IsRunToggled())
}

func (g *GameControls) Advance(elapsed float64) error {
	g.mapRenderer.Advance(elapsed)
	return nil
}

func (g *GameControls) updateLayout() {
	isRightPanelOpen := g.isLeftPanelOpen()
	isLeftPanelOpen := g.isRightPanelOpen()

	if isRightPanelOpen == isLeftPanelOpen {
		g.mapRenderer.ViewportDefault()
	} else if isRightPanelOpen {
		g.mapRenderer.ViewportToLeft()
	} else {
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

func (g *GameControls) isInActiveMenusRect(px int, py int) bool {
	var bottomMenuRect = d2common.Rectangle{Left: 0, Top: 550, Width: 800, Height: 50}

	var leftMenuRect = d2common.Rectangle{Left: 0, Top: 0, Width: 400, Height: 600}

	var rightMenuRect = d2common.Rectangle{Left: 400, Top: 0, Width: 400, Height: 600}

	if bottomMenuRect.IsInRect(px, py) {
		return true
	}

	if g.isLeftPanelOpen() && leftMenuRect.IsInRect(px, py) {
		return true
	}

	if g.isRightPanelOpen() && rightMenuRect.IsInRect(px, py) {
		return true
	}

	return false
}

// TODO: consider caching the panels to single image that is reused.
func (g *GameControls) Render(target d2interface.Surface) error {
	for entityIdx := range *g.mapEngine.Entities() {
		entity := (*g.mapEngine.Entities())[entityIdx]
		if !entity.Selectable() {
			continue
		}

		entScreenXf, entScreenYf := g.mapRenderer.WorldToScreenF(entity.GetPositionF())
		entScreenX := int(math.Floor(entScreenXf))
		entScreenY := int(math.Floor(entScreenYf))

		if ((entScreenX - 20) <= g.lastMouseX) && ((entScreenX + 20) >= g.lastMouseX) &&
			((entScreenY - 80) <= g.lastMouseY) && (entScreenY >= g.lastMouseY) {
			g.nameLabel.SetText(entity.Name())
			g.nameLabel.SetPosition(entScreenX, entScreenY-100)
			g.nameLabel.Render(target)
			entity.Highlight()

			break
		}
	}

	g.inventory.Render(target)
	g.heroStatsPanel.Render(target)

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
	if err := g.menuButton.SetCurrentFrame(0); err != nil {
		return err
	}

	w, _ = g.mainPanel.GetCurrentFrameSize()

	g.menuButton.SetPosition((width/2)-8, height-16)

	if err := g.menuButton.Render(target); err != nil {
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

	w, _ = g.mainPanel.GetCurrentFrameSize()

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

	return nil
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

// Handles what to do when an actionable is hovered
func (g *GameControls) onHoverActionable(item ActionableType) {
	switch item {
	case leftSkill:
		return
	case leftSelec:
		return
	case xp:
		return
	case walkRun:
		return
	case stamina:
		return
	case miniPanel:
		return
	case rightSelec:
		return
	case rightSkill:
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
	case leftSelec:
		log.Println("Left Skill Selector Action Pressed")
	case xp:
		log.Println("XP Action Pressed")
	case walkRun:
		log.Println("Walk/Run Action Pressed")
	case stamina:
		log.Println("Stamina Action Pressed")
	case miniPanel:
		log.Println("Mini Panel Action Pressed")
	case rightSelec:
		log.Println("Right Skill Selector Action Pressed")
	case rightSkill:
		log.Println("Right Skill Action Pressed")
	default:
		log.Printf("Unrecognized ActionableType(%d) being clicked\n", item)
	}
}
