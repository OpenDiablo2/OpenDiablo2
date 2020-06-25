package d2gamescreen

import (
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

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
	SelectSfx                d2audio.SoundEffect
	DeselectSfx              d2audio.SoundEffect
}

func (hri *HeroRenderInfo) Advance(elapsed float64) {
	advanceSprite(hri.IdleSprite, elapsed)
	advanceSprite(hri.IdleSelectedSprite, elapsed)
	advanceSprite(hri.ForwardWalkSprite, elapsed)
	advanceSprite(hri.ForwardWalkSpriteOverlay, elapsed)
	advanceSprite(hri.SelectedSprite, elapsed)
	advanceSprite(hri.SelectedSpriteOverlay, elapsed)
	advanceSprite(hri.BackWalkSprite, elapsed)
	advanceSprite(hri.BackWalkSpriteOverlay, elapsed)
}

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
}

func CreateSelectHeroClass(connectionType d2clientconnectiontype.ClientConnectionType, connectionHost string) *SelectHeroClass {
	result := &SelectHeroClass{
		heroRenderInfo: make(map[d2enum.Hero]*HeroRenderInfo),
		selectedHero:   d2enum.HeroNone,
		connectionType: connectionType,
		connectionHost: connectionHost,
	}
	return result
}

func (v *SelectHeroClass) OnLoad(loading d2screen.LoadingState) {
	d2audio.PlayBGM(d2resource.BGMTitle)
	loading.Progress(0.1)

	v.bgImage = loadSprite(d2resource.CharacterSelectBackground, d2resource.PaletteFechar)
	v.bgImage.SetPosition(0, 0)

	v.headingLabel = d2ui.CreateLabel(d2resource.Font30, d2resource.PaletteUnits)
	fontWidth, _ := v.headingLabel.GetSize()
	v.headingLabel.SetPosition(400-fontWidth/2, 17)
	v.headingLabel.SetText("Select Hero Class")
	v.headingLabel.Alignment = d2ui.LabelAlignCenter

	v.heroClassLabel = d2ui.CreateLabel(d2resource.Font30, d2resource.PaletteUnits)
	v.heroClassLabel.Alignment = d2ui.LabelAlignCenter
	v.heroClassLabel.SetPosition(400, 65)

	v.heroDesc1Label = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroDesc1Label.Alignment = d2ui.LabelAlignCenter
	v.heroDesc1Label.SetPosition(400, 100)
	loading.Progress(0.3)

	v.heroDesc2Label = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroDesc2Label.Alignment = d2ui.LabelAlignCenter
	v.heroDesc2Label.SetPosition(400, 115)

	v.heroDesc3Label = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroDesc3Label.Alignment = d2ui.LabelAlignCenter
	v.heroDesc3Label.SetPosition(400, 130)

	v.campfire = loadSprite(d2resource.CharacterSelectCampfire, d2resource.PaletteFechar)
	v.campfire.SetPosition(380, 335)
	v.campfire.PlayForward()
	v.campfire.SetBlend(true)

	v.exitButton = d2ui.CreateButton(d2ui.ButtonTypeMedium, "EXIT")
	v.exitButton.SetPosition(33, 537)
	v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
	d2ui.AddWidget(&v.exitButton)

	v.okButton = d2ui.CreateButton(d2ui.ButtonTypeMedium, "OK")
	v.okButton.SetPosition(630, 537)
	v.okButton.OnActivated(func() { v.onOkButtonClicked() })
	v.okButton.SetVisible(false)
	v.okButton.SetEnabled(false)
	d2ui.AddWidget(&v.okButton)

	v.heroNameLabel = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.heroNameLabel.Alignment = d2ui.LabelAlignLeft
	v.heroNameLabel.Color = color.RGBA{R: 216, G: 196, B: 128, A: 255}
	v.heroNameLabel.SetText("Character Name")
	v.heroNameLabel.SetPosition(321, 475)
	loading.Progress(0.4)

	v.heroNameTextbox = d2ui.CreateTextbox()
	v.heroNameTextbox.SetPosition(318, 493)
	v.heroNameTextbox.SetVisible(false)
	d2ui.AddWidget(&v.heroNameTextbox)

	v.expansionCheckbox = d2ui.CreateCheckbox(true)
	v.expansionCheckbox.SetPosition(318, 526)
	v.expansionCheckbox.SetVisible(false)
	d2ui.AddWidget(&v.expansionCheckbox)

	v.expansionCharLabel = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.expansionCharLabel.Alignment = d2ui.LabelAlignLeft
	v.expansionCharLabel.Color = color.RGBA{R: 216, G: 196, B: 128, A: 255}
	v.expansionCharLabel.SetText("EXPANSION CHARACTER")
	v.expansionCharLabel.SetPosition(339, 526)

	v.hardcoreCheckbox = d2ui.CreateCheckbox(false)
	v.hardcoreCheckbox.SetPosition(318, 548)
	v.hardcoreCheckbox.SetVisible(false)
	d2ui.AddWidget(&v.hardcoreCheckbox)

	v.hardcoreCharLabel = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.hardcoreCharLabel.Alignment = d2ui.LabelAlignLeft
	v.hardcoreCharLabel.Color = color.RGBA{R: 216, G: 196, B: 128, A: 255}
	v.hardcoreCharLabel.SetText("Hardcore")
	v.hardcoreCharLabel.SetPosition(339, 548)
	loading.Progress(0.5)

	v.heroRenderInfo[d2enum.HeroBarbarian] = &HeroRenderInfo{
		d2enum.HeroStanceIdle,
		loadSprite(d2resource.CharacterSelectBarbarianUnselected, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectBarbarianUnselectedH, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectBarbarianForwardWalk, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectBarbarianForwardWalkOverlay, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectBarbarianSelected, d2resource.PaletteFechar),
		nil,
		loadSprite(d2resource.CharacterSelectBarbarianBackWalk, d2resource.PaletteFechar),
		nil,
		image.Rectangle{Min: image.Point{X: 364, Y: 201}, Max: image.Point{X: 90, Y: 170}},
		loadSoundEffect(d2resource.SFXBarbarianSelect),
		loadSoundEffect(d2resource.SFXBarbarianDeselect),
	}
	v.heroRenderInfo[d2enum.HeroBarbarian].IdleSprite.SetPosition(400, 330)
	v.heroRenderInfo[d2enum.HeroBarbarian].IdleSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroBarbarian].IdleSelectedSprite.SetPosition(400, 330)
	v.heroRenderInfo[d2enum.HeroBarbarian].IdleSelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroBarbarian].ForwardWalkSprite.SetPosition(400, 330)
	v.heroRenderInfo[d2enum.HeroBarbarian].ForwardWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroBarbarian].ForwardWalkSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroBarbarian].ForwardWalkSprite.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroBarbarian].ForwardWalkSpriteOverlay.SetPosition(400, 330)
	v.heroRenderInfo[d2enum.HeroBarbarian].ForwardWalkSpriteOverlay.PlayForward()
	v.heroRenderInfo[d2enum.HeroBarbarian].ForwardWalkSpriteOverlay.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroBarbarian].ForwardWalkSpriteOverlay.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroBarbarian].SelectedSprite.SetPosition(400, 330)
	v.heroRenderInfo[d2enum.HeroBarbarian].SelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroBarbarian].BackWalkSprite.SetPosition(400, 330)
	v.heroRenderInfo[d2enum.HeroBarbarian].BackWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroBarbarian].BackWalkSprite.SetPlayLengthMs(1000)
	v.heroRenderInfo[d2enum.HeroBarbarian].BackWalkSprite.SetPlayLoop(false)

	v.heroRenderInfo[d2enum.HeroSorceress] = &HeroRenderInfo{
		d2enum.HeroStanceIdle,
		loadSprite(d2resource.CharacterSelecSorceressUnselected, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecSorceressUnselectedH, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecSorceressForwardWalk, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecSorceressForwardWalkOverlay, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecSorceressSelected, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecSorceressSelectedOverlay, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecSorceressBackWalk, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecSorceressBackWalkOverlay, d2resource.PaletteFechar),
		image.Rectangle{Min: image.Point{X: 580, Y: 240}, Max: image.Point{X: 65, Y: 160}},
		loadSoundEffect(d2resource.SFXSorceressSelect),
		loadSoundEffect(d2resource.SFXSorceressDeselect),
	}
	v.heroRenderInfo[d2enum.HeroSorceress].IdleSprite.SetPosition(626, 352)
	v.heroRenderInfo[d2enum.HeroSorceress].IdleSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroSorceress].IdleSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroSorceress].IdleSelectedSprite.SetPosition(626, 352)
	v.heroRenderInfo[d2enum.HeroSorceress].IdleSelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroSorceress].IdleSelectedSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroSorceress].ForwardWalkSprite.SetPosition(626, 352)
	v.heroRenderInfo[d2enum.HeroSorceress].ForwardWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroSorceress].ForwardWalkSprite.SetPlayLengthMs(2300)
	v.heroRenderInfo[d2enum.HeroSorceress].ForwardWalkSprite.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroSorceress].ForwardWalkSpriteOverlay.SetBlend(true)
	v.heroRenderInfo[d2enum.HeroSorceress].ForwardWalkSpriteOverlay.SetPosition(626, 352)
	v.heroRenderInfo[d2enum.HeroSorceress].ForwardWalkSpriteOverlay.PlayForward()
	v.heroRenderInfo[d2enum.HeroSorceress].ForwardWalkSpriteOverlay.SetPlayLengthMs(2300)
	v.heroRenderInfo[d2enum.HeroSorceress].ForwardWalkSpriteOverlay.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroSorceress].SelectedSprite.SetPosition(626, 352)
	v.heroRenderInfo[d2enum.HeroSorceress].SelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroSorceress].SelectedSprite.SetPlayLengthMs(450)
	v.heroRenderInfo[d2enum.HeroSorceress].SelectedSpriteOverlay.SetBlend(true)
	v.heroRenderInfo[d2enum.HeroSorceress].SelectedSpriteOverlay.SetPosition(626, 352)
	v.heroRenderInfo[d2enum.HeroSorceress].SelectedSpriteOverlay.PlayForward()
	v.heroRenderInfo[d2enum.HeroSorceress].SelectedSpriteOverlay.SetPlayLengthMs(450)
	v.heroRenderInfo[d2enum.HeroSorceress].BackWalkSprite.SetPosition(626, 352)
	v.heroRenderInfo[d2enum.HeroSorceress].BackWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroSorceress].BackWalkSprite.SetPlayLengthMs(1200)
	v.heroRenderInfo[d2enum.HeroSorceress].BackWalkSprite.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroSorceress].BackWalkSpriteOverlay.SetBlend(true)
	v.heroRenderInfo[d2enum.HeroSorceress].BackWalkSpriteOverlay.SetPosition(626, 352)
	v.heroRenderInfo[d2enum.HeroSorceress].BackWalkSpriteOverlay.PlayForward()
	v.heroRenderInfo[d2enum.HeroSorceress].BackWalkSpriteOverlay.SetPlayLengthMs(1200)
	v.heroRenderInfo[d2enum.HeroSorceress].BackWalkSpriteOverlay.SetPlayLoop(false)
	loading.Progress(0.6)

	v.heroRenderInfo[d2enum.HeroNecromancer] = &HeroRenderInfo{
		d2enum.HeroStanceIdle,
		loadSprite(d2resource.CharacterSelectNecromancerUnselected, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectNecromancerUnselectedH, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecNecromancerForwardWalk, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecNecromancerForwardWalkOverlay, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecNecromancerSelected, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecNecromancerSelectedOverlay, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecNecromancerBackWalk, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecNecromancerBackWalkOverlay, d2resource.PaletteFechar),
		image.Rectangle{Min: image.Point{X: 265, Y: 220}, Max: image.Point{X: 55, Y: 175}},
		loadSoundEffect(d2resource.SFXNecromancerSelect),
		loadSoundEffect(d2resource.SFXNecromancerDeselect),
	}
	v.heroRenderInfo[d2enum.HeroNecromancer].IdleSprite.SetPosition(300, 335)
	v.heroRenderInfo[d2enum.HeroNecromancer].IdleSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroNecromancer].IdleSprite.SetPlayLengthMs(1200)
	v.heroRenderInfo[d2enum.HeroNecromancer].IdleSelectedSprite.SetPosition(300, 335)
	v.heroRenderInfo[d2enum.HeroNecromancer].IdleSelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroNecromancer].IdleSelectedSprite.SetPlayLengthMs(1200)
	v.heroRenderInfo[d2enum.HeroNecromancer].ForwardWalkSprite.SetPosition(300, 335)
	v.heroRenderInfo[d2enum.HeroNecromancer].ForwardWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroNecromancer].ForwardWalkSprite.SetPlayLengthMs(2000)
	v.heroRenderInfo[d2enum.HeroNecromancer].ForwardWalkSprite.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroNecromancer].ForwardWalkSpriteOverlay.SetBlend(true)
	v.heroRenderInfo[d2enum.HeroNecromancer].ForwardWalkSpriteOverlay.SetPosition(300, 335)
	v.heroRenderInfo[d2enum.HeroNecromancer].ForwardWalkSpriteOverlay.PlayForward()
	v.heroRenderInfo[d2enum.HeroNecromancer].ForwardWalkSpriteOverlay.SetPlayLengthMs(2000)
	v.heroRenderInfo[d2enum.HeroNecromancer].ForwardWalkSpriteOverlay.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroNecromancer].SelectedSprite.SetPosition(300, 335)
	v.heroRenderInfo[d2enum.HeroNecromancer].SelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroNecromancer].SelectedSpriteOverlay.SetBlend(true)
	v.heroRenderInfo[d2enum.HeroNecromancer].SelectedSpriteOverlay.SetPosition(300, 335)
	v.heroRenderInfo[d2enum.HeroNecromancer].SelectedSpriteOverlay.PlayForward()
	v.heroRenderInfo[d2enum.HeroNecromancer].BackWalkSprite.SetPosition(300, 335)
	v.heroRenderInfo[d2enum.HeroNecromancer].BackWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroNecromancer].BackWalkSprite.SetPlayLengthMs(1500)
	v.heroRenderInfo[d2enum.HeroNecromancer].BackWalkSprite.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroNecromancer].BackWalkSpriteOverlay.SetBlend(true)
	v.heroRenderInfo[d2enum.HeroNecromancer].BackWalkSpriteOverlay.SetPosition(300, 335)
	v.heroRenderInfo[d2enum.HeroNecromancer].BackWalkSpriteOverlay.PlayForward()
	v.heroRenderInfo[d2enum.HeroNecromancer].BackWalkSpriteOverlay.SetPlayLengthMs(1500)
	v.heroRenderInfo[d2enum.HeroNecromancer].BackWalkSpriteOverlay.SetPlayLoop(false)

	v.heroRenderInfo[d2enum.HeroPaladin] = &HeroRenderInfo{
		d2enum.HeroStanceIdle,
		loadSprite(d2resource.CharacterSelectPaladinUnselected, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectPaladinUnselectedH, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecPaladinForwardWalk, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecPaladinForwardWalkOverlay, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecPaladinSelected, d2resource.PaletteFechar),
		nil,
		loadSprite(d2resource.CharacterSelecPaladinBackWalk, d2resource.PaletteFechar),
		nil,
		image.Rectangle{Min: image.Point{X: 490, Y: 210}, Max: image.Point{X: 65, Y: 180}},
		loadSoundEffect(d2resource.SFXPaladinSelect),
		loadSoundEffect(d2resource.SFXPaladinDeselect),
	}
	v.heroRenderInfo[d2enum.HeroPaladin].IdleSprite.SetPosition(521, 338)
	v.heroRenderInfo[d2enum.HeroPaladin].IdleSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroPaladin].IdleSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroPaladin].IdleSelectedSprite.SetPosition(521, 338)
	v.heroRenderInfo[d2enum.HeroPaladin].IdleSelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroPaladin].IdleSelectedSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroPaladin].ForwardWalkSprite.SetPosition(521, 338)
	v.heroRenderInfo[d2enum.HeroPaladin].ForwardWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroPaladin].ForwardWalkSprite.SetPlayLengthMs(3400)
	v.heroRenderInfo[d2enum.HeroPaladin].ForwardWalkSprite.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroPaladin].ForwardWalkSpriteOverlay.SetPosition(521, 338)
	v.heroRenderInfo[d2enum.HeroPaladin].ForwardWalkSpriteOverlay.PlayForward()
	v.heroRenderInfo[d2enum.HeroPaladin].ForwardWalkSpriteOverlay.SetPlayLengthMs(3400)
	v.heroRenderInfo[d2enum.HeroPaladin].ForwardWalkSpriteOverlay.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroPaladin].SelectedSprite.SetPosition(521, 338)
	v.heroRenderInfo[d2enum.HeroPaladin].SelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroPaladin].SelectedSprite.SetPlayLengthMs(650)
	v.heroRenderInfo[d2enum.HeroPaladin].BackWalkSprite.SetPosition(521, 338)
	v.heroRenderInfo[d2enum.HeroPaladin].BackWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroPaladin].BackWalkSprite.SetPlayLengthMs(1300)
	v.heroRenderInfo[d2enum.HeroPaladin].BackWalkSprite.SetPlayLoop(false)
	loading.Progress(0.7)

	v.heroRenderInfo[d2enum.HeroAmazon] = &HeroRenderInfo{
		d2enum.HeroStanceIdle,
		loadSprite(d2resource.CharacterSelectAmazonUnselected, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectAmazonUnselectedH, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelecAmazonForwardWalk, d2resource.PaletteFechar),
		nil,
		loadSprite(d2resource.CharacterSelecAmazonSelected, d2resource.PaletteFechar),
		nil,
		loadSprite(d2resource.CharacterSelecAmazonBackWalk, d2resource.PaletteFechar),
		nil,
		image.Rectangle{Min: image.Point{X: 70, Y: 220}, Max: image.Point{X: 55, Y: 200}},
		loadSoundEffect(d2resource.SFXAmazonSelect),
		loadSoundEffect(d2resource.SFXAmazonDeselect),
	}
	v.heroRenderInfo[d2enum.HeroAmazon].IdleSprite.SetPosition(100, 339)
	v.heroRenderInfo[d2enum.HeroAmazon].IdleSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAmazon].IdleSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroAmazon].IdleSelectedSprite.SetPosition(100, 339)
	v.heroRenderInfo[d2enum.HeroAmazon].IdleSelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAmazon].IdleSelectedSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroAmazon].ForwardWalkSprite.SetPosition(100, 339)
	v.heroRenderInfo[d2enum.HeroAmazon].ForwardWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAmazon].ForwardWalkSprite.SetPlayLengthMs(2200)
	v.heroRenderInfo[d2enum.HeroAmazon].ForwardWalkSprite.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroAmazon].SelectedSprite.SetPosition(100, 339)
	v.heroRenderInfo[d2enum.HeroAmazon].SelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAmazon].SelectedSprite.SetPlayLengthMs(1350)
	v.heroRenderInfo[d2enum.HeroAmazon].BackWalkSprite.SetPosition(100, 339)
	v.heroRenderInfo[d2enum.HeroAmazon].BackWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAmazon].BackWalkSprite.SetPlayLengthMs(1500)
	v.heroRenderInfo[d2enum.HeroAmazon].BackWalkSprite.SetPlayLoop(false)

	v.heroRenderInfo[d2enum.HeroAssassin] = &HeroRenderInfo{
		d2enum.HeroStanceIdle,
		loadSprite(d2resource.CharacterSelectAssassinUnselected, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectAssassinUnselectedH, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectAssassinForwardWalk, d2resource.PaletteFechar),
		nil,
		loadSprite(d2resource.CharacterSelectAssassinSelected, d2resource.PaletteFechar),
		nil,
		loadSprite(d2resource.CharacterSelectAssassinBackWalk, d2resource.PaletteFechar),
		nil,
		image.Rectangle{Min: image.Point{X: 175, Y: 235}, Max: image.Point{X: 50, Y: 180}},
		loadSoundEffect(d2resource.SFXAssassinSelect),
		loadSoundEffect(d2resource.SFXAssassinDeselect),
	}
	v.heroRenderInfo[d2enum.HeroAssassin].IdleSprite.SetPosition(231, 365)
	v.heroRenderInfo[d2enum.HeroAssassin].IdleSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAssassin].IdleSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroAssassin].IdleSelectedSprite.SetPosition(231, 365)
	v.heroRenderInfo[d2enum.HeroAssassin].IdleSelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAssassin].IdleSelectedSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroAssassin].ForwardWalkSprite.SetPosition(231, 365)
	v.heroRenderInfo[d2enum.HeroAssassin].ForwardWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAssassin].ForwardWalkSprite.SetPlayLengthMs(3800)
	v.heroRenderInfo[d2enum.HeroAssassin].ForwardWalkSprite.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroAssassin].SelectedSprite.SetPosition(231, 365)
	v.heroRenderInfo[d2enum.HeroAssassin].SelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAssassin].SelectedSprite.SetPlayLengthMs(2500)
	v.heroRenderInfo[d2enum.HeroAssassin].BackWalkSprite.SetPosition(231, 365)
	v.heroRenderInfo[d2enum.HeroAssassin].BackWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroAssassin].BackWalkSprite.SetPlayLengthMs(1500)
	v.heroRenderInfo[d2enum.HeroAssassin].BackWalkSprite.SetPlayLoop(false)
	loading.Progress(0.8)

	v.heroRenderInfo[d2enum.HeroDruid] = &HeroRenderInfo{
		d2enum.HeroStanceIdle,
		loadSprite(d2resource.CharacterSelectDruidUnselected, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectDruidUnselectedH, d2resource.PaletteFechar),
		loadSprite(d2resource.CharacterSelectDruidForwardWalk, d2resource.PaletteFechar),
		nil,
		loadSprite(d2resource.CharacterSelectDruidSelected, d2resource.PaletteFechar),
		nil,
		loadSprite(d2resource.CharacterSelectDruidBackWalk, d2resource.PaletteFechar),
		nil,
		image.Rectangle{Min: image.Point{X: 680, Y: 220}, Max: image.Point{X: 70, Y: 195}},
		loadSoundEffect(d2resource.SFXDruidSelect),
		loadSoundEffect(d2resource.SFXDruidDeselect),
	}
	v.heroRenderInfo[d2enum.HeroDruid].IdleSprite.SetPosition(720, 370)
	v.heroRenderInfo[d2enum.HeroDruid].IdleSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroDruid].IdleSprite.SetPlayLengthMs(1500)
	v.heroRenderInfo[d2enum.HeroDruid].IdleSelectedSprite.SetPosition(720, 370)
	v.heroRenderInfo[d2enum.HeroDruid].IdleSelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroDruid].IdleSelectedSprite.SetPlayLengthMs(1500)
	v.heroRenderInfo[d2enum.HeroDruid].ForwardWalkSprite.SetPosition(720, 370)
	v.heroRenderInfo[d2enum.HeroDruid].ForwardWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroDruid].ForwardWalkSprite.SetPlayLengthMs(4800)
	v.heroRenderInfo[d2enum.HeroDruid].ForwardWalkSprite.SetPlayLoop(false)
	v.heroRenderInfo[d2enum.HeroDruid].SelectedSprite.SetPosition(720, 370)
	v.heroRenderInfo[d2enum.HeroDruid].SelectedSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroDruid].SelectedSprite.SetPlayLengthMs(1500)
	v.heroRenderInfo[d2enum.HeroDruid].BackWalkSprite.SetPosition(720, 370)
	v.heroRenderInfo[d2enum.HeroDruid].BackWalkSprite.PlayForward()
	v.heroRenderInfo[d2enum.HeroDruid].BackWalkSprite.SetPlayLengthMs(1500)
	v.heroRenderInfo[d2enum.HeroDruid].BackWalkSprite.SetPlayLoop(false)
}

func (v *SelectHeroClass) OnUnload() error {
	for i := range v.heroRenderInfo {
		v.heroRenderInfo[i].SelectSfx.Stop()
		v.heroRenderInfo[i].DeselectSfx.Stop()
	}
	v.heroRenderInfo = nil
	return nil
}

func (v SelectHeroClass) onExitButtonClicked() {
	d2screen.SetNextScreen(CreateCharacterSelect(v.connectionType, v.connectionHost))
}

func (v SelectHeroClass) onOkButtonClicked() {
	gameState := d2player.CreatePlayerState(v.heroNameTextbox.GetText(), v.selectedHero, *d2datadict.CharStats[v.selectedHero], v.hardcoreCheckbox.GetCheckState())
	gameClient, _ := d2client.Create(d2clientconnectiontype.Local)
	gameClient.Open(v.connectionHost, gameState.FilePath)
	d2screen.SetNextScreen(CreateGame(gameClient))
}

func (v *SelectHeroClass) Render(screen d2render.Surface) error {
	v.bgImage.RenderSegmented(screen, 4, 3, 0)
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
	v.campfire.Render(screen)
	if v.heroNameTextbox.GetVisible() {
		v.heroNameLabel.Render(screen)
		v.expansionCharLabel.Render(screen)
		v.hardcoreCharLabel.Render(screen)
	}

	return nil
}

func (v *SelectHeroClass) Advance(tickTime float64) error {
	canSelect := true
	v.campfire.Advance(tickTime)
	for _, info := range v.heroRenderInfo {
		info.Advance(tickTime)
		if info.Stance != d2enum.HeroStanceIdle && info.Stance != d2enum.HeroStanceIdleSelected && info.Stance != d2enum.HeroStanceSelected {
			canSelect = false
		}
	}
	for heroType, _ := range v.heroRenderInfo {
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
	if !canSelect {
		return
	}
	if renderInfo.Stance == d2enum.HeroStanceSelected {
		return
	}
	mouseX, mouseY := d2ui.CursorPosition()
	b := renderInfo.SelectionBounds
	mouseHover := (mouseX >= b.Min.X) && (mouseX <= b.Min.X+b.Max.X) && (mouseY >= b.Min.Y) && (mouseY <= b.Min.Y+b.Max.Y)
	if mouseHover && d2ui.CursorButtonPressed(d2ui.CursorButtonLeft) {
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

		return
	}

	if mouseHover && renderInfo.Stance != d2enum.HeroStanceIdleSelected {
		renderInfo.IdleSelectedSprite.SetCurrentFrame(renderInfo.IdleSprite.GetCurrentFrame())
		renderInfo.Stance = d2enum.HeroStanceIdleSelected
	} else if !mouseHover && renderInfo.Stance != d2enum.HeroStanceIdle {
		renderInfo.IdleSprite.SetCurrentFrame(renderInfo.IdleSelectedSprite.GetCurrentFrame())
		renderInfo.Stance = d2enum.HeroStanceIdle
	}

	if v.selectedHero == d2enum.HeroNone && mouseHover {
		v.selectedHero = hero
		v.updateHeroText()
	}

}

func (v *SelectHeroClass) renderHero(screen d2render.Surface, hero d2enum.Hero) {
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
	/*
	   if (selectedHero == null)
	                   return;

	               switch (selectedHero.Value)
	               {

	               }

	               heroClassLabel.Location = new Point(400 - (heroClassLabel.TextArea.Width / 2), 65);
	               heroDesc1Label.Location = new Point(400 - (heroDesc1Label.TextArea.Width / 2), 100);
	               heroDesc2Label.Location = new Point(400 - (heroDesc2Label.TextArea.Width / 2), 115);
	               heroDesc3Label.Location = new Point(400 - (heroDesc3Label.TextArea.Width / 2), 130);
	*/
}

func (v *SelectHeroClass) setDescLabels(descKey string) {
	heroDesc := d2common.TranslateString(descKey)
	parts := d2common.SplitIntoLinesWithMaxWidth(heroDesc, 37)
	if len(parts) > 1 {
		v.heroDesc1Label.SetText(parts[0])
	} else {
		v.heroDesc1Label.SetText("")
	}
	if len(parts) > 1 {
		v.heroDesc2Label.SetText(parts[1])
	} else {
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

func drawSprite(sprite *d2ui.Sprite, target d2render.Surface) {
	if sprite != nil {
		sprite.Render(target)
	}
}

func advanceSprite(sprite *d2ui.Sprite, elapsed float64) {
	if sprite != nil {
		sprite.Advance(elapsed)
	}
}

func loadSprite(animationPath, palettePath string) *d2ui.Sprite {
	animation, _ := d2asset.LoadAnimation(animationPath, palettePath)
	sprite, _ := d2ui.LoadSprite(animation)
	return sprite
}

func loadSoundEffect(sfx string) d2audio.SoundEffect {
	result, _ := d2audio.LoadSoundEffect(sfx)
	return result
}
