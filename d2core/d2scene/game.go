package d2scene

import (
	"image/color"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"
	"github.com/OpenDiablo2/D2Shared/d2helper"
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
	mapEngine     *d2mapengine.Engine
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
			// TODO: This needs to be different depending on the act of the player
			v.mapEngine.GenerateMap(d2enum.RegionAct1Town, 1, 0)
			v.mapEngine.SetRegion(0)
			region := v.mapEngine.Region()
			rx, ry := d2helper.IsoToScreen(region.Region.StartX, region.Region.StartY, 0, 0)
			v.mapEngine.CenterCameraOn(rx, ry)
			v.mapEngine.Hero = d2core.CreateHero(
				int32((region.Region.StartX*5)+3),
				int32((region.Region.StartY*5)+3),
				0,
				v.gameState.HeroType,
				v.gameState.Equipment,
				v.fileProvider)
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
	// TODO: Pathfinding

	if v.mapEngine.Hero.AnimatedEntity.LocationX != v.mapEngine.Hero.AnimatedEntity.TargetX ||
		v.mapEngine.Hero.AnimatedEntity.LocationY != v.mapEngine.Hero.AnimatedEntity.TargetY {
		v.mapEngine.Hero.AnimatedEntity.Step(tickTime)
	}

	for _, npc := range v.mapEngine.Region().Region.NPCs {

		if npc.HasPaths &&
			npc.AnimatedEntity.LocationX == npc.AnimatedEntity.TargetX &&
			npc.AnimatedEntity.LocationY == npc.AnimatedEntity.TargetY &&
			npc.AnimatedEntity.Wait() {
			// If at the target, set target to the next path.
			path := npc.NextPath()
			npc.AnimatedEntity.SetTarget(
				float64(path.X),
				float64(path.Y),
				path.Action,
			)
		}

		if npc.AnimatedEntity.LocationX != npc.AnimatedEntity.TargetX ||
			npc.AnimatedEntity.LocationY != npc.AnimatedEntity.TargetY {
			npc.AnimatedEntity.Step(tickTime)
		}

	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		px, py := v.mapEngine.ScreenToIso(ebiten.CursorPosition())
		v.mapEngine.Hero.AnimatedEntity.SetTarget(px*5, py*5, 1)
	}

	rx, ry := d2helper.IsoToScreen(v.mapEngine.Hero.AnimatedEntity.LocationX/5, v.mapEngine.Hero.AnimatedEntity.LocationY/5, 0, 0)
	v.mapEngine.CenterCameraOn(rx, ry)
}
