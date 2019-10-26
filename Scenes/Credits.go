package Scenes

import (
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

// Load is called to load the resources for the credits scene
func (v *Credits) Load() []func() {
	return []func(){
		func() {
			v.creditsBackground = v.fileProvider.LoadSprite(ResourcePaths.CreditsBackground, Palettes.Sky)
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
