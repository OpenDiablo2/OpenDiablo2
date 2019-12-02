package d2scene

import (
	"image/color"
	"math"
	"os"
	"strings"

	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/OpenDiablo2/D2Shared/d2common"
	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	dh "github.com/OpenDiablo2/D2Shared/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon/d2coreinterface"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
	"github.com/hajimehoshi/ebiten"
)

type CharacterSelect struct {
	uiManager              *d2ui.Manager
	soundManager           *d2audio.Manager
	fileProvider           d2interface.FileProvider
	sceneProvider          d2coreinterface.SceneProvider
	background             d2render.Sprite
	newCharButton          d2ui.Button
	convertCharButton      d2ui.Button
	deleteCharButton       d2ui.Button
	exitButton             d2ui.Button
	okButton               d2ui.Button
	deleteCharCancelButton d2ui.Button
	deleteCharOkButton     d2ui.Button
	selectionBox           d2render.Sprite
	okCancelBox            d2render.Sprite
	d2HeroTitle            d2ui.Label
	deleteCharConfirmLabel d2ui.Label
	charScrollbar          d2ui.Scrollbar
	characterNameLabel     [8]d2ui.Label
	characterStatsLabel    [8]d2ui.Label
	characterExpLabel      [8]d2ui.Label
	characterImage         [8]*d2core.Hero
	gameStates             []*d2core.GameState
	selectedCharacter      int
	mouseButtonPressed     bool
	showDeleteConfirmation bool
}

func CreateCharacterSelect(
	fileProvider d2interface.FileProvider,
	sceneProvider d2coreinterface.SceneProvider,
	uiManager *d2ui.Manager,
	soundManager *d2audio.Manager,
) *CharacterSelect {
	result := &CharacterSelect{
		selectedCharacter: -1,
		uiManager:         uiManager,
		sceneProvider:     sceneProvider,
		fileProvider:      fileProvider,
		soundManager:      soundManager,
	}
	return result
}

func (v *CharacterSelect) Load() []func() {
	v.soundManager.PlayBGM(d2resource.BGMTitle)
	return []func(){
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.CharacterSelectionBackground), d2datadict.Palettes[d2enum.Sky])
			v.background = d2render.CreateSpriteFromDC6(dc6)
			v.background.MoveTo(0, 0)
		},
		func() {
			v.newCharButton = d2ui.CreateButton(d2ui.ButtonTypeTall, v.fileProvider, dh.CombineStrings(dh.SplitIntoLinesWithMaxWidth(d2common.TranslateString("#831"), 15)))
			v.newCharButton.MoveTo(33, 468)
			v.newCharButton.OnActivated(func() { v.onNewCharButtonClicked() })
			v.uiManager.AddWidget(&v.newCharButton)
		},
		func() {
			v.convertCharButton = d2ui.CreateButton(d2ui.ButtonTypeTall, v.fileProvider, dh.CombineStrings(dh.SplitIntoLinesWithMaxWidth(d2common.TranslateString("#825"), 15)))
			v.convertCharButton.MoveTo(233, 468)
			v.convertCharButton.SetEnabled(false)
			v.uiManager.AddWidget(&v.convertCharButton)
		},
		func() {
			v.deleteCharButton = d2ui.CreateButton(d2ui.ButtonTypeTall, v.fileProvider, dh.CombineStrings(dh.SplitIntoLinesWithMaxWidth(d2common.TranslateString("#832"), 15)))
			v.deleteCharButton.OnActivated(func() { v.onDeleteCharButtonClicked() })
			v.deleteCharButton.MoveTo(433, 468)
			v.uiManager.AddWidget(&v.deleteCharButton)
		},
		func() {
			v.exitButton = d2ui.CreateButton(d2ui.ButtonTypeMedium, v.fileProvider, d2common.TranslateString("#970"))
			v.exitButton.MoveTo(33, 537)
			v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(&v.exitButton)
		},
		func() {
			v.deleteCharCancelButton = d2ui.CreateButton(d2ui.ButtonTypeOkCancel, v.fileProvider, d2common.TranslateString("#4231"))
			v.deleteCharCancelButton.MoveTo(282, 308)
			v.deleteCharCancelButton.SetVisible(false)
			v.deleteCharCancelButton.OnActivated(func() { v.onDeleteCharacterCancelClicked() })
			v.uiManager.AddWidget(&v.deleteCharCancelButton)
		},
		func() {
			v.deleteCharOkButton = d2ui.CreateButton(d2ui.ButtonTypeOkCancel, v.fileProvider, d2common.TranslateString("#4227"))
			v.deleteCharOkButton.MoveTo(422, 308)
			v.deleteCharOkButton.SetVisible(false)
			v.deleteCharOkButton.OnActivated(func() { v.onDeleteCharacterConfirmClicked() })
			v.uiManager.AddWidget(&v.deleteCharOkButton)
		},
		func() {
			v.okButton = d2ui.CreateButton(d2ui.ButtonTypeMedium, v.fileProvider, d2common.TranslateString("#971"))
			v.okButton.MoveTo(625, 537)
			v.okButton.OnActivated(func() { v.onOkButtonClicked() })
			v.uiManager.AddWidget(&v.okButton)
		},
		func() {
			v.d2HeroTitle = d2ui.CreateLabel(v.fileProvider, d2resource.Font42, d2enum.Units)
			v.d2HeroTitle.MoveTo(320, 23)
			v.d2HeroTitle.Alignment = d2ui.LabelAlignCenter
		},
		func() {
			v.deleteCharConfirmLabel = d2ui.CreateLabel(v.fileProvider, d2resource.Font16, d2enum.Units)
			lines := dh.SplitIntoLinesWithMaxWidth(d2common.TranslateString("#1878"), 29)
			v.deleteCharConfirmLabel.SetText(strings.Join(lines, "\n"))
			v.deleteCharConfirmLabel.Alignment = d2ui.LabelAlignCenter
			v.deleteCharConfirmLabel.MoveTo(400, 185)
		},
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.CharacterSelectionSelectBox), d2datadict.Palettes[d2enum.Sky])
			v.selectionBox = d2render.CreateSpriteFromDC6(dc6)
			v.selectionBox.MoveTo(37, 86)
		},
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.PopUpOkCancel), d2datadict.Palettes[d2enum.Fechar])
			v.okCancelBox = d2render.CreateSpriteFromDC6(dc6)
			v.okCancelBox.MoveTo(270, 175)
		},
		func() {
			v.charScrollbar = d2ui.CreateScrollbar(v.fileProvider, 586, 87, 369)
			v.charScrollbar.OnActivated(func() { v.onScrollUpdate() })
			v.uiManager.AddWidget(&v.charScrollbar)
		},
		func() {
			for i := 0; i < 8; i++ {
				xOffset := 115
				if i&1 > 0 {
					xOffset = 385
				}
				v.characterNameLabel[i] = d2ui.CreateLabel(v.fileProvider, d2resource.Font16, d2enum.Units)
				v.characterNameLabel[i].Color = color.RGBA{188, 168, 140, 255}
				v.characterNameLabel[i].MoveTo(xOffset, 100+((i/2)*95))
				v.characterStatsLabel[i] = d2ui.CreateLabel(v.fileProvider, d2resource.Font16, d2enum.Units)
				v.characterStatsLabel[i].MoveTo(xOffset, 115+((i/2)*95))
				v.characterExpLabel[i] = d2ui.CreateLabel(v.fileProvider, d2resource.Font16, d2enum.Static)
				v.characterExpLabel[i].Color = color.RGBA{24, 255, 0, 255}
				v.characterExpLabel[i].MoveTo(xOffset, 130+((i/2)*95))
			}
			v.refreshGameStates()
		},
	}
}

func (v *CharacterSelect) onScrollUpdate() {
	v.moveSelectionBox()
	v.updateCharacterBoxes()
}

func (v *CharacterSelect) updateCharacterBoxes() {
	expText := d2common.TranslateString("#803")
	for i := 0; i < 8; i++ {
		idx := i + (v.charScrollbar.GetCurrentOffset() * 2)
		if idx >= len(v.gameStates) {
			v.characterNameLabel[i].SetText("")
			v.characterStatsLabel[i].SetText("")
			v.characterExpLabel[i].SetText("")
			v.characterImage[i] = nil
			continue
		}
		v.characterNameLabel[i].SetText(v.gameStates[idx].HeroName)
		v.characterStatsLabel[i].SetText("Level 1 " + v.gameStates[idx].HeroType.String())
		v.characterExpLabel[i].SetText(expText)
		// TODO: Generate or load the object from the actual player data...
		v.characterImage[i] = d2core.CreateHero(
			0,
			0,
			0,
			v.gameStates[idx].HeroType,
			d2core.HeroObjects[v.gameStates[idx].HeroType],
			v.fileProvider,
		)
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
	v.d2HeroTitle.Draw(screen)
	actualSelectionIndex := v.selectedCharacter - (v.charScrollbar.GetCurrentOffset() * 2)
	if v.selectedCharacter > -1 && actualSelectionIndex >= 0 && actualSelectionIndex < 8 {
		v.selectionBox.DrawSegments(screen, 2, 1, 0)
	}
	for i := 0; i < 8; i++ {
		idx := i + (v.charScrollbar.GetCurrentOffset() * 2)
		if idx >= len(v.gameStates) {
			continue
		}
		v.characterNameLabel[i].Draw(screen)
		v.characterStatsLabel[i].Draw(screen)
		v.characterExpLabel[i].Draw(screen)
		v.characterImage[i].Render(screen, v.characterNameLabel[i].X-40, v.characterNameLabel[i].Y+50)
	}
	if v.showDeleteConfirmation {
		ebitenutil.DrawRect(screen, 0.0, 0.0, 800.0, 600.0, color.RGBA{0, 0, 0, 128})
		v.okCancelBox.DrawSegments(screen, 2, 1, 0)
		v.deleteCharConfirmLabel.Draw(screen)
	}
}

func (v *CharacterSelect) moveSelectionBox() {
	if v.selectedCharacter == -1 {
		v.d2HeroTitle.SetText("")
		return
	}
	bw := 272
	bh := 92
	selectedIndex := v.selectedCharacter - (v.charScrollbar.GetCurrentOffset() * 2)
	v.selectionBox.MoveTo(37+((selectedIndex&1)*int(bw)), 86+(int(bh)*(selectedIndex/2)))
	v.d2HeroTitle.SetText(v.gameStates[v.selectedCharacter].HeroName)
}

func (v *CharacterSelect) Update(tickTime float64) {
	if !v.showDeleteConfirmation {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if !v.mouseButtonPressed {
				v.mouseButtonPressed = true
				mx, my := ebiten.CursorPosition()
				bw := 272
				bh := 92
				localMouseX := mx - 37
				localMouseY := my - 86
				if localMouseX > 0 && localMouseX < int(bw*2) && localMouseY >= 0 && localMouseY < int(bh*4) {
					adjustY := localMouseY / int(bh)
					selectedIndex := adjustY * 2
					if localMouseX > int(bw) {
						selectedIndex += 1
					}
					if (v.charScrollbar.GetCurrentOffset()*2)+selectedIndex < len(v.gameStates) {
						v.selectedCharacter = (v.charScrollbar.GetCurrentOffset() * 2) + selectedIndex
						v.moveSelectionBox()
					}
				}
			}
		} else {
			v.mouseButtonPressed = false
		}
	}
}

func (v *CharacterSelect) onDeleteCharButtonClicked() {
	v.toggleDeleteCharacterDialog(true)
}

func (v *CharacterSelect) onDeleteCharacterConfirmClicked() {
	_ = os.Remove(v.gameStates[v.selectedCharacter].FilePath)
	v.charScrollbar.SetCurrentOffset(0)
	v.refreshGameStates()
	v.toggleDeleteCharacterDialog(false)
	v.deleteCharButton.SetEnabled(len(v.gameStates) > 0)
	v.okButton.SetEnabled(len(v.gameStates) > 0)
}

func (v *CharacterSelect) onDeleteCharacterCancelClicked() {
	v.toggleDeleteCharacterDialog(false)
}

func (v *CharacterSelect) toggleDeleteCharacterDialog(showDialog bool) {
	v.showDeleteConfirmation = showDialog
	v.okButton.SetEnabled(!showDialog)
	v.deleteCharButton.SetEnabled(!showDialog)
	v.exitButton.SetEnabled(!showDialog)
	v.newCharButton.SetEnabled(!showDialog)
	v.deleteCharOkButton.SetVisible(showDialog)
	v.deleteCharCancelButton.SetVisible(showDialog)
}

func (v *CharacterSelect) refreshGameStates() {
	v.gameStates = d2core.GetAllGameStates()
	v.updateCharacterBoxes()
	if len(v.gameStates) > 0 {
		v.selectedCharacter = 0
		v.d2HeroTitle.SetText(v.gameStates[0].HeroName)
		v.charScrollbar.SetMaxOffset(int(math.Ceil(float64(len(v.gameStates)-8) / float64(2))))
	} else {
		v.selectedCharacter = -1
		v.charScrollbar.SetMaxOffset(0)
	}
	v.moveSelectionBox()

}

func (v *CharacterSelect) onOkButtonClicked() {
	v.sceneProvider.SetNextScene(CreateGame(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager, v.gameStates[v.selectedCharacter]))
}
