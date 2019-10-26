package Scenes

import (
	"image/color"
	"os"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/Sound"
	"github.com/essial/OpenDiablo2/UI"

	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/hajimehoshi/ebiten"
)

// MainMenu represents the main menu
type MainMenu struct {
	uiManager           *UI.Manager
	soundManager        *Sound.Manager
	fileProvider        Common.FileProvider
	sceneProvider       SceneProvider
	trademarkBackground *Common.Sprite
	background          *Common.Sprite
	diabloLogoLeft      *Common.Sprite
	diabloLogoRight     *Common.Sprite
	diabloLogoLeftBack  *Common.Sprite
	diabloLogoRightBack *Common.Sprite
	exitDiabloButton    *UI.Button
	creditsButton       *UI.Button
	copyrightLabel      *UI.Label
	copyrightLabel2     *UI.Label
	showTrademarkScreen bool
	leftButtonHeld      bool
}

// CreateMainMenu creates an instance of MainMenu
func CreateMainMenu(fileProvider Common.FileProvider, sceneProvider SceneProvider, uiManager *UI.Manager, soundManager *Sound.Manager) *MainMenu {
	result := &MainMenu{
		fileProvider:        fileProvider,
		uiManager:           uiManager,
		soundManager:        soundManager,
		sceneProvider:       sceneProvider,
		showTrademarkScreen: true,
	}
	return result
}

// Load is called to load the resources for the main menu
func (v *MainMenu) Load() []func() {
	v.soundManager.PlayBGM(ResourcePaths.BGMTitle)
	return []func(){
		func() {
			v.copyrightLabel = UI.CreateLabel(v.fileProvider, ResourcePaths.FontFormal12, Palettes.Static)
			v.copyrightLabel.Alignment = UI.LabelAlignCenter
			v.copyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
			v.copyrightLabel.Color = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel.MoveTo(400, 500)
		},
		func() {
			v.copyrightLabel2 = UI.CreateLabel(v.fileProvider, ResourcePaths.FontFormal12, Palettes.Static)
			v.copyrightLabel2.Alignment = UI.LabelAlignCenter
			v.copyrightLabel2.SetText("All Rights Reserved.")
			v.copyrightLabel2.Color = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel2.MoveTo(400, 525)
		},
		func() {
			v.background = v.fileProvider.LoadSprite(ResourcePaths.GameSelectScreen, Palettes.Sky)
			v.background.MoveTo(0, 0)
		},
		func() {
			v.trademarkBackground = v.fileProvider.LoadSprite(ResourcePaths.TrademarkScreen, Palettes.Sky)
			v.trademarkBackground.MoveTo(0, 0)
		},
		func() {
			v.diabloLogoLeft = v.fileProvider.LoadSprite(ResourcePaths.Diablo2LogoFireLeft, Palettes.Units)
			v.diabloLogoLeft.Blend = true
			v.diabloLogoLeft.Animate = true
			v.diabloLogoLeft.MoveTo(400, 120)
		},
		func() {
			v.diabloLogoRight = v.fileProvider.LoadSprite(ResourcePaths.Diablo2LogoFireRight, Palettes.Units)
			v.diabloLogoRight.Blend = true
			v.diabloLogoRight.Animate = true
			v.diabloLogoRight.MoveTo(400, 120)
		},
		func() {
			v.diabloLogoLeftBack = v.fileProvider.LoadSprite(ResourcePaths.Diablo2LogoBlackLeft, Palettes.Units)
			v.diabloLogoLeftBack.MoveTo(400, 120)
		},
		func() {
			v.diabloLogoRightBack = v.fileProvider.LoadSprite(ResourcePaths.Diablo2LogoBlackRight, Palettes.Units)
			v.diabloLogoRightBack.MoveTo(400, 120)
		},
		func() {
			v.exitDiabloButton = UI.CreateButton(UI.ButtonTypeWide, v.fileProvider, "EXIT DIABLO II")
			v.exitDiabloButton.MoveTo(264, 535)
			v.exitDiabloButton.SetVisible(false)
			v.exitDiabloButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(v.exitDiabloButton)
		},
		func() {
			v.creditsButton = UI.CreateButton(UI.ButtonTypeShort, v.fileProvider, "CREDITS")
			v.creditsButton.MoveTo(264, 505)
			v.creditsButton.SetVisible(false)
			v.creditsButton.OnActivated(func() { v.onCreditsButtonClicked() })
			v.uiManager.AddWidget(v.creditsButton)
		},
	}
}

func (v *MainMenu) onExitButtonClicked() {
	os.Exit(0)
}

func (v *MainMenu) onCreditsButtonClicked() {
	v.sceneProvider.SetNextScene(CreateCredits(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager))
}

// Unload unloads the data for the main menu
func (v *MainMenu) Unload() {

}

// Render renders the main menu
func (v *MainMenu) Render(screen *ebiten.Image) {
	if v.showTrademarkScreen {
		v.trademarkBackground.DrawSegments(screen, 4, 3, 0)
	} else {
		v.background.DrawSegments(screen, 4, 3, 0)
	}
	v.diabloLogoLeftBack.Draw(screen)
	v.diabloLogoRightBack.Draw(screen)
	v.diabloLogoLeft.Draw(screen)
	v.diabloLogoRight.Draw(screen)

	if v.showTrademarkScreen {
		v.copyrightLabel.Draw(screen)
		v.copyrightLabel2.Draw(screen)
	} else {

	}
}

// Update runs the update logic on the main menu
func (v *MainMenu) Update() {
	if v.showTrademarkScreen {
		if v.uiManager.CursorButtonPressed(UI.CursorButtonLeft) {
			v.leftButtonHeld = true
			v.showTrademarkScreen = false
			v.exitDiabloButton.SetVisible(true)
			v.creditsButton.SetVisible(true)
		}
		return
	}
}
