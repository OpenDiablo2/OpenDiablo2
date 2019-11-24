package d2scene

import (
	"fmt"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"
	"image/color"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"

	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2render"

	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon/d2coreinterface"

	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/D2Shared/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"

	"github.com/hajimehoshi/ebiten"
)

// MainMenu represents the main menu
type MainMenu struct {
	uiManager           *d2ui.Manager
	soundManager        *d2audio.Manager
	fileProvider        d2interface.FileProvider
	sceneProvider       d2coreinterface.SceneProvider
	trademarkBackground d2render.Sprite
	background          d2render.Sprite
	diabloLogoLeft      d2render.Sprite
	diabloLogoRight     d2render.Sprite
	diabloLogoLeftBack  d2render.Sprite
	diabloLogoRightBack d2render.Sprite
	singlePlayerButton  d2ui.Button
	githubButton        d2ui.Button
	exitDiabloButton    d2ui.Button
	creditsButton       d2ui.Button
	cinematicsButton    d2ui.Button
	mapTestButton       d2ui.Button
	copyrightLabel      d2ui.Label
	copyrightLabel2     d2ui.Label
	openDiabloLabel     d2ui.Label
	versionLabel        d2ui.Label
	commitLabel         d2ui.Label

	ShowTrademarkScreen bool
	leftButtonHeld      bool
}

// CreateMainMenu creates an instance of MainMenu
func CreateMainMenu(fileProvider d2interface.FileProvider, sceneProvider d2coreinterface.SceneProvider, uiManager *d2ui.Manager, soundManager *d2audio.Manager) *MainMenu {
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
	v.soundManager.PlayBGM(d2resource.BGMTitle)
	return []func(){
		func() {
			v.versionLabel = d2ui.CreateLabel(v.fileProvider, d2resource.FontFormal12, d2enum.Static)
			v.versionLabel.Alignment = d2ui.LabelAlignRight
			v.versionLabel.SetText("OpenDiablo2 - " + d2common.BuildInfo.Branch)
			v.versionLabel.Color = color.RGBA{255, 255, 255, 255}
			v.versionLabel.MoveTo(795, -10)
		},
		func() {
			v.commitLabel = d2ui.CreateLabel(v.fileProvider, d2resource.FontFormal10, d2enum.Static)
			v.commitLabel.Alignment = d2ui.LabelAlignLeft
			v.commitLabel.SetText(d2common.BuildInfo.Commit)
			v.commitLabel.Color = color.RGBA{255, 255, 255, 255}
			v.commitLabel.MoveTo(2, 2)
		},
		func() {
			v.copyrightLabel = d2ui.CreateLabel(v.fileProvider, d2resource.FontFormal12, d2enum.Static)
			v.copyrightLabel.Alignment = d2ui.LabelAlignCenter
			v.copyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
			v.copyrightLabel.Color = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel.MoveTo(400, 500)
		},
		func() {
			v.copyrightLabel2 = d2ui.CreateLabel(v.fileProvider, d2resource.FontFormal12, d2enum.Static)
			v.copyrightLabel2.Alignment = d2ui.LabelAlignCenter
			v.copyrightLabel2.SetText(d2common.TranslateString("#1614"))
			v.copyrightLabel2.Color = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel2.MoveTo(400, 525)
		},
		func() {
			v.openDiabloLabel = d2ui.CreateLabel(v.fileProvider, d2resource.FontFormal10, d2enum.Static)
			v.openDiabloLabel.Alignment = d2ui.LabelAlignCenter
			v.openDiabloLabel.SetText("OpenDiablo2 is neither developed by, nor endorsed by Blizzard or its parent company Activision")
			v.openDiabloLabel.Color = color.RGBA{255, 255, 140, 255}
			v.openDiabloLabel.MoveTo(400, 580)
		},
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.GameSelectScreen), d2datadict.Palettes[d2enum.Sky])
			v.background = d2render.CreateSpriteFromDC6(dc6)
			v.background.MoveTo(0, 0)
		},
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.TrademarkScreen), d2datadict.Palettes[d2enum.Sky])
			v.trademarkBackground = d2render.CreateSpriteFromDC6(dc6)
			v.trademarkBackground.MoveTo(0, 0)
		},
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.Diablo2LogoFireLeft), d2datadict.Palettes[d2enum.Units])
			v.diabloLogoLeft = d2render.CreateSpriteFromDC6(dc6)
			v.diabloLogoLeft.Blend = true
			v.diabloLogoLeft.Animate = true
			v.diabloLogoLeft.MoveTo(400, 120)
		},
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.Diablo2LogoFireRight), d2datadict.Palettes[d2enum.Units])
			v.diabloLogoRight = d2render.CreateSpriteFromDC6(dc6)
			v.diabloLogoRight.Blend = true
			v.diabloLogoRight.Animate = true
			v.diabloLogoRight.MoveTo(400, 120)
		},
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.Diablo2LogoBlackLeft), d2datadict.Palettes[d2enum.Units])
			v.diabloLogoLeftBack = d2render.CreateSpriteFromDC6(dc6)
			v.diabloLogoLeftBack.MoveTo(400, 120)
		},
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.Diablo2LogoBlackRight), d2datadict.Palettes[d2enum.Units])
			v.diabloLogoRightBack = d2render.CreateSpriteFromDC6(dc6)
			v.diabloLogoRightBack.MoveTo(400, 120)
		},
		func() {
			v.exitDiabloButton = d2ui.CreateButton(d2ui.ButtonTypeWide, v.fileProvider, d2common.TranslateString("#1625"))
			v.exitDiabloButton.MoveTo(264, 535)
			v.exitDiabloButton.SetVisible(!v.ShowTrademarkScreen)
			v.exitDiabloButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(&v.exitDiabloButton)
		},
		func() {
			v.creditsButton = d2ui.CreateButton(d2ui.ButtonTypeShort, v.fileProvider, d2common.TranslateString("#1627"))
			v.creditsButton.MoveTo(264, 505)
			v.creditsButton.SetVisible(!v.ShowTrademarkScreen)
			v.creditsButton.OnActivated(func() { v.onCreditsButtonClicked() })
			v.uiManager.AddWidget(&v.creditsButton)
		},
		func() {
			v.cinematicsButton = d2ui.CreateButton(d2ui.ButtonTypeShort, v.fileProvider, d2common.TranslateString("#1639"))
			v.cinematicsButton.MoveTo(401, 505)
			v.cinematicsButton.SetVisible(!v.ShowTrademarkScreen)
			v.uiManager.AddWidget(&v.cinematicsButton)
		},
		func() {
			v.singlePlayerButton = d2ui.CreateButton(d2ui.ButtonTypeWide, v.fileProvider, d2common.TranslateString("#1620"))
			v.singlePlayerButton.MoveTo(264, 290)
			v.singlePlayerButton.SetVisible(!v.ShowTrademarkScreen)
			v.singlePlayerButton.OnActivated(func() { v.onSinglePlayerClicked() })
			v.uiManager.AddWidget(&v.singlePlayerButton)
		},
		func() {
			v.githubButton = d2ui.CreateButton(d2ui.ButtonTypeWide, v.fileProvider, "PROJECT WEBSITE")
			v.githubButton.MoveTo(264, 330)
			v.githubButton.SetVisible(!v.ShowTrademarkScreen)
			v.githubButton.OnActivated(func() { v.onGithubButtonClicked() })
			v.uiManager.AddWidget(&v.githubButton)
		},
		func() {
			v.mapTestButton = d2ui.CreateButton(d2ui.ButtonTypeWide, v.fileProvider, "MAP ENGINE TEST")
			v.mapTestButton.MoveTo(264, 450)
			v.mapTestButton.SetVisible(!v.ShowTrademarkScreen)
			v.mapTestButton.OnActivated(func() { v.onMapTestClicked() })
			v.uiManager.AddWidget(&v.mapTestButton)
		},
	}
}

func (v *MainMenu) onMapTestClicked() {
	v.sceneProvider.SetNextScene(CreateMapEngineTest(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager, 0, 1))
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
	if d2core.HasGameStates() {
		v.sceneProvider.SetNextScene(CreateCharacterSelect(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager))
		return
	}
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
		if v.uiManager.CursorButtonPressed(d2ui.CursorButtonLeft) {
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
