package scenes

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/OpenDiablo2/OpenDiablo2/common"
	"github.com/OpenDiablo2/OpenDiablo2/palettedefs"
	"github.com/OpenDiablo2/OpenDiablo2/sound"
	"github.com/OpenDiablo2/OpenDiablo2/ui"

	"github.com/OpenDiablo2/OpenDiablo2/resourcepaths"
	"github.com/hajimehoshi/ebiten"
)

// MainMenu represents the main menu
type MainMenu struct {
	uiManager           *ui.Manager
	soundManager        *sound.Manager
	fileProvider        common.FileProvider
	sceneProvider       SceneProvider
	trademarkBackground *common.Sprite
	background          *common.Sprite
	diabloLogoLeft      *common.Sprite
	diabloLogoRight     *common.Sprite
	diabloLogoLeftBack  *common.Sprite
	diabloLogoRightBack *common.Sprite
	singlePlayerButton  *ui.Button
	githubButton        *ui.Button
	exitDiabloButton    *ui.Button
	creditsButton       *ui.Button
	cinematicsButton    *ui.Button
	mapTestButton       *ui.Button
	copyrightLabel      *ui.Label
	copyrightLabel2     *ui.Label
	openDiabloLabel     *ui.Label
	versionLabel        *ui.Label
	commitLabel         *ui.Label

	ShowTrademarkScreen bool
	leftButtonHeld      bool
}

// CreateMainMenu creates an instance of MainMenu
func CreateMainMenu(fileProvider common.FileProvider, sceneProvider SceneProvider, uiManager *ui.Manager, soundManager *sound.Manager) *MainMenu {
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
	v.soundManager.PlayBGM(resourcepaths.BGMTitle)
	return []func(){
		func() {
			v.versionLabel = ui.CreateLabel(v.fileProvider, resourcepaths.FontFormal12, palettedefs.Static)
			v.versionLabel.Alignment = ui.LabelAlignRight
			v.versionLabel.SetText("OpenDiablo2 - " + common.BuildInfo.Branch)
			v.versionLabel.Color = color.RGBA{255, 255, 255, 255}
			v.versionLabel.MoveTo(795, -10)
		},
		func() {
			v.commitLabel = ui.CreateLabel(v.fileProvider, resourcepaths.FontFormal10, palettedefs.Static)
			v.commitLabel.Alignment = ui.LabelAlignLeft
			v.commitLabel.SetText(common.BuildInfo.Commit)
			v.commitLabel.Color = color.RGBA{255, 255, 255, 255}
			v.commitLabel.MoveTo(2, 2)
		},
		func() {
			v.copyrightLabel = ui.CreateLabel(v.fileProvider, resourcepaths.FontFormal12, palettedefs.Static)
			v.copyrightLabel.Alignment = ui.LabelAlignCenter
			v.copyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
			v.copyrightLabel.Color = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel.MoveTo(400, 500)
		},
		func() {
			v.copyrightLabel2 = ui.CreateLabel(v.fileProvider, resourcepaths.FontFormal12, palettedefs.Static)
			v.copyrightLabel2.Alignment = ui.LabelAlignCenter
			v.copyrightLabel2.SetText(common.TranslateString("#1614"))
			v.copyrightLabel2.Color = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel2.MoveTo(400, 525)
		},
		func() {
			v.openDiabloLabel = ui.CreateLabel(v.fileProvider, resourcepaths.FontFormal10, palettedefs.Static)
			v.openDiabloLabel.Alignment = ui.LabelAlignCenter
			v.openDiabloLabel.SetText("OpenDiablo2 is neither developed by, nor endorsed by Blizzard or its parent company Activision")
			v.openDiabloLabel.Color = color.RGBA{255, 255, 140, 255}
			v.openDiabloLabel.MoveTo(400, 580)
		},
		func() {
			v.background = v.fileProvider.LoadSprite(resourcepaths.GameSelectScreen, palettedefs.Sky)
			v.background.MoveTo(0, 0)
		},
		func() {
			v.trademarkBackground = v.fileProvider.LoadSprite(resourcepaths.TrademarkScreen, palettedefs.Sky)
			v.trademarkBackground.MoveTo(0, 0)
		},
		func() {
			v.diabloLogoLeft = v.fileProvider.LoadSprite(resourcepaths.Diablo2LogoFireLeft, palettedefs.Units)
			v.diabloLogoLeft.Blend = true
			v.diabloLogoLeft.Animate = true
			v.diabloLogoLeft.MoveTo(400, 120)
		},
		func() {
			v.diabloLogoRight = v.fileProvider.LoadSprite(resourcepaths.Diablo2LogoFireRight, palettedefs.Units)
			v.diabloLogoRight.Blend = true
			v.diabloLogoRight.Animate = true
			v.diabloLogoRight.MoveTo(400, 120)
		},
		func() {
			v.diabloLogoLeftBack = v.fileProvider.LoadSprite(resourcepaths.Diablo2LogoBlackLeft, palettedefs.Units)
			v.diabloLogoLeftBack.MoveTo(400, 120)
		},
		func() {
			v.diabloLogoRightBack = v.fileProvider.LoadSprite(resourcepaths.Diablo2LogoBlackRight, palettedefs.Units)
			v.diabloLogoRightBack.MoveTo(400, 120)
		},
		func() {
			v.exitDiabloButton = ui.CreateButton(ui.ButtonTypeWide, v.fileProvider, common.TranslateString("#1625"))
			v.exitDiabloButton.MoveTo(264, 535)
			v.exitDiabloButton.SetVisible(!v.ShowTrademarkScreen)
			v.exitDiabloButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(v.exitDiabloButton)
		},
		func() {
			v.creditsButton = ui.CreateButton(ui.ButtonTypeShort, v.fileProvider, common.TranslateString("#1627"))
			v.creditsButton.MoveTo(264, 505)
			v.creditsButton.SetVisible(!v.ShowTrademarkScreen)
			v.creditsButton.OnActivated(func() { v.onCreditsButtonClicked() })
			v.uiManager.AddWidget(v.creditsButton)
		},
		func() {
			v.cinematicsButton = ui.CreateButton(ui.ButtonTypeShort, v.fileProvider, common.TranslateString("#1639"))
			v.cinematicsButton.MoveTo(401, 505)
			v.cinematicsButton.SetVisible(!v.ShowTrademarkScreen)
			v.uiManager.AddWidget(v.cinematicsButton)
		},
		func() {
			v.singlePlayerButton = ui.CreateButton(ui.ButtonTypeWide, v.fileProvider, common.TranslateString("#1620"))
			v.singlePlayerButton.MoveTo(264, 290)
			v.singlePlayerButton.SetVisible(!v.ShowTrademarkScreen)
			v.singlePlayerButton.OnActivated(func() { v.onSinglePlayerClicked() })
			v.uiManager.AddWidget(v.singlePlayerButton)
		},
		func() {
			v.githubButton = ui.CreateButton(ui.ButtonTypeWide, v.fileProvider, "PROJECT WEBSITE")
			v.githubButton.MoveTo(264, 330)
			v.githubButton.SetVisible(!v.ShowTrademarkScreen)
			v.githubButton.OnActivated(func() { v.onGithubButtonClicked() })
			v.uiManager.AddWidget(v.githubButton)
		},
		func() {
			v.mapTestButton = ui.CreateButton(ui.ButtonTypeWide, v.fileProvider, "MAP ENGINE TEST")
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
	openbrowser("https://www.github.com/OpenDiablo2/OpenDiablo2")
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
		v.versionLabel.Draw(screen)
		v.commitLabel.Draw(screen)
	}
}

// Update runs the update logic on the main menu
func (v *MainMenu) Update(tickTime float64) {
	if v.ShowTrademarkScreen {
		if v.uiManager.CursorButtonPressed(ui.CursorButtonLeft) {
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
