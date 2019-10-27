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
}

type SelectHeroClass struct {
	uiManager      *UI.Manager
	soundManager   *Sound.Manager
	fileProvider   Common.FileProvider
	sceneProvider  SceneProvider
	bgImage        *Common.Sprite
	campfire       *Common.Sprite
	headingLabel   *UI.Label
	heroRenderInfo map[Common.Hero]*HeroRenderInfo
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
	}
}

func (v *SelectHeroClass) Unload() {

}

func (v *SelectHeroClass) Render(screen *ebiten.Image) {
	v.bgImage.DrawSegments(screen, 4, 3, 0)
	v.headingLabel.Draw(screen)
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
	for heroType := range v.heroRenderInfo {
		v.updateHeroSelectionHover(heroType, canSelect)
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
			// PlayHeroDeselected(ri.Key);
			heroInfo.Stance = HeroStanceRetreating
			heroInfo.BackWalkSprite.ResetAnimation()
			if heroInfo.BackWalkSpriteOverlay != nil {
				heroInfo.BackWalkSpriteOverlay.ResetAnimation()
			}
		}
		// selectedHero = hero;
		// UpdateHeroText();
		// PlayHeroSelected(hero);

		return
	}

	if mouseHover {
		renderInfo.Stance = HeroStanceIdleSelected
	} else {
		renderInfo.Stance = HeroStanceIdle
	}

	/*
		   if (selectedHero == null && mouseHover)
		   {
			   selectedHero = hero;
			   UpdateHeroText();
		   }
	*/

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
