package d2gamescreen

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Game struct {
	//pentSpinLeft  *d2ui.Sprite
	//pentSpinRight *d2ui.Sprite
	//testLabel     d2ui.Label
	gameClient           *d2client.GameClient
	mapRenderer          *d2map.MapRenderer
	gameControls         *d2player.GameControls // TODO: Hack
	localPlayer          *d2map.Player
	lastLevelType        int
	ticksSinceLevelCheck float64
}

func CreateGame(gameClient *d2client.GameClient) *Game {
	return &Game{
		gameClient:           gameClient,
		gameControls:         nil,
		localPlayer:          nil,
		lastLevelType:        -1,
		ticksSinceLevelCheck: 0,
		mapRenderer:          d2map.CreateMapRenderer(gameClient.MapEngine),
	}
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
	screen.Clear(color.Black)
	v.mapRenderer.Render(screen)
	if v.gameControls != nil {
		v.gameControls.Render(screen)
	}
	return nil
}

func (v *Game) Advance(tickTime float64) error {
	v.gameClient.MapEngine.Advance(tickTime) // TODO: Hack

	v.ticksSinceLevelCheck += tickTime
	if v.ticksSinceLevelCheck > 1.0 {
		v.ticksSinceLevelCheck = 0
		if v.localPlayer != nil {
			region := v.gameClient.MapEngine.GetRegionAtTile(v.localPlayer.TileX, v.localPlayer.TileY)
			if region != nil {
				levelType := region.GetLevelType().Id
				if levelType != v.lastLevelType {
					v.lastLevelType = levelType
					switch levelType {
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
	if v.localPlayer != nil {
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
