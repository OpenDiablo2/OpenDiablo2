package d2gamescreen

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type mainMenuScreenMode int

const (
	screenModeUnknown mainMenuScreenMode = iota
	screenModeTrademark
	screenModeMainMenu
	screenModeMultiplayer
	screenModeTCPIP
	screenModeServerIP
)

// MainMenu represents the main menu
type MainMenu struct {
	tcpIPBackground     *d2ui.Sprite
	trademarkBackground *d2ui.Sprite
	background          *d2ui.Sprite
	diabloLogoLeft      *d2ui.Sprite
	diabloLogoRight     *d2ui.Sprite
	diabloLogoLeftBack  *d2ui.Sprite
	diabloLogoRightBack *d2ui.Sprite
	serverIPBackground  *d2ui.Sprite
	singlePlayerButton  d2ui.Button
	multiplayerButton   d2ui.Button
	githubButton        d2ui.Button
	exitDiabloButton    d2ui.Button
	creditsButton       d2ui.Button
	cinematicsButton    d2ui.Button
	mapTestButton       d2ui.Button
	networkTCPIPButton  d2ui.Button
	networkCancelButton d2ui.Button
	btnTCPIPCancel      d2ui.Button
	btnTCPIPHostGame    d2ui.Button
	btnTCPIPJoinGame    d2ui.Button
	btnServerIPCancel   d2ui.Button
	btnServerIPOk       d2ui.Button
	copyrightLabel      d2ui.Label
	copyrightLabel2     d2ui.Label
	openDiabloLabel     d2ui.Label
	versionLabel        d2ui.Label
	commitLabel         d2ui.Label
	tcpIPOptionsLabel   d2ui.Label
	tcpJoinGameLabel    d2ui.Label
	tcpJoinGameEntry    d2ui.TextBox
	screenMode          mainMenuScreenMode
	leftButtonHeld      bool
	renderer            d2interface.Renderer
	audioProvider       d2interface.AudioProvider
	terminal            d2interface.Terminal
}

// CreateMainMenu creates an instance of MainMenu
func CreateMainMenu(renderer d2interface.Renderer, audioProvider d2interface.AudioProvider, term d2interface.Terminal) *MainMenu {
	return &MainMenu{
		screenMode:     screenModeUnknown,
		leftButtonHeld: true,
		renderer:       renderer,
		audioProvider:  audioProvider,
		terminal:       term,
	}
}

// OnLoad is called to load the resources for the main menu
func (v *MainMenu) OnLoad(loading d2screen.LoadingState) {
	v.audioProvider.PlayBGM(d2resource.BGMTitle)
	loading.Progress(0.2)

	v.createLabels(loading)
	v.loadBackgroundSprites()
	v.createLogos(loading)
	v.createButtons(loading)

	v.tcpJoinGameEntry = d2ui.CreateTextbox(v.renderer)
	v.tcpJoinGameEntry.SetPosition(318, 245)
	v.tcpJoinGameEntry.SetFilter("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890._:")
	d2ui.AddWidget(&v.tcpJoinGameEntry)
	loading.Progress(0.9)

	if v.screenMode == screenModeUnknown {
		v.setScreenMode(screenModeTrademark)
	} else {
		v.setScreenMode(screenModeMainMenu)
	}

	if err := d2input.BindHandler(v); err != nil {
		fmt.Println("failed to add main menu as event handler")
	}
}

func (v *MainMenu) loadBackgroundSprites() {
	animation, _ := d2asset.LoadAnimation(d2resource.GameSelectScreen, d2resource.PaletteSky)
	v.background, _ = d2ui.LoadSprite(animation)
	v.background.SetPosition(0, 0)

	animation, _ = d2asset.LoadAnimation(d2resource.TrademarkScreen, d2resource.PaletteSky)
	v.trademarkBackground, _ = d2ui.LoadSprite(animation)
	v.trademarkBackground.SetPosition(0, 0)

	animation, _ = d2asset.LoadAnimation(d2resource.TCPIPBackground, d2resource.PaletteSky)
	v.tcpIPBackground, _ = d2ui.LoadSprite(animation)
	v.tcpIPBackground.SetPosition(0, 0)

	animation, _ = d2asset.LoadAnimation(d2resource.PopUpOkCancel, d2resource.PaletteFechar)
	v.serverIPBackground, _ = d2ui.LoadSprite(animation)
	v.serverIPBackground.SetPosition(270, 175)
}

func (v *MainMenu) createLabels(loading d2screen.LoadingState) {
	v.versionLabel = d2ui.CreateLabel(v.renderer, d2resource.FontFormal12, d2resource.PaletteStatic)
	v.versionLabel.Alignment = d2ui.LabelAlignRight
	v.versionLabel.SetText("OpenDiablo2 - " + d2common.BuildInfo.Branch)
	v.versionLabel.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	v.versionLabel.SetPosition(795, -10)

	v.commitLabel = d2ui.CreateLabel(v.renderer, d2resource.FontFormal10, d2resource.PaletteStatic)
	v.commitLabel.Alignment = d2ui.LabelAlignLeft
	v.commitLabel.SetText(d2common.BuildInfo.Commit)
	v.commitLabel.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	v.commitLabel.SetPosition(2, 2)

	v.copyrightLabel = d2ui.CreateLabel(v.renderer, d2resource.FontFormal12, d2resource.PaletteStatic)
	v.copyrightLabel.Alignment = d2ui.LabelAlignCenter
	v.copyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
	v.copyrightLabel.Color = color.RGBA{R: 188, G: 168, B: 140, A: 255}
	v.copyrightLabel.SetPosition(400, 500)
	loading.Progress(0.3)

	v.copyrightLabel2 = d2ui.CreateLabel(v.renderer, d2resource.FontFormal12, d2resource.PaletteStatic)
	v.copyrightLabel2.Alignment = d2ui.LabelAlignCenter
	v.copyrightLabel2.SetText("All Rights Reserved.")
	v.copyrightLabel2.Color = color.RGBA{R: 188, G: 168, B: 140, A: 255}
	v.copyrightLabel2.SetPosition(400, 525)

	v.openDiabloLabel = d2ui.CreateLabel(v.renderer, d2resource.FontFormal10, d2resource.PaletteStatic)
	v.openDiabloLabel.Alignment = d2ui.LabelAlignCenter
	v.openDiabloLabel.SetText("OpenDiablo2 is neither developed by, nor endorsed by Blizzard or its parent company Activision")
	v.openDiabloLabel.Color = color.RGBA{R: 255, G: 255, B: 140, A: 255}
	v.openDiabloLabel.SetPosition(400, 580)
	loading.Progress(0.5)

	v.tcpIPOptionsLabel = d2ui.CreateLabel(v.renderer, d2resource.Font42, d2resource.PaletteUnits)
	v.tcpIPOptionsLabel.SetPosition(400, 23)
	v.tcpIPOptionsLabel.Alignment = d2ui.LabelAlignCenter
	v.tcpIPOptionsLabel.SetText("TCP/IP Options")

	v.tcpJoinGameLabel = d2ui.CreateLabel(v.renderer, d2resource.Font16, d2resource.PaletteUnits)
	v.tcpJoinGameLabel.Alignment = d2ui.LabelAlignCenter
	v.tcpJoinGameLabel.SetText(d2common.CombineStrings(
		d2common.SplitIntoLinesWithMaxWidth("Enter Host IP Address to Join Game", 23)))

	v.tcpJoinGameLabel.Color = color.RGBA{R: 216, G: 196, B: 128, A: 255}
	v.tcpJoinGameLabel.SetPosition(400, 190)
}

func (v *MainMenu) createLogos(loading d2screen.LoadingState) {
	animation, _ := d2asset.LoadAnimation(d2resource.Diablo2LogoFireLeft, d2resource.PaletteUnits)
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
}

func (v *MainMenu) createButtons(loading d2screen.LoadingState) {
	v.exitDiabloButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "EXIT DIABLO II")
	v.exitDiabloButton.SetPosition(264, 535)
	v.exitDiabloButton.OnActivated(func() { v.onExitButtonClicked() })
	d2ui.AddWidget(&v.exitDiabloButton)

	v.creditsButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeShort, "CREDITS")
	v.creditsButton.SetPosition(264, 505)
	v.creditsButton.OnActivated(func() { v.onCreditsButtonClicked() })
	d2ui.AddWidget(&v.creditsButton)

	v.cinematicsButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeShort, "CINEMATICS")
	v.cinematicsButton.SetPosition(401, 505)
	d2ui.AddWidget(&v.cinematicsButton)
	loading.Progress(0.7)

	v.singlePlayerButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "SINGLE PLAYER")
	v.singlePlayerButton.SetPosition(264, 290)
	v.singlePlayerButton.OnActivated(func() { v.onSinglePlayerClicked() })
	d2ui.AddWidget(&v.singlePlayerButton)

	v.githubButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "PROJECT WEBSITE")
	v.githubButton.SetPosition(264, 400)
	v.githubButton.OnActivated(func() { v.onGithubButtonClicked() })
	d2ui.AddWidget(&v.githubButton)

	v.mapTestButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "MAP ENGINE TEST")
	v.mapTestButton.SetPosition(264, 440)
	v.mapTestButton.OnActivated(func() { v.onMapTestClicked() })
	d2ui.AddWidget(&v.mapTestButton)

	v.btnTCPIPCancel = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeMedium, d2common.TranslateString("cancel"))
	v.btnTCPIPCancel.SetPosition(33, 543)
	v.btnTCPIPCancel.OnActivated(func() { v.onTCPIPCancelClicked() })
	d2ui.AddWidget(&v.btnTCPIPCancel)

	v.btnServerIPCancel = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeOkCancel, "CANCEL")
	v.btnServerIPCancel.SetPosition(285, 305)
	v.btnServerIPCancel.OnActivated(func() { v.onBtnTCPIPCancelClicked() })
	d2ui.AddWidget(&v.btnServerIPCancel)

	v.btnServerIPOk = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeOkCancel, "OK")
	v.btnServerIPOk.SetPosition(420, 305)
	v.btnServerIPOk.OnActivated(func() { v.onBtnTCPIPOkClicked() })
	d2ui.AddWidget(&v.btnServerIPOk)

	v.createMultiplayerMenuButtons()
	loading.Progress(0.8)
}

func (v *MainMenu) createMultiplayerMenuButtons() {
	v.multiplayerButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "MULTIPLAYER")
	v.multiplayerButton.SetPosition(264, 330)
	v.multiplayerButton.OnActivated(func() { v.onMultiplayerClicked() })
	d2ui.AddWidget(&v.multiplayerButton)

	v.networkTCPIPButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "TCP/IP GAME")
	v.networkTCPIPButton.SetPosition(264, 280)
	v.networkTCPIPButton.OnActivated(func() { v.onNetworkTCPIPClicked() })
	d2ui.AddWidget(&v.networkTCPIPButton)

	v.networkCancelButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, d2common.TranslateString("cancel"))
	v.networkCancelButton.SetPosition(264, 540)
	v.networkCancelButton.OnActivated(func() { v.onNetworkCancelClicked() })
	d2ui.AddWidget(&v.networkCancelButton)

	v.btnTCPIPHostGame = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "HOST GAME")
	v.btnTCPIPHostGame.SetPosition(264, 280)
	v.btnTCPIPHostGame.OnActivated(func() { v.onTCPIPHostGameClicked() })
	d2ui.AddWidget(&v.btnTCPIPHostGame)

	v.btnTCPIPJoinGame = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "JOIN GAME")
	v.btnTCPIPJoinGame.SetPosition(264, 320)
	v.btnTCPIPJoinGame.OnActivated(func() { v.onTCPIPJoinGameClicked() })
	d2ui.AddWidget(&v.btnTCPIPJoinGame)
}

func (v *MainMenu) onMapTestClicked() {
	d2screen.SetNextScreen(CreateMapEngineTest(0, 1, v.terminal, v.renderer))
}

func (v *MainMenu) onSinglePlayerClicked() {
	// Go here only if existing characters are available to select
	if d2player.HasGameStates() {
		d2screen.SetNextScreen(CreateCharacterSelect(v.renderer, v.audioProvider, d2clientconnectiontype.Local,
			v.tcpJoinGameEntry.GetText(), v.terminal))
		return
	}

	d2screen.SetNextScreen(CreateSelectHeroClass(v.renderer, v.audioProvider,
		d2clientconnectiontype.Local, v.tcpJoinGameEntry.GetText(), v.terminal))
}

func (v *MainMenu) onGithubButtonClicked() {
	url := "https://www.github.com/OpenDiablo2/OpenDiablo2"

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

func (v *MainMenu) onExitButtonClicked() {
	os.Exit(0)
}

func (v *MainMenu) onCreditsButtonClicked() {
	d2screen.SetNextScreen(CreateCredits(v.renderer, v.audioProvider))
}

// Render renders the main menu
func (v *MainMenu) Render(screen d2interface.Surface) error {
	if err := v.renderBackgrounds(screen); err != nil {
		return err
	}

	if err := v.renderLogos(screen); err != nil {
		return err
	}

	if err := v.renderLabels(screen); err != nil {
		return err
	}

	return nil
}

func (v *MainMenu) renderBackgrounds(screen d2interface.Surface) error {
	switch v.screenMode {
	case screenModeTrademark:
		if err := v.trademarkBackground.RenderSegmented(screen, 4, 3, 0); err != nil {
			return err
		}
	case screenModeServerIP:
		if err := v.serverIPBackground.RenderSegmented(screen, 2, 1, 0); err != nil {
			return err
		}
	case screenModeTCPIP:
		if err := v.tcpIPBackground.RenderSegmented(screen, 4, 3, 0); err != nil {
			return err
		}
	default:
		if err := v.background.RenderSegmented(screen, 4, 3, 0); err != nil {
			return err
		}
	}

	return nil
}

func (v *MainMenu) renderLogos(screen d2interface.Surface) error {
	switch v.screenMode {
	case screenModeTrademark, screenModeMainMenu, screenModeMultiplayer:
		if err := v.diabloLogoLeftBack.Render(screen); err != nil {
			return err
		}

		if err := v.diabloLogoRightBack.Render(screen); err != nil {
			return err
		}

		if err := v.diabloLogoLeft.Render(screen); err != nil {
			return err
		}

		if err := v.diabloLogoRight.Render(screen); err != nil {
			return err
		}
	}

	return nil
}

func (v *MainMenu) renderLabels(screen d2interface.Surface) error {
	switch v.screenMode {
	case screenModeServerIP:
		v.tcpIPOptionsLabel.Render(screen)
		v.tcpJoinGameLabel.Render(screen)
	case screenModeTCPIP:
		v.tcpIPOptionsLabel.Render(screen)
	case screenModeTrademark:
		v.copyrightLabel.Render(screen)
		v.copyrightLabel2.Render(screen)
	case screenModeMainMenu:
		v.openDiabloLabel.Render(screen)
		v.versionLabel.Render(screen)
		v.commitLabel.Render(screen)
	}

	return nil
}

// Advance runs the update logic on the main menu
func (v *MainMenu) Advance(tickTime float64) error {
	switch v.screenMode {
	case screenModeMainMenu, screenModeTrademark, screenModeMultiplayer:
		if err := v.diabloLogoLeftBack.Advance(tickTime); err != nil {
			return err
		}

		if err := v.diabloLogoRightBack.Advance(tickTime); err != nil {
			return err
		}

		if err := v.diabloLogoLeft.Advance(tickTime); err != nil {
			return err
		}

		if err := v.diabloLogoRight.Advance(tickTime); err != nil {
			return err
		}
	}

	return nil
}

// OnMouseButtonDown is called when a mouse button is clicked
func (v *MainMenu) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	if v.screenMode == screenModeTrademark && event.Button() == d2interface.MouseButtonLeft {
		v.setScreenMode(screenModeMainMenu)
		return true
	}

	return false
}

func (v *MainMenu) setScreenMode(screenMode mainMenuScreenMode) {
	v.screenMode = screenMode
	isMainMenu := screenMode == screenModeMainMenu
	isMultiplayer := screenMode == screenModeMultiplayer
	isTCPIP := screenMode == screenModeTCPIP
	isServerIP := screenMode == screenModeServerIP

	v.exitDiabloButton.SetVisible(isMainMenu)
	v.creditsButton.SetVisible(isMainMenu)
	v.cinematicsButton.SetVisible(isMainMenu)
	v.singlePlayerButton.SetVisible(isMainMenu)
	v.githubButton.SetVisible(isMainMenu)
	v.mapTestButton.SetVisible(isMainMenu)
	v.multiplayerButton.SetVisible(isMainMenu)
	v.networkTCPIPButton.SetVisible(isMultiplayer)
	v.networkCancelButton.SetVisible(isMultiplayer)
	v.btnTCPIPCancel.SetVisible(isTCPIP)
	v.btnTCPIPHostGame.SetVisible(isTCPIP)
	v.btnTCPIPJoinGame.SetVisible(isTCPIP)
	v.tcpJoinGameEntry.SetVisible(isServerIP)

	if isServerIP {
		v.tcpJoinGameEntry.Activate()
	}

	v.btnServerIPOk.SetVisible(isServerIP)
	v.btnServerIPCancel.SetVisible(isServerIP)
}

func (v *MainMenu) onNetworkCancelClicked() {
	v.setScreenMode(screenModeMainMenu)
}

func (v *MainMenu) onMultiplayerClicked() {
	v.setScreenMode(screenModeMultiplayer)
}

func (v *MainMenu) onNetworkTCPIPClicked() {
	v.setScreenMode(screenModeTCPIP)
}

func (v *MainMenu) onTCPIPCancelClicked() {
	v.setScreenMode(screenModeMultiplayer)
}

func (v *MainMenu) onTCPIPHostGameClicked() {
	d2screen.SetNextScreen(CreateCharacterSelect(v.renderer, v.audioProvider,
		d2clientconnectiontype.LANServer, "", v.terminal))
}

func (v *MainMenu) onTCPIPJoinGameClicked() {
	v.setScreenMode(screenModeServerIP)
}

func (v *MainMenu) onBtnTCPIPCancelClicked() {
	v.setScreenMode(screenModeTCPIP)
}

func (v *MainMenu) onBtnTCPIPOkClicked() {
	d2screen.SetNextScreen(CreateCharacterSelect(v.renderer, v.audioProvider,
		d2clientconnectiontype.LANClient, v.tcpJoinGameEntry.GetText(), v.terminal))
}
