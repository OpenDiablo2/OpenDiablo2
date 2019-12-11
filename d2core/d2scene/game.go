package d2scene

import (
	"image/color"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon/d2coreinterface"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	gameState     *d2core.GameState
	uiManager     *d2ui.Manager
	soundManager  *d2audio.Manager
	fileProvider  d2interface.FileProvider
	sceneProvider d2coreinterface.SceneProvider
	pentSpinLeft  d2render.Sprite
	pentSpinRight d2render.Sprite
	testLabel     d2ui.Label
	mapEngine     *d2mapengine.MapEngine
	hero          *d2core.Hero
}

func CreateGame(
	fileProvider d2interface.FileProvider,
	sceneProvider d2coreinterface.SceneProvider,
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
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.PentSpin), d2datadict.Palettes[d2enum.Sky])
			v.pentSpinLeft = d2render.CreateSpriteFromDC6(dc6)
			v.pentSpinLeft.Animate = true
			v.pentSpinLeft.AnimateBackwards = true
			v.pentSpinLeft.SpecialFrameTime = 475
			v.pentSpinLeft.MoveTo(100, 300)
		},
		func() {
			dc6, _ := d2dc6.LoadDC6(v.fileProvider.LoadFile(d2resource.PentSpin), d2datadict.Palettes[d2enum.Sky])
			v.pentSpinRight = d2render.CreateSpriteFromDC6(dc6)
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
		func() {
			v.mapEngine = d2mapengine.CreateMapEngine(v.gameState, v.soundManager, v.fileProvider)
			v.mapEngine.GenerateMap(d2enum.RegionAct1Town, 1, 0)

			startX, startY := v.mapEngine.GetStartTilePosition()
			v.hero = d2core.CreateHero(
				int32(startX*5)+3,
				int32(startY*5)+3,
				0,
				v.gameState.HeroType,
				v.gameState.Equipment,
				v.fileProvider,
			)
			v.mapEngine.AddEntity(v.hero)
		},
	}
}

func (v *Game) Unload() {
}

func (v Game) Render(screen *ebiten.Image) {
	screen.Fill(color.Black)
	v.mapEngine.Render(screen)
}

func (v *Game) Update(tickTime float64) {
	v.mapEngine.Advance(tickTime)

	rx, ry := v.mapEngine.IsoToWorld(v.hero.AnimatedEntity.LocationX/5, v.hero.AnimatedEntity.LocationY/5)
	v.mapEngine.MoveCameraTo(rx, ry)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		px, py := v.mapEngine.ScreenToIso(ebiten.CursorPosition())
		v.hero.AnimatedEntity.SetTarget(px*5, py*5, 1)
	}
}
