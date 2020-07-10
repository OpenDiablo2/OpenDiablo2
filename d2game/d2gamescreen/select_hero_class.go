package d2gamescreen

import (
	"fmt"
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2script"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
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

func getHeroRenderConfiguration() map[d2enum.Hero]*heroRenderConfig {
	return map[d2enum.Hero]*heroRenderConfig{
		d2enum.HeroBarbarian: createHeroRenderConfig(
			d2resource.CharacterSelectBarbarianUnselected, d2resource.CharacterSelectBarbarianUnselectedH,
			d2resource.CharacterSelectBarbarianForwardWalk, d2resource.CharacterSelectBarbarianForwardWalkOverlay,
			false, d2resource.CharacterSelectBarbarianSelected, "",
			d2resource.CharacterSelectBarbarianBackWalk, "",
			image.Rectangle{Min: image.Point{X: 364, Y: 201}, Max: image.Point{X: 90, Y: 170}},
			d2resource.SFXBarbarianSelect, d2resource.SFXBarbarianDeselect, image.Point{X: 400, Y: 330},
			0, 2500, 1000,
		),
		d2enum.HeroSorceress: createHeroRenderConfig(
			d2resource.CharacterSelectSorceressUnselected, d2resource.CharacterSelectSorceressUnselectedH,
			d2resource.CharacterSelectSorceressForwardWalk, d2resource.CharacterSelectSorceressForwardWalkOverlay,
			true, d2resource.CharacterSelectSorceressSelected, d2resource.CharacterSelectSorceressSelectedOverlay,
			d2resource.CharacterSelectSorceressBackWalk, d2resource.CharacterSelectSorceressBackWalkOverlay,
			image.Rectangle{Min: image.Point{X: 580, Y: 240}, Max: image.Point{X: 65, Y: 160}},
			d2resource.SFXSorceressSelect, d2resource.SFXSorceressDeselect, image.Point{X: 626, Y: 352},
			2500, 2300, 1200,
		),
		d2enum.HeroNecromancer: createHeroRenderConfig(
			d2resource.CharacterSelectNecromancerUnselected, d2resource.CharacterSelectNecromancerUnselectedH,
			d2resource.CharacterSelectNecromancerForwardWalk, d2resource.CharacterSelectNecromancerForwardWalkOverlay,
			true, d2resource.CharacterSelectNecromancerSelected, d2resource.CharacterSelectNecromancerSelectedOverlay,
			d2resource.CharacterSelectNecromancerBackWalk, d2resource.CharacterSelectNecromancerBackWalkOverlay,
			image.Rectangle{Min: image.Point{X: 265, Y: 220}, Max: image.Point{X: 55, Y: 175}},
			d2resource.SFXNecromancerSelect, d2resource.SFXNecromancerDeselect, image.Point{X: 300, Y: 335},
			1200, 2000, 1500,
		),
		d2enum.HeroPaladin: createHeroRenderConfig(
			d2resource.CharacterSelectPaladinUnselected, d2resource.CharacterSelectPaladinUnselectedH,
			d2resource.CharacterSelectPaladinForwardWalk, d2resource.CharacterSelectPaladinForwardWalkOverlay,
			false, d2resource.CharacterSelectPaladinSelected, "",
			d2resource.CharacterSelectPaladinBackWalk, "",
			image.Rectangle{Min: image.Point{X: 490, Y: 210}, Max: image.Point{X: 65, Y: 180}},
			d2resource.SFXPaladinSelect, d2resource.SFXPaladinDeselect, image.Point{X: 521, Y: 338},
			2500, 3400, 1300,
		),
		d2enum.HeroAmazon: createHeroRenderConfig(
			d2resource.CharacterSelectAmazonUnselected, d2resource.CharacterSelectAmazonUnselectedH,
			d2resource.CharacterSelectAmazonForwardWalk, "",
			false, d2resource.CharacterSelectAmazonSelected, "",
			d2resource.CharacterSelectAmazonBackWalk, "",
			image.Rectangle{Min: image.Point{X: 70, Y: 220}, Max: image.Point{X: 55, Y: 200}},
			d2resource.SFXAmazonSelect, d2resource.SFXAmazonDeselect, image.Point{X: 100, Y: 339},
			2500, 2200, 1500,
		),
		d2enum.HeroAssassin: createHeroRenderConfig(
			d2resource.CharacterSelectAssassinUnselected, d2resource.CharacterSelectAssassinUnselectedH,
			d2resource.CharacterSelectAssassinForwardWalk, "",
			false, d2resource.CharacterSelectAssassinSelected, "",
			d2resource.CharacterSelectAssassinBackWalk, "",
			image.Rectangle{Min: image.Point{X: 175, Y: 235}, Max: image.Point{X: 50, Y: 180}},
			d2resource.SFXAssassinSelect, d2resource.SFXAssassinDeselect, image.Point{X: 231, Y: 365},
			2500, 3800, 1500,
		),
		d2enum.HeroDruid: createHeroRenderConfig(
			d2resource.CharacterSelectDruidUnselected, d2resource.CharacterSelectDruidUnselectedH,
			d2resource.CharacterSelectDruidForwardWalk, "",
			false, d2resource.CharacterSelectDruidSelected, "",
			d2resource.CharacterSelectDruidBackWalk, "",
			image.Rectangle{Min: image.Point{X: 680, Y: 220}, Max: image.Point{X: 70, Y: 195}},
			d2resource.SFXDruidSelect, d2resource.SFXDruidDeselect, image.Point{X: 720, Y: 370},
			1500, 4800, 1500,
		),
	}
}

func createHeroRenderConfig(idleAnimationPath, idleSelectedAnimationPath, forwardWalkAnimationPath,
	forwardWalkOverlayAnimationPath string, forwardWalkOverlayBlend bool, selectedAnimationPath,
	selectedOverlayAnimationPath, backWalkAnimationPath, backWalkOverlayAnimationPath string,
	selectionBounds image.Rectangle, selectSfx, deselectSfx string, position image.Point,
	idlePlayLengthMs, forwardWalkPlayLengthMs, backWalkPlayLengthMs int,
) *heroRenderConfig {
	return &heroRenderConfig{
		idleAnimationPath:               idleAnimationPath,
		idleSelectedAnimationPath:       idleSelectedAnimationPath,
		forwardWalkAnimationPath:        forwardWalkAnimationPath,
		forwardWalkOverlayAnimationPath: forwardWalkOverlayAnimationPath,
		forwardWalkOverlayBlend:         forwardWalkOverlayBlend,
		selectedAnimationPath:           selectedAnimationPath,
		selectedOverlayAnimationPath:    selectedOverlayAnimationPath,
		backWalkAnimationPath:           backWalkAnimationPath,
		backWalkOverlayAnimationPath:    backWalkOverlayAnimationPath,
		selectionBounds:                 selectionBounds,
		selectSfx:                       selectSfx,
		deselectSfx:                     deselectSfx,
		position:                        position,
		idlePlayLengthMs:                idlePlayLengthMs,
		forwardWalkPlayLengthMs:         forwardWalkPlayLengthMs,
		backWalkPlayLengthMs:            backWalkPlayLengthMs,
	}
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
}

func (hri *HeroRenderInfo) advance(elapsed float64) {
	advanceSprite(hri.IdleSprite, elapsed)
	advanceSprite(hri.IdleSelectedSprite, elapsed)
	advanceSprite(hri.ForwardWalkSprite, elapsed)
	advanceSprite(hri.ForwardWalkSpriteOverlay, elapsed)
	advanceSprite(hri.SelectedSprite, elapsed)
	advanceSprite(hri.SelectedSpriteOverlay, elapsed)
	advanceSprite(hri.BackWalkSprite, elapsed)
	advanceSprite(hri.BackWalkSpriteOverlay, elapsed)
}

// SelectHeroClass represents the Select Hero Class screen
type SelectHeroClass struct {
	bgImage            *d2ui.Sprite
	campfire           *d2ui.Sprite
	headingLabel       d2ui.Label
	heroClassLabel     d2ui.Label
	heroDesc1Label     d2ui.Label
	heroDesc2Label     d2ui.Label
	heroDesc3Label     d2ui.Label
	heroNameTextbox    d2ui.TextBox
	heroNameLabel      d2ui.Label
	heroRenderInfo     map[d2enum.Hero]*HeroRenderInfo
	selectedHero       d2enum.Hero
	exitButton         d2ui.Button
	okButton           d2ui.Button
	expansionCheckbox  d2ui.Checkbox
	expansionCharLabel d2ui.Label
	hardcoreCheckbox   d2ui.Checkbox
	hardcoreCharLabel  d2ui.Label
	connectionType     d2clientconnectiontype.ClientConnectionType
	connectionHost     string
	audioProvider      d2interface.AudioProvider
	terminal           d2interface.Terminal
	renderer           d2interface.Renderer
	scriptEngine       *d2script.ScriptEngine
}

// CreateSelectHeroClass creates an instance of a SelectHeroClass
func CreateSelectHeroClass(
	renderer d2interface.Renderer,
	audioProvider d2interface.AudioProvider,
	connectionType d2clientconnectiontype.ClientConnectionType,
	connectionHost string,
	terminal d2interface.Terminal,
	scriptEngine *d2script.ScriptEngine,
) *SelectHeroClass {
	result := &SelectHeroClass{
		heroRenderInfo: make(map[d2enum.Hero]*HeroRenderInfo),
		selectedHero:   d2enum.HeroNone,
		connectionType: connectionType,
		connectionHost: connectionHost,
		audioProvider:  audioProvider,
		terminal:       terminal,
		renderer:       renderer,
		scriptEngine:   scriptEngine,
	}

	return result
}

// OnLoad loads the resources for the Select Hero Class screen
func (v *SelectHeroClass) OnLoad(loading d2screen.LoadingState) {
	v.audioProvider.PlayBGM(d2resource.BGMTitle)
	loading.Progress(0.1)

	v.bgImage = loadSprite(d2resource.CharacterSelectBackground, image.Point{X: 0, Y: 0}, 0, true, false)

	loading.Progress(0.3)

	v.createLabels()
	loading.Progress(0.4)
	v.createButtons()

	v.campfire = loadSprite(d2resource.CharacterSelectCampfire, image.Point{X: 380, Y: 335}, 0, true, true)

	v.createCheckboxes(v.renderer)
	loading.Progress(0.5)

	for hero, config := range getHeroRenderConfiguration() {
		position := config.position
		forwardWalkOverlaySprite := loadSprite(
			config.forwardWalkOverlayAnimationPath,
			position,
			config.forwardWalkPlayLengthMs,
			false,
			config.forwardWalkOverlayBlend,
		)
		v.heroRenderInfo[hero] = &HeroRenderInfo{
			Stance:                   d2enum.HeroStanceIdle,
			IdleSprite:               loadSprite(config.idleAnimationPath, position, config.idlePlayLengthMs, true, false),
			IdleSelectedSprite:       loadSprite(config.idleSelectedAnimationPath, position, config.idlePlayLengthMs, true, false),
			ForwardWalkSprite:        loadSprite(config.forwardWalkAnimationPath, position, config.forwardWalkPlayLengthMs, false, false),
			ForwardWalkSpriteOverlay: forwardWalkOverlaySprite,
			SelectedSprite:           loadSprite(config.selectedAnimationPath, position, config.idlePlayLengthMs, true, false),
			SelectedSpriteOverlay:    loadSprite(config.selectedOverlayAnimationPath, position, config.idlePlayLengthMs, true, true),
			BackWalkSprite:           loadSprite(config.backWalkAnimationPath, position, config.backWalkPlayLengthMs, false, false),
			BackWalkSpriteOverlay:    loadSprite(config.backWalkOverlayAnimationPath, position, config.backWalkPlayLengthMs, false, true),
			SelectionBounds:          config.selectionBounds,
			SelectSfx:                v.loadSoundEffect(config.selectSfx),
			DeselectSfx:              v.loadSoundEffect(config.deselectSfx),
		}
	}
}

func (v *SelectHeroClass) createLabels() {
	v.headingLabel = d2ui.CreateLabel(d2resource.Font30, d2resource.PaletteUnits)
	fontWidth, _ := v.headingLabel.GetSize()
	v.headingLabel.SetPosition(400-fontWidth/2, 17)
	v.headingLabel.SetText("Select Hero Class")
	v.headingLabel.Alignment = d2gui.HorizontalAlignCenter

	v.heroClassLabel = d2ui.CreateLabel(d2resource.Font30, d2resource.PaletteUnits)
	v.heroClassLabel.Alignment = d2gui.HorizontalAlignCenter
	v.heroClassLabel.SetPosition(400, 65)

	v.heroDesc1Label = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroDesc1Label.Alignment = d2gui.HorizontalAlignCenter
	v.heroDesc1Label.SetPosition(400, 100)

	v.heroDesc2Label = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroDesc2Label.Alignment = d2gui.HorizontalAlignCenter
	v.heroDesc2Label.SetPosition(400, 115)

	v.heroDesc3Label = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroDesc3Label.Alignment = d2gui.HorizontalAlignCenter
	v.heroDesc3Label.SetPosition(400, 130)

	v.heroNameLabel = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroNameLabel.Alignment = d2gui.HorizontalAlignLeft
	v.heroNameLabel.Color = color.RGBA{R: 216, G: 196, B: 128, A: 255}
	v.heroNameLabel.SetText("Character Name")
	v.heroNameLabel.SetPosition(321, 475)

	v.expansionCharLabel = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.expansionCharLabel.Alignment = d2gui.HorizontalAlignLeft
	v.expansionCharLabel.Color = color.RGBA{R: 216, G: 196, B: 128, A: 255}
	v.expansionCharLabel.SetText("EXPANSION CHARACTER")
	v.expansionCharLabel.SetPosition(339, 526)

	v.hardcoreCharLabel = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.hardcoreCharLabel.Alignment = d2gui.HorizontalAlignLeft
	v.hardcoreCharLabel.Color = color.RGBA{R: 216, G: 196, B: 128, A: 255}
	v.hardcoreCharLabel.SetText("Hardcore")
	v.hardcoreCharLabel.SetPosition(339, 548)
}

func (v *SelectHeroClass) createButtons() {
	v.exitButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeMedium, "EXIT")
	v.exitButton.SetPosition(33, 537)
	v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
	d2ui.AddWidget(&v.exitButton)

	v.okButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeMedium, "OK")
	v.okButton.SetPosition(630, 537)
	v.okButton.OnActivated(func() { v.onOkButtonClicked() })
	v.okButton.SetVisible(false)
	v.okButton.SetEnabled(false)
	d2ui.AddWidget(&v.okButton)
}

func (v *SelectHeroClass) createCheckboxes(renderer d2interface.Renderer) {
	v.heroNameTextbox = d2ui.CreateTextbox(renderer)
	v.heroNameTextbox.SetPosition(318, 493)
	v.heroNameTextbox.SetVisible(false)
	d2ui.AddWidget(&v.heroNameTextbox)

	v.expansionCheckbox = d2ui.CreateCheckbox(v.renderer, true)
	v.expansionCheckbox.SetPosition(318, 526)
	v.expansionCheckbox.SetVisible(false)
	d2ui.AddWidget(&v.expansionCheckbox)

	v.hardcoreCheckbox = d2ui.CreateCheckbox(renderer, false)
	v.hardcoreCheckbox.SetPosition(318, 548)
	v.hardcoreCheckbox.SetVisible(false)
	d2ui.AddWidget(&v.hardcoreCheckbox)
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
	d2screen.SetNextScreen(CreateCharacterSelect(v.renderer, v.audioProvider, v.connectionType,
		v.connectionHost, v.terminal, v.scriptEngine))
}

func (v *SelectHeroClass) onOkButtonClicked() {
	gameState := d2player.CreatePlayerState(
		v.heroNameTextbox.GetText(),
		v.selectedHero,
		d2datadict.CharStats[v.selectedHero],
		v.hardcoreCheckbox.GetCheckState(),
	)
	gameClient, _ := d2client.Create(d2clientconnectiontype.Local, v.scriptEngine)

	if err := gameClient.Open(v.connectionHost, gameState.FilePath); err != nil {
		fmt.Printf("can not connect to the host: %s\n", v.connectionHost)
	}

	d2screen.SetNextScreen(CreateGame(v.renderer, v.audioProvider, gameClient, v.terminal, v.scriptEngine))
}

// Render renders the Select Hero Class screen
func (v *SelectHeroClass) Render(screen d2interface.Surface) error {
	if err := v.bgImage.RenderSegmented(screen, 4, 3, 0); err != nil {
		return err
	}

	v.headingLabel.Render(screen)

	if v.selectedHero != d2enum.HeroNone {
		v.heroClassLabel.Render(screen)
		v.heroDesc1Label.Render(screen)
		v.heroDesc2Label.Render(screen)
		v.heroDesc3Label.Render(screen)
	}

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

	if err := v.campfire.Render(screen); err != nil {
		return err
	}

	if v.heroNameTextbox.GetVisible() {
		v.heroNameLabel.Render(screen)
		v.expansionCharLabel.Render(screen)
		v.hardcoreCharLabel.Render(screen)
	}

	return nil
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

	mouseX, mouseY := d2ui.CursorPosition()
	b := renderInfo.SelectionBounds
	mouseHover := (mouseX >= b.Min.X) && (mouseX <= b.Min.X+b.Max.X) && (mouseY >= b.Min.Y) && (mouseY <= b.Min.Y+b.Max.Y)

	if mouseHover && d2ui.CursorButtonPressed(d2ui.CursorButtonLeft) {
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
			fmt.Printf("could not set current frame to: %d\n", renderInfo.IdleSprite.GetCurrentFrame())
		}

		renderInfo.Stance = d2enum.HeroStanceIdleSelected
	} else if !mouseHover && renderInfo.Stance != d2enum.HeroStanceIdle {
		if err := renderInfo.IdleSprite.SetCurrentFrame(renderInfo.IdleSelectedSprite.GetCurrentFrame()); err != nil {
			fmt.Printf("could not set current frame to: %d\n", renderInfo.IdleSelectedSprite.GetCurrentFrame())
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
		v.heroClassLabel.SetText(d2common.TranslateString("partycharbar"))
		v.setDescLabels("He is unequaled in close-quarters combat and mastery of weapons.")
	case d2enum.HeroNecromancer:
		v.heroClassLabel.SetText(d2common.TranslateString("partycharnec"))
		v.setDescLabels("Summoning undead minions and cursing his enemies are his specialties.")
	case d2enum.HeroPaladin:
		v.heroClassLabel.SetText(d2common.TranslateString("partycharpal"))
		v.setDescLabels("He is a natural party leader, holy man, and blessed warrior.")
	case d2enum.HeroAssassin:
		v.heroClassLabel.SetText(d2common.TranslateString("partycharass"))
		v.setDescLabels("Schooled in the Martial Arts, her mind and body are deadly weapons.")
	case d2enum.HeroSorceress:
		v.heroClassLabel.SetText(d2common.TranslateString("partycharsor"))
		v.setDescLabels("She has mastered the elemental magicks -- fire, lightning, and ice.")
	case d2enum.HeroAmazon:
		v.heroClassLabel.SetText(d2common.TranslateString("partycharama"))
		v.setDescLabels("Skilled with the spear and the bow, she is a very versatile fighter.")
	case d2enum.HeroDruid:
		v.heroClassLabel.SetText(d2common.TranslateString("partychardru"))
		v.setDescLabels("Commanding the forces of nature, he summons wild beasts and raging storms to his side.")
	}
}

func (v *SelectHeroClass) setDescLabels(descKey string) {
	heroDesc := d2common.TranslateString(descKey)
	parts := d2common.SplitIntoLinesWithMaxWidth(heroDesc, 37)

	if len(parts) > 1 {
		v.heroDesc1Label.SetText(parts[0])
		v.heroDesc2Label.SetText(parts[1])
	} else {
		v.heroDesc1Label.SetText("")
		v.heroDesc2Label.SetText("")
	}

	if len(parts) > 2 {
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
		if err := sprite.Render(target); err != nil {
			x, y := sprite.GetPosition()
			fmt.Printf("could not render the sprite to the position(x: %d, y: %d)\n", x, y)
		}
	}
}

func advanceSprite(sprite *d2ui.Sprite, elapsed float64) {
	if sprite != nil {
		if err := sprite.Advance(elapsed); err != nil {
			fmt.Printf("could not advance the sprite\n")
		}
	}
}

func loadSprite(animationPath string, position image.Point, playLength int, playLoop,
	blend bool) *d2ui.Sprite {
	if animationPath == "" {
		return nil
	}

	animation, err := d2asset.LoadAnimation(animationPath, d2resource.PaletteFechar)
	if err != nil {
		fmt.Printf("could not load animation: %s\n", animationPath)
		return nil
	}

	animation.PlayForward()
	animation.SetPlayLoop(playLoop)

	if blend {
		animation.SetEffect(d2enum.DrawEffectModulate)
	}

	if playLength != 0 {
		animation.SetPlayLengthMs(playLength)
	}

	sprite, err := d2ui.LoadSprite(animation)
	if err != nil {
		fmt.Printf("could not load sprite for the animation: %s\n", animationPath)
		return nil
	}

	sprite.SetPosition(position.X, position.Y)

	return sprite
}

func (v *SelectHeroClass) loadSoundEffect(sfx string) d2interface.SoundEffect {
	result, _ := v.audioProvider.LoadSoundEffect(sfx)
	return result
}
