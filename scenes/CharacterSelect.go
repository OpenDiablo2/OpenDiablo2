package scenes

import (
	"github.com/OpenDiablo2/OpenDiablo2/common"
	"github.com/OpenDiablo2/OpenDiablo2/palettedefs"
	"github.com/OpenDiablo2/OpenDiablo2/resourcepaths"
	"github.com/OpenDiablo2/OpenDiablo2/sound"
	"github.com/OpenDiablo2/OpenDiablo2/ui"
	"github.com/hajimehoshi/ebiten"
)

type CharacterSelect struct {
	uiManager         *ui.Manager
	soundManager      *sound.Manager
	fileProvider      common.FileProvider
	sceneProvider     SceneProvider
	background        *common.Sprite
	newCharButton     *ui.Button
	convertCharButton *ui.Button
	deleteCharButton  *ui.Button
	exitButton        *ui.Button
	okButton          *ui.Button
}

func CreateCharacterSelect(
	fileProvider common.FileProvider,
	sceneProvider SceneProvider,
	uiManager *ui.Manager,
	soundManager *sound.Manager,
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
	v.soundManager.PlayBGM(resourcepaths.BGMTitle)
	return []func(){
		func() {
			v.background = v.fileProvider.LoadSprite(resourcepaths.CharacterSelectionBackground, palettedefs.Sky)
			v.background.MoveTo(0, 0)
		},
		func() {
			v.newCharButton = ui.CreateButton(ui.ButtonTypeTall, v.fileProvider, common.CombineStrings(common.SplitIntoLinesWithMaxWidth(common.TranslateString("#831"), 15)))
			v.newCharButton.MoveTo(33, 468)
			v.newCharButton.OnActivated(func() { v.onNewCharButtonClicked() })
			v.uiManager.AddWidget(v.newCharButton)
		},
		func() {
			v.convertCharButton = ui.CreateButton(ui.ButtonTypeTall, v.fileProvider, common.CombineStrings(common.SplitIntoLinesWithMaxWidth(common.TranslateString("#825"), 15)))
			v.convertCharButton.MoveTo(233, 468)
			v.convertCharButton.SetEnabled(false)
			v.uiManager.AddWidget(v.convertCharButton)
		},
		func() {
			v.deleteCharButton = ui.CreateButton(ui.ButtonTypeTall, v.fileProvider, common.CombineStrings(common.SplitIntoLinesWithMaxWidth(common.TranslateString("#832"), 15)))
			v.deleteCharButton.MoveTo(433, 468)
			v.deleteCharButton.SetEnabled(false)
			v.uiManager.AddWidget(v.deleteCharButton)
		},
		func() {
			v.exitButton = ui.CreateButton(ui.ButtonTypeMedium, v.fileProvider, common.TranslateString("#970"))
			v.exitButton.MoveTo(33, 537)
			v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(v.exitButton)
		},
		func() {
			v.okButton = ui.CreateButton(ui.ButtonTypeMedium, v.fileProvider, common.TranslateString("#971"))
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
