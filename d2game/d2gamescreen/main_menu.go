// Package d2gamescreen contains the screens
package d2gamescreen

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
)

type mainMenuScreenMode int

// mainMenuScreenMode types
const (
	ScreenModeUnknown mainMenuScreenMode = iota
	ScreenModeTrademark
	ScreenModeMainMenu
	ScreenModeMultiplayer
	ScreenModeTCPIP
	ScreenModeServerIP
)

const (
	joinGameDialogX, joinGameDialogY         = 318, 245
	serverIPbackgroundX, serverIPbackgroundY = 270, 175
	backgroundX, backgroundY                 = 0, 0
	versionLabelX, versionLabelY             = 795, -10
	commitLabelX, commitLabelY               = 2, 2
	copyrightX, copyrightY                   = 400, 500
	copyright2X, copyright2Y                 = 400, 525
	od2LabelX, od2LabelY                     = 400, 580
	tcpOptionsX, tcpOptionsY                 = 400, 23
	joinGameX, joinGameY                     = 400, 190
	diabloLogoX, diabloLogoY                 = 400, 120
	exitDiabloBtnX, exitDiabloBtnY           = 264, 535
	creditBtnX, creditBtnY                   = 264, 505
	cineBtnX, cineBtnY                       = 401, 505
	singlePlayerBtnX, singlePlayerBtnY       = 264, 290
	githubBtnX, githubBtnY                   = 264, 400
	mapTestBtnX, mapTestBtnY                 = 264, 440
	tcpBtnX, tcpBtnY                         = 33, 543
	srvCancelBtnX, srvCancelBtnY             = 285, 305
	srvOkBtnX, srvOkBtnY                     = 420, 305
	multiplayerBtnX, multiplayerBtnY         = 264, 330
	tcpNetBtnX, tcpNetBtnY                   = 264, 280
	networkCancelBtnX, networkCancelBtnY     = 264, 540
	tcpHostBtnX, tcpHostBtnY                 = 264, 280
	tcpJoinBtnX, tcpJoinBtnY                 = 264, 320
)

const (
	white       = 0xffffffff
	lightYellow = 0xffff8cff
	gold        = 0xd8c480ff
)

const (
	joinGameCharacterFilter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890._:"
)

// BuildInfo contains information about the current build
type BuildInfo struct {
	Branch, Commit string
}

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

	inputManager  d2interface.InputManager
	renderer      d2interface.Renderer
	audioProvider d2interface.AudioProvider
	scriptEngine  *d2script.ScriptEngine
	navigator     Navigator

	buildInfo BuildInfo
}

// CreateMainMenu creates an instance of MainMenu
func CreateMainMenu(
	navigator Navigator,
	renderer d2interface.Renderer,
	inputManager d2interface.InputManager,
	audioProvider d2interface.AudioProvider,
	buildInfo BuildInfo,
) *MainMenu {
	return &MainMenu{
		screenMode:     ScreenModeUnknown,
		leftButtonHeld: true,
		renderer:       renderer,
		inputManager:   inputManager,
		audioProvider:  audioProvider,
		navigator:      navigator,
		buildInfo: buildInfo,
	}
}

// OnLoad is called to load the resources for the main menu
func (v *MainMenu) OnLoad(loading d2screen.LoadingState) {
	v.audioProvider.PlayBGM(d2resource.BGMTitle)
	loading.Progress(twentyPercent)

	v.createLabels(loading)
	v.loadBackgroundSprites()
	v.createLogos(loading)
	v.createButtons(loading)

	v.tcpJoinGameEntry = d2ui.CreateTextbox()
	v.tcpJoinGameEntry.SetPosition(joinGameDialogX, joinGameDialogY)
	v.tcpJoinGameEntry.SetFilter(joinGameCharacterFilter)
	d2ui.AddWidget(&v.tcpJoinGameEntry)
	loading.Progress(ninetyPercent)

	if v.screenMode == ScreenModeUnknown {
		v.SetScreenMode(ScreenModeTrademark)
	} else {
		v.SetScreenMode(ScreenModeMainMenu)
	}

	if err := v.inputManager.BindHandler(v); err != nil {
		fmt.Println("failed to add main menu as event handler")
	}
}

func (v *MainMenu) loadBackgroundSprites() {
	animation, _ := d2asset.LoadAnimation(d2resource.GameSelectScreen, d2resource.PaletteSky)
	v.background, _ = d2ui.LoadSprite(animation)
	v.background.SetPosition(backgroundX, backgroundY)

	animation, _ = d2asset.LoadAnimation(d2resource.TrademarkScreen, d2resource.PaletteSky)
	v.trademarkBackground, _ = d2ui.LoadSprite(animation)
	v.trademarkBackground.SetPosition(backgroundX, backgroundY)

	animation, _ = d2asset.LoadAnimation(d2resource.TCPIPBackground, d2resource.PaletteSky)
	v.tcpIPBackground, _ = d2ui.LoadSprite(animation)
	v.tcpIPBackground.SetPosition(backgroundX, backgroundY)

	animation, _ = d2asset.LoadAnimation(d2resource.PopUpOkCancel, d2resource.PaletteFechar)
	v.serverIPBackground, _ = d2ui.LoadSprite(animation)
	v.serverIPBackground.SetPosition(serverIPbackgroundX, serverIPbackgroundY)
}

func (v *MainMenu) createLabels(loading d2screen.LoadingState) {
	v.versionLabel = d2ui.CreateLabel(d2resource.FontFormal12, d2resource.PaletteStatic)
	v.versionLabel.Alignment = d2gui.HorizontalAlignRight
	v.versionLabel.SetText("OpenDiablo2 - " + v.buildInfo.Branch)
	v.versionLabel.Color = rgbaColor(white)
	v.versionLabel.SetPosition(versionLabelX, versionLabelY)

	v.commitLabel = d2ui.CreateLabel(d2resource.FontFormal10, d2resource.PaletteStatic)
	v.commitLabel.Alignment = d2gui.HorizontalAlignLeft
	v.commitLabel.SetText(v.buildInfo.Commit)
	v.commitLabel.Color = rgbaColor(white)
	v.commitLabel.SetPosition(commitLabelX, commitLabelY)

	v.copyrightLabel = d2ui.CreateLabel(d2resource.FontFormal12, d2resource.PaletteStatic)
	v.copyrightLabel.Alignment = d2gui.HorizontalAlignCenter
	v.copyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
	v.copyrightLabel.Color = rgbaColor(lightBrown)
	v.copyrightLabel.SetPosition(copyrightX, copyrightY)
	loading.Progress(thirtyPercent)

	v.copyrightLabel2 = d2ui.CreateLabel(d2resource.FontFormal12, d2resource.PaletteStatic)
	v.copyrightLabel2.Alignment = d2gui.HorizontalAlignCenter
	v.copyrightLabel2.SetText("All Rights Reserved.")
	v.copyrightLabel2.Color = rgbaColor(lightBrown)
	v.copyrightLabel2.SetPosition(copyright2X, copyright2Y)

	v.openDiabloLabel = d2ui.CreateLabel(d2resource.FontFormal10, d2resource.PaletteStatic)
	v.openDiabloLabel.Alignment = d2gui.HorizontalAlignCenter
	v.openDiabloLabel.SetText("OpenDiablo2 is neither developed by, nor endorsed by Blizzard or its parent company Activision")
	v.openDiabloLabel.Color = rgbaColor(lightYellow)
	v.openDiabloLabel.SetPosition(od2LabelX, od2LabelY)
	loading.Progress(fiftyPercent)

	v.tcpIPOptionsLabel = d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteUnits)
	v.tcpIPOptionsLabel.SetPosition(tcpOptionsX, tcpOptionsY)
	v.tcpIPOptionsLabel.Alignment = d2gui.HorizontalAlignCenter
	v.tcpIPOptionsLabel.SetText("TCP/IP Options")

	v.tcpJoinGameLabel = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	v.tcpJoinGameLabel.Alignment = d2gui.HorizontalAlignCenter
	v.tcpJoinGameLabel.SetText("Enter Host IP Address\nto Join Game")

	v.tcpJoinGameLabel.Color = rgbaColor(gold)
	v.tcpJoinGameLabel.SetPosition(joinGameX, joinGameY)
}

func (v *MainMenu) createLogos(loading d2screen.LoadingState) {
	animation, _ := d2asset.LoadAnimation(d2resource.Diablo2LogoFireLeft, d2resource.PaletteUnits)
	v.diabloLogoLeft, _ = d2ui.LoadSprite(animation)
	v.diabloLogoLeft.SetEffect(d2enum.DrawEffectModulate)
	v.diabloLogoLeft.PlayForward()
	v.diabloLogoLeft.SetPosition(diabloLogoX, diabloLogoY)
	loading.Progress(sixtyPercent)

	animation, _ = d2asset.LoadAnimation(d2resource.Diablo2LogoFireRight, d2resource.PaletteUnits)
	v.diabloLogoRight, _ = d2ui.LoadSprite(animation)
	v.diabloLogoRight.SetEffect(d2enum.DrawEffectModulate)
	v.diabloLogoRight.PlayForward()
	v.diabloLogoRight.SetPosition(diabloLogoX, diabloLogoY)

	animation, _ = d2asset.LoadAnimation(d2resource.Diablo2LogoBlackLeft, d2resource.PaletteUnits)
	v.diabloLogoLeftBack, _ = d2ui.LoadSprite(animation)
	v.diabloLogoLeftBack.SetPosition(diabloLogoX, diabloLogoY)

	animation, _ = d2asset.LoadAnimation(d2resource.Diablo2LogoBlackRight, d2resource.PaletteUnits)
	v.diabloLogoRightBack, _ = d2ui.LoadSprite(animation)
	v.diabloLogoRightBack.SetPosition(diabloLogoX, diabloLogoY)
}

func (v *MainMenu) createButtons(loading d2screen.LoadingState) {
	v.exitDiabloButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "EXIT DIABLO II")
	v.exitDiabloButton.SetPosition(exitDiabloBtnX, exitDiabloBtnY)
	v.exitDiabloButton.OnActivated(func() { v.onExitButtonClicked() })
	d2ui.AddWidget(&v.exitDiabloButton)

	v.creditsButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeShort, "CREDITS")
	v.creditsButton.SetPosition(creditBtnX, creditBtnY)
	v.creditsButton.OnActivated(func() { v.onCreditsButtonClicked() })
	d2ui.AddWidget(&v.creditsButton)

	v.cinematicsButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeShort, "CINEMATICS")
	v.cinematicsButton.SetPosition(cineBtnX, cineBtnY)
	d2ui.AddWidget(&v.cinematicsButton)
	loading.Progress(seventyPercent)

	v.singlePlayerButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "SINGLE PLAYER")
	v.singlePlayerButton.SetPosition(singlePlayerBtnX, singlePlayerBtnY)
	v.singlePlayerButton.OnActivated(func() { v.onSinglePlayerClicked() })
	d2ui.AddWidget(&v.singlePlayerButton)

	v.githubButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "PROJECT WEBSITE")
	v.githubButton.SetPosition(githubBtnX, githubBtnY)
	v.githubButton.OnActivated(func() { v.onGithubButtonClicked() })
	d2ui.AddWidget(&v.githubButton)

	v.mapTestButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "MAP ENGINE TEST")
	v.mapTestButton.SetPosition(mapTestBtnX, mapTestBtnY)
	v.mapTestButton.OnActivated(func() { v.onMapTestClicked() })
	d2ui.AddWidget(&v.mapTestButton)

	v.btnTCPIPCancel = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeMedium, d2common.TranslateString("cancel"))
	v.btnTCPIPCancel.SetPosition(tcpBtnX, tcpBtnY)
	v.btnTCPIPCancel.OnActivated(func() { v.onTCPIPCancelClicked() })
	d2ui.AddWidget(&v.btnTCPIPCancel)

	v.btnServerIPCancel = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeOkCancel, "CANCEL")
	v.btnServerIPCancel.SetPosition(srvCancelBtnX, srvCancelBtnY)
	v.btnServerIPCancel.OnActivated(func() { v.onBtnTCPIPCancelClicked() })
	d2ui.AddWidget(&v.btnServerIPCancel)

	v.btnServerIPOk = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeOkCancel, "OK")
	v.btnServerIPOk.SetPosition(srvOkBtnX, srvOkBtnY)
	v.btnServerIPOk.OnActivated(func() { v.onBtnTCPIPOkClicked() })
	d2ui.AddWidget(&v.btnServerIPOk)

	v.createMultiplayerMenuButtons()
	loading.Progress(eightyPercent)
}

func (v *MainMenu) createMultiplayerMenuButtons() {
	v.multiplayerButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "MULTIPLAYER")
	v.multiplayerButton.SetPosition(multiplayerBtnX, multiplayerBtnY)
	v.multiplayerButton.OnActivated(func() { v.onMultiplayerClicked() })
	d2ui.AddWidget(&v.multiplayerButton)

	v.networkTCPIPButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "TCP/IP GAME")
	v.networkTCPIPButton.SetPosition(tcpNetBtnX, tcpNetBtnY)
	v.networkTCPIPButton.OnActivated(func() { v.onNetworkTCPIPClicked() })
	d2ui.AddWidget(&v.networkTCPIPButton)

	v.networkCancelButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, d2common.TranslateString("cancel"))
	v.networkCancelButton.SetPosition(networkCancelBtnX, networkCancelBtnY)
	v.networkCancelButton.OnActivated(func() { v.onNetworkCancelClicked() })
	d2ui.AddWidget(&v.networkCancelButton)

	v.btnTCPIPHostGame = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "HOST GAME")
	v.btnTCPIPHostGame.SetPosition(tcpHostBtnX, tcpHostBtnY)
	v.btnTCPIPHostGame.OnActivated(func() { v.onTCPIPHostGameClicked() })
	d2ui.AddWidget(&v.btnTCPIPHostGame)

	v.btnTCPIPJoinGame = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeWide, "JOIN GAME")
	v.btnTCPIPJoinGame.SetPosition(tcpJoinBtnX, tcpJoinBtnY)
	v.btnTCPIPJoinGame.OnActivated(func() { v.onTCPIPJoinGameClicked() })
	d2ui.AddWidget(&v.btnTCPIPJoinGame)
}

func (v *MainMenu) onMapTestClicked() {
	v.navigator.ToMapEngineTest(0, 1)
}

func (v *MainMenu) onSinglePlayerClicked() {
	if d2player.HasGameStates() {
		// Go here only if existing characters are available to select
		v.navigator.ToCharacterSelect(d2clientconnectiontype.Local, v.tcpJoinGameEntry.GetText())
	} else {
		v.navigator.ToSelectHero(d2clientconnectiontype.Local, v.tcpJoinGameEntry.GetText())
	}
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
	v.navigator.ToCredits()
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
	case ScreenModeTrademark:
		if err := v.trademarkBackground.RenderSegmented(screen, 4, 3, 0); err != nil {
			return err
		}
	case ScreenModeServerIP:
		if err := v.serverIPBackground.RenderSegmented(screen, 2, 1, 0); err != nil {
			return err
		}
	case ScreenModeTCPIP:
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
	case ScreenModeTrademark, ScreenModeMainMenu, ScreenModeMultiplayer:
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
	case ScreenModeServerIP:
		v.tcpIPOptionsLabel.Render(screen)
		v.tcpJoinGameLabel.Render(screen)
	case ScreenModeTCPIP:
		v.tcpIPOptionsLabel.Render(screen)
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

// Advance runs the update logic on the main menu
func (v *MainMenu) Advance(tickTime float64) error {
	switch v.screenMode {
	case ScreenModeMainMenu, ScreenModeTrademark, ScreenModeMultiplayer:
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
	if v.screenMode == ScreenModeTrademark && event.Button() == d2enum.MouseButtonLeft {
		v.SetScreenMode(ScreenModeMainMenu)
		return true
	}

	return false
}

// SetScreenMode sets the screen mode (which sub-menu the screen is on)
func (v *MainMenu) SetScreenMode(screenMode mainMenuScreenMode) {
	v.screenMode = screenMode
	isMainMenu := screenMode == ScreenModeMainMenu
	isMultiplayer := screenMode == ScreenModeMultiplayer
	isTCPIP := screenMode == ScreenModeTCPIP
	isServerIP := screenMode == ScreenModeServerIP

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
	v.SetScreenMode(ScreenModeMainMenu)
}

func (v *MainMenu) onMultiplayerClicked() {
	v.SetScreenMode(ScreenModeMultiplayer)
}

func (v *MainMenu) onNetworkTCPIPClicked() {
	v.SetScreenMode(ScreenModeTCPIP)
}

func (v *MainMenu) onTCPIPCancelClicked() {
	v.SetScreenMode(ScreenModeMultiplayer)
}

func (v *MainMenu) onTCPIPHostGameClicked() {
	v.navigator.ToCharacterSelect(d2clientconnectiontype.LANServer, "")
}

func (v *MainMenu) onTCPIPJoinGameClicked() {
	v.SetScreenMode(ScreenModeServerIP)
}

func (v *MainMenu) onBtnTCPIPCancelClicked() {
	v.SetScreenMode(ScreenModeTCPIP)
}

func (v *MainMenu) onBtnTCPIPOkClicked() {
	v.navigator.ToCharacterSelect(d2clientconnectiontype.LANClient, v.tcpJoinGameEntry.GetText())
}
