package d2gamescene

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2scene"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// MainMenu represents the main menu
type MainMenu struct {
	trademarkBackground *d2ui.Sprite
	background          *d2ui.Sprite
	diabloLogoLeft      *d2ui.Sprite
	diabloLogoRight     *d2ui.Sprite
	diabloLogoLeftBack  *d2ui.Sprite
	diabloLogoRightBack *d2ui.Sprite
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
func CreateMainMenu() *MainMenu {
	result := &MainMenu{
		ShowTrademarkScreen: true,
		leftButtonHeld:      true,
	}
	return result
}

// Load is called to load the resources for the main menu
func (v *MainMenu) Load() []func() {
	d2audio.PlayBGM(d2resource.BGMTitle)
	return []func(){
		func() {
			v.versionLabel = d2ui.CreateLabel(d2resource.FontFormal12, d2resource.PaletteStatic)
			v.versionLabel.Alignment = d2ui.LabelAlignRight
			v.versionLabel.SetText("OpenDiablo2 - " + d2common.BuildInfo.Branch)
			v.versionLabel.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			v.versionLabel.SetPosition(795, -10)
		},
		func() {
			v.commitLabel = d2ui.CreateLabel(d2resource.FontFormal10, d2resource.PaletteStatic)
			v.commitLabel.Alignment = d2ui.LabelAlignLeft
			v.commitLabel.SetText(d2common.BuildInfo.Commit)
			v.commitLabel.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			v.commitLabel.SetPosition(2, 2)
		},
		func() {
			v.copyrightLabel = d2ui.CreateLabel(d2resource.FontFormal12, d2resource.PaletteStatic)
			v.copyrightLabel.Alignment = d2ui.LabelAlignCenter
			v.copyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
			v.copyrightLabel.Color = color.RGBA{R: 188, G: 168, B: 140, A: 255}
			v.copyrightLabel.SetPosition(400, 500)
		},
		func() {
			v.copyrightLabel2 = d2ui.CreateLabel(d2resource.FontFormal12, d2resource.PaletteStatic)
			v.copyrightLabel2.Alignment = d2ui.LabelAlignCenter
			v.copyrightLabel2.SetText(d2common.TranslateString("#1614"))
			v.copyrightLabel2.Color = color.RGBA{R: 188, G: 168, B: 140, A: 255}
			v.copyrightLabel2.SetPosition(400, 525)
		},
		func() {
			v.openDiabloLabel = d2ui.CreateLabel(d2resource.FontFormal10, d2resource.PaletteStatic)
			v.openDiabloLabel.Alignment = d2ui.LabelAlignCenter
			v.openDiabloLabel.SetText("OpenDiablo2 is neither developed by, nor endorsed by Blizzard or its parent company Activision")
			v.openDiabloLabel.Color = color.RGBA{R: 255, G: 255, B: 140, A: 255}
			v.openDiabloLabel.SetPosition(400, 580)
		},
		func() {
			animation, _ := d2asset.LoadAnimation(d2resource.GameSelectScreen, d2resource.PaletteSky)
			v.background, _ = d2ui.LoadSprite(animation)
			v.background.SetPosition(0, 0)
		},
		func() {
			animation, _ := d2asset.LoadAnimation(d2resource.TrademarkScreen, d2resource.PaletteSky)
			v.trademarkBackground, _ = d2ui.LoadSprite(animation)
			v.trademarkBackground.SetPosition(0, 0)
		},
		func() {
			animation, _ := d2asset.LoadAnimation(d2resource.Diablo2LogoFireLeft, d2resource.PaletteUnits)
			v.diabloLogoLeft, _ = d2ui.LoadSprite(animation)
			v.diabloLogoLeft.SetBlend(true)
			v.diabloLogoLeft.PlayForward()
			v.diabloLogoLeft.SetPosition(400, 120)
		},
		func() {
			animation, _ := d2asset.LoadAnimation(d2resource.Diablo2LogoFireRight, d2resource.PaletteUnits)
			v.diabloLogoRight, _ = d2ui.LoadSprite(animation)
			v.diabloLogoRight.SetBlend(true)
			v.diabloLogoRight.PlayForward()
			v.diabloLogoRight.SetPosition(400, 120)
		},
		func() {
			animation, _ := d2asset.LoadAnimation(d2resource.Diablo2LogoBlackLeft, d2resource.PaletteUnits)
			v.diabloLogoLeftBack, _ = d2ui.LoadSprite(animation)
			v.diabloLogoLeftBack.SetPosition(400, 120)
		},
		func() {
			animation, _ := d2asset.LoadAnimation(d2resource.Diablo2LogoBlackRight, d2resource.PaletteUnits)
			v.diabloLogoRightBack, _ = d2ui.LoadSprite(animation)
			v.diabloLogoRightBack.SetPosition(400, 120)
		},
		func() {
			v.exitDiabloButton = d2ui.CreateButton(d2ui.ButtonTypeWide, d2common.TranslateString("#1625"))
			v.exitDiabloButton.SetPosition(264, 535)
			v.exitDiabloButton.SetVisible(!v.ShowTrademarkScreen)
			v.exitDiabloButton.OnActivated(func() { v.onExitButtonClicked() })
			d2ui.AddWidget(&v.exitDiabloButton)
		},
		func() {
			v.creditsButton = d2ui.CreateButton(d2ui.ButtonTypeShort, d2common.TranslateString("#1627"))
			v.creditsButton.SetPosition(264, 505)
			v.creditsButton.SetVisible(!v.ShowTrademarkScreen)
			v.creditsButton.OnActivated(func() { v.onCreditsButtonClicked() })
			d2ui.AddWidget(&v.creditsButton)
		},
		func() {
			v.cinematicsButton = d2ui.CreateButton(d2ui.ButtonTypeShort, d2common.TranslateString("#1639"))
			v.cinematicsButton.SetPosition(401, 505)
			v.cinematicsButton.SetVisible(!v.ShowTrademarkScreen)
			d2ui.AddWidget(&v.cinematicsButton)
		},
		func() {
			v.singlePlayerButton = d2ui.CreateButton(d2ui.ButtonTypeWide, d2common.TranslateString("#1620"))
			v.singlePlayerButton.SetPosition(264, 290)
			v.singlePlayerButton.SetVisible(!v.ShowTrademarkScreen)
			v.singlePlayerButton.OnActivated(func() { v.onSinglePlayerClicked() })
			d2ui.AddWidget(&v.singlePlayerButton)
		},
		func() {
			v.githubButton = d2ui.CreateButton(d2ui.ButtonTypeWide, "PROJECT WEBSITE")
			v.githubButton.SetPosition(264, 330)
			v.githubButton.SetVisible(!v.ShowTrademarkScreen)
			v.githubButton.OnActivated(func() { v.onGithubButtonClicked() })
			d2ui.AddWidget(&v.githubButton)
		},
		func() {
			v.mapTestButton = d2ui.CreateButton(d2ui.ButtonTypeWide, "MAP ENGINE TEST")
			v.mapTestButton.SetPosition(264, 450)
			v.mapTestButton.SetVisible(!v.ShowTrademarkScreen)
			v.mapTestButton.OnActivated(func() { v.onMapTestClicked() })
			d2ui.AddWidget(&v.mapTestButton)
		},
	}
}

func (v *MainMenu) onMapTestClicked() {
	d2scene.SetNextScene(CreateMapEngineTest(0, 1))
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
	if d2gamestate.HasGameStates() {
		d2scene.SetNextScene(CreateCharacterSelect())
		return
	}
	d2scene.SetNextScene(CreateSelectHeroClass())
}

func (v *MainMenu) onGithubButtonClicked() {
	openbrowser("https://www.github.com/OpenDiablo2/OpenDiablo2")
}

func (v *MainMenu) onExitButtonClicked() {
	os.Exit(0)
}

func (v *MainMenu) onCreditsButtonClicked() {
	d2scene.SetNextScene(CreateCredits())
}

// Unload unloads the data for the main menu
func (v *MainMenu) Unload() {

}

// Render renders the main menu
func (v *MainMenu) Render(screen d2render.Surface) {
	if v.ShowTrademarkScreen {
		v.trademarkBackground.RenderSegmented(screen, 4, 3, 0)
	} else {
		v.background.RenderSegmented(screen, 4, 3, 0)
	}
	v.diabloLogoLeftBack.Render(screen)
	v.diabloLogoRightBack.Render(screen)
	v.diabloLogoLeft.Render(screen)
	v.diabloLogoRight.Render(screen)

	if v.ShowTrademarkScreen {
		v.copyrightLabel.Render(screen)
		v.copyrightLabel2.Render(screen)
	} else {
		v.openDiabloLabel.Render(screen)
		v.versionLabel.Render(screen)
		v.commitLabel.Render(screen)
	}
}

// Update runs the update logic on the main menu
func (v *MainMenu) Advance(tickTime float64) {
	v.diabloLogoLeftBack.Advance(tickTime)
	v.diabloLogoRightBack.Advance(tickTime)
	v.diabloLogoLeft.Advance(tickTime)
	v.diabloLogoRight.Advance(tickTime)

	if v.ShowTrademarkScreen {
		if d2ui.CursorButtonPressed(d2ui.CursorButtonLeft) {
			if v.leftButtonHeld {
				return
			}
			d2ui.WaitForMouseRelease()
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
