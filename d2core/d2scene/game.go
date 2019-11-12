package d2scene

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	gameState     *d2core.GameState
	uiManager     *d2ui.Manager
	soundManager  *d2audio.Manager
	fileProvider  d2interface.FileProvider
	sceneProvider d2interface.SceneProvider
}

func CreateGame(
	fileProvider d2interface.FileProvider,
	sceneProvider d2interface.SceneProvider,
	uiManager *d2ui.Manager,
	soundManager *d2audio.Manager,
	gameState *d2core.GameState,
) *Game {
	result := &Game{
		gameState:     gameState,
		fileProvider:  fileProvider,
		uiManager:     uiManager,
		soundManager:  soundManager,
		sceneProvider: sceneProvider,
	}
	return result
}

func (v *Game) Load() []func() {
	return []func(){
		func() {

		},
	}
}

func (v *Game) Unload() {

}

func (v Game) Render(screen *ebiten.Image) {

}

func (v *Game) Update(tickTime float64) {

}
