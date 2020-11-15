package d2player

import (
	"fmt"
	"math"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	runButtonX = 255
	runButtonY = 570
)

const (
	zoneChangeTextX = screenWidth / 2
	zoneChangeTextY = screenHeight / 4
)

const (
	expBarWidth          = 120.0
	expBarHeight         = 4
	staminaBarWidth      = 102.0
	staminaBarHeight     = 19.0
	hoverLabelOuterPad   = 5
	percentStaminaBarLow = 0.25
)

const (
	hpLabelX = 15
	hpLabelY = 487

	manaLabelX = 785
	manaLabelY = 487

	staminaExperienceY = 535
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

// HUD represents the always visible user interface of the game
type HUD struct {
	actionableRegions  []actionableRegion
	asset              *d2asset.AssetManager
	uiManager          *d2ui.UIManager
	help               *HelpOverlay
	mapEngine          *d2mapengine.MapEngine
	mapRenderer        *d2maprenderer.MapRenderer
	lastMouseX         int
	lastMouseY         int
	hero               *d2mapentity.Player
	mainPanel          *d2ui.Sprite
	globeSprite        *d2ui.Sprite
	menuButton         *d2ui.Sprite
	hpManaStatusSprite *d2ui.Sprite
	leftSkillResource  *SkillResource
	rightSkillResource *SkillResource
	runButton          *d2ui.Button
	zoneChangeText     *d2ui.Label
	miniPanel          *miniPanel
	isZoneTextShown    bool
	hpStatsIsVisible   bool
	manaStatsIsVisible bool
	skillSelectMenu    *SkillSelectMenu
	staminaTooltip     *d2ui.Tooltip
	runWalkTooltip     *d2ui.Tooltip
	experienceTooltip  *d2ui.Tooltip
	healthTooltip      *d2ui.Tooltip
	manaTooltip        *d2ui.Tooltip
	miniPanelTooltip   *d2ui.Tooltip
	nameLabel          *d2ui.Label
	healthGlobe        *globeWidget
	manaGlobe          *globeWidget
	widgetStamina      *d2ui.CustomWidget
	widgetExperience   *d2ui.CustomWidget
	widgetLeftSkill    *d2ui.CustomWidget
	widgetRightSkill   *d2ui.CustomWidget
	panelBackground    *d2ui.CustomWidget
	logger             *d2util.Logger
}

// NewHUD creates a HUD object
func NewHUD(
	asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	hero *d2mapentity.Player,
	help *HelpOverlay,
	miniPanel *miniPanel,
	actionableRegions []actionableRegion,
	mapEngine *d2mapengine.MapEngine,
	l d2util.LogLevel,
	mapRenderer *d2maprenderer.MapRenderer,
) *HUD {
	nameLabel := ui.NewLabel(d2resource.Font16, d2resource.PaletteStatic)
	nameLabel.Alignment = d2ui.HorizontalAlignCenter
	nameLabel.SetText(d2ui.ColorTokenize("", d2ui.ColorTokenServer))

	zoneLabel := ui.NewLabel(d2resource.Font30, d2resource.PaletteUnits)
	zoneLabel.Alignment = d2ui.HorizontalAlignCenter

	healthGlobe := newGlobeWidget(ui, 0, screenHeight, typeHealthGlobe, &hero.Stats.Health, l, &hero.Stats.MaxHealth)
	manaGlobe := newGlobeWidget(ui, screenWidth-manaGlobeScreenOffsetX, screenHeight, typeManaGlobe, &hero.Stats.Mana, l, &hero.Stats.MaxMana)

	hud := &HUD{
		asset:             asset,
		uiManager:         ui,
		hero:              hero,
		help:              help,
		mapEngine:         mapEngine,
		mapRenderer:       mapRenderer,
		miniPanel:         miniPanel,
		actionableRegions: actionableRegions,
		nameLabel:         nameLabel,
		skillSelectMenu:   NewSkillSelectMenu(asset, ui, l, hero),
		zoneChangeText:    zoneLabel,
		healthGlobe:       healthGlobe,
		manaGlobe:         manaGlobe,
	}

	hud.logger = d2util.NewLogger()
	hud.logger.SetPrefix(logPrefix)
	hud.logger.SetLevel(l)

	return hud
}

// Load creates the ui elemets
func (h *HUD) Load() {
	h.loadSprites()

	h.healthGlobe.load()
	h.manaGlobe.load()

	h.loadSkillResources()
	h.loadCustomWidgets()
	h.loadUIButtons()
	h.loadTooltips()
}

func (h *HUD) loadCustomWidgets() {
	// static background
	_, height, err := h.mainPanel.GetFrameSize(0) // health globe is the frame with max height
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	h.panelBackground = h.uiManager.NewCustomWidgetCached(h.renderPanelStatic, screenWidth, height)
	h.panelBackground.SetPosition(0, screenHeight-height)

	// stamina bar
	h.widgetStamina = h.uiManager.NewCustomWidget(h.renderStaminaBar, staminaBarWidth, staminaBarHeight)
	h.widgetStamina.SetPosition(staminaBarOffsetX, staminaBarOffsetY)

	// experience bar
	h.widgetExperience = h.uiManager.NewCustomWidget(h.renderExperienceBar, expBarWidth, expBarHeight)

	// Left skill widget
	leftRenderFunc := func(target d2interface.Surface) {
		x, y := h.widgetLeftSkill.GetPosition()
		h.renderLeftSkill(x, y, target)
	}

	h.widgetLeftSkill = h.uiManager.NewCustomWidget(leftRenderFunc, skillIconWidth, skillIconHeight)
	h.widgetLeftSkill.SetPosition(leftSkillX, screenHeight)

	// Right skill widget
	rightRenderFunc := func(target d2interface.Surface) {
		x, y := h.widgetRightSkill.GetPosition()
		h.renderRightSkill(x, y, target)
	}

	h.widgetRightSkill = h.uiManager.NewCustomWidget(rightRenderFunc, skillIconWidth, skillIconHeight)
	h.widgetRightSkill.SetPosition(rightSkillX, screenHeight)
}

func (h *HUD) loadSkillResources() {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/799
	genericSkillsSprite, err := h.uiManager.NewSprite(d2resource.GenericSkills, d2resource.PaletteSky)
	if err != nil {
		h.logger.Error(err.Error())
	}

	attackIconID := 2

	h.leftSkillResource = &SkillResource{
		SkillIcon:         genericSkillsSprite,
		IconNumber:        attackIconID,
		SkillResourcePath: d2resource.GenericSkills,
	}

	h.rightSkillResource = &SkillResource{
		SkillIcon:         genericSkillsSprite,
		IconNumber:        attackIconID,
		SkillResourcePath: d2resource.GenericSkills,
	}
}

func (h *HUD) loadSprites() {
	var err error

	h.globeSprite, err = h.uiManager.NewSprite(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	if err != nil {
		h.logger.Error(err.Error())
	}

	h.hpManaStatusSprite, err = h.uiManager.NewSprite(d2resource.HealthManaIndicator, d2resource.PaletteSky)
	if err != nil {
		h.logger.Error(err.Error())
	}

	h.menuButton, err = h.uiManager.NewSprite(d2resource.MenuButton, d2resource.PaletteSky)
	if err != nil {
		h.logger.Error(err.Error())
	}

	err = h.menuButton.SetCurrentFrame(frameMenuButton)
	if err != nil {
		h.logger.Error(err.Error())
	}

	h.mainPanel, err = h.uiManager.NewSprite(d2resource.GamePanels, d2resource.PaletteSky)
	if err != nil {
		h.logger.Error(err.Error())
	}
}

func (h *HUD) loadTooltips() {
	// stamina tooltip
	h.staminaTooltip = h.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	rect := &h.actionableRegions[stamina].rect

	halfButtonWidth := rect.Width >> 1
	centerX := rect.Left + halfButtonWidth

	_, labelHeight := h.staminaTooltip.GetSize()
	halfLabelHeight := labelHeight >> 1

	labelX := centerX
	labelY := staminaExperienceY - halfLabelHeight
	h.staminaTooltip.SetPosition(labelX, labelY)

	// runwalk tooltip
	h.runWalkTooltip = h.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYBottom)
	rect = &h.actionableRegions[walkRun].rect

	halfButtonWidth = rect.Width >> 1
	halfButtonHeight := rect.Height >> 1

	centerX = rect.Left + halfButtonWidth
	centerY := rect.Top + halfButtonHeight

	_, labelHeight = h.runWalkTooltip.GetSize()
	labelX = centerX
	labelY = centerY - halfButtonHeight - labelHeight
	h.runWalkTooltip.SetPosition(labelX, labelY)

	// experience tooltip
	h.experienceTooltip = h.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteSky, d2ui.TooltipXCenter, d2ui.TooltipYTop)
	rect = &h.actionableRegions[stamina].rect

	halfButtonWidth = rect.Width >> 1
	centerX = rect.Left + halfButtonWidth

	_, labelHeight = h.experienceTooltip.GetSize()
	halfLabelHeight = labelHeight >> 1

	labelX = centerX
	labelY = staminaExperienceY - halfLabelHeight
	h.experienceTooltip.SetPosition(labelX, labelY)

	// Health tooltip
	h.healthTooltip = h.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteUnits, d2ui.TooltipXLeft, d2ui.TooltipYTop)
	h.healthTooltip.SetPosition(hpLabelX, hpLabelY)
	h.healthTooltip.SetBoxEnabled(false)

	// Health tooltip
	h.manaTooltip = h.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteUnits, d2ui.TooltipXLeft, d2ui.TooltipYTop)
	h.manaTooltip.SetPosition(manaLabelX, manaLabelY)
	h.manaTooltip.SetBoxEnabled(false)

	// minipanel tooltip
	h.miniPanelTooltip = h.uiManager.NewTooltip(d2resource.Font16, d2resource.PaletteUnits, d2ui.TooltipXCenter, d2ui.TooltipYTop)
}

func (h *HUD) loadUIButtons() {
	// Run button
	h.runButton = h.uiManager.NewButton(d2ui.ButtonTypeRun, "")

	h.runButton.SetPosition(runButtonX, runButtonY)
	h.runButton.OnActivated(func() { h.onToggleRunButton(false) })

	if h.hero.IsRunToggled() {
		h.runButton.Toggle()
	}
}

func (h *HUD) onToggleRunButton(noButton bool) {
	if !noButton {
		h.runButton.Toggle()
	}

	h.hero.ToggleRunWalk()

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/800
	h.hero.SetIsRunning(h.hero.IsRunToggled())
}

// NOTE: the positioning of all of the panel elements is coupled to the rendering order :(
// don't change the order in which the render methods are called, as there is an x,y offset
// that is updated between render calls
func (h *HUD) renderPanelStatic(target d2interface.Surface) {
	_, height := target.GetSize()
	offsetX, offsetY := 0, height

	// Main panel background
	if err := h.renderPanel(offsetX, offsetY, target); err != nil {
		h.logger.Error(err.Error())
		return
	}

	// New Stats Button
	w, _ := h.mainPanel.GetCurrentFrameSize()
	offsetX += w + skillIconWidth

	if err := h.renderNewStatsButton(offsetX, offsetY, target); err != nil {
		h.logger.Error(err.Error())
		return
	}

	// Stamina
	w, _ = h.mainPanel.GetCurrentFrameSize()
	offsetX += w

	if err := h.renderStamina(offsetX, offsetY, target); err != nil {
		h.logger.Error(err.Error())
		return
	}

	// Potions
	w, _ = h.mainPanel.GetCurrentFrameSize()
	offsetX += w

	if err := h.renderPotions(offsetX, offsetY, target); err != nil {
		h.logger.Error(err.Error())
		return
	}

	// New Skills Button
	w, _ = h.mainPanel.GetCurrentFrameSize()
	offsetX += w

	if err := h.renderNewSkillsButton(offsetX, offsetY, target); err != nil {
		h.logger.Error(err.Error())
		return
	}

	// Empty Mana Globe
	w, _ = h.mainPanel.GetCurrentFrameSize()
	offsetX += w + skillIconWidth

	if err := h.mainPanel.SetCurrentFrame(frameRightGlobeHolder); err != nil {
		h.logger.Error(err.Error())
		return
	}

	h.mainPanel.SetPosition(offsetX, height)
	h.mainPanel.Render(target)
}

func (h *HUD) renderPanel(x, y int, target d2interface.Surface) error {
	if err := h.mainPanel.SetCurrentFrame(0); err != nil {
		return err
	}

	h.mainPanel.SetPosition(x, y)
	h.mainPanel.Render(target)

	return nil
}

func (h *HUD) renderLeftSkill(x, y int, target d2interface.Surface) {
	newSkillResourcePath := h.getSkillResourceByClass(h.hero.LeftSkill.Charclass)
	if newSkillResourcePath != h.leftSkillResource.SkillResourcePath {
		h.leftSkillResource.SkillResourcePath = newSkillResourcePath
		h.leftSkillResource.SkillIcon, _ = h.uiManager.NewSprite(newSkillResourcePath, d2resource.PaletteSky)
	}

	if err := h.leftSkillResource.SkillIcon.SetCurrentFrame(h.hero.LeftSkill.IconCel); err != nil {
		h.logger.Error(err.Error())
		return
	}

	h.leftSkillResource.SkillIcon.SetPosition(x, y)
	h.leftSkillResource.SkillIcon.Render(target)
}

func (h *HUD) renderRightSkill(x, _ int, target d2interface.Surface) {
	_, height := target.GetSize()

	newSkillResourcePath := h.getSkillResourceByClass(h.hero.RightSkill.Charclass)
	if newSkillResourcePath != h.rightSkillResource.SkillResourcePath {
		h.rightSkillResource.SkillIcon, _ = h.uiManager.NewSprite(newSkillResourcePath, d2resource.PaletteSky)
		h.rightSkillResource.SkillResourcePath = newSkillResourcePath
	}

	if err := h.rightSkillResource.SkillIcon.SetCurrentFrame(h.hero.RightSkill.IconCel); err != nil {
		h.logger.Error(err.Error())
		return
	}

	h.rightSkillResource.SkillIcon.SetPosition(x, height)
	h.rightSkillResource.SkillIcon.Render(target)
}

func (h *HUD) renderNewStatsButton(x, y int, target d2interface.Surface) error {
	if err := h.mainPanel.SetCurrentFrame(frameNewStatsSelector); err != nil {
		return err
	}

	h.mainPanel.SetPosition(x, y)
	h.mainPanel.Render(target)

	return nil
}

func (h *HUD) renderStamina(x, y int, target d2interface.Surface) error {
	if err := h.mainPanel.SetCurrentFrame(frameStamina); err != nil {
		return err
	}

	h.mainPanel.SetPosition(x, y)
	h.mainPanel.Render(target)

	return nil
}

func (h *HUD) renderStaminaBar(target d2interface.Surface) {
	target.PushTranslation(staminaBarOffsetX, staminaBarOffsetY)
	defer target.Pop()

	target.PushEffect(d2enum.DrawEffectModulate)
	defer target.Pop()

	staminaPercent := h.hero.Stats.Stamina / float64(h.hero.Stats.MaxStamina)

	staminaBarColor := d2util.Color(lightBrownAlpha72)
	if staminaPercent < percentStaminaBarLow {
		staminaBarColor = d2util.Color(redAlpha72)
	}

	target.DrawRect(int(staminaPercent*staminaBarWidth), staminaBarHeight, staminaBarColor)
}

func (h *HUD) renderExperienceBar(target d2interface.Surface) {
	target.PushTranslation(experienceBarOffsetX, experienceBarOffsetY)
	defer target.Pop()

	expPercent := float64(h.hero.Stats.Experience) / float64(h.hero.Stats.NextLevelExp)

	target.DrawRect(int(expPercent*expBarWidth), 2, d2util.Color(whiteAlpha100))
}

func (h *HUD) renderMiniPanel(target d2interface.Surface) error {
	width, height := target.GetSize()
	mx, my := h.lastMouseX, h.lastMouseY

	menuButtonFrameIndex := 0
	if h.miniPanel.isOpen {
		menuButtonFrameIndex = 2
	}

	if err := h.menuButton.SetCurrentFrame(menuButtonFrameIndex); err != nil {
		return err
	}

	buttonX, buttonY := (width>>1)+miniPanelButtonOffsetX, height+miniPanelButtonOffsetY

	h.menuButton.SetPosition(buttonX, buttonY)
	h.menuButton.Render(target)
	h.miniPanel.Render(target)

	miniPanelButtons := map[actionableType]string{
		miniPanelCharacter:  "minipanelchar",
		miniPanelInventory:  "minipanelinv",
		miniPanelSkillTree:  "minipaneltree",
		miniPanelAutomap:    "minipanelautomap",
		miniPanelMessageLog: "minipanelmessage",
		miniPanelQuestLog:   "minipanelquest",
		miniPanelGameMenu:   "minipanelmenubtn",
	}

	if !h.miniPanel.IsOpen() {
		return nil
	}

	for miniPanelButton, stringTableKey := range miniPanelButtons {
		if !h.actionableRegions[miniPanelButton].rect.IsInRect(mx, my) {
			continue
		}

		rect := &h.actionableRegions[miniPanelButton].rect
		h.miniPanelTooltip.SetText(h.asset.TranslateString(stringTableKey))

		halfButtonWidth := rect.Width >> 1
		halfButtonHeight := rect.Height >> 1

		centerX := rect.Left + halfButtonWidth
		centerY := rect.Top + halfButtonHeight

		_, labelHeight := h.miniPanelTooltip.GetSize()

		labelX := centerX
		labelY := centerY - halfButtonHeight - labelHeight

		h.miniPanelTooltip.SetPosition(labelX, labelY)
		h.miniPanelTooltip.Render(target)
	}

	return nil
}

func (h *HUD) renderPotions(x, _ int, target d2interface.Surface) error {
	_, height := target.GetSize()

	if err := h.mainPanel.SetCurrentFrame(framePotions); err != nil {
		return err
	}

	h.mainPanel.SetPosition(x, height)
	h.mainPanel.Render(target)

	return nil
}

func (h *HUD) renderNewSkillsButton(x, _ int, target d2interface.Surface) error {
	_, height := target.GetSize()

	if err := h.mainPanel.SetCurrentFrame(frameNewSkillsSelector); err != nil {
		return err
	}

	h.mainPanel.SetPosition(x, height)
	h.mainPanel.Render(target)

	return nil
}

//nolint:golint,dupl // we clean this up later
func (h *HUD) renderHealthTooltip(target d2interface.Surface) {
	mx, my := h.lastMouseX, h.lastMouseY

	// Create and format Health string from string lookup table.
	fmtHealth := h.asset.TranslateString("panelhealth")
	healthCurr, healthMax := h.hero.Stats.Health, h.hero.Stats.MaxHealth
	strPanelHealth := fmt.Sprintf(fmtHealth, healthCurr, healthMax)

	// Display current hp and mana stats hpGlobe or manaGlobe region is clicked
	if !(h.actionableRegions[hpGlobe].rect.IsInRect(mx, my) || h.hpStatsIsVisible) {
		return
	}

	h.healthTooltip.SetText(strPanelHealth)
	h.healthTooltip.Render(target)
}

//nolint:golint,dupl // we clean this up later
func (h *HUD) renderManaTooltip(target d2interface.Surface) {
	mx, my := h.lastMouseX, h.lastMouseY

	// Create and format Mana string from string lookup table.
	fmtMana := h.asset.TranslateString("panelmana")
	manaCurr, manaMax := h.hero.Stats.Mana, h.hero.Stats.MaxMana
	strPanelMana := fmt.Sprintf(fmtMana, manaCurr, manaMax)

	if !(h.actionableRegions[manaGlobe].rect.IsInRect(mx, my) || h.manaStatsIsVisible) {
		return
	}

	h.manaTooltip.SetText(strPanelMana)
	h.manaTooltip.Render(target)
}

func (h *HUD) renderRunWalkTooltip(target d2interface.Surface) {
	mx, my := h.lastMouseX, h.lastMouseY

	// Display run/walk tooltip when hovered.
	// Note that whether the player is walking or running, the tooltip is the same in Diablo 2.
	if !h.actionableRegions[walkRun].rect.IsInRect(mx, my) {
		return
	}

	var stringTableKey string

	if h.hero.IsRunToggled() {
		stringTableKey = "RunOff"
	} else {
		stringTableKey = "RunOn"
	}

	h.runWalkTooltip.SetText(h.asset.TranslateString(stringTableKey))
	h.runWalkTooltip.Render(target)
}

func (h *HUD) renderStaminaTooltip(target d2interface.Surface) {
	mx, my := h.lastMouseX, h.lastMouseY

	// Display stamina tooltip when hovered.
	if !h.actionableRegions[stamina].rect.IsInRect(mx, my) {
		return
	}

	// Create and format Stamina string from string lookup table.
	fmtStamina := h.asset.TranslateString("panelstamina")
	staminaCurr, staminaMax := int(h.hero.Stats.Stamina), h.hero.Stats.MaxStamina
	strPanelStamina := fmt.Sprintf(fmtStamina, staminaCurr, staminaMax)

	h.staminaTooltip.SetText(strPanelStamina)
	h.staminaTooltip.Render(target)
}

func (h *HUD) renderExperienceTooltip(target d2interface.Surface) {
	mx, my := h.lastMouseX, h.lastMouseY

	// Display experience tooltip when hovered.
	if !h.actionableRegions[xp].rect.IsInRect(mx, my) {
		return
	}

	// Create and format Experience string from string lookup table.
	fmtExp := h.asset.TranslateString("panelexp")

	// The English string for "panelexp" is "Experience: %u / %u", however %u doesn't
	// translate well. So we need to rewrite %u into a formatable Go verb. %d is used in other
	// strings, so we go with that, keeping in mind that %u likely referred to
	// an unsigned integer.
	fmtExp = strings.ReplaceAll(fmtExp, "%u", "%d")

	expCurr, expMax := uint(h.hero.Stats.Experience), uint(h.hero.Stats.NextLevelExp)
	strPanelExp := fmt.Sprintf(fmtExp, expCurr, expMax)

	h.experienceTooltip.SetText(strPanelExp)
	h.experienceTooltip.Render(target)
}

func (h *HUD) renderForSelectableEntitiesHovered(target d2interface.Surface) {
	mx, my := h.lastMouseX, h.lastMouseY

	for entityIdx := range h.mapEngine.Entities() {
		entity := (h.mapEngine.Entities())[entityIdx]
		if !entity.Selectable() {
			continue
		}

		entPos := entity.GetPosition()
		entOffset := entPos.RenderOffset()
		entScreenXf, entScreenYf := h.mapRenderer.WorldToScreenF(entity.GetPositionF())
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

			h.nameLabel.SetText(entity.Label())

			xLabel, yLabel := entScreenX-xOff, entScreenY-yOff-entityHeight-hoverLabelOuterPad
			h.nameLabel.SetPosition(xLabel, yLabel)

			h.nameLabel.Render(target)
			entity.Highlight()

			break
		}
	}
}

// Render draws the HUD to the screen
func (h *HUD) Render(target d2interface.Surface) error {
	h.renderForSelectableEntitiesHovered(target)

	h.panelBackground.Render(target)

	h.healthGlobe.Render(target)
	h.widgetLeftSkill.Render(target)
	h.widgetRightSkill.Render(target)
	h.manaGlobe.Render(target)
	h.widgetStamina.Render(target)
	h.widgetExperience.Render(target)

	// Mini Panel and button
	if err := h.renderMiniPanel(target); err != nil {
		return err
	}

	if err := h.help.Render(target); err != nil {
		return err
	}

	if h.isZoneTextShown {
		h.zoneChangeText.SetPosition(zoneChangeTextX, zoneChangeTextY)
		h.zoneChangeText.Render(target)
	}

	h.renderHealthTooltip(target)
	h.renderManaTooltip(target)
	h.renderRunWalkTooltip(target)
	h.renderStaminaTooltip(target)
	h.renderExperienceTooltip(target)

	if h.skillSelectMenu.IsOpen() {
		h.skillSelectMenu.Render(target)
	}

	return nil
}

func (h *HUD) getSkillResourceByClass(class string) string {
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
		h.logger.Error("Unknown class token: '%s'" + class)
	}

	return entry
}

// OnMouseMove handles mouse move events
func (h *HUD) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	mx, my := event.X(), event.Y()
	h.lastMouseX = mx
	h.lastMouseY = my

	h.skillSelectMenu.LeftPanel.HandleMouseMove(mx, my)
	h.skillSelectMenu.RightPanel.HandleMouseMove(mx, my)

	return false
}
