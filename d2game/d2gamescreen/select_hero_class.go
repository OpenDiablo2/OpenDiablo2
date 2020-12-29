package d2gamescreen

import (
	"image"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
)

const (
	millisecondsPerSecond = 1000.0
)

type heroRenderConfig struct {
	idleAnimationPath               string
	idleSelectedAnimationPath       string
	forwardWalkAnimationPath        string
	forwardWalkOverlayAnimationPath string
	forwardWalkOverlayBlend         bool
	selectedAnimationPath           string
	selectedOverlayAnimationPath    string
	backWalkAnimationPath           string
	backWalkOverlayAnimationPath    string
	selectionBounds                 image.Rectangle
	selectSfx                       string
	deselectSfx                     string
	position                        image.Point
	idlePlayLengthMs                int
	forwardWalkPlayLengthMs         int
	backWalkPlayLengthMs            int
}

func point(x, y int) image.Point {
	return image.Point{X: x, Y: y}
}

func rect(x1, y1, x2, y2 int) image.Rectangle {
	return image.Rectangle{Min: point(x1, y1), Max: point(x2, y2)}
}

// animation position, selection box bound, animation play lengths in ms
const (
	barbPosX, barbPosY                                     = 400, 330
	barbRectMinX, barbRectMinY, barbRectMaxX, barbRectMaxY = 364, 201, 90, 170
	barbIdleLength, barbForwardLength, barbBackLength      = 0, 2500, 1000

	sorcPosX, sorcPosY                                     = 626, 352
	sorcRectMinX, sorcRectMinY, sorcRectMaxX, sorcRectMaxY = 580, 240, 65, 160
	sorcIdleLength, sorcForwardLength, sorcBackLength      = 2500, 2300, 1200

	necPosX, necPosY                                   = 300, 335
	necRectMinX, necRectMinY, necRectMaxX, necRectMaxY = 265, 220, 55, 175
	necIdleLength, necForwardLength, necBackLength     = 1200, 2000, 1500

	palPosX, palPosY                                   = 521, 338
	palRectMinX, palRectMinY, palRectMaxX, palRectMaxY = 490, 210, 65, 180
	palIdleLength, palForwardLength, palBackLength     = 2500, 3400, 1300

	amaPosX, amaPosY                                   = 100, 339
	amaRectMinX, amaRectMinY, amaRectMaxX, amaRectMaxY = 70, 220, 55, 200
	amaIdleLength, amaForwardLength, amaBackLength     = 2500, 2200, 1500

	assPosX, assPosY                                   = 231, 365
	assRectMinX, assRectMinY, assRectMaxX, assRectMaxY = 175, 235, 50, 180
	assIdleLength, assForwardLength, assBackLength     = 2500, 3800, 1500

	druPosX, druPosY                                   = 720, 370
	druRectMinX, druRectMinY, druRectMaxX, druRectMaxY = 680, 220, 70, 195
	druIdleLength, druForwardLength, druBackLength     = 1500, 4800, 1500

	campfirePosX, campfirePosY = 380, 335
)

// label and button positions
const (
	headingX, headingY               = 400, 17
	heroClassLabelX, heroClassLabelY = 400, 65
	heroDescLine1X, heroDescLine1Y   = 400, 100
	heroDescLine2X, heroDescLine2Y   = 400, 115
	heroDescLine3X, heroDescLine3Y   = 400, 130
	heroNameLabelX, heroNameLabelY   = 321, 475
	expansionLabelX, expansionLabelY = 339, 526
	hardcoreLabelX, hardcoreLabelY   = 339, 548

	selHeroExitBtnX, selHeroExitBtnY = 33, 537
	selHeroOkBtnX, selHeroOkBtnY     = 630, 537

	heroNameTextBoxX, heoNameTextBoxY       = 318, 493
	expandsionCheckboxX, expansionCheckboxY = 318, 526
	hardcoreCheckoxX, hardcoreCheckboxY     = 318, 548
)

const heroDescCharWidth = 37

//nolint:funlen // this func returns a map of structs and the structs are big, deal with it
func getHeroRenderConfiguration() map[d2enum.Hero]*heroRenderConfig {
	configs := make(map[d2enum.Hero]*heroRenderConfig)

	configs[d2enum.HeroBarbarian] = &heroRenderConfig{
		d2resource.CharacterSelectBarbarianUnselected,
		d2resource.CharacterSelectBarbarianUnselectedH,
		d2resource.CharacterSelectBarbarianForwardWalk,
		d2resource.CharacterSelectBarbarianForwardWalkOverlay,
		false,
		d2resource.CharacterSelectBarbarianSelected,
		"",
		d2resource.CharacterSelectBarbarianBackWalk,
		"",
		rect(barbRectMinX, barbRectMinY, barbRectMaxX, barbRectMaxY),
		d2resource.SFXBarbarianSelect,
		d2resource.SFXBarbarianDeselect,
		point(barbPosX, barbPosY),
		barbIdleLength,
		barbForwardLength,
		barbBackLength,
	}

	configs[d2enum.HeroSorceress] = &heroRenderConfig{
		d2resource.CharacterSelectSorceressUnselected,
		d2resource.CharacterSelectSorceressUnselectedH,
		d2resource.CharacterSelectSorceressForwardWalk,
		d2resource.CharacterSelectSorceressForwardWalkOverlay,
		true,
		d2resource.CharacterSelectSorceressSelected,
		d2resource.CharacterSelectSorceressSelectedOverlay,
		d2resource.CharacterSelectSorceressBackWalk,
		d2resource.CharacterSelectSorceressBackWalkOverlay,
		rect(sorcRectMinX, sorcRectMinY, sorcRectMaxX, sorcRectMaxY),
		d2resource.SFXSorceressSelect,
		d2resource.SFXSorceressDeselect,
		point(sorcPosX, sorcPosY),
		sorcIdleLength,
		sorcForwardLength,
		sorcBackLength,
	}

	configs[d2enum.HeroNecromancer] = &heroRenderConfig{
		d2resource.CharacterSelectNecromancerUnselected,
		d2resource.CharacterSelectNecromancerUnselectedH,
		d2resource.CharacterSelectNecromancerForwardWalk,
		d2resource.CharacterSelectNecromancerForwardWalkOverlay,
		true,
		d2resource.CharacterSelectNecromancerSelected,
		d2resource.CharacterSelectNecromancerSelectedOverlay,
		d2resource.CharacterSelectNecromancerBackWalk,
		d2resource.CharacterSelectNecromancerBackWalkOverlay,
		rect(necRectMinX, necRectMinY, necRectMaxX, necRectMaxY),
		d2resource.SFXNecromancerSelect,
		d2resource.SFXNecromancerDeselect,
		point(necPosX, necPosY),
		necIdleLength,
		necForwardLength,
		necBackLength,
	}

	configs[d2enum.HeroPaladin] = &heroRenderConfig{
		d2resource.CharacterSelectPaladinUnselected,
		d2resource.CharacterSelectPaladinUnselectedH,
		d2resource.CharacterSelectPaladinForwardWalk,
		d2resource.CharacterSelectPaladinForwardWalkOverlay,
		false,
		d2resource.CharacterSelectPaladinSelected,
		"",
		d2resource.CharacterSelectPaladinBackWalk,
		"",
		rect(palRectMinX, palRectMinY, palRectMaxX, palRectMaxY),
		d2resource.SFXPaladinSelect,
		d2resource.SFXPaladinDeselect,
		point(palPosX, palPosY),
		palIdleLength,
		palForwardLength,
		palBackLength,
	}

	configs[d2enum.HeroAmazon] = &heroRenderConfig{
		d2resource.CharacterSelectAmazonUnselected,
		d2resource.CharacterSelectAmazonUnselectedH,
		d2resource.CharacterSelectAmazonForwardWalk,
		"",
		false,
		d2resource.CharacterSelectAmazonSelected,
		"",
		d2resource.CharacterSelectAmazonBackWalk,
		"",
		rect(amaRectMinX, amaRectMinY, amaRectMaxX, amaRectMaxY),
		d2resource.SFXAmazonSelect,
		d2resource.SFXAmazonDeselect,
		point(amaPosX, amaPosY),
		amaIdleLength,
		amaForwardLength,
		amaBackLength,
	}

	configs[d2enum.HeroAssassin] = &heroRenderConfig{
		d2resource.CharacterSelectAssassinUnselected,
		d2resource.CharacterSelectAssassinUnselectedH,
		d2resource.CharacterSelectAssassinForwardWalk,
		"",
		false,
		d2resource.CharacterSelectAssassinSelected,
		"",
		d2resource.CharacterSelectAssassinBackWalk,
		"",
		rect(assRectMinX, assRectMinY, assRectMaxX, assRectMaxY),
		d2resource.SFXAssassinSelect,
		d2resource.SFXAssassinDeselect,
		point(assPosX, assPosY),
		assIdleLength,
		assForwardLength,
		assBackLength,
	}

	configs[d2enum.HeroDruid] = &heroRenderConfig{
		d2resource.CharacterSelectDruidUnselected,
		d2resource.CharacterSelectDruidUnselectedH,
		d2resource.CharacterSelectDruidForwardWalk,
		"",
		false,
		d2resource.CharacterSelectDruidSelected,
		"",
		d2resource.CharacterSelectDruidBackWalk,
		"",
		rect(druRectMinX, druRectMinY, druRectMaxX, druRectMaxY),
		d2resource.SFXDruidSelect,
		d2resource.SFXDruidDeselect,
		point(druPosX, druPosY),
		druIdleLength,
		druForwardLength,
		druBackLength,
	}

	return configs
}

// HeroRenderInfo stores the rendering information of a hero for the Select Hero Class screen
type HeroRenderInfo struct {
	Stance                   d2enum.HeroStance
	IdleSprite               *d2ui.Sprite
	IdleSelectedSprite       *d2ui.Sprite
	ForwardWalkSprite        *d2ui.Sprite
	ForwardWalkSpriteOverlay *d2ui.Sprite
	SelectedSprite           *d2ui.Sprite
	SelectedSpriteOverlay    *d2ui.Sprite
	BackWalkSprite           *d2ui.Sprite
	BackWalkSpriteOverlay    *d2ui.Sprite
	SelectionBounds          image.Rectangle
	SelectSfx                d2interface.SoundEffect
	DeselectSfx              d2interface.SoundEffect
	shc                      *SelectHeroClass
}

func (hri *HeroRenderInfo) advance(elapsed float64) {
	advanceSprite(hri.shc, hri.IdleSprite, elapsed)
	advanceSprite(hri.shc, hri.IdleSelectedSprite, elapsed)
	advanceSprite(hri.shc, hri.ForwardWalkSprite, elapsed)
	advanceSprite(hri.shc, hri.ForwardWalkSpriteOverlay, elapsed)
	advanceSprite(hri.shc, hri.SelectedSprite, elapsed)
	advanceSprite(hri.shc, hri.SelectedSpriteOverlay, elapsed)
	advanceSprite(hri.shc, hri.BackWalkSprite, elapsed)
	advanceSprite(hri.shc, hri.BackWalkSpriteOverlay, elapsed)
}

// CreateSelectHeroClass creates an instance of a SelectHeroClass
func CreateSelectHeroClass(
	navigator d2interface.Navigator,
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	audioProvider d2interface.AudioProvider,
	ui *d2ui.UIManager,
	connectionType d2clientconnectiontype.ClientConnectionType,
	l d2util.LogLevel,
	connectionHost string,
) (*SelectHeroClass, error) {
	playerStateFactory, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	inventoryItemFactory, err := d2inventory.NewInventoryItemFactory(asset)
	if err != nil {
		return nil, err
	}

	selectHeroClass := &SelectHeroClass{
		asset:                asset,
		heroRenderInfo:       make(map[d2enum.Hero]*HeroRenderInfo),
		selectedHero:         d2enum.HeroNone,
		connectionType:       connectionType,
		connectionHost:       connectionHost,
		audioProvider:        audioProvider,
		renderer:             renderer,
		navigator:            navigator,
		uiManager:            ui,
		HeroStateFactory:     playerStateFactory,
		InventoryItemFactory: inventoryItemFactory,
	}

	selectHeroClass.Logger = d2util.NewLogger()
	selectHeroClass.Logger.SetLevel(l)
	selectHeroClass.Logger.SetPrefix(logPrefix)

	return selectHeroClass, nil
}

// SelectHeroClass represents the Select Hero Class screen
type SelectHeroClass struct {
	asset           *d2asset.AssetManager
	uiManager       *d2ui.UIManager
	bgImage         *d2ui.Sprite
	campfire        *d2ui.Sprite
	headingLabel    *d2ui.Label
	heroClassLabel  *d2ui.Label
	heroDesc1Label  *d2ui.Label
	heroDesc2Label  *d2ui.Label
	heroDesc3Label  *d2ui.Label
	heroNameTextbox *d2ui.TextBox
	heroNameLabel   *d2ui.Label
	heroRenderInfo  map[d2enum.Hero]*HeroRenderInfo
	*d2inventory.InventoryItemFactory
	*d2hero.HeroStateFactory
	selectedHero       d2enum.Hero
	exitButton         *d2ui.Button
	okButton           *d2ui.Button
	expansionCheckbox  *d2ui.Checkbox
	expansionCharLabel *d2ui.Label
	hardcoreCheckbox   *d2ui.Checkbox
	hardcoreCharLabel  *d2ui.Label
	connectionType     d2clientconnectiontype.ClientConnectionType
	connectionHost     string

	audioProvider d2interface.AudioProvider
	renderer      d2interface.Renderer
	navigator     d2interface.Navigator

	*d2util.Logger
}

// OnLoad loads the resources for the Select Hero Class screen
func (v *SelectHeroClass) OnLoad(loading d2screen.LoadingState) {
	v.audioProvider.PlayBGM(d2resource.BGMTitle)
	loading.Progress(tenPercent)

	v.bgImage = v.loadSprite(
		d2resource.CharacterSelectBackground,
		point(0, 0),
		0,
		true,
		false,
	)

	loading.Progress(thirtyPercent)

	v.createLabels()
	loading.Progress(fourtyPercent)
	v.createButtons()

	v.campfire = v.loadSprite(
		d2resource.CharacterSelectCampfire,
		point(campfirePosX, campfirePosY),
		0,
		true,
		true,
	)

	v.createCheckboxes()
	loading.Progress(fiftyPercent)

	for hero, config := range getHeroRenderConfiguration() {
		position := config.position
		forwardWalkOverlaySprite := v.loadSprite(
			config.forwardWalkOverlayAnimationPath,
			position,
			config.forwardWalkPlayLengthMs,
			false,
			config.forwardWalkOverlayBlend,
		)
		v.heroRenderInfo[hero] = &HeroRenderInfo{
			Stance: d2enum.HeroStanceIdle,
			IdleSprite: v.loadSprite(config.idleAnimationPath, position,
				config.idlePlayLengthMs, true, false),
			IdleSelectedSprite: v.loadSprite(config.idleSelectedAnimationPath,
				position,
				config.idlePlayLengthMs, true, false),
			ForwardWalkSprite: v.loadSprite(config.forwardWalkAnimationPath, position,
				config.forwardWalkPlayLengthMs, false, false),
			ForwardWalkSpriteOverlay: forwardWalkOverlaySprite,
			SelectedSprite: v.loadSprite(config.selectedAnimationPath, position,
				config.idlePlayLengthMs, true, false),
			SelectedSpriteOverlay: v.loadSprite(config.selectedOverlayAnimationPath, position,
				config.idlePlayLengthMs, true, true),
			BackWalkSprite: v.loadSprite(config.backWalkAnimationPath, position,
				config.backWalkPlayLengthMs, false, false),
			BackWalkSpriteOverlay: v.loadSprite(config.backWalkOverlayAnimationPath, position,
				config.backWalkPlayLengthMs, false, true),
			SelectionBounds: config.selectionBounds,
			SelectSfx:       v.loadSoundEffect(config.selectSfx),
			DeselectSfx:     v.loadSoundEffect(config.deselectSfx),
		}
	}
}

func (v *SelectHeroClass) createLabels() {
	v.headingLabel = v.uiManager.NewLabel(d2resource.Font30, d2resource.PaletteUnits)
	fontWidth, _ := v.headingLabel.GetSize()
	half := 2
	halfFontWidth := fontWidth / half

	v.headingLabel.SetPosition(headingX-halfFontWidth, headingY)
	v.headingLabel.SetText(v.asset.TranslateLabel(d2enum.SelectHeroClassLabel))
	v.headingLabel.Alignment = d2ui.HorizontalAlignCenter

	v.heroClassLabel = v.uiManager.NewLabel(d2resource.Font30, d2resource.PaletteUnits)
	v.heroClassLabel.Alignment = d2ui.HorizontalAlignCenter
	v.heroClassLabel.SetPosition(heroClassLabelX, heroClassLabelY)

	v.heroDesc1Label = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroDesc1Label.Alignment = d2ui.HorizontalAlignCenter
	v.heroDesc1Label.SetPosition(heroDescLine1X, heroDescLine1Y)

	v.heroDesc2Label = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroDesc2Label.Alignment = d2ui.HorizontalAlignCenter
	v.heroDesc2Label.SetPosition(heroDescLine2X, heroDescLine2Y)

	v.heroDesc3Label = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroDesc3Label.Alignment = d2ui.HorizontalAlignCenter
	v.heroDesc3Label.SetPosition(heroDescLine3X, heroDescLine3Y)

	v.heroNameLabel = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroNameLabel.Alignment = d2ui.HorizontalAlignLeft
	v.heroNameLabel.SetText(d2ui.ColorTokenize(v.asset.TranslateLabel(d2enum.CharNameLabel), d2ui.ColorTokenGold))
	v.heroNameLabel.SetPosition(heroNameLabelX, heroNameLabelY)

	v.expansionCharLabel = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.expansionCharLabel.Alignment = d2ui.HorizontalAlignLeft
	v.expansionCharLabel.SetText(d2ui.ColorTokenize(v.asset.TranslateString("#803"), d2ui.ColorTokenGold))
	v.expansionCharLabel.SetPosition(expansionLabelX, expansionLabelY)

	v.hardcoreCharLabel = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.hardcoreCharLabel.Alignment = d2ui.HorizontalAlignLeft
	v.hardcoreCharLabel.SetText(d2ui.ColorTokenize(v.asset.TranslateLabel(d2enum.HardCoreLabel), d2ui.ColorTokenGold))
	v.hardcoreCharLabel.SetPosition(hardcoreLabelX, hardcoreLabelY)
}

func (v *SelectHeroClass) createButtons() {
	v.exitButton = v.uiManager.NewButton(d2ui.ButtonTypeMedium, v.asset.TranslateLabel(d2enum.ExitLabel))
	v.exitButton.SetPosition(selHeroExitBtnX, selHeroExitBtnY)
	v.exitButton.OnActivated(func() { v.onExitButtonClicked() })

	v.okButton = v.uiManager.NewButton(d2ui.ButtonTypeMedium, v.asset.TranslateLabel(d2enum.OKLabel))
	v.okButton.SetPosition(selHeroOkBtnX, selHeroOkBtnY)
	v.okButton.OnActivated(func() { v.onOkButtonClicked() })
	v.okButton.SetVisible(false)
	v.okButton.SetEnabled(false)
}

func (v *SelectHeroClass) createCheckboxes() {
	v.heroNameTextbox = v.uiManager.NewTextbox()
	v.heroNameTextbox.SetPosition(heroNameTextBoxX, heoNameTextBoxY)
	v.heroNameTextbox.SetVisible(false)

	v.expansionCheckbox = v.uiManager.NewCheckbox(true)
	v.expansionCheckbox.SetPosition(expandsionCheckboxX, expansionCheckboxY)
	v.expansionCheckbox.SetVisible(false)

	v.hardcoreCheckbox = v.uiManager.NewCheckbox(false)
	v.hardcoreCheckbox.SetPosition(hardcoreCheckoxX, hardcoreCheckboxY)
	v.hardcoreCheckbox.SetVisible(false)
}

// OnUnload releases the resources of the Select Hero Class screen
func (v *SelectHeroClass) OnUnload() error {
	for i := range v.heroRenderInfo {
		v.heroRenderInfo[i].SelectSfx.Stop()
		v.heroRenderInfo[i].DeselectSfx.Stop()
	}

	v.heroRenderInfo = nil

	return nil
}

func (v *SelectHeroClass) onExitButtonClicked() {
	v.navigator.ToCharacterSelect(v.connectionType, v.connectionHost)
}

func (v *SelectHeroClass) onOkButtonClicked() {
	heroName := v.heroNameTextbox.GetText()
	defaultStats := v.asset.Records.Character.Stats[v.selectedHero]
	statsState := v.CreateHeroStatsState(v.selectedHero, defaultStats)

	playerState, err := v.CreateHeroState(heroName, v.selectedHero, statsState)
	if err != nil {
		v.Errorf("failed to create hero state!, err: %v", err.Error())
		return
	}

	err = v.Save(playerState)
	if err != nil {
		v.Errorf("failed to save game state!, err: %v", err.Error())
		return
	}

	playerState.Equipment = v.InventoryItemFactory.DefaultHeroItems[v.selectedHero]
	v.navigator.ToCreateGame(playerState.FilePath, d2clientconnectiontype.Local, v.connectionHost)
}

// Render renders the Select Hero Class screen
func (v *SelectHeroClass) Render(screen d2interface.Surface) {
	v.bgImage.RenderSegmented(screen, 4, 3, 0)
	v.headingLabel.Render(screen)

	for heroClass, heroInfo := range v.heroRenderInfo {
		if heroInfo.Stance == d2enum.HeroStanceIdle || heroInfo.Stance == d2enum.HeroStanceIdleSelected {
			v.renderHero(screen, heroClass)
		}
	}

	for heroClass, heroInfo := range v.heroRenderInfo {
		if heroInfo.Stance != d2enum.HeroStanceIdle && heroInfo.Stance != d2enum.HeroStanceIdleSelected {
			v.renderHero(screen, heroClass)
		}
	}

	if v.selectedHero != d2enum.HeroNone {
		v.heroClassLabel.Render(screen)
		v.heroDesc1Label.Render(screen)
		v.heroDesc2Label.Render(screen)
		v.heroDesc3Label.Render(screen)
	}

	v.campfire.Render(screen)

	if v.heroNameTextbox.GetVisible() {
		v.heroNameLabel.Render(screen)
		v.expansionCharLabel.Render(screen)
		v.hardcoreCharLabel.Render(screen)
	}
}

// Advance runs the update logic on the Select Hero Class screen
func (v *SelectHeroClass) Advance(tickTime float64) error {
	canSelect := true

	if err := v.campfire.Advance(tickTime); err != nil {
		return err
	}

	for infoIdx := range v.heroRenderInfo {
		v.heroRenderInfo[infoIdx].advance(tickTime)

		if v.heroRenderInfo[infoIdx].Stance != d2enum.HeroStanceIdle &&
			v.heroRenderInfo[infoIdx].Stance != d2enum.HeroStanceIdleSelected &&
			v.heroRenderInfo[infoIdx].Stance != d2enum.HeroStanceSelected {
			canSelect = false
		}
	}

	for heroType := range v.heroRenderInfo {
		v.updateHeroSelectionHover(heroType, canSelect)
	}

	v.okButton.SetEnabled(len(v.heroNameTextbox.GetText()) >= 2 && v.selectedHero != d2enum.HeroNone)

	return nil
}

func (v *SelectHeroClass) updateHeroSelectionHover(hero d2enum.Hero, canSelect bool) {
	renderInfo := v.heroRenderInfo[hero]
	switch renderInfo.Stance {
	case d2enum.HeroStanceApproaching:
		if renderInfo.ForwardWalkSprite.IsOnLastFrame() {
			renderInfo.Stance = d2enum.HeroStanceSelected
			setSpriteToFirstFrame(renderInfo.SelectedSprite)
			setSpriteToFirstFrame(renderInfo.SelectedSpriteOverlay)
		}

		return
	case d2enum.HeroStanceRetreating:
		if renderInfo.BackWalkSprite.IsOnLastFrame() {
			renderInfo.Stance = d2enum.HeroStanceIdle
			setSpriteToFirstFrame(renderInfo.IdleSprite)
		}

		return
	}

	if !canSelect || renderInfo.Stance == d2enum.HeroStanceSelected {
		return
	}

	mouseX, mouseY := v.uiManager.CursorPosition()
	b := renderInfo.SelectionBounds
	mouseHover := (mouseX >= b.Min.X) && (mouseX <= b.Min.X+b.Max.X) && (mouseY >= b.Min.Y) && (mouseY <= b.Min.Y+b.Max.Y)

	if mouseHover && v.uiManager.CursorButtonPressed(d2ui.CursorButtonLeft) {
		v.handleCursorButtonPress(hero, renderInfo)
		return
	}

	v.setCurrentFrame(mouseHover, renderInfo)

	if v.selectedHero == d2enum.HeroNone && mouseHover {
		v.selectedHero = hero
		v.updateHeroText()
	}
}

func (v *SelectHeroClass) handleCursorButtonPress(hero d2enum.Hero, renderInfo *HeroRenderInfo) {
	v.heroNameTextbox.SetVisible(true)
	v.heroNameTextbox.Activate()
	v.okButton.SetVisible(true)
	v.expansionCheckbox.SetVisible(true)
	v.hardcoreCheckbox.SetVisible(true)

	renderInfo.Stance = d2enum.HeroStanceApproaching
	setSpriteToFirstFrame(renderInfo.ForwardWalkSprite)
	setSpriteToFirstFrame(renderInfo.ForwardWalkSpriteOverlay)

	for _, heroInfo := range v.heroRenderInfo {
		if heroInfo.Stance != d2enum.HeroStanceSelected {
			continue
		}

		heroInfo.SelectSfx.Stop()
		heroInfo.DeselectSfx.Play()
		heroInfo.Stance = d2enum.HeroStanceRetreating
		setSpriteToFirstFrame(heroInfo.BackWalkSprite)
		setSpriteToFirstFrame(heroInfo.BackWalkSpriteOverlay)
	}

	v.selectedHero = hero
	v.updateHeroText()
	renderInfo.SelectSfx.Play()
}

func (v *SelectHeroClass) setCurrentFrame(mouseHover bool, renderInfo *HeroRenderInfo) {
	if mouseHover && renderInfo.Stance != d2enum.HeroStanceIdleSelected {
		if err := renderInfo.IdleSelectedSprite.SetCurrentFrame(renderInfo.IdleSprite.GetCurrentFrame()); err != nil {
			v.Errorf("could not set current frame to: %d\n", renderInfo.IdleSprite.GetCurrentFrame())
		}

		renderInfo.Stance = d2enum.HeroStanceIdleSelected
	} else if !mouseHover && renderInfo.Stance != d2enum.HeroStanceIdle {
		if err := renderInfo.IdleSprite.SetCurrentFrame(renderInfo.IdleSelectedSprite.GetCurrentFrame()); err != nil {
			v.Errorf("could not set current frame to: %d\n", renderInfo.IdleSelectedSprite.GetCurrentFrame())
		}

		renderInfo.Stance = d2enum.HeroStanceIdle
	}
}

func (v *SelectHeroClass) renderHero(screen d2interface.Surface, hero d2enum.Hero) {
	renderInfo := v.heroRenderInfo[hero]
	switch renderInfo.Stance {
	case d2enum.HeroStanceIdle:
		drawSprite(renderInfo.IdleSprite, screen)
	case d2enum.HeroStanceIdleSelected:
		drawSprite(renderInfo.IdleSelectedSprite, screen)
	case d2enum.HeroStanceApproaching:
		drawSprite(renderInfo.ForwardWalkSprite, screen)
		drawSprite(renderInfo.ForwardWalkSpriteOverlay, screen)
	case d2enum.HeroStanceSelected:
		drawSprite(renderInfo.SelectedSprite, screen)
		drawSprite(renderInfo.SelectedSpriteOverlay, screen)
	case d2enum.HeroStanceRetreating:
		drawSprite(renderInfo.BackWalkSprite, screen)
		drawSprite(renderInfo.BackWalkSpriteOverlay, screen)
	}
}

func (v *SelectHeroClass) updateHeroText() {
	// v.setDescLabels("") really takes a string translation key, but temporarily disabled.
	switch v.selectedHero {
	case d2enum.HeroNone:
		return
	case d2enum.HeroBarbarian:
		v.heroClassLabel.SetText(v.asset.TranslateString("partycharbar"))
		v.setDescLabels(d2enum.BarbarianDescr, "")
	case d2enum.HeroNecromancer:
		v.heroClassLabel.SetText(v.asset.TranslateString("partycharnec"))
		v.setDescLabels(d2enum.NecromancerDescr, "")
	case d2enum.HeroPaladin:
		v.heroClassLabel.SetText(v.asset.TranslateString("partycharpal"))
		v.setDescLabels(d2enum.PaladinDescr, "")
	case d2enum.HeroAssassin:
		v.heroClassLabel.SetText(v.asset.TranslateString("partycharass"))
		v.setDescLabels(0, "#305")
	case d2enum.HeroSorceress:
		v.heroClassLabel.SetText(v.asset.TranslateString("partycharsor"))
		v.setDescLabels(d2enum.SorceressDescr, "")
	case d2enum.HeroAmazon:
		v.heroClassLabel.SetText(v.asset.TranslateString("partycharama"))
		v.setDescLabels(d2enum.AmazonDescr, "")
	case d2enum.HeroDruid:
		v.heroClassLabel.SetText(v.asset.TranslateString("partychardru"))
		// here is a problem with polish language: in polish string table, there are two items with key "#304"
		v.setDescLabels(0, "#304")
	}
}

const (
	oneLine = 1
	twoLine = 2
)

func (v *SelectHeroClass) setDescLabels(descKey int, key string) {
	var heroDesc string

	if key != "" {
		heroDesc = v.asset.TranslateString(key)
	} else {
		heroDesc = v.asset.TranslateLabel(descKey)
	}

	parts := d2util.SplitIntoLinesWithMaxWidth(heroDesc, heroDescCharWidth)

	numLines := len(parts)

	if numLines > oneLine {
		v.heroDesc1Label.SetText(parts[0])
		v.heroDesc2Label.SetText(parts[1])
	} else {
		v.heroDesc1Label.SetText("")
		v.heroDesc2Label.SetText("")
	}

	if numLines > twoLine {
		v.heroDesc3Label.SetText(parts[2])
	} else {
		v.heroDesc3Label.SetText("")
	}
}

func setSpriteToFirstFrame(sprite *d2ui.Sprite) {
	if sprite != nil {
		sprite.Rewind()
	}
}

func drawSprite(sprite *d2ui.Sprite, target d2interface.Surface) {
	if sprite != nil {
		sprite.Render(target)
	}
}

func advanceSprite(v *SelectHeroClass, sprite *d2ui.Sprite, elapsed float64) {
	if sprite != nil {
		if err := sprite.Advance(elapsed); err != nil {
			v.Error("could not advance the sprite:" + err.Error())
		}
	}
}

func (v *SelectHeroClass) loadSprite(animationPath string, position image.Point,
	playLength int,
	playLoop,
	blend bool) *d2ui.Sprite {
	if animationPath == "" {
		return nil
	}

	sprite, err := v.uiManager.NewSprite(animationPath, d2resource.PaletteFechar)
	if err != nil {
		v.Error("could not load sprite for the animation: %s\n" + animationPath + "with error: " + err.Error())
		return nil
	}

	sprite.PlayForward()
	sprite.SetPlayLoop(playLoop)

	if blend {
		sprite.SetEffect(d2enum.DrawEffectModulate)
	}

	if playLength != 0 {
		sprite.SetPlayLength(float64(playLength) / millisecondsPerSecond)
	}

	sprite.SetPosition(position.X, position.Y)

	return sprite
}

func (v *SelectHeroClass) loadSoundEffect(sfx string) d2interface.SoundEffect {
	result, err := v.audioProvider.LoadSound(sfx, false, false)
	if err != nil {
		v.Error(err.Error())
		return nil
	}

	return result
}
