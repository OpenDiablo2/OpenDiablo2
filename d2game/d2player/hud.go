package d2player

import (
	"fmt"
	"log"
	"strings"
	"image"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
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


type HUD struct {
	actionableRegions      []actionableRegion
	asset *d2asset.AssetManager
	uiManager *d2ui.UIManager
	help *HelpOverlay
	lastMouseX             int
	lastMouseY             int
	hero                   *d2mapentity.Player
	mainPanel              *d2ui.Sprite
	globeSprite            *d2ui.Sprite
	menuButton             *d2ui.Sprite
	hpManaStatusSprite     *d2ui.Sprite
	leftSkillResource      *SkillResource
	rightSkillResource     *SkillResource
	runButton              *d2ui.Button
	zoneChangeText         *d2ui.Label
	miniPanel              *miniPanel
	isZoneTextShown        bool
	nameLabel              *d2ui.Label
	hpStatsIsVisible       bool
	manaStatsIsVisible     bool
	hpManaStatsLabel       *d2ui.Label
	skillSelectMenu        *SkillSelectMenu
}

func NewHUD(
	asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	hero *d2mapentity.Player,
	help *HelpOverlay,
	miniPanel *miniPanel,
	actionableRegions      []actionableRegion,
) *HUD {
	nameLabel := ui.NewLabel(d2resource.Font16, d2resource.PaletteStatic)
	nameLabel.Alignment = d2gui.HorizontalAlignCenter
	nameLabel.SetText(d2ui.ColorTokenize("", d2ui.ColorTokenServer))

	hpManaStatsLabel := ui.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	hpManaStatsLabel.Alignment = d2gui.HorizontalAlignLeft

	zoneLabel := ui.NewLabel(d2resource.Font30, d2resource.PaletteUnits)
	zoneLabel.Alignment = d2gui.HorizontalAlignCenter

	return &HUD {
		asset: asset,
		uiManager: ui,
		hero: hero,
		help: help,
		miniPanel: miniPanel,
		actionableRegions: actionableRegions,
		nameLabel: nameLabel,
		skillSelectMenu: NewSkillSelectMenu(asset, ui, hero),
		hpManaStatsLabel: hpManaStatsLabel,
		zoneChangeText: zoneLabel,
	}
}

func (h *HUD) Load() {
	var err error

	h.globeSprite, err = h.uiManager.NewSprite(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	h.hpManaStatusSprite, err = h.uiManager.NewSprite(d2resource.HealthManaIndicator, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	h.menuButton, err = h.uiManager.NewSprite(d2resource.MenuButton, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	err = h.menuButton.SetCurrentFrame(frameMenuButton)
	if err != nil {
		log.Print(err)
	}


	h.mainPanel, err = h.uiManager.NewSprite(d2resource.GamePanels, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/799
	genericSkillsSprite, err := h.uiManager.NewSprite(d2resource.GenericSkills, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
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

	h.loadUIButtons()

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
func (h *HUD) renderGameControlPanelElements(target d2interface.Surface) error {
	_, height := target.GetSize()
	offsetX, offsetY := 0, 0

	// Main panel background
	offsetY = height
	if err := h.renderPanel(offsetX, offsetY, target); err != nil {
		return err
	}

	// Health globe
	w, _ := h.mainPanel.GetCurrentFrameSize()

	if err := h.renderHealthGlobe(offsetX, offsetY, target); err != nil {
		return err
	}

	// Left Skill
	offsetX += w
	if err := h.renderLeftSkill(offsetX, offsetY, target); err != nil {
		return err
	}

	// New Stats Button
	w, _ = h.leftSkillResource.SkillIcon.GetCurrentFrameSize()
	offsetX += w

	if err := h.renderNewStatsButton(offsetX, offsetY, target); err != nil {
		return err
	}

	// Stamina
	w, _ = h.mainPanel.GetCurrentFrameSize()
	offsetX += w

	if err := h.renderStamina(offsetX, offsetY, target); err != nil {
		return err
	}

	// Stamina status bar
	w, _ = h.mainPanel.GetCurrentFrameSize()
	offsetX += w

	if err := h.renderStaminaBar(target); err != nil {
		return err
	}

	// Experience status bar
	if err := h.renderExperienceBar(target); err != nil {
		return err
	}

	// Mini Panel and button
	if err := h.renderMiniPanel(target); err != nil {
		return err
	}

	// Potions
	if err := h.renderPotions(offsetX, offsetY, target); err != nil {
		return err
	}

	// New Skills Button
	w, _ = h.mainPanel.GetCurrentFrameSize()
	offsetX += w

	if err := h.renderNewSkillsButton(offsetX, offsetY, target); err != nil {
		return err
	}

	// Right skill
	w, _ = h.mainPanel.GetCurrentFrameSize()
	offsetX += w

	if err := h.renderRightSkill(offsetX, offsetY, target); err != nil {
		return err
	}

	// Mana Globe
	w, _ = h.rightSkillResource.SkillIcon.GetCurrentFrameSize()
	offsetX += w

	if err := h.renderManaGlobe(offsetX, offsetY, target); err != nil {
		return err
	}

	return nil
}

func (h *HUD) renderManaGlobe(x, _ int, target d2interface.Surface) error {
	_, height := target.GetSize()

	if err := h.mainPanel.SetCurrentFrame(frameRightGlobeHolder); err != nil {
		return err
	}

	h.mainPanel.SetPosition(x, height)

	h.mainPanel.Render(target)

	// Mana status bar
	manaPercent := float64(h.hero.Stats.Mana) / float64(h.hero.Stats.MaxMana)
	manaBarHeight := int(manaPercent * float64(globeHeight))

	if err := h.hpManaStatusSprite.SetCurrentFrame(frameManaStatus); err != nil {
		return err
	}

	h.hpManaStatusSprite.SetPosition(x+manaStatusOffsetX, height+manaStatusOffsetY)

	manaMaskRect := image.Rect(0, globeHeight-manaBarHeight, globeWidth, globeHeight)
	h.hpManaStatusSprite.RenderSection(target, manaMaskRect)

	// Right globe
	if err := h.globeSprite.SetCurrentFrame(frameRightGlobe); err != nil {
		return err
	}

	h.globeSprite.SetPosition(x+rightGlobeOffsetX, height+rightGlobeOffsetY)

	h.globeSprite.Render(target)
	h.globeSprite.Render(target)

	return nil
}

func (h *HUD) renderHealthGlobe(x, y int, target d2interface.Surface) error {
	healthPercent := float64(h.hero.Stats.Health) / float64(h.hero.Stats.MaxHealth)
	hpBarHeight := int(healthPercent * float64(globeHeight))

	if err := h.hpManaStatusSprite.SetCurrentFrame(0); err != nil {
		return err
	}

	h.hpManaStatusSprite.SetPosition(x+healthStatusOffsetX, y+healthStatusOffsetY)

	healthMaskRect := image.Rect(0, globeHeight-hpBarHeight, globeWidth, globeHeight)
	h.hpManaStatusSprite.RenderSection(target, healthMaskRect)

	// Left globe
	if err := h.globeSprite.SetCurrentFrame(frameHealthStatus); err != nil {
		return err
	}

	h.globeSprite.SetPosition(x+globeSpriteOffsetX, y+globeSpriteOffsetY)
	h.globeSprite.Render(target)

	return nil
}


func (h *HUD) renderPanel(x, y int, target d2interface.Surface) error {
	if err := h.mainPanel.SetCurrentFrame(0); err != nil {
		return err
	}

	h.mainPanel.SetPosition(x, y)
	h.mainPanel.Render(target)

	return nil
}

func (h *HUD) renderLeftSkill(x, y int, target d2interface.Surface) error {
	newSkillResourcePath := h.getSkillResourceByClass(h.hero.LeftSkill.Charclass)
	if newSkillResourcePath != h.leftSkillResource.SkillResourcePath {
		h.leftSkillResource.SkillResourcePath = newSkillResourcePath
		h.leftSkillResource.SkillIcon, _ = h.uiManager.NewSprite(newSkillResourcePath, d2resource.PaletteSky)
	}

	if err := h.leftSkillResource.SkillIcon.SetCurrentFrame(h.hero.LeftSkill.IconCel); err != nil {
		return err
	}

	h.leftSkillResource.SkillIcon.SetPosition(x, y)
	h.leftSkillResource.SkillIcon.Render(target)

	return nil
}

func (h *HUD) renderRightSkill(x, _ int, target d2interface.Surface) error {
	_, height := target.GetSize()

	newSkillResourcePath := h.getSkillResourceByClass(h.hero.RightSkill.Charclass)
	if newSkillResourcePath != h.rightSkillResource.SkillResourcePath {
		h.rightSkillResource.SkillIcon, _ = h.uiManager.NewSprite(newSkillResourcePath, d2resource.PaletteSky)
		h.rightSkillResource.SkillResourcePath = newSkillResourcePath
	}

	if err := h.rightSkillResource.SkillIcon.SetCurrentFrame(h.hero.RightSkill.IconCel); err != nil {
		return err
	}

	h.rightSkillResource.SkillIcon.SetPosition(x, height)
	h.rightSkillResource.SkillIcon.Render(target)

	return nil
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

func (h *HUD) renderStaminaBar(target d2interface.Surface) error {
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

	return nil
}

func (h *HUD) renderExperienceBar(target d2interface.Surface) error {
	target.PushTranslation(experienceBarOffsetX, experienceBarOffsetY)
	defer target.Pop()

	expPercent := float64(h.hero.Stats.Experience) / float64(h.hero.Stats.NextLevelExp)

	target.DrawRect(int(expPercent*expBarWidth), 2, d2util.Color(whiteAlpha100))

	return nil
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
		h.nameLabel.SetText(h.asset.TranslateString(stringTableKey))

		halfButtonWidth := rect.Width >> 1
		halfButtonHeight := rect.Height >> 1

		centerX := rect.Left + halfButtonWidth
		centerY := rect.Top + halfButtonHeight

		_, labelHeight := h.nameLabel.GetSize()

		labelX := centerX
		labelY := centerY - halfButtonHeight - labelHeight

		h.nameLabel.SetPosition(labelX, labelY)
		h.nameLabel.Render(target)
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

	h.hpManaStatsLabel.SetText(strPanelHealth)
	h.hpManaStatsLabel.SetPosition(hpLabelX, hpLabelY)
	h.hpManaStatsLabel.Render(target)
}

func (h *HUD) renderManaTooltip(target d2interface.Surface) {
	mx, my := h.lastMouseX, h.lastMouseY

	// Create and format Mana string from string lookup table.
	fmtMana := h.asset.TranslateString("panelmana")
	manaCurr, manaMax := h.hero.Stats.Mana, h.hero.Stats.MaxMana
	strPanelMana := fmt.Sprintf(fmtMana, manaCurr, manaMax)

	if !(h.actionableRegions[manaGlobe].rect.IsInRect(mx, my) || h.manaStatsIsVisible) {
		return
	}

	h.hpManaStatsLabel.SetText(strPanelMana)
	// In case if the mana value gets higher, we need to shift the
	// label to the left a little, hence widthManaLabel.
	widthManaLabel, _ := h.hpManaStatsLabel.GetSize()
	xManaLabel := manaLabelX - widthManaLabel
	h.hpManaStatsLabel.SetPosition(xManaLabel, manaLabelY)
	h.hpManaStatsLabel.Render(target)
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

	h.nameLabel.SetText(h.asset.TranslateString(stringTableKey))

	rect := &h.actionableRegions[walkRun].rect

	halfButtonWidth := rect.Width >> 1
	halfButtonHeight := rect.Height >> 1

	centerX := rect.Left + halfButtonWidth
	centerY := rect.Top + halfButtonHeight

	_, labelHeight := h.nameLabel.GetSize()

	labelX := centerX
	labelY := centerY - halfButtonHeight - labelHeight

	h.nameLabel.SetPosition(labelX, labelY)
	h.nameLabel.Render(target)
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

	h.nameLabel.SetText(strPanelStamina)

	rect := &h.actionableRegions[stamina].rect

	halfButtonWidth := rect.Width >> 1
	centerX := rect.Left + halfButtonWidth

	_, labelHeight := h.nameLabel.GetSize()
	halfLabelHeight := labelHeight >> 1

	labelX := centerX
	labelY := staminaExperienceY - halfLabelHeight

	h.nameLabel.SetPosition(labelX, labelY)
	h.nameLabel.Render(target)
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

	h.nameLabel.SetText(strPanelExp)
	rect := &h.actionableRegions[stamina].rect

	halfButtonWidth := rect.Width >> 1
	centerX := rect.Left + halfButtonWidth

	_, labelHeight := h.nameLabel.GetSize()
	halfLabelHeight := labelHeight >> 1

	labelX := centerX
	labelY := staminaExperienceY - halfLabelHeight

	h.nameLabel.SetPosition(labelX, labelY)
	h.nameLabel.Render(target)
}



func (h *HUD) Render(target d2interface.Surface) error {
	if err := h.renderGameControlPanelElements(target); err != nil {
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
		log.Fatalf("Unknown class token: '%s'", class)
	}

	return entry
}

func (h *HUD) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	mx, my := event.X(), event.Y()
	h.lastMouseX = mx
	h.lastMouseY = my

	h.skillSelectMenu.LeftPanel.HandleMouseMove(mx, my)
	h.skillSelectMenu.RightPanel.HandleMouseMove(mx, my)

	return false
}

