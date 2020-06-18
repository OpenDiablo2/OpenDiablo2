package d2gamescreen

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Game struct {
	//pentSpinLeft  *d2ui.Sprite
	//pentSpinRight *d2ui.Sprite
	//testLabel     d2ui.Label
	gameClient   *d2client.GameClient
	mapRenderer  *d2map.MapRenderer
	gameControls *d2player.GameControls // TODO: Hack
	localPlayer  *d2map.Player
}

func CreateGame(gameClient *d2client.GameClient) *Game {
	return &Game{
		gameClient:   gameClient,
		gameControls: nil,
		localPlayer:  nil,
		mapRenderer:  d2map.CreateMapRenderer(gameClient.MapEngine),
	}
}

func (v *Game) OnLoad() error {
	return nil
}

func (v *Game) OnUnload() error {
	d2input.UnbindHandler(v.gameControls) // TODO: hack
	return nil
}

func (v *Game) Render(screen d2render.Surface) error {
	screen.Clear(color.Black)
	v.mapRenderer.Render(screen)
	if v.gameControls != nil {
		v.gameControls.Render(screen)
	}
	return nil
}

func (v *Game) Advance(tickTime float64) error {
	v.gameClient.MapEngine.Advance(tickTime) // TODO: Hack

	// Bind the game controls to the player once it exists
	if v.gameControls == nil {
		for _, player := range v.gameClient.Players {
			if player.Id != v.gameClient.PlayerId {
				continue
			}
			v.localPlayer = player
			v.gameControls = d2player.NewGameControls(player, v.gameClient.MapEngine, v.mapRenderer, v.gameClient)
			v.gameControls.Load()
			d2input.BindHandler(v.gameControls)

			break
		}
	}

	// Update the camera to focus on the player
	if v.localPlayer != nil {
		rx, ry := v.mapRenderer.WorldToOrtho(v.localPlayer.AnimatedComposite.LocationX/5, v.localPlayer.AnimatedComposite.LocationY/5)
		v.mapRenderer.MoveCameraTo(rx, ry)
	}
	return nil
}
