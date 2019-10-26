package Scenes

import (
	"encoding/binary"
	"strings"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/essial/OpenDiablo2/Sound"
	"github.com/essial/OpenDiablo2/UI"
	"github.com/hajimehoshi/ebiten"
)

// Credits represents the credits scene
type Credits struct {
	uiManager         *UI.Manager
	soundManager      *Sound.Manager
	fileProvider      Common.FileProvider
	sceneProvider     SceneProvider
	creditsBackground *Common.Sprite
	exitButton        *UI.Button
	creditsText       []string
}

// CreateCredits creates an instance of the credits scene
func CreateCredits(fileProvider Common.FileProvider, sceneProvider SceneProvider, uiManager *UI.Manager, soundManager *Sound.Manager) *Credits {
	result := &Credits{
		fileProvider:  fileProvider,
		uiManager:     uiManager,
		soundManager:  soundManager,
		sceneProvider: sceneProvider,
	}
	return result
}

func utf16BytesToString(b []byte, o binary.ByteOrder) string {
	utf := make([]uint16, (len(b)+(2-1))/2)
	for i := 0; i+(2-1) < len(b); i += 2 {
		utf[i/2] = o.Uint16(b[i:])
	}
	if len(b)/2 < len(utf) {
		utf[len(utf)-1] = utf8.RuneError
	}
	return string(utf16.Decode(utf))
}

// Load is called to load the resources for the credits scene
func (v *Credits) Load() []func() {
	return []func(){
		func() {
			v.creditsBackground = v.fileProvider.LoadSprite(ResourcePaths.CreditsBackground, Palettes.Sky)
			v.creditsBackground.MoveTo(0, 0)
		},
		func() {
			v.exitButton = UI.CreateButton(UI.ButtonTypeMedium, v.fileProvider, "EXIT")
			v.exitButton.MoveTo(30, 550)
			v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(v.exitButton)
		},
		func() {
			fileData := utf16BytesToString(v.fileProvider.LoadFile(ResourcePaths.CreditsText), binary.LittleEndian)
			v.creditsText = strings.Split(fileData, "\r\n")
		},
	}
}

// Unload unloads the data for the credits scene
func (v *Credits) Unload() {

}

// Render renders the credits scene
func (v *Credits) Render(screen *ebiten.Image) {
	v.creditsBackground.DrawSegments(screen, 4, 3, 0)
}

// Update runs the update logic on the credits scene
func (v *Credits) Update() {

}

func (v *Credits) onExitButtonClicked() {
	mainMenu := CreateMainMenu(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager)
	mainMenu.ShowTrademarkScreen = false
	v.sceneProvider.SetNextScene(mainMenu)
}
