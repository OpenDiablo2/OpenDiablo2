package d2gamescreen

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2video"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	cinematicsX, cinematicsY               = 240, 90
	cinematicsLabelX, cinematicsLabelY     = 400, 110
	a1BtnX, a1BtnY                         = 264, 160
	a2BtnX, a2BtnY                         = 264, 205
	a3BtnX, a3BtnY                         = 264, 250
	a4BtnX, a4BtnY                         = 264, 295
	endCreditClassBtnX, endCreditClassBtnY = 264, 340
	a5BtnX, a5BtnY                         = 264, 385
	endCreditExpBtnX, endCreditExpBtnY     = 264, 430
	cinematicsExitBtnX, cinematicsExitBtnY = 340, 470
)

// CreateCinematics creates an instance of the credits screen
func CreateCinematics(
	navigator d2interface.Navigator,
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	aup d2interface.AudioProvider,
	l d2util.LogLevel,
	ui *d2ui.UIManager) *Cinematics {
	cinematics := &Cinematics{
		asset:         asset,
		renderer:      renderer,
		navigator:     navigator,
		uiManager:     ui,
		audioProvider: aup,
	}

	cinematics.Logger = d2util.NewLogger()
	cinematics.Logger.SetPrefix(logPrefix)
	cinematics.Logger.SetLevel(l)

	return cinematics
}

// Cinematics represents the cinematics screen
type Cinematics struct {
	cinematicsBackground *d2ui.Sprite
	background           *d2ui.Sprite
	a1Btn                *d2ui.Button
	a2Btn                *d2ui.Button
	a3Btn                *d2ui.Button
	a4Btn                *d2ui.Button
	a5Btn                *d2ui.Button
	endCreditClassBtn    *d2ui.Button
	endCreditExpBtn      *d2ui.Button
	cinematicsExitBtn    *d2ui.Button
	cinematicsLabel      *d2ui.Label

	asset         *d2asset.AssetManager
	renderer      d2interface.Renderer
	navigator     d2interface.Navigator
	uiManager     *d2ui.UIManager
	videoDecoder  *d2video.BinkDecoder
	audioProvider d2interface.AudioProvider

	*d2util.Logger
}

// OnLoad is called to load the resources for the credits screen
func (v *Cinematics) OnLoad(_ d2screen.LoadingState) {
	var err error

	v.audioProvider.PlayBGM("")

	v.background, err = v.uiManager.NewSprite(d2resource.GameSelectScreen, d2resource.PaletteSky)

	if err != nil {
		v.Error(err.Error())
	}

	v.background.SetPosition(backgroundX, backgroundY)

	v.cinematicsBackground, err = v.uiManager.NewSprite(d2resource.CinematicsBackground, d2resource.PaletteSky)

	if err != nil {
		v.Error(err.Error())
	}

	v.cinematicsBackground.SetPosition(cinematicsX, cinematicsY)

	v.createButtons()

	v.cinematicsLabel = v.uiManager.NewLabel(d2resource.Font30, d2resource.PaletteStatic)
	v.cinematicsLabel.Alignment = d2ui.HorizontalAlignCenter
	v.cinematicsLabel.SetText(v.asset.TranslateLabel(d2enum.SelectCinematicLabel))
	v.cinematicsLabel.Color[0] = rgbaColor(lightBrown)
	v.cinematicsLabel.SetPosition(cinematicsLabelX, cinematicsLabelY)
}

func (v *Cinematics) createButtons() {
	v.cinematicsExitBtn = v.uiManager.NewButton(d2ui.ButtonTypeMedium,
		v.asset.TranslateString(v.asset.TranslateLabel(d2enum.CancelLabel)))
	v.cinematicsExitBtn.SetPosition(cinematicsExitBtnX, cinematicsExitBtnY)
	v.cinematicsExitBtn.OnActivated(func() { v.onCinematicsExitBtnClicked() })

	v.a1Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, v.asset.TranslateString("act1X"))
	v.a1Btn.SetPosition(a1BtnX, a1BtnY)
	v.a1Btn.OnActivated(func() { v.onA1BtnClicked() })

	v.a2Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, v.asset.TranslateString("act2X"))
	v.a2Btn.SetPosition(a2BtnX, a2BtnY)
	v.a2Btn.OnActivated(func() { v.onA2BtnClicked() })

	v.a3Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, v.asset.TranslateString("act3X"))
	v.a3Btn.SetPosition(a3BtnX, a3BtnY)
	v.a3Btn.OnActivated(func() { v.onA3BtnClicked() })

	v.a4Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, v.asset.TranslateString("act4X"))
	v.a4Btn.SetPosition(a4BtnX, a4BtnY)
	v.a4Btn.OnActivated(func() { v.onA4BtnClicked() })

	v.endCreditClassBtn = v.uiManager.NewButton(d2ui.ButtonTypeWide, v.asset.TranslateString("strepilogueX"))
	v.endCreditClassBtn.SetPosition(endCreditClassBtnX, endCreditClassBtnY)
	v.endCreditClassBtn.OnActivated(func() { v.onEndCreditClassBtnClicked() })

	v.a5Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, v.asset.TranslateString("act5X"))
	v.a5Btn.SetPosition(a5BtnX, a5BtnY)
	v.a5Btn.OnActivated(func() { v.onA5BtnClicked() })

	v.endCreditExpBtn = v.uiManager.NewButton(d2ui.ButtonTypeWide, v.asset.TranslateString("strlastcinematic"))
	v.endCreditExpBtn.SetPosition(endCreditExpBtnX, endCreditExpBtnY)
	v.endCreditExpBtn.OnActivated(func() { v.onEndCreditExpBtnClicked() })
}

func (v *Cinematics) onCinematicsExitBtnClicked() {
	v.navigator.ToMainMenu()
}

func (v *Cinematics) onA1BtnClicked() {
	v.playVideo(d2resource.Act1Intro)
}

func (v *Cinematics) onA2BtnClicked() {
	v.playVideo(d2resource.Act2Intro)
}

func (v *Cinematics) onA3BtnClicked() {
	v.playVideo(d2resource.Act3Intro)
}

func (v *Cinematics) onA4BtnClicked() {
	v.playVideo(d2resource.Act4Intro)
}

func (v *Cinematics) onEndCreditClassBtnClicked() {
	v.playVideo(d2resource.Act4Outro)
}

func (v *Cinematics) onA5BtnClicked() {
	v.playVideo(d2resource.Act5Intro)
}

func (v *Cinematics) onEndCreditExpBtnClicked() {
	v.playVideo(d2resource.Act5Outro)
}

func (v *Cinematics) playVideo(path string) {
	videoBytes, err := v.asset.LoadFile(path)
	if err != nil {
		v.Error(err.Error())
		return
	}

	v.videoDecoder = d2video.CreateBinkDecoder(videoBytes)
}

// Render renders the credits screen
func (v *Cinematics) Render(screen d2interface.Surface) {
	v.background.RenderSegmented(screen, 4, 3, 0)
	v.cinematicsBackground.RenderSegmented(screen, 2, 2, 0)
	v.cinematicsLabel.Render(screen)
}
