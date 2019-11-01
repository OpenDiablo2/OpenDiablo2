package Scenes

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/PaletteDefs"
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
	singlePlayerButton  *UI.Button
	githubButton        *UI.Button
	exitDiabloButton    *UI.Button
	creditsButton       *UI.Button
	cinematicsButton    *UI.Button
	mapTestButton       *UI.Button
	copyrightLabel      *UI.Label
	copyrightLabel2     *UI.Label
	openDiabloLabel     *UI.Label

	ShowTrademarkScreen bool
	leftButtonHeld      bool
}

// CreateMainMenu creates an instance of MainMenu
func CreateMainMenu(fileProvider Common.FileProvider, sceneProvider SceneProvider, uiManager *UI.Manager, soundManager *Sound.Manager) *MainMenu {
	result := &MainMenu{
		fileProvider:        fileProvider,
		uiManager:           uiManager,
		soundManager:        soundManager,
		sceneProvider:       sceneProvider,
		ShowTrademarkScreen: true,
		leftButtonHeld:      true,
	}
	return result
}

// Load is called to load the resources for the main menu
func (v *MainMenu) Load() []func() {
	v.soundManager.PlayBGM(ResourcePaths.BGMTitle)
	return []func(){
		func() {
			v.copyrightLabel = UI.CreateLabel(v.fileProvider, ResourcePaths.FontFormal12, PaletteDefs.Static)
			v.copyrightLabel.Alignment = UI.LabelAlignCenter
			v.copyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
			v.copyrightLabel.Color = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel.MoveTo(400, 500)
		},
		func() {
			v.copyrightLabel2 = UI.CreateLabel(v.fileProvider, ResourcePaths.FontFormal12, PaletteDefs.Static)
			v.copyrightLabel2.Alignment = UI.LabelAlignCenter
			v.copyrightLabel2.SetText(Common.TranslateString("#1614"))
			v.copyrightLabel2.Color = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel2.MoveTo(400, 525)
		},
		func() {
			v.openDiabloLabel = UI.CreateLabel(v.fileProvider, ResourcePaths.FontFormal10, PaletteDefs.Static)
			v.openDiabloLabel.Alignment = UI.LabelAlignCenter
			v.openDiabloLabel.SetText("OpenDiablo2 is neither developed by, nor endorsed by Blizzard or its parent company Activision")
			v.openDiabloLabel.Color = color.RGBA{255, 255, 140, 255}
			v.openDiabloLabel.MoveTo(400, 580)
		},
		func() {
			v.background = v.fileProvider.LoadSprite(ResourcePaths.GameSelectScreen, PaletteDefs.Sky)
			v.background.MoveTo(0, 0)
		},
		func() {
			v.trademarkBackground = v.fileProvider.LoadSprite(ResourcePaths.TrademarkScreen, PaletteDefs.Sky)
			v.trademarkBackground.MoveTo(0, 0)
		},
		func() {
			v.diabloLogoLeft = v.fileProvider.LoadSprite(ResourcePaths.Diablo2LogoFireLeft, PaletteDefs.Units)
			v.diabloLogoLeft.Blend = true
			v.diabloLogoLeft.Animate = true
			v.diabloLogoLeft.MoveTo(400, 120)
		},
		func() {
			v.diabloLogoRight = v.fileProvider.LoadSprite(ResourcePaths.Diablo2LogoFireRight, PaletteDefs.Units)
			v.diabloLogoRight.Blend = true
			v.diabloLogoRight.Animate = true
			v.diabloLogoRight.MoveTo(400, 120)
		},
		func() {
			v.diabloLogoLeftBack = v.fileProvider.LoadSprite(ResourcePaths.Diablo2LogoBlackLeft, PaletteDefs.Units)
			v.diabloLogoLeftBack.MoveTo(400, 120)
		},
		func() {
			v.diabloLogoRightBack = v.fileProvider.LoadSprite(ResourcePaths.Diablo2LogoBlackRight, PaletteDefs.Units)
			v.diabloLogoRightBack.MoveTo(400, 120)
		},
		func() {
			v.exitDiabloButton = UI.CreateButton(UI.ButtonTypeWide, v.fileProvider, Common.TranslateString("#1625"))
			v.exitDiabloButton.MoveTo(264, 535)
			v.exitDiabloButton.SetVisible(!v.ShowTrademarkScreen)
			v.exitDiabloButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(v.exitDiabloButton)
		},
		func() {
			v.creditsButton = UI.CreateButton(UI.ButtonTypeShort, v.fileProvider, Common.TranslateString("#1627"))
			v.creditsButton.MoveTo(264, 505)
			v.creditsButton.SetVisible(!v.ShowTrademarkScreen)
			v.creditsButton.OnActivated(func() { v.onCreditsButtonClicked() })
			v.uiManager.AddWidget(v.creditsButton)
		},
		func() {
			v.cinematicsButton = UI.CreateButton(UI.ButtonTypeShort, v.fileProvider, Common.TranslateString("#1639"))
			v.cinematicsButton.MoveTo(401, 505)
			v.cinematicsButton.SetVisible(!v.ShowTrademarkScreen)
			v.uiManager.AddWidget(v.cinematicsButton)
		},
		func() {
			v.singlePlayerButton = UI.CreateButton(UI.ButtonTypeWide, v.fileProvider, Common.TranslateString("#1620"))
			v.singlePlayerButton.MoveTo(264, 290)
			v.singlePlayerButton.SetVisible(!v.ShowTrademarkScreen)
			v.singlePlayerButton.OnActivated(func() { v.onSinglePlayerClicked() })
			v.uiManager.AddWidget(v.singlePlayerButton)
		},
		func() {
			v.githubButton = UI.CreateButton(UI.ButtonTypeWide, v.fileProvider, "PROJECT WEBSITE")
			v.githubButton.MoveTo(264, 330)
			v.githubButton.SetVisible(!v.ShowTrademarkScreen)
			v.githubButton.OnActivated(func() { v.onGithubButtonClicked() })
			v.uiManager.AddWidget(v.githubButton)
		},
		func() {
			v.mapTestButton = UI.CreateButton(UI.ButtonTypeWide, v.fileProvider, "MAP ENGINE TEST")
			v.mapTestButton.MoveTo(264, 450)
			v.mapTestButton.SetVisible(!v.ShowTrademarkScreen)
			v.mapTestButton.OnActivated(func() { v.onMapTestClicked() })
			v.uiManager.AddWidget(v.mapTestButton)
		},
	}
}

func (v *MainMenu) onMapTestClicked() {
	v.sceneProvider.SetNextScene(CreateMapEngineTest(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager))
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func (v *MainMenu) onSinglePlayerClicked() {
	// Go here only if existing characters are available to select
	//v.sceneProvider.SetNextScene(CreateCharacterSelect(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager))
	v.sceneProvider.SetNextScene(CreateSelectHeroClass(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager))
}

func (v *MainMenu) onGithubButtonClicked() {
	openbrowser("https://www.github.com/essial/OpenDiablo2")
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
	if v.ShowTrademarkScreen {
		v.trademarkBackground.DrawSegments(screen, 4, 3, 0)
	} else {
		v.background.DrawSegments(screen, 4, 3, 0)
	}
	v.diabloLogoLeftBack.Draw(screen)
	v.diabloLogoRightBack.Draw(screen)
	v.diabloLogoLeft.Draw(screen)
	v.diabloLogoRight.Draw(screen)

	if v.ShowTrademarkScreen {
		v.copyrightLabel.Draw(screen)
		v.copyrightLabel2.Draw(screen)
	} else {
		v.openDiabloLabel.Draw(screen)
	}
}

// Update runs the update logic on the main menu
func (v *MainMenu) Update(tickTime float64) {
	if v.ShowTrademarkScreen {
		if v.uiManager.CursorButtonPressed(UI.CursorButtonLeft) {
			if v.leftButtonHeld {
				return
			}
			v.uiManager.WaitForMouseRelease()
			v.leftButtonHeld = true
			v.ShowTrademarkScreen = false
			v.exitDiabloButton.SetVisible(true)
			v.creditsButton.SetVisible(true)
			v.cinematicsButton.SetVisible(true)
			v.singlePlayerButton.SetVisible(true)
			v.githubButton.SetVisible(true)
			v.mapTestButton.SetVisible(true)
		} else {
			v.leftButtonHeld = false
		}
		return
	}
}
