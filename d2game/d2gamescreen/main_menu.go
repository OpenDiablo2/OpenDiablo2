package d2gamescreen

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type MainMenuScreenMode int

const (
	ScreenModeUnknown MainMenuScreenMode = iota
	ScreenModeTrademark
	ScreenModeMainMenu
	ScreenModeMultiplayer
	ScreenModeTcpIp
	ScreenModeServerIp
)

// MainMenu represents the main menu
type MainMenu struct {
	tcpIpBackground     *d2ui.Sprite
	trademarkBackground *d2ui.Sprite
	background          *d2ui.Sprite
	diabloLogoLeft      *d2ui.Sprite
	diabloLogoRight     *d2ui.Sprite
	diabloLogoLeftBack  *d2ui.Sprite
	diabloLogoRightBack *d2ui.Sprite
	serverIpBackground  *d2ui.Sprite
	singlePlayerButton  d2ui.Button
	multiplayerButton   d2ui.Button
	githubButton        d2ui.Button
	exitDiabloButton    d2ui.Button
	creditsButton       d2ui.Button
	cinematicsButton    d2ui.Button
	mapTestButton       d2ui.Button
	networkTcpIpButton  d2ui.Button
	networkCancelButton d2ui.Button
	btnTcpIpCancel      d2ui.Button
	btnTcpIpHostGame    d2ui.Button
	btnTcpIpJoinGame    d2ui.Button
	btnServerIpCancel   d2ui.Button
	btnServerIpOk       d2ui.Button
	copyrightLabel      d2ui.Label
	copyrightLabel2     d2ui.Label
	openDiabloLabel     d2ui.Label
	versionLabel        d2ui.Label
	commitLabel         d2ui.Label
	tcpIpOptionsLabel   d2ui.Label
	tcpJoinGameLabel    d2ui.Label
	tcpJoinGameEntry    d2ui.TextBox
	screenMode          MainMenuScreenMode
	leftButtonHeld      bool
}

// CreateMainMenu creates an instance of MainMenu
func CreateMainMenu() *MainMenu {
	return &MainMenu{
		screenMode:     ScreenModeUnknown,
		leftButtonHeld: true,
	}
}

// Load is called to load the resources for the main menu
func (v *MainMenu) OnLoad(loading d2screen.LoadingState) {
	d2audio.PlayBGM(d2resource.BGMTitle)
	loading.Progress(0.2)

	v.versionLabel = d2ui.CreateLabel(d2resource.FontFormal12, d2resource.PaletteStatic)
	v.versionLabel.Alignment = d2ui.LabelAlignRight
	v.versionLabel.SetText("OpenDiablo2 - " + d2common.BuildInfo.Branch)
	v.versionLabel.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	v.versionLabel.SetPosition(795, -10)

	v.commitLabel = d2ui.CreateLabel(d2resource.FontFormal10, d2resource.PaletteStatic)
	v.commitLabel.Alignment = d2ui.LabelAlignLeft
	v.commitLabel.SetText(d2common.BuildInfo.Commit)
	v.commitLabel.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	v.commitLabel.SetPosition(2, 2)

	v.copyrightLabel = d2ui.CreateLabel(d2resource.FontFormal12, d2resource.PaletteStatic)
	v.copyrightLabel.Alignment = d2ui.LabelAlignCenter
	v.copyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
	v.copyrightLabel.Color = color.RGBA{R: 188, G: 168, B: 140, A: 255}
	v.copyrightLabel.SetPosition(400, 500)
	loading.Progress(0.3)

	v.copyrightLabel2 = d2ui.CreateLabel(d2resource.FontFormal12, d2resource.PaletteStatic)
	v.copyrightLabel2.Alignment = d2ui.LabelAlignCenter
	v.copyrightLabel2.SetText("All Rights Reserved.")
	v.copyrightLabel2.Color = color.RGBA{R: 188, G: 168, B: 140, A: 255}
	v.copyrightLabel2.SetPosition(400, 525)

	v.openDiabloLabel = d2ui.CreateLabel(d2resource.FontFormal10, d2resource.PaletteStatic)
	v.openDiabloLabel.Alignment = d2ui.LabelAlignCenter
	v.openDiabloLabel.SetText("OpenDiablo2 is neither developed by, nor endorsed by Blizzard or its parent company Activision")
	v.openDiabloLabel.Color = color.RGBA{R: 255, G: 255, B: 140, A: 255}
	v.openDiabloLabel.SetPosition(400, 580)
	loading.Progress(0.5)

	animation, _ := d2asset.LoadAnimation(d2resource.GameSelectScreen, d2resource.PaletteSky)
	v.background, _ = d2ui.LoadSprite(animation)
	v.background.SetPosition(0, 0)

	animation, _ = d2asset.LoadAnimation(d2resource.TrademarkScreen, d2resource.PaletteSky)
	v.trademarkBackground, _ = d2ui.LoadSprite(animation)
	v.trademarkBackground.SetPosition(0, 0)

	animation, _ = d2asset.LoadAnimation(d2resource.TcpIpBackground, d2resource.PaletteSky)
	v.tcpIpBackground, _ = d2ui.LoadSprite(animation)
	v.tcpIpBackground.SetPosition(0, 0)

	animation, _ = d2asset.LoadAnimation(d2resource.Diablo2LogoFireLeft, d2resource.PaletteUnits)
	v.diabloLogoLeft, _ = d2ui.LoadSprite(animation)
	v.diabloLogoLeft.SetBlend(true)
	v.diabloLogoLeft.PlayForward()
	v.diabloLogoLeft.SetPosition(400, 120)
	loading.Progress(0.6)

	animation, _ = d2asset.LoadAnimation(d2resource.Diablo2LogoFireRight, d2resource.PaletteUnits)
	v.diabloLogoRight, _ = d2ui.LoadSprite(animation)
	v.diabloLogoRight.SetBlend(true)
	v.diabloLogoRight.PlayForward()
	v.diabloLogoRight.SetPosition(400, 120)

	animation, _ = d2asset.LoadAnimation(d2resource.Diablo2LogoBlackLeft, d2resource.PaletteUnits)
	v.diabloLogoLeftBack, _ = d2ui.LoadSprite(animation)
	v.diabloLogoLeftBack.SetPosition(400, 120)

	animation, _ = d2asset.LoadAnimation(d2resource.Diablo2LogoBlackRight, d2resource.PaletteUnits)
	v.diabloLogoRightBack, _ = d2ui.LoadSprite(animation)
	v.diabloLogoRightBack.SetPosition(400, 120)

	v.exitDiabloButton = d2ui.CreateButton(d2ui.ButtonTypeWide, "EXIT DIABLO II")
	v.exitDiabloButton.SetPosition(264, 535)
	v.exitDiabloButton.OnActivated(func() { v.onExitButtonClicked() })
	d2ui.AddWidget(&v.exitDiabloButton)

	v.creditsButton = d2ui.CreateButton(d2ui.ButtonTypeShort, "CREDITS")
	v.creditsButton.SetPosition(264, 505)
	v.creditsButton.OnActivated(func() { v.onCreditsButtonClicked() })
	d2ui.AddWidget(&v.creditsButton)

	v.cinematicsButton = d2ui.CreateButton(d2ui.ButtonTypeShort, "CINEMATICS")
	v.cinematicsButton.SetPosition(401, 505)
	d2ui.AddWidget(&v.cinematicsButton)
	loading.Progress(0.7)

	v.singlePlayerButton = d2ui.CreateButton(d2ui.ButtonTypeWide, "SINGLE PLAYER")
	v.singlePlayerButton.SetPosition(264, 290)
	v.singlePlayerButton.OnActivated(func() { v.onSinglePlayerClicked() })
	d2ui.AddWidget(&v.singlePlayerButton)

	v.multiplayerButton = d2ui.CreateButton(d2ui.ButtonTypeWide, "MULTIPLAYER")
	v.multiplayerButton.SetPosition(264, 330)
	v.multiplayerButton.OnActivated(func() { v.onMultiplayerClicked() })
	d2ui.AddWidget(&v.multiplayerButton)

	v.githubButton = d2ui.CreateButton(d2ui.ButtonTypeWide, "PROJECT WEBSITE")
	v.githubButton.SetPosition(264, 400)
	v.githubButton.OnActivated(func() { v.onGithubButtonClicked() })
	d2ui.AddWidget(&v.githubButton)

	v.mapTestButton = d2ui.CreateButton(d2ui.ButtonTypeWide, "MAP ENGINE TEST")
	v.mapTestButton.SetPosition(264, 440)
	v.mapTestButton.OnActivated(func() { v.onMapTestClicked() })
	d2ui.AddWidget(&v.mapTestButton)

	v.networkTcpIpButton = d2ui.CreateButton(d2ui.ButtonTypeWide, "TCP/IP GAME")
	v.networkTcpIpButton.SetPosition(264, 280)
	v.networkTcpIpButton.OnActivated(func() { v.onNetworkTcpIpClicked() })
	d2ui.AddWidget(&v.networkTcpIpButton)

	v.networkCancelButton = d2ui.CreateButton(d2ui.ButtonTypeWide, d2common.TranslateString("cancel"))
	v.networkCancelButton.SetPosition(264, 540)
	v.networkCancelButton.OnActivated(func() { v.onNetworkCancelClicked() })
	d2ui.AddWidget(&v.networkCancelButton)

	v.btnTcpIpCancel = d2ui.CreateButton(d2ui.ButtonTypeMedium, d2common.TranslateString("cancel"))
	v.btnTcpIpCancel.SetPosition(33, 543)
	v.btnTcpIpCancel.OnActivated(func() { v.onTcpIpCancelClicked() })
	d2ui.AddWidget(&v.btnTcpIpCancel)

	v.btnTcpIpHostGame = d2ui.CreateButton(d2ui.ButtonTypeWide, "HOST GAME")
	v.btnTcpIpHostGame.SetPosition(264, 280)
	v.btnTcpIpHostGame.OnActivated(func() { v.onTcpIpHostGameClicked() })
	d2ui.AddWidget(&v.btnTcpIpHostGame)

	v.btnTcpIpJoinGame = d2ui.CreateButton(d2ui.ButtonTypeWide, "JOIN GAME")
	v.btnTcpIpJoinGame.SetPosition(264, 320)
	v.btnTcpIpJoinGame.OnActivated(func() { v.onTcpIpJoinGameClicked() })
	d2ui.AddWidget(&v.btnTcpIpJoinGame)
	loading.Progress(0.8)

	v.tcpIpOptionsLabel = d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteUnits)
	v.tcpIpOptionsLabel.SetPosition(400, 23)
	v.tcpIpOptionsLabel.Alignment = d2ui.LabelAlignCenter
	v.tcpIpOptionsLabel.SetText("TCP/IP Options")

	animation, _ = d2asset.LoadAnimation(d2resource.PopUpOkCancel, d2resource.PaletteFechar)
	v.serverIpBackground, _ = d2ui.LoadSprite(animation)
	v.serverIpBackground.SetPosition(270, 175)

	v.tcpJoinGameLabel = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.tcpJoinGameLabel.Alignment = d2ui.LabelAlignCenter
	v.tcpJoinGameLabel.SetText(d2common.CombineStrings(d2common.
		SplitIntoLinesWithMaxWidth("Enter Host IP Address to Join Game", 23)))
	v.tcpJoinGameLabel.Color = color.RGBA{R: 216, G: 196, B: 128, A: 255}
	v.tcpJoinGameLabel.SetPosition(400, 190)

	v.tcpJoinGameEntry = d2ui.CreateTextbox()
	v.tcpJoinGameEntry.SetPosition(318, 245)
	v.tcpJoinGameEntry.SetFilter("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890._:")
	d2ui.AddWidget(&v.tcpJoinGameEntry)
	loading.Progress(0.9)

	v.btnServerIpCancel = d2ui.CreateButton(d2ui.ButtonTypeOkCancel, "CANCEL")
	v.btnServerIpCancel.SetPosition(285, 305)
	v.btnServerIpCancel.OnActivated(func() { v.onBtnTcpIpCancelClicked() })
	d2ui.AddWidget(&v.btnServerIpCancel)

	v.btnServerIpOk = d2ui.CreateButton(d2ui.ButtonTypeOkCancel, "OK")
	v.btnServerIpOk.SetPosition(420, 305)
	v.btnServerIpOk.OnActivated(func() { v.onBtnTcpIpOkClicked() })
	d2ui.AddWidget(&v.btnServerIpOk)

	if v.screenMode == ScreenModeUnknown {
		v.SetScreenMode(ScreenModeTrademark)
	} else {
		v.SetScreenMode(ScreenModeMainMenu)
	}

	d2input.BindHandler(v)
}

func (v *MainMenu) onMapTestClicked() {
	d2screen.SetNextScreen(CreateMapEngineTest(0, 1))
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
	if d2player.HasGameStates() {
		d2screen.SetNextScreen(CreateCharacterSelect(d2clientconnectiontype.Local, v.tcpJoinGameEntry.GetText()))
		return
	}
	d2screen.SetNextScreen(CreateSelectHeroClass(d2clientconnectiontype.Local, v.tcpJoinGameEntry.GetText()))
}

func (v *MainMenu) onGithubButtonClicked() {
	openbrowser("https://www.github.com/OpenDiablo2/OpenDiablo2")
}

func (v *MainMenu) onExitButtonClicked() {
	os.Exit(0)
}

func (v *MainMenu) onCreditsButtonClicked() {
	d2screen.SetNextScreen(CreateCredits())
}

// Render renders the main menu
func (v *MainMenu) Render(screen d2render.Surface) error {
	switch v.screenMode {
	case ScreenModeTrademark:
		v.trademarkBackground.RenderSegmented(screen, 4, 3, 0)
	case ScreenModeServerIp:
		fallthrough
	case ScreenModeTcpIp:
		v.tcpIpBackground.RenderSegmented(screen, 4, 3, 0)
	default:
		v.background.RenderSegmented(screen, 4, 3, 0)
	}

	switch v.screenMode {
	case ScreenModeTrademark:
		fallthrough
	case ScreenModeMainMenu:
		fallthrough
	case ScreenModeMultiplayer:
		v.diabloLogoLeftBack.Render(screen)
		v.diabloLogoRightBack.Render(screen)
		v.diabloLogoLeft.Render(screen)
		v.diabloLogoRight.Render(screen)
	}

	switch v.screenMode {
	case ScreenModeServerIp:
		v.tcpIpOptionsLabel.Render(screen)
		v.serverIpBackground.RenderSegmented(screen, 2, 1, 0)
		v.tcpJoinGameLabel.Render(screen)
	case ScreenModeTcpIp:
		v.tcpIpOptionsLabel.Render(screen)
	case ScreenModeTrademark:
		v.copyrightLabel.Render(screen)
		v.copyrightLabel2.Render(screen)
	case ScreenModeMainMenu:
		v.openDiabloLabel.Render(screen)
		v.versionLabel.Render(screen)
		v.commitLabel.Render(screen)
	}

	return nil
}

// Update runs the update logic on the main menu
func (v *MainMenu) Advance(tickTime float64) error {
	switch v.screenMode {
	case ScreenModeMainMenu:
		fallthrough
	case ScreenModeTrademark:
		fallthrough
	case ScreenModeMultiplayer:
		v.diabloLogoLeftBack.Advance(tickTime)
		v.diabloLogoRightBack.Advance(tickTime)
		v.diabloLogoLeft.Advance(tickTime)
		v.diabloLogoRight.Advance(tickTime)
	}

	return nil
}

func (v *MainMenu) OnMouseButtonDown(event d2input.MouseEvent) bool {
	if v.screenMode == ScreenModeTrademark && event.Button == d2input.MouseButtonLeft {
		v.SetScreenMode(ScreenModeMainMenu)
		return true
	}
	return false
}

func (v *MainMenu) SetScreenMode(screenMode MainMenuScreenMode) {
	v.screenMode = screenMode
	isMainMenu := screenMode == ScreenModeMainMenu
	isMultiplayer := screenMode == ScreenModeMultiplayer
	isTcpIp := screenMode == ScreenModeTcpIp
	isServerIp := screenMode == ScreenModeServerIp
	v.exitDiabloButton.SetVisible(isMainMenu)
	v.creditsButton.SetVisible(isMainMenu)
	v.cinematicsButton.SetVisible(isMainMenu)
	v.singlePlayerButton.SetVisible(isMainMenu)
	v.githubButton.SetVisible(isMainMenu)
	v.mapTestButton.SetVisible(isMainMenu)
	v.multiplayerButton.SetVisible(isMainMenu)
	v.networkTcpIpButton.SetVisible(isMultiplayer)
	v.networkCancelButton.SetVisible(isMultiplayer)
	v.btnTcpIpCancel.SetVisible(isTcpIp)
	v.btnTcpIpHostGame.SetVisible(isTcpIp)
	v.btnTcpIpJoinGame.SetVisible(isTcpIp)
	v.tcpJoinGameEntry.SetVisible(isServerIp)
	if isServerIp {
		v.tcpJoinGameEntry.Activate()
	}
	v.btnServerIpOk.SetVisible(isServerIp)
	v.btnServerIpCancel.SetVisible(isServerIp)
}

func (v *MainMenu) onNetworkCancelClicked() {
	v.SetScreenMode(ScreenModeMainMenu)
}

func (v *MainMenu) onMultiplayerClicked() {
	v.SetScreenMode(ScreenModeMultiplayer)
}

func (v *MainMenu) onNetworkTcpIpClicked() {
	v.SetScreenMode(ScreenModeTcpIp)
}

func (v *MainMenu) onTcpIpCancelClicked() {
	v.SetScreenMode(ScreenModeMultiplayer)
}

func (v *MainMenu) onTcpIpHostGameClicked() {
	d2screen.SetNextScreen(CreateCharacterSelect(d2clientconnectiontype.LANServer, ""))
}

func (v *MainMenu) onTcpIpJoinGameClicked() {
	v.SetScreenMode(ScreenModeServerIp)
}

func (v *MainMenu) onBtnTcpIpCancelClicked() {
	v.SetScreenMode(ScreenModeTcpIp)
}

func (v *MainMenu) onBtnTcpIpOkClicked() {
	d2screen.SetNextScreen(CreateCharacterSelect(d2clientconnectiontype.LANClient, v.tcpJoinGameEntry.GetText()))
}
