package d2scene

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	gameState     *d2core.GameState
	uiManager     *d2ui.Manager
	soundManager  *d2audio.Manager
	fileProvider  d2interface.FileProvider
	sceneProvider d2interface.SceneProvider
	pentSpinLeft  d2render.Sprite
	pentSpinRight d2render.Sprite
	testLabel     d2ui.Label
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
			v.pentSpinLeft = d2render.CreateSprite(v.fileProvider.LoadFile(d2resource.PentSpin), d2datadict.Palettes[d2enum.Sky])
			v.pentSpinLeft.Animate = true
			v.pentSpinLeft.AnimateBackwards = true
			v.pentSpinLeft.SpecialFrameTime = 475
			v.pentSpinLeft.MoveTo(100, 300)
		},
		func() {
			v.pentSpinRight = d2render.CreateSprite(v.fileProvider.LoadFile(d2resource.PentSpin), d2datadict.Palettes[d2enum.Sky])
			v.pentSpinRight.Animate = true
			v.pentSpinRight.SpecialFrameTime = 475
			v.pentSpinRight.MoveTo(650, 300)
		},
		func() {
			v.testLabel = d2ui.CreateLabel(v.fileProvider, d2resource.Font42, d2enum.Units)
			v.testLabel.Alignment = d2ui.LabelAlignCenter
			v.testLabel.SetText("Soon :tm:")
			v.testLabel.MoveTo(400, 250)
		},
	}
}

func (v *Game) Unload() {

}

func (v Game) Render(screen *ebiten.Image) {
	screen.Fill(color.Black)
	v.pentSpinLeft.Draw(screen)
	v.pentSpinRight.Draw(screen)
	v.testLabel.Draw(screen)
}

func (v *Game) Update(tickTime float64) {

}
