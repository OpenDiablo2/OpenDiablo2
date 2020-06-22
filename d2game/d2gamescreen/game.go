package d2gamescreen

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

type Game struct {
	//pentSpinLeft  *d2ui.Sprite
	//pentSpinRight *d2ui.Sprite
	//testLabel     d2ui.Label
	gameClient           *d2client.GameClient
	mapRenderer          *d2maprenderer.MapRenderer
	gameControls         *d2player.GameControls // TODO: Hack
	localPlayer          *d2mapentity.Player
	lastLevelType        int
	ticksSinceLevelCheck float64
}

func CreateGame(gameClient *d2client.GameClient) *Game {
	result := &Game{
		gameClient:           gameClient,
		gameControls:         nil,
		localPlayer:          nil,
		lastLevelType:        -1,
		ticksSinceLevelCheck: 0,
		mapRenderer:          d2maprenderer.CreateMapRenderer(gameClient.MapEngine),
	}
	return result
}

func (v *Game) OnLoad() error {
	d2audio.PlayBGM("")
	return nil
}

func (v *Game) OnUnload() error {
	d2input.UnbindHandler(v.gameControls) // TODO: hack
	return nil
}

func (v *Game) Render(screen d2render.Surface) error {
	if v.gameClient.RegenMap {
		v.gameClient.RegenMap = false
		v.mapRenderer.RegenerateTileCache()
	}
	screen.Clear(color.Black)
	v.mapRenderer.Render(screen)
	if v.gameControls != nil {
		v.gameControls.Render(screen)
	}
	return nil
}

func (v *Game) Advance(tickTime float64) error {
	if !v.gameControls.InEscapeMenu() || len(v.gameClient.Players) != 1 {
		v.gameClient.MapEngine.Advance(tickTime) // TODO: Hack
	}

	if v.gameControls != nil {
		v.gameControls.Advance(tickTime)
	}

	v.ticksSinceLevelCheck += tickTime
	if v.ticksSinceLevelCheck > 1.0 {
		v.ticksSinceLevelCheck = 0
		if v.localPlayer != nil {
			tile := v.gameClient.MapEngine.TileAt(v.localPlayer.TileX, v.localPlayer.TileY)
			if tile != nil {
				switch tile.RegionType {
				case 1: // Rogue encampent
					v.localPlayer.SetIsInTown(true)
					d2audio.PlayBGM("/data/global/music/Act1/town1.wav")
				case 2: // Blood Moore
					v.localPlayer.SetIsInTown(false)
					d2audio.PlayBGM("/data/global/music/Act1/wild.wav")
				}
			}
		}
	}

	// Bind the game controls to the player once it exists
	if v.gameControls == nil {
		for _, player := range v.gameClient.Players {
			if player.Id != v.gameClient.PlayerId {
				continue
			}
			v.localPlayer = player
			v.gameControls = d2player.NewGameControls(player, v.gameClient.MapEngine, v.mapRenderer, v)
			v.gameControls.Load()
			d2input.BindHandler(v.gameControls)

			break
		}
	}

	// Update the camera to focus on the player
	if v.localPlayer != nil && !v.gameControls.FreeCam {
		rx, ry := v.mapRenderer.WorldToOrtho(v.localPlayer.AnimatedComposite.LocationX/5, v.localPlayer.AnimatedComposite.LocationY/5)
		v.mapRenderer.MoveCameraTo(rx, ry)
	}
	return nil
}

func (v *Game) OnPlayerMove(x, y float64) {
	heroPosX := v.localPlayer.AnimatedComposite.LocationX / 5.0
	heroPosY := v.localPlayer.AnimatedComposite.LocationY / 5.0
	v.gameClient.SendPacketToServer(d2netpacket.CreateMovePlayerPacket(v.gameClient.PlayerId, heroPosX, heroPosY, x, y))
}
