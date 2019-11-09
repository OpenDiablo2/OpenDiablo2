package scenes

import (
	"image"

	"github.com/OpenDiablo2/OpenDiablo2/common"
	"github.com/OpenDiablo2/OpenDiablo2/palettedefs"
	"github.com/OpenDiablo2/OpenDiablo2/resourcepaths"
	"github.com/OpenDiablo2/OpenDiablo2/sound"
	"github.com/OpenDiablo2/OpenDiablo2/ui"
	"github.com/hajimehoshi/ebiten"
)

type HeroStance int

const (
	HeroStanceIdle         HeroStance = 0
	HeroStanceIdleSelected HeroStance = 1
	HeroStanceApproaching  HeroStance = 2
	HeroStanceSelected     HeroStance = 3
	HeroStanceRetreating   HeroStance = 4
)

type HeroRenderInfo struct {
	Stance                   HeroStance
	IdleSprite               *common.Sprite
	IdleSelectedSprite       *common.Sprite
	ForwardWalkSprite        *common.Sprite
	ForwardWalkSpriteOverlay *common.Sprite
	SelectedSprite           *common.Sprite
	SelectedSpriteOverlay    *common.Sprite
	BackWalkSprite           *common.Sprite
	BackWalkSpriteOverlay    *common.Sprite
	SelectionBounds          image.Rectangle
	SelectSfx                *sound.SoundEffect
	DeselectSfx              *sound.SoundEffect
}

type SelectHeroClass struct {
	uiManager      *ui.Manager
	soundManager   *sound.Manager
	fileProvider   common.FileProvider
	sceneProvider  SceneProvider
	bgImage        *common.Sprite
	campfire       *common.Sprite
	headingLabel   *ui.Label
	heroClassLabel *ui.Label
	heroDesc1Label *ui.Label
	heroDesc2Label *ui.Label
	heroDesc3Label *ui.Label
	heroRenderInfo map[common.Hero]*HeroRenderInfo
	selectedHero   common.Hero
	exitButton     *ui.Button
}

func CreateSelectHeroClass(
	fileProvider common.FileProvider,
	sceneProvider SceneProvider,
	uiManager *ui.Manager, soundManager *sound.Manager,
) *SelectHeroClass {
	result := &SelectHeroClass{
		uiManager:      uiManager,
		sceneProvider:  sceneProvider,
		fileProvider:   fileProvider,
		soundManager:   soundManager,
		heroRenderInfo: make(map[common.Hero]*HeroRenderInfo),
		selectedHero:   common.HeroNone,
	}
	return result
}

func (v *SelectHeroClass) Load() []func() {
	v.soundManager.PlayBGM(resourcepaths.BGMTitle)
	return []func(){
		func() {
			v.bgImage = v.fileProvider.LoadSprite(resourcepaths.CharacterSelectBackground, palettedefs.Fechar)
			v.bgImage.MoveTo(0, 0)
		},
		func() {
			v.headingLabel = ui.CreateLabel(v.fileProvider, resourcepaths.Font30, palettedefs.Units)
			fontWidth, _ := v.headingLabel.GetSize()
			v.headingLabel.MoveTo(400-int(fontWidth/2), 17)
			v.headingLabel.SetText("Select Hero Class")
			v.headingLabel.Alignment = ui.LabelAlignCenter
		},
		func() {
			v.heroClassLabel = ui.CreateLabel(v.fileProvider, resourcepaths.Font30, palettedefs.Units)
			v.heroClassLabel.Alignment = ui.LabelAlignCenter
			v.heroClassLabel.MoveTo(400, 65)
		},
		func() {
			v.heroDesc1Label = ui.CreateLabel(v.fileProvider, resourcepaths.Font16, palettedefs.Units)
			v.heroDesc1Label.Alignment = ui.LabelAlignCenter
			v.heroDesc1Label.MoveTo(400, 100)
		},
		func() {
			v.heroDesc2Label = ui.CreateLabel(v.fileProvider, resourcepaths.Font16, palettedefs.Units)
			v.heroDesc2Label.Alignment = ui.LabelAlignCenter
			v.heroDesc2Label.MoveTo(400, 115)
		},
		func() {
			v.heroDesc3Label = ui.CreateLabel(v.fileProvider, resourcepaths.Font16, palettedefs.Units)
			v.heroDesc3Label.Alignment = ui.LabelAlignCenter
			v.heroDesc3Label.MoveTo(400, 130)
		},
		func() {
			v.campfire = v.fileProvider.LoadSprite(resourcepaths.CharacterSelectCampfire, palettedefs.Fechar)
			v.campfire.MoveTo(380, 335)
			v.campfire.Animate = true
			v.campfire.Blend = true
		},
		func() {
			v.exitButton = ui.CreateButton(ui.ButtonTypeMedium, v.fileProvider, common.TranslateString("#970"))
			v.exitButton.MoveTo(33, 537)
			v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(v.exitButton)
		},
		func() {
			v.heroRenderInfo[common.HeroBarbarian] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectBarbarianUnselected, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectBarbarianUnselectedH, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectBarbarianForwardWalk, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectBarbarianForwardWalkOverlay, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectBarbarianSelected, palettedefs.Fechar),
				nil,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectBarbarianBackWalk, palettedefs.Fechar),
				nil,
				image.Rectangle{Min: image.Point{364, 201}, Max: image.Point{90, 170}},
				v.soundManager.LoadSoundEffect(resourcepaths.SFXBarbarianSelect),
				v.soundManager.LoadSoundEffect(resourcepaths.SFXBarbarianDeselect),
			}
			v.heroRenderInfo[common.HeroBarbarian].IdleSprite.MoveTo(400, 330)
			v.heroRenderInfo[common.HeroBarbarian].IdleSprite.Animate = true
			v.heroRenderInfo[common.HeroBarbarian].IdleSelectedSprite.MoveTo(400, 330)
			v.heroRenderInfo[common.HeroBarbarian].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroBarbarian].ForwardWalkSprite.MoveTo(400, 330)
			v.heroRenderInfo[common.HeroBarbarian].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroBarbarian].ForwardWalkSprite.SpecialFrameTime = 2500
			v.heroRenderInfo[common.HeroBarbarian].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroBarbarian].ForwardWalkSpriteOverlay.MoveTo(400, 330)
			v.heroRenderInfo[common.HeroBarbarian].ForwardWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[common.HeroBarbarian].ForwardWalkSpriteOverlay.SpecialFrameTime = 2500
			v.heroRenderInfo[common.HeroBarbarian].ForwardWalkSpriteOverlay.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroBarbarian].SelectedSprite.MoveTo(400, 330)
			v.heroRenderInfo[common.HeroBarbarian].SelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroBarbarian].BackWalkSprite.MoveTo(400, 330)
			v.heroRenderInfo[common.HeroBarbarian].BackWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroBarbarian].BackWalkSprite.SpecialFrameTime = 1000
			v.heroRenderInfo[common.HeroBarbarian].BackWalkSprite.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[common.HeroSorceress] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecSorceressUnselected, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecSorceressUnselectedH, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecSorceressForwardWalk, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecSorceressForwardWalkOverlay, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecSorceressSelected, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecSorceressSelectedOverlay, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecSorceressBackWalk, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecSorceressBackWalkOverlay, palettedefs.Fechar),
				image.Rectangle{Min: image.Point{580, 240}, Max: image.Point{65, 160}},
				v.soundManager.LoadSoundEffect(resourcepaths.SFXSorceressSelect),
				v.soundManager.LoadSoundEffect(resourcepaths.SFXSorceressDeselect),
			}
			v.heroRenderInfo[common.HeroSorceress].IdleSprite.MoveTo(626, 352)
			v.heroRenderInfo[common.HeroSorceress].IdleSprite.Animate = true
			v.heroRenderInfo[common.HeroSorceress].IdleSelectedSprite.MoveTo(626, 352)
			v.heroRenderInfo[common.HeroSorceress].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroSorceress].ForwardWalkSprite.MoveTo(626, 352)
			v.heroRenderInfo[common.HeroSorceress].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroSorceress].ForwardWalkSprite.SpecialFrameTime = 2300
			v.heroRenderInfo[common.HeroSorceress].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroSorceress].ForwardWalkSpriteOverlay.Blend = true
			v.heroRenderInfo[common.HeroSorceress].ForwardWalkSpriteOverlay.MoveTo(626, 352)
			v.heroRenderInfo[common.HeroSorceress].ForwardWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[common.HeroSorceress].ForwardWalkSpriteOverlay.SpecialFrameTime = 2300
			v.heroRenderInfo[common.HeroSorceress].ForwardWalkSpriteOverlay.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroSorceress].SelectedSprite.MoveTo(626, 352)
			v.heroRenderInfo[common.HeroSorceress].SelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroSorceress].SelectedSpriteOverlay.Blend = true
			v.heroRenderInfo[common.HeroSorceress].SelectedSpriteOverlay.MoveTo(626, 352)
			v.heroRenderInfo[common.HeroSorceress].SelectedSpriteOverlay.Animate = true
			v.heroRenderInfo[common.HeroSorceress].BackWalkSprite.MoveTo(626, 352)
			v.heroRenderInfo[common.HeroSorceress].BackWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroSorceress].BackWalkSprite.SpecialFrameTime = 1200
			v.heroRenderInfo[common.HeroSorceress].BackWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroSorceress].BackWalkSpriteOverlay.Blend = true
			v.heroRenderInfo[common.HeroSorceress].BackWalkSpriteOverlay.MoveTo(626, 352)
			v.heroRenderInfo[common.HeroSorceress].BackWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[common.HeroSorceress].BackWalkSpriteOverlay.SpecialFrameTime = 1200
			v.heroRenderInfo[common.HeroSorceress].BackWalkSpriteOverlay.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[common.HeroNecromancer] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectNecromancerUnselected, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectNecromancerUnselectedH, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecNecromancerForwardWalk, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecNecromancerForwardWalkOverlay, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecNecromancerSelected, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecNecromancerSelectedOverlay, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecNecromancerBackWalk, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecNecromancerBackWalkOverlay, palettedefs.Fechar),
				image.Rectangle{Min: image.Point{265, 220}, Max: image.Point{55, 175}},
				v.soundManager.LoadSoundEffect(resourcepaths.SFXNecromancerSelect),
				v.soundManager.LoadSoundEffect(resourcepaths.SFXNecromancerDeselect),
			}
			v.heroRenderInfo[common.HeroNecromancer].IdleSprite.MoveTo(300, 335)
			v.heroRenderInfo[common.HeroNecromancer].IdleSprite.Animate = true
			v.heroRenderInfo[common.HeroNecromancer].IdleSelectedSprite.MoveTo(300, 335)
			v.heroRenderInfo[common.HeroNecromancer].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroNecromancer].ForwardWalkSprite.MoveTo(300, 335)
			v.heroRenderInfo[common.HeroNecromancer].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroNecromancer].ForwardWalkSprite.SpecialFrameTime = 2000
			v.heroRenderInfo[common.HeroNecromancer].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroNecromancer].ForwardWalkSpriteOverlay.Blend = true
			v.heroRenderInfo[common.HeroNecromancer].ForwardWalkSpriteOverlay.MoveTo(300, 335)
			v.heroRenderInfo[common.HeroNecromancer].ForwardWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[common.HeroNecromancer].ForwardWalkSpriteOverlay.SpecialFrameTime = 2000
			v.heroRenderInfo[common.HeroNecromancer].ForwardWalkSpriteOverlay.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroNecromancer].SelectedSprite.MoveTo(300, 335)
			v.heroRenderInfo[common.HeroNecromancer].SelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroNecromancer].SelectedSpriteOverlay.Blend = true
			v.heroRenderInfo[common.HeroNecromancer].SelectedSpriteOverlay.MoveTo(300, 335)
			v.heroRenderInfo[common.HeroNecromancer].SelectedSpriteOverlay.Animate = true
			v.heroRenderInfo[common.HeroNecromancer].BackWalkSprite.MoveTo(300, 335)
			v.heroRenderInfo[common.HeroNecromancer].BackWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroNecromancer].BackWalkSprite.SpecialFrameTime = 1500
			v.heroRenderInfo[common.HeroNecromancer].BackWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroNecromancer].BackWalkSpriteOverlay.Blend = true
			v.heroRenderInfo[common.HeroNecromancer].BackWalkSpriteOverlay.MoveTo(300, 335)
			v.heroRenderInfo[common.HeroNecromancer].BackWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[common.HeroNecromancer].BackWalkSpriteOverlay.SpecialFrameTime = 1500
			v.heroRenderInfo[common.HeroNecromancer].BackWalkSpriteOverlay.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[common.HeroPaladin] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectPaladinUnselected, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectPaladinUnselectedH, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecPaladinForwardWalk, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecPaladinForwardWalkOverlay, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecPaladinSelected, palettedefs.Fechar),
				nil,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecPaladinBackWalk, palettedefs.Fechar),
				nil,
				image.Rectangle{Min: image.Point{490, 210}, Max: image.Point{65, 180}},
				v.soundManager.LoadSoundEffect(resourcepaths.SFXPaladinSelect),
				v.soundManager.LoadSoundEffect(resourcepaths.SFXPaladinDeselect),
			}
			v.heroRenderInfo[common.HeroPaladin].IdleSprite.MoveTo(521, 338)
			v.heroRenderInfo[common.HeroPaladin].IdleSprite.Animate = true
			v.heroRenderInfo[common.HeroPaladin].IdleSelectedSprite.MoveTo(521, 338)
			v.heroRenderInfo[common.HeroPaladin].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroPaladin].ForwardWalkSprite.MoveTo(521, 338)
			v.heroRenderInfo[common.HeroPaladin].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroPaladin].ForwardWalkSprite.SpecialFrameTime = 3400
			v.heroRenderInfo[common.HeroPaladin].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroPaladin].ForwardWalkSpriteOverlay.MoveTo(521, 338)
			v.heroRenderInfo[common.HeroPaladin].ForwardWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[common.HeroPaladin].ForwardWalkSpriteOverlay.SpecialFrameTime = 3400
			v.heroRenderInfo[common.HeroPaladin].ForwardWalkSpriteOverlay.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroPaladin].SelectedSprite.MoveTo(521, 338)
			v.heroRenderInfo[common.HeroPaladin].SelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroPaladin].BackWalkSprite.MoveTo(521, 338)
			v.heroRenderInfo[common.HeroPaladin].BackWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroPaladin].BackWalkSprite.SpecialFrameTime = 1300
			v.heroRenderInfo[common.HeroPaladin].BackWalkSprite.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[common.HeroAmazon] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectAmazonUnselected, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectAmazonUnselectedH, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecAmazonForwardWalk, palettedefs.Fechar),
				nil,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecAmazonSelected, palettedefs.Fechar),
				nil,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelecAmazonBackWalk, palettedefs.Fechar),
				nil,
				image.Rectangle{Min: image.Point{70, 220}, Max: image.Point{55, 200}},
				v.soundManager.LoadSoundEffect(resourcepaths.SFXAmazonSelect),
				v.soundManager.LoadSoundEffect(resourcepaths.SFXAmazonDeselect),
			}
			v.heroRenderInfo[common.HeroAmazon].IdleSprite.MoveTo(100, 339)
			v.heroRenderInfo[common.HeroAmazon].IdleSprite.Animate = true
			v.heroRenderInfo[common.HeroAmazon].IdleSelectedSprite.MoveTo(100, 339)
			v.heroRenderInfo[common.HeroAmazon].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroAmazon].ForwardWalkSprite.MoveTo(100, 339)
			v.heroRenderInfo[common.HeroAmazon].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroAmazon].ForwardWalkSprite.SpecialFrameTime = 2200
			v.heroRenderInfo[common.HeroAmazon].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroAmazon].SelectedSprite.MoveTo(100, 339)
			v.heroRenderInfo[common.HeroAmazon].SelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroAmazon].BackWalkSprite.MoveTo(100, 339)
			v.heroRenderInfo[common.HeroAmazon].BackWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroAmazon].BackWalkSprite.SpecialFrameTime = 1500
			v.heroRenderInfo[common.HeroAmazon].BackWalkSprite.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[common.HeroAssassin] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectAssassinUnselected, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectAssassinUnselectedH, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectAssassinForwardWalk, palettedefs.Fechar),
				nil,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectAssassinSelected, palettedefs.Fechar),
				nil,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectAssassinBackWalk, palettedefs.Fechar),
				nil,
				image.Rectangle{Min: image.Point{175, 235}, Max: image.Point{50, 180}},
				v.soundManager.LoadSoundEffect(resourcepaths.SFXAssassinSelect),
				v.soundManager.LoadSoundEffect(resourcepaths.SFXAssassinDeselect),
			}
			v.heroRenderInfo[common.HeroAssassin].IdleSprite.MoveTo(231, 365)
			v.heroRenderInfo[common.HeroAssassin].IdleSprite.Animate = true
			v.heroRenderInfo[common.HeroAssassin].IdleSelectedSprite.MoveTo(231, 365)
			v.heroRenderInfo[common.HeroAssassin].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroAssassin].ForwardWalkSprite.MoveTo(231, 365)
			v.heroRenderInfo[common.HeroAssassin].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroAssassin].ForwardWalkSprite.SpecialFrameTime = 3800
			v.heroRenderInfo[common.HeroAssassin].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroAssassin].SelectedSprite.MoveTo(231, 365)
			v.heroRenderInfo[common.HeroAssassin].SelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroAssassin].BackWalkSprite.MoveTo(231, 365)
			v.heroRenderInfo[common.HeroAssassin].BackWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroAssassin].BackWalkSprite.SpecialFrameTime = 1500
			v.heroRenderInfo[common.HeroAssassin].BackWalkSprite.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[common.HeroDruid] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectDruidUnselected, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectDruidUnselectedH, palettedefs.Fechar),
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectDruidForwardWalk, palettedefs.Fechar),
				nil,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectDruidSelected, palettedefs.Fechar),
				nil,
				v.fileProvider.LoadSprite(resourcepaths.CharacterSelectDruidBackWalk, palettedefs.Fechar),
				nil,
				image.Rectangle{Min: image.Point{680, 220}, Max: image.Point{70, 195}},
				v.soundManager.LoadSoundEffect(resourcepaths.SFXDruidSelect),
				v.soundManager.LoadSoundEffect(resourcepaths.SFXDruidDeselect),
			}
			v.heroRenderInfo[common.HeroDruid].IdleSprite.MoveTo(720, 370)
			v.heroRenderInfo[common.HeroDruid].IdleSprite.Animate = true
			v.heroRenderInfo[common.HeroDruid].IdleSelectedSprite.MoveTo(720, 370)
			v.heroRenderInfo[common.HeroDruid].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroDruid].ForwardWalkSprite.MoveTo(720, 370)
			v.heroRenderInfo[common.HeroDruid].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroDruid].ForwardWalkSprite.SpecialFrameTime = 4800
			v.heroRenderInfo[common.HeroDruid].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[common.HeroDruid].SelectedSprite.MoveTo(720, 370)
			v.heroRenderInfo[common.HeroDruid].SelectedSprite.Animate = true
			v.heroRenderInfo[common.HeroDruid].BackWalkSprite.MoveTo(720, 370)
			v.heroRenderInfo[common.HeroDruid].BackWalkSprite.Animate = true
			v.heroRenderInfo[common.HeroDruid].BackWalkSprite.SpecialFrameTime = 1500
			v.heroRenderInfo[common.HeroDruid].BackWalkSprite.StopOnLastFrame = true
		},
	}
}

func (v *SelectHeroClass) Unload() {
	v.heroRenderInfo = nil
}

func (v *SelectHeroClass) onExitButtonClicked() {
	v.sceneProvider.SetNextScene(CreateCharacterSelect(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager))
}

func (v *SelectHeroClass) Render(screen *ebiten.Image) {
	v.bgImage.DrawSegments(screen, 4, 3, 0)
	v.headingLabel.Draw(screen)
	if v.selectedHero != common.HeroNone {
		v.heroClassLabel.Draw(screen)
		v.heroDesc1Label.Draw(screen)
		v.heroDesc2Label.Draw(screen)
		v.heroDesc3Label.Draw(screen)
	}
	for heroClass, heroInfo := range v.heroRenderInfo {
		if heroInfo.Stance == HeroStanceIdle || heroInfo.Stance == HeroStanceIdleSelected {
			v.renderHero(screen, heroClass)
		}
	}
	for heroClass, heroInfo := range v.heroRenderInfo {
		if heroInfo.Stance != HeroStanceIdle && heroInfo.Stance != HeroStanceIdleSelected {
			v.renderHero(screen, heroClass)
		}
	}
	v.campfire.Draw(screen)
}

func (v *SelectHeroClass) Update(tickTime float64) {
	canSelect := true
	for _, info := range v.heroRenderInfo {
		if info.Stance != HeroStanceIdle && info.Stance != HeroStanceIdleSelected && info.Stance != HeroStanceSelected {
			canSelect = false
			break
		}
	}
	allIdle := true
	for heroType, data := range v.heroRenderInfo {
		if allIdle && data.Stance != HeroStanceIdle {
			allIdle = false
		}
		v.updateHeroSelectionHover(heroType, canSelect)
	}
	if v.selectedHero != common.HeroNone && allIdle {
		v.selectedHero = common.HeroNone
	}
}

func (v *SelectHeroClass) updateHeroSelectionHover(hero common.Hero, canSelect bool) {
	renderInfo := v.heroRenderInfo[hero]
	switch renderInfo.Stance {
	case HeroStanceApproaching:
		if renderInfo.ForwardWalkSprite.OnLastFrame() {
			renderInfo.Stance = HeroStanceSelected
			renderInfo.SelectedSprite.ResetAnimation()
			if renderInfo.SelectedSpriteOverlay != nil {
				renderInfo.SelectedSpriteOverlay.ResetAnimation()
			}
		}
		return
	case HeroStanceRetreating:
		if renderInfo.BackWalkSprite.OnLastFrame() {
			renderInfo.Stance = HeroStanceIdle
			renderInfo.IdleSprite.ResetAnimation()
		}
		return
	}
	if !canSelect {
		return
	}
	if renderInfo.Stance == HeroStanceSelected {
		return
	}
	mouseX := v.uiManager.CursorX
	mouseY := v.uiManager.CursorY
	b := renderInfo.SelectionBounds
	mouseHover := (mouseX >= b.Min.X) && (mouseX <= b.Min.X+b.Max.X) && (mouseY >= b.Min.Y) && (mouseY <= b.Min.Y+b.Max.Y)
	if mouseHover && v.uiManager.CursorButtonPressed(ui.CursorButtonLeft) {
		// showEntryUi = true;
		renderInfo.Stance = HeroStanceApproaching
		renderInfo.ForwardWalkSprite.ResetAnimation()
		if renderInfo.ForwardWalkSpriteOverlay != nil {
			renderInfo.ForwardWalkSpriteOverlay.ResetAnimation()
		}
		for _, heroInfo := range v.heroRenderInfo {
			if heroInfo.Stance != HeroStanceSelected {
				continue
			}
			heroInfo.SelectSfx.Stop()
			heroInfo.DeselectSfx.Play()
			heroInfo.Stance = HeroStanceRetreating
			heroInfo.BackWalkSprite.ResetAnimation()
			if heroInfo.BackWalkSpriteOverlay != nil {
				heroInfo.BackWalkSpriteOverlay.ResetAnimation()
			}
		}
		v.selectedHero = hero
		v.updateHeroText()
		renderInfo.SelectSfx.Play()

		return
	}

	if mouseHover {
		renderInfo.Stance = HeroStanceIdleSelected
	} else {
		renderInfo.Stance = HeroStanceIdle
	}

	if v.selectedHero == common.HeroNone && mouseHover {
		v.selectedHero = hero
		v.updateHeroText()
	}

}

func (v *SelectHeroClass) renderHero(screen *ebiten.Image, hero common.Hero) {
	renderInfo := v.heroRenderInfo[hero]
	switch renderInfo.Stance {
	case HeroStanceIdle:
		renderInfo.IdleSprite.Draw(screen)
	case HeroStanceIdleSelected:
		renderInfo.IdleSelectedSprite.Draw(screen)
	case HeroStanceApproaching:
		renderInfo.ForwardWalkSprite.Draw(screen)
		if renderInfo.ForwardWalkSpriteOverlay != nil {
			renderInfo.ForwardWalkSpriteOverlay.Draw(screen)
		}
	case HeroStanceSelected:
		renderInfo.SelectedSprite.Draw(screen)
		if renderInfo.SelectedSpriteOverlay != nil {
			renderInfo.SelectedSpriteOverlay.Draw(screen)
		}
	case HeroStanceRetreating:
		renderInfo.BackWalkSprite.Draw(screen)
		if renderInfo.BackWalkSpriteOverlay != nil {
			renderInfo.BackWalkSpriteOverlay.Draw(screen)
		}
	}
}

func (v *SelectHeroClass) updateHeroText() {
	switch v.selectedHero {
	case common.HeroNone:
		return
	case common.HeroBarbarian:
		v.heroClassLabel.SetText(common.TranslateString("partycharbar"))
		v.setDescLabels("#1709")
	case common.HeroNecromancer:
		v.heroClassLabel.SetText(common.TranslateString("partycharnec"))
		v.setDescLabels("#1704")
	case common.HeroPaladin:
		v.heroClassLabel.SetText(common.TranslateString("partycharpal"))
		v.setDescLabels("#1711")
	case common.HeroAssassin:
		v.heroClassLabel.SetText(common.TranslateString("partycharass"))
		v.setDescLabels("#305")
	case common.HeroSorceress:
		v.heroClassLabel.SetText(common.TranslateString("partycharsor"))
		v.setDescLabels("#1710")
	case common.HeroAmazon:
		v.heroClassLabel.SetText(common.TranslateString("partycharama"))
		v.setDescLabels("#1698")
	case common.HeroDruid:
		v.heroClassLabel.SetText(common.TranslateString("partychardru"))
		v.setDescLabels("#304")
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
	heroDesc := common.TranslateString(descKey)
	parts := common.SplitIntoLinesWithMaxWidth(heroDesc, 37)
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
