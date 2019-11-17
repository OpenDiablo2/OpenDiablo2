package d2scene

import (
	"image/color"
	"log"

	"github.com/OpenDiablo2/D2Shared/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon/d2coreinterface"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
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
		func() {
			v.mapEngine = d2mapengine.CreateMapEngine(v.gameState, v.soundManager, v.fileProvider)
			// TODO: This needs to be different depending on the act of the player
			v.mapEngine.GenerateMap(d2enum.RegionAct1Town, 1, 0)
			region := v.mapEngine.GetRegion(0)
			rx, ry := d2helper.IsoToScreen(int(region.Region.StartX), int(region.Region.StartY), 0, 0)
			v.mapEngine.CenterCameraOn(float64(rx), float64(ry))
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
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		px, py := d2helper.ScreenToIso(float64(mx)-v.mapEngine.OffsetX, float64(my)-v.mapEngine.OffsetY)
		angle := 359 - d2helper.GetAngleBetween(
			v.mapEngine.Hero.AnimatedEntity.LocationX,
			v.mapEngine.Hero.AnimatedEntity.LocationY,
			px,
			py,
		)

		directionIndex := int((float64(angle) / 360.0) * 16.0)
		newDirection := d2render.DirectionLookup[directionIndex]
		if newDirection != v.mapEngine.Hero.AnimatedEntity.GetDirection() {
			v.mapEngine.Hero.AnimatedEntity.SetMode(d2enum.AnimationModePlayerTownNeutral.String(), v.mapEngine.Hero.Equipment.RightHand.GetWeaponClass(), newDirection)
			log.Printf("Angle: %d -> %d", directionIndex, newDirection)
		}
	}

	rx, ry := d2helper.IsoToScreen(int(v.mapEngine.Hero.AnimatedEntity.LocationX), int(v.mapEngine.Hero.AnimatedEntity.LocationY), 0, 0)
	v.mapEngine.CenterCameraOn(float64(rx), float64(ry))
}
