package d2gamescreen

import (
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

type Game struct {
	gameClient           *d2client.GameClient
	mapRenderer          *d2maprenderer.MapRenderer
	gameControls         *d2player.GameControls // TODO: Hack
	localPlayer          *d2mapentity.Player
	lastRegionType       d2enum.RegionIdType
	audioProvider        d2interface.AudioProvider
	terminal             d2interface.Terminal
	ticksSinceLevelCheck float64
	escapeMenu           *EscapeMenu
}

func CreateGame(audioProvider d2interface.AudioProvider, gameClient *d2client.GameClient, term d2interface.Terminal) *Game {
	result := &Game{
		gameClient:           gameClient,
		gameControls:         nil,
		localPlayer:          nil,
		lastRegionType:       d2enum.RegionNone,
		ticksSinceLevelCheck: 0,
		mapRenderer:          d2maprenderer.CreateMapRenderer(gameClient.MapEngine, term),
		escapeMenu:           NewEscapeMenu(audioProvider, term),
		audioProvider:        audioProvider,
		terminal:             term,
	}
	result.escapeMenu.onLoad()
	d2input.BindHandler(result.escapeMenu)

	return result
}

func (v *Game) OnLoad(loading d2screen.LoadingState) {
	v.audioProvider.PlayBGM("")
}

func (v *Game) OnUnload() error {
	d2input.UnbindHandler(v.gameControls) // TODO: hack
	d2input.UnbindHandler(v.escapeMenu)   // TODO: hack
	v.gameClient.Close()

	return nil
}

func (v *Game) Render(screen d2interface.Surface) error {
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

var hideZoneTextAfterSeconds = 2.0

func (v *Game) Advance(tickTime float64) error {
	if (v.escapeMenu != nil && !v.escapeMenu.isOpen) || len(v.gameClient.Players) != 1 {
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
				musicInfo := d2common.GetMusicDef(tile.RegionType)
				v.audioProvider.PlayBGM(musicInfo.MusicFile)

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
			v.gameControls = d2player.NewGameControls(player, v.gameClient.MapEngine, v.mapRenderer, v, v.terminal)
			v.gameControls.Load()
			d2input.BindHandler(v.gameControls)

			break
		}
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

func (v *Game) OnPlayerCast(missleID int, targetX, targetY float64) {
	v.gameClient.SendPacketToServer(d2netpacket.CreateCastPacket(v.gameClient.PlayerId, missleID, targetX, targetY))
}
