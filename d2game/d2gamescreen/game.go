package d2gamescreen

import (
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"
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
	lastRegionType       d2enum.RegionIdType
	ticksSinceLevelCheck float64
	escapeMenu           *EscapeMenu
}

func CreateGame(gameClient *d2client.GameClient) *Game {
	result := &Game{
		gameClient:           gameClient,
		gameControls:         nil,
		localPlayer:          nil,
		lastRegionType:       d2enum.RegionNone,
		ticksSinceLevelCheck: 0,
		mapRenderer:          d2maprenderer.CreateMapRenderer(gameClient.MapEngine),
		escapeMenu:           NewEscapeMenu(),
	}
	result.escapeMenu.OnLoad()
	d2input.BindHandler(result.escapeMenu)
	return result
}

func (v *Game) OnLoad(loading d2screen.LoadingState) {
	d2audio.PlayBGM("")
}

func (v *Game) OnUnload() error {
	d2input.UnbindHandler(v.gameControls) // TODO: hack
	v.gameClient.Close()
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

	if v.escapeMenu != nil {
		v.escapeMenu.Render(screen)
	}

	return nil
}

var hideZoneTextAfterSeconds = 2.0

func (v *Game) Advance(tickTime float64) error {
	if (v.escapeMenu != nil && !v.escapeMenu.IsOpen()) || len(v.gameClient.Players) != 1 {
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
				case d2enum.RegionAct1Town: // Rogue encampent
					v.localPlayer.SetIsInTown(true)
					d2audio.PlayBGM("/data/global/music/Act1/town1.wav")
				case d2enum.RegionAct1Wilderness: // Blood Moore
					v.localPlayer.SetIsInTown(false)
					d2audio.PlayBGM("/data/global/music/Act1/wild.wav")
				}

				// skip showing zone change text the first time we enter the world
				if v.lastRegionType != d2enum.RegionNone && v.lastRegionType != tile.RegionType {
					//TODO: Should not be using RegionType as an index - this will return incorrect LevelDetails record for most of the zones.
					v.gameControls.SetZoneChangeText(fmt.Sprintf("Entering The %s", d2datadict.LevelDetails[int(tile.RegionType)].LevelDisplayName))
					v.gameControls.ShowZoneChangeText()
					v.gameControls.HideZoneChangeTextAfter(hideZoneTextAfterSeconds)
				}
				v.lastRegionType = tile.RegionType
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

	if v.escapeMenu.IsOpen() {
		v.escapeMenu.Advance(tickTime)
	}

	// Update the camera to focus on the player
	if v.localPlayer != nil && !v.gameControls.FreeCam {
		rx, ry := v.mapRenderer.WorldToOrtho(v.localPlayer.LocationX/5, v.localPlayer.LocationY/5)
		v.mapRenderer.MoveCameraTo(rx, ry)
	}
	return nil
}

func (v *Game) OnPlayerMove(x, y float64) {
	heroPosX := v.localPlayer.LocationX / 5.0
	heroPosY := v.localPlayer.LocationY / 5.0
	v.gameClient.SendPacketToServer(d2netpacket.CreateMovePlayerPacket(v.gameClient.PlayerId, heroPosX, heroPosY, x, y))
}
