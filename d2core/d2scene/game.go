package d2scene

import (
	"image/color"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon/d2coreinterface"
	"github.com/OpenDiablo2/OpenDiablo2/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
)

type Game struct {
	gameState     *d2core.GameState
	uiManager     *d2ui.Manager
	soundManager  *d2audio.Manager
	sceneProvider d2coreinterface.SceneProvider
	pentSpinLeft  *d2render.Sprite
	pentSpinRight *d2render.Sprite
	testLabel     d2ui.Label
	mapEngine     *d2mapengine.MapEngine
	hero          *d2core.Hero
	gameControls  *d2player.GameControls
}

func CreateGame(
	sceneProvider d2coreinterface.SceneProvider,
	uiManager *d2ui.Manager,
	soundManager *d2audio.Manager,
	gameState *d2core.GameState,
) *Game {
	result := &Game{
		gameState:     gameState,
		uiManager:     uiManager,
		soundManager:  soundManager,
		sceneProvider: sceneProvider,
	}
	return result
}

func (v *Game) Load() []func() {
	return []func(){
		func() {
			v.pentSpinLeft, _ = d2render.LoadSprite(d2resource.PentSpin, d2resource.PaletteSky)
			v.pentSpinLeft.PlayBackward()
			v.pentSpinLeft.SetPlayLengthMs(475)
			v.pentSpinLeft.SetPosition(100, 300)
		},
		func() {
			v.pentSpinRight, _ = d2render.LoadSprite(d2resource.PentSpin, d2resource.PaletteSky)
			v.pentSpinRight.PlayForward()
			v.pentSpinRight.SetPlayLengthMs(475)
			v.pentSpinRight.SetPosition(650, 300)
		},
		func() {
			v.testLabel = d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteUnits)
			v.testLabel.Alignment = d2ui.LabelAlignCenter
			v.testLabel.SetText("Soon :tm:")
			v.testLabel.SetPosition(400, 250)
		},
		func() {
			v.mapEngine = d2mapengine.CreateMapEngine(v.gameState, v.soundManager)
			v.mapEngine.GenerateMap(d2enum.RegionAct1Town, 1, 0)

			startX, startY := v.mapEngine.GetStartPosition()
			v.hero = d2core.CreateHero(
				int32(startX*5)+3,
				int32(startY*5)+3,
				0,
				v.gameState.HeroType,
				v.gameState.Equipment,
			)
			v.mapEngine.AddEntity(v.hero)
		},
		func() {
			v.gameControls = d2player.NewGameControls(v.hero, v.mapEngine)
			v.gameControls.Load()
		},
	}
}

func (v *Game) Unload() {
}

func (v Game) Render(screen *d2surface.Surface) {
	screen.Clear(color.Black)
	v.mapEngine.Render(screen)
	v.gameControls.Render(screen)
}

func (v *Game) Advance(tickTime float64) {
	v.mapEngine.Advance(tickTime)

	rx, ry := v.mapEngine.WorldToOrtho(v.hero.AnimatedEntity.LocationX/5, v.hero.AnimatedEntity.LocationY/5)
	v.mapEngine.MoveCameraTo(rx, ry)

	v.gameControls.Move(tickTime)
}
