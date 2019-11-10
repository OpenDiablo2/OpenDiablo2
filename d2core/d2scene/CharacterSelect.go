package d2scene

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"
	dh "github.com/OpenDiablo2/OpenDiablo2/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
	"github.com/hajimehoshi/ebiten"
)

type CharacterSelect struct {
	uiManager         *d2ui.Manager
	soundManager      *d2audio.Manager
	fileProvider      d2interface.FileProvider
	sceneProvider     d2interface.SceneProvider
	background        d2render.Sprite
	newCharButton     *d2ui.Button
	convertCharButton *d2ui.Button
	deleteCharButton  *d2ui.Button
	exitButton        *d2ui.Button
	okButton          *d2ui.Button
}

func CreateCharacterSelect(
	fileProvider d2interface.FileProvider,
	sceneProvider d2interface.SceneProvider,
	uiManager *d2ui.Manager,
	soundManager *d2audio.Manager,
) *CharacterSelect {
	result := &CharacterSelect{
		uiManager:     uiManager,
		sceneProvider: sceneProvider,
		fileProvider:  fileProvider,
		soundManager:  soundManager,
	}
	return result
}

func (v *CharacterSelect) Load() []func() {
	v.soundManager.PlayBGM(d2resource.BGMTitle)
	return []func(){
		func() {
			v.background = d2render.CreateSprite(v.fileProvider.LoadFile(d2resource.CharacterSelectionBackground), d2datadict.Palettes[d2enum.Sky])
			v.background.MoveTo(0, 0)
		},
		func() {
			v.newCharButton = d2ui.CreateButton(d2ui.ButtonTypeTall, v.fileProvider, dh.CombineStrings(dh.SplitIntoLinesWithMaxWidth(d2common.TranslateString("#831"), 15)))
			v.newCharButton.MoveTo(33, 468)
			v.newCharButton.OnActivated(func() { v.onNewCharButtonClicked() })
			v.uiManager.AddWidget(v.newCharButton)
		},
		func() {
			v.convertCharButton = d2ui.CreateButton(d2ui.ButtonTypeTall, v.fileProvider, dh.CombineStrings(dh.SplitIntoLinesWithMaxWidth(d2common.TranslateString("#825"), 15)))
			v.convertCharButton.MoveTo(233, 468)
			v.convertCharButton.SetEnabled(false)
			v.uiManager.AddWidget(v.convertCharButton)
		},
		func() {
			v.deleteCharButton = d2ui.CreateButton(d2ui.ButtonTypeTall, v.fileProvider, dh.CombineStrings(dh.SplitIntoLinesWithMaxWidth(d2common.TranslateString("#832"), 15)))
			v.deleteCharButton.MoveTo(433, 468)
			v.deleteCharButton.SetEnabled(false)
			v.uiManager.AddWidget(v.deleteCharButton)
		},
		func() {
			v.exitButton = d2ui.CreateButton(d2ui.ButtonTypeMedium, v.fileProvider, d2common.TranslateString("#970"))
			v.exitButton.MoveTo(33, 537)
			v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(v.exitButton)
		},
		func() {
			v.okButton = d2ui.CreateButton(d2ui.ButtonTypeMedium, v.fileProvider, d2common.TranslateString("#971"))
			v.okButton.MoveTo(625, 537)
			v.okButton.SetEnabled(false)
			v.uiManager.AddWidget(v.okButton)
		},
	}
}

func (v *CharacterSelect) onNewCharButtonClicked() {
	v.sceneProvider.SetNextScene(CreateSelectHeroClass(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager))
}

func (v *CharacterSelect) onExitButtonClicked() {
	mainMenu := CreateMainMenu(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager)
	mainMenu.ShowTrademarkScreen = false
	v.sceneProvider.SetNextScene(mainMenu)
}

func (v *CharacterSelect) Unload() {
}

func (v *CharacterSelect) Render(screen *ebiten.Image) {
	v.background.DrawSegments(screen, 4, 3, 0)
}

func (v *CharacterSelect) Update(tickTime float64) {
}
