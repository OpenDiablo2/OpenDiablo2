package Scenes

import (
	"image"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/essial/OpenDiablo2/Sound"
	"github.com/essial/OpenDiablo2/UI"
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
	IdleSprite               *Common.Sprite
	IdleSelectedSprite       *Common.Sprite
	ForwardWalkSprite        *Common.Sprite
	ForwardWalkSpriteOverlay *Common.Sprite
	SelectedSprite           *Common.Sprite
	SelectedSpriteOverlay    *Common.Sprite
	BackWalkSprite           *Common.Sprite
	BackWalkSpriteOverlay    *Common.Sprite
	SelectionBounds          image.Rectangle
	SelectSfx                *Sound.SoundEffect
	DeselectSfx              *Sound.SoundEffect
}

type SelectHeroClass struct {
	uiManager      *UI.Manager
	soundManager   *Sound.Manager
	fileProvider   Common.FileProvider
	sceneProvider  SceneProvider
	bgImage        *Common.Sprite
	campfire       *Common.Sprite
	headingLabel   *UI.Label
	heroClassLabel *UI.Label
	heroDesc1Label *UI.Label
	heroDesc2Label *UI.Label
	heroDesc3Label *UI.Label
	heroRenderInfo map[Common.Hero]*HeroRenderInfo
	selectedHero   Common.Hero
}

func CreateSelectHeroClass(
	fileProvider Common.FileProvider,
	sceneProvider SceneProvider,
	uiManager *UI.Manager, soundManager *Sound.Manager,
) *SelectHeroClass {
	result := &SelectHeroClass{
		uiManager:      uiManager,
		sceneProvider:  sceneProvider,
		fileProvider:   fileProvider,
		soundManager:   soundManager,
		heroRenderInfo: make(map[Common.Hero]*HeroRenderInfo),
		selectedHero:   Common.HeroNone,
	}
	return result
}

func (v *SelectHeroClass) Load() []func() {
	v.soundManager.PlayBGM(ResourcePaths.BGMTitle)
	return []func(){
		func() {
			v.bgImage = v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectBackground, Palettes.Fechar)
			v.bgImage.MoveTo(0, 0)
		},
		func() {
			v.headingLabel = UI.CreateLabel(v.fileProvider, ResourcePaths.Font30, Palettes.Units)
			fontWidth, _ := v.headingLabel.GetSize()
			v.headingLabel.MoveTo(400-int(fontWidth/2), 17)
			v.headingLabel.SetText("Select Hero Class")
			v.headingLabel.Alignment = UI.LabelAlignCenter
		},
		func() {
			v.heroClassLabel = UI.CreateLabel(v.fileProvider, ResourcePaths.Font30, Palettes.Units)
			v.heroClassLabel.Alignment = UI.LabelAlignCenter
			v.heroClassLabel.MoveTo(400, 65)
		},
		func() {
			v.heroDesc1Label = UI.CreateLabel(v.fileProvider, ResourcePaths.Font16, Palettes.Units)
			v.heroDesc1Label.Alignment = UI.LabelAlignCenter
			v.heroDesc1Label.MoveTo(400, 100)
		},
		func() {
			v.heroDesc2Label = UI.CreateLabel(v.fileProvider, ResourcePaths.Font16, Palettes.Units)
			v.heroDesc2Label.Alignment = UI.LabelAlignCenter
			v.heroDesc2Label.MoveTo(400, 115)
		},
		func() {
			v.heroDesc3Label = UI.CreateLabel(v.fileProvider, ResourcePaths.Font16, Palettes.Units)
			v.heroDesc3Label.Alignment = UI.LabelAlignCenter
			v.heroDesc3Label.MoveTo(400, 130)
		},
		func() {
			v.campfire = v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectCampfire, Palettes.Fechar)
			v.campfire.MoveTo(380, 335)
			v.campfire.Animate = true
		},
		func() {
			v.heroRenderInfo[Common.HeroBarbarian] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectBarbarianUnselected, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectBarbarianUnselectedH, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectBarbarianForwardWalk, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectBarbarianForwardWalkOverlay, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectBarbarianSelected, Palettes.Fechar),
				nil,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectBarbarianBackWalk, Palettes.Fechar),
				nil,
				image.Rectangle{Min: image.Point{364, 201}, Max: image.Point{90, 170}},
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXBarbarianSelect),
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXBarbarianDeselect),
			}
			v.heroRenderInfo[Common.HeroBarbarian].IdleSprite.MoveTo(400, 330)
			v.heroRenderInfo[Common.HeroBarbarian].IdleSprite.Animate = true
			v.heroRenderInfo[Common.HeroBarbarian].IdleSelectedSprite.MoveTo(400, 330)
			v.heroRenderInfo[Common.HeroBarbarian].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroBarbarian].ForwardWalkSprite.MoveTo(400, 330)
			v.heroRenderInfo[Common.HeroBarbarian].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroBarbarian].ForwardWalkSprite.SpecialFrameTime = 2500
			v.heroRenderInfo[Common.HeroBarbarian].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroBarbarian].ForwardWalkSpriteOverlay.MoveTo(400, 330)
			v.heroRenderInfo[Common.HeroBarbarian].ForwardWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[Common.HeroBarbarian].ForwardWalkSpriteOverlay.SpecialFrameTime = 2500
			v.heroRenderInfo[Common.HeroBarbarian].ForwardWalkSpriteOverlay.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroBarbarian].SelectedSprite.MoveTo(400, 330)
			v.heroRenderInfo[Common.HeroBarbarian].SelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroBarbarian].BackWalkSprite.MoveTo(400, 330)
			v.heroRenderInfo[Common.HeroBarbarian].BackWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroBarbarian].BackWalkSprite.SpecialFrameTime = 1000
			v.heroRenderInfo[Common.HeroBarbarian].BackWalkSprite.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[Common.HeroSorceress] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecSorceressUnselected, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecSorceressUnselectedH, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecSorceressForwardWalk, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecSorceressForwardWalkOverlay, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecSorceressSelected, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecSorceressSelectedOverlay, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecSorceressBackWalk, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecSorceressBackWalkOverlay, Palettes.Fechar),
				image.Rectangle{Min: image.Point{580, 240}, Max: image.Point{65, 160}},
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXSorceressSelect),
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXSorceressDeselect),
			}
			v.heroRenderInfo[Common.HeroSorceress].IdleSprite.MoveTo(626, 352)
			v.heroRenderInfo[Common.HeroSorceress].IdleSprite.Animate = true
			v.heroRenderInfo[Common.HeroSorceress].IdleSelectedSprite.MoveTo(626, 352)
			v.heroRenderInfo[Common.HeroSorceress].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroSorceress].ForwardWalkSprite.MoveTo(626, 352)
			v.heroRenderInfo[Common.HeroSorceress].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroSorceress].ForwardWalkSprite.SpecialFrameTime = 2300
			v.heroRenderInfo[Common.HeroSorceress].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroSorceress].ForwardWalkSpriteOverlay.Blend = true
			v.heroRenderInfo[Common.HeroSorceress].ForwardWalkSpriteOverlay.MoveTo(626, 352)
			v.heroRenderInfo[Common.HeroSorceress].ForwardWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[Common.HeroSorceress].ForwardWalkSpriteOverlay.SpecialFrameTime = 2300
			v.heroRenderInfo[Common.HeroSorceress].ForwardWalkSpriteOverlay.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroSorceress].SelectedSprite.MoveTo(626, 352)
			v.heroRenderInfo[Common.HeroSorceress].SelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroSorceress].SelectedSpriteOverlay.Blend = true
			v.heroRenderInfo[Common.HeroSorceress].SelectedSpriteOverlay.MoveTo(626, 352)
			v.heroRenderInfo[Common.HeroSorceress].SelectedSpriteOverlay.Animate = true
			v.heroRenderInfo[Common.HeroSorceress].BackWalkSprite.MoveTo(626, 352)
			v.heroRenderInfo[Common.HeroSorceress].BackWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroSorceress].BackWalkSprite.SpecialFrameTime = 1200
			v.heroRenderInfo[Common.HeroSorceress].BackWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroSorceress].BackWalkSpriteOverlay.Blend = true
			v.heroRenderInfo[Common.HeroSorceress].BackWalkSpriteOverlay.MoveTo(626, 352)
			v.heroRenderInfo[Common.HeroSorceress].BackWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[Common.HeroSorceress].BackWalkSpriteOverlay.SpecialFrameTime = 1200
			v.heroRenderInfo[Common.HeroSorceress].BackWalkSpriteOverlay.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[Common.HeroNecromancer] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectNecromancerUnselected, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectNecromancerUnselectedH, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecNecromancerForwardWalk, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecNecromancerForwardWalkOverlay, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecNecromancerSelected, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecNecromancerSelectedOverlay, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecNecromancerBackWalk, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecNecromancerBackWalkOverlay, Palettes.Fechar),
				image.Rectangle{Min: image.Point{265, 220}, Max: image.Point{55, 175}},
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXNecromancerSelect),
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXNecromancerDeselect),
			}
			v.heroRenderInfo[Common.HeroNecromancer].IdleSprite.MoveTo(300, 335)
			v.heroRenderInfo[Common.HeroNecromancer].IdleSprite.Animate = true
			v.heroRenderInfo[Common.HeroNecromancer].IdleSelectedSprite.MoveTo(300, 335)
			v.heroRenderInfo[Common.HeroNecromancer].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroNecromancer].ForwardWalkSprite.MoveTo(300, 335)
			v.heroRenderInfo[Common.HeroNecromancer].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroNecromancer].ForwardWalkSprite.SpecialFrameTime = 2000
			v.heroRenderInfo[Common.HeroNecromancer].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroNecromancer].ForwardWalkSpriteOverlay.Blend = true
			v.heroRenderInfo[Common.HeroNecromancer].ForwardWalkSpriteOverlay.MoveTo(300, 335)
			v.heroRenderInfo[Common.HeroNecromancer].ForwardWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[Common.HeroNecromancer].ForwardWalkSpriteOverlay.SpecialFrameTime = 2000
			v.heroRenderInfo[Common.HeroNecromancer].ForwardWalkSpriteOverlay.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroNecromancer].SelectedSprite.MoveTo(300, 335)
			v.heroRenderInfo[Common.HeroNecromancer].SelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroNecromancer].SelectedSpriteOverlay.Blend = true
			v.heroRenderInfo[Common.HeroNecromancer].SelectedSpriteOverlay.MoveTo(300, 335)
			v.heroRenderInfo[Common.HeroNecromancer].SelectedSpriteOverlay.Animate = true
			v.heroRenderInfo[Common.HeroNecromancer].BackWalkSprite.MoveTo(300, 335)
			v.heroRenderInfo[Common.HeroNecromancer].BackWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroNecromancer].BackWalkSprite.SpecialFrameTime = 1500
			v.heroRenderInfo[Common.HeroNecromancer].BackWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroNecromancer].BackWalkSpriteOverlay.Blend = true
			v.heroRenderInfo[Common.HeroNecromancer].BackWalkSpriteOverlay.MoveTo(300, 335)
			v.heroRenderInfo[Common.HeroNecromancer].BackWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[Common.HeroNecromancer].BackWalkSpriteOverlay.SpecialFrameTime = 1500
			v.heroRenderInfo[Common.HeroNecromancer].BackWalkSpriteOverlay.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[Common.HeroPaladin] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectPaladinUnselected, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectPaladinUnselectedH, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecPaladinForwardWalk, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecPaladinForwardWalkOverlay, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecPaladinSelected, Palettes.Fechar),
				nil,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecPaladinBackWalk, Palettes.Fechar),
				nil,
				image.Rectangle{Min: image.Point{490, 210}, Max: image.Point{65, 180}},
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXPaladinSelect),
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXPaladinDeselect),
			}
			v.heroRenderInfo[Common.HeroPaladin].IdleSprite.MoveTo(521, 338)
			v.heroRenderInfo[Common.HeroPaladin].IdleSprite.Animate = true
			v.heroRenderInfo[Common.HeroPaladin].IdleSelectedSprite.MoveTo(521, 338)
			v.heroRenderInfo[Common.HeroPaladin].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroPaladin].ForwardWalkSprite.MoveTo(521, 338)
			v.heroRenderInfo[Common.HeroPaladin].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroPaladin].ForwardWalkSprite.SpecialFrameTime = 3400
			v.heroRenderInfo[Common.HeroPaladin].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroPaladin].ForwardWalkSpriteOverlay.MoveTo(521, 338)
			v.heroRenderInfo[Common.HeroPaladin].ForwardWalkSpriteOverlay.Animate = true
			v.heroRenderInfo[Common.HeroPaladin].ForwardWalkSpriteOverlay.SpecialFrameTime = 3400
			v.heroRenderInfo[Common.HeroPaladin].ForwardWalkSpriteOverlay.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroPaladin].SelectedSprite.MoveTo(521, 338)
			v.heroRenderInfo[Common.HeroPaladin].SelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroPaladin].BackWalkSprite.MoveTo(521, 338)
			v.heroRenderInfo[Common.HeroPaladin].BackWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroPaladin].BackWalkSprite.SpecialFrameTime = 1300
			v.heroRenderInfo[Common.HeroPaladin].BackWalkSprite.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[Common.HeroAmazon] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectAmazonUnselected, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectAmazonUnselectedH, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecAmazonForwardWalk, Palettes.Fechar),
				nil,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecAmazonSelected, Palettes.Fechar),
				nil,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelecAmazonBackWalk, Palettes.Fechar),
				nil,
				image.Rectangle{Min: image.Point{70, 220}, Max: image.Point{55, 200}},
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXAmazonSelect),
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXAmazonDeselect),
			}
			v.heroRenderInfo[Common.HeroAmazon].IdleSprite.MoveTo(100, 339)
			v.heroRenderInfo[Common.HeroAmazon].IdleSprite.Animate = true
			v.heroRenderInfo[Common.HeroAmazon].IdleSelectedSprite.MoveTo(100, 339)
			v.heroRenderInfo[Common.HeroAmazon].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroAmazon].ForwardWalkSprite.MoveTo(100, 339)
			v.heroRenderInfo[Common.HeroAmazon].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroAmazon].ForwardWalkSprite.SpecialFrameTime = 2200
			v.heroRenderInfo[Common.HeroAmazon].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroAmazon].SelectedSprite.MoveTo(100, 339)
			v.heroRenderInfo[Common.HeroAmazon].SelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroAmazon].BackWalkSprite.MoveTo(100, 339)
			v.heroRenderInfo[Common.HeroAmazon].BackWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroAmazon].BackWalkSprite.SpecialFrameTime = 1500
			v.heroRenderInfo[Common.HeroAmazon].BackWalkSprite.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[Common.HeroAssassin] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectAssassinUnselected, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectAssassinUnselectedH, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectAssassinForwardWalk, Palettes.Fechar),
				nil,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectAssassinSelected, Palettes.Fechar),
				nil,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectAssassinBackWalk, Palettes.Fechar),
				nil,
				image.Rectangle{Min: image.Point{175, 235}, Max: image.Point{50, 180}},
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXAssassinSelect),
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXAssassinDeselect),
			}
			v.heroRenderInfo[Common.HeroAssassin].IdleSprite.MoveTo(231, 365)
			v.heroRenderInfo[Common.HeroAssassin].IdleSprite.Animate = true
			v.heroRenderInfo[Common.HeroAssassin].IdleSelectedSprite.MoveTo(231, 365)
			v.heroRenderInfo[Common.HeroAssassin].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroAssassin].ForwardWalkSprite.MoveTo(231, 365)
			v.heroRenderInfo[Common.HeroAssassin].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroAssassin].ForwardWalkSprite.SpecialFrameTime = 3800
			v.heroRenderInfo[Common.HeroAssassin].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroAssassin].SelectedSprite.MoveTo(231, 365)
			v.heroRenderInfo[Common.HeroAssassin].SelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroAssassin].BackWalkSprite.MoveTo(231, 365)
			v.heroRenderInfo[Common.HeroAssassin].BackWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroAssassin].BackWalkSprite.SpecialFrameTime = 1500
			v.heroRenderInfo[Common.HeroAssassin].BackWalkSprite.StopOnLastFrame = true
		},
		func() {
			v.heroRenderInfo[Common.HeroDruid] = &HeroRenderInfo{
				HeroStanceIdle,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectDruidUnselected, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectDruidUnselectedH, Palettes.Fechar),
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectDruidForwardWalk, Palettes.Fechar),
				nil,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectDruidSelected, Palettes.Fechar),
				nil,
				v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectDruidBackWalk, Palettes.Fechar),
				nil,
				image.Rectangle{Min: image.Point{680, 220}, Max: image.Point{70, 195}},
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXDruidSelect),
				v.soundManager.LoadSoundEffect(ResourcePaths.SFXDruidDeselect),
			}
			v.heroRenderInfo[Common.HeroDruid].IdleSprite.MoveTo(720, 370)
			v.heroRenderInfo[Common.HeroDruid].IdleSprite.Animate = true
			v.heroRenderInfo[Common.HeroDruid].IdleSelectedSprite.MoveTo(720, 370)
			v.heroRenderInfo[Common.HeroDruid].IdleSelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroDruid].ForwardWalkSprite.MoveTo(720, 370)
			v.heroRenderInfo[Common.HeroDruid].ForwardWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroDruid].ForwardWalkSprite.SpecialFrameTime = 4800
			v.heroRenderInfo[Common.HeroDruid].ForwardWalkSprite.StopOnLastFrame = true
			v.heroRenderInfo[Common.HeroDruid].SelectedSprite.MoveTo(720, 370)
			v.heroRenderInfo[Common.HeroDruid].SelectedSprite.Animate = true
			v.heroRenderInfo[Common.HeroDruid].BackWalkSprite.MoveTo(720, 370)
			v.heroRenderInfo[Common.HeroDruid].BackWalkSprite.Animate = true
			v.heroRenderInfo[Common.HeroDruid].BackWalkSprite.SpecialFrameTime = 1500
			v.heroRenderInfo[Common.HeroDruid].BackWalkSprite.StopOnLastFrame = true
		},
	}
}

func (v *SelectHeroClass) Unload() {

}

func (v *SelectHeroClass) Render(screen *ebiten.Image) {
	v.bgImage.DrawSegments(screen, 4, 3, 0)
	v.headingLabel.Draw(screen)
	if v.selectedHero != Common.HeroNone {
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
	if v.selectedHero != Common.HeroNone && allIdle {
		v.selectedHero = Common.HeroNone
	}
}

func (v *SelectHeroClass) updateHeroSelectionHover(hero Common.Hero, canSelect bool) {
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
	if mouseHover && v.uiManager.CursorButtonPressed(UI.CursorButtonLeft) {
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

	if v.selectedHero == Common.HeroNone && mouseHover {
		v.selectedHero = hero
		v.updateHeroText()
	}

}

func (v *SelectHeroClass) renderHero(screen *ebiten.Image, hero Common.Hero) {
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
	case Common.HeroNone:
		return
	case Common.HeroBarbarian:
		v.heroClassLabel.SetText(Common.TranslateString("partycharbar"))
		v.setDescLabels("#1709")
	case Common.HeroNecromancer:
		v.heroClassLabel.SetText(Common.TranslateString("partycharnec"))
		v.setDescLabels("#1704")
	case Common.HeroPaladin:
		v.heroClassLabel.SetText(Common.TranslateString("partycharpal"))
		v.setDescLabels("#1711")
	case Common.HeroAssassin:
		v.heroClassLabel.SetText(Common.TranslateString("partycharass"))
		v.setDescLabels("#305")
	case Common.HeroSorceress:
		v.heroClassLabel.SetText(Common.TranslateString("partycharsor"))
		v.setDescLabels("#1710")
	case Common.HeroAmazon:
		v.heroClassLabel.SetText(Common.TranslateString("partycharama"))
		v.setDescLabels("#1698")
	case Common.HeroDruid:
		v.heroClassLabel.SetText(Common.TranslateString("partychardru"))
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
	heroDesc := Common.TranslateString(descKey)
	parts := Common.SplitIntoLinesWithMaxWidth(heroDesc, 37)
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
