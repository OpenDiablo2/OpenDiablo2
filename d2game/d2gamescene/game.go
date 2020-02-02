package d2gamescene

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
)

type Game struct {
	gameState     *d2gamestate.GameState
	pentSpinLeft  *d2ui.Sprite
	pentSpinRight *d2ui.Sprite
	testLabel     d2ui.Label
	mapEngine     *d2map.MapEngine
	hero          *d2map.Hero
	gameControls  *d2player.GameControls
}

func CreateGame(
	gameState *d2gamestate.GameState,
) *Game {
	result := &Game{
		gameState: gameState,
	}
	return result
}

func (v *Game) Load() []func() {
	return []func(){
		func() {
			animation, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteSky)
			v.pentSpinLeft, _ = d2ui.LoadSprite(animation)
			v.pentSpinLeft.PlayBackward()
			v.pentSpinLeft.SetPlayLengthMs(475)
			v.pentSpinLeft.SetPosition(100, 300)
		},
		func() {
			animation, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteSky)
			v.pentSpinRight, _ = d2ui.LoadSprite(animation)
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
			v.mapEngine = d2map.CreateMapEngine(v.gameState)
			v.mapEngine.GenerateMap(d2enum.RegionAct1Town, 1, 0)

			startX, startY := v.mapEngine.GetStartPosition()
			v.hero = d2map.CreateHero(
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
			d2input.BindHandler(v.gameControls)
		},
	}
}

func (v *Game) Unload() {
	d2input.UnbindHandler(v.gameControls)
}

func (v Game) Render(screen d2render.Surface) {
	screen.Clear(color.Black)
	v.mapEngine.Render(screen)
	v.gameControls.Render(screen)
}

func (v *Game) Advance(tickTime float64) {
	v.mapEngine.Advance(tickTime)

	rx, ry := v.mapEngine.WorldToOrtho(v.hero.AnimatedEntity.LocationX/5, v.hero.AnimatedEntity.LocationY/5)
	v.mapEngine.MoveCameraTo(rx, ry)
}
