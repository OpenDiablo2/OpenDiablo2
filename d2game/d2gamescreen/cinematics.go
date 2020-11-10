package d2gamescreen

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2video"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
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

// Cinematics represents the cinematics screen
type Cinematics struct {
	cinematicsBackground *d2ui.Sprite
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
}

// CreateCinematics creates an instance of the credits screen
func CreateCinematics(
	navigator d2interface.Navigator,
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	aup d2interface.AudioProvider,
	ui *d2ui.UIManager) *Cinematics {
	result := &Cinematics{
		asset:         asset,
		renderer:      renderer,
		navigator:     navigator,
		uiManager:     ui,
		audioProvider: aup,
	}

	return result
}

// OnLoad is called to load the resources for the credits screen
func (v *Cinematics) OnLoad(loading d2screen.LoadingState) {
	var err error

	v.audioProvider.PlayBGM("")

	v.cinematicsBackground, err = v.uiManager.NewSprite(d2resource.CinematicsBackground, d2resource.PaletteSky)

	if err != nil {
		log.Print(err)
	}

	v.cinematicsBackground.SetPosition(cinematicsX, cinematicsY)

	loading.Progress(twentyPercent)

	v.createButtons()
	loading.Progress(fourtyPercent)

	v.cinematicsLabel = v.uiManager.NewLabel(d2resource.Font30, d2resource.PaletteStatic)
	v.cinematicsLabel.Alignment = d2gui.HorizontalAlignCenter
	v.cinematicsLabel.SetText("SELECT CINEMATIC")
	v.cinematicsLabel.Color[0] = rgbaColor(lightBrown)
	v.cinematicsLabel.SetPosition(cinematicsLabelX, cinematicsLabelY)

	loading.Progress(sixtyPercent)
	loading.Progress(eightyPercent)
}

func (v *Cinematics) createButtons() {
	/*CINEMATICS NAMES:
	Act 1: The Sister's Lament
	Act 2: Dessert Journay
	Act 3: Mephisto's Jungle
	Act 4: Enter Hell
	Act 5: Search For Ball
	end Credit Classic: Terror's End
	end Credit Expansion: Destruction's End
	*/
	v.cinematicsExitBtn = v.uiManager.NewButton(d2ui.ButtonTypeMedium, "CANCEL")
	v.cinematicsExitBtn.SetPosition(cinematicsExitBtnX, cinematicsExitBtnY)
	v.cinematicsExitBtn.OnActivated(func() { v.onCinematicsExitBtnClicked() })

	v.a1Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, "THE SISTER'S LAMENT")
	v.a1Btn.SetPosition(a1BtnX, a1BtnY)
	v.a1Btn.OnActivated(func() { v.onA1BtnClicked() })

	v.a2Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, "DESSERT JOURNAY")
	v.a2Btn.SetPosition(a2BtnX, a2BtnY)
	v.a2Btn.OnActivated(func() { v.onA2BtnClicked() })

	v.a3Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, "MEPHISTO'S JUNGLE")
	v.a3Btn.SetPosition(a3BtnX, a3BtnY)
	v.a3Btn.OnActivated(func() { v.onA3BtnClicked() })

	v.a4Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, "ENTER HELL")
	v.a4Btn.SetPosition(a4BtnX, a4BtnY)
	v.a4Btn.OnActivated(func() { v.onA4BtnClicked() })

	v.endCreditClassBtn = v.uiManager.NewButton(d2ui.ButtonTypeWide, "TERROR'S END")
	v.endCreditClassBtn.SetPosition(endCreditClassBtnX, endCreditClassBtnY)
	v.endCreditClassBtn.OnActivated(func() { v.onEndCreditClassBtnClicked() })

	v.a5Btn = v.uiManager.NewButton(d2ui.ButtonTypeWide, "SEARCH FOR BAAL")
	v.a5Btn.SetPosition(a5BtnX, a5BtnY)
	v.a5Btn.OnActivated(func() { v.onA5BtnClicked() })

	v.endCreditExpBtn = v.uiManager.NewButton(d2ui.ButtonTypeWide, "DESTRUCTION'S END")
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
		log.Print(err)
		return
	}

	v.videoDecoder = d2video.CreateBinkDecoder(videoBytes)
}

// Render renders the credits screen
func (v *Cinematics) Render(screen d2interface.Surface) {
	err := v.cinematicsBackground.RenderSegmented(screen, 2, 2, 0)
	if err != nil {
		return
	}

	v.cinematicsLabel.Render(screen)
}
